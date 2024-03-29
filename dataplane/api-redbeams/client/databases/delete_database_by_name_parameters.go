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

// NewDeleteDatabaseByNameParams creates a new DeleteDatabaseByNameParams object
// with the default values initialized.
func NewDeleteDatabaseByNameParams() *DeleteDatabaseByNameParams {
	var ()
	return &DeleteDatabaseByNameParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteDatabaseByNameParamsWithTimeout creates a new DeleteDatabaseByNameParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteDatabaseByNameParamsWithTimeout(timeout time.Duration) *DeleteDatabaseByNameParams {
	var ()
	return &DeleteDatabaseByNameParams{

		timeout: timeout,
	}
}

// NewDeleteDatabaseByNameParamsWithContext creates a new DeleteDatabaseByNameParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteDatabaseByNameParamsWithContext(ctx context.Context) *DeleteDatabaseByNameParams {
	var ()
	return &DeleteDatabaseByNameParams{

		Context: ctx,
	}
}

// NewDeleteDatabaseByNameParamsWithHTTPClient creates a new DeleteDatabaseByNameParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteDatabaseByNameParamsWithHTTPClient(client *http.Client) *DeleteDatabaseByNameParams {
	var ()
	return &DeleteDatabaseByNameParams{
		HTTPClient: client,
	}
}

/*
DeleteDatabaseByNameParams contains all the parameters to send to the API endpoint
for the delete database by name operation typically these are written to a http.Request
*/
type DeleteDatabaseByNameParams struct {

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

// WithTimeout adds the timeout to the delete database by name params
func (o *DeleteDatabaseByNameParams) WithTimeout(timeout time.Duration) *DeleteDatabaseByNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete database by name params
func (o *DeleteDatabaseByNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete database by name params
func (o *DeleteDatabaseByNameParams) WithContext(ctx context.Context) *DeleteDatabaseByNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete database by name params
func (o *DeleteDatabaseByNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete database by name params
func (o *DeleteDatabaseByNameParams) WithHTTPClient(client *http.Client) *DeleteDatabaseByNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete database by name params
func (o *DeleteDatabaseByNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEnvironmentCrn adds the environmentCrn to the delete database by name params
func (o *DeleteDatabaseByNameParams) WithEnvironmentCrn(environmentCrn string) *DeleteDatabaseByNameParams {
	o.SetEnvironmentCrn(environmentCrn)
	return o
}

// SetEnvironmentCrn adds the environmentCrn to the delete database by name params
func (o *DeleteDatabaseByNameParams) SetEnvironmentCrn(environmentCrn string) {
	o.EnvironmentCrn = environmentCrn
}

// WithName adds the name to the delete database by name params
func (o *DeleteDatabaseByNameParams) WithName(name string) *DeleteDatabaseByNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the delete database by name params
func (o *DeleteDatabaseByNameParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteDatabaseByNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
