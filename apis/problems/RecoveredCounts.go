package problems

// RecoveredCounts TODO: documentation
type RecoveredCounts struct {
	Application    int `json:"APPLICATION,omitempty" xml:"APPLICATION,attr"`
	Environment    int `json:"ENVIRONMENT,omitempty" xml:"ENVIRONMENT,attr"`
	Infrastructure int `json:"INFRASTRUCTURE,omitempty" xml:"INFRASTRUCTURE,attr"`
	Service        int `json:"SERVICE,omitempty" xml:"SERVICE,attr"`
}
