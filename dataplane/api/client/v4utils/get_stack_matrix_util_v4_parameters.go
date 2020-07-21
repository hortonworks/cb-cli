// Code generated by go-swagger; DO NOT EDIT.

package v4utils

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
)

// NewGetStackMatrixUtilV4Params creates a new GetStackMatrixUtilV4Params object
// with the default values initialized.
func NewGetStackMatrixUtilV4Params() *GetStackMatrixUtilV4Params {
	var ()
	return &GetStackMatrixUtilV4Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetStackMatrixUtilV4ParamsWithTimeout creates a new GetStackMatrixUtilV4Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetStackMatrixUtilV4ParamsWithTimeout(timeout time.Duration) *GetStackMatrixUtilV4Params {
	var ()
	return &GetStackMatrixUtilV4Params{

		timeout: timeout,
	}
}

// NewGetStackMatrixUtilV4ParamsWithContext creates a new GetStackMatrixUtilV4Params object
// with the default values initialized, and the ability to set a context for a request
func NewGetStackMatrixUtilV4ParamsWithContext(ctx context.Context) *GetStackMatrixUtilV4Params {
	var ()
	return &GetStackMatrixUtilV4Params{

		Context: ctx,
	}
}

// NewGetStackMatrixUtilV4ParamsWithHTTPClient creates a new GetStackMatrixUtilV4Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetStackMatrixUtilV4ParamsWithHTTPClient(client *http.Client) *GetStackMatrixUtilV4Params {
	var ()
	return &GetStackMatrixUtilV4Params{
		HTTPClient: client,
	}
}

/*GetStackMatrixUtilV4Params contains all the parameters to send to the API endpoint
for the get stack matrix util v4 operation typically these are written to a http.Request
*/
type GetStackMatrixUtilV4Params struct {

	/*ImageCatalogName*/
	ImageCatalogName *string
	/*Platform*/
	Platform *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get stack matrix util v4 params
func (o *GetStackMatrixUtilV4Params) WithTimeout(timeout time.Duration) *GetStackMatrixUtilV4Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get stack matrix util v4 params
func (o *GetStackMatrixUtilV4Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get stack matrix util v4 params
func (o *GetStackMatrixUtilV4Params) WithContext(ctx context.Context) *GetStackMatrixUtilV4Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get stack matrix util v4 params
func (o *GetStackMatrixUtilV4Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get stack matrix util v4 params
func (o *GetStackMatrixUtilV4Params) WithHTTPClient(client *http.Client) *GetStackMatrixUtilV4Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get stack matrix util v4 params
func (o *GetStackMatrixUtilV4Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithImageCatalogName adds the imageCatalogName to the get stack matrix util v4 params
func (o *GetStackMatrixUtilV4Params) WithImageCatalogName(imageCatalogName *string) *GetStackMatrixUtilV4Params {
	o.SetImageCatalogName(imageCatalogName)
	return o
}

// SetImageCatalogName adds the imageCatalogName to the get stack matrix util v4 params
func (o *GetStackMatrixUtilV4Params) SetImageCatalogName(imageCatalogName *string) {
	o.ImageCatalogName = imageCatalogName
}

// WithPlatform adds the platform to the get stack matrix util v4 params
func (o *GetStackMatrixUtilV4Params) WithPlatform(platform *string) *GetStackMatrixUtilV4Params {
	o.SetPlatform(platform)
	return o
}

// SetPlatform adds the platform to the get stack matrix util v4 params
func (o *GetStackMatrixUtilV4Params) SetPlatform(platform *string) {
	o.Platform = platform
}

// WriteToRequest writes these params to a swagger request
func (o *GetStackMatrixUtilV4Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ImageCatalogName != nil {

		// query param imageCatalogName
		var qrImageCatalogName string
		if o.ImageCatalogName != nil {
			qrImageCatalogName = *o.ImageCatalogName
		}
		qImageCatalogName := qrImageCatalogName
		if qImageCatalogName != "" {
			if err := r.SetQueryParam("imageCatalogName", qImageCatalogName); err != nil {
				return err
			}
		}

	}

	if o.Platform != nil {

		// query param platform
		var qrPlatform string
		if o.Platform != nil {
			qrPlatform = *o.Platform
		}
		qPlatform := qrPlatform
		if qPlatform != "" {
			if err := r.SetQueryParam("platform", qPlatform); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
