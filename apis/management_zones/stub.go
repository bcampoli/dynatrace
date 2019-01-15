package managementzones

// Stub is a ManagementZoneStub
// Only holds the ID and Name of the Management Zone
type Stub struct {
	ID   string `json:"id,omitempty"`   // ID is the unique identifier of the Management Zone
	Name string `json:"name,omitempty"` // Name is the humand readable name of the Management Zone
}
