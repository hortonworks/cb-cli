// Code generated by go-swagger; DO NOT EDIT.

package v1structured_events

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new v1structured events API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for v1structured events API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetAuditEventsZipForResource lists audit events for the given resource in zip file

Audit event operations
*/
func (a *Client) GetAuditEventsZipForResource(params *GetAuditEventsZipForResourceParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAuditEventsZipForResourceParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getAuditEventsZipForResource",
		Method:             "GET",
		PathPattern:        "/v1/structured_events/zip",
		ProducesMediaTypes: []string{"application/json", "application/octet-stream"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetAuditEventsZipForResourceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
GetCDPAuditEventsForResource lists audit events for the given resource

Audit event operations
*/
func (a *Client) GetCDPAuditEventsForResource(params *GetCDPAuditEventsForResourceParams) (*GetCDPAuditEventsForResourceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCDPAuditEventsForResourceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getCDPAuditEventsForResource",
		Method:             "GET",
		PathPattern:        "/v1/structured_events",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetCDPAuditEventsForResourceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetCDPAuditEventsForResourceOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
