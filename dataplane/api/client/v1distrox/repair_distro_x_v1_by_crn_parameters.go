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

// NewRepairDistroXV1ByCrnParams creates a new RepairDistroXV1ByCrnParams object
// with the default values initialized.
func NewRepairDistroXV1ByCrnParams() *RepairDistroXV1ByCrnParams {
	var ()
	return &RepairDistroXV1ByCrnParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRepairDistroXV1ByCrnParamsWithTimeout creates a new RepairDistroXV1ByCrnParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRepairDistroXV1ByCrnParamsWithTimeout(timeout time.Duration) *RepairDistroXV1ByCrnParams {
	var ()
	return &RepairDistroXV1ByCrnParams{

		timeout: timeout,
	}
}

// NewRepairDistroXV1ByCrnParamsWithContext creates a new RepairDistroXV1ByCrnParams object
// with the default values initialized, and the ability to set a context for a request
func NewRepairDistroXV1ByCrnParamsWithContext(ctx context.Context) *RepairDistroXV1ByCrnParams {
	var ()
	return &RepairDistroXV1ByCrnParams{

		Context: ctx,
	}
}

// NewRepairDistroXV1ByCrnParamsWithHTTPClient creates a new RepairDistroXV1ByCrnParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRepairDistroXV1ByCrnParamsWithHTTPClient(client *http.Client) *RepairDistroXV1ByCrnParams {
	var ()
	return &RepairDistroXV1ByCrnParams{
		HTTPClient: client,
	}
}

/*
RepairDistroXV1ByCrnParams contains all the parameters to send to the API endpoint
for the repair distro x v1 by crn operation typically these are written to a http.Request
*/
type RepairDistroXV1ByCrnParams struct {

	/*Body*/
	Body *model.DistroXRepairV1Request
	/*Crn*/
	Crn string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the repair distro x v1 by crn params
func (o *RepairDistroXV1ByCrnParams) WithTimeout(timeout time.Duration) *RepairDistroXV1ByCrnParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the repair distro x v1 by crn params
func (o *RepairDistroXV1ByCrnParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the repair distro x v1 by crn params
func (o *RepairDistroXV1ByCrnParams) WithContext(ctx context.Context) *RepairDistroXV1ByCrnParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the repair distro x v1 by crn params
func (o *RepairDistroXV1ByCrnParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the repair distro x v1 by crn params
func (o *RepairDistroXV1ByCrnParams) WithHTTPClient(client *http.Client) *RepairDistroXV1ByCrnParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the repair distro x v1 by crn params
func (o *RepairDistroXV1ByCrnParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the repair distro x v1 by crn params
func (o *RepairDistroXV1ByCrnParams) WithBody(body *model.DistroXRepairV1Request) *RepairDistroXV1ByCrnParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the repair distro x v1 by crn params
func (o *RepairDistroXV1ByCrnParams) SetBody(body *model.DistroXRepairV1Request) {
	o.Body = body
}

// WithCrn adds the crn to the repair distro x v1 by crn params
func (o *RepairDistroXV1ByCrnParams) WithCrn(crn string) *RepairDistroXV1ByCrnParams {
	o.SetCrn(crn)
	return o
}

// SetCrn adds the crn to the repair distro x v1 by crn params
func (o *RepairDistroXV1ByCrnParams) SetCrn(crn string) {
	o.Crn = crn
}

// WriteToRequest writes these params to a swagger request
func (o *RepairDistroXV1ByCrnParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param crn
	if err := r.SetPathParam("crn", o.Crn); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
