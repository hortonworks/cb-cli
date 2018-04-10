// Code generated by go-swagger; DO NOT EDIT.

package v1mpacks

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

// NewPostManagementPackParams creates a new PostManagementPackParams object
// with the default values initialized.
func NewPostManagementPackParams() *PostManagementPackParams {
	var ()
	return &PostManagementPackParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostManagementPackParamsWithTimeout creates a new PostManagementPackParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostManagementPackParamsWithTimeout(timeout time.Duration) *PostManagementPackParams {
	var ()
	return &PostManagementPackParams{

		timeout: timeout,
	}
}

// NewPostManagementPackParamsWithContext creates a new PostManagementPackParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostManagementPackParamsWithContext(ctx context.Context) *PostManagementPackParams {
	var ()
	return &PostManagementPackParams{

		Context: ctx,
	}
}

// NewPostManagementPackParamsWithHTTPClient creates a new PostManagementPackParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostManagementPackParamsWithHTTPClient(client *http.Client) *PostManagementPackParams {
	var ()
	return &PostManagementPackParams{
		HTTPClient: client,
	}
}

/*PostManagementPackParams contains all the parameters to send to the API endpoint
for the post management pack operation typically these are written to a http.Request
*/
type PostManagementPackParams struct {

	/*Body*/
	Body *models_cloudbreak.MpackRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post management pack params
func (o *PostManagementPackParams) WithTimeout(timeout time.Duration) *PostManagementPackParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post management pack params
func (o *PostManagementPackParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post management pack params
func (o *PostManagementPackParams) WithContext(ctx context.Context) *PostManagementPackParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post management pack params
func (o *PostManagementPackParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post management pack params
func (o *PostManagementPackParams) WithHTTPClient(client *http.Client) *PostManagementPackParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post management pack params
func (o *PostManagementPackParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post management pack params
func (o *PostManagementPackParams) WithBody(body *models_cloudbreak.MpackRequest) *PostManagementPackParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post management pack params
func (o *PostManagementPackParams) SetBody(body *models_cloudbreak.MpackRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostManagementPackParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body == nil {
		o.Body = new(models_cloudbreak.MpackRequest)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
