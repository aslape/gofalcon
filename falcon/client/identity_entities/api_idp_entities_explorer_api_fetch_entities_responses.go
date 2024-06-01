// Code generated by go-swagger; DO NOT EDIT.

package identity_entities

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/crowdstrike/gofalcon/falcon/models"
)

// APIIdpEntitiesExplorerAPIFetchEntitiesReader is a Reader for the APIIdpEntitiesExplorerAPIFetchEntities structure.
type APIIdpEntitiesExplorerAPIFetchEntitiesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *APIIdpEntitiesExplorerAPIFetchEntitiesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAPIIdpEntitiesExplorerAPIFetchEntitiesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewAPIIdpEntitiesExplorerAPIFetchEntitiesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewAPIIdpEntitiesExplorerAPIFetchEntitiesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /idp-entities-explorer/entities/entities/v1] api.idp-entities-explorer.api.fetch-entities", response, response.Code())
	}
}

// NewAPIIdpEntitiesExplorerAPIFetchEntitiesOK creates a APIIdpEntitiesExplorerAPIFetchEntitiesOK with default headers values
func NewAPIIdpEntitiesExplorerAPIFetchEntitiesOK() *APIIdpEntitiesExplorerAPIFetchEntitiesOK {
	return &APIIdpEntitiesExplorerAPIFetchEntitiesOK{}
}

/*
APIIdpEntitiesExplorerAPIFetchEntitiesOK describes a response with status code 200, with default header values.

OK
*/
type APIIdpEntitiesExplorerAPIFetchEntitiesOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.EntitiesEntitiesFetchResponse
}

// IsSuccess returns true when this api idp entities explorer Api fetch entities o k response has a 2xx status code
func (o *APIIdpEntitiesExplorerAPIFetchEntitiesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this api idp entities explorer Api fetch entities o k response has a 3xx status code
func (o *APIIdpEntitiesExplorerAPIFetchEntitiesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this api idp entities explorer Api fetch entities o k response has a 4xx status code
func (o *APIIdpEntitiesExplorerAPIFetchEntitiesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this api idp entities explorer Api fetch entities o k response has a 5xx status code
func (o *APIIdpEntitiesExplorerAPIFetchEntitiesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this api idp entities explorer Api fetch entities o k response a status code equal to that given
func (o *APIIdpEntitiesExplorerAPIFetchEntitiesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the api idp entities explorer Api fetch entities o k response
func (o *APIIdpEntitiesExplorerAPIFetchEntitiesOK) Code() int {
	return 200
}

func (o *APIIdpEntitiesExplorerAPIFetchEntitiesOK) Error() string {
	return fmt.Sprintf("[POST /idp-entities-explorer/entities/entities/v1][%d] apiIdpEntitiesExplorerApiFetchEntitiesOK  %+v", 200, o.Payload)
}

func (o *APIIdpEntitiesExplorerAPIFetchEntitiesOK) String() string {
	return fmt.Sprintf("[POST /idp-entities-explorer/entities/entities/v1][%d] apiIdpEntitiesExplorerApiFetchEntitiesOK  %+v", 200, o.Payload)
}

func (o *APIIdpEntitiesExplorerAPIFetchEntitiesOK) GetPayload() *models.EntitiesEntitiesFetchResponse {
	return o.Payload
}

func (o *APIIdpEntitiesExplorerAPIFetchEntitiesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-CS-TRACEID
	hdrXCSTRACEID := response.GetHeader("X-CS-TRACEID")

	if hdrXCSTRACEID != "" {
		o.XCSTRACEID = hdrXCSTRACEID
	}

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	o.Payload = new(models.EntitiesEntitiesFetchResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAPIIdpEntitiesExplorerAPIFetchEntitiesForbidden creates a APIIdpEntitiesExplorerAPIFetchEntitiesForbidden with default headers values
func NewAPIIdpEntitiesExplorerAPIFetchEntitiesForbidden() *APIIdpEntitiesExplorerAPIFetchEntitiesForbidden {
	return &APIIdpEntitiesExplorerAPIFetchEntitiesForbidden{}
}

/*
APIIdpEntitiesExplorerAPIFetchEntitiesForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type APIIdpEntitiesExplorerAPIFetchEntitiesForbidden struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaReplyMetaOnly
}

// IsSuccess returns true when this api idp entities explorer Api fetch entities forbidden response has a 2xx status code
func (o *APIIdpEntitiesExplorerAPIFetchEntitiesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this api idp entities explorer Api fetch entities forbidden response has a 3xx status code
func (o *APIIdpEntitiesExplorerAPIFetchEntitiesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this api idp entities explorer Api fetch entities forbidden response has a 4xx status code
func (o *APIIdpEntitiesExplorerAPIFetchEntitiesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this api idp entities explorer Api fetch entities forbidden response has a 5xx status code
func (o *APIIdpEntitiesExplorerAPIFetchEntitiesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this api idp entities explorer Api fetch entities forbidden response a status code equal to that given
func (o *APIIdpEntitiesExplorerAPIFetchEntitiesForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the api idp entities explorer Api fetch entities forbidden response
func (o *APIIdpEntitiesExplorerAPIFetchEntitiesForbidden) Code() int {
	return 403
}

func (o *APIIdpEntitiesExplorerAPIFetchEntitiesForbidden) Error() string {
	return fmt.Sprintf("[POST /idp-entities-explorer/entities/entities/v1][%d] apiIdpEntitiesExplorerApiFetchEntitiesForbidden  %+v", 403, o.Payload)
}

func (o *APIIdpEntitiesExplorerAPIFetchEntitiesForbidden) String() string {
	return fmt.Sprintf("[POST /idp-entities-explorer/entities/entities/v1][%d] apiIdpEntitiesExplorerApiFetchEntitiesForbidden  %+v", 403, o.Payload)
}

func (o *APIIdpEntitiesExplorerAPIFetchEntitiesForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *APIIdpEntitiesExplorerAPIFetchEntitiesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-CS-TRACEID
	hdrXCSTRACEID := response.GetHeader("X-CS-TRACEID")

	if hdrXCSTRACEID != "" {
		o.XCSTRACEID = hdrXCSTRACEID
	}

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	o.Payload = new(models.MsaReplyMetaOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAPIIdpEntitiesExplorerAPIFetchEntitiesTooManyRequests creates a APIIdpEntitiesExplorerAPIFetchEntitiesTooManyRequests with default headers values
func NewAPIIdpEntitiesExplorerAPIFetchEntitiesTooManyRequests() *APIIdpEntitiesExplorerAPIFetchEntitiesTooManyRequests {
	return &APIIdpEntitiesExplorerAPIFetchEntitiesTooManyRequests{}
}

/*
APIIdpEntitiesExplorerAPIFetchEntitiesTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type APIIdpEntitiesExplorerAPIFetchEntitiesTooManyRequests struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	/* Too many requests, retry after this time (as milliseconds since epoch)
	 */
	XRateLimitRetryAfter int64

	Payload *models.MsaReplyMetaOnly
}

// IsSuccess returns true when this api idp entities explorer Api fetch entities too many requests response has a 2xx status code
func (o *APIIdpEntitiesExplorerAPIFetchEntitiesTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this api idp entities explorer Api fetch entities too many requests response has a 3xx status code
func (o *APIIdpEntitiesExplorerAPIFetchEntitiesTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this api idp entities explorer Api fetch entities too many requests response has a 4xx status code
func (o *APIIdpEntitiesExplorerAPIFetchEntitiesTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this api idp entities explorer Api fetch entities too many requests response has a 5xx status code
func (o *APIIdpEntitiesExplorerAPIFetchEntitiesTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this api idp entities explorer Api fetch entities too many requests response a status code equal to that given
func (o *APIIdpEntitiesExplorerAPIFetchEntitiesTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the api idp entities explorer Api fetch entities too many requests response
func (o *APIIdpEntitiesExplorerAPIFetchEntitiesTooManyRequests) Code() int {
	return 429
}

func (o *APIIdpEntitiesExplorerAPIFetchEntitiesTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /idp-entities-explorer/entities/entities/v1][%d] apiIdpEntitiesExplorerApiFetchEntitiesTooManyRequests  %+v", 429, o.Payload)
}

func (o *APIIdpEntitiesExplorerAPIFetchEntitiesTooManyRequests) String() string {
	return fmt.Sprintf("[POST /idp-entities-explorer/entities/entities/v1][%d] apiIdpEntitiesExplorerApiFetchEntitiesTooManyRequests  %+v", 429, o.Payload)
}

func (o *APIIdpEntitiesExplorerAPIFetchEntitiesTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *APIIdpEntitiesExplorerAPIFetchEntitiesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-CS-TRACEID
	hdrXCSTRACEID := response.GetHeader("X-CS-TRACEID")

	if hdrXCSTRACEID != "" {
		o.XCSTRACEID = hdrXCSTRACEID
	}

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	// hydrates response header X-RateLimit-RetryAfter
	hdrXRateLimitRetryAfter := response.GetHeader("X-RateLimit-RetryAfter")

	if hdrXRateLimitRetryAfter != "" {
		valxRateLimitRetryAfter, err := swag.ConvertInt64(hdrXRateLimitRetryAfter)
		if err != nil {
			return errors.InvalidType("X-RateLimit-RetryAfter", "header", "int64", hdrXRateLimitRetryAfter)
		}
		o.XRateLimitRetryAfter = valxRateLimitRetryAfter
	}

	o.Payload = new(models.MsaReplyMetaOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}