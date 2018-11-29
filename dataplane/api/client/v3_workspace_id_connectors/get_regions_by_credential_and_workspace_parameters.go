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

// NewGetRegionsByCredentialAndWorkspaceParams creates a new GetRegionsByCredentialAndWorkspaceParams object
// with the default values initialized.
func NewGetRegionsByCredentialAndWorkspaceParams() *GetRegionsByCredentialAndWorkspaceParams {
	var ()
	return &GetRegionsByCredentialAndWorkspaceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRegionsByCredentialAndWorkspaceParamsWithTimeout creates a new GetRegionsByCredentialAndWorkspaceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRegionsByCredentialAndWorkspaceParamsWithTimeout(timeout time.Duration) *GetRegionsByCredentialAndWorkspaceParams {
	var ()
	return &GetRegionsByCredentialAndWorkspaceParams{

		timeout: timeout,
	}
}

// NewGetRegionsByCredentialAndWorkspaceParamsWithContext creates a new GetRegionsByCredentialAndWorkspaceParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRegionsByCredentialAndWorkspaceParamsWithContext(ctx context.Context) *GetRegionsByCredentialAndWorkspaceParams {
	var ()
	return &GetRegionsByCredentialAndWorkspaceParams{

		Context: ctx,
	}
}

// NewGetRegionsByCredentialAndWorkspaceParamsWithHTTPClient creates a new GetRegionsByCredentialAndWorkspaceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRegionsByCredentialAndWorkspaceParamsWithHTTPClient(client *http.Client) *GetRegionsByCredentialAndWorkspaceParams {
	var ()
	return &GetRegionsByCredentialAndWorkspaceParams{
		HTTPClient: client,
	}
}

/*GetRegionsByCredentialAndWorkspaceParams contains all the parameters to send to the API endpoint
for the get regions by credential and workspace operation typically these are written to a http.Request
*/
type GetRegionsByCredentialAndWorkspaceParams struct {

	/*Body*/
	Body *model.PlatformResourceRequestJSON
	/*WorkspaceID*/
	WorkspaceID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get regions by credential and workspace params
func (o *GetRegionsByCredentialAndWorkspaceParams) WithTimeout(timeout time.Duration) *GetRegionsByCredentialAndWorkspaceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get regions by credential and workspace params
func (o *GetRegionsByCredentialAndWorkspaceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get regions by credential and workspace params
func (o *GetRegionsByCredentialAndWorkspaceParams) WithContext(ctx context.Context) *GetRegionsByCredentialAndWorkspaceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get regions by credential and workspace params
func (o *GetRegionsByCredentialAndWorkspaceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get regions by credential and workspace params
func (o *GetRegionsByCredentialAndWorkspaceParams) WithHTTPClient(client *http.Client) *GetRegionsByCredentialAndWorkspaceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get regions by credential and workspace params
func (o *GetRegionsByCredentialAndWorkspaceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the get regions by credential and workspace params
func (o *GetRegionsByCredentialAndWorkspaceParams) WithBody(body *model.PlatformResourceRequestJSON) *GetRegionsByCredentialAndWorkspaceParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the get regions by credential and workspace params
func (o *GetRegionsByCredentialAndWorkspaceParams) SetBody(body *model.PlatformResourceRequestJSON) {
	o.Body = body
}

// WithWorkspaceID adds the workspaceID to the get regions by credential and workspace params
func (o *GetRegionsByCredentialAndWorkspaceParams) WithWorkspaceID(workspaceID int64) *GetRegionsByCredentialAndWorkspaceParams {
	o.SetWorkspaceID(workspaceID)
	return o
}

// SetWorkspaceID adds the workspaceId to the get regions by credential and workspace params
func (o *GetRegionsByCredentialAndWorkspaceParams) SetWorkspaceID(workspaceID int64) {
	o.WorkspaceID = workspaceID
}

// WriteToRequest writes these params to a swagger request
func (o *GetRegionsByCredentialAndWorkspaceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
