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

// BackupDatabaseInternalReader is a Reader for the BackupDatabaseInternal structure.
type BackupDatabaseInternalReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BackupDatabaseInternalReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewBackupDatabaseInternalOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewBackupDatabaseInternalOK creates a BackupDatabaseInternalOK with default headers values
func NewBackupDatabaseInternalOK() *BackupDatabaseInternalOK {
	return &BackupDatabaseInternalOK{}
}

/*
BackupDatabaseInternalOK handles this case with default header values.

successful operation
*/
type BackupDatabaseInternalOK struct {
	Payload *model.SdxDatabaseBackupResponse
}

func (o *BackupDatabaseInternalOK) Error() string {
	return fmt.Sprintf("[POST /sdx/{name}/backupDatabaseInternal][%d] backupDatabaseInternalOK  %+v", 200, o.Payload)
}

func (o *BackupDatabaseInternalOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.SdxDatabaseBackupResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
