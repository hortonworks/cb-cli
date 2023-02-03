// Code generated by go-swagger; DO NOT EDIT.

package v4custom_configurations

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
)

// NewGetRoleTypesParams creates a new GetRoleTypesParams object
// with the default values initialized.
func NewGetRoleTypesParams() *GetRoleTypesParams {

	return &GetRoleTypesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRoleTypesParamsWithTimeout creates a new GetRoleTypesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRoleTypesParamsWithTimeout(timeout time.Duration) *GetRoleTypesParams {

	return &GetRoleTypesParams{

		timeout: timeout,
	}
}

// NewGetRoleTypesParamsWithContext creates a new GetRoleTypesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRoleTypesParamsWithContext(ctx context.Context) *GetRoleTypesParams {

	return &GetRoleTypesParams{

		Context: ctx,
	}
}

// NewGetRoleTypesParamsWithHTTPClient creates a new GetRoleTypesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRoleTypesParamsWithHTTPClient(client *http.Client) *GetRoleTypesParams {

	return &GetRoleTypesParams{
		HTTPClient: client,
	}
}

/*
GetRoleTypesParams contains all the parameters to send to the API endpoint
for the get role types operation typically these are written to a http.Request
*/
type GetRoleTypesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get role types params
func (o *GetRoleTypesParams) WithTimeout(timeout time.Duration) *GetRoleTypesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get role types params
func (o *GetRoleTypesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get role types params
func (o *GetRoleTypesParams) WithContext(ctx context.Context) *GetRoleTypesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get role types params
func (o *GetRoleTypesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get role types params
func (o *GetRoleTypesParams) WithHTTPClient(client *http.Client) *GetRoleTypesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get role types params
func (o *GetRoleTypesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetRoleTypesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
