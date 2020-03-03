// Code generated by go-swagger; DO NOT EDIT.

package v1env

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

// NewStartEnvironmentByNameV1Params creates a new StartEnvironmentByNameV1Params object
// with the default values initialized.
func NewStartEnvironmentByNameV1Params() *StartEnvironmentByNameV1Params {
	var ()
	return &StartEnvironmentByNameV1Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewStartEnvironmentByNameV1ParamsWithTimeout creates a new StartEnvironmentByNameV1Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewStartEnvironmentByNameV1ParamsWithTimeout(timeout time.Duration) *StartEnvironmentByNameV1Params {
	var ()
	return &StartEnvironmentByNameV1Params{

		timeout: timeout,
	}
}

// NewStartEnvironmentByNameV1ParamsWithContext creates a new StartEnvironmentByNameV1Params object
// with the default values initialized, and the ability to set a context for a request
func NewStartEnvironmentByNameV1ParamsWithContext(ctx context.Context) *StartEnvironmentByNameV1Params {
	var ()
	return &StartEnvironmentByNameV1Params{

		Context: ctx,
	}
}

// NewStartEnvironmentByNameV1ParamsWithHTTPClient creates a new StartEnvironmentByNameV1Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewStartEnvironmentByNameV1ParamsWithHTTPClient(client *http.Client) *StartEnvironmentByNameV1Params {
	var ()
	return &StartEnvironmentByNameV1Params{
		HTTPClient: client,
	}
}

/*StartEnvironmentByNameV1Params contains all the parameters to send to the API endpoint
for the start environment by name v1 operation typically these are written to a http.Request
*/
type StartEnvironmentByNameV1Params struct {

	/*Name*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the start environment by name v1 params
func (o *StartEnvironmentByNameV1Params) WithTimeout(timeout time.Duration) *StartEnvironmentByNameV1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the start environment by name v1 params
func (o *StartEnvironmentByNameV1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the start environment by name v1 params
func (o *StartEnvironmentByNameV1Params) WithContext(ctx context.Context) *StartEnvironmentByNameV1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the start environment by name v1 params
func (o *StartEnvironmentByNameV1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the start environment by name v1 params
func (o *StartEnvironmentByNameV1Params) WithHTTPClient(client *http.Client) *StartEnvironmentByNameV1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the start environment by name v1 params
func (o *StartEnvironmentByNameV1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the start environment by name v1 params
func (o *StartEnvironmentByNameV1Params) WithName(name string) *StartEnvironmentByNameV1Params {
	o.SetName(name)
	return o
}

// SetName adds the name to the start environment by name v1 params
func (o *StartEnvironmentByNameV1Params) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *StartEnvironmentByNameV1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
