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

// NewDeleteStackInWorkspaceV4Params creates a new DeleteStackInWorkspaceV4Params object
// with the default values initialized.
func NewDeleteStackInWorkspaceV4Params() *DeleteStackInWorkspaceV4Params {
	var (
		deleteDependenciesDefault = bool(false)
		forcedDefault             = bool(false)
	)
	return &DeleteStackInWorkspaceV4Params{
		DeleteDependencies: &deleteDependenciesDefault,
		Forced:             &forcedDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteStackInWorkspaceV4ParamsWithTimeout creates a new DeleteStackInWorkspaceV4Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteStackInWorkspaceV4ParamsWithTimeout(timeout time.Duration) *DeleteStackInWorkspaceV4Params {
	var (
		deleteDependenciesDefault = bool(false)
		forcedDefault             = bool(false)
	)
	return &DeleteStackInWorkspaceV4Params{
		DeleteDependencies: &deleteDependenciesDefault,
		Forced:             &forcedDefault,

		timeout: timeout,
	}
}

// NewDeleteStackInWorkspaceV4ParamsWithContext creates a new DeleteStackInWorkspaceV4Params object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteStackInWorkspaceV4ParamsWithContext(ctx context.Context) *DeleteStackInWorkspaceV4Params {
	var (
		deleteDependenciesDefault = bool(false)
		forcedDefault             = bool(false)
	)
	return &DeleteStackInWorkspaceV4Params{
		DeleteDependencies: &deleteDependenciesDefault,
		Forced:             &forcedDefault,

		Context: ctx,
	}
}

// NewDeleteStackInWorkspaceV4ParamsWithHTTPClient creates a new DeleteStackInWorkspaceV4Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteStackInWorkspaceV4ParamsWithHTTPClient(client *http.Client) *DeleteStackInWorkspaceV4Params {
	var (
		deleteDependenciesDefault = bool(false)
		forcedDefault             = bool(false)
	)
	return &DeleteStackInWorkspaceV4Params{
		DeleteDependencies: &deleteDependenciesDefault,
		Forced:             &forcedDefault,
		HTTPClient:         client,
	}
}

/*DeleteStackInWorkspaceV4Params contains all the parameters to send to the API endpoint
for the delete stack in workspace v4 operation typically these are written to a http.Request
*/
type DeleteStackInWorkspaceV4Params struct {

	/*DeleteDependencies*/
	DeleteDependencies *bool
	/*Forced*/
	Forced *bool
	/*Name*/
	Name string
	/*WorkspaceID*/
	WorkspaceID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete stack in workspace v4 params
func (o *DeleteStackInWorkspaceV4Params) WithTimeout(timeout time.Duration) *DeleteStackInWorkspaceV4Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete stack in workspace v4 params
func (o *DeleteStackInWorkspaceV4Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete stack in workspace v4 params
func (o *DeleteStackInWorkspaceV4Params) WithContext(ctx context.Context) *DeleteStackInWorkspaceV4Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete stack in workspace v4 params
func (o *DeleteStackInWorkspaceV4Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete stack in workspace v4 params
func (o *DeleteStackInWorkspaceV4Params) WithHTTPClient(client *http.Client) *DeleteStackInWorkspaceV4Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete stack in workspace v4 params
func (o *DeleteStackInWorkspaceV4Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDeleteDependencies adds the deleteDependencies to the delete stack in workspace v4 params
func (o *DeleteStackInWorkspaceV4Params) WithDeleteDependencies(deleteDependencies *bool) *DeleteStackInWorkspaceV4Params {
	o.SetDeleteDependencies(deleteDependencies)
	return o
}

// SetDeleteDependencies adds the deleteDependencies to the delete stack in workspace v4 params
func (o *DeleteStackInWorkspaceV4Params) SetDeleteDependencies(deleteDependencies *bool) {
	o.DeleteDependencies = deleteDependencies
}

// WithForced adds the forced to the delete stack in workspace v4 params
func (o *DeleteStackInWorkspaceV4Params) WithForced(forced *bool) *DeleteStackInWorkspaceV4Params {
	o.SetForced(forced)
	return o
}

// SetForced adds the forced to the delete stack in workspace v4 params
func (o *DeleteStackInWorkspaceV4Params) SetForced(forced *bool) {
	o.Forced = forced
}

// WithName adds the name to the delete stack in workspace v4 params
func (o *DeleteStackInWorkspaceV4Params) WithName(name string) *DeleteStackInWorkspaceV4Params {
	o.SetName(name)
	return o
}

// SetName adds the name to the delete stack in workspace v4 params
func (o *DeleteStackInWorkspaceV4Params) SetName(name string) {
	o.Name = name
}

// WithWorkspaceID adds the workspaceID to the delete stack in workspace v4 params
func (o *DeleteStackInWorkspaceV4Params) WithWorkspaceID(workspaceID int64) *DeleteStackInWorkspaceV4Params {
	o.SetWorkspaceID(workspaceID)
	return o
}

// SetWorkspaceID adds the workspaceId to the delete stack in workspace v4 params
func (o *DeleteStackInWorkspaceV4Params) SetWorkspaceID(workspaceID int64) {
	o.WorkspaceID = workspaceID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteStackInWorkspaceV4Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DeleteDependencies != nil {

		// query param deleteDependencies
		var qrDeleteDependencies bool
		if o.DeleteDependencies != nil {
			qrDeleteDependencies = *o.DeleteDependencies
		}
		qDeleteDependencies := swag.FormatBool(qrDeleteDependencies)
		if qDeleteDependencies != "" {
			if err := r.SetQueryParam("deleteDependencies", qDeleteDependencies); err != nil {
				return err
			}
		}

	}

	if o.Forced != nil {

		// query param forced
		var qrForced bool
		if o.Forced != nil {
			qrForced = *o.Forced
		}
		qForced := swag.FormatBool(qrForced)
		if qForced != "" {
			if err := r.SetQueryParam("forced", qForced); err != nil {
				return err
			}
		}

	}

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
