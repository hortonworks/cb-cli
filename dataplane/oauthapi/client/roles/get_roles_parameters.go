// Code generated by go-swagger; DO NOT EDIT.

package roles

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

	"github.com/hortonworks/cb-cli/dataplane/oauthapi/model"
)

// NewGetRolesParams creates a new GetRolesParams object
// with the default values initialized.
func NewGetRolesParams() *GetRolesParams {
	var ()
	return &GetRolesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRolesParamsWithTimeout creates a new GetRolesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRolesParamsWithTimeout(timeout time.Duration) *GetRolesParams {
	var ()
	return &GetRolesParams{

		timeout: timeout,
	}
}

// NewGetRolesParamsWithContext creates a new GetRolesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRolesParamsWithContext(ctx context.Context) *GetRolesParams {
	var ()
	return &GetRolesParams{

		Context: ctx,
	}
}

// NewGetRolesParamsWithHTTPClient creates a new GetRolesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRolesParamsWithHTTPClient(client *http.Client) *GetRolesParams {
	var ()
	return &GetRolesParams{
		HTTPClient: client,
	}
}

/*GetRolesParams contains all the parameters to send to the API endpoint
for the get roles operation typically these are written to a http.Request
*/
type GetRolesParams struct {

	/*Body*/
	Body *model.UserContext

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get roles params
func (o *GetRolesParams) WithTimeout(timeout time.Duration) *GetRolesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get roles params
func (o *GetRolesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get roles params
func (o *GetRolesParams) WithContext(ctx context.Context) *GetRolesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get roles params
func (o *GetRolesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get roles params
func (o *GetRolesParams) WithHTTPClient(client *http.Client) *GetRolesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get roles params
func (o *GetRolesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the get roles params
func (o *GetRolesParams) WithBody(body *model.UserContext) *GetRolesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the get roles params
func (o *GetRolesParams) SetBody(body *model.UserContext) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *GetRolesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body == nil {
		o.Body = new(model.UserContext)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
