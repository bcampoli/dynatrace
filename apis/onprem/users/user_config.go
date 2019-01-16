package users

// UserConfig holds the properties of a configured user
// within an OnPremise installation of Dynatrace
type UserConfig struct {
	ID                string   `json:"id,omitempty"`                // User ID
	EMail             string   `json:"email,omitempty"`             // User's email address
	FirstName         string   `json:"firstName,omitempty"`         // User's first name
	LastName          string   `json:"lastName,omitempty"`          // User's last name
	PasswordClearText string   `json:"passwordClearText,omitempty"` // User's password in a clear text; used only to set initial password
	Groups            []string `json:"groups,omitempty"`            // List of user's user groups
}
