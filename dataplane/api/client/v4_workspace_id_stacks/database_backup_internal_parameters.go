// Code generated by go-swagger; DO NOT EDIT.

package v4_workspace_id_stacks

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

// NewDatabaseBackupInternalParams creates a new DatabaseBackupInternalParams object
// with the default values initialized.
func NewDatabaseBackupInternalParams() *DatabaseBackupInternalParams {
	var ()
	return &DatabaseBackupInternalParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDatabaseBackupInternalParamsWithTimeout creates a new DatabaseBackupInternalParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDatabaseBackupInternalParamsWithTimeout(timeout time.Duration) *DatabaseBackupInternalParams {
	var ()
	return &DatabaseBackupInternalParams{

		timeout: timeout,
	}
}

// NewDatabaseBackupInternalParamsWithContext creates a new DatabaseBackupInternalParams object
// with the default values initialized, and the ability to set a context for a request
func NewDatabaseBackupInternalParamsWithContext(ctx context.Context) *DatabaseBackupInternalParams {
	var ()
	return &DatabaseBackupInternalParams{

		Context: ctx,
	}
}

// NewDatabaseBackupInternalParamsWithHTTPClient creates a new DatabaseBackupInternalParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDatabaseBackupInternalParamsWithHTTPClient(client *http.Client) *DatabaseBackupInternalParams {
	var ()
	return &DatabaseBackupInternalParams{
		HTTPClient: client,
	}
}

/*
DatabaseBackupInternalParams contains all the parameters to send to the API endpoint
for the database backup internal operation typically these are written to a http.Request
*/
type DatabaseBackupInternalParams struct {

	/*BackupID*/
	BackupID *string
	/*BackupLocation*/
	BackupLocation *string
	/*CloseConnections*/
	CloseConnections *bool
	/*InitiatorUserCrn*/
	InitiatorUserCrn *string
	/*Name*/
	Name string
	/*SkipDatabaseNames*/
	SkipDatabaseNames []string
	/*WorkspaceID*/
	WorkspaceID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the database backup internal params
func (o *DatabaseBackupInternalParams) WithTimeout(timeout time.Duration) *DatabaseBackupInternalParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the database backup internal params
func (o *DatabaseBackupInternalParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the database backup internal params
func (o *DatabaseBackupInternalParams) WithContext(ctx context.Context) *DatabaseBackupInternalParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the database backup internal params
func (o *DatabaseBackupInternalParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the database backup internal params
func (o *DatabaseBackupInternalParams) WithHTTPClient(client *http.Client) *DatabaseBackupInternalParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the database backup internal params
func (o *DatabaseBackupInternalParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBackupID adds the backupID to the database backup internal params
func (o *DatabaseBackupInternalParams) WithBackupID(backupID *string) *DatabaseBackupInternalParams {
	o.SetBackupID(backupID)
	return o
}

// SetBackupID adds the backupId to the database backup internal params
func (o *DatabaseBackupInternalParams) SetBackupID(backupID *string) {
	o.BackupID = backupID
}

// WithBackupLocation adds the backupLocation to the database backup internal params
func (o *DatabaseBackupInternalParams) WithBackupLocation(backupLocation *string) *DatabaseBackupInternalParams {
	o.SetBackupLocation(backupLocation)
	return o
}

// SetBackupLocation adds the backupLocation to the database backup internal params
func (o *DatabaseBackupInternalParams) SetBackupLocation(backupLocation *string) {
	o.BackupLocation = backupLocation
}

// WithCloseConnections adds the closeConnections to the database backup internal params
func (o *DatabaseBackupInternalParams) WithCloseConnections(closeConnections *bool) *DatabaseBackupInternalParams {
	o.SetCloseConnections(closeConnections)
	return o
}

// SetCloseConnections adds the closeConnections to the database backup internal params
func (o *DatabaseBackupInternalParams) SetCloseConnections(closeConnections *bool) {
	o.CloseConnections = closeConnections
}

// WithInitiatorUserCrn adds the initiatorUserCrn to the database backup internal params
func (o *DatabaseBackupInternalParams) WithInitiatorUserCrn(initiatorUserCrn *string) *DatabaseBackupInternalParams {
	o.SetInitiatorUserCrn(initiatorUserCrn)
	return o
}

// SetInitiatorUserCrn adds the initiatorUserCrn to the database backup internal params
func (o *DatabaseBackupInternalParams) SetInitiatorUserCrn(initiatorUserCrn *string) {
	o.InitiatorUserCrn = initiatorUserCrn
}

// WithName adds the name to the database backup internal params
func (o *DatabaseBackupInternalParams) WithName(name string) *DatabaseBackupInternalParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the database backup internal params
func (o *DatabaseBackupInternalParams) SetName(name string) {
	o.Name = name
}

// WithSkipDatabaseNames adds the skipDatabaseNames to the database backup internal params
func (o *DatabaseBackupInternalParams) WithSkipDatabaseNames(skipDatabaseNames []string) *DatabaseBackupInternalParams {
	o.SetSkipDatabaseNames(skipDatabaseNames)
	return o
}

// SetSkipDatabaseNames adds the skipDatabaseNames to the database backup internal params
func (o *DatabaseBackupInternalParams) SetSkipDatabaseNames(skipDatabaseNames []string) {
	o.SkipDatabaseNames = skipDatabaseNames
}

// WithWorkspaceID adds the workspaceID to the database backup internal params
func (o *DatabaseBackupInternalParams) WithWorkspaceID(workspaceID int64) *DatabaseBackupInternalParams {
	o.SetWorkspaceID(workspaceID)
	return o
}

// SetWorkspaceID adds the workspaceId to the database backup internal params
func (o *DatabaseBackupInternalParams) SetWorkspaceID(workspaceID int64) {
	o.WorkspaceID = workspaceID
}

// WriteToRequest writes these params to a swagger request
func (o *DatabaseBackupInternalParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.BackupLocation != nil {

		// query param backupLocation
		var qrBackupLocation string
		if o.BackupLocation != nil {
			qrBackupLocation = *o.BackupLocation
		}
		qBackupLocation := qrBackupLocation
		if qBackupLocation != "" {
			if err := r.SetQueryParam("backupLocation", qBackupLocation); err != nil {
				return err
			}
		}

	}

	if o.CloseConnections != nil {

		// query param closeConnections
		var qrCloseConnections bool
		if o.CloseConnections != nil {
			qrCloseConnections = *o.CloseConnections
		}
		qCloseConnections := swag.FormatBool(qrCloseConnections)
		if qCloseConnections != "" {
			if err := r.SetQueryParam("closeConnections", qCloseConnections); err != nil {
				return err
			}
		}

	}

	if o.InitiatorUserCrn != nil {

		// query param initiatorUserCrn
		var qrInitiatorUserCrn string
		if o.InitiatorUserCrn != nil {
			qrInitiatorUserCrn = *o.InitiatorUserCrn
		}
		qInitiatorUserCrn := qrInitiatorUserCrn
		if qInitiatorUserCrn != "" {
			if err := r.SetQueryParam("initiatorUserCrn", qInitiatorUserCrn); err != nil {
				return err
			}
		}

	}

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	valuesSkipDatabaseNames := o.SkipDatabaseNames

	joinedSkipDatabaseNames := swag.JoinByFormat(valuesSkipDatabaseNames, "multi")
	// query array param skipDatabaseNames
	if err := r.SetQueryParam("skipDatabaseNames", joinedSkipDatabaseNames...); err != nil {
		return err
	}

	// path param workspaceId
	if err := r.SetPathParam("workspaceId", swag.FormatInt64(o.WorkspaceID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
