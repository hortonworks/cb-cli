// Code generated by go-swagger; DO NOT EDIT.

package sdx

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api-sdx/model"
)

// BackupDatalakeStatusReader is a Reader for the BackupDatalakeStatus structure.
type BackupDatalakeStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BackupDatalakeStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewBackupDatalakeStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewBackupDatalakeStatusOK creates a BackupDatalakeStatusOK with default headers values
func NewBackupDatalakeStatusOK() *BackupDatalakeStatusOK {
	return &BackupDatalakeStatusOK{}
}

/*
BackupDatalakeStatusOK handles this case with default header values.

successful operation
*/
type BackupDatalakeStatusOK struct {
	Payload *model.SdxBackupStatusResponse
}

func (o *BackupDatalakeStatusOK) Error() string {
	return fmt.Sprintf("[POST /sdx/{name}/backupDatalakeStatus][%d] backupDatalakeStatusOK  %+v", 200, o.Payload)
}

func (o *BackupDatalakeStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.SdxBackupStatusResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
