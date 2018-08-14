package xml

// Handler TODO: documentation
type Handler interface {
	Handle(xml string) error
}
