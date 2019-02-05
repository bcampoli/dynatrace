package managementzones

// ZonePermission TODO: documentation
type ZonePermission struct {
	ManagementZoneID string       `json:"mzId,omitempty"`        // The ID of the required management zone
	Permissions      []Permission `json:"permissions,omitempty"` // The list of permissions for the required management zone
}
