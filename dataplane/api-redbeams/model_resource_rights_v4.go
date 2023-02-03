/*
Redbeams API

API for working with databases and database servers

API version: 2.66.0-b41-1-ge11bb40
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package model

import (
	"encoding/json"
)

// checks if the ResourceRightsV4 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResourceRightsV4{}

// ResourceRightsV4 struct for ResourceRightsV4
type ResourceRightsV4 struct {
	ResourceCrn *string  `json:"resourceCrn,omitempty"`
	Rights      []string `json:"rights,omitempty"`
}

// NewResourceRightsV4 instantiates a new ResourceRightsV4 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceRightsV4() *ResourceRightsV4 {
	this := ResourceRightsV4{}
	return &this
}

// NewResourceRightsV4WithDefaults instantiates a new ResourceRightsV4 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceRightsV4WithDefaults() *ResourceRightsV4 {
	this := ResourceRightsV4{}
	return &this
}

// GetResourceCrn returns the ResourceCrn field value if set, zero value otherwise.
func (o *ResourceRightsV4) GetResourceCrn() string {
	if o == nil || isNil(o.ResourceCrn) {
		var ret string
		return ret
	}
	return *o.ResourceCrn
}

// GetResourceCrnOk returns a tuple with the ResourceCrn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceRightsV4) GetResourceCrnOk() (*string, bool) {
	if o == nil || isNil(o.ResourceCrn) {
		return nil, false
	}
	return o.ResourceCrn, true
}

// HasResourceCrn returns a boolean if a field has been set.
func (o *ResourceRightsV4) HasResourceCrn() bool {
	if o != nil && !isNil(o.ResourceCrn) {
		return true
	}

	return false
}

// SetResourceCrn gets a reference to the given string and assigns it to the ResourceCrn field.
func (o *ResourceRightsV4) SetResourceCrn(v string) {
	o.ResourceCrn = &v
}

// GetRights returns the Rights field value if set, zero value otherwise.
func (o *ResourceRightsV4) GetRights() []string {
	if o == nil || isNil(o.Rights) {
		var ret []string
		return ret
	}
	return o.Rights
}

// GetRightsOk returns a tuple with the Rights field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceRightsV4) GetRightsOk() ([]string, bool) {
	if o == nil || isNil(o.Rights) {
		return nil, false
	}
	return o.Rights, true
}

// HasRights returns a boolean if a field has been set.
func (o *ResourceRightsV4) HasRights() bool {
	if o != nil && !isNil(o.Rights) {
		return true
	}

	return false
}

// SetRights gets a reference to the given []string and assigns it to the Rights field.
func (o *ResourceRightsV4) SetRights(v []string) {
	o.Rights = v
}

func (o ResourceRightsV4) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResourceRightsV4) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ResourceCrn) {
		toSerialize["resourceCrn"] = o.ResourceCrn
	}
	if !isNil(o.Rights) {
		toSerialize["rights"] = o.Rights
	}
	return toSerialize, nil
}

type NullableResourceRightsV4 struct {
	value *ResourceRightsV4
	isSet bool
}

func (v NullableResourceRightsV4) Get() *ResourceRightsV4 {
	return v.value
}

func (v *NullableResourceRightsV4) Set(val *ResourceRightsV4) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceRightsV4) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceRightsV4) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceRightsV4(val *ResourceRightsV4) *NullableResourceRightsV4 {
	return &NullableResourceRightsV4{value: val, isSet: true}
}

func (v NullableResourceRightsV4) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceRightsV4) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
