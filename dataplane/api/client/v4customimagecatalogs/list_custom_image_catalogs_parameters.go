// Code generated by go-swagger; DO NOT EDIT.

package v4customimagecatalogs

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

// NewListCustomImageCatalogsParams creates a new ListCustomImageCatalogsParams object
// with the default values initialized.
func NewListCustomImageCatalogsParams() *ListCustomImageCatalogsParams {
	var ()
	return &ListCustomImageCatalogsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListCustomImageCatalogsParamsWithTimeout creates a new ListCustomImageCatalogsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListCustomImageCatalogsParamsWithTimeout(timeout time.Duration) *ListCustomImageCatalogsParams {
	var ()
	return &ListCustomImageCatalogsParams{

		timeout: timeout,
	}
}

// NewListCustomImageCatalogsParamsWithContext creates a new ListCustomImageCatalogsParams object
// with the default values initialized, and the ability to set a context for a request
func NewListCustomImageCatalogsParamsWithContext(ctx context.Context) *ListCustomImageCatalogsParams {
	var ()
	return &ListCustomImageCatalogsParams{

		Context: ctx,
	}
}

// NewListCustomImageCatalogsParamsWithHTTPClient creates a new ListCustomImageCatalogsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListCustomImageCatalogsParamsWithHTTPClient(client *http.Client) *ListCustomImageCatalogsParams {
	var ()
	return &ListCustomImageCatalogsParams{
		HTTPClient: client,
	}
}

/*
ListCustomImageCatalogsParams contains all the parameters to send to the API endpoint
for the list custom image catalogs operation typically these are written to a http.Request
*/
type ListCustomImageCatalogsParams struct {

	/*AccountID*/
	AccountID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list custom image catalogs params
func (o *ListCustomImageCatalogsParams) WithTimeout(timeout time.Duration) *ListCustomImageCatalogsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list custom image catalogs params
func (o *ListCustomImageCatalogsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list custom image catalogs params
func (o *ListCustomImageCatalogsParams) WithContext(ctx context.Context) *ListCustomImageCatalogsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list custom image catalogs params
func (o *ListCustomImageCatalogsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list custom image catalogs params
func (o *ListCustomImageCatalogsParams) WithHTTPClient(client *http.Client) *ListCustomImageCatalogsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list custom image catalogs params
func (o *ListCustomImageCatalogsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountID adds the accountID to the list custom image catalogs params
func (o *ListCustomImageCatalogsParams) WithAccountID(accountID *string) *ListCustomImageCatalogsParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the list custom image catalogs params
func (o *ListCustomImageCatalogsParams) SetAccountID(accountID *string) {
	o.AccountID = accountID
}

// WriteToRequest writes these params to a swagger request
func (o *ListCustomImageCatalogsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AccountID != nil {

		// query param accountId
		var qrAccountID string
		if o.AccountID != nil {
			qrAccountID = *o.AccountID
		}
		qAccountID := qrAccountID
		if qAccountID != "" {
			if err := r.SetQueryParam("accountId", qAccountID); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
