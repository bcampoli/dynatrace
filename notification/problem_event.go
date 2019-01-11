package notification

import "github.com/dtcookie/dynatrace/apis/problems"

// ProblemEvent TODO: documentation
type ProblemEvent struct {
	URI          string            `json:"-"`
	Notification *Default          `json:"notification,omitempty"`
	Problem      *problems.Problem `json:"details,omitempty"`
}
