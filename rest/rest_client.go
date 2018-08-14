package rest

import (
	"io/ioutil"
	"net/http"
)

// Client TODO: documentation
type Client struct {
	credentials *Credentials
}

// NewClient TODO: documentation
func NewClient(credentials *Credentials) *Client {
	client := Client{}
	client.credentials = credentials
	return &client
}

// Get TODO: documentation
func (client *Client) Get(path string) ([]byte, error) {
	var err error
	var httpResponse *http.Response
	var request *http.Request
	if request, err = http.NewRequest(http.MethodGet, client.credentials.BaseURL()+path, nil); err != nil {
		return make([]byte, 0), err
	}
	if err = client.credentials.Authenticate(request); err != nil {
		return make([]byte, 0), err
	}
	httpClient := &http.Client{}
	if httpResponse, err = httpClient.Do(request); err != nil {
		return make([]byte, 0), err
	}
	return ioutil.ReadAll(httpResponse.Body)
}
