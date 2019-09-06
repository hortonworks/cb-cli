// Code generated by go-swagger; DO NOT EDIT.

package v3_workspace_id_stacks

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

// NewDeleteMultipleInstancesStackV3Params creates a new DeleteMultipleInstancesStackV3Params object
// with the default values initialized.
func NewDeleteMultipleInstancesStackV3Params() *DeleteMultipleInstancesStackV3Params {
	var (
		forcedDefault = bool(false)
	)
	return &DeleteMultipleInstancesStackV3Params{
		Forced: &forcedDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteMultipleInstancesStackV3ParamsWithTimeout creates a new DeleteMultipleInstancesStackV3Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteMultipleInstancesStackV3ParamsWithTimeout(timeout time.Duration) *DeleteMultipleInstancesStackV3Params {
	var (
		forcedDefault = bool(false)
	)
	return &DeleteMultipleInstancesStackV3Params{
		Forced: &forcedDefault,

		timeout: timeout,
	}
}

// NewDeleteMultipleInstancesStackV3ParamsWithContext creates a new DeleteMultipleInstancesStackV3Params object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteMultipleInstancesStackV3ParamsWithContext(ctx context.Context) *DeleteMultipleInstancesStackV3Params {
	var (
		forcedDefault = bool(false)
	)
	return &DeleteMultipleInstancesStackV3Params{
		Forced: &forcedDefault,

		Context: ctx,
	}
}

// NewDeleteMultipleInstancesStackV3ParamsWithHTTPClient creates a new DeleteMultipleInstancesStackV3Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteMultipleInstancesStackV3ParamsWithHTTPClient(client *http.Client) *DeleteMultipleInstancesStackV3Params {
	var (
		forcedDefault = bool(false)
	)
	return &DeleteMultipleInstancesStackV3Params{
		Forced:     &forcedDefault,
		HTTPClient: client,
	}
}

/*DeleteMultipleInstancesStackV3Params contains all the parameters to send to the API endpoint
for the delete multiple instances stack v3 operation typically these are written to a http.Request
*/
type DeleteMultipleInstancesStackV3Params struct {

	/*Forced*/
	Forced *bool
	/*ID*/
	ID []string
	/*Name*/
	Name string
	/*WorkspaceID*/
	WorkspaceID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete multiple instances stack v3 params
func (o *DeleteMultipleInstancesStackV3Params) WithTimeout(timeout time.Duration) *DeleteMultipleInstancesStackV3Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete multiple instances stack v3 params
func (o *DeleteMultipleInstancesStackV3Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete multiple instances stack v3 params
func (o *DeleteMultipleInstancesStackV3Params) WithContext(ctx context.Context) *DeleteMultipleInstancesStackV3Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete multiple instances stack v3 params
func (o *DeleteMultipleInstancesStackV3Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete multiple instances stack v3 params
func (o *DeleteMultipleInstancesStackV3Params) WithHTTPClient(client *http.Client) *DeleteMultipleInstancesStackV3Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete multiple instances stack v3 params
func (o *DeleteMultipleInstancesStackV3Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithForced adds the forced to the delete multiple instances stack v3 params
func (o *DeleteMultipleInstancesStackV3Params) WithForced(forced *bool) *DeleteMultipleInstancesStackV3Params {
	o.SetForced(forced)
	return o
}

// SetForced adds the forced to the delete multiple instances stack v3 params
func (o *DeleteMultipleInstancesStackV3Params) SetForced(forced *bool) {
	o.Forced = forced
}

// WithID adds the id to the delete multiple instances stack v3 params
func (o *DeleteMultipleInstancesStackV3Params) WithID(id []string) *DeleteMultipleInstancesStackV3Params {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete multiple instances stack v3 params
func (o *DeleteMultipleInstancesStackV3Params) SetID(id []string) {
	o.ID = id
}

// WithName adds the name to the delete multiple instances stack v3 params
func (o *DeleteMultipleInstancesStackV3Params) WithName(name string) *DeleteMultipleInstancesStackV3Params {
	o.SetName(name)
	return o
}

// SetName adds the name to the delete multiple instances stack v3 params
func (o *DeleteMultipleInstancesStackV3Params) SetName(name string) {
	o.Name = name
}

// WithWorkspaceID adds the workspaceID to the delete multiple instances stack v3 params
func (o *DeleteMultipleInstancesStackV3Params) WithWorkspaceID(workspaceID int64) *DeleteMultipleInstancesStackV3Params {
	o.SetWorkspaceID(workspaceID)
	return o
}

// SetWorkspaceID adds the workspaceId to the delete multiple instances stack v3 params
func (o *DeleteMultipleInstancesStackV3Params) SetWorkspaceID(workspaceID int64) {
	o.WorkspaceID = workspaceID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteMultipleInstancesStackV3Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Forced != nil {

		// query param forced
		var qrForced bool
		if o.Forced != nil {
			qrForced = *o.Forced
		}
		qForced := swag.FormatBool(qrForced)
		if qForced != "" {
			if err := r.SetQueryParam("forced", qForced); err != nil {
				return err
			}
		}

	}

	valuesID := o.ID

	joinedID := swag.JoinByFormat(valuesID, "multi")
	// query array param id
	if err := r.SetQueryParam("id", joinedID...); err != nil {
		return err
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
