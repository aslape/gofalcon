// Code generated by go-swagger; DO NOT EDIT.

package kubernetes_protection

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

// ContainerVulnerabilitiesBySeverityCountReader is a Reader for the ContainerVulnerabilitiesBySeverityCount structure.
type ContainerVulnerabilitiesBySeverityCountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ContainerVulnerabilitiesBySeverityCountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewContainerVulnerabilitiesBySeverityCountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewContainerVulnerabilitiesBySeverityCountForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewContainerVulnerabilitiesBySeverityCountTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewContainerVulnerabilitiesBySeverityCountInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /container-security/aggregates/containers/vulnerability-count-by-severity/v1] ContainerVulnerabilitiesBySeverityCount", response, response.Code())
	}
}

// NewContainerVulnerabilitiesBySeverityCountOK creates a ContainerVulnerabilitiesBySeverityCountOK with default headers values
func NewContainerVulnerabilitiesBySeverityCountOK() *ContainerVulnerabilitiesBySeverityCountOK {
	return &ContainerVulnerabilitiesBySeverityCountOK{}
}

/*
ContainerVulnerabilitiesBySeverityCountOK describes a response with status code 200, with default header values.

OK
*/
type ContainerVulnerabilitiesBySeverityCountOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ModelsAggregateValuesByFieldResponse
}

// IsSuccess returns true when this container vulnerabilities by severity count o k response has a 2xx status code
func (o *ContainerVulnerabilitiesBySeverityCountOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this container vulnerabilities by severity count o k response has a 3xx status code
func (o *ContainerVulnerabilitiesBySeverityCountOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this container vulnerabilities by severity count o k response has a 4xx status code
func (o *ContainerVulnerabilitiesBySeverityCountOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this container vulnerabilities by severity count o k response has a 5xx status code
func (o *ContainerVulnerabilitiesBySeverityCountOK) IsServerError() bool {
	return false
}

// IsCode returns true when this container vulnerabilities by severity count o k response a status code equal to that given
func (o *ContainerVulnerabilitiesBySeverityCountOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the container vulnerabilities by severity count o k response
func (o *ContainerVulnerabilitiesBySeverityCountOK) Code() int {
	return 200
}

func (o *ContainerVulnerabilitiesBySeverityCountOK) Error() string {
	return fmt.Sprintf("[GET /container-security/aggregates/containers/vulnerability-count-by-severity/v1][%d] containerVulnerabilitiesBySeverityCountOK  %+v", 200, o.Payload)
}

func (o *ContainerVulnerabilitiesBySeverityCountOK) String() string {
	return fmt.Sprintf("[GET /container-security/aggregates/containers/vulnerability-count-by-severity/v1][%d] containerVulnerabilitiesBySeverityCountOK  %+v", 200, o.Payload)
}

func (o *ContainerVulnerabilitiesBySeverityCountOK) GetPayload() *models.ModelsAggregateValuesByFieldResponse {
	return o.Payload
}

func (o *ContainerVulnerabilitiesBySeverityCountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.ModelsAggregateValuesByFieldResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewContainerVulnerabilitiesBySeverityCountForbidden creates a ContainerVulnerabilitiesBySeverityCountForbidden with default headers values
func NewContainerVulnerabilitiesBySeverityCountForbidden() *ContainerVulnerabilitiesBySeverityCountForbidden {
	return &ContainerVulnerabilitiesBySeverityCountForbidden{}
}

/*
ContainerVulnerabilitiesBySeverityCountForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ContainerVulnerabilitiesBySeverityCountForbidden struct {

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

// IsSuccess returns true when this container vulnerabilities by severity count forbidden response has a 2xx status code
func (o *ContainerVulnerabilitiesBySeverityCountForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this container vulnerabilities by severity count forbidden response has a 3xx status code
func (o *ContainerVulnerabilitiesBySeverityCountForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this container vulnerabilities by severity count forbidden response has a 4xx status code
func (o *ContainerVulnerabilitiesBySeverityCountForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this container vulnerabilities by severity count forbidden response has a 5xx status code
func (o *ContainerVulnerabilitiesBySeverityCountForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this container vulnerabilities by severity count forbidden response a status code equal to that given
func (o *ContainerVulnerabilitiesBySeverityCountForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the container vulnerabilities by severity count forbidden response
func (o *ContainerVulnerabilitiesBySeverityCountForbidden) Code() int {
	return 403
}

func (o *ContainerVulnerabilitiesBySeverityCountForbidden) Error() string {
	return fmt.Sprintf("[GET /container-security/aggregates/containers/vulnerability-count-by-severity/v1][%d] containerVulnerabilitiesBySeverityCountForbidden  %+v", 403, o.Payload)
}

func (o *ContainerVulnerabilitiesBySeverityCountForbidden) String() string {
	return fmt.Sprintf("[GET /container-security/aggregates/containers/vulnerability-count-by-severity/v1][%d] containerVulnerabilitiesBySeverityCountForbidden  %+v", 403, o.Payload)
}

func (o *ContainerVulnerabilitiesBySeverityCountForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *ContainerVulnerabilitiesBySeverityCountForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewContainerVulnerabilitiesBySeverityCountTooManyRequests creates a ContainerVulnerabilitiesBySeverityCountTooManyRequests with default headers values
func NewContainerVulnerabilitiesBySeverityCountTooManyRequests() *ContainerVulnerabilitiesBySeverityCountTooManyRequests {
	return &ContainerVulnerabilitiesBySeverityCountTooManyRequests{}
}

/*
ContainerVulnerabilitiesBySeverityCountTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type ContainerVulnerabilitiesBySeverityCountTooManyRequests struct {

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

// IsSuccess returns true when this container vulnerabilities by severity count too many requests response has a 2xx status code
func (o *ContainerVulnerabilitiesBySeverityCountTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this container vulnerabilities by severity count too many requests response has a 3xx status code
func (o *ContainerVulnerabilitiesBySeverityCountTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this container vulnerabilities by severity count too many requests response has a 4xx status code
func (o *ContainerVulnerabilitiesBySeverityCountTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this container vulnerabilities by severity count too many requests response has a 5xx status code
func (o *ContainerVulnerabilitiesBySeverityCountTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this container vulnerabilities by severity count too many requests response a status code equal to that given
func (o *ContainerVulnerabilitiesBySeverityCountTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the container vulnerabilities by severity count too many requests response
func (o *ContainerVulnerabilitiesBySeverityCountTooManyRequests) Code() int {
	return 429
}

func (o *ContainerVulnerabilitiesBySeverityCountTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /container-security/aggregates/containers/vulnerability-count-by-severity/v1][%d] containerVulnerabilitiesBySeverityCountTooManyRequests  %+v", 429, o.Payload)
}

func (o *ContainerVulnerabilitiesBySeverityCountTooManyRequests) String() string {
	return fmt.Sprintf("[GET /container-security/aggregates/containers/vulnerability-count-by-severity/v1][%d] containerVulnerabilitiesBySeverityCountTooManyRequests  %+v", 429, o.Payload)
}

func (o *ContainerVulnerabilitiesBySeverityCountTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *ContainerVulnerabilitiesBySeverityCountTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewContainerVulnerabilitiesBySeverityCountInternalServerError creates a ContainerVulnerabilitiesBySeverityCountInternalServerError with default headers values
func NewContainerVulnerabilitiesBySeverityCountInternalServerError() *ContainerVulnerabilitiesBySeverityCountInternalServerError {
	return &ContainerVulnerabilitiesBySeverityCountInternalServerError{}
}

/*
ContainerVulnerabilitiesBySeverityCountInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type ContainerVulnerabilitiesBySeverityCountInternalServerError struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.CoreEntitiesResponse
}

// IsSuccess returns true when this container vulnerabilities by severity count internal server error response has a 2xx status code
func (o *ContainerVulnerabilitiesBySeverityCountInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this container vulnerabilities by severity count internal server error response has a 3xx status code
func (o *ContainerVulnerabilitiesBySeverityCountInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this container vulnerabilities by severity count internal server error response has a 4xx status code
func (o *ContainerVulnerabilitiesBySeverityCountInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this container vulnerabilities by severity count internal server error response has a 5xx status code
func (o *ContainerVulnerabilitiesBySeverityCountInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this container vulnerabilities by severity count internal server error response a status code equal to that given
func (o *ContainerVulnerabilitiesBySeverityCountInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the container vulnerabilities by severity count internal server error response
func (o *ContainerVulnerabilitiesBySeverityCountInternalServerError) Code() int {
	return 500
}

func (o *ContainerVulnerabilitiesBySeverityCountInternalServerError) Error() string {
	return fmt.Sprintf("[GET /container-security/aggregates/containers/vulnerability-count-by-severity/v1][%d] containerVulnerabilitiesBySeverityCountInternalServerError  %+v", 500, o.Payload)
}

func (o *ContainerVulnerabilitiesBySeverityCountInternalServerError) String() string {
	return fmt.Sprintf("[GET /container-security/aggregates/containers/vulnerability-count-by-severity/v1][%d] containerVulnerabilitiesBySeverityCountInternalServerError  %+v", 500, o.Payload)
}

func (o *ContainerVulnerabilitiesBySeverityCountInternalServerError) GetPayload() *models.CoreEntitiesResponse {
	return o.Payload
}

func (o *ContainerVulnerabilitiesBySeverityCountInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.CoreEntitiesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
