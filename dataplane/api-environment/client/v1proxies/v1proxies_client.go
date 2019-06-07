// Code generated by go-swagger; DO NOT EDIT.

package v1proxies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new v1proxies API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for v1proxies API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CreateProxyConfigV1 creates proxy configuration

An proxy Configuration describe a connection to an external proxy server which provides internet access cluster members. It's applied for package manager and Ambari too
*/
func (a *Client) CreateProxyConfigV1(params *CreateProxyConfigV1Params) (*CreateProxyConfigV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateProxyConfigV1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createProxyConfigV1",
		Method:             "POST",
		PathPattern:        "/v1/proxies",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateProxyConfigV1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateProxyConfigV1OK), nil

}

/*
DeleteProxyConfigByCrnV1 deletes proxy configuration by crn

An proxy Configuration describe a connection to an external proxy server which provides internet access cluster members. It's applied for package manager and Ambari too
*/
func (a *Client) DeleteProxyConfigByCrnV1(params *DeleteProxyConfigByCrnV1Params) (*DeleteProxyConfigByCrnV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteProxyConfigByCrnV1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteProxyConfigByCrnV1",
		Method:             "DELETE",
		PathPattern:        "/v1/proxies/crn/{crn}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteProxyConfigByCrnV1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteProxyConfigByCrnV1OK), nil

}

/*
DeleteProxyConfigByNameV1 deletes proxy configuration by name

An proxy Configuration describe a connection to an external proxy server which provides internet access cluster members. It's applied for package manager and Ambari too
*/
func (a *Client) DeleteProxyConfigByNameV1(params *DeleteProxyConfigByNameV1Params) (*DeleteProxyConfigByNameV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteProxyConfigByNameV1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteProxyConfigByNameV1",
		Method:             "DELETE",
		PathPattern:        "/v1/proxies/name/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteProxyConfigByNameV1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteProxyConfigByNameV1OK), nil

}

/*
DeleteProxyConfigsV1 deletes multiple proxy configurations by name

An proxy Configuration describe a connection to an external proxy server which provides internet access cluster members. It's applied for package manager and Ambari too
*/
func (a *Client) DeleteProxyConfigsV1(params *DeleteProxyConfigsV1Params) (*DeleteProxyConfigsV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteProxyConfigsV1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteProxyConfigsV1",
		Method:             "DELETE",
		PathPattern:        "/v1/proxies",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteProxyConfigsV1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteProxyConfigsV1OK), nil

}

/*
GetProxyConfigByCrnV1 gets proxy configuration

An proxy Configuration describe a connection to an external proxy server which provides internet access cluster members. It's applied for package manager and Ambari too
*/
func (a *Client) GetProxyConfigByCrnV1(params *GetProxyConfigByCrnV1Params) (*GetProxyConfigByCrnV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetProxyConfigByCrnV1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getProxyConfigByCrnV1",
		Method:             "GET",
		PathPattern:        "/v1/proxies/crn/{crn}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetProxyConfigByCrnV1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetProxyConfigByCrnV1OK), nil

}

/*
GetProxyConfigByNameV1 gets proxy configuration

An proxy Configuration describe a connection to an external proxy server which provides internet access cluster members. It's applied for package manager and Ambari too
*/
func (a *Client) GetProxyConfigByNameV1(params *GetProxyConfigByNameV1Params) (*GetProxyConfigByNameV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetProxyConfigByNameV1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getProxyConfigByNameV1",
		Method:             "GET",
		PathPattern:        "/v1/proxies/name/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetProxyConfigByNameV1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetProxyConfigByNameV1OK), nil

}

/*
GetProxyRequestFromNameV1 gets request by name

An proxy Configuration describe a connection to an external proxy server which provides internet access cluster members. It's applied for package manager and Ambari too
*/
func (a *Client) GetProxyRequestFromNameV1(params *GetProxyRequestFromNameV1Params) (*GetProxyRequestFromNameV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetProxyRequestFromNameV1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getProxyRequestFromNameV1",
		Method:             "GET",
		PathPattern:        "/v1/proxies/{name}/request",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetProxyRequestFromNameV1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetProxyRequestFromNameV1OK), nil

}

/*
ListProxyConfigsV1 lists proxy configurations

An proxy Configuration describe a connection to an external proxy server which provides internet access cluster members. It's applied for package manager and Ambari too
*/
func (a *Client) ListProxyConfigsV1(params *ListProxyConfigsV1Params) (*ListProxyConfigsV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListProxyConfigsV1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listProxyConfigsV1",
		Method:             "GET",
		PathPattern:        "/v1/proxies",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListProxyConfigsV1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListProxyConfigsV1OK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
