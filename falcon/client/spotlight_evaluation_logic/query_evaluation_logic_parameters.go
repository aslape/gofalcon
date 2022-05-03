// Code generated by go-swagger; DO NOT EDIT.

package spotlight_evaluation_logic

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

// NewQueryEvaluationLogicParams creates a new QueryEvaluationLogicParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewQueryEvaluationLogicParams() *QueryEvaluationLogicParams {
	return &QueryEvaluationLogicParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewQueryEvaluationLogicParamsWithTimeout creates a new QueryEvaluationLogicParams object
// with the ability to set a timeout on a request.
func NewQueryEvaluationLogicParamsWithTimeout(timeout time.Duration) *QueryEvaluationLogicParams {
	return &QueryEvaluationLogicParams{
		timeout: timeout,
	}
}

// NewQueryEvaluationLogicParamsWithContext creates a new QueryEvaluationLogicParams object
// with the ability to set a context for a request.
func NewQueryEvaluationLogicParamsWithContext(ctx context.Context) *QueryEvaluationLogicParams {
	return &QueryEvaluationLogicParams{
		Context: ctx,
	}
}

// NewQueryEvaluationLogicParamsWithHTTPClient creates a new QueryEvaluationLogicParams object
// with the ability to set a custom HTTPClient for a request.
func NewQueryEvaluationLogicParamsWithHTTPClient(client *http.Client) *QueryEvaluationLogicParams {
	return &QueryEvaluationLogicParams{
		HTTPClient: client,
	}
}

/* QueryEvaluationLogicParams contains all the parameters to send to the API endpoint
   for the query evaluation logic operation.

   Typically these are written to a http.Request.
*/
type QueryEvaluationLogicParams struct {

	/* After.

	   A pagination token used with the `limit` parameter to manage pagination of results. On your first request, don't provide an `after` token. On subsequent requests, provide the `after` token from the previous response to continue from that place in the results.
	*/
	After *string

	/* Filter.

	   FQL query specifying the filter parameters.
	*/
	Filter string

	/* Limit.

	   Maximum number of entities to return.
	*/
	Limit *int64

	/* Sort.

	   Sort evaluation logic by their properties.
	*/
	Sort *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the query evaluation logic params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *QueryEvaluationLogicParams) WithDefaults() *QueryEvaluationLogicParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the query evaluation logic params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *QueryEvaluationLogicParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the query evaluation logic params
func (o *QueryEvaluationLogicParams) WithTimeout(timeout time.Duration) *QueryEvaluationLogicParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the query evaluation logic params
func (o *QueryEvaluationLogicParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the query evaluation logic params
func (o *QueryEvaluationLogicParams) WithContext(ctx context.Context) *QueryEvaluationLogicParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the query evaluation logic params
func (o *QueryEvaluationLogicParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the query evaluation logic params
func (o *QueryEvaluationLogicParams) WithHTTPClient(client *http.Client) *QueryEvaluationLogicParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the query evaluation logic params
func (o *QueryEvaluationLogicParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAfter adds the after to the query evaluation logic params
func (o *QueryEvaluationLogicParams) WithAfter(after *string) *QueryEvaluationLogicParams {
	o.SetAfter(after)
	return o
}

// SetAfter adds the after to the query evaluation logic params
func (o *QueryEvaluationLogicParams) SetAfter(after *string) {
	o.After = after
}

// WithFilter adds the filter to the query evaluation logic params
func (o *QueryEvaluationLogicParams) WithFilter(filter string) *QueryEvaluationLogicParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the query evaluation logic params
func (o *QueryEvaluationLogicParams) SetFilter(filter string) {
	o.Filter = filter
}

// WithLimit adds the limit to the query evaluation logic params
func (o *QueryEvaluationLogicParams) WithLimit(limit *int64) *QueryEvaluationLogicParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the query evaluation logic params
func (o *QueryEvaluationLogicParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithSort adds the sort to the query evaluation logic params
func (o *QueryEvaluationLogicParams) WithSort(sort *string) *QueryEvaluationLogicParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the query evaluation logic params
func (o *QueryEvaluationLogicParams) SetSort(sort *string) {
	o.Sort = sort
}

// WriteToRequest writes these params to a swagger request
func (o *QueryEvaluationLogicParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.After != nil {

		// query param after
		var qrAfter string

		if o.After != nil {
			qrAfter = *o.After
		}
		qAfter := qrAfter
		if qAfter != "" {

			if err := r.SetQueryParam("after", qAfter); err != nil {
				return err
			}
		}
	}

	// query param filter
	qrFilter := o.Filter
	qFilter := qrFilter
	if qFilter != "" {

		if err := r.SetQueryParam("filter", qFilter); err != nil {
			return err
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