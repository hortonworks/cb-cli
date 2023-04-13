// Code generated by go-swagger; DO NOT EDIT.

package v1distroxcarbon_dioxide

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new v1distroxcarbon dioxide API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for v1distroxcarbon dioxide API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
ListDistroXCO2V1 gets c o2 cost calculation for distro clusters

CO2 cost of clusters
*/
func (a *Client) ListDistroXCO2V1(params *ListDistroXCO2V1Params) (*ListDistroXCO2V1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListDistroXCO2V1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listDistroXCO2V1",
		Method:             "PUT",
		PathPattern:        "/v1/distrox/carbon_dioxide",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListDistroXCO2V1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListDistroXCO2V1OK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
