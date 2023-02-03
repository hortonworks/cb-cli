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

// NewListStackInWorkspaceV4Params creates a new ListStackInWorkspaceV4Params object
// with the default values initialized.
func NewListStackInWorkspaceV4Params() *ListStackInWorkspaceV4Params {
	var (
		onlyDatalakesDefault = bool(false)
	)
	return &ListStackInWorkspaceV4Params{
		OnlyDatalakes: &onlyDatalakesDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewListStackInWorkspaceV4ParamsWithTimeout creates a new ListStackInWorkspaceV4Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewListStackInWorkspaceV4ParamsWithTimeout(timeout time.Duration) *ListStackInWorkspaceV4Params {
	var (
		onlyDatalakesDefault = bool(false)
	)
	return &ListStackInWorkspaceV4Params{
		OnlyDatalakes: &onlyDatalakesDefault,

		timeout: timeout,
	}
}

// NewListStackInWorkspaceV4ParamsWithContext creates a new ListStackInWorkspaceV4Params object
// with the default values initialized, and the ability to set a context for a request
func NewListStackInWorkspaceV4ParamsWithContext(ctx context.Context) *ListStackInWorkspaceV4Params {
	var (
		onlyDatalakesDefault = bool(false)
	)
	return &ListStackInWorkspaceV4Params{
		OnlyDatalakes: &onlyDatalakesDefault,

		Context: ctx,
	}
}

// NewListStackInWorkspaceV4ParamsWithHTTPClient creates a new ListStackInWorkspaceV4Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListStackInWorkspaceV4ParamsWithHTTPClient(client *http.Client) *ListStackInWorkspaceV4Params {
	var (
		onlyDatalakesDefault = bool(false)
	)
	return &ListStackInWorkspaceV4Params{
		OnlyDatalakes: &onlyDatalakesDefault,
		HTTPClient:    client,
	}
}

/*
ListStackInWorkspaceV4Params contains all the parameters to send to the API endpoint
for the list stack in workspace v4 operation typically these are written to a http.Request
*/
type ListStackInWorkspaceV4Params struct {

	/*Environment*/
	Environment *string
	/*OnlyDatalakes*/
	OnlyDatalakes *bool
	/*WorkspaceID*/
	WorkspaceID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list stack in workspace v4 params
func (o *ListStackInWorkspaceV4Params) WithTimeout(timeout time.Duration) *ListStackInWorkspaceV4Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list stack in workspace v4 params
func (o *ListStackInWorkspaceV4Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list stack in workspace v4 params
func (o *ListStackInWorkspaceV4Params) WithContext(ctx context.Context) *ListStackInWorkspaceV4Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list stack in workspace v4 params
func (o *ListStackInWorkspaceV4Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list stack in workspace v4 params
func (o *ListStackInWorkspaceV4Params) WithHTTPClient(client *http.Client) *ListStackInWorkspaceV4Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list stack in workspace v4 params
func (o *ListStackInWorkspaceV4Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEnvironment adds the environment to the list stack in workspace v4 params
func (o *ListStackInWorkspaceV4Params) WithEnvironment(environment *string) *ListStackInWorkspaceV4Params {
	o.SetEnvironment(environment)
	return o
}

// SetEnvironment adds the environment to the list stack in workspace v4 params
func (o *ListStackInWorkspaceV4Params) SetEnvironment(environment *string) {
	o.Environment = environment
}

// WithOnlyDatalakes adds the onlyDatalakes to the list stack in workspace v4 params
func (o *ListStackInWorkspaceV4Params) WithOnlyDatalakes(onlyDatalakes *bool) *ListStackInWorkspaceV4Params {
	o.SetOnlyDatalakes(onlyDatalakes)
	return o
}

// SetOnlyDatalakes adds the onlyDatalakes to the list stack in workspace v4 params
func (o *ListStackInWorkspaceV4Params) SetOnlyDatalakes(onlyDatalakes *bool) {
	o.OnlyDatalakes = onlyDatalakes
}

// WithWorkspaceID adds the workspaceID to the list stack in workspace v4 params
func (o *ListStackInWorkspaceV4Params) WithWorkspaceID(workspaceID int64) *ListStackInWorkspaceV4Params {
	o.SetWorkspaceID(workspaceID)
	return o
}

// SetWorkspaceID adds the workspaceId to the list stack in workspace v4 params
func (o *ListStackInWorkspaceV4Params) SetWorkspaceID(workspaceID int64) {
	o.WorkspaceID = workspaceID
}

// WriteToRequest writes these params to a swagger request
func (o *ListStackInWorkspaceV4Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Environment != nil {

		// query param environment
		var qrEnvironment string
		if o.Environment != nil {
			qrEnvironment = *o.Environment
		}
		qEnvironment := qrEnvironment
		if qEnvironment != "" {
			if err := r.SetQueryParam("environment", qEnvironment); err != nil {
				return err
			}
		}

	}

	if o.OnlyDatalakes != nil {

		// query param onlyDatalakes
		var qrOnlyDatalakes bool
		if o.OnlyDatalakes != nil {
			qrOnlyDatalakes = *o.OnlyDatalakes
		}
		qOnlyDatalakes := swag.FormatBool(qrOnlyDatalakes)
		if qOnlyDatalakes != "" {
			if err := r.SetQueryParam("onlyDatalakes", qOnlyDatalakes); err != nil {
				return err
			}
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
