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

	strfmt "github.com/go-openapi/strfmt"
)

// NewListDistroXV1Params creates a new ListDistroXV1Params object
// with the default values initialized.
func NewListDistroXV1Params() *ListDistroXV1Params {
	var ()
	return &ListDistroXV1Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewListDistroXV1ParamsWithTimeout creates a new ListDistroXV1Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewListDistroXV1ParamsWithTimeout(timeout time.Duration) *ListDistroXV1Params {
	var ()
	return &ListDistroXV1Params{

		timeout: timeout,
	}
}

// NewListDistroXV1ParamsWithContext creates a new ListDistroXV1Params object
// with the default values initialized, and the ability to set a context for a request
func NewListDistroXV1ParamsWithContext(ctx context.Context) *ListDistroXV1Params {
	var ()
	return &ListDistroXV1Params{

		Context: ctx,
	}
}

// NewListDistroXV1ParamsWithHTTPClient creates a new ListDistroXV1Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListDistroXV1ParamsWithHTTPClient(client *http.Client) *ListDistroXV1Params {
	var ()
	return &ListDistroXV1Params{
		HTTPClient: client,
	}
}

/*
ListDistroXV1Params contains all the parameters to send to the API endpoint
for the list distro x v1 operation typically these are written to a http.Request
*/
type ListDistroXV1Params struct {

	/*EnvironmentCrn*/
	EnvironmentCrn *string
	/*EnvironmentName*/
	EnvironmentName *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list distro x v1 params
func (o *ListDistroXV1Params) WithTimeout(timeout time.Duration) *ListDistroXV1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list distro x v1 params
func (o *ListDistroXV1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list distro x v1 params
func (o *ListDistroXV1Params) WithContext(ctx context.Context) *ListDistroXV1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list distro x v1 params
func (o *ListDistroXV1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list distro x v1 params
func (o *ListDistroXV1Params) WithHTTPClient(client *http.Client) *ListDistroXV1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list distro x v1 params
func (o *ListDistroXV1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEnvironmentCrn adds the environmentCrn to the list distro x v1 params
func (o *ListDistroXV1Params) WithEnvironmentCrn(environmentCrn *string) *ListDistroXV1Params {
	o.SetEnvironmentCrn(environmentCrn)
	return o
}

// SetEnvironmentCrn adds the environmentCrn to the list distro x v1 params
func (o *ListDistroXV1Params) SetEnvironmentCrn(environmentCrn *string) {
	o.EnvironmentCrn = environmentCrn
}

// WithEnvironmentName adds the environmentName to the list distro x v1 params
func (o *ListDistroXV1Params) WithEnvironmentName(environmentName *string) *ListDistroXV1Params {
	o.SetEnvironmentName(environmentName)
	return o
}

// SetEnvironmentName adds the environmentName to the list distro x v1 params
func (o *ListDistroXV1Params) SetEnvironmentName(environmentName *string) {
	o.EnvironmentName = environmentName
}

// WriteToRequest writes these params to a swagger request
func (o *ListDistroXV1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.EnvironmentCrn != nil {

		// query param environmentCrn
		var qrEnvironmentCrn string
		if o.EnvironmentCrn != nil {
			qrEnvironmentCrn = *o.EnvironmentCrn
		}
		qEnvironmentCrn := qrEnvironmentCrn
		if qEnvironmentCrn != "" {
			if err := r.SetQueryParam("environmentCrn", qEnvironmentCrn); err != nil {
				return err
			}
		}

	}

	if o.EnvironmentName != nil {

		// query param environmentName
		var qrEnvironmentName string
		if o.EnvironmentName != nil {
			qrEnvironmentName = *o.EnvironmentName
		}
		qEnvironmentName := qrEnvironmentName
		if qEnvironmentName != "" {
			if err := r.SetQueryParam("environmentName", qEnvironmentName); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
