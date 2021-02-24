// Code generated by go-swagger; DO NOT EDIT.

package v1env

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api-environment/model"
)

// NewChangeTelemetryFeaturesInEnvironmentV1ByNameParams creates a new ChangeTelemetryFeaturesInEnvironmentV1ByNameParams object
// with the default values initialized.
func NewChangeTelemetryFeaturesInEnvironmentV1ByNameParams() *ChangeTelemetryFeaturesInEnvironmentV1ByNameParams {
	var ()
	return &ChangeTelemetryFeaturesInEnvironmentV1ByNameParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewChangeTelemetryFeaturesInEnvironmentV1ByNameParamsWithTimeout creates a new ChangeTelemetryFeaturesInEnvironmentV1ByNameParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewChangeTelemetryFeaturesInEnvironmentV1ByNameParamsWithTimeout(timeout time.Duration) *ChangeTelemetryFeaturesInEnvironmentV1ByNameParams {
	var ()
	return &ChangeTelemetryFeaturesInEnvironmentV1ByNameParams{

		timeout: timeout,
	}
}

// NewChangeTelemetryFeaturesInEnvironmentV1ByNameParamsWithContext creates a new ChangeTelemetryFeaturesInEnvironmentV1ByNameParams object
// with the default values initialized, and the ability to set a context for a request
func NewChangeTelemetryFeaturesInEnvironmentV1ByNameParamsWithContext(ctx context.Context) *ChangeTelemetryFeaturesInEnvironmentV1ByNameParams {
	var ()
	return &ChangeTelemetryFeaturesInEnvironmentV1ByNameParams{

		Context: ctx,
	}
}

// NewChangeTelemetryFeaturesInEnvironmentV1ByNameParamsWithHTTPClient creates a new ChangeTelemetryFeaturesInEnvironmentV1ByNameParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewChangeTelemetryFeaturesInEnvironmentV1ByNameParamsWithHTTPClient(client *http.Client) *ChangeTelemetryFeaturesInEnvironmentV1ByNameParams {
	var ()
	return &ChangeTelemetryFeaturesInEnvironmentV1ByNameParams{
		HTTPClient: client,
	}
}

/*ChangeTelemetryFeaturesInEnvironmentV1ByNameParams contains all the parameters to send to the API endpoint
for the change telemetry features in environment v1 by name operation typically these are written to a http.Request
*/
type ChangeTelemetryFeaturesInEnvironmentV1ByNameParams struct {

	/*Body*/
	Body *model.FeaturesRequest
	/*Name*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the change telemetry features in environment v1 by name params
func (o *ChangeTelemetryFeaturesInEnvironmentV1ByNameParams) WithTimeout(timeout time.Duration) *ChangeTelemetryFeaturesInEnvironmentV1ByNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the change telemetry features in environment v1 by name params
func (o *ChangeTelemetryFeaturesInEnvironmentV1ByNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the change telemetry features in environment v1 by name params
func (o *ChangeTelemetryFeaturesInEnvironmentV1ByNameParams) WithContext(ctx context.Context) *ChangeTelemetryFeaturesInEnvironmentV1ByNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the change telemetry features in environment v1 by name params
func (o *ChangeTelemetryFeaturesInEnvironmentV1ByNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the change telemetry features in environment v1 by name params
func (o *ChangeTelemetryFeaturesInEnvironmentV1ByNameParams) WithHTTPClient(client *http.Client) *ChangeTelemetryFeaturesInEnvironmentV1ByNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the change telemetry features in environment v1 by name params
func (o *ChangeTelemetryFeaturesInEnvironmentV1ByNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the change telemetry features in environment v1 by name params
func (o *ChangeTelemetryFeaturesInEnvironmentV1ByNameParams) WithBody(body *model.FeaturesRequest) *ChangeTelemetryFeaturesInEnvironmentV1ByNameParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the change telemetry features in environment v1 by name params
func (o *ChangeTelemetryFeaturesInEnvironmentV1ByNameParams) SetBody(body *model.FeaturesRequest) {
	o.Body = body
}

// WithName adds the name to the change telemetry features in environment v1 by name params
func (o *ChangeTelemetryFeaturesInEnvironmentV1ByNameParams) WithName(name string) *ChangeTelemetryFeaturesInEnvironmentV1ByNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the change telemetry features in environment v1 by name params
func (o *ChangeTelemetryFeaturesInEnvironmentV1ByNameParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *ChangeTelemetryFeaturesInEnvironmentV1ByNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
