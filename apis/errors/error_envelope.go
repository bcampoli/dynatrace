package errors

// ErrorEnvelope TODO documentation
type ErrorEnvelope struct {
	// Error TODO documentation
	Error RESTError `json:"error,omitempty"`
}

// RESTError TODO documentation
type RESTError struct {
	// Code TODO documentation
	Code int `json:"code,omitempty"`
	// Message TODO documentation
	Message string `json:"message,omitempty"`
}
