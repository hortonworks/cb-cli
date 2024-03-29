// Code generated by go-swagger; DO NOT EDIT.

package v1events

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

// NewGetStructuredEventsInAccountParams creates a new GetStructuredEventsInAccountParams object
// with the default values initialized.
func NewGetStructuredEventsInAccountParams() *GetStructuredEventsInAccountParams {
	var ()
	return &GetStructuredEventsInAccountParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetStructuredEventsInAccountParamsWithTimeout creates a new GetStructuredEventsInAccountParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetStructuredEventsInAccountParamsWithTimeout(timeout time.Duration) *GetStructuredEventsInAccountParams {
	var ()
	return &GetStructuredEventsInAccountParams{

		timeout: timeout,
	}
}

// NewGetStructuredEventsInAccountParamsWithContext creates a new GetStructuredEventsInAccountParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetStructuredEventsInAccountParamsWithContext(ctx context.Context) *GetStructuredEventsInAccountParams {
	var ()
	return &GetStructuredEventsInAccountParams{

		Context: ctx,
	}
}

// NewGetStructuredEventsInAccountParamsWithHTTPClient creates a new GetStructuredEventsInAccountParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetStructuredEventsInAccountParamsWithHTTPClient(client *http.Client) *GetStructuredEventsInAccountParams {
	var ()
	return &GetStructuredEventsInAccountParams{
		HTTPClient: client,
	}
}

/*
GetStructuredEventsInAccountParams contains all the parameters to send to the API endpoint
for the get structured events in account operation typically these are written to a http.Request
*/
type GetStructuredEventsInAccountParams struct {

	/*Name*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get structured events in account params
func (o *GetStructuredEventsInAccountParams) WithTimeout(timeout time.Duration) *GetStructuredEventsInAccountParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get structured events in account params
func (o *GetStructuredEventsInAccountParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get structured events in account params
func (o *GetStructuredEventsInAccountParams) WithContext(ctx context.Context) *GetStructuredEventsInAccountParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get structured events in account params
func (o *GetStructuredEventsInAccountParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get structured events in account params
func (o *GetStructuredEventsInAccountParams) WithHTTPClient(client *http.Client) *GetStructuredEventsInAccountParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get structured events in account params
func (o *GetStructuredEventsInAccountParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the get structured events in account params
func (o *GetStructuredEventsInAccountParams) WithName(name string) *GetStructuredEventsInAccountParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get structured events in account params
func (o *GetStructuredEventsInAccountParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *GetStructuredEventsInAccountParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
