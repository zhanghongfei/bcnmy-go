package metax

import (
    "encoding/json"
    "io"
    "net/http"
)

func (b *Bcnmy) asyncHttpx(req *http.Request, resp interface{}, errorCh chan error, responseCh chan interface{}) {
	go func() {
		res, err := b.httpClient.Do(req)
		if err != nil {
			b.logger.WithError(err).Error("HttpClient request to failed")
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
		if err := json.Unmarshal(replyData, resp); err != nil {
			b.logger.WithError(err).Error("json unmarshal body data failed")
			errorCh <- err
			return
		}
		responseCh <- resp
	}()
}
