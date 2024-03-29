// Code generated by go-swagger; DO NOT EDIT.

package v1ldaps

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

// NewGetLdapConfigForClusterV1Params creates a new GetLdapConfigForClusterV1Params object
// with the default values initialized.
func NewGetLdapConfigForClusterV1Params() *GetLdapConfigForClusterV1Params {
	var ()
	return &GetLdapConfigForClusterV1Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLdapConfigForClusterV1ParamsWithTimeout creates a new GetLdapConfigForClusterV1Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLdapConfigForClusterV1ParamsWithTimeout(timeout time.Duration) *GetLdapConfigForClusterV1Params {
	var ()
	return &GetLdapConfigForClusterV1Params{

		timeout: timeout,
	}
}

// NewGetLdapConfigForClusterV1ParamsWithContext creates a new GetLdapConfigForClusterV1Params object
// with the default values initialized, and the ability to set a context for a request
func NewGetLdapConfigForClusterV1ParamsWithContext(ctx context.Context) *GetLdapConfigForClusterV1Params {
	var ()
	return &GetLdapConfigForClusterV1Params{

		Context: ctx,
	}
}

// NewGetLdapConfigForClusterV1ParamsWithHTTPClient creates a new GetLdapConfigForClusterV1Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLdapConfigForClusterV1ParamsWithHTTPClient(client *http.Client) *GetLdapConfigForClusterV1Params {
	var ()
	return &GetLdapConfigForClusterV1Params{
		HTTPClient: client,
	}
}

/*
GetLdapConfigForClusterV1Params contains all the parameters to send to the API endpoint
for the get ldap config for cluster v1 operation typically these are written to a http.Request
*/
type GetLdapConfigForClusterV1Params struct {

	/*ClusterName*/
	ClusterName *string
	/*EnvironmentCrn*/
	EnvironmentCrn *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get ldap config for cluster v1 params
func (o *GetLdapConfigForClusterV1Params) WithTimeout(timeout time.Duration) *GetLdapConfigForClusterV1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get ldap config for cluster v1 params
func (o *GetLdapConfigForClusterV1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get ldap config for cluster v1 params
func (o *GetLdapConfigForClusterV1Params) WithContext(ctx context.Context) *GetLdapConfigForClusterV1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get ldap config for cluster v1 params
func (o *GetLdapConfigForClusterV1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get ldap config for cluster v1 params
func (o *GetLdapConfigForClusterV1Params) WithHTTPClient(client *http.Client) *GetLdapConfigForClusterV1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get ldap config for cluster v1 params
func (o *GetLdapConfigForClusterV1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterName adds the clusterName to the get ldap config for cluster v1 params
func (o *GetLdapConfigForClusterV1Params) WithClusterName(clusterName *string) *GetLdapConfigForClusterV1Params {
	o.SetClusterName(clusterName)
	return o
}

// SetClusterName adds the clusterName to the get ldap config for cluster v1 params
func (o *GetLdapConfigForClusterV1Params) SetClusterName(clusterName *string) {
	o.ClusterName = clusterName
}

// WithEnvironmentCrn adds the environmentCrn to the get ldap config for cluster v1 params
func (o *GetLdapConfigForClusterV1Params) WithEnvironmentCrn(environmentCrn *string) *GetLdapConfigForClusterV1Params {
	o.SetEnvironmentCrn(environmentCrn)
	return o
}

// SetEnvironmentCrn adds the environmentCrn to the get ldap config for cluster v1 params
func (o *GetLdapConfigForClusterV1Params) SetEnvironmentCrn(environmentCrn *string) {
	o.EnvironmentCrn = environmentCrn
}

// WriteToRequest writes these params to a swagger request
func (o *GetLdapConfigForClusterV1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ClusterName != nil {

		// query param clusterName
		var qrClusterName string
		if o.ClusterName != nil {
			qrClusterName = *o.ClusterName
		}
		qClusterName := qrClusterName
		if qClusterName != "" {
			if err := r.SetQueryParam("clusterName", qClusterName); err != nil {
				return err
			}
		}

	}

	if o.EnvironmentCrn != nil {

		// query param environmentCrn
		var qrEnvironmentCrn string
		if o.EnvironmentCrn != nil {
			qrEnvironmentCrn = *o.EnvironmentCrn
		}
		qEnvironmentCrn := qrEnvironmentCrn
		if qEnvironmentCrn != "" {
			if err := r.SetQueryParam("environmentCrn", qEnvironmentCrn); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
