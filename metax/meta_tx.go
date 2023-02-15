package metax

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"math/big"
	"net/http"
	"time"

	ethereum "github.com/ethereum/go-ethereum"
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
	From          string                 `json:"from"`
	To            string                 `json:"to"`
	ApiID         string                 `json:"apiId"`
	Params        map[string]interface{} `json:"params"`
	SignatureType string                 `json:"signatureType"`
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
	metaTxCh := make(chan *MetaTxResponse, 1)
	errorCh := make(chan error)
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
	go func() {
		res, err := b.httpClient.Do(req)
		if err != nil {
			b.logger.WithError(err).Error("HttpClient request to SendMetaTxNative failed")
			errorCh <- err
			return
		}
		defer res.Body.Close()
		replyData, err := io.ReadAll(res.Body)
		if err != nil {
			b.logger.WithError(err).Error("io read request body failed")
			errorCh <- err
			return
		}
		var ret *MetaTxResponse
		if err := json.Unmarshal(replyData, &ret); err != nil {
			b.logger.WithError(err).Error("json unmarshal body data failed")
			errorCh <- err
			return
		}
		metaTxCh <- &ret
	}()
	var resp *MetaTxResponse
	select {
	case resp = <-metaTxCh:
	case err := <-errorCh:
		b.logger.WithError(err).Error(err.Error())
		return nil, err
	}
	if resp.TxHash == common.HexToHash("0x0") {
		err := fmt.Errorf("%s", resp)
		b.logger.WithError(err).Error("TxHash is 0x0")
		return nil, err
	}
	return resp, nil
}

func (b *Bcnmy) RawTransact(signer *Signer, method string, params ...interface{}) (*types.Transaction, error) {
	apiId, ok := b.apiID[method]
	if !ok {
		err := fmt.Errorf("ApiId not found for %s", method)
		b.logger.Error(err.Error())
		return nil, err
	}
	data, err := b.abi.Pack(method, params...)
	if err != nil {
		b.logger.WithError(err).Error("Abi Pack failed")
		return nil, err
	}

	callMsg, err := ethereum.CallMsg{
		From: signer.Address,
		To:   &b.address,
		Data: data,
	}
	callOpts = bind.CallOpts{
		Context: b.ctx,
		From:    signer.Address,
	}
	estimateGas, err := b.ethClient.EstimateGas(b.ctx, callMsg)
	if err != nil {
		b.logger.WithError(err).Error("EstimateGas failed")
		return nil, err
	}
	batchNonce, err := b.trustedForwarder.Contract.GetNonce(callOpts, account.Address, b.batchId)
	if err != nil {
		b.logger.WithError(err).Errorf("GetNonce from %s failed", b.batchId)
		return nil, err
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
		Data:          hexutil.Encode(data),
	}

	typedData := apitypes.TypedData{
		Type:        SignedTypes,
		PrimaryType: ForwardRequestType,
		Domain: apitypes.TypedDataDomain{
			Name:              ForwardRequestName,
			Version:           Version,
			VerifyingContract: b.address.Hex(),
			Salt:              hexutil.Encode(common.LeftPadBytes(m.chainID.Bytes(), 32)),
		},
		Message: metaTxMessage.TypedData(),
	}
	signature, err := signer.SignTypedData(typedData)
	if err != nil {
		b.logger.WithError(err).Error("Signer signTypeData failed")
		return nil, err
	}

	domainSeparator, err := typedData.HashStruct(EIP712DoaminType, typedData.Domain.Map())
	if err != nil {
		b.logger.WithError(err).Error("EIP712Domain Separator hash failed")
		return nil, err
	}

	req := &MetaTxRequest{
		From:  signer.Address.Hex(),
		To:    b.address.Hex(),
		ApiID: b.apiId,
		Params: []interface{}{
			metaTxMessage,
			hexutil.Encode(domainSeparator),
			hexutil.Encode(signature),
		},
		SignatureType: SignatureTypeEIP712,
	}
	resp, err := b.SendMetaNativeTx(req)
	if err != nil {
		b.logger.WithError(err).Errorf("Transaction failed: %v", err)
		return nil, err
	}
	tx, _, err := b.ethClient.TransactionByHash(b.ctx, resp.TxHash)
	return tx, err
}

// / Backend using this method, handle frontend passing signature, MetaTxMessage and
// / ForwardRequestType data Hash value
func (b *Bcnmy) EnhanceTransact(signature []byte, metaTxMessage *MetaTxMessage, typedDataHash string) (*types.Transaction, error) {
	typedData := apitypes.TypedData{
		Type:        SignedTypes,
		PrimaryType: ForwardRequestType,
		Domain: apitypes.TypedDataDomain{
			Name:              ForwardRequestName,
			Version:           Version,
			VerifyingContract: b.address.Hex(),
			Salt:              hexutil.Encode(common.LeftPadBytes(m.chainID.Bytes(), 32)),
		},
		Message: metaTxMessage.TypedData(),
	}
	hash, err := typedData.HashStruct(typedData.PrimaryType, typedData.Message)
	if err != nil {
		b.logger.WithError(err).Errorf("HashStruct failed to hash typedData, %v", err)
		return nil, err
	}
	if hash.String() != typedDataHash {
		err := fmt.Errorf("Hash string not match parameter typedDataHash %s", typedDataHash)
		b.logger.WithError(err).Errorf("%v", err)
		return nil, err
	}

	domainSeparator, err := typedData.HashStruct(EIP712DoaminType, typedData.Domain.Map())
	if err != nil {
		b.logger.WithError(err).Error("EIP712Domain Separator hash failed")
		return nil, err
	}

	req := &MetaTxRequest{
		From:  signer.Address.Hex(),
		To:    b.address.Hex(),
		ApiID: b.apiId,
		Params: []interface{}{
			metaTxMessage,
			hexutil.Encode(domainSeparator),
			hexutil.Encode(signature),
		},
		SignatureType: SignatureTypeEIP712,
	}
	resp, err := b.SendMetaNativeTx(req)
	if err != nil {
		b.logger.WithError(err).Errorf("Transaction failed: %v", err)
		return nil, err
	}
	tx, _, err := b.ethClient.TransactionByHash(b.ctx, resp.TxHash)
	return tx, err
}
