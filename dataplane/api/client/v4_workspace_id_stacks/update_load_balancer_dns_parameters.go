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

// NewUpdateLoadBalancerDNSParams creates a new UpdateLoadBalancerDNSParams object
// with the default values initialized.
func NewUpdateLoadBalancerDNSParams() *UpdateLoadBalancerDNSParams {
	var ()
	return &UpdateLoadBalancerDNSParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateLoadBalancerDNSParamsWithTimeout creates a new UpdateLoadBalancerDNSParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateLoadBalancerDNSParamsWithTimeout(timeout time.Duration) *UpdateLoadBalancerDNSParams {
	var ()
	return &UpdateLoadBalancerDNSParams{

		timeout: timeout,
	}
}

// NewUpdateLoadBalancerDNSParamsWithContext creates a new UpdateLoadBalancerDNSParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateLoadBalancerDNSParamsWithContext(ctx context.Context) *UpdateLoadBalancerDNSParams {
	var ()
	return &UpdateLoadBalancerDNSParams{

		Context: ctx,
	}
}

// NewUpdateLoadBalancerDNSParamsWithHTTPClient creates a new UpdateLoadBalancerDNSParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateLoadBalancerDNSParamsWithHTTPClient(client *http.Client) *UpdateLoadBalancerDNSParams {
	var ()
	return &UpdateLoadBalancerDNSParams{
		HTTPClient: client,
	}
}

/*UpdateLoadBalancerDNSParams contains all the parameters to send to the API endpoint
for the update load balancer DNS operation typically these are written to a http.Request
*/
type UpdateLoadBalancerDNSParams struct {

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

// WithTimeout adds the timeout to the update load balancer DNS params
func (o *UpdateLoadBalancerDNSParams) WithTimeout(timeout time.Duration) *UpdateLoadBalancerDNSParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update load balancer DNS params
func (o *UpdateLoadBalancerDNSParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update load balancer DNS params
func (o *UpdateLoadBalancerDNSParams) WithContext(ctx context.Context) *UpdateLoadBalancerDNSParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update load balancer DNS params
func (o *UpdateLoadBalancerDNSParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update load balancer DNS params
func (o *UpdateLoadBalancerDNSParams) WithHTTPClient(client *http.Client) *UpdateLoadBalancerDNSParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update load balancer DNS params
func (o *UpdateLoadBalancerDNSParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInitiatorUserCrn adds the initiatorUserCrn to the update load balancer DNS params
func (o *UpdateLoadBalancerDNSParams) WithInitiatorUserCrn(initiatorUserCrn *string) *UpdateLoadBalancerDNSParams {
	o.SetInitiatorUserCrn(initiatorUserCrn)
	return o
}

// SetInitiatorUserCrn adds the initiatorUserCrn to the update load balancer DNS params
func (o *UpdateLoadBalancerDNSParams) SetInitiatorUserCrn(initiatorUserCrn *string) {
	o.InitiatorUserCrn = initiatorUserCrn
}

// WithName adds the name to the update load balancer DNS params
func (o *UpdateLoadBalancerDNSParams) WithName(name string) *UpdateLoadBalancerDNSParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the update load balancer DNS params
func (o *UpdateLoadBalancerDNSParams) SetName(name string) {
	o.Name = name
}

// WithWorkspaceID adds the workspaceID to the update load balancer DNS params
func (o *UpdateLoadBalancerDNSParams) WithWorkspaceID(workspaceID int64) *UpdateLoadBalancerDNSParams {
	o.SetWorkspaceID(workspaceID)
	return o
}

// SetWorkspaceID adds the workspaceId to the update load balancer DNS params
func (o *UpdateLoadBalancerDNSParams) SetWorkspaceID(workspaceID int64) {
	o.WorkspaceID = workspaceID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateLoadBalancerDNSParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
