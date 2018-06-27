package httprequest

import (
	"strings"
	"io/ioutil"
	"net/http"
	"encoding/json"
	"bytes"
	"bufio"
)

type HttpRequest struct {
	Host       string
	ApiVersion string
}

const (
	MainNet  = "https://mainnet.nebulas.io"
	TestNet  = "https://testnet.nebulas.io"
	LocalNet = "http://localhost:8685"

	APIVersion1 = "v1"
)

func NewHttpRequest(host string, apiVersion string) *HttpRequest {
	return &HttpRequest{host, apiVersion}
}

func (req *HttpRequest) CreateUrl(api string) string {
	return req.Host + "/" + req.ApiVersion + api
}

func (req *HttpRequest) Get(api string, params map[string]string) ([]byte, error) {
	url := req.CreateUrl(api)
	if params != nil {
		url += "?"
		for k, v := range params {
			url += k + "=" + v + "&"
		}
		url = strings.TrimSuffix(url, "&")
	}

	response, err := http.Get(url)
	defer response.Body.Close()

	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	return body, nil
}

func (req *HttpRequest) Post(api string, reqBody interface{}) ([]byte, error) {
	url := req.CreateUrl(api)
	contentType := "application/json"

	jsonReqBody, err := json.Marshal(reqBody)

	if err != nil {
		return nil, err
	}

	response, err := http.Post(url, contentType, bytes.NewBuffer(jsonReqBody))
	defer response.Body.Close()

	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	return body, nil
}

func (req *HttpRequest) PostStream(api string, reqBody interface{}, callback func([]byte)) (error) {
	url := req.CreateUrl(api)
	contentType := "application/json"

	jsonReqBody, err := json.Marshal(reqBody)

	if err != nil {
		return err
	}

	response, err := http.Post(url, contentType, bytes.NewBuffer(jsonReqBody))
	defer response.Body.Close()

	if err != nil {
		return err
	}

	reader := bufio.NewReader(response.Body)
	for {
		line, err := reader.ReadBytes('\n')

		if err != nil {
			break
		}

		callback(line)
	}

	if err != nil {
		return err
	}

	return nil
}
