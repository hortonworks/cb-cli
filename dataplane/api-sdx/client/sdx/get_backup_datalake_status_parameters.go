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

// NewGetBackupDatalakeStatusParams creates a new GetBackupDatalakeStatusParams object
// with the default values initialized.
func NewGetBackupDatalakeStatusParams() *GetBackupDatalakeStatusParams {
	var ()
	return &GetBackupDatalakeStatusParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetBackupDatalakeStatusParamsWithTimeout creates a new GetBackupDatalakeStatusParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetBackupDatalakeStatusParamsWithTimeout(timeout time.Duration) *GetBackupDatalakeStatusParams {
	var ()
	return &GetBackupDatalakeStatusParams{

		timeout: timeout,
	}
}

// NewGetBackupDatalakeStatusParamsWithContext creates a new GetBackupDatalakeStatusParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetBackupDatalakeStatusParamsWithContext(ctx context.Context) *GetBackupDatalakeStatusParams {
	var ()
	return &GetBackupDatalakeStatusParams{

		Context: ctx,
	}
}

// NewGetBackupDatalakeStatusParamsWithHTTPClient creates a new GetBackupDatalakeStatusParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetBackupDatalakeStatusParamsWithHTTPClient(client *http.Client) *GetBackupDatalakeStatusParams {
	var ()
	return &GetBackupDatalakeStatusParams{
		HTTPClient: client,
	}
}

/*GetBackupDatalakeStatusParams contains all the parameters to send to the API endpoint
for the get backup datalake status operation typically these are written to a http.Request
*/
type GetBackupDatalakeStatusParams struct {

	/*BackupID
	  optional: datalake backup id

	*/
	BackupID *string
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

// WithTimeout adds the timeout to the get backup datalake status params
func (o *GetBackupDatalakeStatusParams) WithTimeout(timeout time.Duration) *GetBackupDatalakeStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get backup datalake status params
func (o *GetBackupDatalakeStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get backup datalake status params
func (o *GetBackupDatalakeStatusParams) WithContext(ctx context.Context) *GetBackupDatalakeStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get backup datalake status params
func (o *GetBackupDatalakeStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get backup datalake status params
func (o *GetBackupDatalakeStatusParams) WithHTTPClient(client *http.Client) *GetBackupDatalakeStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get backup datalake status params
func (o *GetBackupDatalakeStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBackupID adds the backupID to the get backup datalake status params
func (o *GetBackupDatalakeStatusParams) WithBackupID(backupID *string) *GetBackupDatalakeStatusParams {
	o.SetBackupID(backupID)
	return o
}

// SetBackupID adds the backupId to the get backup datalake status params
func (o *GetBackupDatalakeStatusParams) SetBackupID(backupID *string) {
	o.BackupID = backupID
}

// WithBackupName adds the backupName to the get backup datalake status params
func (o *GetBackupDatalakeStatusParams) WithBackupName(backupName *string) *GetBackupDatalakeStatusParams {
	o.SetBackupName(backupName)
	return o
}

// SetBackupName adds the backupName to the get backup datalake status params
func (o *GetBackupDatalakeStatusParams) SetBackupName(backupName *string) {
	o.BackupName = backupName
}

// WithName adds the name to the get backup datalake status params
func (o *GetBackupDatalakeStatusParams) WithName(name string) *GetBackupDatalakeStatusParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get backup datalake status params
func (o *GetBackupDatalakeStatusParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *GetBackupDatalakeStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.BackupID != nil {

		// query param backupId
		var qrBackupID string
		if o.BackupID != nil {
			qrBackupID = *o.BackupID
		}
		qBackupID := qrBackupID
		if qBackupID != "" {
			if err := r.SetQueryParam("backupId", qBackupID); err != nil {
				return err
			}
		}

	}

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
