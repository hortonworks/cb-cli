// Code generated by go-swagger; DO NOT EDIT.

package v1envcost

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

	model "github.com/hortonworks/cb-cli/dataplane/api-environment/model"
)

// NewListEnvironmentCostV1Params creates a new ListEnvironmentCostV1Params object
// with the default values initialized.
func NewListEnvironmentCostV1Params() *ListEnvironmentCostV1Params {
	var ()
	return &ListEnvironmentCostV1Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewListEnvironmentCostV1ParamsWithTimeout creates a new ListEnvironmentCostV1Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewListEnvironmentCostV1ParamsWithTimeout(timeout time.Duration) *ListEnvironmentCostV1Params {
	var ()
	return &ListEnvironmentCostV1Params{

		timeout: timeout,
	}
}

// NewListEnvironmentCostV1ParamsWithContext creates a new ListEnvironmentCostV1Params object
// with the default values initialized, and the ability to set a context for a request
func NewListEnvironmentCostV1ParamsWithContext(ctx context.Context) *ListEnvironmentCostV1Params {
	var ()
	return &ListEnvironmentCostV1Params{

		Context: ctx,
	}
}

// NewListEnvironmentCostV1ParamsWithHTTPClient creates a new ListEnvironmentCostV1Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListEnvironmentCostV1ParamsWithHTTPClient(client *http.Client) *ListEnvironmentCostV1Params {
	var ()
	return &ListEnvironmentCostV1Params{
		HTTPClient: client,
	}
}

/*
ListEnvironmentCostV1Params contains all the parameters to send to the API endpoint
for the list environment cost v1 operation typically these are written to a http.Request
*/
type ListEnvironmentCostV1Params struct {

	/*Body*/
	Body *model.EnvironmentRealTimeCostV1Request

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list environment cost v1 params
func (o *ListEnvironmentCostV1Params) WithTimeout(timeout time.Duration) *ListEnvironmentCostV1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list environment cost v1 params
func (o *ListEnvironmentCostV1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list environment cost v1 params
func (o *ListEnvironmentCostV1Params) WithContext(ctx context.Context) *ListEnvironmentCostV1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list environment cost v1 params
func (o *ListEnvironmentCostV1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list environment cost v1 params
func (o *ListEnvironmentCostV1Params) WithHTTPClient(client *http.Client) *ListEnvironmentCostV1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list environment cost v1 params
func (o *ListEnvironmentCostV1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the list environment cost v1 params
func (o *ListEnvironmentCostV1Params) WithBody(body *model.EnvironmentRealTimeCostV1Request) *ListEnvironmentCostV1Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the list environment cost v1 params
func (o *ListEnvironmentCostV1Params) SetBody(body *model.EnvironmentRealTimeCostV1Request) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ListEnvironmentCostV1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
