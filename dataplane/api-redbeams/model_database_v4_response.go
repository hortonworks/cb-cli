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

// checks if the DatabaseV4Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DatabaseV4Response{}

// DatabaseV4Response Response containing information about a database that was acted upon, e.g., retrieved, deleted, listed
type DatabaseV4Response struct {
	// Name of the database
	Name string `json:"name"`
	// Description of the database
	Description *string `json:"description,omitempty"`
	// JDBC connection URL in the form of jdbc:<db-type>:<driver-specific-part>
	ConnectionURL string `json:"connectionURL"`
	// Type of database, i.e., the service name that will use the database (HIVE, DRUID, SUPERSET, RANGER, ...)
	Type string `json:"type"`
	// Name of the JDBC connection driver (for example: 'org.postgresql.Driver')
	ConnectionDriver string `json:"connectionDriver"`
	// CRN of the environment of the database
	EnvironmentCrn string `json:"environmentCrn"`
	// CRN of the database
	Crn *string `json:"crn,omitempty"`
	// Creation date / time of the database, in epoch milliseconds
	CreationDate *int64 `json:"creationDate,omitempty"`
	// Name of the database vendor (MYSQL, POSTGRES...)
	DatabaseEngine string `json:"databaseEngine"`
	// Display name of the database vendor (MySQL, PostgreSQL, ...)
	DatabaseEngineDisplayName string          `json:"databaseEngineDisplayName"`
	ConnectionUserName        *SecretResponse `json:"connectionUserName,omitempty"`
	ConnectionPassword        *SecretResponse `json:"connectionPassword,omitempty"`
	// Ownership status of the database
	ResourceStatus *string `json:"resourceStatus,omitempty"`
}

// NewDatabaseV4Response instantiates a new DatabaseV4Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDatabaseV4Response(name string, connectionURL string, type_ string, connectionDriver string, environmentCrn string, databaseEngine string, databaseEngineDisplayName string) *DatabaseV4Response {
	this := DatabaseV4Response{}
	this.Name = name
	this.ConnectionURL = connectionURL
	this.Type = type_
	this.ConnectionDriver = connectionDriver
	this.EnvironmentCrn = environmentCrn
	this.DatabaseEngine = databaseEngine
	this.DatabaseEngineDisplayName = databaseEngineDisplayName
	return &this
}

// NewDatabaseV4ResponseWithDefaults instantiates a new DatabaseV4Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatabaseV4ResponseWithDefaults() *DatabaseV4Response {
	this := DatabaseV4Response{}
	return &this
}

// GetName returns the Name field value
func (o *DatabaseV4Response) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DatabaseV4Response) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DatabaseV4Response) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DatabaseV4Response) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseV4Response) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DatabaseV4Response) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DatabaseV4Response) SetDescription(v string) {
	o.Description = &v
}

// GetConnectionURL returns the ConnectionURL field value
func (o *DatabaseV4Response) GetConnectionURL() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConnectionURL
}

// GetConnectionURLOk returns a tuple with the ConnectionURL field value
// and a boolean to check if the value has been set.
func (o *DatabaseV4Response) GetConnectionURLOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConnectionURL, true
}

// SetConnectionURL sets field value
func (o *DatabaseV4Response) SetConnectionURL(v string) {
	o.ConnectionURL = v
}

// GetType returns the Type field value
func (o *DatabaseV4Response) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DatabaseV4Response) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *DatabaseV4Response) SetType(v string) {
	o.Type = v
}

// GetConnectionDriver returns the ConnectionDriver field value
func (o *DatabaseV4Response) GetConnectionDriver() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConnectionDriver
}

// GetConnectionDriverOk returns a tuple with the ConnectionDriver field value
// and a boolean to check if the value has been set.
func (o *DatabaseV4Response) GetConnectionDriverOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConnectionDriver, true
}

// SetConnectionDriver sets field value
func (o *DatabaseV4Response) SetConnectionDriver(v string) {
	o.ConnectionDriver = v
}

// GetEnvironmentCrn returns the EnvironmentCrn field value
func (o *DatabaseV4Response) GetEnvironmentCrn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EnvironmentCrn
}

// GetEnvironmentCrnOk returns a tuple with the EnvironmentCrn field value
// and a boolean to check if the value has been set.
func (o *DatabaseV4Response) GetEnvironmentCrnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnvironmentCrn, true
}

// SetEnvironmentCrn sets field value
func (o *DatabaseV4Response) SetEnvironmentCrn(v string) {
	o.EnvironmentCrn = v
}

// GetCrn returns the Crn field value if set, zero value otherwise.
func (o *DatabaseV4Response) GetCrn() string {
	if o == nil || isNil(o.Crn) {
		var ret string
		return ret
	}
	return *o.Crn
}

// GetCrnOk returns a tuple with the Crn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseV4Response) GetCrnOk() (*string, bool) {
	if o == nil || isNil(o.Crn) {
		return nil, false
	}
	return o.Crn, true
}

// HasCrn returns a boolean if a field has been set.
func (o *DatabaseV4Response) HasCrn() bool {
	if o != nil && !isNil(o.Crn) {
		return true
	}

	return false
}

// SetCrn gets a reference to the given string and assigns it to the Crn field.
func (o *DatabaseV4Response) SetCrn(v string) {
	o.Crn = &v
}

// GetCreationDate returns the CreationDate field value if set, zero value otherwise.
func (o *DatabaseV4Response) GetCreationDate() int64 {
	if o == nil || isNil(o.CreationDate) {
		var ret int64
		return ret
	}
	return *o.CreationDate
}

// GetCreationDateOk returns a tuple with the CreationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseV4Response) GetCreationDateOk() (*int64, bool) {
	if o == nil || isNil(o.CreationDate) {
		return nil, false
	}
	return o.CreationDate, true
}

// HasCreationDate returns a boolean if a field has been set.
func (o *DatabaseV4Response) HasCreationDate() bool {
	if o != nil && !isNil(o.CreationDate) {
		return true
	}

	return false
}

// SetCreationDate gets a reference to the given int64 and assigns it to the CreationDate field.
func (o *DatabaseV4Response) SetCreationDate(v int64) {
	o.CreationDate = &v
}

// GetDatabaseEngine returns the DatabaseEngine field value
func (o *DatabaseV4Response) GetDatabaseEngine() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DatabaseEngine
}

// GetDatabaseEngineOk returns a tuple with the DatabaseEngine field value
// and a boolean to check if the value has been set.
func (o *DatabaseV4Response) GetDatabaseEngineOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DatabaseEngine, true
}

// SetDatabaseEngine sets field value
func (o *DatabaseV4Response) SetDatabaseEngine(v string) {
	o.DatabaseEngine = v
}

// GetDatabaseEngineDisplayName returns the DatabaseEngineDisplayName field value
func (o *DatabaseV4Response) GetDatabaseEngineDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DatabaseEngineDisplayName
}

// GetDatabaseEngineDisplayNameOk returns a tuple with the DatabaseEngineDisplayName field value
// and a boolean to check if the value has been set.
func (o *DatabaseV4Response) GetDatabaseEngineDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DatabaseEngineDisplayName, true
}

// SetDatabaseEngineDisplayName sets field value
func (o *DatabaseV4Response) SetDatabaseEngineDisplayName(v string) {
	o.DatabaseEngineDisplayName = v
}

// GetConnectionUserName returns the ConnectionUserName field value if set, zero value otherwise.
func (o *DatabaseV4Response) GetConnectionUserName() SecretResponse {
	if o == nil || isNil(o.ConnectionUserName) {
		var ret SecretResponse
		return ret
	}
	return *o.ConnectionUserName
}

// GetConnectionUserNameOk returns a tuple with the ConnectionUserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseV4Response) GetConnectionUserNameOk() (*SecretResponse, bool) {
	if o == nil || isNil(o.ConnectionUserName) {
		return nil, false
	}
	return o.ConnectionUserName, true
}

// HasConnectionUserName returns a boolean if a field has been set.
func (o *DatabaseV4Response) HasConnectionUserName() bool {
	if o != nil && !isNil(o.ConnectionUserName) {
		return true
	}

	return false
}

// SetConnectionUserName gets a reference to the given SecretResponse and assigns it to the ConnectionUserName field.
func (o *DatabaseV4Response) SetConnectionUserName(v SecretResponse) {
	o.ConnectionUserName = &v
}

// GetConnectionPassword returns the ConnectionPassword field value if set, zero value otherwise.
func (o *DatabaseV4Response) GetConnectionPassword() SecretResponse {
	if o == nil || isNil(o.ConnectionPassword) {
		var ret SecretResponse
		return ret
	}
	return *o.ConnectionPassword
}

// GetConnectionPasswordOk returns a tuple with the ConnectionPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseV4Response) GetConnectionPasswordOk() (*SecretResponse, bool) {
	if o == nil || isNil(o.ConnectionPassword) {
		return nil, false
	}
	return o.ConnectionPassword, true
}

// HasConnectionPassword returns a boolean if a field has been set.
func (o *DatabaseV4Response) HasConnectionPassword() bool {
	if o != nil && !isNil(o.ConnectionPassword) {
		return true
	}

	return false
}

// SetConnectionPassword gets a reference to the given SecretResponse and assigns it to the ConnectionPassword field.
func (o *DatabaseV4Response) SetConnectionPassword(v SecretResponse) {
	o.ConnectionPassword = &v
}

// GetResourceStatus returns the ResourceStatus field value if set, zero value otherwise.
func (o *DatabaseV4Response) GetResourceStatus() string {
	if o == nil || isNil(o.ResourceStatus) {
		var ret string
		return ret
	}
	return *o.ResourceStatus
}

// GetResourceStatusOk returns a tuple with the ResourceStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseV4Response) GetResourceStatusOk() (*string, bool) {
	if o == nil || isNil(o.ResourceStatus) {
		return nil, false
	}
	return o.ResourceStatus, true
}

// HasResourceStatus returns a boolean if a field has been set.
func (o *DatabaseV4Response) HasResourceStatus() bool {
	if o != nil && !isNil(o.ResourceStatus) {
		return true
	}

	return false
}

// SetResourceStatus gets a reference to the given string and assigns it to the ResourceStatus field.
func (o *DatabaseV4Response) SetResourceStatus(v string) {
	o.ResourceStatus = &v
}

func (o DatabaseV4Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DatabaseV4Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["connectionURL"] = o.ConnectionURL
	toSerialize["type"] = o.Type
	toSerialize["connectionDriver"] = o.ConnectionDriver
	toSerialize["environmentCrn"] = o.EnvironmentCrn
	if !isNil(o.Crn) {
		toSerialize["crn"] = o.Crn
	}
	if !isNil(o.CreationDate) {
		toSerialize["creationDate"] = o.CreationDate
	}
	toSerialize["databaseEngine"] = o.DatabaseEngine
	toSerialize["databaseEngineDisplayName"] = o.DatabaseEngineDisplayName
	if !isNil(o.ConnectionUserName) {
		toSerialize["connectionUserName"] = o.ConnectionUserName
	}
	if !isNil(o.ConnectionPassword) {
		toSerialize["connectionPassword"] = o.ConnectionPassword
	}
	if !isNil(o.ResourceStatus) {
		toSerialize["resourceStatus"] = o.ResourceStatus
	}
	return toSerialize, nil
}

type NullableDatabaseV4Response struct {
	value *DatabaseV4Response
	isSet bool
}

func (v NullableDatabaseV4Response) Get() *DatabaseV4Response {
	return v.value
}

func (v *NullableDatabaseV4Response) Set(val *DatabaseV4Response) {
	v.value = val
	v.isSet = true
}

func (v NullableDatabaseV4Response) IsSet() bool {
	return v.isSet
}

func (v *NullableDatabaseV4Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatabaseV4Response(val *DatabaseV4Response) *NullableDatabaseV4Response {
	return &NullableDatabaseV4Response{value: val, isSet: true}
}

func (v NullableDatabaseV4Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatabaseV4Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
