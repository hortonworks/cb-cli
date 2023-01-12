# AwsDatabaseServerV4Parameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupRetentionPeriod** | Pointer to **int32** | Time to retain backups, in days | [optional] 
**EngineVersion** | Pointer to **string** | Version of the database engine (vendor) | [optional] 
**MultiAZ** | Pointer to **string** | Whether to use a multi-AZ deployment | [optional] 
**StorageType** | Pointer to **string** | Storage type | [optional] 

## Methods

### NewAwsDatabaseServerV4Parameters

`func NewAwsDatabaseServerV4Parameters() *AwsDatabaseServerV4Parameters`

NewAwsDatabaseServerV4Parameters instantiates a new AwsDatabaseServerV4Parameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsDatabaseServerV4ParametersWithDefaults

`func NewAwsDatabaseServerV4ParametersWithDefaults() *AwsDatabaseServerV4Parameters`

NewAwsDatabaseServerV4ParametersWithDefaults instantiates a new AwsDatabaseServerV4Parameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupRetentionPeriod

`func (o *AwsDatabaseServerV4Parameters) GetBackupRetentionPeriod() int32`

GetBackupRetentionPeriod returns the BackupRetentionPeriod field if non-nil, zero value otherwise.

### GetBackupRetentionPeriodOk

`func (o *AwsDatabaseServerV4Parameters) GetBackupRetentionPeriodOk() (*int32, bool)`

GetBackupRetentionPeriodOk returns a tuple with the BackupRetentionPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupRetentionPeriod

`func (o *AwsDatabaseServerV4Parameters) SetBackupRetentionPeriod(v int32)`

SetBackupRetentionPeriod sets BackupRetentionPeriod field to given value.

### HasBackupRetentionPeriod

`func (o *AwsDatabaseServerV4Parameters) HasBackupRetentionPeriod() bool`

HasBackupRetentionPeriod returns a boolean if a field has been set.

### GetEngineVersion

`func (o *AwsDatabaseServerV4Parameters) GetEngineVersion() string`

GetEngineVersion returns the EngineVersion field if non-nil, zero value otherwise.

### GetEngineVersionOk

`func (o *AwsDatabaseServerV4Parameters) GetEngineVersionOk() (*string, bool)`

GetEngineVersionOk returns a tuple with the EngineVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineVersion

`func (o *AwsDatabaseServerV4Parameters) SetEngineVersion(v string)`

SetEngineVersion sets EngineVersion field to given value.

### HasEngineVersion

`func (o *AwsDatabaseServerV4Parameters) HasEngineVersion() bool`

HasEngineVersion returns a boolean if a field has been set.

### GetMultiAZ

`func (o *AwsDatabaseServerV4Parameters) GetMultiAZ() string`

GetMultiAZ returns the MultiAZ field if non-nil, zero value otherwise.

### GetMultiAZOk

`func (o *AwsDatabaseServerV4Parameters) GetMultiAZOk() (*string, bool)`

GetMultiAZOk returns a tuple with the MultiAZ field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiAZ

`func (o *AwsDatabaseServerV4Parameters) SetMultiAZ(v string)`

SetMultiAZ sets MultiAZ field to given value.

### HasMultiAZ

`func (o *AwsDatabaseServerV4Parameters) HasMultiAZ() bool`

HasMultiAZ returns a boolean if a field has been set.

### GetStorageType

`func (o *AwsDatabaseServerV4Parameters) GetStorageType() string`

GetStorageType returns the StorageType field if non-nil, zero value otherwise.

### GetStorageTypeOk

`func (o *AwsDatabaseServerV4Parameters) GetStorageTypeOk() (*string, bool)`

GetStorageTypeOk returns a tuple with the StorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageType

`func (o *AwsDatabaseServerV4Parameters) SetStorageType(v string)`

SetStorageType sets StorageType field to given value.

### HasStorageType

`func (o *AwsDatabaseServerV4Parameters) HasStorageType() bool`

HasStorageType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


