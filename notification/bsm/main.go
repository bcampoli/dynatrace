package main

import (
	"fmt"

	"github.com/dtcookie/dynatrace/apis/problems"
	"github.com/dtcookie/dynatrace/notification"
	"github.com/dtcookie/dynatrace/rest"
)

// PrintHandler TODO: documentation
type PrintHandler struct {
	notification.Handler
}

// ListenPort TODO: documentation
func (handler *PrintHandler) ListenPort() int {
	return 80
}

// Credentials TODO: documentation
func (handler *PrintHandler) Credentials() rest.Credentials {
	return rest.NewSaasCredentials("siz65484", "KN7jh2l6ROOxdtYJk3KX_")
}

// Handle TODO: documentation
func (handler *PrintHandler) Handle(problem *problems.Problem) error {
	fmt.Println(toXML(problem))
	return nil
}

func main() {
	notification.Launch(&PrintHandler{})
}
