// Code generated by go-swagger; DO NOT EDIT.

package v2stacks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new v2stacks API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for v2stacks API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DeleteInstanceStackV2 deletes instance resource from stack

Stacks are template instances - a running cloud infrastructure created based on a template. Stacks are always launched on behalf of a cloud user account. Stacks support a wide range of resources, allowing you to build a highly available, reliable, and scalable infrastructure for your application needs.
*/
func (a *Client) DeleteInstanceStackV2(params *DeleteInstanceStackV2Params) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteInstanceStackV2Params()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteInstanceStackV2",
		Method:             "DELETE",
		PathPattern:        "/v2/stacks/{stackId}/{instanceId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteInstanceStackV2Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
DeletePrivateStackV2 deletes private stack by name

Stacks are template instances - a running cloud infrastructure created based on a template. Stacks are always launched on behalf of a cloud user account. Stacks support a wide range of resources, allowing you to build a highly available, reliable, and scalable infrastructure for your application needs.
*/
func (a *Client) DeletePrivateStackV2(params *DeletePrivateStackV2Params) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeletePrivateStackV2Params()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deletePrivateStackV2",
		Method:             "DELETE",
		PathPattern:        "/v2/stacks/user/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeletePrivateStackV2Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
DeletePublicStackV2 deletes public owned or private stack by name

Stacks are template instances - a running cloud infrastructure created based on a template. Stacks are always launched on behalf of a cloud user account. Stacks support a wide range of resources, allowing you to build a highly available, reliable, and scalable infrastructure for your application needs.
*/
func (a *Client) DeletePublicStackV2(params *DeletePublicStackV2Params) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeletePublicStackV2Params()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deletePublicStackV2",
		Method:             "DELETE",
		PathPattern:        "/v2/stacks/account/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeletePublicStackV2Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
DeleteStackV2 deletes stack by id

Stacks are template instances - a running cloud infrastructure created based on a template. Stacks are always launched on behalf of a cloud user account. Stacks support a wide range of resources, allowing you to build a highly available, reliable, and scalable infrastructure for your application needs.
*/
func (a *Client) DeleteStackV2(params *DeleteStackV2Params) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteStackV2Params()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteStackV2",
		Method:             "DELETE",
		PathPattern:        "/v2/stacks/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteStackV2Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
GetAllStackV2 retrieves all stacks

Stacks are template instances - a running cloud infrastructure created based on a template. Stacks are always launched on behalf of a cloud user account. Stacks support a wide range of resources, allowing you to build a highly available, reliable, and scalable infrastructure for your application needs.
*/
func (a *Client) GetAllStackV2(params *GetAllStackV2Params) (*GetAllStackV2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAllStackV2Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getAllStackV2",
		Method:             "GET",
		PathPattern:        "/v2/stacks/all",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetAllStackV2Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetAllStackV2OK), nil

}

/*
GetCertificateStackV2 retrieves the TLS certificate used by the gateway

Stacks are template instances - a running cloud infrastructure created based on a template. Stacks are always launched on behalf of a cloud user account. Stacks support a wide range of resources, allowing you to build a highly available, reliable, and scalable infrastructure for your application needs.
*/
func (a *Client) GetCertificateStackV2(params *GetCertificateStackV2Params) (*GetCertificateStackV2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCertificateStackV2Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getCertificateStackV2",
		Method:             "GET",
		PathPattern:        "/v2/stacks/{id}/certificate",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetCertificateStackV2Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetCertificateStackV2OK), nil

}

/*
GetClusterRequestFromName retrieves stack request by stack name

Stacks are template instances - a running cloud infrastructure created based on a template. Stacks are always launched on behalf of a cloud user account. Stacks support a wide range of resources, allowing you to build a highly available, reliable, and scalable infrastructure for your application needs.
*/
func (a *Client) GetClusterRequestFromName(params *GetClusterRequestFromNameParams) (*GetClusterRequestFromNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetClusterRequestFromNameParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getClusterRequestFromName",
		Method:             "GET",
		PathPattern:        "/v2/stacks/{name}/request",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetClusterRequestFromNameReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetClusterRequestFromNameOK), nil

}

/*
GetPrivateStackV2 retrieves a private stack by name

Stacks are template instances - a running cloud infrastructure created based on a template. Stacks are always launched on behalf of a cloud user account. Stacks support a wide range of resources, allowing you to build a highly available, reliable, and scalable infrastructure for your application needs.
*/
func (a *Client) GetPrivateStackV2(params *GetPrivateStackV2Params) (*GetPrivateStackV2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPrivateStackV2Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getPrivateStackV2",
		Method:             "GET",
		PathPattern:        "/v2/stacks/user/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetPrivateStackV2Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPrivateStackV2OK), nil

}

/*
GetPrivatesStackV2 retrieves private stack

Stacks are template instances - a running cloud infrastructure created based on a template. Stacks are always launched on behalf of a cloud user account. Stacks support a wide range of resources, allowing you to build a highly available, reliable, and scalable infrastructure for your application needs.
*/
func (a *Client) GetPrivatesStackV2(params *GetPrivatesStackV2Params) (*GetPrivatesStackV2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPrivatesStackV2Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getPrivatesStackV2",
		Method:             "GET",
		PathPattern:        "/v2/stacks/user",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetPrivatesStackV2Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPrivatesStackV2OK), nil

}

/*
GetPublicStackV2 retrieves a public or private owned stack by name

Stacks are template instances - a running cloud infrastructure created based on a template. Stacks are always launched on behalf of a cloud user account. Stacks support a wide range of resources, allowing you to build a highly available, reliable, and scalable infrastructure for your application needs.
*/
func (a *Client) GetPublicStackV2(params *GetPublicStackV2Params) (*GetPublicStackV2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPublicStackV2Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getPublicStackV2",
		Method:             "GET",
		PathPattern:        "/v2/stacks/account/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetPublicStackV2Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPublicStackV2OK), nil

}

/*
GetPublicsStackV2 retrieves public and private owned stacks

Stacks are template instances - a running cloud infrastructure created based on a template. Stacks are always launched on behalf of a cloud user account. Stacks support a wide range of resources, allowing you to build a highly available, reliable, and scalable infrastructure for your application needs.
*/
func (a *Client) GetPublicsStackV2(params *GetPublicsStackV2Params) (*GetPublicsStackV2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPublicsStackV2Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getPublicsStackV2",
		Method:             "GET",
		PathPattern:        "/v2/stacks/account",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetPublicsStackV2Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPublicsStackV2OK), nil

}

/*
GetStackForAmbariV2 retrieves stack by ambari address

Stacks are template instances - a running cloud infrastructure created based on a template. Stacks are always launched on behalf of a cloud user account. Stacks support a wide range of resources, allowing you to build a highly available, reliable, and scalable infrastructure for your application needs.
*/
func (a *Client) GetStackForAmbariV2(params *GetStackForAmbariV2Params) (*GetStackForAmbariV2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetStackForAmbariV2Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getStackForAmbariV2",
		Method:             "POST",
		PathPattern:        "/v2/stacks/ambari",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetStackForAmbariV2Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetStackForAmbariV2OK), nil

}

/*
GetStackV2 retrieves stack by id

Stacks are template instances - a running cloud infrastructure created based on a template. Stacks are always launched on behalf of a cloud user account. Stacks support a wide range of resources, allowing you to build a highly available, reliable, and scalable infrastructure for your application needs.
*/
func (a *Client) GetStackV2(params *GetStackV2Params) (*GetStackV2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetStackV2Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getStackV2",
		Method:             "GET",
		PathPattern:        "/v2/stacks/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetStackV2Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetStackV2OK), nil

}

/*
PostPrivateStackV2 creates stack as private resource

Stacks are template instances - a running cloud infrastructure created based on a template. Stacks are always launched on behalf of a cloud user account. Stacks support a wide range of resources, allowing you to build a highly available, reliable, and scalable infrastructure for your application needs.
*/
func (a *Client) PostPrivateStackV2(params *PostPrivateStackV2Params) (*PostPrivateStackV2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostPrivateStackV2Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "postPrivateStackV2",
		Method:             "POST",
		PathPattern:        "/v2/stacks/user",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostPrivateStackV2Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostPrivateStackV2OK), nil

}

/*
PostPublicStackV2 creates stack as public resource

Stacks are template instances - a running cloud infrastructure created based on a template. Stacks are always launched on behalf of a cloud user account. Stacks support a wide range of resources, allowing you to build a highly available, reliable, and scalable infrastructure for your application needs.
*/
func (a *Client) PostPublicStackV2(params *PostPublicStackV2Params) (*PostPublicStackV2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostPublicStackV2Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "postPublicStackV2",
		Method:             "POST",
		PathPattern:        "/v2/stacks/account",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostPublicStackV2Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostPublicStackV2OK), nil

}

/*
PutpasswordStackV2 updates stack by name

Stacks are template instances - a running cloud infrastructure created based on a template. Stacks are always launched on behalf of a cloud user account. Stacks support a wide range of resources, allowing you to build a highly available, reliable, and scalable infrastructure for your application needs.
*/
func (a *Client) PutpasswordStackV2(params *PutpasswordStackV2Params) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutpasswordStackV2Params()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "putpasswordStackV2",
		Method:             "PUT",
		PathPattern:        "/v2/stacks/ambari_password/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PutpasswordStackV2Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
PutreinstallStackV2 updates stack by name

Stacks are template instances - a running cloud infrastructure created based on a template. Stacks are always launched on behalf of a cloud user account. Stacks support a wide range of resources, allowing you to build a highly available, reliable, and scalable infrastructure for your application needs.
*/
func (a *Client) PutreinstallStackV2(params *PutreinstallStackV2Params) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutreinstallStackV2Params()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "putreinstallStackV2",
		Method:             "PUT",
		PathPattern:        "/v2/stacks/reinstall/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PutreinstallStackV2Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
PutrepairStackV2 updates stack by name

Stacks are template instances - a running cloud infrastructure created based on a template. Stacks are always launched on behalf of a cloud user account. Stacks support a wide range of resources, allowing you to build a highly available, reliable, and scalable infrastructure for your application needs.
*/
func (a *Client) PutrepairStackV2(params *PutrepairStackV2Params) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutrepairStackV2Params()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "putrepairStackV2",
		Method:             "PUT",
		PathPattern:        "/v2/stacks/repair/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PutrepairStackV2Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
PutscalingStackV2 updates stack by name

Stacks are template instances - a running cloud infrastructure created based on a template. Stacks are always launched on behalf of a cloud user account. Stacks support a wide range of resources, allowing you to build a highly available, reliable, and scalable infrastructure for your application needs.
*/
func (a *Client) PutscalingStackV2(params *PutscalingStackV2Params) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutscalingStackV2Params()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "putscalingStackV2",
		Method:             "PUT",
		PathPattern:        "/v2/stacks/scaling/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PutscalingStackV2Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
PutstartStackV2 updates stack by name

Stacks are template instances - a running cloud infrastructure created based on a template. Stacks are always launched on behalf of a cloud user account. Stacks support a wide range of resources, allowing you to build a highly available, reliable, and scalable infrastructure for your application needs.
*/
func (a *Client) PutstartStackV2(params *PutstartStackV2Params) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutstartStackV2Params()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "putstartStackV2",
		Method:             "PUT",
		PathPattern:        "/v2/stacks/start/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PutstartStackV2Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
PutstopStackV2 updates stack by name

Stacks are template instances - a running cloud infrastructure created based on a template. Stacks are always launched on behalf of a cloud user account. Stacks support a wide range of resources, allowing you to build a highly available, reliable, and scalable infrastructure for your application needs.
*/
func (a *Client) PutstopStackV2(params *PutstopStackV2Params) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutstopStackV2Params()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "putstopStackV2",
		Method:             "PUT",
		PathPattern:        "/v2/stacks/stop/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PutstopStackV2Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
PutsyncStackV2 updates stack by name

Stacks are template instances - a running cloud infrastructure created based on a template. Stacks are always launched on behalf of a cloud user account. Stacks support a wide range of resources, allowing you to build a highly available, reliable, and scalable infrastructure for your application needs.
*/
func (a *Client) PutsyncStackV2(params *PutsyncStackV2Params) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutsyncStackV2Params()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "putsyncStackV2",
		Method:             "PUT",
		PathPattern:        "/v2/stacks/sync/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PutsyncStackV2Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
StatusStackV2 retrieves stack status by stack id

Stacks are template instances - a running cloud infrastructure created based on a template. Stacks are always launched on behalf of a cloud user account. Stacks support a wide range of resources, allowing you to build a highly available, reliable, and scalable infrastructure for your application needs.
*/
func (a *Client) StatusStackV2(params *StatusStackV2Params) (*StatusStackV2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStatusStackV2Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "statusStackV2",
		Method:             "GET",
		PathPattern:        "/v2/stacks/{id}/status",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &StatusStackV2Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*StatusStackV2OK), nil

}

/*
ValidateStackV2 validates stack

Stacks are template instances - a running cloud infrastructure created based on a template. Stacks are always launched on behalf of a cloud user account. Stacks support a wide range of resources, allowing you to build a highly available, reliable, and scalable infrastructure for your application needs.
*/
func (a *Client) ValidateStackV2(params *ValidateStackV2Params) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewValidateStackV2Params()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "validateStackV2",
		Method:             "POST",
		PathPattern:        "/v2/stacks/validate",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ValidateStackV2Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
VariantsStackV2 retrieves available platform variants

Stacks are template instances - a running cloud infrastructure created based on a template. Stacks are always launched on behalf of a cloud user account. Stacks support a wide range of resources, allowing you to build a highly available, reliable, and scalable infrastructure for your application needs.
*/
func (a *Client) VariantsStackV2(params *VariantsStackV2Params) (*VariantsStackV2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVariantsStackV2Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "variantsStackV2",
		Method:             "GET",
		PathPattern:        "/v2/stacks/platformVariants",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &VariantsStackV2Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*VariantsStackV2OK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
