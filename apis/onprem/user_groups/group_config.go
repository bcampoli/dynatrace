package usersgroups

// GroupConfig TODO: documentation
type GroupConfig struct {
	ID                  string              `json:"id,omitempty"`                  // Group ID
	Name                string              `name:"name,omitempty"`                // Group name
	LDAPGroupNames      []string            `json:"ldapGroupNames,omitempty"`      // LDAP group names
	AccessRight         map[string][]string `json:"accessRight,omitempty"`         // Access rights
	IsClusterAdminGroup bool                `json:"isClusterAdminGroup,omitempty"` // If true, then the cluster has administrator rights
}
