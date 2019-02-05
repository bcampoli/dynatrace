package usersgroups

import (
	"encoding/json"

	resterrors "github.com/dtcookie/dynatrace/apis/errors"
	"github.com/dtcookie/dynatrace/rest"
)

// API is able to make REST API Calls to the Users API of an
// OnPremise Dynatrace Cluster
type API struct {
	client *rest.Client
}

// NewAPI creates a preconfigured API for accessing the Users API
// of an OnPremise Dynatrace Cluster
func NewAPI(config *rest.Config, credentials *rest.Credentials) *API {
	return &API{client: rest.NewClient(config, credentials)}
}

// All queries for the currently configured users
func (api *API) All() ([]GroupConfig, error) {
	var err error
	var bytes []byte

	if bytes, err = api.client.GET("/api/v1.0/onpremise/groups", 200); err != nil {
		return nil, resterrors.Resolve(bytes, err)
	}
	var response []GroupConfig
	if err = json.Unmarshal(bytes, &response); err != nil {
		return nil, err
	}
	return response, nil
}

// Create queries for the currently configured users
func (api *API) Create(groupConfig GroupConfig) (*GroupConfig, error) {
	var err error
	var bytes []byte

	if bytes, err = api.client.POST("/api/v1.0/onpremise/groups", groupConfig, 200); err != nil {
		return nil, resterrors.Resolve(bytes, err)
	}
	var response GroupConfig
	if err = json.Unmarshal(bytes, &response); err != nil {
		return nil, err
	}
	return &response, nil
}
