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

// NewGetCredentialRequestFromNameParams creates a new GetCredentialRequestFromNameParams object
// with the default values initialized.
func NewGetCredentialRequestFromNameParams() *GetCredentialRequestFromNameParams {
	var ()
	return &GetCredentialRequestFromNameParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCredentialRequestFromNameParamsWithTimeout creates a new GetCredentialRequestFromNameParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCredentialRequestFromNameParamsWithTimeout(timeout time.Duration) *GetCredentialRequestFromNameParams {
	var ()
	return &GetCredentialRequestFromNameParams{

		timeout: timeout,
	}
}

// NewGetCredentialRequestFromNameParamsWithContext creates a new GetCredentialRequestFromNameParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCredentialRequestFromNameParamsWithContext(ctx context.Context) *GetCredentialRequestFromNameParams {
	var ()
	return &GetCredentialRequestFromNameParams{

		Context: ctx,
	}
}

// NewGetCredentialRequestFromNameParamsWithHTTPClient creates a new GetCredentialRequestFromNameParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCredentialRequestFromNameParamsWithHTTPClient(client *http.Client) *GetCredentialRequestFromNameParams {
	var ()
	return &GetCredentialRequestFromNameParams{
		HTTPClient: client,
	}
}

/*GetCredentialRequestFromNameParams contains all the parameters to send to the API endpoint
for the get credential request from name operation typically these are written to a http.Request
*/
type GetCredentialRequestFromNameParams struct {

	/*Name*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get credential request from name params
func (o *GetCredentialRequestFromNameParams) WithTimeout(timeout time.Duration) *GetCredentialRequestFromNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get credential request from name params
func (o *GetCredentialRequestFromNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get credential request from name params
func (o *GetCredentialRequestFromNameParams) WithContext(ctx context.Context) *GetCredentialRequestFromNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get credential request from name params
func (o *GetCredentialRequestFromNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get credential request from name params
func (o *GetCredentialRequestFromNameParams) WithHTTPClient(client *http.Client) *GetCredentialRequestFromNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get credential request from name params
func (o *GetCredentialRequestFromNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the get credential request from name params
func (o *GetCredentialRequestFromNameParams) WithName(name string) *GetCredentialRequestFromNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get credential request from name params
func (o *GetCredentialRequestFromNameParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *GetCredentialRequestFromNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
