// Code generated by go-swagger; DO NOT EDIT.

package v3_workspace_id_blueprints

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

// NewDeleteBlueprintInWorkspaceParams creates a new DeleteBlueprintInWorkspaceParams object
// with the default values initialized.
func NewDeleteBlueprintInWorkspaceParams() *DeleteBlueprintInWorkspaceParams {
	var ()
	return &DeleteBlueprintInWorkspaceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteBlueprintInWorkspaceParamsWithTimeout creates a new DeleteBlueprintInWorkspaceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteBlueprintInWorkspaceParamsWithTimeout(timeout time.Duration) *DeleteBlueprintInWorkspaceParams {
	var ()
	return &DeleteBlueprintInWorkspaceParams{

		timeout: timeout,
	}
}

// NewDeleteBlueprintInWorkspaceParamsWithContext creates a new DeleteBlueprintInWorkspaceParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteBlueprintInWorkspaceParamsWithContext(ctx context.Context) *DeleteBlueprintInWorkspaceParams {
	var ()
	return &DeleteBlueprintInWorkspaceParams{

		Context: ctx,
	}
}

// NewDeleteBlueprintInWorkspaceParamsWithHTTPClient creates a new DeleteBlueprintInWorkspaceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteBlueprintInWorkspaceParamsWithHTTPClient(client *http.Client) *DeleteBlueprintInWorkspaceParams {
	var ()
	return &DeleteBlueprintInWorkspaceParams{
		HTTPClient: client,
	}
}

/*DeleteBlueprintInWorkspaceParams contains all the parameters to send to the API endpoint
for the delete blueprint in workspace operation typically these are written to a http.Request
*/
type DeleteBlueprintInWorkspaceParams struct {

	/*Name*/
	Name string
	/*WorkspaceID*/
	WorkspaceID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete blueprint in workspace params
func (o *DeleteBlueprintInWorkspaceParams) WithTimeout(timeout time.Duration) *DeleteBlueprintInWorkspaceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete blueprint in workspace params
func (o *DeleteBlueprintInWorkspaceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete blueprint in workspace params
func (o *DeleteBlueprintInWorkspaceParams) WithContext(ctx context.Context) *DeleteBlueprintInWorkspaceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete blueprint in workspace params
func (o *DeleteBlueprintInWorkspaceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete blueprint in workspace params
func (o *DeleteBlueprintInWorkspaceParams) WithHTTPClient(client *http.Client) *DeleteBlueprintInWorkspaceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete blueprint in workspace params
func (o *DeleteBlueprintInWorkspaceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the delete blueprint in workspace params
func (o *DeleteBlueprintInWorkspaceParams) WithName(name string) *DeleteBlueprintInWorkspaceParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the delete blueprint in workspace params
func (o *DeleteBlueprintInWorkspaceParams) SetName(name string) {
	o.Name = name
}

// WithWorkspaceID adds the workspaceID to the delete blueprint in workspace params
func (o *DeleteBlueprintInWorkspaceParams) WithWorkspaceID(workspaceID int64) *DeleteBlueprintInWorkspaceParams {
	o.SetWorkspaceID(workspaceID)
	return o
}

// SetWorkspaceID adds the workspaceId to the delete blueprint in workspace params
func (o *DeleteBlueprintInWorkspaceParams) SetWorkspaceID(workspaceID int64) {
	o.WorkspaceID = workspaceID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteBlueprintInWorkspaceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
