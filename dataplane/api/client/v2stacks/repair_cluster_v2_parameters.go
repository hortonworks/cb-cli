// Code generated by go-swagger; DO NOT EDIT.

package v2stacks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api/model"
)

// NewRepairClusterV2Params creates a new RepairClusterV2Params object
// with the default values initialized.
func NewRepairClusterV2Params() *RepairClusterV2Params {
	var ()
	return &RepairClusterV2Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewRepairClusterV2ParamsWithTimeout creates a new RepairClusterV2Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewRepairClusterV2ParamsWithTimeout(timeout time.Duration) *RepairClusterV2Params {
	var ()
	return &RepairClusterV2Params{

		timeout: timeout,
	}
}

// NewRepairClusterV2ParamsWithContext creates a new RepairClusterV2Params object
// with the default values initialized, and the ability to set a context for a request
func NewRepairClusterV2ParamsWithContext(ctx context.Context) *RepairClusterV2Params {
	var ()
	return &RepairClusterV2Params{

		Context: ctx,
	}
}

// NewRepairClusterV2ParamsWithHTTPClient creates a new RepairClusterV2Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRepairClusterV2ParamsWithHTTPClient(client *http.Client) *RepairClusterV2Params {
	var ()
	return &RepairClusterV2Params{
		HTTPClient: client,
	}
}

/*RepairClusterV2Params contains all the parameters to send to the API endpoint
for the repair cluster v2 operation typically these are written to a http.Request
*/
type RepairClusterV2Params struct {

	/*Body*/
	Body *model.ClusterRepairRequest
	/*Name*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the repair cluster v2 params
func (o *RepairClusterV2Params) WithTimeout(timeout time.Duration) *RepairClusterV2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the repair cluster v2 params
func (o *RepairClusterV2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the repair cluster v2 params
func (o *RepairClusterV2Params) WithContext(ctx context.Context) *RepairClusterV2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the repair cluster v2 params
func (o *RepairClusterV2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the repair cluster v2 params
func (o *RepairClusterV2Params) WithHTTPClient(client *http.Client) *RepairClusterV2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the repair cluster v2 params
func (o *RepairClusterV2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the repair cluster v2 params
func (o *RepairClusterV2Params) WithBody(body *model.ClusterRepairRequest) *RepairClusterV2Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the repair cluster v2 params
func (o *RepairClusterV2Params) SetBody(body *model.ClusterRepairRequest) {
	o.Body = body
}

// WithName adds the name to the repair cluster v2 params
func (o *RepairClusterV2Params) WithName(name string) *RepairClusterV2Params {
	o.SetName(name)
	return o
}

// SetName adds the name to the repair cluster v2 params
func (o *RepairClusterV2Params) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *RepairClusterV2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
