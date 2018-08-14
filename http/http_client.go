package http

import (
	"bytes"
	"io/ioutil"
	nethttp "net/http"
)

// Client TODO: documentation
type Client struct {
}

// Post TODO: documentation
func (client *Client) Post(URL string, body []byte) error {
	var err error
	var response *nethttp.Response
	var request *nethttp.Request
	if request, err = nethttp.NewRequest(nethttp.MethodPost, URL, bytes.NewBuffer(body)); err != nil {
		return err
	}
	request.Header.Set("Content-Type", "application/xml")

	httpClient := &nethttp.Client{}
	if response, err = httpClient.Do(request); err != nil {
		return err
	}
	defer response.Body.Close()
	ioutil.ReadAll(response.Body)
	return nil
}
