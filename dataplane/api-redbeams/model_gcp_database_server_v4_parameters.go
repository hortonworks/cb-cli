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

// checks if the GcpDatabaseServerV4Parameters type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GcpDatabaseServerV4Parameters{}

// GcpDatabaseServerV4Parameters GCP-specific parameters for the database server
type GcpDatabaseServerV4Parameters struct {
	// Time to retain backups, in days
	BackupRetentionDays *int32 `json:"backupRetentionDays,omitempty"`
	// The version of the database software to use
	EngineVersion *string `json:"engineVersion,omitempty"`
}

// NewGcpDatabaseServerV4Parameters instantiates a new GcpDatabaseServerV4Parameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGcpDatabaseServerV4Parameters() *GcpDatabaseServerV4Parameters {
	this := GcpDatabaseServerV4Parameters{}
	return &this
}

// NewGcpDatabaseServerV4ParametersWithDefaults instantiates a new GcpDatabaseServerV4Parameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGcpDatabaseServerV4ParametersWithDefaults() *GcpDatabaseServerV4Parameters {
	this := GcpDatabaseServerV4Parameters{}
	return &this
}

// GetBackupRetentionDays returns the BackupRetentionDays field value if set, zero value otherwise.
func (o *GcpDatabaseServerV4Parameters) GetBackupRetentionDays() int32 {
	if o == nil || isNil(o.BackupRetentionDays) {
		var ret int32
		return ret
	}
	return *o.BackupRetentionDays
}

// GetBackupRetentionDaysOk returns a tuple with the BackupRetentionDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpDatabaseServerV4Parameters) GetBackupRetentionDaysOk() (*int32, bool) {
	if o == nil || isNil(o.BackupRetentionDays) {
		return nil, false
	}
	return o.BackupRetentionDays, true
}

// HasBackupRetentionDays returns a boolean if a field has been set.
func (o *GcpDatabaseServerV4Parameters) HasBackupRetentionDays() bool {
	if o != nil && !isNil(o.BackupRetentionDays) {
		return true
	}

	return false
}

// SetBackupRetentionDays gets a reference to the given int32 and assigns it to the BackupRetentionDays field.
func (o *GcpDatabaseServerV4Parameters) SetBackupRetentionDays(v int32) {
	o.BackupRetentionDays = &v
}

// GetEngineVersion returns the EngineVersion field value if set, zero value otherwise.
func (o *GcpDatabaseServerV4Parameters) GetEngineVersion() string {
	if o == nil || isNil(o.EngineVersion) {
		var ret string
		return ret
	}
	return *o.EngineVersion
}

// GetEngineVersionOk returns a tuple with the EngineVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpDatabaseServerV4Parameters) GetEngineVersionOk() (*string, bool) {
	if o == nil || isNil(o.EngineVersion) {
		return nil, false
	}
	return o.EngineVersion, true
}

// HasEngineVersion returns a boolean if a field has been set.
func (o *GcpDatabaseServerV4Parameters) HasEngineVersion() bool {
	if o != nil && !isNil(o.EngineVersion) {
		return true
	}

	return false
}

// SetEngineVersion gets a reference to the given string and assigns it to the EngineVersion field.
func (o *GcpDatabaseServerV4Parameters) SetEngineVersion(v string) {
	o.EngineVersion = &v
}

func (o GcpDatabaseServerV4Parameters) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GcpDatabaseServerV4Parameters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.BackupRetentionDays) {
		toSerialize["backupRetentionDays"] = o.BackupRetentionDays
	}
	if !isNil(o.EngineVersion) {
		toSerialize["engineVersion"] = o.EngineVersion
	}
	return toSerialize, nil
}

type NullableGcpDatabaseServerV4Parameters struct {
	value *GcpDatabaseServerV4Parameters
	isSet bool
}

func (v NullableGcpDatabaseServerV4Parameters) Get() *GcpDatabaseServerV4Parameters {
	return v.value
}

func (v *NullableGcpDatabaseServerV4Parameters) Set(val *GcpDatabaseServerV4Parameters) {
	v.value = val
	v.isSet = true
}

func (v NullableGcpDatabaseServerV4Parameters) IsSet() bool {
	return v.isSet
}

func (v *NullableGcpDatabaseServerV4Parameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGcpDatabaseServerV4Parameters(val *GcpDatabaseServerV4Parameters) *NullableGcpDatabaseServerV4Parameters {
	return &NullableGcpDatabaseServerV4Parameters{value: val, isSet: true}
}

func (v NullableGcpDatabaseServerV4Parameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGcpDatabaseServerV4Parameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
