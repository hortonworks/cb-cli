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

// NewUpdateLoadBalancersInternalParams creates a new UpdateLoadBalancersInternalParams object
// with the default values initialized.
func NewUpdateLoadBalancersInternalParams() *UpdateLoadBalancersInternalParams {
	var ()
	return &UpdateLoadBalancersInternalParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateLoadBalancersInternalParamsWithTimeout creates a new UpdateLoadBalancersInternalParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateLoadBalancersInternalParamsWithTimeout(timeout time.Duration) *UpdateLoadBalancersInternalParams {
	var ()
	return &UpdateLoadBalancersInternalParams{

		timeout: timeout,
	}
}

// NewUpdateLoadBalancersInternalParamsWithContext creates a new UpdateLoadBalancersInternalParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateLoadBalancersInternalParamsWithContext(ctx context.Context) *UpdateLoadBalancersInternalParams {
	var ()
	return &UpdateLoadBalancersInternalParams{

		Context: ctx,
	}
}

// NewUpdateLoadBalancersInternalParamsWithHTTPClient creates a new UpdateLoadBalancersInternalParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateLoadBalancersInternalParamsWithHTTPClient(client *http.Client) *UpdateLoadBalancersInternalParams {
	var ()
	return &UpdateLoadBalancersInternalParams{
		HTTPClient: client,
	}
}

/*
UpdateLoadBalancersInternalParams contains all the parameters to send to the API endpoint
for the update load balancers internal operation typically these are written to a http.Request
*/
type UpdateLoadBalancersInternalParams struct {

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

// WithTimeout adds the timeout to the update load balancers internal params
func (o *UpdateLoadBalancersInternalParams) WithTimeout(timeout time.Duration) *UpdateLoadBalancersInternalParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update load balancers internal params
func (o *UpdateLoadBalancersInternalParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update load balancers internal params
func (o *UpdateLoadBalancersInternalParams) WithContext(ctx context.Context) *UpdateLoadBalancersInternalParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update load balancers internal params
func (o *UpdateLoadBalancersInternalParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update load balancers internal params
func (o *UpdateLoadBalancersInternalParams) WithHTTPClient(client *http.Client) *UpdateLoadBalancersInternalParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update load balancers internal params
func (o *UpdateLoadBalancersInternalParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInitiatorUserCrn adds the initiatorUserCrn to the update load balancers internal params
func (o *UpdateLoadBalancersInternalParams) WithInitiatorUserCrn(initiatorUserCrn *string) *UpdateLoadBalancersInternalParams {
	o.SetInitiatorUserCrn(initiatorUserCrn)
	return o
}

// SetInitiatorUserCrn adds the initiatorUserCrn to the update load balancers internal params
func (o *UpdateLoadBalancersInternalParams) SetInitiatorUserCrn(initiatorUserCrn *string) {
	o.InitiatorUserCrn = initiatorUserCrn
}

// WithName adds the name to the update load balancers internal params
func (o *UpdateLoadBalancersInternalParams) WithName(name string) *UpdateLoadBalancersInternalParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the update load balancers internal params
func (o *UpdateLoadBalancersInternalParams) SetName(name string) {
	o.Name = name
}

// WithWorkspaceID adds the workspaceID to the update load balancers internal params
func (o *UpdateLoadBalancersInternalParams) WithWorkspaceID(workspaceID int64) *UpdateLoadBalancersInternalParams {
	o.SetWorkspaceID(workspaceID)
	return o
}

// SetWorkspaceID adds the workspaceId to the update load balancers internal params
func (o *UpdateLoadBalancersInternalParams) SetWorkspaceID(workspaceID int64) {
	o.WorkspaceID = workspaceID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateLoadBalancersInternalParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param workspaceId
	if err := r.SetPathParam("workspaceId", swag.FormatInt64(o.WorkspaceID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
