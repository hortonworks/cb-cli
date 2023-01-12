# FlowStateTransitionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | Pointer to **string** |  | [optional] 
**NextEvent** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**ElapsedTimeInSeconds** | Pointer to **float64** |  | [optional] 

## Methods

### NewFlowStateTransitionResponse

`func NewFlowStateTransitionResponse() *FlowStateTransitionResponse`

NewFlowStateTransitionResponse instantiates a new FlowStateTransitionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlowStateTransitionResponseWithDefaults

`func NewFlowStateTransitionResponseWithDefaults() *FlowStateTransitionResponse`

NewFlowStateTransitionResponseWithDefaults instantiates a new FlowStateTransitionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *FlowStateTransitionResponse) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *FlowStateTransitionResponse) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *FlowStateTransitionResponse) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *FlowStateTransitionResponse) HasState() bool`

HasState returns a boolean if a field has been set.

### GetNextEvent

`func (o *FlowStateTransitionResponse) GetNextEvent() string`

GetNextEvent returns the NextEvent field if non-nil, zero value otherwise.

### GetNextEventOk

`func (o *FlowStateTransitionResponse) GetNextEventOk() (*string, bool)`

GetNextEventOk returns a tuple with the NextEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextEvent

`func (o *FlowStateTransitionResponse) SetNextEvent(v string)`

SetNextEvent sets NextEvent field to given value.

### HasNextEvent

`func (o *FlowStateTransitionResponse) HasNextEvent() bool`

HasNextEvent returns a boolean if a field has been set.

### GetStatus

`func (o *FlowStateTransitionResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FlowStateTransitionResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FlowStateTransitionResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *FlowStateTransitionResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetElapsedTimeInSeconds

`func (o *FlowStateTransitionResponse) GetElapsedTimeInSeconds() float64`

GetElapsedTimeInSeconds returns the ElapsedTimeInSeconds field if non-nil, zero value otherwise.

### GetElapsedTimeInSecondsOk

`func (o *FlowStateTransitionResponse) GetElapsedTimeInSecondsOk() (*float64, bool)`

GetElapsedTimeInSecondsOk returns a tuple with the ElapsedTimeInSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElapsedTimeInSeconds

`func (o *FlowStateTransitionResponse) SetElapsedTimeInSeconds(v float64)`

SetElapsedTimeInSeconds sets ElapsedTimeInSeconds field to given value.

### HasElapsedTimeInSeconds

`func (o *FlowStateTransitionResponse) HasElapsedTimeInSeconds() bool`

HasElapsedTimeInSeconds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


