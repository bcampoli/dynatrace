package main

import (
	"encoding/json"
	"fmt"

	"github.com/dtcookie/dynatrace/http"
	"github.com/dtcookie/dynatrace/notification"
)

// BSMhandler TODO: documentation
type BSMhandler struct {
	notification.Handler
	Target string
	client *http.Client
}

func (handler *BSMhandler) Handle(event *notification.ProblemEvent) error {
	var err error
	var jsonstr string
	if jsonstr, err = toJSON(event); err != nil {
		return err
	}
	if false {
		return handler.client.Post(handler.Target, []byte(jsonstr))
	}
	fmt.Println(jsonstr)
	fmt.Println()

	xmlStr := ""

	xmlStr = xmlStr + "<xml>"
	xmlStr = xmlStr + "<BSMCEvent>"
	xmlStr = xmlStr + "  <Title>" + event.Notification.Title + "</Title>"
	xmlStr = xmlStr + "  <Description>" + event.Notification.URL + "</Description>"
	xmlStr = xmlStr + "  <PID>" + event.Notification.PID + "</PID>"
	xmlStr = xmlStr + "  <Severity>" + "Critical" + "</Severity>"
	xmlStr = xmlStr + "  <RelatedCI>" + "UNKNOWN" + "</RelatedCI>"
	// xmlStr = xmlStr + "  <timeStamp>12/13/17 08:59:57 AM</timeStamp>"
	// xmlStr = xmlStr + "  <impact>neglectable</impact>"
	// xmlStr = xmlStr + "  <category>DT Managed</category>"
	// xmlStr = xmlStr + "  <relatedEntity>dcobel</relatedEntity>"
	xmlStr = xmlStr + "</BSMCEvent>"
	xmlStr = xmlStr + "</xml>"

	fmt.Println(xmlStr)

	return nil
}

func toJSON(v interface{}) (string, error) {
	var err error
	var bytes []byte
	if bytes, err = json.MarshalIndent(v, "", "  "); err != nil {
		return "", err
	}
	return string(bytes), nil
}
