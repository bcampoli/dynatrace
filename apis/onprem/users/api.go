package users

import (
	"encoding/json"
	"errors"
	"fmt"

	resterrors "github.com/dtcookie/dynatrace/apis/errors"
	"github.com/dtcookie/dynatrace/rest"
)

// API is able to make REST API Calls to the Users API of an
// OnPremise Dynatrace Cluster
type API struct {
	client *rest.Client
}

// NewAPI creates a preconfigured API for accessing the Users API
// of an OnPremise Dynatrace Cluster
func NewAPI(config *rest.Config, credentials *rest.Credentials) *API {
	return &API{client: rest.NewClient(config, credentials)}
}

// GetUsers queries for the currently configured users
func (api *API) GetUsers() ([]UserConfig, error) {
	var err error
	var bytes []byte

	if bytes, err = api.client.GET("/api/v1.0/onpremise/users", 200); err != nil {
		return nil, resterrors.Resolve(bytes, err)
	}
	var response []UserConfig
	if err = json.Unmarshal(bytes, &response); err != nil {
		return nil, err
	}
	return response, nil
}

// Create TODO: documentation
func (api *API) Create(config *UserConfig) (*UserConfig, error) {
	var err error
	var bytes []byte
	// if bytes, err = api.client.POST("/api/v1.0/onpremise/users", config, 200); err != nil {
	// 	return nil, resterrors.Resolve(bytes, err)
	// }
	if bytes, err = api.client.NewPOST("/api/v1.0/onpremise/users", config).Expect(200).OnResponse(func(statusCode int) error {
		switch statusCode {
		case 200:
			return nil
		case 400:
			return errors.New("All values (ID, email, first name, last name) must be set")
		case 406:
			return errors.New("Unacceptable request")
		case 522:
			return errors.New("Couldn’t create user")
		case 523:
			return errors.New("User already exists")
		case 524:
			return errors.New("Email address already registered")
		default:
			return fmt.Errorf("Error Code %d", statusCode)
		}
	}).Send(); err != nil {
		return nil, err
	}
	var response UserConfig
	if err = json.Unmarshal(bytes, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// Update TODO: documentation
func (api *API) Update(config *UserConfig) (*UserConfig, error) {
	var err error
	var bytes []byte
	if bytes, err = api.client.NewPUT("/api/v1.0/onpremise/users", config).Expect(200).OnResponse(func(statusCode int) error {
		switch statusCode {
		case 200:
			return nil
		case 400:
			return errors.New("All values (ID, email, first name, last name) must be set")
		case 406:
			return errors.New("Unacceptable request")
		case 524:
			return errors.New("Email address already registered")
		default:
			return fmt.Errorf("Error Code %d", statusCode)
		}
	}).Send(); err != nil {
		return nil, err
	}
	var response UserConfig
	if err = json.Unmarshal(bytes, &response); err != nil {
		return nil, err
	}
	return &response, nil
}
