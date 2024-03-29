// Code generated by go-swagger; DO NOT EDIT.

package v4_workspace_id_recipes

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

// NewGetRecipeRequestFromNameInWorkspaceParams creates a new GetRecipeRequestFromNameInWorkspaceParams object
// with the default values initialized.
func NewGetRecipeRequestFromNameInWorkspaceParams() *GetRecipeRequestFromNameInWorkspaceParams {
	var ()
	return &GetRecipeRequestFromNameInWorkspaceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRecipeRequestFromNameInWorkspaceParamsWithTimeout creates a new GetRecipeRequestFromNameInWorkspaceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRecipeRequestFromNameInWorkspaceParamsWithTimeout(timeout time.Duration) *GetRecipeRequestFromNameInWorkspaceParams {
	var ()
	return &GetRecipeRequestFromNameInWorkspaceParams{

		timeout: timeout,
	}
}

// NewGetRecipeRequestFromNameInWorkspaceParamsWithContext creates a new GetRecipeRequestFromNameInWorkspaceParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRecipeRequestFromNameInWorkspaceParamsWithContext(ctx context.Context) *GetRecipeRequestFromNameInWorkspaceParams {
	var ()
	return &GetRecipeRequestFromNameInWorkspaceParams{

		Context: ctx,
	}
}

// NewGetRecipeRequestFromNameInWorkspaceParamsWithHTTPClient creates a new GetRecipeRequestFromNameInWorkspaceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRecipeRequestFromNameInWorkspaceParamsWithHTTPClient(client *http.Client) *GetRecipeRequestFromNameInWorkspaceParams {
	var ()
	return &GetRecipeRequestFromNameInWorkspaceParams{
		HTTPClient: client,
	}
}

/*
GetRecipeRequestFromNameInWorkspaceParams contains all the parameters to send to the API endpoint
for the get recipe request from name in workspace operation typically these are written to a http.Request
*/
type GetRecipeRequestFromNameInWorkspaceParams struct {

	/*Name*/
	Name string
	/*WorkspaceID*/
	WorkspaceID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get recipe request from name in workspace params
func (o *GetRecipeRequestFromNameInWorkspaceParams) WithTimeout(timeout time.Duration) *GetRecipeRequestFromNameInWorkspaceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get recipe request from name in workspace params
func (o *GetRecipeRequestFromNameInWorkspaceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get recipe request from name in workspace params
func (o *GetRecipeRequestFromNameInWorkspaceParams) WithContext(ctx context.Context) *GetRecipeRequestFromNameInWorkspaceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get recipe request from name in workspace params
func (o *GetRecipeRequestFromNameInWorkspaceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get recipe request from name in workspace params
func (o *GetRecipeRequestFromNameInWorkspaceParams) WithHTTPClient(client *http.Client) *GetRecipeRequestFromNameInWorkspaceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get recipe request from name in workspace params
func (o *GetRecipeRequestFromNameInWorkspaceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the get recipe request from name in workspace params
func (o *GetRecipeRequestFromNameInWorkspaceParams) WithName(name string) *GetRecipeRequestFromNameInWorkspaceParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get recipe request from name in workspace params
func (o *GetRecipeRequestFromNameInWorkspaceParams) SetName(name string) {
	o.Name = name
}

// WithWorkspaceID adds the workspaceID to the get recipe request from name in workspace params
func (o *GetRecipeRequestFromNameInWorkspaceParams) WithWorkspaceID(workspaceID int64) *GetRecipeRequestFromNameInWorkspaceParams {
	o.SetWorkspaceID(workspaceID)
	return o
}

// SetWorkspaceID adds the workspaceId to the get recipe request from name in workspace params
func (o *GetRecipeRequestFromNameInWorkspaceParams) SetWorkspaceID(workspaceID int64) {
	o.WorkspaceID = workspaceID
}

// WriteToRequest writes these params to a swagger request
func (o *GetRecipeRequestFromNameInWorkspaceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
