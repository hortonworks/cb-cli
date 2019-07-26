// Code generated by go-swagger; DO NOT EDIT.

package v4_workspace_id_audits

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

// NewGetAuditEventsInWorkspaceParams creates a new GetAuditEventsInWorkspaceParams object
// with the default values initialized.
func NewGetAuditEventsInWorkspaceParams() *GetAuditEventsInWorkspaceParams {
	var ()
	return &GetAuditEventsInWorkspaceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAuditEventsInWorkspaceParamsWithTimeout creates a new GetAuditEventsInWorkspaceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAuditEventsInWorkspaceParamsWithTimeout(timeout time.Duration) *GetAuditEventsInWorkspaceParams {
	var ()
	return &GetAuditEventsInWorkspaceParams{

		timeout: timeout,
	}
}

// NewGetAuditEventsInWorkspaceParamsWithContext creates a new GetAuditEventsInWorkspaceParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAuditEventsInWorkspaceParamsWithContext(ctx context.Context) *GetAuditEventsInWorkspaceParams {
	var ()
	return &GetAuditEventsInWorkspaceParams{

		Context: ctx,
	}
}

// NewGetAuditEventsInWorkspaceParamsWithHTTPClient creates a new GetAuditEventsInWorkspaceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAuditEventsInWorkspaceParamsWithHTTPClient(client *http.Client) *GetAuditEventsInWorkspaceParams {
	var ()
	return &GetAuditEventsInWorkspaceParams{
		HTTPClient: client,
	}
}

/*GetAuditEventsInWorkspaceParams contains all the parameters to send to the API endpoint
for the get audit events in workspace operation typically these are written to a http.Request
*/
type GetAuditEventsInWorkspaceParams struct {

	/*ResourceID*/
	ResourceID *int64
	/*ResourceType*/
	ResourceType *string
	/*WorkspaceID*/
	WorkspaceID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get audit events in workspace params
func (o *GetAuditEventsInWorkspaceParams) WithTimeout(timeout time.Duration) *GetAuditEventsInWorkspaceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get audit events in workspace params
func (o *GetAuditEventsInWorkspaceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get audit events in workspace params
func (o *GetAuditEventsInWorkspaceParams) WithContext(ctx context.Context) *GetAuditEventsInWorkspaceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get audit events in workspace params
func (o *GetAuditEventsInWorkspaceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get audit events in workspace params
func (o *GetAuditEventsInWorkspaceParams) WithHTTPClient(client *http.Client) *GetAuditEventsInWorkspaceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get audit events in workspace params
func (o *GetAuditEventsInWorkspaceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithResourceID adds the resourceID to the get audit events in workspace params
func (o *GetAuditEventsInWorkspaceParams) WithResourceID(resourceID *int64) *GetAuditEventsInWorkspaceParams {
	o.SetResourceID(resourceID)
	return o
}

// SetResourceID adds the resourceId to the get audit events in workspace params
func (o *GetAuditEventsInWorkspaceParams) SetResourceID(resourceID *int64) {
	o.ResourceID = resourceID
}

// WithResourceType adds the resourceType to the get audit events in workspace params
func (o *GetAuditEventsInWorkspaceParams) WithResourceType(resourceType *string) *GetAuditEventsInWorkspaceParams {
	o.SetResourceType(resourceType)
	return o
}

// SetResourceType adds the resourceType to the get audit events in workspace params
func (o *GetAuditEventsInWorkspaceParams) SetResourceType(resourceType *string) {
	o.ResourceType = resourceType
}

// WithWorkspaceID adds the workspaceID to the get audit events in workspace params
func (o *GetAuditEventsInWorkspaceParams) WithWorkspaceID(workspaceID int64) *GetAuditEventsInWorkspaceParams {
	o.SetWorkspaceID(workspaceID)
	return o
}

// SetWorkspaceID adds the workspaceId to the get audit events in workspace params
func (o *GetAuditEventsInWorkspaceParams) SetWorkspaceID(workspaceID int64) {
	o.WorkspaceID = workspaceID
}

// WriteToRequest writes these params to a swagger request
func (o *GetAuditEventsInWorkspaceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ResourceID != nil {

		// query param resourceId
		var qrResourceID int64
		if o.ResourceID != nil {
			qrResourceID = *o.ResourceID
		}
		qResourceID := swag.FormatInt64(qrResourceID)
		if qResourceID != "" {
			if err := r.SetQueryParam("resourceId", qResourceID); err != nil {
				return err
			}
		}

	}

	if o.ResourceType != nil {

		// query param resourceType
		var qrResourceType string
		if o.ResourceType != nil {
			qrResourceType = *o.ResourceType
		}
		qResourceType := qrResourceType
		if qResourceType != "" {
			if err := r.SetQueryParam("resourceType", qResourceType); err != nil {
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
