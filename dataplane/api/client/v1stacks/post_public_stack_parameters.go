// Code generated by go-swagger; DO NOT EDIT.

package v1stacks

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

// NewPostPublicStackParams creates a new PostPublicStackParams object
// with the default values initialized.
func NewPostPublicStackParams() *PostPublicStackParams {
	var ()
	return &PostPublicStackParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostPublicStackParamsWithTimeout creates a new PostPublicStackParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostPublicStackParamsWithTimeout(timeout time.Duration) *PostPublicStackParams {
	var ()
	return &PostPublicStackParams{

		timeout: timeout,
	}
}

// NewPostPublicStackParamsWithContext creates a new PostPublicStackParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostPublicStackParamsWithContext(ctx context.Context) *PostPublicStackParams {
	var ()
	return &PostPublicStackParams{

		Context: ctx,
	}
}

// NewPostPublicStackParamsWithHTTPClient creates a new PostPublicStackParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostPublicStackParamsWithHTTPClient(client *http.Client) *PostPublicStackParams {
	var ()
	return &PostPublicStackParams{
		HTTPClient: client,
	}
}

/*PostPublicStackParams contains all the parameters to send to the API endpoint
for the post public stack operation typically these are written to a http.Request
*/
type PostPublicStackParams struct {

	/*Body*/
	Body *model.StackRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post public stack params
func (o *PostPublicStackParams) WithTimeout(timeout time.Duration) *PostPublicStackParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post public stack params
func (o *PostPublicStackParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post public stack params
func (o *PostPublicStackParams) WithContext(ctx context.Context) *PostPublicStackParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post public stack params
func (o *PostPublicStackParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post public stack params
func (o *PostPublicStackParams) WithHTTPClient(client *http.Client) *PostPublicStackParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post public stack params
func (o *PostPublicStackParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post public stack params
func (o *PostPublicStackParams) WithBody(body *model.StackRequest) *PostPublicStackParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post public stack params
func (o *PostPublicStackParams) SetBody(body *model.StackRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostPublicStackParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
