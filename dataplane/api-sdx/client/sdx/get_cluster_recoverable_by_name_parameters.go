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

// NewGetClusterRecoverableByNameParams creates a new GetClusterRecoverableByNameParams object
// with the default values initialized.
func NewGetClusterRecoverableByNameParams() *GetClusterRecoverableByNameParams {
	var ()
	return &GetClusterRecoverableByNameParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetClusterRecoverableByNameParamsWithTimeout creates a new GetClusterRecoverableByNameParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetClusterRecoverableByNameParamsWithTimeout(timeout time.Duration) *GetClusterRecoverableByNameParams {
	var ()
	return &GetClusterRecoverableByNameParams{

		timeout: timeout,
	}
}

// NewGetClusterRecoverableByNameParamsWithContext creates a new GetClusterRecoverableByNameParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetClusterRecoverableByNameParamsWithContext(ctx context.Context) *GetClusterRecoverableByNameParams {
	var ()
	return &GetClusterRecoverableByNameParams{

		Context: ctx,
	}
}

// NewGetClusterRecoverableByNameParamsWithHTTPClient creates a new GetClusterRecoverableByNameParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetClusterRecoverableByNameParamsWithHTTPClient(client *http.Client) *GetClusterRecoverableByNameParams {
	var ()
	return &GetClusterRecoverableByNameParams{
		HTTPClient: client,
	}
}

/*GetClusterRecoverableByNameParams contains all the parameters to send to the API endpoint
for the get cluster recoverable by name operation typically these are written to a http.Request
*/
type GetClusterRecoverableByNameParams struct {

	/*Name*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get cluster recoverable by name params
func (o *GetClusterRecoverableByNameParams) WithTimeout(timeout time.Duration) *GetClusterRecoverableByNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get cluster recoverable by name params
func (o *GetClusterRecoverableByNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get cluster recoverable by name params
func (o *GetClusterRecoverableByNameParams) WithContext(ctx context.Context) *GetClusterRecoverableByNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get cluster recoverable by name params
func (o *GetClusterRecoverableByNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get cluster recoverable by name params
func (o *GetClusterRecoverableByNameParams) WithHTTPClient(client *http.Client) *GetClusterRecoverableByNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get cluster recoverable by name params
func (o *GetClusterRecoverableByNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the get cluster recoverable by name params
func (o *GetClusterRecoverableByNameParams) WithName(name string) *GetClusterRecoverableByNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get cluster recoverable by name params
func (o *GetClusterRecoverableByNameParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *GetClusterRecoverableByNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
