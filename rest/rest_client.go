package rest

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
)

// Client TODO: documentation
type Client struct {
	config      *Config
	credentials *Credentials
}

// NewClient TODO: documentation
func NewClient(config *Config, credentials *Credentials) *Client {
	client := Client{}
	client.credentials = credentials
	client.config = config
	return &client
}

func (client *Client) createHTTPClient() *http.Client {
	var httpClient *http.Client
	if client.config.NoProxy {
		if client.config.Insecure {
			httpClient = &http.Client{
				Transport: &http.Transport{
					TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
					Proxy:           http.ProxyURL(nil),
				},
			}
		} else {
			httpClient = &http.Client{
				Transport: &http.Transport{
					Proxy: http.ProxyURL(nil),
				},
			}
		}
	} else {
		if client.config.Insecure {
			httpClient = &http.Client{
				Transport: &http.Transport{
					TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
				},
			}
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

// Get TODO: documentation
func (client *Client) Get(path string) ([]byte, error) {
	var err error
	var httpResponse *http.Response
	var request *http.Request
	var bytes []byte

	url := client.getURL(path)
	if request, err = http.NewRequest(http.MethodGet, url, nil); err != nil {
		return make([]byte, 0), err
	}
	if err = client.credentials.Authenticate(request); err != nil {
		return make([]byte, 0), err
	}

	httpClient := client.createHTTPClient()

	if httpResponse, err = httpClient.Do(request); err != nil {
		return make([]byte, 0), err
	}
	if httpResponse.StatusCode != http.StatusOK {
		finalError := errors.New(http.StatusText(httpResponse.StatusCode) + " (GET " + url + ")")
		if bytes, err = ioutil.ReadAll(httpResponse.Body); err != nil {
			return nil, finalError
		}
		return bytes, finalError
	}
	return ioutil.ReadAll(httpResponse.Body)
}

// Post TODO: documentation
func (client *Client) Post(path string, payload interface{}) ([]byte, error) {
	var err error
	var request *http.Request
	var response *http.Response
	var requestbody []byte
	var body []byte

	var httpClient *http.Client
	httpClient = client.createHTTPClient()

	if requestbody, err = json.Marshal(payload); err != nil {
		return nil, err
	}

	url := client.getURL(path)
	if request, err = http.NewRequest("POST", url, bytes.NewReader(requestbody)); err != nil {
		return nil, err
	}

	if err = client.credentials.Authenticate(request); err != nil {
		return make([]byte, 0), err
	}

	request.Header.Add("Content-Type", "application/json")
	if response, err = httpClient.Do(request); err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if body, err = ioutil.ReadAll(response.Body); err != nil {
		return nil, err
	}
	return body, nil
}
