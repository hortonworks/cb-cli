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

// NewHasFlowRunningByChainIDParams creates a new HasFlowRunningByChainIDParams object
// with the default values initialized.
func NewHasFlowRunningByChainIDParams() *HasFlowRunningByChainIDParams {
	var ()
	return &HasFlowRunningByChainIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewHasFlowRunningByChainIDParamsWithTimeout creates a new HasFlowRunningByChainIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewHasFlowRunningByChainIDParamsWithTimeout(timeout time.Duration) *HasFlowRunningByChainIDParams {
	var ()
	return &HasFlowRunningByChainIDParams{

		timeout: timeout,
	}
}

// NewHasFlowRunningByChainIDParamsWithContext creates a new HasFlowRunningByChainIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewHasFlowRunningByChainIDParamsWithContext(ctx context.Context) *HasFlowRunningByChainIDParams {
	var ()
	return &HasFlowRunningByChainIDParams{

		Context: ctx,
	}
}

// NewHasFlowRunningByChainIDParamsWithHTTPClient creates a new HasFlowRunningByChainIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewHasFlowRunningByChainIDParamsWithHTTPClient(client *http.Client) *HasFlowRunningByChainIDParams {
	var ()
	return &HasFlowRunningByChainIDParams{
		HTTPClient: client,
	}
}

/*
HasFlowRunningByChainIDParams contains all the parameters to send to the API endpoint
for the has flow running by chain Id operation typically these are written to a http.Request
*/
type HasFlowRunningByChainIDParams struct {

	/*ChainID*/
	ChainID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the has flow running by chain Id params
func (o *HasFlowRunningByChainIDParams) WithTimeout(timeout time.Duration) *HasFlowRunningByChainIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the has flow running by chain Id params
func (o *HasFlowRunningByChainIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the has flow running by chain Id params
func (o *HasFlowRunningByChainIDParams) WithContext(ctx context.Context) *HasFlowRunningByChainIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the has flow running by chain Id params
func (o *HasFlowRunningByChainIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the has flow running by chain Id params
func (o *HasFlowRunningByChainIDParams) WithHTTPClient(client *http.Client) *HasFlowRunningByChainIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the has flow running by chain Id params
func (o *HasFlowRunningByChainIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithChainID adds the chainID to the has flow running by chain Id params
func (o *HasFlowRunningByChainIDParams) WithChainID(chainID string) *HasFlowRunningByChainIDParams {
	o.SetChainID(chainID)
	return o
}

// SetChainID adds the chainId to the has flow running by chain Id params
func (o *HasFlowRunningByChainIDParams) SetChainID(chainID string) {
	o.ChainID = chainID
}

// WriteToRequest writes these params to a swagger request
func (o *HasFlowRunningByChainIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param chainId
	if err := r.SetPathParam("chainId", o.ChainID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
