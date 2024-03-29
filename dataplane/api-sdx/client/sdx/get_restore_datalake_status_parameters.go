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

// NewGetRestoreDatalakeStatusParams creates a new GetRestoreDatalakeStatusParams object
// with the default values initialized.
func NewGetRestoreDatalakeStatusParams() *GetRestoreDatalakeStatusParams {
	var ()
	return &GetRestoreDatalakeStatusParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRestoreDatalakeStatusParamsWithTimeout creates a new GetRestoreDatalakeStatusParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRestoreDatalakeStatusParamsWithTimeout(timeout time.Duration) *GetRestoreDatalakeStatusParams {
	var ()
	return &GetRestoreDatalakeStatusParams{

		timeout: timeout,
	}
}

// NewGetRestoreDatalakeStatusParamsWithContext creates a new GetRestoreDatalakeStatusParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRestoreDatalakeStatusParamsWithContext(ctx context.Context) *GetRestoreDatalakeStatusParams {
	var ()
	return &GetRestoreDatalakeStatusParams{

		Context: ctx,
	}
}

// NewGetRestoreDatalakeStatusParamsWithHTTPClient creates a new GetRestoreDatalakeStatusParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRestoreDatalakeStatusParamsWithHTTPClient(client *http.Client) *GetRestoreDatalakeStatusParams {
	var ()
	return &GetRestoreDatalakeStatusParams{
		HTTPClient: client,
	}
}

/*
GetRestoreDatalakeStatusParams contains all the parameters to send to the API endpoint
for the get restore datalake status operation typically these are written to a http.Request
*/
type GetRestoreDatalakeStatusParams struct {

	/*BackupName
	  optional: datalake backup name

	*/
	BackupName *string
	/*Name
	  required: datalake name

	*/
	Name string
	/*RestoreID
	  optional: datalake restore id

	*/
	RestoreID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get restore datalake status params
func (o *GetRestoreDatalakeStatusParams) WithTimeout(timeout time.Duration) *GetRestoreDatalakeStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get restore datalake status params
func (o *GetRestoreDatalakeStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get restore datalake status params
func (o *GetRestoreDatalakeStatusParams) WithContext(ctx context.Context) *GetRestoreDatalakeStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get restore datalake status params
func (o *GetRestoreDatalakeStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get restore datalake status params
func (o *GetRestoreDatalakeStatusParams) WithHTTPClient(client *http.Client) *GetRestoreDatalakeStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get restore datalake status params
func (o *GetRestoreDatalakeStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBackupName adds the backupName to the get restore datalake status params
func (o *GetRestoreDatalakeStatusParams) WithBackupName(backupName *string) *GetRestoreDatalakeStatusParams {
	o.SetBackupName(backupName)
	return o
}

// SetBackupName adds the backupName to the get restore datalake status params
func (o *GetRestoreDatalakeStatusParams) SetBackupName(backupName *string) {
	o.BackupName = backupName
}

// WithName adds the name to the get restore datalake status params
func (o *GetRestoreDatalakeStatusParams) WithName(name string) *GetRestoreDatalakeStatusParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get restore datalake status params
func (o *GetRestoreDatalakeStatusParams) SetName(name string) {
	o.Name = name
}

// WithRestoreID adds the restoreID to the get restore datalake status params
func (o *GetRestoreDatalakeStatusParams) WithRestoreID(restoreID *string) *GetRestoreDatalakeStatusParams {
	o.SetRestoreID(restoreID)
	return o
}

// SetRestoreID adds the restoreId to the get restore datalake status params
func (o *GetRestoreDatalakeStatusParams) SetRestoreID(restoreID *string) {
	o.RestoreID = restoreID
}

// WriteToRequest writes these params to a swagger request
func (o *GetRestoreDatalakeStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.RestoreID != nil {

		// query param restoreId
		var qrRestoreID string
		if o.RestoreID != nil {
			qrRestoreID = *o.RestoreID
		}
		qRestoreID := qrRestoreID
		if qRestoreID != "" {
			if err := r.SetQueryParam("restoreId", qRestoreID); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
