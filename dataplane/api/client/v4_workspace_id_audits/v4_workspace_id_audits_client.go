// Code generated by go-swagger; DO NOT EDIT.

package v4_workspace_id_audits

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new v4 workspace id audits API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for v4 workspace id audits API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetAuditEventByWorkspace gets audit event in workspace by id

Audit event operations
*/
func (a *Client) GetAuditEventByWorkspace(params *GetAuditEventByWorkspaceParams) (*GetAuditEventByWorkspaceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAuditEventByWorkspaceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getAuditEventByWorkspace",
		Method:             "GET",
		PathPattern:        "/v4/{workspaceId}/audits/{auditId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetAuditEventByWorkspaceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetAuditEventByWorkspaceOK), nil

}

/*
GetAuditEventsInWorkspace lists audit events for the given workspace please use the API filter by resource crn because the resource type and id combination is deprecated

Audit event operations
*/
func (a *Client) GetAuditEventsInWorkspace(params *GetAuditEventsInWorkspaceParams) (*GetAuditEventsInWorkspaceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAuditEventsInWorkspaceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getAuditEventsInWorkspace",
		Method:             "GET",
		PathPattern:        "/v4/{workspaceId}/audits",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetAuditEventsInWorkspaceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetAuditEventsInWorkspaceOK), nil

}

/*
GetAuditEventsZipInWorkspace lists audit events for the given workspace in zip file please use the API filter by resource crn because the resource type and id combination is deprecated

Audit event operations
*/
func (a *Client) GetAuditEventsZipInWorkspace(params *GetAuditEventsZipInWorkspaceParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAuditEventsZipInWorkspaceParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getAuditEventsZipInWorkspace",
		Method:             "GET",
		PathPattern:        "/v4/{workspaceId}/audits/zip",
		ProducesMediaTypes: []string{"application/json", "application/octet-stream"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetAuditEventsZipInWorkspaceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
