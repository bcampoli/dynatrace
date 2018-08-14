package rest

import (
	"net/http"
)

// Credentials TODO: documentation
type Credentials struct {
	APIBaseURL string `json:"api-base-url,omitempty"`
	APIToken   string `json:"api-token,omitempty"`
}

// NewCredentials TODO: documentation
func NewCredentials(apiBaseURL string, apiToken string) *Credentials {
	return &Credentials{
		APIBaseURL: apiBaseURL,
		APIToken:   apiToken,
	}
}

// Authenticate TODO: documentation
func (credentials *Credentials) Authenticate(request *http.Request) error {
	request.Header.Set("Authorization", "Api-Token "+credentials.APIToken)
	return nil
}
