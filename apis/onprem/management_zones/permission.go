package managementzones

import (
	"bytes"
	"encoding/json"
)

// Permission TODO: documentation
type Permission struct {
	name string
}

// Permissions TODO: documentation
var Permissions = struct {
	DemoUser                 Permission
	LogViewer                Permission
	ManageSettings           Permission
	Viewer                   Permission
	ViewSensitiveRequestData Permission
}{
	Permission{"DEMO_USER"},
	Permission{"LOG_VIEWER"},
	Permission{"MANAGE_SETTINGS"},
	Permission{"VIEWER"},
	Permission{"VIEW_SENSITIVE_REQUEST_DATA"},
}

// MarshalJSON marshals the enum as a quoted json string
func (permission Permission) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(permission.name)
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

// UnmarshalJSON unmashals a quoted json string to the enum value
func (permission *Permission) UnmarshalJSON(b []byte) error {
	var j string
	if err := json.Unmarshal(b, &j); err != nil {
		return err
	}
	permission.name = j
	return nil
}
