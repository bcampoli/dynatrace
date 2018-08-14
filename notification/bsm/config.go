package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/dtcookie/dynatrace/notification"
)

func parseConfig(handler *BSMhandler) *notification.Config {
	var err error
	var config *notification.Config

	flagSet := flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	flagSet.StringVar(&handler.Target, "target", "", "")
	if config, err = notification.ParseConfig(flagSet); err != nil {
		if !strings.HasPrefix(err.Error(), "flag provided but not defined") {
			fmt.Println(err.Error())
			usage()
		}
		return nil
	}

	if handler.Target == "" {
		fmt.Println("no target specified")
		usage()
		return nil
	}

	return config
}

func usage() {
	fmt.Println()
	fmt.Println("USAGE: bsm [-environment <environment-id>] [-api-token <api-token>] [-cluster <cluster-url>] [-listen <listen-port>] [-config <config-json-file>")
	fmt.Println("  Hint: you can also define the environment variables DT_NOTIFICATION_ENVIRONMENT, DT_NOTIFICATION_API_TOKEN, DT_NOTIFICATION_CLUSTER and DT_NOTIFICATION_LISTEN_PORT")
	fmt.Println("  Hint: you can also specify the -config flag referring to a JSON file containing the parameters")
}
