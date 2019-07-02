// Code generated by go-swagger; DO NOT EDIT.

package v1kerberosmgmt

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

	model "github.com/hortonworks/cb-cli/dataplane/api-freeipa/model"
)

// NewGetServiceKeytabV1Params creates a new GetServiceKeytabV1Params object
// with the default values initialized.
func NewGetServiceKeytabV1Params() *GetServiceKeytabV1Params {
	var ()
	return &GetServiceKeytabV1Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetServiceKeytabV1ParamsWithTimeout creates a new GetServiceKeytabV1Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetServiceKeytabV1ParamsWithTimeout(timeout time.Duration) *GetServiceKeytabV1Params {
	var ()
	return &GetServiceKeytabV1Params{

		timeout: timeout,
	}
}

// NewGetServiceKeytabV1ParamsWithContext creates a new GetServiceKeytabV1Params object
// with the default values initialized, and the ability to set a context for a request
func NewGetServiceKeytabV1ParamsWithContext(ctx context.Context) *GetServiceKeytabV1Params {
	var ()
	return &GetServiceKeytabV1Params{

		Context: ctx,
	}
}

// NewGetServiceKeytabV1ParamsWithHTTPClient creates a new GetServiceKeytabV1Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetServiceKeytabV1ParamsWithHTTPClient(client *http.Client) *GetServiceKeytabV1Params {
	var ()
	return &GetServiceKeytabV1Params{
		HTTPClient: client,
	}
}

/*GetServiceKeytabV1Params contains all the parameters to send to the API endpoint
for the get service keytab v1 operation typically these are written to a http.Request
*/
type GetServiceKeytabV1Params struct {

	/*Body*/
	Body *model.ServiceKeytabV1Request

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get service keytab v1 params
func (o *GetServiceKeytabV1Params) WithTimeout(timeout time.Duration) *GetServiceKeytabV1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get service keytab v1 params
func (o *GetServiceKeytabV1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get service keytab v1 params
func (o *GetServiceKeytabV1Params) WithContext(ctx context.Context) *GetServiceKeytabV1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get service keytab v1 params
func (o *GetServiceKeytabV1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get service keytab v1 params
func (o *GetServiceKeytabV1Params) WithHTTPClient(client *http.Client) *GetServiceKeytabV1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get service keytab v1 params
func (o *GetServiceKeytabV1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the get service keytab v1 params
func (o *GetServiceKeytabV1Params) WithBody(body *model.ServiceKeytabV1Request) *GetServiceKeytabV1Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the get service keytab v1 params
func (o *GetServiceKeytabV1Params) SetBody(body *model.ServiceKeytabV1Request) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *GetServiceKeytabV1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
