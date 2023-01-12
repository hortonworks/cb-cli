# GcpDatabaseServerV4Parameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupRetentionDays** | Pointer to **int32** | Time to retain backups, in days | [optional] 
**EngineVersion** | Pointer to **string** | The version of the database software to use | [optional] 

## Methods

### NewGcpDatabaseServerV4Parameters

`func NewGcpDatabaseServerV4Parameters() *GcpDatabaseServerV4Parameters`

NewGcpDatabaseServerV4Parameters instantiates a new GcpDatabaseServerV4Parameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGcpDatabaseServerV4ParametersWithDefaults

`func NewGcpDatabaseServerV4ParametersWithDefaults() *GcpDatabaseServerV4Parameters`

NewGcpDatabaseServerV4ParametersWithDefaults instantiates a new GcpDatabaseServerV4Parameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupRetentionDays

`func (o *GcpDatabaseServerV4Parameters) GetBackupRetentionDays() int32`

GetBackupRetentionDays returns the BackupRetentionDays field if non-nil, zero value otherwise.

### GetBackupRetentionDaysOk

`func (o *GcpDatabaseServerV4Parameters) GetBackupRetentionDaysOk() (*int32, bool)`

GetBackupRetentionDaysOk returns a tuple with the BackupRetentionDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupRetentionDays

`func (o *GcpDatabaseServerV4Parameters) SetBackupRetentionDays(v int32)`

SetBackupRetentionDays sets BackupRetentionDays field to given value.

### HasBackupRetentionDays

`func (o *GcpDatabaseServerV4Parameters) HasBackupRetentionDays() bool`

HasBackupRetentionDays returns a boolean if a field has been set.

### GetEngineVersion

`func (o *GcpDatabaseServerV4Parameters) GetEngineVersion() string`

GetEngineVersion returns the EngineVersion field if non-nil, zero value otherwise.

### GetEngineVersionOk

`func (o *GcpDatabaseServerV4Parameters) GetEngineVersionOk() (*string, bool)`

GetEngineVersionOk returns a tuple with the EngineVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineVersion

`func (o *GcpDatabaseServerV4Parameters) SetEngineVersion(v string)`

SetEngineVersion sets EngineVersion field to given value.

### HasEngineVersion

`func (o *GcpDatabaseServerV4Parameters) HasEngineVersion() bool`

HasEngineVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


