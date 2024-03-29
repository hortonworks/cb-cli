// Code generated by go-swagger; DO NOT EDIT.

package v4events

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

// NewGetStructuredEventsInWorkspaceByCrnParams creates a new GetStructuredEventsInWorkspaceByCrnParams object
// with the default values initialized.
func NewGetStructuredEventsInWorkspaceByCrnParams() *GetStructuredEventsInWorkspaceByCrnParams {
	var (
		onlyAliveDefault = bool(true)
	)
	return &GetStructuredEventsInWorkspaceByCrnParams{
		OnlyAlive: &onlyAliveDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetStructuredEventsInWorkspaceByCrnParamsWithTimeout creates a new GetStructuredEventsInWorkspaceByCrnParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetStructuredEventsInWorkspaceByCrnParamsWithTimeout(timeout time.Duration) *GetStructuredEventsInWorkspaceByCrnParams {
	var (
		onlyAliveDefault = bool(true)
	)
	return &GetStructuredEventsInWorkspaceByCrnParams{
		OnlyAlive: &onlyAliveDefault,

		timeout: timeout,
	}
}

// NewGetStructuredEventsInWorkspaceByCrnParamsWithContext creates a new GetStructuredEventsInWorkspaceByCrnParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetStructuredEventsInWorkspaceByCrnParamsWithContext(ctx context.Context) *GetStructuredEventsInWorkspaceByCrnParams {
	var (
		onlyAliveDefault = bool(true)
	)
	return &GetStructuredEventsInWorkspaceByCrnParams{
		OnlyAlive: &onlyAliveDefault,

		Context: ctx,
	}
}

// NewGetStructuredEventsInWorkspaceByCrnParamsWithHTTPClient creates a new GetStructuredEventsInWorkspaceByCrnParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetStructuredEventsInWorkspaceByCrnParamsWithHTTPClient(client *http.Client) *GetStructuredEventsInWorkspaceByCrnParams {
	var (
		onlyAliveDefault = bool(true)
	)
	return &GetStructuredEventsInWorkspaceByCrnParams{
		OnlyAlive:  &onlyAliveDefault,
		HTTPClient: client,
	}
}

/*
GetStructuredEventsInWorkspaceByCrnParams contains all the parameters to send to the API endpoint
for the get structured events in workspace by crn operation typically these are written to a http.Request
*/
type GetStructuredEventsInWorkspaceByCrnParams struct {

	/*Crn*/
	Crn string
	/*OnlyAlive*/
	OnlyAlive *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get structured events in workspace by crn params
func (o *GetStructuredEventsInWorkspaceByCrnParams) WithTimeout(timeout time.Duration) *GetStructuredEventsInWorkspaceByCrnParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get structured events in workspace by crn params
func (o *GetStructuredEventsInWorkspaceByCrnParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get structured events in workspace by crn params
func (o *GetStructuredEventsInWorkspaceByCrnParams) WithContext(ctx context.Context) *GetStructuredEventsInWorkspaceByCrnParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get structured events in workspace by crn params
func (o *GetStructuredEventsInWorkspaceByCrnParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get structured events in workspace by crn params
func (o *GetStructuredEventsInWorkspaceByCrnParams) WithHTTPClient(client *http.Client) *GetStructuredEventsInWorkspaceByCrnParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get structured events in workspace by crn params
func (o *GetStructuredEventsInWorkspaceByCrnParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCrn adds the crn to the get structured events in workspace by crn params
func (o *GetStructuredEventsInWorkspaceByCrnParams) WithCrn(crn string) *GetStructuredEventsInWorkspaceByCrnParams {
	o.SetCrn(crn)
	return o
}

// SetCrn adds the crn to the get structured events in workspace by crn params
func (o *GetStructuredEventsInWorkspaceByCrnParams) SetCrn(crn string) {
	o.Crn = crn
}

// WithOnlyAlive adds the onlyAlive to the get structured events in workspace by crn params
func (o *GetStructuredEventsInWorkspaceByCrnParams) WithOnlyAlive(onlyAlive *bool) *GetStructuredEventsInWorkspaceByCrnParams {
	o.SetOnlyAlive(onlyAlive)
	return o
}

// SetOnlyAlive adds the onlyAlive to the get structured events in workspace by crn params
func (o *GetStructuredEventsInWorkspaceByCrnParams) SetOnlyAlive(onlyAlive *bool) {
	o.OnlyAlive = onlyAlive
}

// WriteToRequest writes these params to a swagger request
func (o *GetStructuredEventsInWorkspaceByCrnParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param crn
	if err := r.SetPathParam("crn", o.Crn); err != nil {
		return err
	}

	if o.OnlyAlive != nil {

		// query param onlyAlive
		var qrOnlyAlive bool
		if o.OnlyAlive != nil {
			qrOnlyAlive = *o.OnlyAlive
		}
		qOnlyAlive := swag.FormatBool(qrOnlyAlive)
		if qOnlyAlive != "" {
			if err := r.SetQueryParam("onlyAlive", qOnlyAlive); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
