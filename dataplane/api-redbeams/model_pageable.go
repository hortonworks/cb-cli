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

// checks if the Pageable type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Pageable{}

// Pageable struct for Pageable
type Pageable struct {
	Sort       *Sort  `json:"sort,omitempty"`
	Offset     *int64 `json:"offset,omitempty"`
	Paged      *bool  `json:"paged,omitempty"`
	Unpaged    *bool  `json:"unpaged,omitempty"`
	PageNumber *int32 `json:"pageNumber,omitempty"`
	PageSize   *int32 `json:"pageSize,omitempty"`
}

// NewPageable instantiates a new Pageable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPageable() *Pageable {
	this := Pageable{}
	return &this
}

// NewPageableWithDefaults instantiates a new Pageable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPageableWithDefaults() *Pageable {
	this := Pageable{}
	return &this
}

// GetSort returns the Sort field value if set, zero value otherwise.
func (o *Pageable) GetSort() Sort {
	if o == nil || isNil(o.Sort) {
		var ret Sort
		return ret
	}
	return *o.Sort
}

// GetSortOk returns a tuple with the Sort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pageable) GetSortOk() (*Sort, bool) {
	if o == nil || isNil(o.Sort) {
		return nil, false
	}
	return o.Sort, true
}

// HasSort returns a boolean if a field has been set.
func (o *Pageable) HasSort() bool {
	if o != nil && !isNil(o.Sort) {
		return true
	}

	return false
}

// SetSort gets a reference to the given Sort and assigns it to the Sort field.
func (o *Pageable) SetSort(v Sort) {
	o.Sort = &v
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *Pageable) GetOffset() int64 {
	if o == nil || isNil(o.Offset) {
		var ret int64
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pageable) GetOffsetOk() (*int64, bool) {
	if o == nil || isNil(o.Offset) {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *Pageable) HasOffset() bool {
	if o != nil && !isNil(o.Offset) {
		return true
	}

	return false
}

// SetOffset gets a reference to the given int64 and assigns it to the Offset field.
func (o *Pageable) SetOffset(v int64) {
	o.Offset = &v
}

// GetPaged returns the Paged field value if set, zero value otherwise.
func (o *Pageable) GetPaged() bool {
	if o == nil || isNil(o.Paged) {
		var ret bool
		return ret
	}
	return *o.Paged
}

// GetPagedOk returns a tuple with the Paged field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pageable) GetPagedOk() (*bool, bool) {
	if o == nil || isNil(o.Paged) {
		return nil, false
	}
	return o.Paged, true
}

// HasPaged returns a boolean if a field has been set.
func (o *Pageable) HasPaged() bool {
	if o != nil && !isNil(o.Paged) {
		return true
	}

	return false
}

// SetPaged gets a reference to the given bool and assigns it to the Paged field.
func (o *Pageable) SetPaged(v bool) {
	o.Paged = &v
}

// GetUnpaged returns the Unpaged field value if set, zero value otherwise.
func (o *Pageable) GetUnpaged() bool {
	if o == nil || isNil(o.Unpaged) {
		var ret bool
		return ret
	}
	return *o.Unpaged
}

// GetUnpagedOk returns a tuple with the Unpaged field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pageable) GetUnpagedOk() (*bool, bool) {
	if o == nil || isNil(o.Unpaged) {
		return nil, false
	}
	return o.Unpaged, true
}

// HasUnpaged returns a boolean if a field has been set.
func (o *Pageable) HasUnpaged() bool {
	if o != nil && !isNil(o.Unpaged) {
		return true
	}

	return false
}

// SetUnpaged gets a reference to the given bool and assigns it to the Unpaged field.
func (o *Pageable) SetUnpaged(v bool) {
	o.Unpaged = &v
}

// GetPageNumber returns the PageNumber field value if set, zero value otherwise.
func (o *Pageable) GetPageNumber() int32 {
	if o == nil || isNil(o.PageNumber) {
		var ret int32
		return ret
	}
	return *o.PageNumber
}

// GetPageNumberOk returns a tuple with the PageNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pageable) GetPageNumberOk() (*int32, bool) {
	if o == nil || isNil(o.PageNumber) {
		return nil, false
	}
	return o.PageNumber, true
}

// HasPageNumber returns a boolean if a field has been set.
func (o *Pageable) HasPageNumber() bool {
	if o != nil && !isNil(o.PageNumber) {
		return true
	}

	return false
}

// SetPageNumber gets a reference to the given int32 and assigns it to the PageNumber field.
func (o *Pageable) SetPageNumber(v int32) {
	o.PageNumber = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *Pageable) GetPageSize() int32 {
	if o == nil || isNil(o.PageSize) {
		var ret int32
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pageable) GetPageSizeOk() (*int32, bool) {
	if o == nil || isNil(o.PageSize) {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *Pageable) HasPageSize() bool {
	if o != nil && !isNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int32 and assigns it to the PageSize field.
func (o *Pageable) SetPageSize(v int32) {
	o.PageSize = &v
}

func (o Pageable) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Pageable) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Sort) {
		toSerialize["sort"] = o.Sort
	}
	if !isNil(o.Offset) {
		toSerialize["offset"] = o.Offset
	}
	if !isNil(o.Paged) {
		toSerialize["paged"] = o.Paged
	}
	if !isNil(o.Unpaged) {
		toSerialize["unpaged"] = o.Unpaged
	}
	if !isNil(o.PageNumber) {
		toSerialize["pageNumber"] = o.PageNumber
	}
	if !isNil(o.PageSize) {
		toSerialize["pageSize"] = o.PageSize
	}
	return toSerialize, nil
}

type NullablePageable struct {
	value *Pageable
	isSet bool
}

func (v NullablePageable) Get() *Pageable {
	return v.value
}

func (v *NullablePageable) Set(val *Pageable) {
	v.value = val
	v.isSet = true
}

func (v NullablePageable) IsSet() bool {
	return v.isSet
}

func (v *NullablePageable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePageable(val *Pageable) *NullablePageable {
	return &NullablePageable{value: val, isSet: true}
}

func (v NullablePageable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePageable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
