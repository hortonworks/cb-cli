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

// NewDeleteRecipesInWorkspaceParams creates a new DeleteRecipesInWorkspaceParams object
// with the default values initialized.
func NewDeleteRecipesInWorkspaceParams() *DeleteRecipesInWorkspaceParams {
	var ()
	return &DeleteRecipesInWorkspaceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteRecipesInWorkspaceParamsWithTimeout creates a new DeleteRecipesInWorkspaceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteRecipesInWorkspaceParamsWithTimeout(timeout time.Duration) *DeleteRecipesInWorkspaceParams {
	var ()
	return &DeleteRecipesInWorkspaceParams{

		timeout: timeout,
	}
}

// NewDeleteRecipesInWorkspaceParamsWithContext creates a new DeleteRecipesInWorkspaceParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteRecipesInWorkspaceParamsWithContext(ctx context.Context) *DeleteRecipesInWorkspaceParams {
	var ()
	return &DeleteRecipesInWorkspaceParams{

		Context: ctx,
	}
}

// NewDeleteRecipesInWorkspaceParamsWithHTTPClient creates a new DeleteRecipesInWorkspaceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteRecipesInWorkspaceParamsWithHTTPClient(client *http.Client) *DeleteRecipesInWorkspaceParams {
	var ()
	return &DeleteRecipesInWorkspaceParams{
		HTTPClient: client,
	}
}

/*DeleteRecipesInWorkspaceParams contains all the parameters to send to the API endpoint
for the delete recipes in workspace operation typically these are written to a http.Request
*/
type DeleteRecipesInWorkspaceParams struct {

	/*Body*/
	Body []string
	/*WorkspaceID*/
	WorkspaceID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete recipes in workspace params
func (o *DeleteRecipesInWorkspaceParams) WithTimeout(timeout time.Duration) *DeleteRecipesInWorkspaceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete recipes in workspace params
func (o *DeleteRecipesInWorkspaceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete recipes in workspace params
func (o *DeleteRecipesInWorkspaceParams) WithContext(ctx context.Context) *DeleteRecipesInWorkspaceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete recipes in workspace params
func (o *DeleteRecipesInWorkspaceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete recipes in workspace params
func (o *DeleteRecipesInWorkspaceParams) WithHTTPClient(client *http.Client) *DeleteRecipesInWorkspaceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete recipes in workspace params
func (o *DeleteRecipesInWorkspaceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the delete recipes in workspace params
func (o *DeleteRecipesInWorkspaceParams) WithBody(body []string) *DeleteRecipesInWorkspaceParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the delete recipes in workspace params
func (o *DeleteRecipesInWorkspaceParams) SetBody(body []string) {
	o.Body = body
}

// WithWorkspaceID adds the workspaceID to the delete recipes in workspace params
func (o *DeleteRecipesInWorkspaceParams) WithWorkspaceID(workspaceID int64) *DeleteRecipesInWorkspaceParams {
	o.SetWorkspaceID(workspaceID)
	return o
}

// SetWorkspaceID adds the workspaceId to the delete recipes in workspace params
func (o *DeleteRecipesInWorkspaceParams) SetWorkspaceID(workspaceID int64) {
	o.WorkspaceID = workspaceID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteRecipesInWorkspaceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
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
