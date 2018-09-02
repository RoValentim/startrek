package urls

import (
	"errors"
        "io"
        "time"

        "encoding/json"

	"io/ioutil"

	"net/http"
	"net/http/cookiejar"
	"net/url"

	"shared/defines"
	"shared/logger"
)

func GetDataFromService(logId, url string, headers map[string]string, timeout int) (map[string]interface{}, error) {
	logger.Log(logId, "--- urls.GetDataFromService", logger.FATAL)
        req := &http.Client{Timeout: time.Duration(timeout) * time.Second}

        request, err := http.NewRequest("GET", url, nil)
        if err != nil {
                return nil, err
        }

	for k, v := range headers {
		request.Header.Set(k, v)
	}

        resp, err  := req.Do(request)
	body, _    := request.GetBody()
	reqBody, _ := ioutil.ReadAll(body)
        if err != nil {
		logger.Log(logId, request.URL       , logger.DEBUG)
		logger.Log(logId, request.Header    , logger.DEBUG)
		logger.Log(logId, string(reqBody[:]), logger.FATAL)
		logger.Log(logId, request           , logger.DEBUG)

                return nil, err
        }
	defer resp.Body.Close()

	var d map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&d)

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		logger.Log(logId, request.URL       , logger.DEBUG)
		logger.Log(logId, request.Header    , logger.DEBUG)
		logger.Log(logId, string(reqBody[:]), logger.FATAL)
		logger.Log(logId, request           , logger.DEBUG)

		return d, errors.New(resp.Status)
	}

	return d, nil
}

func PostData(logId, uri string, data io.Reader, headers map[string]string, cookies []*http.Cookie, timeout int) ([]*http.Cookie, map[string]interface{}, error) {
	logger.Log(logId, "--- urls.PostData", logger.FATAL)

	var cooks []*http.Cookie
	var client *http.Client

	expiration := time.Now().Add(defines.TimeoutToken * time.Minute)

	request, err := http.NewRequest("POST", uri, data)
	if err != nil {
		return nil, nil, err
	}

	for k, v := range headers {
		request.Header.Set(k, v)
	}

	cookieLen := len(cookies)
	for i := 0; i < cookieLen; i++ {
		cook := http.Cookie{Name: cookies[i].Name, Value: cookies[i].Value, Expires: expiration}
		cooks = append(cooks, &cook)
	}

	jar, _ := cookiejar.New(nil)
	u,   _ := url.Parse(uri)

	if jar != nil {
		jar.SetCookies(u, cooks)
		client = &http.Client{Jar: jar, Timeout: time.Duration(timeout) * time.Second}
	} else {
		client = &http.Client{Timeout: time.Duration(timeout) * time.Second}
	}

	resp, err  := client.Do(request)
	if err != nil {
		return nil, nil, err
	}

	body,    _ := request.GetBody() // resp.Body
	reqBody, _ := ioutil.ReadAll(body)
	if err != nil {
		logger.Log(logId, request.URL       , logger.DEBUG)
		logger.Log(logId, request.Header    , logger.DEBUG)
		logger.Log(logId, string(reqBody[:]), logger.FATAL)
		logger.Log(logId, request           , logger.DEBUG)

		return nil, nil, err
	}
	defer resp.Body.Close()
	logger.Log(logId, reqBody, logger.FATAL)

	var d map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&d)
	logger.Log(logId, d, logger.FATAL)

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		logger.Log(logId, request.URL       , logger.DEBUG)
		logger.Log(logId, request.Header    , logger.DEBUG)
		logger.Log(logId, string(reqBody[:]), logger.FATAL)
		logger.Log(logId, request           , logger.DEBUG)

		return nil, d, errors.New(resp.Status)
	}

	logger.Log(logId, jar.Cookies(u), logger.DEBUG)
	if jar == nil {
		return nil, d, nil
	}

	return jar.Cookies(u), d, nil
}
