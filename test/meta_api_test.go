package test

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"

	metax "github.com/oblzh/bcnmy-go/metax"
)

type MetaAPIResponse struct {
	Log     string        `json:"log"`
	Flag    int           `json:"flag"`
	Total   int           `json:"total"`
	ListAPI []MetaAPIInfo `json:"listApis"`
}

type MetaAPIInfo struct {
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

func TestMetaAPI(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, metax.MetaAPIURL, nil)
	metaApiCh := make(chan *MetaAPIResponse)
	errorCh := make(chan error)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("x-api-key", os.Getenv("apiKey"))
	httpClient := &http.Client{}

	go func() {
		res, err := httpClient.Do(req)
		if err != nil {
			fmt.Println("HttpClient request to MetaAPI failed")
			errorCh <- err
			return
		}
		defer res.Body.Close()
		data, err := io.ReadAll(res.Body)
		if err != nil {
			fmt.Println("io read request body failed")
			errorCh <- err
			return
		}
		var ret MetaAPIResponse
		if err := json.Unmarshal(data, &ret); err != nil {
			fmt.Println("json unmarshal body data failed")
			errorCh <- err
			return
		}
		metaApiCh <- &ret
	}()
	var resp *MetaAPIResponse
	select {
	case resp = <-metaApiCh:
		{
		}
	case err := <-errorCh:
		fmt.Println(err.Error())
	}
	if resp != nil && resp.Flag != 143 {
		panic(fmt.Errorf("%v", resp))
	}
}
