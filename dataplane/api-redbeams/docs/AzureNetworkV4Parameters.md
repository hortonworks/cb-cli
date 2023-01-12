# AzureNetworkV4Parameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subnets** | Pointer to **string** | Comma-separated list of fully-qualified subnets with connectivity to the database server | [optional] 

## Methods

### NewAzureNetworkV4Parameters

`func NewAzureNetworkV4Parameters() *AzureNetworkV4Parameters`

NewAzureNetworkV4Parameters instantiates a new AzureNetworkV4Parameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureNetworkV4ParametersWithDefaults

`func NewAzureNetworkV4ParametersWithDefaults() *AzureNetworkV4Parameters`

NewAzureNetworkV4ParametersWithDefaults instantiates a new AzureNetworkV4Parameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubnets

`func (o *AzureNetworkV4Parameters) GetSubnets() string`

GetSubnets returns the Subnets field if non-nil, zero value otherwise.

### GetSubnetsOk

`func (o *AzureNetworkV4Parameters) GetSubnetsOk() (*string, bool)`

GetSubnetsOk returns a tuple with the Subnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnets

`func (o *AzureNetworkV4Parameters) SetSubnets(v string)`

SetSubnets sets Subnets field to given value.

### HasSubnets

`func (o *AzureNetworkV4Parameters) HasSubnets() bool`

HasSubnets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


