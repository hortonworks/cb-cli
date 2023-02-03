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

// NewCheckUpgradeRdsByNameInternalParams creates a new CheckUpgradeRdsByNameInternalParams object
// with the default values initialized.
func NewCheckUpgradeRdsByNameInternalParams() *CheckUpgradeRdsByNameInternalParams {
	var ()
	return &CheckUpgradeRdsByNameInternalParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCheckUpgradeRdsByNameInternalParamsWithTimeout creates a new CheckUpgradeRdsByNameInternalParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCheckUpgradeRdsByNameInternalParamsWithTimeout(timeout time.Duration) *CheckUpgradeRdsByNameInternalParams {
	var ()
	return &CheckUpgradeRdsByNameInternalParams{

		timeout: timeout,
	}
}

// NewCheckUpgradeRdsByNameInternalParamsWithContext creates a new CheckUpgradeRdsByNameInternalParams object
// with the default values initialized, and the ability to set a context for a request
func NewCheckUpgradeRdsByNameInternalParamsWithContext(ctx context.Context) *CheckUpgradeRdsByNameInternalParams {
	var ()
	return &CheckUpgradeRdsByNameInternalParams{

		Context: ctx,
	}
}

// NewCheckUpgradeRdsByNameInternalParamsWithHTTPClient creates a new CheckUpgradeRdsByNameInternalParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCheckUpgradeRdsByNameInternalParamsWithHTTPClient(client *http.Client) *CheckUpgradeRdsByNameInternalParams {
	var ()
	return &CheckUpgradeRdsByNameInternalParams{
		HTTPClient: client,
	}
}

/*
CheckUpgradeRdsByNameInternalParams contains all the parameters to send to the API endpoint
for the check upgrade rds by name internal operation typically these are written to a http.Request
*/
type CheckUpgradeRdsByNameInternalParams struct {

	/*InitiatorUserCrn*/
	InitiatorUserCrn *string
	/*Name*/
	Name string
	/*TargetVersion*/
	TargetVersion *string
	/*WorkspaceID*/
	WorkspaceID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the check upgrade rds by name internal params
func (o *CheckUpgradeRdsByNameInternalParams) WithTimeout(timeout time.Duration) *CheckUpgradeRdsByNameInternalParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the check upgrade rds by name internal params
func (o *CheckUpgradeRdsByNameInternalParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the check upgrade rds by name internal params
func (o *CheckUpgradeRdsByNameInternalParams) WithContext(ctx context.Context) *CheckUpgradeRdsByNameInternalParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the check upgrade rds by name internal params
func (o *CheckUpgradeRdsByNameInternalParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the check upgrade rds by name internal params
func (o *CheckUpgradeRdsByNameInternalParams) WithHTTPClient(client *http.Client) *CheckUpgradeRdsByNameInternalParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the check upgrade rds by name internal params
func (o *CheckUpgradeRdsByNameInternalParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInitiatorUserCrn adds the initiatorUserCrn to the check upgrade rds by name internal params
func (o *CheckUpgradeRdsByNameInternalParams) WithInitiatorUserCrn(initiatorUserCrn *string) *CheckUpgradeRdsByNameInternalParams {
	o.SetInitiatorUserCrn(initiatorUserCrn)
	return o
}

// SetInitiatorUserCrn adds the initiatorUserCrn to the check upgrade rds by name internal params
func (o *CheckUpgradeRdsByNameInternalParams) SetInitiatorUserCrn(initiatorUserCrn *string) {
	o.InitiatorUserCrn = initiatorUserCrn
}

// WithName adds the name to the check upgrade rds by name internal params
func (o *CheckUpgradeRdsByNameInternalParams) WithName(name string) *CheckUpgradeRdsByNameInternalParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the check upgrade rds by name internal params
func (o *CheckUpgradeRdsByNameInternalParams) SetName(name string) {
	o.Name = name
}

// WithTargetVersion adds the targetVersion to the check upgrade rds by name internal params
func (o *CheckUpgradeRdsByNameInternalParams) WithTargetVersion(targetVersion *string) *CheckUpgradeRdsByNameInternalParams {
	o.SetTargetVersion(targetVersion)
	return o
}

// SetTargetVersion adds the targetVersion to the check upgrade rds by name internal params
func (o *CheckUpgradeRdsByNameInternalParams) SetTargetVersion(targetVersion *string) {
	o.TargetVersion = targetVersion
}

// WithWorkspaceID adds the workspaceID to the check upgrade rds by name internal params
func (o *CheckUpgradeRdsByNameInternalParams) WithWorkspaceID(workspaceID int64) *CheckUpgradeRdsByNameInternalParams {
	o.SetWorkspaceID(workspaceID)
	return o
}

// SetWorkspaceID adds the workspaceId to the check upgrade rds by name internal params
func (o *CheckUpgradeRdsByNameInternalParams) SetWorkspaceID(workspaceID int64) {
	o.WorkspaceID = workspaceID
}

// WriteToRequest writes these params to a swagger request
func (o *CheckUpgradeRdsByNameInternalParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.TargetVersion != nil {

		// query param targetVersion
		var qrTargetVersion string
		if o.TargetVersion != nil {
			qrTargetVersion = *o.TargetVersion
		}
		qTargetVersion := qrTargetVersion
		if qTargetVersion != "" {
			if err := r.SetQueryParam("targetVersion", qTargetVersion); err != nil {
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
