// Code generated by go-swagger; DO NOT EDIT.

package v4_workspace_id_clustertemplates

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

// NewDeleteClusterTemplateByCrnInWorkspaceParams creates a new DeleteClusterTemplateByCrnInWorkspaceParams object
// with the default values initialized.
func NewDeleteClusterTemplateByCrnInWorkspaceParams() *DeleteClusterTemplateByCrnInWorkspaceParams {
	var ()
	return &DeleteClusterTemplateByCrnInWorkspaceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteClusterTemplateByCrnInWorkspaceParamsWithTimeout creates a new DeleteClusterTemplateByCrnInWorkspaceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteClusterTemplateByCrnInWorkspaceParamsWithTimeout(timeout time.Duration) *DeleteClusterTemplateByCrnInWorkspaceParams {
	var ()
	return &DeleteClusterTemplateByCrnInWorkspaceParams{

		timeout: timeout,
	}
}

// NewDeleteClusterTemplateByCrnInWorkspaceParamsWithContext creates a new DeleteClusterTemplateByCrnInWorkspaceParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteClusterTemplateByCrnInWorkspaceParamsWithContext(ctx context.Context) *DeleteClusterTemplateByCrnInWorkspaceParams {
	var ()
	return &DeleteClusterTemplateByCrnInWorkspaceParams{

		Context: ctx,
	}
}

// NewDeleteClusterTemplateByCrnInWorkspaceParamsWithHTTPClient creates a new DeleteClusterTemplateByCrnInWorkspaceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteClusterTemplateByCrnInWorkspaceParamsWithHTTPClient(client *http.Client) *DeleteClusterTemplateByCrnInWorkspaceParams {
	var ()
	return &DeleteClusterTemplateByCrnInWorkspaceParams{
		HTTPClient: client,
	}
}

/*DeleteClusterTemplateByCrnInWorkspaceParams contains all the parameters to send to the API endpoint
for the delete cluster template by crn in workspace operation typically these are written to a http.Request
*/
type DeleteClusterTemplateByCrnInWorkspaceParams struct {

	/*Crn*/
	Crn string
	/*WorkspaceID*/
	WorkspaceID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete cluster template by crn in workspace params
func (o *DeleteClusterTemplateByCrnInWorkspaceParams) WithTimeout(timeout time.Duration) *DeleteClusterTemplateByCrnInWorkspaceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete cluster template by crn in workspace params
func (o *DeleteClusterTemplateByCrnInWorkspaceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete cluster template by crn in workspace params
func (o *DeleteClusterTemplateByCrnInWorkspaceParams) WithContext(ctx context.Context) *DeleteClusterTemplateByCrnInWorkspaceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete cluster template by crn in workspace params
func (o *DeleteClusterTemplateByCrnInWorkspaceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete cluster template by crn in workspace params
func (o *DeleteClusterTemplateByCrnInWorkspaceParams) WithHTTPClient(client *http.Client) *DeleteClusterTemplateByCrnInWorkspaceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete cluster template by crn in workspace params
func (o *DeleteClusterTemplateByCrnInWorkspaceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCrn adds the crn to the delete cluster template by crn in workspace params
func (o *DeleteClusterTemplateByCrnInWorkspaceParams) WithCrn(crn string) *DeleteClusterTemplateByCrnInWorkspaceParams {
	o.SetCrn(crn)
	return o
}

// SetCrn adds the crn to the delete cluster template by crn in workspace params
func (o *DeleteClusterTemplateByCrnInWorkspaceParams) SetCrn(crn string) {
	o.Crn = crn
}

// WithWorkspaceID adds the workspaceID to the delete cluster template by crn in workspace params
func (o *DeleteClusterTemplateByCrnInWorkspaceParams) WithWorkspaceID(workspaceID int64) *DeleteClusterTemplateByCrnInWorkspaceParams {
	o.SetWorkspaceID(workspaceID)
	return o
}

// SetWorkspaceID adds the workspaceId to the delete cluster template by crn in workspace params
func (o *DeleteClusterTemplateByCrnInWorkspaceParams) SetWorkspaceID(workspaceID int64) {
	o.WorkspaceID = workspaceID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteClusterTemplateByCrnInWorkspaceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param crn
	if err := r.SetPathParam("crn", o.Crn); err != nil {
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
