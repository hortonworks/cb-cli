// Code generated by go-swagger; DO NOT EDIT.

package v1stacks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetPublicStackParams creates a new GetPublicStackParams object
// with the default values initialized.
func NewGetPublicStackParams() *GetPublicStackParams {
	var ()
	return &GetPublicStackParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPublicStackParamsWithTimeout creates a new GetPublicStackParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPublicStackParamsWithTimeout(timeout time.Duration) *GetPublicStackParams {
	var ()
	return &GetPublicStackParams{

		timeout: timeout,
	}
}

// NewGetPublicStackParamsWithContext creates a new GetPublicStackParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPublicStackParamsWithContext(ctx context.Context) *GetPublicStackParams {
	var ()
	return &GetPublicStackParams{

		Context: ctx,
	}
}

// NewGetPublicStackParamsWithHTTPClient creates a new GetPublicStackParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPublicStackParamsWithHTTPClient(client *http.Client) *GetPublicStackParams {
	var ()
	return &GetPublicStackParams{
		HTTPClient: client,
	}
}

/*GetPublicStackParams contains all the parameters to send to the API endpoint
for the get public stack operation typically these are written to a http.Request
*/
type GetPublicStackParams struct {

	/*Entry*/
	Entry []string
	/*Name*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get public stack params
func (o *GetPublicStackParams) WithTimeout(timeout time.Duration) *GetPublicStackParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get public stack params
func (o *GetPublicStackParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get public stack params
func (o *GetPublicStackParams) WithContext(ctx context.Context) *GetPublicStackParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get public stack params
func (o *GetPublicStackParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get public stack params
func (o *GetPublicStackParams) WithHTTPClient(client *http.Client) *GetPublicStackParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get public stack params
func (o *GetPublicStackParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEntry adds the entry to the get public stack params
func (o *GetPublicStackParams) WithEntry(entry []string) *GetPublicStackParams {
	o.SetEntry(entry)
	return o
}

// SetEntry adds the entry to the get public stack params
func (o *GetPublicStackParams) SetEntry(entry []string) {
	o.Entry = entry
}

// WithName adds the name to the get public stack params
func (o *GetPublicStackParams) WithName(name string) *GetPublicStackParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get public stack params
func (o *GetPublicStackParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *GetPublicStackParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	valuesEntry := o.Entry

	joinedEntry := swag.JoinByFormat(valuesEntry, "multi")
	// query array param entry
	if err := r.SetQueryParam("entry", joinedEntry...); err != nil {
		return err
	}

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}