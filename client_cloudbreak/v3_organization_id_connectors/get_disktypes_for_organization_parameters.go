// Code generated by go-swagger; DO NOT EDIT.

package v3_organization_id_connectors

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
)

// NewGetDisktypesForOrganizationParams creates a new GetDisktypesForOrganizationParams object
// with the default values initialized.
func NewGetDisktypesForOrganizationParams() *GetDisktypesForOrganizationParams {
	var ()
	return &GetDisktypesForOrganizationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDisktypesForOrganizationParamsWithTimeout creates a new GetDisktypesForOrganizationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDisktypesForOrganizationParamsWithTimeout(timeout time.Duration) *GetDisktypesForOrganizationParams {
	var ()
	return &GetDisktypesForOrganizationParams{

		timeout: timeout,
	}
}

// NewGetDisktypesForOrganizationParamsWithContext creates a new GetDisktypesForOrganizationParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDisktypesForOrganizationParamsWithContext(ctx context.Context) *GetDisktypesForOrganizationParams {
	var ()
	return &GetDisktypesForOrganizationParams{

		Context: ctx,
	}
}

// NewGetDisktypesForOrganizationParamsWithHTTPClient creates a new GetDisktypesForOrganizationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDisktypesForOrganizationParamsWithHTTPClient(client *http.Client) *GetDisktypesForOrganizationParams {
	var ()
	return &GetDisktypesForOrganizationParams{
		HTTPClient: client,
	}
}

/*GetDisktypesForOrganizationParams contains all the parameters to send to the API endpoint
for the get disktypes for organization operation typically these are written to a http.Request
*/
type GetDisktypesForOrganizationParams struct {

	/*OrganizationID*/
	OrganizationID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get disktypes for organization params
func (o *GetDisktypesForOrganizationParams) WithTimeout(timeout time.Duration) *GetDisktypesForOrganizationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get disktypes for organization params
func (o *GetDisktypesForOrganizationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get disktypes for organization params
func (o *GetDisktypesForOrganizationParams) WithContext(ctx context.Context) *GetDisktypesForOrganizationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get disktypes for organization params
func (o *GetDisktypesForOrganizationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get disktypes for organization params
func (o *GetDisktypesForOrganizationParams) WithHTTPClient(client *http.Client) *GetDisktypesForOrganizationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get disktypes for organization params
func (o *GetDisktypesForOrganizationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrganizationID adds the organizationID to the get disktypes for organization params
func (o *GetDisktypesForOrganizationParams) WithOrganizationID(organizationID int64) *GetDisktypesForOrganizationParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the get disktypes for organization params
func (o *GetDisktypesForOrganizationParams) SetOrganizationID(organizationID int64) {
	o.OrganizationID = organizationID
}

// WriteToRequest writes these params to a swagger request
func (o *GetDisktypesForOrganizationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param organizationId
	if err := r.SetPathParam("organizationId", swag.FormatInt64(o.OrganizationID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
