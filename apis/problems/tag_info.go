package problems

// TagInfo TODO: documentation
type TagInfo struct {
	Key     string     `json:"key,omitempty" xml:"key,attr"`
	Context tagContext `json:"context,omitempty" xml:"context"`
	Value   string     `json:"value,omitempty" xml:"value,attr"`
}
