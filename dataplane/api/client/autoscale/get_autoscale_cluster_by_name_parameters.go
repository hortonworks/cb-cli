// Code generated by go-swagger; DO NOT EDIT.

package autoscale

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

// NewGetAutoscaleClusterByNameParams creates a new GetAutoscaleClusterByNameParams object
// with the default values initialized.
func NewGetAutoscaleClusterByNameParams() *GetAutoscaleClusterByNameParams {
	var ()
	return &GetAutoscaleClusterByNameParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAutoscaleClusterByNameParamsWithTimeout creates a new GetAutoscaleClusterByNameParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAutoscaleClusterByNameParamsWithTimeout(timeout time.Duration) *GetAutoscaleClusterByNameParams {
	var ()
	return &GetAutoscaleClusterByNameParams{

		timeout: timeout,
	}
}

// NewGetAutoscaleClusterByNameParamsWithContext creates a new GetAutoscaleClusterByNameParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAutoscaleClusterByNameParamsWithContext(ctx context.Context) *GetAutoscaleClusterByNameParams {
	var ()
	return &GetAutoscaleClusterByNameParams{

		Context: ctx,
	}
}

// NewGetAutoscaleClusterByNameParamsWithHTTPClient creates a new GetAutoscaleClusterByNameParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAutoscaleClusterByNameParamsWithHTTPClient(client *http.Client) *GetAutoscaleClusterByNameParams {
	var ()
	return &GetAutoscaleClusterByNameParams{
		HTTPClient: client,
	}
}

/*
GetAutoscaleClusterByNameParams contains all the parameters to send to the API endpoint
for the get autoscale cluster by name operation typically these are written to a http.Request
*/
type GetAutoscaleClusterByNameParams struct {

	/*Name*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get autoscale cluster by name params
func (o *GetAutoscaleClusterByNameParams) WithTimeout(timeout time.Duration) *GetAutoscaleClusterByNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get autoscale cluster by name params
func (o *GetAutoscaleClusterByNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get autoscale cluster by name params
func (o *GetAutoscaleClusterByNameParams) WithContext(ctx context.Context) *GetAutoscaleClusterByNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get autoscale cluster by name params
func (o *GetAutoscaleClusterByNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get autoscale cluster by name params
func (o *GetAutoscaleClusterByNameParams) WithHTTPClient(client *http.Client) *GetAutoscaleClusterByNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get autoscale cluster by name params
func (o *GetAutoscaleClusterByNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the get autoscale cluster by name params
func (o *GetAutoscaleClusterByNameParams) WithName(name string) *GetAutoscaleClusterByNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get autoscale cluster by name params
func (o *GetAutoscaleClusterByNameParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *GetAutoscaleClusterByNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
