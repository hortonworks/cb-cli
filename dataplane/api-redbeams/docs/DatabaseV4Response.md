# DatabaseV4Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the database | 
**Description** | Pointer to **string** | Description of the database | [optional] 
**ConnectionURL** | **string** | JDBC connection URL in the form of jdbc:&lt;db-type&gt;:&lt;driver-specific-part&gt; | 
**Type** | **string** | Type of database, i.e., the service name that will use the database (HIVE, DRUID, SUPERSET, RANGER, ...) | 
**ConnectionDriver** | **string** | Name of the JDBC connection driver (for example: &#39;org.postgresql.Driver&#39;) | 
**EnvironmentCrn** | **string** | CRN of the environment of the database | 
**Crn** | Pointer to **string** | CRN of the database | [optional] 
**CreationDate** | Pointer to **int64** | Creation date / time of the database, in epoch milliseconds | [optional] 
**DatabaseEngine** | **string** | Name of the database vendor (MYSQL, POSTGRES...) | 
**DatabaseEngineDisplayName** | **string** | Display name of the database vendor (MySQL, PostgreSQL, ...) | 
**ConnectionUserName** | Pointer to [**SecretResponse**](SecretResponse.md) |  | [optional] 
**ConnectionPassword** | Pointer to [**SecretResponse**](SecretResponse.md) |  | [optional] 
**ResourceStatus** | Pointer to **string** | Ownership status of the database | [optional] 

## Methods

### NewDatabaseV4Response

`func NewDatabaseV4Response(name string, connectionURL string, type_ string, connectionDriver string, environmentCrn string, databaseEngine string, databaseEngineDisplayName string, ) *DatabaseV4Response`

NewDatabaseV4Response instantiates a new DatabaseV4Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabaseV4ResponseWithDefaults

`func NewDatabaseV4ResponseWithDefaults() *DatabaseV4Response`

NewDatabaseV4ResponseWithDefaults instantiates a new DatabaseV4Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DatabaseV4Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DatabaseV4Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DatabaseV4Response) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *DatabaseV4Response) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DatabaseV4Response) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DatabaseV4Response) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DatabaseV4Response) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetConnectionURL

`func (o *DatabaseV4Response) GetConnectionURL() string`

GetConnectionURL returns the ConnectionURL field if non-nil, zero value otherwise.

### GetConnectionURLOk

`func (o *DatabaseV4Response) GetConnectionURLOk() (*string, bool)`

GetConnectionURLOk returns a tuple with the ConnectionURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionURL

`func (o *DatabaseV4Response) SetConnectionURL(v string)`

SetConnectionURL sets ConnectionURL field to given value.


### GetType

`func (o *DatabaseV4Response) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DatabaseV4Response) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DatabaseV4Response) SetType(v string)`

SetType sets Type field to given value.


### GetConnectionDriver

`func (o *DatabaseV4Response) GetConnectionDriver() string`

GetConnectionDriver returns the ConnectionDriver field if non-nil, zero value otherwise.

### GetConnectionDriverOk

`func (o *DatabaseV4Response) GetConnectionDriverOk() (*string, bool)`

GetConnectionDriverOk returns a tuple with the ConnectionDriver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionDriver

`func (o *DatabaseV4Response) SetConnectionDriver(v string)`

SetConnectionDriver sets ConnectionDriver field to given value.


### GetEnvironmentCrn

`func (o *DatabaseV4Response) GetEnvironmentCrn() string`

GetEnvironmentCrn returns the EnvironmentCrn field if non-nil, zero value otherwise.

### GetEnvironmentCrnOk

`func (o *DatabaseV4Response) GetEnvironmentCrnOk() (*string, bool)`

GetEnvironmentCrnOk returns a tuple with the EnvironmentCrn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentCrn

`func (o *DatabaseV4Response) SetEnvironmentCrn(v string)`

SetEnvironmentCrn sets EnvironmentCrn field to given value.


### GetCrn

`func (o *DatabaseV4Response) GetCrn() string`

GetCrn returns the Crn field if non-nil, zero value otherwise.

### GetCrnOk

`func (o *DatabaseV4Response) GetCrnOk() (*string, bool)`

GetCrnOk returns a tuple with the Crn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrn

`func (o *DatabaseV4Response) SetCrn(v string)`

SetCrn sets Crn field to given value.

### HasCrn

`func (o *DatabaseV4Response) HasCrn() bool`

HasCrn returns a boolean if a field has been set.

### GetCreationDate

`func (o *DatabaseV4Response) GetCreationDate() int64`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *DatabaseV4Response) GetCreationDateOk() (*int64, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *DatabaseV4Response) SetCreationDate(v int64)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *DatabaseV4Response) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.

### GetDatabaseEngine

`func (o *DatabaseV4Response) GetDatabaseEngine() string`

GetDatabaseEngine returns the DatabaseEngine field if non-nil, zero value otherwise.

### GetDatabaseEngineOk

`func (o *DatabaseV4Response) GetDatabaseEngineOk() (*string, bool)`

GetDatabaseEngineOk returns a tuple with the DatabaseEngine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseEngine

`func (o *DatabaseV4Response) SetDatabaseEngine(v string)`

SetDatabaseEngine sets DatabaseEngine field to given value.


### GetDatabaseEngineDisplayName

`func (o *DatabaseV4Response) GetDatabaseEngineDisplayName() string`

GetDatabaseEngineDisplayName returns the DatabaseEngineDisplayName field if non-nil, zero value otherwise.

### GetDatabaseEngineDisplayNameOk

`func (o *DatabaseV4Response) GetDatabaseEngineDisplayNameOk() (*string, bool)`

GetDatabaseEngineDisplayNameOk returns a tuple with the DatabaseEngineDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseEngineDisplayName

`func (o *DatabaseV4Response) SetDatabaseEngineDisplayName(v string)`

SetDatabaseEngineDisplayName sets DatabaseEngineDisplayName field to given value.


### GetConnectionUserName

`func (o *DatabaseV4Response) GetConnectionUserName() SecretResponse`

GetConnectionUserName returns the ConnectionUserName field if non-nil, zero value otherwise.

### GetConnectionUserNameOk

`func (o *DatabaseV4Response) GetConnectionUserNameOk() (*SecretResponse, bool)`

GetConnectionUserNameOk returns a tuple with the ConnectionUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionUserName

`func (o *DatabaseV4Response) SetConnectionUserName(v SecretResponse)`

SetConnectionUserName sets ConnectionUserName field to given value.

### HasConnectionUserName

`func (o *DatabaseV4Response) HasConnectionUserName() bool`

HasConnectionUserName returns a boolean if a field has been set.

### GetConnectionPassword

`func (o *DatabaseV4Response) GetConnectionPassword() SecretResponse`

GetConnectionPassword returns the ConnectionPassword field if non-nil, zero value otherwise.

### GetConnectionPasswordOk

`func (o *DatabaseV4Response) GetConnectionPasswordOk() (*SecretResponse, bool)`

GetConnectionPasswordOk returns a tuple with the ConnectionPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionPassword

`func (o *DatabaseV4Response) SetConnectionPassword(v SecretResponse)`

SetConnectionPassword sets ConnectionPassword field to given value.

### HasConnectionPassword

`func (o *DatabaseV4Response) HasConnectionPassword() bool`

HasConnectionPassword returns a boolean if a field has been set.

### GetResourceStatus

`func (o *DatabaseV4Response) GetResourceStatus() string`

GetResourceStatus returns the ResourceStatus field if non-nil, zero value otherwise.

### GetResourceStatusOk

`func (o *DatabaseV4Response) GetResourceStatusOk() (*string, bool)`

GetResourceStatusOk returns a tuple with the ResourceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceStatus

`func (o *DatabaseV4Response) SetResourceStatus(v string)`

SetResourceStatus sets ResourceStatus field to given value.

### HasResourceStatus

`func (o *DatabaseV4Response) HasResourceStatus() bool`

HasResourceStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


