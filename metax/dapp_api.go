package metax

import (
	"context"
	"fmt"
	"net/http"
)

type DappAPIResponse struct {
	Log  string `json:"log"`
	Flag int    `json:"flag"`
	Dapp Dapp   `json:"dapp"`
}

type Dapp struct {
	NetworkId       string    `json:"networkId"`
	DappName        string    `json:"dappName"`
	DappLimit       DappLimit `json:"dappLimit"`
	DappLimitStatus int       `json:"dappLimitStatus"`
	UserLimit       DappLimit `json:"userLimit"`
	UserLimitStatus int       `json:"userLimitStatus"`
	ApiLimit        DappLimit `json:"apiLimit"`
	ApiLimitStatus  int       `json:"apiLimitStatus"`
}

type DappLimit struct {
	Type              int     `json:"type"`
	Value             float32 `json:"value"`
	DurationValue     int     `json:"durationValue"`
	DurationUnit      string  `json:"durationUnit"`
	LimitStartTime    int64   `json:"limitStartTime"`
	LimitDurationInMs int64   `json:"limitDurationInMs"`
}

func (b *Bcnmy) GetDappAPI(ctx context.Context) (*DappAPIResponse, error) {
	responseCh := make(chan interface{}, 1)
	errorCh := make(chan error)
	defer close(errorCh)
	defer close(responseCh)

	req, err := http.NewRequest(http.MethodGet, DappURL, nil)
	if err != nil {
		b.logger.WithError(err).Error("DappAPI NewRequest failed")
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("x-api-key", b.apiKey)
	var resp DappAPIResponse
	b.asyncHttpx(req, &resp, errorCh, responseCh)
	select {
	case ret := <-responseCh:
		resp, ok := ret.(*DappAPIResponse)
		if !ok {
			return nil, fmt.Errorf("DappAPI failed")
		}
		if resp.Flag != 143 {
			err := fmt.Errorf("%v", resp)
			b.logger.WithError(err).Error(resp.Log)
			return nil, err
		}
		return resp, nil
	case err := <-errorCh:
		b.logger.Error(err.Error())
		return nil, err
	}
}
