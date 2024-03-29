// Code generated by go-swagger; DO NOT EDIT.

package v1distrox

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

// NewDeleteInstancesDistroXV1ByCrnParams creates a new DeleteInstancesDistroXV1ByCrnParams object
// with the default values initialized.
func NewDeleteInstancesDistroXV1ByCrnParams() *DeleteInstancesDistroXV1ByCrnParams {
	var (
		forcedDefault = bool(false)
	)
	return &DeleteInstancesDistroXV1ByCrnParams{
		Forced: &forcedDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteInstancesDistroXV1ByCrnParamsWithTimeout creates a new DeleteInstancesDistroXV1ByCrnParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteInstancesDistroXV1ByCrnParamsWithTimeout(timeout time.Duration) *DeleteInstancesDistroXV1ByCrnParams {
	var (
		forcedDefault = bool(false)
	)
	return &DeleteInstancesDistroXV1ByCrnParams{
		Forced: &forcedDefault,

		timeout: timeout,
	}
}

// NewDeleteInstancesDistroXV1ByCrnParamsWithContext creates a new DeleteInstancesDistroXV1ByCrnParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteInstancesDistroXV1ByCrnParamsWithContext(ctx context.Context) *DeleteInstancesDistroXV1ByCrnParams {
	var (
		forcedDefault = bool(false)
	)
	return &DeleteInstancesDistroXV1ByCrnParams{
		Forced: &forcedDefault,

		Context: ctx,
	}
}

// NewDeleteInstancesDistroXV1ByCrnParamsWithHTTPClient creates a new DeleteInstancesDistroXV1ByCrnParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteInstancesDistroXV1ByCrnParamsWithHTTPClient(client *http.Client) *DeleteInstancesDistroXV1ByCrnParams {
	var (
		forcedDefault = bool(false)
	)
	return &DeleteInstancesDistroXV1ByCrnParams{
		Forced:     &forcedDefault,
		HTTPClient: client,
	}
}

/*
DeleteInstancesDistroXV1ByCrnParams contains all the parameters to send to the API endpoint
for the delete instances distro x v1 by crn operation typically these are written to a http.Request
*/
type DeleteInstancesDistroXV1ByCrnParams struct {

	/*Body*/
	Body *model.MultipleInstanceDeleteRequest
	/*Crn*/
	Crn string
	/*Forced*/
	Forced *bool
	/*ID*/
	ID []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete instances distro x v1 by crn params
func (o *DeleteInstancesDistroXV1ByCrnParams) WithTimeout(timeout time.Duration) *DeleteInstancesDistroXV1ByCrnParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete instances distro x v1 by crn params
func (o *DeleteInstancesDistroXV1ByCrnParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete instances distro x v1 by crn params
func (o *DeleteInstancesDistroXV1ByCrnParams) WithContext(ctx context.Context) *DeleteInstancesDistroXV1ByCrnParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete instances distro x v1 by crn params
func (o *DeleteInstancesDistroXV1ByCrnParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete instances distro x v1 by crn params
func (o *DeleteInstancesDistroXV1ByCrnParams) WithHTTPClient(client *http.Client) *DeleteInstancesDistroXV1ByCrnParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete instances distro x v1 by crn params
func (o *DeleteInstancesDistroXV1ByCrnParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the delete instances distro x v1 by crn params
func (o *DeleteInstancesDistroXV1ByCrnParams) WithBody(body *model.MultipleInstanceDeleteRequest) *DeleteInstancesDistroXV1ByCrnParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the delete instances distro x v1 by crn params
func (o *DeleteInstancesDistroXV1ByCrnParams) SetBody(body *model.MultipleInstanceDeleteRequest) {
	o.Body = body
}

// WithCrn adds the crn to the delete instances distro x v1 by crn params
func (o *DeleteInstancesDistroXV1ByCrnParams) WithCrn(crn string) *DeleteInstancesDistroXV1ByCrnParams {
	o.SetCrn(crn)
	return o
}

// SetCrn adds the crn to the delete instances distro x v1 by crn params
func (o *DeleteInstancesDistroXV1ByCrnParams) SetCrn(crn string) {
	o.Crn = crn
}

// WithForced adds the forced to the delete instances distro x v1 by crn params
func (o *DeleteInstancesDistroXV1ByCrnParams) WithForced(forced *bool) *DeleteInstancesDistroXV1ByCrnParams {
	o.SetForced(forced)
	return o
}

// SetForced adds the forced to the delete instances distro x v1 by crn params
func (o *DeleteInstancesDistroXV1ByCrnParams) SetForced(forced *bool) {
	o.Forced = forced
}

// WithID adds the id to the delete instances distro x v1 by crn params
func (o *DeleteInstancesDistroXV1ByCrnParams) WithID(id []string) *DeleteInstancesDistroXV1ByCrnParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete instances distro x v1 by crn params
func (o *DeleteInstancesDistroXV1ByCrnParams) SetID(id []string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteInstancesDistroXV1ByCrnParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

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

	valuesID := o.ID

	joinedID := swag.JoinByFormat(valuesID, "multi")
	// query array param id
	if err := r.SetQueryParam("id", joinedID...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
