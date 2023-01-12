# FlowLogResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceId** | Pointer to **int64** |  | [optional] 
**Created** | Pointer to **int64** |  | [optional] 
**FlowId** | Pointer to **string** |  | [optional] 
**FlowChainId** | Pointer to **string** |  | [optional] 
**NextEvent** | Pointer to **string** |  | [optional] 
**CurrentState** | Pointer to **string** |  | [optional] 
**Finalized** | Pointer to **bool** |  | [optional] 
**NodeId** | Pointer to **string** |  | [optional] 
**StateStatus** | Pointer to **string** |  | [optional] 
**FlowTriggerUserCrn** | Pointer to **string** |  | [optional] 
**EndTime** | Pointer to **int64** |  | [optional] 

## Methods

### NewFlowLogResponse

`func NewFlowLogResponse() *FlowLogResponse`

NewFlowLogResponse instantiates a new FlowLogResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlowLogResponseWithDefaults

`func NewFlowLogResponseWithDefaults() *FlowLogResponse`

NewFlowLogResponseWithDefaults instantiates a new FlowLogResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceId

`func (o *FlowLogResponse) GetResourceId() int64`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *FlowLogResponse) GetResourceIdOk() (*int64, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *FlowLogResponse) SetResourceId(v int64)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *FlowLogResponse) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### GetCreated

`func (o *FlowLogResponse) GetCreated() int64`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *FlowLogResponse) GetCreatedOk() (*int64, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *FlowLogResponse) SetCreated(v int64)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *FlowLogResponse) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetFlowId

`func (o *FlowLogResponse) GetFlowId() string`

GetFlowId returns the FlowId field if non-nil, zero value otherwise.

### GetFlowIdOk

`func (o *FlowLogResponse) GetFlowIdOk() (*string, bool)`

GetFlowIdOk returns a tuple with the FlowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowId

`func (o *FlowLogResponse) SetFlowId(v string)`

SetFlowId sets FlowId field to given value.

### HasFlowId

`func (o *FlowLogResponse) HasFlowId() bool`

HasFlowId returns a boolean if a field has been set.

### GetFlowChainId

`func (o *FlowLogResponse) GetFlowChainId() string`

GetFlowChainId returns the FlowChainId field if non-nil, zero value otherwise.

### GetFlowChainIdOk

`func (o *FlowLogResponse) GetFlowChainIdOk() (*string, bool)`

GetFlowChainIdOk returns a tuple with the FlowChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowChainId

`func (o *FlowLogResponse) SetFlowChainId(v string)`

SetFlowChainId sets FlowChainId field to given value.

### HasFlowChainId

`func (o *FlowLogResponse) HasFlowChainId() bool`

HasFlowChainId returns a boolean if a field has been set.

### GetNextEvent

`func (o *FlowLogResponse) GetNextEvent() string`

GetNextEvent returns the NextEvent field if non-nil, zero value otherwise.

### GetNextEventOk

`func (o *FlowLogResponse) GetNextEventOk() (*string, bool)`

GetNextEventOk returns a tuple with the NextEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextEvent

`func (o *FlowLogResponse) SetNextEvent(v string)`

SetNextEvent sets NextEvent field to given value.

### HasNextEvent

`func (o *FlowLogResponse) HasNextEvent() bool`

HasNextEvent returns a boolean if a field has been set.

### GetCurrentState

`func (o *FlowLogResponse) GetCurrentState() string`

GetCurrentState returns the CurrentState field if non-nil, zero value otherwise.

### GetCurrentStateOk

`func (o *FlowLogResponse) GetCurrentStateOk() (*string, bool)`

GetCurrentStateOk returns a tuple with the CurrentState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentState

`func (o *FlowLogResponse) SetCurrentState(v string)`

SetCurrentState sets CurrentState field to given value.

### HasCurrentState

`func (o *FlowLogResponse) HasCurrentState() bool`

HasCurrentState returns a boolean if a field has been set.

### GetFinalized

`func (o *FlowLogResponse) GetFinalized() bool`

GetFinalized returns the Finalized field if non-nil, zero value otherwise.

### GetFinalizedOk

`func (o *FlowLogResponse) GetFinalizedOk() (*bool, bool)`

GetFinalizedOk returns a tuple with the Finalized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalized

`func (o *FlowLogResponse) SetFinalized(v bool)`

SetFinalized sets Finalized field to given value.

### HasFinalized

`func (o *FlowLogResponse) HasFinalized() bool`

HasFinalized returns a boolean if a field has been set.

### GetNodeId

`func (o *FlowLogResponse) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *FlowLogResponse) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *FlowLogResponse) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.

### HasNodeId

`func (o *FlowLogResponse) HasNodeId() bool`

HasNodeId returns a boolean if a field has been set.

### GetStateStatus

`func (o *FlowLogResponse) GetStateStatus() string`

GetStateStatus returns the StateStatus field if non-nil, zero value otherwise.

### GetStateStatusOk

`func (o *FlowLogResponse) GetStateStatusOk() (*string, bool)`

GetStateStatusOk returns a tuple with the StateStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateStatus

`func (o *FlowLogResponse) SetStateStatus(v string)`

SetStateStatus sets StateStatus field to given value.

### HasStateStatus

`func (o *FlowLogResponse) HasStateStatus() bool`

HasStateStatus returns a boolean if a field has been set.

### GetFlowTriggerUserCrn

`func (o *FlowLogResponse) GetFlowTriggerUserCrn() string`

GetFlowTriggerUserCrn returns the FlowTriggerUserCrn field if non-nil, zero value otherwise.

### GetFlowTriggerUserCrnOk

`func (o *FlowLogResponse) GetFlowTriggerUserCrnOk() (*string, bool)`

GetFlowTriggerUserCrnOk returns a tuple with the FlowTriggerUserCrn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowTriggerUserCrn

`func (o *FlowLogResponse) SetFlowTriggerUserCrn(v string)`

SetFlowTriggerUserCrn sets FlowTriggerUserCrn field to given value.

### HasFlowTriggerUserCrn

`func (o *FlowLogResponse) HasFlowTriggerUserCrn() bool`

HasFlowTriggerUserCrn returns a boolean if a field has been set.

### GetEndTime

`func (o *FlowLogResponse) GetEndTime() int64`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *FlowLogResponse) GetEndTimeOk() (*int64, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *FlowLogResponse) SetEndTime(v int64)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *FlowLogResponse) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


