# CheckResourceRightsV4Request

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rights** | Pointer to **[]string** |  | [optional] 
**ResourceRights** | Pointer to [**[]ResourceRightsV4**](ResourceRightsV4.md) |  | [optional] 

## Methods

### NewCheckResourceRightsV4Request

`func NewCheckResourceRightsV4Request() *CheckResourceRightsV4Request`

NewCheckResourceRightsV4Request instantiates a new CheckResourceRightsV4Request object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckResourceRightsV4RequestWithDefaults

`func NewCheckResourceRightsV4RequestWithDefaults() *CheckResourceRightsV4Request`

NewCheckResourceRightsV4RequestWithDefaults instantiates a new CheckResourceRightsV4Request object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRights

`func (o *CheckResourceRightsV4Request) GetRights() []string`

GetRights returns the Rights field if non-nil, zero value otherwise.

### GetRightsOk

`func (o *CheckResourceRightsV4Request) GetRightsOk() (*[]string, bool)`

GetRightsOk returns a tuple with the Rights field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRights

`func (o *CheckResourceRightsV4Request) SetRights(v []string)`

SetRights sets Rights field to given value.

### HasRights

`func (o *CheckResourceRightsV4Request) HasRights() bool`

HasRights returns a boolean if a field has been set.

### GetResourceRights

`func (o *CheckResourceRightsV4Request) GetResourceRights() []ResourceRightsV4`

GetResourceRights returns the ResourceRights field if non-nil, zero value otherwise.

### GetResourceRightsOk

`func (o *CheckResourceRightsV4Request) GetResourceRightsOk() (*[]ResourceRightsV4, bool)`

GetResourceRightsOk returns a tuple with the ResourceRights field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceRights

`func (o *CheckResourceRightsV4Request) SetResourceRights(v []ResourceRightsV4)`

SetResourceRights sets ResourceRights field to given value.

### HasResourceRights

`func (o *CheckResourceRightsV4Request) HasResourceRights() bool`

HasResourceRights returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


