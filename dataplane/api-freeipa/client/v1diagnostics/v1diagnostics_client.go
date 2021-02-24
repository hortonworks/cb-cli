// Code generated by go-swagger; DO NOT EDIT.

package v1diagnostics

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new v1diagnostics API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for v1diagnostics API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CancelDiagnosticsCollectionsV1 cancels the not finished diagnostics collections

FreeIPA is an integrated Identity and Authentication solution that can be used for any of CM, CDP services.
*/
func (a *Client) CancelDiagnosticsCollectionsV1(params *CancelDiagnosticsCollectionsV1Params) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCancelDiagnosticsCollectionsV1Params()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "cancelDiagnosticsCollectionsV1",
		Method:             "POST",
		PathPattern:        "/v1/diagnostics/{environmentCrn}/collections/cancel",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CancelDiagnosticsCollectionsV1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
CollectFreeIpaDiagnosticsV1 initiates the collection of diagnostical data on the free IP a host

FreeIPA is an integrated Identity and Authentication solution that can be used for any of CM, CDP services.
*/
func (a *Client) CollectFreeIpaDiagnosticsV1(params *CollectFreeIpaDiagnosticsV1Params) (*CollectFreeIpaDiagnosticsV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCollectFreeIpaDiagnosticsV1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "collectFreeIpaDiagnosticsV1",
		Method:             "POST",
		PathPattern:        "/v1/diagnostics",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CollectFreeIpaDiagnosticsV1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CollectFreeIpaDiagnosticsV1OK), nil

}

/*
GetFreeIpaVMLogsV1 returns a list of log paths on the free ipa VM

FreeIPA is an integrated Identity and Authentication solution that can be used for any of CM, CDP services.
*/
func (a *Client) GetFreeIpaVMLogsV1(params *GetFreeIpaVMLogsV1Params) (*GetFreeIpaVMLogsV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFreeIpaVMLogsV1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getFreeIpaVmLogsV1",
		Method:             "GET",
		PathPattern:        "/v1/diagnostics/logs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetFreeIpaVMLogsV1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetFreeIpaVMLogsV1OK), nil

}

/*
ListDiagnosticsCollectionsV1 returns a list of diagnostics collections for the specified free ipa cluster

FreeIPA is an integrated Identity and Authentication solution that can be used for any of CM, CDP services.
*/
func (a *Client) ListDiagnosticsCollectionsV1(params *ListDiagnosticsCollectionsV1Params) (*ListDiagnosticsCollectionsV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListDiagnosticsCollectionsV1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listDiagnosticsCollectionsV1",
		Method:             "GET",
		PathPattern:        "/v1/diagnostics/{environmentCrn}/collections",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListDiagnosticsCollectionsV1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListDiagnosticsCollectionsV1OK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
