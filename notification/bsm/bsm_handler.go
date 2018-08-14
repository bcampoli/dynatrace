package main

import (
	"github.com/dtcookie/dynatrace/http"
	"github.com/dtcookie/dynatrace/notification/xml"
)

// SOAPHandler TODO: documentation
type BSMhandler struct {
	xml.Handler
	Target string
	client *http.Client
}

// Handle TODO: documentation
func (handler *BSMhandler) Handle(xml string) error {
	return handler.client.Post(handler.Target, []byte(xml))
}
