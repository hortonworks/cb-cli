// Code generated by go-swagger; DO NOT EDIT.

package v1credentials

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

// NewDeletePublicCredentialParams creates a new DeletePublicCredentialParams object
// with the default values initialized.
func NewDeletePublicCredentialParams() *DeletePublicCredentialParams {
	var ()
	return &DeletePublicCredentialParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeletePublicCredentialParamsWithTimeout creates a new DeletePublicCredentialParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeletePublicCredentialParamsWithTimeout(timeout time.Duration) *DeletePublicCredentialParams {
	var ()
	return &DeletePublicCredentialParams{

		timeout: timeout,
	}
}

// NewDeletePublicCredentialParamsWithContext creates a new DeletePublicCredentialParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeletePublicCredentialParamsWithContext(ctx context.Context) *DeletePublicCredentialParams {
	var ()
	return &DeletePublicCredentialParams{

		Context: ctx,
	}
}

// NewDeletePublicCredentialParamsWithHTTPClient creates a new DeletePublicCredentialParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeletePublicCredentialParamsWithHTTPClient(client *http.Client) *DeletePublicCredentialParams {
	var ()
	return &DeletePublicCredentialParams{
		HTTPClient: client,
	}
}

/*DeletePublicCredentialParams contains all the parameters to send to the API endpoint
for the delete public credential operation typically these are written to a http.Request
*/
type DeletePublicCredentialParams struct {

	/*Name*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete public credential params
func (o *DeletePublicCredentialParams) WithTimeout(timeout time.Duration) *DeletePublicCredentialParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete public credential params
func (o *DeletePublicCredentialParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete public credential params
func (o *DeletePublicCredentialParams) WithContext(ctx context.Context) *DeletePublicCredentialParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete public credential params
func (o *DeletePublicCredentialParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete public credential params
func (o *DeletePublicCredentialParams) WithHTTPClient(client *http.Client) *DeletePublicCredentialParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete public credential params
func (o *DeletePublicCredentialParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the delete public credential params
func (o *DeletePublicCredentialParams) WithName(name string) *DeletePublicCredentialParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the delete public credential params
func (o *DeletePublicCredentialParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *DeletePublicCredentialParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
