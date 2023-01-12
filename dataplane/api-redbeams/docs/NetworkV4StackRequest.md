# NetworkV4StackRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aws** | Pointer to [**AwsNetworkV4Parameters**](AwsNetworkV4Parameters.md) |  | [optional] 
**Azure** | Pointer to [**AzureNetworkV4Parameters**](AzureNetworkV4Parameters.md) |  | [optional] 
**Gcp** | Pointer to [**GcpNetworkV4Parameters**](GcpNetworkV4Parameters.md) |  | [optional] 

## Methods

### NewNetworkV4StackRequest

`func NewNetworkV4StackRequest() *NetworkV4StackRequest`

NewNetworkV4StackRequest instantiates a new NetworkV4StackRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkV4StackRequestWithDefaults

`func NewNetworkV4StackRequestWithDefaults() *NetworkV4StackRequest`

NewNetworkV4StackRequestWithDefaults instantiates a new NetworkV4StackRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAws

`func (o *NetworkV4StackRequest) GetAws() AwsNetworkV4Parameters`

GetAws returns the Aws field if non-nil, zero value otherwise.

### GetAwsOk

`func (o *NetworkV4StackRequest) GetAwsOk() (*AwsNetworkV4Parameters, bool)`

GetAwsOk returns a tuple with the Aws field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAws

`func (o *NetworkV4StackRequest) SetAws(v AwsNetworkV4Parameters)`

SetAws sets Aws field to given value.

### HasAws

`func (o *NetworkV4StackRequest) HasAws() bool`

HasAws returns a boolean if a field has been set.

### GetAzure

`func (o *NetworkV4StackRequest) GetAzure() AzureNetworkV4Parameters`

GetAzure returns the Azure field if non-nil, zero value otherwise.

### GetAzureOk

`func (o *NetworkV4StackRequest) GetAzureOk() (*AzureNetworkV4Parameters, bool)`

GetAzureOk returns a tuple with the Azure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzure

`func (o *NetworkV4StackRequest) SetAzure(v AzureNetworkV4Parameters)`

SetAzure sets Azure field to given value.

### HasAzure

`func (o *NetworkV4StackRequest) HasAzure() bool`

HasAzure returns a boolean if a field has been set.

### GetGcp

`func (o *NetworkV4StackRequest) GetGcp() GcpNetworkV4Parameters`

GetGcp returns the Gcp field if non-nil, zero value otherwise.

### GetGcpOk

`func (o *NetworkV4StackRequest) GetGcpOk() (*GcpNetworkV4Parameters, bool)`

GetGcpOk returns a tuple with the Gcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcp

`func (o *NetworkV4StackRequest) SetGcp(v GcpNetworkV4Parameters)`

SetGcp sets Gcp field to given value.

### HasGcp

`func (o *NetworkV4StackRequest) HasGcp() bool`

HasGcp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


