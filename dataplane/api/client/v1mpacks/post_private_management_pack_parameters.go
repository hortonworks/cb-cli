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

	model "github.com/hortonworks/cb-cli/dataplane/api/model"
)

// NewPostPrivateManagementPackParams creates a new PostPrivateManagementPackParams object
// with the default values initialized.
func NewPostPrivateManagementPackParams() *PostPrivateManagementPackParams {
	var ()
	return &PostPrivateManagementPackParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostPrivateManagementPackParamsWithTimeout creates a new PostPrivateManagementPackParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostPrivateManagementPackParamsWithTimeout(timeout time.Duration) *PostPrivateManagementPackParams {
	var ()
	return &PostPrivateManagementPackParams{

		timeout: timeout,
	}
}

// NewPostPrivateManagementPackParamsWithContext creates a new PostPrivateManagementPackParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostPrivateManagementPackParamsWithContext(ctx context.Context) *PostPrivateManagementPackParams {
	var ()
	return &PostPrivateManagementPackParams{

		Context: ctx,
	}
}

// NewPostPrivateManagementPackParamsWithHTTPClient creates a new PostPrivateManagementPackParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostPrivateManagementPackParamsWithHTTPClient(client *http.Client) *PostPrivateManagementPackParams {
	var ()
	return &PostPrivateManagementPackParams{
		HTTPClient: client,
	}
}

/*PostPrivateManagementPackParams contains all the parameters to send to the API endpoint
for the post private management pack operation typically these are written to a http.Request
*/
type PostPrivateManagementPackParams struct {

	/*Body*/
	Body *model.ManagementPackRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post private management pack params
func (o *PostPrivateManagementPackParams) WithTimeout(timeout time.Duration) *PostPrivateManagementPackParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post private management pack params
func (o *PostPrivateManagementPackParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post private management pack params
func (o *PostPrivateManagementPackParams) WithContext(ctx context.Context) *PostPrivateManagementPackParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post private management pack params
func (o *PostPrivateManagementPackParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post private management pack params
func (o *PostPrivateManagementPackParams) WithHTTPClient(client *http.Client) *PostPrivateManagementPackParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post private management pack params
func (o *PostPrivateManagementPackParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post private management pack params
func (o *PostPrivateManagementPackParams) WithBody(body *model.ManagementPackRequest) *PostPrivateManagementPackParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post private management pack params
func (o *PostPrivateManagementPackParams) SetBody(body *model.ManagementPackRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostPrivateManagementPackParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
