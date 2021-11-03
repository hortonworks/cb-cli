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

// NewRefreshStackRecipesParams creates a new RefreshStackRecipesParams object
// with the default values initialized.
func NewRefreshStackRecipesParams() *RefreshStackRecipesParams {
	var ()
	return &RefreshStackRecipesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRefreshStackRecipesParamsWithTimeout creates a new RefreshStackRecipesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRefreshStackRecipesParamsWithTimeout(timeout time.Duration) *RefreshStackRecipesParams {
	var ()
	return &RefreshStackRecipesParams{

		timeout: timeout,
	}
}

// NewRefreshStackRecipesParamsWithContext creates a new RefreshStackRecipesParams object
// with the default values initialized, and the ability to set a context for a request
func NewRefreshStackRecipesParamsWithContext(ctx context.Context) *RefreshStackRecipesParams {
	var ()
	return &RefreshStackRecipesParams{

		Context: ctx,
	}
}

// NewRefreshStackRecipesParamsWithHTTPClient creates a new RefreshStackRecipesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRefreshStackRecipesParamsWithHTTPClient(client *http.Client) *RefreshStackRecipesParams {
	var ()
	return &RefreshStackRecipesParams{
		HTTPClient: client,
	}
}

/*RefreshStackRecipesParams contains all the parameters to send to the API endpoint
for the refresh stack recipes operation typically these are written to a http.Request
*/
type RefreshStackRecipesParams struct {

	/*AccountID*/
	AccountID *string
	/*Body*/
	Body *model.UpdateRecipesV4Request
	/*Name*/
	Name string
	/*WorkspaceID*/
	WorkspaceID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the refresh stack recipes params
func (o *RefreshStackRecipesParams) WithTimeout(timeout time.Duration) *RefreshStackRecipesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the refresh stack recipes params
func (o *RefreshStackRecipesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the refresh stack recipes params
func (o *RefreshStackRecipesParams) WithContext(ctx context.Context) *RefreshStackRecipesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the refresh stack recipes params
func (o *RefreshStackRecipesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the refresh stack recipes params
func (o *RefreshStackRecipesParams) WithHTTPClient(client *http.Client) *RefreshStackRecipesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the refresh stack recipes params
func (o *RefreshStackRecipesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountID adds the accountID to the refresh stack recipes params
func (o *RefreshStackRecipesParams) WithAccountID(accountID *string) *RefreshStackRecipesParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the refresh stack recipes params
func (o *RefreshStackRecipesParams) SetAccountID(accountID *string) {
	o.AccountID = accountID
}

// WithBody adds the body to the refresh stack recipes params
func (o *RefreshStackRecipesParams) WithBody(body *model.UpdateRecipesV4Request) *RefreshStackRecipesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the refresh stack recipes params
func (o *RefreshStackRecipesParams) SetBody(body *model.UpdateRecipesV4Request) {
	o.Body = body
}

// WithName adds the name to the refresh stack recipes params
func (o *RefreshStackRecipesParams) WithName(name string) *RefreshStackRecipesParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the refresh stack recipes params
func (o *RefreshStackRecipesParams) SetName(name string) {
	o.Name = name
}

// WithWorkspaceID adds the workspaceID to the refresh stack recipes params
func (o *RefreshStackRecipesParams) WithWorkspaceID(workspaceID int64) *RefreshStackRecipesParams {
	o.SetWorkspaceID(workspaceID)
	return o
}

// SetWorkspaceID adds the workspaceId to the refresh stack recipes params
func (o *RefreshStackRecipesParams) SetWorkspaceID(workspaceID int64) {
	o.WorkspaceID = workspaceID
}

// WriteToRequest writes these params to a swagger request
func (o *RefreshStackRecipesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
