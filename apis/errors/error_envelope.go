package errors

import (
	"encoding/json"
	"errors"
)

// ErrorEnvelope is potentially the JSON response code
// when a REST API call fails
type ErrorEnvelope struct {
	Error RESTError `json:"error"` // the actual error object
}

// RESTError is potentially the JSON response code
// when a REST API call fails
type RESTError struct {
	Code    int    `json:"code"`    // error code
	Message string `json:"message"` // error message
}

// Resolve checks whether the given bytes are a JSON representation
// of a Dynatrace API REST Error or Error Envelope
// If that's the case a new error will be generated as a replacement
// for the given err (which is likely a low level HTTP error)
func Resolve(bytes []byte, err error) error {
	if bytes == nil {
		return err
	}
	var envelope ErrorEnvelope
	if innerError := json.Unmarshal(bytes, &envelope); innerError == nil {
		if envelope.Error.Message != "" {
			return errors.New(envelope.Error.Message)
		}
		return err
	}
	var resterror RESTError
	if innerError := json.Unmarshal(bytes, &resterror); innerError == nil {
		if resterror.Message != "" {
			return errors.New(resterror.Message)
		}
		return err
	}
	return err
}
