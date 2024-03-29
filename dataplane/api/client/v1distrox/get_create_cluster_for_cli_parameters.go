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

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api/model"
)

// NewGetCreateClusterForCliParams creates a new GetCreateClusterForCliParams object
// with the default values initialized.
func NewGetCreateClusterForCliParams() *GetCreateClusterForCliParams {
	var ()
	return &GetCreateClusterForCliParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCreateClusterForCliParamsWithTimeout creates a new GetCreateClusterForCliParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCreateClusterForCliParamsWithTimeout(timeout time.Duration) *GetCreateClusterForCliParams {
	var ()
	return &GetCreateClusterForCliParams{

		timeout: timeout,
	}
}

// NewGetCreateClusterForCliParamsWithContext creates a new GetCreateClusterForCliParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCreateClusterForCliParamsWithContext(ctx context.Context) *GetCreateClusterForCliParams {
	var ()
	return &GetCreateClusterForCliParams{

		Context: ctx,
	}
}

// NewGetCreateClusterForCliParamsWithHTTPClient creates a new GetCreateClusterForCliParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCreateClusterForCliParamsWithHTTPClient(client *http.Client) *GetCreateClusterForCliParams {
	var ()
	return &GetCreateClusterForCliParams{
		HTTPClient: client,
	}
}

/*
GetCreateClusterForCliParams contains all the parameters to send to the API endpoint
for the get create cluster for cli operation typically these are written to a http.Request
*/
type GetCreateClusterForCliParams struct {

	/*Body*/
	Body *model.DistroXV1Request

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get create cluster for cli params
func (o *GetCreateClusterForCliParams) WithTimeout(timeout time.Duration) *GetCreateClusterForCliParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get create cluster for cli params
func (o *GetCreateClusterForCliParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get create cluster for cli params
func (o *GetCreateClusterForCliParams) WithContext(ctx context.Context) *GetCreateClusterForCliParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get create cluster for cli params
func (o *GetCreateClusterForCliParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get create cluster for cli params
func (o *GetCreateClusterForCliParams) WithHTTPClient(client *http.Client) *GetCreateClusterForCliParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get create cluster for cli params
func (o *GetCreateClusterForCliParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the get create cluster for cli params
func (o *GetCreateClusterForCliParams) WithBody(body *model.DistroXV1Request) *GetCreateClusterForCliParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the get create cluster for cli params
func (o *GetCreateClusterForCliParams) SetBody(body *model.DistroXV1Request) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *GetCreateClusterForCliParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
