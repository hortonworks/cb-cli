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

// NewSyncDistroXV1ByCrnParams creates a new SyncDistroXV1ByCrnParams object
// with the default values initialized.
func NewSyncDistroXV1ByCrnParams() *SyncDistroXV1ByCrnParams {
	var ()
	return &SyncDistroXV1ByCrnParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSyncDistroXV1ByCrnParamsWithTimeout creates a new SyncDistroXV1ByCrnParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSyncDistroXV1ByCrnParamsWithTimeout(timeout time.Duration) *SyncDistroXV1ByCrnParams {
	var ()
	return &SyncDistroXV1ByCrnParams{

		timeout: timeout,
	}
}

// NewSyncDistroXV1ByCrnParamsWithContext creates a new SyncDistroXV1ByCrnParams object
// with the default values initialized, and the ability to set a context for a request
func NewSyncDistroXV1ByCrnParamsWithContext(ctx context.Context) *SyncDistroXV1ByCrnParams {
	var ()
	return &SyncDistroXV1ByCrnParams{

		Context: ctx,
	}
}

// NewSyncDistroXV1ByCrnParamsWithHTTPClient creates a new SyncDistroXV1ByCrnParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSyncDistroXV1ByCrnParamsWithHTTPClient(client *http.Client) *SyncDistroXV1ByCrnParams {
	var ()
	return &SyncDistroXV1ByCrnParams{
		HTTPClient: client,
	}
}

/*
SyncDistroXV1ByCrnParams contains all the parameters to send to the API endpoint
for the sync distro x v1 by crn operation typically these are written to a http.Request
*/
type SyncDistroXV1ByCrnParams struct {

	/*Crn*/
	Crn string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the sync distro x v1 by crn params
func (o *SyncDistroXV1ByCrnParams) WithTimeout(timeout time.Duration) *SyncDistroXV1ByCrnParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the sync distro x v1 by crn params
func (o *SyncDistroXV1ByCrnParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the sync distro x v1 by crn params
func (o *SyncDistroXV1ByCrnParams) WithContext(ctx context.Context) *SyncDistroXV1ByCrnParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the sync distro x v1 by crn params
func (o *SyncDistroXV1ByCrnParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the sync distro x v1 by crn params
func (o *SyncDistroXV1ByCrnParams) WithHTTPClient(client *http.Client) *SyncDistroXV1ByCrnParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the sync distro x v1 by crn params
func (o *SyncDistroXV1ByCrnParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCrn adds the crn to the sync distro x v1 by crn params
func (o *SyncDistroXV1ByCrnParams) WithCrn(crn string) *SyncDistroXV1ByCrnParams {
	o.SetCrn(crn)
	return o
}

// SetCrn adds the crn to the sync distro x v1 by crn params
func (o *SyncDistroXV1ByCrnParams) SetCrn(crn string) {
	o.Crn = crn
}

// WriteToRequest writes these params to a swagger request
func (o *SyncDistroXV1ByCrnParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
