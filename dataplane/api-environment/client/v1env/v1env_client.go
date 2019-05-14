// Code generated by go-swagger; DO NOT EDIT.

package v1env

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new v1env API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for v1env API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
ChangeCredentialInEnvironmentV1 changes the credential of the environment and the clusters in the environment

Environment consists of a credential and various other resources and enables users to quickly create clusters in given regions in a given cloud provider.
*/
func (a *Client) ChangeCredentialInEnvironmentV1(params *ChangeCredentialInEnvironmentV1Params) (*ChangeCredentialInEnvironmentV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewChangeCredentialInEnvironmentV1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "changeCredentialInEnvironmentV1",
		Method:             "PUT",
		PathPattern:        "/v1/env/{name}/change_credential",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ChangeCredentialInEnvironmentV1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ChangeCredentialInEnvironmentV1OK), nil

}

/*
CreateEnvironmentV1 creates an environment

Environment consists of a credential and various other resources and enables users to quickly create clusters in given regions in a given cloud provider.
*/
func (a *Client) CreateEnvironmentV1(params *CreateEnvironmentV1Params) (*CreateEnvironmentV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateEnvironmentV1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createEnvironmentV1",
		Method:             "POST",
		PathPattern:        "/v1/env",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateEnvironmentV1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateEnvironmentV1OK), nil

}

/*
DeleteEnvironments deletes multiple environment only possible if no cluster is running in the environments

Environment consists of a credential and various other resources and enables users to quickly create clusters in given regions in a given cloud provider.
*/
func (a *Client) DeleteEnvironments(params *DeleteEnvironmentsParams) (*DeleteEnvironmentsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteEnvironmentsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteEnvironments",
		Method:             "DELETE",
		PathPattern:        "/v1/env",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteEnvironmentsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteEnvironmentsOK), nil

}

/*
EditEnvironmentV1 edits and environment location regions and description can be changed

Environment consists of a credential and various other resources and enables users to quickly create clusters in given regions in a given cloud provider.
*/
func (a *Client) EditEnvironmentV1(params *EditEnvironmentV1Params) (*EditEnvironmentV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEditEnvironmentV1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "editEnvironmentV1",
		Method:             "PUT",
		PathPattern:        "/v1/env/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &EditEnvironmentV1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*EditEnvironmentV1OK), nil

}

/*
GetEnvironmentV1 gets an environment

Environment consists of a credential and various other resources and enables users to quickly create clusters in given regions in a given cloud provider.
*/
func (a *Client) GetEnvironmentV1(params *GetEnvironmentV1Params) (*GetEnvironmentV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetEnvironmentV1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getEnvironmentV1",
		Method:             "GET",
		PathPattern:        "/v1/env/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetEnvironmentV1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetEnvironmentV1OK), nil

}

/*
GetWelcomeMessage welcomes
*/
func (a *Client) GetWelcomeMessage(params *GetWelcomeMessageParams) (*GetWelcomeMessageOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetWelcomeMessageParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getWelcomeMessage",
		Method:             "GET",
		PathPattern:        "/v1/env/welcome",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetWelcomeMessageReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetWelcomeMessageOK), nil

}

/*
ListEnvironmentV1 lists all environments

Environment consists of a credential and various other resources and enables users to quickly create clusters in given regions in a given cloud provider.
*/
func (a *Client) ListEnvironmentV1(params *ListEnvironmentV1Params) (*ListEnvironmentV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListEnvironmentV1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listEnvironmentV1",
		Method:             "GET",
		PathPattern:        "/v1/env",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListEnvironmentV1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListEnvironmentV1OK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
