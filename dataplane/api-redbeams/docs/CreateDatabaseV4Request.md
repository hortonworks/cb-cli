# CreateDatabaseV4Request

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExistingDatabaseServerCrn** | **string** | CRN of the database server | 
**DatabaseName** | **string** | Name of the database | 
**Type** | **string** | Type of database, i.e., the service name that will use the database (HIVE, DRUID, SUPERSET, RANGER, ...) | 
**DatabaseDescription** | Pointer to **string** | Description of the database | [optional] 

## Methods

### NewCreateDatabaseV4Request

`func NewCreateDatabaseV4Request(existingDatabaseServerCrn string, databaseName string, type_ string, ) *CreateDatabaseV4Request`

NewCreateDatabaseV4Request instantiates a new CreateDatabaseV4Request object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDatabaseV4RequestWithDefaults

`func NewCreateDatabaseV4RequestWithDefaults() *CreateDatabaseV4Request`

NewCreateDatabaseV4RequestWithDefaults instantiates a new CreateDatabaseV4Request object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExistingDatabaseServerCrn

`func (o *CreateDatabaseV4Request) GetExistingDatabaseServerCrn() string`

GetExistingDatabaseServerCrn returns the ExistingDatabaseServerCrn field if non-nil, zero value otherwise.

### GetExistingDatabaseServerCrnOk

`func (o *CreateDatabaseV4Request) GetExistingDatabaseServerCrnOk() (*string, bool)`

GetExistingDatabaseServerCrnOk returns a tuple with the ExistingDatabaseServerCrn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistingDatabaseServerCrn

`func (o *CreateDatabaseV4Request) SetExistingDatabaseServerCrn(v string)`

SetExistingDatabaseServerCrn sets ExistingDatabaseServerCrn field to given value.


### GetDatabaseName

`func (o *CreateDatabaseV4Request) GetDatabaseName() string`

GetDatabaseName returns the DatabaseName field if non-nil, zero value otherwise.

### GetDatabaseNameOk

`func (o *CreateDatabaseV4Request) GetDatabaseNameOk() (*string, bool)`

GetDatabaseNameOk returns a tuple with the DatabaseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseName

`func (o *CreateDatabaseV4Request) SetDatabaseName(v string)`

SetDatabaseName sets DatabaseName field to given value.


### GetType

`func (o *CreateDatabaseV4Request) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateDatabaseV4Request) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateDatabaseV4Request) SetType(v string)`

SetType sets Type field to given value.


### GetDatabaseDescription

`func (o *CreateDatabaseV4Request) GetDatabaseDescription() string`

GetDatabaseDescription returns the DatabaseDescription field if non-nil, zero value otherwise.

### GetDatabaseDescriptionOk

`func (o *CreateDatabaseV4Request) GetDatabaseDescriptionOk() (*string, bool)`

GetDatabaseDescriptionOk returns a tuple with the DatabaseDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseDescription

`func (o *CreateDatabaseV4Request) SetDatabaseDescription(v string)`

SetDatabaseDescription sets DatabaseDescription field to given value.

### HasDatabaseDescription

`func (o *CreateDatabaseV4Request) HasDatabaseDescription() bool`

HasDatabaseDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


