// Code generated by go-swagger; DO NOT EDIT.

package sdx

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

// NewDeleteSdxParams creates a new DeleteSdxParams object
// with the default values initialized.
func NewDeleteSdxParams() *DeleteSdxParams {
	var (
		forcedDefault = bool(false)
	)
	return &DeleteSdxParams{
		Forced: &forcedDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteSdxParamsWithTimeout creates a new DeleteSdxParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteSdxParamsWithTimeout(timeout time.Duration) *DeleteSdxParams {
	var (
		forcedDefault = bool(false)
	)
	return &DeleteSdxParams{
		Forced: &forcedDefault,

		timeout: timeout,
	}
}

// NewDeleteSdxParamsWithContext creates a new DeleteSdxParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteSdxParamsWithContext(ctx context.Context) *DeleteSdxParams {
	var (
		forcedDefault = bool(false)
	)
	return &DeleteSdxParams{
		Forced: &forcedDefault,

		Context: ctx,
	}
}

// NewDeleteSdxParamsWithHTTPClient creates a new DeleteSdxParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteSdxParamsWithHTTPClient(client *http.Client) *DeleteSdxParams {
	var (
		forcedDefault = bool(false)
	)
	return &DeleteSdxParams{
		Forced:     &forcedDefault,
		HTTPClient: client,
	}
}

/*
DeleteSdxParams contains all the parameters to send to the API endpoint
for the delete sdx operation typically these are written to a http.Request
*/
type DeleteSdxParams struct {

	/*Forced*/
	Forced *bool
	/*Name*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete sdx params
func (o *DeleteSdxParams) WithTimeout(timeout time.Duration) *DeleteSdxParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete sdx params
func (o *DeleteSdxParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete sdx params
func (o *DeleteSdxParams) WithContext(ctx context.Context) *DeleteSdxParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete sdx params
func (o *DeleteSdxParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete sdx params
func (o *DeleteSdxParams) WithHTTPClient(client *http.Client) *DeleteSdxParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete sdx params
func (o *DeleteSdxParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithForced adds the forced to the delete sdx params
func (o *DeleteSdxParams) WithForced(forced *bool) *DeleteSdxParams {
	o.SetForced(forced)
	return o
}

// SetForced adds the forced to the delete sdx params
func (o *DeleteSdxParams) SetForced(forced *bool) {
	o.Forced = forced
}

// WithName adds the name to the delete sdx params
func (o *DeleteSdxParams) WithName(name string) *DeleteSdxParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the delete sdx params
func (o *DeleteSdxParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteSdxParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
