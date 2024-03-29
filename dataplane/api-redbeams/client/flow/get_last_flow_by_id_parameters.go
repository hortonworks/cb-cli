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

// NewGetLastFlowByIDParams creates a new GetLastFlowByIDParams object
// with the default values initialized.
func NewGetLastFlowByIDParams() *GetLastFlowByIDParams {
	var ()
	return &GetLastFlowByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLastFlowByIDParamsWithTimeout creates a new GetLastFlowByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLastFlowByIDParamsWithTimeout(timeout time.Duration) *GetLastFlowByIDParams {
	var ()
	return &GetLastFlowByIDParams{

		timeout: timeout,
	}
}

// NewGetLastFlowByIDParamsWithContext creates a new GetLastFlowByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLastFlowByIDParamsWithContext(ctx context.Context) *GetLastFlowByIDParams {
	var ()
	return &GetLastFlowByIDParams{

		Context: ctx,
	}
}

// NewGetLastFlowByIDParamsWithHTTPClient creates a new GetLastFlowByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLastFlowByIDParamsWithHTTPClient(client *http.Client) *GetLastFlowByIDParams {
	var ()
	return &GetLastFlowByIDParams{
		HTTPClient: client,
	}
}

/*
GetLastFlowByIDParams contains all the parameters to send to the API endpoint
for the get last flow by Id operation typically these are written to a http.Request
*/
type GetLastFlowByIDParams struct {

	/*FlowID*/
	FlowID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get last flow by Id params
func (o *GetLastFlowByIDParams) WithTimeout(timeout time.Duration) *GetLastFlowByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get last flow by Id params
func (o *GetLastFlowByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get last flow by Id params
func (o *GetLastFlowByIDParams) WithContext(ctx context.Context) *GetLastFlowByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get last flow by Id params
func (o *GetLastFlowByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get last flow by Id params
func (o *GetLastFlowByIDParams) WithHTTPClient(client *http.Client) *GetLastFlowByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get last flow by Id params
func (o *GetLastFlowByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFlowID adds the flowID to the get last flow by Id params
func (o *GetLastFlowByIDParams) WithFlowID(flowID string) *GetLastFlowByIDParams {
	o.SetFlowID(flowID)
	return o
}

// SetFlowID adds the flowId to the get last flow by Id params
func (o *GetLastFlowByIDParams) SetFlowID(flowID string) {
	o.FlowID = flowID
}

// WriteToRequest writes these params to a swagger request
func (o *GetLastFlowByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
