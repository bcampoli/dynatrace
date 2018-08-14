package main

import (
	"github.com/dtcookie/dynatrace/notification"
)

func main() {
	var config *notification.Config
	var handler SOAPHandler

	if config = parseConfig(&handler); config == nil {
		return
	}

	notification.Listen(config, &handler)
}
