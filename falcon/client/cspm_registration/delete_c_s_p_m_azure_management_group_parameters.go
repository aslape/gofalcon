// Code generated by go-swagger; DO NOT EDIT.

package cspm_registration

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

// NewDeleteCSPMAzureManagementGroupParams creates a new DeleteCSPMAzureManagementGroupParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteCSPMAzureManagementGroupParams() *DeleteCSPMAzureManagementGroupParams {
	return &DeleteCSPMAzureManagementGroupParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteCSPMAzureManagementGroupParamsWithTimeout creates a new DeleteCSPMAzureManagementGroupParams object
// with the ability to set a timeout on a request.
func NewDeleteCSPMAzureManagementGroupParamsWithTimeout(timeout time.Duration) *DeleteCSPMAzureManagementGroupParams {
	return &DeleteCSPMAzureManagementGroupParams{
		timeout: timeout,
	}
}

// NewDeleteCSPMAzureManagementGroupParamsWithContext creates a new DeleteCSPMAzureManagementGroupParams object
// with the ability to set a context for a request.
func NewDeleteCSPMAzureManagementGroupParamsWithContext(ctx context.Context) *DeleteCSPMAzureManagementGroupParams {
	return &DeleteCSPMAzureManagementGroupParams{
		Context: ctx,
	}
}

// NewDeleteCSPMAzureManagementGroupParamsWithHTTPClient creates a new DeleteCSPMAzureManagementGroupParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteCSPMAzureManagementGroupParamsWithHTTPClient(client *http.Client) *DeleteCSPMAzureManagementGroupParams {
	return &DeleteCSPMAzureManagementGroupParams{
		HTTPClient: client,
	}
}

/*
DeleteCSPMAzureManagementGroupParams contains all the parameters to send to the API endpoint

	for the delete c s p m azure management group operation.

	Typically these are written to a http.Request.
*/
type DeleteCSPMAzureManagementGroupParams struct {

	/* TenantIds.

	   Tenant ids to remove
	*/
	TenantIds []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete c s p m azure management group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteCSPMAzureManagementGroupParams) WithDefaults() *DeleteCSPMAzureManagementGroupParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete c s p m azure management group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteCSPMAzureManagementGroupParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete c s p m azure management group params
func (o *DeleteCSPMAzureManagementGroupParams) WithTimeout(timeout time.Duration) *DeleteCSPMAzureManagementGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete c s p m azure management group params
func (o *DeleteCSPMAzureManagementGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete c s p m azure management group params
func (o *DeleteCSPMAzureManagementGroupParams) WithContext(ctx context.Context) *DeleteCSPMAzureManagementGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete c s p m azure management group params
func (o *DeleteCSPMAzureManagementGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete c s p m azure management group params
func (o *DeleteCSPMAzureManagementGroupParams) WithHTTPClient(client *http.Client) *DeleteCSPMAzureManagementGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete c s p m azure management group params
func (o *DeleteCSPMAzureManagementGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTenantIds adds the tenantIds to the delete c s p m azure management group params
func (o *DeleteCSPMAzureManagementGroupParams) WithTenantIds(tenantIds []string) *DeleteCSPMAzureManagementGroupParams {
	o.SetTenantIds(tenantIds)
	return o
}

// SetTenantIds adds the tenantIds to the delete c s p m azure management group params
func (o *DeleteCSPMAzureManagementGroupParams) SetTenantIds(tenantIds []string) {
	o.TenantIds = tenantIds
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteCSPMAzureManagementGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.TenantIds != nil {

		// binding items for tenant_ids
		joinedTenantIds := o.bindParamTenantIds(reg)

		// query array param tenant_ids
		if err := r.SetQueryParam("tenant_ids", joinedTenantIds...); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamDeleteCSPMAzureManagementGroup binds the parameter tenant_ids
func (o *DeleteCSPMAzureManagementGroupParams) bindParamTenantIds(formats strfmt.Registry) []string {
	tenantIdsIR := o.TenantIds

	var tenantIdsIC []string
	for _, tenantIdsIIR := range tenantIdsIR { // explode []string

		tenantIdsIIV := tenantIdsIIR // string as string
		tenantIdsIC = append(tenantIdsIC, tenantIdsIIV)
	}

	// items.CollectionFormat: "multi"
	tenantIdsIS := swag.JoinByFormat(tenantIdsIC, "multi")

	return tenantIdsIS
}
