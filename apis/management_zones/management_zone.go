package managementzones

// ManagementZone is the full configuration of a Management Zone
type ManagementZone struct {
	Name  string `json:"name"`  // the humand readable name of the Management Zone
	Rules []Rule `json:"rules"` // the rules a Management Zone is based on
}

// Rule TODO: documentation
type Rule struct {
	Type             string      `json:"type"`
	Enabled          bool        `json:"enabled"`
	PropagationTypes []string    `json:"propagationTypes"`
	Conditions       []Condition `json:"conditions"`
}

// Condition TODO: documentation
type Condition struct {
	Key            ConditionKey   `json:"key"`
	ComparisonInfo ComparisonInfo `json:"comparisonInfo"`
}

// ConditionKey TODO: documentation
type ConditionKey struct {
	Attribute string `json:"attribute"`
}

// ComparisonInfo TODO: documentation
type ComparisonInfo struct {
	Type          string `json:"type"`
	Operator      string `json:"operator"`
	Value         string `json:"value"`
	Negate        bool   `json:"negate"`
	CaseSensitive bool   `json:"caseSensitive"`
}
