package main

import (
	"github.com/dtcookie/dynatrace/apis/problems"
	"github.com/dtcookie/dynatrace/notification"
)

// SOAPHandler TODO: documentation
type SOAPHandler struct {
	notification.Handler
	Target string
	client *HTTPClient
}

// Handle TODO: documentation
func (handler *SOAPHandler) Handle(problem *problems.Problem) error {
	return handler.client.Post(handler.Target, []byte(toXML(problem)))
}
