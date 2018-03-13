// Code generated by go-swagger; DO NOT EDIT.

package v1credentials

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

// NewPutPublicCredentialParams creates a new PutPublicCredentialParams object
// with the default values initialized.
func NewPutPublicCredentialParams() *PutPublicCredentialParams {
	var ()
	return &PutPublicCredentialParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutPublicCredentialParamsWithTimeout creates a new PutPublicCredentialParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutPublicCredentialParamsWithTimeout(timeout time.Duration) *PutPublicCredentialParams {
	var ()
	return &PutPublicCredentialParams{

		timeout: timeout,
	}
}

// NewPutPublicCredentialParamsWithContext creates a new PutPublicCredentialParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutPublicCredentialParamsWithContext(ctx context.Context) *PutPublicCredentialParams {
	var ()
	return &PutPublicCredentialParams{

		Context: ctx,
	}
}

// NewPutPublicCredentialParamsWithHTTPClient creates a new PutPublicCredentialParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutPublicCredentialParamsWithHTTPClient(client *http.Client) *PutPublicCredentialParams {
	var ()
	return &PutPublicCredentialParams{
		HTTPClient: client,
	}
}

/*PutPublicCredentialParams contains all the parameters to send to the API endpoint
for the put public credential operation typically these are written to a http.Request
*/
type PutPublicCredentialParams struct {

	/*Body*/
	Body *models_cloudbreak.CredentialRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put public credential params
func (o *PutPublicCredentialParams) WithTimeout(timeout time.Duration) *PutPublicCredentialParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put public credential params
func (o *PutPublicCredentialParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put public credential params
func (o *PutPublicCredentialParams) WithContext(ctx context.Context) *PutPublicCredentialParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put public credential params
func (o *PutPublicCredentialParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put public credential params
func (o *PutPublicCredentialParams) WithHTTPClient(client *http.Client) *PutPublicCredentialParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put public credential params
func (o *PutPublicCredentialParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put public credential params
func (o *PutPublicCredentialParams) WithBody(body *models_cloudbreak.CredentialRequest) *PutPublicCredentialParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put public credential params
func (o *PutPublicCredentialParams) SetBody(body *models_cloudbreak.CredentialRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PutPublicCredentialParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body == nil {
		o.Body = new(models_cloudbreak.CredentialRequest)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
