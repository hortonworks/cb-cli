# DatabaseTestV4Request

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExistingDatabase** | Pointer to [**DatabaseV4Identifiers**](DatabaseV4Identifiers.md) |  | [optional] 
**Database** | Pointer to [**DatabaseV4Request**](DatabaseV4Request.md) |  | [optional] 

## Methods

### NewDatabaseTestV4Request

`func NewDatabaseTestV4Request() *DatabaseTestV4Request`

NewDatabaseTestV4Request instantiates a new DatabaseTestV4Request object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabaseTestV4RequestWithDefaults

`func NewDatabaseTestV4RequestWithDefaults() *DatabaseTestV4Request`

NewDatabaseTestV4RequestWithDefaults instantiates a new DatabaseTestV4Request object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExistingDatabase

`func (o *DatabaseTestV4Request) GetExistingDatabase() DatabaseV4Identifiers`

GetExistingDatabase returns the ExistingDatabase field if non-nil, zero value otherwise.

### GetExistingDatabaseOk

`func (o *DatabaseTestV4Request) GetExistingDatabaseOk() (*DatabaseV4Identifiers, bool)`

GetExistingDatabaseOk returns a tuple with the ExistingDatabase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistingDatabase

`func (o *DatabaseTestV4Request) SetExistingDatabase(v DatabaseV4Identifiers)`

SetExistingDatabase sets ExistingDatabase field to given value.

### HasExistingDatabase

`func (o *DatabaseTestV4Request) HasExistingDatabase() bool`

HasExistingDatabase returns a boolean if a field has been set.

### GetDatabase

`func (o *DatabaseTestV4Request) GetDatabase() DatabaseV4Request`

GetDatabase returns the Database field if non-nil, zero value otherwise.

### GetDatabaseOk

`func (o *DatabaseTestV4Request) GetDatabaseOk() (*DatabaseV4Request, bool)`

GetDatabaseOk returns a tuple with the Database field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabase

`func (o *DatabaseTestV4Request) SetDatabase(v DatabaseV4Request)`

SetDatabase sets Database field to given value.

### HasDatabase

`func (o *DatabaseTestV4Request) HasDatabase() bool`

HasDatabase returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


