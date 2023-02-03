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

// NewRegisterDatabaseServerParams creates a new RegisterDatabaseServerParams object
// with the default values initialized.
func NewRegisterDatabaseServerParams() *RegisterDatabaseServerParams {
	var ()
	return &RegisterDatabaseServerParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRegisterDatabaseServerParamsWithTimeout creates a new RegisterDatabaseServerParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRegisterDatabaseServerParamsWithTimeout(timeout time.Duration) *RegisterDatabaseServerParams {
	var ()
	return &RegisterDatabaseServerParams{

		timeout: timeout,
	}
}

// NewRegisterDatabaseServerParamsWithContext creates a new RegisterDatabaseServerParams object
// with the default values initialized, and the ability to set a context for a request
func NewRegisterDatabaseServerParamsWithContext(ctx context.Context) *RegisterDatabaseServerParams {
	var ()
	return &RegisterDatabaseServerParams{

		Context: ctx,
	}
}

// NewRegisterDatabaseServerParamsWithHTTPClient creates a new RegisterDatabaseServerParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRegisterDatabaseServerParamsWithHTTPClient(client *http.Client) *RegisterDatabaseServerParams {
	var ()
	return &RegisterDatabaseServerParams{
		HTTPClient: client,
	}
}

/*
RegisterDatabaseServerParams contains all the parameters to send to the API endpoint
for the register database server operation typically these are written to a http.Request
*/
type RegisterDatabaseServerParams struct {

	/*Body
	  Request containing information about a database server to be registered

	*/
	Body *model.DatabaseServerV4Request

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the register database server params
func (o *RegisterDatabaseServerParams) WithTimeout(timeout time.Duration) *RegisterDatabaseServerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the register database server params
func (o *RegisterDatabaseServerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the register database server params
func (o *RegisterDatabaseServerParams) WithContext(ctx context.Context) *RegisterDatabaseServerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the register database server params
func (o *RegisterDatabaseServerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the register database server params
func (o *RegisterDatabaseServerParams) WithHTTPClient(client *http.Client) *RegisterDatabaseServerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the register database server params
func (o *RegisterDatabaseServerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the register database server params
func (o *RegisterDatabaseServerParams) WithBody(body *model.DatabaseServerV4Request) *RegisterDatabaseServerParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the register database server params
func (o *RegisterDatabaseServerParams) SetBody(body *model.DatabaseServerV4Request) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *RegisterDatabaseServerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
