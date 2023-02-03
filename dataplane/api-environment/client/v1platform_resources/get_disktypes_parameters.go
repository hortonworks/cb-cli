// Code generated by go-swagger; DO NOT EDIT.

package v1platform_resources

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

// NewGetDisktypesParams creates a new GetDisktypesParams object
// with the default values initialized.
func NewGetDisktypesParams() *GetDisktypesParams {

	return &GetDisktypesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDisktypesParamsWithTimeout creates a new GetDisktypesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDisktypesParamsWithTimeout(timeout time.Duration) *GetDisktypesParams {

	return &GetDisktypesParams{

		timeout: timeout,
	}
}

// NewGetDisktypesParamsWithContext creates a new GetDisktypesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDisktypesParamsWithContext(ctx context.Context) *GetDisktypesParams {

	return &GetDisktypesParams{

		Context: ctx,
	}
}

// NewGetDisktypesParamsWithHTTPClient creates a new GetDisktypesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDisktypesParamsWithHTTPClient(client *http.Client) *GetDisktypesParams {

	return &GetDisktypesParams{
		HTTPClient: client,
	}
}

/*
GetDisktypesParams contains all the parameters to send to the API endpoint
for the get disktypes operation typically these are written to a http.Request
*/
type GetDisktypesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get disktypes params
func (o *GetDisktypesParams) WithTimeout(timeout time.Duration) *GetDisktypesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get disktypes params
func (o *GetDisktypesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get disktypes params
func (o *GetDisktypesParams) WithContext(ctx context.Context) *GetDisktypesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get disktypes params
func (o *GetDisktypesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get disktypes params
func (o *GetDisktypesParams) WithHTTPClient(client *http.Client) *GetDisktypesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get disktypes params
func (o *GetDisktypesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetDisktypesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
