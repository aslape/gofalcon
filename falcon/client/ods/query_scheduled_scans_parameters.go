// Code generated by go-swagger; DO NOT EDIT.

package ods

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

// NewQueryScheduledScansParams creates a new QueryScheduledScansParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewQueryScheduledScansParams() *QueryScheduledScansParams {
	return &QueryScheduledScansParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewQueryScheduledScansParamsWithTimeout creates a new QueryScheduledScansParams object
// with the ability to set a timeout on a request.
func NewQueryScheduledScansParamsWithTimeout(timeout time.Duration) *QueryScheduledScansParams {
	return &QueryScheduledScansParams{
		timeout: timeout,
	}
}

// NewQueryScheduledScansParamsWithContext creates a new QueryScheduledScansParams object
// with the ability to set a context for a request.
func NewQueryScheduledScansParamsWithContext(ctx context.Context) *QueryScheduledScansParams {
	return &QueryScheduledScansParams{
		Context: ctx,
	}
}

// NewQueryScheduledScansParamsWithHTTPClient creates a new QueryScheduledScansParams object
// with the ability to set a custom HTTPClient for a request.
func NewQueryScheduledScansParamsWithHTTPClient(client *http.Client) *QueryScheduledScansParams {
	return &QueryScheduledScansParams{
		HTTPClient: client,
	}
}

/*
QueryScheduledScansParams contains all the parameters to send to the API endpoint

	for the query scheduled scans operation.

	Typically these are written to a http.Request.
*/
type QueryScheduledScansParams struct {

	/* XCSUSERUUID.

	   The user ID
	*/
	XCSUSERUUID string

	/* Filter.

	   A FQL compatible query string. Terms: [id description initiated_from status schedule.start_timestamp schedule.Interval created_on created_by last_updated deleted]
	*/
	Filter string

	/* Limit.

	   The max number of resources to return

	   Default: 500
	*/
	Limit *int64

	/* Offset.

	   Index of the starting resource
	*/
	Offset *int64

	/* Sort.

	   The property to sort on, followed by a |, followed by the sort direction, either "asc" or "desc"

	   Default: "schedule.start_timestamp|desc"
	*/
	Sort string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the query scheduled scans params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *QueryScheduledScansParams) WithDefaults() *QueryScheduledScansParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the query scheduled scans params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *QueryScheduledScansParams) SetDefaults() {
	var (
		limitDefault = int64(500)

		offsetDefault = int64(0)

		sortDefault = string("schedule.start_timestamp|desc")
	)

	val := QueryScheduledScansParams{
		Limit:  &limitDefault,
		Offset: &offsetDefault,
		Sort:   sortDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the query scheduled scans params
func (o *QueryScheduledScansParams) WithTimeout(timeout time.Duration) *QueryScheduledScansParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the query scheduled scans params
func (o *QueryScheduledScansParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the query scheduled scans params
func (o *QueryScheduledScansParams) WithContext(ctx context.Context) *QueryScheduledScansParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the query scheduled scans params
func (o *QueryScheduledScansParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the query scheduled scans params
func (o *QueryScheduledScansParams) WithHTTPClient(client *http.Client) *QueryScheduledScansParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the query scheduled scans params
func (o *QueryScheduledScansParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXCSUSERUUID adds the xCSUSERUUID to the query scheduled scans params
func (o *QueryScheduledScansParams) WithXCSUSERUUID(xCSUSERUUID string) *QueryScheduledScansParams {
	o.SetXCSUSERUUID(xCSUSERUUID)
	return o
}

// SetXCSUSERUUID adds the xCSUSERUuid to the query scheduled scans params
func (o *QueryScheduledScansParams) SetXCSUSERUUID(xCSUSERUUID string) {
	o.XCSUSERUUID = xCSUSERUUID
}

// WithFilter adds the filter to the query scheduled scans params
func (o *QueryScheduledScansParams) WithFilter(filter string) *QueryScheduledScansParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the query scheduled scans params
func (o *QueryScheduledScansParams) SetFilter(filter string) {
	o.Filter = filter
}

// WithLimit adds the limit to the query scheduled scans params
func (o *QueryScheduledScansParams) WithLimit(limit *int64) *QueryScheduledScansParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the query scheduled scans params
func (o *QueryScheduledScansParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the query scheduled scans params
func (o *QueryScheduledScansParams) WithOffset(offset *int64) *QueryScheduledScansParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the query scheduled scans params
func (o *QueryScheduledScansParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithSort adds the sort to the query scheduled scans params
func (o *QueryScheduledScansParams) WithSort(sort string) *QueryScheduledScansParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the query scheduled scans params
func (o *QueryScheduledScansParams) SetSort(sort string) {
	o.Sort = sort
}

// WriteToRequest writes these params to a swagger request
func (o *QueryScheduledScansParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param X-CS-USERUUID
	if err := r.SetHeaderParam("X-CS-USERUUID", o.XCSUSERUUID); err != nil {
		return err
	}

	// query param filter
	qrFilter := o.Filter
	qFilter := qrFilter

	if err := r.SetQueryParam("filter", qFilter); err != nil {
		return err
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
		var qrOffset int64

		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {

			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}
	}

	// query param sort
	qrSort := o.Sort
	qSort := qrSort

	if err := r.SetQueryParam("sort", qSort); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
