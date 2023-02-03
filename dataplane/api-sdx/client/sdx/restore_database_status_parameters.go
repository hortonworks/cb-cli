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

	strfmt "github.com/go-openapi/strfmt"
)

// NewRestoreDatabaseStatusParams creates a new RestoreDatabaseStatusParams object
// with the default values initialized.
func NewRestoreDatabaseStatusParams() *RestoreDatabaseStatusParams {
	var ()
	return &RestoreDatabaseStatusParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRestoreDatabaseStatusParamsWithTimeout creates a new RestoreDatabaseStatusParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRestoreDatabaseStatusParamsWithTimeout(timeout time.Duration) *RestoreDatabaseStatusParams {
	var ()
	return &RestoreDatabaseStatusParams{

		timeout: timeout,
	}
}

// NewRestoreDatabaseStatusParamsWithContext creates a new RestoreDatabaseStatusParams object
// with the default values initialized, and the ability to set a context for a request
func NewRestoreDatabaseStatusParamsWithContext(ctx context.Context) *RestoreDatabaseStatusParams {
	var ()
	return &RestoreDatabaseStatusParams{

		Context: ctx,
	}
}

// NewRestoreDatabaseStatusParamsWithHTTPClient creates a new RestoreDatabaseStatusParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRestoreDatabaseStatusParamsWithHTTPClient(client *http.Client) *RestoreDatabaseStatusParams {
	var ()
	return &RestoreDatabaseStatusParams{
		HTTPClient: client,
	}
}

/*
RestoreDatabaseStatusParams contains all the parameters to send to the API endpoint
for the restore database status operation typically these are written to a http.Request
*/
type RestoreDatabaseStatusParams struct {

	/*Name*/
	Name string
	/*OperationID*/
	OperationID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the restore database status params
func (o *RestoreDatabaseStatusParams) WithTimeout(timeout time.Duration) *RestoreDatabaseStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the restore database status params
func (o *RestoreDatabaseStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the restore database status params
func (o *RestoreDatabaseStatusParams) WithContext(ctx context.Context) *RestoreDatabaseStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the restore database status params
func (o *RestoreDatabaseStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the restore database status params
func (o *RestoreDatabaseStatusParams) WithHTTPClient(client *http.Client) *RestoreDatabaseStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the restore database status params
func (o *RestoreDatabaseStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the restore database status params
func (o *RestoreDatabaseStatusParams) WithName(name string) *RestoreDatabaseStatusParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the restore database status params
func (o *RestoreDatabaseStatusParams) SetName(name string) {
	o.Name = name
}

// WithOperationID adds the operationID to the restore database status params
func (o *RestoreDatabaseStatusParams) WithOperationID(operationID *string) *RestoreDatabaseStatusParams {
	o.SetOperationID(operationID)
	return o
}

// SetOperationID adds the operationId to the restore database status params
func (o *RestoreDatabaseStatusParams) SetOperationID(operationID *string) {
	o.OperationID = operationID
}

// WriteToRequest writes these params to a swagger request
func (o *RestoreDatabaseStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if o.OperationID != nil {

		// query param operationId
		var qrOperationID string
		if o.OperationID != nil {
			qrOperationID = *o.OperationID
		}
		qOperationID := qrOperationID
		if qOperationID != "" {
			if err := r.SetQueryParam("operationId", qOperationID); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
