package urls

import (
	"errors"
        "io"
        "time"

        "encoding/json"
	"net/http"

	"shared/logger"
)

func SendRequest(logId, method, uri string, data io.Reader, headers map[string]string, timeout int) (map[string]interface{}, error) {
	logger.Log(logId, "--- urls.SendRequest", logger.FATAL)

	var client *http.Client

	request, err := http.NewRequest(method, uri, data)
	if err != nil {
		return nil, err
	}

	for k, v := range headers {
		request.Header.Set(k, v)
	}

	client = &http.Client{Timeout: time.Duration(timeout) * time.Second}

	resp, err  := client.Do(request)
	defer resp.Body.Close()

	if err != nil {
		return nil, err
	}

	var d map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&d)
	logger.Log(logId, d, logger.FATAL)

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		logger.Log(logId, request.URL       , logger.DEBUG)
		logger.Log(logId, request.Header    , logger.DEBUG)
		logger.Log(logId, request           , logger.DEBUG)

		return d, errors.New(resp.Status)
	}

	return d, nil
}
