// Code generated by go-swagger; DO NOT EDIT.

package v1distrox

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

// NewDeleteInstanceDistroXV1ByCrnParams creates a new DeleteInstanceDistroXV1ByCrnParams object
// with the default values initialized.
func NewDeleteInstanceDistroXV1ByCrnParams() *DeleteInstanceDistroXV1ByCrnParams {
	var (
		forcedDefault = bool(false)
	)
	return &DeleteInstanceDistroXV1ByCrnParams{
		Forced: &forcedDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteInstanceDistroXV1ByCrnParamsWithTimeout creates a new DeleteInstanceDistroXV1ByCrnParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteInstanceDistroXV1ByCrnParamsWithTimeout(timeout time.Duration) *DeleteInstanceDistroXV1ByCrnParams {
	var (
		forcedDefault = bool(false)
	)
	return &DeleteInstanceDistroXV1ByCrnParams{
		Forced: &forcedDefault,

		timeout: timeout,
	}
}

// NewDeleteInstanceDistroXV1ByCrnParamsWithContext creates a new DeleteInstanceDistroXV1ByCrnParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteInstanceDistroXV1ByCrnParamsWithContext(ctx context.Context) *DeleteInstanceDistroXV1ByCrnParams {
	var (
		forcedDefault = bool(false)
	)
	return &DeleteInstanceDistroXV1ByCrnParams{
		Forced: &forcedDefault,

		Context: ctx,
	}
}

// NewDeleteInstanceDistroXV1ByCrnParamsWithHTTPClient creates a new DeleteInstanceDistroXV1ByCrnParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteInstanceDistroXV1ByCrnParamsWithHTTPClient(client *http.Client) *DeleteInstanceDistroXV1ByCrnParams {
	var (
		forcedDefault = bool(false)
	)
	return &DeleteInstanceDistroXV1ByCrnParams{
		Forced:     &forcedDefault,
		HTTPClient: client,
	}
}

/*DeleteInstanceDistroXV1ByCrnParams contains all the parameters to send to the API endpoint
for the delete instance distro x v1 by crn operation typically these are written to a http.Request
*/
type DeleteInstanceDistroXV1ByCrnParams struct {

	/*Crn*/
	Crn string
	/*Forced*/
	Forced *bool
	/*InstanceID*/
	InstanceID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete instance distro x v1 by crn params
func (o *DeleteInstanceDistroXV1ByCrnParams) WithTimeout(timeout time.Duration) *DeleteInstanceDistroXV1ByCrnParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete instance distro x v1 by crn params
func (o *DeleteInstanceDistroXV1ByCrnParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete instance distro x v1 by crn params
func (o *DeleteInstanceDistroXV1ByCrnParams) WithContext(ctx context.Context) *DeleteInstanceDistroXV1ByCrnParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete instance distro x v1 by crn params
func (o *DeleteInstanceDistroXV1ByCrnParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete instance distro x v1 by crn params
func (o *DeleteInstanceDistroXV1ByCrnParams) WithHTTPClient(client *http.Client) *DeleteInstanceDistroXV1ByCrnParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete instance distro x v1 by crn params
func (o *DeleteInstanceDistroXV1ByCrnParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCrn adds the crn to the delete instance distro x v1 by crn params
func (o *DeleteInstanceDistroXV1ByCrnParams) WithCrn(crn string) *DeleteInstanceDistroXV1ByCrnParams {
	o.SetCrn(crn)
	return o
}

// SetCrn adds the crn to the delete instance distro x v1 by crn params
func (o *DeleteInstanceDistroXV1ByCrnParams) SetCrn(crn string) {
	o.Crn = crn
}

// WithForced adds the forced to the delete instance distro x v1 by crn params
func (o *DeleteInstanceDistroXV1ByCrnParams) WithForced(forced *bool) *DeleteInstanceDistroXV1ByCrnParams {
	o.SetForced(forced)
	return o
}

// SetForced adds the forced to the delete instance distro x v1 by crn params
func (o *DeleteInstanceDistroXV1ByCrnParams) SetForced(forced *bool) {
	o.Forced = forced
}

// WithInstanceID adds the instanceID to the delete instance distro x v1 by crn params
func (o *DeleteInstanceDistroXV1ByCrnParams) WithInstanceID(instanceID *string) *DeleteInstanceDistroXV1ByCrnParams {
	o.SetInstanceID(instanceID)
	return o
}

// SetInstanceID adds the instanceId to the delete instance distro x v1 by crn params
func (o *DeleteInstanceDistroXV1ByCrnParams) SetInstanceID(instanceID *string) {
	o.InstanceID = instanceID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteInstanceDistroXV1ByCrnParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param crn
	if err := r.SetPathParam("crn", o.Crn); err != nil {
		return err
	}

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

	if o.InstanceID != nil {

		// query param instanceId
		var qrInstanceID string
		if o.InstanceID != nil {
			qrInstanceID = *o.InstanceID
		}
		qInstanceID := qrInstanceID
		if qInstanceID != "" {
			if err := r.SetQueryParam("instanceId", qInstanceID); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
