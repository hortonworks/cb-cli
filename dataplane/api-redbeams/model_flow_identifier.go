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

// checks if the FlowIdentifier type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FlowIdentifier{}

// FlowIdentifier The id of the flow or flow chain that was triggered as part of the process.
type FlowIdentifier struct {
	Type       *string `json:"type,omitempty"`
	PollableId *string `json:"pollableId,omitempty"`
}

// NewFlowIdentifier instantiates a new FlowIdentifier object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFlowIdentifier() *FlowIdentifier {
	this := FlowIdentifier{}
	return &this
}

// NewFlowIdentifierWithDefaults instantiates a new FlowIdentifier object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFlowIdentifierWithDefaults() *FlowIdentifier {
	this := FlowIdentifier{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *FlowIdentifier) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlowIdentifier) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *FlowIdentifier) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *FlowIdentifier) SetType(v string) {
	o.Type = &v
}

// GetPollableId returns the PollableId field value if set, zero value otherwise.
func (o *FlowIdentifier) GetPollableId() string {
	if o == nil || isNil(o.PollableId) {
		var ret string
		return ret
	}
	return *o.PollableId
}

// GetPollableIdOk returns a tuple with the PollableId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlowIdentifier) GetPollableIdOk() (*string, bool) {
	if o == nil || isNil(o.PollableId) {
		return nil, false
	}
	return o.PollableId, true
}

// HasPollableId returns a boolean if a field has been set.
func (o *FlowIdentifier) HasPollableId() bool {
	if o != nil && !isNil(o.PollableId) {
		return true
	}

	return false
}

// SetPollableId gets a reference to the given string and assigns it to the PollableId field.
func (o *FlowIdentifier) SetPollableId(v string) {
	o.PollableId = &v
}

func (o FlowIdentifier) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FlowIdentifier) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.PollableId) {
		toSerialize["pollableId"] = o.PollableId
	}
	return toSerialize, nil
}

type NullableFlowIdentifier struct {
	value *FlowIdentifier
	isSet bool
}

func (v NullableFlowIdentifier) Get() *FlowIdentifier {
	return v.value
}

func (v *NullableFlowIdentifier) Set(val *FlowIdentifier) {
	v.value = val
	v.isSet = true
}

func (v NullableFlowIdentifier) IsSet() bool {
	return v.isSet
}

func (v *NullableFlowIdentifier) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFlowIdentifier(val *FlowIdentifier) *NullableFlowIdentifier {
	return &NullableFlowIdentifier{value: val, isSet: true}
}

func (v NullableFlowIdentifier) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFlowIdentifier) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
