// Code generated by go-swagger; DO NOT EDIT.

package v4_workspace_id_stacks

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

// NewUpdatePillarConfigurationByNameParams creates a new UpdatePillarConfigurationByNameParams object
// with the default values initialized.
func NewUpdatePillarConfigurationByNameParams() *UpdatePillarConfigurationByNameParams {
	var ()
	return &UpdatePillarConfigurationByNameParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdatePillarConfigurationByNameParamsWithTimeout creates a new UpdatePillarConfigurationByNameParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdatePillarConfigurationByNameParamsWithTimeout(timeout time.Duration) *UpdatePillarConfigurationByNameParams {
	var ()
	return &UpdatePillarConfigurationByNameParams{

		timeout: timeout,
	}
}

// NewUpdatePillarConfigurationByNameParamsWithContext creates a new UpdatePillarConfigurationByNameParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdatePillarConfigurationByNameParamsWithContext(ctx context.Context) *UpdatePillarConfigurationByNameParams {
	var ()
	return &UpdatePillarConfigurationByNameParams{

		Context: ctx,
	}
}

// NewUpdatePillarConfigurationByNameParamsWithHTTPClient creates a new UpdatePillarConfigurationByNameParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdatePillarConfigurationByNameParamsWithHTTPClient(client *http.Client) *UpdatePillarConfigurationByNameParams {
	var ()
	return &UpdatePillarConfigurationByNameParams{
		HTTPClient: client,
	}
}

/*
UpdatePillarConfigurationByNameParams contains all the parameters to send to the API endpoint
for the update pillar configuration by name operation typically these are written to a http.Request
*/
type UpdatePillarConfigurationByNameParams struct {

	/*Name*/
	Name string
	/*WorkspaceID*/
	WorkspaceID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update pillar configuration by name params
func (o *UpdatePillarConfigurationByNameParams) WithTimeout(timeout time.Duration) *UpdatePillarConfigurationByNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update pillar configuration by name params
func (o *UpdatePillarConfigurationByNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update pillar configuration by name params
func (o *UpdatePillarConfigurationByNameParams) WithContext(ctx context.Context) *UpdatePillarConfigurationByNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update pillar configuration by name params
func (o *UpdatePillarConfigurationByNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update pillar configuration by name params
func (o *UpdatePillarConfigurationByNameParams) WithHTTPClient(client *http.Client) *UpdatePillarConfigurationByNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update pillar configuration by name params
func (o *UpdatePillarConfigurationByNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the update pillar configuration by name params
func (o *UpdatePillarConfigurationByNameParams) WithName(name string) *UpdatePillarConfigurationByNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the update pillar configuration by name params
func (o *UpdatePillarConfigurationByNameParams) SetName(name string) {
	o.Name = name
}

// WithWorkspaceID adds the workspaceID to the update pillar configuration by name params
func (o *UpdatePillarConfigurationByNameParams) WithWorkspaceID(workspaceID int64) *UpdatePillarConfigurationByNameParams {
	o.SetWorkspaceID(workspaceID)
	return o
}

// SetWorkspaceID adds the workspaceId to the update pillar configuration by name params
func (o *UpdatePillarConfigurationByNameParams) SetWorkspaceID(workspaceID int64) {
	o.WorkspaceID = workspaceID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdatePillarConfigurationByNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
