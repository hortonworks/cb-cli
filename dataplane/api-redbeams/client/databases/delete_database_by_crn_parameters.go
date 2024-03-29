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

// NewDeleteDatabaseByCrnParams creates a new DeleteDatabaseByCrnParams object
// with the default values initialized.
func NewDeleteDatabaseByCrnParams() *DeleteDatabaseByCrnParams {
	var ()
	return &DeleteDatabaseByCrnParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteDatabaseByCrnParamsWithTimeout creates a new DeleteDatabaseByCrnParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteDatabaseByCrnParamsWithTimeout(timeout time.Duration) *DeleteDatabaseByCrnParams {
	var ()
	return &DeleteDatabaseByCrnParams{

		timeout: timeout,
	}
}

// NewDeleteDatabaseByCrnParamsWithContext creates a new DeleteDatabaseByCrnParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteDatabaseByCrnParamsWithContext(ctx context.Context) *DeleteDatabaseByCrnParams {
	var ()
	return &DeleteDatabaseByCrnParams{

		Context: ctx,
	}
}

// NewDeleteDatabaseByCrnParamsWithHTTPClient creates a new DeleteDatabaseByCrnParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteDatabaseByCrnParamsWithHTTPClient(client *http.Client) *DeleteDatabaseByCrnParams {
	var ()
	return &DeleteDatabaseByCrnParams{
		HTTPClient: client,
	}
}

/*
DeleteDatabaseByCrnParams contains all the parameters to send to the API endpoint
for the delete database by crn operation typically these are written to a http.Request
*/
type DeleteDatabaseByCrnParams struct {

	/*Crn
	  CRN of the database

	*/
	Crn string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete database by crn params
func (o *DeleteDatabaseByCrnParams) WithTimeout(timeout time.Duration) *DeleteDatabaseByCrnParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete database by crn params
func (o *DeleteDatabaseByCrnParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete database by crn params
func (o *DeleteDatabaseByCrnParams) WithContext(ctx context.Context) *DeleteDatabaseByCrnParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete database by crn params
func (o *DeleteDatabaseByCrnParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete database by crn params
func (o *DeleteDatabaseByCrnParams) WithHTTPClient(client *http.Client) *DeleteDatabaseByCrnParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete database by crn params
func (o *DeleteDatabaseByCrnParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCrn adds the crn to the delete database by crn params
func (o *DeleteDatabaseByCrnParams) WithCrn(crn string) *DeleteDatabaseByCrnParams {
	o.SetCrn(crn)
	return o
}

// SetCrn adds the crn to the delete database by crn params
func (o *DeleteDatabaseByCrnParams) SetCrn(crn string) {
	o.Crn = crn
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteDatabaseByCrnParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param crn
	if err := r.SetPathParam("crn", o.Crn); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
