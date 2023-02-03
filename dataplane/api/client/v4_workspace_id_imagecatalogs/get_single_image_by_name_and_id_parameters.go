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

// NewGetSingleImageByNameAndIDParams creates a new GetSingleImageByNameAndIDParams object
// with the default values initialized.
func NewGetSingleImageByNameAndIDParams() *GetSingleImageByNameAndIDParams {
	var ()
	return &GetSingleImageByNameAndIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSingleImageByNameAndIDParamsWithTimeout creates a new GetSingleImageByNameAndIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSingleImageByNameAndIDParamsWithTimeout(timeout time.Duration) *GetSingleImageByNameAndIDParams {
	var ()
	return &GetSingleImageByNameAndIDParams{

		timeout: timeout,
	}
}

// NewGetSingleImageByNameAndIDParamsWithContext creates a new GetSingleImageByNameAndIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSingleImageByNameAndIDParamsWithContext(ctx context.Context) *GetSingleImageByNameAndIDParams {
	var ()
	return &GetSingleImageByNameAndIDParams{

		Context: ctx,
	}
}

// NewGetSingleImageByNameAndIDParamsWithHTTPClient creates a new GetSingleImageByNameAndIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSingleImageByNameAndIDParamsWithHTTPClient(client *http.Client) *GetSingleImageByNameAndIDParams {
	var ()
	return &GetSingleImageByNameAndIDParams{
		HTTPClient: client,
	}
}

/*
GetSingleImageByNameAndIDParams contains all the parameters to send to the API endpoint
for the get single image by name and Id operation typically these are written to a http.Request
*/
type GetSingleImageByNameAndIDParams struct {

	/*ImageID*/
	ImageID *string
	/*Name*/
	Name string
	/*WorkspaceID*/
	WorkspaceID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get single image by name and Id params
func (o *GetSingleImageByNameAndIDParams) WithTimeout(timeout time.Duration) *GetSingleImageByNameAndIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get single image by name and Id params
func (o *GetSingleImageByNameAndIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get single image by name and Id params
func (o *GetSingleImageByNameAndIDParams) WithContext(ctx context.Context) *GetSingleImageByNameAndIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get single image by name and Id params
func (o *GetSingleImageByNameAndIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get single image by name and Id params
func (o *GetSingleImageByNameAndIDParams) WithHTTPClient(client *http.Client) *GetSingleImageByNameAndIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get single image by name and Id params
func (o *GetSingleImageByNameAndIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithImageID adds the imageID to the get single image by name and Id params
func (o *GetSingleImageByNameAndIDParams) WithImageID(imageID *string) *GetSingleImageByNameAndIDParams {
	o.SetImageID(imageID)
	return o
}

// SetImageID adds the imageId to the get single image by name and Id params
func (o *GetSingleImageByNameAndIDParams) SetImageID(imageID *string) {
	o.ImageID = imageID
}

// WithName adds the name to the get single image by name and Id params
func (o *GetSingleImageByNameAndIDParams) WithName(name string) *GetSingleImageByNameAndIDParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get single image by name and Id params
func (o *GetSingleImageByNameAndIDParams) SetName(name string) {
	o.Name = name
}

// WithWorkspaceID adds the workspaceID to the get single image by name and Id params
func (o *GetSingleImageByNameAndIDParams) WithWorkspaceID(workspaceID int64) *GetSingleImageByNameAndIDParams {
	o.SetWorkspaceID(workspaceID)
	return o
}

// SetWorkspaceID adds the workspaceId to the get single image by name and Id params
func (o *GetSingleImageByNameAndIDParams) SetWorkspaceID(workspaceID int64) {
	o.WorkspaceID = workspaceID
}

// WriteToRequest writes these params to a swagger request
func (o *GetSingleImageByNameAndIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ImageID != nil {

		// query param imageId
		var qrImageID string
		if o.ImageID != nil {
			qrImageID = *o.ImageID
		}
		qImageID := qrImageID
		if qImageID != "" {
			if err := r.SetQueryParam("imageId", qImageID); err != nil {
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
