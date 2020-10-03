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

// NewGetStackCmVMLogsParams creates a new GetStackCmVMLogsParams object
// with the default values initialized.
func NewGetStackCmVMLogsParams() *GetStackCmVMLogsParams {

	return &GetStackCmVMLogsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetStackCmVMLogsParamsWithTimeout creates a new GetStackCmVMLogsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetStackCmVMLogsParamsWithTimeout(timeout time.Duration) *GetStackCmVMLogsParams {

	return &GetStackCmVMLogsParams{

		timeout: timeout,
	}
}

// NewGetStackCmVMLogsParamsWithContext creates a new GetStackCmVMLogsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetStackCmVMLogsParamsWithContext(ctx context.Context) *GetStackCmVMLogsParams {

	return &GetStackCmVMLogsParams{

		Context: ctx,
	}
}

// NewGetStackCmVMLogsParamsWithHTTPClient creates a new GetStackCmVMLogsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetStackCmVMLogsParamsWithHTTPClient(client *http.Client) *GetStackCmVMLogsParams {

	return &GetStackCmVMLogsParams{
		HTTPClient: client,
	}
}

/*GetStackCmVMLogsParams contains all the parameters to send to the API endpoint
for the get stack cm Vm logs operation typically these are written to a http.Request
*/
type GetStackCmVMLogsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get stack cm Vm logs params
func (o *GetStackCmVMLogsParams) WithTimeout(timeout time.Duration) *GetStackCmVMLogsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get stack cm Vm logs params
func (o *GetStackCmVMLogsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get stack cm Vm logs params
func (o *GetStackCmVMLogsParams) WithContext(ctx context.Context) *GetStackCmVMLogsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get stack cm Vm logs params
func (o *GetStackCmVMLogsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get stack cm Vm logs params
func (o *GetStackCmVMLogsParams) WithHTTPClient(client *http.Client) *GetStackCmVMLogsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get stack cm Vm logs params
func (o *GetStackCmVMLogsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetStackCmVMLogsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
