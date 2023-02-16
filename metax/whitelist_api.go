package metax

import (
	"bytes"
	"encoding/json"
	"fmt"
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
	responseCh := make(chan interface{}, 1)
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
	var resp AddDestinationResponse
	b.asyncHttpx(req, &resp, errorCh, responseCh)
	select {
	case ret := <-responseCh:
		resp, ok := ret.(*AddDestinationResponse)
		if !ok {
			return nil, fmt.Errorf("AddDestination failed")
		}
		return resp, nil
	case err := <-errorCh:
		b.logger.WithError(err).Error(err.Error())
		return nil, err
	}
}

func (b *Bcnmy) AddProxyContracts(data *AddProxyContractsRequest) (*ProxyContractsResponse, error) {
	responseCh := make(chan interface{}, 1)
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
	var resp ProxyContractsResponse
	b.asyncHttpx(req, &resp, errorCh, responseCh)
	select {
	case ret := <-responseCh:
		resp, ok := ret.(*ProxyContractsResponse)
		if !ok {
			return nil, fmt.Errorf("AddProxyContracts failed")
		}
		return resp, nil
	case err := <-errorCh:
		b.logger.WithError(err).Error(err.Error())
		return nil, err
	}
}

func (b *Bcnmy) PatchProxyContracts(data *PatchProxyContractsRequest) (*ProxyContractsResponse, error) {
	responseCh := make(chan interface{}, 1)
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
	var resp ProxyContractsResponse
	b.asyncHttpx(req, &resp, errorCh, responseCh)
	select {
	case ret := <-responseCh:
		resp, ok := ret.(*ProxyContractsResponse)
		if !ok {
			return nil, fmt.Errorf("PatchProxyContracts failed")
		}
		return resp, nil
	case err := <-errorCh:
		b.logger.WithError(err).Error(err.Error())
		return nil, err
	}
}

func (b *Bcnmy) GetProxyContracts() (*GetProxyContractsResponse, error) {
	responseCh := make(chan interface{}, 1)
	errorCh := make(chan error)
	req, err := http.NewRequest(http.MethodGet, ProxyContractsURL, nil)
	if err != nil {
		b.logger.WithError(err).Error("GetProxyContracts NewRequest failed")
		return nil, err
	}
	req.Header.Set("Authorization", b.GetAuthorization())
	var resp GetProxyContractsResponse
	b.asyncHttpx(req, &resp, errorCh, responseCh)
	select {
	case ret := <-responseCh:
		resp, ok := ret.(*GetProxyContractsResponse)
		if !ok {
			return nil, fmt.Errorf("GetProxyContracts failed")
		}
		return resp, nil
	case err := <-errorCh:
		b.logger.WithError(err).Error(err.Error())
		return nil, err
	}
}
