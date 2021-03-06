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
GetOperationProgressByResourceCrn gets flow operation progress details for resource by resource crn

Environment consists of a credential and various other resources and enables users to quickly create clusters in given regions in a given cloud provider.
*/
func (a *Client) GetOperationProgressByResourceCrn(params *GetOperationProgressByResourceCrnParams) (*GetOperationProgressByResourceCrnOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOperationProgressByResourceCrnParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getOperationProgressByResourceCrn",
		Method:             "GET",
		PathPattern:        "/v1/operation/resource/crn/{resourceCrn}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetOperationProgressByResourceCrnReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetOperationProgressByResourceCrnOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
