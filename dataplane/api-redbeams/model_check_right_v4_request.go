/*
Redbeams API

API for working with databases and database servers

API version: 2.66.0-b41-1-ge11bb40
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the CheckRightV4Request type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CheckRightV4Request{}

// CheckRightV4Request struct for CheckRightV4Request
type CheckRightV4Request struct {
	Rights []string `json:"rights,omitempty"`
}

// NewCheckRightV4Request instantiates a new CheckRightV4Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckRightV4Request() *CheckRightV4Request {
	this := CheckRightV4Request{}
	return &this
}

// NewCheckRightV4RequestWithDefaults instantiates a new CheckRightV4Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckRightV4RequestWithDefaults() *CheckRightV4Request {
	this := CheckRightV4Request{}
	return &this
}

// GetRights returns the Rights field value if set, zero value otherwise.
func (o *CheckRightV4Request) GetRights() []string {
	if o == nil || isNil(o.Rights) {
		var ret []string
		return ret
	}
	return o.Rights
}

// GetRightsOk returns a tuple with the Rights field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckRightV4Request) GetRightsOk() ([]string, bool) {
	if o == nil || isNil(o.Rights) {
		return nil, false
	}
	return o.Rights, true
}

// HasRights returns a boolean if a field has been set.
func (o *CheckRightV4Request) HasRights() bool {
	if o != nil && !isNil(o.Rights) {
		return true
	}

	return false
}

// SetRights gets a reference to the given []string and assigns it to the Rights field.
func (o *CheckRightV4Request) SetRights(v []string) {
	o.Rights = v
}

func (o CheckRightV4Request) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CheckRightV4Request) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Rights) {
		toSerialize["rights"] = o.Rights
	}
	return toSerialize, nil
}

type NullableCheckRightV4Request struct {
	value *CheckRightV4Request
	isSet bool
}

func (v NullableCheckRightV4Request) Get() *CheckRightV4Request {
	return v.value
}

func (v *NullableCheckRightV4Request) Set(val *CheckRightV4Request) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckRightV4Request) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckRightV4Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckRightV4Request(val *CheckRightV4Request) *NullableCheckRightV4Request {
	return &NullableCheckRightV4Request{value: val, isSet: true}
}

func (v NullableCheckRightV4Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckRightV4Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
