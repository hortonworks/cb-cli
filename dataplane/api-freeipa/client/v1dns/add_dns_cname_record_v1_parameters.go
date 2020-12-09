// Code generated by go-swagger; DO NOT EDIT.

package v1dns

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

// NewAddDNSCnameRecordV1Params creates a new AddDNSCnameRecordV1Params object
// with the default values initialized.
func NewAddDNSCnameRecordV1Params() *AddDNSCnameRecordV1Params {
	var ()
	return &AddDNSCnameRecordV1Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddDNSCnameRecordV1ParamsWithTimeout creates a new AddDNSCnameRecordV1Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddDNSCnameRecordV1ParamsWithTimeout(timeout time.Duration) *AddDNSCnameRecordV1Params {
	var ()
	return &AddDNSCnameRecordV1Params{

		timeout: timeout,
	}
}

// NewAddDNSCnameRecordV1ParamsWithContext creates a new AddDNSCnameRecordV1Params object
// with the default values initialized, and the ability to set a context for a request
func NewAddDNSCnameRecordV1ParamsWithContext(ctx context.Context) *AddDNSCnameRecordV1Params {
	var ()
	return &AddDNSCnameRecordV1Params{

		Context: ctx,
	}
}

// NewAddDNSCnameRecordV1ParamsWithHTTPClient creates a new AddDNSCnameRecordV1Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddDNSCnameRecordV1ParamsWithHTTPClient(client *http.Client) *AddDNSCnameRecordV1Params {
	var ()
	return &AddDNSCnameRecordV1Params{
		HTTPClient: client,
	}
}

/*AddDNSCnameRecordV1Params contains all the parameters to send to the API endpoint
for the add Dns cname record v1 operation typically these are written to a http.Request
*/
type AddDNSCnameRecordV1Params struct {

	/*Body*/
	Body *model.AddDNSCnameRecordV1Request

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the add Dns cname record v1 params
func (o *AddDNSCnameRecordV1Params) WithTimeout(timeout time.Duration) *AddDNSCnameRecordV1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add Dns cname record v1 params
func (o *AddDNSCnameRecordV1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add Dns cname record v1 params
func (o *AddDNSCnameRecordV1Params) WithContext(ctx context.Context) *AddDNSCnameRecordV1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add Dns cname record v1 params
func (o *AddDNSCnameRecordV1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add Dns cname record v1 params
func (o *AddDNSCnameRecordV1Params) WithHTTPClient(client *http.Client) *AddDNSCnameRecordV1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add Dns cname record v1 params
func (o *AddDNSCnameRecordV1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the add Dns cname record v1 params
func (o *AddDNSCnameRecordV1Params) WithBody(body *model.AddDNSCnameRecordV1Request) *AddDNSCnameRecordV1Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add Dns cname record v1 params
func (o *AddDNSCnameRecordV1Params) SetBody(body *model.AddDNSCnameRecordV1Request) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *AddDNSCnameRecordV1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
