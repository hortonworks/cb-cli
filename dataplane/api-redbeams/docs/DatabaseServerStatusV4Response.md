# DatabaseServerStatusV4Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnvironmentCrn** | **string** | CRN of the environment of the database server | 
**Name** | **string** | Name of the database server | 
**ResourceCrn** | **string** | CRN of the database server | 
**Status** | **string** | Status of the database server stack | 
**StatusReason** | **string** | Additional status information about the database server stack | 

## Methods

### NewDatabaseServerStatusV4Response

`func NewDatabaseServerStatusV4Response(environmentCrn string, name string, resourceCrn string, status string, statusReason string, ) *DatabaseServerStatusV4Response`

NewDatabaseServerStatusV4Response instantiates a new DatabaseServerStatusV4Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabaseServerStatusV4ResponseWithDefaults

`func NewDatabaseServerStatusV4ResponseWithDefaults() *DatabaseServerStatusV4Response`

NewDatabaseServerStatusV4ResponseWithDefaults instantiates a new DatabaseServerStatusV4Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironmentCrn

`func (o *DatabaseServerStatusV4Response) GetEnvironmentCrn() string`

GetEnvironmentCrn returns the EnvironmentCrn field if non-nil, zero value otherwise.

### GetEnvironmentCrnOk

`func (o *DatabaseServerStatusV4Response) GetEnvironmentCrnOk() (*string, bool)`

GetEnvironmentCrnOk returns a tuple with the EnvironmentCrn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentCrn

`func (o *DatabaseServerStatusV4Response) SetEnvironmentCrn(v string)`

SetEnvironmentCrn sets EnvironmentCrn field to given value.


### GetName

`func (o *DatabaseServerStatusV4Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DatabaseServerStatusV4Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DatabaseServerStatusV4Response) SetName(v string)`

SetName sets Name field to given value.


### GetResourceCrn

`func (o *DatabaseServerStatusV4Response) GetResourceCrn() string`

GetResourceCrn returns the ResourceCrn field if non-nil, zero value otherwise.

### GetResourceCrnOk

`func (o *DatabaseServerStatusV4Response) GetResourceCrnOk() (*string, bool)`

GetResourceCrnOk returns a tuple with the ResourceCrn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceCrn

`func (o *DatabaseServerStatusV4Response) SetResourceCrn(v string)`

SetResourceCrn sets ResourceCrn field to given value.


### GetStatus

`func (o *DatabaseServerStatusV4Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DatabaseServerStatusV4Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DatabaseServerStatusV4Response) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetStatusReason

`func (o *DatabaseServerStatusV4Response) GetStatusReason() string`

GetStatusReason returns the StatusReason field if non-nil, zero value otherwise.

### GetStatusReasonOk

`func (o *DatabaseServerStatusV4Response) GetStatusReasonOk() (*string, bool)`

GetStatusReasonOk returns a tuple with the StatusReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusReason

`func (o *DatabaseServerStatusV4Response) SetStatusReason(v string)`

SetStatusReason sets StatusReason field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


