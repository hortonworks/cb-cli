# DatabaseServerV4Request

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the database server | 
**Description** | Pointer to **string** | Description of the database server | [optional] 
**Host** | **string** | Host of the database server | 
**Port** | **int32** | Port of the database server | 
**DatabaseVendor** | **string** | Name of the database vendor (MYSQL, POSTGRES, ...) | 
**ConnectionDriver** | Pointer to **string** | Name of the JDBC connection driver (for example: &#39;org.postgresql.Driver&#39;) | [optional] 
**EnvironmentCrn** | **string** | CRN of the environment of the database server | 
**ConnectionUserName** | **string** | Username of the administrative user of the database server | 
**ConnectionPassword** | **string** | Password of the administrative user of the database server | 

## Methods

### NewDatabaseServerV4Request

`func NewDatabaseServerV4Request(name string, host string, port int32, databaseVendor string, environmentCrn string, connectionUserName string, connectionPassword string, ) *DatabaseServerV4Request`

NewDatabaseServerV4Request instantiates a new DatabaseServerV4Request object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabaseServerV4RequestWithDefaults

`func NewDatabaseServerV4RequestWithDefaults() *DatabaseServerV4Request`

NewDatabaseServerV4RequestWithDefaults instantiates a new DatabaseServerV4Request object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DatabaseServerV4Request) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DatabaseServerV4Request) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DatabaseServerV4Request) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *DatabaseServerV4Request) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DatabaseServerV4Request) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DatabaseServerV4Request) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DatabaseServerV4Request) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetHost

`func (o *DatabaseServerV4Request) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *DatabaseServerV4Request) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *DatabaseServerV4Request) SetHost(v string)`

SetHost sets Host field to given value.


### GetPort

`func (o *DatabaseServerV4Request) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *DatabaseServerV4Request) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *DatabaseServerV4Request) SetPort(v int32)`

SetPort sets Port field to given value.


### GetDatabaseVendor

`func (o *DatabaseServerV4Request) GetDatabaseVendor() string`

GetDatabaseVendor returns the DatabaseVendor field if non-nil, zero value otherwise.

### GetDatabaseVendorOk

`func (o *DatabaseServerV4Request) GetDatabaseVendorOk() (*string, bool)`

GetDatabaseVendorOk returns a tuple with the DatabaseVendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseVendor

`func (o *DatabaseServerV4Request) SetDatabaseVendor(v string)`

SetDatabaseVendor sets DatabaseVendor field to given value.


### GetConnectionDriver

`func (o *DatabaseServerV4Request) GetConnectionDriver() string`

GetConnectionDriver returns the ConnectionDriver field if non-nil, zero value otherwise.

### GetConnectionDriverOk

`func (o *DatabaseServerV4Request) GetConnectionDriverOk() (*string, bool)`

GetConnectionDriverOk returns a tuple with the ConnectionDriver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionDriver

`func (o *DatabaseServerV4Request) SetConnectionDriver(v string)`

SetConnectionDriver sets ConnectionDriver field to given value.

### HasConnectionDriver

`func (o *DatabaseServerV4Request) HasConnectionDriver() bool`

HasConnectionDriver returns a boolean if a field has been set.

### GetEnvironmentCrn

`func (o *DatabaseServerV4Request) GetEnvironmentCrn() string`

GetEnvironmentCrn returns the EnvironmentCrn field if non-nil, zero value otherwise.

### GetEnvironmentCrnOk

`func (o *DatabaseServerV4Request) GetEnvironmentCrnOk() (*string, bool)`

GetEnvironmentCrnOk returns a tuple with the EnvironmentCrn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentCrn

`func (o *DatabaseServerV4Request) SetEnvironmentCrn(v string)`

SetEnvironmentCrn sets EnvironmentCrn field to given value.


### GetConnectionUserName

`func (o *DatabaseServerV4Request) GetConnectionUserName() string`

GetConnectionUserName returns the ConnectionUserName field if non-nil, zero value otherwise.

### GetConnectionUserNameOk

`func (o *DatabaseServerV4Request) GetConnectionUserNameOk() (*string, bool)`

GetConnectionUserNameOk returns a tuple with the ConnectionUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionUserName

`func (o *DatabaseServerV4Request) SetConnectionUserName(v string)`

SetConnectionUserName sets ConnectionUserName field to given value.


### GetConnectionPassword

`func (o *DatabaseServerV4Request) GetConnectionPassword() string`

GetConnectionPassword returns the ConnectionPassword field if non-nil, zero value otherwise.

### GetConnectionPasswordOk

`func (o *DatabaseServerV4Request) GetConnectionPasswordOk() (*string, bool)`

GetConnectionPasswordOk returns a tuple with the ConnectionPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionPassword

`func (o *DatabaseServerV4Request) SetConnectionPassword(v string)`

SetConnectionPassword sets ConnectionPassword field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


