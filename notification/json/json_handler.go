package json

// Handler TODO: documentation
type Handler interface {
	Handle(jsonstr string) error
}
