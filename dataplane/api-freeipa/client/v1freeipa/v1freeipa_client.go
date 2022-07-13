// Code generated by go-swagger; DO NOT EDIT.

package v1freeipa

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new v1freeipa API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for v1freeipa API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
AttachChildEnvironmentV1 registers a child environment

FreeIPA is an integrated Identity and Authentication solution that can be used for any of CM, CDP services.
*/
func (a *Client) AttachChildEnvironmentV1(params *AttachChildEnvironmentV1Params, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAttachChildEnvironmentV1Params()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "attachChildEnvironmentV1",
		Method:             "POST",
		PathPattern:        "/v1/freeipa/attach_child_environment",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AttachChildEnvironmentV1Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
ChangeImageCatalog changes the image catalog used for creating instances

FreeIPA is an integrated Identity and Authentication solution that can be used for any of CM, CDP services.
*/
func (a *Client) ChangeImageCatalog(params *ChangeImageCatalogParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewChangeImageCatalogParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "changeImageCatalog",
		Method:             "PUT",
		PathPattern:        "/v1/freeipa/change_image_catalog",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ChangeImageCatalogReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
ChangeImageV1 changes the image used for creating instances

FreeIPA is an integrated Identity and Authentication solution that can be used for any of CM, CDP services.
*/
func (a *Client) ChangeImageV1(params *ChangeImageV1Params, authInfo runtime.ClientAuthInfoWriter) (*ChangeImageV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewChangeImageV1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "changeImageV1",
		Method:             "PUT",
		PathPattern:        "/v1/freeipa/change_image",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ChangeImageV1Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ChangeImageV1OK), nil

}

/*
CleanupV1 cleans out users hosts and related DNS entries

FreeIPA is an integrated Identity and Authentication solution that can be used for any of CM, CDP services.
*/
func (a *Client) CleanupV1(params *CleanupV1Params, authInfo runtime.ClientAuthInfoWriter) (*CleanupV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCleanupV1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "cleanupV1",
		Method:             "POST",
		PathPattern:        "/v1/freeipa/cleanup",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CleanupV1Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CleanupV1OK), nil

}

/*
ClusterProxyDeregisterV1 deregisters free IP a stack with given environment c r n with cluster proxy

FreeIPA is an integrated Identity and Authentication solution that can be used for any of CM, CDP services.
*/
func (a *Client) ClusterProxyDeregisterV1(params *ClusterProxyDeregisterV1Params, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewClusterProxyDeregisterV1Params()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "clusterProxyDeregisterV1",
		Method:             "POST",
		PathPattern:        "/v1/freeipa/cluster-proxy/deregister",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ClusterProxyDeregisterV1Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
ClusterProxyRegisterV1 registers free IP a stack with given environment c r n with cluster proxy

FreeIPA is an integrated Identity and Authentication solution that can be used for any of CM, CDP services.
*/
func (a *Client) ClusterProxyRegisterV1(params *ClusterProxyRegisterV1Params, authInfo runtime.ClientAuthInfoWriter) (*ClusterProxyRegisterV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewClusterProxyRegisterV1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "clusterProxyRegisterV1",
		Method:             "POST",
		PathPattern:        "/v1/freeipa/cluster-proxy/register",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ClusterProxyRegisterV1Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ClusterProxyRegisterV1OK), nil

}

/*
CreateBindUserV1 creates kerberos and ldap bind users for cluster

FreeIPA is an integrated Identity and Authentication solution that can be used for any of CM, CDP services.
*/
func (a *Client) CreateBindUserV1(params *CreateBindUserV1Params, authInfo runtime.ClientAuthInfoWriter) (*CreateBindUserV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateBindUserV1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createBindUserV1",
		Method:             "POST",
		PathPattern:        "/v1/freeipa/binduser/create",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateBindUserV1Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateBindUserV1OK), nil

}

/*
CreateE2ETestBindUserV1 creates kerberos and ldap bind users for cluster

FreeIPA is an integrated Identity and Authentication solution that can be used for any of CM, CDP services.
*/
func (a *Client) CreateE2ETestBindUserV1(params *CreateE2ETestBindUserV1Params, authInfo runtime.ClientAuthInfoWriter) (*CreateE2ETestBindUserV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateE2ETestBindUserV1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createE2ETestBindUserV1",
		Method:             "POST",
		PathPattern:        "/v1/freeipa/binduser/create/e2etest",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateE2ETestBindUserV1Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateE2ETestBindUserV1OK), nil

}

/*
CreateFreeIpaV1 creates free IP a stack

FreeIPA is an integrated Identity and Authentication solution that can be used for any of CM, CDP services.
*/
func (a *Client) CreateFreeIpaV1(params *CreateFreeIpaV1Params, authInfo runtime.ClientAuthInfoWriter) (*CreateFreeIpaV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateFreeIpaV1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createFreeIpaV1",
		Method:             "POST",
		PathPattern:        "/v1/freeipa",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateFreeIpaV1Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateFreeIpaV1OK), nil

}

/*
DeleteFreeIpaByEnvironmentV1 deletes free IP a stack by environment c r n

FreeIPA is an integrated Identity and Authentication solution that can be used for any of CM, CDP services.
*/
func (a *Client) DeleteFreeIpaByEnvironmentV1(params *DeleteFreeIpaByEnvironmentV1Params, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteFreeIpaByEnvironmentV1Params()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteFreeIpaByEnvironmentV1",
		Method:             "DELETE",
		PathPattern:        "/v1/freeipa",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteFreeIpaByEnvironmentV1Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
DetachChildEnvironmentV1 deregisters a child environment

FreeIPA is an integrated Identity and Authentication solution that can be used for any of CM, CDP services.
*/
func (a *Client) DetachChildEnvironmentV1(params *DetachChildEnvironmentV1Params, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDetachChildEnvironmentV1Params()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "detachChildEnvironmentV1",
		Method:             "POST",
		PathPattern:        "/v1/freeipa/detach_child_environment",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DetachChildEnvironmentV1Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
DownscaleFreeIpaV1 downscales free IP a instances

FreeIPA is an integrated Identity and Authentication solution that can be used for any of CM, CDP services.
*/
func (a *Client) DownscaleFreeIpaV1(params *DownscaleFreeIpaV1Params, authInfo runtime.ClientAuthInfoWriter) (*DownscaleFreeIpaV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDownscaleFreeIpaV1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "downscaleFreeIpaV1",
		Method:             "PUT",
		PathPattern:        "/v1/freeipa/downscale",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DownscaleFreeIpaV1Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DownscaleFreeIpaV1OK), nil

}

/*
GenerateImageCatalog generates an image catalog that only contains the currently used image for creating instances

FreeIPA is an integrated Identity and Authentication solution that can be used for any of CM, CDP services.
*/
func (a *Client) GenerateImageCatalog(params *GenerateImageCatalogParams, authInfo runtime.ClientAuthInfoWriter) (*GenerateImageCatalogOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGenerateImageCatalogParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "generateImageCatalog",
		Method:             "GET",
		PathPattern:        "/v1/freeipa/generate_image_catalog",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GenerateImageCatalogReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GenerateImageCatalogOK), nil

}

/*
GetAllFreeIpaByEnvironmentV1 gets all free IP a stacks by environment c r n

FreeIPA is an integrated Identity and Authentication solution that can be used for any of CM, CDP services.
*/
func (a *Client) GetAllFreeIpaByEnvironmentV1(params *GetAllFreeIpaByEnvironmentV1Params, authInfo runtime.ClientAuthInfoWriter) (*GetAllFreeIpaByEnvironmentV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAllFreeIpaByEnvironmentV1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getAllFreeIpaByEnvironmentV1",
		Method:             "GET",
		PathPattern:        "/v1/freeipa/all",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetAllFreeIpaByEnvironmentV1Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetAllFreeIpaByEnvironmentV1OK), nil

}

/*
GetFreeIpaByEnvironmentV1 gets free IP a stack by environment c r n

FreeIPA is an integrated Identity and Authentication solution that can be used for any of CM, CDP services.
*/
func (a *Client) GetFreeIpaByEnvironmentV1(params *GetFreeIpaByEnvironmentV1Params, authInfo runtime.ClientAuthInfoWriter) (*GetFreeIpaByEnvironmentV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFreeIpaByEnvironmentV1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getFreeIpaByEnvironmentV1",
		Method:             "GET",
		PathPattern:        "/v1/freeipa",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetFreeIpaByEnvironmentV1Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetFreeIpaByEnvironmentV1OK), nil

}

/*
GetFreeIpaRootCertificateByEnvironmentV1 gets free IP a root certificate by environment c r n

FreeIPA is an integrated Identity and Authentication solution that can be used for any of CM, CDP services.
*/
func (a *Client) GetFreeIpaRootCertificateByEnvironmentV1(params *GetFreeIpaRootCertificateByEnvironmentV1Params, authInfo runtime.ClientAuthInfoWriter) (*GetFreeIpaRootCertificateByEnvironmentV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFreeIpaRootCertificateByEnvironmentV1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getFreeIpaRootCertificateByEnvironmentV1",
		Method:             "GET",
		PathPattern:        "/v1/freeipa/ca.crt",
		ProducesMediaTypes: []string{"text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetFreeIpaRootCertificateByEnvironmentV1Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetFreeIpaRootCertificateByEnvironmentV1OK), nil

}

/*
GetFreeIpaUpgradeOptionsV1 gets available images for free IP a upgrade if catalog is defined use the catalog as image source

FreeIPA is an integrated Identity and Authentication solution that can be used for any of CM, CDP services.
*/
func (a *Client) GetFreeIpaUpgradeOptionsV1(params *GetFreeIpaUpgradeOptionsV1Params) (*GetFreeIpaUpgradeOptionsV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFreeIpaUpgradeOptionsV1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getFreeIpaUpgradeOptionsV1",
		Method:             "GET",
		PathPattern:        "/v1/freeipa/upgrade/options",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetFreeIpaUpgradeOptionsV1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetFreeIpaUpgradeOptionsV1OK), nil

}

/*
GetRecommendationV1 gets recommendation that advises cloud resources for free IP a based on the given credential c r n

FreeIPA is an integrated Identity and Authentication solution that can be used for any of CM, CDP services.
*/
func (a *Client) GetRecommendationV1(params *GetRecommendationV1Params, authInfo runtime.ClientAuthInfoWriter) (*GetRecommendationV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRecommendationV1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getRecommendationV1",
		Method:             "GET",
		PathPattern:        "/v1/freeipa/get_recommendation",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetRecommendationV1Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetRecommendationV1OK), nil

}

/*
HealthV1 provides a detailed health of the free IP a stack

FreeIPA is an integrated Identity and Authentication solution that can be used for any of CM, CDP services.
*/
func (a *Client) HealthV1(params *HealthV1Params, authInfo runtime.ClientAuthInfoWriter) (*HealthV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewHealthV1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "healthV1",
		Method:             "GET",
		PathPattern:        "/v1/freeipa/health",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &HealthV1Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*HealthV1OK), nil

}

/*
InternalCleanupV1 cleans out users hosts and related DNS entries using internal actor

FreeIPA is an integrated Identity and Authentication solution that can be used for any of CM, CDP services.
*/
func (a *Client) InternalCleanupV1(params *InternalCleanupV1Params, authInfo runtime.ClientAuthInfoWriter) (*InternalCleanupV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewInternalCleanupV1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "internalCleanupV1",
		Method:             "POST",
		PathPattern:        "/v1/freeipa/internal_cleanup",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &InternalCleanupV1Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*InternalCleanupV1OK), nil

}

/*
InternalGetAllFreeIpaByEnvironmentV1 gets all free IP a stacks by environment c r n and account ID using the internal actor

FreeIPA is an integrated Identity and Authentication solution that can be used for any of CM, CDP services.
*/
func (a *Client) InternalGetAllFreeIpaByEnvironmentV1(params *InternalGetAllFreeIpaByEnvironmentV1Params, authInfo runtime.ClientAuthInfoWriter) (*InternalGetAllFreeIpaByEnvironmentV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewInternalGetAllFreeIpaByEnvironmentV1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "internalGetAllFreeIpaByEnvironmentV1",
		Method:             "GET",
		PathPattern:        "/v1/freeipa/internal/all",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &InternalGetAllFreeIpaByEnvironmentV1Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*InternalGetAllFreeIpaByEnvironmentV1OK), nil

}

/*
InternalGetFreeIpaByEnvironmentV1 gets free IP a stack by environment c r n and account ID using the internal actor

FreeIPA is an integrated Identity and Authentication solution that can be used for any of CM, CDP services.
*/
func (a *Client) InternalGetFreeIpaByEnvironmentV1(params *InternalGetFreeIpaByEnvironmentV1Params, authInfo runtime.ClientAuthInfoWriter) (*InternalGetFreeIpaByEnvironmentV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewInternalGetFreeIpaByEnvironmentV1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "internalGetFreeIpaByEnvironmentV1",
		Method:             "GET",
		PathPattern:        "/v1/freeipa/internal",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &InternalGetFreeIpaByEnvironmentV1Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*InternalGetFreeIpaByEnvironmentV1OK), nil

}

/*
InternalListFreeIpaClustersByAccountV1 lists all free IP a stacks by account using the internal actor

FreeIPA is an integrated Identity and Authentication solution that can be used for any of CM, CDP services.
*/
func (a *Client) InternalListFreeIpaClustersByAccountV1(params *InternalListFreeIpaClustersByAccountV1Params, authInfo runtime.ClientAuthInfoWriter) (*InternalListFreeIpaClustersByAccountV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewInternalListFreeIpaClustersByAccountV1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "internalListFreeIpaClustersByAccountV1",
		Method:             "GET",
		PathPattern:        "/v1/freeipa/internal/list",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &InternalListFreeIpaClustersByAccountV1Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*InternalListFreeIpaClustersByAccountV1OK), nil

}

/*
InternalUpgradeCcmByEnvironmentV1 initiates the c c m tunnel type upgrade to the latest available version for free IP a stack by environment c r n using the internal actor

FreeIPA is an integrated Identity and Authentication solution that can be used for any of CM, CDP services.
*/
func (a *Client) InternalUpgradeCcmByEnvironmentV1(params *InternalUpgradeCcmByEnvironmentV1Params, authInfo runtime.ClientAuthInfoWriter) (*InternalUpgradeCcmByEnvironmentV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewInternalUpgradeCcmByEnvironmentV1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "internalUpgradeCcmByEnvironmentV1",
		Method:             "PUT",
		PathPattern:        "/v1/freeipa/internal/upgrade_ccm",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &InternalUpgradeCcmByEnvironmentV1Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*InternalUpgradeCcmByEnvironmentV1OK), nil

}

/*
ListFreeIpaClustersByAccountV1 lists all free IP a stacks by account

FreeIPA is an integrated Identity and Authentication solution that can be used for any of CM, CDP services.
*/
func (a *Client) ListFreeIpaClustersByAccountV1(params *ListFreeIpaClustersByAccountV1Params, authInfo runtime.ClientAuthInfoWriter) (*ListFreeIpaClustersByAccountV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListFreeIpaClustersByAccountV1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listFreeIpaClustersByAccountV1",
		Method:             "GET",
		PathPattern:        "/v1/freeipa/list",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListFreeIpaClustersByAccountV1Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListFreeIpaClustersByAccountV1OK), nil

}

/*
ListRetryableFlowsV1 lists retryable failed flows

FreeIPA is an integrated Identity and Authentication solution that can be used for any of CM, CDP services.
*/
func (a *Client) ListRetryableFlowsV1(params *ListRetryableFlowsV1Params, authInfo runtime.ClientAuthInfoWriter) (*ListRetryableFlowsV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListRetryableFlowsV1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listRetryableFlowsV1",
		Method:             "GET",
		PathPattern:        "/v1/freeipa/retry",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListRetryableFlowsV1Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListRetryableFlowsV1OK), nil

}

/*
RebootV1 reboots one or more instances

FreeIPA is an integrated Identity and Authentication solution that can be used for any of CM, CDP services.
*/
func (a *Client) RebootV1(params *RebootV1Params, authInfo runtime.ClientAuthInfoWriter) (*RebootV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRebootV1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "rebootV1",
		Method:             "POST",
		PathPattern:        "/v1/freeipa/reboot",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RebootV1Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*RebootV1OK), nil

}

/*
RebuildV1 rebuilds the free IP a cluster

FreeIPA is an integrated Identity and Authentication solution that can be used for any of CM, CDP services.
*/
func (a *Client) RebuildV1(params *RebuildV1Params, authInfo runtime.ClientAuthInfoWriter) (*RebuildV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRebuildV1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "rebuildV1",
		Method:             "POST",
		PathPattern:        "/v1/freeipa/rebuild",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RebuildV1Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*RebuildV1OK), nil

}

/*
RepairV1 repairs one or more instances

FreeIPA is an integrated Identity and Authentication solution that can be used for any of CM, CDP services.
*/
func (a *Client) RepairV1(params *RepairV1Params, authInfo runtime.ClientAuthInfoWriter) (*RepairV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRepairV1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "repairV1",
		Method:             "POST",
		PathPattern:        "/v1/freeipa/repair",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RepairV1Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*RepairV1OK), nil

}

/*
RetryV1 retries the latest failed operation

FreeIPA is an integrated Identity and Authentication solution that can be used for any of CM, CDP services.
*/
func (a *Client) RetryV1(params *RetryV1Params, authInfo runtime.ClientAuthInfoWriter) (*RetryV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRetryV1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "retryV1",
		Method:             "PUT",
		PathPattern:        "/v1/freeipa/retry",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RetryV1Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*RetryV1OK), nil

}

/*
RotateSaltPasswordV1 rotates salt stack user password of free IP a stacks that attached to the given environment c r n

FreeIPA is an integrated Identity and Authentication solution that can be used for any of CM, CDP services.
*/
func (a *Client) RotateSaltPasswordV1(params *RotateSaltPasswordV1Params, authInfo runtime.ClientAuthInfoWriter) (*RotateSaltPasswordV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRotateSaltPasswordV1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "rotateSaltPasswordV1",
		Method:             "POST",
		PathPattern:        "/v1/freeipa/rotate_salt_password",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RotateSaltPasswordV1Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*RotateSaltPasswordV1OK), nil

}

/*
StartV1 starts all free IP a stacks that attached to the given environment c r n

FreeIPA is an integrated Identity and Authentication solution that can be used for any of CM, CDP services.
*/
func (a *Client) StartV1(params *StartV1Params, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStartV1Params()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "startV1",
		Method:             "PUT",
		PathPattern:        "/v1/freeipa/start",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &StartV1Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
StopV1 stops all free IP a stacks that attached to the given environment c r n

FreeIPA is an integrated Identity and Authentication solution that can be used for any of CM, CDP services.
*/
func (a *Client) StopV1(params *StopV1Params, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStopV1Params()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "stopV1",
		Method:             "PUT",
		PathPattern:        "/v1/freeipa/stop",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &StopV1Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
UpdateSaltV1 updates salt states on free IP a instances

FreeIPA is an integrated Identity and Authentication solution that can be used for any of CM, CDP services.
*/
func (a *Client) UpdateSaltV1(params *UpdateSaltV1Params, authInfo runtime.ClientAuthInfoWriter) (*UpdateSaltV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateSaltV1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateSaltV1",
		Method:             "PUT",
		PathPattern:        "/v1/freeipa/salt_update",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateSaltV1Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateSaltV1OK), nil

}

/*
UpgradeFreeIpaV1 upgrades free IP a to the latest or defined image

FreeIPA is an integrated Identity and Authentication solution that can be used for any of CM, CDP services.
*/
func (a *Client) UpgradeFreeIpaV1(params *UpgradeFreeIpaV1Params) (*UpgradeFreeIpaV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpgradeFreeIpaV1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "upgradeFreeIpaV1",
		Method:             "POST",
		PathPattern:        "/v1/freeipa/upgrade",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpgradeFreeIpaV1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpgradeFreeIpaV1OK), nil

}

/*
UpscaleFreeIpaV1 upscales free IP a instances

FreeIPA is an integrated Identity and Authentication solution that can be used for any of CM, CDP services.
*/
func (a *Client) UpscaleFreeIpaV1(params *UpscaleFreeIpaV1Params, authInfo runtime.ClientAuthInfoWriter) (*UpscaleFreeIpaV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpscaleFreeIpaV1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "upscaleFreeIpaV1",
		Method:             "PUT",
		PathPattern:        "/v1/freeipa/upscale",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpscaleFreeIpaV1Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpscaleFreeIpaV1OK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
