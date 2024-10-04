package common

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"
)

var sleepTimes int
var mu sync.Mutex

func DoRequest[R any](l AccessTokenLoader, c *http.Client, req *http.Request) (*R, error) {
	if sleepTimes > 0 {
		mu.Lock()
		sleepTimes = 0
		mu.Unlock()
	}

	if l != nil {
		token, err := l.LoadAccessToken()
		if err != nil {
			return nil, err
		}
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", *token))
	}

	var reqBodyBytes []byte
	if req.Body != nil {
		body, err := io.ReadAll(req.Body)
		if err != nil {
			return nil, err
		}
		reqBodyBytes = body
		req.Body = io.NopCloser(bytes.NewReader(reqBodyBytes))
	}

	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("close body error: ", err)
		}
	}(resp.Body)
	buff, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusTooManyRequests {
			xRetryAfter, err := strconv.Atoi(resp.Header.Get("x-retry-after"))
			if err != nil {
				return nil, err
			}

			log.Println(req.URL.Path, "x-retry-after", xRetryAfter)

			mu.Lock()
			sleepTimes = xRetryAfter
			time.Sleep(time.Duration(sleepTimes) * time.Millisecond)
			mu.Unlock()

			if reqBodyBytes != nil {
				req.Body = io.NopCloser(bytes.NewReader(reqBodyBytes))
			}
			return DoRequest[R](l, c, req)
		} else {
			return nil, fmt.Errorf(
				"http status code: %v | http response body: %v",
				resp.StatusCode, string(buff),
			)
		}
	}
	var result R
	err = json.Unmarshal(buff, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func DoFormRequest[R any](l AccessTokenLoader, c *http.Client, apiUrl string, formValues url.Values) (*R, error) {
	req, err := http.NewRequest(
		"POST",
		apiUrl,
		strings.NewReader(formValues.Encode()),
	)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	return DoRequest[R](l, c, req)
}

func DoJsonRequest[R any](l AccessTokenLoader, c *http.Client, apiUrl string, params interface{}) (*R, error) {
	buff, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(
		"POST",
		apiUrl,
		strings.NewReader(string(buff)),
	)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	return DoRequest[R](l, c, req)
}

func DoGetQuery[R any](l AccessTokenLoader, c *http.Client, apiUrl string, queryValues url.Values) (*R, error) {
	if len(queryValues) > 0 {
		apiUrl += "?" + queryValues.Encode()
	}
	req, err := http.NewRequest("GET", apiUrl, nil)
	if err != nil {
		return nil, err
	}
	return DoRequest[R](l, c, req)
}

func DoPostNoBody[R any](l AccessTokenLoader, c *http.Client, apiUrl string) (*R, error) {
	req, err := http.NewRequest("POST", apiUrl, nil)
	if err != nil {
		return nil, err
	}
	return DoRequest[R](l, c, req)
}
