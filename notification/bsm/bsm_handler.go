package main

import (
	"fmt"

	"github.com/dtcookie/dynatrace/http"
	"github.com/dtcookie/dynatrace/notification/json"
)

// BSMhandler TODO: documentation
type BSMhandler struct {
	json.Handler
	Target string
	client *http.Client
}

// Handle TODO: documentation
func (handler *BSMhandler) Handle(jsonstr string) error {
	if false {
		return handler.client.Post(handler.Target, []byte(jsonstr))
	}
	fmt.Println(jsonstr)
	return nil
}
