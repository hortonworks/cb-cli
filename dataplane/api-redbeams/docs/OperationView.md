# OperationView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OperationId** | Pointer to **string** |  | [optional] 
**OperationType** | Pointer to **string** |  | [optional] 
**OperationResource** | Pointer to **string** |  | [optional] 
**Operations** | Pointer to [**[]FlowProgressResponse**](FlowProgressResponse.md) |  | [optional] 
**SubOperations** | Pointer to [**map[string]OperationView**](OperationView.md) |  | [optional] 
**SubOperationConditions** | Pointer to **map[string]string** |  | [optional] 
**ProgressStatus** | Pointer to **string** |  | [optional] 
**Progress** | Pointer to **int32** |  | [optional] 

## Methods

### NewOperationView

`func NewOperationView() *OperationView`

NewOperationView instantiates a new OperationView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOperationViewWithDefaults

`func NewOperationViewWithDefaults() *OperationView`

NewOperationViewWithDefaults instantiates a new OperationView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperationId

`func (o *OperationView) GetOperationId() string`

GetOperationId returns the OperationId field if non-nil, zero value otherwise.

### GetOperationIdOk

`func (o *OperationView) GetOperationIdOk() (*string, bool)`

GetOperationIdOk returns a tuple with the OperationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationId

`func (o *OperationView) SetOperationId(v string)`

SetOperationId sets OperationId field to given value.

### HasOperationId

`func (o *OperationView) HasOperationId() bool`

HasOperationId returns a boolean if a field has been set.

### GetOperationType

`func (o *OperationView) GetOperationType() string`

GetOperationType returns the OperationType field if non-nil, zero value otherwise.

### GetOperationTypeOk

`func (o *OperationView) GetOperationTypeOk() (*string, bool)`

GetOperationTypeOk returns a tuple with the OperationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationType

`func (o *OperationView) SetOperationType(v string)`

SetOperationType sets OperationType field to given value.

### HasOperationType

`func (o *OperationView) HasOperationType() bool`

HasOperationType returns a boolean if a field has been set.

### GetOperationResource

`func (o *OperationView) GetOperationResource() string`

GetOperationResource returns the OperationResource field if non-nil, zero value otherwise.

### GetOperationResourceOk

`func (o *OperationView) GetOperationResourceOk() (*string, bool)`

GetOperationResourceOk returns a tuple with the OperationResource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationResource

`func (o *OperationView) SetOperationResource(v string)`

SetOperationResource sets OperationResource field to given value.

### HasOperationResource

`func (o *OperationView) HasOperationResource() bool`

HasOperationResource returns a boolean if a field has been set.

### GetOperations

`func (o *OperationView) GetOperations() []FlowProgressResponse`

GetOperations returns the Operations field if non-nil, zero value otherwise.

### GetOperationsOk

`func (o *OperationView) GetOperationsOk() (*[]FlowProgressResponse, bool)`

GetOperationsOk returns a tuple with the Operations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperations

`func (o *OperationView) SetOperations(v []FlowProgressResponse)`

SetOperations sets Operations field to given value.

### HasOperations

`func (o *OperationView) HasOperations() bool`

HasOperations returns a boolean if a field has been set.

### GetSubOperations

`func (o *OperationView) GetSubOperations() map[string]OperationView`

GetSubOperations returns the SubOperations field if non-nil, zero value otherwise.

### GetSubOperationsOk

`func (o *OperationView) GetSubOperationsOk() (*map[string]OperationView, bool)`

GetSubOperationsOk returns a tuple with the SubOperations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubOperations

`func (o *OperationView) SetSubOperations(v map[string]OperationView)`

SetSubOperations sets SubOperations field to given value.

### HasSubOperations

`func (o *OperationView) HasSubOperations() bool`

HasSubOperations returns a boolean if a field has been set.

### GetSubOperationConditions

`func (o *OperationView) GetSubOperationConditions() map[string]string`

GetSubOperationConditions returns the SubOperationConditions field if non-nil, zero value otherwise.

### GetSubOperationConditionsOk

`func (o *OperationView) GetSubOperationConditionsOk() (*map[string]string, bool)`

GetSubOperationConditionsOk returns a tuple with the SubOperationConditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubOperationConditions

`func (o *OperationView) SetSubOperationConditions(v map[string]string)`

SetSubOperationConditions sets SubOperationConditions field to given value.

### HasSubOperationConditions

`func (o *OperationView) HasSubOperationConditions() bool`

HasSubOperationConditions returns a boolean if a field has been set.

### GetProgressStatus

`func (o *OperationView) GetProgressStatus() string`

GetProgressStatus returns the ProgressStatus field if non-nil, zero value otherwise.

### GetProgressStatusOk

`func (o *OperationView) GetProgressStatusOk() (*string, bool)`

GetProgressStatusOk returns a tuple with the ProgressStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgressStatus

`func (o *OperationView) SetProgressStatus(v string)`

SetProgressStatus sets ProgressStatus field to given value.

### HasProgressStatus

`func (o *OperationView) HasProgressStatus() bool`

HasProgressStatus returns a boolean if a field has been set.

### GetProgress

`func (o *OperationView) GetProgress() int32`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *OperationView) GetProgressOk() (*int32, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *OperationView) SetProgress(v int32)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *OperationView) HasProgress() bool`

HasProgress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


