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

// NewDeleteMultipleInstancesStackV4Params creates a new DeleteMultipleInstancesStackV4Params object
// with the default values initialized.
func NewDeleteMultipleInstancesStackV4Params() *DeleteMultipleInstancesStackV4Params {
	var (
		forcedDefault = bool(false)
	)
	return &DeleteMultipleInstancesStackV4Params{
		Forced: &forcedDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteMultipleInstancesStackV4ParamsWithTimeout creates a new DeleteMultipleInstancesStackV4Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteMultipleInstancesStackV4ParamsWithTimeout(timeout time.Duration) *DeleteMultipleInstancesStackV4Params {
	var (
		forcedDefault = bool(false)
	)
	return &DeleteMultipleInstancesStackV4Params{
		Forced: &forcedDefault,

		timeout: timeout,
	}
}

// NewDeleteMultipleInstancesStackV4ParamsWithContext creates a new DeleteMultipleInstancesStackV4Params object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteMultipleInstancesStackV4ParamsWithContext(ctx context.Context) *DeleteMultipleInstancesStackV4Params {
	var (
		forcedDefault = bool(false)
	)
	return &DeleteMultipleInstancesStackV4Params{
		Forced: &forcedDefault,

		Context: ctx,
	}
}

// NewDeleteMultipleInstancesStackV4ParamsWithHTTPClient creates a new DeleteMultipleInstancesStackV4Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteMultipleInstancesStackV4ParamsWithHTTPClient(client *http.Client) *DeleteMultipleInstancesStackV4Params {
	var (
		forcedDefault = bool(false)
	)
	return &DeleteMultipleInstancesStackV4Params{
		Forced:     &forcedDefault,
		HTTPClient: client,
	}
}

/*
DeleteMultipleInstancesStackV4Params contains all the parameters to send to the API endpoint
for the delete multiple instances stack v4 operation typically these are written to a http.Request
*/
type DeleteMultipleInstancesStackV4Params struct {

	/*AccountID*/
	AccountID *string
	/*Forced*/
	Forced *bool
	/*ID*/
	ID []string
	/*Name*/
	Name string
	/*WorkspaceID*/
	WorkspaceID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete multiple instances stack v4 params
func (o *DeleteMultipleInstancesStackV4Params) WithTimeout(timeout time.Duration) *DeleteMultipleInstancesStackV4Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete multiple instances stack v4 params
func (o *DeleteMultipleInstancesStackV4Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete multiple instances stack v4 params
func (o *DeleteMultipleInstancesStackV4Params) WithContext(ctx context.Context) *DeleteMultipleInstancesStackV4Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete multiple instances stack v4 params
func (o *DeleteMultipleInstancesStackV4Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete multiple instances stack v4 params
func (o *DeleteMultipleInstancesStackV4Params) WithHTTPClient(client *http.Client) *DeleteMultipleInstancesStackV4Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete multiple instances stack v4 params
func (o *DeleteMultipleInstancesStackV4Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountID adds the accountID to the delete multiple instances stack v4 params
func (o *DeleteMultipleInstancesStackV4Params) WithAccountID(accountID *string) *DeleteMultipleInstancesStackV4Params {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the delete multiple instances stack v4 params
func (o *DeleteMultipleInstancesStackV4Params) SetAccountID(accountID *string) {
	o.AccountID = accountID
}

// WithForced adds the forced to the delete multiple instances stack v4 params
func (o *DeleteMultipleInstancesStackV4Params) WithForced(forced *bool) *DeleteMultipleInstancesStackV4Params {
	o.SetForced(forced)
	return o
}

// SetForced adds the forced to the delete multiple instances stack v4 params
func (o *DeleteMultipleInstancesStackV4Params) SetForced(forced *bool) {
	o.Forced = forced
}

// WithID adds the id to the delete multiple instances stack v4 params
func (o *DeleteMultipleInstancesStackV4Params) WithID(id []string) *DeleteMultipleInstancesStackV4Params {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete multiple instances stack v4 params
func (o *DeleteMultipleInstancesStackV4Params) SetID(id []string) {
	o.ID = id
}

// WithName adds the name to the delete multiple instances stack v4 params
func (o *DeleteMultipleInstancesStackV4Params) WithName(name string) *DeleteMultipleInstancesStackV4Params {
	o.SetName(name)
	return o
}

// SetName adds the name to the delete multiple instances stack v4 params
func (o *DeleteMultipleInstancesStackV4Params) SetName(name string) {
	o.Name = name
}

// WithWorkspaceID adds the workspaceID to the delete multiple instances stack v4 params
func (o *DeleteMultipleInstancesStackV4Params) WithWorkspaceID(workspaceID int64) *DeleteMultipleInstancesStackV4Params {
	o.SetWorkspaceID(workspaceID)
	return o
}

// SetWorkspaceID adds the workspaceId to the delete multiple instances stack v4 params
func (o *DeleteMultipleInstancesStackV4Params) SetWorkspaceID(workspaceID int64) {
	o.WorkspaceID = workspaceID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteMultipleInstancesStackV4Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	valuesID := o.ID

	joinedID := swag.JoinByFormat(valuesID, "multi")
	// query array param id
	if err := r.SetQueryParam("id", joinedID...); err != nil {
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
