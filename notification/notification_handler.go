package notification

import "github.com/dtcookie/dynatrace/apis/problems"
import "github.com/dtcookie/dynatrace/rest"

// Handler TODO: documentation
type Handler interface {
	ListenPort() int
	Credentials() rest.Credentials
	Handle(problem *problems.Problem) error
}