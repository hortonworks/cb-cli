// Code generated by go-swagger; DO NOT EDIT.

package v1rdsconfigs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeletePublicRdsParams creates a new DeletePublicRdsParams object
// with the default values initialized.
func NewDeletePublicRdsParams() *DeletePublicRdsParams {
	var ()
	return &DeletePublicRdsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeletePublicRdsParamsWithTimeout creates a new DeletePublicRdsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeletePublicRdsParamsWithTimeout(timeout time.Duration) *DeletePublicRdsParams {
	var ()
	return &DeletePublicRdsParams{

		timeout: timeout,
	}
}

// NewDeletePublicRdsParamsWithContext creates a new DeletePublicRdsParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeletePublicRdsParamsWithContext(ctx context.Context) *DeletePublicRdsParams {
	var ()
	return &DeletePublicRdsParams{

		Context: ctx,
	}
}

// NewDeletePublicRdsParamsWithHTTPClient creates a new DeletePublicRdsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeletePublicRdsParamsWithHTTPClient(client *http.Client) *DeletePublicRdsParams {
	var ()
	return &DeletePublicRdsParams{
		HTTPClient: client,
	}
}

/*DeletePublicRdsParams contains all the parameters to send to the API endpoint
for the delete public rds operation typically these are written to a http.Request
*/
type DeletePublicRdsParams struct {

	/*Name*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete public rds params
func (o *DeletePublicRdsParams) WithTimeout(timeout time.Duration) *DeletePublicRdsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete public rds params
func (o *DeletePublicRdsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete public rds params
func (o *DeletePublicRdsParams) WithContext(ctx context.Context) *DeletePublicRdsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete public rds params
func (o *DeletePublicRdsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete public rds params
func (o *DeletePublicRdsParams) WithHTTPClient(client *http.Client) *DeletePublicRdsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete public rds params
func (o *DeletePublicRdsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the delete public rds params
func (o *DeletePublicRdsParams) WithName(name string) *DeletePublicRdsParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the delete public rds params
func (o *DeletePublicRdsParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *DeletePublicRdsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
