# Pageable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sort** | Pointer to [**Sort**](Sort.md) |  | [optional] 
**Offset** | Pointer to **int64** |  | [optional] 
**Paged** | Pointer to **bool** |  | [optional] 
**Unpaged** | Pointer to **bool** |  | [optional] 
**PageNumber** | Pointer to **int32** |  | [optional] 
**PageSize** | Pointer to **int32** |  | [optional] 

## Methods

### NewPageable

`func NewPageable() *Pageable`

NewPageable instantiates a new Pageable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPageableWithDefaults

`func NewPageableWithDefaults() *Pageable`

NewPageableWithDefaults instantiates a new Pageable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSort

`func (o *Pageable) GetSort() Sort`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *Pageable) GetSortOk() (*Sort, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *Pageable) SetSort(v Sort)`

SetSort sets Sort field to given value.

### HasSort

`func (o *Pageable) HasSort() bool`

HasSort returns a boolean if a field has been set.

### GetOffset

`func (o *Pageable) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *Pageable) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *Pageable) SetOffset(v int64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *Pageable) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetPaged

`func (o *Pageable) GetPaged() bool`

GetPaged returns the Paged field if non-nil, zero value otherwise.

### GetPagedOk

`func (o *Pageable) GetPagedOk() (*bool, bool)`

GetPagedOk returns a tuple with the Paged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaged

`func (o *Pageable) SetPaged(v bool)`

SetPaged sets Paged field to given value.

### HasPaged

`func (o *Pageable) HasPaged() bool`

HasPaged returns a boolean if a field has been set.

### GetUnpaged

`func (o *Pageable) GetUnpaged() bool`

GetUnpaged returns the Unpaged field if non-nil, zero value otherwise.

### GetUnpagedOk

`func (o *Pageable) GetUnpagedOk() (*bool, bool)`

GetUnpagedOk returns a tuple with the Unpaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnpaged

`func (o *Pageable) SetUnpaged(v bool)`

SetUnpaged sets Unpaged field to given value.

### HasUnpaged

`func (o *Pageable) HasUnpaged() bool`

HasUnpaged returns a boolean if a field has been set.

### GetPageNumber

`func (o *Pageable) GetPageNumber() int32`

GetPageNumber returns the PageNumber field if non-nil, zero value otherwise.

### GetPageNumberOk

`func (o *Pageable) GetPageNumberOk() (*int32, bool)`

GetPageNumberOk returns a tuple with the PageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageNumber

`func (o *Pageable) SetPageNumber(v int32)`

SetPageNumber sets PageNumber field to given value.

### HasPageNumber

`func (o *Pageable) HasPageNumber() bool`

HasPageNumber returns a boolean if a field has been set.

### GetPageSize

`func (o *Pageable) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *Pageable) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *Pageable) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *Pageable) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


