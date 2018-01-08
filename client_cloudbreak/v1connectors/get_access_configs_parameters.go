// Code generated by go-swagger; DO NOT EDIT.

package v1connectors

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

// NewGetAccessConfigsParams creates a new GetAccessConfigsParams object
// with the default values initialized.
func NewGetAccessConfigsParams() *GetAccessConfigsParams {
	var ()
	return &GetAccessConfigsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAccessConfigsParamsWithTimeout creates a new GetAccessConfigsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAccessConfigsParamsWithTimeout(timeout time.Duration) *GetAccessConfigsParams {
	var ()
	return &GetAccessConfigsParams{

		timeout: timeout,
	}
}

// NewGetAccessConfigsParamsWithContext creates a new GetAccessConfigsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAccessConfigsParamsWithContext(ctx context.Context) *GetAccessConfigsParams {
	var ()
	return &GetAccessConfigsParams{

		Context: ctx,
	}
}

// NewGetAccessConfigsParamsWithHTTPClient creates a new GetAccessConfigsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAccessConfigsParamsWithHTTPClient(client *http.Client) *GetAccessConfigsParams {
	var ()
	return &GetAccessConfigsParams{
		HTTPClient: client,
	}
}

/*GetAccessConfigsParams contains all the parameters to send to the API endpoint
for the get access configs operation typically these are written to a http.Request
*/
type GetAccessConfigsParams struct {

	/*Body*/
	Body *models_cloudbreak.PlatformResourceRequestJSON

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get access configs params
func (o *GetAccessConfigsParams) WithTimeout(timeout time.Duration) *GetAccessConfigsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get access configs params
func (o *GetAccessConfigsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get access configs params
func (o *GetAccessConfigsParams) WithContext(ctx context.Context) *GetAccessConfigsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get access configs params
func (o *GetAccessConfigsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get access configs params
func (o *GetAccessConfigsParams) WithHTTPClient(client *http.Client) *GetAccessConfigsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get access configs params
func (o *GetAccessConfigsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the get access configs params
func (o *GetAccessConfigsParams) WithBody(body *models_cloudbreak.PlatformResourceRequestJSON) *GetAccessConfigsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the get access configs params
func (o *GetAccessConfigsParams) SetBody(body *models_cloudbreak.PlatformResourceRequestJSON) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *GetAccessConfigsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body == nil {
		o.Body = new(models_cloudbreak.PlatformResourceRequestJSON)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
