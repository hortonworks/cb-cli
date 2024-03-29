// Code generated by go-swagger; DO NOT EDIT.

package sdxcost

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new sdxcost API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for sdxcost API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
ListSdxCost costs of s d x clusters
*/
func (a *Client) ListSdxCost(params *ListSdxCostParams) (*ListSdxCostOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListSdxCostParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listSdxCost",
		Method:             "PUT",
		PathPattern:        "/sdx/cost",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListSdxCostReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListSdxCostOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
