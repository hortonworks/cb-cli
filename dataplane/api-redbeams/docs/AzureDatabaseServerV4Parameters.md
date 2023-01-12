# AzureDatabaseServerV4Parameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupRetentionDays** | Pointer to **int32** | Time to retain backups, in days | [optional] 
**DbVersion** | Pointer to **string** | The version of the database software to use | [optional] 
**GeoRedundantBackup** | Pointer to **bool** | Whether backups are geographically redundant | [optional] 
**SkuCapacity** | Pointer to **int32** | The number of vCPUs assigned to the database server | [optional] 
**SkuFamily** | Pointer to **string** | The family of hardware used for the database server | [optional] 
**SkuTier** | Pointer to **string** | The tier of SKU for the database server | [optional] 
**StorageAutoGrow** | Pointer to **bool** | Whether the database server will automatically grow storage when necessary | [optional] 

## Methods

### NewAzureDatabaseServerV4Parameters

`func NewAzureDatabaseServerV4Parameters() *AzureDatabaseServerV4Parameters`

NewAzureDatabaseServerV4Parameters instantiates a new AzureDatabaseServerV4Parameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureDatabaseServerV4ParametersWithDefaults

`func NewAzureDatabaseServerV4ParametersWithDefaults() *AzureDatabaseServerV4Parameters`

NewAzureDatabaseServerV4ParametersWithDefaults instantiates a new AzureDatabaseServerV4Parameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupRetentionDays

`func (o *AzureDatabaseServerV4Parameters) GetBackupRetentionDays() int32`

GetBackupRetentionDays returns the BackupRetentionDays field if non-nil, zero value otherwise.

### GetBackupRetentionDaysOk

`func (o *AzureDatabaseServerV4Parameters) GetBackupRetentionDaysOk() (*int32, bool)`

GetBackupRetentionDaysOk returns a tuple with the BackupRetentionDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupRetentionDays

`func (o *AzureDatabaseServerV4Parameters) SetBackupRetentionDays(v int32)`

SetBackupRetentionDays sets BackupRetentionDays field to given value.

### HasBackupRetentionDays

`func (o *AzureDatabaseServerV4Parameters) HasBackupRetentionDays() bool`

HasBackupRetentionDays returns a boolean if a field has been set.

### GetDbVersion

`func (o *AzureDatabaseServerV4Parameters) GetDbVersion() string`

GetDbVersion returns the DbVersion field if non-nil, zero value otherwise.

### GetDbVersionOk

`func (o *AzureDatabaseServerV4Parameters) GetDbVersionOk() (*string, bool)`

GetDbVersionOk returns a tuple with the DbVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbVersion

`func (o *AzureDatabaseServerV4Parameters) SetDbVersion(v string)`

SetDbVersion sets DbVersion field to given value.

### HasDbVersion

`func (o *AzureDatabaseServerV4Parameters) HasDbVersion() bool`

HasDbVersion returns a boolean if a field has been set.

### GetGeoRedundantBackup

`func (o *AzureDatabaseServerV4Parameters) GetGeoRedundantBackup() bool`

GetGeoRedundantBackup returns the GeoRedundantBackup field if non-nil, zero value otherwise.

### GetGeoRedundantBackupOk

`func (o *AzureDatabaseServerV4Parameters) GetGeoRedundantBackupOk() (*bool, bool)`

GetGeoRedundantBackupOk returns a tuple with the GeoRedundantBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeoRedundantBackup

`func (o *AzureDatabaseServerV4Parameters) SetGeoRedundantBackup(v bool)`

SetGeoRedundantBackup sets GeoRedundantBackup field to given value.

### HasGeoRedundantBackup

`func (o *AzureDatabaseServerV4Parameters) HasGeoRedundantBackup() bool`

HasGeoRedundantBackup returns a boolean if a field has been set.

### GetSkuCapacity

`func (o *AzureDatabaseServerV4Parameters) GetSkuCapacity() int32`

GetSkuCapacity returns the SkuCapacity field if non-nil, zero value otherwise.

### GetSkuCapacityOk

`func (o *AzureDatabaseServerV4Parameters) GetSkuCapacityOk() (*int32, bool)`

GetSkuCapacityOk returns a tuple with the SkuCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkuCapacity

`func (o *AzureDatabaseServerV4Parameters) SetSkuCapacity(v int32)`

SetSkuCapacity sets SkuCapacity field to given value.

### HasSkuCapacity

`func (o *AzureDatabaseServerV4Parameters) HasSkuCapacity() bool`

HasSkuCapacity returns a boolean if a field has been set.

### GetSkuFamily

`func (o *AzureDatabaseServerV4Parameters) GetSkuFamily() string`

GetSkuFamily returns the SkuFamily field if non-nil, zero value otherwise.

### GetSkuFamilyOk

`func (o *AzureDatabaseServerV4Parameters) GetSkuFamilyOk() (*string, bool)`

GetSkuFamilyOk returns a tuple with the SkuFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkuFamily

`func (o *AzureDatabaseServerV4Parameters) SetSkuFamily(v string)`

SetSkuFamily sets SkuFamily field to given value.

### HasSkuFamily

`func (o *AzureDatabaseServerV4Parameters) HasSkuFamily() bool`

HasSkuFamily returns a boolean if a field has been set.

### GetSkuTier

`func (o *AzureDatabaseServerV4Parameters) GetSkuTier() string`

GetSkuTier returns the SkuTier field if non-nil, zero value otherwise.

### GetSkuTierOk

`func (o *AzureDatabaseServerV4Parameters) GetSkuTierOk() (*string, bool)`

GetSkuTierOk returns a tuple with the SkuTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkuTier

`func (o *AzureDatabaseServerV4Parameters) SetSkuTier(v string)`

SetSkuTier sets SkuTier field to given value.

### HasSkuTier

`func (o *AzureDatabaseServerV4Parameters) HasSkuTier() bool`

HasSkuTier returns a boolean if a field has been set.

### GetStorageAutoGrow

`func (o *AzureDatabaseServerV4Parameters) GetStorageAutoGrow() bool`

GetStorageAutoGrow returns the StorageAutoGrow field if non-nil, zero value otherwise.

### GetStorageAutoGrowOk

`func (o *AzureDatabaseServerV4Parameters) GetStorageAutoGrowOk() (*bool, bool)`

GetStorageAutoGrowOk returns a tuple with the StorageAutoGrow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageAutoGrow

`func (o *AzureDatabaseServerV4Parameters) SetStorageAutoGrow(v bool)`

SetStorageAutoGrow sets StorageAutoGrow field to given value.

### HasStorageAutoGrow

`func (o *AzureDatabaseServerV4Parameters) HasStorageAutoGrow() bool`

HasStorageAutoGrow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


