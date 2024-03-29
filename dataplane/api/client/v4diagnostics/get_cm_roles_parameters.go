// Code generated by go-swagger; DO NOT EDIT.

package v4diagnostics

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

// NewGetCmRolesParams creates a new GetCmRolesParams object
// with the default values initialized.
func NewGetCmRolesParams() *GetCmRolesParams {
	var ()
	return &GetCmRolesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCmRolesParamsWithTimeout creates a new GetCmRolesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCmRolesParamsWithTimeout(timeout time.Duration) *GetCmRolesParams {
	var ()
	return &GetCmRolesParams{

		timeout: timeout,
	}
}

// NewGetCmRolesParamsWithContext creates a new GetCmRolesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCmRolesParamsWithContext(ctx context.Context) *GetCmRolesParams {
	var ()
	return &GetCmRolesParams{

		Context: ctx,
	}
}

// NewGetCmRolesParamsWithHTTPClient creates a new GetCmRolesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCmRolesParamsWithHTTPClient(client *http.Client) *GetCmRolesParams {
	var ()
	return &GetCmRolesParams{
		HTTPClient: client,
	}
}

/*
GetCmRolesParams contains all the parameters to send to the API endpoint
for the get cm roles operation typically these are written to a http.Request
*/
type GetCmRolesParams struct {

	/*StackCrn*/
	StackCrn string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get cm roles params
func (o *GetCmRolesParams) WithTimeout(timeout time.Duration) *GetCmRolesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get cm roles params
func (o *GetCmRolesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get cm roles params
func (o *GetCmRolesParams) WithContext(ctx context.Context) *GetCmRolesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get cm roles params
func (o *GetCmRolesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get cm roles params
func (o *GetCmRolesParams) WithHTTPClient(client *http.Client) *GetCmRolesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get cm roles params
func (o *GetCmRolesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithStackCrn adds the stackCrn to the get cm roles params
func (o *GetCmRolesParams) WithStackCrn(stackCrn string) *GetCmRolesParams {
	o.SetStackCrn(stackCrn)
	return o
}

// SetStackCrn adds the stackCrn to the get cm roles params
func (o *GetCmRolesParams) SetStackCrn(stackCrn string) {
	o.StackCrn = stackCrn
}

// WriteToRequest writes these params to a swagger request
func (o *GetCmRolesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param stackCrn
	if err := r.SetPathParam("stackCrn", o.StackCrn); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
