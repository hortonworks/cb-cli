# DatabaseV4Identifiers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the database | 
**EnvironmentCrn** | **string** | CRN of the environment of the database | 

## Methods

### NewDatabaseV4Identifiers

`func NewDatabaseV4Identifiers(name string, environmentCrn string, ) *DatabaseV4Identifiers`

NewDatabaseV4Identifiers instantiates a new DatabaseV4Identifiers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabaseV4IdentifiersWithDefaults

`func NewDatabaseV4IdentifiersWithDefaults() *DatabaseV4Identifiers`

NewDatabaseV4IdentifiersWithDefaults instantiates a new DatabaseV4Identifiers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DatabaseV4Identifiers) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DatabaseV4Identifiers) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DatabaseV4Identifiers) SetName(v string)`

SetName sets Name field to given value.


### GetEnvironmentCrn

`func (o *DatabaseV4Identifiers) GetEnvironmentCrn() string`

GetEnvironmentCrn returns the EnvironmentCrn field if non-nil, zero value otherwise.

### GetEnvironmentCrnOk

`func (o *DatabaseV4Identifiers) GetEnvironmentCrnOk() (*string, bool)`

GetEnvironmentCrnOk returns a tuple with the EnvironmentCrn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentCrn

`func (o *DatabaseV4Identifiers) SetEnvironmentCrn(v string)`

SetEnvironmentCrn sets EnvironmentCrn field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


