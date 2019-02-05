package managementzones

import (
	"encoding/json"
	"errors"

	resterrors "github.com/dtcookie/dynatrace/apis/errors"
	"github.com/dtcookie/dynatrace/rest"
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

func unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

func resolveError(bytes []byte, err error) error {
	if bytes != nil {
		var errorEnvelope resterrors.ErrorEnvelope
		var innerError error
		if innerError = json.Unmarshal(bytes, &errorEnvelope); innerError == nil {
			if errorEnvelope.Error.Message != "" {
				return errors.New(errorEnvelope.Error.Message)
			} else {
				return err
			}
		}
	}
	return err
}

// GetManagementZones queries for the existing Management Zones
// delivers only Stubs, not the actual configuration
func (api *API) GetManagementZones() ([]Stub, error) {
	var err error
	var bytes []byte

	if bytes, err = api.client.GET("/api/config/v1/managementZones", 200); err != nil {
		return nil, resolveError(bytes, err)
	}
	var response GetManagementZonesResponse
	if err = unmarshal(bytes, &response); err != nil {
		return nil, err
	}
	return response.Values, nil
}

// CreateManagementZone creates a new Management Zone
func (api *API) CreateManagementZone(zone ManagementZone) (*Stub, error) {
	var err error
	var bytes []byte

	if bytes, err = api.client.POST("/api/config/v1/managementZones", &zone, 201); err != nil {
		return nil, resolveError(bytes, err)
	}
	var stub Stub
	if err = unmarshal(bytes, &stub); err != nil {
		return nil, err
	}
	return &stub, nil
}
