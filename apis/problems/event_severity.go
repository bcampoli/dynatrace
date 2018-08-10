package problems

// EventSeverity TODO: documentation
type EventSeverity struct {
	Context eventSeverityContext `json:"context,omitempty" xml:"context,attr,omitempty"`
	Value   float64              `json:"value,omitempty" xml:"value,attr,omitempty"`
	Unit    string               `json:"unit,omitempty" xml:"unit,attr,omitempty"`
}
