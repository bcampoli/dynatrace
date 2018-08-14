package main

import (
	"fmt"

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
	if false {
		return handler.client.Post(handler.Target, []byte(xml))
	}
	fmt.Println(xml)
	return nil
}
