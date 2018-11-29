// Code generated by go-swagger; DO NOT EDIT.

package v1subscriptions

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

// NewSubscribeSubscriptionParams creates a new SubscribeSubscriptionParams object
// with the default values initialized.
func NewSubscribeSubscriptionParams() *SubscribeSubscriptionParams {
	var ()
	return &SubscribeSubscriptionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSubscribeSubscriptionParamsWithTimeout creates a new SubscribeSubscriptionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSubscribeSubscriptionParamsWithTimeout(timeout time.Duration) *SubscribeSubscriptionParams {
	var ()
	return &SubscribeSubscriptionParams{

		timeout: timeout,
	}
}

// NewSubscribeSubscriptionParamsWithContext creates a new SubscribeSubscriptionParams object
// with the default values initialized, and the ability to set a context for a request
func NewSubscribeSubscriptionParamsWithContext(ctx context.Context) *SubscribeSubscriptionParams {
	var ()
	return &SubscribeSubscriptionParams{

		Context: ctx,
	}
}

// NewSubscribeSubscriptionParamsWithHTTPClient creates a new SubscribeSubscriptionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSubscribeSubscriptionParamsWithHTTPClient(client *http.Client) *SubscribeSubscriptionParams {
	var ()
	return &SubscribeSubscriptionParams{
		HTTPClient: client,
	}
}

/*SubscribeSubscriptionParams contains all the parameters to send to the API endpoint
for the subscribe subscription operation typically these are written to a http.Request
*/
type SubscribeSubscriptionParams struct {

	/*Body*/
	Body *model.SubscriptionRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the subscribe subscription params
func (o *SubscribeSubscriptionParams) WithTimeout(timeout time.Duration) *SubscribeSubscriptionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the subscribe subscription params
func (o *SubscribeSubscriptionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the subscribe subscription params
func (o *SubscribeSubscriptionParams) WithContext(ctx context.Context) *SubscribeSubscriptionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the subscribe subscription params
func (o *SubscribeSubscriptionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the subscribe subscription params
func (o *SubscribeSubscriptionParams) WithHTTPClient(client *http.Client) *SubscribeSubscriptionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the subscribe subscription params
func (o *SubscribeSubscriptionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the subscribe subscription params
func (o *SubscribeSubscriptionParams) WithBody(body *model.SubscriptionRequest) *SubscribeSubscriptionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the subscribe subscription params
func (o *SubscribeSubscriptionParams) SetBody(body *model.SubscriptionRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *SubscribeSubscriptionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
