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

// NewUpdateDbEngineVersionParams creates a new UpdateDbEngineVersionParams object
// with the default values initialized.
func NewUpdateDbEngineVersionParams() *UpdateDbEngineVersionParams {
	var ()
	return &UpdateDbEngineVersionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateDbEngineVersionParamsWithTimeout creates a new UpdateDbEngineVersionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateDbEngineVersionParamsWithTimeout(timeout time.Duration) *UpdateDbEngineVersionParams {
	var ()
	return &UpdateDbEngineVersionParams{

		timeout: timeout,
	}
}

// NewUpdateDbEngineVersionParamsWithContext creates a new UpdateDbEngineVersionParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateDbEngineVersionParamsWithContext(ctx context.Context) *UpdateDbEngineVersionParams {
	var ()
	return &UpdateDbEngineVersionParams{

		Context: ctx,
	}
}

// NewUpdateDbEngineVersionParamsWithHTTPClient creates a new UpdateDbEngineVersionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateDbEngineVersionParamsWithHTTPClient(client *http.Client) *UpdateDbEngineVersionParams {
	var ()
	return &UpdateDbEngineVersionParams{
		HTTPClient: client,
	}
}

/*
UpdateDbEngineVersionParams contains all the parameters to send to the API endpoint
for the update db engine version operation typically these are written to a http.Request
*/
type UpdateDbEngineVersionParams struct {

	/*Crn*/
	Crn string
	/*DbEngineVersion*/
	DbEngineVersion *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update db engine version params
func (o *UpdateDbEngineVersionParams) WithTimeout(timeout time.Duration) *UpdateDbEngineVersionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update db engine version params
func (o *UpdateDbEngineVersionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update db engine version params
func (o *UpdateDbEngineVersionParams) WithContext(ctx context.Context) *UpdateDbEngineVersionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update db engine version params
func (o *UpdateDbEngineVersionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update db engine version params
func (o *UpdateDbEngineVersionParams) WithHTTPClient(client *http.Client) *UpdateDbEngineVersionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update db engine version params
func (o *UpdateDbEngineVersionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCrn adds the crn to the update db engine version params
func (o *UpdateDbEngineVersionParams) WithCrn(crn string) *UpdateDbEngineVersionParams {
	o.SetCrn(crn)
	return o
}

// SetCrn adds the crn to the update db engine version params
func (o *UpdateDbEngineVersionParams) SetCrn(crn string) {
	o.Crn = crn
}

// WithDbEngineVersion adds the dbEngineVersion to the update db engine version params
func (o *UpdateDbEngineVersionParams) WithDbEngineVersion(dbEngineVersion *string) *UpdateDbEngineVersionParams {
	o.SetDbEngineVersion(dbEngineVersion)
	return o
}

// SetDbEngineVersion adds the dbEngineVersion to the update db engine version params
func (o *UpdateDbEngineVersionParams) SetDbEngineVersion(dbEngineVersion *string) {
	o.DbEngineVersion = dbEngineVersion
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateDbEngineVersionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param crn
	if err := r.SetPathParam("crn", o.Crn); err != nil {
		return err
	}

	if o.DbEngineVersion != nil {

		// query param dbEngineVersion
		var qrDbEngineVersion string
		if o.DbEngineVersion != nil {
			qrDbEngineVersion = *o.DbEngineVersion
		}
		qDbEngineVersion := qrDbEngineVersion
		if qDbEngineVersion != "" {
			if err := r.SetQueryParam("dbEngineVersion", qDbEngineVersion); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
