// Code generated by go-swagger; DO NOT EDIT.

package v3_workspace_id_ldapconfigs

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

// NewGetLdapConfigInWorkspaceParams creates a new GetLdapConfigInWorkspaceParams object
// with the default values initialized.
func NewGetLdapConfigInWorkspaceParams() *GetLdapConfigInWorkspaceParams {
	var ()
	return &GetLdapConfigInWorkspaceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLdapConfigInWorkspaceParamsWithTimeout creates a new GetLdapConfigInWorkspaceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLdapConfigInWorkspaceParamsWithTimeout(timeout time.Duration) *GetLdapConfigInWorkspaceParams {
	var ()
	return &GetLdapConfigInWorkspaceParams{

		timeout: timeout,
	}
}

// NewGetLdapConfigInWorkspaceParamsWithContext creates a new GetLdapConfigInWorkspaceParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLdapConfigInWorkspaceParamsWithContext(ctx context.Context) *GetLdapConfigInWorkspaceParams {
	var ()
	return &GetLdapConfigInWorkspaceParams{

		Context: ctx,
	}
}

// NewGetLdapConfigInWorkspaceParamsWithHTTPClient creates a new GetLdapConfigInWorkspaceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLdapConfigInWorkspaceParamsWithHTTPClient(client *http.Client) *GetLdapConfigInWorkspaceParams {
	var ()
	return &GetLdapConfigInWorkspaceParams{
		HTTPClient: client,
	}
}

/*GetLdapConfigInWorkspaceParams contains all the parameters to send to the API endpoint
for the get ldap config in workspace operation typically these are written to a http.Request
*/
type GetLdapConfigInWorkspaceParams struct {

	/*Name*/
	Name string
	/*WorkspaceID*/
	WorkspaceID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get ldap config in workspace params
func (o *GetLdapConfigInWorkspaceParams) WithTimeout(timeout time.Duration) *GetLdapConfigInWorkspaceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get ldap config in workspace params
func (o *GetLdapConfigInWorkspaceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get ldap config in workspace params
func (o *GetLdapConfigInWorkspaceParams) WithContext(ctx context.Context) *GetLdapConfigInWorkspaceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get ldap config in workspace params
func (o *GetLdapConfigInWorkspaceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get ldap config in workspace params
func (o *GetLdapConfigInWorkspaceParams) WithHTTPClient(client *http.Client) *GetLdapConfigInWorkspaceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get ldap config in workspace params
func (o *GetLdapConfigInWorkspaceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the get ldap config in workspace params
func (o *GetLdapConfigInWorkspaceParams) WithName(name string) *GetLdapConfigInWorkspaceParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get ldap config in workspace params
func (o *GetLdapConfigInWorkspaceParams) SetName(name string) {
	o.Name = name
}

// WithWorkspaceID adds the workspaceID to the get ldap config in workspace params
func (o *GetLdapConfigInWorkspaceParams) WithWorkspaceID(workspaceID int64) *GetLdapConfigInWorkspaceParams {
	o.SetWorkspaceID(workspaceID)
	return o
}

// SetWorkspaceID adds the workspaceId to the get ldap config in workspace params
func (o *GetLdapConfigInWorkspaceParams) SetWorkspaceID(workspaceID int64) {
	o.WorkspaceID = workspaceID
}

// WriteToRequest writes these params to a swagger request
func (o *GetLdapConfigInWorkspaceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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
