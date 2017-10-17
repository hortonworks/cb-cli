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

// NewGetStackParams creates a new GetStackParams object
// with the default values initialized.
func NewGetStackParams() *GetStackParams {
	var ()
	return &GetStackParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetStackParamsWithTimeout creates a new GetStackParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetStackParamsWithTimeout(timeout time.Duration) *GetStackParams {
	var ()
	return &GetStackParams{

		timeout: timeout,
	}
}

// NewGetStackParamsWithContext creates a new GetStackParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetStackParamsWithContext(ctx context.Context) *GetStackParams {
	var ()
	return &GetStackParams{

		Context: ctx,
	}
}

// NewGetStackParamsWithHTTPClient creates a new GetStackParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetStackParamsWithHTTPClient(client *http.Client) *GetStackParams {
	var ()
	return &GetStackParams{
		HTTPClient: client,
	}
}

/*GetStackParams contains all the parameters to send to the API endpoint
for the get stack operation typically these are written to a http.Request
*/
type GetStackParams struct {

	/*Entry*/
	Entry []string
	/*ID*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get stack params
func (o *GetStackParams) WithTimeout(timeout time.Duration) *GetStackParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get stack params
func (o *GetStackParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get stack params
func (o *GetStackParams) WithContext(ctx context.Context) *GetStackParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get stack params
func (o *GetStackParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get stack params
func (o *GetStackParams) WithHTTPClient(client *http.Client) *GetStackParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get stack params
func (o *GetStackParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEntry adds the entry to the get stack params
func (o *GetStackParams) WithEntry(entry []string) *GetStackParams {
	o.SetEntry(entry)
	return o
}

// SetEntry adds the entry to the get stack params
func (o *GetStackParams) SetEntry(entry []string) {
	o.Entry = entry
}

// WithID adds the id to the get stack params
func (o *GetStackParams) WithID(id int64) *GetStackParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get stack params
func (o *GetStackParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetStackParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
