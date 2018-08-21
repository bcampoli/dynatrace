package main

import (
	"github.com/dtcookie/dynatrace/notification"
	"github.com/dtcookie/dynatrace/notification/json"
)

func main() {
	var config *notification.Config
	var handler BSMhandler

	if config = parseConfig(&handler); config == nil {
		return
	}

	notification.Listen(config, json.NewJSONAdapter(&handler))
}
