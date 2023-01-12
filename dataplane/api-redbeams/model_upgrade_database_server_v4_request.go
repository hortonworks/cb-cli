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

// checks if the UpgradeDatabaseServerV4Request type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpgradeDatabaseServerV4Request{}

// UpgradeDatabaseServerV4Request Request for upgrading a database server in a provider to a higher major version
type UpgradeDatabaseServerV4Request struct {
	// The major version to which the database server should be upgraded
	UpgradeTargetMajorVersion string `json:"upgradeTargetMajorVersion"`
}

// NewUpgradeDatabaseServerV4Request instantiates a new UpgradeDatabaseServerV4Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpgradeDatabaseServerV4Request(upgradeTargetMajorVersion string) *UpgradeDatabaseServerV4Request {
	this := UpgradeDatabaseServerV4Request{}
	this.UpgradeTargetMajorVersion = upgradeTargetMajorVersion
	return &this
}

// NewUpgradeDatabaseServerV4RequestWithDefaults instantiates a new UpgradeDatabaseServerV4Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpgradeDatabaseServerV4RequestWithDefaults() *UpgradeDatabaseServerV4Request {
	this := UpgradeDatabaseServerV4Request{}
	return &this
}

// GetUpgradeTargetMajorVersion returns the UpgradeTargetMajorVersion field value
func (o *UpgradeDatabaseServerV4Request) GetUpgradeTargetMajorVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpgradeTargetMajorVersion
}

// GetUpgradeTargetMajorVersionOk returns a tuple with the UpgradeTargetMajorVersion field value
// and a boolean to check if the value has been set.
func (o *UpgradeDatabaseServerV4Request) GetUpgradeTargetMajorVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpgradeTargetMajorVersion, true
}

// SetUpgradeTargetMajorVersion sets field value
func (o *UpgradeDatabaseServerV4Request) SetUpgradeTargetMajorVersion(v string) {
	o.UpgradeTargetMajorVersion = v
}

func (o UpgradeDatabaseServerV4Request) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpgradeDatabaseServerV4Request) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["upgradeTargetMajorVersion"] = o.UpgradeTargetMajorVersion
	return toSerialize, nil
}

type NullableUpgradeDatabaseServerV4Request struct {
	value *UpgradeDatabaseServerV4Request
	isSet bool
}

func (v NullableUpgradeDatabaseServerV4Request) Get() *UpgradeDatabaseServerV4Request {
	return v.value
}

func (v *NullableUpgradeDatabaseServerV4Request) Set(val *UpgradeDatabaseServerV4Request) {
	v.value = val
	v.isSet = true
}

func (v NullableUpgradeDatabaseServerV4Request) IsSet() bool {
	return v.isSet
}

func (v *NullableUpgradeDatabaseServerV4Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpgradeDatabaseServerV4Request(val *UpgradeDatabaseServerV4Request) *NullableUpgradeDatabaseServerV4Request {
	return &NullableUpgradeDatabaseServerV4Request{value: val, isSet: true}
}

func (v NullableUpgradeDatabaseServerV4Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpgradeDatabaseServerV4Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
