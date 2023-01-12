# DatabaseServerTestV4Request

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExistingDatabaseServerCrn** | Pointer to **string** | CRN of registered database server to be tested for connectivity | [optional] 
**DatabaseServer** | Pointer to [**DatabaseServerV4Request**](DatabaseServerV4Request.md) |  | [optional] 

## Methods

### NewDatabaseServerTestV4Request

`func NewDatabaseServerTestV4Request() *DatabaseServerTestV4Request`

NewDatabaseServerTestV4Request instantiates a new DatabaseServerTestV4Request object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabaseServerTestV4RequestWithDefaults

`func NewDatabaseServerTestV4RequestWithDefaults() *DatabaseServerTestV4Request`

NewDatabaseServerTestV4RequestWithDefaults instantiates a new DatabaseServerTestV4Request object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExistingDatabaseServerCrn

`func (o *DatabaseServerTestV4Request) GetExistingDatabaseServerCrn() string`

GetExistingDatabaseServerCrn returns the ExistingDatabaseServerCrn field if non-nil, zero value otherwise.

### GetExistingDatabaseServerCrnOk

`func (o *DatabaseServerTestV4Request) GetExistingDatabaseServerCrnOk() (*string, bool)`

GetExistingDatabaseServerCrnOk returns a tuple with the ExistingDatabaseServerCrn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistingDatabaseServerCrn

`func (o *DatabaseServerTestV4Request) SetExistingDatabaseServerCrn(v string)`

SetExistingDatabaseServerCrn sets ExistingDatabaseServerCrn field to given value.

### HasExistingDatabaseServerCrn

`func (o *DatabaseServerTestV4Request) HasExistingDatabaseServerCrn() bool`

HasExistingDatabaseServerCrn returns a boolean if a field has been set.

### GetDatabaseServer

`func (o *DatabaseServerTestV4Request) GetDatabaseServer() DatabaseServerV4Request`

GetDatabaseServer returns the DatabaseServer field if non-nil, zero value otherwise.

### GetDatabaseServerOk

`func (o *DatabaseServerTestV4Request) GetDatabaseServerOk() (*DatabaseServerV4Request, bool)`

GetDatabaseServerOk returns a tuple with the DatabaseServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseServer

`func (o *DatabaseServerTestV4Request) SetDatabaseServer(v DatabaseServerV4Request)`

SetDatabaseServer sets DatabaseServer field to given value.

### HasDatabaseServer

`func (o *DatabaseServerTestV4Request) HasDatabaseServer() bool`

HasDatabaseServer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


