// Code generated by go-swagger; DO NOT EDIT.

package internalsdx

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

// NewModifyInternalSdxProxyConfigParams creates a new ModifyInternalSdxProxyConfigParams object
// with the default values initialized.
func NewModifyInternalSdxProxyConfigParams() *ModifyInternalSdxProxyConfigParams {
	var ()
	return &ModifyInternalSdxProxyConfigParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewModifyInternalSdxProxyConfigParamsWithTimeout creates a new ModifyInternalSdxProxyConfigParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewModifyInternalSdxProxyConfigParamsWithTimeout(timeout time.Duration) *ModifyInternalSdxProxyConfigParams {
	var ()
	return &ModifyInternalSdxProxyConfigParams{

		timeout: timeout,
	}
}

// NewModifyInternalSdxProxyConfigParamsWithContext creates a new ModifyInternalSdxProxyConfigParams object
// with the default values initialized, and the ability to set a context for a request
func NewModifyInternalSdxProxyConfigParamsWithContext(ctx context.Context) *ModifyInternalSdxProxyConfigParams {
	var ()
	return &ModifyInternalSdxProxyConfigParams{

		Context: ctx,
	}
}

// NewModifyInternalSdxProxyConfigParamsWithHTTPClient creates a new ModifyInternalSdxProxyConfigParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewModifyInternalSdxProxyConfigParamsWithHTTPClient(client *http.Client) *ModifyInternalSdxProxyConfigParams {
	var ()
	return &ModifyInternalSdxProxyConfigParams{
		HTTPClient: client,
	}
}

/*ModifyInternalSdxProxyConfigParams contains all the parameters to send to the API endpoint
for the modify internal sdx proxy config operation typically these are written to a http.Request
*/
type ModifyInternalSdxProxyConfigParams struct {

	/*Crn*/
	Crn string
	/*InitiatorUserCrn*/
	InitiatorUserCrn *string
	/*PreviousProxy*/
	PreviousProxy *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the modify internal sdx proxy config params
func (o *ModifyInternalSdxProxyConfigParams) WithTimeout(timeout time.Duration) *ModifyInternalSdxProxyConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the modify internal sdx proxy config params
func (o *ModifyInternalSdxProxyConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the modify internal sdx proxy config params
func (o *ModifyInternalSdxProxyConfigParams) WithContext(ctx context.Context) *ModifyInternalSdxProxyConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the modify internal sdx proxy config params
func (o *ModifyInternalSdxProxyConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the modify internal sdx proxy config params
func (o *ModifyInternalSdxProxyConfigParams) WithHTTPClient(client *http.Client) *ModifyInternalSdxProxyConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the modify internal sdx proxy config params
func (o *ModifyInternalSdxProxyConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCrn adds the crn to the modify internal sdx proxy config params
func (o *ModifyInternalSdxProxyConfigParams) WithCrn(crn string) *ModifyInternalSdxProxyConfigParams {
	o.SetCrn(crn)
	return o
}

// SetCrn adds the crn to the modify internal sdx proxy config params
func (o *ModifyInternalSdxProxyConfigParams) SetCrn(crn string) {
	o.Crn = crn
}

// WithInitiatorUserCrn adds the initiatorUserCrn to the modify internal sdx proxy config params
func (o *ModifyInternalSdxProxyConfigParams) WithInitiatorUserCrn(initiatorUserCrn *string) *ModifyInternalSdxProxyConfigParams {
	o.SetInitiatorUserCrn(initiatorUserCrn)
	return o
}

// SetInitiatorUserCrn adds the initiatorUserCrn to the modify internal sdx proxy config params
func (o *ModifyInternalSdxProxyConfigParams) SetInitiatorUserCrn(initiatorUserCrn *string) {
	o.InitiatorUserCrn = initiatorUserCrn
}

// WithPreviousProxy adds the previousProxy to the modify internal sdx proxy config params
func (o *ModifyInternalSdxProxyConfigParams) WithPreviousProxy(previousProxy *string) *ModifyInternalSdxProxyConfigParams {
	o.SetPreviousProxy(previousProxy)
	return o
}

// SetPreviousProxy adds the previousProxy to the modify internal sdx proxy config params
func (o *ModifyInternalSdxProxyConfigParams) SetPreviousProxy(previousProxy *string) {
	o.PreviousProxy = previousProxy
}

// WriteToRequest writes these params to a swagger request
func (o *ModifyInternalSdxProxyConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param crn
	if err := r.SetPathParam("crn", o.Crn); err != nil {
		return err
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

	if o.PreviousProxy != nil {

		// query param previousProxy
		var qrPreviousProxy string
		if o.PreviousProxy != nil {
			qrPreviousProxy = *o.PreviousProxy
		}
		qPreviousProxy := qrPreviousProxy
		if qPreviousProxy != "" {
			if err := r.SetQueryParam("previousProxy", qPreviousProxy); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
