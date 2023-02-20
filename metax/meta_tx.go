package metax

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"time"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
)

type MetaTxMessage struct {
	From          common.Address `json:"from"`
	To            common.Address `json:"to"`
	Token         common.Address `json:"token"`
	TxGas         uint64         `json:"txGas"`
	TokenGasPrice string         `json:"tokenGasPrice"`
	BatchId       *big.Int       `json:"batchId"`
	BatchNonce    *big.Int       `json:"batchNonce"`
	Deadline      *big.Int       `json:"deadline"`
	Data          string         `json:"data"`
}

type MetaTxRequest struct {
	From          string        `json:"from"`
	To            string        `json:"to"`
	ApiID         string        `json:"apiId"`
	Params        []interface{} `json:"params"`
	SignatureType string        `json:"signatureType"`
}

type MetaTxResponse struct {
	TxHash common.Hash `json:"txHash"`
	Log    string      `json:"log"`
	Flag   int         `json:"flag"`
}

func (m *MetaTxMessage) TypedData() apitypes.TypedDataMessage {
	return apitypes.TypedDataMessage{
		"from":          m.From.Hex(),
		"to":            m.To.Hex(),
		"token":         m.Token.Hex(),
		"txGas":         hexutil.EncodeUint64(m.TxGas),
		"tokenGasPrice": m.TokenGasPrice,
		"batchId":       m.BatchId.String(),
		"batchNonce":    m.BatchNonce.String(),
		"deadline":      m.Deadline.String(),
		"data":          m.Data,
	}
}

func (b *Bcnmy) SendMetaNativeTx(data *MetaTxRequest) (*MetaTxResponse, error) {
	responseCh := make(chan interface{}, 1)
	errorCh := make(chan error)
	defer close(errorCh)
	defer close(responseCh)

	body, err := json.Marshal(data)
	if err != nil {
		b.logger.WithError(err).Error("json marshal `MetaTxRequest` data failed")
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPost, MetaTxNativeURL, bytes.NewBuffer(body))
	if err != nil {
		b.logger.WithError(err).Error("SendMetaNativeTx NewRequest failed")
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("x-api-key", b.apiKey)
	var resp MetaTxResponse
	b.asyncHttpx(req, &resp, errorCh, responseCh)
	select {
	case ret := <-responseCh:
		resp, ok := ret.(*MetaTxResponse)
		if !ok {
			return nil, fmt.Errorf("MetaAPI failed")
		}
		if resp.TxHash == common.HexToHash("0x0") {
			err := fmt.Errorf("%v", resp)
			b.logger.WithError(err).Error("TxHash is 0x0")
			return nil, err
		}
		return resp, nil
	case err := <-errorCh:
		b.logger.WithError(err).Error(err.Error())
		return nil, err
	}
}

func (b *Bcnmy) RawTransact(signer *Signer, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	apiId, ok := b.apiID[method]
	if !ok {
		err := fmt.Errorf("ApiId %s not found for %s", apiId.ID, method)
		b.logger.Error(err.Error())
		return nil, nil, err
	}
	funcSig, err := b.abi.Pack(method, params...)
	if err != nil {
		b.logger.WithError(err).Error("Abi Pack failed")
		return nil, nil, err
	}

	callMsg := ethereum.CallMsg{
		From: signer.Address,
		To:   &b.address,
		Data: funcSig,
	}
	callOpts := bind.CallOpts{
		Context: b.ctx,
		From:    signer.Address,
	}
	estimateGas, err := b.ethClient.EstimateGas(b.ctx, callMsg)
	if err != nil {
		b.logger.WithError(err).Error("EstimateGas failed")
		return nil, nil, err
	}
	batchNonce, err := b.trustedForwarder.Contract.GetNonce(&callOpts, signer.Address, b.batchId)
	if err != nil {
		b.logger.WithError(err).Errorf("GetNonce from %s failed", b.batchId)
		return nil, nil, err
	}

	metaTxMessage := &MetaTxMessage{
		From:          signer.Address,
		To:            b.address,
		Token:         common.HexToAddress("0x0"),
		TxGas:         estimateGas,
		TokenGasPrice: "0",
		BatchId:       b.batchId,
		BatchNonce:    batchNonce,
		Deadline:      big.NewInt(time.Now().Add(time.Hour).Unix()),
		Data:          hexutil.Encode(funcSig),
	}

	typedData := apitypes.TypedData{
		Types:       SignedTypes,
		PrimaryType: ForwardRequestType,
		Domain: apitypes.TypedDataDomain{
			Name:              ForwardRequestName,
			Version:           Version,
			VerifyingContract: b.trustedForwarder.Address.Hex(),
			Salt:              hexutil.Encode(common.LeftPadBytes(b.chainId.Bytes(), 32)),
		},
		Message: metaTxMessage.TypedData(),
	}
	signature, err := signer.SignTypedData(typedData)
	if err != nil {
		b.logger.WithError(err).Error("Signer signTypeData failed")
		return nil, nil, err
	}

	domainSeparator, err := typedData.HashStruct(EIP712DomainType, typedData.Domain.Map())
	if err != nil {
		b.logger.WithError(err).Error("EIP712Domain Separator hash failed")
		return nil, nil, err
	}

	req := &MetaTxRequest{
		From:  signer.Address.Hex(),
		To:    b.address.Hex(),
		ApiID: apiId.ID,
		Params: []interface{}{
			metaTxMessage,
			hexutil.Encode(domainSeparator),
			hexutil.Encode(signature),
		},
		SignatureType: SignatureEIP712Type,
	}
	resp, err := b.SendMetaNativeTx(req)
	if err != nil {
		b.logger.WithError(err).Errorf("Transaction failed: %v", err)
		return nil, nil, err
	}

	tx, _, err := b.ethClient.TransactionByHash(b.ctx, resp.TxHash)
	if err != nil {
		b.logger.WithError(err).Errorf("Checking TransactionByHash failed: %v", err)
		return nil, nil, err
	}
	receipt, err := bind.WaitMined(context.Background(), b.ethClient, tx)
	return tx, receipt, err
}

// / Backend using this method, handle frontend passing signature, MetaTxMessage and
// / ForwardRequestType data Hash value
func (b *Bcnmy) EnhanceTransact(from string, method string, signature []byte, metaTxMessage *MetaTxMessage, typedDataHash string) (*types.Transaction, *types.Receipt, error) {
	apiId, ok := b.apiID[method]
	if !ok {
		err := fmt.Errorf("ApiId %s not found for %s", apiId.ID, method)
		b.logger.Error(err.Error())
		return nil, nil, err
	}
	typedData := apitypes.TypedData{
		Types:       SignedTypes,
		PrimaryType: ForwardRequestType,
		Domain: apitypes.TypedDataDomain{
			Name:              ForwardRequestName,
			Version:           Version,
			VerifyingContract: b.trustedForwarder.Address.Hex(),
			Salt:              hexutil.Encode(common.LeftPadBytes(b.chainId.Bytes(), 32)),
		},
		Message: metaTxMessage.TypedData(),
	}
	hash, err := typedData.HashStruct(typedData.PrimaryType, typedData.Message)
	if err != nil {
		b.logger.WithError(err).Errorf("HashStruct failed to hash typedData, %v", err)
		return nil, nil, err
	}
	if hash.String() != typedDataHash {
		err := fmt.Errorf("Hash string not match parameter typedDataHash %s", typedDataHash)
		b.logger.WithError(err).Errorf("%v", err)
		return nil, nil, err
	}

	domainSeparator, err := typedData.HashStruct(EIP712DomainType, typedData.Domain.Map())
	if err != nil {
		b.logger.WithError(err).Error("EIP712Domain Separator hash failed")
		return nil, nil, err
	}

	req := &MetaTxRequest{
		From:  from,
		To:    b.address.Hex(),
		ApiID: apiId.ID,
		Params: []interface{}{
			metaTxMessage,
			hexutil.Encode(domainSeparator),
			hexutil.Encode(signature),
		},
		SignatureType: SignatureEIP712Type,
	}
	resp, err := b.SendMetaNativeTx(req)
	if err != nil {
		b.logger.WithError(err).Errorf("Transaction failed: %v", err)
		return nil, nil, err
	}

	tx, _, err := b.ethClient.TransactionByHash(b.ctx, resp.TxHash)
	if err != nil {
		b.logger.WithError(err).Errorf("Checking TransactionByHash failed: %v", err)
		return nil, nil, err
	}
	receipt, err := bind.WaitMined(context.Background(), b.ethClient, tx)
	return tx, receipt, err
}

func (b *Bcnmy) Pack(method string, params ...interface{}) ([]byte, error) {
	data, err := b.abi.Pack(method, params...)
	if err != nil {
		b.logger.WithError(err).Error("Abi Pack failed")
		return nil, err
	}
	return data, err
}
