// Code generated by go-swagger; DO NOT EDIT.

package flow_logs

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

// NewGetFlowLogsByResourceNameParams creates a new GetFlowLogsByResourceNameParams object
// with the default values initialized.
func NewGetFlowLogsByResourceNameParams() *GetFlowLogsByResourceNameParams {
	var ()
	return &GetFlowLogsByResourceNameParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetFlowLogsByResourceNameParamsWithTimeout creates a new GetFlowLogsByResourceNameParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetFlowLogsByResourceNameParamsWithTimeout(timeout time.Duration) *GetFlowLogsByResourceNameParams {
	var ()
	return &GetFlowLogsByResourceNameParams{

		timeout: timeout,
	}
}

// NewGetFlowLogsByResourceNameParamsWithContext creates a new GetFlowLogsByResourceNameParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetFlowLogsByResourceNameParamsWithContext(ctx context.Context) *GetFlowLogsByResourceNameParams {
	var ()
	return &GetFlowLogsByResourceNameParams{

		Context: ctx,
	}
}

// NewGetFlowLogsByResourceNameParamsWithHTTPClient creates a new GetFlowLogsByResourceNameParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetFlowLogsByResourceNameParamsWithHTTPClient(client *http.Client) *GetFlowLogsByResourceNameParams {
	var ()
	return &GetFlowLogsByResourceNameParams{
		HTTPClient: client,
	}
}

/*GetFlowLogsByResourceNameParams contains all the parameters to send to the API endpoint
for the get flow logs by resource name operation typically these are written to a http.Request
*/
type GetFlowLogsByResourceNameParams struct {

	/*ResourceName*/
	ResourceName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get flow logs by resource name params
func (o *GetFlowLogsByResourceNameParams) WithTimeout(timeout time.Duration) *GetFlowLogsByResourceNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get flow logs by resource name params
func (o *GetFlowLogsByResourceNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get flow logs by resource name params
func (o *GetFlowLogsByResourceNameParams) WithContext(ctx context.Context) *GetFlowLogsByResourceNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get flow logs by resource name params
func (o *GetFlowLogsByResourceNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get flow logs by resource name params
func (o *GetFlowLogsByResourceNameParams) WithHTTPClient(client *http.Client) *GetFlowLogsByResourceNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get flow logs by resource name params
func (o *GetFlowLogsByResourceNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithResourceName adds the resourceName to the get flow logs by resource name params
func (o *GetFlowLogsByResourceNameParams) WithResourceName(resourceName string) *GetFlowLogsByResourceNameParams {
	o.SetResourceName(resourceName)
	return o
}

// SetResourceName adds the resourceName to the get flow logs by resource name params
func (o *GetFlowLogsByResourceNameParams) SetResourceName(resourceName string) {
	o.ResourceName = resourceName
}

// WriteToRequest writes these params to a swagger request
func (o *GetFlowLogsByResourceNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param resourceName
	if err := r.SetPathParam("resourceName", o.ResourceName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
