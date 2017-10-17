// Code generated by go-swagger; DO NOT EDIT.

package v1topologies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new v1topologies API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for v1topologies API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DeleteTopology deletes topology by id

A topology gives system administrators an easy way to associate compute nodes with data centers and racks.
*/
func (a *Client) DeleteTopology(params *DeleteTopologyParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteTopologyParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteTopology",
		Method:             "DELETE",
		PathPattern:        "/v1/topologies/account/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteTopologyReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
GetPublicsTopology retrieves topoligies

A topology gives system administrators an easy way to associate compute nodes with data centers and racks.
*/
func (a *Client) GetPublicsTopology(params *GetPublicsTopologyParams) (*GetPublicsTopologyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPublicsTopologyParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getPublicsTopology",
		Method:             "GET",
		PathPattern:        "/v1/topologies/account",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetPublicsTopologyReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPublicsTopologyOK), nil

}

/*
GetTopology retrieves topology by id

A template gives developers and systems administrators an easy way to create and manage a collection of cloud infrastructure related resources, maintaining and updating them in an orderly and predictable fashion. Templates are cloud specific - and on top of the infrastructural setup they collect the information such as the used machine images, the datacenter location, instance types, and can capture and control region-specific infrastructure variations. We support heterogenous clusters - this one Hadoop cluster can be built by combining different templates.
*/
func (a *Client) GetTopology(params *GetTopologyParams) (*GetTopologyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTopologyParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getTopology",
		Method:             "GET",
		PathPattern:        "/v1/topologies/account/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetTopologyReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetTopologyOK), nil

}

/*
PostPublicTopology creates topology as public resource

A topology gives system administrators an easy way to associate compute nodes with data centers and racks.
*/
func (a *Client) PostPublicTopology(params *PostPublicTopologyParams) (*PostPublicTopologyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostPublicTopologyParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "postPublicTopology",
		Method:             "POST",
		PathPattern:        "/v1/topologies/account",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostPublicTopologyReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostPublicTopologyOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
