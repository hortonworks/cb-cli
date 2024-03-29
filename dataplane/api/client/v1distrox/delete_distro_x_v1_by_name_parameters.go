// Code generated by go-swagger; DO NOT EDIT.

package v1distrox

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

// NewDeleteDistroXV1ByNameParams creates a new DeleteDistroXV1ByNameParams object
// with the default values initialized.
func NewDeleteDistroXV1ByNameParams() *DeleteDistroXV1ByNameParams {
	var (
		forcedDefault = bool(false)
	)
	return &DeleteDistroXV1ByNameParams{
		Forced: &forcedDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteDistroXV1ByNameParamsWithTimeout creates a new DeleteDistroXV1ByNameParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteDistroXV1ByNameParamsWithTimeout(timeout time.Duration) *DeleteDistroXV1ByNameParams {
	var (
		forcedDefault = bool(false)
	)
	return &DeleteDistroXV1ByNameParams{
		Forced: &forcedDefault,

		timeout: timeout,
	}
}

// NewDeleteDistroXV1ByNameParamsWithContext creates a new DeleteDistroXV1ByNameParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteDistroXV1ByNameParamsWithContext(ctx context.Context) *DeleteDistroXV1ByNameParams {
	var (
		forcedDefault = bool(false)
	)
	return &DeleteDistroXV1ByNameParams{
		Forced: &forcedDefault,

		Context: ctx,
	}
}

// NewDeleteDistroXV1ByNameParamsWithHTTPClient creates a new DeleteDistroXV1ByNameParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteDistroXV1ByNameParamsWithHTTPClient(client *http.Client) *DeleteDistroXV1ByNameParams {
	var (
		forcedDefault = bool(false)
	)
	return &DeleteDistroXV1ByNameParams{
		Forced:     &forcedDefault,
		HTTPClient: client,
	}
}

/*
DeleteDistroXV1ByNameParams contains all the parameters to send to the API endpoint
for the delete distro x v1 by name operation typically these are written to a http.Request
*/
type DeleteDistroXV1ByNameParams struct {

	/*Forced*/
	Forced *bool
	/*Name*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete distro x v1 by name params
func (o *DeleteDistroXV1ByNameParams) WithTimeout(timeout time.Duration) *DeleteDistroXV1ByNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete distro x v1 by name params
func (o *DeleteDistroXV1ByNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete distro x v1 by name params
func (o *DeleteDistroXV1ByNameParams) WithContext(ctx context.Context) *DeleteDistroXV1ByNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete distro x v1 by name params
func (o *DeleteDistroXV1ByNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete distro x v1 by name params
func (o *DeleteDistroXV1ByNameParams) WithHTTPClient(client *http.Client) *DeleteDistroXV1ByNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete distro x v1 by name params
func (o *DeleteDistroXV1ByNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithForced adds the forced to the delete distro x v1 by name params
func (o *DeleteDistroXV1ByNameParams) WithForced(forced *bool) *DeleteDistroXV1ByNameParams {
	o.SetForced(forced)
	return o
}

// SetForced adds the forced to the delete distro x v1 by name params
func (o *DeleteDistroXV1ByNameParams) SetForced(forced *bool) {
	o.Forced = forced
}

// WithName adds the name to the delete distro x v1 by name params
func (o *DeleteDistroXV1ByNameParams) WithName(name string) *DeleteDistroXV1ByNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the delete distro x v1 by name params
func (o *DeleteDistroXV1ByNameParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteDistroXV1ByNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
