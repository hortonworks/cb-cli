// Code generated by go-swagger; DO NOT EDIT.

package v4_workspace_id_recipes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new v4 workspace id recipes API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for v4 workspace id recipes API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CreateRecipeInWorkspace creates recipe in workspace

Recipes are basically script extensions to a cluster that run on a set of nodes before or after the Ambari cluster installation.
*/
func (a *Client) CreateRecipeInWorkspace(params *CreateRecipeInWorkspaceParams) (*CreateRecipeInWorkspaceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateRecipeInWorkspaceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createRecipeInWorkspace",
		Method:             "POST",
		PathPattern:        "/v4/{workspaceId}/recipes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateRecipeInWorkspaceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateRecipeInWorkspaceOK), nil

}

/*
DeleteRecipeByCrnInWorkspace deletes recipe by crn in workspace

Recipes are basically script extensions to a cluster that run on a set of nodes before or after the Ambari cluster installation.
*/
func (a *Client) DeleteRecipeByCrnInWorkspace(params *DeleteRecipeByCrnInWorkspaceParams) (*DeleteRecipeByCrnInWorkspaceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteRecipeByCrnInWorkspaceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteRecipeByCrnInWorkspace",
		Method:             "DELETE",
		PathPattern:        "/v4/{workspaceId}/recipes/crn/{crn}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteRecipeByCrnInWorkspaceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteRecipeByCrnInWorkspaceOK), nil

}

/*
DeleteRecipeInWorkspace deletes recipe by name in workspace

Recipes are basically script extensions to a cluster that run on a set of nodes before or after the Ambari cluster installation.
*/
func (a *Client) DeleteRecipeInWorkspace(params *DeleteRecipeInWorkspaceParams) (*DeleteRecipeInWorkspaceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteRecipeInWorkspaceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteRecipeInWorkspace",
		Method:             "DELETE",
		PathPattern:        "/v4/{workspaceId}/recipes/name/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteRecipeInWorkspaceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteRecipeInWorkspaceOK), nil

}

/*
DeleteRecipesInWorkspace deletes multiple recipes by name in workspace

Recipes are basically script extensions to a cluster that run on a set of nodes before or after the Ambari cluster installation.
*/
func (a *Client) DeleteRecipesInWorkspace(params *DeleteRecipesInWorkspaceParams) (*DeleteRecipesInWorkspaceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteRecipesInWorkspaceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteRecipesInWorkspace",
		Method:             "DELETE",
		PathPattern:        "/v4/{workspaceId}/recipes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteRecipesInWorkspaceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteRecipesInWorkspaceOK), nil

}

/*
GetCreateRecipeRequestForCli produces cli command input

Recipes are basically script extensions to a cluster that run on a set of nodes before or after the Ambari cluster installation.
*/
func (a *Client) GetCreateRecipeRequestForCli(params *GetCreateRecipeRequestForCliParams) (*GetCreateRecipeRequestForCliOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCreateRecipeRequestForCliParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getCreateRecipeRequestForCli",
		Method:             "POST",
		PathPattern:        "/v4/{workspaceId}/recipes/cli_create",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetCreateRecipeRequestForCliReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetCreateRecipeRequestForCliOK), nil

}

/*
GetRecipeByCrnInWorkspace gets recipe by crn in workspace

Recipes are basically script extensions to a cluster that run on a set of nodes before or after the Ambari cluster installation.
*/
func (a *Client) GetRecipeByCrnInWorkspace(params *GetRecipeByCrnInWorkspaceParams) (*GetRecipeByCrnInWorkspaceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRecipeByCrnInWorkspaceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getRecipeByCrnInWorkspace",
		Method:             "GET",
		PathPattern:        "/v4/{workspaceId}/recipes/crn/{crn}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetRecipeByCrnInWorkspaceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetRecipeByCrnInWorkspaceOK), nil

}

/*
GetRecipeInWorkspace gets recipe by name in workspace

Recipes are basically script extensions to a cluster that run on a set of nodes before or after the Ambari cluster installation.
*/
func (a *Client) GetRecipeInWorkspace(params *GetRecipeInWorkspaceParams) (*GetRecipeInWorkspaceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRecipeInWorkspaceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getRecipeInWorkspace",
		Method:             "GET",
		PathPattern:        "/v4/{workspaceId}/recipes/name/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetRecipeInWorkspaceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetRecipeInWorkspaceOK), nil

}

/*
GetRecipeRequestFromNameInWorkspace retrieves recipe request by recipe name

Recipes are basically script extensions to a cluster that run on a set of nodes before or after the Ambari cluster installation.
*/
func (a *Client) GetRecipeRequestFromNameInWorkspace(params *GetRecipeRequestFromNameInWorkspaceParams) (*GetRecipeRequestFromNameInWorkspaceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRecipeRequestFromNameInWorkspaceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getRecipeRequestFromNameInWorkspace",
		Method:             "GET",
		PathPattern:        "/v4/{workspaceId}/recipes/{name}/request",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetRecipeRequestFromNameInWorkspaceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetRecipeRequestFromNameInWorkspaceOK), nil

}

/*
ListRecipesByWorkspace lists recipes for the given workspace

Recipes are basically script extensions to a cluster that run on a set of nodes before or after the Ambari cluster installation.
*/
func (a *Client) ListRecipesByWorkspace(params *ListRecipesByWorkspaceParams) (*ListRecipesByWorkspaceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListRecipesByWorkspaceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listRecipesByWorkspace",
		Method:             "GET",
		PathPattern:        "/v4/{workspaceId}/recipes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListRecipesByWorkspaceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListRecipesByWorkspaceOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
