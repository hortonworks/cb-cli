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

// NewGetRecipeRequestsFromNamesInWorkspaceParams creates a new GetRecipeRequestsFromNamesInWorkspaceParams object
// with the default values initialized.
func NewGetRecipeRequestsFromNamesInWorkspaceParams() *GetRecipeRequestsFromNamesInWorkspaceParams {
	var ()
	return &GetRecipeRequestsFromNamesInWorkspaceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRecipeRequestsFromNamesInWorkspaceParamsWithTimeout creates a new GetRecipeRequestsFromNamesInWorkspaceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRecipeRequestsFromNamesInWorkspaceParamsWithTimeout(timeout time.Duration) *GetRecipeRequestsFromNamesInWorkspaceParams {
	var ()
	return &GetRecipeRequestsFromNamesInWorkspaceParams{

		timeout: timeout,
	}
}

// NewGetRecipeRequestsFromNamesInWorkspaceParamsWithContext creates a new GetRecipeRequestsFromNamesInWorkspaceParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRecipeRequestsFromNamesInWorkspaceParamsWithContext(ctx context.Context) *GetRecipeRequestsFromNamesInWorkspaceParams {
	var ()
	return &GetRecipeRequestsFromNamesInWorkspaceParams{

		Context: ctx,
	}
}

// NewGetRecipeRequestsFromNamesInWorkspaceParamsWithHTTPClient creates a new GetRecipeRequestsFromNamesInWorkspaceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRecipeRequestsFromNamesInWorkspaceParamsWithHTTPClient(client *http.Client) *GetRecipeRequestsFromNamesInWorkspaceParams {
	var ()
	return &GetRecipeRequestsFromNamesInWorkspaceParams{
		HTTPClient: client,
	}
}

/*
GetRecipeRequestsFromNamesInWorkspaceParams contains all the parameters to send to the API endpoint
for the get recipe requests from names in workspace operation typically these are written to a http.Request
*/
type GetRecipeRequestsFromNamesInWorkspaceParams struct {

	/*Names*/
	Names []string
	/*WorkspaceID*/
	WorkspaceID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get recipe requests from names in workspace params
func (o *GetRecipeRequestsFromNamesInWorkspaceParams) WithTimeout(timeout time.Duration) *GetRecipeRequestsFromNamesInWorkspaceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get recipe requests from names in workspace params
func (o *GetRecipeRequestsFromNamesInWorkspaceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get recipe requests from names in workspace params
func (o *GetRecipeRequestsFromNamesInWorkspaceParams) WithContext(ctx context.Context) *GetRecipeRequestsFromNamesInWorkspaceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get recipe requests from names in workspace params
func (o *GetRecipeRequestsFromNamesInWorkspaceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get recipe requests from names in workspace params
func (o *GetRecipeRequestsFromNamesInWorkspaceParams) WithHTTPClient(client *http.Client) *GetRecipeRequestsFromNamesInWorkspaceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get recipe requests from names in workspace params
func (o *GetRecipeRequestsFromNamesInWorkspaceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNames adds the names to the get recipe requests from names in workspace params
func (o *GetRecipeRequestsFromNamesInWorkspaceParams) WithNames(names []string) *GetRecipeRequestsFromNamesInWorkspaceParams {
	o.SetNames(names)
	return o
}

// SetNames adds the names to the get recipe requests from names in workspace params
func (o *GetRecipeRequestsFromNamesInWorkspaceParams) SetNames(names []string) {
	o.Names = names
}

// WithWorkspaceID adds the workspaceID to the get recipe requests from names in workspace params
func (o *GetRecipeRequestsFromNamesInWorkspaceParams) WithWorkspaceID(workspaceID int64) *GetRecipeRequestsFromNamesInWorkspaceParams {
	o.SetWorkspaceID(workspaceID)
	return o
}

// SetWorkspaceID adds the workspaceId to the get recipe requests from names in workspace params
func (o *GetRecipeRequestsFromNamesInWorkspaceParams) SetWorkspaceID(workspaceID int64) {
	o.WorkspaceID = workspaceID
}

// WriteToRequest writes these params to a swagger request
func (o *GetRecipeRequestsFromNamesInWorkspaceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	valuesNames := o.Names

	joinedNames := swag.JoinByFormat(valuesNames, "multi")
	// query array param names
	if err := r.SetQueryParam("names", joinedNames...); err != nil {
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
