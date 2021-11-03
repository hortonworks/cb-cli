// Code generated by go-swagger; DO NOT EDIT.

package sdx

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

	model "github.com/hortonworks/cb-cli/dataplane/api-sdx/model"
)

// NewAttachRecipesByCrnParams creates a new AttachRecipesByCrnParams object
// with the default values initialized.
func NewAttachRecipesByCrnParams() *AttachRecipesByCrnParams {
	var ()
	return &AttachRecipesByCrnParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAttachRecipesByCrnParamsWithTimeout creates a new AttachRecipesByCrnParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAttachRecipesByCrnParamsWithTimeout(timeout time.Duration) *AttachRecipesByCrnParams {
	var ()
	return &AttachRecipesByCrnParams{

		timeout: timeout,
	}
}

// NewAttachRecipesByCrnParamsWithContext creates a new AttachRecipesByCrnParams object
// with the default values initialized, and the ability to set a context for a request
func NewAttachRecipesByCrnParamsWithContext(ctx context.Context) *AttachRecipesByCrnParams {
	var ()
	return &AttachRecipesByCrnParams{

		Context: ctx,
	}
}

// NewAttachRecipesByCrnParamsWithHTTPClient creates a new AttachRecipesByCrnParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAttachRecipesByCrnParamsWithHTTPClient(client *http.Client) *AttachRecipesByCrnParams {
	var ()
	return &AttachRecipesByCrnParams{
		HTTPClient: client,
	}
}

/*AttachRecipesByCrnParams contains all the parameters to send to the API endpoint
for the attach recipes by crn operation typically these are written to a http.Request
*/
type AttachRecipesByCrnParams struct {

	/*Body*/
	Body *model.AttachRecipeV4Request
	/*Crn*/
	Crn string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the attach recipes by crn params
func (o *AttachRecipesByCrnParams) WithTimeout(timeout time.Duration) *AttachRecipesByCrnParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the attach recipes by crn params
func (o *AttachRecipesByCrnParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the attach recipes by crn params
func (o *AttachRecipesByCrnParams) WithContext(ctx context.Context) *AttachRecipesByCrnParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the attach recipes by crn params
func (o *AttachRecipesByCrnParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the attach recipes by crn params
func (o *AttachRecipesByCrnParams) WithHTTPClient(client *http.Client) *AttachRecipesByCrnParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the attach recipes by crn params
func (o *AttachRecipesByCrnParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the attach recipes by crn params
func (o *AttachRecipesByCrnParams) WithBody(body *model.AttachRecipeV4Request) *AttachRecipesByCrnParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the attach recipes by crn params
func (o *AttachRecipesByCrnParams) SetBody(body *model.AttachRecipeV4Request) {
	o.Body = body
}

// WithCrn adds the crn to the attach recipes by crn params
func (o *AttachRecipesByCrnParams) WithCrn(crn string) *AttachRecipesByCrnParams {
	o.SetCrn(crn)
	return o
}

// SetCrn adds the crn to the attach recipes by crn params
func (o *AttachRecipesByCrnParams) SetCrn(crn string) {
	o.Crn = crn
}

// WriteToRequest writes these params to a swagger request
func (o *AttachRecipesByCrnParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
