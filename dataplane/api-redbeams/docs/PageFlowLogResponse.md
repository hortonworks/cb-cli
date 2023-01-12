# PageFlowLogResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalPages** | Pointer to **int32** |  | [optional] 
**TotalElements** | Pointer to **int64** |  | [optional] 
**Sort** | Pointer to [**Sort**](Sort.md) |  | [optional] 
**NumberOfElements** | Pointer to **int32** |  | [optional] 
**Size** | Pointer to **int32** |  | [optional] 
**Content** | Pointer to [**[]FlowLogResponse**](FlowLogResponse.md) |  | [optional] 
**Number** | Pointer to **int32** |  | [optional] 
**First** | Pointer to **bool** |  | [optional] 
**Last** | Pointer to **bool** |  | [optional] 
**Pageable** | Pointer to [**Pageable**](Pageable.md) |  | [optional] 
**Empty** | Pointer to **bool** |  | [optional] 

## Methods

### NewPageFlowLogResponse

`func NewPageFlowLogResponse() *PageFlowLogResponse`

NewPageFlowLogResponse instantiates a new PageFlowLogResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPageFlowLogResponseWithDefaults

`func NewPageFlowLogResponseWithDefaults() *PageFlowLogResponse`

NewPageFlowLogResponseWithDefaults instantiates a new PageFlowLogResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalPages

`func (o *PageFlowLogResponse) GetTotalPages() int32`

GetTotalPages returns the TotalPages field if non-nil, zero value otherwise.

### GetTotalPagesOk

`func (o *PageFlowLogResponse) GetTotalPagesOk() (*int32, bool)`

GetTotalPagesOk returns a tuple with the TotalPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPages

`func (o *PageFlowLogResponse) SetTotalPages(v int32)`

SetTotalPages sets TotalPages field to given value.

### HasTotalPages

`func (o *PageFlowLogResponse) HasTotalPages() bool`

HasTotalPages returns a boolean if a field has been set.

### GetTotalElements

`func (o *PageFlowLogResponse) GetTotalElements() int64`

GetTotalElements returns the TotalElements field if non-nil, zero value otherwise.

### GetTotalElementsOk

`func (o *PageFlowLogResponse) GetTotalElementsOk() (*int64, bool)`

GetTotalElementsOk returns a tuple with the TotalElements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalElements

`func (o *PageFlowLogResponse) SetTotalElements(v int64)`

SetTotalElements sets TotalElements field to given value.

### HasTotalElements

`func (o *PageFlowLogResponse) HasTotalElements() bool`

HasTotalElements returns a boolean if a field has been set.

### GetSort

`func (o *PageFlowLogResponse) GetSort() Sort`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *PageFlowLogResponse) GetSortOk() (*Sort, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *PageFlowLogResponse) SetSort(v Sort)`

SetSort sets Sort field to given value.

### HasSort

`func (o *PageFlowLogResponse) HasSort() bool`

HasSort returns a boolean if a field has been set.

### GetNumberOfElements

`func (o *PageFlowLogResponse) GetNumberOfElements() int32`

GetNumberOfElements returns the NumberOfElements field if non-nil, zero value otherwise.

### GetNumberOfElementsOk

`func (o *PageFlowLogResponse) GetNumberOfElementsOk() (*int32, bool)`

GetNumberOfElementsOk returns a tuple with the NumberOfElements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfElements

`func (o *PageFlowLogResponse) SetNumberOfElements(v int32)`

SetNumberOfElements sets NumberOfElements field to given value.

### HasNumberOfElements

`func (o *PageFlowLogResponse) HasNumberOfElements() bool`

HasNumberOfElements returns a boolean if a field has been set.

### GetSize

`func (o *PageFlowLogResponse) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *PageFlowLogResponse) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *PageFlowLogResponse) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *PageFlowLogResponse) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetContent

`func (o *PageFlowLogResponse) GetContent() []FlowLogResponse`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *PageFlowLogResponse) GetContentOk() (*[]FlowLogResponse, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *PageFlowLogResponse) SetContent(v []FlowLogResponse)`

SetContent sets Content field to given value.

### HasContent

`func (o *PageFlowLogResponse) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetNumber

`func (o *PageFlowLogResponse) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *PageFlowLogResponse) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *PageFlowLogResponse) SetNumber(v int32)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *PageFlowLogResponse) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetFirst

`func (o *PageFlowLogResponse) GetFirst() bool`

GetFirst returns the First field if non-nil, zero value otherwise.

### GetFirstOk

`func (o *PageFlowLogResponse) GetFirstOk() (*bool, bool)`

GetFirstOk returns a tuple with the First field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirst

`func (o *PageFlowLogResponse) SetFirst(v bool)`

SetFirst sets First field to given value.

### HasFirst

`func (o *PageFlowLogResponse) HasFirst() bool`

HasFirst returns a boolean if a field has been set.

### GetLast

`func (o *PageFlowLogResponse) GetLast() bool`

GetLast returns the Last field if non-nil, zero value otherwise.

### GetLastOk

`func (o *PageFlowLogResponse) GetLastOk() (*bool, bool)`

GetLastOk returns a tuple with the Last field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast

`func (o *PageFlowLogResponse) SetLast(v bool)`

SetLast sets Last field to given value.

### HasLast

`func (o *PageFlowLogResponse) HasLast() bool`

HasLast returns a boolean if a field has been set.

### GetPageable

`func (o *PageFlowLogResponse) GetPageable() Pageable`

GetPageable returns the Pageable field if non-nil, zero value otherwise.

### GetPageableOk

`func (o *PageFlowLogResponse) GetPageableOk() (*Pageable, bool)`

GetPageableOk returns a tuple with the Pageable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageable

`func (o *PageFlowLogResponse) SetPageable(v Pageable)`

SetPageable sets Pageable field to given value.

### HasPageable

`func (o *PageFlowLogResponse) HasPageable() bool`

HasPageable returns a boolean if a field has been set.

### GetEmpty

`func (o *PageFlowLogResponse) GetEmpty() bool`

GetEmpty returns the Empty field if non-nil, zero value otherwise.

### GetEmptyOk

`func (o *PageFlowLogResponse) GetEmptyOk() (*bool, bool)`

GetEmptyOk returns a tuple with the Empty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmpty

`func (o *PageFlowLogResponse) SetEmpty(v bool)`

SetEmpty sets Empty field to given value.

### HasEmpty

`func (o *PageFlowLogResponse) HasEmpty() bool`

HasEmpty returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


