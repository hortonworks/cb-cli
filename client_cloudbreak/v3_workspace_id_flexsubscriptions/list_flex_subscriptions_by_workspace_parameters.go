// Code generated by go-swagger; DO NOT EDIT.

package v3_workspace_id_flexsubscriptions

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

// NewListFlexSubscriptionsByWorkspaceParams creates a new ListFlexSubscriptionsByWorkspaceParams object
// with the default values initialized.
func NewListFlexSubscriptionsByWorkspaceParams() *ListFlexSubscriptionsByWorkspaceParams {
	var ()
	return &ListFlexSubscriptionsByWorkspaceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListFlexSubscriptionsByWorkspaceParamsWithTimeout creates a new ListFlexSubscriptionsByWorkspaceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListFlexSubscriptionsByWorkspaceParamsWithTimeout(timeout time.Duration) *ListFlexSubscriptionsByWorkspaceParams {
	var ()
	return &ListFlexSubscriptionsByWorkspaceParams{

		timeout: timeout,
	}
}

// NewListFlexSubscriptionsByWorkspaceParamsWithContext creates a new ListFlexSubscriptionsByWorkspaceParams object
// with the default values initialized, and the ability to set a context for a request
func NewListFlexSubscriptionsByWorkspaceParamsWithContext(ctx context.Context) *ListFlexSubscriptionsByWorkspaceParams {
	var ()
	return &ListFlexSubscriptionsByWorkspaceParams{

		Context: ctx,
	}
}

// NewListFlexSubscriptionsByWorkspaceParamsWithHTTPClient creates a new ListFlexSubscriptionsByWorkspaceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListFlexSubscriptionsByWorkspaceParamsWithHTTPClient(client *http.Client) *ListFlexSubscriptionsByWorkspaceParams {
	var ()
	return &ListFlexSubscriptionsByWorkspaceParams{
		HTTPClient: client,
	}
}

/*ListFlexSubscriptionsByWorkspaceParams contains all the parameters to send to the API endpoint
for the list flex subscriptions by workspace operation typically these are written to a http.Request
*/
type ListFlexSubscriptionsByWorkspaceParams struct {

	/*WorkspaceID*/
	WorkspaceID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list flex subscriptions by workspace params
func (o *ListFlexSubscriptionsByWorkspaceParams) WithTimeout(timeout time.Duration) *ListFlexSubscriptionsByWorkspaceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list flex subscriptions by workspace params
func (o *ListFlexSubscriptionsByWorkspaceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list flex subscriptions by workspace params
func (o *ListFlexSubscriptionsByWorkspaceParams) WithContext(ctx context.Context) *ListFlexSubscriptionsByWorkspaceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list flex subscriptions by workspace params
func (o *ListFlexSubscriptionsByWorkspaceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list flex subscriptions by workspace params
func (o *ListFlexSubscriptionsByWorkspaceParams) WithHTTPClient(client *http.Client) *ListFlexSubscriptionsByWorkspaceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list flex subscriptions by workspace params
func (o *ListFlexSubscriptionsByWorkspaceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithWorkspaceID adds the workspaceID to the list flex subscriptions by workspace params
func (o *ListFlexSubscriptionsByWorkspaceParams) WithWorkspaceID(workspaceID int64) *ListFlexSubscriptionsByWorkspaceParams {
	o.SetWorkspaceID(workspaceID)
	return o
}

// SetWorkspaceID adds the workspaceId to the list flex subscriptions by workspace params
func (o *ListFlexSubscriptionsByWorkspaceParams) SetWorkspaceID(workspaceID int64) {
	o.WorkspaceID = workspaceID
}

// WriteToRequest writes these params to a swagger request
func (o *ListFlexSubscriptionsByWorkspaceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param workspaceId
	if err := r.SetPathParam("workspaceId", swag.FormatInt64(o.WorkspaceID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
