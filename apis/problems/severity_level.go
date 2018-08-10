package problems

import (
	"encoding/xml"
	"fmt"
	"strconv"
)

// severityLevel TODO: documentation
type severityLevel struct {
	fmt.Stringer
	name string
}

type severityLevels struct {
	Availability          severityLevel
	Error                 severityLevel
	Performance           severityLevel
	ResourceContention    severityLevel
	CustomAlert           severityLevel
	MonitoringUnavailable severityLevel
}

// SeverityLevels TODO: documentation
var SeverityLevels = severityLevels{
	Availability: severityLevel{
		name: "AVAILABILITY",
	},
	Error: severityLevel{
		name: "ERROR",
	},
	Performance: severityLevel{
		name: "PERFORMANCE",
	},
	ResourceContention: severityLevel{
		name: "RESOURCE_CONTENTION",
	},
	CustomAlert: severityLevel{
		name: "CUSTOM_ALERT",
	},
	MonitoringUnavailable: severityLevel{
		name: "MONITORING_UNAVAILABLE",
	},
}

func (level *severityLevel) String() string {
	return level.name
}

func (level *severityLevel) MarshalJSON() ([]byte, error) {
	return []byte(strconv.Quote(level.name)), nil
}

func (level *severityLevel) UnmarshalJSON(data []byte) error {
	quoted, err := strconv.Unquote(string(data))
	level.name = quoted
	return err
}

func (level *severityLevel) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	return xml.Attr{Name: name, Value: level.name}, nil
}

func (level *severityLevel) UnmarshalXMLAttr(attr xml.Attr) error {
	level.name = attr.Value
	return nil
}
