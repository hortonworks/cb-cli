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

// NewPrepareDistroxClusterUpgradeByCrnParams creates a new PrepareDistroxClusterUpgradeByCrnParams object
// with the default values initialized.
func NewPrepareDistroxClusterUpgradeByCrnParams() *PrepareDistroxClusterUpgradeByCrnParams {
	var ()
	return &PrepareDistroxClusterUpgradeByCrnParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPrepareDistroxClusterUpgradeByCrnParamsWithTimeout creates a new PrepareDistroxClusterUpgradeByCrnParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPrepareDistroxClusterUpgradeByCrnParamsWithTimeout(timeout time.Duration) *PrepareDistroxClusterUpgradeByCrnParams {
	var ()
	return &PrepareDistroxClusterUpgradeByCrnParams{

		timeout: timeout,
	}
}

// NewPrepareDistroxClusterUpgradeByCrnParamsWithContext creates a new PrepareDistroxClusterUpgradeByCrnParams object
// with the default values initialized, and the ability to set a context for a request
func NewPrepareDistroxClusterUpgradeByCrnParamsWithContext(ctx context.Context) *PrepareDistroxClusterUpgradeByCrnParams {
	var ()
	return &PrepareDistroxClusterUpgradeByCrnParams{

		Context: ctx,
	}
}

// NewPrepareDistroxClusterUpgradeByCrnParamsWithHTTPClient creates a new PrepareDistroxClusterUpgradeByCrnParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPrepareDistroxClusterUpgradeByCrnParamsWithHTTPClient(client *http.Client) *PrepareDistroxClusterUpgradeByCrnParams {
	var ()
	return &PrepareDistroxClusterUpgradeByCrnParams{
		HTTPClient: client,
	}
}

/*PrepareDistroxClusterUpgradeByCrnParams contains all the parameters to send to the API endpoint
for the prepare distrox cluster upgrade by crn operation typically these are written to a http.Request
*/
type PrepareDistroxClusterUpgradeByCrnParams struct {

	/*Body*/
	Body *model.DistroXUpgradeV1Request
	/*Crn*/
	Crn string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the prepare distrox cluster upgrade by crn params
func (o *PrepareDistroxClusterUpgradeByCrnParams) WithTimeout(timeout time.Duration) *PrepareDistroxClusterUpgradeByCrnParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the prepare distrox cluster upgrade by crn params
func (o *PrepareDistroxClusterUpgradeByCrnParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the prepare distrox cluster upgrade by crn params
func (o *PrepareDistroxClusterUpgradeByCrnParams) WithContext(ctx context.Context) *PrepareDistroxClusterUpgradeByCrnParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the prepare distrox cluster upgrade by crn params
func (o *PrepareDistroxClusterUpgradeByCrnParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the prepare distrox cluster upgrade by crn params
func (o *PrepareDistroxClusterUpgradeByCrnParams) WithHTTPClient(client *http.Client) *PrepareDistroxClusterUpgradeByCrnParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the prepare distrox cluster upgrade by crn params
func (o *PrepareDistroxClusterUpgradeByCrnParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the prepare distrox cluster upgrade by crn params
func (o *PrepareDistroxClusterUpgradeByCrnParams) WithBody(body *model.DistroXUpgradeV1Request) *PrepareDistroxClusterUpgradeByCrnParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the prepare distrox cluster upgrade by crn params
func (o *PrepareDistroxClusterUpgradeByCrnParams) SetBody(body *model.DistroXUpgradeV1Request) {
	o.Body = body
}

// WithCrn adds the crn to the prepare distrox cluster upgrade by crn params
func (o *PrepareDistroxClusterUpgradeByCrnParams) WithCrn(crn string) *PrepareDistroxClusterUpgradeByCrnParams {
	o.SetCrn(crn)
	return o
}

// SetCrn adds the crn to the prepare distrox cluster upgrade by crn params
func (o *PrepareDistroxClusterUpgradeByCrnParams) SetCrn(crn string) {
	o.Crn = crn
}

// WriteToRequest writes these params to a swagger request
func (o *PrepareDistroxClusterUpgradeByCrnParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
