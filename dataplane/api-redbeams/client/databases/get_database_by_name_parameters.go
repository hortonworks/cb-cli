// Code generated by go-swagger; DO NOT EDIT.

package databases

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

// NewGetDatabaseByNameParams creates a new GetDatabaseByNameParams object
// with the default values initialized.
func NewGetDatabaseByNameParams() *GetDatabaseByNameParams {
	var ()
	return &GetDatabaseByNameParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDatabaseByNameParamsWithTimeout creates a new GetDatabaseByNameParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDatabaseByNameParamsWithTimeout(timeout time.Duration) *GetDatabaseByNameParams {
	var ()
	return &GetDatabaseByNameParams{

		timeout: timeout,
	}
}

// NewGetDatabaseByNameParamsWithContext creates a new GetDatabaseByNameParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDatabaseByNameParamsWithContext(ctx context.Context) *GetDatabaseByNameParams {
	var ()
	return &GetDatabaseByNameParams{

		Context: ctx,
	}
}

// NewGetDatabaseByNameParamsWithHTTPClient creates a new GetDatabaseByNameParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDatabaseByNameParamsWithHTTPClient(client *http.Client) *GetDatabaseByNameParams {
	var ()
	return &GetDatabaseByNameParams{
		HTTPClient: client,
	}
}

/*
GetDatabaseByNameParams contains all the parameters to send to the API endpoint
for the get database by name operation typically these are written to a http.Request
*/
type GetDatabaseByNameParams struct {

	/*EnvironmentCrn
	  CRN of the environment of the database(s)

	*/
	EnvironmentCrn string
	/*Name
	  Name of the database

	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get database by name params
func (o *GetDatabaseByNameParams) WithTimeout(timeout time.Duration) *GetDatabaseByNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get database by name params
func (o *GetDatabaseByNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get database by name params
func (o *GetDatabaseByNameParams) WithContext(ctx context.Context) *GetDatabaseByNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get database by name params
func (o *GetDatabaseByNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get database by name params
func (o *GetDatabaseByNameParams) WithHTTPClient(client *http.Client) *GetDatabaseByNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get database by name params
func (o *GetDatabaseByNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEnvironmentCrn adds the environmentCrn to the get database by name params
func (o *GetDatabaseByNameParams) WithEnvironmentCrn(environmentCrn string) *GetDatabaseByNameParams {
	o.SetEnvironmentCrn(environmentCrn)
	return o
}

// SetEnvironmentCrn adds the environmentCrn to the get database by name params
func (o *GetDatabaseByNameParams) SetEnvironmentCrn(environmentCrn string) {
	o.EnvironmentCrn = environmentCrn
}

// WithName adds the name to the get database by name params
func (o *GetDatabaseByNameParams) WithName(name string) *GetDatabaseByNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get database by name params
func (o *GetDatabaseByNameParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *GetDatabaseByNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
