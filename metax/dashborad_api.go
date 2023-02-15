package metax

import (
	"bytes"
	"encoding/json"
	"io"
	"math/big"
	"net/http"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

type GeneralResponse struct {
	Code         int    `json:"code"`
	Message      string `json:"message"`
	ResponseCode int    `json:"responseCode"`
}

type CreateDappRequest struct {
	DappName             string `json:"dappName"`
	NetworkId            string `json:"networkId"`
	EnableBiconomyWallet bool   `json:"enableBiconomyWallet"`
}

type CreateDappResponse struct {
	ApiKey     string   `json:"apiKey"`
	fundingKey *big.Int `json:"fundingKey"`
}

type AddContractRequest struct {
	ContractName        string         `json:"contractName"`
	ContractAddress     common.Address `json:"contractAddress"`
	ContractType        string         `json:"contractType"`        // SCW for contract wallet or SC for contract
	WalletType          string         `json:"walletType"`          // SCW or GNOSIS or blank
	MetaTransactionType string         `json:"metaTransactionType"` // DEFAULT, TRUSTED_FORWARDER, ERC20_FORWARDER
	ABI                 abi.ABI        `json:"abi"`
}

type AddMethodRequest struct {
	ApiType         string         `json:"apiType"`
	MethodType      string         `json:"methodType"`
	Name            string         `json:"name"`
	ContractAddress common.Address `json:"contractAddress"`
	Method          string         `json:"method"`
}

type AddMethodResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	ApiIds  []struct {
		ApiId  string `json:"apiId"`
		Method string `json:"method"`
		Name   string `json:"name"`
	} `json:"apiIds"`
}

func (b *Bcnmy) CreateDapp(data *CreateDappRequest) (*CreateDappResponse, error) {
	errorCh := make(chan error)
	responseCh := make(chan *CreateDappResponse)
	body, err := json.Marshal(data)
	if err != nil {
		b.logger.WithError(err).Error("json marshal `CreateDappRequest` data failed")
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPost, CreateDappPublicURL, bytes.NewBuffer(body))
	if err != nil {
		b.logger.WithError(err).Error("CreateDapp NewRequest failed")
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("authToken", b.authToken)
	go func() {
		res, err := b.httpClient.Do(req)
		if err != nil {
			b.logger.WithError(err).Error("HttpClient request to CreateDapp failed")
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
		var ret *CreateDappResponse
		if err := json.Unmarshal(replyData, &ret); err != nil {
			b.logger.WithError(err).Error("json unmarshal body data failed")
			errorCh <- err
			return
		}
		responseCh <- &ret
	}()
	var resp *CreateDappResponse
	select {
	case resp = <-responseCh:
		return resp, nil
	case err := <-errorCh:
		b.logger.WithError(err).Error(err.Error())
		return nil, err
	}
}

func (b *Bcnmy) AddContract(data *AddContractRequest) (*GeneralResponse, error) {
	errorCh := make(chan error)
	responseCh := make(chan *GeneralResponse)
	body, err := json.Marshal(data)
	if err != nil {
		b.logger.WithError(err).Error("json marshal `AddContractRequest` data failed")
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPost, AddContractURL, bytes.NewBuffer(body))
	if err != nil {
		b.logger.WithError(err).Error("AddContract NewRequest failed")
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("authToken", b.authToken)
	req.Header.Set("apiKey", b.apiKey)
	go func() {
		res, err := b.httpClient.Do(req)
		if err != nil {
			b.logger.WithError(err).Error("HttpClient request to AddContract failed")
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
		var ret *GeneralResponse
		if err := json.Unmarshal(replyData, &ret); err != nil {
			b.logger.WithError(err).Error("json unmarshal body data failed")
			errorCh <- err
			return
		}
		responseCh <- &ret
	}()
	var resp *GeneralResponse
	select {
	case resp = <-responseCh:
		return resp, nil
	case err := <-errorCh:
		b.logger.WithError(err).Error(err.Error())
		return nil, err
	}
}

func (b *Bcnmy) AddMethod(data *AddMethodRequest) (*AddMethodResponse, error) {
	errorCh := make(chan error)
	responseCh := make(chan *AddMethodResponse)
	body, err := json.Marshal(data)
	if err != nil {
		b.logger.WithError(err).Error("json marshal `AddMethodRequest` data failed")
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPost, AddMethodURL, bytes.NewBuffer(body))
	if err != nil {
		b.logger.WithError(err).Error("AddMethod NewRequest failed")
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("authToken", b.authToken)
	req.Header.Set("apiKey", b.apiKey)
	go func() {
		res, err := b.httpClient.Do(req)
		if err != nil {
			b.logger.WithError(err).Error("HttpClient request to AddMethod failed")
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
		var ret *AddMethodResponse
		if err := json.Unmarshal(replyData, &ret); err != nil {
			b.logger.WithError(err).Error("json unmarshal body data failed")
			errorCh <- err
			return
		}
		responseCh <- &ret
	}()
	var resp *AddMethodResponse
	select {
	case resp = <-responseCh:
		return resp, nil
	case err := <-errorCh:
		b.logger.WithError(err).Error(err.Error())
		return nil, err
	}
}

func (b *Bcnmy) DeleteContract(data *DeleteContractRequest) (*GeneralResponse, error) {
	body, err := json.Marshal(data)
	if err != nil {
		b.logger.WithError(err).Error("json marshal `DeleteContract` data failed")
		return nil, err
	}
	req, err := http.NewRequest(http.MethodDelete, DeleteContractURL, bytes.NewBuffer(body))
	if err != nil {
		b.logger.WithError(err).Error("DeleteContract NewRequest failed")
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("authToken", b.authToken)
	req.Header.Set("apiKey", b.apiKey)

	res, err := b.httpClient.Do(req)
	if err != nil {
		b.logger.WithError(err).Error("HttpClient request to DeleteContract failed")
		return nil, err
	}
	defer res.Body.Close()
	replyData, err := io.ReadAll(res.Body)
	var ret *GeneralResponse
	if err := json.Unmarshal(replyData, &ret); err != nil {
		b.logger.WithError(err).Error("json unmarshal body data failed")
		return nil, err
	}
	return ret, nil
}

func (b *Bcnmy) DeleteMethod(data *DeleteMethodRequest) (*GeneralResponse, error) {
	body, err := json.Marshal(data)
	if err != nil {
		b.logger.WithError(err).Error("json marshal `DeleteMethod` data failed")
		return nil, err
	}
	req, err := http.NewRequest(http.MethodDelete, DeleteMethodURL, bytes.NewBuffer(body))
	if err != nil {
		b.logger.WithError(err).Error("DeleteMethod NewRequest failed")
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("authToken", b.authToken)
	req.Header.Set("apiKey", b.apiKey)

	res, err := b.httpClient.Do(req)
	if err != nil {
		b.logger.WithError(err).Error("HttpClient request to DeleteMethod failed")
		return nil, err
	}
	defer res.Body.Close()
	replyData, err := io.ReadAll(res.Body)
	var ret *GeneralResponse
	if err := json.Unmarshal(replyData, &ret); err != nil {
		b.logger.WithError(err).Error("json unmarshal body data failed")
		return nil, err
	}
	return ret, nil
}
