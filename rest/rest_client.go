package rest

import (
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
)

// Client TODO: documentation
type Client struct {
	credentials *Credentials
}

// NewClient TODO: documentation
func NewClient(credentials *Credentials) *Client {
	// http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	// tr := &http.Transport{
	// 	TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	// }
	client := Client{}
	client.credentials = credentials
	return &client
}

// Get TODO: documentation
func (client *Client) Get(path string) ([]byte, error) {
	var err error
	var httpResponse *http.Response
	var request *http.Request
	var bytes []byte

	apiBaseURL := client.credentials.APIBaseURL
	if !strings.HasSuffix(apiBaseURL, "/") {
		apiBaseURL = apiBaseURL + "/"
	}
	if strings.HasPrefix(path, "/") {
		path = strings.TrimPrefix(path, "/")
	}
	if request, err = http.NewRequest(http.MethodGet, apiBaseURL+path, nil); err != nil {
		return make([]byte, 0), err
	}
	if err = client.credentials.Authenticate(request); err != nil {
		return make([]byte, 0), err
	}
	httpClient := &http.Client{
	// Transport: &http.Transport{
	// 	TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	// 	Proxy:           http.ProxyURL(nil),
	// },
	}
	if httpResponse, err = httpClient.Do(request); err != nil {
		return make([]byte, 0), err
	}
	if httpResponse.StatusCode != http.StatusOK {
		finalError := errors.New(http.StatusText(httpResponse.StatusCode) + " (GET " + apiBaseURL + path + ")")
		if bytes, err = ioutil.ReadAll(httpResponse.Body); err != nil {
			return nil, finalError
		} else {
			return bytes, finalError
		}
	}
	return ioutil.ReadAll(httpResponse.Body)
}
