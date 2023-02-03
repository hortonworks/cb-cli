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

// NewAttachRecipeInternalParams creates a new AttachRecipeInternalParams object
// with the default values initialized.
func NewAttachRecipeInternalParams() *AttachRecipeInternalParams {
	var ()
	return &AttachRecipeInternalParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAttachRecipeInternalParamsWithTimeout creates a new AttachRecipeInternalParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAttachRecipeInternalParamsWithTimeout(timeout time.Duration) *AttachRecipeInternalParams {
	var ()
	return &AttachRecipeInternalParams{

		timeout: timeout,
	}
}

// NewAttachRecipeInternalParamsWithContext creates a new AttachRecipeInternalParams object
// with the default values initialized, and the ability to set a context for a request
func NewAttachRecipeInternalParamsWithContext(ctx context.Context) *AttachRecipeInternalParams {
	var ()
	return &AttachRecipeInternalParams{

		Context: ctx,
	}
}

// NewAttachRecipeInternalParamsWithHTTPClient creates a new AttachRecipeInternalParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAttachRecipeInternalParamsWithHTTPClient(client *http.Client) *AttachRecipeInternalParams {
	var ()
	return &AttachRecipeInternalParams{
		HTTPClient: client,
	}
}

/*
AttachRecipeInternalParams contains all the parameters to send to the API endpoint
for the attach recipe internal operation typically these are written to a http.Request
*/
type AttachRecipeInternalParams struct {

	/*Body*/
	Body *model.AttachRecipeV4Request
	/*InitiatorUserCrn*/
	InitiatorUserCrn *string
	/*Name*/
	Name string
	/*WorkspaceID*/
	WorkspaceID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the attach recipe internal params
func (o *AttachRecipeInternalParams) WithTimeout(timeout time.Duration) *AttachRecipeInternalParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the attach recipe internal params
func (o *AttachRecipeInternalParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the attach recipe internal params
func (o *AttachRecipeInternalParams) WithContext(ctx context.Context) *AttachRecipeInternalParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the attach recipe internal params
func (o *AttachRecipeInternalParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the attach recipe internal params
func (o *AttachRecipeInternalParams) WithHTTPClient(client *http.Client) *AttachRecipeInternalParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the attach recipe internal params
func (o *AttachRecipeInternalParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the attach recipe internal params
func (o *AttachRecipeInternalParams) WithBody(body *model.AttachRecipeV4Request) *AttachRecipeInternalParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the attach recipe internal params
func (o *AttachRecipeInternalParams) SetBody(body *model.AttachRecipeV4Request) {
	o.Body = body
}

// WithInitiatorUserCrn adds the initiatorUserCrn to the attach recipe internal params
func (o *AttachRecipeInternalParams) WithInitiatorUserCrn(initiatorUserCrn *string) *AttachRecipeInternalParams {
	o.SetInitiatorUserCrn(initiatorUserCrn)
	return o
}

// SetInitiatorUserCrn adds the initiatorUserCrn to the attach recipe internal params
func (o *AttachRecipeInternalParams) SetInitiatorUserCrn(initiatorUserCrn *string) {
	o.InitiatorUserCrn = initiatorUserCrn
}

// WithName adds the name to the attach recipe internal params
func (o *AttachRecipeInternalParams) WithName(name string) *AttachRecipeInternalParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the attach recipe internal params
func (o *AttachRecipeInternalParams) SetName(name string) {
	o.Name = name
}

// WithWorkspaceID adds the workspaceID to the attach recipe internal params
func (o *AttachRecipeInternalParams) WithWorkspaceID(workspaceID int64) *AttachRecipeInternalParams {
	o.SetWorkspaceID(workspaceID)
	return o
}

// SetWorkspaceID adds the workspaceId to the attach recipe internal params
func (o *AttachRecipeInternalParams) SetWorkspaceID(workspaceID int64) {
	o.WorkspaceID = workspaceID
}

// WriteToRequest writes these params to a swagger request
func (o *AttachRecipeInternalParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
