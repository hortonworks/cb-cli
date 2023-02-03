// Code generated by go-swagger; DO NOT EDIT.

package v4_workspace_id_imagecatalogs

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

// NewUpdateImageCatalogInWorkspaceParams creates a new UpdateImageCatalogInWorkspaceParams object
// with the default values initialized.
func NewUpdateImageCatalogInWorkspaceParams() *UpdateImageCatalogInWorkspaceParams {
	var ()
	return &UpdateImageCatalogInWorkspaceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateImageCatalogInWorkspaceParamsWithTimeout creates a new UpdateImageCatalogInWorkspaceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateImageCatalogInWorkspaceParamsWithTimeout(timeout time.Duration) *UpdateImageCatalogInWorkspaceParams {
	var ()
	return &UpdateImageCatalogInWorkspaceParams{

		timeout: timeout,
	}
}

// NewUpdateImageCatalogInWorkspaceParamsWithContext creates a new UpdateImageCatalogInWorkspaceParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateImageCatalogInWorkspaceParamsWithContext(ctx context.Context) *UpdateImageCatalogInWorkspaceParams {
	var ()
	return &UpdateImageCatalogInWorkspaceParams{

		Context: ctx,
	}
}

// NewUpdateImageCatalogInWorkspaceParamsWithHTTPClient creates a new UpdateImageCatalogInWorkspaceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateImageCatalogInWorkspaceParamsWithHTTPClient(client *http.Client) *UpdateImageCatalogInWorkspaceParams {
	var ()
	return &UpdateImageCatalogInWorkspaceParams{
		HTTPClient: client,
	}
}

/*
UpdateImageCatalogInWorkspaceParams contains all the parameters to send to the API endpoint
for the update image catalog in workspace operation typically these are written to a http.Request
*/
type UpdateImageCatalogInWorkspaceParams struct {

	/*Body*/
	Body *model.UpdateImageCatalogV4Request
	/*WorkspaceID*/
	WorkspaceID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update image catalog in workspace params
func (o *UpdateImageCatalogInWorkspaceParams) WithTimeout(timeout time.Duration) *UpdateImageCatalogInWorkspaceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update image catalog in workspace params
func (o *UpdateImageCatalogInWorkspaceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update image catalog in workspace params
func (o *UpdateImageCatalogInWorkspaceParams) WithContext(ctx context.Context) *UpdateImageCatalogInWorkspaceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update image catalog in workspace params
func (o *UpdateImageCatalogInWorkspaceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update image catalog in workspace params
func (o *UpdateImageCatalogInWorkspaceParams) WithHTTPClient(client *http.Client) *UpdateImageCatalogInWorkspaceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update image catalog in workspace params
func (o *UpdateImageCatalogInWorkspaceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update image catalog in workspace params
func (o *UpdateImageCatalogInWorkspaceParams) WithBody(body *model.UpdateImageCatalogV4Request) *UpdateImageCatalogInWorkspaceParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update image catalog in workspace params
func (o *UpdateImageCatalogInWorkspaceParams) SetBody(body *model.UpdateImageCatalogV4Request) {
	o.Body = body
}

// WithWorkspaceID adds the workspaceID to the update image catalog in workspace params
func (o *UpdateImageCatalogInWorkspaceParams) WithWorkspaceID(workspaceID int64) *UpdateImageCatalogInWorkspaceParams {
	o.SetWorkspaceID(workspaceID)
	return o
}

// SetWorkspaceID adds the workspaceId to the update image catalog in workspace params
func (o *UpdateImageCatalogInWorkspaceParams) SetWorkspaceID(workspaceID int64) {
	o.WorkspaceID = workspaceID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateImageCatalogInWorkspaceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
