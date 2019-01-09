// Code generated by go-swagger; DO NOT EDIT.

package v4utils

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

// NewRepositoryConfigsValidationV4Params creates a new RepositoryConfigsValidationV4Params object
// with the default values initialized.
func NewRepositoryConfigsValidationV4Params() *RepositoryConfigsValidationV4Params {
	var ()
	return &RepositoryConfigsValidationV4Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewRepositoryConfigsValidationV4ParamsWithTimeout creates a new RepositoryConfigsValidationV4Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewRepositoryConfigsValidationV4ParamsWithTimeout(timeout time.Duration) *RepositoryConfigsValidationV4Params {
	var ()
	return &RepositoryConfigsValidationV4Params{

		timeout: timeout,
	}
}

// NewRepositoryConfigsValidationV4ParamsWithContext creates a new RepositoryConfigsValidationV4Params object
// with the default values initialized, and the ability to set a context for a request
func NewRepositoryConfigsValidationV4ParamsWithContext(ctx context.Context) *RepositoryConfigsValidationV4Params {
	var ()
	return &RepositoryConfigsValidationV4Params{

		Context: ctx,
	}
}

// NewRepositoryConfigsValidationV4ParamsWithHTTPClient creates a new RepositoryConfigsValidationV4Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRepositoryConfigsValidationV4ParamsWithHTTPClient(client *http.Client) *RepositoryConfigsValidationV4Params {
	var ()
	return &RepositoryConfigsValidationV4Params{
		HTTPClient: client,
	}
}

/*RepositoryConfigsValidationV4Params contains all the parameters to send to the API endpoint
for the repository configs validation v4 operation typically these are written to a http.Request
*/
type RepositoryConfigsValidationV4Params struct {

	/*Body*/
	Body *model.RepoConfigValidationV4Request

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the repository configs validation v4 params
func (o *RepositoryConfigsValidationV4Params) WithTimeout(timeout time.Duration) *RepositoryConfigsValidationV4Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the repository configs validation v4 params
func (o *RepositoryConfigsValidationV4Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the repository configs validation v4 params
func (o *RepositoryConfigsValidationV4Params) WithContext(ctx context.Context) *RepositoryConfigsValidationV4Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the repository configs validation v4 params
func (o *RepositoryConfigsValidationV4Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the repository configs validation v4 params
func (o *RepositoryConfigsValidationV4Params) WithHTTPClient(client *http.Client) *RepositoryConfigsValidationV4Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the repository configs validation v4 params
func (o *RepositoryConfigsValidationV4Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the repository configs validation v4 params
func (o *RepositoryConfigsValidationV4Params) WithBody(body *model.RepoConfigValidationV4Request) *RepositoryConfigsValidationV4Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the repository configs validation v4 params
func (o *RepositoryConfigsValidationV4Params) SetBody(body *model.RepoConfigValidationV4Request) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *RepositoryConfigsValidationV4Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
