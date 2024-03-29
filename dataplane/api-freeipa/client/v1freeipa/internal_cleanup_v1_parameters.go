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

// NewInternalCleanupV1Params creates a new InternalCleanupV1Params object
// with the default values initialized.
func NewInternalCleanupV1Params() *InternalCleanupV1Params {
	var ()
	return &InternalCleanupV1Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewInternalCleanupV1ParamsWithTimeout creates a new InternalCleanupV1Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewInternalCleanupV1ParamsWithTimeout(timeout time.Duration) *InternalCleanupV1Params {
	var ()
	return &InternalCleanupV1Params{

		timeout: timeout,
	}
}

// NewInternalCleanupV1ParamsWithContext creates a new InternalCleanupV1Params object
// with the default values initialized, and the ability to set a context for a request
func NewInternalCleanupV1ParamsWithContext(ctx context.Context) *InternalCleanupV1Params {
	var ()
	return &InternalCleanupV1Params{

		Context: ctx,
	}
}

// NewInternalCleanupV1ParamsWithHTTPClient creates a new InternalCleanupV1Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewInternalCleanupV1ParamsWithHTTPClient(client *http.Client) *InternalCleanupV1Params {
	var ()
	return &InternalCleanupV1Params{
		HTTPClient: client,
	}
}

/*
InternalCleanupV1Params contains all the parameters to send to the API endpoint
for the internal cleanup v1 operation typically these are written to a http.Request
*/
type InternalCleanupV1Params struct {

	/*AccountID*/
	AccountID *string
	/*Body*/
	Body *model.CleanupV1Request

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the internal cleanup v1 params
func (o *InternalCleanupV1Params) WithTimeout(timeout time.Duration) *InternalCleanupV1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the internal cleanup v1 params
func (o *InternalCleanupV1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the internal cleanup v1 params
func (o *InternalCleanupV1Params) WithContext(ctx context.Context) *InternalCleanupV1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the internal cleanup v1 params
func (o *InternalCleanupV1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the internal cleanup v1 params
func (o *InternalCleanupV1Params) WithHTTPClient(client *http.Client) *InternalCleanupV1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the internal cleanup v1 params
func (o *InternalCleanupV1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountID adds the accountID to the internal cleanup v1 params
func (o *InternalCleanupV1Params) WithAccountID(accountID *string) *InternalCleanupV1Params {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the internal cleanup v1 params
func (o *InternalCleanupV1Params) SetAccountID(accountID *string) {
	o.AccountID = accountID
}

// WithBody adds the body to the internal cleanup v1 params
func (o *InternalCleanupV1Params) WithBody(body *model.CleanupV1Request) *InternalCleanupV1Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the internal cleanup v1 params
func (o *InternalCleanupV1Params) SetBody(body *model.CleanupV1Request) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *InternalCleanupV1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AccountID != nil {

		// query param accountId
		var qrAccountID string
		if o.AccountID != nil {
			qrAccountID = *o.AccountID
		}
		qAccountID := qrAccountID
		if qAccountID != "" {
			if err := r.SetQueryParam("accountId", qAccountID); err != nil {
				return err
			}
		}

	}

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
