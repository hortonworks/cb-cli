// Code generated by go-swagger; DO NOT EDIT.

package v1freeipacost

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new v1freeipacost API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for v1freeipacost API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
ListFreeIpaCostV1 shows free IP a cost

FreeIPA is an integrated Identity and Authentication solution that can be used for any of CM, CDP services.
*/
func (a *Client) ListFreeIpaCostV1(params *ListFreeIpaCostV1Params, authInfo runtime.ClientAuthInfoWriter) (*ListFreeIpaCostV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListFreeIpaCostV1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listFreeIpaCostV1",
		Method:             "PUT",
		PathPattern:        "/v1/freeipa/cost",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListFreeIpaCostV1Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListFreeIpaCostV1OK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
