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

	"github.com/hortonworks/cb-cli/models_cloudbreak"
)

// NewGetAccessConfigsForWorkspaceParams creates a new GetAccessConfigsForWorkspaceParams object
// with the default values initialized.
func NewGetAccessConfigsForWorkspaceParams() *GetAccessConfigsForWorkspaceParams {
	var ()
	return &GetAccessConfigsForWorkspaceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAccessConfigsForWorkspaceParamsWithTimeout creates a new GetAccessConfigsForWorkspaceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAccessConfigsForWorkspaceParamsWithTimeout(timeout time.Duration) *GetAccessConfigsForWorkspaceParams {
	var ()
	return &GetAccessConfigsForWorkspaceParams{

		timeout: timeout,
	}
}

// NewGetAccessConfigsForWorkspaceParamsWithContext creates a new GetAccessConfigsForWorkspaceParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAccessConfigsForWorkspaceParamsWithContext(ctx context.Context) *GetAccessConfigsForWorkspaceParams {
	var ()
	return &GetAccessConfigsForWorkspaceParams{

		Context: ctx,
	}
}

// NewGetAccessConfigsForWorkspaceParamsWithHTTPClient creates a new GetAccessConfigsForWorkspaceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAccessConfigsForWorkspaceParamsWithHTTPClient(client *http.Client) *GetAccessConfigsForWorkspaceParams {
	var ()
	return &GetAccessConfigsForWorkspaceParams{
		HTTPClient: client,
	}
}

/*GetAccessConfigsForWorkspaceParams contains all the parameters to send to the API endpoint
for the get access configs for workspace operation typically these are written to a http.Request
*/
type GetAccessConfigsForWorkspaceParams struct {

	/*Body*/
	Body *models_cloudbreak.PlatformResourceRequestJSON
	/*WorkspaceID*/
	WorkspaceID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get access configs for workspace params
func (o *GetAccessConfigsForWorkspaceParams) WithTimeout(timeout time.Duration) *GetAccessConfigsForWorkspaceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get access configs for workspace params
func (o *GetAccessConfigsForWorkspaceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get access configs for workspace params
func (o *GetAccessConfigsForWorkspaceParams) WithContext(ctx context.Context) *GetAccessConfigsForWorkspaceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get access configs for workspace params
func (o *GetAccessConfigsForWorkspaceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get access configs for workspace params
func (o *GetAccessConfigsForWorkspaceParams) WithHTTPClient(client *http.Client) *GetAccessConfigsForWorkspaceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get access configs for workspace params
func (o *GetAccessConfigsForWorkspaceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the get access configs for workspace params
func (o *GetAccessConfigsForWorkspaceParams) WithBody(body *models_cloudbreak.PlatformResourceRequestJSON) *GetAccessConfigsForWorkspaceParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the get access configs for workspace params
func (o *GetAccessConfigsForWorkspaceParams) SetBody(body *models_cloudbreak.PlatformResourceRequestJSON) {
	o.Body = body
}

// WithWorkspaceID adds the workspaceID to the get access configs for workspace params
func (o *GetAccessConfigsForWorkspaceParams) WithWorkspaceID(workspaceID int64) *GetAccessConfigsForWorkspaceParams {
	o.SetWorkspaceID(workspaceID)
	return o
}

// SetWorkspaceID adds the workspaceId to the get access configs for workspace params
func (o *GetAccessConfigsForWorkspaceParams) SetWorkspaceID(workspaceID int64) {
	o.WorkspaceID = workspaceID
}

// WriteToRequest writes these params to a swagger request
func (o *GetAccessConfigsForWorkspaceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body == nil {
		o.Body = new(models_cloudbreak.PlatformResourceRequestJSON)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
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
