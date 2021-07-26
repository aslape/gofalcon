// Code generated by go-swagger; DO NOT EDIT.

package report_executions

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

// NewReportExecutionsQueryParams creates a new ReportExecutionsQueryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewReportExecutionsQueryParams() *ReportExecutionsQueryParams {
	return &ReportExecutionsQueryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewReportExecutionsQueryParamsWithTimeout creates a new ReportExecutionsQueryParams object
// with the ability to set a timeout on a request.
func NewReportExecutionsQueryParamsWithTimeout(timeout time.Duration) *ReportExecutionsQueryParams {
	return &ReportExecutionsQueryParams{
		timeout: timeout,
	}
}

// NewReportExecutionsQueryParamsWithContext creates a new ReportExecutionsQueryParams object
// with the ability to set a context for a request.
func NewReportExecutionsQueryParamsWithContext(ctx context.Context) *ReportExecutionsQueryParams {
	return &ReportExecutionsQueryParams{
		Context: ctx,
	}
}

// NewReportExecutionsQueryParamsWithHTTPClient creates a new ReportExecutionsQueryParams object
// with the ability to set a custom HTTPClient for a request.
func NewReportExecutionsQueryParamsWithHTTPClient(client *http.Client) *ReportExecutionsQueryParams {
	return &ReportExecutionsQueryParams{
		HTTPClient: client,
	}
}

/* ReportExecutionsQueryParams contains all the parameters to send to the API endpoint
   for the report executions query operation.

   Typically these are written to a http.Request.
*/
type ReportExecutionsQueryParams struct {

	/* XCSUSERID.

	   The user id (not used with API client)
	*/
	XCSUSERID *string

	/* XCSUSERUUID.

	   The user uuid (not used with API client)
	*/
	XCSUSERUUID *string

	/* Filter.

	   FQL query specifying the filter parameters. Filter term criteria: type, scheduled_report_id, status. Filter range criteria: created_on, last_updated_on, expiration_on; use any common date format, such as '2010-05-15T14:55:21.892315096Z'.
	*/
	Filter *string

	/* Limit.

	   Number of ids to return.
	*/
	Limit *int64

	/* Offset.

	   Starting index of overall result set from which to return ids.
	*/
	Offset *string

	/* Q.

	   Match query criteria, which includes all the filter string fields
	*/
	Q *string

	/* Sort.

	   Possible order by fields: created_on, last_updated_on
	*/
	Sort *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the report executions query params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReportExecutionsQueryParams) WithDefaults() *ReportExecutionsQueryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the report executions query params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReportExecutionsQueryParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the report executions query params
func (o *ReportExecutionsQueryParams) WithTimeout(timeout time.Duration) *ReportExecutionsQueryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the report executions query params
func (o *ReportExecutionsQueryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the report executions query params
func (o *ReportExecutionsQueryParams) WithContext(ctx context.Context) *ReportExecutionsQueryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the report executions query params
func (o *ReportExecutionsQueryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the report executions query params
func (o *ReportExecutionsQueryParams) WithHTTPClient(client *http.Client) *ReportExecutionsQueryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the report executions query params
func (o *ReportExecutionsQueryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXCSUSERID adds the xCSUSERID to the report executions query params
func (o *ReportExecutionsQueryParams) WithXCSUSERID(xCSUSERID *string) *ReportExecutionsQueryParams {
	o.SetXCSUSERID(xCSUSERID)
	return o
}

// SetXCSUSERID adds the xCSUSERId to the report executions query params
func (o *ReportExecutionsQueryParams) SetXCSUSERID(xCSUSERID *string) {
	o.XCSUSERID = xCSUSERID
}

// WithXCSUSERUUID adds the xCSUSERUUID to the report executions query params
func (o *ReportExecutionsQueryParams) WithXCSUSERUUID(xCSUSERUUID *string) *ReportExecutionsQueryParams {
	o.SetXCSUSERUUID(xCSUSERUUID)
	return o
}

// SetXCSUSERUUID adds the xCSUSERUuid to the report executions query params
func (o *ReportExecutionsQueryParams) SetXCSUSERUUID(xCSUSERUUID *string) {
	o.XCSUSERUUID = xCSUSERUUID
}

// WithFilter adds the filter to the report executions query params
func (o *ReportExecutionsQueryParams) WithFilter(filter *string) *ReportExecutionsQueryParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the report executions query params
func (o *ReportExecutionsQueryParams) SetFilter(filter *string) {
	o.Filter = filter
}

// WithLimit adds the limit to the report executions query params
func (o *ReportExecutionsQueryParams) WithLimit(limit *int64) *ReportExecutionsQueryParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the report executions query params
func (o *ReportExecutionsQueryParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the report executions query params
func (o *ReportExecutionsQueryParams) WithOffset(offset *string) *ReportExecutionsQueryParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the report executions query params
func (o *ReportExecutionsQueryParams) SetOffset(offset *string) {
	o.Offset = offset
}

// WithQ adds the q to the report executions query params
func (o *ReportExecutionsQueryParams) WithQ(q *string) *ReportExecutionsQueryParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the report executions query params
func (o *ReportExecutionsQueryParams) SetQ(q *string) {
	o.Q = q
}

// WithSort adds the sort to the report executions query params
func (o *ReportExecutionsQueryParams) WithSort(sort *string) *ReportExecutionsQueryParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the report executions query params
func (o *ReportExecutionsQueryParams) SetSort(sort *string) {
	o.Sort = sort
}

// WriteToRequest writes these params to a swagger request
func (o *ReportExecutionsQueryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.XCSUSERID != nil {

		// header param X-CS-USERID
		if err := r.SetHeaderParam("X-CS-USERID", *o.XCSUSERID); err != nil {
			return err
		}
	}

	if o.XCSUSERUUID != nil {

		// header param X-CS-USERUUID
		if err := r.SetHeaderParam("X-CS-USERUUID", *o.XCSUSERUUID); err != nil {
			return err
		}
	}

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
