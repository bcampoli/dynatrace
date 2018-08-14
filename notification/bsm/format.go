package main

import (
	"encoding/json"
	"encoding/xml"
)

func toJSON(v interface{}) string {
	var err error
	var bytes []byte
	if bytes, err = json.Marshal(v); err != nil {
		return err.Error()
	}
	return string(bytes)
}

func toXML(v interface{}) string {
	var err error
	var bytes []byte
	if bytes, err = xml.MarshalIndent(v, "", "  "); err != nil {
		return err.Error()
	}
	return string(bytes)
}
