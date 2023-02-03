// Code generated by go-swagger; DO NOT EDIT.

package database_servers

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

// NewDeleteMultipleDatabaseServersByCrnParams creates a new DeleteMultipleDatabaseServersByCrnParams object
// with the default values initialized.
func NewDeleteMultipleDatabaseServersByCrnParams() *DeleteMultipleDatabaseServersByCrnParams {
	var (
		forceDefault = bool(false)
	)
	return &DeleteMultipleDatabaseServersByCrnParams{
		Force: &forceDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteMultipleDatabaseServersByCrnParamsWithTimeout creates a new DeleteMultipleDatabaseServersByCrnParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteMultipleDatabaseServersByCrnParamsWithTimeout(timeout time.Duration) *DeleteMultipleDatabaseServersByCrnParams {
	var (
		forceDefault = bool(false)
	)
	return &DeleteMultipleDatabaseServersByCrnParams{
		Force: &forceDefault,

		timeout: timeout,
	}
}

// NewDeleteMultipleDatabaseServersByCrnParamsWithContext creates a new DeleteMultipleDatabaseServersByCrnParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteMultipleDatabaseServersByCrnParamsWithContext(ctx context.Context) *DeleteMultipleDatabaseServersByCrnParams {
	var (
		forceDefault = bool(false)
	)
	return &DeleteMultipleDatabaseServersByCrnParams{
		Force: &forceDefault,

		Context: ctx,
	}
}

// NewDeleteMultipleDatabaseServersByCrnParamsWithHTTPClient creates a new DeleteMultipleDatabaseServersByCrnParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteMultipleDatabaseServersByCrnParamsWithHTTPClient(client *http.Client) *DeleteMultipleDatabaseServersByCrnParams {
	var (
		forceDefault = bool(false)
	)
	return &DeleteMultipleDatabaseServersByCrnParams{
		Force:      &forceDefault,
		HTTPClient: client,
	}
}

/*
DeleteMultipleDatabaseServersByCrnParams contains all the parameters to send to the API endpoint
for the delete multiple database servers by crn operation typically these are written to a http.Request
*/
type DeleteMultipleDatabaseServersByCrnParams struct {

	/*Body
	  CRNs of the database servers

	*/
	Body []string
	/*Force*/
	Force *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete multiple database servers by crn params
func (o *DeleteMultipleDatabaseServersByCrnParams) WithTimeout(timeout time.Duration) *DeleteMultipleDatabaseServersByCrnParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete multiple database servers by crn params
func (o *DeleteMultipleDatabaseServersByCrnParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete multiple database servers by crn params
func (o *DeleteMultipleDatabaseServersByCrnParams) WithContext(ctx context.Context) *DeleteMultipleDatabaseServersByCrnParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete multiple database servers by crn params
func (o *DeleteMultipleDatabaseServersByCrnParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete multiple database servers by crn params
func (o *DeleteMultipleDatabaseServersByCrnParams) WithHTTPClient(client *http.Client) *DeleteMultipleDatabaseServersByCrnParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete multiple database servers by crn params
func (o *DeleteMultipleDatabaseServersByCrnParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the delete multiple database servers by crn params
func (o *DeleteMultipleDatabaseServersByCrnParams) WithBody(body []string) *DeleteMultipleDatabaseServersByCrnParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the delete multiple database servers by crn params
func (o *DeleteMultipleDatabaseServersByCrnParams) SetBody(body []string) {
	o.Body = body
}

// WithForce adds the force to the delete multiple database servers by crn params
func (o *DeleteMultipleDatabaseServersByCrnParams) WithForce(force *bool) *DeleteMultipleDatabaseServersByCrnParams {
	o.SetForce(force)
	return o
}

// SetForce adds the force to the delete multiple database servers by crn params
func (o *DeleteMultipleDatabaseServersByCrnParams) SetForce(force *bool) {
	o.Force = force
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteMultipleDatabaseServersByCrnParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if o.Force != nil {

		// query param force
		var qrForce bool
		if o.Force != nil {
			qrForce = *o.Force
		}
		qForce := swag.FormatBool(qrForce)
		if qForce != "" {
			if err := r.SetQueryParam("force", qForce); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
