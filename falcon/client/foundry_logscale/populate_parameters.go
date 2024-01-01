// Code generated by go-swagger; DO NOT EDIT.

package foundry_logscale

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewPopulateParams creates a new PopulateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPopulateParams() *PopulateParams {
	return &PopulateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPopulateParamsWithTimeout creates a new PopulateParams object
// with the ability to set a timeout on a request.
func NewPopulateParamsWithTimeout(timeout time.Duration) *PopulateParams {
	return &PopulateParams{
		timeout: timeout,
	}
}

// NewPopulateParamsWithContext creates a new PopulateParams object
// with the ability to set a context for a request.
func NewPopulateParamsWithContext(ctx context.Context) *PopulateParams {
	return &PopulateParams{
		Context: ctx,
	}
}

// NewPopulateParamsWithHTTPClient creates a new PopulateParams object
// with the ability to set a custom HTTPClient for a request.
func NewPopulateParamsWithHTTPClient(client *http.Client) *PopulateParams {
	return &PopulateParams{
		HTTPClient: client,
	}
}

/*
PopulateParams contains all the parameters to send to the API endpoint

	for the populate operation.

	Typically these are written to a http.Request.
*/
type PopulateParams struct {

	/* AppID.

	   Application ID.
	*/
	AppID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the populate params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PopulateParams) WithDefaults() *PopulateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the populate params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PopulateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the populate params
func (o *PopulateParams) WithTimeout(timeout time.Duration) *PopulateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the populate params
func (o *PopulateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the populate params
func (o *PopulateParams) WithContext(ctx context.Context) *PopulateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the populate params
func (o *PopulateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the populate params
func (o *PopulateParams) WithHTTPClient(client *http.Client) *PopulateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the populate params
func (o *PopulateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppID adds the appID to the populate params
func (o *PopulateParams) WithAppID(appID *string) *PopulateParams {
	o.SetAppID(appID)
	return o
}

// SetAppID adds the appId to the populate params
func (o *PopulateParams) SetAppID(appID *string) {
	o.AppID = appID
}

// WriteToRequest writes these params to a swagger request
func (o *PopulateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AppID != nil {

		// query param app_id
		var qrAppID string

		if o.AppID != nil {
			qrAppID = *o.AppID
		}
		qAppID := qrAppID
		if qAppID != "" {

			if err := r.SetQueryParam("app_id", qAppID); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}