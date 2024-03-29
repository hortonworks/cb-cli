// Code generated by go-swagger; DO NOT EDIT.

package v4_workspace_id_blueprints

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

	model "github.com/hortonworks/cb-cli/dataplane/api/model"
)

// NewCreateBlueprintInWorkspaceParams creates a new CreateBlueprintInWorkspaceParams object
// with the default values initialized.
func NewCreateBlueprintInWorkspaceParams() *CreateBlueprintInWorkspaceParams {
	var ()
	return &CreateBlueprintInWorkspaceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateBlueprintInWorkspaceParamsWithTimeout creates a new CreateBlueprintInWorkspaceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateBlueprintInWorkspaceParamsWithTimeout(timeout time.Duration) *CreateBlueprintInWorkspaceParams {
	var ()
	return &CreateBlueprintInWorkspaceParams{

		timeout: timeout,
	}
}

// NewCreateBlueprintInWorkspaceParamsWithContext creates a new CreateBlueprintInWorkspaceParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateBlueprintInWorkspaceParamsWithContext(ctx context.Context) *CreateBlueprintInWorkspaceParams {
	var ()
	return &CreateBlueprintInWorkspaceParams{

		Context: ctx,
	}
}

// NewCreateBlueprintInWorkspaceParamsWithHTTPClient creates a new CreateBlueprintInWorkspaceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateBlueprintInWorkspaceParamsWithHTTPClient(client *http.Client) *CreateBlueprintInWorkspaceParams {
	var ()
	return &CreateBlueprintInWorkspaceParams{
		HTTPClient: client,
	}
}

/*
CreateBlueprintInWorkspaceParams contains all the parameters to send to the API endpoint
for the create blueprint in workspace operation typically these are written to a http.Request
*/
type CreateBlueprintInWorkspaceParams struct {

	/*Body*/
	Body *model.BlueprintV4Request
	/*WorkspaceID*/
	WorkspaceID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create blueprint in workspace params
func (o *CreateBlueprintInWorkspaceParams) WithTimeout(timeout time.Duration) *CreateBlueprintInWorkspaceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create blueprint in workspace params
func (o *CreateBlueprintInWorkspaceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create blueprint in workspace params
func (o *CreateBlueprintInWorkspaceParams) WithContext(ctx context.Context) *CreateBlueprintInWorkspaceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create blueprint in workspace params
func (o *CreateBlueprintInWorkspaceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create blueprint in workspace params
func (o *CreateBlueprintInWorkspaceParams) WithHTTPClient(client *http.Client) *CreateBlueprintInWorkspaceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create blueprint in workspace params
func (o *CreateBlueprintInWorkspaceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create blueprint in workspace params
func (o *CreateBlueprintInWorkspaceParams) WithBody(body *model.BlueprintV4Request) *CreateBlueprintInWorkspaceParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create blueprint in workspace params
func (o *CreateBlueprintInWorkspaceParams) SetBody(body *model.BlueprintV4Request) {
	o.Body = body
}

// WithWorkspaceID adds the workspaceID to the create blueprint in workspace params
func (o *CreateBlueprintInWorkspaceParams) WithWorkspaceID(workspaceID int64) *CreateBlueprintInWorkspaceParams {
	o.SetWorkspaceID(workspaceID)
	return o
}

// SetWorkspaceID adds the workspaceId to the create blueprint in workspace params
func (o *CreateBlueprintInWorkspaceParams) SetWorkspaceID(workspaceID int64) {
	o.WorkspaceID = workspaceID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateBlueprintInWorkspaceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
