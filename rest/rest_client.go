package rest

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// Client TODO: documentation
type Client struct {
	config      *Config
	credentials *Credentials
	httpClient  *http.Client
}

// NewClient TODO: documentation
func NewClient(config *Config, credentials *Credentials) *Client {
	client := Client{}
	client.credentials = credentials
	client.config = config
	client.httpClient = createHTTPClient(config)
	return &client
}

func createHTTPClient(config *Config) *http.Client {
	var httpClient *http.Client
	if config.NoProxy {
		if config.Insecure {
			httpClient = &http.Client{
				Transport: &http.Transport{
					TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
					Proxy:           http.ProxyURL(nil)}}
		} else {
			httpClient = &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(nil)}}
		}
	} else {
		if config.Insecure {
			httpClient = &http.Client{Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}}
		} else {
			httpClient = &http.Client{}
		}
	}
	return httpClient
}

func (client *Client) getURL(path string) string {
	apiBaseURL := client.credentials.APIBaseURL
	if !strings.HasSuffix(apiBaseURL, "/") {
		apiBaseURL = apiBaseURL + "/"
	}
	if strings.HasPrefix(path, "/") {
		path = strings.TrimPrefix(path, "/")
	}
	return apiBaseURL + path
}

// GET TODO: documentation
func (client *Client) GET(path string, expectedStatusCode int) ([]byte, error) {
	var err error
	var httpResponse *http.Response
	var request *http.Request

	url := client.getURL(path)
	if request, err = http.NewRequest(http.MethodGet, url, nil); err != nil {
		return make([]byte, 0), err
	}
	if err = client.credentials.Authenticate(request); err != nil {
		return make([]byte, 0), err
	}

	if httpResponse, err = client.httpClient.Do(request); err != nil {
		return make([]byte, 0), err
	}
	return readHTTPResponse(httpResponse, http.MethodGet, url, expectedStatusCode)
}

// POST TODO: documentation
func (client *Client) POST(path string, payload interface{}, expectedStatusCode int) ([]byte, error) {
	return client.send(path, http.MethodPost, payload, expectedStatusCode)
}

// DELETE TODO: documentation
func (client *Client) DELETE(path string, expectedStatusCode int) ([]byte, error) {
	var err error
	var httpResponse *http.Response
	var request *http.Request

	url := client.getURL(path)
	if request, err = http.NewRequest(http.MethodDelete, url, nil); err != nil {
		return make([]byte, 0), err
	}
	if err = client.credentials.Authenticate(request); err != nil {
		return make([]byte, 0), err
	}

	if httpResponse, err = client.httpClient.Do(request); err != nil {
		return make([]byte, 0), err
	}
	return readHTTPResponse(httpResponse, http.MethodDelete, url, expectedStatusCode)
}

// PUT TODO: documentation
func (client *Client) PUT(path string, payload interface{}, expectedStatusCode int) ([]byte, error) {
	return client.send(path, http.MethodPut, payload, expectedStatusCode)
}

func (client *Client) send(path string, method string, payload interface{}, expectedStatusCode int) ([]byte, error) {
	var err error
	var request *http.Request
	var httpResponse *http.Response
	var requestbody []byte

	if requestbody, err = json.Marshal(payload); err != nil {
		return nil, err
	}

	url := client.getURL(path)
	if request, err = http.NewRequest(method, url, bytes.NewReader(requestbody)); err != nil {
		return nil, err
	}

	if err = client.credentials.Authenticate(request); err != nil {
		return make([]byte, 0), err
	}

	request.Header.Add("Content-Type", "application/json")
	if httpResponse, err = client.httpClient.Do(request); err != nil {
		return nil, err
	}
	return readHTTPResponse(httpResponse, method, url, expectedStatusCode)
}

func readHTTPResponse(httpResponse *http.Response, method string, url string, expectedStatusCode int) ([]byte, error) {
	var err error
	var body []byte
	defer httpResponse.Body.Close()

	if httpResponse.StatusCode != expectedStatusCode {
		finalError := fmt.Errorf("%s (%s) %s", http.StatusText(httpResponse.StatusCode), method, url)
		if body, err = ioutil.ReadAll(httpResponse.Body); err != nil {
			return nil, finalError
		}
		return body, finalError
	}
	if body, err = ioutil.ReadAll(httpResponse.Body); err != nil {
		return nil, err
	}
	return body, nil
}
