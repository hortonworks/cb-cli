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

// NewPostStackInWorkspaceV4InternalParams creates a new PostStackInWorkspaceV4InternalParams object
// with the default values initialized.
func NewPostStackInWorkspaceV4InternalParams() *PostStackInWorkspaceV4InternalParams {
	var ()
	return &PostStackInWorkspaceV4InternalParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostStackInWorkspaceV4InternalParamsWithTimeout creates a new PostStackInWorkspaceV4InternalParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostStackInWorkspaceV4InternalParamsWithTimeout(timeout time.Duration) *PostStackInWorkspaceV4InternalParams {
	var ()
	return &PostStackInWorkspaceV4InternalParams{

		timeout: timeout,
	}
}

// NewPostStackInWorkspaceV4InternalParamsWithContext creates a new PostStackInWorkspaceV4InternalParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostStackInWorkspaceV4InternalParamsWithContext(ctx context.Context) *PostStackInWorkspaceV4InternalParams {
	var ()
	return &PostStackInWorkspaceV4InternalParams{

		Context: ctx,
	}
}

// NewPostStackInWorkspaceV4InternalParamsWithHTTPClient creates a new PostStackInWorkspaceV4InternalParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostStackInWorkspaceV4InternalParamsWithHTTPClient(client *http.Client) *PostStackInWorkspaceV4InternalParams {
	var ()
	return &PostStackInWorkspaceV4InternalParams{
		HTTPClient: client,
	}
}

/*
PostStackInWorkspaceV4InternalParams contains all the parameters to send to the API endpoint
for the post stack in workspace v4 internal operation typically these are written to a http.Request
*/
type PostStackInWorkspaceV4InternalParams struct {

	/*Body*/
	Body *model.StackV4Request
	/*InitiatorUserCrn*/
	InitiatorUserCrn *string
	/*WorkspaceID*/
	WorkspaceID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post stack in workspace v4 internal params
func (o *PostStackInWorkspaceV4InternalParams) WithTimeout(timeout time.Duration) *PostStackInWorkspaceV4InternalParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post stack in workspace v4 internal params
func (o *PostStackInWorkspaceV4InternalParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post stack in workspace v4 internal params
func (o *PostStackInWorkspaceV4InternalParams) WithContext(ctx context.Context) *PostStackInWorkspaceV4InternalParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post stack in workspace v4 internal params
func (o *PostStackInWorkspaceV4InternalParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post stack in workspace v4 internal params
func (o *PostStackInWorkspaceV4InternalParams) WithHTTPClient(client *http.Client) *PostStackInWorkspaceV4InternalParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post stack in workspace v4 internal params
func (o *PostStackInWorkspaceV4InternalParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post stack in workspace v4 internal params
func (o *PostStackInWorkspaceV4InternalParams) WithBody(body *model.StackV4Request) *PostStackInWorkspaceV4InternalParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post stack in workspace v4 internal params
func (o *PostStackInWorkspaceV4InternalParams) SetBody(body *model.StackV4Request) {
	o.Body = body
}

// WithInitiatorUserCrn adds the initiatorUserCrn to the post stack in workspace v4 internal params
func (o *PostStackInWorkspaceV4InternalParams) WithInitiatorUserCrn(initiatorUserCrn *string) *PostStackInWorkspaceV4InternalParams {
	o.SetInitiatorUserCrn(initiatorUserCrn)
	return o
}

// SetInitiatorUserCrn adds the initiatorUserCrn to the post stack in workspace v4 internal params
func (o *PostStackInWorkspaceV4InternalParams) SetInitiatorUserCrn(initiatorUserCrn *string) {
	o.InitiatorUserCrn = initiatorUserCrn
}

// WithWorkspaceID adds the workspaceID to the post stack in workspace v4 internal params
func (o *PostStackInWorkspaceV4InternalParams) WithWorkspaceID(workspaceID int64) *PostStackInWorkspaceV4InternalParams {
	o.SetWorkspaceID(workspaceID)
	return o
}

// SetWorkspaceID adds the workspaceId to the post stack in workspace v4 internal params
func (o *PostStackInWorkspaceV4InternalParams) SetWorkspaceID(workspaceID int64) {
	o.WorkspaceID = workspaceID
}

// WriteToRequest writes these params to a swagger request
func (o *PostStackInWorkspaceV4InternalParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
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
