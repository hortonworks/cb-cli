# ApiAuthorizationInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | Pointer to **string** |  | [optional] 
**HttpMethod** | Pointer to **string** |  | [optional] 
**NewAuthorization** | Pointer to [**NewAuthorizationInfo**](NewAuthorizationInfo.md) |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 

## Methods

### NewApiAuthorizationInfo

`func NewApiAuthorizationInfo() *ApiAuthorizationInfo`

NewApiAuthorizationInfo instantiates a new ApiAuthorizationInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAuthorizationInfoWithDefaults

`func NewApiAuthorizationInfoWithDefaults() *ApiAuthorizationInfo`

NewApiAuthorizationInfoWithDefaults instantiates a new ApiAuthorizationInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *ApiAuthorizationInfo) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ApiAuthorizationInfo) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ApiAuthorizationInfo) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *ApiAuthorizationInfo) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetHttpMethod

`func (o *ApiAuthorizationInfo) GetHttpMethod() string`

GetHttpMethod returns the HttpMethod field if non-nil, zero value otherwise.

### GetHttpMethodOk

`func (o *ApiAuthorizationInfo) GetHttpMethodOk() (*string, bool)`

GetHttpMethodOk returns a tuple with the HttpMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpMethod

`func (o *ApiAuthorizationInfo) SetHttpMethod(v string)`

SetHttpMethod sets HttpMethod field to given value.

### HasHttpMethod

`func (o *ApiAuthorizationInfo) HasHttpMethod() bool`

HasHttpMethod returns a boolean if a field has been set.

### GetNewAuthorization

`func (o *ApiAuthorizationInfo) GetNewAuthorization() NewAuthorizationInfo`

GetNewAuthorization returns the NewAuthorization field if non-nil, zero value otherwise.

### GetNewAuthorizationOk

`func (o *ApiAuthorizationInfo) GetNewAuthorizationOk() (*NewAuthorizationInfo, bool)`

GetNewAuthorizationOk returns a tuple with the NewAuthorization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewAuthorization

`func (o *ApiAuthorizationInfo) SetNewAuthorization(v NewAuthorizationInfo)`

SetNewAuthorization sets NewAuthorization field to given value.

### HasNewAuthorization

`func (o *ApiAuthorizationInfo) HasNewAuthorization() bool`

HasNewAuthorization returns a boolean if a field has been set.

### GetMessage

`func (o *ApiAuthorizationInfo) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ApiAuthorizationInfo) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ApiAuthorizationInfo) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ApiAuthorizationInfo) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


