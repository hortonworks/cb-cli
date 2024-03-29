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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetEventsByResourceNameInAccountParams creates a new GetEventsByResourceNameInAccountParams object
// with the default values initialized.
func NewGetEventsByResourceNameInAccountParams() *GetEventsByResourceNameInAccountParams {
	var (
		pageDefault = int32(0)
		sizeDefault = int32(100)
	)
	return &GetEventsByResourceNameInAccountParams{
		Page: &pageDefault,
		Size: &sizeDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetEventsByResourceNameInAccountParamsWithTimeout creates a new GetEventsByResourceNameInAccountParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetEventsByResourceNameInAccountParamsWithTimeout(timeout time.Duration) *GetEventsByResourceNameInAccountParams {
	var (
		pageDefault = int32(0)
		sizeDefault = int32(100)
	)
	return &GetEventsByResourceNameInAccountParams{
		Page: &pageDefault,
		Size: &sizeDefault,

		timeout: timeout,
	}
}

// NewGetEventsByResourceNameInAccountParamsWithContext creates a new GetEventsByResourceNameInAccountParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetEventsByResourceNameInAccountParamsWithContext(ctx context.Context) *GetEventsByResourceNameInAccountParams {
	var (
		pageDefault = int32(0)
		sizeDefault = int32(100)
	)
	return &GetEventsByResourceNameInAccountParams{
		Page: &pageDefault,
		Size: &sizeDefault,

		Context: ctx,
	}
}

// NewGetEventsByResourceNameInAccountParamsWithHTTPClient creates a new GetEventsByResourceNameInAccountParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetEventsByResourceNameInAccountParamsWithHTTPClient(client *http.Client) *GetEventsByResourceNameInAccountParams {
	var (
		pageDefault = int32(0)
		sizeDefault = int32(100)
	)
	return &GetEventsByResourceNameInAccountParams{
		Page:       &pageDefault,
		Size:       &sizeDefault,
		HTTPClient: client,
	}
}

/*
GetEventsByResourceNameInAccountParams contains all the parameters to send to the API endpoint
for the get events by resource name in account operation typically these are written to a http.Request
*/
type GetEventsByResourceNameInAccountParams struct {

	/*Name*/
	Name string
	/*Page*/
	Page *int32
	/*Size*/
	Size *int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get events by resource name in account params
func (o *GetEventsByResourceNameInAccountParams) WithTimeout(timeout time.Duration) *GetEventsByResourceNameInAccountParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get events by resource name in account params
func (o *GetEventsByResourceNameInAccountParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get events by resource name in account params
func (o *GetEventsByResourceNameInAccountParams) WithContext(ctx context.Context) *GetEventsByResourceNameInAccountParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get events by resource name in account params
func (o *GetEventsByResourceNameInAccountParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get events by resource name in account params
func (o *GetEventsByResourceNameInAccountParams) WithHTTPClient(client *http.Client) *GetEventsByResourceNameInAccountParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get events by resource name in account params
func (o *GetEventsByResourceNameInAccountParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the get events by resource name in account params
func (o *GetEventsByResourceNameInAccountParams) WithName(name string) *GetEventsByResourceNameInAccountParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get events by resource name in account params
func (o *GetEventsByResourceNameInAccountParams) SetName(name string) {
	o.Name = name
}

// WithPage adds the page to the get events by resource name in account params
func (o *GetEventsByResourceNameInAccountParams) WithPage(page *int32) *GetEventsByResourceNameInAccountParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get events by resource name in account params
func (o *GetEventsByResourceNameInAccountParams) SetPage(page *int32) {
	o.Page = page
}

// WithSize adds the size to the get events by resource name in account params
func (o *GetEventsByResourceNameInAccountParams) WithSize(size *int32) *GetEventsByResourceNameInAccountParams {
	o.SetSize(size)
	return o
}

// SetSize adds the size to the get events by resource name in account params
func (o *GetEventsByResourceNameInAccountParams) SetSize(size *int32) {
	o.Size = size
}

// WriteToRequest writes these params to a swagger request
func (o *GetEventsByResourceNameInAccountParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if o.Page != nil {

		// query param page
		var qrPage int32
		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt32(qrPage)
		if qPage != "" {
			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}

	}

	if o.Size != nil {

		// query param size
		var qrSize int32
		if o.Size != nil {
			qrSize = *o.Size
		}
		qSize := swag.FormatInt32(qrSize)
		if qSize != "" {
			if err := r.SetQueryParam("size", qSize); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
