// Code generated by go-swagger; DO NOT EDIT.

package v4databaseservers

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

// NewDeleteMultipleDatabaseServersParams creates a new DeleteMultipleDatabaseServersParams object
// with the default values initialized.
func NewDeleteMultipleDatabaseServersParams() *DeleteMultipleDatabaseServersParams {
	var ()
	return &DeleteMultipleDatabaseServersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteMultipleDatabaseServersParamsWithTimeout creates a new DeleteMultipleDatabaseServersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteMultipleDatabaseServersParamsWithTimeout(timeout time.Duration) *DeleteMultipleDatabaseServersParams {
	var ()
	return &DeleteMultipleDatabaseServersParams{

		timeout: timeout,
	}
}

// NewDeleteMultipleDatabaseServersParamsWithContext creates a new DeleteMultipleDatabaseServersParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteMultipleDatabaseServersParamsWithContext(ctx context.Context) *DeleteMultipleDatabaseServersParams {
	var ()
	return &DeleteMultipleDatabaseServersParams{

		Context: ctx,
	}
}

// NewDeleteMultipleDatabaseServersParamsWithHTTPClient creates a new DeleteMultipleDatabaseServersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteMultipleDatabaseServersParamsWithHTTPClient(client *http.Client) *DeleteMultipleDatabaseServersParams {
	var ()
	return &DeleteMultipleDatabaseServersParams{
		HTTPClient: client,
	}
}

/*DeleteMultipleDatabaseServersParams contains all the parameters to send to the API endpoint
for the delete multiple database servers operation typically these are written to a http.Request
*/
type DeleteMultipleDatabaseServersParams struct {

	/*Body*/
	Body []string
	/*EnvironmentID*/
	EnvironmentID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete multiple database servers params
func (o *DeleteMultipleDatabaseServersParams) WithTimeout(timeout time.Duration) *DeleteMultipleDatabaseServersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete multiple database servers params
func (o *DeleteMultipleDatabaseServersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete multiple database servers params
func (o *DeleteMultipleDatabaseServersParams) WithContext(ctx context.Context) *DeleteMultipleDatabaseServersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete multiple database servers params
func (o *DeleteMultipleDatabaseServersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete multiple database servers params
func (o *DeleteMultipleDatabaseServersParams) WithHTTPClient(client *http.Client) *DeleteMultipleDatabaseServersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete multiple database servers params
func (o *DeleteMultipleDatabaseServersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the delete multiple database servers params
func (o *DeleteMultipleDatabaseServersParams) WithBody(body []string) *DeleteMultipleDatabaseServersParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the delete multiple database servers params
func (o *DeleteMultipleDatabaseServersParams) SetBody(body []string) {
	o.Body = body
}

// WithEnvironmentID adds the environmentID to the delete multiple database servers params
func (o *DeleteMultipleDatabaseServersParams) WithEnvironmentID(environmentID string) *DeleteMultipleDatabaseServersParams {
	o.SetEnvironmentID(environmentID)
	return o
}

// SetEnvironmentID adds the environmentId to the delete multiple database servers params
func (o *DeleteMultipleDatabaseServersParams) SetEnvironmentID(environmentID string) {
	o.EnvironmentID = environmentID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteMultipleDatabaseServersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// query param environmentId
	qrEnvironmentID := o.EnvironmentID
	qEnvironmentID := qrEnvironmentID
	if qEnvironmentID != "" {
		if err := r.SetQueryParam("environmentId", qEnvironmentID); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
