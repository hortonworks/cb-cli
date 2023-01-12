# AllocateDatabaseServerV4Request

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the database stack | [optional] 
**EnvironmentCrn** | **string** | CRN of the environment of the database server | 
**ClusterCrn** | **string** | CRN of the cluster of the database server | 
**Network** | Pointer to [**NetworkV4StackRequest**](NetworkV4StackRequest.md) |  | [optional] 
**DatabaseServer** | [**DatabaseServerV4StackRequest**](DatabaseServerV4StackRequest.md) |  | 
**Aws** | Pointer to **map[string]interface{}** | AWS-specific parameters for the database stack | [optional] 
**Azure** | Pointer to **map[string]interface{}** |  | [optional] 
**Gcp** | Pointer to **map[string]interface{}** | Azure-specific parameters for the database stack | [optional] 
**SslConfig** | Pointer to [**SslConfigV4Request**](SslConfigV4Request.md) |  | [optional] 
**Tags** | Pointer to **map[string]string** | UserDefined tags for the DB | [optional] 

## Methods

### NewAllocateDatabaseServerV4Request

`func NewAllocateDatabaseServerV4Request(environmentCrn string, clusterCrn string, databaseServer DatabaseServerV4StackRequest, ) *AllocateDatabaseServerV4Request`

NewAllocateDatabaseServerV4Request instantiates a new AllocateDatabaseServerV4Request object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAllocateDatabaseServerV4RequestWithDefaults

`func NewAllocateDatabaseServerV4RequestWithDefaults() *AllocateDatabaseServerV4Request`

NewAllocateDatabaseServerV4RequestWithDefaults instantiates a new AllocateDatabaseServerV4Request object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AllocateDatabaseServerV4Request) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AllocateDatabaseServerV4Request) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AllocateDatabaseServerV4Request) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AllocateDatabaseServerV4Request) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEnvironmentCrn

`func (o *AllocateDatabaseServerV4Request) GetEnvironmentCrn() string`

GetEnvironmentCrn returns the EnvironmentCrn field if non-nil, zero value otherwise.

### GetEnvironmentCrnOk

`func (o *AllocateDatabaseServerV4Request) GetEnvironmentCrnOk() (*string, bool)`

GetEnvironmentCrnOk returns a tuple with the EnvironmentCrn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentCrn

`func (o *AllocateDatabaseServerV4Request) SetEnvironmentCrn(v string)`

SetEnvironmentCrn sets EnvironmentCrn field to given value.


### GetClusterCrn

`func (o *AllocateDatabaseServerV4Request) GetClusterCrn() string`

GetClusterCrn returns the ClusterCrn field if non-nil, zero value otherwise.

### GetClusterCrnOk

`func (o *AllocateDatabaseServerV4Request) GetClusterCrnOk() (*string, bool)`

GetClusterCrnOk returns a tuple with the ClusterCrn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterCrn

`func (o *AllocateDatabaseServerV4Request) SetClusterCrn(v string)`

SetClusterCrn sets ClusterCrn field to given value.


### GetNetwork

`func (o *AllocateDatabaseServerV4Request) GetNetwork() NetworkV4StackRequest`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *AllocateDatabaseServerV4Request) GetNetworkOk() (*NetworkV4StackRequest, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *AllocateDatabaseServerV4Request) SetNetwork(v NetworkV4StackRequest)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *AllocateDatabaseServerV4Request) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetDatabaseServer

`func (o *AllocateDatabaseServerV4Request) GetDatabaseServer() DatabaseServerV4StackRequest`

GetDatabaseServer returns the DatabaseServer field if non-nil, zero value otherwise.

### GetDatabaseServerOk

`func (o *AllocateDatabaseServerV4Request) GetDatabaseServerOk() (*DatabaseServerV4StackRequest, bool)`

GetDatabaseServerOk returns a tuple with the DatabaseServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseServer

`func (o *AllocateDatabaseServerV4Request) SetDatabaseServer(v DatabaseServerV4StackRequest)`

SetDatabaseServer sets DatabaseServer field to given value.


### GetAws

`func (o *AllocateDatabaseServerV4Request) GetAws() map[string]interface{}`

GetAws returns the Aws field if non-nil, zero value otherwise.

### GetAwsOk

`func (o *AllocateDatabaseServerV4Request) GetAwsOk() (*map[string]interface{}, bool)`

GetAwsOk returns a tuple with the Aws field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAws

`func (o *AllocateDatabaseServerV4Request) SetAws(v map[string]interface{})`

SetAws sets Aws field to given value.

### HasAws

`func (o *AllocateDatabaseServerV4Request) HasAws() bool`

HasAws returns a boolean if a field has been set.

### GetAzure

`func (o *AllocateDatabaseServerV4Request) GetAzure() map[string]interface{}`

GetAzure returns the Azure field if non-nil, zero value otherwise.

### GetAzureOk

`func (o *AllocateDatabaseServerV4Request) GetAzureOk() (*map[string]interface{}, bool)`

GetAzureOk returns a tuple with the Azure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzure

`func (o *AllocateDatabaseServerV4Request) SetAzure(v map[string]interface{})`

SetAzure sets Azure field to given value.

### HasAzure

`func (o *AllocateDatabaseServerV4Request) HasAzure() bool`

HasAzure returns a boolean if a field has been set.

### GetGcp

`func (o *AllocateDatabaseServerV4Request) GetGcp() map[string]interface{}`

GetGcp returns the Gcp field if non-nil, zero value otherwise.

### GetGcpOk

`func (o *AllocateDatabaseServerV4Request) GetGcpOk() (*map[string]interface{}, bool)`

GetGcpOk returns a tuple with the Gcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcp

`func (o *AllocateDatabaseServerV4Request) SetGcp(v map[string]interface{})`

SetGcp sets Gcp field to given value.

### HasGcp

`func (o *AllocateDatabaseServerV4Request) HasGcp() bool`

HasGcp returns a boolean if a field has been set.

### GetSslConfig

`func (o *AllocateDatabaseServerV4Request) GetSslConfig() SslConfigV4Request`

GetSslConfig returns the SslConfig field if non-nil, zero value otherwise.

### GetSslConfigOk

`func (o *AllocateDatabaseServerV4Request) GetSslConfigOk() (*SslConfigV4Request, bool)`

GetSslConfigOk returns a tuple with the SslConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslConfig

`func (o *AllocateDatabaseServerV4Request) SetSslConfig(v SslConfigV4Request)`

SetSslConfig sets SslConfig field to given value.

### HasSslConfig

`func (o *AllocateDatabaseServerV4Request) HasSslConfig() bool`

HasSslConfig returns a boolean if a field has been set.

### GetTags

`func (o *AllocateDatabaseServerV4Request) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AllocateDatabaseServerV4Request) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AllocateDatabaseServerV4Request) SetTags(v map[string]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *AllocateDatabaseServerV4Request) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


