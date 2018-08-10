package problems

import (
	"encoding/json"

	"github.com/dtcookie/dynatrace/rest"
)

// API TODO: documentation
type API struct {
	client *rest.Client
}

// NewAPI TODO: documentation
func NewAPI(credentials rest.Credentials) *API {
	return &API{
		client: rest.NewClient(credentials),
	}
}

// Get TODO: documentation
func (api *API) Get(ID string) (*Problem, error) {
	var err error
	var bytes []byte
	var problemResult problemResult
	if bytes, err = api.client.Get("/api/v1/problem/details/" + ID); err != nil {
		return nil, err
	}
	// fmt.Println(string(bytes))
	if err = json.Unmarshal(bytes, &problemResult); err != nil {
		return nil, err
	}
	return &problemResult.Result, nil
}

