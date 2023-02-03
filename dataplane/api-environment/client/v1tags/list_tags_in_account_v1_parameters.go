// Code generated by go-swagger; DO NOT EDIT.

package v1tags

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

// NewListTagsInAccountV1Params creates a new ListTagsInAccountV1Params object
// with the default values initialized.
func NewListTagsInAccountV1Params() *ListTagsInAccountV1Params {
	var ()
	return &ListTagsInAccountV1Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewListTagsInAccountV1ParamsWithTimeout creates a new ListTagsInAccountV1Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewListTagsInAccountV1ParamsWithTimeout(timeout time.Duration) *ListTagsInAccountV1Params {
	var ()
	return &ListTagsInAccountV1Params{

		timeout: timeout,
	}
}

// NewListTagsInAccountV1ParamsWithContext creates a new ListTagsInAccountV1Params object
// with the default values initialized, and the ability to set a context for a request
func NewListTagsInAccountV1ParamsWithContext(ctx context.Context) *ListTagsInAccountV1Params {
	var ()
	return &ListTagsInAccountV1Params{

		Context: ctx,
	}
}

// NewListTagsInAccountV1ParamsWithHTTPClient creates a new ListTagsInAccountV1Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListTagsInAccountV1ParamsWithHTTPClient(client *http.Client) *ListTagsInAccountV1Params {
	var ()
	return &ListTagsInAccountV1Params{
		HTTPClient: client,
	}
}

/*
ListTagsInAccountV1Params contains all the parameters to send to the API endpoint
for the list tags in account v1 operation typically these are written to a http.Request
*/
type ListTagsInAccountV1Params struct {

	/*AccountID*/
	AccountID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list tags in account v1 params
func (o *ListTagsInAccountV1Params) WithTimeout(timeout time.Duration) *ListTagsInAccountV1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list tags in account v1 params
func (o *ListTagsInAccountV1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list tags in account v1 params
func (o *ListTagsInAccountV1Params) WithContext(ctx context.Context) *ListTagsInAccountV1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list tags in account v1 params
func (o *ListTagsInAccountV1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list tags in account v1 params
func (o *ListTagsInAccountV1Params) WithHTTPClient(client *http.Client) *ListTagsInAccountV1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list tags in account v1 params
func (o *ListTagsInAccountV1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountID adds the accountID to the list tags in account v1 params
func (o *ListTagsInAccountV1Params) WithAccountID(accountID string) *ListTagsInAccountV1Params {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the list tags in account v1 params
func (o *ListTagsInAccountV1Params) SetAccountID(accountID string) {
	o.AccountID = accountID
}

// WriteToRequest writes these params to a swagger request
func (o *ListTagsInAccountV1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param accountId
	if err := r.SetPathParam("accountId", o.AccountID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
