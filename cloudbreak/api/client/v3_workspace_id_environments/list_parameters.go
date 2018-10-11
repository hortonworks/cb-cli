// Code generated by go-swagger; DO NOT EDIT.

package v3_workspace_id_environments

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

// NewListParams creates a new ListParams object
// with the default values initialized.
func NewListParams() *ListParams {
	var ()
	return &ListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListParamsWithTimeout creates a new ListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListParamsWithTimeout(timeout time.Duration) *ListParams {
	var ()
	return &ListParams{

		timeout: timeout,
	}
}

// NewListParamsWithContext creates a new ListParams object
// with the default values initialized, and the ability to set a context for a request
func NewListParamsWithContext(ctx context.Context) *ListParams {
	var ()
	return &ListParams{

		Context: ctx,
	}
}

// NewListParamsWithHTTPClient creates a new ListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListParamsWithHTTPClient(client *http.Client) *ListParams {
	var ()
	return &ListParams{
		HTTPClient: client,
	}
}

/*ListParams contains all the parameters to send to the API endpoint
for the list operation typically these are written to a http.Request
*/
type ListParams struct {

	/*WorkspaceID*/
	WorkspaceID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list params
func (o *ListParams) WithTimeout(timeout time.Duration) *ListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list params
func (o *ListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list params
func (o *ListParams) WithContext(ctx context.Context) *ListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list params
func (o *ListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list params
func (o *ListParams) WithHTTPClient(client *http.Client) *ListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list params
func (o *ListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithWorkspaceID adds the workspaceID to the list params
func (o *ListParams) WithWorkspaceID(workspaceID int64) *ListParams {
	o.SetWorkspaceID(workspaceID)
	return o
}

// SetWorkspaceID adds the workspaceId to the list params
func (o *ListParams) SetWorkspaceID(workspaceID int64) {
	o.WorkspaceID = workspaceID
}

// WriteToRequest writes these params to a swagger request
func (o *ListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param workspaceId
	if err := r.SetPathParam("workspaceId", swag.FormatInt64(o.WorkspaceID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
