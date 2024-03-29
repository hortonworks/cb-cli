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

// NewGetDatalakeEventsZipForResourceParams creates a new GetDatalakeEventsZipForResourceParams object
// with the default values initialized.
func NewGetDatalakeEventsZipForResourceParams() *GetDatalakeEventsZipForResourceParams {
	var ()
	return &GetDatalakeEventsZipForResourceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDatalakeEventsZipForResourceParamsWithTimeout creates a new GetDatalakeEventsZipForResourceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDatalakeEventsZipForResourceParamsWithTimeout(timeout time.Duration) *GetDatalakeEventsZipForResourceParams {
	var ()
	return &GetDatalakeEventsZipForResourceParams{

		timeout: timeout,
	}
}

// NewGetDatalakeEventsZipForResourceParamsWithContext creates a new GetDatalakeEventsZipForResourceParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDatalakeEventsZipForResourceParamsWithContext(ctx context.Context) *GetDatalakeEventsZipForResourceParams {
	var ()
	return &GetDatalakeEventsZipForResourceParams{

		Context: ctx,
	}
}

// NewGetDatalakeEventsZipForResourceParamsWithHTTPClient creates a new GetDatalakeEventsZipForResourceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDatalakeEventsZipForResourceParamsWithHTTPClient(client *http.Client) *GetDatalakeEventsZipForResourceParams {
	var ()
	return &GetDatalakeEventsZipForResourceParams{
		HTTPClient: client,
	}
}

/*
GetDatalakeEventsZipForResourceParams contains all the parameters to send to the API endpoint
for the get datalake events zip for resource operation typically these are written to a http.Request
*/
type GetDatalakeEventsZipForResourceParams struct {

	/*EnvironmentCrn*/
	EnvironmentCrn string
	/*Types*/
	Types []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get datalake events zip for resource params
func (o *GetDatalakeEventsZipForResourceParams) WithTimeout(timeout time.Duration) *GetDatalakeEventsZipForResourceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get datalake events zip for resource params
func (o *GetDatalakeEventsZipForResourceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get datalake events zip for resource params
func (o *GetDatalakeEventsZipForResourceParams) WithContext(ctx context.Context) *GetDatalakeEventsZipForResourceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get datalake events zip for resource params
func (o *GetDatalakeEventsZipForResourceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get datalake events zip for resource params
func (o *GetDatalakeEventsZipForResourceParams) WithHTTPClient(client *http.Client) *GetDatalakeEventsZipForResourceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get datalake events zip for resource params
func (o *GetDatalakeEventsZipForResourceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEnvironmentCrn adds the environmentCrn to the get datalake events zip for resource params
func (o *GetDatalakeEventsZipForResourceParams) WithEnvironmentCrn(environmentCrn string) *GetDatalakeEventsZipForResourceParams {
	o.SetEnvironmentCrn(environmentCrn)
	return o
}

// SetEnvironmentCrn adds the environmentCrn to the get datalake events zip for resource params
func (o *GetDatalakeEventsZipForResourceParams) SetEnvironmentCrn(environmentCrn string) {
	o.EnvironmentCrn = environmentCrn
}

// WithTypes adds the types to the get datalake events zip for resource params
func (o *GetDatalakeEventsZipForResourceParams) WithTypes(types []string) *GetDatalakeEventsZipForResourceParams {
	o.SetTypes(types)
	return o
}

// SetTypes adds the types to the get datalake events zip for resource params
func (o *GetDatalakeEventsZipForResourceParams) SetTypes(types []string) {
	o.Types = types
}

// WriteToRequest writes these params to a swagger request
func (o *GetDatalakeEventsZipForResourceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param environmentCrn
	qrEnvironmentCrn := o.EnvironmentCrn
	qEnvironmentCrn := qrEnvironmentCrn
	if qEnvironmentCrn != "" {
		if err := r.SetQueryParam("environmentCrn", qEnvironmentCrn); err != nil {
			return err
		}
	}

	valuesTypes := o.Types

	joinedTypes := swag.JoinByFormat(valuesTypes, "multi")
	// query array param types
	if err := r.SetQueryParam("types", joinedTypes...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
