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

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetStructuredEventsZipInWorkspaceParams creates a new GetStructuredEventsZipInWorkspaceParams object
// with the default values initialized.
func NewGetStructuredEventsZipInWorkspaceParams() *GetStructuredEventsZipInWorkspaceParams {
	var ()
	return &GetStructuredEventsZipInWorkspaceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetStructuredEventsZipInWorkspaceParamsWithTimeout creates a new GetStructuredEventsZipInWorkspaceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetStructuredEventsZipInWorkspaceParamsWithTimeout(timeout time.Duration) *GetStructuredEventsZipInWorkspaceParams {
	var ()
	return &GetStructuredEventsZipInWorkspaceParams{

		timeout: timeout,
	}
}

// NewGetStructuredEventsZipInWorkspaceParamsWithContext creates a new GetStructuredEventsZipInWorkspaceParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetStructuredEventsZipInWorkspaceParamsWithContext(ctx context.Context) *GetStructuredEventsZipInWorkspaceParams {
	var ()
	return &GetStructuredEventsZipInWorkspaceParams{

		Context: ctx,
	}
}

// NewGetStructuredEventsZipInWorkspaceParamsWithHTTPClient creates a new GetStructuredEventsZipInWorkspaceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetStructuredEventsZipInWorkspaceParamsWithHTTPClient(client *http.Client) *GetStructuredEventsZipInWorkspaceParams {
	var ()
	return &GetStructuredEventsZipInWorkspaceParams{
		HTTPClient: client,
	}
}

/*GetStructuredEventsZipInWorkspaceParams contains all the parameters to send to the API endpoint
for the get structured events zip in workspace operation typically these are written to a http.Request
*/
type GetStructuredEventsZipInWorkspaceParams struct {

	/*Name*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get structured events zip in workspace params
func (o *GetStructuredEventsZipInWorkspaceParams) WithTimeout(timeout time.Duration) *GetStructuredEventsZipInWorkspaceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get structured events zip in workspace params
func (o *GetStructuredEventsZipInWorkspaceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get structured events zip in workspace params
func (o *GetStructuredEventsZipInWorkspaceParams) WithContext(ctx context.Context) *GetStructuredEventsZipInWorkspaceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get structured events zip in workspace params
func (o *GetStructuredEventsZipInWorkspaceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get structured events zip in workspace params
func (o *GetStructuredEventsZipInWorkspaceParams) WithHTTPClient(client *http.Client) *GetStructuredEventsZipInWorkspaceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get structured events zip in workspace params
func (o *GetStructuredEventsZipInWorkspaceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the get structured events zip in workspace params
func (o *GetStructuredEventsZipInWorkspaceParams) WithName(name string) *GetStructuredEventsZipInWorkspaceParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get structured events zip in workspace params
func (o *GetStructuredEventsZipInWorkspaceParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *GetStructuredEventsZipInWorkspaceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
