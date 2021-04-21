// Code generated by go-swagger; DO NOT EDIT.

package v4utils

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

	model "github.com/hortonworks/cb-cli/dataplane/api/model"
)

// NewCheckRightOnResourcesParams creates a new CheckRightOnResourcesParams object
// with the default values initialized.
func NewCheckRightOnResourcesParams() *CheckRightOnResourcesParams {
	var ()
	return &CheckRightOnResourcesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCheckRightOnResourcesParamsWithTimeout creates a new CheckRightOnResourcesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCheckRightOnResourcesParamsWithTimeout(timeout time.Duration) *CheckRightOnResourcesParams {
	var ()
	return &CheckRightOnResourcesParams{

		timeout: timeout,
	}
}

// NewCheckRightOnResourcesParamsWithContext creates a new CheckRightOnResourcesParams object
// with the default values initialized, and the ability to set a context for a request
func NewCheckRightOnResourcesParamsWithContext(ctx context.Context) *CheckRightOnResourcesParams {
	var ()
	return &CheckRightOnResourcesParams{

		Context: ctx,
	}
}

// NewCheckRightOnResourcesParamsWithHTTPClient creates a new CheckRightOnResourcesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCheckRightOnResourcesParamsWithHTTPClient(client *http.Client) *CheckRightOnResourcesParams {
	var ()
	return &CheckRightOnResourcesParams{
		HTTPClient: client,
	}
}

/*CheckRightOnResourcesParams contains all the parameters to send to the API endpoint
for the check right on resources operation typically these are written to a http.Request
*/
type CheckRightOnResourcesParams struct {

	/*Body*/
	Body *model.CheckRightOnResourcesV4Request

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the check right on resources params
func (o *CheckRightOnResourcesParams) WithTimeout(timeout time.Duration) *CheckRightOnResourcesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the check right on resources params
func (o *CheckRightOnResourcesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the check right on resources params
func (o *CheckRightOnResourcesParams) WithContext(ctx context.Context) *CheckRightOnResourcesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the check right on resources params
func (o *CheckRightOnResourcesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the check right on resources params
func (o *CheckRightOnResourcesParams) WithHTTPClient(client *http.Client) *CheckRightOnResourcesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the check right on resources params
func (o *CheckRightOnResourcesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the check right on resources params
func (o *CheckRightOnResourcesParams) WithBody(body *model.CheckRightOnResourcesV4Request) *CheckRightOnResourcesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the check right on resources params
func (o *CheckRightOnResourcesParams) SetBody(body *model.CheckRightOnResourcesV4Request) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CheckRightOnResourcesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
