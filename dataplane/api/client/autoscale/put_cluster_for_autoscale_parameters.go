// Code generated by go-swagger; DO NOT EDIT.

package autoscale

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

// NewPutClusterForAutoscaleParams creates a new PutClusterForAutoscaleParams object
// with the default values initialized.
func NewPutClusterForAutoscaleParams() *PutClusterForAutoscaleParams {
	var ()
	return &PutClusterForAutoscaleParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutClusterForAutoscaleParamsWithTimeout creates a new PutClusterForAutoscaleParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutClusterForAutoscaleParamsWithTimeout(timeout time.Duration) *PutClusterForAutoscaleParams {
	var ()
	return &PutClusterForAutoscaleParams{

		timeout: timeout,
	}
}

// NewPutClusterForAutoscaleParamsWithContext creates a new PutClusterForAutoscaleParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutClusterForAutoscaleParamsWithContext(ctx context.Context) *PutClusterForAutoscaleParams {
	var ()
	return &PutClusterForAutoscaleParams{

		Context: ctx,
	}
}

// NewPutClusterForAutoscaleParamsWithHTTPClient creates a new PutClusterForAutoscaleParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutClusterForAutoscaleParamsWithHTTPClient(client *http.Client) *PutClusterForAutoscaleParams {
	var ()
	return &PutClusterForAutoscaleParams{
		HTTPClient: client,
	}
}

/*PutClusterForAutoscaleParams contains all the parameters to send to the API endpoint
for the put cluster for autoscale operation typically these are written to a http.Request
*/
type PutClusterForAutoscaleParams struct {

	/*Body*/
	Body *model.UpdateClusterV4Request
	/*Crn*/
	Crn string
	/*UserID*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put cluster for autoscale params
func (o *PutClusterForAutoscaleParams) WithTimeout(timeout time.Duration) *PutClusterForAutoscaleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put cluster for autoscale params
func (o *PutClusterForAutoscaleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put cluster for autoscale params
func (o *PutClusterForAutoscaleParams) WithContext(ctx context.Context) *PutClusterForAutoscaleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put cluster for autoscale params
func (o *PutClusterForAutoscaleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put cluster for autoscale params
func (o *PutClusterForAutoscaleParams) WithHTTPClient(client *http.Client) *PutClusterForAutoscaleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put cluster for autoscale params
func (o *PutClusterForAutoscaleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put cluster for autoscale params
func (o *PutClusterForAutoscaleParams) WithBody(body *model.UpdateClusterV4Request) *PutClusterForAutoscaleParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put cluster for autoscale params
func (o *PutClusterForAutoscaleParams) SetBody(body *model.UpdateClusterV4Request) {
	o.Body = body
}

// WithCrn adds the crn to the put cluster for autoscale params
func (o *PutClusterForAutoscaleParams) WithCrn(crn string) *PutClusterForAutoscaleParams {
	o.SetCrn(crn)
	return o
}

// SetCrn adds the crn to the put cluster for autoscale params
func (o *PutClusterForAutoscaleParams) SetCrn(crn string) {
	o.Crn = crn
}

// WithUserID adds the userID to the put cluster for autoscale params
func (o *PutClusterForAutoscaleParams) WithUserID(userID string) *PutClusterForAutoscaleParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the put cluster for autoscale params
func (o *PutClusterForAutoscaleParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *PutClusterForAutoscaleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param userId
	if err := r.SetPathParam("userId", o.UserID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
