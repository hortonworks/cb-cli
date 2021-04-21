// Code generated by go-swagger; DO NOT EDIT.

package v1freeipatest

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

	model "github.com/hortonworks/cb-cli/dataplane/api-freeipa/model"
)

// NewCheckGroupsParams creates a new CheckGroupsParams object
// with the default values initialized.
func NewCheckGroupsParams() *CheckGroupsParams {
	var ()
	return &CheckGroupsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCheckGroupsParamsWithTimeout creates a new CheckGroupsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCheckGroupsParamsWithTimeout(timeout time.Duration) *CheckGroupsParams {
	var ()
	return &CheckGroupsParams{

		timeout: timeout,
	}
}

// NewCheckGroupsParamsWithContext creates a new CheckGroupsParams object
// with the default values initialized, and the ability to set a context for a request
func NewCheckGroupsParamsWithContext(ctx context.Context) *CheckGroupsParams {
	var ()
	return &CheckGroupsParams{

		Context: ctx,
	}
}

// NewCheckGroupsParamsWithHTTPClient creates a new CheckGroupsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCheckGroupsParamsWithHTTPClient(client *http.Client) *CheckGroupsParams {
	var ()
	return &CheckGroupsParams{
		HTTPClient: client,
	}
}

/*CheckGroupsParams contains all the parameters to send to the API endpoint
for the check groups operation typically these are written to a http.Request
*/
type CheckGroupsParams struct {

	/*Body*/
	Body *model.CheckGroupsV1Request

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the check groups params
func (o *CheckGroupsParams) WithTimeout(timeout time.Duration) *CheckGroupsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the check groups params
func (o *CheckGroupsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the check groups params
func (o *CheckGroupsParams) WithContext(ctx context.Context) *CheckGroupsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the check groups params
func (o *CheckGroupsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the check groups params
func (o *CheckGroupsParams) WithHTTPClient(client *http.Client) *CheckGroupsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the check groups params
func (o *CheckGroupsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the check groups params
func (o *CheckGroupsParams) WithBody(body *model.CheckGroupsV1Request) *CheckGroupsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the check groups params
func (o *CheckGroupsParams) SetBody(body *model.CheckGroupsV1Request) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CheckGroupsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
