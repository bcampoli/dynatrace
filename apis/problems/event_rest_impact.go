package problems

// EventRestImpact TODO: documentation
type EventRestImpact struct {
	// The ID of the affected Dynatrace entity.
	// You can find the ID in the URL of the corresponding Dynatrace entity page, for example, `HOST-007`.
	EntityID string `xml:"EntityID,attr,omitempty"`
	// The name of the affected Dynatrace entity.
	EntityName string `xml:"EntityName,attr,omitempty"`
	// The severity of the event.
	SeverityLevel severityLevel `xml:"severityLevel,attr,omitempty"`
	// The impact level of the event.
	ImpactLevel impactLevel `xml:"impactLevel,attr,omitempty"`
	// The type of the event.
	EventType eventType `xml:"eventType,attr,omitempty"`
}
