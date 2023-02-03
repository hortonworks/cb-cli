// Code generated by go-swagger; DO NOT EDIT.

package v1operation

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetOperationProgressByEnvironmentCrnParams creates a new GetOperationProgressByEnvironmentCrnParams object
// with the default values initialized.
func NewGetOperationProgressByEnvironmentCrnParams() *GetOperationProgressByEnvironmentCrnParams {
	var (
		detailedDefault = bool(false)
	)
	return &GetOperationProgressByEnvironmentCrnParams{
		Detailed: &detailedDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetOperationProgressByEnvironmentCrnParamsWithTimeout creates a new GetOperationProgressByEnvironmentCrnParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetOperationProgressByEnvironmentCrnParamsWithTimeout(timeout time.Duration) *GetOperationProgressByEnvironmentCrnParams {
	var (
		detailedDefault = bool(false)
	)
	return &GetOperationProgressByEnvironmentCrnParams{
		Detailed: &detailedDefault,

		timeout: timeout,
	}
}

// NewGetOperationProgressByEnvironmentCrnParamsWithContext creates a new GetOperationProgressByEnvironmentCrnParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetOperationProgressByEnvironmentCrnParamsWithContext(ctx context.Context) *GetOperationProgressByEnvironmentCrnParams {
	var (
		detailedDefault = bool(false)
	)
	return &GetOperationProgressByEnvironmentCrnParams{
		Detailed: &detailedDefault,

		Context: ctx,
	}
}

// NewGetOperationProgressByEnvironmentCrnParamsWithHTTPClient creates a new GetOperationProgressByEnvironmentCrnParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetOperationProgressByEnvironmentCrnParamsWithHTTPClient(client *http.Client) *GetOperationProgressByEnvironmentCrnParams {
	var (
		detailedDefault = bool(false)
	)
	return &GetOperationProgressByEnvironmentCrnParams{
		Detailed:   &detailedDefault,
		HTTPClient: client,
	}
}

/*
GetOperationProgressByEnvironmentCrnParams contains all the parameters to send to the API endpoint
for the get operation progress by environment crn operation typically these are written to a http.Request
*/
type GetOperationProgressByEnvironmentCrnParams struct {

	/*Detailed*/
	Detailed *bool
	/*EnvironmentCrn*/
	EnvironmentCrn string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get operation progress by environment crn params
func (o *GetOperationProgressByEnvironmentCrnParams) WithTimeout(timeout time.Duration) *GetOperationProgressByEnvironmentCrnParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get operation progress by environment crn params
func (o *GetOperationProgressByEnvironmentCrnParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get operation progress by environment crn params
func (o *GetOperationProgressByEnvironmentCrnParams) WithContext(ctx context.Context) *GetOperationProgressByEnvironmentCrnParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get operation progress by environment crn params
func (o *GetOperationProgressByEnvironmentCrnParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get operation progress by environment crn params
func (o *GetOperationProgressByEnvironmentCrnParams) WithHTTPClient(client *http.Client) *GetOperationProgressByEnvironmentCrnParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get operation progress by environment crn params
func (o *GetOperationProgressByEnvironmentCrnParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDetailed adds the detailed to the get operation progress by environment crn params
func (o *GetOperationProgressByEnvironmentCrnParams) WithDetailed(detailed *bool) *GetOperationProgressByEnvironmentCrnParams {
	o.SetDetailed(detailed)
	return o
}

// SetDetailed adds the detailed to the get operation progress by environment crn params
func (o *GetOperationProgressByEnvironmentCrnParams) SetDetailed(detailed *bool) {
	o.Detailed = detailed
}

// WithEnvironmentCrn adds the environmentCrn to the get operation progress by environment crn params
func (o *GetOperationProgressByEnvironmentCrnParams) WithEnvironmentCrn(environmentCrn string) *GetOperationProgressByEnvironmentCrnParams {
	o.SetEnvironmentCrn(environmentCrn)
	return o
}

// SetEnvironmentCrn adds the environmentCrn to the get operation progress by environment crn params
func (o *GetOperationProgressByEnvironmentCrnParams) SetEnvironmentCrn(environmentCrn string) {
	o.EnvironmentCrn = environmentCrn
}

// WriteToRequest writes these params to a swagger request
func (o *GetOperationProgressByEnvironmentCrnParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Detailed != nil {

		// query param detailed
		var qrDetailed bool
		if o.Detailed != nil {
			qrDetailed = *o.Detailed
		}
		qDetailed := swag.FormatBool(qrDetailed)
		if qDetailed != "" {
			if err := r.SetQueryParam("detailed", qDetailed); err != nil {
				return err
			}
		}

	}

	// path param environmentCrn
	if err := r.SetPathParam("environmentCrn", o.EnvironmentCrn); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
