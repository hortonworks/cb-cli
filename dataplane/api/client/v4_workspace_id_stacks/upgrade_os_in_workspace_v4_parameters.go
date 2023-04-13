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

// NewUpgradeOsInWorkspaceV4Params creates a new UpgradeOsInWorkspaceV4Params object
// with the default values initialized.
func NewUpgradeOsInWorkspaceV4Params() *UpgradeOsInWorkspaceV4Params {
	var ()
	return &UpgradeOsInWorkspaceV4Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpgradeOsInWorkspaceV4ParamsWithTimeout creates a new UpgradeOsInWorkspaceV4Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpgradeOsInWorkspaceV4ParamsWithTimeout(timeout time.Duration) *UpgradeOsInWorkspaceV4Params {
	var ()
	return &UpgradeOsInWorkspaceV4Params{

		timeout: timeout,
	}
}

// NewUpgradeOsInWorkspaceV4ParamsWithContext creates a new UpgradeOsInWorkspaceV4Params object
// with the default values initialized, and the ability to set a context for a request
func NewUpgradeOsInWorkspaceV4ParamsWithContext(ctx context.Context) *UpgradeOsInWorkspaceV4Params {
	var ()
	return &UpgradeOsInWorkspaceV4Params{

		Context: ctx,
	}
}

// NewUpgradeOsInWorkspaceV4ParamsWithHTTPClient creates a new UpgradeOsInWorkspaceV4Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpgradeOsInWorkspaceV4ParamsWithHTTPClient(client *http.Client) *UpgradeOsInWorkspaceV4Params {
	var ()
	return &UpgradeOsInWorkspaceV4Params{
		HTTPClient: client,
	}
}

/*
UpgradeOsInWorkspaceV4Params contains all the parameters to send to the API endpoint
for the upgrade os in workspace v4 operation typically these are written to a http.Request
*/
type UpgradeOsInWorkspaceV4Params struct {

	/*AccountID*/
	AccountID *string
	/*KeepVariant*/
	KeepVariant *bool
	/*Name*/
	Name string
	/*WorkspaceID*/
	WorkspaceID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the upgrade os in workspace v4 params
func (o *UpgradeOsInWorkspaceV4Params) WithTimeout(timeout time.Duration) *UpgradeOsInWorkspaceV4Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the upgrade os in workspace v4 params
func (o *UpgradeOsInWorkspaceV4Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the upgrade os in workspace v4 params
func (o *UpgradeOsInWorkspaceV4Params) WithContext(ctx context.Context) *UpgradeOsInWorkspaceV4Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the upgrade os in workspace v4 params
func (o *UpgradeOsInWorkspaceV4Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the upgrade os in workspace v4 params
func (o *UpgradeOsInWorkspaceV4Params) WithHTTPClient(client *http.Client) *UpgradeOsInWorkspaceV4Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the upgrade os in workspace v4 params
func (o *UpgradeOsInWorkspaceV4Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountID adds the accountID to the upgrade os in workspace v4 params
func (o *UpgradeOsInWorkspaceV4Params) WithAccountID(accountID *string) *UpgradeOsInWorkspaceV4Params {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the upgrade os in workspace v4 params
func (o *UpgradeOsInWorkspaceV4Params) SetAccountID(accountID *string) {
	o.AccountID = accountID
}

// WithKeepVariant adds the keepVariant to the upgrade os in workspace v4 params
func (o *UpgradeOsInWorkspaceV4Params) WithKeepVariant(keepVariant *bool) *UpgradeOsInWorkspaceV4Params {
	o.SetKeepVariant(keepVariant)
	return o
}

// SetKeepVariant adds the keepVariant to the upgrade os in workspace v4 params
func (o *UpgradeOsInWorkspaceV4Params) SetKeepVariant(keepVariant *bool) {
	o.KeepVariant = keepVariant
}

// WithName adds the name to the upgrade os in workspace v4 params
func (o *UpgradeOsInWorkspaceV4Params) WithName(name string) *UpgradeOsInWorkspaceV4Params {
	o.SetName(name)
	return o
}

// SetName adds the name to the upgrade os in workspace v4 params
func (o *UpgradeOsInWorkspaceV4Params) SetName(name string) {
	o.Name = name
}

// WithWorkspaceID adds the workspaceID to the upgrade os in workspace v4 params
func (o *UpgradeOsInWorkspaceV4Params) WithWorkspaceID(workspaceID int64) *UpgradeOsInWorkspaceV4Params {
	o.SetWorkspaceID(workspaceID)
	return o
}

// SetWorkspaceID adds the workspaceId to the upgrade os in workspace v4 params
func (o *UpgradeOsInWorkspaceV4Params) SetWorkspaceID(workspaceID int64) {
	o.WorkspaceID = workspaceID
}

// WriteToRequest writes these params to a swagger request
func (o *UpgradeOsInWorkspaceV4Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.KeepVariant != nil {

		// query param keepVariant
		var qrKeepVariant bool
		if o.KeepVariant != nil {
			qrKeepVariant = *o.KeepVariant
		}
		qKeepVariant := swag.FormatBool(qrKeepVariant)
		if qKeepVariant != "" {
			if err := r.SetQueryParam("keepVariant", qKeepVariant); err != nil {
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
