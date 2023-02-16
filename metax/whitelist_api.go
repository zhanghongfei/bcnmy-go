package metax

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type AddDestinationRequest struct {
	DestinationAddresses []string `json:"destinationAddresses"`
}

type AddDestinationResponse struct {
	Code               int      `json:"code"`
	Message            string   `json:"message"`
	RegisteredCount    int      `json:"registeredCount"`
	DuplicateContracts []string `json:"duplicateContracts"`
	InvalidContracts   []string `json:"invalidContracts"`
}

type AddProxyContractsRequest struct {
	Addresses []string `json:"addresses"`
}

type ProxyContractsResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type GetProxyContractsResponse struct {
	ProxyContractsResponse
	Total     int `json:"total"`
	Addresses []struct {
		Address string `json:"address"`
		Status  bool   `json:"status"`
	} `json:"addresses"`
}

type PatchProxyContractsRequest struct {
	Status  int    `json:"status"` // 0 => inactive, 1 => active
	Address string `json:"address"`
}

func (b *Bcnmy) AddDestinationAddresses(data *AddDestinationRequest) (*AddDestinationResponse, error) {
	responseCh := make(chan *AddDestinationResponse, 1)
	errorCh := make(chan error)
	body, err := json.Marshal(data)
	if err != nil {
		b.logger.WithError(err).Error("json marshal `AddDestinationRequest` data failed")
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPost, AddDestinationAddressesURL, bytes.NewBuffer(body))
	if err != nil {
		b.logger.WithError(err).Error("AddDestinationAddresses NewRequest failed")
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Authorization", b.GetAuthorization())
	go func() {
		res, err := b.httpClient.Do(req)
		if err != nil {
			b.logger.WithError(err).Error("HttpClient request to AddDestinationAddresses failed")
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
		var ret *AddDestinationResponse
		if err := json.Unmarshal(replyData, &ret); err != nil {
			b.logger.WithError(err).Error("json unmarshal body data failed")
			errorCh <- err
			return
		}
		responseCh <- ret
	}()
	var resp *AddDestinationResponse
	select {
	case resp = <-responseCh:
		return resp, nil
	case err := <-errorCh:
		b.logger.WithError(err).Error(err.Error())
		return nil, err
	}
}

func (b *Bcnmy) AddProxyContracts(data *AddProxyContractsRequest) (*ProxyContractsResponse, error) {
	responseCh := make(chan *ProxyContractsResponse, 1)
	errorCh := make(chan error)
	body, err := json.Marshal(data)
	if err != nil {
		b.logger.WithError(err).Error("json marshal `AddProxyContractsRequest` data failed")
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPost, ProxyContractsURL, bytes.NewBuffer(body))
	if err != nil {
		b.logger.WithError(err).Error("AddProxyContracts NewRequest failed")
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Authorization", b.GetAuthorization())
	go func() {
		res, err := b.httpClient.Do(req)
		if err != nil {
			b.logger.WithError(err).Error("HttpClient request to AddProxyContracts failed")
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
		var ret *ProxyContractsResponse
		if err := json.Unmarshal(replyData, &ret); err != nil {
			b.logger.WithError(err).Error("json unmarshal body data failed")
			errorCh <- err
			return
		}
		responseCh <- ret
	}()
	var resp *ProxyContractsResponse
	select {
	case resp = <-responseCh:
		return resp, nil
	case err := <-errorCh:
		b.logger.WithError(err).Error(err.Error())
		return nil, err
	}
}

func (b *Bcnmy) PatchProxyContracts(data *PatchProxyContractsRequest) (*ProxyContractsResponse, error) {
	responseCh := make(chan *ProxyContractsResponse, 1)
	errorCh := make(chan error)
	body, err := json.Marshal(data)
	if err != nil {
		b.logger.WithError(err).Error("json marshal `PatchProxyContractsRequest` data failed")
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPatch, ProxyContractsURL, bytes.NewBuffer(body))
	if err != nil {
		b.logger.WithError(err).Error("PatchProxyContracts NewRequest failed")
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Authorization", b.GetAuthorization())
	go func() {
		res, err := b.httpClient.Do(req)
		if err != nil {
			b.logger.WithError(err).Error("HttpClient request to PatchProxyContracts failed")
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
		var ret *ProxyContractsResponse
		if err := json.Unmarshal(replyData, &ret); err != nil {
			b.logger.WithError(err).Error("json unmarshal body data failed")
			errorCh <- err
			return
		}
		responseCh <- ret
	}()
	var resp *ProxyContractsResponse
	select {
	case resp = <-responseCh:
		return resp, nil
	case err := <-errorCh:
		b.logger.WithError(err).Error(err.Error())
		return nil, err
	}
}

func (b *Bcnmy) GetProxyContracts() (*GetProxyContractsResponse, error) {
	responseCh := make(chan *GetProxyContractsResponse, 1)
	errorCh := make(chan error)
	req, err := http.NewRequest(http.MethodGet, ProxyContractsURL, nil)
	if err != nil {
		b.logger.WithError(err).Error("GetProxyContracts NewRequest failed")
		return nil, err
	}
	req.Header.Set("Authorization", b.GetAuthorization())
	go func() {
		res, err := b.httpClient.Do(req)
		if err != nil {
			b.logger.WithError(err).Error("HttpClient request to GetProxyContracts failed")
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
		var ret *GetProxyContractsResponse
		if err := json.Unmarshal(replyData, &ret); err != nil {
			b.logger.WithError(err).Error("json unmarshal body data failed")
			errorCh <- err
			return
		}
		responseCh <- ret
	}()
	var resp *GetProxyContractsResponse
	select {
	case resp = <-responseCh:
		return resp, nil
	case err := <-errorCh:
		b.logger.WithError(err).Error(err.Error())
		return nil, err
	}
}
