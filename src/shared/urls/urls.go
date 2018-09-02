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

func DeleteDataFromService(logId, url string, headers map[string]string, timeout int) (map[string]interface{}, error) {
	logger.Log(logId, "--- urls.DeleteDataFromService", logger.FATAL)
	req := &http.Client{Timeout: time.Duration(timeout) * time.Second}

	request, err := http.NewRequest("DELETE", url, nil)
        if err != nil {
                return nil, err
        }

	for k, v := range headers {
		request.Header.Set(k, v)
	}

        resp, err  := req.Do(request)
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		logger.Log(logId, request.URL       , logger.DEBUG)
		logger.Log(logId, request.Header    , logger.DEBUG)
		logger.Log(logId, request           , logger.DEBUG)

		return nil, errors.New(resp.Status)
	}

	body, _      := request.GetBody()
	reqBody, err := ioutil.ReadAll(body)
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

func GetDataFromService(logId, url string, timeout int) (string, []*http.Cookie, int, error) {
	logger.Log(logId, "--- urls.GetDataFromService", logger.FATAL)
        req := &http.Client{Timeout: time.Duration(timeout) * time.Second}

        resp, err  := req.Get(url)
        if err != nil {
                return "", nil, 1, err
        }
	defer resp.Body.Close()

	cookies := resp.Cookies()
	logger.Log(logId, cookies, logger.FATAL)

        var data map[string]interface{}
        json.NewDecoder(resp.Body).Decode(&data)

	if data != nil {
		logger.Log(logId, data, logger.FATAL)

		if data["status"] == nil || data["data"] == nil {
			return "", nil, 1, err
			return "", nil, 1, errors.New("Data in invalid format from service")
		}

		if resp.StatusCode < 200 || resp.StatusCode > 299 {
			return "", nil, int(data["status"].(float64)), errors.New(resp.Status)
		}

		return data["data"].(string), cookies, int(data["status"].(float64)), nil
	} else {
		responseBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return "", nil, 1, err
		}

		responseString := string(responseBody)
		logger.Log(logId, responseString, logger.FATAL)

		if resp.StatusCode < 200 || resp.StatusCode > 299 {
			return "", nil, 1, errors.New(resp.Status)
		}

		return responseString, cookies, 0, nil
	}
}

func GetDataFromServiceWithCustomHeader(logId, url string, headers map[string]string, timeout int) (map[string]interface{}, error) {
	logger.Log(logId, "--- urls.GetDataFromServiceWithCustomHeader", logger.FATAL)
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

func GetStatusCodeFromService(logId, url string, timeout int) (int, error) {
	logger.Log(logId, "--- urls.GetStatusCodeFromService", logger.FATAL)
        req := &http.Client{Timeout: time.Duration(timeout) * time.Second}

        request, err := http.NewRequest("GET", url, nil)
        if err != nil {
                return 0, err
        }

        resp, err  := req.Do(request)
	body, _    := request.GetBody()
	reqBody, _ := ioutil.ReadAll(body)
        if err != nil {
		logger.Log(logId, request.URL       , logger.DEBUG)
		logger.Log(logId, request.Header    , logger.DEBUG)
		logger.Log(logId, string(reqBody[:]), logger.FATAL)
		logger.Log(logId, request           , logger.DEBUG)

                return 0, err
        }
	defer resp.Body.Close()

	var d map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&d)

	return resp.StatusCode, nil
}

func UpdateDataFromService(logId, url string, timeout int, form io.Reader) (int, string, error) {
	logger.Log(logId, "--- urls.UpdateDataFromService", logger.FATAL)
        req := &http.Client{Timeout: time.Duration(timeout) * time.Second}

        request, err := http.NewRequest("PUT", url, form)

        if err != nil {
                return 1, "", err
        }

        request.ContentLength = 23

        resp, err  := req.Do(request)
	body, _    := request.GetBody()
	reqBody, _ := ioutil.ReadAll(body)
        if err != nil {
		logger.Log(logId, request.URL       , logger.DEBUG)
		logger.Log(logId, request.Header    , logger.DEBUG)
		logger.Log(logId, string(reqBody[:]), logger.FATAL)
		logger.Log(logId, request           , logger.DEBUG)

                return 1, "", err
        }
	defer resp.Body.Close()

        var d map[string]interface{}
        json.NewDecoder(resp.Body).Decode(&d)

	if d["status"] == nil || d["data"] == nil {
		logger.Log(logId, request.URL       , logger.DEBUG)
		logger.Log(logId, request.Header    , logger.DEBUG)
		logger.Log(logId, string(reqBody[:]), logger.FATAL)
		logger.Log(logId, request           , logger.DEBUG)

		return 1, "", errors.New("Data in invalid format from service")
	}

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		logger.Log(logId, request.URL       , logger.DEBUG)
		logger.Log(logId, request.Header    , logger.DEBUG)
		logger.Log(logId, string(reqBody[:]), logger.FATAL)
		logger.Log(logId, request           , logger.DEBUG)

		return int(d["status"].(float64)), d["data"].(string), errors.New(resp.Status)
	}

        return int(d["status"].(float64)), d["data"].(string), nil
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
