// Code generated by go-swagger; DO NOT EDIT.

package v1credentials

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

// NewCodeGrantFlowOnExistingCredentialV1Params creates a new CodeGrantFlowOnExistingCredentialV1Params object
// with the default values initialized.
func NewCodeGrantFlowOnExistingCredentialV1Params() *CodeGrantFlowOnExistingCredentialV1Params {
	var ()
	return &CodeGrantFlowOnExistingCredentialV1Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewCodeGrantFlowOnExistingCredentialV1ParamsWithTimeout creates a new CodeGrantFlowOnExistingCredentialV1Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewCodeGrantFlowOnExistingCredentialV1ParamsWithTimeout(timeout time.Duration) *CodeGrantFlowOnExistingCredentialV1Params {
	var ()
	return &CodeGrantFlowOnExistingCredentialV1Params{

		timeout: timeout,
	}
}

// NewCodeGrantFlowOnExistingCredentialV1ParamsWithContext creates a new CodeGrantFlowOnExistingCredentialV1Params object
// with the default values initialized, and the ability to set a context for a request
func NewCodeGrantFlowOnExistingCredentialV1ParamsWithContext(ctx context.Context) *CodeGrantFlowOnExistingCredentialV1Params {
	var ()
	return &CodeGrantFlowOnExistingCredentialV1Params{

		Context: ctx,
	}
}

// NewCodeGrantFlowOnExistingCredentialV1ParamsWithHTTPClient creates a new CodeGrantFlowOnExistingCredentialV1Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCodeGrantFlowOnExistingCredentialV1ParamsWithHTTPClient(client *http.Client) *CodeGrantFlowOnExistingCredentialV1Params {
	var ()
	return &CodeGrantFlowOnExistingCredentialV1Params{
		HTTPClient: client,
	}
}

/*
CodeGrantFlowOnExistingCredentialV1Params contains all the parameters to send to the API endpoint
for the code grant flow on existing credential v1 operation typically these are written to a http.Request
*/
type CodeGrantFlowOnExistingCredentialV1Params struct {

	/*Name*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the code grant flow on existing credential v1 params
func (o *CodeGrantFlowOnExistingCredentialV1Params) WithTimeout(timeout time.Duration) *CodeGrantFlowOnExistingCredentialV1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the code grant flow on existing credential v1 params
func (o *CodeGrantFlowOnExistingCredentialV1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the code grant flow on existing credential v1 params
func (o *CodeGrantFlowOnExistingCredentialV1Params) WithContext(ctx context.Context) *CodeGrantFlowOnExistingCredentialV1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the code grant flow on existing credential v1 params
func (o *CodeGrantFlowOnExistingCredentialV1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the code grant flow on existing credential v1 params
func (o *CodeGrantFlowOnExistingCredentialV1Params) WithHTTPClient(client *http.Client) *CodeGrantFlowOnExistingCredentialV1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the code grant flow on existing credential v1 params
func (o *CodeGrantFlowOnExistingCredentialV1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the code grant flow on existing credential v1 params
func (o *CodeGrantFlowOnExistingCredentialV1Params) WithName(name string) *CodeGrantFlowOnExistingCredentialV1Params {
	o.SetName(name)
	return o
}

// SetName adds the name to the code grant flow on existing credential v1 params
func (o *CodeGrantFlowOnExistingCredentialV1Params) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *CodeGrantFlowOnExistingCredentialV1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
