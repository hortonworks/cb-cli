// Code generated by go-swagger; DO NOT EDIT.

package v4_workspace_id_mpacks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new v4 workspace id mpacks API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for v4 workspace id mpacks API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CreateManagementPackInWorkspace creates management pack in workspace

An Apache Ambari Management Pack (Mpack) can bundle multiple service definitions, stack definitions, stack add-on service definitions, view definitions services so that releasing these artifacts don’t enforce an Apache Ambari release.
*/
func (a *Client) CreateManagementPackInWorkspace(params *CreateManagementPackInWorkspaceParams) (*CreateManagementPackInWorkspaceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateManagementPackInWorkspaceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createManagementPackInWorkspace",
		Method:             "POST",
		PathPattern:        "/v4/{workspaceId}/mpacks",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateManagementPackInWorkspaceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateManagementPackInWorkspaceOK), nil

}

/*
DeleteManagementPackInWorkspace deletes management pack by name in workspace

An Apache Ambari Management Pack (Mpack) can bundle multiple service definitions, stack definitions, stack add-on service definitions, view definitions services so that releasing these artifacts don’t enforce an Apache Ambari release.
*/
func (a *Client) DeleteManagementPackInWorkspace(params *DeleteManagementPackInWorkspaceParams) (*DeleteManagementPackInWorkspaceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteManagementPackInWorkspaceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteManagementPackInWorkspace",
		Method:             "DELETE",
		PathPattern:        "/v4/{workspaceId}/mpacks/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteManagementPackInWorkspaceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteManagementPackInWorkspaceOK), nil

}

/*
DeleteManagementPacksInWorkspace deletes multiple management packs by name in workspace

An Apache Ambari Management Pack (Mpack) can bundle multiple service definitions, stack definitions, stack add-on service definitions, view definitions services so that releasing these artifacts don’t enforce an Apache Ambari release.
*/
func (a *Client) DeleteManagementPacksInWorkspace(params *DeleteManagementPacksInWorkspaceParams) (*DeleteManagementPacksInWorkspaceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteManagementPacksInWorkspaceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteManagementPacksInWorkspace",
		Method:             "DELETE",
		PathPattern:        "/v4/{workspaceId}/mpacks",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteManagementPacksInWorkspaceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteManagementPacksInWorkspaceOK), nil

}

/*
GetManagementPackInWorkspace gets management pack by name in workspace

An Apache Ambari Management Pack (Mpack) can bundle multiple service definitions, stack definitions, stack add-on service definitions, view definitions services so that releasing these artifacts don’t enforce an Apache Ambari release.
*/
func (a *Client) GetManagementPackInWorkspace(params *GetManagementPackInWorkspaceParams) (*GetManagementPackInWorkspaceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetManagementPackInWorkspaceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getManagementPackInWorkspace",
		Method:             "GET",
		PathPattern:        "/v4/{workspaceId}/mpacks/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetManagementPackInWorkspaceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetManagementPackInWorkspaceOK), nil

}

/*
ListManagementPacksByWorkspace lists management packs for the given workspace

An Apache Ambari Management Pack (Mpack) can bundle multiple service definitions, stack definitions, stack add-on service definitions, view definitions services so that releasing these artifacts don’t enforce an Apache Ambari release.
*/
func (a *Client) ListManagementPacksByWorkspace(params *ListManagementPacksByWorkspaceParams) (*ListManagementPacksByWorkspaceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListManagementPacksByWorkspaceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listManagementPacksByWorkspace",
		Method:             "GET",
		PathPattern:        "/v4/{workspaceId}/mpacks",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListManagementPacksByWorkspaceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListManagementPacksByWorkspaceOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
