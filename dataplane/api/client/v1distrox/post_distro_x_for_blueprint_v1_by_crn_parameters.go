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

// NewPostDistroXForBlueprintV1ByCrnParams creates a new PostDistroXForBlueprintV1ByCrnParams object
// with the default values initialized.
func NewPostDistroXForBlueprintV1ByCrnParams() *PostDistroXForBlueprintV1ByCrnParams {
	var ()
	return &PostDistroXForBlueprintV1ByCrnParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostDistroXForBlueprintV1ByCrnParamsWithTimeout creates a new PostDistroXForBlueprintV1ByCrnParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostDistroXForBlueprintV1ByCrnParamsWithTimeout(timeout time.Duration) *PostDistroXForBlueprintV1ByCrnParams {
	var ()
	return &PostDistroXForBlueprintV1ByCrnParams{

		timeout: timeout,
	}
}

// NewPostDistroXForBlueprintV1ByCrnParamsWithContext creates a new PostDistroXForBlueprintV1ByCrnParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostDistroXForBlueprintV1ByCrnParamsWithContext(ctx context.Context) *PostDistroXForBlueprintV1ByCrnParams {
	var ()
	return &PostDistroXForBlueprintV1ByCrnParams{

		Context: ctx,
	}
}

// NewPostDistroXForBlueprintV1ByCrnParamsWithHTTPClient creates a new PostDistroXForBlueprintV1ByCrnParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostDistroXForBlueprintV1ByCrnParamsWithHTTPClient(client *http.Client) *PostDistroXForBlueprintV1ByCrnParams {
	var ()
	return &PostDistroXForBlueprintV1ByCrnParams{
		HTTPClient: client,
	}
}

/*PostDistroXForBlueprintV1ByCrnParams contains all the parameters to send to the API endpoint
for the post distro x for blueprint v1 by crn operation typically these are written to a http.Request
*/
type PostDistroXForBlueprintV1ByCrnParams struct {

	/*Body*/
	Body *model.DistroXV1Request
	/*Crn*/
	Crn string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post distro x for blueprint v1 by crn params
func (o *PostDistroXForBlueprintV1ByCrnParams) WithTimeout(timeout time.Duration) *PostDistroXForBlueprintV1ByCrnParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post distro x for blueprint v1 by crn params
func (o *PostDistroXForBlueprintV1ByCrnParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post distro x for blueprint v1 by crn params
func (o *PostDistroXForBlueprintV1ByCrnParams) WithContext(ctx context.Context) *PostDistroXForBlueprintV1ByCrnParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post distro x for blueprint v1 by crn params
func (o *PostDistroXForBlueprintV1ByCrnParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post distro x for blueprint v1 by crn params
func (o *PostDistroXForBlueprintV1ByCrnParams) WithHTTPClient(client *http.Client) *PostDistroXForBlueprintV1ByCrnParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post distro x for blueprint v1 by crn params
func (o *PostDistroXForBlueprintV1ByCrnParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post distro x for blueprint v1 by crn params
func (o *PostDistroXForBlueprintV1ByCrnParams) WithBody(body *model.DistroXV1Request) *PostDistroXForBlueprintV1ByCrnParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post distro x for blueprint v1 by crn params
func (o *PostDistroXForBlueprintV1ByCrnParams) SetBody(body *model.DistroXV1Request) {
	o.Body = body
}

// WithCrn adds the crn to the post distro x for blueprint v1 by crn params
func (o *PostDistroXForBlueprintV1ByCrnParams) WithCrn(crn string) *PostDistroXForBlueprintV1ByCrnParams {
	o.SetCrn(crn)
	return o
}

// SetCrn adds the crn to the post distro x for blueprint v1 by crn params
func (o *PostDistroXForBlueprintV1ByCrnParams) SetCrn(crn string) {
	o.Crn = crn
}

// WriteToRequest writes these params to a swagger request
func (o *PostDistroXForBlueprintV1ByCrnParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param crn
	if err := r.SetPathParam("crn", o.Crn); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
