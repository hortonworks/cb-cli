// Code generated by go-swagger; DO NOT EDIT.

package v1blueprints

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new v1blueprints API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for v1blueprints API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DeleteBlueprint deletes blueprint by id

Ambari Blueprints are a declarative definition of a Hadoop cluster. With a Blueprint, you specify a stack, the component layout and the configurations to materialize a Hadoop cluster instance. Hostgroups defined in blueprints can be associated to different templates, thus you can spin up a highly available cluster running on different instance types. This will give you the option to group your Hadoop services based on resource needs (e.g. high I/O, CPU or memory) and create an infrastructure which fits your workload best.
*/
func (a *Client) DeleteBlueprint(params *DeleteBlueprintParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteBlueprintParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteBlueprint",
		Method:             "DELETE",
		PathPattern:        "/v1/blueprints/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteBlueprintReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
DeletePrivateBlueprint deletes private blueprint by name

Ambari Blueprints are a declarative definition of a Hadoop cluster. With a Blueprint, you specify a stack, the component layout and the configurations to materialize a Hadoop cluster instance. Hostgroups defined in blueprints can be associated to different templates, thus you can spin up a highly available cluster running on different instance types. This will give you the option to group your Hadoop services based on resource needs (e.g. high I/O, CPU or memory) and create an infrastructure which fits your workload best.
*/
func (a *Client) DeletePrivateBlueprint(params *DeletePrivateBlueprintParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeletePrivateBlueprintParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deletePrivateBlueprint",
		Method:             "DELETE",
		PathPattern:        "/v1/blueprints/user/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeletePrivateBlueprintReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
DeletePublicBlueprint deletes public owned or private blueprint by name

Ambari Blueprints are a declarative definition of a Hadoop cluster. With a Blueprint, you specify a stack, the component layout and the configurations to materialize a Hadoop cluster instance. Hostgroups defined in blueprints can be associated to different templates, thus you can spin up a highly available cluster running on different instance types. This will give you the option to group your Hadoop services based on resource needs (e.g. high I/O, CPU or memory) and create an infrastructure which fits your workload best.
*/
func (a *Client) DeletePublicBlueprint(params *DeletePublicBlueprintParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeletePublicBlueprintParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deletePublicBlueprint",
		Method:             "DELETE",
		PathPattern:        "/v1/blueprints/account/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeletePublicBlueprintReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
GetBlueprint retrieves blueprint by id

Ambari Blueprints are a declarative definition of a Hadoop cluster. With a Blueprint, you specify a stack, the component layout and the configurations to materialize a Hadoop cluster instance. Hostgroups defined in blueprints can be associated to different templates, thus you can spin up a highly available cluster running on different instance types. This will give you the option to group your Hadoop services based on resource needs (e.g. high I/O, CPU or memory) and create an infrastructure which fits your workload best.
*/
func (a *Client) GetBlueprint(params *GetBlueprintParams) (*GetBlueprintOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetBlueprintParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getBlueprint",
		Method:             "GET",
		PathPattern:        "/v1/blueprints/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetBlueprintReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetBlueprintOK), nil

}

/*
GetBlueprintRequestFromID retrieves validation request by blueprint id

Ambari Blueprints are a declarative definition of a Hadoop cluster. With a Blueprint, you specify a stack, the component layout and the configurations to materialize a Hadoop cluster instance. Hostgroups defined in blueprints can be associated to different templates, thus you can spin up a highly available cluster running on different instance types. This will give you the option to group your Hadoop services based on resource needs (e.g. high I/O, CPU or memory) and create an infrastructure which fits your workload best.
*/
func (a *Client) GetBlueprintRequestFromID(params *GetBlueprintRequestFromIDParams) (*GetBlueprintRequestFromIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetBlueprintRequestFromIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getBlueprintRequestFromId",
		Method:             "GET",
		PathPattern:        "/v1/blueprints/{id}/request",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetBlueprintRequestFromIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetBlueprintRequestFromIDOK), nil

}

/*
GetPrivateBlueprint retrieves a private blueprint by name

Ambari Blueprints are a declarative definition of a Hadoop cluster. With a Blueprint, you specify a stack, the component layout and the configurations to materialize a Hadoop cluster instance. Hostgroups defined in blueprints can be associated to different templates, thus you can spin up a highly available cluster running on different instance types. This will give you the option to group your Hadoop services based on resource needs (e.g. high I/O, CPU or memory) and create an infrastructure which fits your workload best.
*/
func (a *Client) GetPrivateBlueprint(params *GetPrivateBlueprintParams) (*GetPrivateBlueprintOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPrivateBlueprintParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getPrivateBlueprint",
		Method:             "GET",
		PathPattern:        "/v1/blueprints/user/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetPrivateBlueprintReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPrivateBlueprintOK), nil

}

/*
GetPrivatesBlueprint retrieves private blueprints

Ambari Blueprints are a declarative definition of a Hadoop cluster. With a Blueprint, you specify a stack, the component layout and the configurations to materialize a Hadoop cluster instance. Hostgroups defined in blueprints can be associated to different templates, thus you can spin up a highly available cluster running on different instance types. This will give you the option to group your Hadoop services based on resource needs (e.g. high I/O, CPU or memory) and create an infrastructure which fits your workload best.
*/
func (a *Client) GetPrivatesBlueprint(params *GetPrivatesBlueprintParams) (*GetPrivatesBlueprintOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPrivatesBlueprintParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getPrivatesBlueprint",
		Method:             "GET",
		PathPattern:        "/v1/blueprints/user",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetPrivatesBlueprintReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPrivatesBlueprintOK), nil

}

/*
GetPublicBlueprint retrieves a public or private owned blueprint by name

Ambari Blueprints are a declarative definition of a Hadoop cluster. With a Blueprint, you specify a stack, the component layout and the configurations to materialize a Hadoop cluster instance. Hostgroups defined in blueprints can be associated to different templates, thus you can spin up a highly available cluster running on different instance types. This will give you the option to group your Hadoop services based on resource needs (e.g. high I/O, CPU or memory) and create an infrastructure which fits your workload best.
*/
func (a *Client) GetPublicBlueprint(params *GetPublicBlueprintParams) (*GetPublicBlueprintOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPublicBlueprintParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getPublicBlueprint",
		Method:             "GET",
		PathPattern:        "/v1/blueprints/account/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetPublicBlueprintReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPublicBlueprintOK), nil

}

/*
GetPublicsBlueprint retrieves public and private owned blueprints

Ambari Blueprints are a declarative definition of a Hadoop cluster. With a Blueprint, you specify a stack, the component layout and the configurations to materialize a Hadoop cluster instance. Hostgroups defined in blueprints can be associated to different templates, thus you can spin up a highly available cluster running on different instance types. This will give you the option to group your Hadoop services based on resource needs (e.g. high I/O, CPU or memory) and create an infrastructure which fits your workload best.
*/
func (a *Client) GetPublicsBlueprint(params *GetPublicsBlueprintParams) (*GetPublicsBlueprintOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPublicsBlueprintParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getPublicsBlueprint",
		Method:             "GET",
		PathPattern:        "/v1/blueprints/account",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetPublicsBlueprintReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPublicsBlueprintOK), nil

}

/*
PostPrivateBlueprint creates blueprint as private resource

Ambari Blueprints are a declarative definition of a Hadoop cluster. With a Blueprint, you specify a stack, the component layout and the configurations to materialize a Hadoop cluster instance. Hostgroups defined in blueprints can be associated to different templates, thus you can spin up a highly available cluster running on different instance types. This will give you the option to group your Hadoop services based on resource needs (e.g. high I/O, CPU or memory) and create an infrastructure which fits your workload best.
*/
func (a *Client) PostPrivateBlueprint(params *PostPrivateBlueprintParams) (*PostPrivateBlueprintOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostPrivateBlueprintParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "postPrivateBlueprint",
		Method:             "POST",
		PathPattern:        "/v1/blueprints/user",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostPrivateBlueprintReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostPrivateBlueprintOK), nil

}

/*
PostPublicBlueprint creates blueprint as public resource

Ambari Blueprints are a declarative definition of a Hadoop cluster. With a Blueprint, you specify a stack, the component layout and the configurations to materialize a Hadoop cluster instance. Hostgroups defined in blueprints can be associated to different templates, thus you can spin up a highly available cluster running on different instance types. This will give you the option to group your Hadoop services based on resource needs (e.g. high I/O, CPU or memory) and create an infrastructure which fits your workload best.
*/
func (a *Client) PostPublicBlueprint(params *PostPublicBlueprintParams) (*PostPublicBlueprintOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostPublicBlueprintParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "postPublicBlueprint",
		Method:             "POST",
		PathPattern:        "/v1/blueprints/account",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostPublicBlueprintReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostPublicBlueprintOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
