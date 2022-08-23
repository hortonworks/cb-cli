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

// NewPrepareClusterUpgradeByCrnInternalParams creates a new PrepareClusterUpgradeByCrnInternalParams object
// with the default values initialized.
func NewPrepareClusterUpgradeByCrnInternalParams() *PrepareClusterUpgradeByCrnInternalParams {
	var ()
	return &PrepareClusterUpgradeByCrnInternalParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPrepareClusterUpgradeByCrnInternalParamsWithTimeout creates a new PrepareClusterUpgradeByCrnInternalParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPrepareClusterUpgradeByCrnInternalParamsWithTimeout(timeout time.Duration) *PrepareClusterUpgradeByCrnInternalParams {
	var ()
	return &PrepareClusterUpgradeByCrnInternalParams{

		timeout: timeout,
	}
}

// NewPrepareClusterUpgradeByCrnInternalParamsWithContext creates a new PrepareClusterUpgradeByCrnInternalParams object
// with the default values initialized, and the ability to set a context for a request
func NewPrepareClusterUpgradeByCrnInternalParamsWithContext(ctx context.Context) *PrepareClusterUpgradeByCrnInternalParams {
	var ()
	return &PrepareClusterUpgradeByCrnInternalParams{

		Context: ctx,
	}
}

// NewPrepareClusterUpgradeByCrnInternalParamsWithHTTPClient creates a new PrepareClusterUpgradeByCrnInternalParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPrepareClusterUpgradeByCrnInternalParamsWithHTTPClient(client *http.Client) *PrepareClusterUpgradeByCrnInternalParams {
	var ()
	return &PrepareClusterUpgradeByCrnInternalParams{
		HTTPClient: client,
	}
}

/*PrepareClusterUpgradeByCrnInternalParams contains all the parameters to send to the API endpoint
for the prepare cluster upgrade by crn internal operation typically these are written to a http.Request
*/
type PrepareClusterUpgradeByCrnInternalParams struct {

	/*Crn*/
	Crn string
	/*ImageID*/
	ImageID *string
	/*InitiatorUserCrn*/
	InitiatorUserCrn *string
	/*WorkspaceID*/
	WorkspaceID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the prepare cluster upgrade by crn internal params
func (o *PrepareClusterUpgradeByCrnInternalParams) WithTimeout(timeout time.Duration) *PrepareClusterUpgradeByCrnInternalParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the prepare cluster upgrade by crn internal params
func (o *PrepareClusterUpgradeByCrnInternalParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the prepare cluster upgrade by crn internal params
func (o *PrepareClusterUpgradeByCrnInternalParams) WithContext(ctx context.Context) *PrepareClusterUpgradeByCrnInternalParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the prepare cluster upgrade by crn internal params
func (o *PrepareClusterUpgradeByCrnInternalParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the prepare cluster upgrade by crn internal params
func (o *PrepareClusterUpgradeByCrnInternalParams) WithHTTPClient(client *http.Client) *PrepareClusterUpgradeByCrnInternalParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the prepare cluster upgrade by crn internal params
func (o *PrepareClusterUpgradeByCrnInternalParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCrn adds the crn to the prepare cluster upgrade by crn internal params
func (o *PrepareClusterUpgradeByCrnInternalParams) WithCrn(crn string) *PrepareClusterUpgradeByCrnInternalParams {
	o.SetCrn(crn)
	return o
}

// SetCrn adds the crn to the prepare cluster upgrade by crn internal params
func (o *PrepareClusterUpgradeByCrnInternalParams) SetCrn(crn string) {
	o.Crn = crn
}

// WithImageID adds the imageID to the prepare cluster upgrade by crn internal params
func (o *PrepareClusterUpgradeByCrnInternalParams) WithImageID(imageID *string) *PrepareClusterUpgradeByCrnInternalParams {
	o.SetImageID(imageID)
	return o
}

// SetImageID adds the imageId to the prepare cluster upgrade by crn internal params
func (o *PrepareClusterUpgradeByCrnInternalParams) SetImageID(imageID *string) {
	o.ImageID = imageID
}

// WithInitiatorUserCrn adds the initiatorUserCrn to the prepare cluster upgrade by crn internal params
func (o *PrepareClusterUpgradeByCrnInternalParams) WithInitiatorUserCrn(initiatorUserCrn *string) *PrepareClusterUpgradeByCrnInternalParams {
	o.SetInitiatorUserCrn(initiatorUserCrn)
	return o
}

// SetInitiatorUserCrn adds the initiatorUserCrn to the prepare cluster upgrade by crn internal params
func (o *PrepareClusterUpgradeByCrnInternalParams) SetInitiatorUserCrn(initiatorUserCrn *string) {
	o.InitiatorUserCrn = initiatorUserCrn
}

// WithWorkspaceID adds the workspaceID to the prepare cluster upgrade by crn internal params
func (o *PrepareClusterUpgradeByCrnInternalParams) WithWorkspaceID(workspaceID int64) *PrepareClusterUpgradeByCrnInternalParams {
	o.SetWorkspaceID(workspaceID)
	return o
}

// SetWorkspaceID adds the workspaceId to the prepare cluster upgrade by crn internal params
func (o *PrepareClusterUpgradeByCrnInternalParams) SetWorkspaceID(workspaceID int64) {
	o.WorkspaceID = workspaceID
}

// WriteToRequest writes these params to a swagger request
func (o *PrepareClusterUpgradeByCrnInternalParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param crn
	if err := r.SetPathParam("crn", o.Crn); err != nil {
		return err
	}

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

	// path param workspaceId
	if err := r.SetPathParam("workspaceId", swag.FormatInt64(o.WorkspaceID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
