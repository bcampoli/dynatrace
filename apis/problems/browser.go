package problems

import (
	"encoding/xml"
	"fmt"
	"strconv"
)

type browser struct {
	fmt.Stringer
	name string
}

type browsers struct {
	Infrastructure browser
	Service        browser
	Application    browser
	Environment    browser
}

// Browsers TODO: documentation
var Browsers = browsers{
	Infrastructure: browser{
		name: "INFRASTRUCTURE",
	},
	Service: browser{
		name: "SERVICE",
	},
	Application: browser{
		name: "APPLICATION",
	},
	Environment: browser{
		name: "ENVIRONMENT",
	},
}

func (browser *browser) String() string {
	return browser.name
}

func (browser *browser) MarshalJSON() ([]byte, error) {
	return []byte(strconv.Quote(browser.name)), nil
}

func (browser *browser) UnmarshalJSON(data []byte) error {
	quoted, err := strconv.Unquote(string(data))
	browser.name = quoted
	return err
}

func (browser *browser) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	return xml.Attr{Name: name, Value: browser.name}, nil
}

func (browser *browser) UnmarshalXMLAttr(attr xml.Attr) error {
	browser.name = attr.Value
	return nil
}
