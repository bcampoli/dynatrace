package rest

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"strings"

	"github.com/dtcookie/dynatrace/libdtlog"
)

var Verbose = false

var jar = createJar()

func createJar() *cookiejar.Jar {
	jar, _ := cookiejar.New(nil)
	return jar
}

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
	httpClient.Jar = jar
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
	if Verbose {
		dtlog.Println(fmt.Sprintf("GET %s", url))
	}
	if request, err = http.NewRequest(http.MethodGet, url, nil); err != nil {
		return make([]byte, 0), err
	}
	if err = client.credentials.Authenticate(request); err != nil {
		return make([]byte, 0), err
	}

	if httpResponse, err = client.httpClient.Do(request); err != nil {
		return make([]byte, 0), err
	}
	return readHTTPResponse(httpResponse, http.MethodGet, url, expectedStatusCode, nil)
}

// NewPOST TODO: documentation
func (client *Client) NewPOST(path string, payload interface{}) *Post {
	return newPost(client, path, payload)
}

// NewPUT TODO: documentation
func (client *Client) NewPUT(path string, payload interface{}) *Put {
	return newPut(client, path, payload)
}

// POST TODO: documentation
func (client *Client) POST(path string, payload interface{}, expectedStatusCode int) ([]byte, error) {
	return client.send(path, http.MethodPost, payload, expectedStatusCode, nil)
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
	return readHTTPResponse(httpResponse, http.MethodDelete, url, expectedStatusCode, nil)
}

// PUT TODO: documentation
func (client *Client) PUT(path string, payload interface{}, expectedStatusCode int) ([]byte, error) {
	return client.send(path, http.MethodPut, payload, expectedStatusCode, nil)
}

func (client *Client) send(path string, method string, payload interface{}, expectedStatusCode int, onResponse func(int) error) ([]byte, error) {
	var err error
	var request *http.Request
	var httpResponse *http.Response
	var requestbody []byte

	if requestbody, err = json.Marshal(payload); err != nil {
		return nil, err
	}

	url := client.getURL(path)
	if Verbose {
		dtlog.Println(fmt.Sprintf("%s %s", strings.ToUpper(method), url))
	}
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
	return readHTTPResponse(httpResponse, method, url, expectedStatusCode, onResponse)
}

func readHTTPResponse(httpResponse *http.Response, method string, url string, expectedStatusCode int, onResponse func(int) error) ([]byte, error) {
	var err error
	var body []byte
	defer httpResponse.Body.Close()

	if Verbose {
		dtlog.Println(fmt.Sprintf("  %d %s", httpResponse.StatusCode, http.StatusText(httpResponse.StatusCode)))
	}

	if onResponse != nil {
		if err = onResponse(httpResponse.StatusCode); err != nil {
			return nil, err
		}
	}

	if httpResponse.StatusCode != expectedStatusCode {
		finalError := fmt.Errorf("%s (%s) %s", http.StatusText(httpResponse.StatusCode), method, url)
		if body, err = ioutil.ReadAll(httpResponse.Body); err != nil {
			return nil, finalError
		}
		if Verbose && (body != nil) && len(body) > 0 {
			dtlog.Println("  " + string(body))
		}
		return body, finalError
	}
	if body, err = ioutil.ReadAll(httpResponse.Body); err != nil {
		return nil, err
	}
	if Verbose && (body != nil) && len(body) > 0 {
		dtlog.Println("  " + string(body))
	}
	return body, nil
}
