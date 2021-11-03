// Code generated by go-swagger; DO NOT EDIT.

package v4_workspace_id_stacks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api/model"
)

// NewDetachStackRecipeParams creates a new DetachStackRecipeParams object
// with the default values initialized.
func NewDetachStackRecipeParams() *DetachStackRecipeParams {
	var ()
	return &DetachStackRecipeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDetachStackRecipeParamsWithTimeout creates a new DetachStackRecipeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDetachStackRecipeParamsWithTimeout(timeout time.Duration) *DetachStackRecipeParams {
	var ()
	return &DetachStackRecipeParams{

		timeout: timeout,
	}
}

// NewDetachStackRecipeParamsWithContext creates a new DetachStackRecipeParams object
// with the default values initialized, and the ability to set a context for a request
func NewDetachStackRecipeParamsWithContext(ctx context.Context) *DetachStackRecipeParams {
	var ()
	return &DetachStackRecipeParams{

		Context: ctx,
	}
}

// NewDetachStackRecipeParamsWithHTTPClient creates a new DetachStackRecipeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDetachStackRecipeParamsWithHTTPClient(client *http.Client) *DetachStackRecipeParams {
	var ()
	return &DetachStackRecipeParams{
		HTTPClient: client,
	}
}

/*DetachStackRecipeParams contains all the parameters to send to the API endpoint
for the detach stack recipe operation typically these are written to a http.Request
*/
type DetachStackRecipeParams struct {

	/*AccountID*/
	AccountID *string
	/*Body*/
	Body *model.DetachRecipeV4Request
	/*Name*/
	Name string
	/*WorkspaceID*/
	WorkspaceID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the detach stack recipe params
func (o *DetachStackRecipeParams) WithTimeout(timeout time.Duration) *DetachStackRecipeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the detach stack recipe params
func (o *DetachStackRecipeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the detach stack recipe params
func (o *DetachStackRecipeParams) WithContext(ctx context.Context) *DetachStackRecipeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the detach stack recipe params
func (o *DetachStackRecipeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the detach stack recipe params
func (o *DetachStackRecipeParams) WithHTTPClient(client *http.Client) *DetachStackRecipeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the detach stack recipe params
func (o *DetachStackRecipeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountID adds the accountID to the detach stack recipe params
func (o *DetachStackRecipeParams) WithAccountID(accountID *string) *DetachStackRecipeParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the detach stack recipe params
func (o *DetachStackRecipeParams) SetAccountID(accountID *string) {
	o.AccountID = accountID
}

// WithBody adds the body to the detach stack recipe params
func (o *DetachStackRecipeParams) WithBody(body *model.DetachRecipeV4Request) *DetachStackRecipeParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the detach stack recipe params
func (o *DetachStackRecipeParams) SetBody(body *model.DetachRecipeV4Request) {
	o.Body = body
}

// WithName adds the name to the detach stack recipe params
func (o *DetachStackRecipeParams) WithName(name string) *DetachStackRecipeParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the detach stack recipe params
func (o *DetachStackRecipeParams) SetName(name string) {
	o.Name = name
}

// WithWorkspaceID adds the workspaceID to the detach stack recipe params
func (o *DetachStackRecipeParams) WithWorkspaceID(workspaceID int64) *DetachStackRecipeParams {
	o.SetWorkspaceID(workspaceID)
	return o
}

// SetWorkspaceID adds the workspaceId to the detach stack recipe params
func (o *DetachStackRecipeParams) SetWorkspaceID(workspaceID int64) {
	o.WorkspaceID = workspaceID
}

// WriteToRequest writes these params to a swagger request
func (o *DetachStackRecipeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AccountID != nil {

		// query param accountId
		var qrAccountID string
		if o.AccountID != nil {
			qrAccountID = *o.AccountID
		}
		qAccountID := qrAccountID
		if qAccountID != "" {
			if err := r.SetQueryParam("accountId", qAccountID); err != nil {
				return err
			}
		}

	}

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
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
