# DatabaseServerV4StackRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceType** | Pointer to **string** | Instance type of the database server | [optional] 
**DatabaseVendor** | Pointer to **string** | Name of the database vendor (MYSQL, POSTGRES, ...) | [optional] 
**ConnectionDriver** | Pointer to **string** | Name of the JDBC connection driver (for example: &#39;org.postgresql.Driver&#39;) | [optional] 
**StorageSize** | Pointer to **int64** | Storage size of the database server, in GB | [optional] 
**RootUserName** | Pointer to **string** | Username of the administrative user of the database server | [optional] 
**RootUserPassword** | Pointer to **string** | Password of the administrative user of the database server | [optional] 
**Port** | Pointer to **int32** | Port of the database server | [optional] 
**Aws** | Pointer to [**AwsDatabaseServerV4Parameters**](AwsDatabaseServerV4Parameters.md) |  | [optional] 
**Azure** | Pointer to [**AzureDatabaseServerV4Parameters**](AzureDatabaseServerV4Parameters.md) |  | [optional] 
**Gcp** | Pointer to [**GcpDatabaseServerV4Parameters**](GcpDatabaseServerV4Parameters.md) |  | [optional] 
**SecurityGroup** | Pointer to [**SecurityGroupV4StackRequest**](SecurityGroupV4StackRequest.md) |  | [optional] 

## Methods

### NewDatabaseServerV4StackRequest

`func NewDatabaseServerV4StackRequest() *DatabaseServerV4StackRequest`

NewDatabaseServerV4StackRequest instantiates a new DatabaseServerV4StackRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabaseServerV4StackRequestWithDefaults

`func NewDatabaseServerV4StackRequestWithDefaults() *DatabaseServerV4StackRequest`

NewDatabaseServerV4StackRequestWithDefaults instantiates a new DatabaseServerV4StackRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstanceType

`func (o *DatabaseServerV4StackRequest) GetInstanceType() string`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *DatabaseServerV4StackRequest) GetInstanceTypeOk() (*string, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *DatabaseServerV4StackRequest) SetInstanceType(v string)`

SetInstanceType sets InstanceType field to given value.

### HasInstanceType

`func (o *DatabaseServerV4StackRequest) HasInstanceType() bool`

HasInstanceType returns a boolean if a field has been set.

### GetDatabaseVendor

`func (o *DatabaseServerV4StackRequest) GetDatabaseVendor() string`

GetDatabaseVendor returns the DatabaseVendor field if non-nil, zero value otherwise.

### GetDatabaseVendorOk

`func (o *DatabaseServerV4StackRequest) GetDatabaseVendorOk() (*string, bool)`

GetDatabaseVendorOk returns a tuple with the DatabaseVendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseVendor

`func (o *DatabaseServerV4StackRequest) SetDatabaseVendor(v string)`

SetDatabaseVendor sets DatabaseVendor field to given value.

### HasDatabaseVendor

`func (o *DatabaseServerV4StackRequest) HasDatabaseVendor() bool`

HasDatabaseVendor returns a boolean if a field has been set.

### GetConnectionDriver

`func (o *DatabaseServerV4StackRequest) GetConnectionDriver() string`

GetConnectionDriver returns the ConnectionDriver field if non-nil, zero value otherwise.

### GetConnectionDriverOk

`func (o *DatabaseServerV4StackRequest) GetConnectionDriverOk() (*string, bool)`

GetConnectionDriverOk returns a tuple with the ConnectionDriver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionDriver

`func (o *DatabaseServerV4StackRequest) SetConnectionDriver(v string)`

SetConnectionDriver sets ConnectionDriver field to given value.

### HasConnectionDriver

`func (o *DatabaseServerV4StackRequest) HasConnectionDriver() bool`

HasConnectionDriver returns a boolean if a field has been set.

### GetStorageSize

`func (o *DatabaseServerV4StackRequest) GetStorageSize() int64`

GetStorageSize returns the StorageSize field if non-nil, zero value otherwise.

### GetStorageSizeOk

`func (o *DatabaseServerV4StackRequest) GetStorageSizeOk() (*int64, bool)`

GetStorageSizeOk returns a tuple with the StorageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageSize

`func (o *DatabaseServerV4StackRequest) SetStorageSize(v int64)`

SetStorageSize sets StorageSize field to given value.

### HasStorageSize

`func (o *DatabaseServerV4StackRequest) HasStorageSize() bool`

HasStorageSize returns a boolean if a field has been set.

### GetRootUserName

`func (o *DatabaseServerV4StackRequest) GetRootUserName() string`

GetRootUserName returns the RootUserName field if non-nil, zero value otherwise.

### GetRootUserNameOk

`func (o *DatabaseServerV4StackRequest) GetRootUserNameOk() (*string, bool)`

GetRootUserNameOk returns a tuple with the RootUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootUserName

`func (o *DatabaseServerV4StackRequest) SetRootUserName(v string)`

SetRootUserName sets RootUserName field to given value.

### HasRootUserName

`func (o *DatabaseServerV4StackRequest) HasRootUserName() bool`

HasRootUserName returns a boolean if a field has been set.

### GetRootUserPassword

`func (o *DatabaseServerV4StackRequest) GetRootUserPassword() string`

GetRootUserPassword returns the RootUserPassword field if non-nil, zero value otherwise.

### GetRootUserPasswordOk

`func (o *DatabaseServerV4StackRequest) GetRootUserPasswordOk() (*string, bool)`

GetRootUserPasswordOk returns a tuple with the RootUserPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootUserPassword

`func (o *DatabaseServerV4StackRequest) SetRootUserPassword(v string)`

SetRootUserPassword sets RootUserPassword field to given value.

### HasRootUserPassword

`func (o *DatabaseServerV4StackRequest) HasRootUserPassword() bool`

HasRootUserPassword returns a boolean if a field has been set.

### GetPort

`func (o *DatabaseServerV4StackRequest) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *DatabaseServerV4StackRequest) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *DatabaseServerV4StackRequest) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *DatabaseServerV4StackRequest) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetAws

`func (o *DatabaseServerV4StackRequest) GetAws() AwsDatabaseServerV4Parameters`

GetAws returns the Aws field if non-nil, zero value otherwise.

### GetAwsOk

`func (o *DatabaseServerV4StackRequest) GetAwsOk() (*AwsDatabaseServerV4Parameters, bool)`

GetAwsOk returns a tuple with the Aws field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAws

`func (o *DatabaseServerV4StackRequest) SetAws(v AwsDatabaseServerV4Parameters)`

SetAws sets Aws field to given value.

### HasAws

`func (o *DatabaseServerV4StackRequest) HasAws() bool`

HasAws returns a boolean if a field has been set.

### GetAzure

`func (o *DatabaseServerV4StackRequest) GetAzure() AzureDatabaseServerV4Parameters`

GetAzure returns the Azure field if non-nil, zero value otherwise.

### GetAzureOk

`func (o *DatabaseServerV4StackRequest) GetAzureOk() (*AzureDatabaseServerV4Parameters, bool)`

GetAzureOk returns a tuple with the Azure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzure

`func (o *DatabaseServerV4StackRequest) SetAzure(v AzureDatabaseServerV4Parameters)`

SetAzure sets Azure field to given value.

### HasAzure

`func (o *DatabaseServerV4StackRequest) HasAzure() bool`

HasAzure returns a boolean if a field has been set.

### GetGcp

`func (o *DatabaseServerV4StackRequest) GetGcp() GcpDatabaseServerV4Parameters`

GetGcp returns the Gcp field if non-nil, zero value otherwise.

### GetGcpOk

`func (o *DatabaseServerV4StackRequest) GetGcpOk() (*GcpDatabaseServerV4Parameters, bool)`

GetGcpOk returns a tuple with the Gcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcp

`func (o *DatabaseServerV4StackRequest) SetGcp(v GcpDatabaseServerV4Parameters)`

SetGcp sets Gcp field to given value.

### HasGcp

`func (o *DatabaseServerV4StackRequest) HasGcp() bool`

HasGcp returns a boolean if a field has been set.

### GetSecurityGroup

`func (o *DatabaseServerV4StackRequest) GetSecurityGroup() SecurityGroupV4StackRequest`

GetSecurityGroup returns the SecurityGroup field if non-nil, zero value otherwise.

### GetSecurityGroupOk

`func (o *DatabaseServerV4StackRequest) GetSecurityGroupOk() (*SecurityGroupV4StackRequest, bool)`

GetSecurityGroupOk returns a tuple with the SecurityGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroup

`func (o *DatabaseServerV4StackRequest) SetSecurityGroup(v SecurityGroupV4StackRequest)`

SetSecurityGroup sets SecurityGroup field to given value.

### HasSecurityGroup

`func (o *DatabaseServerV4StackRequest) HasSecurityGroup() bool`

HasSecurityGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


