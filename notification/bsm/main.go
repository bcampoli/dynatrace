package main

import (
	"fmt"

	"github.com/dtcookie/dynatrace/apis/problems"
	"github.com/dtcookie/dynatrace/notification"
)

// PrintHandler TODO: documentation
type PrintHandler struct {
	notification.Handler
}

// Handle TODO: documentation
func (handler *PrintHandler) Handle(problem *problems.Problem) error {
	fmt.Println(toXML(problem))
	return nil
}

func main() {
	var err error
	var config *notification.Config

	if config, err = notification.ParseConfig(); err != nil {
		fmt.Println(err.Error())
		fmt.Println("USAGE: bsm [-environment <environment-id>] [-api-token <api-token>] [-cluster <cluster-url>] [-listen <listen-port>]")
		return
	}
	notification.Listen(config, &PrintHandler{})
}
