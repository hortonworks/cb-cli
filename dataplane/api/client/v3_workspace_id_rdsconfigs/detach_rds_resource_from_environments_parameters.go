// Code generated by go-swagger; DO NOT EDIT.

package v3_workspace_id_rdsconfigs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDetachRdsResourceFromEnvironmentsParams creates a new DetachRdsResourceFromEnvironmentsParams object
// with the default values initialized.
func NewDetachRdsResourceFromEnvironmentsParams() *DetachRdsResourceFromEnvironmentsParams {
	var ()
	return &DetachRdsResourceFromEnvironmentsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDetachRdsResourceFromEnvironmentsParamsWithTimeout creates a new DetachRdsResourceFromEnvironmentsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDetachRdsResourceFromEnvironmentsParamsWithTimeout(timeout time.Duration) *DetachRdsResourceFromEnvironmentsParams {
	var ()
	return &DetachRdsResourceFromEnvironmentsParams{

		timeout: timeout,
	}
}

// NewDetachRdsResourceFromEnvironmentsParamsWithContext creates a new DetachRdsResourceFromEnvironmentsParams object
// with the default values initialized, and the ability to set a context for a request
func NewDetachRdsResourceFromEnvironmentsParamsWithContext(ctx context.Context) *DetachRdsResourceFromEnvironmentsParams {
	var ()
	return &DetachRdsResourceFromEnvironmentsParams{

		Context: ctx,
	}
}

// NewDetachRdsResourceFromEnvironmentsParamsWithHTTPClient creates a new DetachRdsResourceFromEnvironmentsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDetachRdsResourceFromEnvironmentsParamsWithHTTPClient(client *http.Client) *DetachRdsResourceFromEnvironmentsParams {
	var ()
	return &DetachRdsResourceFromEnvironmentsParams{
		HTTPClient: client,
	}
}

/*DetachRdsResourceFromEnvironmentsParams contains all the parameters to send to the API endpoint
for the detach rds resource from environments operation typically these are written to a http.Request
*/
type DetachRdsResourceFromEnvironmentsParams struct {

	/*Body*/
	Body []string
	/*Name*/
	Name string
	/*WorkspaceID*/
	WorkspaceID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the detach rds resource from environments params
func (o *DetachRdsResourceFromEnvironmentsParams) WithTimeout(timeout time.Duration) *DetachRdsResourceFromEnvironmentsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the detach rds resource from environments params
func (o *DetachRdsResourceFromEnvironmentsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the detach rds resource from environments params
func (o *DetachRdsResourceFromEnvironmentsParams) WithContext(ctx context.Context) *DetachRdsResourceFromEnvironmentsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the detach rds resource from environments params
func (o *DetachRdsResourceFromEnvironmentsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the detach rds resource from environments params
func (o *DetachRdsResourceFromEnvironmentsParams) WithHTTPClient(client *http.Client) *DetachRdsResourceFromEnvironmentsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the detach rds resource from environments params
func (o *DetachRdsResourceFromEnvironmentsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the detach rds resource from environments params
func (o *DetachRdsResourceFromEnvironmentsParams) WithBody(body []string) *DetachRdsResourceFromEnvironmentsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the detach rds resource from environments params
func (o *DetachRdsResourceFromEnvironmentsParams) SetBody(body []string) {
	o.Body = body
}

// WithName adds the name to the detach rds resource from environments params
func (o *DetachRdsResourceFromEnvironmentsParams) WithName(name string) *DetachRdsResourceFromEnvironmentsParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the detach rds resource from environments params
func (o *DetachRdsResourceFromEnvironmentsParams) SetName(name string) {
	o.Name = name
}

// WithWorkspaceID adds the workspaceID to the detach rds resource from environments params
func (o *DetachRdsResourceFromEnvironmentsParams) WithWorkspaceID(workspaceID int64) *DetachRdsResourceFromEnvironmentsParams {
	o.SetWorkspaceID(workspaceID)
	return o
}

// SetWorkspaceID adds the workspaceId to the detach rds resource from environments params
func (o *DetachRdsResourceFromEnvironmentsParams) SetWorkspaceID(workspaceID int64) {
	o.WorkspaceID = workspaceID
}

// WriteToRequest writes these params to a swagger request
func (o *DetachRdsResourceFromEnvironmentsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
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
