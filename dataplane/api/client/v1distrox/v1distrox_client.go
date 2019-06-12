// Code generated by go-swagger; DO NOT EDIT.

package v1distrox

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new v1distrox API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for v1distrox API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
ChangeImageDistroXV1 checks image in stack by name

Stacks are template instances - a running cloud infrastructure created based on a template. Stacks are always launched on behalf of a cloud user account. Stacks support a wide range of resources, allowing you to build a highly available, reliable, and scalable infrastructure for your application needs.
*/
func (a *Client) ChangeImageDistroXV1(params *ChangeImageDistroXV1Params) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewChangeImageDistroXV1Params()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "changeImageDistroXV1",
		Method:             "PUT",
		PathPattern:        "/v1/distrox/{name}/change_image",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ChangeImageDistroXV1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
DeleteDistroXV1 deletes an workspace by name

Stacks are template instances - a running cloud infrastructure created based on a template. Stacks are always launched on behalf of a cloud user account. Stacks support a wide range of resources, allowing you to build a highly available, reliable, and scalable infrastructure for your application needs.
*/
func (a *Client) DeleteDistroXV1(params *DeleteDistroXV1Params) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteDistroXV1Params()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteDistroXV1",
		Method:             "DELETE",
		PathPattern:        "/v1/distrox/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteDistroXV1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
DeleteInstanceDistroXV1 deletes instance from the stack s cluster

Stacks are template instances - a running cloud infrastructure created based on a template. Stacks are always launched on behalf of a cloud user account. Stacks support a wide range of resources, allowing you to build a highly available, reliable, and scalable infrastructure for your application needs.
*/
func (a *Client) DeleteInstanceDistroXV1(params *DeleteInstanceDistroXV1Params) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteInstanceDistroXV1Params()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteInstanceDistroXV1",
		Method:             "DELETE",
		PathPattern:        "/v1/distrox/{name}/instance",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteInstanceDistroXV1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
DeleteWithKerberosDistroXV1 deletes the stack with kerberos cluster by name

Clusters are materialised Hadoop services on a given infrastructure. They are built based on a blueprint (running the components and services specified) and on a configured infrastructure Stack. Once a cluster is created and launched, it can be used the usual way as any Hadoop cluster. We suggest to start with the Cluster's Ambari UI for an overview of your cluster.
*/
func (a *Client) DeleteWithKerberosDistroXV1(params *DeleteWithKerberosDistroXV1Params) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteWithKerberosDistroXV1Params()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteWithKerberosDistroXV1",
		Method:             "DELETE",
		PathPattern:        "/v1/distrox/{name}/cluster",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteWithKerberosDistroXV1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
GetDistroXRequestFromNameV1 gets stack request by name

Stacks are template instances - a running cloud infrastructure created based on a template. Stacks are always launched on behalf of a cloud user account. Stacks support a wide range of resources, allowing you to build a highly available, reliable, and scalable infrastructure for your application needs.
*/
func (a *Client) GetDistroXRequestFromNameV1(params *GetDistroXRequestFromNameV1Params) (*GetDistroXRequestFromNameV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDistroXRequestFromNameV1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getDistroXRequestFromNameV1",
		Method:             "GET",
		PathPattern:        "/v1/distrox/{name}/request",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetDistroXRequestFromNameV1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetDistroXRequestFromNameV1OK), nil

}

/*
GetDistroXV1 gets stack by name

Stacks are template instances - a running cloud infrastructure created based on a template. Stacks are always launched on behalf of a cloud user account. Stacks support a wide range of resources, allowing you to build a highly available, reliable, and scalable infrastructure for your application needs.
*/
func (a *Client) GetDistroXV1(params *GetDistroXV1Params) (*GetDistroXV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDistroXV1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getDistroXV1",
		Method:             "GET",
		PathPattern:        "/v1/distrox/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetDistroXV1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetDistroXV1OK), nil

}

/*
ListDistroXV1 lists stacks

Stacks are template instances - a running cloud infrastructure created based on a template. Stacks are always launched on behalf of a cloud user account. Stacks support a wide range of resources, allowing you to build a highly available, reliable, and scalable infrastructure for your application needs.
*/
func (a *Client) ListDistroXV1(params *ListDistroXV1Params) (*ListDistroXV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListDistroXV1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listDistroXV1",
		Method:             "GET",
		PathPattern:        "/v1/distrox",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListDistroXV1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListDistroXV1OK), nil

}

/*
PostDistroXForBlueprintV1 posts stack for blueprint

Stacks are template instances - a running cloud infrastructure created based on a template. Stacks are always launched on behalf of a cloud user account. Stacks support a wide range of resources, allowing you to build a highly available, reliable, and scalable infrastructure for your application needs.
*/
func (a *Client) PostDistroXForBlueprintV1(params *PostDistroXForBlueprintV1Params) (*PostDistroXForBlueprintV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostDistroXForBlueprintV1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "postDistroXForBlueprintV1",
		Method:             "POST",
		PathPattern:        "/v1/distrox/{name}/blueprint",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostDistroXForBlueprintV1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostDistroXForBlueprintV1OK), nil

}

/*
PostDistroXV1 creates stack

Stacks are template instances - a running cloud infrastructure created based on a template. Stacks are always launched on behalf of a cloud user account. Stacks support a wide range of resources, allowing you to build a highly available, reliable, and scalable infrastructure for your application needs.
*/
func (a *Client) PostDistroXV1(params *PostDistroXV1Params) (*PostDistroXV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostDistroXV1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "postDistroXV1",
		Method:             "POST",
		PathPattern:        "/v1/distrox",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostDistroXV1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostDistroXV1OK), nil

}

/*
PutScalingDistroXV1 scales the stack by name

Stacks are template instances - a running cloud infrastructure created based on a template. Stacks are always launched on behalf of a cloud user account. Stacks support a wide range of resources, allowing you to build a highly available, reliable, and scalable infrastructure for your application needs.
*/
func (a *Client) PutScalingDistroXV1(params *PutScalingDistroXV1Params) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutScalingDistroXV1Params()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "putScalingDistroXV1",
		Method:             "PUT",
		PathPattern:        "/v1/distrox/{name}/scaling",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PutScalingDistroXV1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
RepairDistroXV1 repairs the stack by name

Removing the failed nodes and starting new nodes to substitute them.
*/
func (a *Client) RepairDistroXV1(params *RepairDistroXV1Params) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRepairDistroXV1Params()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "repairDistroXV1",
		Method:             "POST",
		PathPattern:        "/v1/distrox/{name}/manual_repair",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RepairDistroXV1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
RetryDistroXV1 retries the stack by name

Failed or interrupted stack and cluster operations can be retried, after the cause of the failure was eliminated. The operations will continue at the state, where the previous process failed.
*/
func (a *Client) RetryDistroXV1(params *RetryDistroXV1Params) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRetryDistroXV1Params()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "retryDistroXV1",
		Method:             "PUT",
		PathPattern:        "/v1/distrox/{name}/retry",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RetryDistroXV1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
SetDistroXMaintenanceMode sets maintenance mode for the cluster

Setting maintenance mode for the cluster in order to be able to update Ambari and/or the Hadoop stack.
*/
func (a *Client) SetDistroXMaintenanceMode(params *SetDistroXMaintenanceModeParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSetDistroXMaintenanceModeParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "setDistroXMaintenanceMode",
		Method:             "PUT",
		PathPattern:        "/v1/distrox/{name}/maintenance",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SetDistroXMaintenanceModeReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
StartDistroXV1 starts the stack by name

Stacks are template instances - a running cloud infrastructure created based on a template. Stacks are always launched on behalf of a cloud user account. Stacks support a wide range of resources, allowing you to build a highly available, reliable, and scalable infrastructure for your application needs.
*/
func (a *Client) StartDistroXV1(params *StartDistroXV1Params) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStartDistroXV1Params()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "startDistroXV1",
		Method:             "PUT",
		PathPattern:        "/v1/distrox/{name}/start",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &StartDistroXV1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
StatusDistroXV1 retrieves stack status by stack name

Stacks are template instances - a running cloud infrastructure created based on a template. Stacks are always launched on behalf of a cloud user account. Stacks support a wide range of resources, allowing you to build a highly available, reliable, and scalable infrastructure for your application needs.
*/
func (a *Client) StatusDistroXV1(params *StatusDistroXV1Params) (*StatusDistroXV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStatusDistroXV1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "statusDistroXV1",
		Method:             "GET",
		PathPattern:        "/v1/distrox/{name}/status",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &StatusDistroXV1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*StatusDistroXV1OK), nil

}

/*
StopDistroXV1 stops the stack by name

Stacks are template instances - a running cloud infrastructure created based on a template. Stacks are always launched on behalf of a cloud user account. Stacks support a wide range of resources, allowing you to build a highly available, reliable, and scalable infrastructure for your application needs.
*/
func (a *Client) StopDistroXV1(params *StopDistroXV1Params) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStopDistroXV1Params()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "stopDistroXV1",
		Method:             "PUT",
		PathPattern:        "/v1/distrox/{name}/stop",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &StopDistroXV1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
SyncDistroXV1 syncs the stack by name

Stacks are template instances - a running cloud infrastructure created based on a template. Stacks are always launched on behalf of a cloud user account. Stacks support a wide range of resources, allowing you to build a highly available, reliable, and scalable infrastructure for your application needs.
*/
func (a *Client) SyncDistroXV1(params *SyncDistroXV1Params) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSyncDistroXV1Params()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "syncDistroXV1",
		Method:             "PUT",
		PathPattern:        "/v1/distrox/{name}/sync",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SyncDistroXV1Reader{formats: a.formats},
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
