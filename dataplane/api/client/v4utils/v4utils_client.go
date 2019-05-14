// Code generated by go-swagger; DO NOT EDIT.

package v4utils

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new v4utils API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for v4utils API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CheckClientVersionV4 checks the client version
*/
func (a *Client) CheckClientVersionV4(params *CheckClientVersionV4Params) (*CheckClientVersionV4OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCheckClientVersionV4Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "checkClientVersionV4",
		Method:             "GET",
		PathPattern:        "/v4/utils/client",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CheckClientVersionV4Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CheckClientVersionV4OK), nil

}

/*
GetCloudStorageMatrixV4 returns supported cloud storage for stack version

Define stack version at least at patch level eg. 2.6.0
*/
func (a *Client) GetCloudStorageMatrixV4(params *GetCloudStorageMatrixV4Params) (*GetCloudStorageMatrixV4OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCloudStorageMatrixV4Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getCloudStorageMatrixV4",
		Method:             "GET",
		PathPattern:        "/v4/utils/cloud_storage_matrix",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetCloudStorageMatrixV4Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetCloudStorageMatrixV4OK), nil

}

/*
GetDefaultSecurityRules gets default security rules

Security Rules operations
*/
func (a *Client) GetDefaultSecurityRules(params *GetDefaultSecurityRulesParams) (*GetDefaultSecurityRulesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDefaultSecurityRulesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getDefaultSecurityRules",
		Method:             "GET",
		PathPattern:        "/v4/utils/default_security_rules",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetDefaultSecurityRulesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetDefaultSecurityRulesOK), nil

}

/*
GetDeploymentInfo retrieves account preferences for admin user

Account related preferences that could be managed by the account admins and different restrictions could be added to Cloudbreak resources.
*/
func (a *Client) GetDeploymentInfo(params *GetDeploymentInfoParams) (*GetDeploymentInfoOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDeploymentInfoParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getDeploymentInfo",
		Method:             "GET",
		PathPattern:        "/v4/utils/deployment",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetDeploymentInfoReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetDeploymentInfoOK), nil

}

/*
GetStackMatrixUtilV4 returns default ambari details for distinct h d p and h d f
*/
func (a *Client) GetStackMatrixUtilV4(params *GetStackMatrixUtilV4Params) (*GetStackMatrixUtilV4OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetStackMatrixUtilV4Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getStackMatrixUtilV4",
		Method:             "GET",
		PathPattern:        "/v4/utils/stack_matrix",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetStackMatrixUtilV4Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetStackMatrixUtilV4OK), nil

}

/*
PostNotificationTest triggers a new notification to the notification system could be validated from the begins

Account related preferences that could be managed by the account admins and different restrictions could be added to Cloudbreak resources.
*/
func (a *Client) PostNotificationTest(params *PostNotificationTestParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostNotificationTestParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "postNotificationTest",
		Method:             "POST",
		PathPattern:        "/v4/utils/notification_test",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostNotificationTestReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
RepositoryConfigsValidationV4 validates repository configs fields check their availability

Repository configs validation related operations
*/
func (a *Client) RepositoryConfigsValidationV4(params *RepositoryConfigsValidationV4Params) (*RepositoryConfigsValidationV4OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRepositoryConfigsValidationV4Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "repositoryConfigsValidationV4",
		Method:             "POST",
		PathPattern:        "/v4/utils/validate_repository",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RepositoryConfigsValidationV4Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*RepositoryConfigsValidationV4OK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
