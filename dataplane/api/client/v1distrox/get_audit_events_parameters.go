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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetAuditEventsParams creates a new GetAuditEventsParams object
// with the default values initialized.
func NewGetAuditEventsParams() *GetAuditEventsParams {
	var (
		pageDefault = int32(0)
		sizeDefault = int32(100)
	)
	return &GetAuditEventsParams{
		Page: &pageDefault,
		Size: &sizeDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAuditEventsParamsWithTimeout creates a new GetAuditEventsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAuditEventsParamsWithTimeout(timeout time.Duration) *GetAuditEventsParams {
	var (
		pageDefault = int32(0)
		sizeDefault = int32(100)
	)
	return &GetAuditEventsParams{
		Page: &pageDefault,
		Size: &sizeDefault,

		timeout: timeout,
	}
}

// NewGetAuditEventsParamsWithContext creates a new GetAuditEventsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAuditEventsParamsWithContext(ctx context.Context) *GetAuditEventsParams {
	var (
		pageDefault = int32(0)
		sizeDefault = int32(100)
	)
	return &GetAuditEventsParams{
		Page: &pageDefault,
		Size: &sizeDefault,

		Context: ctx,
	}
}

// NewGetAuditEventsParamsWithHTTPClient creates a new GetAuditEventsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAuditEventsParamsWithHTTPClient(client *http.Client) *GetAuditEventsParams {
	var (
		pageDefault = int32(0)
		sizeDefault = int32(100)
	)
	return &GetAuditEventsParams{
		Page:       &pageDefault,
		Size:       &sizeDefault,
		HTTPClient: client,
	}
}

/*GetAuditEventsParams contains all the parameters to send to the API endpoint
for the get audit events operation typically these are written to a http.Request
*/
type GetAuditEventsParams struct {

	/*Page*/
	Page *int32
	/*ResourceCrn*/
	ResourceCrn string
	/*Size*/
	Size *int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get audit events params
func (o *GetAuditEventsParams) WithTimeout(timeout time.Duration) *GetAuditEventsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get audit events params
func (o *GetAuditEventsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get audit events params
func (o *GetAuditEventsParams) WithContext(ctx context.Context) *GetAuditEventsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get audit events params
func (o *GetAuditEventsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get audit events params
func (o *GetAuditEventsParams) WithHTTPClient(client *http.Client) *GetAuditEventsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get audit events params
func (o *GetAuditEventsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPage adds the page to the get audit events params
func (o *GetAuditEventsParams) WithPage(page *int32) *GetAuditEventsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get audit events params
func (o *GetAuditEventsParams) SetPage(page *int32) {
	o.Page = page
}

// WithResourceCrn adds the resourceCrn to the get audit events params
func (o *GetAuditEventsParams) WithResourceCrn(resourceCrn string) *GetAuditEventsParams {
	o.SetResourceCrn(resourceCrn)
	return o
}

// SetResourceCrn adds the resourceCrn to the get audit events params
func (o *GetAuditEventsParams) SetResourceCrn(resourceCrn string) {
	o.ResourceCrn = resourceCrn
}

// WithSize adds the size to the get audit events params
func (o *GetAuditEventsParams) WithSize(size *int32) *GetAuditEventsParams {
	o.SetSize(size)
	return o
}

// SetSize adds the size to the get audit events params
func (o *GetAuditEventsParams) SetSize(size *int32) {
	o.Size = size
}

// WriteToRequest writes these params to a swagger request
func (o *GetAuditEventsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Page != nil {

		// query param page
		var qrPage int32
		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt32(qrPage)
		if qPage != "" {
			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}

	}

	// query param resourceCrn
	qrResourceCrn := o.ResourceCrn
	qResourceCrn := qrResourceCrn
	if qResourceCrn != "" {
		if err := r.SetQueryParam("resourceCrn", qResourceCrn); err != nil {
			return err
		}
	}

	if o.Size != nil {

		// query param size
		var qrSize int32
		if o.Size != nil {
			qrSize = *o.Size
		}
		qSize := swag.FormatInt32(qrSize)
		if qSize != "" {
			if err := r.SetQueryParam("size", qSize); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
