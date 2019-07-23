// Code generated by go-swagger; DO NOT EDIT.

package v1distrox

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api/model"
)

// NewPostDistroXForBlueprintV1ByNameParams creates a new PostDistroXForBlueprintV1ByNameParams object
// with the default values initialized.
func NewPostDistroXForBlueprintV1ByNameParams() *PostDistroXForBlueprintV1ByNameParams {
	var ()
	return &PostDistroXForBlueprintV1ByNameParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostDistroXForBlueprintV1ByNameParamsWithTimeout creates a new PostDistroXForBlueprintV1ByNameParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostDistroXForBlueprintV1ByNameParamsWithTimeout(timeout time.Duration) *PostDistroXForBlueprintV1ByNameParams {
	var ()
	return &PostDistroXForBlueprintV1ByNameParams{

		timeout: timeout,
	}
}

// NewPostDistroXForBlueprintV1ByNameParamsWithContext creates a new PostDistroXForBlueprintV1ByNameParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostDistroXForBlueprintV1ByNameParamsWithContext(ctx context.Context) *PostDistroXForBlueprintV1ByNameParams {
	var ()
	return &PostDistroXForBlueprintV1ByNameParams{

		Context: ctx,
	}
}

// NewPostDistroXForBlueprintV1ByNameParamsWithHTTPClient creates a new PostDistroXForBlueprintV1ByNameParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostDistroXForBlueprintV1ByNameParamsWithHTTPClient(client *http.Client) *PostDistroXForBlueprintV1ByNameParams {
	var ()
	return &PostDistroXForBlueprintV1ByNameParams{
		HTTPClient: client,
	}
}

/*PostDistroXForBlueprintV1ByNameParams contains all the parameters to send to the API endpoint
for the post distro x for blueprint v1 by name operation typically these are written to a http.Request
*/
type PostDistroXForBlueprintV1ByNameParams struct {

	/*Body*/
	Body *model.DistroXV1Request
	/*Name*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post distro x for blueprint v1 by name params
func (o *PostDistroXForBlueprintV1ByNameParams) WithTimeout(timeout time.Duration) *PostDistroXForBlueprintV1ByNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post distro x for blueprint v1 by name params
func (o *PostDistroXForBlueprintV1ByNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post distro x for blueprint v1 by name params
func (o *PostDistroXForBlueprintV1ByNameParams) WithContext(ctx context.Context) *PostDistroXForBlueprintV1ByNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post distro x for blueprint v1 by name params
func (o *PostDistroXForBlueprintV1ByNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post distro x for blueprint v1 by name params
func (o *PostDistroXForBlueprintV1ByNameParams) WithHTTPClient(client *http.Client) *PostDistroXForBlueprintV1ByNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post distro x for blueprint v1 by name params
func (o *PostDistroXForBlueprintV1ByNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post distro x for blueprint v1 by name params
func (o *PostDistroXForBlueprintV1ByNameParams) WithBody(body *model.DistroXV1Request) *PostDistroXForBlueprintV1ByNameParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post distro x for blueprint v1 by name params
func (o *PostDistroXForBlueprintV1ByNameParams) SetBody(body *model.DistroXV1Request) {
	o.Body = body
}

// WithName adds the name to the post distro x for blueprint v1 by name params
func (o *PostDistroXForBlueprintV1ByNameParams) WithName(name string) *PostDistroXForBlueprintV1ByNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the post distro x for blueprint v1 by name params
func (o *PostDistroXForBlueprintV1ByNameParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *PostDistroXForBlueprintV1ByNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
