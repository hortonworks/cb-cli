// Code generated by go-swagger; DO NOT EDIT.

package v1imagecatalogs

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

// NewGetPublicImageCatalogsByNameParams creates a new GetPublicImageCatalogsByNameParams object
// with the default values initialized.
func NewGetPublicImageCatalogsByNameParams() *GetPublicImageCatalogsByNameParams {
	var ()
	return &GetPublicImageCatalogsByNameParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPublicImageCatalogsByNameParamsWithTimeout creates a new GetPublicImageCatalogsByNameParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPublicImageCatalogsByNameParamsWithTimeout(timeout time.Duration) *GetPublicImageCatalogsByNameParams {
	var ()
	return &GetPublicImageCatalogsByNameParams{

		timeout: timeout,
	}
}

// NewGetPublicImageCatalogsByNameParamsWithContext creates a new GetPublicImageCatalogsByNameParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPublicImageCatalogsByNameParamsWithContext(ctx context.Context) *GetPublicImageCatalogsByNameParams {
	var ()
	return &GetPublicImageCatalogsByNameParams{

		Context: ctx,
	}
}

// NewGetPublicImageCatalogsByNameParamsWithHTTPClient creates a new GetPublicImageCatalogsByNameParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPublicImageCatalogsByNameParamsWithHTTPClient(client *http.Client) *GetPublicImageCatalogsByNameParams {
	var ()
	return &GetPublicImageCatalogsByNameParams{
		HTTPClient: client,
	}
}

/*GetPublicImageCatalogsByNameParams contains all the parameters to send to the API endpoint
for the get public image catalogs by name operation typically these are written to a http.Request
*/
type GetPublicImageCatalogsByNameParams struct {

	/*Name*/
	Name string
	/*WithImages*/
	WithImages *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get public image catalogs by name params
func (o *GetPublicImageCatalogsByNameParams) WithTimeout(timeout time.Duration) *GetPublicImageCatalogsByNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get public image catalogs by name params
func (o *GetPublicImageCatalogsByNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get public image catalogs by name params
func (o *GetPublicImageCatalogsByNameParams) WithContext(ctx context.Context) *GetPublicImageCatalogsByNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get public image catalogs by name params
func (o *GetPublicImageCatalogsByNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get public image catalogs by name params
func (o *GetPublicImageCatalogsByNameParams) WithHTTPClient(client *http.Client) *GetPublicImageCatalogsByNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get public image catalogs by name params
func (o *GetPublicImageCatalogsByNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the get public image catalogs by name params
func (o *GetPublicImageCatalogsByNameParams) WithName(name string) *GetPublicImageCatalogsByNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get public image catalogs by name params
func (o *GetPublicImageCatalogsByNameParams) SetName(name string) {
	o.Name = name
}

// WithWithImages adds the withImages to the get public image catalogs by name params
func (o *GetPublicImageCatalogsByNameParams) WithWithImages(withImages *bool) *GetPublicImageCatalogsByNameParams {
	o.SetWithImages(withImages)
	return o
}

// SetWithImages adds the withImages to the get public image catalogs by name params
func (o *GetPublicImageCatalogsByNameParams) SetWithImages(withImages *bool) {
	o.WithImages = withImages
}

// WriteToRequest writes these params to a swagger request
func (o *GetPublicImageCatalogsByNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if o.WithImages != nil {

		// query param withImages
		var qrWithImages bool
		if o.WithImages != nil {
			qrWithImages = *o.WithImages
		}
		qWithImages := swag.FormatBool(qrWithImages)
		if qWithImages != "" {
			if err := r.SetQueryParam("withImages", qWithImages); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
