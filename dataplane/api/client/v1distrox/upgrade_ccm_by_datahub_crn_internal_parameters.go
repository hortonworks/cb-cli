// Code generated by go-swagger; DO NOT EDIT.

package v1distrox

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

// NewUpgradeCcmByDatahubCrnInternalParams creates a new UpgradeCcmByDatahubCrnInternalParams object
// with the default values initialized.
func NewUpgradeCcmByDatahubCrnInternalParams() *UpgradeCcmByDatahubCrnInternalParams {
	var ()
	return &UpgradeCcmByDatahubCrnInternalParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpgradeCcmByDatahubCrnInternalParamsWithTimeout creates a new UpgradeCcmByDatahubCrnInternalParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpgradeCcmByDatahubCrnInternalParamsWithTimeout(timeout time.Duration) *UpgradeCcmByDatahubCrnInternalParams {
	var ()
	return &UpgradeCcmByDatahubCrnInternalParams{

		timeout: timeout,
	}
}

// NewUpgradeCcmByDatahubCrnInternalParamsWithContext creates a new UpgradeCcmByDatahubCrnInternalParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpgradeCcmByDatahubCrnInternalParamsWithContext(ctx context.Context) *UpgradeCcmByDatahubCrnInternalParams {
	var ()
	return &UpgradeCcmByDatahubCrnInternalParams{

		Context: ctx,
	}
}

// NewUpgradeCcmByDatahubCrnInternalParamsWithHTTPClient creates a new UpgradeCcmByDatahubCrnInternalParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpgradeCcmByDatahubCrnInternalParamsWithHTTPClient(client *http.Client) *UpgradeCcmByDatahubCrnInternalParams {
	var ()
	return &UpgradeCcmByDatahubCrnInternalParams{
		HTTPClient: client,
	}
}

/*
UpgradeCcmByDatahubCrnInternalParams contains all the parameters to send to the API endpoint
for the upgrade ccm by datahub crn internal operation typically these are written to a http.Request
*/
type UpgradeCcmByDatahubCrnInternalParams struct {

	/*Crn*/
	Crn string
	/*InitiatorUserCrn*/
	InitiatorUserCrn *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the upgrade ccm by datahub crn internal params
func (o *UpgradeCcmByDatahubCrnInternalParams) WithTimeout(timeout time.Duration) *UpgradeCcmByDatahubCrnInternalParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the upgrade ccm by datahub crn internal params
func (o *UpgradeCcmByDatahubCrnInternalParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the upgrade ccm by datahub crn internal params
func (o *UpgradeCcmByDatahubCrnInternalParams) WithContext(ctx context.Context) *UpgradeCcmByDatahubCrnInternalParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the upgrade ccm by datahub crn internal params
func (o *UpgradeCcmByDatahubCrnInternalParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the upgrade ccm by datahub crn internal params
func (o *UpgradeCcmByDatahubCrnInternalParams) WithHTTPClient(client *http.Client) *UpgradeCcmByDatahubCrnInternalParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the upgrade ccm by datahub crn internal params
func (o *UpgradeCcmByDatahubCrnInternalParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCrn adds the crn to the upgrade ccm by datahub crn internal params
func (o *UpgradeCcmByDatahubCrnInternalParams) WithCrn(crn string) *UpgradeCcmByDatahubCrnInternalParams {
	o.SetCrn(crn)
	return o
}

// SetCrn adds the crn to the upgrade ccm by datahub crn internal params
func (o *UpgradeCcmByDatahubCrnInternalParams) SetCrn(crn string) {
	o.Crn = crn
}

// WithInitiatorUserCrn adds the initiatorUserCrn to the upgrade ccm by datahub crn internal params
func (o *UpgradeCcmByDatahubCrnInternalParams) WithInitiatorUserCrn(initiatorUserCrn *string) *UpgradeCcmByDatahubCrnInternalParams {
	o.SetInitiatorUserCrn(initiatorUserCrn)
	return o
}

// SetInitiatorUserCrn adds the initiatorUserCrn to the upgrade ccm by datahub crn internal params
func (o *UpgradeCcmByDatahubCrnInternalParams) SetInitiatorUserCrn(initiatorUserCrn *string) {
	o.InitiatorUserCrn = initiatorUserCrn
}

// WriteToRequest writes these params to a swagger request
func (o *UpgradeCcmByDatahubCrnInternalParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
