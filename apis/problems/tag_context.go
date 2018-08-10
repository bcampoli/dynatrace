package problems

import (
	"encoding/xml"
	"fmt"
	"strconv"
)

type tagContext struct {
	fmt.Stringer
	name string
}

type tagContexts struct {
	Contextless  tagContext
	AWS          tagContext
	Environment  tagContext
	CloudFoundry tagContext
	Kubernetes   tagContext
	GoogleCloud  tagContext
}

// TagContexts TODO: documentation
var TagContexts = tagContexts{
	Contextless: tagContext{
		name: "CONTEXTLESS",
	},
	AWS: tagContext{
		name: "AWS",
	},
	Environment: tagContext{
		name: "ENVIRONMENT",
	},
	CloudFoundry: tagContext{
		name: "CLOUD_FOUNDRY",
	},
	Kubernetes: tagContext{
		name: "KUBERNETES",
	},
	GoogleCloud: tagContext{
		name: "GOOGLE_CLOUD",
	},
}

func (context *tagContext) String() string {
	return context.name
}

func (context *tagContext) MarshalJSON() ([]byte, error) {
	return []byte(strconv.Quote(context.name)), nil
}

func (context *tagContext) UnmarshalJSON(data []byte) error {
	quoted, err := strconv.Unquote(string(data))
	context.name = quoted
	return err
}

func (context *tagContext) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	return xml.Attr{Name: name, Value: context.name}, nil
}

func (context *tagContext) UnmarshalXMLAttr(attr xml.Attr) error {
	context.name = attr.Value
	return nil
}
