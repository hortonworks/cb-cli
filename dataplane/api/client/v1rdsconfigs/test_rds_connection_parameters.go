// Code generated by go-swagger; DO NOT EDIT.

package v1rdsconfigs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api/model"
)

// NewTestRdsConnectionParams creates a new TestRdsConnectionParams object
// with the default values initialized.
func NewTestRdsConnectionParams() *TestRdsConnectionParams {
	var ()
	return &TestRdsConnectionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewTestRdsConnectionParamsWithTimeout creates a new TestRdsConnectionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTestRdsConnectionParamsWithTimeout(timeout time.Duration) *TestRdsConnectionParams {
	var ()
	return &TestRdsConnectionParams{

		timeout: timeout,
	}
}

// NewTestRdsConnectionParamsWithContext creates a new TestRdsConnectionParams object
// with the default values initialized, and the ability to set a context for a request
func NewTestRdsConnectionParamsWithContext(ctx context.Context) *TestRdsConnectionParams {
	var ()
	return &TestRdsConnectionParams{

		Context: ctx,
	}
}

// NewTestRdsConnectionParamsWithHTTPClient creates a new TestRdsConnectionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTestRdsConnectionParamsWithHTTPClient(client *http.Client) *TestRdsConnectionParams {
	var ()
	return &TestRdsConnectionParams{
		HTTPClient: client,
	}
}

/*TestRdsConnectionParams contains all the parameters to send to the API endpoint
for the test rds connection operation typically these are written to a http.Request
*/
type TestRdsConnectionParams struct {

	/*Body*/
	Body *model.RdsTestRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the test rds connection params
func (o *TestRdsConnectionParams) WithTimeout(timeout time.Duration) *TestRdsConnectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the test rds connection params
func (o *TestRdsConnectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the test rds connection params
func (o *TestRdsConnectionParams) WithContext(ctx context.Context) *TestRdsConnectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the test rds connection params
func (o *TestRdsConnectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the test rds connection params
func (o *TestRdsConnectionParams) WithHTTPClient(client *http.Client) *TestRdsConnectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the test rds connection params
func (o *TestRdsConnectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the test rds connection params
func (o *TestRdsConnectionParams) WithBody(body *model.RdsTestRequest) *TestRdsConnectionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the test rds connection params
func (o *TestRdsConnectionParams) SetBody(body *model.RdsTestRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *TestRdsConnectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
