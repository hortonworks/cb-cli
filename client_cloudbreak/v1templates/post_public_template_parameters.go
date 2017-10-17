// Code generated by go-swagger; DO NOT EDIT.

package v1templates

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

// NewPostPublicTemplateParams creates a new PostPublicTemplateParams object
// with the default values initialized.
func NewPostPublicTemplateParams() *PostPublicTemplateParams {
	var ()
	return &PostPublicTemplateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostPublicTemplateParamsWithTimeout creates a new PostPublicTemplateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostPublicTemplateParamsWithTimeout(timeout time.Duration) *PostPublicTemplateParams {
	var ()
	return &PostPublicTemplateParams{

		timeout: timeout,
	}
}

// NewPostPublicTemplateParamsWithContext creates a new PostPublicTemplateParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostPublicTemplateParamsWithContext(ctx context.Context) *PostPublicTemplateParams {
	var ()
	return &PostPublicTemplateParams{

		Context: ctx,
	}
}

// NewPostPublicTemplateParamsWithHTTPClient creates a new PostPublicTemplateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostPublicTemplateParamsWithHTTPClient(client *http.Client) *PostPublicTemplateParams {
	var ()
	return &PostPublicTemplateParams{
		HTTPClient: client,
	}
}

/*PostPublicTemplateParams contains all the parameters to send to the API endpoint
for the post public template operation typically these are written to a http.Request
*/
type PostPublicTemplateParams struct {

	/*Body*/
	Body *models_cloudbreak.TemplateRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post public template params
func (o *PostPublicTemplateParams) WithTimeout(timeout time.Duration) *PostPublicTemplateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post public template params
func (o *PostPublicTemplateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post public template params
func (o *PostPublicTemplateParams) WithContext(ctx context.Context) *PostPublicTemplateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post public template params
func (o *PostPublicTemplateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post public template params
func (o *PostPublicTemplateParams) WithHTTPClient(client *http.Client) *PostPublicTemplateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post public template params
func (o *PostPublicTemplateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post public template params
func (o *PostPublicTemplateParams) WithBody(body *models_cloudbreak.TemplateRequest) *PostPublicTemplateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post public template params
func (o *PostPublicTemplateParams) SetBody(body *models_cloudbreak.TemplateRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostPublicTemplateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body == nil {
		o.Body = new(models_cloudbreak.TemplateRequest)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
