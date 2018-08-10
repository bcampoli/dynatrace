package problems

// Problem TODO: documentation
type Problem struct {

	// The ID of the problem.
	ID string `json:"id,omitempty" xml:"id,attr,omitempty"`
	// Timestamp of the problem start, in UTC milliseconds.
	StartTime int64 `json:"startTime,omitempty" xml:"startTime,attr,omitempty"`
	// Timestamp of the problem end, in UTC milliseconds. `-1` if the problem is still open.
	EndTime int64 `json:"endTime,omitempty" xml:"endTime,attr,omitempty"`
	// The name of the problem, displayed in the UI.
	DisplayName string `json:"displayName,omitempty" xml:"displayName,attr,omitempty"`
	// The impact level of the problem. It shows what is affected by the problem: infrastructure, service, or application.
	ImpactLevel impactLevel `json:"impactLevel,omitempty" xml:"impactLevel,attr,omitempty"`
	// The status of the problem.
	Status eventStatus `json:"status,omitempty" xml:"status,attr,omitempty"`
	// The severity of the problem.
	SeverityLevel string `json:"severityLevel,omitempty" xml:"severityLevel,attr,omitempty"`
	// Number of comments to the problem.
	CommentCount int32 `json:"commentCount,omitempty" xml:"commentCount,attr,omitempty"`
	// Tags of entities affected by the problem.
	TagsOfAffectedEntities []TagInfo `json:"tagsOfAffectedEntities,omitempty" xml:"TagsOfAffectedEntities>TagInfo"`
	// List of events related to the problem.
	RankedEvents []Event `json:"rankedEvents,omitempty" xml:"rankedEvents>Event"`
	// Provides impact information of the events in an aggregated form. For a more detailed impact analysis, look at 'rankedEvents'
	RankedImpacts []EventRestImpact `json:"rankedImpacts,omitempty" xml:"rankedImpacts>eventRestImpact"`
	// The number of affected entities per impact level.
	AffectedCounts *AffectedCounts `json:"affectedCounts,omitempty" xml:"affectedCounts"`
	// The number of entities which were affected, but recovered, per impact level.
	RecoveredCounts *RecoveredCounts `json:"recoveredCounts,omitempty" xml:"recoveredCounts"`
	// Whether Dynatrace has found at least one possible root cause for the problem.
	HasRootCause bool `json:"hasRootCause" xml:"hasRootCause"`
}
