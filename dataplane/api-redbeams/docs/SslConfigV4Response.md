# SslConfigV4Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SslCertificates** | Pointer to **[]string** | Set of relevant SSL certificates for the database server, including the active one | [optional] 
**SslCertificateType** | Pointer to **string** | SSL certificate type | [optional] 
**SslMode** | Pointer to **string** | SSL enforcement mode for the database server | [optional] 
**SslCertificateActiveVersion** | Pointer to **int32** | Version number of the SSL certificate currently active for the database server | [optional] 
**SslCertificateHighestAvailableVersion** | Pointer to **int32** | Highest version number of the SSL certificate available for the database server; does not necessarily equal the active version | [optional] 
**SslCertificateActiveCloudProviderIdentifier** | Pointer to **string** | Cloud provider specific identifier of the SSL certificate currently active for the database server | [optional] 

## Methods

### NewSslConfigV4Response

`func NewSslConfigV4Response() *SslConfigV4Response`

NewSslConfigV4Response instantiates a new SslConfigV4Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSslConfigV4ResponseWithDefaults

`func NewSslConfigV4ResponseWithDefaults() *SslConfigV4Response`

NewSslConfigV4ResponseWithDefaults instantiates a new SslConfigV4Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSslCertificates

`func (o *SslConfigV4Response) GetSslCertificates() []string`

GetSslCertificates returns the SslCertificates field if non-nil, zero value otherwise.

### GetSslCertificatesOk

`func (o *SslConfigV4Response) GetSslCertificatesOk() (*[]string, bool)`

GetSslCertificatesOk returns a tuple with the SslCertificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslCertificates

`func (o *SslConfigV4Response) SetSslCertificates(v []string)`

SetSslCertificates sets SslCertificates field to given value.

### HasSslCertificates

`func (o *SslConfigV4Response) HasSslCertificates() bool`

HasSslCertificates returns a boolean if a field has been set.

### GetSslCertificateType

`func (o *SslConfigV4Response) GetSslCertificateType() string`

GetSslCertificateType returns the SslCertificateType field if non-nil, zero value otherwise.

### GetSslCertificateTypeOk

`func (o *SslConfigV4Response) GetSslCertificateTypeOk() (*string, bool)`

GetSslCertificateTypeOk returns a tuple with the SslCertificateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslCertificateType

`func (o *SslConfigV4Response) SetSslCertificateType(v string)`

SetSslCertificateType sets SslCertificateType field to given value.

### HasSslCertificateType

`func (o *SslConfigV4Response) HasSslCertificateType() bool`

HasSslCertificateType returns a boolean if a field has been set.

### GetSslMode

`func (o *SslConfigV4Response) GetSslMode() string`

GetSslMode returns the SslMode field if non-nil, zero value otherwise.

### GetSslModeOk

`func (o *SslConfigV4Response) GetSslModeOk() (*string, bool)`

GetSslModeOk returns a tuple with the SslMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslMode

`func (o *SslConfigV4Response) SetSslMode(v string)`

SetSslMode sets SslMode field to given value.

### HasSslMode

`func (o *SslConfigV4Response) HasSslMode() bool`

HasSslMode returns a boolean if a field has been set.

### GetSslCertificateActiveVersion

`func (o *SslConfigV4Response) GetSslCertificateActiveVersion() int32`

GetSslCertificateActiveVersion returns the SslCertificateActiveVersion field if non-nil, zero value otherwise.

### GetSslCertificateActiveVersionOk

`func (o *SslConfigV4Response) GetSslCertificateActiveVersionOk() (*int32, bool)`

GetSslCertificateActiveVersionOk returns a tuple with the SslCertificateActiveVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslCertificateActiveVersion

`func (o *SslConfigV4Response) SetSslCertificateActiveVersion(v int32)`

SetSslCertificateActiveVersion sets SslCertificateActiveVersion field to given value.

### HasSslCertificateActiveVersion

`func (o *SslConfigV4Response) HasSslCertificateActiveVersion() bool`

HasSslCertificateActiveVersion returns a boolean if a field has been set.

### GetSslCertificateHighestAvailableVersion

`func (o *SslConfigV4Response) GetSslCertificateHighestAvailableVersion() int32`

GetSslCertificateHighestAvailableVersion returns the SslCertificateHighestAvailableVersion field if non-nil, zero value otherwise.

### GetSslCertificateHighestAvailableVersionOk

`func (o *SslConfigV4Response) GetSslCertificateHighestAvailableVersionOk() (*int32, bool)`

GetSslCertificateHighestAvailableVersionOk returns a tuple with the SslCertificateHighestAvailableVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslCertificateHighestAvailableVersion

`func (o *SslConfigV4Response) SetSslCertificateHighestAvailableVersion(v int32)`

SetSslCertificateHighestAvailableVersion sets SslCertificateHighestAvailableVersion field to given value.

### HasSslCertificateHighestAvailableVersion

`func (o *SslConfigV4Response) HasSslCertificateHighestAvailableVersion() bool`

HasSslCertificateHighestAvailableVersion returns a boolean if a field has been set.

### GetSslCertificateActiveCloudProviderIdentifier

`func (o *SslConfigV4Response) GetSslCertificateActiveCloudProviderIdentifier() string`

GetSslCertificateActiveCloudProviderIdentifier returns the SslCertificateActiveCloudProviderIdentifier field if non-nil, zero value otherwise.

### GetSslCertificateActiveCloudProviderIdentifierOk

`func (o *SslConfigV4Response) GetSslCertificateActiveCloudProviderIdentifierOk() (*string, bool)`

GetSslCertificateActiveCloudProviderIdentifierOk returns a tuple with the SslCertificateActiveCloudProviderIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslCertificateActiveCloudProviderIdentifier

`func (o *SslConfigV4Response) SetSslCertificateActiveCloudProviderIdentifier(v string)`

SetSslCertificateActiveCloudProviderIdentifier sets SslCertificateActiveCloudProviderIdentifier field to given value.

### HasSslCertificateActiveCloudProviderIdentifier

`func (o *SslConfigV4Response) HasSslCertificateActiveCloudProviderIdentifier() bool`

HasSslCertificateActiveCloudProviderIdentifier returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


