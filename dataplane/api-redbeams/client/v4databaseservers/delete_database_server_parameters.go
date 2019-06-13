// Code generated by go-swagger; DO NOT EDIT.

package v4databaseservers

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

// NewDeleteDatabaseServerParams creates a new DeleteDatabaseServerParams object
// with the default values initialized.
func NewDeleteDatabaseServerParams() *DeleteDatabaseServerParams {
	var ()
	return &DeleteDatabaseServerParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteDatabaseServerParamsWithTimeout creates a new DeleteDatabaseServerParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteDatabaseServerParamsWithTimeout(timeout time.Duration) *DeleteDatabaseServerParams {
	var ()
	return &DeleteDatabaseServerParams{

		timeout: timeout,
	}
}

// NewDeleteDatabaseServerParamsWithContext creates a new DeleteDatabaseServerParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteDatabaseServerParamsWithContext(ctx context.Context) *DeleteDatabaseServerParams {
	var ()
	return &DeleteDatabaseServerParams{

		Context: ctx,
	}
}

// NewDeleteDatabaseServerParamsWithHTTPClient creates a new DeleteDatabaseServerParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteDatabaseServerParamsWithHTTPClient(client *http.Client) *DeleteDatabaseServerParams {
	var ()
	return &DeleteDatabaseServerParams{
		HTTPClient: client,
	}
}

/*DeleteDatabaseServerParams contains all the parameters to send to the API endpoint
for the delete database server operation typically these are written to a http.Request
*/
type DeleteDatabaseServerParams struct {

	/*EnvironmentID*/
	EnvironmentID string
	/*Name*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete database server params
func (o *DeleteDatabaseServerParams) WithTimeout(timeout time.Duration) *DeleteDatabaseServerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete database server params
func (o *DeleteDatabaseServerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete database server params
func (o *DeleteDatabaseServerParams) WithContext(ctx context.Context) *DeleteDatabaseServerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete database server params
func (o *DeleteDatabaseServerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete database server params
func (o *DeleteDatabaseServerParams) WithHTTPClient(client *http.Client) *DeleteDatabaseServerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete database server params
func (o *DeleteDatabaseServerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEnvironmentID adds the environmentID to the delete database server params
func (o *DeleteDatabaseServerParams) WithEnvironmentID(environmentID string) *DeleteDatabaseServerParams {
	o.SetEnvironmentID(environmentID)
	return o
}

// SetEnvironmentID adds the environmentId to the delete database server params
func (o *DeleteDatabaseServerParams) SetEnvironmentID(environmentID string) {
	o.EnvironmentID = environmentID
}

// WithName adds the name to the delete database server params
func (o *DeleteDatabaseServerParams) WithName(name string) *DeleteDatabaseServerParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the delete database server params
func (o *DeleteDatabaseServerParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteDatabaseServerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param environmentId
	qrEnvironmentID := o.EnvironmentID
	qEnvironmentID := qrEnvironmentID
	if qEnvironmentID != "" {
		if err := r.SetQueryParam("environmentId", qEnvironmentID); err != nil {
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
