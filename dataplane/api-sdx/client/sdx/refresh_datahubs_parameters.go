// Code generated by go-swagger; DO NOT EDIT.

package sdx

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

// NewRefreshDatahubsParams creates a new RefreshDatahubsParams object
// with the default values initialized.
func NewRefreshDatahubsParams() *RefreshDatahubsParams {
	var ()
	return &RefreshDatahubsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRefreshDatahubsParamsWithTimeout creates a new RefreshDatahubsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRefreshDatahubsParamsWithTimeout(timeout time.Duration) *RefreshDatahubsParams {
	var ()
	return &RefreshDatahubsParams{

		timeout: timeout,
	}
}

// NewRefreshDatahubsParamsWithContext creates a new RefreshDatahubsParams object
// with the default values initialized, and the ability to set a context for a request
func NewRefreshDatahubsParamsWithContext(ctx context.Context) *RefreshDatahubsParams {
	var ()
	return &RefreshDatahubsParams{

		Context: ctx,
	}
}

// NewRefreshDatahubsParamsWithHTTPClient creates a new RefreshDatahubsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRefreshDatahubsParamsWithHTTPClient(client *http.Client) *RefreshDatahubsParams {
	var ()
	return &RefreshDatahubsParams{
		HTTPClient: client,
	}
}

/*
RefreshDatahubsParams contains all the parameters to send to the API endpoint
for the refresh datahubs operation typically these are written to a http.Request
*/
type RefreshDatahubsParams struct {

	/*DatahubName*/
	DatahubName *string
	/*DatalakeName*/
	DatalakeName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the refresh datahubs params
func (o *RefreshDatahubsParams) WithTimeout(timeout time.Duration) *RefreshDatahubsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the refresh datahubs params
func (o *RefreshDatahubsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the refresh datahubs params
func (o *RefreshDatahubsParams) WithContext(ctx context.Context) *RefreshDatahubsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the refresh datahubs params
func (o *RefreshDatahubsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the refresh datahubs params
func (o *RefreshDatahubsParams) WithHTTPClient(client *http.Client) *RefreshDatahubsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the refresh datahubs params
func (o *RefreshDatahubsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDatahubName adds the datahubName to the refresh datahubs params
func (o *RefreshDatahubsParams) WithDatahubName(datahubName *string) *RefreshDatahubsParams {
	o.SetDatahubName(datahubName)
	return o
}

// SetDatahubName adds the datahubName to the refresh datahubs params
func (o *RefreshDatahubsParams) SetDatahubName(datahubName *string) {
	o.DatahubName = datahubName
}

// WithDatalakeName adds the datalakeName to the refresh datahubs params
func (o *RefreshDatahubsParams) WithDatalakeName(datalakeName string) *RefreshDatahubsParams {
	o.SetDatalakeName(datalakeName)
	return o
}

// SetDatalakeName adds the datalakeName to the refresh datahubs params
func (o *RefreshDatahubsParams) SetDatalakeName(datalakeName string) {
	o.DatalakeName = datalakeName
}

// WriteToRequest writes these params to a swagger request
func (o *RefreshDatahubsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DatahubName != nil {

		// query param datahubName
		var qrDatahubName string
		if o.DatahubName != nil {
			qrDatahubName = *o.DatahubName
		}
		qDatahubName := qrDatahubName
		if qDatahubName != "" {
			if err := r.SetQueryParam("datahubName", qDatahubName); err != nil {
				return err
			}
		}

	}

	// path param datalakeName
	if err := r.SetPathParam("datalakeName", o.DatalakeName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
