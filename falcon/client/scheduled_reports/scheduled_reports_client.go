// Code generated by go-swagger; DO NOT EDIT.

package scheduled_reports

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new scheduled reports API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for scheduled reports API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	ScheduledReportsGet(params *ScheduledReportsGetParams, opts ...ClientOption) (*ScheduledReportsGetOK, error)

	ScheduledReportsQuery(params *ScheduledReportsQueryParams, opts ...ClientOption) (*ScheduledReportsQueryOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  ScheduledReportsGet retrieves scheduled reports for the provided report i ds
*/
func (a *Client) ScheduledReportsGet(params *ScheduledReportsGetParams, opts ...ClientOption) (*ScheduledReportsGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewScheduledReportsGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "scheduled-reports.get",
		Method:             "GET",
		PathPattern:        "/reports/entities/scheduled-reports/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ScheduledReportsGetReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ScheduledReportsGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ScheduledReportsGetDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ScheduledReportsQuery finds all report i ds matching the query with filter
*/
func (a *Client) ScheduledReportsQuery(params *ScheduledReportsQueryParams, opts ...ClientOption) (*ScheduledReportsQueryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewScheduledReportsQueryParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "scheduled-reports.query",
		Method:             "GET",
		PathPattern:        "/reports/queries/scheduled-reports/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ScheduledReportsQueryReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ScheduledReportsQueryOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ScheduledReportsQueryDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
