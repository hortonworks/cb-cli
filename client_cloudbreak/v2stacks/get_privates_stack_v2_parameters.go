// Code generated by go-swagger; DO NOT EDIT.

package v2stacks

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

// NewGetPrivatesStackV2Params creates a new GetPrivatesStackV2Params object
// with the default values initialized.
func NewGetPrivatesStackV2Params() *GetPrivatesStackV2Params {

	return &GetPrivatesStackV2Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPrivatesStackV2ParamsWithTimeout creates a new GetPrivatesStackV2Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPrivatesStackV2ParamsWithTimeout(timeout time.Duration) *GetPrivatesStackV2Params {

	return &GetPrivatesStackV2Params{

		timeout: timeout,
	}
}

// NewGetPrivatesStackV2ParamsWithContext creates a new GetPrivatesStackV2Params object
// with the default values initialized, and the ability to set a context for a request
func NewGetPrivatesStackV2ParamsWithContext(ctx context.Context) *GetPrivatesStackV2Params {

	return &GetPrivatesStackV2Params{

		Context: ctx,
	}
}

// NewGetPrivatesStackV2ParamsWithHTTPClient creates a new GetPrivatesStackV2Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPrivatesStackV2ParamsWithHTTPClient(client *http.Client) *GetPrivatesStackV2Params {

	return &GetPrivatesStackV2Params{
		HTTPClient: client,
	}
}

/*GetPrivatesStackV2Params contains all the parameters to send to the API endpoint
for the get privates stack v2 operation typically these are written to a http.Request
*/
type GetPrivatesStackV2Params struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get privates stack v2 params
func (o *GetPrivatesStackV2Params) WithTimeout(timeout time.Duration) *GetPrivatesStackV2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get privates stack v2 params
func (o *GetPrivatesStackV2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get privates stack v2 params
func (o *GetPrivatesStackV2Params) WithContext(ctx context.Context) *GetPrivatesStackV2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get privates stack v2 params
func (o *GetPrivatesStackV2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get privates stack v2 params
func (o *GetPrivatesStackV2Params) WithHTTPClient(client *http.Client) *GetPrivatesStackV2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get privates stack v2 params
func (o *GetPrivatesStackV2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetPrivatesStackV2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
