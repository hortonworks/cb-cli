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
)

// NewDeleteWithKerberosParams creates a new DeleteWithKerberosParams object
// with the default values initialized.
func NewDeleteWithKerberosParams() *DeleteWithKerberosParams {
	var (
		forcedDefault = bool(false)
	)
	return &DeleteWithKerberosParams{
		Forced: &forcedDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteWithKerberosParamsWithTimeout creates a new DeleteWithKerberosParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteWithKerberosParamsWithTimeout(timeout time.Duration) *DeleteWithKerberosParams {
	var (
		forcedDefault = bool(false)
	)
	return &DeleteWithKerberosParams{
		Forced: &forcedDefault,

		timeout: timeout,
	}
}

// NewDeleteWithKerberosParamsWithContext creates a new DeleteWithKerberosParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteWithKerberosParamsWithContext(ctx context.Context) *DeleteWithKerberosParams {
	var (
		forcedDefault = bool(false)
	)
	return &DeleteWithKerberosParams{
		Forced: &forcedDefault,

		Context: ctx,
	}
}

// NewDeleteWithKerberosParamsWithHTTPClient creates a new DeleteWithKerberosParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteWithKerberosParamsWithHTTPClient(client *http.Client) *DeleteWithKerberosParams {
	var (
		forcedDefault = bool(false)
	)
	return &DeleteWithKerberosParams{
		Forced:     &forcedDefault,
		HTTPClient: client,
	}
}

/*
DeleteWithKerberosParams contains all the parameters to send to the API endpoint
for the delete with kerberos operation typically these are written to a http.Request
*/
type DeleteWithKerberosParams struct {

	/*AccountID*/
	AccountID *string
	/*Forced*/
	Forced *bool
	/*Name*/
	Name string
	/*WorkspaceID*/
	WorkspaceID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete with kerberos params
func (o *DeleteWithKerberosParams) WithTimeout(timeout time.Duration) *DeleteWithKerberosParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete with kerberos params
func (o *DeleteWithKerberosParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete with kerberos params
func (o *DeleteWithKerberosParams) WithContext(ctx context.Context) *DeleteWithKerberosParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete with kerberos params
func (o *DeleteWithKerberosParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete with kerberos params
func (o *DeleteWithKerberosParams) WithHTTPClient(client *http.Client) *DeleteWithKerberosParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete with kerberos params
func (o *DeleteWithKerberosParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountID adds the accountID to the delete with kerberos params
func (o *DeleteWithKerberosParams) WithAccountID(accountID *string) *DeleteWithKerberosParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the delete with kerberos params
func (o *DeleteWithKerberosParams) SetAccountID(accountID *string) {
	o.AccountID = accountID
}

// WithForced adds the forced to the delete with kerberos params
func (o *DeleteWithKerberosParams) WithForced(forced *bool) *DeleteWithKerberosParams {
	o.SetForced(forced)
	return o
}

// SetForced adds the forced to the delete with kerberos params
func (o *DeleteWithKerberosParams) SetForced(forced *bool) {
	o.Forced = forced
}

// WithName adds the name to the delete with kerberos params
func (o *DeleteWithKerberosParams) WithName(name string) *DeleteWithKerberosParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the delete with kerberos params
func (o *DeleteWithKerberosParams) SetName(name string) {
	o.Name = name
}

// WithWorkspaceID adds the workspaceID to the delete with kerberos params
func (o *DeleteWithKerberosParams) WithWorkspaceID(workspaceID int64) *DeleteWithKerberosParams {
	o.SetWorkspaceID(workspaceID)
	return o
}

// SetWorkspaceID adds the workspaceId to the delete with kerberos params
func (o *DeleteWithKerberosParams) SetWorkspaceID(workspaceID int64) {
	o.WorkspaceID = workspaceID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteWithKerberosParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Forced != nil {

		// query param forced
		var qrForced bool
		if o.Forced != nil {
			qrForced = *o.Forced
		}
		qForced := swag.FormatBool(qrForced)
		if qForced != "" {
			if err := r.SetQueryParam("forced", qForced); err != nil {
				return err
			}
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
