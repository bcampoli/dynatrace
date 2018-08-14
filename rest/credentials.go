package rest

import (
	"net/http"
)

// Credentials TODO: documentation
type Credentials struct {
	Cluster       string `json:"cluster,omitempty"`
	EnvironmentID string `json:"environment,omitempty"`
	APIToken      string `json:"api-token,omitempty"`
}

// NewSaasCredentials TODO: documentation
func NewSaasCredentials(environmentID string, apiToken string) *Credentials {
	return &Credentials{
		Cluster:       "",
		EnvironmentID: environmentID,
		APIToken:      apiToken,
	}
}

// NewManagedCredentials TODO: documentation
func NewManagedCredentials(cluster string, environmentID string, apiToken string) *Credentials {
	return &Credentials{
		Cluster:       cluster,
		EnvironmentID: environmentID,
		APIToken:      apiToken,
	}
}

// BaseURL TODO: documentation
func (credentials *Credentials) BaseURL() string {
	if credentials.Cluster == "" {
		return "https://" + credentials.EnvironmentID + ".live.dynatrace.com"
	}
	return "https://" + credentials.Cluster + credentials.EnvironmentID
}

// Authenticate TODO: documentation
func (credentials *Credentials) Authenticate(request *http.Request) error {
	request.Header.Set("Authorization", "Api-Token "+credentials.APIToken)
	return nil
}
