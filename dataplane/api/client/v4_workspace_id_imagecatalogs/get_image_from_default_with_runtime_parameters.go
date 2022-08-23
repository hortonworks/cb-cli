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
)

// NewGetImageFromDefaultWithRuntimeParams creates a new GetImageFromDefaultWithRuntimeParams object
// with the default values initialized.
func NewGetImageFromDefaultWithRuntimeParams() *GetImageFromDefaultWithRuntimeParams {
	var ()
	return &GetImageFromDefaultWithRuntimeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetImageFromDefaultWithRuntimeParamsWithTimeout creates a new GetImageFromDefaultWithRuntimeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetImageFromDefaultWithRuntimeParamsWithTimeout(timeout time.Duration) *GetImageFromDefaultWithRuntimeParams {
	var ()
	return &GetImageFromDefaultWithRuntimeParams{

		timeout: timeout,
	}
}

// NewGetImageFromDefaultWithRuntimeParamsWithContext creates a new GetImageFromDefaultWithRuntimeParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetImageFromDefaultWithRuntimeParamsWithContext(ctx context.Context) *GetImageFromDefaultWithRuntimeParams {
	var ()
	return &GetImageFromDefaultWithRuntimeParams{

		Context: ctx,
	}
}

// NewGetImageFromDefaultWithRuntimeParamsWithHTTPClient creates a new GetImageFromDefaultWithRuntimeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetImageFromDefaultWithRuntimeParamsWithHTTPClient(client *http.Client) *GetImageFromDefaultWithRuntimeParams {
	var ()
	return &GetImageFromDefaultWithRuntimeParams{
		HTTPClient: client,
	}
}

/*GetImageFromDefaultWithRuntimeParams contains all the parameters to send to the API endpoint
for the get image from default with runtime operation typically these are written to a http.Request
*/
type GetImageFromDefaultWithRuntimeParams struct {

	/*GovCloud*/
	GovCloud bool
	/*Provider*/
	Provider string
	/*Runtime*/
	Runtime string
	/*Type*/
	Type string
	/*WorkspaceID*/
	WorkspaceID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get image from default with runtime params
func (o *GetImageFromDefaultWithRuntimeParams) WithTimeout(timeout time.Duration) *GetImageFromDefaultWithRuntimeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get image from default with runtime params
func (o *GetImageFromDefaultWithRuntimeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get image from default with runtime params
func (o *GetImageFromDefaultWithRuntimeParams) WithContext(ctx context.Context) *GetImageFromDefaultWithRuntimeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get image from default with runtime params
func (o *GetImageFromDefaultWithRuntimeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get image from default with runtime params
func (o *GetImageFromDefaultWithRuntimeParams) WithHTTPClient(client *http.Client) *GetImageFromDefaultWithRuntimeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get image from default with runtime params
func (o *GetImageFromDefaultWithRuntimeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGovCloud adds the govCloud to the get image from default with runtime params
func (o *GetImageFromDefaultWithRuntimeParams) WithGovCloud(govCloud bool) *GetImageFromDefaultWithRuntimeParams {
	o.SetGovCloud(govCloud)
	return o
}

// SetGovCloud adds the govCloud to the get image from default with runtime params
func (o *GetImageFromDefaultWithRuntimeParams) SetGovCloud(govCloud bool) {
	o.GovCloud = govCloud
}

// WithProvider adds the provider to the get image from default with runtime params
func (o *GetImageFromDefaultWithRuntimeParams) WithProvider(provider string) *GetImageFromDefaultWithRuntimeParams {
	o.SetProvider(provider)
	return o
}

// SetProvider adds the provider to the get image from default with runtime params
func (o *GetImageFromDefaultWithRuntimeParams) SetProvider(provider string) {
	o.Provider = provider
}

// WithRuntime adds the runtime to the get image from default with runtime params
func (o *GetImageFromDefaultWithRuntimeParams) WithRuntime(runtime string) *GetImageFromDefaultWithRuntimeParams {
	o.SetRuntime(runtime)
	return o
}

// SetRuntime adds the runtime to the get image from default with runtime params
func (o *GetImageFromDefaultWithRuntimeParams) SetRuntime(runtime string) {
	o.Runtime = runtime
}

// WithType adds the typeVar to the get image from default with runtime params
func (o *GetImageFromDefaultWithRuntimeParams) WithType(typeVar string) *GetImageFromDefaultWithRuntimeParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the get image from default with runtime params
func (o *GetImageFromDefaultWithRuntimeParams) SetType(typeVar string) {
	o.Type = typeVar
}

// WithWorkspaceID adds the workspaceID to the get image from default with runtime params
func (o *GetImageFromDefaultWithRuntimeParams) WithWorkspaceID(workspaceID int64) *GetImageFromDefaultWithRuntimeParams {
	o.SetWorkspaceID(workspaceID)
	return o
}

// SetWorkspaceID adds the workspaceId to the get image from default with runtime params
func (o *GetImageFromDefaultWithRuntimeParams) SetWorkspaceID(workspaceID int64) {
	o.WorkspaceID = workspaceID
}

// WriteToRequest writes these params to a swagger request
func (o *GetImageFromDefaultWithRuntimeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param govCloud
	if err := r.SetPathParam("govCloud", swag.FormatBool(o.GovCloud)); err != nil {
		return err
	}

	// path param provider
	if err := r.SetPathParam("provider", o.Provider); err != nil {
		return err
	}

	// path param runtime
	if err := r.SetPathParam("runtime", o.Runtime); err != nil {
		return err
	}

	// path param type
	if err := r.SetPathParam("type", o.Type); err != nil {
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
