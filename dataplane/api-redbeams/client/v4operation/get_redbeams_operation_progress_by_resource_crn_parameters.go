// Code generated by go-swagger; DO NOT EDIT.

package v4operation

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

// NewGetRedbeamsOperationProgressByResourceCrnParams creates a new GetRedbeamsOperationProgressByResourceCrnParams object
// with the default values initialized.
func NewGetRedbeamsOperationProgressByResourceCrnParams() *GetRedbeamsOperationProgressByResourceCrnParams {
	var (
		detailedDefault = bool(false)
	)
	return &GetRedbeamsOperationProgressByResourceCrnParams{
		Detailed: &detailedDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRedbeamsOperationProgressByResourceCrnParamsWithTimeout creates a new GetRedbeamsOperationProgressByResourceCrnParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRedbeamsOperationProgressByResourceCrnParamsWithTimeout(timeout time.Duration) *GetRedbeamsOperationProgressByResourceCrnParams {
	var (
		detailedDefault = bool(false)
	)
	return &GetRedbeamsOperationProgressByResourceCrnParams{
		Detailed: &detailedDefault,

		timeout: timeout,
	}
}

// NewGetRedbeamsOperationProgressByResourceCrnParamsWithContext creates a new GetRedbeamsOperationProgressByResourceCrnParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRedbeamsOperationProgressByResourceCrnParamsWithContext(ctx context.Context) *GetRedbeamsOperationProgressByResourceCrnParams {
	var (
		detailedDefault = bool(false)
	)
	return &GetRedbeamsOperationProgressByResourceCrnParams{
		Detailed: &detailedDefault,

		Context: ctx,
	}
}

// NewGetRedbeamsOperationProgressByResourceCrnParamsWithHTTPClient creates a new GetRedbeamsOperationProgressByResourceCrnParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRedbeamsOperationProgressByResourceCrnParamsWithHTTPClient(client *http.Client) *GetRedbeamsOperationProgressByResourceCrnParams {
	var (
		detailedDefault = bool(false)
	)
	return &GetRedbeamsOperationProgressByResourceCrnParams{
		Detailed:   &detailedDefault,
		HTTPClient: client,
	}
}

/*
GetRedbeamsOperationProgressByResourceCrnParams contains all the parameters to send to the API endpoint
for the get redbeams operation progress by resource crn operation typically these are written to a http.Request
*/
type GetRedbeamsOperationProgressByResourceCrnParams struct {

	/*Detailed*/
	Detailed *bool
	/*ResourceCrn*/
	ResourceCrn string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get redbeams operation progress by resource crn params
func (o *GetRedbeamsOperationProgressByResourceCrnParams) WithTimeout(timeout time.Duration) *GetRedbeamsOperationProgressByResourceCrnParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get redbeams operation progress by resource crn params
func (o *GetRedbeamsOperationProgressByResourceCrnParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get redbeams operation progress by resource crn params
func (o *GetRedbeamsOperationProgressByResourceCrnParams) WithContext(ctx context.Context) *GetRedbeamsOperationProgressByResourceCrnParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get redbeams operation progress by resource crn params
func (o *GetRedbeamsOperationProgressByResourceCrnParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get redbeams operation progress by resource crn params
func (o *GetRedbeamsOperationProgressByResourceCrnParams) WithHTTPClient(client *http.Client) *GetRedbeamsOperationProgressByResourceCrnParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get redbeams operation progress by resource crn params
func (o *GetRedbeamsOperationProgressByResourceCrnParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDetailed adds the detailed to the get redbeams operation progress by resource crn params
func (o *GetRedbeamsOperationProgressByResourceCrnParams) WithDetailed(detailed *bool) *GetRedbeamsOperationProgressByResourceCrnParams {
	o.SetDetailed(detailed)
	return o
}

// SetDetailed adds the detailed to the get redbeams operation progress by resource crn params
func (o *GetRedbeamsOperationProgressByResourceCrnParams) SetDetailed(detailed *bool) {
	o.Detailed = detailed
}

// WithResourceCrn adds the resourceCrn to the get redbeams operation progress by resource crn params
func (o *GetRedbeamsOperationProgressByResourceCrnParams) WithResourceCrn(resourceCrn string) *GetRedbeamsOperationProgressByResourceCrnParams {
	o.SetResourceCrn(resourceCrn)
	return o
}

// SetResourceCrn adds the resourceCrn to the get redbeams operation progress by resource crn params
func (o *GetRedbeamsOperationProgressByResourceCrnParams) SetResourceCrn(resourceCrn string) {
	o.ResourceCrn = resourceCrn
}

// WriteToRequest writes these params to a swagger request
func (o *GetRedbeamsOperationProgressByResourceCrnParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param resourceCrn
	if err := r.SetPathParam("resourceCrn", o.ResourceCrn); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
