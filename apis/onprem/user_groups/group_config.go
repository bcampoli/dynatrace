package usersgroups

// GroupConfig TODO: documentation
type GroupConfig struct {
	ID                  string              `json:"id,omitempty"`        // Group ID
	Name                string              `json:"name"`                // Group name
	LDAPGroupNames      []string            `json:"ldapGroupNames"`      // LDAP group names
	AccessRight         map[string][]string `json:"accessRight"`         // Access rights
	IsClusterAdminGroup bool                `json:"isClusterAdminGroup"` // If true, then the cluster has administrator rights
}
