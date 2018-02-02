// Code generated by go-swagger; DO NOT EDIT.

package v1stacks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new v1stacks API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for v1stacks API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DeleteCluster deletes cluster on a specific stack

Clusters are materialised Hadoop services on a given infrastructure. They are built based on a Blueprint (running the components and services specified) and on a configured infrastructure Stack. Once a cluster is created and launched, it can be used the usual way as any Hadoop cluster. We suggest to start with the Cluster's Ambari UI for an overview of your cluster.
*/
func (a *Client) DeleteCluster(params *DeleteClusterParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteClusterParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteCluster",
		Method:             "DELETE",
		PathPattern:        "/v1/stacks/{id}/cluster",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteClusterReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
DeleteInstanceStack deletes instance resource from stack

Stacks are template instances - a running cloud infrastructure created based on a template. Stacks are always launched on behalf of a cloud user account. Stacks support a wide range of resources, allowing you to build a highly available, reliable, and scalable infrastructure for your application needs.
*/
func (a *Client) DeleteInstanceStack(params *DeleteInstanceStackParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteInstanceStackParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteInstanceStack",
		Method:             "DELETE",
		PathPattern:        "/v1/stacks/{stackId}/{instanceId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteInstanceStackReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
DeletePrivateStack deletes private stack by name

Stacks are template instances - a running cloud infrastructure created based on a template. Stacks are always launched on behalf of a cloud user account. Stacks support a wide range of resources, allowing you to build a highly available, reliable, and scalable infrastructure for your application needs.
*/
func (a *Client) DeletePrivateStack(params *DeletePrivateStackParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeletePrivateStackParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deletePrivateStack",
		Method:             "DELETE",
		PathPattern:        "/v1/stacks/user/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeletePrivateStackReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
DeletePublicStack deletes public owned or private stack by name

Stacks are template instances - a running cloud infrastructure created based on a template. Stacks are always launched on behalf of a cloud user account. Stacks support a wide range of resources, allowing you to build a highly available, reliable, and scalable infrastructure for your application needs.
*/
func (a *Client) DeletePublicStack(params *DeletePublicStackParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeletePublicStackParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deletePublicStack",
		Method:             "DELETE",
		PathPattern:        "/v1/stacks/account/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeletePublicStackReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
DeleteStack deletes stack by id

Stacks are template instances - a running cloud infrastructure created based on a template. Stacks are always launched on behalf of a cloud user account. Stacks support a wide range of resources, allowing you to build a highly available, reliable, and scalable infrastructure for your application needs.
*/
func (a *Client) DeleteStack(params *DeleteStackParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteStackParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteStack",
		Method:             "DELETE",
		PathPattern:        "/v1/stacks/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteStackReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
FailureReportCluster failures report

Endpoint to report the failed nodes in the given cluster. If recovery mode for the node's hostgroup is AUTO then autorecovery would be started. If recovery mode for the node's hostgroup is MANUAL, the nodes will be marked as unhealthy.
*/
func (a *Client) FailureReportCluster(params *FailureReportClusterParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFailureReportClusterParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "failureReportCluster",
		Method:             "POST",
		PathPattern:        "/v1/stacks/{id}/cluster/failurereport",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &FailureReportClusterReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
GetAllStack retrieves all stacks

Stacks are template instances - a running cloud infrastructure created based on a template. Stacks are always launched on behalf of a cloud user account. Stacks support a wide range of resources, allowing you to build a highly available, reliable, and scalable infrastructure for your application needs.
*/
func (a *Client) GetAllStack(params *GetAllStackParams) (*GetAllStackOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAllStackParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getAllStack",
		Method:             "GET",
		PathPattern:        "/v1/stacks/all",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetAllStackReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetAllStackOK), nil

}

/*
GetCertificateStack retrieves the TLS certificate used by the gateway

Stacks are template instances - a running cloud infrastructure created based on a template. Stacks are always launched on behalf of a cloud user account. Stacks support a wide range of resources, allowing you to build a highly available, reliable, and scalable infrastructure for your application needs.
*/
func (a *Client) GetCertificateStack(params *GetCertificateStackParams) (*GetCertificateStackOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCertificateStackParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getCertificateStack",
		Method:             "GET",
		PathPattern:        "/v1/stacks/{id}/certificate",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetCertificateStackReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetCertificateStackOK), nil

}

/*
GetCluster retrieves cluster by stack id

Clusters are materialised Hadoop services on a given infrastructure. They are built based on a Blueprint (running the components and services specified) and on a configured infrastructure Stack. Once a cluster is created and launched, it can be used the usual way as any Hadoop cluster. We suggest to start with the Cluster's Ambari UI for an overview of your cluster.
*/
func (a *Client) GetCluster(params *GetClusterParams) (*GetClusterOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetClusterParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getCluster",
		Method:             "GET",
		PathPattern:        "/v1/stacks/{id}/cluster",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetClusterReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetClusterOK), nil

}

/*
GetConfigsCluster gets cluster properties with blueprint outputs

Clusters are materialised Hadoop services on a given infrastructure. They are built based on a Blueprint (running the components and services specified) and on a configured infrastructure Stack. Once a cluster is created and launched, it can be used the usual way as any Hadoop cluster. We suggest to start with the Cluster's Ambari UI for an overview of your cluster.
*/
func (a *Client) GetConfigsCluster(params *GetConfigsClusterParams) (*GetConfigsClusterOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetConfigsClusterParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getConfigsCluster",
		Method:             "POST",
		PathPattern:        "/v1/stacks/{id}/cluster/config",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetConfigsClusterReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetConfigsClusterOK), nil

}

/*
GetFullCluster retrieves cluster by stack id

Clusters are materialised Hadoop services on a given infrastructure. They are built based on a Blueprint (running the components and services specified) and on a configured infrastructure Stack. Once a cluster is created and launched, it can be used the usual way as any Hadoop cluster. We suggest to start with the Cluster's Ambari UI for an overview of your cluster.
*/
func (a *Client) GetFullCluster(params *GetFullClusterParams) (*GetFullClusterOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFullClusterParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getFullCluster",
		Method:             "GET",
		PathPattern:        "/v1/stacks/{id}/cluster/full",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetFullClusterReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetFullClusterOK), nil

}

/*
GetPrivateCluster retrieves cluster by stack name private

Clusters are materialised Hadoop services on a given infrastructure. They are built based on a Blueprint (running the components and services specified) and on a configured infrastructure Stack. Once a cluster is created and launched, it can be used the usual way as any Hadoop cluster. We suggest to start with the Cluster's Ambari UI for an overview of your cluster.
*/
func (a *Client) GetPrivateCluster(params *GetPrivateClusterParams) (*GetPrivateClusterOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPrivateClusterParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getPrivateCluster",
		Method:             "GET",
		PathPattern:        "/v1/stacks/user/{name}/cluster",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetPrivateClusterReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPrivateClusterOK), nil

}

/*
GetPrivateStack retrieves a private stack by name

Stacks are template instances - a running cloud infrastructure created based on a template. Stacks are always launched on behalf of a cloud user account. Stacks support a wide range of resources, allowing you to build a highly available, reliable, and scalable infrastructure for your application needs.
*/
func (a *Client) GetPrivateStack(params *GetPrivateStackParams) (*GetPrivateStackOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPrivateStackParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getPrivateStack",
		Method:             "GET",
		PathPattern:        "/v1/stacks/user/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetPrivateStackReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPrivateStackOK), nil

}

/*
GetPrivatesStack retrieves private stack

Stacks are template instances - a running cloud infrastructure created based on a template. Stacks are always launched on behalf of a cloud user account. Stacks support a wide range of resources, allowing you to build a highly available, reliable, and scalable infrastructure for your application needs.
*/
func (a *Client) GetPrivatesStack(params *GetPrivatesStackParams) (*GetPrivatesStackOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPrivatesStackParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getPrivatesStack",
		Method:             "GET",
		PathPattern:        "/v1/stacks/user",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetPrivatesStackReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPrivatesStackOK), nil

}

/*
GetPublicCluster retrieves cluster by stack name public

Clusters are materialised Hadoop services on a given infrastructure. They are built based on a Blueprint (running the components and services specified) and on a configured infrastructure Stack. Once a cluster is created and launched, it can be used the usual way as any Hadoop cluster. We suggest to start with the Cluster's Ambari UI for an overview of your cluster.
*/
func (a *Client) GetPublicCluster(params *GetPublicClusterParams) (*GetPublicClusterOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPublicClusterParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getPublicCluster",
		Method:             "GET",
		PathPattern:        "/v1/stacks/account/{name}/cluster",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetPublicClusterReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPublicClusterOK), nil

}

/*
GetPublicStack retrieves a public or private owned stack by name

Stacks are template instances - a running cloud infrastructure created based on a template. Stacks are always launched on behalf of a cloud user account. Stacks support a wide range of resources, allowing you to build a highly available, reliable, and scalable infrastructure for your application needs.
*/
func (a *Client) GetPublicStack(params *GetPublicStackParams) (*GetPublicStackOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPublicStackParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getPublicStack",
		Method:             "GET",
		PathPattern:        "/v1/stacks/account/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetPublicStackReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPublicStackOK), nil

}

/*
GetPublicsStack retrieves public and private owned stacks

Stacks are template instances - a running cloud infrastructure created based on a template. Stacks are always launched on behalf of a cloud user account. Stacks support a wide range of resources, allowing you to build a highly available, reliable, and scalable infrastructure for your application needs.
*/
func (a *Client) GetPublicsStack(params *GetPublicsStackParams) (*GetPublicsStackOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPublicsStackParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getPublicsStack",
		Method:             "GET",
		PathPattern:        "/v1/stacks/account",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetPublicsStackReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPublicsStackOK), nil

}

/*
GetStack retrieves stack by id

Stacks are template instances - a running cloud infrastructure created based on a template. Stacks are always launched on behalf of a cloud user account. Stacks support a wide range of resources, allowing you to build a highly available, reliable, and scalable infrastructure for your application needs.
*/
func (a *Client) GetStack(params *GetStackParams) (*GetStackOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetStackParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getStack",
		Method:             "GET",
		PathPattern:        "/v1/stacks/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetStackReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetStackOK), nil

}

/*
GetStackForAmbari retrieves stack by ambari address

Stacks are template instances - a running cloud infrastructure created based on a template. Stacks are always launched on behalf of a cloud user account. Stacks support a wide range of resources, allowing you to build a highly available, reliable, and scalable infrastructure for your application needs.
*/
func (a *Client) GetStackForAmbari(params *GetStackForAmbariParams) (*GetStackForAmbariOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetStackForAmbariParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getStackForAmbari",
		Method:             "POST",
		PathPattern:        "/v1/stacks/ambari",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetStackForAmbariReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetStackForAmbariOK), nil

}

/*
PostCluster creates cluster for stack

Clusters are materialised Hadoop services on a given infrastructure. They are built based on a Blueprint (running the components and services specified) and on a configured infrastructure Stack. Once a cluster is created and launched, it can be used the usual way as any Hadoop cluster. We suggest to start with the Cluster's Ambari UI for an overview of your cluster.
*/
func (a *Client) PostCluster(params *PostClusterParams) (*PostClusterOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostClusterParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "postCluster",
		Method:             "POST",
		PathPattern:        "/v1/stacks/{id}/cluster",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostClusterReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostClusterOK), nil

}

/*
PutCluster updates cluster by stack id

Clusters are materialised Hadoop services on a given infrastructure. They are built based on a Blueprint (running the components and services specified) and on a configured infrastructure Stack. Once a cluster is created and launched, it can be used the usual way as any Hadoop cluster. We suggest to start with the Cluster's Ambari UI for an overview of your cluster.
*/
func (a *Client) PutCluster(params *PutClusterParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutClusterParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "putCluster",
		Method:             "PUT",
		PathPattern:        "/v1/stacks/{id}/cluster",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PutClusterReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
PutStack updates stack by id

Stacks are template instances - a running cloud infrastructure created based on a template. Stacks are always launched on behalf of a cloud user account. Stacks support a wide range of resources, allowing you to build a highly available, reliable, and scalable infrastructure for your application needs.
*/
func (a *Client) PutStack(params *PutStackParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutStackParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "putStack",
		Method:             "PUT",
		PathPattern:        "/v1/stacks/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PutStackReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
RepairCluster repairs the cluster

Removing the failed nodes and starting new nodes to substitute them.
*/
func (a *Client) RepairCluster(params *RepairClusterParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRepairClusterParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "repairCluster",
		Method:             "POST",
		PathPattern:        "/v1/stacks/{id}/cluster/manualrepair",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RepairClusterReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
StatusStack retrieves stack status by stack id

Stacks are template instances - a running cloud infrastructure created based on a template. Stacks are always launched on behalf of a cloud user account. Stacks support a wide range of resources, allowing you to build a highly available, reliable, and scalable infrastructure for your application needs.
*/
func (a *Client) StatusStack(params *StatusStackParams) (*StatusStackOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStatusStackParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "statusStack",
		Method:             "GET",
		PathPattern:        "/v1/stacks/{id}/status",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &StatusStackReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*StatusStackOK), nil

}

/*
UpgradeCluster upgrades the ambari version

Ambari is used to provision the Hadoop clusters.
*/
func (a *Client) UpgradeCluster(params *UpgradeClusterParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpgradeClusterParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "upgradeCluster",
		Method:             "POST",
		PathPattern:        "/v1/stacks/{id}/cluster/upgrade",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpgradeClusterReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
ValidateStack validates stack

Stacks are template instances - a running cloud infrastructure created based on a template. Stacks are always launched on behalf of a cloud user account. Stacks support a wide range of resources, allowing you to build a highly available, reliable, and scalable infrastructure for your application needs.
*/
func (a *Client) ValidateStack(params *ValidateStackParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewValidateStackParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "validateStack",
		Method:             "POST",
		PathPattern:        "/v1/stacks/validate",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ValidateStackReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
VariantsStack retrieves available platform variants

Stacks are template instances - a running cloud infrastructure created based on a template. Stacks are always launched on behalf of a cloud user account. Stacks support a wide range of resources, allowing you to build a highly available, reliable, and scalable infrastructure for your application needs.
*/
func (a *Client) VariantsStack(params *VariantsStackParams) (*VariantsStackOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVariantsStackParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "variantsStack",
		Method:             "GET",
		PathPattern:        "/v1/stacks/platformVariants",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &VariantsStackReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*VariantsStackOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
