package metax

import (
	"fmt"
	"math/big"
	"net/http"
	"net/url"
	"strings"
)

type UniqueUserDataRequest struct {
	StartDate string `json:"startDate"` /// Format (“MM-DD-YYYY”) example: 21st Jan 2022 would be 01-21-2022
	EndDate   string `json:"endDate"`   /// Format (“MM-DD-YYYY”)
}

type UniqueUserDataResponse struct {
	GeneralResponse
	UniqueUserData []struct {
		Date      string   `json:"date"`
		Count     int      `json:"count"`
		Addresses []string `json:"addresses"`
	}
}

type UserLimitRequest struct {
	SignerAddress string `json:"signerAddress"`
	ApiId         string `json:"apiId"`
}

type UserLimitResponse struct {
	GeneralResponse
	UserLimitData struct {
		LimitLeft struct {
			SignerAddress        string `json:"signerAddress"`
			TransactionLimitLeft int    `json:"transactionLimitLeft"`
			TransactionCount     int    `json:"transactionCount"`
			AreLimitsConsumed    bool   `json:"areLimitsConsumed"`
			UserTransactionLimit int    `json:"userTransactionLimit"`
		} `json:"limitLeft"`
		LimitType        string   `json:"limitType"`
		LimitStartTime   *big.Int `json:"limitStartTime"`
		LimitEndTime     *big.Int `json:"limitEndTime"`
		TimePeriodInDays int      `json:"timePeriodInDays"`
	} `json:"userLimitData"`
}

//type GasTankBalanceResponse struct {
//GeneralResponse
//DappGasTankData struct {
//EffectiveBalanceInWei          *big.Int `json:"effectiveBalanceInWei"`
//EffectiveBalanceInStandardForm string   `json:"effectiveBalanceInStandardForm"`
//IsBelowThreshold               bool     `json:"isBelowThreshold"`
//IsInGracePeriod                bool     `json:"isInGracePeriod"`
//} `json:"dappGasTankData"`
//}

func (b *Bcnmy) GetUniqueUserData(data *UniqueUserDataRequest) (*UniqueUserDataResponse, error) {
	responseCh := make(chan interface{}, 1)
	errorCh := make(chan error)
	defer close(errorCh)
	defer close(responseCh)

	body := url.Values{
		"startDate": {data.StartDate},
		"endDate":   {data.EndDate},
	}
	req, err := http.NewRequest(http.MethodGet, UniqueUserDataURL, strings.NewReader(body.Encode()))
	if err != nil {
		b.logger.WithError(err).Error("GetUniqueUserData NewRequest failed")
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("authToken", b.authToken)
	req.Header.Set("apiKey", b.apiKey)
	var resp UniqueUserDataResponse
	b.asyncHttpx(req, &resp, errorCh, responseCh)
	select {
	case ret := <-responseCh:
		resp, ok := ret.(*UniqueUserDataResponse)
		if !ok {
			return nil, fmt.Errorf("UniqueUserData failed")
		}
		return resp, nil
	case err := <-errorCh:
		b.logger.Error(err.Error())
		return nil, err
	}
}

func (b *Bcnmy) GetUserLimit(data *UserLimitRequest) (*UserLimitResponse, error) {
	responseCh := make(chan interface{}, 1)
	errorCh := make(chan error)
	defer close(errorCh)
	defer close(responseCh)

	body := url.Values{
		"signerAddress": {data.SignerAddress},
		"apiId":         {data.ApiId},
	}
	req, err := http.NewRequest(http.MethodGet, UserLimitURL, strings.NewReader(body.Encode()))
	if err != nil {
		b.logger.WithError(err).Error("GetUserLimit NewRequest failed")
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("authToken", b.authToken)
	req.Header.Set("apiKey", b.apiKey)
	var resp UserLimitResponse
	b.asyncHttpx(req, &resp, errorCh, responseCh)
	select {
	case ret := <-responseCh:
		resp, ok := ret.(*UserLimitResponse)
		if !ok {
			return nil, fmt.Errorf("UserLimit failed")
		}
		return resp, nil
	case err := <-errorCh:
		b.logger.Error(err.Error())
		return nil, err
	}
}

/// always returns 401 UnAuthorized bad code, remove but mark it.
//func (b *Bcnmy) GetGasTankBalance() (*GasTankBalanceResponse, error) {
//responseCh := make(chan interface{}, 1)
//errorCh := make(chan error)
//defer close(errorCh)
//defer close(responseCh)

//req, err := http.NewRequest(http.MethodGet, GasTankBalanceURL, nil)
//if err != nil {
//b.logger.WithError(err).Error("GetGasTankBalance NewRequest failed")
//return nil, err
//}
//req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
//fmt.Println(b.authToken)
//fmt.Println(b.apiKey)
//req.Header.Set("authToken", b.authToken)
//req.Header.Set("apiKey", b.apiKey)
//var resp GasTankBalanceResponse
//b.asyncHttpx(req, &resp, errorCh, responseCh)
//select {
//case ret := <-responseCh:
//resp, ok := ret.(*GasTankBalanceResponse)
//if !ok {
//return nil, fmt.Errorf("GasTankBalance failed")
//}
//return resp, nil
//case err := <-errorCh:
//b.logger.Error(err.Error())
//return nil, err
//}
//}
