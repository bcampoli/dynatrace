package cluster

import (
	"encoding/json"
	"errors"

	resterrors "github.com/bcampoli/dynatrace/apis/errors"
	"github.com/bcampoli/dynatrace/rest"
)

// API TODO: documentation
type API struct {
	client *rest.Client
}

// NewAPI TODO: documentation
func NewAPI(config *rest.Config, credentials *rest.Credentials) *API {
	return &API{
		client: rest.NewClient(config, credentials),
	}
}

// Get TODO: documentation
func (api *API) Get() (string, error) {
	var err error
	var bytes []byte
	var version Version
	var errorEnvelope resterrors.ErrorEnvelope

	if bytes, err = api.client.GET("/api/v1/config/clusterversion", 200); err != nil {
		if bytes != nil {
			var innerError error
			if innerError = json.Unmarshal(bytes, &errorEnvelope); innerError == nil {
				if errorEnvelope.Error.Message != "" {
					return "", errors.New(errorEnvelope.Error.Message)
				} else {
					return "", err
				}
			}
		}
		return "", err
	}
	if err = json.Unmarshal(bytes, &version); err != nil {
		return "", err
	}
	return version.Version, nil
}
