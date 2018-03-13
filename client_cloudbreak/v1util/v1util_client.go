// Code generated by go-swagger; DO NOT EDIT.

package v1util

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new v1util API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for v1util API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CheckClientVersion checks the client version
*/
func (a *Client) CheckClientVersion(params *CheckClientVersionParams) (*CheckClientVersionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCheckClientVersionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "checkClientVersion",
		Method:             "GET",
		PathPattern:        "/v1/util/client/{version}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CheckClientVersionReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CheckClientVersionOK), nil

}

/*
CreateRDSDatabaseUtil creates a database for the service in the r d s if the connection could be created
*/
func (a *Client) CreateRDSDatabaseUtil(params *CreateRDSDatabaseUtilParams) (*CreateRDSDatabaseUtilOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateRDSDatabaseUtilParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createRDSDatabaseUtil",
		Method:             "POST",
		PathPattern:        "/v1/util/rds-database",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateRDSDatabaseUtilReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateRDSDatabaseUtilOK), nil

}

/*
TestAmbariDatabaseUtil tests a database connection parameters
*/
func (a *Client) TestAmbariDatabaseUtil(params *TestAmbariDatabaseUtilParams) (*TestAmbariDatabaseUtilOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTestAmbariDatabaseUtilParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "testAmbariDatabaseUtil",
		Method:             "POST",
		PathPattern:        "/v1/util/ambari-database",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &TestAmbariDatabaseUtilReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*TestAmbariDatabaseUtilOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
