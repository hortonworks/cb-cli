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

	model "github.com/hortonworks/cb-cli/dataplane/api/model"
)

// NewRotateAutoTLSCertificatesByCrnParams creates a new RotateAutoTLSCertificatesByCrnParams object
// with the default values initialized.
func NewRotateAutoTLSCertificatesByCrnParams() *RotateAutoTLSCertificatesByCrnParams {
	var ()
	return &RotateAutoTLSCertificatesByCrnParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRotateAutoTLSCertificatesByCrnParamsWithTimeout creates a new RotateAutoTLSCertificatesByCrnParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRotateAutoTLSCertificatesByCrnParamsWithTimeout(timeout time.Duration) *RotateAutoTLSCertificatesByCrnParams {
	var ()
	return &RotateAutoTLSCertificatesByCrnParams{

		timeout: timeout,
	}
}

// NewRotateAutoTLSCertificatesByCrnParamsWithContext creates a new RotateAutoTLSCertificatesByCrnParams object
// with the default values initialized, and the ability to set a context for a request
func NewRotateAutoTLSCertificatesByCrnParamsWithContext(ctx context.Context) *RotateAutoTLSCertificatesByCrnParams {
	var ()
	return &RotateAutoTLSCertificatesByCrnParams{

		Context: ctx,
	}
}

// NewRotateAutoTLSCertificatesByCrnParamsWithHTTPClient creates a new RotateAutoTLSCertificatesByCrnParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRotateAutoTLSCertificatesByCrnParamsWithHTTPClient(client *http.Client) *RotateAutoTLSCertificatesByCrnParams {
	var ()
	return &RotateAutoTLSCertificatesByCrnParams{
		HTTPClient: client,
	}
}

/*
RotateAutoTLSCertificatesByCrnParams contains all the parameters to send to the API endpoint
for the rotate auto Tls certificates by crn operation typically these are written to a http.Request
*/
type RotateAutoTLSCertificatesByCrnParams struct {

	/*Body*/
	Body *model.CertificatesRotationV4Request
	/*Crn*/
	Crn string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the rotate auto Tls certificates by crn params
func (o *RotateAutoTLSCertificatesByCrnParams) WithTimeout(timeout time.Duration) *RotateAutoTLSCertificatesByCrnParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the rotate auto Tls certificates by crn params
func (o *RotateAutoTLSCertificatesByCrnParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the rotate auto Tls certificates by crn params
func (o *RotateAutoTLSCertificatesByCrnParams) WithContext(ctx context.Context) *RotateAutoTLSCertificatesByCrnParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the rotate auto Tls certificates by crn params
func (o *RotateAutoTLSCertificatesByCrnParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the rotate auto Tls certificates by crn params
func (o *RotateAutoTLSCertificatesByCrnParams) WithHTTPClient(client *http.Client) *RotateAutoTLSCertificatesByCrnParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the rotate auto Tls certificates by crn params
func (o *RotateAutoTLSCertificatesByCrnParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the rotate auto Tls certificates by crn params
func (o *RotateAutoTLSCertificatesByCrnParams) WithBody(body *model.CertificatesRotationV4Request) *RotateAutoTLSCertificatesByCrnParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the rotate auto Tls certificates by crn params
func (o *RotateAutoTLSCertificatesByCrnParams) SetBody(body *model.CertificatesRotationV4Request) {
	o.Body = body
}

// WithCrn adds the crn to the rotate auto Tls certificates by crn params
func (o *RotateAutoTLSCertificatesByCrnParams) WithCrn(crn string) *RotateAutoTLSCertificatesByCrnParams {
	o.SetCrn(crn)
	return o
}

// SetCrn adds the crn to the rotate auto Tls certificates by crn params
func (o *RotateAutoTLSCertificatesByCrnParams) SetCrn(crn string) {
	o.Crn = crn
}

// WriteToRequest writes these params to a swagger request
func (o *RotateAutoTLSCertificatesByCrnParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param crn
	if err := r.SetPathParam("crn", o.Crn); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
