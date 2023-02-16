package metax

import (
	"context"
	"fmt"
	"net/http"
)

type MetaAPIResponse struct {
	Log     string        `json:"log"`
	Flag    int           `json:"flag"`
	Total   int           `json:"total"`
	ListAPI []MetaAPIInfo `json:"listApis"`
}

type MetaAPIInfo struct {
	/// need to filter non contractAdress
	ContractAddress   string      `json:"contractAddress"`
	ID                string      `json:"id"`
	Name              string      `json:"name"`
	URL               string      `json:"url"`
	Version           int         `json:"version"`
	Method            string      `json:"method"`
	MethodType        string      `json:"methodType"`
	APIType           string      `json:"apiType"`
	MetaTxLimitStatus int         `json:"metaTxLimitStatus"`
	MetaTxLimit       MetaTxLimit `json:"metaTxLimit"`
}

type MetaTxLimit struct {
	Type              int     `json:"type"`
	Value             float32 `json:"value"`
	DurationValue     int     `json:"durationValue"`
	Day               string  `json:"day"`
	LimitStartTime    int64   `json:"limitStartTime"`
	LimitDurationInMs int64   `json:"limitDurationInMs"`
}

func (b *Bcnmy) GetMetaAPI(ctx context.Context) (*MetaAPIResponse, error) {
	responseCh := make(chan interface{}, 1)
	errorCh := make(chan error)
	req, err := http.NewRequest(http.MethodGet, MetaAPIURL, nil)
	if err != nil {
		b.logger.WithError(err).Error("MetaAPI NewRequest failed")
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("x-api-key", b.apiKey)
	var resp MetaAPIResponse
    b.asyncHttpx(req, &resp, errorCh, responseCh)
	select {
    case ret := <-responseCh:
        resp, ok := ret.(*MetaAPIResponse)
        if !ok {
            return nil, fmt.Errorf("MetaAPI failed")
        }
        if resp.Flag != 143 {
            err := fmt.Errorf("%v", resp)
            b.logger.WithError(err).Error(resp.Log)
            return nil, err
        }
        return resp, nil
	case err := <-errorCh:
		b.logger.WithError(err).Error(err.Error())
		return nil, err
	}
}
