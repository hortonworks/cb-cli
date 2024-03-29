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

// NewRepairDistroXV1ByNameParams creates a new RepairDistroXV1ByNameParams object
// with the default values initialized.
func NewRepairDistroXV1ByNameParams() *RepairDistroXV1ByNameParams {
	var ()
	return &RepairDistroXV1ByNameParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRepairDistroXV1ByNameParamsWithTimeout creates a new RepairDistroXV1ByNameParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRepairDistroXV1ByNameParamsWithTimeout(timeout time.Duration) *RepairDistroXV1ByNameParams {
	var ()
	return &RepairDistroXV1ByNameParams{

		timeout: timeout,
	}
}

// NewRepairDistroXV1ByNameParamsWithContext creates a new RepairDistroXV1ByNameParams object
// with the default values initialized, and the ability to set a context for a request
func NewRepairDistroXV1ByNameParamsWithContext(ctx context.Context) *RepairDistroXV1ByNameParams {
	var ()
	return &RepairDistroXV1ByNameParams{

		Context: ctx,
	}
}

// NewRepairDistroXV1ByNameParamsWithHTTPClient creates a new RepairDistroXV1ByNameParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRepairDistroXV1ByNameParamsWithHTTPClient(client *http.Client) *RepairDistroXV1ByNameParams {
	var ()
	return &RepairDistroXV1ByNameParams{
		HTTPClient: client,
	}
}

/*
RepairDistroXV1ByNameParams contains all the parameters to send to the API endpoint
for the repair distro x v1 by name operation typically these are written to a http.Request
*/
type RepairDistroXV1ByNameParams struct {

	/*Body*/
	Body *model.DistroXRepairV1Request
	/*Name*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the repair distro x v1 by name params
func (o *RepairDistroXV1ByNameParams) WithTimeout(timeout time.Duration) *RepairDistroXV1ByNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the repair distro x v1 by name params
func (o *RepairDistroXV1ByNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the repair distro x v1 by name params
func (o *RepairDistroXV1ByNameParams) WithContext(ctx context.Context) *RepairDistroXV1ByNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the repair distro x v1 by name params
func (o *RepairDistroXV1ByNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the repair distro x v1 by name params
func (o *RepairDistroXV1ByNameParams) WithHTTPClient(client *http.Client) *RepairDistroXV1ByNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the repair distro x v1 by name params
func (o *RepairDistroXV1ByNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the repair distro x v1 by name params
func (o *RepairDistroXV1ByNameParams) WithBody(body *model.DistroXRepairV1Request) *RepairDistroXV1ByNameParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the repair distro x v1 by name params
func (o *RepairDistroXV1ByNameParams) SetBody(body *model.DistroXRepairV1Request) {
	o.Body = body
}

// WithName adds the name to the repair distro x v1 by name params
func (o *RepairDistroXV1ByNameParams) WithName(name string) *RepairDistroXV1ByNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the repair distro x v1 by name params
func (o *RepairDistroXV1ByNameParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *RepairDistroXV1ByNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
