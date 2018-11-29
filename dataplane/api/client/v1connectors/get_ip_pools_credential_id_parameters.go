// Code generated by go-swagger; DO NOT EDIT.

package v1connectors

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api/model"
)

// NewGetIPPoolsCredentialIDParams creates a new GetIPPoolsCredentialIDParams object
// with the default values initialized.
func NewGetIPPoolsCredentialIDParams() *GetIPPoolsCredentialIDParams {
	var ()
	return &GetIPPoolsCredentialIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetIPPoolsCredentialIDParamsWithTimeout creates a new GetIPPoolsCredentialIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetIPPoolsCredentialIDParamsWithTimeout(timeout time.Duration) *GetIPPoolsCredentialIDParams {
	var ()
	return &GetIPPoolsCredentialIDParams{

		timeout: timeout,
	}
}

// NewGetIPPoolsCredentialIDParamsWithContext creates a new GetIPPoolsCredentialIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetIPPoolsCredentialIDParamsWithContext(ctx context.Context) *GetIPPoolsCredentialIDParams {
	var ()
	return &GetIPPoolsCredentialIDParams{

		Context: ctx,
	}
}

// NewGetIPPoolsCredentialIDParamsWithHTTPClient creates a new GetIPPoolsCredentialIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetIPPoolsCredentialIDParamsWithHTTPClient(client *http.Client) *GetIPPoolsCredentialIDParams {
	var ()
	return &GetIPPoolsCredentialIDParams{
		HTTPClient: client,
	}
}

/*GetIPPoolsCredentialIDParams contains all the parameters to send to the API endpoint
for the get Ip pools credential Id operation typically these are written to a http.Request
*/
type GetIPPoolsCredentialIDParams struct {

	/*Body*/
	Body *model.PlatformResourceRequestJSON

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get Ip pools credential Id params
func (o *GetIPPoolsCredentialIDParams) WithTimeout(timeout time.Duration) *GetIPPoolsCredentialIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get Ip pools credential Id params
func (o *GetIPPoolsCredentialIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get Ip pools credential Id params
func (o *GetIPPoolsCredentialIDParams) WithContext(ctx context.Context) *GetIPPoolsCredentialIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get Ip pools credential Id params
func (o *GetIPPoolsCredentialIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get Ip pools credential Id params
func (o *GetIPPoolsCredentialIDParams) WithHTTPClient(client *http.Client) *GetIPPoolsCredentialIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get Ip pools credential Id params
func (o *GetIPPoolsCredentialIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the get Ip pools credential Id params
func (o *GetIPPoolsCredentialIDParams) WithBody(body *model.PlatformResourceRequestJSON) *GetIPPoolsCredentialIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the get Ip pools credential Id params
func (o *GetIPPoolsCredentialIDParams) SetBody(body *model.PlatformResourceRequestJSON) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *GetIPPoolsCredentialIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
