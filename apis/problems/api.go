package problems

import (
	"encoding/json"
	"errors"

	resterrors "github.com/dtcookie/dynatrace/apis/errors"
	"github.com/dtcookie/dynatrace/rest"
)

// API TODO: documentation
type API struct {
	client *rest.Client
}

// NewAPI TODO: documentation
func NewAPI(config *rest.Config, credentials *rest.Credentials) *API {
	return &API{
		client: rest.NewClient(config, credentials),
	}
}

// Get TODO: documentation
func (api *API) Get(ID string) (*Problem, error) {
	var err error
	var bytes []byte
	var problemResult problemResult
	var errorEnvelope resterrors.ErrorEnvelope

	if bytes, err = api.client.GET("/api/v1/problem/details/"+ID, 200); err != nil {
		if bytes != nil {
			var innerError error
			if innerError = json.Unmarshal(bytes, &errorEnvelope); innerError == nil {
				if errorEnvelope.Error.Message != "" {
					return nil, errors.New(errorEnvelope.Error.Message)
				} else {
					return nil, err
				}
			}
		}
		return nil, err
	}
	// fmt.Println(string(bytes))
	if err = json.Unmarshal(bytes, &problemResult); err != nil {
		return nil, err
	}
	return &problemResult.Result, nil
}
