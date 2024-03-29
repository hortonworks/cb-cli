// Code generated by go-swagger; DO NOT EDIT.

package sdx

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

// NewIsStoppableParams creates a new IsStoppableParams object
// with the default values initialized.
func NewIsStoppableParams() *IsStoppableParams {
	var ()
	return &IsStoppableParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewIsStoppableParamsWithTimeout creates a new IsStoppableParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewIsStoppableParamsWithTimeout(timeout time.Duration) *IsStoppableParams {
	var ()
	return &IsStoppableParams{

		timeout: timeout,
	}
}

// NewIsStoppableParamsWithContext creates a new IsStoppableParams object
// with the default values initialized, and the ability to set a context for a request
func NewIsStoppableParamsWithContext(ctx context.Context) *IsStoppableParams {
	var ()
	return &IsStoppableParams{

		Context: ctx,
	}
}

// NewIsStoppableParamsWithHTTPClient creates a new IsStoppableParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewIsStoppableParamsWithHTTPClient(client *http.Client) *IsStoppableParams {
	var ()
	return &IsStoppableParams{
		HTTPClient: client,
	}
}

/*
IsStoppableParams contains all the parameters to send to the API endpoint
for the is stoppable operation typically these are written to a http.Request
*/
type IsStoppableParams struct {

	/*Crn*/
	Crn string
	/*InitiatorUserCrn*/
	InitiatorUserCrn *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the is stoppable params
func (o *IsStoppableParams) WithTimeout(timeout time.Duration) *IsStoppableParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the is stoppable params
func (o *IsStoppableParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the is stoppable params
func (o *IsStoppableParams) WithContext(ctx context.Context) *IsStoppableParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the is stoppable params
func (o *IsStoppableParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the is stoppable params
func (o *IsStoppableParams) WithHTTPClient(client *http.Client) *IsStoppableParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the is stoppable params
func (o *IsStoppableParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCrn adds the crn to the is stoppable params
func (o *IsStoppableParams) WithCrn(crn string) *IsStoppableParams {
	o.SetCrn(crn)
	return o
}

// SetCrn adds the crn to the is stoppable params
func (o *IsStoppableParams) SetCrn(crn string) {
	o.Crn = crn
}

// WithInitiatorUserCrn adds the initiatorUserCrn to the is stoppable params
func (o *IsStoppableParams) WithInitiatorUserCrn(initiatorUserCrn *string) *IsStoppableParams {
	o.SetInitiatorUserCrn(initiatorUserCrn)
	return o
}

// SetInitiatorUserCrn adds the initiatorUserCrn to the is stoppable params
func (o *IsStoppableParams) SetInitiatorUserCrn(initiatorUserCrn *string) {
	o.InitiatorUserCrn = initiatorUserCrn
}

// WriteToRequest writes these params to a swagger request
func (o *IsStoppableParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param crn
	if err := r.SetPathParam("crn", o.Crn); err != nil {
		return err
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
