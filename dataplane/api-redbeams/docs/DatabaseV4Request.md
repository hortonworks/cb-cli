# DatabaseV4Request

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the database | 
**Description** | Pointer to **string** | Description of the database | [optional] 
**ConnectionURL** | **string** | JDBC connection URL in the form of jdbc:&lt;db-type&gt;:&lt;driver-specific-part&gt; | 
**Type** | **string** | Type of database, i.e., the service name that will use the database (HIVE, DRUID, SUPERSET, RANGER, ...) | 
**ConnectionDriver** | Pointer to **string** | Name of the JDBC connection driver (for example: &#39;org.postgresql.Driver&#39;) | [optional] 
**EnvironmentCrn** | **string** | CRN of the environment of the database | 
**ConnectionUserName** | **string** | Username to use for authentication | 
**ConnectionPassword** | **string** | Password to use for authentication | 

## Methods

### NewDatabaseV4Request

`func NewDatabaseV4Request(name string, connectionURL string, type_ string, environmentCrn string, connectionUserName string, connectionPassword string, ) *DatabaseV4Request`

NewDatabaseV4Request instantiates a new DatabaseV4Request object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabaseV4RequestWithDefaults

`func NewDatabaseV4RequestWithDefaults() *DatabaseV4Request`

NewDatabaseV4RequestWithDefaults instantiates a new DatabaseV4Request object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DatabaseV4Request) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DatabaseV4Request) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DatabaseV4Request) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *DatabaseV4Request) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DatabaseV4Request) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DatabaseV4Request) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DatabaseV4Request) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetConnectionURL

`func (o *DatabaseV4Request) GetConnectionURL() string`

GetConnectionURL returns the ConnectionURL field if non-nil, zero value otherwise.

### GetConnectionURLOk

`func (o *DatabaseV4Request) GetConnectionURLOk() (*string, bool)`

GetConnectionURLOk returns a tuple with the ConnectionURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionURL

`func (o *DatabaseV4Request) SetConnectionURL(v string)`

SetConnectionURL sets ConnectionURL field to given value.


### GetType

`func (o *DatabaseV4Request) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DatabaseV4Request) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DatabaseV4Request) SetType(v string)`

SetType sets Type field to given value.


### GetConnectionDriver

`func (o *DatabaseV4Request) GetConnectionDriver() string`

GetConnectionDriver returns the ConnectionDriver field if non-nil, zero value otherwise.

### GetConnectionDriverOk

`func (o *DatabaseV4Request) GetConnectionDriverOk() (*string, bool)`

GetConnectionDriverOk returns a tuple with the ConnectionDriver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionDriver

`func (o *DatabaseV4Request) SetConnectionDriver(v string)`

SetConnectionDriver sets ConnectionDriver field to given value.

### HasConnectionDriver

`func (o *DatabaseV4Request) HasConnectionDriver() bool`

HasConnectionDriver returns a boolean if a field has been set.

### GetEnvironmentCrn

`func (o *DatabaseV4Request) GetEnvironmentCrn() string`

GetEnvironmentCrn returns the EnvironmentCrn field if non-nil, zero value otherwise.

### GetEnvironmentCrnOk

`func (o *DatabaseV4Request) GetEnvironmentCrnOk() (*string, bool)`

GetEnvironmentCrnOk returns a tuple with the EnvironmentCrn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentCrn

`func (o *DatabaseV4Request) SetEnvironmentCrn(v string)`

SetEnvironmentCrn sets EnvironmentCrn field to given value.


### GetConnectionUserName

`func (o *DatabaseV4Request) GetConnectionUserName() string`

GetConnectionUserName returns the ConnectionUserName field if non-nil, zero value otherwise.

### GetConnectionUserNameOk

`func (o *DatabaseV4Request) GetConnectionUserNameOk() (*string, bool)`

GetConnectionUserNameOk returns a tuple with the ConnectionUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionUserName

`func (o *DatabaseV4Request) SetConnectionUserName(v string)`

SetConnectionUserName sets ConnectionUserName field to given value.


### GetConnectionPassword

`func (o *DatabaseV4Request) GetConnectionPassword() string`

GetConnectionPassword returns the ConnectionPassword field if non-nil, zero value otherwise.

### GetConnectionPasswordOk

`func (o *DatabaseV4Request) GetConnectionPasswordOk() (*string, bool)`

GetConnectionPasswordOk returns a tuple with the ConnectionPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionPassword

`func (o *DatabaseV4Request) SetConnectionPassword(v string)`

SetConnectionPassword sets ConnectionPassword field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


