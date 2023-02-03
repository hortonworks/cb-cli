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

// NewCreateDatabaseOnServerParams creates a new CreateDatabaseOnServerParams object
// with the default values initialized.
func NewCreateDatabaseOnServerParams() *CreateDatabaseOnServerParams {
	var ()
	return &CreateDatabaseOnServerParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateDatabaseOnServerParamsWithTimeout creates a new CreateDatabaseOnServerParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateDatabaseOnServerParamsWithTimeout(timeout time.Duration) *CreateDatabaseOnServerParams {
	var ()
	return &CreateDatabaseOnServerParams{

		timeout: timeout,
	}
}

// NewCreateDatabaseOnServerParamsWithContext creates a new CreateDatabaseOnServerParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateDatabaseOnServerParamsWithContext(ctx context.Context) *CreateDatabaseOnServerParams {
	var ()
	return &CreateDatabaseOnServerParams{

		Context: ctx,
	}
}

// NewCreateDatabaseOnServerParamsWithHTTPClient creates a new CreateDatabaseOnServerParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateDatabaseOnServerParamsWithHTTPClient(client *http.Client) *CreateDatabaseOnServerParams {
	var ()
	return &CreateDatabaseOnServerParams{
		HTTPClient: client,
	}
}

/*
CreateDatabaseOnServerParams contains all the parameters to send to the API endpoint
for the create database on server operation typically these are written to a http.Request
*/
type CreateDatabaseOnServerParams struct {

	/*Body
	  Request for creating a new database on a registered database server

	*/
	Body *model.CreateDatabaseV4Request

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create database on server params
func (o *CreateDatabaseOnServerParams) WithTimeout(timeout time.Duration) *CreateDatabaseOnServerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create database on server params
func (o *CreateDatabaseOnServerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create database on server params
func (o *CreateDatabaseOnServerParams) WithContext(ctx context.Context) *CreateDatabaseOnServerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create database on server params
func (o *CreateDatabaseOnServerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create database on server params
func (o *CreateDatabaseOnServerParams) WithHTTPClient(client *http.Client) *CreateDatabaseOnServerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create database on server params
func (o *CreateDatabaseOnServerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create database on server params
func (o *CreateDatabaseOnServerParams) WithBody(body *model.CreateDatabaseV4Request) *CreateDatabaseOnServerParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create database on server params
func (o *CreateDatabaseOnServerParams) SetBody(body *model.CreateDatabaseV4Request) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateDatabaseOnServerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
