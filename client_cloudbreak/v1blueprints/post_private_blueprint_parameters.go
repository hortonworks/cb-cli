// Code generated by go-swagger; DO NOT EDIT.

package v1blueprints

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

// NewPostPrivateBlueprintParams creates a new PostPrivateBlueprintParams object
// with the default values initialized.
func NewPostPrivateBlueprintParams() *PostPrivateBlueprintParams {
	var ()
	return &PostPrivateBlueprintParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostPrivateBlueprintParamsWithTimeout creates a new PostPrivateBlueprintParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostPrivateBlueprintParamsWithTimeout(timeout time.Duration) *PostPrivateBlueprintParams {
	var ()
	return &PostPrivateBlueprintParams{

		timeout: timeout,
	}
}

// NewPostPrivateBlueprintParamsWithContext creates a new PostPrivateBlueprintParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostPrivateBlueprintParamsWithContext(ctx context.Context) *PostPrivateBlueprintParams {
	var ()
	return &PostPrivateBlueprintParams{

		Context: ctx,
	}
}

// NewPostPrivateBlueprintParamsWithHTTPClient creates a new PostPrivateBlueprintParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostPrivateBlueprintParamsWithHTTPClient(client *http.Client) *PostPrivateBlueprintParams {
	var ()
	return &PostPrivateBlueprintParams{
		HTTPClient: client,
	}
}

/*PostPrivateBlueprintParams contains all the parameters to send to the API endpoint
for the post private blueprint operation typically these are written to a http.Request
*/
type PostPrivateBlueprintParams struct {

	/*Body*/
	Body *models_cloudbreak.BlueprintRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post private blueprint params
func (o *PostPrivateBlueprintParams) WithTimeout(timeout time.Duration) *PostPrivateBlueprintParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post private blueprint params
func (o *PostPrivateBlueprintParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post private blueprint params
func (o *PostPrivateBlueprintParams) WithContext(ctx context.Context) *PostPrivateBlueprintParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post private blueprint params
func (o *PostPrivateBlueprintParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post private blueprint params
func (o *PostPrivateBlueprintParams) WithHTTPClient(client *http.Client) *PostPrivateBlueprintParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post private blueprint params
func (o *PostPrivateBlueprintParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post private blueprint params
func (o *PostPrivateBlueprintParams) WithBody(body *models_cloudbreak.BlueprintRequest) *PostPrivateBlueprintParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post private blueprint params
func (o *PostPrivateBlueprintParams) SetBody(body *models_cloudbreak.BlueprintRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostPrivateBlueprintParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body == nil {
		o.Body = new(models_cloudbreak.BlueprintRequest)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
