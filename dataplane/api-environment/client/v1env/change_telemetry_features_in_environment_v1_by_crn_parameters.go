// Code generated by go-swagger; DO NOT EDIT.

package v1env

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

	model "github.com/hortonworks/cb-cli/dataplane/api-environment/model"
)

// NewChangeTelemetryFeaturesInEnvironmentV1ByCrnParams creates a new ChangeTelemetryFeaturesInEnvironmentV1ByCrnParams object
// with the default values initialized.
func NewChangeTelemetryFeaturesInEnvironmentV1ByCrnParams() *ChangeTelemetryFeaturesInEnvironmentV1ByCrnParams {
	var ()
	return &ChangeTelemetryFeaturesInEnvironmentV1ByCrnParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewChangeTelemetryFeaturesInEnvironmentV1ByCrnParamsWithTimeout creates a new ChangeTelemetryFeaturesInEnvironmentV1ByCrnParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewChangeTelemetryFeaturesInEnvironmentV1ByCrnParamsWithTimeout(timeout time.Duration) *ChangeTelemetryFeaturesInEnvironmentV1ByCrnParams {
	var ()
	return &ChangeTelemetryFeaturesInEnvironmentV1ByCrnParams{

		timeout: timeout,
	}
}

// NewChangeTelemetryFeaturesInEnvironmentV1ByCrnParamsWithContext creates a new ChangeTelemetryFeaturesInEnvironmentV1ByCrnParams object
// with the default values initialized, and the ability to set a context for a request
func NewChangeTelemetryFeaturesInEnvironmentV1ByCrnParamsWithContext(ctx context.Context) *ChangeTelemetryFeaturesInEnvironmentV1ByCrnParams {
	var ()
	return &ChangeTelemetryFeaturesInEnvironmentV1ByCrnParams{

		Context: ctx,
	}
}

// NewChangeTelemetryFeaturesInEnvironmentV1ByCrnParamsWithHTTPClient creates a new ChangeTelemetryFeaturesInEnvironmentV1ByCrnParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewChangeTelemetryFeaturesInEnvironmentV1ByCrnParamsWithHTTPClient(client *http.Client) *ChangeTelemetryFeaturesInEnvironmentV1ByCrnParams {
	var ()
	return &ChangeTelemetryFeaturesInEnvironmentV1ByCrnParams{
		HTTPClient: client,
	}
}

/*
ChangeTelemetryFeaturesInEnvironmentV1ByCrnParams contains all the parameters to send to the API endpoint
for the change telemetry features in environment v1 by crn operation typically these are written to a http.Request
*/
type ChangeTelemetryFeaturesInEnvironmentV1ByCrnParams struct {

	/*Body*/
	Body *model.FeaturesRequest
	/*Crn*/
	Crn string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the change telemetry features in environment v1 by crn params
func (o *ChangeTelemetryFeaturesInEnvironmentV1ByCrnParams) WithTimeout(timeout time.Duration) *ChangeTelemetryFeaturesInEnvironmentV1ByCrnParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the change telemetry features in environment v1 by crn params
func (o *ChangeTelemetryFeaturesInEnvironmentV1ByCrnParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the change telemetry features in environment v1 by crn params
func (o *ChangeTelemetryFeaturesInEnvironmentV1ByCrnParams) WithContext(ctx context.Context) *ChangeTelemetryFeaturesInEnvironmentV1ByCrnParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the change telemetry features in environment v1 by crn params
func (o *ChangeTelemetryFeaturesInEnvironmentV1ByCrnParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the change telemetry features in environment v1 by crn params
func (o *ChangeTelemetryFeaturesInEnvironmentV1ByCrnParams) WithHTTPClient(client *http.Client) *ChangeTelemetryFeaturesInEnvironmentV1ByCrnParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the change telemetry features in environment v1 by crn params
func (o *ChangeTelemetryFeaturesInEnvironmentV1ByCrnParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the change telemetry features in environment v1 by crn params
func (o *ChangeTelemetryFeaturesInEnvironmentV1ByCrnParams) WithBody(body *model.FeaturesRequest) *ChangeTelemetryFeaturesInEnvironmentV1ByCrnParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the change telemetry features in environment v1 by crn params
func (o *ChangeTelemetryFeaturesInEnvironmentV1ByCrnParams) SetBody(body *model.FeaturesRequest) {
	o.Body = body
}

// WithCrn adds the crn to the change telemetry features in environment v1 by crn params
func (o *ChangeTelemetryFeaturesInEnvironmentV1ByCrnParams) WithCrn(crn string) *ChangeTelemetryFeaturesInEnvironmentV1ByCrnParams {
	o.SetCrn(crn)
	return o
}

// SetCrn adds the crn to the change telemetry features in environment v1 by crn params
func (o *ChangeTelemetryFeaturesInEnvironmentV1ByCrnParams) SetCrn(crn string) {
	o.Crn = crn
}

// WriteToRequest writes these params to a swagger request
func (o *ChangeTelemetryFeaturesInEnvironmentV1ByCrnParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param crn
	if err := r.SetPathParam("crn", o.Crn); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
