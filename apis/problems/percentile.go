package problems

import (
	"encoding/xml"
	"fmt"
	"strconv"
)

type percentile struct {
	fmt.Stringer
	name string
}

type percentiles struct {
	Fiftieth  percentile
	Nineteeth percentile
}

// Percentiles TODO: documentation
var Percentiles = percentiles{
	Fiftieth: percentile{
		name: "50th",
	},
	Nineteeth: percentile{
		name: "90th",
	},
}

func (percentile *percentile) String() string {
	return percentile.name
}

func (percentile *percentile) MarshalJSON() ([]byte, error) {
	return []byte(strconv.Quote(percentile.name)), nil
}

func (percentile *percentile) UnmarshalJSON(data []byte) error {
	quoted, err := strconv.Unquote(string(data))
	percentile.name = quoted
	return err
}

func (percentile *percentile) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	return xml.Attr{Name: name, Value: percentile.name}, nil
}

func (percentile *percentile) UnmarshalXMLAttr(attr xml.Attr) error {
	percentile.name = attr.Value
	return nil
}
