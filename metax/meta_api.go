package metax

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
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
	metaApiCh := make(chan *MetaAPIResponse, 1)
	errorCh := make(chan error)
	req, err := http.NewRequest(http.MethodGet, MetaAPIURL, nil)
	if err != nil {
		b.logger.WithError(err).Error("MetaAPI NewRequest failed")
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("x-api-key", b.apiKey)
	go func() {
		res, err := b.httpClient.Do(req)
		if err != nil {
			b.logger.WithError(err).Error("HttpClient request to MetaAPI failed")
			errorCh <- err
			return
		}
		defer res.Body.Close()
		data, err := io.ReadAll(res.Body)
		if err != nil {
			b.logger.WithError(err).Error("io read request body failed")
			errorCh <- err
			return
		}
		var ret MetaAPIResponse
		if err := json.Unmarshal(data, &ret); err != nil {
			b.logger.WithError(err).Error("json unmarshal body data failed")
			errorCh <- err
			return
		}
		metaApiCh <- &ret
	}()

	var resp *MetaAPIResponse
	select {
	case resp = <-metaApiCh:
	case err := <-errorCh:
		b.logger.WithError(err).Error(err.Error())
		return nil, err
	}
	if resp.Flag != 143 {
		err := fmt.Errorf("%v", resp)
		b.logger.WithError(err).Error(resp.Log)
		return nil, err
	}
	return resp, nil
}
