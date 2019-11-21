// Code generated by go-swagger; DO NOT EDIT.

package v1operation

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new v1operation API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for v1operation API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetOperationStatusV1 gets the status of an operation

Operation management endpoint
*/
func (a *Client) GetOperationStatusV1(params *GetOperationStatusV1Params) (*GetOperationStatusV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOperationStatusV1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getOperationStatusV1",
		Method:             "GET",
		PathPattern:        "/v1/operation",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetOperationStatusV1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetOperationStatusV1OK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
