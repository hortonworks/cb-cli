// Code generated by go-swagger; DO NOT EDIT.

package v4_workspace_id_clustertemplates

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

// NewGetClusterTemplateByNameInWorkspaceParams creates a new GetClusterTemplateByNameInWorkspaceParams object
// with the default values initialized.
func NewGetClusterTemplateByNameInWorkspaceParams() *GetClusterTemplateByNameInWorkspaceParams {
	var ()
	return &GetClusterTemplateByNameInWorkspaceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetClusterTemplateByNameInWorkspaceParamsWithTimeout creates a new GetClusterTemplateByNameInWorkspaceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetClusterTemplateByNameInWorkspaceParamsWithTimeout(timeout time.Duration) *GetClusterTemplateByNameInWorkspaceParams {
	var ()
	return &GetClusterTemplateByNameInWorkspaceParams{

		timeout: timeout,
	}
}

// NewGetClusterTemplateByNameInWorkspaceParamsWithContext creates a new GetClusterTemplateByNameInWorkspaceParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetClusterTemplateByNameInWorkspaceParamsWithContext(ctx context.Context) *GetClusterTemplateByNameInWorkspaceParams {
	var ()
	return &GetClusterTemplateByNameInWorkspaceParams{

		Context: ctx,
	}
}

// NewGetClusterTemplateByNameInWorkspaceParamsWithHTTPClient creates a new GetClusterTemplateByNameInWorkspaceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetClusterTemplateByNameInWorkspaceParamsWithHTTPClient(client *http.Client) *GetClusterTemplateByNameInWorkspaceParams {
	var ()
	return &GetClusterTemplateByNameInWorkspaceParams{
		HTTPClient: client,
	}
}

/*GetClusterTemplateByNameInWorkspaceParams contains all the parameters to send to the API endpoint
for the get cluster template by name in workspace operation typically these are written to a http.Request
*/
type GetClusterTemplateByNameInWorkspaceParams struct {

	/*Name*/
	Name string
	/*WorkspaceID*/
	WorkspaceID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get cluster template by name in workspace params
func (o *GetClusterTemplateByNameInWorkspaceParams) WithTimeout(timeout time.Duration) *GetClusterTemplateByNameInWorkspaceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get cluster template by name in workspace params
func (o *GetClusterTemplateByNameInWorkspaceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get cluster template by name in workspace params
func (o *GetClusterTemplateByNameInWorkspaceParams) WithContext(ctx context.Context) *GetClusterTemplateByNameInWorkspaceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get cluster template by name in workspace params
func (o *GetClusterTemplateByNameInWorkspaceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get cluster template by name in workspace params
func (o *GetClusterTemplateByNameInWorkspaceParams) WithHTTPClient(client *http.Client) *GetClusterTemplateByNameInWorkspaceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get cluster template by name in workspace params
func (o *GetClusterTemplateByNameInWorkspaceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the get cluster template by name in workspace params
func (o *GetClusterTemplateByNameInWorkspaceParams) WithName(name string) *GetClusterTemplateByNameInWorkspaceParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get cluster template by name in workspace params
func (o *GetClusterTemplateByNameInWorkspaceParams) SetName(name string) {
	o.Name = name
}

// WithWorkspaceID adds the workspaceID to the get cluster template by name in workspace params
func (o *GetClusterTemplateByNameInWorkspaceParams) WithWorkspaceID(workspaceID int64) *GetClusterTemplateByNameInWorkspaceParams {
	o.SetWorkspaceID(workspaceID)
	return o
}

// SetWorkspaceID adds the workspaceId to the get cluster template by name in workspace params
func (o *GetClusterTemplateByNameInWorkspaceParams) SetWorkspaceID(workspaceID int64) {
	o.WorkspaceID = workspaceID
}

// WriteToRequest writes these params to a swagger request
func (o *GetClusterTemplateByNameInWorkspaceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
