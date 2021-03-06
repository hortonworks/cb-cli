// Code generated by go-swagger; DO NOT EDIT.

package v4dbconfig

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new v4dbconfig API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for v4dbconfig API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetDbConfig gets database config
*/
func (a *Client) GetDbConfig(params *GetDbConfigParams) (*GetDbConfigOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDbConfigParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getDbConfig",
		Method:             "GET",
		PathPattern:        "/v4/dbconfig/connectionparams",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetDbConfigReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetDbConfigOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
