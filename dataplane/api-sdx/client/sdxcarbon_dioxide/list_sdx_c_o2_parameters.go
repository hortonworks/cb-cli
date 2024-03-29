// Code generated by go-swagger; DO NOT EDIT.

package sdxcarbon_dioxide

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

// NewListSdxCO2Params creates a new ListSdxCO2Params object
// with the default values initialized.
func NewListSdxCO2Params() *ListSdxCO2Params {
	var ()
	return &ListSdxCO2Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewListSdxCO2ParamsWithTimeout creates a new ListSdxCO2Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewListSdxCO2ParamsWithTimeout(timeout time.Duration) *ListSdxCO2Params {
	var ()
	return &ListSdxCO2Params{

		timeout: timeout,
	}
}

// NewListSdxCO2ParamsWithContext creates a new ListSdxCO2Params object
// with the default values initialized, and the ability to set a context for a request
func NewListSdxCO2ParamsWithContext(ctx context.Context) *ListSdxCO2Params {
	var ()
	return &ListSdxCO2Params{

		Context: ctx,
	}
}

// NewListSdxCO2ParamsWithHTTPClient creates a new ListSdxCO2Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListSdxCO2ParamsWithHTTPClient(client *http.Client) *ListSdxCO2Params {
	var ()
	return &ListSdxCO2Params{
		HTTPClient: client,
	}
}

/*
ListSdxCO2Params contains all the parameters to send to the API endpoint
for the list sdx c o2 operation typically these are written to a http.Request
*/
type ListSdxCO2Params struct {

	/*Body*/
	Body []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list sdx c o2 params
func (o *ListSdxCO2Params) WithTimeout(timeout time.Duration) *ListSdxCO2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list sdx c o2 params
func (o *ListSdxCO2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list sdx c o2 params
func (o *ListSdxCO2Params) WithContext(ctx context.Context) *ListSdxCO2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list sdx c o2 params
func (o *ListSdxCO2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list sdx c o2 params
func (o *ListSdxCO2Params) WithHTTPClient(client *http.Client) *ListSdxCO2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list sdx c o2 params
func (o *ListSdxCO2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the list sdx c o2 params
func (o *ListSdxCO2Params) WithBody(body []string) *ListSdxCO2Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the list sdx c o2 params
func (o *ListSdxCO2Params) SetBody(body []string) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ListSdxCO2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
