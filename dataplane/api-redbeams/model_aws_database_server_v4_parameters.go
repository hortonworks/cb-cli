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

// checks if the AwsDatabaseServerV4Parameters type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AwsDatabaseServerV4Parameters{}

// AwsDatabaseServerV4Parameters AWS-specific parameters for the database server
type AwsDatabaseServerV4Parameters struct {
	// Time to retain backups, in days
	BackupRetentionPeriod *int32 `json:"backupRetentionPeriod,omitempty"`
	// Version of the database engine (vendor)
	EngineVersion *string `json:"engineVersion,omitempty"`
	// Whether to use a multi-AZ deployment
	MultiAZ *string `json:"multiAZ,omitempty"`
	// Storage type
	StorageType *string `json:"storageType,omitempty"`
}

// NewAwsDatabaseServerV4Parameters instantiates a new AwsDatabaseServerV4Parameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAwsDatabaseServerV4Parameters() *AwsDatabaseServerV4Parameters {
	this := AwsDatabaseServerV4Parameters{}
	return &this
}

// NewAwsDatabaseServerV4ParametersWithDefaults instantiates a new AwsDatabaseServerV4Parameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAwsDatabaseServerV4ParametersWithDefaults() *AwsDatabaseServerV4Parameters {
	this := AwsDatabaseServerV4Parameters{}
	return &this
}

// GetBackupRetentionPeriod returns the BackupRetentionPeriod field value if set, zero value otherwise.
func (o *AwsDatabaseServerV4Parameters) GetBackupRetentionPeriod() int32 {
	if o == nil || isNil(o.BackupRetentionPeriod) {
		var ret int32
		return ret
	}
	return *o.BackupRetentionPeriod
}

// GetBackupRetentionPeriodOk returns a tuple with the BackupRetentionPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsDatabaseServerV4Parameters) GetBackupRetentionPeriodOk() (*int32, bool) {
	if o == nil || isNil(o.BackupRetentionPeriod) {
		return nil, false
	}
	return o.BackupRetentionPeriod, true
}

// HasBackupRetentionPeriod returns a boolean if a field has been set.
func (o *AwsDatabaseServerV4Parameters) HasBackupRetentionPeriod() bool {
	if o != nil && !isNil(o.BackupRetentionPeriod) {
		return true
	}

	return false
}

// SetBackupRetentionPeriod gets a reference to the given int32 and assigns it to the BackupRetentionPeriod field.
func (o *AwsDatabaseServerV4Parameters) SetBackupRetentionPeriod(v int32) {
	o.BackupRetentionPeriod = &v
}

// GetEngineVersion returns the EngineVersion field value if set, zero value otherwise.
func (o *AwsDatabaseServerV4Parameters) GetEngineVersion() string {
	if o == nil || isNil(o.EngineVersion) {
		var ret string
		return ret
	}
	return *o.EngineVersion
}

// GetEngineVersionOk returns a tuple with the EngineVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsDatabaseServerV4Parameters) GetEngineVersionOk() (*string, bool) {
	if o == nil || isNil(o.EngineVersion) {
		return nil, false
	}
	return o.EngineVersion, true
}

// HasEngineVersion returns a boolean if a field has been set.
func (o *AwsDatabaseServerV4Parameters) HasEngineVersion() bool {
	if o != nil && !isNil(o.EngineVersion) {
		return true
	}

	return false
}

// SetEngineVersion gets a reference to the given string and assigns it to the EngineVersion field.
func (o *AwsDatabaseServerV4Parameters) SetEngineVersion(v string) {
	o.EngineVersion = &v
}

// GetMultiAZ returns the MultiAZ field value if set, zero value otherwise.
func (o *AwsDatabaseServerV4Parameters) GetMultiAZ() string {
	if o == nil || isNil(o.MultiAZ) {
		var ret string
		return ret
	}
	return *o.MultiAZ
}

// GetMultiAZOk returns a tuple with the MultiAZ field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsDatabaseServerV4Parameters) GetMultiAZOk() (*string, bool) {
	if o == nil || isNil(o.MultiAZ) {
		return nil, false
	}
	return o.MultiAZ, true
}

// HasMultiAZ returns a boolean if a field has been set.
func (o *AwsDatabaseServerV4Parameters) HasMultiAZ() bool {
	if o != nil && !isNil(o.MultiAZ) {
		return true
	}

	return false
}

// SetMultiAZ gets a reference to the given string and assigns it to the MultiAZ field.
func (o *AwsDatabaseServerV4Parameters) SetMultiAZ(v string) {
	o.MultiAZ = &v
}

// GetStorageType returns the StorageType field value if set, zero value otherwise.
func (o *AwsDatabaseServerV4Parameters) GetStorageType() string {
	if o == nil || isNil(o.StorageType) {
		var ret string
		return ret
	}
	return *o.StorageType
}

// GetStorageTypeOk returns a tuple with the StorageType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsDatabaseServerV4Parameters) GetStorageTypeOk() (*string, bool) {
	if o == nil || isNil(o.StorageType) {
		return nil, false
	}
	return o.StorageType, true
}

// HasStorageType returns a boolean if a field has been set.
func (o *AwsDatabaseServerV4Parameters) HasStorageType() bool {
	if o != nil && !isNil(o.StorageType) {
		return true
	}

	return false
}

// SetStorageType gets a reference to the given string and assigns it to the StorageType field.
func (o *AwsDatabaseServerV4Parameters) SetStorageType(v string) {
	o.StorageType = &v
}

func (o AwsDatabaseServerV4Parameters) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AwsDatabaseServerV4Parameters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.BackupRetentionPeriod) {
		toSerialize["backupRetentionPeriod"] = o.BackupRetentionPeriod
	}
	if !isNil(o.EngineVersion) {
		toSerialize["engineVersion"] = o.EngineVersion
	}
	if !isNil(o.MultiAZ) {
		toSerialize["multiAZ"] = o.MultiAZ
	}
	if !isNil(o.StorageType) {
		toSerialize["storageType"] = o.StorageType
	}
	return toSerialize, nil
}

type NullableAwsDatabaseServerV4Parameters struct {
	value *AwsDatabaseServerV4Parameters
	isSet bool
}

func (v NullableAwsDatabaseServerV4Parameters) Get() *AwsDatabaseServerV4Parameters {
	return v.value
}

func (v *NullableAwsDatabaseServerV4Parameters) Set(val *AwsDatabaseServerV4Parameters) {
	v.value = val
	v.isSet = true
}

func (v NullableAwsDatabaseServerV4Parameters) IsSet() bool {
	return v.isSet
}

func (v *NullableAwsDatabaseServerV4Parameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAwsDatabaseServerV4Parameters(val *AwsDatabaseServerV4Parameters) *NullableAwsDatabaseServerV4Parameters {
	return &NullableAwsDatabaseServerV4Parameters{value: val, isSet: true}
}

func (v NullableAwsDatabaseServerV4Parameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAwsDatabaseServerV4Parameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
