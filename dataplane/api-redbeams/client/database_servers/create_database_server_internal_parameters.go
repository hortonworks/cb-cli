// Code generated by go-swagger; DO NOT EDIT.

package database_servers

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

	model "github.com/hortonworks/cb-cli/dataplane/api-redbeams/model"
)

// NewCreateDatabaseServerInternalParams creates a new CreateDatabaseServerInternalParams object
// with the default values initialized.
func NewCreateDatabaseServerInternalParams() *CreateDatabaseServerInternalParams {
	var ()
	return &CreateDatabaseServerInternalParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateDatabaseServerInternalParamsWithTimeout creates a new CreateDatabaseServerInternalParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateDatabaseServerInternalParamsWithTimeout(timeout time.Duration) *CreateDatabaseServerInternalParams {
	var ()
	return &CreateDatabaseServerInternalParams{

		timeout: timeout,
	}
}

// NewCreateDatabaseServerInternalParamsWithContext creates a new CreateDatabaseServerInternalParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateDatabaseServerInternalParamsWithContext(ctx context.Context) *CreateDatabaseServerInternalParams {
	var ()
	return &CreateDatabaseServerInternalParams{

		Context: ctx,
	}
}

// NewCreateDatabaseServerInternalParamsWithHTTPClient creates a new CreateDatabaseServerInternalParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateDatabaseServerInternalParamsWithHTTPClient(client *http.Client) *CreateDatabaseServerInternalParams {
	var ()
	return &CreateDatabaseServerInternalParams{
		HTTPClient: client,
	}
}

/*
CreateDatabaseServerInternalParams contains all the parameters to send to the API endpoint
for the create database server internal operation typically these are written to a http.Request
*/
type CreateDatabaseServerInternalParams struct {

	/*Body
	  Request for allocating a new database server in a provider

	*/
	Body *model.AllocateDatabaseServerV4Request
	/*InitiatorUserCrn*/
	InitiatorUserCrn *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create database server internal params
func (o *CreateDatabaseServerInternalParams) WithTimeout(timeout time.Duration) *CreateDatabaseServerInternalParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create database server internal params
func (o *CreateDatabaseServerInternalParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create database server internal params
func (o *CreateDatabaseServerInternalParams) WithContext(ctx context.Context) *CreateDatabaseServerInternalParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create database server internal params
func (o *CreateDatabaseServerInternalParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create database server internal params
func (o *CreateDatabaseServerInternalParams) WithHTTPClient(client *http.Client) *CreateDatabaseServerInternalParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create database server internal params
func (o *CreateDatabaseServerInternalParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create database server internal params
func (o *CreateDatabaseServerInternalParams) WithBody(body *model.AllocateDatabaseServerV4Request) *CreateDatabaseServerInternalParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create database server internal params
func (o *CreateDatabaseServerInternalParams) SetBody(body *model.AllocateDatabaseServerV4Request) {
	o.Body = body
}

// WithInitiatorUserCrn adds the initiatorUserCrn to the create database server internal params
func (o *CreateDatabaseServerInternalParams) WithInitiatorUserCrn(initiatorUserCrn *string) *CreateDatabaseServerInternalParams {
	o.SetInitiatorUserCrn(initiatorUserCrn)
	return o
}

// SetInitiatorUserCrn adds the initiatorUserCrn to the create database server internal params
func (o *CreateDatabaseServerInternalParams) SetInitiatorUserCrn(initiatorUserCrn *string) {
	o.InitiatorUserCrn = initiatorUserCrn
}

// WriteToRequest writes these params to a swagger request
func (o *CreateDatabaseServerInternalParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
