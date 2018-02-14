// Code generated by go-swagger; DO NOT EDIT.

package v1repositoryconfigs

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

// NewPostRepositoryConfigsValidationParams creates a new PostRepositoryConfigsValidationParams object
// with the default values initialized.
func NewPostRepositoryConfigsValidationParams() *PostRepositoryConfigsValidationParams {
	var ()
	return &PostRepositoryConfigsValidationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostRepositoryConfigsValidationParamsWithTimeout creates a new PostRepositoryConfigsValidationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostRepositoryConfigsValidationParamsWithTimeout(timeout time.Duration) *PostRepositoryConfigsValidationParams {
	var ()
	return &PostRepositoryConfigsValidationParams{

		timeout: timeout,
	}
}

// NewPostRepositoryConfigsValidationParamsWithContext creates a new PostRepositoryConfigsValidationParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostRepositoryConfigsValidationParamsWithContext(ctx context.Context) *PostRepositoryConfigsValidationParams {
	var ()
	return &PostRepositoryConfigsValidationParams{

		Context: ctx,
	}
}

// NewPostRepositoryConfigsValidationParamsWithHTTPClient creates a new PostRepositoryConfigsValidationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostRepositoryConfigsValidationParamsWithHTTPClient(client *http.Client) *PostRepositoryConfigsValidationParams {
	var ()
	return &PostRepositoryConfigsValidationParams{
		HTTPClient: client,
	}
}

/*PostRepositoryConfigsValidationParams contains all the parameters to send to the API endpoint
for the post repository configs validation operation typically these are written to a http.Request
*/
type PostRepositoryConfigsValidationParams struct {

	/*Body*/
	Body *models_cloudbreak.RepoConfigValidationRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post repository configs validation params
func (o *PostRepositoryConfigsValidationParams) WithTimeout(timeout time.Duration) *PostRepositoryConfigsValidationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post repository configs validation params
func (o *PostRepositoryConfigsValidationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post repository configs validation params
func (o *PostRepositoryConfigsValidationParams) WithContext(ctx context.Context) *PostRepositoryConfigsValidationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post repository configs validation params
func (o *PostRepositoryConfigsValidationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post repository configs validation params
func (o *PostRepositoryConfigsValidationParams) WithHTTPClient(client *http.Client) *PostRepositoryConfigsValidationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post repository configs validation params
func (o *PostRepositoryConfigsValidationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post repository configs validation params
func (o *PostRepositoryConfigsValidationParams) WithBody(body *models_cloudbreak.RepoConfigValidationRequest) *PostRepositoryConfigsValidationParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post repository configs validation params
func (o *PostRepositoryConfigsValidationParams) SetBody(body *models_cloudbreak.RepoConfigValidationRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostRepositoryConfigsValidationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body == nil {
		o.Body = new(models_cloudbreak.RepoConfigValidationRequest)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
