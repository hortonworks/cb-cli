// Code generated by go-swagger; DO NOT EDIT.

package autoscale

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

// NewGetRecommendationParams creates a new GetRecommendationParams object
// with the default values initialized.
func NewGetRecommendationParams() *GetRecommendationParams {
	var ()
	return &GetRecommendationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRecommendationParamsWithTimeout creates a new GetRecommendationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRecommendationParamsWithTimeout(timeout time.Duration) *GetRecommendationParams {
	var ()
	return &GetRecommendationParams{

		timeout: timeout,
	}
}

// NewGetRecommendationParamsWithContext creates a new GetRecommendationParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRecommendationParamsWithContext(ctx context.Context) *GetRecommendationParams {
	var ()
	return &GetRecommendationParams{

		Context: ctx,
	}
}

// NewGetRecommendationParamsWithHTTPClient creates a new GetRecommendationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRecommendationParamsWithHTTPClient(client *http.Client) *GetRecommendationParams {
	var ()
	return &GetRecommendationParams{
		HTTPClient: client,
	}
}

/*GetRecommendationParams contains all the parameters to send to the API endpoint
for the get recommendation operation typically these are written to a http.Request
*/
type GetRecommendationParams struct {

	/*Crn*/
	Crn string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get recommendation params
func (o *GetRecommendationParams) WithTimeout(timeout time.Duration) *GetRecommendationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get recommendation params
func (o *GetRecommendationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get recommendation params
func (o *GetRecommendationParams) WithContext(ctx context.Context) *GetRecommendationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get recommendation params
func (o *GetRecommendationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get recommendation params
func (o *GetRecommendationParams) WithHTTPClient(client *http.Client) *GetRecommendationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get recommendation params
func (o *GetRecommendationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCrn adds the crn to the get recommendation params
func (o *GetRecommendationParams) WithCrn(crn string) *GetRecommendationParams {
	o.SetCrn(crn)
	return o
}

// SetCrn adds the crn to the get recommendation params
func (o *GetRecommendationParams) SetCrn(crn string) {
	o.Crn = crn
}

// WriteToRequest writes these params to a swagger request
func (o *GetRecommendationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param crn
	if err := r.SetPathParam("crn", o.Crn); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
