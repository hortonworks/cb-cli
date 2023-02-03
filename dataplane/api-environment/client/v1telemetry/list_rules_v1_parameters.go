// Code generated by go-swagger; DO NOT EDIT.

package v1telemetry

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

// NewListRulesV1Params creates a new ListRulesV1Params object
// with the default values initialized.
func NewListRulesV1Params() *ListRulesV1Params {

	return &ListRulesV1Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewListRulesV1ParamsWithTimeout creates a new ListRulesV1Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewListRulesV1ParamsWithTimeout(timeout time.Duration) *ListRulesV1Params {

	return &ListRulesV1Params{

		timeout: timeout,
	}
}

// NewListRulesV1ParamsWithContext creates a new ListRulesV1Params object
// with the default values initialized, and the ability to set a context for a request
func NewListRulesV1ParamsWithContext(ctx context.Context) *ListRulesV1Params {

	return &ListRulesV1Params{

		Context: ctx,
	}
}

// NewListRulesV1ParamsWithHTTPClient creates a new ListRulesV1Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListRulesV1ParamsWithHTTPClient(client *http.Client) *ListRulesV1Params {

	return &ListRulesV1Params{
		HTTPClient: client,
	}
}

/*
ListRulesV1Params contains all the parameters to send to the API endpoint
for the list rules v1 operation typically these are written to a http.Request
*/
type ListRulesV1Params struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list rules v1 params
func (o *ListRulesV1Params) WithTimeout(timeout time.Duration) *ListRulesV1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list rules v1 params
func (o *ListRulesV1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list rules v1 params
func (o *ListRulesV1Params) WithContext(ctx context.Context) *ListRulesV1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list rules v1 params
func (o *ListRulesV1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list rules v1 params
func (o *ListRulesV1Params) WithHTTPClient(client *http.Client) *ListRulesV1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list rules v1 params
func (o *ListRulesV1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ListRulesV1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
