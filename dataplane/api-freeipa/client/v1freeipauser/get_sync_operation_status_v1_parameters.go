// Code generated by go-swagger; DO NOT EDIT.

package v1freeipauser

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

// NewGetSyncOperationStatusV1Params creates a new GetSyncOperationStatusV1Params object
// with the default values initialized.
func NewGetSyncOperationStatusV1Params() *GetSyncOperationStatusV1Params {
	var ()
	return &GetSyncOperationStatusV1Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSyncOperationStatusV1ParamsWithTimeout creates a new GetSyncOperationStatusV1Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSyncOperationStatusV1ParamsWithTimeout(timeout time.Duration) *GetSyncOperationStatusV1Params {
	var ()
	return &GetSyncOperationStatusV1Params{

		timeout: timeout,
	}
}

// NewGetSyncOperationStatusV1ParamsWithContext creates a new GetSyncOperationStatusV1Params object
// with the default values initialized, and the ability to set a context for a request
func NewGetSyncOperationStatusV1ParamsWithContext(ctx context.Context) *GetSyncOperationStatusV1Params {
	var ()
	return &GetSyncOperationStatusV1Params{

		Context: ctx,
	}
}

// NewGetSyncOperationStatusV1ParamsWithHTTPClient creates a new GetSyncOperationStatusV1Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSyncOperationStatusV1ParamsWithHTTPClient(client *http.Client) *GetSyncOperationStatusV1Params {
	var ()
	return &GetSyncOperationStatusV1Params{
		HTTPClient: client,
	}
}

/*
GetSyncOperationStatusV1Params contains all the parameters to send to the API endpoint
for the get sync operation status v1 operation typically these are written to a http.Request
*/
type GetSyncOperationStatusV1Params struct {

	/*OperationID*/
	OperationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get sync operation status v1 params
func (o *GetSyncOperationStatusV1Params) WithTimeout(timeout time.Duration) *GetSyncOperationStatusV1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get sync operation status v1 params
func (o *GetSyncOperationStatusV1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get sync operation status v1 params
func (o *GetSyncOperationStatusV1Params) WithContext(ctx context.Context) *GetSyncOperationStatusV1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get sync operation status v1 params
func (o *GetSyncOperationStatusV1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get sync operation status v1 params
func (o *GetSyncOperationStatusV1Params) WithHTTPClient(client *http.Client) *GetSyncOperationStatusV1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get sync operation status v1 params
func (o *GetSyncOperationStatusV1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOperationID adds the operationID to the get sync operation status v1 params
func (o *GetSyncOperationStatusV1Params) WithOperationID(operationID string) *GetSyncOperationStatusV1Params {
	o.SetOperationID(operationID)
	return o
}

// SetOperationID adds the operationId to the get sync operation status v1 params
func (o *GetSyncOperationStatusV1Params) SetOperationID(operationID string) {
	o.OperationID = operationID
}

// WriteToRequest writes these params to a swagger request
func (o *GetSyncOperationStatusV1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param operationId
	qrOperationID := o.OperationID
	qOperationID := qrOperationID
	if qOperationID != "" {
		if err := r.SetQueryParam("operationId", qOperationID); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
