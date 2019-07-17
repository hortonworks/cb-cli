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
ChangeCredentialInEnvironmentV1 changes the credential of the environment and the clusters in the environment of a given name

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
		PathPattern:        "/v1/env/name/{name}/change_credential",
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
ChangeCredentialInEnvironmentV1ByCrn changes the credential of the environment and the clusters in the environment of a given c r n

Environment consists of a credential and various other resources and enables users to quickly create clusters in given regions in a given cloud provider.
*/
func (a *Client) ChangeCredentialInEnvironmentV1ByCrn(params *ChangeCredentialInEnvironmentV1ByCrnParams) (*ChangeCredentialInEnvironmentV1ByCrnOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewChangeCredentialInEnvironmentV1ByCrnParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "changeCredentialInEnvironmentV1ByCrn",
		Method:             "PUT",
		PathPattern:        "/v1/env/crn/{crn}/change_credential",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ChangeCredentialInEnvironmentV1ByCrnReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ChangeCredentialInEnvironmentV1ByCrnOK), nil

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
DeleteEnvironmentV1ByCrn deletes an environment by c r n only possible if no cluster is running in the environment

Environment consists of a credential and various other resources and enables users to quickly create clusters in given regions in a given cloud provider.
*/
func (a *Client) DeleteEnvironmentV1ByCrn(params *DeleteEnvironmentV1ByCrnParams) (*DeleteEnvironmentV1ByCrnOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteEnvironmentV1ByCrnParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteEnvironmentV1ByCrn",
		Method:             "DELETE",
		PathPattern:        "/v1/env/crn/{crn}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteEnvironmentV1ByCrnReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteEnvironmentV1ByCrnOK), nil

}

/*
DeleteEnvironmentV1ByName deletes an environment by name only possible if no cluster is running in the environment

Environment consists of a credential and various other resources and enables users to quickly create clusters in given regions in a given cloud provider.
*/
func (a *Client) DeleteEnvironmentV1ByName(params *DeleteEnvironmentV1ByNameParams) (*DeleteEnvironmentV1ByNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteEnvironmentV1ByNameParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteEnvironmentV1ByName",
		Method:             "DELETE",
		PathPattern:        "/v1/env/name/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteEnvironmentV1ByNameReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteEnvironmentV1ByNameOK), nil

}

/*
DeleteEnvironmentsByCrn deletes multiple environment by c r ns only possible if no cluster is running in the environments

Environment consists of a credential and various other resources and enables users to quickly create clusters in given regions in a given cloud provider.
*/
func (a *Client) DeleteEnvironmentsByCrn(params *DeleteEnvironmentsByCrnParams) (*DeleteEnvironmentsByCrnOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteEnvironmentsByCrnParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteEnvironmentsByCrn",
		Method:             "DELETE",
		PathPattern:        "/v1/env/crn",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteEnvironmentsByCrnReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteEnvironmentsByCrnOK), nil

}

/*
DeleteEnvironmentsByName deletes multiple environment by namns only possible if no cluster is running in the environments

Environment consists of a credential and various other resources and enables users to quickly create clusters in given regions in a given cloud provider.
*/
func (a *Client) DeleteEnvironmentsByName(params *DeleteEnvironmentsByNameParams) (*DeleteEnvironmentsByNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteEnvironmentsByNameParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteEnvironmentsByName",
		Method:             "DELETE",
		PathPattern:        "/v1/env/name",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteEnvironmentsByNameReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteEnvironmentsByNameOK), nil

}

/*
EditEnvironmentV1 edits an environment by name location regions and description can be changed

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
		PathPattern:        "/v1/env/name/{name}",
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
EditEnvironmentV1ByCrn edits an environment by c r n location regions and description can be changed

Environment consists of a credential and various other resources and enables users to quickly create clusters in given regions in a given cloud provider.
*/
func (a *Client) EditEnvironmentV1ByCrn(params *EditEnvironmentV1ByCrnParams) (*EditEnvironmentV1ByCrnOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEditEnvironmentV1ByCrnParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "editEnvironmentV1ByCrn",
		Method:             "PUT",
		PathPattern:        "/v1/env/crn/{crn}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &EditEnvironmentV1ByCrnReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*EditEnvironmentV1ByCrnOK), nil

}

/*
GetEnvironmentV1ByCrn gets an environment by c r n

Environment consists of a credential and various other resources and enables users to quickly create clusters in given regions in a given cloud provider.
*/
func (a *Client) GetEnvironmentV1ByCrn(params *GetEnvironmentV1ByCrnParams) (*GetEnvironmentV1ByCrnOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetEnvironmentV1ByCrnParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getEnvironmentV1ByCrn",
		Method:             "GET",
		PathPattern:        "/v1/env/crn/{crn}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetEnvironmentV1ByCrnReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetEnvironmentV1ByCrnOK), nil

}

/*
GetEnvironmentV1ByName gets an environment by name

Environment consists of a credential and various other resources and enables users to quickly create clusters in given regions in a given cloud provider.
*/
func (a *Client) GetEnvironmentV1ByName(params *GetEnvironmentV1ByNameParams) (*GetEnvironmentV1ByNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetEnvironmentV1ByNameParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getEnvironmentV1ByName",
		Method:             "GET",
		PathPattern:        "/v1/env/name/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetEnvironmentV1ByNameReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetEnvironmentV1ByNameOK), nil

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
