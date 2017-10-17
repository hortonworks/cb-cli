// Code generated by go-swagger; DO NOT EDIT.

package v1networks

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

// NewGetPublicsNetworkParams creates a new GetPublicsNetworkParams object
// with the default values initialized.
func NewGetPublicsNetworkParams() *GetPublicsNetworkParams {

	return &GetPublicsNetworkParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPublicsNetworkParamsWithTimeout creates a new GetPublicsNetworkParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPublicsNetworkParamsWithTimeout(timeout time.Duration) *GetPublicsNetworkParams {

	return &GetPublicsNetworkParams{

		timeout: timeout,
	}
}

// NewGetPublicsNetworkParamsWithContext creates a new GetPublicsNetworkParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPublicsNetworkParamsWithContext(ctx context.Context) *GetPublicsNetworkParams {

	return &GetPublicsNetworkParams{

		Context: ctx,
	}
}

// NewGetPublicsNetworkParamsWithHTTPClient creates a new GetPublicsNetworkParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPublicsNetworkParamsWithHTTPClient(client *http.Client) *GetPublicsNetworkParams {

	return &GetPublicsNetworkParams{
		HTTPClient: client,
	}
}

/*GetPublicsNetworkParams contains all the parameters to send to the API endpoint
for the get publics network operation typically these are written to a http.Request
*/
type GetPublicsNetworkParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get publics network params
func (o *GetPublicsNetworkParams) WithTimeout(timeout time.Duration) *GetPublicsNetworkParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get publics network params
func (o *GetPublicsNetworkParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get publics network params
func (o *GetPublicsNetworkParams) WithContext(ctx context.Context) *GetPublicsNetworkParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get publics network params
func (o *GetPublicsNetworkParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get publics network params
func (o *GetPublicsNetworkParams) WithHTTPClient(client *http.Client) *GetPublicsNetworkParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get publics network params
func (o *GetPublicsNetworkParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetPublicsNetworkParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
