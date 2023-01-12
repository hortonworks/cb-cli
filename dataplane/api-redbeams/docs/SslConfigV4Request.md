# SslConfigV4Request

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SslMode** | Pointer to **string** | SSL enforcement mode for the database server | [optional] 

## Methods

### NewSslConfigV4Request

`func NewSslConfigV4Request() *SslConfigV4Request`

NewSslConfigV4Request instantiates a new SslConfigV4Request object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSslConfigV4RequestWithDefaults

`func NewSslConfigV4RequestWithDefaults() *SslConfigV4Request`

NewSslConfigV4RequestWithDefaults instantiates a new SslConfigV4Request object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSslMode

`func (o *SslConfigV4Request) GetSslMode() string`

GetSslMode returns the SslMode field if non-nil, zero value otherwise.

### GetSslModeOk

`func (o *SslConfigV4Request) GetSslModeOk() (*string, bool)`

GetSslModeOk returns a tuple with the SslMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslMode

`func (o *SslConfigV4Request) SetSslMode(v string)`

SetSslMode sets SslMode field to given value.

### HasSslMode

`func (o *SslConfigV4Request) HasSslMode() bool`

HasSslMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


