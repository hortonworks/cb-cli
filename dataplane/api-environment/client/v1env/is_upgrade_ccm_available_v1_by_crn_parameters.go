// Code generated by go-swagger; DO NOT EDIT.

package v1env

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

// NewIsUpgradeCcmAvailableV1ByCrnParams creates a new IsUpgradeCcmAvailableV1ByCrnParams object
// with the default values initialized.
func NewIsUpgradeCcmAvailableV1ByCrnParams() *IsUpgradeCcmAvailableV1ByCrnParams {
	var ()
	return &IsUpgradeCcmAvailableV1ByCrnParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewIsUpgradeCcmAvailableV1ByCrnParamsWithTimeout creates a new IsUpgradeCcmAvailableV1ByCrnParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewIsUpgradeCcmAvailableV1ByCrnParamsWithTimeout(timeout time.Duration) *IsUpgradeCcmAvailableV1ByCrnParams {
	var ()
	return &IsUpgradeCcmAvailableV1ByCrnParams{

		timeout: timeout,
	}
}

// NewIsUpgradeCcmAvailableV1ByCrnParamsWithContext creates a new IsUpgradeCcmAvailableV1ByCrnParams object
// with the default values initialized, and the ability to set a context for a request
func NewIsUpgradeCcmAvailableV1ByCrnParamsWithContext(ctx context.Context) *IsUpgradeCcmAvailableV1ByCrnParams {
	var ()
	return &IsUpgradeCcmAvailableV1ByCrnParams{

		Context: ctx,
	}
}

// NewIsUpgradeCcmAvailableV1ByCrnParamsWithHTTPClient creates a new IsUpgradeCcmAvailableV1ByCrnParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewIsUpgradeCcmAvailableV1ByCrnParamsWithHTTPClient(client *http.Client) *IsUpgradeCcmAvailableV1ByCrnParams {
	var ()
	return &IsUpgradeCcmAvailableV1ByCrnParams{
		HTTPClient: client,
	}
}

/*
IsUpgradeCcmAvailableV1ByCrnParams contains all the parameters to send to the API endpoint
for the is upgrade ccm available v1 by crn operation typically these are written to a http.Request
*/
type IsUpgradeCcmAvailableV1ByCrnParams struct {

	/*Crn*/
	Crn string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the is upgrade ccm available v1 by crn params
func (o *IsUpgradeCcmAvailableV1ByCrnParams) WithTimeout(timeout time.Duration) *IsUpgradeCcmAvailableV1ByCrnParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the is upgrade ccm available v1 by crn params
func (o *IsUpgradeCcmAvailableV1ByCrnParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the is upgrade ccm available v1 by crn params
func (o *IsUpgradeCcmAvailableV1ByCrnParams) WithContext(ctx context.Context) *IsUpgradeCcmAvailableV1ByCrnParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the is upgrade ccm available v1 by crn params
func (o *IsUpgradeCcmAvailableV1ByCrnParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the is upgrade ccm available v1 by crn params
func (o *IsUpgradeCcmAvailableV1ByCrnParams) WithHTTPClient(client *http.Client) *IsUpgradeCcmAvailableV1ByCrnParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the is upgrade ccm available v1 by crn params
func (o *IsUpgradeCcmAvailableV1ByCrnParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCrn adds the crn to the is upgrade ccm available v1 by crn params
func (o *IsUpgradeCcmAvailableV1ByCrnParams) WithCrn(crn string) *IsUpgradeCcmAvailableV1ByCrnParams {
	o.SetCrn(crn)
	return o
}

// SetCrn adds the crn to the is upgrade ccm available v1 by crn params
func (o *IsUpgradeCcmAvailableV1ByCrnParams) SetCrn(crn string) {
	o.Crn = crn
}

// WriteToRequest writes these params to a swagger request
func (o *IsUpgradeCcmAvailableV1ByCrnParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param crn
	if err := r.SetPathParam("crn", o.Crn); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
