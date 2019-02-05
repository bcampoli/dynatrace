package managementzones

// PermissionsForGroup TODO: documentation
type PermissionsForGroup struct {
	GroupID                string                  `json:"groupId,omitempty"`                     // Group ID
	EnvironmentPermissions []EnvironmentPermission `json:"mzPermissionsPerEnvironment,omitempty"` // List of management zone permissions per environment
}
