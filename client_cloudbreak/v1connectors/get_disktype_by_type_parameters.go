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
)

// NewGetDisktypeByTypeParams creates a new GetDisktypeByTypeParams object
// with the default values initialized.
func NewGetDisktypeByTypeParams() *GetDisktypeByTypeParams {
	var ()
	return &GetDisktypeByTypeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDisktypeByTypeParamsWithTimeout creates a new GetDisktypeByTypeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDisktypeByTypeParamsWithTimeout(timeout time.Duration) *GetDisktypeByTypeParams {
	var ()
	return &GetDisktypeByTypeParams{

		timeout: timeout,
	}
}

// NewGetDisktypeByTypeParamsWithContext creates a new GetDisktypeByTypeParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDisktypeByTypeParamsWithContext(ctx context.Context) *GetDisktypeByTypeParams {
	var ()
	return &GetDisktypeByTypeParams{

		Context: ctx,
	}
}

// NewGetDisktypeByTypeParamsWithHTTPClient creates a new GetDisktypeByTypeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDisktypeByTypeParamsWithHTTPClient(client *http.Client) *GetDisktypeByTypeParams {
	var ()
	return &GetDisktypeByTypeParams{
		HTTPClient: client,
	}
}

/*GetDisktypeByTypeParams contains all the parameters to send to the API endpoint
for the get disktype by type operation typically these are written to a http.Request
*/
type GetDisktypeByTypeParams struct {

	/*Type*/
	Type string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get disktype by type params
func (o *GetDisktypeByTypeParams) WithTimeout(timeout time.Duration) *GetDisktypeByTypeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get disktype by type params
func (o *GetDisktypeByTypeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get disktype by type params
func (o *GetDisktypeByTypeParams) WithContext(ctx context.Context) *GetDisktypeByTypeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get disktype by type params
func (o *GetDisktypeByTypeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get disktype by type params
func (o *GetDisktypeByTypeParams) WithHTTPClient(client *http.Client) *GetDisktypeByTypeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get disktype by type params
func (o *GetDisktypeByTypeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithType adds the typeVar to the get disktype by type params
func (o *GetDisktypeByTypeParams) WithType(typeVar string) *GetDisktypeByTypeParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the get disktype by type params
func (o *GetDisktypeByTypeParams) SetType(typeVar string) {
	o.Type = typeVar
}

// WriteToRequest writes these params to a swagger request
func (o *GetDisktypeByTypeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param type
	if err := r.SetPathParam("type", o.Type); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
