// Code generated by go-swagger; DO NOT EDIT.

package v3_organization_id_filesystems

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new v3 organization id filesystems API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for v3 organization id filesystems API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetFileSystemParametersV3 returns filesystem parameters
*/
func (a *Client) GetFileSystemParametersV3(params *GetFileSystemParametersV3Params) (*GetFileSystemParametersV3OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFileSystemParametersV3Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getFileSystemParametersV3",
		Method:             "POST",
		PathPattern:        "/v3/{organizationId}/filesystems/parameters",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetFileSystemParametersV3Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetFileSystemParametersV3OK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
