// Code generated by go-swagger; DO NOT EDIT.

package flow

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

// NewGetFlowLogsByFlowIDParams creates a new GetFlowLogsByFlowIDParams object
// with the default values initialized.
func NewGetFlowLogsByFlowIDParams() *GetFlowLogsByFlowIDParams {
	var ()
	return &GetFlowLogsByFlowIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetFlowLogsByFlowIDParamsWithTimeout creates a new GetFlowLogsByFlowIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetFlowLogsByFlowIDParamsWithTimeout(timeout time.Duration) *GetFlowLogsByFlowIDParams {
	var ()
	return &GetFlowLogsByFlowIDParams{

		timeout: timeout,
	}
}

// NewGetFlowLogsByFlowIDParamsWithContext creates a new GetFlowLogsByFlowIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetFlowLogsByFlowIDParamsWithContext(ctx context.Context) *GetFlowLogsByFlowIDParams {
	var ()
	return &GetFlowLogsByFlowIDParams{

		Context: ctx,
	}
}

// NewGetFlowLogsByFlowIDParamsWithHTTPClient creates a new GetFlowLogsByFlowIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetFlowLogsByFlowIDParamsWithHTTPClient(client *http.Client) *GetFlowLogsByFlowIDParams {
	var ()
	return &GetFlowLogsByFlowIDParams{
		HTTPClient: client,
	}
}

/*GetFlowLogsByFlowIDParams contains all the parameters to send to the API endpoint
for the get flow logs by flow Id operation typically these are written to a http.Request
*/
type GetFlowLogsByFlowIDParams struct {

	/*FlowID*/
	FlowID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get flow logs by flow Id params
func (o *GetFlowLogsByFlowIDParams) WithTimeout(timeout time.Duration) *GetFlowLogsByFlowIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get flow logs by flow Id params
func (o *GetFlowLogsByFlowIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get flow logs by flow Id params
func (o *GetFlowLogsByFlowIDParams) WithContext(ctx context.Context) *GetFlowLogsByFlowIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get flow logs by flow Id params
func (o *GetFlowLogsByFlowIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get flow logs by flow Id params
func (o *GetFlowLogsByFlowIDParams) WithHTTPClient(client *http.Client) *GetFlowLogsByFlowIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get flow logs by flow Id params
func (o *GetFlowLogsByFlowIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFlowID adds the flowID to the get flow logs by flow Id params
func (o *GetFlowLogsByFlowIDParams) WithFlowID(flowID string) *GetFlowLogsByFlowIDParams {
	o.SetFlowID(flowID)
	return o
}

// SetFlowID adds the flowId to the get flow logs by flow Id params
func (o *GetFlowLogsByFlowIDParams) SetFlowID(flowID string) {
	o.FlowID = flowID
}

// WriteToRequest writes these params to a swagger request
func (o *GetFlowLogsByFlowIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param flowId
	if err := r.SetPathParam("flowId", o.FlowID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
