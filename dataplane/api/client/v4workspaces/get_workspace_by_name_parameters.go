// Code generated by go-swagger; DO NOT EDIT.

package v4workspaces

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetWorkspaceByNameParams creates a new GetWorkspaceByNameParams object
// with the default values initialized.
func NewGetWorkspaceByNameParams() *GetWorkspaceByNameParams {
	var ()
	return &GetWorkspaceByNameParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetWorkspaceByNameParamsWithTimeout creates a new GetWorkspaceByNameParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetWorkspaceByNameParamsWithTimeout(timeout time.Duration) *GetWorkspaceByNameParams {
	var ()
	return &GetWorkspaceByNameParams{

		timeout: timeout,
	}
}

// NewGetWorkspaceByNameParamsWithContext creates a new GetWorkspaceByNameParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetWorkspaceByNameParamsWithContext(ctx context.Context) *GetWorkspaceByNameParams {
	var ()
	return &GetWorkspaceByNameParams{

		Context: ctx,
	}
}

// NewGetWorkspaceByNameParamsWithHTTPClient creates a new GetWorkspaceByNameParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetWorkspaceByNameParamsWithHTTPClient(client *http.Client) *GetWorkspaceByNameParams {
	var ()
	return &GetWorkspaceByNameParams{
		HTTPClient: client,
	}
}

/*GetWorkspaceByNameParams contains all the parameters to send to the API endpoint
for the get workspace by name operation typically these are written to a http.Request
*/
type GetWorkspaceByNameParams struct {

	/*Name*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get workspace by name params
func (o *GetWorkspaceByNameParams) WithTimeout(timeout time.Duration) *GetWorkspaceByNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get workspace by name params
func (o *GetWorkspaceByNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get workspace by name params
func (o *GetWorkspaceByNameParams) WithContext(ctx context.Context) *GetWorkspaceByNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get workspace by name params
func (o *GetWorkspaceByNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get workspace by name params
func (o *GetWorkspaceByNameParams) WithHTTPClient(client *http.Client) *GetWorkspaceByNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get workspace by name params
func (o *GetWorkspaceByNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the get workspace by name params
func (o *GetWorkspaceByNameParams) WithName(name string) *GetWorkspaceByNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get workspace by name params
func (o *GetWorkspaceByNameParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *GetWorkspaceByNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
