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

// NewQueryActionsV1Params creates a new QueryActionsV1Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewQueryActionsV1Params() *QueryActionsV1Params {
	return &QueryActionsV1Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewQueryActionsV1ParamsWithTimeout creates a new QueryActionsV1Params object
// with the ability to set a timeout on a request.
func NewQueryActionsV1ParamsWithTimeout(timeout time.Duration) *QueryActionsV1Params {
	return &QueryActionsV1Params{
		timeout: timeout,
	}
}

// NewQueryActionsV1ParamsWithContext creates a new QueryActionsV1Params object
// with the ability to set a context for a request.
func NewQueryActionsV1ParamsWithContext(ctx context.Context) *QueryActionsV1Params {
	return &QueryActionsV1Params{
		Context: ctx,
	}
}

// NewQueryActionsV1ParamsWithHTTPClient creates a new QueryActionsV1Params object
// with the ability to set a custom HTTPClient for a request.
func NewQueryActionsV1ParamsWithHTTPClient(client *http.Client) *QueryActionsV1Params {
	return &QueryActionsV1Params{
		HTTPClient: client,
	}
}

/* QueryActionsV1Params contains all the parameters to send to the API endpoint
   for the query actions v1 operation.

   Typically these are written to a http.Request.
*/
type QueryActionsV1Params struct {

	/* Filter.

	   FQL query to filter actions by. Possible filter properties are: [id cid user_uuid rule_id type frequency recipients status created_timestamp updated_timestamp]
	*/
	Filter *string

	/* Limit.

	   Number of IDs to return.
	*/
	Limit *int64

	/* Offset.

	   Starting index of overall result set from which to return IDs.
	*/
	Offset *string

	/* Q.

	   Free text search across all indexed fields
	*/
	Q *string

	/* Sort.

	   Possible order by fields: created_timestamp, updated_timestamp. Ex: 'updated_timestamp|desc'.
	*/
	Sort *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the query actions v1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *QueryActionsV1Params) WithDefaults() *QueryActionsV1Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the query actions v1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *QueryActionsV1Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the query actions v1 params
func (o *QueryActionsV1Params) WithTimeout(timeout time.Duration) *QueryActionsV1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the query actions v1 params
func (o *QueryActionsV1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the query actions v1 params
func (o *QueryActionsV1Params) WithContext(ctx context.Context) *QueryActionsV1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the query actions v1 params
func (o *QueryActionsV1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the query actions v1 params
func (o *QueryActionsV1Params) WithHTTPClient(client *http.Client) *QueryActionsV1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the query actions v1 params
func (o *QueryActionsV1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilter adds the filter to the query actions v1 params
func (o *QueryActionsV1Params) WithFilter(filter *string) *QueryActionsV1Params {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the query actions v1 params
func (o *QueryActionsV1Params) SetFilter(filter *string) {
	o.Filter = filter
}

// WithLimit adds the limit to the query actions v1 params
func (o *QueryActionsV1Params) WithLimit(limit *int64) *QueryActionsV1Params {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the query actions v1 params
func (o *QueryActionsV1Params) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the query actions v1 params
func (o *QueryActionsV1Params) WithOffset(offset *string) *QueryActionsV1Params {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the query actions v1 params
func (o *QueryActionsV1Params) SetOffset(offset *string) {
	o.Offset = offset
}

// WithQ adds the q to the query actions v1 params
func (o *QueryActionsV1Params) WithQ(q *string) *QueryActionsV1Params {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the query actions v1 params
func (o *QueryActionsV1Params) SetQ(q *string) {
	o.Q = q
}

// WithSort adds the sort to the query actions v1 params
func (o *QueryActionsV1Params) WithSort(sort *string) *QueryActionsV1Params {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the query actions v1 params
func (o *QueryActionsV1Params) SetSort(sort *string) {
	o.Sort = sort
}

// WriteToRequest writes these params to a swagger request
func (o *QueryActionsV1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Filter != nil {

		// query param filter
		var qrFilter string

		if o.Filter != nil {
			qrFilter = *o.Filter
		}
		qFilter := qrFilter
		if qFilter != "" {

			if err := r.SetQueryParam("filter", qFilter); err != nil {
				return err
			}
		}
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64

		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {

			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}
	}

	if o.Offset != nil {

		// query param offset
		var qrOffset string

		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := qrOffset
		if qOffset != "" {

			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}
	}

	if o.Q != nil {

		// query param q
		var qrQ string

		if o.Q != nil {
			qrQ = *o.Q
		}
		qQ := qrQ
		if qQ != "" {

			if err := r.SetQueryParam("q", qQ); err != nil {
				return err
			}
		}
	}

	if o.Sort != nil {

		// query param sort
		var qrSort string

		if o.Sort != nil {
			qrSort = *o.Sort
		}
		qSort := qrSort
		if qSort != "" {

			if err := r.SetQueryParam("sort", qSort); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
