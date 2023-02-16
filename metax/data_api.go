package metax

import (
	"encoding/json"
	"io"
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

type GasTankBalanceResponse struct {
	GeneralResponse
	DappGasTankData struct {
		EffectiveBalanceInWei          *big.Int `json:"effectiveBalanceInWei"`
		EffectiveBalanceInStandardForm string   `json:"effectiveBalanceInStandardForm"`
		IsBelowThreshold               bool     `json:"isBelowThreshold"`
		IsInGracePeriod                bool     `json:"isInGracePeriod"`
	} `json:"dappGasTankData"`
}

func (b *Bcnmy) GetUniqueUserData(data *UniqueUserDataRequest) (*UniqueUserDataResponse, error) {
	responseCh := make(chan *UniqueUserDataResponse, 1)
	errorCh := make(chan error)
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
	go func() {
		res, err := b.httpClient.Do(req)
		if err != nil {
			b.logger.WithError(err).Error("HttpClient request to GetUniqueUserData failed")
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
		var ret *UniqueUserDataResponse
		if err := json.Unmarshal(replyData, &ret); err != nil {
			b.logger.WithError(err).Error("json unmarshal body data failed")
			errorCh <- err
			return
		}
		responseCh <- ret
	}()
	var resp *UniqueUserDataResponse
	select {
	case resp = <-responseCh:
		return resp, nil
	case err := <-errorCh:
		b.logger.WithError(err).Error(err.Error())
		return nil, err
	}
}

func (b *Bcnmy) GetUserLimit(data *UserLimitRequest) (*UserLimitResponse, error) {
	responseCh := make(chan *UserLimitResponse, 1)
	errorCh := make(chan error)
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
	go func() {
		res, err := b.httpClient.Do(req)
		if err != nil {
			b.logger.WithError(err).Error("HttpClient request to GetUserLimit failed")
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
		var ret *UserLimitResponse
		if err := json.Unmarshal(replyData, &ret); err != nil {
			b.logger.WithError(err).Error("json unmarshal body data failed")
			errorCh <- err
			return
		}
		responseCh <- ret
	}()
	var resp *UserLimitResponse
	select {
	case resp = <-responseCh:
		return resp, nil
	case err := <-errorCh:
		b.logger.WithError(err).Error(err.Error())
		return nil, err
	}
}

func (b *Bcnmy) GetGasTankBalance() (*GasTankBalanceResponse, error) {
	responseCh := make(chan *GasTankBalanceResponse, 1)
	errorCh := make(chan error)
	req, err := http.NewRequest(http.MethodGet, UserLimitURL, nil)
	if err != nil {
		b.logger.WithError(err).Error("GetGasTankBalance NewRequest failed")
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("authToken", b.authToken)
	req.Header.Set("apiKey", b.apiKey)
	go func() {
		res, err := b.httpClient.Do(req)
		if err != nil {
			b.logger.WithError(err).Error("HttpClient request to GetGasTankBalance failed")
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
		var ret *GasTankBalanceResponse
		if err := json.Unmarshal(replyData, &ret); err != nil {
			b.logger.WithError(err).Error("json unmarshal body data failed")
			errorCh <- err
			return
		}
		responseCh <- ret
	}()
	var resp *GasTankBalanceResponse
	select {
	case resp = <-responseCh:
		return resp, nil
	case err := <-errorCh:
		b.logger.WithError(err).Error(err.Error())
		return nil, err
	}
}
