// Code generated by go-swagger; DO NOT EDIT.

package v4_workspace_id_environments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new v4 workspace id environments API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for v4 workspace id environments API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
AttachResourcesToEnvironment attaches resources to an environment

Environment consists of a credential and various other resources and enables users to quickly create clusters in given regions in a given cloud provider.
*/
func (a *Client) AttachResourcesToEnvironment(params *AttachResourcesToEnvironmentParams) (*AttachResourcesToEnvironmentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAttachResourcesToEnvironmentParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "attachResourcesToEnvironment",
		Method:             "PUT",
		PathPattern:        "/v4/{workspaceId}/environments/{crn}/attach",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AttachResourcesToEnvironmentReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AttachResourcesToEnvironmentOK), nil

}

/*
CreateEnvironment creates an environment

Environment consists of a credential and various other resources and enables users to quickly create clusters in given regions in a given cloud provider.
*/
func (a *Client) CreateEnvironment(params *CreateEnvironmentParams) (*CreateEnvironmentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateEnvironmentParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createEnvironment",
		Method:             "POST",
		PathPattern:        "/v4/{workspaceId}/environments",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateEnvironmentReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateEnvironmentOK), nil

}

/*
DeleteEnvironment deletes an environment only possible if no cluster is running in the environment

Environment consists of a credential and various other resources and enables users to quickly create clusters in given regions in a given cloud provider.
*/
func (a *Client) DeleteEnvironment(params *DeleteEnvironmentParams) (*DeleteEnvironmentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteEnvironmentParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteEnvironment",
		Method:             "DELETE",
		PathPattern:        "/v4/{workspaceId}/environments/{crn}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteEnvironmentReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteEnvironmentOK), nil

}

/*
DetachResourcesFromEnvironment detaches resources from an environment

Environment consists of a credential and various other resources and enables users to quickly create clusters in given regions in a given cloud provider.
*/
func (a *Client) DetachResourcesFromEnvironment(params *DetachResourcesFromEnvironmentParams) (*DetachResourcesFromEnvironmentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDetachResourcesFromEnvironmentParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "detachResourcesFromEnvironment",
		Method:             "PUT",
		PathPattern:        "/v4/{workspaceId}/environments/{crn}/detach",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DetachResourcesFromEnvironmentReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DetachResourcesFromEnvironmentOK), nil

}

/*
EditEnvironment edits and environment location regions and description can be changed

Environment consists of a credential and various other resources and enables users to quickly create clusters in given regions in a given cloud provider.
*/
func (a *Client) EditEnvironment(params *EditEnvironmentParams) (*EditEnvironmentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEditEnvironmentParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "editEnvironment",
		Method:             "PUT",
		PathPattern:        "/v4/{workspaceId}/environments/{crn}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &EditEnvironmentReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*EditEnvironmentOK), nil

}

/*
GetEnvironment gets an environment

Environment consists of a credential and various other resources and enables users to quickly create clusters in given regions in a given cloud provider.
*/
func (a *Client) GetEnvironment(params *GetEnvironmentParams) (*GetEnvironmentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetEnvironmentParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getEnvironment",
		Method:             "GET",
		PathPattern:        "/v4/{workspaceId}/environments/{crn}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetEnvironmentReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetEnvironmentOK), nil

}

/*
ListEnvironment lists all environments in the workspace

Environment consists of a credential and various other resources and enables users to quickly create clusters in given regions in a given cloud provider.
*/
func (a *Client) ListEnvironment(params *ListEnvironmentParams) (*ListEnvironmentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListEnvironmentParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listEnvironment",
		Method:             "GET",
		PathPattern:        "/v4/{workspaceId}/environments",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListEnvironmentReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListEnvironmentOK), nil

}

/*
RegisterDatalakePrerequisites returns datalake prerequisites
*/
func (a *Client) RegisterDatalakePrerequisites(params *RegisterDatalakePrerequisitesParams) (*RegisterDatalakePrerequisitesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRegisterDatalakePrerequisitesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "registerDatalakePrerequisites",
		Method:             "POST",
		PathPattern:        "/v4/{workspaceId}/environments/{crn}/register_datalake_prerequisites",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RegisterDatalakePrerequisitesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*RegisterDatalakePrerequisitesOK), nil

}

/*
RegisterExternalDatalake registers external datalake

Environment consists of a credential and various other resources and enables users to quickly create clusters in given regions in a given cloud provider.
*/
func (a *Client) RegisterExternalDatalake(params *RegisterExternalDatalakeParams) (*RegisterExternalDatalakeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRegisterExternalDatalakeParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "registerExternalDatalake",
		Method:             "PUT",
		PathPattern:        "/v4/{workspaceId}/environments/{crn}/register_datalake",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RegisterExternalDatalakeReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*RegisterExternalDatalakeOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
