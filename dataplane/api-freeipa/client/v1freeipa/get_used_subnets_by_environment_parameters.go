// Code generated by go-swagger; DO NOT EDIT.

package v1freeipa

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

// NewGetUsedSubnetsByEnvironmentParams creates a new GetUsedSubnetsByEnvironmentParams object
// with the default values initialized.
func NewGetUsedSubnetsByEnvironmentParams() *GetUsedSubnetsByEnvironmentParams {
	var ()
	return &GetUsedSubnetsByEnvironmentParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetUsedSubnetsByEnvironmentParamsWithTimeout creates a new GetUsedSubnetsByEnvironmentParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetUsedSubnetsByEnvironmentParamsWithTimeout(timeout time.Duration) *GetUsedSubnetsByEnvironmentParams {
	var ()
	return &GetUsedSubnetsByEnvironmentParams{

		timeout: timeout,
	}
}

// NewGetUsedSubnetsByEnvironmentParamsWithContext creates a new GetUsedSubnetsByEnvironmentParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetUsedSubnetsByEnvironmentParamsWithContext(ctx context.Context) *GetUsedSubnetsByEnvironmentParams {
	var ()
	return &GetUsedSubnetsByEnvironmentParams{

		Context: ctx,
	}
}

// NewGetUsedSubnetsByEnvironmentParamsWithHTTPClient creates a new GetUsedSubnetsByEnvironmentParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetUsedSubnetsByEnvironmentParamsWithHTTPClient(client *http.Client) *GetUsedSubnetsByEnvironmentParams {
	var ()
	return &GetUsedSubnetsByEnvironmentParams{
		HTTPClient: client,
	}
}

/*
GetUsedSubnetsByEnvironmentParams contains all the parameters to send to the API endpoint
for the get used subnets by environment operation typically these are written to a http.Request
*/
type GetUsedSubnetsByEnvironmentParams struct {

	/*EnvironmentCrn*/
	EnvironmentCrn *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get used subnets by environment params
func (o *GetUsedSubnetsByEnvironmentParams) WithTimeout(timeout time.Duration) *GetUsedSubnetsByEnvironmentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get used subnets by environment params
func (o *GetUsedSubnetsByEnvironmentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get used subnets by environment params
func (o *GetUsedSubnetsByEnvironmentParams) WithContext(ctx context.Context) *GetUsedSubnetsByEnvironmentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get used subnets by environment params
func (o *GetUsedSubnetsByEnvironmentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get used subnets by environment params
func (o *GetUsedSubnetsByEnvironmentParams) WithHTTPClient(client *http.Client) *GetUsedSubnetsByEnvironmentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get used subnets by environment params
func (o *GetUsedSubnetsByEnvironmentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEnvironmentCrn adds the environmentCrn to the get used subnets by environment params
func (o *GetUsedSubnetsByEnvironmentParams) WithEnvironmentCrn(environmentCrn *string) *GetUsedSubnetsByEnvironmentParams {
	o.SetEnvironmentCrn(environmentCrn)
	return o
}

// SetEnvironmentCrn adds the environmentCrn to the get used subnets by environment params
func (o *GetUsedSubnetsByEnvironmentParams) SetEnvironmentCrn(environmentCrn *string) {
	o.EnvironmentCrn = environmentCrn
}

// WriteToRequest writes these params to a swagger request
func (o *GetUsedSubnetsByEnvironmentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
