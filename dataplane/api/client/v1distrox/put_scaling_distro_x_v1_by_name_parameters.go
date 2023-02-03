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

	model "github.com/hortonworks/cb-cli/dataplane/api/model"
)

// NewPutScalingDistroXV1ByNameParams creates a new PutScalingDistroXV1ByNameParams object
// with the default values initialized.
func NewPutScalingDistroXV1ByNameParams() *PutScalingDistroXV1ByNameParams {
	var ()
	return &PutScalingDistroXV1ByNameParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutScalingDistroXV1ByNameParamsWithTimeout creates a new PutScalingDistroXV1ByNameParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutScalingDistroXV1ByNameParamsWithTimeout(timeout time.Duration) *PutScalingDistroXV1ByNameParams {
	var ()
	return &PutScalingDistroXV1ByNameParams{

		timeout: timeout,
	}
}

// NewPutScalingDistroXV1ByNameParamsWithContext creates a new PutScalingDistroXV1ByNameParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutScalingDistroXV1ByNameParamsWithContext(ctx context.Context) *PutScalingDistroXV1ByNameParams {
	var ()
	return &PutScalingDistroXV1ByNameParams{

		Context: ctx,
	}
}

// NewPutScalingDistroXV1ByNameParamsWithHTTPClient creates a new PutScalingDistroXV1ByNameParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutScalingDistroXV1ByNameParamsWithHTTPClient(client *http.Client) *PutScalingDistroXV1ByNameParams {
	var ()
	return &PutScalingDistroXV1ByNameParams{
		HTTPClient: client,
	}
}

/*
PutScalingDistroXV1ByNameParams contains all the parameters to send to the API endpoint
for the put scaling distro x v1 by name operation typically these are written to a http.Request
*/
type PutScalingDistroXV1ByNameParams struct {

	/*Body*/
	Body *model.DistroXScaleV1Request
	/*Name*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put scaling distro x v1 by name params
func (o *PutScalingDistroXV1ByNameParams) WithTimeout(timeout time.Duration) *PutScalingDistroXV1ByNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put scaling distro x v1 by name params
func (o *PutScalingDistroXV1ByNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put scaling distro x v1 by name params
func (o *PutScalingDistroXV1ByNameParams) WithContext(ctx context.Context) *PutScalingDistroXV1ByNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put scaling distro x v1 by name params
func (o *PutScalingDistroXV1ByNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put scaling distro x v1 by name params
func (o *PutScalingDistroXV1ByNameParams) WithHTTPClient(client *http.Client) *PutScalingDistroXV1ByNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put scaling distro x v1 by name params
func (o *PutScalingDistroXV1ByNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put scaling distro x v1 by name params
func (o *PutScalingDistroXV1ByNameParams) WithBody(body *model.DistroXScaleV1Request) *PutScalingDistroXV1ByNameParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put scaling distro x v1 by name params
func (o *PutScalingDistroXV1ByNameParams) SetBody(body *model.DistroXScaleV1Request) {
	o.Body = body
}

// WithName adds the name to the put scaling distro x v1 by name params
func (o *PutScalingDistroXV1ByNameParams) WithName(name string) *PutScalingDistroXV1ByNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the put scaling distro x v1 by name params
func (o *PutScalingDistroXV1ByNameParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *PutScalingDistroXV1ByNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
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
