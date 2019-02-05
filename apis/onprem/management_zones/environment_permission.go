package managementzones

// EnvironmentPermission TODO: documentation
type EnvironmentPermission struct {
	EnvironmentUUID string           `json:"environmentUuid,omitempty"` // Environment UUID
	ZonePermissions []ZonePermission `json:"mzPermissions,omitempty"`   // List of management zone models with permissions
}
