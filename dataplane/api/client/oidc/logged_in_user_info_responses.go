// Code generated by go-swagger; DO NOT EDIT.

package oidc

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/hortonworks/cb-cli/dataplane/api/model"
)

// LoggedInUserInfoReader is a Reader for the LoggedInUserInfo structure.
type LoggedInUserInfoReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LoggedInUserInfoReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewLoggedInUserInfoOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewLoggedInUserInfoBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewLoggedInUserInfoUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewLoggedInUserInfoForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewLoggedInUserInfoInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewLoggedInUserInfoOK creates a LoggedInUserInfoOK with default headers values
func NewLoggedInUserInfoOK() *LoggedInUserInfoOK {
	return &LoggedInUserInfoOK{}
}

/*LoggedInUserInfoOK handles this case with default header values.

successful operation
*/
type LoggedInUserInfoOK struct {
	Payload *model.UserInfo
}

func (o *LoggedInUserInfoOK) Error() string {
	return fmt.Sprintf("[GET /oidc/userinfo][%d] loggedInUserInfoOK  %+v", 200, o.Payload)
}

func (o *LoggedInUserInfoOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.UserInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLoggedInUserInfoBadRequest creates a LoggedInUserInfoBadRequest with default headers values
func NewLoggedInUserInfoBadRequest() *LoggedInUserInfoBadRequest {
	return &LoggedInUserInfoBadRequest{}
}

/*LoggedInUserInfoBadRequest handles this case with default header values.

Bad Request
*/
type LoggedInUserInfoBadRequest struct {
}

func (o *LoggedInUserInfoBadRequest) Error() string {
	return fmt.Sprintf("[GET /oidc/userinfo][%d] loggedInUserInfoBadRequest ", 400)
}

func (o *LoggedInUserInfoBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewLoggedInUserInfoUnauthorized creates a LoggedInUserInfoUnauthorized with default headers values
func NewLoggedInUserInfoUnauthorized() *LoggedInUserInfoUnauthorized {
	return &LoggedInUserInfoUnauthorized{}
}

/*LoggedInUserInfoUnauthorized handles this case with default header values.

Unauthorized
*/
type LoggedInUserInfoUnauthorized struct {
}

func (o *LoggedInUserInfoUnauthorized) Error() string {
	return fmt.Sprintf("[GET /oidc/userinfo][%d] loggedInUserInfoUnauthorized ", 401)
}

func (o *LoggedInUserInfoUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewLoggedInUserInfoForbidden creates a LoggedInUserInfoForbidden with default headers values
func NewLoggedInUserInfoForbidden() *LoggedInUserInfoForbidden {
	return &LoggedInUserInfoForbidden{}
}

/*LoggedInUserInfoForbidden handles this case with default header values.

Forbidden
*/
type LoggedInUserInfoForbidden struct {
}

func (o *LoggedInUserInfoForbidden) Error() string {
	return fmt.Sprintf("[GET /oidc/userinfo][%d] loggedInUserInfoForbidden ", 403)
}

func (o *LoggedInUserInfoForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewLoggedInUserInfoInternalServerError creates a LoggedInUserInfoInternalServerError with default headers values
func NewLoggedInUserInfoInternalServerError() *LoggedInUserInfoInternalServerError {
	return &LoggedInUserInfoInternalServerError{}
}

/*LoggedInUserInfoInternalServerError handles this case with default header values.

Internal server error
*/
type LoggedInUserInfoInternalServerError struct {
}

func (o *LoggedInUserInfoInternalServerError) Error() string {
	return fmt.Sprintf("[GET /oidc/userinfo][%d] loggedInUserInfoInternalServerError ", 500)
}

func (o *LoggedInUserInfoInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
