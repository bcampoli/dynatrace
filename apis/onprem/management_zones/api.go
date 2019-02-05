package managementzones

import (
	"encoding/json"

	resterrors "github.com/dtcookie/dynatrace/apis/errors"
	"github.com/dtcookie/dynatrace/rest"
)

// API TODO: documentation
type API struct {
	client *rest.Client
}

// NewAPI creates a preconfigured API for accessing the Users API
// of an OnPremise Dynatrace Cluster
func NewAPI(config *rest.Config, credentials *rest.Credentials) *API {
	return &API{client: rest.NewClient(config, credentials)}
}

// GetPermissionsForGroup queries for the configured management zone permissions for a given group
func (api *API) GetPermissionsForGroup(groupID string) ([]PermissionsForGroup, error) {
	var err error
	var bytes []byte

	if bytes, err = api.client.GET("/api/v1.0/onpremise/groups/managementZones/"+groupID, 200); err != nil {
		return nil, resterrors.Resolve(bytes, err)
	}
	var response []PermissionsForGroup
	if err = json.Unmarshal(bytes, &response); err != nil {
		return nil, err
	}
	return response, nil
}

// GetPermissions queries for the configured management zone permissions for all groups
func (api *API) GetPermissions() ([]PermissionsForGroup, error) {
	var err error
	var bytes []byte

	if bytes, err = api.client.GET("/api/v1.0/onpremise/groups/managementZones", 200); err != nil {
		return nil, resterrors.Resolve(bytes, err)
	}
	var response []PermissionsForGroup
	if err = json.Unmarshal(bytes, &response); err != nil {
		return nil, err
	}
	return response, nil
}

// SetPermissionsForGroup Get management zone permissions for a given group
func (api *API) SetPermissionsForGroup(permissions PermissionsForGroup) ([]PermissionsForGroup, error) {
	var err error
	var bytes []byte

	if bytes, err = api.client.PUT("/api/v1.0/onpremise/groups/managementZones", permissions, 200); err != nil {
		return nil, resterrors.Resolve(bytes, err)
	}
	var response []PermissionsForGroup
	if err = json.Unmarshal(bytes, &response); err != nil {
		return nil, err
	}
	return response, nil
}
