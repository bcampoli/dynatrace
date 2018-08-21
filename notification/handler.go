package notification

// Handler TODO: documentation
type Handler interface {
	Handle(event *ProblemEvent) error
}

// NewHandler TODO: documentation
func NewHandler(handler Handler) Handler {
	return &hub{next: handler}
}

type hub struct {
	Handler
	next Handler
}

func (hub *hub) Handle(event *ProblemEvent) error {
	return hub.next.Handle(event)
}
