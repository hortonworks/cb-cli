# FlowCheckResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FlowId** | Pointer to **string** |  | [optional] 
**FlowChainId** | Pointer to **string** |  | [optional] 
**HasActiveFlow** | Pointer to **bool** |  | [optional] 
**LatestFlowFinalizedAndFailed** | Pointer to **bool** |  | [optional] 
**EndTime** | Pointer to **int64** |  | [optional] 

## Methods

### NewFlowCheckResponse

`func NewFlowCheckResponse() *FlowCheckResponse`

NewFlowCheckResponse instantiates a new FlowCheckResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlowCheckResponseWithDefaults

`func NewFlowCheckResponseWithDefaults() *FlowCheckResponse`

NewFlowCheckResponseWithDefaults instantiates a new FlowCheckResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFlowId

`func (o *FlowCheckResponse) GetFlowId() string`

GetFlowId returns the FlowId field if non-nil, zero value otherwise.

### GetFlowIdOk

`func (o *FlowCheckResponse) GetFlowIdOk() (*string, bool)`

GetFlowIdOk returns a tuple with the FlowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowId

`func (o *FlowCheckResponse) SetFlowId(v string)`

SetFlowId sets FlowId field to given value.

### HasFlowId

`func (o *FlowCheckResponse) HasFlowId() bool`

HasFlowId returns a boolean if a field has been set.

### GetFlowChainId

`func (o *FlowCheckResponse) GetFlowChainId() string`

GetFlowChainId returns the FlowChainId field if non-nil, zero value otherwise.

### GetFlowChainIdOk

`func (o *FlowCheckResponse) GetFlowChainIdOk() (*string, bool)`

GetFlowChainIdOk returns a tuple with the FlowChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowChainId

`func (o *FlowCheckResponse) SetFlowChainId(v string)`

SetFlowChainId sets FlowChainId field to given value.

### HasFlowChainId

`func (o *FlowCheckResponse) HasFlowChainId() bool`

HasFlowChainId returns a boolean if a field has been set.

### GetHasActiveFlow

`func (o *FlowCheckResponse) GetHasActiveFlow() bool`

GetHasActiveFlow returns the HasActiveFlow field if non-nil, zero value otherwise.

### GetHasActiveFlowOk

`func (o *FlowCheckResponse) GetHasActiveFlowOk() (*bool, bool)`

GetHasActiveFlowOk returns a tuple with the HasActiveFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasActiveFlow

`func (o *FlowCheckResponse) SetHasActiveFlow(v bool)`

SetHasActiveFlow sets HasActiveFlow field to given value.

### HasHasActiveFlow

`func (o *FlowCheckResponse) HasHasActiveFlow() bool`

HasHasActiveFlow returns a boolean if a field has been set.

### GetLatestFlowFinalizedAndFailed

`func (o *FlowCheckResponse) GetLatestFlowFinalizedAndFailed() bool`

GetLatestFlowFinalizedAndFailed returns the LatestFlowFinalizedAndFailed field if non-nil, zero value otherwise.

### GetLatestFlowFinalizedAndFailedOk

`func (o *FlowCheckResponse) GetLatestFlowFinalizedAndFailedOk() (*bool, bool)`

GetLatestFlowFinalizedAndFailedOk returns a tuple with the LatestFlowFinalizedAndFailed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestFlowFinalizedAndFailed

`func (o *FlowCheckResponse) SetLatestFlowFinalizedAndFailed(v bool)`

SetLatestFlowFinalizedAndFailed sets LatestFlowFinalizedAndFailed field to given value.

### HasLatestFlowFinalizedAndFailed

`func (o *FlowCheckResponse) HasLatestFlowFinalizedAndFailed() bool`

HasLatestFlowFinalizedAndFailed returns a boolean if a field has been set.

### GetEndTime

`func (o *FlowCheckResponse) GetEndTime() int64`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *FlowCheckResponse) GetEndTimeOk() (*int64, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *FlowCheckResponse) SetEndTime(v int64)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *FlowCheckResponse) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


