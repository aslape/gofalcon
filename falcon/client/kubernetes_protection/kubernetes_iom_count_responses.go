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

// KubernetesIomCountReader is a Reader for the KubernetesIomCount structure.
type KubernetesIomCountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KubernetesIomCountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewKubernetesIomCountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewKubernetesIomCountForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewKubernetesIomCountTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewKubernetesIomCountInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /container-security/aggregates/kubernetes-ioms/count/v1] KubernetesIomCount", response, response.Code())
	}
}

// NewKubernetesIomCountOK creates a KubernetesIomCountOK with default headers values
func NewKubernetesIomCountOK() *KubernetesIomCountOK {
	return &KubernetesIomCountOK{}
}

/*
KubernetesIomCountOK describes a response with status code 200, with default header values.

OK
*/
type KubernetesIomCountOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.K8siomsKubernetesIOMCountValue
}

// IsSuccess returns true when this kubernetes iom count o k response has a 2xx status code
func (o *KubernetesIomCountOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this kubernetes iom count o k response has a 3xx status code
func (o *KubernetesIomCountOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes iom count o k response has a 4xx status code
func (o *KubernetesIomCountOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this kubernetes iom count o k response has a 5xx status code
func (o *KubernetesIomCountOK) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes iom count o k response a status code equal to that given
func (o *KubernetesIomCountOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the kubernetes iom count o k response
func (o *KubernetesIomCountOK) Code() int {
	return 200
}

func (o *KubernetesIomCountOK) Error() string {
	return fmt.Sprintf("[GET /container-security/aggregates/kubernetes-ioms/count/v1][%d] kubernetesIomCountOK  %+v", 200, o.Payload)
}

func (o *KubernetesIomCountOK) String() string {
	return fmt.Sprintf("[GET /container-security/aggregates/kubernetes-ioms/count/v1][%d] kubernetesIomCountOK  %+v", 200, o.Payload)
}

func (o *KubernetesIomCountOK) GetPayload() *models.K8siomsKubernetesIOMCountValue {
	return o.Payload
}

func (o *KubernetesIomCountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.K8siomsKubernetesIOMCountValue)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesIomCountForbidden creates a KubernetesIomCountForbidden with default headers values
func NewKubernetesIomCountForbidden() *KubernetesIomCountForbidden {
	return &KubernetesIomCountForbidden{}
}

/*
KubernetesIomCountForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type KubernetesIomCountForbidden struct {

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

// IsSuccess returns true when this kubernetes iom count forbidden response has a 2xx status code
func (o *KubernetesIomCountForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes iom count forbidden response has a 3xx status code
func (o *KubernetesIomCountForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes iom count forbidden response has a 4xx status code
func (o *KubernetesIomCountForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes iom count forbidden response has a 5xx status code
func (o *KubernetesIomCountForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes iom count forbidden response a status code equal to that given
func (o *KubernetesIomCountForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the kubernetes iom count forbidden response
func (o *KubernetesIomCountForbidden) Code() int {
	return 403
}

func (o *KubernetesIomCountForbidden) Error() string {
	return fmt.Sprintf("[GET /container-security/aggregates/kubernetes-ioms/count/v1][%d] kubernetesIomCountForbidden  %+v", 403, o.Payload)
}

func (o *KubernetesIomCountForbidden) String() string {
	return fmt.Sprintf("[GET /container-security/aggregates/kubernetes-ioms/count/v1][%d] kubernetesIomCountForbidden  %+v", 403, o.Payload)
}

func (o *KubernetesIomCountForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *KubernetesIomCountForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewKubernetesIomCountTooManyRequests creates a KubernetesIomCountTooManyRequests with default headers values
func NewKubernetesIomCountTooManyRequests() *KubernetesIomCountTooManyRequests {
	return &KubernetesIomCountTooManyRequests{}
}

/*
KubernetesIomCountTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type KubernetesIomCountTooManyRequests struct {

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

// IsSuccess returns true when this kubernetes iom count too many requests response has a 2xx status code
func (o *KubernetesIomCountTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes iom count too many requests response has a 3xx status code
func (o *KubernetesIomCountTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes iom count too many requests response has a 4xx status code
func (o *KubernetesIomCountTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes iom count too many requests response has a 5xx status code
func (o *KubernetesIomCountTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes iom count too many requests response a status code equal to that given
func (o *KubernetesIomCountTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the kubernetes iom count too many requests response
func (o *KubernetesIomCountTooManyRequests) Code() int {
	return 429
}

func (o *KubernetesIomCountTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /container-security/aggregates/kubernetes-ioms/count/v1][%d] kubernetesIomCountTooManyRequests  %+v", 429, o.Payload)
}

func (o *KubernetesIomCountTooManyRequests) String() string {
	return fmt.Sprintf("[GET /container-security/aggregates/kubernetes-ioms/count/v1][%d] kubernetesIomCountTooManyRequests  %+v", 429, o.Payload)
}

func (o *KubernetesIomCountTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *KubernetesIomCountTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewKubernetesIomCountInternalServerError creates a KubernetesIomCountInternalServerError with default headers values
func NewKubernetesIomCountInternalServerError() *KubernetesIomCountInternalServerError {
	return &KubernetesIomCountInternalServerError{}
}

/*
KubernetesIomCountInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type KubernetesIomCountInternalServerError struct {

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

// IsSuccess returns true when this kubernetes iom count internal server error response has a 2xx status code
func (o *KubernetesIomCountInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes iom count internal server error response has a 3xx status code
func (o *KubernetesIomCountInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes iom count internal server error response has a 4xx status code
func (o *KubernetesIomCountInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this kubernetes iom count internal server error response has a 5xx status code
func (o *KubernetesIomCountInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this kubernetes iom count internal server error response a status code equal to that given
func (o *KubernetesIomCountInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the kubernetes iom count internal server error response
func (o *KubernetesIomCountInternalServerError) Code() int {
	return 500
}

func (o *KubernetesIomCountInternalServerError) Error() string {
	return fmt.Sprintf("[GET /container-security/aggregates/kubernetes-ioms/count/v1][%d] kubernetesIomCountInternalServerError  %+v", 500, o.Payload)
}

func (o *KubernetesIomCountInternalServerError) String() string {
	return fmt.Sprintf("[GET /container-security/aggregates/kubernetes-ioms/count/v1][%d] kubernetesIomCountInternalServerError  %+v", 500, o.Payload)
}

func (o *KubernetesIomCountInternalServerError) GetPayload() *models.CoreEntitiesResponse {
	return o.Payload
}

func (o *KubernetesIomCountInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
