# FlowProgressResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FlowId** | Pointer to **string** |  | [optional] 
**FlowChainId** | Pointer to **string** |  | [optional] 
**ResourceCrn** | Pointer to **string** |  | [optional] 
**Created** | Pointer to **int64** |  | [optional] 
**Transitions** | Pointer to [**[]FlowStateTransitionResponse**](FlowStateTransitionResponse.md) |  | [optional] 
**Progress** | Pointer to **int32** |  | [optional] 
**ElapsedTimeInSeconds** | Pointer to **float64** |  | [optional] 
**Finalized** | Pointer to **bool** |  | [optional] 
**MaxNumberOfTransitions** | Pointer to **int32** |  | [optional] 

## Methods

### NewFlowProgressResponse

`func NewFlowProgressResponse() *FlowProgressResponse`

NewFlowProgressResponse instantiates a new FlowProgressResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlowProgressResponseWithDefaults

`func NewFlowProgressResponseWithDefaults() *FlowProgressResponse`

NewFlowProgressResponseWithDefaults instantiates a new FlowProgressResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFlowId

`func (o *FlowProgressResponse) GetFlowId() string`

GetFlowId returns the FlowId field if non-nil, zero value otherwise.

### GetFlowIdOk

`func (o *FlowProgressResponse) GetFlowIdOk() (*string, bool)`

GetFlowIdOk returns a tuple with the FlowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowId

`func (o *FlowProgressResponse) SetFlowId(v string)`

SetFlowId sets FlowId field to given value.

### HasFlowId

`func (o *FlowProgressResponse) HasFlowId() bool`

HasFlowId returns a boolean if a field has been set.

### GetFlowChainId

`func (o *FlowProgressResponse) GetFlowChainId() string`

GetFlowChainId returns the FlowChainId field if non-nil, zero value otherwise.

### GetFlowChainIdOk

`func (o *FlowProgressResponse) GetFlowChainIdOk() (*string, bool)`

GetFlowChainIdOk returns a tuple with the FlowChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowChainId

`func (o *FlowProgressResponse) SetFlowChainId(v string)`

SetFlowChainId sets FlowChainId field to given value.

### HasFlowChainId

`func (o *FlowProgressResponse) HasFlowChainId() bool`

HasFlowChainId returns a boolean if a field has been set.

### GetResourceCrn

`func (o *FlowProgressResponse) GetResourceCrn() string`

GetResourceCrn returns the ResourceCrn field if non-nil, zero value otherwise.

### GetResourceCrnOk

`func (o *FlowProgressResponse) GetResourceCrnOk() (*string, bool)`

GetResourceCrnOk returns a tuple with the ResourceCrn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceCrn

`func (o *FlowProgressResponse) SetResourceCrn(v string)`

SetResourceCrn sets ResourceCrn field to given value.

### HasResourceCrn

`func (o *FlowProgressResponse) HasResourceCrn() bool`

HasResourceCrn returns a boolean if a field has been set.

### GetCreated

`func (o *FlowProgressResponse) GetCreated() int64`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *FlowProgressResponse) GetCreatedOk() (*int64, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *FlowProgressResponse) SetCreated(v int64)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *FlowProgressResponse) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetTransitions

`func (o *FlowProgressResponse) GetTransitions() []FlowStateTransitionResponse`

GetTransitions returns the Transitions field if non-nil, zero value otherwise.

### GetTransitionsOk

`func (o *FlowProgressResponse) GetTransitionsOk() (*[]FlowStateTransitionResponse, bool)`

GetTransitionsOk returns a tuple with the Transitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitions

`func (o *FlowProgressResponse) SetTransitions(v []FlowStateTransitionResponse)`

SetTransitions sets Transitions field to given value.

### HasTransitions

`func (o *FlowProgressResponse) HasTransitions() bool`

HasTransitions returns a boolean if a field has been set.

### GetProgress

`func (o *FlowProgressResponse) GetProgress() int32`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *FlowProgressResponse) GetProgressOk() (*int32, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *FlowProgressResponse) SetProgress(v int32)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *FlowProgressResponse) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### GetElapsedTimeInSeconds

`func (o *FlowProgressResponse) GetElapsedTimeInSeconds() float64`

GetElapsedTimeInSeconds returns the ElapsedTimeInSeconds field if non-nil, zero value otherwise.

### GetElapsedTimeInSecondsOk

`func (o *FlowProgressResponse) GetElapsedTimeInSecondsOk() (*float64, bool)`

GetElapsedTimeInSecondsOk returns a tuple with the ElapsedTimeInSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElapsedTimeInSeconds

`func (o *FlowProgressResponse) SetElapsedTimeInSeconds(v float64)`

SetElapsedTimeInSeconds sets ElapsedTimeInSeconds field to given value.

### HasElapsedTimeInSeconds

`func (o *FlowProgressResponse) HasElapsedTimeInSeconds() bool`

HasElapsedTimeInSeconds returns a boolean if a field has been set.

### GetFinalized

`func (o *FlowProgressResponse) GetFinalized() bool`

GetFinalized returns the Finalized field if non-nil, zero value otherwise.

### GetFinalizedOk

`func (o *FlowProgressResponse) GetFinalizedOk() (*bool, bool)`

GetFinalizedOk returns a tuple with the Finalized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalized

`func (o *FlowProgressResponse) SetFinalized(v bool)`

SetFinalized sets Finalized field to given value.

### HasFinalized

`func (o *FlowProgressResponse) HasFinalized() bool`

HasFinalized returns a boolean if a field has been set.

### GetMaxNumberOfTransitions

`func (o *FlowProgressResponse) GetMaxNumberOfTransitions() int32`

GetMaxNumberOfTransitions returns the MaxNumberOfTransitions field if non-nil, zero value otherwise.

### GetMaxNumberOfTransitionsOk

`func (o *FlowProgressResponse) GetMaxNumberOfTransitionsOk() (*int32, bool)`

GetMaxNumberOfTransitionsOk returns a tuple with the MaxNumberOfTransitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNumberOfTransitions

`func (o *FlowProgressResponse) SetMaxNumberOfTransitions(v int32)`

SetMaxNumberOfTransitions sets MaxNumberOfTransitions field to given value.

### HasMaxNumberOfTransitions

`func (o *FlowProgressResponse) HasMaxNumberOfTransitions() bool`

HasMaxNumberOfTransitions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


