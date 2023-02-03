// Code generated by go-swagger; DO NOT EDIT.

package v1dns

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

// NewDeleteDNSZoneBySubnetIDV1Params creates a new DeleteDNSZoneBySubnetIDV1Params object
// with the default values initialized.
func NewDeleteDNSZoneBySubnetIDV1Params() *DeleteDNSZoneBySubnetIDV1Params {
	var ()
	return &DeleteDNSZoneBySubnetIDV1Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteDNSZoneBySubnetIDV1ParamsWithTimeout creates a new DeleteDNSZoneBySubnetIDV1Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteDNSZoneBySubnetIDV1ParamsWithTimeout(timeout time.Duration) *DeleteDNSZoneBySubnetIDV1Params {
	var ()
	return &DeleteDNSZoneBySubnetIDV1Params{

		timeout: timeout,
	}
}

// NewDeleteDNSZoneBySubnetIDV1ParamsWithContext creates a new DeleteDNSZoneBySubnetIDV1Params object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteDNSZoneBySubnetIDV1ParamsWithContext(ctx context.Context) *DeleteDNSZoneBySubnetIDV1Params {
	var ()
	return &DeleteDNSZoneBySubnetIDV1Params{

		Context: ctx,
	}
}

// NewDeleteDNSZoneBySubnetIDV1ParamsWithHTTPClient creates a new DeleteDNSZoneBySubnetIDV1Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteDNSZoneBySubnetIDV1ParamsWithHTTPClient(client *http.Client) *DeleteDNSZoneBySubnetIDV1Params {
	var ()
	return &DeleteDNSZoneBySubnetIDV1Params{
		HTTPClient: client,
	}
}

/*
DeleteDNSZoneBySubnetIDV1Params contains all the parameters to send to the API endpoint
for the delete Dns zone by subnet Id v1 operation typically these are written to a http.Request
*/
type DeleteDNSZoneBySubnetIDV1Params struct {

	/*Environment*/
	Environment *string
	/*NetworkID*/
	NetworkID *string
	/*SubnetID*/
	SubnetID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete Dns zone by subnet Id v1 params
func (o *DeleteDNSZoneBySubnetIDV1Params) WithTimeout(timeout time.Duration) *DeleteDNSZoneBySubnetIDV1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete Dns zone by subnet Id v1 params
func (o *DeleteDNSZoneBySubnetIDV1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete Dns zone by subnet Id v1 params
func (o *DeleteDNSZoneBySubnetIDV1Params) WithContext(ctx context.Context) *DeleteDNSZoneBySubnetIDV1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete Dns zone by subnet Id v1 params
func (o *DeleteDNSZoneBySubnetIDV1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete Dns zone by subnet Id v1 params
func (o *DeleteDNSZoneBySubnetIDV1Params) WithHTTPClient(client *http.Client) *DeleteDNSZoneBySubnetIDV1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete Dns zone by subnet Id v1 params
func (o *DeleteDNSZoneBySubnetIDV1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEnvironment adds the environment to the delete Dns zone by subnet Id v1 params
func (o *DeleteDNSZoneBySubnetIDV1Params) WithEnvironment(environment *string) *DeleteDNSZoneBySubnetIDV1Params {
	o.SetEnvironment(environment)
	return o
}

// SetEnvironment adds the environment to the delete Dns zone by subnet Id v1 params
func (o *DeleteDNSZoneBySubnetIDV1Params) SetEnvironment(environment *string) {
	o.Environment = environment
}

// WithNetworkID adds the networkID to the delete Dns zone by subnet Id v1 params
func (o *DeleteDNSZoneBySubnetIDV1Params) WithNetworkID(networkID *string) *DeleteDNSZoneBySubnetIDV1Params {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the delete Dns zone by subnet Id v1 params
func (o *DeleteDNSZoneBySubnetIDV1Params) SetNetworkID(networkID *string) {
	o.NetworkID = networkID
}

// WithSubnetID adds the subnetID to the delete Dns zone by subnet Id v1 params
func (o *DeleteDNSZoneBySubnetIDV1Params) WithSubnetID(subnetID *string) *DeleteDNSZoneBySubnetIDV1Params {
	o.SetSubnetID(subnetID)
	return o
}

// SetSubnetID adds the subnetId to the delete Dns zone by subnet Id v1 params
func (o *DeleteDNSZoneBySubnetIDV1Params) SetSubnetID(subnetID *string) {
	o.SubnetID = subnetID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteDNSZoneBySubnetIDV1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Environment != nil {

		// query param environment
		var qrEnvironment string
		if o.Environment != nil {
			qrEnvironment = *o.Environment
		}
		qEnvironment := qrEnvironment
		if qEnvironment != "" {
			if err := r.SetQueryParam("environment", qEnvironment); err != nil {
				return err
			}
		}

	}

	if o.NetworkID != nil {

		// query param networkId
		var qrNetworkID string
		if o.NetworkID != nil {
			qrNetworkID = *o.NetworkID
		}
		qNetworkID := qrNetworkID
		if qNetworkID != "" {
			if err := r.SetQueryParam("networkId", qNetworkID); err != nil {
				return err
			}
		}

	}

	if o.SubnetID != nil {

		// query param subnetId
		var qrSubnetID string
		if o.SubnetID != nil {
			qrSubnetID = *o.SubnetID
		}
		qSubnetID := qrSubnetID
		if qSubnetID != "" {
			if err := r.SetQueryParam("subnetId", qSubnetID); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
