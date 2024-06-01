// Code generated by go-swagger; DO NOT EDIT.

package cspm_registration

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

// GetCSPMGCPValidateAccountsExtReader is a Reader for the GetCSPMGCPValidateAccountsExt structure.
type GetCSPMGCPValidateAccountsExtReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCSPMGCPValidateAccountsExtReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCSPMGCPValidateAccountsExtOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetCSPMGCPValidateAccountsExtBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetCSPMGCPValidateAccountsExtForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetCSPMGCPValidateAccountsExtNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetCSPMGCPValidateAccountsExtTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetCSPMGCPValidateAccountsExtInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /cloud-connect-cspm-gcp/entities/account/validate/v1] GetCSPMGCPValidateAccountsExt", response, response.Code())
	}
}

// NewGetCSPMGCPValidateAccountsExtOK creates a GetCSPMGCPValidateAccountsExtOK with default headers values
func NewGetCSPMGCPValidateAccountsExtOK() *GetCSPMGCPValidateAccountsExtOK {
	return &GetCSPMGCPValidateAccountsExtOK{}
}

/*
GetCSPMGCPValidateAccountsExtOK describes a response with status code 200, with default header values.

OK
*/
type GetCSPMGCPValidateAccountsExtOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.RegistrationGCPAccountValidationResponseV1
}

// IsSuccess returns true when this get c s p m g c p validate accounts ext o k response has a 2xx status code
func (o *GetCSPMGCPValidateAccountsExtOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get c s p m g c p validate accounts ext o k response has a 3xx status code
func (o *GetCSPMGCPValidateAccountsExtOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get c s p m g c p validate accounts ext o k response has a 4xx status code
func (o *GetCSPMGCPValidateAccountsExtOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get c s p m g c p validate accounts ext o k response has a 5xx status code
func (o *GetCSPMGCPValidateAccountsExtOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get c s p m g c p validate accounts ext o k response a status code equal to that given
func (o *GetCSPMGCPValidateAccountsExtOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get c s p m g c p validate accounts ext o k response
func (o *GetCSPMGCPValidateAccountsExtOK) Code() int {
	return 200
}

func (o *GetCSPMGCPValidateAccountsExtOK) Error() string {
	return fmt.Sprintf("[POST /cloud-connect-cspm-gcp/entities/account/validate/v1][%d] getCSPMGCPValidateAccountsExtOK  %+v", 200, o.Payload)
}

func (o *GetCSPMGCPValidateAccountsExtOK) String() string {
	return fmt.Sprintf("[POST /cloud-connect-cspm-gcp/entities/account/validate/v1][%d] getCSPMGCPValidateAccountsExtOK  %+v", 200, o.Payload)
}

func (o *GetCSPMGCPValidateAccountsExtOK) GetPayload() *models.RegistrationGCPAccountValidationResponseV1 {
	return o.Payload
}

func (o *GetCSPMGCPValidateAccountsExtOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.RegistrationGCPAccountValidationResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCSPMGCPValidateAccountsExtBadRequest creates a GetCSPMGCPValidateAccountsExtBadRequest with default headers values
func NewGetCSPMGCPValidateAccountsExtBadRequest() *GetCSPMGCPValidateAccountsExtBadRequest {
	return &GetCSPMGCPValidateAccountsExtBadRequest{}
}

/*
GetCSPMGCPValidateAccountsExtBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetCSPMGCPValidateAccountsExtBadRequest struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.RegistrationGCPAccountValidationResponseV1
}

// IsSuccess returns true when this get c s p m g c p validate accounts ext bad request response has a 2xx status code
func (o *GetCSPMGCPValidateAccountsExtBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get c s p m g c p validate accounts ext bad request response has a 3xx status code
func (o *GetCSPMGCPValidateAccountsExtBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get c s p m g c p validate accounts ext bad request response has a 4xx status code
func (o *GetCSPMGCPValidateAccountsExtBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get c s p m g c p validate accounts ext bad request response has a 5xx status code
func (o *GetCSPMGCPValidateAccountsExtBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get c s p m g c p validate accounts ext bad request response a status code equal to that given
func (o *GetCSPMGCPValidateAccountsExtBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get c s p m g c p validate accounts ext bad request response
func (o *GetCSPMGCPValidateAccountsExtBadRequest) Code() int {
	return 400
}

func (o *GetCSPMGCPValidateAccountsExtBadRequest) Error() string {
	return fmt.Sprintf("[POST /cloud-connect-cspm-gcp/entities/account/validate/v1][%d] getCSPMGCPValidateAccountsExtBadRequest  %+v", 400, o.Payload)
}

func (o *GetCSPMGCPValidateAccountsExtBadRequest) String() string {
	return fmt.Sprintf("[POST /cloud-connect-cspm-gcp/entities/account/validate/v1][%d] getCSPMGCPValidateAccountsExtBadRequest  %+v", 400, o.Payload)
}

func (o *GetCSPMGCPValidateAccountsExtBadRequest) GetPayload() *models.RegistrationGCPAccountValidationResponseV1 {
	return o.Payload
}

func (o *GetCSPMGCPValidateAccountsExtBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.RegistrationGCPAccountValidationResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCSPMGCPValidateAccountsExtForbidden creates a GetCSPMGCPValidateAccountsExtForbidden with default headers values
func NewGetCSPMGCPValidateAccountsExtForbidden() *GetCSPMGCPValidateAccountsExtForbidden {
	return &GetCSPMGCPValidateAccountsExtForbidden{}
}

/*
GetCSPMGCPValidateAccountsExtForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetCSPMGCPValidateAccountsExtForbidden struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaspecResponseFields
}

// IsSuccess returns true when this get c s p m g c p validate accounts ext forbidden response has a 2xx status code
func (o *GetCSPMGCPValidateAccountsExtForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get c s p m g c p validate accounts ext forbidden response has a 3xx status code
func (o *GetCSPMGCPValidateAccountsExtForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get c s p m g c p validate accounts ext forbidden response has a 4xx status code
func (o *GetCSPMGCPValidateAccountsExtForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get c s p m g c p validate accounts ext forbidden response has a 5xx status code
func (o *GetCSPMGCPValidateAccountsExtForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get c s p m g c p validate accounts ext forbidden response a status code equal to that given
func (o *GetCSPMGCPValidateAccountsExtForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get c s p m g c p validate accounts ext forbidden response
func (o *GetCSPMGCPValidateAccountsExtForbidden) Code() int {
	return 403
}

func (o *GetCSPMGCPValidateAccountsExtForbidden) Error() string {
	return fmt.Sprintf("[POST /cloud-connect-cspm-gcp/entities/account/validate/v1][%d] getCSPMGCPValidateAccountsExtForbidden  %+v", 403, o.Payload)
}

func (o *GetCSPMGCPValidateAccountsExtForbidden) String() string {
	return fmt.Sprintf("[POST /cloud-connect-cspm-gcp/entities/account/validate/v1][%d] getCSPMGCPValidateAccountsExtForbidden  %+v", 403, o.Payload)
}

func (o *GetCSPMGCPValidateAccountsExtForbidden) GetPayload() *models.MsaspecResponseFields {
	return o.Payload
}

func (o *GetCSPMGCPValidateAccountsExtForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaspecResponseFields)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCSPMGCPValidateAccountsExtNotFound creates a GetCSPMGCPValidateAccountsExtNotFound with default headers values
func NewGetCSPMGCPValidateAccountsExtNotFound() *GetCSPMGCPValidateAccountsExtNotFound {
	return &GetCSPMGCPValidateAccountsExtNotFound{}
}

/*
GetCSPMGCPValidateAccountsExtNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetCSPMGCPValidateAccountsExtNotFound struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.RegistrationGCPAccountValidationResponseV1
}

// IsSuccess returns true when this get c s p m g c p validate accounts ext not found response has a 2xx status code
func (o *GetCSPMGCPValidateAccountsExtNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get c s p m g c p validate accounts ext not found response has a 3xx status code
func (o *GetCSPMGCPValidateAccountsExtNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get c s p m g c p validate accounts ext not found response has a 4xx status code
func (o *GetCSPMGCPValidateAccountsExtNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get c s p m g c p validate accounts ext not found response has a 5xx status code
func (o *GetCSPMGCPValidateAccountsExtNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get c s p m g c p validate accounts ext not found response a status code equal to that given
func (o *GetCSPMGCPValidateAccountsExtNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get c s p m g c p validate accounts ext not found response
func (o *GetCSPMGCPValidateAccountsExtNotFound) Code() int {
	return 404
}

func (o *GetCSPMGCPValidateAccountsExtNotFound) Error() string {
	return fmt.Sprintf("[POST /cloud-connect-cspm-gcp/entities/account/validate/v1][%d] getCSPMGCPValidateAccountsExtNotFound  %+v", 404, o.Payload)
}

func (o *GetCSPMGCPValidateAccountsExtNotFound) String() string {
	return fmt.Sprintf("[POST /cloud-connect-cspm-gcp/entities/account/validate/v1][%d] getCSPMGCPValidateAccountsExtNotFound  %+v", 404, o.Payload)
}

func (o *GetCSPMGCPValidateAccountsExtNotFound) GetPayload() *models.RegistrationGCPAccountValidationResponseV1 {
	return o.Payload
}

func (o *GetCSPMGCPValidateAccountsExtNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.RegistrationGCPAccountValidationResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCSPMGCPValidateAccountsExtTooManyRequests creates a GetCSPMGCPValidateAccountsExtTooManyRequests with default headers values
func NewGetCSPMGCPValidateAccountsExtTooManyRequests() *GetCSPMGCPValidateAccountsExtTooManyRequests {
	return &GetCSPMGCPValidateAccountsExtTooManyRequests{}
}

/*
GetCSPMGCPValidateAccountsExtTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetCSPMGCPValidateAccountsExtTooManyRequests struct {

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

// IsSuccess returns true when this get c s p m g c p validate accounts ext too many requests response has a 2xx status code
func (o *GetCSPMGCPValidateAccountsExtTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get c s p m g c p validate accounts ext too many requests response has a 3xx status code
func (o *GetCSPMGCPValidateAccountsExtTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get c s p m g c p validate accounts ext too many requests response has a 4xx status code
func (o *GetCSPMGCPValidateAccountsExtTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get c s p m g c p validate accounts ext too many requests response has a 5xx status code
func (o *GetCSPMGCPValidateAccountsExtTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get c s p m g c p validate accounts ext too many requests response a status code equal to that given
func (o *GetCSPMGCPValidateAccountsExtTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the get c s p m g c p validate accounts ext too many requests response
func (o *GetCSPMGCPValidateAccountsExtTooManyRequests) Code() int {
	return 429
}

func (o *GetCSPMGCPValidateAccountsExtTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /cloud-connect-cspm-gcp/entities/account/validate/v1][%d] getCSPMGCPValidateAccountsExtTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetCSPMGCPValidateAccountsExtTooManyRequests) String() string {
	return fmt.Sprintf("[POST /cloud-connect-cspm-gcp/entities/account/validate/v1][%d] getCSPMGCPValidateAccountsExtTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetCSPMGCPValidateAccountsExtTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetCSPMGCPValidateAccountsExtTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCSPMGCPValidateAccountsExtInternalServerError creates a GetCSPMGCPValidateAccountsExtInternalServerError with default headers values
func NewGetCSPMGCPValidateAccountsExtInternalServerError() *GetCSPMGCPValidateAccountsExtInternalServerError {
	return &GetCSPMGCPValidateAccountsExtInternalServerError{}
}

/*
GetCSPMGCPValidateAccountsExtInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetCSPMGCPValidateAccountsExtInternalServerError struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.RegistrationGCPAccountValidationResponseV1
}

// IsSuccess returns true when this get c s p m g c p validate accounts ext internal server error response has a 2xx status code
func (o *GetCSPMGCPValidateAccountsExtInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get c s p m g c p validate accounts ext internal server error response has a 3xx status code
func (o *GetCSPMGCPValidateAccountsExtInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get c s p m g c p validate accounts ext internal server error response has a 4xx status code
func (o *GetCSPMGCPValidateAccountsExtInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get c s p m g c p validate accounts ext internal server error response has a 5xx status code
func (o *GetCSPMGCPValidateAccountsExtInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get c s p m g c p validate accounts ext internal server error response a status code equal to that given
func (o *GetCSPMGCPValidateAccountsExtInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get c s p m g c p validate accounts ext internal server error response
func (o *GetCSPMGCPValidateAccountsExtInternalServerError) Code() int {
	return 500
}

func (o *GetCSPMGCPValidateAccountsExtInternalServerError) Error() string {
	return fmt.Sprintf("[POST /cloud-connect-cspm-gcp/entities/account/validate/v1][%d] getCSPMGCPValidateAccountsExtInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCSPMGCPValidateAccountsExtInternalServerError) String() string {
	return fmt.Sprintf("[POST /cloud-connect-cspm-gcp/entities/account/validate/v1][%d] getCSPMGCPValidateAccountsExtInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCSPMGCPValidateAccountsExtInternalServerError) GetPayload() *models.RegistrationGCPAccountValidationResponseV1 {
	return o.Payload
}

func (o *GetCSPMGCPValidateAccountsExtInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.RegistrationGCPAccountValidationResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}