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
func (a *Client) AttachChildEnvironmentV1(params *AttachChildEnvironmentV1Params) error {
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
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
CleanupV1 cleans out users hosts and related DNS entries

FreeIPA is an integrated Identity and Authentication solution that can be used for any of CM, CDP services.
*/
func (a *Client) CleanupV1(params *CleanupV1Params) (*CleanupV1OK, error) {
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
func (a *Client) ClusterProxyDeregisterV1(params *ClusterProxyDeregisterV1Params) error {
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
func (a *Client) ClusterProxyRegisterV1(params *ClusterProxyRegisterV1Params) (*ClusterProxyRegisterV1OK, error) {
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
func (a *Client) CreateBindUserV1(params *CreateBindUserV1Params) (*CreateBindUserV1OK, error) {
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
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateBindUserV1OK), nil

}

/*
CreateFreeIpaV1 creates free ipa stack

FreeIPA is an integrated Identity and Authentication solution that can be used for any of CM, CDP services.
*/
func (a *Client) CreateFreeIpaV1(params *CreateFreeIpaV1Params) (*CreateFreeIpaV1OK, error) {
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
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateFreeIpaV1OK), nil

}

/*
DeleteFreeIpaByEnvironmentV1 deletes free IP a stack by envid

FreeIPA is an integrated Identity and Authentication solution that can be used for any of CM, CDP services.
*/
func (a *Client) DeleteFreeIpaByEnvironmentV1(params *DeleteFreeIpaByEnvironmentV1Params) error {
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
func (a *Client) DetachChildEnvironmentV1(params *DetachChildEnvironmentV1Params) error {
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
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
GetFreeIpaByEnvironmentV1 gets free IP a stack by envid

FreeIPA is an integrated Identity and Authentication solution that can be used for any of CM, CDP services.
*/
func (a *Client) GetFreeIpaByEnvironmentV1(params *GetFreeIpaByEnvironmentV1Params) (*GetFreeIpaByEnvironmentV1OK, error) {
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
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetFreeIpaByEnvironmentV1OK), nil

}

/*
GetFreeIpaRootCertificateByEnvironmentV1 gets free IP a root certificate by envid

FreeIPA is an integrated Identity and Authentication solution that can be used for any of CM, CDP services.
*/
func (a *Client) GetFreeIpaRootCertificateByEnvironmentV1(params *GetFreeIpaRootCertificateByEnvironmentV1Params) (*GetFreeIpaRootCertificateByEnvironmentV1OK, error) {
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
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetFreeIpaRootCertificateByEnvironmentV1OK), nil

}

/*
HealthV1 provides a detailed health of the free IP a stack

FreeIPA is an integrated Identity and Authentication solution that can be used for any of CM, CDP services.
*/
func (a *Client) HealthV1(params *HealthV1Params) (*HealthV1OK, error) {
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
func (a *Client) InternalCleanupV1(params *InternalCleanupV1Params) (*InternalCleanupV1OK, error) {
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
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*InternalCleanupV1OK), nil

}

/*
InternalGetFreeIpaByEnvironmentV1 gets free IP a stack by envid and account id

FreeIPA is an integrated Identity and Authentication solution that can be used for any of CM, CDP services.
*/
func (a *Client) InternalGetFreeIpaByEnvironmentV1(params *InternalGetFreeIpaByEnvironmentV1Params) (*InternalGetFreeIpaByEnvironmentV1OK, error) {
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
func (a *Client) InternalListFreeIpaClustersByAccountV1(params *InternalListFreeIpaClustersByAccountV1Params) (*InternalListFreeIpaClustersByAccountV1OK, error) {
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
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*InternalListFreeIpaClustersByAccountV1OK), nil

}

/*
ListFreeIpaClustersByAccountV1 lists all free IP a stacks by account

FreeIPA is an integrated Identity and Authentication solution that can be used for any of CM, CDP services.
*/
func (a *Client) ListFreeIpaClustersByAccountV1(params *ListFreeIpaClustersByAccountV1Params) (*ListFreeIpaClustersByAccountV1OK, error) {
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
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListFreeIpaClustersByAccountV1OK), nil

}

/*
RebootV1 reboots one or more instances

FreeIPA is an integrated Identity and Authentication solution that can be used for any of CM, CDP services.
*/
func (a *Client) RebootV1(params *RebootV1Params) (*RebootV1OK, error) {
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
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*RebootV1OK), nil

}

/*
RepairV1 repairs one or more instances

FreeIPA is an integrated Identity and Authentication solution that can be used for any of CM, CDP services.
*/
func (a *Client) RepairV1(params *RepairV1Params) (*RepairV1OK, error) {
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
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*RepairV1OK), nil

}

/*
StartV1 starts all free IP a stacks that attached to the given environment c r n

FreeIPA is an integrated Identity and Authentication solution that can be used for any of CM, CDP services.
*/
func (a *Client) StartV1(params *StartV1Params) error {
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
func (a *Client) StopV1(params *StopV1Params) error {
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
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
