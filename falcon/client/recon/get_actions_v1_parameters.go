// Code generated by go-swagger; DO NOT EDIT.

package recon

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
	"github.com/go-openapi/swag"
)

// NewGetActionsV1Params creates a new GetActionsV1Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetActionsV1Params() *GetActionsV1Params {
	return &GetActionsV1Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetActionsV1ParamsWithTimeout creates a new GetActionsV1Params object
// with the ability to set a timeout on a request.
func NewGetActionsV1ParamsWithTimeout(timeout time.Duration) *GetActionsV1Params {
	return &GetActionsV1Params{
		timeout: timeout,
	}
}

// NewGetActionsV1ParamsWithContext creates a new GetActionsV1Params object
// with the ability to set a context for a request.
func NewGetActionsV1ParamsWithContext(ctx context.Context) *GetActionsV1Params {
	return &GetActionsV1Params{
		Context: ctx,
	}
}

// NewGetActionsV1ParamsWithHTTPClient creates a new GetActionsV1Params object
// with the ability to set a custom HTTPClient for a request.
func NewGetActionsV1ParamsWithHTTPClient(client *http.Client) *GetActionsV1Params {
	return &GetActionsV1Params{
		HTTPClient: client,
	}
}

/* GetActionsV1Params contains all the parameters to send to the API endpoint
   for the get actions v1 operation.

   Typically these are written to a http.Request.
*/
type GetActionsV1Params struct {

	/* Ids.

	   Action IDs.
	*/
	Ids []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get actions v1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetActionsV1Params) WithDefaults() *GetActionsV1Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get actions v1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetActionsV1Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get actions v1 params
func (o *GetActionsV1Params) WithTimeout(timeout time.Duration) *GetActionsV1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get actions v1 params
func (o *GetActionsV1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get actions v1 params
func (o *GetActionsV1Params) WithContext(ctx context.Context) *GetActionsV1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get actions v1 params
func (o *GetActionsV1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get actions v1 params
func (o *GetActionsV1Params) WithHTTPClient(client *http.Client) *GetActionsV1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get actions v1 params
func (o *GetActionsV1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIds adds the ids to the get actions v1 params
func (o *GetActionsV1Params) WithIds(ids []string) *GetActionsV1Params {
	o.SetIds(ids)
	return o
}

// SetIds adds the ids to the get actions v1 params
func (o *GetActionsV1Params) SetIds(ids []string) {
	o.Ids = ids
}

// WriteToRequest writes these params to a swagger request
func (o *GetActionsV1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Ids != nil {

		// binding items for ids
		joinedIds := o.bindParamIds(reg)

		// query array param ids
		if err := r.SetQueryParam("ids", joinedIds...); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetActionsV1 binds the parameter ids
func (o *GetActionsV1Params) bindParamIds(formats strfmt.Registry) []string {
	idsIR := o.Ids

	var idsIC []string
	for _, idsIIR := range idsIR { // explode []string

		idsIIV := idsIIR // string as string
		idsIC = append(idsIC, idsIIV)
	}

	// items.CollectionFormat: "multi"
	idsIS := swag.JoinByFormat(idsIC, "multi")

	return idsIS
}
