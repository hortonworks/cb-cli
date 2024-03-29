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

// NewRepairStackInWorkspaceV4Params creates a new RepairStackInWorkspaceV4Params object
// with the default values initialized.
func NewRepairStackInWorkspaceV4Params() *RepairStackInWorkspaceV4Params {
	var ()
	return &RepairStackInWorkspaceV4Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewRepairStackInWorkspaceV4ParamsWithTimeout creates a new RepairStackInWorkspaceV4Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewRepairStackInWorkspaceV4ParamsWithTimeout(timeout time.Duration) *RepairStackInWorkspaceV4Params {
	var ()
	return &RepairStackInWorkspaceV4Params{

		timeout: timeout,
	}
}

// NewRepairStackInWorkspaceV4ParamsWithContext creates a new RepairStackInWorkspaceV4Params object
// with the default values initialized, and the ability to set a context for a request
func NewRepairStackInWorkspaceV4ParamsWithContext(ctx context.Context) *RepairStackInWorkspaceV4Params {
	var ()
	return &RepairStackInWorkspaceV4Params{

		Context: ctx,
	}
}

// NewRepairStackInWorkspaceV4ParamsWithHTTPClient creates a new RepairStackInWorkspaceV4Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRepairStackInWorkspaceV4ParamsWithHTTPClient(client *http.Client) *RepairStackInWorkspaceV4Params {
	var ()
	return &RepairStackInWorkspaceV4Params{
		HTTPClient: client,
	}
}

/*
RepairStackInWorkspaceV4Params contains all the parameters to send to the API endpoint
for the repair stack in workspace v4 operation typically these are written to a http.Request
*/
type RepairStackInWorkspaceV4Params struct {

	/*AccountID*/
	AccountID *string
	/*Body*/
	Body *model.ClusterRepairV4Request
	/*Name*/
	Name string
	/*WorkspaceID*/
	WorkspaceID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the repair stack in workspace v4 params
func (o *RepairStackInWorkspaceV4Params) WithTimeout(timeout time.Duration) *RepairStackInWorkspaceV4Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the repair stack in workspace v4 params
func (o *RepairStackInWorkspaceV4Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the repair stack in workspace v4 params
func (o *RepairStackInWorkspaceV4Params) WithContext(ctx context.Context) *RepairStackInWorkspaceV4Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the repair stack in workspace v4 params
func (o *RepairStackInWorkspaceV4Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the repair stack in workspace v4 params
func (o *RepairStackInWorkspaceV4Params) WithHTTPClient(client *http.Client) *RepairStackInWorkspaceV4Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the repair stack in workspace v4 params
func (o *RepairStackInWorkspaceV4Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountID adds the accountID to the repair stack in workspace v4 params
func (o *RepairStackInWorkspaceV4Params) WithAccountID(accountID *string) *RepairStackInWorkspaceV4Params {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the repair stack in workspace v4 params
func (o *RepairStackInWorkspaceV4Params) SetAccountID(accountID *string) {
	o.AccountID = accountID
}

// WithBody adds the body to the repair stack in workspace v4 params
func (o *RepairStackInWorkspaceV4Params) WithBody(body *model.ClusterRepairV4Request) *RepairStackInWorkspaceV4Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the repair stack in workspace v4 params
func (o *RepairStackInWorkspaceV4Params) SetBody(body *model.ClusterRepairV4Request) {
	o.Body = body
}

// WithName adds the name to the repair stack in workspace v4 params
func (o *RepairStackInWorkspaceV4Params) WithName(name string) *RepairStackInWorkspaceV4Params {
	o.SetName(name)
	return o
}

// SetName adds the name to the repair stack in workspace v4 params
func (o *RepairStackInWorkspaceV4Params) SetName(name string) {
	o.Name = name
}

// WithWorkspaceID adds the workspaceID to the repair stack in workspace v4 params
func (o *RepairStackInWorkspaceV4Params) WithWorkspaceID(workspaceID int64) *RepairStackInWorkspaceV4Params {
	o.SetWorkspaceID(workspaceID)
	return o
}

// SetWorkspaceID adds the workspaceId to the repair stack in workspace v4 params
func (o *RepairStackInWorkspaceV4Params) SetWorkspaceID(workspaceID int64) {
	o.WorkspaceID = workspaceID
}

// WriteToRequest writes these params to a swagger request
func (o *RepairStackInWorkspaceV4Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
