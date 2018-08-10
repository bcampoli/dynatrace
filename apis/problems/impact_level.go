package problems

import (
	"encoding/xml"
	"fmt"
	"strconv"
)

type impactLevel struct {
	fmt.Stringer
	name string
}

type impactLevels struct {
	Infrastructure impactLevel
	Service        impactLevel
	Application    impactLevel
	Environment    impactLevel
}

// ImpactLevels TODO: documentation
var ImpactLevels = impactLevels{
	Infrastructure: impactLevel{
		name: "INFRASTRUCTURE",
	},
	Service: impactLevel{
		name: "SERVICE",
	},
	Application: impactLevel{
		name: "APPLICATION",
	},
	Environment: impactLevel{
		name: "ENVIRONMENT",
	},
}

func (level *impactLevel) String() string {
	return level.name
}

func (level *impactLevel) MarshalJSON() ([]byte, error) {
	return []byte(strconv.Quote(level.name)), nil
}

func (level *impactLevel) UnmarshalJSON(data []byte) error {
	quoted, err := strconv.Unquote(string(data))
	level.name = quoted
	return err
}

func (level *impactLevel) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	return xml.Attr{Name: name, Value: level.name}, nil
}

func (level *impactLevel) UnmarshalXMLAttr(attr xml.Attr) error {
	level.name = attr.Value
	return nil
}
