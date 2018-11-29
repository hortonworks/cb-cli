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

// NewGetIPPoolsCredentialIDForWorkspaceParams creates a new GetIPPoolsCredentialIDForWorkspaceParams object
// with the default values initialized.
func NewGetIPPoolsCredentialIDForWorkspaceParams() *GetIPPoolsCredentialIDForWorkspaceParams {
	var ()
	return &GetIPPoolsCredentialIDForWorkspaceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetIPPoolsCredentialIDForWorkspaceParamsWithTimeout creates a new GetIPPoolsCredentialIDForWorkspaceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetIPPoolsCredentialIDForWorkspaceParamsWithTimeout(timeout time.Duration) *GetIPPoolsCredentialIDForWorkspaceParams {
	var ()
	return &GetIPPoolsCredentialIDForWorkspaceParams{

		timeout: timeout,
	}
}

// NewGetIPPoolsCredentialIDForWorkspaceParamsWithContext creates a new GetIPPoolsCredentialIDForWorkspaceParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetIPPoolsCredentialIDForWorkspaceParamsWithContext(ctx context.Context) *GetIPPoolsCredentialIDForWorkspaceParams {
	var ()
	return &GetIPPoolsCredentialIDForWorkspaceParams{

		Context: ctx,
	}
}

// NewGetIPPoolsCredentialIDForWorkspaceParamsWithHTTPClient creates a new GetIPPoolsCredentialIDForWorkspaceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetIPPoolsCredentialIDForWorkspaceParamsWithHTTPClient(client *http.Client) *GetIPPoolsCredentialIDForWorkspaceParams {
	var ()
	return &GetIPPoolsCredentialIDForWorkspaceParams{
		HTTPClient: client,
	}
}

/*GetIPPoolsCredentialIDForWorkspaceParams contains all the parameters to send to the API endpoint
for the get Ip pools credential Id for workspace operation typically these are written to a http.Request
*/
type GetIPPoolsCredentialIDForWorkspaceParams struct {

	/*Body*/
	Body *model.PlatformResourceRequestJSON
	/*WorkspaceID*/
	WorkspaceID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get Ip pools credential Id for workspace params
func (o *GetIPPoolsCredentialIDForWorkspaceParams) WithTimeout(timeout time.Duration) *GetIPPoolsCredentialIDForWorkspaceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get Ip pools credential Id for workspace params
func (o *GetIPPoolsCredentialIDForWorkspaceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get Ip pools credential Id for workspace params
func (o *GetIPPoolsCredentialIDForWorkspaceParams) WithContext(ctx context.Context) *GetIPPoolsCredentialIDForWorkspaceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get Ip pools credential Id for workspace params
func (o *GetIPPoolsCredentialIDForWorkspaceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get Ip pools credential Id for workspace params
func (o *GetIPPoolsCredentialIDForWorkspaceParams) WithHTTPClient(client *http.Client) *GetIPPoolsCredentialIDForWorkspaceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get Ip pools credential Id for workspace params
func (o *GetIPPoolsCredentialIDForWorkspaceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the get Ip pools credential Id for workspace params
func (o *GetIPPoolsCredentialIDForWorkspaceParams) WithBody(body *model.PlatformResourceRequestJSON) *GetIPPoolsCredentialIDForWorkspaceParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the get Ip pools credential Id for workspace params
func (o *GetIPPoolsCredentialIDForWorkspaceParams) SetBody(body *model.PlatformResourceRequestJSON) {
	o.Body = body
}

// WithWorkspaceID adds the workspaceID to the get Ip pools credential Id for workspace params
func (o *GetIPPoolsCredentialIDForWorkspaceParams) WithWorkspaceID(workspaceID int64) *GetIPPoolsCredentialIDForWorkspaceParams {
	o.SetWorkspaceID(workspaceID)
	return o
}

// SetWorkspaceID adds the workspaceId to the get Ip pools credential Id for workspace params
func (o *GetIPPoolsCredentialIDForWorkspaceParams) SetWorkspaceID(workspaceID int64) {
	o.WorkspaceID = workspaceID
}

// WriteToRequest writes these params to a swagger request
func (o *GetIPPoolsCredentialIDForWorkspaceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
