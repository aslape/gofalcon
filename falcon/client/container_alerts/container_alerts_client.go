// Code generated by go-swagger; DO NOT EDIT.

package container_alerts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new container alerts API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for container alerts API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	ReadContainerAlertsCount(params *ReadContainerAlertsCountParams, opts ...ClientOption) (*ReadContainerAlertsCountOK, error)

	ReadContainerAlertsCountBySeverity(params *ReadContainerAlertsCountBySeverityParams, opts ...ClientOption) (*ReadContainerAlertsCountBySeverityOK, error)

	SearchAndReadContainerAlerts(params *SearchAndReadContainerAlertsParams, opts ...ClientOption) (*SearchAndReadContainerAlertsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
ReadContainerAlertsCount searches container alerts by the provided search criteria
*/
func (a *Client) ReadContainerAlertsCount(params *ReadContainerAlertsCountParams, opts ...ClientOption) (*ReadContainerAlertsCountOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReadContainerAlertsCountParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ReadContainerAlertsCount",
		Method:             "GET",
		PathPattern:        "/container-security/aggregates/container-alerts/count/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ReadContainerAlertsCountReader{formats: a.formats},
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
	success, ok := result.(*ReadContainerAlertsCountOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ReadContainerAlertsCount: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ReadContainerAlertsCountBySeverity gets container alerts counts by severity
*/
func (a *Client) ReadContainerAlertsCountBySeverity(params *ReadContainerAlertsCountBySeverityParams, opts ...ClientOption) (*ReadContainerAlertsCountBySeverityOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReadContainerAlertsCountBySeverityParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ReadContainerAlertsCountBySeverity",
		Method:             "GET",
		PathPattern:        "/container-security/aggregates/container-alerts/count-by-severity/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ReadContainerAlertsCountBySeverityReader{formats: a.formats},
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
	success, ok := result.(*ReadContainerAlertsCountBySeverityOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ReadContainerAlertsCountBySeverity: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
SearchAndReadContainerAlerts searches container alerts by the provided search criteria
*/
func (a *Client) SearchAndReadContainerAlerts(params *SearchAndReadContainerAlertsParams, opts ...ClientOption) (*SearchAndReadContainerAlertsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSearchAndReadContainerAlertsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "SearchAndReadContainerAlerts",
		Method:             "GET",
		PathPattern:        "/container-security/combined/container-alerts/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SearchAndReadContainerAlertsReader{formats: a.formats},
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
	success, ok := result.(*SearchAndReadContainerAlertsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for SearchAndReadContainerAlerts: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}