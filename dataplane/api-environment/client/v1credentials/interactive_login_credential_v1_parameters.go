// Code generated by go-swagger; DO NOT EDIT.

package v1credentials

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

	model "github.com/hortonworks/cb-cli/dataplane/api-environment/model"
)

// NewInteractiveLoginCredentialV1Params creates a new InteractiveLoginCredentialV1Params object
// with the default values initialized.
func NewInteractiveLoginCredentialV1Params() *InteractiveLoginCredentialV1Params {
	var ()
	return &InteractiveLoginCredentialV1Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewInteractiveLoginCredentialV1ParamsWithTimeout creates a new InteractiveLoginCredentialV1Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewInteractiveLoginCredentialV1ParamsWithTimeout(timeout time.Duration) *InteractiveLoginCredentialV1Params {
	var ()
	return &InteractiveLoginCredentialV1Params{

		timeout: timeout,
	}
}

// NewInteractiveLoginCredentialV1ParamsWithContext creates a new InteractiveLoginCredentialV1Params object
// with the default values initialized, and the ability to set a context for a request
func NewInteractiveLoginCredentialV1ParamsWithContext(ctx context.Context) *InteractiveLoginCredentialV1Params {
	var ()
	return &InteractiveLoginCredentialV1Params{

		Context: ctx,
	}
}

// NewInteractiveLoginCredentialV1ParamsWithHTTPClient creates a new InteractiveLoginCredentialV1Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewInteractiveLoginCredentialV1ParamsWithHTTPClient(client *http.Client) *InteractiveLoginCredentialV1Params {
	var ()
	return &InteractiveLoginCredentialV1Params{
		HTTPClient: client,
	}
}

/*InteractiveLoginCredentialV1Params contains all the parameters to send to the API endpoint
for the interactive login credential v1 operation typically these are written to a http.Request
*/
type InteractiveLoginCredentialV1Params struct {

	/*Body*/
	Body *model.CredentialV1Request

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the interactive login credential v1 params
func (o *InteractiveLoginCredentialV1Params) WithTimeout(timeout time.Duration) *InteractiveLoginCredentialV1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the interactive login credential v1 params
func (o *InteractiveLoginCredentialV1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the interactive login credential v1 params
func (o *InteractiveLoginCredentialV1Params) WithContext(ctx context.Context) *InteractiveLoginCredentialV1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the interactive login credential v1 params
func (o *InteractiveLoginCredentialV1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the interactive login credential v1 params
func (o *InteractiveLoginCredentialV1Params) WithHTTPClient(client *http.Client) *InteractiveLoginCredentialV1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the interactive login credential v1 params
func (o *InteractiveLoginCredentialV1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the interactive login credential v1 params
func (o *InteractiveLoginCredentialV1Params) WithBody(body *model.CredentialV1Request) *InteractiveLoginCredentialV1Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the interactive login credential v1 params
func (o *InteractiveLoginCredentialV1Params) SetBody(body *model.CredentialV1Request) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *InteractiveLoginCredentialV1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
