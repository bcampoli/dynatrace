package json

import (
	encjson "encoding/json"

	"github.com/dtcookie/dynatrace/notification"
)

type jsonAdapter struct {
	notification.Handler
	handler Handler
}

// NewJSONAdapter TODO: documentation
func NewJSONAdapter(handler Handler) notification.Handler {
	return &jsonAdapter{handler: handler}
}

// Handle TODO: documentation
func (adapter *jsonAdapter) Handle(event *notification.ProblemEvent) error {
	var err error
	var jsonstr string
	if jsonstr, err = toJSON(event); err != nil {
		return err
	}
	return adapter.handler.Handle(jsonstr)
}

func toJSON(v interface{}) (string, error) {
	var err error
	var bytes []byte
	if bytes, err = encjson.MarshalIndent(v, "", "  "); err != nil {
		return "", err
	}
	return string(bytes), nil
}
