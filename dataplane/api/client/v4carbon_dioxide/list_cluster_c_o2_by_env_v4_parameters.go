// Code generated by go-swagger; DO NOT EDIT.

package v4carbon_dioxide

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

// NewListClusterCO2ByEnvV4Params creates a new ListClusterCO2ByEnvV4Params object
// with the default values initialized.
func NewListClusterCO2ByEnvV4Params() *ListClusterCO2ByEnvV4Params {
	var ()
	return &ListClusterCO2ByEnvV4Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewListClusterCO2ByEnvV4ParamsWithTimeout creates a new ListClusterCO2ByEnvV4Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewListClusterCO2ByEnvV4ParamsWithTimeout(timeout time.Duration) *ListClusterCO2ByEnvV4Params {
	var ()
	return &ListClusterCO2ByEnvV4Params{

		timeout: timeout,
	}
}

// NewListClusterCO2ByEnvV4ParamsWithContext creates a new ListClusterCO2ByEnvV4Params object
// with the default values initialized, and the ability to set a context for a request
func NewListClusterCO2ByEnvV4ParamsWithContext(ctx context.Context) *ListClusterCO2ByEnvV4Params {
	var ()
	return &ListClusterCO2ByEnvV4Params{

		Context: ctx,
	}
}

// NewListClusterCO2ByEnvV4ParamsWithHTTPClient creates a new ListClusterCO2ByEnvV4Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListClusterCO2ByEnvV4ParamsWithHTTPClient(client *http.Client) *ListClusterCO2ByEnvV4Params {
	var ()
	return &ListClusterCO2ByEnvV4Params{
		HTTPClient: client,
	}
}

/*
ListClusterCO2ByEnvV4Params contains all the parameters to send to the API endpoint
for the list cluster c o2 by env v4 operation typically these are written to a http.Request
*/
type ListClusterCO2ByEnvV4Params struct {

	/*Body*/
	Body *model.ClusterCO2V4Request
	/*InitiatorUserCrn*/
	InitiatorUserCrn *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list cluster c o2 by env v4 params
func (o *ListClusterCO2ByEnvV4Params) WithTimeout(timeout time.Duration) *ListClusterCO2ByEnvV4Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list cluster c o2 by env v4 params
func (o *ListClusterCO2ByEnvV4Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list cluster c o2 by env v4 params
func (o *ListClusterCO2ByEnvV4Params) WithContext(ctx context.Context) *ListClusterCO2ByEnvV4Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list cluster c o2 by env v4 params
func (o *ListClusterCO2ByEnvV4Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list cluster c o2 by env v4 params
func (o *ListClusterCO2ByEnvV4Params) WithHTTPClient(client *http.Client) *ListClusterCO2ByEnvV4Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list cluster c o2 by env v4 params
func (o *ListClusterCO2ByEnvV4Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the list cluster c o2 by env v4 params
func (o *ListClusterCO2ByEnvV4Params) WithBody(body *model.ClusterCO2V4Request) *ListClusterCO2ByEnvV4Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the list cluster c o2 by env v4 params
func (o *ListClusterCO2ByEnvV4Params) SetBody(body *model.ClusterCO2V4Request) {
	o.Body = body
}

// WithInitiatorUserCrn adds the initiatorUserCrn to the list cluster c o2 by env v4 params
func (o *ListClusterCO2ByEnvV4Params) WithInitiatorUserCrn(initiatorUserCrn *string) *ListClusterCO2ByEnvV4Params {
	o.SetInitiatorUserCrn(initiatorUserCrn)
	return o
}

// SetInitiatorUserCrn adds the initiatorUserCrn to the list cluster c o2 by env v4 params
func (o *ListClusterCO2ByEnvV4Params) SetInitiatorUserCrn(initiatorUserCrn *string) {
	o.InitiatorUserCrn = initiatorUserCrn
}

// WriteToRequest writes these params to a swagger request
func (o *ListClusterCO2ByEnvV4Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if o.InitiatorUserCrn != nil {

		// query param initiatorUserCrn
		var qrInitiatorUserCrn string
		if o.InitiatorUserCrn != nil {
			qrInitiatorUserCrn = *o.InitiatorUserCrn
		}
		qInitiatorUserCrn := qrInitiatorUserCrn
		if qInitiatorUserCrn != "" {
			if err := r.SetQueryParam("initiatorUserCrn", qInitiatorUserCrn); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
