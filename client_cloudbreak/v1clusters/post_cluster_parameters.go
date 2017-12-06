// Code generated by go-swagger; DO NOT EDIT.

package v1clusters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/hortonworks/cb-cli/models_cloudbreak"
)

// NewPostClusterParams creates a new PostClusterParams object
// with the default values initialized.
func NewPostClusterParams() *PostClusterParams {
	var ()
	return &PostClusterParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostClusterParamsWithTimeout creates a new PostClusterParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostClusterParamsWithTimeout(timeout time.Duration) *PostClusterParams {
	var ()
	return &PostClusterParams{

		timeout: timeout,
	}
}

// NewPostClusterParamsWithContext creates a new PostClusterParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostClusterParamsWithContext(ctx context.Context) *PostClusterParams {
	var ()
	return &PostClusterParams{

		Context: ctx,
	}
}

// NewPostClusterParamsWithHTTPClient creates a new PostClusterParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostClusterParamsWithHTTPClient(client *http.Client) *PostClusterParams {
	var ()
	return &PostClusterParams{
		HTTPClient: client,
	}
}

/*PostClusterParams contains all the parameters to send to the API endpoint
for the post cluster operation typically these are written to a http.Request
*/
type PostClusterParams struct {

	/*Body*/
	Body *models_cloudbreak.ClusterRequest
	/*ID*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post cluster params
func (o *PostClusterParams) WithTimeout(timeout time.Duration) *PostClusterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post cluster params
func (o *PostClusterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post cluster params
func (o *PostClusterParams) WithContext(ctx context.Context) *PostClusterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post cluster params
func (o *PostClusterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post cluster params
func (o *PostClusterParams) WithHTTPClient(client *http.Client) *PostClusterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post cluster params
func (o *PostClusterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post cluster params
func (o *PostClusterParams) WithBody(body *models_cloudbreak.ClusterRequest) *PostClusterParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post cluster params
func (o *PostClusterParams) SetBody(body *models_cloudbreak.ClusterRequest) {
	o.Body = body
}

// WithID adds the id to the post cluster params
func (o *PostClusterParams) WithID(id int64) *PostClusterParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the post cluster params
func (o *PostClusterParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PostClusterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body == nil {
		o.Body = new(models_cloudbreak.ClusterRequest)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
