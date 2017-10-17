// Code generated by go-swagger; DO NOT EDIT.

package v1users

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

	"github.com/hortonworks/hdc-cli/models_cloudbreak"
)

// NewModifyProfileParams creates a new ModifyProfileParams object
// with the default values initialized.
func NewModifyProfileParams() *ModifyProfileParams {
	var ()
	return &ModifyProfileParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewModifyProfileParamsWithTimeout creates a new ModifyProfileParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewModifyProfileParamsWithTimeout(timeout time.Duration) *ModifyProfileParams {
	var ()
	return &ModifyProfileParams{

		timeout: timeout,
	}
}

// NewModifyProfileParamsWithContext creates a new ModifyProfileParams object
// with the default values initialized, and the ability to set a context for a request
func NewModifyProfileParamsWithContext(ctx context.Context) *ModifyProfileParams {
	var ()
	return &ModifyProfileParams{

		Context: ctx,
	}
}

// NewModifyProfileParamsWithHTTPClient creates a new ModifyProfileParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewModifyProfileParamsWithHTTPClient(client *http.Client) *ModifyProfileParams {
	var ()
	return &ModifyProfileParams{
		HTTPClient: client,
	}
}

/*ModifyProfileParams contains all the parameters to send to the API endpoint
for the modify profile operation typically these are written to a http.Request
*/
type ModifyProfileParams struct {

	/*Body*/
	Body *models_cloudbreak.UserProfileRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the modify profile params
func (o *ModifyProfileParams) WithTimeout(timeout time.Duration) *ModifyProfileParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the modify profile params
func (o *ModifyProfileParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the modify profile params
func (o *ModifyProfileParams) WithContext(ctx context.Context) *ModifyProfileParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the modify profile params
func (o *ModifyProfileParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the modify profile params
func (o *ModifyProfileParams) WithHTTPClient(client *http.Client) *ModifyProfileParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the modify profile params
func (o *ModifyProfileParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the modify profile params
func (o *ModifyProfileParams) WithBody(body *models_cloudbreak.UserProfileRequest) *ModifyProfileParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the modify profile params
func (o *ModifyProfileParams) SetBody(body *models_cloudbreak.UserProfileRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ModifyProfileParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body == nil {
		o.Body = new(models_cloudbreak.UserProfileRequest)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
