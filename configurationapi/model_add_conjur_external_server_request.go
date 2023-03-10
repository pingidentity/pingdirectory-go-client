/*
PingData Config - OpenAPI 3.0

This is the PingData Configuration API

API version: 0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configurationapi

import (
	"encoding/json"
)

// checks if the AddConjurExternalServerRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddConjurExternalServerRequest{}

// AddConjurExternalServerRequest struct for AddConjurExternalServerRequest
type AddConjurExternalServerRequest struct {
	// Name of the new External Server
	ServerName string                              `json:"serverName"`
	Schemas    []EnumconjurExternalServerSchemaUrn `json:"schemas"`
	// The base URL needed to access the CyberArk Conjur server. The base URL should consist of the protocol (\"http\" or \"https\"), the server address (resolvable name or IP address), and the port number. For example, \"https://conjur.example.com:8443/\".
	ConjurServerBaseURI []string `json:"conjurServerBaseURI"`
	// The mechanism used to authenticate to the Conjur server.
	ConjurAuthenticationMethod string `json:"conjurAuthenticationMethod"`
	// The name of the account with which the desired secrets are associated.
	ConjurAccountName string `json:"conjurAccountName"`
	// The path to a file containing the information needed to trust the certificate presented by the Conjur servers.
	TrustStoreFile *string `json:"trustStoreFile,omitempty"`
	// The PIN needed to access the contents of the trust store. This is only required if a trust store file is required, and if that trust store requires a PIN to access its contents.
	TrustStorePin *string `json:"trustStorePin,omitempty"`
	// The store type for the specified trust store file. The value should likely be one of \"JKS\", \"PKCS12\", or \"BCFKS\".
	TrustStoreType *string `json:"trustStoreType,omitempty"`
	// A description for this External Server
	Description *string `json:"description,omitempty"`
}

// NewAddConjurExternalServerRequest instantiates a new AddConjurExternalServerRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddConjurExternalServerRequest(serverName string, schemas []EnumconjurExternalServerSchemaUrn, conjurServerBaseURI []string, conjurAuthenticationMethod string, conjurAccountName string) *AddConjurExternalServerRequest {
	this := AddConjurExternalServerRequest{}
	this.ServerName = serverName
	this.Schemas = schemas
	this.ConjurServerBaseURI = conjurServerBaseURI
	this.ConjurAuthenticationMethod = conjurAuthenticationMethod
	this.ConjurAccountName = conjurAccountName
	return &this
}

// NewAddConjurExternalServerRequestWithDefaults instantiates a new AddConjurExternalServerRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddConjurExternalServerRequestWithDefaults() *AddConjurExternalServerRequest {
	this := AddConjurExternalServerRequest{}
	return &this
}

// GetServerName returns the ServerName field value
func (o *AddConjurExternalServerRequest) GetServerName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServerName
}

// GetServerNameOk returns a tuple with the ServerName field value
// and a boolean to check if the value has been set.
func (o *AddConjurExternalServerRequest) GetServerNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServerName, true
}

// SetServerName sets field value
func (o *AddConjurExternalServerRequest) SetServerName(v string) {
	o.ServerName = v
}

// GetSchemas returns the Schemas field value
func (o *AddConjurExternalServerRequest) GetSchemas() []EnumconjurExternalServerSchemaUrn {
	if o == nil {
		var ret []EnumconjurExternalServerSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddConjurExternalServerRequest) GetSchemasOk() ([]EnumconjurExternalServerSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddConjurExternalServerRequest) SetSchemas(v []EnumconjurExternalServerSchemaUrn) {
	o.Schemas = v
}

// GetConjurServerBaseURI returns the ConjurServerBaseURI field value
func (o *AddConjurExternalServerRequest) GetConjurServerBaseURI() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ConjurServerBaseURI
}

// GetConjurServerBaseURIOk returns a tuple with the ConjurServerBaseURI field value
// and a boolean to check if the value has been set.
func (o *AddConjurExternalServerRequest) GetConjurServerBaseURIOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConjurServerBaseURI, true
}

// SetConjurServerBaseURI sets field value
func (o *AddConjurExternalServerRequest) SetConjurServerBaseURI(v []string) {
	o.ConjurServerBaseURI = v
}

// GetConjurAuthenticationMethod returns the ConjurAuthenticationMethod field value
func (o *AddConjurExternalServerRequest) GetConjurAuthenticationMethod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConjurAuthenticationMethod
}

// GetConjurAuthenticationMethodOk returns a tuple with the ConjurAuthenticationMethod field value
// and a boolean to check if the value has been set.
func (o *AddConjurExternalServerRequest) GetConjurAuthenticationMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConjurAuthenticationMethod, true
}

// SetConjurAuthenticationMethod sets field value
func (o *AddConjurExternalServerRequest) SetConjurAuthenticationMethod(v string) {
	o.ConjurAuthenticationMethod = v
}

// GetConjurAccountName returns the ConjurAccountName field value
func (o *AddConjurExternalServerRequest) GetConjurAccountName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConjurAccountName
}

// GetConjurAccountNameOk returns a tuple with the ConjurAccountName field value
// and a boolean to check if the value has been set.
func (o *AddConjurExternalServerRequest) GetConjurAccountNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConjurAccountName, true
}

// SetConjurAccountName sets field value
func (o *AddConjurExternalServerRequest) SetConjurAccountName(v string) {
	o.ConjurAccountName = v
}

// GetTrustStoreFile returns the TrustStoreFile field value if set, zero value otherwise.
func (o *AddConjurExternalServerRequest) GetTrustStoreFile() string {
	if o == nil || IsNil(o.TrustStoreFile) {
		var ret string
		return ret
	}
	return *o.TrustStoreFile
}

// GetTrustStoreFileOk returns a tuple with the TrustStoreFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddConjurExternalServerRequest) GetTrustStoreFileOk() (*string, bool) {
	if o == nil || IsNil(o.TrustStoreFile) {
		return nil, false
	}
	return o.TrustStoreFile, true
}

// HasTrustStoreFile returns a boolean if a field has been set.
func (o *AddConjurExternalServerRequest) HasTrustStoreFile() bool {
	if o != nil && !IsNil(o.TrustStoreFile) {
		return true
	}

	return false
}

// SetTrustStoreFile gets a reference to the given string and assigns it to the TrustStoreFile field.
func (o *AddConjurExternalServerRequest) SetTrustStoreFile(v string) {
	o.TrustStoreFile = &v
}

// GetTrustStorePin returns the TrustStorePin field value if set, zero value otherwise.
func (o *AddConjurExternalServerRequest) GetTrustStorePin() string {
	if o == nil || IsNil(o.TrustStorePin) {
		var ret string
		return ret
	}
	return *o.TrustStorePin
}

// GetTrustStorePinOk returns a tuple with the TrustStorePin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddConjurExternalServerRequest) GetTrustStorePinOk() (*string, bool) {
	if o == nil || IsNil(o.TrustStorePin) {
		return nil, false
	}
	return o.TrustStorePin, true
}

// HasTrustStorePin returns a boolean if a field has been set.
func (o *AddConjurExternalServerRequest) HasTrustStorePin() bool {
	if o != nil && !IsNil(o.TrustStorePin) {
		return true
	}

	return false
}

// SetTrustStorePin gets a reference to the given string and assigns it to the TrustStorePin field.
func (o *AddConjurExternalServerRequest) SetTrustStorePin(v string) {
	o.TrustStorePin = &v
}

// GetTrustStoreType returns the TrustStoreType field value if set, zero value otherwise.
func (o *AddConjurExternalServerRequest) GetTrustStoreType() string {
	if o == nil || IsNil(o.TrustStoreType) {
		var ret string
		return ret
	}
	return *o.TrustStoreType
}

// GetTrustStoreTypeOk returns a tuple with the TrustStoreType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddConjurExternalServerRequest) GetTrustStoreTypeOk() (*string, bool) {
	if o == nil || IsNil(o.TrustStoreType) {
		return nil, false
	}
	return o.TrustStoreType, true
}

// HasTrustStoreType returns a boolean if a field has been set.
func (o *AddConjurExternalServerRequest) HasTrustStoreType() bool {
	if o != nil && !IsNil(o.TrustStoreType) {
		return true
	}

	return false
}

// SetTrustStoreType gets a reference to the given string and assigns it to the TrustStoreType field.
func (o *AddConjurExternalServerRequest) SetTrustStoreType(v string) {
	o.TrustStoreType = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddConjurExternalServerRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddConjurExternalServerRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddConjurExternalServerRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddConjurExternalServerRequest) SetDescription(v string) {
	o.Description = &v
}

func (o AddConjurExternalServerRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddConjurExternalServerRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["serverName"] = o.ServerName
	toSerialize["schemas"] = o.Schemas
	toSerialize["conjurServerBaseURI"] = o.ConjurServerBaseURI
	toSerialize["conjurAuthenticationMethod"] = o.ConjurAuthenticationMethod
	toSerialize["conjurAccountName"] = o.ConjurAccountName
	if !IsNil(o.TrustStoreFile) {
		toSerialize["trustStoreFile"] = o.TrustStoreFile
	}
	if !IsNil(o.TrustStorePin) {
		toSerialize["trustStorePin"] = o.TrustStorePin
	}
	if !IsNil(o.TrustStoreType) {
		toSerialize["trustStoreType"] = o.TrustStoreType
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return toSerialize, nil
}

type NullableAddConjurExternalServerRequest struct {
	value *AddConjurExternalServerRequest
	isSet bool
}

func (v NullableAddConjurExternalServerRequest) Get() *AddConjurExternalServerRequest {
	return v.value
}

func (v *NullableAddConjurExternalServerRequest) Set(val *AddConjurExternalServerRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddConjurExternalServerRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddConjurExternalServerRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddConjurExternalServerRequest(val *AddConjurExternalServerRequest) *NullableAddConjurExternalServerRequest {
	return &NullableAddConjurExternalServerRequest{value: val, isSet: true}
}

func (v NullableAddConjurExternalServerRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddConjurExternalServerRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
