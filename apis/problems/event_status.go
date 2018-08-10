package problems

import (
	"encoding/xml"
	"fmt"
	"strconv"
)

type eventStatus struct {
	fmt.Stringer
	name string
}

type eventStati struct {
	Open   eventStatus
	Closed eventStatus
	all    []eventStatus
}

// EventStatus TODO: documentation
var EventStatus = eventStati{
	Open: eventStatus{
		name: "OPEN",
	},
	Closed: eventStatus{
		name: "CLOSED",
	},
}

func (status *eventStatus) String() string {
	return status.name
}

func (status *eventStatus) MarshalJSON() ([]byte, error) {
	return []byte(strconv.Quote(status.name)), nil
}

func (status *eventStatus) UnmarshalJSON(data []byte) error {
	quoted, err := strconv.Unquote(string(data))
	status.name = quoted
	return err
}

func (status *eventStatus) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	return xml.Attr{Name: name, Value: status.name}, nil
}

func (status *eventStatus) UnmarshalXMLAttr(attr xml.Attr) error {
	status.name = attr.Value
	return nil
}
