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

// NewStopDistroXV1ByNamesParams creates a new StopDistroXV1ByNamesParams object
// with the default values initialized.
func NewStopDistroXV1ByNamesParams() *StopDistroXV1ByNamesParams {
	var ()
	return &StopDistroXV1ByNamesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewStopDistroXV1ByNamesParamsWithTimeout creates a new StopDistroXV1ByNamesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewStopDistroXV1ByNamesParamsWithTimeout(timeout time.Duration) *StopDistroXV1ByNamesParams {
	var ()
	return &StopDistroXV1ByNamesParams{

		timeout: timeout,
	}
}

// NewStopDistroXV1ByNamesParamsWithContext creates a new StopDistroXV1ByNamesParams object
// with the default values initialized, and the ability to set a context for a request
func NewStopDistroXV1ByNamesParamsWithContext(ctx context.Context) *StopDistroXV1ByNamesParams {
	var ()
	return &StopDistroXV1ByNamesParams{

		Context: ctx,
	}
}

// NewStopDistroXV1ByNamesParamsWithHTTPClient creates a new StopDistroXV1ByNamesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewStopDistroXV1ByNamesParamsWithHTTPClient(client *http.Client) *StopDistroXV1ByNamesParams {
	var ()
	return &StopDistroXV1ByNamesParams{
		HTTPClient: client,
	}
}

/*
StopDistroXV1ByNamesParams contains all the parameters to send to the API endpoint
for the stop distro x v1 by names operation typically these are written to a http.Request
*/
type StopDistroXV1ByNamesParams struct {

	/*Names*/
	Names []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the stop distro x v1 by names params
func (o *StopDistroXV1ByNamesParams) WithTimeout(timeout time.Duration) *StopDistroXV1ByNamesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the stop distro x v1 by names params
func (o *StopDistroXV1ByNamesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the stop distro x v1 by names params
func (o *StopDistroXV1ByNamesParams) WithContext(ctx context.Context) *StopDistroXV1ByNamesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the stop distro x v1 by names params
func (o *StopDistroXV1ByNamesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the stop distro x v1 by names params
func (o *StopDistroXV1ByNamesParams) WithHTTPClient(client *http.Client) *StopDistroXV1ByNamesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the stop distro x v1 by names params
func (o *StopDistroXV1ByNamesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNames adds the names to the stop distro x v1 by names params
func (o *StopDistroXV1ByNamesParams) WithNames(names []string) *StopDistroXV1ByNamesParams {
	o.SetNames(names)
	return o
}

// SetNames adds the names to the stop distro x v1 by names params
func (o *StopDistroXV1ByNamesParams) SetNames(names []string) {
	o.Names = names
}

// WriteToRequest writes these params to a swagger request
func (o *StopDistroXV1ByNamesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	valuesNames := o.Names

	joinedNames := swag.JoinByFormat(valuesNames, "multi")
	// query array param names
	if err := r.SetQueryParam("names", joinedNames...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
