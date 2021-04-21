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

	model "github.com/hortonworks/cb-cli/dataplane/api/model"
)

// NewPostDistroXInternalV1Params creates a new PostDistroXInternalV1Params object
// with the default values initialized.
func NewPostDistroXInternalV1Params() *PostDistroXInternalV1Params {
	var ()
	return &PostDistroXInternalV1Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostDistroXInternalV1ParamsWithTimeout creates a new PostDistroXInternalV1Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostDistroXInternalV1ParamsWithTimeout(timeout time.Duration) *PostDistroXInternalV1Params {
	var ()
	return &PostDistroXInternalV1Params{

		timeout: timeout,
	}
}

// NewPostDistroXInternalV1ParamsWithContext creates a new PostDistroXInternalV1Params object
// with the default values initialized, and the ability to set a context for a request
func NewPostDistroXInternalV1ParamsWithContext(ctx context.Context) *PostDistroXInternalV1Params {
	var ()
	return &PostDistroXInternalV1Params{

		Context: ctx,
	}
}

// NewPostDistroXInternalV1ParamsWithHTTPClient creates a new PostDistroXInternalV1Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostDistroXInternalV1ParamsWithHTTPClient(client *http.Client) *PostDistroXInternalV1Params {
	var ()
	return &PostDistroXInternalV1Params{
		HTTPClient: client,
	}
}

/*PostDistroXInternalV1Params contains all the parameters to send to the API endpoint
for the post distro x internal v1 operation typically these are written to a http.Request
*/
type PostDistroXInternalV1Params struct {

	/*AccountID*/
	AccountID *string
	/*Body*/
	Body *model.DistroXV1Request

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post distro x internal v1 params
func (o *PostDistroXInternalV1Params) WithTimeout(timeout time.Duration) *PostDistroXInternalV1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post distro x internal v1 params
func (o *PostDistroXInternalV1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post distro x internal v1 params
func (o *PostDistroXInternalV1Params) WithContext(ctx context.Context) *PostDistroXInternalV1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post distro x internal v1 params
func (o *PostDistroXInternalV1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post distro x internal v1 params
func (o *PostDistroXInternalV1Params) WithHTTPClient(client *http.Client) *PostDistroXInternalV1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post distro x internal v1 params
func (o *PostDistroXInternalV1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountID adds the accountID to the post distro x internal v1 params
func (o *PostDistroXInternalV1Params) WithAccountID(accountID *string) *PostDistroXInternalV1Params {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the post distro x internal v1 params
func (o *PostDistroXInternalV1Params) SetAccountID(accountID *string) {
	o.AccountID = accountID
}

// WithBody adds the body to the post distro x internal v1 params
func (o *PostDistroXInternalV1Params) WithBody(body *model.DistroXV1Request) *PostDistroXInternalV1Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post distro x internal v1 params
func (o *PostDistroXInternalV1Params) SetBody(body *model.DistroXV1Request) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostDistroXInternalV1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
