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
)

// NewInternalListFreeIpaClustersByAccountV1Params creates a new InternalListFreeIpaClustersByAccountV1Params object
// with the default values initialized.
func NewInternalListFreeIpaClustersByAccountV1Params() *InternalListFreeIpaClustersByAccountV1Params {
	var ()
	return &InternalListFreeIpaClustersByAccountV1Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewInternalListFreeIpaClustersByAccountV1ParamsWithTimeout creates a new InternalListFreeIpaClustersByAccountV1Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewInternalListFreeIpaClustersByAccountV1ParamsWithTimeout(timeout time.Duration) *InternalListFreeIpaClustersByAccountV1Params {
	var ()
	return &InternalListFreeIpaClustersByAccountV1Params{

		timeout: timeout,
	}
}

// NewInternalListFreeIpaClustersByAccountV1ParamsWithContext creates a new InternalListFreeIpaClustersByAccountV1Params object
// with the default values initialized, and the ability to set a context for a request
func NewInternalListFreeIpaClustersByAccountV1ParamsWithContext(ctx context.Context) *InternalListFreeIpaClustersByAccountV1Params {
	var ()
	return &InternalListFreeIpaClustersByAccountV1Params{

		Context: ctx,
	}
}

// NewInternalListFreeIpaClustersByAccountV1ParamsWithHTTPClient creates a new InternalListFreeIpaClustersByAccountV1Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewInternalListFreeIpaClustersByAccountV1ParamsWithHTTPClient(client *http.Client) *InternalListFreeIpaClustersByAccountV1Params {
	var ()
	return &InternalListFreeIpaClustersByAccountV1Params{
		HTTPClient: client,
	}
}

/*
InternalListFreeIpaClustersByAccountV1Params contains all the parameters to send to the API endpoint
for the internal list free ipa clusters by account v1 operation typically these are written to a http.Request
*/
type InternalListFreeIpaClustersByAccountV1Params struct {

	/*AccountID*/
	AccountID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the internal list free ipa clusters by account v1 params
func (o *InternalListFreeIpaClustersByAccountV1Params) WithTimeout(timeout time.Duration) *InternalListFreeIpaClustersByAccountV1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the internal list free ipa clusters by account v1 params
func (o *InternalListFreeIpaClustersByAccountV1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the internal list free ipa clusters by account v1 params
func (o *InternalListFreeIpaClustersByAccountV1Params) WithContext(ctx context.Context) *InternalListFreeIpaClustersByAccountV1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the internal list free ipa clusters by account v1 params
func (o *InternalListFreeIpaClustersByAccountV1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the internal list free ipa clusters by account v1 params
func (o *InternalListFreeIpaClustersByAccountV1Params) WithHTTPClient(client *http.Client) *InternalListFreeIpaClustersByAccountV1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the internal list free ipa clusters by account v1 params
func (o *InternalListFreeIpaClustersByAccountV1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountID adds the accountID to the internal list free ipa clusters by account v1 params
func (o *InternalListFreeIpaClustersByAccountV1Params) WithAccountID(accountID *string) *InternalListFreeIpaClustersByAccountV1Params {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the internal list free ipa clusters by account v1 params
func (o *InternalListFreeIpaClustersByAccountV1Params) SetAccountID(accountID *string) {
	o.AccountID = accountID
}

// WriteToRequest writes these params to a swagger request
func (o *InternalListFreeIpaClustersByAccountV1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
