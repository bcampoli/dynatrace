package notification

// Default TODO: documentation
type Default struct {
	PID                string `json:"PID,omitempty"`
	ProblemID          string `json:"ProblemID,omitempty"`
	State              string `json:"State,omitempty"`
	Title              string `json:"ProblemTitle,omitempty"`
	URL                string `json:"ProblemURL,omitempty"`
	Impact             string `json:"ProblemImpact,omitempty"`
	ImpactedEntity     string `json:"ImpactedEntity,omitempty"`
	Severity           string `json:"ProblemSeverity,omitempty"`
	Tags               string `json:"Tags,omitempty"`
	ProblemDetailsText string `json:"ProblemDetailsText,omitempty"`
}
