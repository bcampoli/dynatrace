package notification

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/dtcookie/dynatrace/apis/problems"
)

func newListener(config *Config, handler Handler) *listener {
	return &listener{config: config, handler: handler}
}

// listener TODO: documentation
type listener struct {
	handler Handler
	config  *Config
}

func (listener *listener) listen() {
	http.HandleFunc("/", listener.handleHTTP)
	fmt.Println(fmt.Sprintf("Listening on port %d for incoming problem notifications.", listener.config.ListenPort))
	http.ListenAndServe(fmt.Sprintf(":%d", listener.config.ListenPort), nil)
}

func (listener *listener) handleHTTP(w http.ResponseWriter, request *http.Request) {
	var err error
	var body []byte

	if request.Method != http.MethodPost {
		fmt.Println(request.Method + " responding with " + http.StatusText(http.StatusMethodNotAllowed))
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	if request.ContentLength == 0 {
		fmt.Println("responding with " + http.StatusText(http.StatusBadRequest))
		http.Error(w, http.StatusText(http.StatusBadRequest)+": missing request body", http.StatusBadRequest)
		return
	}
	var contentType string
	contentType = request.Header.Get("content-type")
	if contentType != "application/json" {
		fmt.Println("responding with " + http.StatusText(http.StatusBadRequest) + ": expected content-type 'application/json'")
		http.Error(w, http.StatusText(http.StatusBadRequest)+": expected content-type 'application/json'", http.StatusBadRequest)
		return
	}
	if body, err = ioutil.ReadAll(request.Body); err != nil {
		fmt.Println(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError)+": "+err.Error(), http.StatusInternalServerError)
		return
	}

	var defNotification Default
	if err = json.Unmarshal(body, &defNotification); err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("received problem notification")
	fmt.Println("  ... " + toJSON(defNotification))
	if (defNotification.Title == "Dynatrace problem notification test run") || (defNotification.PID == "999999") {
		http.Error(w, http.StatusText(http.StatusNoContent), http.StatusNoContent)
		return
	}

	if defNotification.PID == "" {
		fmt.Println("  ... " + "no PID included. Cannot query for details, sorry.")
		http.Error(w, http.StatusText(http.StatusNoContent), http.StatusNoContent)
		return
	}
	fmt.Println("  ... querying for problem details")
	problemAPI := problems.NewAPI(listener.config.Credentials)
	var problem *problems.Problem
	if problem, err = problemAPI.Get(defNotification.PID); err != nil {
		fmt.Println(err.Error())
		return
	}
	// fmt.Println(toXML(problem))

	listener.handler.Handle(problem)

	http.Error(w, http.StatusText(http.StatusNoContent), http.StatusNoContent)
}

func toJSON(v interface{}) string {
	var err error
	var bytes []byte
	if bytes, err = json.Marshal(v); err != nil {
		return err.Error()
	}
	return string(bytes)
}

func toXML(v interface{}) string {
	var err error
	var bytes []byte
	if bytes, err = xml.MarshalIndent(v, "", "  "); err != nil {
		return err.Error()
	}
	return string(bytes)
}
