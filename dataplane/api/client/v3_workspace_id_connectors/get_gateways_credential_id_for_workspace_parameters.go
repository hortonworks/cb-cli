// Code generated by go-swagger; DO NOT EDIT.

package v3_workspace_id_connectors

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api/model"
)

// NewGetGatewaysCredentialIDForWorkspaceParams creates a new GetGatewaysCredentialIDForWorkspaceParams object
// with the default values initialized.
func NewGetGatewaysCredentialIDForWorkspaceParams() *GetGatewaysCredentialIDForWorkspaceParams {
	var ()
	return &GetGatewaysCredentialIDForWorkspaceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetGatewaysCredentialIDForWorkspaceParamsWithTimeout creates a new GetGatewaysCredentialIDForWorkspaceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetGatewaysCredentialIDForWorkspaceParamsWithTimeout(timeout time.Duration) *GetGatewaysCredentialIDForWorkspaceParams {
	var ()
	return &GetGatewaysCredentialIDForWorkspaceParams{

		timeout: timeout,
	}
}

// NewGetGatewaysCredentialIDForWorkspaceParamsWithContext creates a new GetGatewaysCredentialIDForWorkspaceParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetGatewaysCredentialIDForWorkspaceParamsWithContext(ctx context.Context) *GetGatewaysCredentialIDForWorkspaceParams {
	var ()
	return &GetGatewaysCredentialIDForWorkspaceParams{

		Context: ctx,
	}
}

// NewGetGatewaysCredentialIDForWorkspaceParamsWithHTTPClient creates a new GetGatewaysCredentialIDForWorkspaceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetGatewaysCredentialIDForWorkspaceParamsWithHTTPClient(client *http.Client) *GetGatewaysCredentialIDForWorkspaceParams {
	var ()
	return &GetGatewaysCredentialIDForWorkspaceParams{
		HTTPClient: client,
	}
}

/*GetGatewaysCredentialIDForWorkspaceParams contains all the parameters to send to the API endpoint
for the get gateways credential Id for workspace operation typically these are written to a http.Request
*/
type GetGatewaysCredentialIDForWorkspaceParams struct {

	/*Body*/
	Body *model.PlatformResourceRequestJSON
	/*WorkspaceID*/
	WorkspaceID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get gateways credential Id for workspace params
func (o *GetGatewaysCredentialIDForWorkspaceParams) WithTimeout(timeout time.Duration) *GetGatewaysCredentialIDForWorkspaceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get gateways credential Id for workspace params
func (o *GetGatewaysCredentialIDForWorkspaceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get gateways credential Id for workspace params
func (o *GetGatewaysCredentialIDForWorkspaceParams) WithContext(ctx context.Context) *GetGatewaysCredentialIDForWorkspaceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get gateways credential Id for workspace params
func (o *GetGatewaysCredentialIDForWorkspaceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get gateways credential Id for workspace params
func (o *GetGatewaysCredentialIDForWorkspaceParams) WithHTTPClient(client *http.Client) *GetGatewaysCredentialIDForWorkspaceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get gateways credential Id for workspace params
func (o *GetGatewaysCredentialIDForWorkspaceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the get gateways credential Id for workspace params
func (o *GetGatewaysCredentialIDForWorkspaceParams) WithBody(body *model.PlatformResourceRequestJSON) *GetGatewaysCredentialIDForWorkspaceParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the get gateways credential Id for workspace params
func (o *GetGatewaysCredentialIDForWorkspaceParams) SetBody(body *model.PlatformResourceRequestJSON) {
	o.Body = body
}

// WithWorkspaceID adds the workspaceID to the get gateways credential Id for workspace params
func (o *GetGatewaysCredentialIDForWorkspaceParams) WithWorkspaceID(workspaceID int64) *GetGatewaysCredentialIDForWorkspaceParams {
	o.SetWorkspaceID(workspaceID)
	return o
}

// SetWorkspaceID adds the workspaceId to the get gateways credential Id for workspace params
func (o *GetGatewaysCredentialIDForWorkspaceParams) SetWorkspaceID(workspaceID int64) {
	o.WorkspaceID = workspaceID
}

// WriteToRequest writes these params to a swagger request
func (o *GetGatewaysCredentialIDForWorkspaceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param workspaceId
	if err := r.SetPathParam("workspaceId", swag.FormatInt64(o.WorkspaceID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
