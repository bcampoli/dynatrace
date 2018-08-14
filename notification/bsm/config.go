package main

import "github.com/dtcookie/dynatrace/rest"

// Config TODO: documentation
type Config struct {
	ListenPort  int
	Credentials rest.Credentials
}
