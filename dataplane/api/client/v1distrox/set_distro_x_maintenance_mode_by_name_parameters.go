// Code generated by go-swagger; DO NOT EDIT.

package v1distrox

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

// NewSetDistroXMaintenanceModeByNameParams creates a new SetDistroXMaintenanceModeByNameParams object
// with the default values initialized.
func NewSetDistroXMaintenanceModeByNameParams() *SetDistroXMaintenanceModeByNameParams {
	var ()
	return &SetDistroXMaintenanceModeByNameParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSetDistroXMaintenanceModeByNameParamsWithTimeout creates a new SetDistroXMaintenanceModeByNameParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSetDistroXMaintenanceModeByNameParamsWithTimeout(timeout time.Duration) *SetDistroXMaintenanceModeByNameParams {
	var ()
	return &SetDistroXMaintenanceModeByNameParams{

		timeout: timeout,
	}
}

// NewSetDistroXMaintenanceModeByNameParamsWithContext creates a new SetDistroXMaintenanceModeByNameParams object
// with the default values initialized, and the ability to set a context for a request
func NewSetDistroXMaintenanceModeByNameParamsWithContext(ctx context.Context) *SetDistroXMaintenanceModeByNameParams {
	var ()
	return &SetDistroXMaintenanceModeByNameParams{

		Context: ctx,
	}
}

// NewSetDistroXMaintenanceModeByNameParamsWithHTTPClient creates a new SetDistroXMaintenanceModeByNameParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSetDistroXMaintenanceModeByNameParamsWithHTTPClient(client *http.Client) *SetDistroXMaintenanceModeByNameParams {
	var ()
	return &SetDistroXMaintenanceModeByNameParams{
		HTTPClient: client,
	}
}

/*SetDistroXMaintenanceModeByNameParams contains all the parameters to send to the API endpoint
for the set distro x maintenance mode by name operation typically these are written to a http.Request
*/
type SetDistroXMaintenanceModeByNameParams struct {

	/*Body*/
	Body *model.DistroXMaintenanceModeV1Request
	/*Name*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the set distro x maintenance mode by name params
func (o *SetDistroXMaintenanceModeByNameParams) WithTimeout(timeout time.Duration) *SetDistroXMaintenanceModeByNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the set distro x maintenance mode by name params
func (o *SetDistroXMaintenanceModeByNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the set distro x maintenance mode by name params
func (o *SetDistroXMaintenanceModeByNameParams) WithContext(ctx context.Context) *SetDistroXMaintenanceModeByNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the set distro x maintenance mode by name params
func (o *SetDistroXMaintenanceModeByNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the set distro x maintenance mode by name params
func (o *SetDistroXMaintenanceModeByNameParams) WithHTTPClient(client *http.Client) *SetDistroXMaintenanceModeByNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the set distro x maintenance mode by name params
func (o *SetDistroXMaintenanceModeByNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the set distro x maintenance mode by name params
func (o *SetDistroXMaintenanceModeByNameParams) WithBody(body *model.DistroXMaintenanceModeV1Request) *SetDistroXMaintenanceModeByNameParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the set distro x maintenance mode by name params
func (o *SetDistroXMaintenanceModeByNameParams) SetBody(body *model.DistroXMaintenanceModeV1Request) {
	o.Body = body
}

// WithName adds the name to the set distro x maintenance mode by name params
func (o *SetDistroXMaintenanceModeByNameParams) WithName(name string) *SetDistroXMaintenanceModeByNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the set distro x maintenance mode by name params
func (o *SetDistroXMaintenanceModeByNameParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *SetDistroXMaintenanceModeByNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
