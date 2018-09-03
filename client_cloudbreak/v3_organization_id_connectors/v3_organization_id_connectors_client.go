// Code generated by go-swagger; DO NOT EDIT.

package v3_organization_id_connectors

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new v3 organization id connectors API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for v3 organization id connectors API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CreateRecommendationForOrganization creates a recommendation that advises cloud resources for the given blueprint

Each cloud provider has it's own specific resources like instance types and disk types. These endpoints are collecting them.
*/
func (a *Client) CreateRecommendationForOrganization(params *CreateRecommendationForOrganizationParams) (*CreateRecommendationForOrganizationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateRecommendationForOrganizationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createRecommendationForOrganization",
		Method:             "POST",
		PathPattern:        "/v3/{organizationId}/connectors/recommendation",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateRecommendationForOrganizationReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateRecommendationForOrganizationOK), nil

}

/*
GetAccessConfigsForOrganization retrives access configs with properties

Each cloud provider has it's own specific resources like instance types and disk types. These endpoints are collecting them.
*/
func (a *Client) GetAccessConfigsForOrganization(params *GetAccessConfigsForOrganizationParams) (*GetAccessConfigsForOrganizationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAccessConfigsForOrganizationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getAccessConfigsForOrganization",
		Method:             "POST",
		PathPattern:        "/v3/{organizationId}/connectors/accessconfigs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetAccessConfigsForOrganizationReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetAccessConfigsForOrganizationOK), nil

}

/*
GetDisktypesForOrganization retrives available disk types

Each cloud provider has it's own specific resources like instance types and disk types. These endpoints are collecting them.
*/
func (a *Client) GetDisktypesForOrganization(params *GetDisktypesForOrganizationParams) (*GetDisktypesForOrganizationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDisktypesForOrganizationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getDisktypesForOrganization",
		Method:             "GET",
		PathPattern:        "/v3/{organizationId}/connectors/disktypes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetDisktypesForOrganizationReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetDisktypesForOrganizationOK), nil

}

/*
GetEncryptionKeysForOrganization retrives encryption keys with properties

Each cloud provider has it's own specific resources like instance types and disk types. These endpoints are collecting them.
*/
func (a *Client) GetEncryptionKeysForOrganization(params *GetEncryptionKeysForOrganizationParams) (*GetEncryptionKeysForOrganizationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetEncryptionKeysForOrganizationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getEncryptionKeysForOrganization",
		Method:             "POST",
		PathPattern:        "/v3/{organizationId}/connectors/encryptionkeys",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetEncryptionKeysForOrganizationReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetEncryptionKeysForOrganizationOK), nil

}

/*
GetGatewaysCredentialIDForOrganization retrives gateways with properties

Each cloud provider has it's own specific resources like instance types and disk types. These endpoints are collecting them.
*/
func (a *Client) GetGatewaysCredentialIDForOrganization(params *GetGatewaysCredentialIDForOrganizationParams) (*GetGatewaysCredentialIDForOrganizationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetGatewaysCredentialIDForOrganizationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getGatewaysCredentialIdForOrganization",
		Method:             "POST",
		PathPattern:        "/v3/{organizationId}/connectors/gateways",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetGatewaysCredentialIDForOrganizationReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetGatewaysCredentialIDForOrganizationOK), nil

}

/*
GetIPPoolsCredentialIDForOrganization retrives ip pools with properties

Each cloud provider has it's own specific resources like instance types and disk types. These endpoints are collecting them.
*/
func (a *Client) GetIPPoolsCredentialIDForOrganization(params *GetIPPoolsCredentialIDForOrganizationParams) (*GetIPPoolsCredentialIDForOrganizationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetIPPoolsCredentialIDForOrganizationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getIpPoolsCredentialIdForOrganization",
		Method:             "POST",
		PathPattern:        "/v3/{organizationId}/connectors/ippools",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetIPPoolsCredentialIDForOrganizationReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetIPPoolsCredentialIDForOrganizationOK), nil

}

/*
GetPlatformNetworksForOrganization retrives network properties

Each cloud provider has it's own specific resources like instance types and disk types. These endpoints are collecting them.
*/
func (a *Client) GetPlatformNetworksForOrganization(params *GetPlatformNetworksForOrganizationParams) (*GetPlatformNetworksForOrganizationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPlatformNetworksForOrganizationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getPlatformNetworksForOrganization",
		Method:             "POST",
		PathPattern:        "/v3/{organizationId}/connectors/networks",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetPlatformNetworksForOrganizationReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPlatformNetworksForOrganizationOK), nil

}

/*
GetPlatformSShKeysForOrganization retrives sshkeys properties

Each cloud provider has it's own specific resources like instance types and disk types. These endpoints are collecting them.
*/
func (a *Client) GetPlatformSShKeysForOrganization(params *GetPlatformSShKeysForOrganizationParams) (*GetPlatformSShKeysForOrganizationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPlatformSShKeysForOrganizationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getPlatformSShKeysForOrganization",
		Method:             "POST",
		PathPattern:        "/v3/{organizationId}/connectors/sshkeys",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetPlatformSShKeysForOrganizationReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPlatformSShKeysForOrganizationOK), nil

}

/*
GetPlatformSecurityGroupsForOrganization retrives securitygroups properties

Each cloud provider has it's own specific resources like instance types and disk types. These endpoints are collecting them.
*/
func (a *Client) GetPlatformSecurityGroupsForOrganization(params *GetPlatformSecurityGroupsForOrganizationParams) (*GetPlatformSecurityGroupsForOrganizationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPlatformSecurityGroupsForOrganizationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getPlatformSecurityGroupsForOrganization",
		Method:             "POST",
		PathPattern:        "/v3/{organizationId}/connectors/securitygroups",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetPlatformSecurityGroupsForOrganizationReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPlatformSecurityGroupsForOrganizationOK), nil

}

/*
GetRegionsByCredentialAndOrganization retrives regions by type

Each cloud provider has it's own specific resources like instance types and disk types. These endpoints are collecting them.
*/
func (a *Client) GetRegionsByCredentialAndOrganization(params *GetRegionsByCredentialAndOrganizationParams) (*GetRegionsByCredentialAndOrganizationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRegionsByCredentialAndOrganizationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getRegionsByCredentialAndOrganization",
		Method:             "POST",
		PathPattern:        "/v3/{organizationId}/connectors/regions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetRegionsByCredentialAndOrganizationReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetRegionsByCredentialAndOrganizationOK), nil

}

/*
GetVMTypesByCredentialAndOrganization retrives vmtype properties by credential

Each cloud provider has it's own specific resources like instance types and disk types. These endpoints are collecting them.
*/
func (a *Client) GetVMTypesByCredentialAndOrganization(params *GetVMTypesByCredentialAndOrganizationParams) (*GetVMTypesByCredentialAndOrganizationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetVMTypesByCredentialAndOrganizationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getVmTypesByCredentialAndOrganization",
		Method:             "POST",
		PathPattern:        "/v3/{organizationId}/connectors/vmtypes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetVMTypesByCredentialAndOrganizationReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetVMTypesByCredentialAndOrganizationOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
