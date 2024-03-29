// Code generated by go-swagger; DO NOT EDIT.

package v1credentialsaudit

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

// NewCreateAuditCredentialV1Params creates a new CreateAuditCredentialV1Params object
// with the default values initialized.
func NewCreateAuditCredentialV1Params() *CreateAuditCredentialV1Params {
	var ()
	return &CreateAuditCredentialV1Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateAuditCredentialV1ParamsWithTimeout creates a new CreateAuditCredentialV1Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateAuditCredentialV1ParamsWithTimeout(timeout time.Duration) *CreateAuditCredentialV1Params {
	var ()
	return &CreateAuditCredentialV1Params{

		timeout: timeout,
	}
}

// NewCreateAuditCredentialV1ParamsWithContext creates a new CreateAuditCredentialV1Params object
// with the default values initialized, and the ability to set a context for a request
func NewCreateAuditCredentialV1ParamsWithContext(ctx context.Context) *CreateAuditCredentialV1Params {
	var ()
	return &CreateAuditCredentialV1Params{

		Context: ctx,
	}
}

// NewCreateAuditCredentialV1ParamsWithHTTPClient creates a new CreateAuditCredentialV1Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateAuditCredentialV1ParamsWithHTTPClient(client *http.Client) *CreateAuditCredentialV1Params {
	var ()
	return &CreateAuditCredentialV1Params{
		HTTPClient: client,
	}
}

/*
CreateAuditCredentialV1Params contains all the parameters to send to the API endpoint
for the create audit credential v1 operation typically these are written to a http.Request
*/
type CreateAuditCredentialV1Params struct {

	/*Body*/
	Body *model.CredentialV1Request

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create audit credential v1 params
func (o *CreateAuditCredentialV1Params) WithTimeout(timeout time.Duration) *CreateAuditCredentialV1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create audit credential v1 params
func (o *CreateAuditCredentialV1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create audit credential v1 params
func (o *CreateAuditCredentialV1Params) WithContext(ctx context.Context) *CreateAuditCredentialV1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create audit credential v1 params
func (o *CreateAuditCredentialV1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create audit credential v1 params
func (o *CreateAuditCredentialV1Params) WithHTTPClient(client *http.Client) *CreateAuditCredentialV1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create audit credential v1 params
func (o *CreateAuditCredentialV1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create audit credential v1 params
func (o *CreateAuditCredentialV1Params) WithBody(body *model.CredentialV1Request) *CreateAuditCredentialV1Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create audit credential v1 params
func (o *CreateAuditCredentialV1Params) SetBody(body *model.CredentialV1Request) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateAuditCredentialV1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
