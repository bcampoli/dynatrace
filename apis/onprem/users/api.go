package users

import (
	"encoding/json"

	resterrors "github.com/dtcookie/dynatrace/apis/errors"
	"github.com/dtcookie/dynatrace/rest"
)

// EndPoint is the HTTP path commonly used for this API
const EndPoint = "/api/v1.0/onpremise/users"

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

// GetUsers queries for the currently configured users
func (api *API) GetUsers() ([]UserConfig, error) {
	var err error
	var bytes []byte

	if bytes, err = api.client.Get(EndPoint); err != nil {
		return nil, resterrors.Resolve(bytes, err)
	}
	var response []UserConfig
	if err = json.Unmarshal(bytes, &response); err != nil {
		return nil, err
	}
	return response, nil
}
