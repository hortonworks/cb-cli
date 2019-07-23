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
)

// NewDeleteWithKerberosDistroXV1ByCrnParams creates a new DeleteWithKerberosDistroXV1ByCrnParams object
// with the default values initialized.
func NewDeleteWithKerberosDistroXV1ByCrnParams() *DeleteWithKerberosDistroXV1ByCrnParams {
	var ()
	return &DeleteWithKerberosDistroXV1ByCrnParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteWithKerberosDistroXV1ByCrnParamsWithTimeout creates a new DeleteWithKerberosDistroXV1ByCrnParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteWithKerberosDistroXV1ByCrnParamsWithTimeout(timeout time.Duration) *DeleteWithKerberosDistroXV1ByCrnParams {
	var ()
	return &DeleteWithKerberosDistroXV1ByCrnParams{

		timeout: timeout,
	}
}

// NewDeleteWithKerberosDistroXV1ByCrnParamsWithContext creates a new DeleteWithKerberosDistroXV1ByCrnParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteWithKerberosDistroXV1ByCrnParamsWithContext(ctx context.Context) *DeleteWithKerberosDistroXV1ByCrnParams {
	var ()
	return &DeleteWithKerberosDistroXV1ByCrnParams{

		Context: ctx,
	}
}

// NewDeleteWithKerberosDistroXV1ByCrnParamsWithHTTPClient creates a new DeleteWithKerberosDistroXV1ByCrnParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteWithKerberosDistroXV1ByCrnParamsWithHTTPClient(client *http.Client) *DeleteWithKerberosDistroXV1ByCrnParams {
	var ()
	return &DeleteWithKerberosDistroXV1ByCrnParams{
		HTTPClient: client,
	}
}

/*DeleteWithKerberosDistroXV1ByCrnParams contains all the parameters to send to the API endpoint
for the delete with kerberos distro x v1 by crn operation typically these are written to a http.Request
*/
type DeleteWithKerberosDistroXV1ByCrnParams struct {

	/*Crn*/
	Crn string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete with kerberos distro x v1 by crn params
func (o *DeleteWithKerberosDistroXV1ByCrnParams) WithTimeout(timeout time.Duration) *DeleteWithKerberosDistroXV1ByCrnParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete with kerberos distro x v1 by crn params
func (o *DeleteWithKerberosDistroXV1ByCrnParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete with kerberos distro x v1 by crn params
func (o *DeleteWithKerberosDistroXV1ByCrnParams) WithContext(ctx context.Context) *DeleteWithKerberosDistroXV1ByCrnParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete with kerberos distro x v1 by crn params
func (o *DeleteWithKerberosDistroXV1ByCrnParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete with kerberos distro x v1 by crn params
func (o *DeleteWithKerberosDistroXV1ByCrnParams) WithHTTPClient(client *http.Client) *DeleteWithKerberosDistroXV1ByCrnParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete with kerberos distro x v1 by crn params
func (o *DeleteWithKerberosDistroXV1ByCrnParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCrn adds the crn to the delete with kerberos distro x v1 by crn params
func (o *DeleteWithKerberosDistroXV1ByCrnParams) WithCrn(crn string) *DeleteWithKerberosDistroXV1ByCrnParams {
	o.SetCrn(crn)
	return o
}

// SetCrn adds the crn to the delete with kerberos distro x v1 by crn params
func (o *DeleteWithKerberosDistroXV1ByCrnParams) SetCrn(crn string) {
	o.Crn = crn
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteWithKerberosDistroXV1ByCrnParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param crn
	if err := r.SetPathParam("crn", o.Crn); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
