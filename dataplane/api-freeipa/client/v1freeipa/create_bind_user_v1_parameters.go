// Code generated by go-swagger; DO NOT EDIT.

package v1freeipa

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

// NewCreateBindUserV1Params creates a new CreateBindUserV1Params object
// with the default values initialized.
func NewCreateBindUserV1Params() *CreateBindUserV1Params {
	var ()
	return &CreateBindUserV1Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateBindUserV1ParamsWithTimeout creates a new CreateBindUserV1Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateBindUserV1ParamsWithTimeout(timeout time.Duration) *CreateBindUserV1Params {
	var ()
	return &CreateBindUserV1Params{

		timeout: timeout,
	}
}

// NewCreateBindUserV1ParamsWithContext creates a new CreateBindUserV1Params object
// with the default values initialized, and the ability to set a context for a request
func NewCreateBindUserV1ParamsWithContext(ctx context.Context) *CreateBindUserV1Params {
	var ()
	return &CreateBindUserV1Params{

		Context: ctx,
	}
}

// NewCreateBindUserV1ParamsWithHTTPClient creates a new CreateBindUserV1Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateBindUserV1ParamsWithHTTPClient(client *http.Client) *CreateBindUserV1Params {
	var ()
	return &CreateBindUserV1Params{
		HTTPClient: client,
	}
}

/*
CreateBindUserV1Params contains all the parameters to send to the API endpoint
for the create bind user v1 operation typically these are written to a http.Request
*/
type CreateBindUserV1Params struct {

	/*Body*/
	Body *model.BindUserCreateV1Request
	/*InitiatorUserCrn*/
	InitiatorUserCrn *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create bind user v1 params
func (o *CreateBindUserV1Params) WithTimeout(timeout time.Duration) *CreateBindUserV1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create bind user v1 params
func (o *CreateBindUserV1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create bind user v1 params
func (o *CreateBindUserV1Params) WithContext(ctx context.Context) *CreateBindUserV1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create bind user v1 params
func (o *CreateBindUserV1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create bind user v1 params
func (o *CreateBindUserV1Params) WithHTTPClient(client *http.Client) *CreateBindUserV1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create bind user v1 params
func (o *CreateBindUserV1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create bind user v1 params
func (o *CreateBindUserV1Params) WithBody(body *model.BindUserCreateV1Request) *CreateBindUserV1Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create bind user v1 params
func (o *CreateBindUserV1Params) SetBody(body *model.BindUserCreateV1Request) {
	o.Body = body
}

// WithInitiatorUserCrn adds the initiatorUserCrn to the create bind user v1 params
func (o *CreateBindUserV1Params) WithInitiatorUserCrn(initiatorUserCrn *string) *CreateBindUserV1Params {
	o.SetInitiatorUserCrn(initiatorUserCrn)
	return o
}

// SetInitiatorUserCrn adds the initiatorUserCrn to the create bind user v1 params
func (o *CreateBindUserV1Params) SetInitiatorUserCrn(initiatorUserCrn *string) {
	o.InitiatorUserCrn = initiatorUserCrn
}

// WriteToRequest writes these params to a swagger request
func (o *CreateBindUserV1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if o.InitiatorUserCrn != nil {

		// query param initiatorUserCrn
		var qrInitiatorUserCrn string
		if o.InitiatorUserCrn != nil {
			qrInitiatorUserCrn = *o.InitiatorUserCrn
		}
		qInitiatorUserCrn := qrInitiatorUserCrn
		if qInitiatorUserCrn != "" {
			if err := r.SetQueryParam("initiatorUserCrn", qInitiatorUserCrn); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
