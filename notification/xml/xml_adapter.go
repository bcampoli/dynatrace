package xml

import (
	"encoding/xml"

	"github.com/dtcookie/dynatrace/apis/problems"
	"github.com/dtcookie/dynatrace/notification"
)

type xmlAdapter struct {
	notification.Handler
	handler Handler
}

// NewXMLAdapter TODO: documentation
func NewXMLAdapter(handler Handler) notification.Handler {
	return &xmlAdapter{handler: handler}
}

// Handle TODO: documentation
func (adapter *xmlAdapter) Handle(problem *problems.Problem) error {
	var err error
	var xml string
	if xml, err = toXML(problem); err != nil {
		return err
	}
	return adapter.handler.Handle(xml)
	// return handler.client.Post(handler.Target, []byte(toXML(problem)))
}

func toXML(v interface{}) (string, error) {
	var err error
	var bytes []byte
	if bytes, err = xml.MarshalIndent(v, "", "  "); err != nil {
		return "", err
	}
	return string(bytes), nil
}
