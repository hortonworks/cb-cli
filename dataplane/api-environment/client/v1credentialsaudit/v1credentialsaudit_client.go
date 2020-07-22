// Code generated by go-swagger; DO NOT EDIT.

package v1credentialsaudit

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new v1credentialsaudit API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for v1credentialsaudit API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CreateAuditCredentialV1 creates credential

Cloudbreak is launching Hadoop clusters on the user's behalf - on different cloud providers. One key point is that Cloudbreak does not store your Cloud provider account details (such as username, password, keys, private SSL certificates, etc). We work around the concept that Identity and Access Management is fully controlled by you - the end user. The Cloudbreak deployer is purely acting on behalf of the end user - without having access to the user's account.
*/
func (a *Client) CreateAuditCredentialV1(params *CreateAuditCredentialV1Params) (*CreateAuditCredentialV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateAuditCredentialV1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createAuditCredentialV1",
		Method:             "POST",
		PathPattern:        "/v1/credentials/audit",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateAuditCredentialV1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateAuditCredentialV1OK), nil

}

/*
GetAuditCredentialByResourceCrnV1 gets credential by crn

Cloudbreak is launching Hadoop clusters on the user's behalf - on different cloud providers. One key point is that Cloudbreak does not store your Cloud provider account details (such as username, password, keys, private SSL certificates, etc). We work around the concept that Identity and Access Management is fully controlled by you - the end user. The Cloudbreak deployer is purely acting on behalf of the end user - without having access to the user's account.
*/
func (a *Client) GetAuditCredentialByResourceCrnV1(params *GetAuditCredentialByResourceCrnV1Params) (*GetAuditCredentialByResourceCrnV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAuditCredentialByResourceCrnV1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getAuditCredentialByResourceCrnV1",
		Method:             "GET",
		PathPattern:        "/v1/credentials/audit/crn/{crn}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetAuditCredentialByResourceCrnV1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetAuditCredentialByResourceCrnV1OK), nil

}

/*
GetAuditPrerequisitesForCloudPlatform gets credential prerequisites for cloud platform

Cloudbreak is launching Hadoop clusters on the user's behalf - on different cloud providers. One key point is that Cloudbreak does not store your Cloud provider account details (such as username, password, keys, private SSL certificates, etc). We work around the concept that Identity and Access Management is fully controlled by you - the end user. The Cloudbreak deployer is purely acting on behalf of the end user - without having access to the user's account.
*/
func (a *Client) GetAuditPrerequisitesForCloudPlatform(params *GetAuditPrerequisitesForCloudPlatformParams) (*GetAuditPrerequisitesForCloudPlatformOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAuditPrerequisitesForCloudPlatformParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getAuditPrerequisitesForCloudPlatform",
		Method:             "GET",
		PathPattern:        "/v1/credentials/audit/prerequisites/{cloudPlatform}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetAuditPrerequisitesForCloudPlatformReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetAuditPrerequisitesForCloudPlatformOK), nil

}

/*
ListAuditCredentialsV1 lists credentials

Cloudbreak is launching Hadoop clusters on the user's behalf - on different cloud providers. One key point is that Cloudbreak does not store your Cloud provider account details (such as username, password, keys, private SSL certificates, etc). We work around the concept that Identity and Access Management is fully controlled by you - the end user. The Cloudbreak deployer is purely acting on behalf of the end user - without having access to the user's account.
*/
func (a *Client) ListAuditCredentialsV1(params *ListAuditCredentialsV1Params) (*ListAuditCredentialsV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListAuditCredentialsV1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listAuditCredentialsV1",
		Method:             "GET",
		PathPattern:        "/v1/credentials/audit",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListAuditCredentialsV1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListAuditCredentialsV1OK), nil

}

/*
PutAuditCredentialV1 modifies public credential resource

Cloudbreak is launching Hadoop clusters on the user's behalf - on different cloud providers. One key point is that Cloudbreak does not store your Cloud provider account details (such as username, password, keys, private SSL certificates, etc). We work around the concept that Identity and Access Management is fully controlled by you - the end user. The Cloudbreak deployer is purely acting on behalf of the end user - without having access to the user's account.
*/
func (a *Client) PutAuditCredentialV1(params *PutAuditCredentialV1Params) (*PutAuditCredentialV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutAuditCredentialV1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "putAuditCredentialV1",
		Method:             "PUT",
		PathPattern:        "/v1/credentials/audit",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PutAuditCredentialV1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PutAuditCredentialV1OK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
