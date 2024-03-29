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

// NewDeleteMultipleDatabasesByCrnParams creates a new DeleteMultipleDatabasesByCrnParams object
// with the default values initialized.
func NewDeleteMultipleDatabasesByCrnParams() *DeleteMultipleDatabasesByCrnParams {
	var ()
	return &DeleteMultipleDatabasesByCrnParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteMultipleDatabasesByCrnParamsWithTimeout creates a new DeleteMultipleDatabasesByCrnParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteMultipleDatabasesByCrnParamsWithTimeout(timeout time.Duration) *DeleteMultipleDatabasesByCrnParams {
	var ()
	return &DeleteMultipleDatabasesByCrnParams{

		timeout: timeout,
	}
}

// NewDeleteMultipleDatabasesByCrnParamsWithContext creates a new DeleteMultipleDatabasesByCrnParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteMultipleDatabasesByCrnParamsWithContext(ctx context.Context) *DeleteMultipleDatabasesByCrnParams {
	var ()
	return &DeleteMultipleDatabasesByCrnParams{

		Context: ctx,
	}
}

// NewDeleteMultipleDatabasesByCrnParamsWithHTTPClient creates a new DeleteMultipleDatabasesByCrnParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteMultipleDatabasesByCrnParamsWithHTTPClient(client *http.Client) *DeleteMultipleDatabasesByCrnParams {
	var ()
	return &DeleteMultipleDatabasesByCrnParams{
		HTTPClient: client,
	}
}

/*
DeleteMultipleDatabasesByCrnParams contains all the parameters to send to the API endpoint
for the delete multiple databases by crn operation typically these are written to a http.Request
*/
type DeleteMultipleDatabasesByCrnParams struct {

	/*Body
	  CRNs of the databases

	*/
	Body []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete multiple databases by crn params
func (o *DeleteMultipleDatabasesByCrnParams) WithTimeout(timeout time.Duration) *DeleteMultipleDatabasesByCrnParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete multiple databases by crn params
func (o *DeleteMultipleDatabasesByCrnParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete multiple databases by crn params
func (o *DeleteMultipleDatabasesByCrnParams) WithContext(ctx context.Context) *DeleteMultipleDatabasesByCrnParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete multiple databases by crn params
func (o *DeleteMultipleDatabasesByCrnParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete multiple databases by crn params
func (o *DeleteMultipleDatabasesByCrnParams) WithHTTPClient(client *http.Client) *DeleteMultipleDatabasesByCrnParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete multiple databases by crn params
func (o *DeleteMultipleDatabasesByCrnParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the delete multiple databases by crn params
func (o *DeleteMultipleDatabasesByCrnParams) WithBody(body []string) *DeleteMultipleDatabasesByCrnParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the delete multiple databases by crn params
func (o *DeleteMultipleDatabasesByCrnParams) SetBody(body []string) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteMultipleDatabasesByCrnParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
