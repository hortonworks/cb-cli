// Code generated by go-swagger; DO NOT EDIT.

package sdx

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

// NewGetDatalakeRestoreIDParams creates a new GetDatalakeRestoreIDParams object
// with the default values initialized.
func NewGetDatalakeRestoreIDParams() *GetDatalakeRestoreIDParams {
	var ()
	return &GetDatalakeRestoreIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDatalakeRestoreIDParamsWithTimeout creates a new GetDatalakeRestoreIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDatalakeRestoreIDParamsWithTimeout(timeout time.Duration) *GetDatalakeRestoreIDParams {
	var ()
	return &GetDatalakeRestoreIDParams{

		timeout: timeout,
	}
}

// NewGetDatalakeRestoreIDParamsWithContext creates a new GetDatalakeRestoreIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDatalakeRestoreIDParamsWithContext(ctx context.Context) *GetDatalakeRestoreIDParams {
	var ()
	return &GetDatalakeRestoreIDParams{

		Context: ctx,
	}
}

// NewGetDatalakeRestoreIDParamsWithHTTPClient creates a new GetDatalakeRestoreIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDatalakeRestoreIDParamsWithHTTPClient(client *http.Client) *GetDatalakeRestoreIDParams {
	var ()
	return &GetDatalakeRestoreIDParams{
		HTTPClient: client,
	}
}

/*
GetDatalakeRestoreIDParams contains all the parameters to send to the API endpoint
for the get datalake restore Id operation typically these are written to a http.Request
*/
type GetDatalakeRestoreIDParams struct {

	/*BackupName
	  optional: datalake backup name

	*/
	BackupName *string
	/*Name
	  required: datalake name

	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get datalake restore Id params
func (o *GetDatalakeRestoreIDParams) WithTimeout(timeout time.Duration) *GetDatalakeRestoreIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get datalake restore Id params
func (o *GetDatalakeRestoreIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get datalake restore Id params
func (o *GetDatalakeRestoreIDParams) WithContext(ctx context.Context) *GetDatalakeRestoreIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get datalake restore Id params
func (o *GetDatalakeRestoreIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get datalake restore Id params
func (o *GetDatalakeRestoreIDParams) WithHTTPClient(client *http.Client) *GetDatalakeRestoreIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get datalake restore Id params
func (o *GetDatalakeRestoreIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBackupName adds the backupName to the get datalake restore Id params
func (o *GetDatalakeRestoreIDParams) WithBackupName(backupName *string) *GetDatalakeRestoreIDParams {
	o.SetBackupName(backupName)
	return o
}

// SetBackupName adds the backupName to the get datalake restore Id params
func (o *GetDatalakeRestoreIDParams) SetBackupName(backupName *string) {
	o.BackupName = backupName
}

// WithName adds the name to the get datalake restore Id params
func (o *GetDatalakeRestoreIDParams) WithName(name string) *GetDatalakeRestoreIDParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get datalake restore Id params
func (o *GetDatalakeRestoreIDParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *GetDatalakeRestoreIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.BackupName != nil {

		// query param backupName
		var qrBackupName string
		if o.BackupName != nil {
			qrBackupName = *o.BackupName
		}
		qBackupName := qrBackupName
		if qBackupName != "" {
			if err := r.SetQueryParam("backupName", qBackupName); err != nil {
				return err
			}
		}

	}

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
