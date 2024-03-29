// Code generated by go-swagger; DO NOT EDIT.

package v4diagnostics

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

// NewCollectCmDiagnosticsParams creates a new CollectCmDiagnosticsParams object
// with the default values initialized.
func NewCollectCmDiagnosticsParams() *CollectCmDiagnosticsParams {
	var ()
	return &CollectCmDiagnosticsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCollectCmDiagnosticsParamsWithTimeout creates a new CollectCmDiagnosticsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCollectCmDiagnosticsParamsWithTimeout(timeout time.Duration) *CollectCmDiagnosticsParams {
	var ()
	return &CollectCmDiagnosticsParams{

		timeout: timeout,
	}
}

// NewCollectCmDiagnosticsParamsWithContext creates a new CollectCmDiagnosticsParams object
// with the default values initialized, and the ability to set a context for a request
func NewCollectCmDiagnosticsParamsWithContext(ctx context.Context) *CollectCmDiagnosticsParams {
	var ()
	return &CollectCmDiagnosticsParams{

		Context: ctx,
	}
}

// NewCollectCmDiagnosticsParamsWithHTTPClient creates a new CollectCmDiagnosticsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCollectCmDiagnosticsParamsWithHTTPClient(client *http.Client) *CollectCmDiagnosticsParams {
	var ()
	return &CollectCmDiagnosticsParams{
		HTTPClient: client,
	}
}

/*
CollectCmDiagnosticsParams contains all the parameters to send to the API endpoint
for the collect cm diagnostics operation typically these are written to a http.Request
*/
type CollectCmDiagnosticsParams struct {

	/*Body*/
	Body *model.CmDiagnosticsCollectionRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the collect cm diagnostics params
func (o *CollectCmDiagnosticsParams) WithTimeout(timeout time.Duration) *CollectCmDiagnosticsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the collect cm diagnostics params
func (o *CollectCmDiagnosticsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the collect cm diagnostics params
func (o *CollectCmDiagnosticsParams) WithContext(ctx context.Context) *CollectCmDiagnosticsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the collect cm diagnostics params
func (o *CollectCmDiagnosticsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the collect cm diagnostics params
func (o *CollectCmDiagnosticsParams) WithHTTPClient(client *http.Client) *CollectCmDiagnosticsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the collect cm diagnostics params
func (o *CollectCmDiagnosticsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the collect cm diagnostics params
func (o *CollectCmDiagnosticsParams) WithBody(body *model.CmDiagnosticsCollectionRequest) *CollectCmDiagnosticsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the collect cm diagnostics params
func (o *CollectCmDiagnosticsParams) SetBody(body *model.CmDiagnosticsCollectionRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CollectCmDiagnosticsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
