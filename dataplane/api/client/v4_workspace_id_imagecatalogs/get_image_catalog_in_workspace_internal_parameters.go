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

// NewGetImageCatalogInWorkspaceInternalParams creates a new GetImageCatalogInWorkspaceInternalParams object
// with the default values initialized.
func NewGetImageCatalogInWorkspaceInternalParams() *GetImageCatalogInWorkspaceInternalParams {
	var (
		withImagesDefault = bool(false)
	)
	return &GetImageCatalogInWorkspaceInternalParams{
		WithImages: &withImagesDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetImageCatalogInWorkspaceInternalParamsWithTimeout creates a new GetImageCatalogInWorkspaceInternalParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetImageCatalogInWorkspaceInternalParamsWithTimeout(timeout time.Duration) *GetImageCatalogInWorkspaceInternalParams {
	var (
		withImagesDefault = bool(false)
	)
	return &GetImageCatalogInWorkspaceInternalParams{
		WithImages: &withImagesDefault,

		timeout: timeout,
	}
}

// NewGetImageCatalogInWorkspaceInternalParamsWithContext creates a new GetImageCatalogInWorkspaceInternalParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetImageCatalogInWorkspaceInternalParamsWithContext(ctx context.Context) *GetImageCatalogInWorkspaceInternalParams {
	var (
		withImagesDefault = bool(false)
	)
	return &GetImageCatalogInWorkspaceInternalParams{
		WithImages: &withImagesDefault,

		Context: ctx,
	}
}

// NewGetImageCatalogInWorkspaceInternalParamsWithHTTPClient creates a new GetImageCatalogInWorkspaceInternalParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetImageCatalogInWorkspaceInternalParamsWithHTTPClient(client *http.Client) *GetImageCatalogInWorkspaceInternalParams {
	var (
		withImagesDefault = bool(false)
	)
	return &GetImageCatalogInWorkspaceInternalParams{
		WithImages: &withImagesDefault,
		HTTPClient: client,
	}
}

/*
GetImageCatalogInWorkspaceInternalParams contains all the parameters to send to the API endpoint
for the get image catalog in workspace internal operation typically these are written to a http.Request
*/
type GetImageCatalogInWorkspaceInternalParams struct {

	/*InitiatorUserCrn*/
	InitiatorUserCrn *string
	/*Name*/
	Name string
	/*WithImages*/
	WithImages *bool
	/*WorkspaceID*/
	WorkspaceID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get image catalog in workspace internal params
func (o *GetImageCatalogInWorkspaceInternalParams) WithTimeout(timeout time.Duration) *GetImageCatalogInWorkspaceInternalParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get image catalog in workspace internal params
func (o *GetImageCatalogInWorkspaceInternalParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get image catalog in workspace internal params
func (o *GetImageCatalogInWorkspaceInternalParams) WithContext(ctx context.Context) *GetImageCatalogInWorkspaceInternalParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get image catalog in workspace internal params
func (o *GetImageCatalogInWorkspaceInternalParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get image catalog in workspace internal params
func (o *GetImageCatalogInWorkspaceInternalParams) WithHTTPClient(client *http.Client) *GetImageCatalogInWorkspaceInternalParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get image catalog in workspace internal params
func (o *GetImageCatalogInWorkspaceInternalParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInitiatorUserCrn adds the initiatorUserCrn to the get image catalog in workspace internal params
func (o *GetImageCatalogInWorkspaceInternalParams) WithInitiatorUserCrn(initiatorUserCrn *string) *GetImageCatalogInWorkspaceInternalParams {
	o.SetInitiatorUserCrn(initiatorUserCrn)
	return o
}

// SetInitiatorUserCrn adds the initiatorUserCrn to the get image catalog in workspace internal params
func (o *GetImageCatalogInWorkspaceInternalParams) SetInitiatorUserCrn(initiatorUserCrn *string) {
	o.InitiatorUserCrn = initiatorUserCrn
}

// WithName adds the name to the get image catalog in workspace internal params
func (o *GetImageCatalogInWorkspaceInternalParams) WithName(name string) *GetImageCatalogInWorkspaceInternalParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get image catalog in workspace internal params
func (o *GetImageCatalogInWorkspaceInternalParams) SetName(name string) {
	o.Name = name
}

// WithWithImages adds the withImages to the get image catalog in workspace internal params
func (o *GetImageCatalogInWorkspaceInternalParams) WithWithImages(withImages *bool) *GetImageCatalogInWorkspaceInternalParams {
	o.SetWithImages(withImages)
	return o
}

// SetWithImages adds the withImages to the get image catalog in workspace internal params
func (o *GetImageCatalogInWorkspaceInternalParams) SetWithImages(withImages *bool) {
	o.WithImages = withImages
}

// WithWorkspaceID adds the workspaceID to the get image catalog in workspace internal params
func (o *GetImageCatalogInWorkspaceInternalParams) WithWorkspaceID(workspaceID int64) *GetImageCatalogInWorkspaceInternalParams {
	o.SetWorkspaceID(workspaceID)
	return o
}

// SetWorkspaceID adds the workspaceId to the get image catalog in workspace internal params
func (o *GetImageCatalogInWorkspaceInternalParams) SetWorkspaceID(workspaceID int64) {
	o.WorkspaceID = workspaceID
}

// WriteToRequest writes these params to a swagger request
func (o *GetImageCatalogInWorkspaceInternalParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.InitiatorUserCrn != nil {

		// query param initiatorUserCrn
		var qrInitiatorUserCrn string
		if o.InitiatorUserCrn != nil {
			qrInitiatorUserCrn = *o.InitiatorUserCrn
		}
		qInitiatorUserCrn := qrInitiatorUserCrn
		if qInitiatorUserCrn != "" {
			if err := r.SetQueryParam("initiatorUserCrn", qInitiatorUserCrn); err != nil {
				return err
			}
		}

	}

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if o.WithImages != nil {

		// query param withImages
		var qrWithImages bool
		if o.WithImages != nil {
			qrWithImages = *o.WithImages
		}
		qWithImages := swag.FormatBool(qrWithImages)
		if qWithImages != "" {
			if err := r.SetQueryParam("withImages", qWithImages); err != nil {
				return err
			}
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
