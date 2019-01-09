// Code generated by go-swagger; DO NOT EDIT.

package v4_workspace_id_stacks

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

// NewStartStackInWorkspaceV4Params creates a new StartStackInWorkspaceV4Params object
// with the default values initialized.
func NewStartStackInWorkspaceV4Params() *StartStackInWorkspaceV4Params {
	var ()
	return &StartStackInWorkspaceV4Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewStartStackInWorkspaceV4ParamsWithTimeout creates a new StartStackInWorkspaceV4Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewStartStackInWorkspaceV4ParamsWithTimeout(timeout time.Duration) *StartStackInWorkspaceV4Params {
	var ()
	return &StartStackInWorkspaceV4Params{

		timeout: timeout,
	}
}

// NewStartStackInWorkspaceV4ParamsWithContext creates a new StartStackInWorkspaceV4Params object
// with the default values initialized, and the ability to set a context for a request
func NewStartStackInWorkspaceV4ParamsWithContext(ctx context.Context) *StartStackInWorkspaceV4Params {
	var ()
	return &StartStackInWorkspaceV4Params{

		Context: ctx,
	}
}

// NewStartStackInWorkspaceV4ParamsWithHTTPClient creates a new StartStackInWorkspaceV4Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewStartStackInWorkspaceV4ParamsWithHTTPClient(client *http.Client) *StartStackInWorkspaceV4Params {
	var ()
	return &StartStackInWorkspaceV4Params{
		HTTPClient: client,
	}
}

/*StartStackInWorkspaceV4Params contains all the parameters to send to the API endpoint
for the start stack in workspace v4 operation typically these are written to a http.Request
*/
type StartStackInWorkspaceV4Params struct {

	/*Name*/
	Name string
	/*WorkspaceID*/
	WorkspaceID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the start stack in workspace v4 params
func (o *StartStackInWorkspaceV4Params) WithTimeout(timeout time.Duration) *StartStackInWorkspaceV4Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the start stack in workspace v4 params
func (o *StartStackInWorkspaceV4Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the start stack in workspace v4 params
func (o *StartStackInWorkspaceV4Params) WithContext(ctx context.Context) *StartStackInWorkspaceV4Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the start stack in workspace v4 params
func (o *StartStackInWorkspaceV4Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the start stack in workspace v4 params
func (o *StartStackInWorkspaceV4Params) WithHTTPClient(client *http.Client) *StartStackInWorkspaceV4Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the start stack in workspace v4 params
func (o *StartStackInWorkspaceV4Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the start stack in workspace v4 params
func (o *StartStackInWorkspaceV4Params) WithName(name string) *StartStackInWorkspaceV4Params {
	o.SetName(name)
	return o
}

// SetName adds the name to the start stack in workspace v4 params
func (o *StartStackInWorkspaceV4Params) SetName(name string) {
	o.Name = name
}

// WithWorkspaceID adds the workspaceID to the start stack in workspace v4 params
func (o *StartStackInWorkspaceV4Params) WithWorkspaceID(workspaceID int64) *StartStackInWorkspaceV4Params {
	o.SetWorkspaceID(workspaceID)
	return o
}

// SetWorkspaceID adds the workspaceId to the start stack in workspace v4 params
func (o *StartStackInWorkspaceV4Params) SetWorkspaceID(workspaceID int64) {
	o.WorkspaceID = workspaceID
}

// WriteToRequest writes these params to a swagger request
func (o *StartStackInWorkspaceV4Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	// path param workspaceId
	if err := r.SetPathParam("workspaceId", swag.FormatInt64(o.WorkspaceID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
