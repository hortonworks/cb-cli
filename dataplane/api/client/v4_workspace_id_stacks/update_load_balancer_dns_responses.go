// Code generated by go-swagger; DO NOT EDIT.

package v4_workspace_id_stacks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// UpdateLoadBalancerDNSReader is a Reader for the UpdateLoadBalancerDNS structure.
type UpdateLoadBalancerDNSReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateLoadBalancerDNSReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {

	result := NewUpdateLoadBalancerDNSDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result

}

// NewUpdateLoadBalancerDNSDefault creates a UpdateLoadBalancerDNSDefault with default headers values
func NewUpdateLoadBalancerDNSDefault(code int) *UpdateLoadBalancerDNSDefault {
	return &UpdateLoadBalancerDNSDefault{
		_statusCode: code,
	}
}

/*
UpdateLoadBalancerDNSDefault handles this case with default header values.

successful operation
*/
type UpdateLoadBalancerDNSDefault struct {
	_statusCode int
}

// Code gets the status code for the update load balancer DNS default response
func (o *UpdateLoadBalancerDNSDefault) Code() int {
	return o._statusCode
}

func (o *UpdateLoadBalancerDNSDefault) Error() string {
	return fmt.Sprintf("[PUT /v4/{workspaceId}/stacks/{name}/update_load_balancer_dns][%d] updateLoadBalancerDNS default ", o._statusCode)
}

func (o *UpdateLoadBalancerDNSDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
