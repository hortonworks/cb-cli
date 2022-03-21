// Code generated by go-swagger; DO NOT EDIT.

package v1credentials

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetPrerequisitesForCloudPlatformParams creates a new GetPrerequisitesForCloudPlatformParams object
// with the default values initialized.
func NewGetPrerequisitesForCloudPlatformParams() *GetPrerequisitesForCloudPlatformParams {
	var ()
	return &GetPrerequisitesForCloudPlatformParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPrerequisitesForCloudPlatformParamsWithTimeout creates a new GetPrerequisitesForCloudPlatformParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPrerequisitesForCloudPlatformParamsWithTimeout(timeout time.Duration) *GetPrerequisitesForCloudPlatformParams {
	var ()
	return &GetPrerequisitesForCloudPlatformParams{

		timeout: timeout,
	}
}

// NewGetPrerequisitesForCloudPlatformParamsWithContext creates a new GetPrerequisitesForCloudPlatformParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPrerequisitesForCloudPlatformParamsWithContext(ctx context.Context) *GetPrerequisitesForCloudPlatformParams {
	var ()
	return &GetPrerequisitesForCloudPlatformParams{

		Context: ctx,
	}
}

// NewGetPrerequisitesForCloudPlatformParamsWithHTTPClient creates a new GetPrerequisitesForCloudPlatformParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPrerequisitesForCloudPlatformParamsWithHTTPClient(client *http.Client) *GetPrerequisitesForCloudPlatformParams {
	var ()
	return &GetPrerequisitesForCloudPlatformParams{
		HTTPClient: client,
	}
}

/*GetPrerequisitesForCloudPlatformParams contains all the parameters to send to the API endpoint
for the get prerequisites for cloud platform operation typically these are written to a http.Request
*/
type GetPrerequisitesForCloudPlatformParams struct {

	/*CloudPlatform*/
	CloudPlatform string
	/*DeploymentAddress*/
	DeploymentAddress *string
	/*GovCloud*/
	GovCloud *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get prerequisites for cloud platform params
func (o *GetPrerequisitesForCloudPlatformParams) WithTimeout(timeout time.Duration) *GetPrerequisitesForCloudPlatformParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get prerequisites for cloud platform params
func (o *GetPrerequisitesForCloudPlatformParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get prerequisites for cloud platform params
func (o *GetPrerequisitesForCloudPlatformParams) WithContext(ctx context.Context) *GetPrerequisitesForCloudPlatformParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get prerequisites for cloud platform params
func (o *GetPrerequisitesForCloudPlatformParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get prerequisites for cloud platform params
func (o *GetPrerequisitesForCloudPlatformParams) WithHTTPClient(client *http.Client) *GetPrerequisitesForCloudPlatformParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get prerequisites for cloud platform params
func (o *GetPrerequisitesForCloudPlatformParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCloudPlatform adds the cloudPlatform to the get prerequisites for cloud platform params
func (o *GetPrerequisitesForCloudPlatformParams) WithCloudPlatform(cloudPlatform string) *GetPrerequisitesForCloudPlatformParams {
	o.SetCloudPlatform(cloudPlatform)
	return o
}

// SetCloudPlatform adds the cloudPlatform to the get prerequisites for cloud platform params
func (o *GetPrerequisitesForCloudPlatformParams) SetCloudPlatform(cloudPlatform string) {
	o.CloudPlatform = cloudPlatform
}

// WithDeploymentAddress adds the deploymentAddress to the get prerequisites for cloud platform params
func (o *GetPrerequisitesForCloudPlatformParams) WithDeploymentAddress(deploymentAddress *string) *GetPrerequisitesForCloudPlatformParams {
	o.SetDeploymentAddress(deploymentAddress)
	return o
}

// SetDeploymentAddress adds the deploymentAddress to the get prerequisites for cloud platform params
func (o *GetPrerequisitesForCloudPlatformParams) SetDeploymentAddress(deploymentAddress *string) {
	o.DeploymentAddress = deploymentAddress
}

// WithGovCloud adds the govCloud to the get prerequisites for cloud platform params
func (o *GetPrerequisitesForCloudPlatformParams) WithGovCloud(govCloud *bool) *GetPrerequisitesForCloudPlatformParams {
	o.SetGovCloud(govCloud)
	return o
}

// SetGovCloud adds the govCloud to the get prerequisites for cloud platform params
func (o *GetPrerequisitesForCloudPlatformParams) SetGovCloud(govCloud *bool) {
	o.GovCloud = govCloud
}

// WriteToRequest writes these params to a swagger request
func (o *GetPrerequisitesForCloudPlatformParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cloudPlatform
	if err := r.SetPathParam("cloudPlatform", o.CloudPlatform); err != nil {
		return err
	}

	if o.DeploymentAddress != nil {

		// query param deploymentAddress
		var qrDeploymentAddress string
		if o.DeploymentAddress != nil {
			qrDeploymentAddress = *o.DeploymentAddress
		}
		qDeploymentAddress := qrDeploymentAddress
		if qDeploymentAddress != "" {
			if err := r.SetQueryParam("deploymentAddress", qDeploymentAddress); err != nil {
				return err
			}
		}

	}

	if o.GovCloud != nil {

		// query param govCloud
		var qrGovCloud bool
		if o.GovCloud != nil {
			qrGovCloud = *o.GovCloud
		}
		qGovCloud := swag.FormatBool(qrGovCloud)
		if qGovCloud != "" {
			if err := r.SetQueryParam("govCloud", qGovCloud); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
