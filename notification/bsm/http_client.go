package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

// HTTPClient TODO: documentation
type HTTPClient struct {
}

// Post TODO: documentation
func (client *HTTPClient) Post(URL string, body []byte) error {
	var err error
	var response *http.Response
	var request *http.Request
	if request, err = http.NewRequest(http.MethodPost, URL, bytes.NewBuffer(body)); err != nil {
		return err
	}
	request.Header.Set("Content-Type", "application/xml")

	httpClient := &http.Client{}
	if response, err = httpClient.Do(request); err != nil {
		return err
	}
	defer response.Body.Close()
	ioutil.ReadAll(response.Body)
	return nil
}
