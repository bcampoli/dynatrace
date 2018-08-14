package notification

import "github.com/dtcookie/dynatrace/apis/problems"

// Handler TODO: documentation
type Handler interface {
	Handle(problem *problems.Problem) error
}

// NewHandler TODO: documentation
func NewHandler(handler Handler) Handler {
	return &hub{next: handler}
}

type hub struct {
	Handler
	next Handler
}

func (hub *hub) Handle(problem *problems.Problem) error {
	return hub.next.Handle(problem)
}
