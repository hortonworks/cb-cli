// Code generated by go-swagger; DO NOT EDIT.

package v1stacks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewStatusStackParams creates a new StatusStackParams object
// with the default values initialized.
func NewStatusStackParams() *StatusStackParams {
	var ()
	return &StatusStackParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewStatusStackParamsWithTimeout creates a new StatusStackParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewStatusStackParamsWithTimeout(timeout time.Duration) *StatusStackParams {
	var ()
	return &StatusStackParams{

		timeout: timeout,
	}
}

// NewStatusStackParamsWithContext creates a new StatusStackParams object
// with the default values initialized, and the ability to set a context for a request
func NewStatusStackParamsWithContext(ctx context.Context) *StatusStackParams {
	var ()
	return &StatusStackParams{

		Context: ctx,
	}
}

// NewStatusStackParamsWithHTTPClient creates a new StatusStackParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewStatusStackParamsWithHTTPClient(client *http.Client) *StatusStackParams {
	var ()
	return &StatusStackParams{
		HTTPClient: client,
	}
}

/*StatusStackParams contains all the parameters to send to the API endpoint
for the status stack operation typically these are written to a http.Request
*/
type StatusStackParams struct {

	/*ID*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the status stack params
func (o *StatusStackParams) WithTimeout(timeout time.Duration) *StatusStackParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the status stack params
func (o *StatusStackParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the status stack params
func (o *StatusStackParams) WithContext(ctx context.Context) *StatusStackParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the status stack params
func (o *StatusStackParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the status stack params
func (o *StatusStackParams) WithHTTPClient(client *http.Client) *StatusStackParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the status stack params
func (o *StatusStackParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the status stack params
func (o *StatusStackParams) WithID(id int64) *StatusStackParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the status stack params
func (o *StatusStackParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *StatusStackParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
