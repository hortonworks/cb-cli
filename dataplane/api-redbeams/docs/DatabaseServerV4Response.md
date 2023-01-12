# DatabaseServerV4Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the database server | 
**Description** | Pointer to **string** | Description of the database server | [optional] 
**Host** | **string** | Host of the database server | 
**Port** | **int32** | Port of the database server | 
**DatabaseVendor** | **string** | Name of the database vendor (MYSQL, POSTGRES, ...) | 
**ConnectionDriver** | **string** | Name of the JDBC connection driver (for example: &#39;org.postgresql.Driver&#39;) | 
**EnvironmentCrn** | **string** | CRN of the environment of the database server | 
**Id** | Pointer to **int64** | Internal ID of the database server | [optional] 
**Crn** | Pointer to **string** | CRN of the database server | [optional] 
**DatabaseVendorDisplayName** | **string** | Display name of the database vendor (MySQL, PostgreSQL, ...) | 
**ConnectionUserName** | Pointer to [**SecretResponse**](SecretResponse.md) |  | [optional] 
**ConnectionPassword** | Pointer to [**SecretResponse**](SecretResponse.md) |  | [optional] 
**CreationDate** | Pointer to **int64** | Creation date / time of the database server, in epoch milliseconds | [optional] 
**ResourceStatus** | Pointer to **string** | Ownership status of the database server | [optional] 
**Status** | Pointer to **string** | Status of the database server stack | [optional] 
**StatusReason** | Pointer to **string** | Additional status information about the database server stack | [optional] 
**SslConfig** | Pointer to [**SslConfigV4Response**](SslConfigV4Response.md) |  | [optional] 
**ClusterCrn** | Pointer to **string** | CRN of the cluster of the database server | [optional] 
**MajorVersion** | Pointer to **string** | Major version of the database server engine | [optional] 

## Methods

### NewDatabaseServerV4Response

`func NewDatabaseServerV4Response(name string, host string, port int32, databaseVendor string, connectionDriver string, environmentCrn string, databaseVendorDisplayName string, ) *DatabaseServerV4Response`

NewDatabaseServerV4Response instantiates a new DatabaseServerV4Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabaseServerV4ResponseWithDefaults

`func NewDatabaseServerV4ResponseWithDefaults() *DatabaseServerV4Response`

NewDatabaseServerV4ResponseWithDefaults instantiates a new DatabaseServerV4Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DatabaseServerV4Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DatabaseServerV4Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DatabaseServerV4Response) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *DatabaseServerV4Response) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DatabaseServerV4Response) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DatabaseServerV4Response) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DatabaseServerV4Response) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetHost

`func (o *DatabaseServerV4Response) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *DatabaseServerV4Response) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *DatabaseServerV4Response) SetHost(v string)`

SetHost sets Host field to given value.


### GetPort

`func (o *DatabaseServerV4Response) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *DatabaseServerV4Response) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *DatabaseServerV4Response) SetPort(v int32)`

SetPort sets Port field to given value.


### GetDatabaseVendor

`func (o *DatabaseServerV4Response) GetDatabaseVendor() string`

GetDatabaseVendor returns the DatabaseVendor field if non-nil, zero value otherwise.

### GetDatabaseVendorOk

`func (o *DatabaseServerV4Response) GetDatabaseVendorOk() (*string, bool)`

GetDatabaseVendorOk returns a tuple with the DatabaseVendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseVendor

`func (o *DatabaseServerV4Response) SetDatabaseVendor(v string)`

SetDatabaseVendor sets DatabaseVendor field to given value.


### GetConnectionDriver

`func (o *DatabaseServerV4Response) GetConnectionDriver() string`

GetConnectionDriver returns the ConnectionDriver field if non-nil, zero value otherwise.

### GetConnectionDriverOk

`func (o *DatabaseServerV4Response) GetConnectionDriverOk() (*string, bool)`

GetConnectionDriverOk returns a tuple with the ConnectionDriver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionDriver

`func (o *DatabaseServerV4Response) SetConnectionDriver(v string)`

SetConnectionDriver sets ConnectionDriver field to given value.


### GetEnvironmentCrn

`func (o *DatabaseServerV4Response) GetEnvironmentCrn() string`

GetEnvironmentCrn returns the EnvironmentCrn field if non-nil, zero value otherwise.

### GetEnvironmentCrnOk

`func (o *DatabaseServerV4Response) GetEnvironmentCrnOk() (*string, bool)`

GetEnvironmentCrnOk returns a tuple with the EnvironmentCrn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentCrn

`func (o *DatabaseServerV4Response) SetEnvironmentCrn(v string)`

SetEnvironmentCrn sets EnvironmentCrn field to given value.


### GetId

`func (o *DatabaseServerV4Response) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DatabaseServerV4Response) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DatabaseServerV4Response) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *DatabaseServerV4Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCrn

`func (o *DatabaseServerV4Response) GetCrn() string`

GetCrn returns the Crn field if non-nil, zero value otherwise.

### GetCrnOk

`func (o *DatabaseServerV4Response) GetCrnOk() (*string, bool)`

GetCrnOk returns a tuple with the Crn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrn

`func (o *DatabaseServerV4Response) SetCrn(v string)`

SetCrn sets Crn field to given value.

### HasCrn

`func (o *DatabaseServerV4Response) HasCrn() bool`

HasCrn returns a boolean if a field has been set.

### GetDatabaseVendorDisplayName

`func (o *DatabaseServerV4Response) GetDatabaseVendorDisplayName() string`

GetDatabaseVendorDisplayName returns the DatabaseVendorDisplayName field if non-nil, zero value otherwise.

### GetDatabaseVendorDisplayNameOk

`func (o *DatabaseServerV4Response) GetDatabaseVendorDisplayNameOk() (*string, bool)`

GetDatabaseVendorDisplayNameOk returns a tuple with the DatabaseVendorDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseVendorDisplayName

`func (o *DatabaseServerV4Response) SetDatabaseVendorDisplayName(v string)`

SetDatabaseVendorDisplayName sets DatabaseVendorDisplayName field to given value.


### GetConnectionUserName

`func (o *DatabaseServerV4Response) GetConnectionUserName() SecretResponse`

GetConnectionUserName returns the ConnectionUserName field if non-nil, zero value otherwise.

### GetConnectionUserNameOk

`func (o *DatabaseServerV4Response) GetConnectionUserNameOk() (*SecretResponse, bool)`

GetConnectionUserNameOk returns a tuple with the ConnectionUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionUserName

`func (o *DatabaseServerV4Response) SetConnectionUserName(v SecretResponse)`

SetConnectionUserName sets ConnectionUserName field to given value.

### HasConnectionUserName

`func (o *DatabaseServerV4Response) HasConnectionUserName() bool`

HasConnectionUserName returns a boolean if a field has been set.

### GetConnectionPassword

`func (o *DatabaseServerV4Response) GetConnectionPassword() SecretResponse`

GetConnectionPassword returns the ConnectionPassword field if non-nil, zero value otherwise.

### GetConnectionPasswordOk

`func (o *DatabaseServerV4Response) GetConnectionPasswordOk() (*SecretResponse, bool)`

GetConnectionPasswordOk returns a tuple with the ConnectionPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionPassword

`func (o *DatabaseServerV4Response) SetConnectionPassword(v SecretResponse)`

SetConnectionPassword sets ConnectionPassword field to given value.

### HasConnectionPassword

`func (o *DatabaseServerV4Response) HasConnectionPassword() bool`

HasConnectionPassword returns a boolean if a field has been set.

### GetCreationDate

`func (o *DatabaseServerV4Response) GetCreationDate() int64`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *DatabaseServerV4Response) GetCreationDateOk() (*int64, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *DatabaseServerV4Response) SetCreationDate(v int64)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *DatabaseServerV4Response) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.

### GetResourceStatus

`func (o *DatabaseServerV4Response) GetResourceStatus() string`

GetResourceStatus returns the ResourceStatus field if non-nil, zero value otherwise.

### GetResourceStatusOk

`func (o *DatabaseServerV4Response) GetResourceStatusOk() (*string, bool)`

GetResourceStatusOk returns a tuple with the ResourceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceStatus

`func (o *DatabaseServerV4Response) SetResourceStatus(v string)`

SetResourceStatus sets ResourceStatus field to given value.

### HasResourceStatus

`func (o *DatabaseServerV4Response) HasResourceStatus() bool`

HasResourceStatus returns a boolean if a field has been set.

### GetStatus

`func (o *DatabaseServerV4Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DatabaseServerV4Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DatabaseServerV4Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DatabaseServerV4Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusReason

`func (o *DatabaseServerV4Response) GetStatusReason() string`

GetStatusReason returns the StatusReason field if non-nil, zero value otherwise.

### GetStatusReasonOk

`func (o *DatabaseServerV4Response) GetStatusReasonOk() (*string, bool)`

GetStatusReasonOk returns a tuple with the StatusReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusReason

`func (o *DatabaseServerV4Response) SetStatusReason(v string)`

SetStatusReason sets StatusReason field to given value.

### HasStatusReason

`func (o *DatabaseServerV4Response) HasStatusReason() bool`

HasStatusReason returns a boolean if a field has been set.

### GetSslConfig

`func (o *DatabaseServerV4Response) GetSslConfig() SslConfigV4Response`

GetSslConfig returns the SslConfig field if non-nil, zero value otherwise.

### GetSslConfigOk

`func (o *DatabaseServerV4Response) GetSslConfigOk() (*SslConfigV4Response, bool)`

GetSslConfigOk returns a tuple with the SslConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslConfig

`func (o *DatabaseServerV4Response) SetSslConfig(v SslConfigV4Response)`

SetSslConfig sets SslConfig field to given value.

### HasSslConfig

`func (o *DatabaseServerV4Response) HasSslConfig() bool`

HasSslConfig returns a boolean if a field has been set.

### GetClusterCrn

`func (o *DatabaseServerV4Response) GetClusterCrn() string`

GetClusterCrn returns the ClusterCrn field if non-nil, zero value otherwise.

### GetClusterCrnOk

`func (o *DatabaseServerV4Response) GetClusterCrnOk() (*string, bool)`

GetClusterCrnOk returns a tuple with the ClusterCrn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterCrn

`func (o *DatabaseServerV4Response) SetClusterCrn(v string)`

SetClusterCrn sets ClusterCrn field to given value.

### HasClusterCrn

`func (o *DatabaseServerV4Response) HasClusterCrn() bool`

HasClusterCrn returns a boolean if a field has been set.

### GetMajorVersion

`func (o *DatabaseServerV4Response) GetMajorVersion() string`

GetMajorVersion returns the MajorVersion field if non-nil, zero value otherwise.

### GetMajorVersionOk

`func (o *DatabaseServerV4Response) GetMajorVersionOk() (*string, bool)`

GetMajorVersionOk returns a tuple with the MajorVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMajorVersion

`func (o *DatabaseServerV4Response) SetMajorVersion(v string)`

SetMajorVersion sets MajorVersion field to given value.

### HasMajorVersion

`func (o *DatabaseServerV4Response) HasMajorVersion() bool`

HasMajorVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


