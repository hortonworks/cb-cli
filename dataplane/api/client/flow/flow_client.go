// Code generated by go-swagger; DO NOT EDIT.

package flow

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new flow API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for flow API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetFlowChainsStatusesByChainIds gets flow check responses for parent chains input size max 50

Flow check log operations
*/
func (a *Client) GetFlowChainsStatusesByChainIds(params *GetFlowChainsStatusesByChainIdsParams) (*GetFlowChainsStatusesByChainIdsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFlowChainsStatusesByChainIdsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getFlowChainsStatusesByChainIds",
		Method:             "GET",
		PathPattern:        "/flow/check/chainIds",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetFlowChainsStatusesByChainIdsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetFlowChainsStatusesByChainIdsOK), nil

}

/*
GetFlowLogsByFlowID gets flow logs by flow id

Flow log operations
*/
func (a *Client) GetFlowLogsByFlowID(params *GetFlowLogsByFlowIDParams) (*GetFlowLogsByFlowIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFlowLogsByFlowIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getFlowLogsByFlowId",
		Method:             "GET",
		PathPattern:        "/flow/logs/{flowId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetFlowLogsByFlowIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetFlowLogsByFlowIDOK), nil

}

/*
GetFlowLogsByFlowIds gets flow logs by a list of flow ids input size max 50

Flow log operations
*/
func (a *Client) GetFlowLogsByFlowIds(params *GetFlowLogsByFlowIdsParams) (*GetFlowLogsByFlowIdsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFlowLogsByFlowIdsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getFlowLogsByFlowIds",
		Method:             "GET",
		PathPattern:        "/flow/logs/flowIds",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetFlowLogsByFlowIdsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetFlowLogsByFlowIdsOK), nil

}

/*
GetFlowLogsByResourceCrn gets flow logs for resource by resource c r n

Flow log operations
*/
func (a *Client) GetFlowLogsByResourceCrn(params *GetFlowLogsByResourceCrnParams) (*GetFlowLogsByResourceCrnOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFlowLogsByResourceCrnParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getFlowLogsByResourceCrn",
		Method:             "GET",
		PathPattern:        "/flow/logs/resource/crn/{resourceCrn}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetFlowLogsByResourceCrnReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetFlowLogsByResourceCrnOK), nil

}

/*
GetFlowLogsByResourceName gets flow logs for resource by resource name

Flow log operations
*/
func (a *Client) GetFlowLogsByResourceName(params *GetFlowLogsByResourceNameParams) (*GetFlowLogsByResourceNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFlowLogsByResourceNameParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getFlowLogsByResourceName",
		Method:             "GET",
		PathPattern:        "/flow/logs/resource/name/{resourceName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetFlowLogsByResourceNameReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetFlowLogsByResourceNameOK), nil

}

/*
GetLastFlowByID gets last flow log by flow id

Flow log operations
*/
func (a *Client) GetLastFlowByID(params *GetLastFlowByIDParams) (*GetLastFlowByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetLastFlowByIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getLastFlowById",
		Method:             "GET",
		PathPattern:        "/flow/logs/{flowId}/last",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetLastFlowByIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetLastFlowByIDOK), nil

}

/*
GetLastFlowByResourceCrn gets last flow log for resource by resource c r n

Flow log operations
*/
func (a *Client) GetLastFlowByResourceCrn(params *GetLastFlowByResourceCrnParams) (*GetLastFlowByResourceCrnOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetLastFlowByResourceCrnParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getLastFlowByResourceCrn",
		Method:             "GET",
		PathPattern:        "/flow/logs/resource/crn/{resourceCrn}/last",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetLastFlowByResourceCrnReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetLastFlowByResourceCrnOK), nil

}

/*
GetLastFlowByResourceName gets last flow log for resource by resource name

Flow log operations
*/
func (a *Client) GetLastFlowByResourceName(params *GetLastFlowByResourceNameParams) (*GetLastFlowByResourceNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetLastFlowByResourceNameParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getLastFlowByResourceName",
		Method:             "GET",
		PathPattern:        "/flow/logs/resource/name/{resourceName}/last",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetLastFlowByResourceNameReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetLastFlowByResourceNameOK), nil

}

/*
HasFlowRunningByChainID checks if there is a running flow for chain id

Flow log operations
*/
func (a *Client) HasFlowRunningByChainID(params *HasFlowRunningByChainIDParams) (*HasFlowRunningByChainIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewHasFlowRunningByChainIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "hasFlowRunningByChainId",
		Method:             "GET",
		PathPattern:        "/flow/check/chainId/{chainId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &HasFlowRunningByChainIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*HasFlowRunningByChainIDOK), nil

}

/*
HasFlowRunningByFlowID checks if there is a running flow for flow id

Flow log operations
*/
func (a *Client) HasFlowRunningByFlowID(params *HasFlowRunningByFlowIDParams) (*HasFlowRunningByFlowIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewHasFlowRunningByFlowIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "hasFlowRunningByFlowId",
		Method:             "GET",
		PathPattern:        "/flow/check/flowId/{flowId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &HasFlowRunningByFlowIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*HasFlowRunningByFlowIDOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
