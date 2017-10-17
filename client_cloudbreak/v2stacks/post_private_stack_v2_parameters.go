// Code generated by go-swagger; DO NOT EDIT.

package v2stacks

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

	"github.com/hortonworks/cb-cli/models_cloudbreak"
)

// NewPostPrivateStackV2Params creates a new PostPrivateStackV2Params object
// with the default values initialized.
func NewPostPrivateStackV2Params() *PostPrivateStackV2Params {
	var ()
	return &PostPrivateStackV2Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostPrivateStackV2ParamsWithTimeout creates a new PostPrivateStackV2Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostPrivateStackV2ParamsWithTimeout(timeout time.Duration) *PostPrivateStackV2Params {
	var ()
	return &PostPrivateStackV2Params{

		timeout: timeout,
	}
}

// NewPostPrivateStackV2ParamsWithContext creates a new PostPrivateStackV2Params object
// with the default values initialized, and the ability to set a context for a request
func NewPostPrivateStackV2ParamsWithContext(ctx context.Context) *PostPrivateStackV2Params {
	var ()
	return &PostPrivateStackV2Params{

		Context: ctx,
	}
}

// NewPostPrivateStackV2ParamsWithHTTPClient creates a new PostPrivateStackV2Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostPrivateStackV2ParamsWithHTTPClient(client *http.Client) *PostPrivateStackV2Params {
	var ()
	return &PostPrivateStackV2Params{
		HTTPClient: client,
	}
}

/*PostPrivateStackV2Params contains all the parameters to send to the API endpoint
for the post private stack v2 operation typically these are written to a http.Request
*/
type PostPrivateStackV2Params struct {

	/*Body*/
	Body *models_cloudbreak.StackV2Request

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post private stack v2 params
func (o *PostPrivateStackV2Params) WithTimeout(timeout time.Duration) *PostPrivateStackV2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post private stack v2 params
func (o *PostPrivateStackV2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post private stack v2 params
func (o *PostPrivateStackV2Params) WithContext(ctx context.Context) *PostPrivateStackV2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post private stack v2 params
func (o *PostPrivateStackV2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post private stack v2 params
func (o *PostPrivateStackV2Params) WithHTTPClient(client *http.Client) *PostPrivateStackV2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post private stack v2 params
func (o *PostPrivateStackV2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post private stack v2 params
func (o *PostPrivateStackV2Params) WithBody(body *models_cloudbreak.StackV2Request) *PostPrivateStackV2Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post private stack v2 params
func (o *PostPrivateStackV2Params) SetBody(body *models_cloudbreak.StackV2Request) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostPrivateStackV2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body == nil {
		o.Body = new(models_cloudbreak.StackV2Request)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
