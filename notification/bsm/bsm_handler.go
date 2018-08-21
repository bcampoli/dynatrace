package main

import (
	"encoding/json"
	"fmt"

	"github.com/dtcookie/dynatrace/http"
	"github.com/dtcookie/dynatrace/notification"
)

// BSMhandler TODO: documentation
type BSMhandler struct {
	notification.Handler
	Target string
	client *http.Client
}

func (handler *BSMhandler) Handle(event *notification.ProblemEvent) error {
	var err error
	var jsonstr string
	if jsonstr, err = toJSON(event); err != nil {
		return err
	}
	if false {
		return handler.client.Post(handler.Target, []byte(jsonstr))
	}
	fmt.Println(jsonstr)
	return nil
}

func toJSON(v interface{}) (string, error) {
	var err error
	var bytes []byte
	if bytes, err = json.MarshalIndent(v, "", "  "); err != nil {
		return "", err
	}
	return string(bytes), nil
}
