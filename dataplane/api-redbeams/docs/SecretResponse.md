# SecretResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnginePath** | Pointer to **string** | Engine path of the secret. | [optional] 
**SecretPath** | Pointer to **string** | Path of the secret. | [optional] 

## Methods

### NewSecretResponse

`func NewSecretResponse() *SecretResponse`

NewSecretResponse instantiates a new SecretResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecretResponseWithDefaults

`func NewSecretResponseWithDefaults() *SecretResponse`

NewSecretResponseWithDefaults instantiates a new SecretResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnginePath

`func (o *SecretResponse) GetEnginePath() string`

GetEnginePath returns the EnginePath field if non-nil, zero value otherwise.

### GetEnginePathOk

`func (o *SecretResponse) GetEnginePathOk() (*string, bool)`

GetEnginePathOk returns a tuple with the EnginePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnginePath

`func (o *SecretResponse) SetEnginePath(v string)`

SetEnginePath sets EnginePath field to given value.

### HasEnginePath

`func (o *SecretResponse) HasEnginePath() bool`

HasEnginePath returns a boolean if a field has been set.

### GetSecretPath

`func (o *SecretResponse) GetSecretPath() string`

GetSecretPath returns the SecretPath field if non-nil, zero value otherwise.

### GetSecretPathOk

`func (o *SecretResponse) GetSecretPathOk() (*string, bool)`

GetSecretPathOk returns a tuple with the SecretPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretPath

`func (o *SecretResponse) SetSecretPath(v string)`

SetSecretPath sets SecretPath field to given value.

### HasSecretPath

`func (o *SecretResponse) HasSecretPath() bool`

HasSecretPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


