package managementzones

// GetManagementZonesResponse is the expected JSON response
// when querying for the currently configured Management Zones via Dynatrace REST API
// The response contains just Stubs, not the full configuration
type GetManagementZonesResponse struct {
	Values []Stub `json:"values,omitempty"` // the Stubs of the currently configured Management Zones
}
