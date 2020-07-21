// Code generated by go-swagger; DO NOT EDIT.

package diagnostics

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new diagnostics API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for diagnostics API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CollectStackCmDiagnostics initiates the collection of diagnostical data on the cloudera manager host
*/
func (a *Client) CollectStackCmDiagnostics(params *CollectStackCmDiagnosticsParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCollectStackCmDiagnosticsParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "collectStackCmDiagnostics",
		Method:             "POST",
		PathPattern:        "/diagnostics",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CollectStackCmDiagnosticsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
GetStackCmVMLogs returns a list of log paths on the cloudera manager VM
*/
func (a *Client) GetStackCmVMLogs(params *GetStackCmVMLogsParams) (*GetStackCmVMLogsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetStackCmVMLogsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getStackCmVmLogs",
		Method:             "GET",
		PathPattern:        "/diagnostics/logs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetStackCmVMLogsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetStackCmVMLogsOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
