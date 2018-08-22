package main

import "encoding/xml"

// Event TODO: documentation
type Event struct {
	XMLName       xml.Name `xml:"Event"`
	Title         string   `json:"title,omitempty" xml:"title,omitempty"`
	Description   string   `json:"description,omitempty" xml:"description,omitempty"`
	PID           string   `json:"PID,omitempty" xml:"PID,omitempty"`
	Severity      string   `json:"severity,omitempty" xml:"severity,omitempty"`
	RelatedEntity string   `json:"relatedEntity,omitempty" xml:"relatedEntity,omitempty"`
}
