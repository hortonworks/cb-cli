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
)

// NewGetAuditCredentialByResourceCrnV1Params creates a new GetAuditCredentialByResourceCrnV1Params object
// with the default values initialized.
func NewGetAuditCredentialByResourceCrnV1Params() *GetAuditCredentialByResourceCrnV1Params {
	var ()
	return &GetAuditCredentialByResourceCrnV1Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAuditCredentialByResourceCrnV1ParamsWithTimeout creates a new GetAuditCredentialByResourceCrnV1Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAuditCredentialByResourceCrnV1ParamsWithTimeout(timeout time.Duration) *GetAuditCredentialByResourceCrnV1Params {
	var ()
	return &GetAuditCredentialByResourceCrnV1Params{

		timeout: timeout,
	}
}

// NewGetAuditCredentialByResourceCrnV1ParamsWithContext creates a new GetAuditCredentialByResourceCrnV1Params object
// with the default values initialized, and the ability to set a context for a request
func NewGetAuditCredentialByResourceCrnV1ParamsWithContext(ctx context.Context) *GetAuditCredentialByResourceCrnV1Params {
	var ()
	return &GetAuditCredentialByResourceCrnV1Params{

		Context: ctx,
	}
}

// NewGetAuditCredentialByResourceCrnV1ParamsWithHTTPClient creates a new GetAuditCredentialByResourceCrnV1Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAuditCredentialByResourceCrnV1ParamsWithHTTPClient(client *http.Client) *GetAuditCredentialByResourceCrnV1Params {
	var ()
	return &GetAuditCredentialByResourceCrnV1Params{
		HTTPClient: client,
	}
}

/*
GetAuditCredentialByResourceCrnV1Params contains all the parameters to send to the API endpoint
for the get audit credential by resource crn v1 operation typically these are written to a http.Request
*/
type GetAuditCredentialByResourceCrnV1Params struct {

	/*Crn*/
	Crn string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get audit credential by resource crn v1 params
func (o *GetAuditCredentialByResourceCrnV1Params) WithTimeout(timeout time.Duration) *GetAuditCredentialByResourceCrnV1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get audit credential by resource crn v1 params
func (o *GetAuditCredentialByResourceCrnV1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get audit credential by resource crn v1 params
func (o *GetAuditCredentialByResourceCrnV1Params) WithContext(ctx context.Context) *GetAuditCredentialByResourceCrnV1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get audit credential by resource crn v1 params
func (o *GetAuditCredentialByResourceCrnV1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get audit credential by resource crn v1 params
func (o *GetAuditCredentialByResourceCrnV1Params) WithHTTPClient(client *http.Client) *GetAuditCredentialByResourceCrnV1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get audit credential by resource crn v1 params
func (o *GetAuditCredentialByResourceCrnV1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCrn adds the crn to the get audit credential by resource crn v1 params
func (o *GetAuditCredentialByResourceCrnV1Params) WithCrn(crn string) *GetAuditCredentialByResourceCrnV1Params {
	o.SetCrn(crn)
	return o
}

// SetCrn adds the crn to the get audit credential by resource crn v1 params
func (o *GetAuditCredentialByResourceCrnV1Params) SetCrn(crn string) {
	o.Crn = crn
}

// WriteToRequest writes these params to a swagger request
func (o *GetAuditCredentialByResourceCrnV1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
