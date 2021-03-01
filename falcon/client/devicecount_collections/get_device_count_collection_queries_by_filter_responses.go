// Code generated by go-swagger; DO NOT EDIT.

package devicecount_collections

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

// GetDeviceCountCollectionQueriesByFilterReader is a Reader for the GetDeviceCountCollectionQueriesByFilter structure.
type GetDeviceCountCollectionQueriesByFilterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDeviceCountCollectionQueriesByFilterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDeviceCountCollectionQueriesByFilterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetDeviceCountCollectionQueriesByFilterForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetDeviceCountCollectionQueriesByFilterTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetDeviceCountCollectionQueriesByFilterDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetDeviceCountCollectionQueriesByFilterOK creates a GetDeviceCountCollectionQueriesByFilterOK with default headers values
func NewGetDeviceCountCollectionQueriesByFilterOK() *GetDeviceCountCollectionQueriesByFilterOK {
	return &GetDeviceCountCollectionQueriesByFilterOK{}
}

/* GetDeviceCountCollectionQueriesByFilterOK describes a response with status code 200, with default header values.

OK
*/
type GetDeviceCountCollectionQueriesByFilterOK struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaQueryResponse
}

func (o *GetDeviceCountCollectionQueriesByFilterOK) Error() string {
	return fmt.Sprintf("[GET /falcon-complete-dashboards/queries/devicecount-collections/v1][%d] getDeviceCountCollectionQueriesByFilterOK  %+v", 200, o.Payload)
}
func (o *GetDeviceCountCollectionQueriesByFilterOK) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *GetDeviceCountCollectionQueriesByFilterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDeviceCountCollectionQueriesByFilterForbidden creates a GetDeviceCountCollectionQueriesByFilterForbidden with default headers values
func NewGetDeviceCountCollectionQueriesByFilterForbidden() *GetDeviceCountCollectionQueriesByFilterForbidden {
	return &GetDeviceCountCollectionQueriesByFilterForbidden{}
}

/* GetDeviceCountCollectionQueriesByFilterForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetDeviceCountCollectionQueriesByFilterForbidden struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaReplyMetaOnly
}

func (o *GetDeviceCountCollectionQueriesByFilterForbidden) Error() string {
	return fmt.Sprintf("[GET /falcon-complete-dashboards/queries/devicecount-collections/v1][%d] getDeviceCountCollectionQueriesByFilterForbidden  %+v", 403, o.Payload)
}
func (o *GetDeviceCountCollectionQueriesByFilterForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetDeviceCountCollectionQueriesByFilterForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetDeviceCountCollectionQueriesByFilterTooManyRequests creates a GetDeviceCountCollectionQueriesByFilterTooManyRequests with default headers values
func NewGetDeviceCountCollectionQueriesByFilterTooManyRequests() *GetDeviceCountCollectionQueriesByFilterTooManyRequests {
	return &GetDeviceCountCollectionQueriesByFilterTooManyRequests{}
}

/* GetDeviceCountCollectionQueriesByFilterTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetDeviceCountCollectionQueriesByFilterTooManyRequests struct {

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

func (o *GetDeviceCountCollectionQueriesByFilterTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /falcon-complete-dashboards/queries/devicecount-collections/v1][%d] getDeviceCountCollectionQueriesByFilterTooManyRequests  %+v", 429, o.Payload)
}
func (o *GetDeviceCountCollectionQueriesByFilterTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetDeviceCountCollectionQueriesByFilterTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetDeviceCountCollectionQueriesByFilterDefault creates a GetDeviceCountCollectionQueriesByFilterDefault with default headers values
func NewGetDeviceCountCollectionQueriesByFilterDefault(code int) *GetDeviceCountCollectionQueriesByFilterDefault {
	return &GetDeviceCountCollectionQueriesByFilterDefault{
		_statusCode: code,
	}
}

/* GetDeviceCountCollectionQueriesByFilterDefault describes a response with status code -1, with default header values.

OK
*/
type GetDeviceCountCollectionQueriesByFilterDefault struct {
	_statusCode int

	Payload *models.MsaQueryResponse
}

// Code gets the status code for the get device count collection queries by filter default response
func (o *GetDeviceCountCollectionQueriesByFilterDefault) Code() int {
	return o._statusCode
}

func (o *GetDeviceCountCollectionQueriesByFilterDefault) Error() string {
	return fmt.Sprintf("[GET /falcon-complete-dashboards/queries/devicecount-collections/v1][%d] GetDeviceCountCollectionQueriesByFilter default  %+v", o._statusCode, o.Payload)
}
func (o *GetDeviceCountCollectionQueriesByFilterDefault) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *GetDeviceCountCollectionQueriesByFilterDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MsaQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}