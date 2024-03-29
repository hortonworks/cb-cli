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

// NewRenewDistroXCertificateParams creates a new RenewDistroXCertificateParams object
// with the default values initialized.
func NewRenewDistroXCertificateParams() *RenewDistroXCertificateParams {
	var ()
	return &RenewDistroXCertificateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRenewDistroXCertificateParamsWithTimeout creates a new RenewDistroXCertificateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRenewDistroXCertificateParamsWithTimeout(timeout time.Duration) *RenewDistroXCertificateParams {
	var ()
	return &RenewDistroXCertificateParams{

		timeout: timeout,
	}
}

// NewRenewDistroXCertificateParamsWithContext creates a new RenewDistroXCertificateParams object
// with the default values initialized, and the ability to set a context for a request
func NewRenewDistroXCertificateParamsWithContext(ctx context.Context) *RenewDistroXCertificateParams {
	var ()
	return &RenewDistroXCertificateParams{

		Context: ctx,
	}
}

// NewRenewDistroXCertificateParamsWithHTTPClient creates a new RenewDistroXCertificateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRenewDistroXCertificateParamsWithHTTPClient(client *http.Client) *RenewDistroXCertificateParams {
	var ()
	return &RenewDistroXCertificateParams{
		HTTPClient: client,
	}
}

/*
RenewDistroXCertificateParams contains all the parameters to send to the API endpoint
for the renew distro x certificate operation typically these are written to a http.Request
*/
type RenewDistroXCertificateParams struct {

	/*Crn*/
	Crn string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the renew distro x certificate params
func (o *RenewDistroXCertificateParams) WithTimeout(timeout time.Duration) *RenewDistroXCertificateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the renew distro x certificate params
func (o *RenewDistroXCertificateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the renew distro x certificate params
func (o *RenewDistroXCertificateParams) WithContext(ctx context.Context) *RenewDistroXCertificateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the renew distro x certificate params
func (o *RenewDistroXCertificateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the renew distro x certificate params
func (o *RenewDistroXCertificateParams) WithHTTPClient(client *http.Client) *RenewDistroXCertificateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the renew distro x certificate params
func (o *RenewDistroXCertificateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCrn adds the crn to the renew distro x certificate params
func (o *RenewDistroXCertificateParams) WithCrn(crn string) *RenewDistroXCertificateParams {
	o.SetCrn(crn)
	return o
}

// SetCrn adds the crn to the renew distro x certificate params
func (o *RenewDistroXCertificateParams) SetCrn(crn string) {
	o.Crn = crn
}

// WriteToRequest writes these params to a swagger request
func (o *RenewDistroXCertificateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
