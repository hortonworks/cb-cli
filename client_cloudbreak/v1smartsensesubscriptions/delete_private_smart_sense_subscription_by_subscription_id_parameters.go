// Code generated by go-swagger; DO NOT EDIT.

package v1smartsensesubscriptions

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

// NewDeletePrivateSmartSenseSubscriptionBySubscriptionIDParams creates a new DeletePrivateSmartSenseSubscriptionBySubscriptionIDParams object
// with the default values initialized.
func NewDeletePrivateSmartSenseSubscriptionBySubscriptionIDParams() *DeletePrivateSmartSenseSubscriptionBySubscriptionIDParams {
	var ()
	return &DeletePrivateSmartSenseSubscriptionBySubscriptionIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeletePrivateSmartSenseSubscriptionBySubscriptionIDParamsWithTimeout creates a new DeletePrivateSmartSenseSubscriptionBySubscriptionIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeletePrivateSmartSenseSubscriptionBySubscriptionIDParamsWithTimeout(timeout time.Duration) *DeletePrivateSmartSenseSubscriptionBySubscriptionIDParams {
	var ()
	return &DeletePrivateSmartSenseSubscriptionBySubscriptionIDParams{

		timeout: timeout,
	}
}

// NewDeletePrivateSmartSenseSubscriptionBySubscriptionIDParamsWithContext creates a new DeletePrivateSmartSenseSubscriptionBySubscriptionIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeletePrivateSmartSenseSubscriptionBySubscriptionIDParamsWithContext(ctx context.Context) *DeletePrivateSmartSenseSubscriptionBySubscriptionIDParams {
	var ()
	return &DeletePrivateSmartSenseSubscriptionBySubscriptionIDParams{

		Context: ctx,
	}
}

// NewDeletePrivateSmartSenseSubscriptionBySubscriptionIDParamsWithHTTPClient creates a new DeletePrivateSmartSenseSubscriptionBySubscriptionIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeletePrivateSmartSenseSubscriptionBySubscriptionIDParamsWithHTTPClient(client *http.Client) *DeletePrivateSmartSenseSubscriptionBySubscriptionIDParams {
	var ()
	return &DeletePrivateSmartSenseSubscriptionBySubscriptionIDParams{
		HTTPClient: client,
	}
}

/*DeletePrivateSmartSenseSubscriptionBySubscriptionIDParams contains all the parameters to send to the API endpoint
for the delete private smart sense subscription by subscription Id operation typically these are written to a http.Request
*/
type DeletePrivateSmartSenseSubscriptionBySubscriptionIDParams struct {

	/*SubscriptionID*/
	SubscriptionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete private smart sense subscription by subscription Id params
func (o *DeletePrivateSmartSenseSubscriptionBySubscriptionIDParams) WithTimeout(timeout time.Duration) *DeletePrivateSmartSenseSubscriptionBySubscriptionIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete private smart sense subscription by subscription Id params
func (o *DeletePrivateSmartSenseSubscriptionBySubscriptionIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete private smart sense subscription by subscription Id params
func (o *DeletePrivateSmartSenseSubscriptionBySubscriptionIDParams) WithContext(ctx context.Context) *DeletePrivateSmartSenseSubscriptionBySubscriptionIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete private smart sense subscription by subscription Id params
func (o *DeletePrivateSmartSenseSubscriptionBySubscriptionIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete private smart sense subscription by subscription Id params
func (o *DeletePrivateSmartSenseSubscriptionBySubscriptionIDParams) WithHTTPClient(client *http.Client) *DeletePrivateSmartSenseSubscriptionBySubscriptionIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete private smart sense subscription by subscription Id params
func (o *DeletePrivateSmartSenseSubscriptionBySubscriptionIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSubscriptionID adds the subscriptionID to the delete private smart sense subscription by subscription Id params
func (o *DeletePrivateSmartSenseSubscriptionBySubscriptionIDParams) WithSubscriptionID(subscriptionID string) *DeletePrivateSmartSenseSubscriptionBySubscriptionIDParams {
	o.SetSubscriptionID(subscriptionID)
	return o
}

// SetSubscriptionID adds the subscriptionId to the delete private smart sense subscription by subscription Id params
func (o *DeletePrivateSmartSenseSubscriptionBySubscriptionIDParams) SetSubscriptionID(subscriptionID string) {
	o.SubscriptionID = subscriptionID
}

// WriteToRequest writes these params to a swagger request
func (o *DeletePrivateSmartSenseSubscriptionBySubscriptionIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param subscriptionId
	if err := r.SetPathParam("subscriptionId", o.SubscriptionID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}