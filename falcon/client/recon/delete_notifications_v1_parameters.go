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

// NewDeleteNotificationsV1Params creates a new DeleteNotificationsV1Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteNotificationsV1Params() *DeleteNotificationsV1Params {
	return &DeleteNotificationsV1Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteNotificationsV1ParamsWithTimeout creates a new DeleteNotificationsV1Params object
// with the ability to set a timeout on a request.
func NewDeleteNotificationsV1ParamsWithTimeout(timeout time.Duration) *DeleteNotificationsV1Params {
	return &DeleteNotificationsV1Params{
		timeout: timeout,
	}
}

// NewDeleteNotificationsV1ParamsWithContext creates a new DeleteNotificationsV1Params object
// with the ability to set a context for a request.
func NewDeleteNotificationsV1ParamsWithContext(ctx context.Context) *DeleteNotificationsV1Params {
	return &DeleteNotificationsV1Params{
		Context: ctx,
	}
}

// NewDeleteNotificationsV1ParamsWithHTTPClient creates a new DeleteNotificationsV1Params object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteNotificationsV1ParamsWithHTTPClient(client *http.Client) *DeleteNotificationsV1Params {
	return &DeleteNotificationsV1Params{
		HTTPClient: client,
	}
}

/* DeleteNotificationsV1Params contains all the parameters to send to the API endpoint
   for the delete notifications v1 operation.

   Typically these are written to a http.Request.
*/
type DeleteNotificationsV1Params struct {

	/* Ids.

	   Notifications IDs.
	*/
	Ids []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete notifications v1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteNotificationsV1Params) WithDefaults() *DeleteNotificationsV1Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete notifications v1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteNotificationsV1Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete notifications v1 params
func (o *DeleteNotificationsV1Params) WithTimeout(timeout time.Duration) *DeleteNotificationsV1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete notifications v1 params
func (o *DeleteNotificationsV1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete notifications v1 params
func (o *DeleteNotificationsV1Params) WithContext(ctx context.Context) *DeleteNotificationsV1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete notifications v1 params
func (o *DeleteNotificationsV1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete notifications v1 params
func (o *DeleteNotificationsV1Params) WithHTTPClient(client *http.Client) *DeleteNotificationsV1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete notifications v1 params
func (o *DeleteNotificationsV1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIds adds the ids to the delete notifications v1 params
func (o *DeleteNotificationsV1Params) WithIds(ids []string) *DeleteNotificationsV1Params {
	o.SetIds(ids)
	return o
}

// SetIds adds the ids to the delete notifications v1 params
func (o *DeleteNotificationsV1Params) SetIds(ids []string) {
	o.Ids = ids
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteNotificationsV1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

// bindParamDeleteNotificationsV1 binds the parameter ids
func (o *DeleteNotificationsV1Params) bindParamIds(formats strfmt.Registry) []string {
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
