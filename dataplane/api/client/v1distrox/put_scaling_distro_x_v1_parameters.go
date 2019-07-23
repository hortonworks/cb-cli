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

// NewPutScalingDistroXV1Params creates a new PutScalingDistroXV1Params object
// with the default values initialized.
func NewPutScalingDistroXV1Params() *PutScalingDistroXV1Params {
	var ()
	return &PutScalingDistroXV1Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutScalingDistroXV1ParamsWithTimeout creates a new PutScalingDistroXV1Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutScalingDistroXV1ParamsWithTimeout(timeout time.Duration) *PutScalingDistroXV1Params {
	var ()
	return &PutScalingDistroXV1Params{

		timeout: timeout,
	}
}

// NewPutScalingDistroXV1ParamsWithContext creates a new PutScalingDistroXV1Params object
// with the default values initialized, and the ability to set a context for a request
func NewPutScalingDistroXV1ParamsWithContext(ctx context.Context) *PutScalingDistroXV1Params {
	var ()
	return &PutScalingDistroXV1Params{

		Context: ctx,
	}
}

// NewPutScalingDistroXV1ParamsWithHTTPClient creates a new PutScalingDistroXV1Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutScalingDistroXV1ParamsWithHTTPClient(client *http.Client) *PutScalingDistroXV1Params {
	var ()
	return &PutScalingDistroXV1Params{
		HTTPClient: client,
	}
}

/*PutScalingDistroXV1Params contains all the parameters to send to the API endpoint
for the put scaling distro x v1 operation typically these are written to a http.Request
*/
type PutScalingDistroXV1Params struct {

	/*Body*/
	Body *model.DistroXScaleV1Request
	/*Name*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put scaling distro x v1 params
func (o *PutScalingDistroXV1Params) WithTimeout(timeout time.Duration) *PutScalingDistroXV1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put scaling distro x v1 params
func (o *PutScalingDistroXV1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put scaling distro x v1 params
func (o *PutScalingDistroXV1Params) WithContext(ctx context.Context) *PutScalingDistroXV1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put scaling distro x v1 params
func (o *PutScalingDistroXV1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put scaling distro x v1 params
func (o *PutScalingDistroXV1Params) WithHTTPClient(client *http.Client) *PutScalingDistroXV1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put scaling distro x v1 params
func (o *PutScalingDistroXV1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put scaling distro x v1 params
func (o *PutScalingDistroXV1Params) WithBody(body *model.DistroXScaleV1Request) *PutScalingDistroXV1Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put scaling distro x v1 params
func (o *PutScalingDistroXV1Params) SetBody(body *model.DistroXScaleV1Request) {
	o.Body = body
}

// WithName adds the name to the put scaling distro x v1 params
func (o *PutScalingDistroXV1Params) WithName(name string) *PutScalingDistroXV1Params {
	o.SetName(name)
	return o
}

// SetName adds the name to the put scaling distro x v1 params
func (o *PutScalingDistroXV1Params) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *PutScalingDistroXV1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
