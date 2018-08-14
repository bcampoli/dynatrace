package main

import (
	"fmt"
	"strings"

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
		if !strings.HasPrefix(err.Error(), "flag provided but not defined") {
			fmt.Println(err.Error())
		}
		fmt.Println()
		fmt.Println("USAGE: bsm [-environment <environment-id>] [-api-token <api-token>] [-cluster <cluster-url>] [-listen <listen-port>] [-config <config-json-file>")
		fmt.Println("  Hint: you can also define the environment variables DT_NOTIFICATION_ENVIRONMENT, DT_NOTIFICATION_API_TOKEN, DT_NOTIFICATION_CLUSTER and DT_NOTIFICATION_LISTEN_PORT")
		fmt.Println("  Hint: you can also specify the -config flag referring to a JSON file containing the parameters")
		return
	}
	notification.Listen(config, &PrintHandler{})
}
