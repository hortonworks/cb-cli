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

	model "github.com/hortonworks/cb-cli/dataplane/api/model"
)

// NewRotateSaltPasswordForStackInWorkspaceV4InternalParams creates a new RotateSaltPasswordForStackInWorkspaceV4InternalParams object
// with the default values initialized.
func NewRotateSaltPasswordForStackInWorkspaceV4InternalParams() *RotateSaltPasswordForStackInWorkspaceV4InternalParams {
	var ()
	return &RotateSaltPasswordForStackInWorkspaceV4InternalParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRotateSaltPasswordForStackInWorkspaceV4InternalParamsWithTimeout creates a new RotateSaltPasswordForStackInWorkspaceV4InternalParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRotateSaltPasswordForStackInWorkspaceV4InternalParamsWithTimeout(timeout time.Duration) *RotateSaltPasswordForStackInWorkspaceV4InternalParams {
	var ()
	return &RotateSaltPasswordForStackInWorkspaceV4InternalParams{

		timeout: timeout,
	}
}

// NewRotateSaltPasswordForStackInWorkspaceV4InternalParamsWithContext creates a new RotateSaltPasswordForStackInWorkspaceV4InternalParams object
// with the default values initialized, and the ability to set a context for a request
func NewRotateSaltPasswordForStackInWorkspaceV4InternalParamsWithContext(ctx context.Context) *RotateSaltPasswordForStackInWorkspaceV4InternalParams {
	var ()
	return &RotateSaltPasswordForStackInWorkspaceV4InternalParams{

		Context: ctx,
	}
}

// NewRotateSaltPasswordForStackInWorkspaceV4InternalParamsWithHTTPClient creates a new RotateSaltPasswordForStackInWorkspaceV4InternalParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRotateSaltPasswordForStackInWorkspaceV4InternalParamsWithHTTPClient(client *http.Client) *RotateSaltPasswordForStackInWorkspaceV4InternalParams {
	var ()
	return &RotateSaltPasswordForStackInWorkspaceV4InternalParams{
		HTTPClient: client,
	}
}

/*
RotateSaltPasswordForStackInWorkspaceV4InternalParams contains all the parameters to send to the API endpoint
for the rotate salt password for stack in workspace v4 internal operation typically these are written to a http.Request
*/
type RotateSaltPasswordForStackInWorkspaceV4InternalParams struct {

	/*Body*/
	Body *model.RotateSaltPasswordRequest
	/*Crn*/
	Crn string
	/*InitiatorUserCrn*/
	InitiatorUserCrn *string
	/*WorkspaceID*/
	WorkspaceID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the rotate salt password for stack in workspace v4 internal params
func (o *RotateSaltPasswordForStackInWorkspaceV4InternalParams) WithTimeout(timeout time.Duration) *RotateSaltPasswordForStackInWorkspaceV4InternalParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the rotate salt password for stack in workspace v4 internal params
func (o *RotateSaltPasswordForStackInWorkspaceV4InternalParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the rotate salt password for stack in workspace v4 internal params
func (o *RotateSaltPasswordForStackInWorkspaceV4InternalParams) WithContext(ctx context.Context) *RotateSaltPasswordForStackInWorkspaceV4InternalParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the rotate salt password for stack in workspace v4 internal params
func (o *RotateSaltPasswordForStackInWorkspaceV4InternalParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the rotate salt password for stack in workspace v4 internal params
func (o *RotateSaltPasswordForStackInWorkspaceV4InternalParams) WithHTTPClient(client *http.Client) *RotateSaltPasswordForStackInWorkspaceV4InternalParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the rotate salt password for stack in workspace v4 internal params
func (o *RotateSaltPasswordForStackInWorkspaceV4InternalParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the rotate salt password for stack in workspace v4 internal params
func (o *RotateSaltPasswordForStackInWorkspaceV4InternalParams) WithBody(body *model.RotateSaltPasswordRequest) *RotateSaltPasswordForStackInWorkspaceV4InternalParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the rotate salt password for stack in workspace v4 internal params
func (o *RotateSaltPasswordForStackInWorkspaceV4InternalParams) SetBody(body *model.RotateSaltPasswordRequest) {
	o.Body = body
}

// WithCrn adds the crn to the rotate salt password for stack in workspace v4 internal params
func (o *RotateSaltPasswordForStackInWorkspaceV4InternalParams) WithCrn(crn string) *RotateSaltPasswordForStackInWorkspaceV4InternalParams {
	o.SetCrn(crn)
	return o
}

// SetCrn adds the crn to the rotate salt password for stack in workspace v4 internal params
func (o *RotateSaltPasswordForStackInWorkspaceV4InternalParams) SetCrn(crn string) {
	o.Crn = crn
}

// WithInitiatorUserCrn adds the initiatorUserCrn to the rotate salt password for stack in workspace v4 internal params
func (o *RotateSaltPasswordForStackInWorkspaceV4InternalParams) WithInitiatorUserCrn(initiatorUserCrn *string) *RotateSaltPasswordForStackInWorkspaceV4InternalParams {
	o.SetInitiatorUserCrn(initiatorUserCrn)
	return o
}

// SetInitiatorUserCrn adds the initiatorUserCrn to the rotate salt password for stack in workspace v4 internal params
func (o *RotateSaltPasswordForStackInWorkspaceV4InternalParams) SetInitiatorUserCrn(initiatorUserCrn *string) {
	o.InitiatorUserCrn = initiatorUserCrn
}

// WithWorkspaceID adds the workspaceID to the rotate salt password for stack in workspace v4 internal params
func (o *RotateSaltPasswordForStackInWorkspaceV4InternalParams) WithWorkspaceID(workspaceID int64) *RotateSaltPasswordForStackInWorkspaceV4InternalParams {
	o.SetWorkspaceID(workspaceID)
	return o
}

// SetWorkspaceID adds the workspaceId to the rotate salt password for stack in workspace v4 internal params
func (o *RotateSaltPasswordForStackInWorkspaceV4InternalParams) SetWorkspaceID(workspaceID int64) {
	o.WorkspaceID = workspaceID
}

// WriteToRequest writes these params to a swagger request
func (o *RotateSaltPasswordForStackInWorkspaceV4InternalParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param crn
	if err := r.SetPathParam("crn", o.Crn); err != nil {
		return err
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
