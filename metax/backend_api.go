package metax

import (
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"net/url"
)

/*
{
    "log": "Dapps fetched for user with userId: f93482",
    "code": 200,
    "data": {
        "dapps": [
            {
                "_id": "64339de",
                "networkId": "137",
                "dappName": "nt",
                "active": true,
                "apiKeyPrefix": "",
                "apiKey": "",
                "dappLimit": {
                    "durationValue": 1,
                    "durationUnit": "day",
                    "value": 1000,
                    "limitDurationInMs": 86400000,
                    "limitStartTime": 1681430400000,
                    "type": 1
                },
                "dappLimitStatus": 1,
                "userLimit": {
                    "durationValue": 1,
                    "durationUnit": "day",
                    "value": 10,
                    "limitDurationInMs": 86400000,
                    "limitStartTime": 1681084800000,
                    "type": 1
                },
                "userLimitStatus": 1,
                "apiLimit": {
                    "durationValue": 1,
                    "durationUnit": "day",
                    "value": 100,
                    "limitDurationInMs": 86400000,
                    "limitStartTime": 1681084800000,
                    "type": 1
                },
                "apiLimitStatus": 1,
                "createdOn": 1681104143290,
                "createdBy": "61ff93482",
                "allowedDomains": [],
                "updatedOn": 1681104143290,
                "updatedBy": "6bff93482",
                "fundingKey": 10333,
                "gasThreshold": 28206553200000000,
                "effectiveBalance": 19198843608675230000,
                "lastOutstandingBalance": 0,
                "gasTankBalance": 19.19884360867523,
                "gasTankDepletionRate": 0,
                "oustandingBalance": 0,
                "transactionCountThisMonth": 0,
                "transactionCountInTotal": 150
            }
        ]
    }
}
*/

type DappResponse struct {
	Log  string `json:"log"`
	Code int    `json:"code"`
	Data Data   `json:"data"`
}

type Data struct {
	Dapps []Dapp `json:"dapps"`
}

type Dapp struct {
	NetworkId                 string    `json:"networkId"`
	DappName                  string    `json:"dappName"`
	DappLimit                 DappLimit `json:"dappLimit"`
	DappLimitStatus           int       `json:"dappLimitStatus"`
	UserLimit                 DappLimit `json:"userLimit"`
	UserLimitStatus           int       `json:"userLimitStatus"`
	ApiLimit                  DappLimit `json:"apiLimit"`
	ApiLimitStatus            int       `json:"apiLimitStatus"`
	GasTankBalance            float64   `json:"gasTankBalance"`
	EffectiveBalance          *big.Int  `json:"effectiveBalance"`
	TransactionCountThisMonth int       `json:"transactionCountThisMonth"`
	TransactionCountInTotal   int       `json:"transactionCountInTotal"`
}

type DappLimit struct {
	Type              int     `json:"type"`
	Value             float32 `json:"value"`
	DurationValue     float32 `json:"durationValue"`
	DurationUnit      string  `json:"durationUnit"`
	LimitStartTime    int64   `json:"limitStartTime"`
	LimitDurationInMs int64   `json:"limitDurationInMs"`
}

type LoginResponse struct {
	Message string `json:"message"`
}

func (b *Bcnmy) BackendLogin() (*LoginResponse, error) {
	body := url.Values{
		"email":    {b.email},
		"password": {b.password},
	}
	loginResp, err := b.backendHttpClient.PostForm(BackendLoginURL, body)
	if err != nil {
		b.logger.WithError(err).Error("BackendLogin error")
		return nil, err
	}
	defer loginResp.Body.Close()
	replyData, err := io.ReadAll(loginResp.Body)
	if err != nil {
		b.logger.WithError(err).Error("BackendLogin Read body error")
		return nil, err
	}
	if loginResp.StatusCode != 200 {
		err = fmt.Errorf("BackendLogin got %v", loginResp.Status)
		b.logger.WithError(err).Error("BackendLogin status error")
		return nil, err
	}
	var resp LoginResponse
	if err = json.Unmarshal(replyData, &resp); err != nil {
		b.logger.WithError(err).Error("BackendLogin json unmarshal body data failed")
		return nil, err
	}
	return &resp, nil
}

func (b *Bcnmy) BackendDappList() (*DappResponse, error) {
	responseCh := make(chan interface{}, 1)
	errorCh := make(chan error)
	defer close(errorCh)
	defer close(responseCh)

	req, err := http.NewRequest(http.MethodGet, BackendDappURL, nil)
	if err != nil {
		b.logger.WithError(err).Error("BackendDappList NewRequest failed")
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	var resp DappResponse
	b.backendAsyncHttpx(req, &resp, errorCh, responseCh)
	select {
	case ret := <-responseCh:
		resp, ok := ret.(*DappResponse)
		if !ok {
			return nil, fmt.Errorf("BackendDappList failed")
		}
		return resp, nil
	case err := <-errorCh:
		b.logger.Error(err.Error())
		return nil, err
	}
}

func (b *Bcnmy) GetBackendDapps() (*DappResponse, error) {
	resp, err := b.BackendDappList()
	if err != nil && err.Error() == "401" {
		loginResp, loginErr := b.BackendLogin()
		if loginErr != nil {
			b.logger.WithError(loginErr).Errorf("GetBackendDapps login error %+v", loginResp)
			return nil, loginErr
		}
		resp, err = b.BackendDappList()
	}
	return resp, err
}
