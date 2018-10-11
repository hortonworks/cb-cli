// Code generated by go-swagger; DO NOT EDIT.

package v2stacks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeletePublicStackV2Params creates a new DeletePublicStackV2Params object
// with the default values initialized.
func NewDeletePublicStackV2Params() *DeletePublicStackV2Params {
	var (
		deleteDependenciesDefault = bool(false)
		forcedDefault             = bool(false)
	)
	return &DeletePublicStackV2Params{
		DeleteDependencies: &deleteDependenciesDefault,
		Forced:             &forcedDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewDeletePublicStackV2ParamsWithTimeout creates a new DeletePublicStackV2Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeletePublicStackV2ParamsWithTimeout(timeout time.Duration) *DeletePublicStackV2Params {
	var (
		deleteDependenciesDefault = bool(false)
		forcedDefault             = bool(false)
	)
	return &DeletePublicStackV2Params{
		DeleteDependencies: &deleteDependenciesDefault,
		Forced:             &forcedDefault,

		timeout: timeout,
	}
}

// NewDeletePublicStackV2ParamsWithContext creates a new DeletePublicStackV2Params object
// with the default values initialized, and the ability to set a context for a request
func NewDeletePublicStackV2ParamsWithContext(ctx context.Context) *DeletePublicStackV2Params {
	var (
		deleteDependenciesDefault = bool(false)
		forcedDefault             = bool(false)
	)
	return &DeletePublicStackV2Params{
		DeleteDependencies: &deleteDependenciesDefault,
		Forced:             &forcedDefault,

		Context: ctx,
	}
}

// NewDeletePublicStackV2ParamsWithHTTPClient creates a new DeletePublicStackV2Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeletePublicStackV2ParamsWithHTTPClient(client *http.Client) *DeletePublicStackV2Params {
	var (
		deleteDependenciesDefault = bool(false)
		forcedDefault             = bool(false)
	)
	return &DeletePublicStackV2Params{
		DeleteDependencies: &deleteDependenciesDefault,
		Forced:             &forcedDefault,
		HTTPClient:         client,
	}
}

/*DeletePublicStackV2Params contains all the parameters to send to the API endpoint
for the delete public stack v2 operation typically these are written to a http.Request
*/
type DeletePublicStackV2Params struct {

	/*DeleteDependencies*/
	DeleteDependencies *bool
	/*Forced*/
	Forced *bool
	/*Name*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete public stack v2 params
func (o *DeletePublicStackV2Params) WithTimeout(timeout time.Duration) *DeletePublicStackV2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete public stack v2 params
func (o *DeletePublicStackV2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete public stack v2 params
func (o *DeletePublicStackV2Params) WithContext(ctx context.Context) *DeletePublicStackV2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete public stack v2 params
func (o *DeletePublicStackV2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete public stack v2 params
func (o *DeletePublicStackV2Params) WithHTTPClient(client *http.Client) *DeletePublicStackV2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete public stack v2 params
func (o *DeletePublicStackV2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDeleteDependencies adds the deleteDependencies to the delete public stack v2 params
func (o *DeletePublicStackV2Params) WithDeleteDependencies(deleteDependencies *bool) *DeletePublicStackV2Params {
	o.SetDeleteDependencies(deleteDependencies)
	return o
}

// SetDeleteDependencies adds the deleteDependencies to the delete public stack v2 params
func (o *DeletePublicStackV2Params) SetDeleteDependencies(deleteDependencies *bool) {
	o.DeleteDependencies = deleteDependencies
}

// WithForced adds the forced to the delete public stack v2 params
func (o *DeletePublicStackV2Params) WithForced(forced *bool) *DeletePublicStackV2Params {
	o.SetForced(forced)
	return o
}

// SetForced adds the forced to the delete public stack v2 params
func (o *DeletePublicStackV2Params) SetForced(forced *bool) {
	o.Forced = forced
}

// WithName adds the name to the delete public stack v2 params
func (o *DeletePublicStackV2Params) WithName(name string) *DeletePublicStackV2Params {
	o.SetName(name)
	return o
}

// SetName adds the name to the delete public stack v2 params
func (o *DeletePublicStackV2Params) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *DeletePublicStackV2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DeleteDependencies != nil {

		// query param deleteDependencies
		var qrDeleteDependencies bool
		if o.DeleteDependencies != nil {
			qrDeleteDependencies = *o.DeleteDependencies
		}
		qDeleteDependencies := swag.FormatBool(qrDeleteDependencies)
		if qDeleteDependencies != "" {
			if err := r.SetQueryParam("deleteDependencies", qDeleteDependencies); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}