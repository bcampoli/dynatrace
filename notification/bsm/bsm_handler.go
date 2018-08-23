package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"strings"

	"github.com/dtcookie/dynatrace/http"
	"github.com/dtcookie/dynatrace/notification"
)

// BSMhandler TODO: documentation
type BSMhandler struct {
	notification.Handler
	Target string
	client *http.Client
}

// Handle TODO: documentation
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

	var tags string
	tags = strings.Trim(event.Notification.Tags, " ")
	var bcode string
	barcodes := strings.Split(tags, ",")
	for _, barcode := range barcodes {
		fmt.Println("barcode: " + barcode)
		barcode = strings.Trim(barcode, " ")
		fmt.Println("barcode (trimmed): " + barcode)
		if strings.HasPrefix(barcode, "APP=") {
			idxOpenBrace := strings.LastIndex(barcode, "(")
			if (idxOpenBrace > 0) && (idxOpenBrace < len(barcode)) {
				barcode = barcode[idxOpenBrace+1 : len(barcode)]
				if strings.HasSuffix(barcode, ")") {
					barcode = barcode[0 : len(barcode)-1]
					bcode = barcode
					break
				}
			}
		}
	}

	if len(bcode) != 3 {
		fmt.Println("No barcode found in " + event.Notification.Tags)
		return nil
	}

	bsmEvent := Event{
		Title:         event.Notification.Title,
		Description:   "For detailed information visit: " + event.Notification.URL,
		PID:           event.Notification.PID,
		Severity:      event.Notification.State,
		RelatedEntity: bcode,
	}

	xmlStr, err := toXML(&bsmEvent)
	fmt.Println(xmlStr)

	if err != nil {
		fmt.Println("Sending to " + handler.Target)
		return handler.client.Post(handler.Target, []byte(xmlStr))
	}

	return err
}

func toJSON(v interface{}) (string, error) {
	var err error
	var bytes []byte
	if bytes, err = json.MarshalIndent(v, "", "  "); err != nil {
		return "", err
	}
	return string(bytes), nil
}

func toXML(v interface{}) (string, error) {
	var err error
	var bytes []byte
	if bytes, err = xml.MarshalIndent(v, "", "  "); err != nil {
		return "", err
	}
	return string(bytes), nil
}
