package managementzones

import (
	"encoding/json"

	resterrors "github.com/bcampoli/dynatrace/apis/errors"
	"github.com/bcampoli/dynatrace/rest"
)

// API TODO: documentation
type API struct {
	client *rest.Client
}

// NewAPI TODO: documentation
func NewAPI(config *rest.Config, credentials *rest.Credentials) *API {
	return &API{client: rest.NewClient(config, credentials)}
}

// WithClient TODO: documentation
func (api *API) WithClient(client *rest.Client) *API {
	api.client = client
	return api
}

type getManagementZonesResponse struct {
	Values []Stub `json:"values,omitempty"` // the Stubs of the currently configured Management Zones
}

// List queries for the existing Management Zones
// delivers only Stubs, not the actual configuration
func (api *API) List() ([]Stub, error) {
	var err error
	var response getManagementZonesResponse
	var bytes []byte

	if bytes, err = api.client.GET("/api/config/v1/managementZones", 200); err != nil {
		return nil, resterrors.Resolve(bytes, err)
	}
	if err = json.Unmarshal(bytes, &response); err != nil {
		return nil, err
	}
	return response.Values, nil
}

// Get queries for the existing Management Zones
// delivers only Stubs, not the actual configuration
func (api *API) Get(ID string) (*Stub, error) {
	var err error
	var response Stub
	var bytes []byte

	if bytes, err = api.client.GET("/api/config/v1/managementZones/"+ID, 200); err != nil {
		return nil, resterrors.Resolve(bytes, err)
	}
	if err = json.Unmarshal(bytes, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// Create creates a new Management Zone
func (api *API) Create(zone ManagementZone) (*Stub, error) {
	var err error
	var bytes []byte

	if bytes, err = api.client.POST("/api/config/v1/managementZones", &zone, 201); err != nil {
		return nil, resterrors.Resolve(bytes, err)
	}
	var stub Stub
	if err = json.Unmarshal(bytes, &stub); err != nil {
		return nil, err
	}
	return &stub, nil
}
