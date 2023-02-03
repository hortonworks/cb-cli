// Code generated by go-swagger; DO NOT EDIT.

package v1ldaps

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

// NewTestLdapConfigV1Params creates a new TestLdapConfigV1Params object
// with the default values initialized.
func NewTestLdapConfigV1Params() *TestLdapConfigV1Params {
	var ()
	return &TestLdapConfigV1Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewTestLdapConfigV1ParamsWithTimeout creates a new TestLdapConfigV1Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewTestLdapConfigV1ParamsWithTimeout(timeout time.Duration) *TestLdapConfigV1Params {
	var ()
	return &TestLdapConfigV1Params{

		timeout: timeout,
	}
}

// NewTestLdapConfigV1ParamsWithContext creates a new TestLdapConfigV1Params object
// with the default values initialized, and the ability to set a context for a request
func NewTestLdapConfigV1ParamsWithContext(ctx context.Context) *TestLdapConfigV1Params {
	var ()
	return &TestLdapConfigV1Params{

		Context: ctx,
	}
}

// NewTestLdapConfigV1ParamsWithHTTPClient creates a new TestLdapConfigV1Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTestLdapConfigV1ParamsWithHTTPClient(client *http.Client) *TestLdapConfigV1Params {
	var ()
	return &TestLdapConfigV1Params{
		HTTPClient: client,
	}
}

/*
TestLdapConfigV1Params contains all the parameters to send to the API endpoint
for the test ldap config v1 operation typically these are written to a http.Request
*/
type TestLdapConfigV1Params struct {

	/*Body*/
	Body *model.TestLdapConfigV1Request

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the test ldap config v1 params
func (o *TestLdapConfigV1Params) WithTimeout(timeout time.Duration) *TestLdapConfigV1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the test ldap config v1 params
func (o *TestLdapConfigV1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the test ldap config v1 params
func (o *TestLdapConfigV1Params) WithContext(ctx context.Context) *TestLdapConfigV1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the test ldap config v1 params
func (o *TestLdapConfigV1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the test ldap config v1 params
func (o *TestLdapConfigV1Params) WithHTTPClient(client *http.Client) *TestLdapConfigV1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the test ldap config v1 params
func (o *TestLdapConfigV1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the test ldap config v1 params
func (o *TestLdapConfigV1Params) WithBody(body *model.TestLdapConfigV1Request) *TestLdapConfigV1Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the test ldap config v1 params
func (o *TestLdapConfigV1Params) SetBody(body *model.TestLdapConfigV1Request) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *TestLdapConfigV1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
