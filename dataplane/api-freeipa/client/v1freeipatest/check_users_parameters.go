// Code generated by go-swagger; DO NOT EDIT.

package v1freeipatest

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api-freeipa/model"
)

// NewCheckUsersParams creates a new CheckUsersParams object
// with the default values initialized.
func NewCheckUsersParams() *CheckUsersParams {
	var ()
	return &CheckUsersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCheckUsersParamsWithTimeout creates a new CheckUsersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCheckUsersParamsWithTimeout(timeout time.Duration) *CheckUsersParams {
	var ()
	return &CheckUsersParams{

		timeout: timeout,
	}
}

// NewCheckUsersParamsWithContext creates a new CheckUsersParams object
// with the default values initialized, and the ability to set a context for a request
func NewCheckUsersParamsWithContext(ctx context.Context) *CheckUsersParams {
	var ()
	return &CheckUsersParams{

		Context: ctx,
	}
}

// NewCheckUsersParamsWithHTTPClient creates a new CheckUsersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCheckUsersParamsWithHTTPClient(client *http.Client) *CheckUsersParams {
	var ()
	return &CheckUsersParams{
		HTTPClient: client,
	}
}

/*
CheckUsersParams contains all the parameters to send to the API endpoint
for the check users operation typically these are written to a http.Request
*/
type CheckUsersParams struct {

	/*Body*/
	Body *model.CheckUsersV1Request

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the check users params
func (o *CheckUsersParams) WithTimeout(timeout time.Duration) *CheckUsersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the check users params
func (o *CheckUsersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the check users params
func (o *CheckUsersParams) WithContext(ctx context.Context) *CheckUsersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the check users params
func (o *CheckUsersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the check users params
func (o *CheckUsersParams) WithHTTPClient(client *http.Client) *CheckUsersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the check users params
func (o *CheckUsersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the check users params
func (o *CheckUsersParams) WithBody(body *model.CheckUsersV1Request) *CheckUsersParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the check users params
func (o *CheckUsersParams) SetBody(body *model.CheckUsersV1Request) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CheckUsersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
