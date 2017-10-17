// Code generated by go-swagger; DO NOT EDIT.

package v1securitygroups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteSecurityGroupParams creates a new DeleteSecurityGroupParams object
// with the default values initialized.
func NewDeleteSecurityGroupParams() *DeleteSecurityGroupParams {
	var ()
	return &DeleteSecurityGroupParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteSecurityGroupParamsWithTimeout creates a new DeleteSecurityGroupParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteSecurityGroupParamsWithTimeout(timeout time.Duration) *DeleteSecurityGroupParams {
	var ()
	return &DeleteSecurityGroupParams{

		timeout: timeout,
	}
}

// NewDeleteSecurityGroupParamsWithContext creates a new DeleteSecurityGroupParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteSecurityGroupParamsWithContext(ctx context.Context) *DeleteSecurityGroupParams {
	var ()
	return &DeleteSecurityGroupParams{

		Context: ctx,
	}
}

// NewDeleteSecurityGroupParamsWithHTTPClient creates a new DeleteSecurityGroupParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteSecurityGroupParamsWithHTTPClient(client *http.Client) *DeleteSecurityGroupParams {
	var ()
	return &DeleteSecurityGroupParams{
		HTTPClient: client,
	}
}

/*DeleteSecurityGroupParams contains all the parameters to send to the API endpoint
for the delete security group operation typically these are written to a http.Request
*/
type DeleteSecurityGroupParams struct {

	/*ID*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete security group params
func (o *DeleteSecurityGroupParams) WithTimeout(timeout time.Duration) *DeleteSecurityGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete security group params
func (o *DeleteSecurityGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete security group params
func (o *DeleteSecurityGroupParams) WithContext(ctx context.Context) *DeleteSecurityGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete security group params
func (o *DeleteSecurityGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete security group params
func (o *DeleteSecurityGroupParams) WithHTTPClient(client *http.Client) *DeleteSecurityGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete security group params
func (o *DeleteSecurityGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete security group params
func (o *DeleteSecurityGroupParams) WithID(id int64) *DeleteSecurityGroupParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete security group params
func (o *DeleteSecurityGroupParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteSecurityGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
