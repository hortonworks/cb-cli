// Code generated by go-swagger; DO NOT EDIT.

package v3_organization_id_blueprints

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/hortonworks/cb-cli/models_cloudbreak"
)

// NewCreateBlueprintInOrganizationParams creates a new CreateBlueprintInOrganizationParams object
// with the default values initialized.
func NewCreateBlueprintInOrganizationParams() *CreateBlueprintInOrganizationParams {
	var ()
	return &CreateBlueprintInOrganizationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateBlueprintInOrganizationParamsWithTimeout creates a new CreateBlueprintInOrganizationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateBlueprintInOrganizationParamsWithTimeout(timeout time.Duration) *CreateBlueprintInOrganizationParams {
	var ()
	return &CreateBlueprintInOrganizationParams{

		timeout: timeout,
	}
}

// NewCreateBlueprintInOrganizationParamsWithContext creates a new CreateBlueprintInOrganizationParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateBlueprintInOrganizationParamsWithContext(ctx context.Context) *CreateBlueprintInOrganizationParams {
	var ()
	return &CreateBlueprintInOrganizationParams{

		Context: ctx,
	}
}

// NewCreateBlueprintInOrganizationParamsWithHTTPClient creates a new CreateBlueprintInOrganizationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateBlueprintInOrganizationParamsWithHTTPClient(client *http.Client) *CreateBlueprintInOrganizationParams {
	var ()
	return &CreateBlueprintInOrganizationParams{
		HTTPClient: client,
	}
}

/*CreateBlueprintInOrganizationParams contains all the parameters to send to the API endpoint
for the create blueprint in organization operation typically these are written to a http.Request
*/
type CreateBlueprintInOrganizationParams struct {

	/*Body*/
	Body *models_cloudbreak.BlueprintRequest
	/*OrganizationID*/
	OrganizationID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create blueprint in organization params
func (o *CreateBlueprintInOrganizationParams) WithTimeout(timeout time.Duration) *CreateBlueprintInOrganizationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create blueprint in organization params
func (o *CreateBlueprintInOrganizationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create blueprint in organization params
func (o *CreateBlueprintInOrganizationParams) WithContext(ctx context.Context) *CreateBlueprintInOrganizationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create blueprint in organization params
func (o *CreateBlueprintInOrganizationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create blueprint in organization params
func (o *CreateBlueprintInOrganizationParams) WithHTTPClient(client *http.Client) *CreateBlueprintInOrganizationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create blueprint in organization params
func (o *CreateBlueprintInOrganizationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create blueprint in organization params
func (o *CreateBlueprintInOrganizationParams) WithBody(body *models_cloudbreak.BlueprintRequest) *CreateBlueprintInOrganizationParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create blueprint in organization params
func (o *CreateBlueprintInOrganizationParams) SetBody(body *models_cloudbreak.BlueprintRequest) {
	o.Body = body
}

// WithOrganizationID adds the organizationID to the create blueprint in organization params
func (o *CreateBlueprintInOrganizationParams) WithOrganizationID(organizationID int64) *CreateBlueprintInOrganizationParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the create blueprint in organization params
func (o *CreateBlueprintInOrganizationParams) SetOrganizationID(organizationID int64) {
	o.OrganizationID = organizationID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateBlueprintInOrganizationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body == nil {
		o.Body = new(models_cloudbreak.BlueprintRequest)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param organizationId
	if err := r.SetPathParam("organizationId", swag.FormatInt64(o.OrganizationID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
