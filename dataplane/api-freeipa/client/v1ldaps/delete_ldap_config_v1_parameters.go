// Code generated by go-swagger; DO NOT EDIT.

package v1ldaps

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
)

// NewDeleteLdapConfigV1Params creates a new DeleteLdapConfigV1Params object
// with the default values initialized.
func NewDeleteLdapConfigV1Params() *DeleteLdapConfigV1Params {
	var ()
	return &DeleteLdapConfigV1Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteLdapConfigV1ParamsWithTimeout creates a new DeleteLdapConfigV1Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteLdapConfigV1ParamsWithTimeout(timeout time.Duration) *DeleteLdapConfigV1Params {
	var ()
	return &DeleteLdapConfigV1Params{

		timeout: timeout,
	}
}

// NewDeleteLdapConfigV1ParamsWithContext creates a new DeleteLdapConfigV1Params object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteLdapConfigV1ParamsWithContext(ctx context.Context) *DeleteLdapConfigV1Params {
	var ()
	return &DeleteLdapConfigV1Params{

		Context: ctx,
	}
}

// NewDeleteLdapConfigV1ParamsWithHTTPClient creates a new DeleteLdapConfigV1Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteLdapConfigV1ParamsWithHTTPClient(client *http.Client) *DeleteLdapConfigV1Params {
	var ()
	return &DeleteLdapConfigV1Params{
		HTTPClient: client,
	}
}

/*DeleteLdapConfigV1Params contains all the parameters to send to the API endpoint
for the delete ldap config v1 operation typically these are written to a http.Request
*/
type DeleteLdapConfigV1Params struct {

	/*Environment*/
	Environment *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete ldap config v1 params
func (o *DeleteLdapConfigV1Params) WithTimeout(timeout time.Duration) *DeleteLdapConfigV1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete ldap config v1 params
func (o *DeleteLdapConfigV1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete ldap config v1 params
func (o *DeleteLdapConfigV1Params) WithContext(ctx context.Context) *DeleteLdapConfigV1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete ldap config v1 params
func (o *DeleteLdapConfigV1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete ldap config v1 params
func (o *DeleteLdapConfigV1Params) WithHTTPClient(client *http.Client) *DeleteLdapConfigV1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete ldap config v1 params
func (o *DeleteLdapConfigV1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEnvironment adds the environment to the delete ldap config v1 params
func (o *DeleteLdapConfigV1Params) WithEnvironment(environment *string) *DeleteLdapConfigV1Params {
	o.SetEnvironment(environment)
	return o
}

// SetEnvironment adds the environment to the delete ldap config v1 params
func (o *DeleteLdapConfigV1Params) SetEnvironment(environment *string) {
	o.Environment = environment
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteLdapConfigV1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Environment != nil {

		// query param environment
		var qrEnvironment string
		if o.Environment != nil {
			qrEnvironment = *o.Environment
		}
		qEnvironment := qrEnvironment
		if qEnvironment != "" {
			if err := r.SetQueryParam("environment", qEnvironment); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
