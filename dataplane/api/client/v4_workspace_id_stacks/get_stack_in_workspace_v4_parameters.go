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

// NewGetStackInWorkspaceV4Params creates a new GetStackInWorkspaceV4Params object
// with the default values initialized.
func NewGetStackInWorkspaceV4Params() *GetStackInWorkspaceV4Params {
	var ()
	return &GetStackInWorkspaceV4Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetStackInWorkspaceV4ParamsWithTimeout creates a new GetStackInWorkspaceV4Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetStackInWorkspaceV4ParamsWithTimeout(timeout time.Duration) *GetStackInWorkspaceV4Params {
	var ()
	return &GetStackInWorkspaceV4Params{

		timeout: timeout,
	}
}

// NewGetStackInWorkspaceV4ParamsWithContext creates a new GetStackInWorkspaceV4Params object
// with the default values initialized, and the ability to set a context for a request
func NewGetStackInWorkspaceV4ParamsWithContext(ctx context.Context) *GetStackInWorkspaceV4Params {
	var ()
	return &GetStackInWorkspaceV4Params{

		Context: ctx,
	}
}

// NewGetStackInWorkspaceV4ParamsWithHTTPClient creates a new GetStackInWorkspaceV4Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetStackInWorkspaceV4ParamsWithHTTPClient(client *http.Client) *GetStackInWorkspaceV4Params {
	var ()
	return &GetStackInWorkspaceV4Params{
		HTTPClient: client,
	}
}

/*GetStackInWorkspaceV4Params contains all the parameters to send to the API endpoint
for the get stack in workspace v4 operation typically these are written to a http.Request
*/
type GetStackInWorkspaceV4Params struct {

	/*AccountID*/
	AccountID *string
	/*Entries*/
	Entries []string
	/*Name*/
	Name string
	/*WorkspaceID*/
	WorkspaceID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get stack in workspace v4 params
func (o *GetStackInWorkspaceV4Params) WithTimeout(timeout time.Duration) *GetStackInWorkspaceV4Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get stack in workspace v4 params
func (o *GetStackInWorkspaceV4Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get stack in workspace v4 params
func (o *GetStackInWorkspaceV4Params) WithContext(ctx context.Context) *GetStackInWorkspaceV4Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get stack in workspace v4 params
func (o *GetStackInWorkspaceV4Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get stack in workspace v4 params
func (o *GetStackInWorkspaceV4Params) WithHTTPClient(client *http.Client) *GetStackInWorkspaceV4Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get stack in workspace v4 params
func (o *GetStackInWorkspaceV4Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountID adds the accountID to the get stack in workspace v4 params
func (o *GetStackInWorkspaceV4Params) WithAccountID(accountID *string) *GetStackInWorkspaceV4Params {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the get stack in workspace v4 params
func (o *GetStackInWorkspaceV4Params) SetAccountID(accountID *string) {
	o.AccountID = accountID
}

// WithEntries adds the entries to the get stack in workspace v4 params
func (o *GetStackInWorkspaceV4Params) WithEntries(entries []string) *GetStackInWorkspaceV4Params {
	o.SetEntries(entries)
	return o
}

// SetEntries adds the entries to the get stack in workspace v4 params
func (o *GetStackInWorkspaceV4Params) SetEntries(entries []string) {
	o.Entries = entries
}

// WithName adds the name to the get stack in workspace v4 params
func (o *GetStackInWorkspaceV4Params) WithName(name string) *GetStackInWorkspaceV4Params {
	o.SetName(name)
	return o
}

// SetName adds the name to the get stack in workspace v4 params
func (o *GetStackInWorkspaceV4Params) SetName(name string) {
	o.Name = name
}

// WithWorkspaceID adds the workspaceID to the get stack in workspace v4 params
func (o *GetStackInWorkspaceV4Params) WithWorkspaceID(workspaceID int64) *GetStackInWorkspaceV4Params {
	o.SetWorkspaceID(workspaceID)
	return o
}

// SetWorkspaceID adds the workspaceId to the get stack in workspace v4 params
func (o *GetStackInWorkspaceV4Params) SetWorkspaceID(workspaceID int64) {
	o.WorkspaceID = workspaceID
}

// WriteToRequest writes these params to a swagger request
func (o *GetStackInWorkspaceV4Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	valuesEntries := o.Entries

	joinedEntries := swag.JoinByFormat(valuesEntries, "multi")
	// query array param entries
	if err := r.SetQueryParam("entries", joinedEntries...); err != nil {
		return err
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
