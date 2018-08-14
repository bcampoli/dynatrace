package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/dtcookie/dynatrace/apis/problems"
	"github.com/dtcookie/dynatrace/notification"
)

// PrintHandler TODO: documentation
type PrintHandler struct {
	notification.Handler
}

// Post TODO: documentation
func (handler *PrintHandler) Post(URL string, problem *problems.Problem) error {
	var err error
	var response *http.Response
	var request *http.Request
	if request, err = http.NewRequest(http.MethodPost, URL, bytes.NewBuffer([]byte(toXML(problem)))); err != nil {
		return err
	}
	request.Header.Set("Content-Type", "application/xml")

	httpClient := &http.Client{}
	if response, err = httpClient.Do(request); err != nil {
		return err
	}
	defer response.Body.Close()
	ioutil.ReadAll(response.Body)
	return nil
}

// Handle TODO: documentation
func (handler *PrintHandler) Handle(problem *problems.Problem) error {
	return handler.Post("http://postb.in/R06fBRsq", problem)
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
	notification.Listen(config, new(PrintHandler))
}
