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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewRestoreDatalakeParams creates a new RestoreDatalakeParams object
// with the default values initialized.
func NewRestoreDatalakeParams() *RestoreDatalakeParams {
	var ()
	return &RestoreDatalakeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRestoreDatalakeParamsWithTimeout creates a new RestoreDatalakeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRestoreDatalakeParamsWithTimeout(timeout time.Duration) *RestoreDatalakeParams {
	var ()
	return &RestoreDatalakeParams{

		timeout: timeout,
	}
}

// NewRestoreDatalakeParamsWithContext creates a new RestoreDatalakeParams object
// with the default values initialized, and the ability to set a context for a request
func NewRestoreDatalakeParamsWithContext(ctx context.Context) *RestoreDatalakeParams {
	var ()
	return &RestoreDatalakeParams{

		Context: ctx,
	}
}

// NewRestoreDatalakeParamsWithHTTPClient creates a new RestoreDatalakeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRestoreDatalakeParamsWithHTTPClient(client *http.Client) *RestoreDatalakeParams {
	var ()
	return &RestoreDatalakeParams{
		HTTPClient: client,
	}
}

/*
RestoreDatalakeParams contains all the parameters to send to the API endpoint
for the restore datalake operation typically these are written to a http.Request
*/
type RestoreDatalakeParams struct {

	/*BackupID*/
	BackupID *string
	/*BackupLocationOverride*/
	BackupLocationOverride *string
	/*FullDrMaxDurationInMin*/
	FullDrMaxDurationInMin *int32
	/*Name*/
	Name string
	/*SkipAtlasMetadata*/
	SkipAtlasMetadata *bool
	/*SkipRangerAudits*/
	SkipRangerAudits *bool
	/*SkipRangerMetadata*/
	SkipRangerMetadata *bool
	/*SkipValidation*/
	SkipValidation *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the restore datalake params
func (o *RestoreDatalakeParams) WithTimeout(timeout time.Duration) *RestoreDatalakeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the restore datalake params
func (o *RestoreDatalakeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the restore datalake params
func (o *RestoreDatalakeParams) WithContext(ctx context.Context) *RestoreDatalakeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the restore datalake params
func (o *RestoreDatalakeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the restore datalake params
func (o *RestoreDatalakeParams) WithHTTPClient(client *http.Client) *RestoreDatalakeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the restore datalake params
func (o *RestoreDatalakeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBackupID adds the backupID to the restore datalake params
func (o *RestoreDatalakeParams) WithBackupID(backupID *string) *RestoreDatalakeParams {
	o.SetBackupID(backupID)
	return o
}

// SetBackupID adds the backupId to the restore datalake params
func (o *RestoreDatalakeParams) SetBackupID(backupID *string) {
	o.BackupID = backupID
}

// WithBackupLocationOverride adds the backupLocationOverride to the restore datalake params
func (o *RestoreDatalakeParams) WithBackupLocationOverride(backupLocationOverride *string) *RestoreDatalakeParams {
	o.SetBackupLocationOverride(backupLocationOverride)
	return o
}

// SetBackupLocationOverride adds the backupLocationOverride to the restore datalake params
func (o *RestoreDatalakeParams) SetBackupLocationOverride(backupLocationOverride *string) {
	o.BackupLocationOverride = backupLocationOverride
}

// WithFullDrMaxDurationInMin adds the fullDrMaxDurationInMin to the restore datalake params
func (o *RestoreDatalakeParams) WithFullDrMaxDurationInMin(fullDrMaxDurationInMin *int32) *RestoreDatalakeParams {
	o.SetFullDrMaxDurationInMin(fullDrMaxDurationInMin)
	return o
}

// SetFullDrMaxDurationInMin adds the fullDrMaxDurationInMin to the restore datalake params
func (o *RestoreDatalakeParams) SetFullDrMaxDurationInMin(fullDrMaxDurationInMin *int32) {
	o.FullDrMaxDurationInMin = fullDrMaxDurationInMin
}

// WithName adds the name to the restore datalake params
func (o *RestoreDatalakeParams) WithName(name string) *RestoreDatalakeParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the restore datalake params
func (o *RestoreDatalakeParams) SetName(name string) {
	o.Name = name
}

// WithSkipAtlasMetadata adds the skipAtlasMetadata to the restore datalake params
func (o *RestoreDatalakeParams) WithSkipAtlasMetadata(skipAtlasMetadata *bool) *RestoreDatalakeParams {
	o.SetSkipAtlasMetadata(skipAtlasMetadata)
	return o
}

// SetSkipAtlasMetadata adds the skipAtlasMetadata to the restore datalake params
func (o *RestoreDatalakeParams) SetSkipAtlasMetadata(skipAtlasMetadata *bool) {
	o.SkipAtlasMetadata = skipAtlasMetadata
}

// WithSkipRangerAudits adds the skipRangerAudits to the restore datalake params
func (o *RestoreDatalakeParams) WithSkipRangerAudits(skipRangerAudits *bool) *RestoreDatalakeParams {
	o.SetSkipRangerAudits(skipRangerAudits)
	return o
}

// SetSkipRangerAudits adds the skipRangerAudits to the restore datalake params
func (o *RestoreDatalakeParams) SetSkipRangerAudits(skipRangerAudits *bool) {
	o.SkipRangerAudits = skipRangerAudits
}

// WithSkipRangerMetadata adds the skipRangerMetadata to the restore datalake params
func (o *RestoreDatalakeParams) WithSkipRangerMetadata(skipRangerMetadata *bool) *RestoreDatalakeParams {
	o.SetSkipRangerMetadata(skipRangerMetadata)
	return o
}

// SetSkipRangerMetadata adds the skipRangerMetadata to the restore datalake params
func (o *RestoreDatalakeParams) SetSkipRangerMetadata(skipRangerMetadata *bool) {
	o.SkipRangerMetadata = skipRangerMetadata
}

// WithSkipValidation adds the skipValidation to the restore datalake params
func (o *RestoreDatalakeParams) WithSkipValidation(skipValidation *bool) *RestoreDatalakeParams {
	o.SetSkipValidation(skipValidation)
	return o
}

// SetSkipValidation adds the skipValidation to the restore datalake params
func (o *RestoreDatalakeParams) SetSkipValidation(skipValidation *bool) {
	o.SkipValidation = skipValidation
}

// WriteToRequest writes these params to a swagger request
func (o *RestoreDatalakeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.BackupLocationOverride != nil {

		// query param backupLocationOverride
		var qrBackupLocationOverride string
		if o.BackupLocationOverride != nil {
			qrBackupLocationOverride = *o.BackupLocationOverride
		}
		qBackupLocationOverride := qrBackupLocationOverride
		if qBackupLocationOverride != "" {
			if err := r.SetQueryParam("backupLocationOverride", qBackupLocationOverride); err != nil {
				return err
			}
		}

	}

	if o.FullDrMaxDurationInMin != nil {

		// query param fullDrMaxDurationInMin
		var qrFullDrMaxDurationInMin int32
		if o.FullDrMaxDurationInMin != nil {
			qrFullDrMaxDurationInMin = *o.FullDrMaxDurationInMin
		}
		qFullDrMaxDurationInMin := swag.FormatInt32(qrFullDrMaxDurationInMin)
		if qFullDrMaxDurationInMin != "" {
			if err := r.SetQueryParam("fullDrMaxDurationInMin", qFullDrMaxDurationInMin); err != nil {
				return err
			}
		}

	}

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if o.SkipAtlasMetadata != nil {

		// query param skipAtlasMetadata
		var qrSkipAtlasMetadata bool
		if o.SkipAtlasMetadata != nil {
			qrSkipAtlasMetadata = *o.SkipAtlasMetadata
		}
		qSkipAtlasMetadata := swag.FormatBool(qrSkipAtlasMetadata)
		if qSkipAtlasMetadata != "" {
			if err := r.SetQueryParam("skipAtlasMetadata", qSkipAtlasMetadata); err != nil {
				return err
			}
		}

	}

	if o.SkipRangerAudits != nil {

		// query param skipRangerAudits
		var qrSkipRangerAudits bool
		if o.SkipRangerAudits != nil {
			qrSkipRangerAudits = *o.SkipRangerAudits
		}
		qSkipRangerAudits := swag.FormatBool(qrSkipRangerAudits)
		if qSkipRangerAudits != "" {
			if err := r.SetQueryParam("skipRangerAudits", qSkipRangerAudits); err != nil {
				return err
			}
		}

	}

	if o.SkipRangerMetadata != nil {

		// query param skipRangerMetadata
		var qrSkipRangerMetadata bool
		if o.SkipRangerMetadata != nil {
			qrSkipRangerMetadata = *o.SkipRangerMetadata
		}
		qSkipRangerMetadata := swag.FormatBool(qrSkipRangerMetadata)
		if qSkipRangerMetadata != "" {
			if err := r.SetQueryParam("skipRangerMetadata", qSkipRangerMetadata); err != nil {
				return err
			}
		}

	}

	if o.SkipValidation != nil {

		// query param skipValidation
		var qrSkipValidation bool
		if o.SkipValidation != nil {
			qrSkipValidation = *o.SkipValidation
		}
		qSkipValidation := swag.FormatBool(qrSkipValidation)
		if qSkipValidation != "" {
			if err := r.SetQueryParam("skipValidation", qSkipValidation); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
