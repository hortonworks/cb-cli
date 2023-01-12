/*
Redbeams API

API for working with databases and database servers

API version: 2.66.0-b41-1-ge11bb40
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the SslConfigV4Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SslConfigV4Response{}

// SslConfigV4Response Response for the SSL config of a database server
type SslConfigV4Response struct {
	// Set of relevant SSL certificates for the database server, including the active one
	SslCertificates []string `json:"sslCertificates,omitempty"`
	// SSL certificate type
	SslCertificateType *string `json:"sslCertificateType,omitempty"`
	// SSL enforcement mode for the database server
	SslMode *string `json:"sslMode,omitempty"`
	// Version number of the SSL certificate currently active for the database server
	SslCertificateActiveVersion *int32 `json:"sslCertificateActiveVersion,omitempty"`
	// Highest version number of the SSL certificate available for the database server; does not necessarily equal the active version
	SslCertificateHighestAvailableVersion *int32 `json:"sslCertificateHighestAvailableVersion,omitempty"`
	// Cloud provider specific identifier of the SSL certificate currently active for the database server
	SslCertificateActiveCloudProviderIdentifier *string `json:"sslCertificateActiveCloudProviderIdentifier,omitempty"`
}

// NewSslConfigV4Response instantiates a new SslConfigV4Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSslConfigV4Response() *SslConfigV4Response {
	this := SslConfigV4Response{}
	return &this
}

// NewSslConfigV4ResponseWithDefaults instantiates a new SslConfigV4Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSslConfigV4ResponseWithDefaults() *SslConfigV4Response {
	this := SslConfigV4Response{}
	return &this
}

// GetSslCertificates returns the SslCertificates field value if set, zero value otherwise.
func (o *SslConfigV4Response) GetSslCertificates() []string {
	if o == nil || isNil(o.SslCertificates) {
		var ret []string
		return ret
	}
	return o.SslCertificates
}

// GetSslCertificatesOk returns a tuple with the SslCertificates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SslConfigV4Response) GetSslCertificatesOk() ([]string, bool) {
	if o == nil || isNil(o.SslCertificates) {
		return nil, false
	}
	return o.SslCertificates, true
}

// HasSslCertificates returns a boolean if a field has been set.
func (o *SslConfigV4Response) HasSslCertificates() bool {
	if o != nil && !isNil(o.SslCertificates) {
		return true
	}

	return false
}

// SetSslCertificates gets a reference to the given []string and assigns it to the SslCertificates field.
func (o *SslConfigV4Response) SetSslCertificates(v []string) {
	o.SslCertificates = v
}

// GetSslCertificateType returns the SslCertificateType field value if set, zero value otherwise.
func (o *SslConfigV4Response) GetSslCertificateType() string {
	if o == nil || isNil(o.SslCertificateType) {
		var ret string
		return ret
	}
	return *o.SslCertificateType
}

// GetSslCertificateTypeOk returns a tuple with the SslCertificateType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SslConfigV4Response) GetSslCertificateTypeOk() (*string, bool) {
	if o == nil || isNil(o.SslCertificateType) {
		return nil, false
	}
	return o.SslCertificateType, true
}

// HasSslCertificateType returns a boolean if a field has been set.
func (o *SslConfigV4Response) HasSslCertificateType() bool {
	if o != nil && !isNil(o.SslCertificateType) {
		return true
	}

	return false
}

// SetSslCertificateType gets a reference to the given string and assigns it to the SslCertificateType field.
func (o *SslConfigV4Response) SetSslCertificateType(v string) {
	o.SslCertificateType = &v
}

// GetSslMode returns the SslMode field value if set, zero value otherwise.
func (o *SslConfigV4Response) GetSslMode() string {
	if o == nil || isNil(o.SslMode) {
		var ret string
		return ret
	}
	return *o.SslMode
}

// GetSslModeOk returns a tuple with the SslMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SslConfigV4Response) GetSslModeOk() (*string, bool) {
	if o == nil || isNil(o.SslMode) {
		return nil, false
	}
	return o.SslMode, true
}

// HasSslMode returns a boolean if a field has been set.
func (o *SslConfigV4Response) HasSslMode() bool {
	if o != nil && !isNil(o.SslMode) {
		return true
	}

	return false
}

// SetSslMode gets a reference to the given string and assigns it to the SslMode field.
func (o *SslConfigV4Response) SetSslMode(v string) {
	o.SslMode = &v
}

// GetSslCertificateActiveVersion returns the SslCertificateActiveVersion field value if set, zero value otherwise.
func (o *SslConfigV4Response) GetSslCertificateActiveVersion() int32 {
	if o == nil || isNil(o.SslCertificateActiveVersion) {
		var ret int32
		return ret
	}
	return *o.SslCertificateActiveVersion
}

// GetSslCertificateActiveVersionOk returns a tuple with the SslCertificateActiveVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SslConfigV4Response) GetSslCertificateActiveVersionOk() (*int32, bool) {
	if o == nil || isNil(o.SslCertificateActiveVersion) {
		return nil, false
	}
	return o.SslCertificateActiveVersion, true
}

// HasSslCertificateActiveVersion returns a boolean if a field has been set.
func (o *SslConfigV4Response) HasSslCertificateActiveVersion() bool {
	if o != nil && !isNil(o.SslCertificateActiveVersion) {
		return true
	}

	return false
}

// SetSslCertificateActiveVersion gets a reference to the given int32 and assigns it to the SslCertificateActiveVersion field.
func (o *SslConfigV4Response) SetSslCertificateActiveVersion(v int32) {
	o.SslCertificateActiveVersion = &v
}

// GetSslCertificateHighestAvailableVersion returns the SslCertificateHighestAvailableVersion field value if set, zero value otherwise.
func (o *SslConfigV4Response) GetSslCertificateHighestAvailableVersion() int32 {
	if o == nil || isNil(o.SslCertificateHighestAvailableVersion) {
		var ret int32
		return ret
	}
	return *o.SslCertificateHighestAvailableVersion
}

// GetSslCertificateHighestAvailableVersionOk returns a tuple with the SslCertificateHighestAvailableVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SslConfigV4Response) GetSslCertificateHighestAvailableVersionOk() (*int32, bool) {
	if o == nil || isNil(o.SslCertificateHighestAvailableVersion) {
		return nil, false
	}
	return o.SslCertificateHighestAvailableVersion, true
}

// HasSslCertificateHighestAvailableVersion returns a boolean if a field has been set.
func (o *SslConfigV4Response) HasSslCertificateHighestAvailableVersion() bool {
	if o != nil && !isNil(o.SslCertificateHighestAvailableVersion) {
		return true
	}

	return false
}

// SetSslCertificateHighestAvailableVersion gets a reference to the given int32 and assigns it to the SslCertificateHighestAvailableVersion field.
func (o *SslConfigV4Response) SetSslCertificateHighestAvailableVersion(v int32) {
	o.SslCertificateHighestAvailableVersion = &v
}

// GetSslCertificateActiveCloudProviderIdentifier returns the SslCertificateActiveCloudProviderIdentifier field value if set, zero value otherwise.
func (o *SslConfigV4Response) GetSslCertificateActiveCloudProviderIdentifier() string {
	if o == nil || isNil(o.SslCertificateActiveCloudProviderIdentifier) {
		var ret string
		return ret
	}
	return *o.SslCertificateActiveCloudProviderIdentifier
}

// GetSslCertificateActiveCloudProviderIdentifierOk returns a tuple with the SslCertificateActiveCloudProviderIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SslConfigV4Response) GetSslCertificateActiveCloudProviderIdentifierOk() (*string, bool) {
	if o == nil || isNil(o.SslCertificateActiveCloudProviderIdentifier) {
		return nil, false
	}
	return o.SslCertificateActiveCloudProviderIdentifier, true
}

// HasSslCertificateActiveCloudProviderIdentifier returns a boolean if a field has been set.
func (o *SslConfigV4Response) HasSslCertificateActiveCloudProviderIdentifier() bool {
	if o != nil && !isNil(o.SslCertificateActiveCloudProviderIdentifier) {
		return true
	}

	return false
}

// SetSslCertificateActiveCloudProviderIdentifier gets a reference to the given string and assigns it to the SslCertificateActiveCloudProviderIdentifier field.
func (o *SslConfigV4Response) SetSslCertificateActiveCloudProviderIdentifier(v string) {
	o.SslCertificateActiveCloudProviderIdentifier = &v
}

func (o SslConfigV4Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SslConfigV4Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.SslCertificates) {
		toSerialize["sslCertificates"] = o.SslCertificates
	}
	if !isNil(o.SslCertificateType) {
		toSerialize["sslCertificateType"] = o.SslCertificateType
	}
	if !isNil(o.SslMode) {
		toSerialize["sslMode"] = o.SslMode
	}
	if !isNil(o.SslCertificateActiveVersion) {
		toSerialize["sslCertificateActiveVersion"] = o.SslCertificateActiveVersion
	}
	if !isNil(o.SslCertificateHighestAvailableVersion) {
		toSerialize["sslCertificateHighestAvailableVersion"] = o.SslCertificateHighestAvailableVersion
	}
	if !isNil(o.SslCertificateActiveCloudProviderIdentifier) {
		toSerialize["sslCertificateActiveCloudProviderIdentifier"] = o.SslCertificateActiveCloudProviderIdentifier
	}
	return toSerialize, nil
}

type NullableSslConfigV4Response struct {
	value *SslConfigV4Response
	isSet bool
}

func (v NullableSslConfigV4Response) Get() *SslConfigV4Response {
	return v.value
}

func (v *NullableSslConfigV4Response) Set(val *SslConfigV4Response) {
	v.value = val
	v.isSet = true
}

func (v NullableSslConfigV4Response) IsSet() bool {
	return v.isSet
}

func (v *NullableSslConfigV4Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSslConfigV4Response(val *SslConfigV4Response) *NullableSslConfigV4Response {
	return &NullableSslConfigV4Response{value: val, isSet: true}
}

func (v NullableSslConfigV4Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSslConfigV4Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
