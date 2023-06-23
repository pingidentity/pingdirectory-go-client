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

// checks if the AddHttpProxyExternalServerRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddHttpProxyExternalServerRequest{}

// AddHttpProxyExternalServerRequest struct for AddHttpProxyExternalServerRequest
type AddHttpProxyExternalServerRequest struct {
	// Name of the new External Server
	ServerName string                                 `json:"serverName"`
	Schemas    []EnumhttpProxyExternalServerSchemaUrn `json:"schemas"`
	// The host name or IP address of the HTTP Proxy External Server.
	ServerHostName string `json:"serverHostName"`
	// The port on which the HTTP Proxy External Server is listening for connections.
	ServerPort int64 `json:"serverPort"`
	// The username to use to authenticate to the HTTP Proxy External Server.
	BasicAuthenticationUsername *string `json:"basicAuthenticationUsername,omitempty"`
	// A passphrase provider that provides access to the password to use to authenticate to the HTTP Proxy External Server.
	BasicAuthenticationPassphraseProvider *string `json:"basicAuthenticationPassphraseProvider,omitempty"`
	// A description for this External Server
	Description *string `json:"description,omitempty"`
}

// NewAddHttpProxyExternalServerRequest instantiates a new AddHttpProxyExternalServerRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddHttpProxyExternalServerRequest(serverName string, schemas []EnumhttpProxyExternalServerSchemaUrn, serverHostName string, serverPort int64) *AddHttpProxyExternalServerRequest {
	this := AddHttpProxyExternalServerRequest{}
	this.ServerName = serverName
	this.Schemas = schemas
	this.ServerHostName = serverHostName
	this.ServerPort = serverPort
	return &this
}

// NewAddHttpProxyExternalServerRequestWithDefaults instantiates a new AddHttpProxyExternalServerRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddHttpProxyExternalServerRequestWithDefaults() *AddHttpProxyExternalServerRequest {
	this := AddHttpProxyExternalServerRequest{}
	return &this
}

// GetServerName returns the ServerName field value
func (o *AddHttpProxyExternalServerRequest) GetServerName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServerName
}

// GetServerNameOk returns a tuple with the ServerName field value
// and a boolean to check if the value has been set.
func (o *AddHttpProxyExternalServerRequest) GetServerNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServerName, true
}

// SetServerName sets field value
func (o *AddHttpProxyExternalServerRequest) SetServerName(v string) {
	o.ServerName = v
}

// GetSchemas returns the Schemas field value
func (o *AddHttpProxyExternalServerRequest) GetSchemas() []EnumhttpProxyExternalServerSchemaUrn {
	if o == nil {
		var ret []EnumhttpProxyExternalServerSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddHttpProxyExternalServerRequest) GetSchemasOk() ([]EnumhttpProxyExternalServerSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddHttpProxyExternalServerRequest) SetSchemas(v []EnumhttpProxyExternalServerSchemaUrn) {
	o.Schemas = v
}

// GetServerHostName returns the ServerHostName field value
func (o *AddHttpProxyExternalServerRequest) GetServerHostName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServerHostName
}

// GetServerHostNameOk returns a tuple with the ServerHostName field value
// and a boolean to check if the value has been set.
func (o *AddHttpProxyExternalServerRequest) GetServerHostNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServerHostName, true
}

// SetServerHostName sets field value
func (o *AddHttpProxyExternalServerRequest) SetServerHostName(v string) {
	o.ServerHostName = v
}

// GetServerPort returns the ServerPort field value
func (o *AddHttpProxyExternalServerRequest) GetServerPort() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ServerPort
}

// GetServerPortOk returns a tuple with the ServerPort field value
// and a boolean to check if the value has been set.
func (o *AddHttpProxyExternalServerRequest) GetServerPortOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServerPort, true
}

// SetServerPort sets field value
func (o *AddHttpProxyExternalServerRequest) SetServerPort(v int64) {
	o.ServerPort = v
}

// GetBasicAuthenticationUsername returns the BasicAuthenticationUsername field value if set, zero value otherwise.
func (o *AddHttpProxyExternalServerRequest) GetBasicAuthenticationUsername() string {
	if o == nil || IsNil(o.BasicAuthenticationUsername) {
		var ret string
		return ret
	}
	return *o.BasicAuthenticationUsername
}

// GetBasicAuthenticationUsernameOk returns a tuple with the BasicAuthenticationUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddHttpProxyExternalServerRequest) GetBasicAuthenticationUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.BasicAuthenticationUsername) {
		return nil, false
	}
	return o.BasicAuthenticationUsername, true
}

// HasBasicAuthenticationUsername returns a boolean if a field has been set.
func (o *AddHttpProxyExternalServerRequest) HasBasicAuthenticationUsername() bool {
	if o != nil && !IsNil(o.BasicAuthenticationUsername) {
		return true
	}

	return false
}

// SetBasicAuthenticationUsername gets a reference to the given string and assigns it to the BasicAuthenticationUsername field.
func (o *AddHttpProxyExternalServerRequest) SetBasicAuthenticationUsername(v string) {
	o.BasicAuthenticationUsername = &v
}

// GetBasicAuthenticationPassphraseProvider returns the BasicAuthenticationPassphraseProvider field value if set, zero value otherwise.
func (o *AddHttpProxyExternalServerRequest) GetBasicAuthenticationPassphraseProvider() string {
	if o == nil || IsNil(o.BasicAuthenticationPassphraseProvider) {
		var ret string
		return ret
	}
	return *o.BasicAuthenticationPassphraseProvider
}

// GetBasicAuthenticationPassphraseProviderOk returns a tuple with the BasicAuthenticationPassphraseProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddHttpProxyExternalServerRequest) GetBasicAuthenticationPassphraseProviderOk() (*string, bool) {
	if o == nil || IsNil(o.BasicAuthenticationPassphraseProvider) {
		return nil, false
	}
	return o.BasicAuthenticationPassphraseProvider, true
}

// HasBasicAuthenticationPassphraseProvider returns a boolean if a field has been set.
func (o *AddHttpProxyExternalServerRequest) HasBasicAuthenticationPassphraseProvider() bool {
	if o != nil && !IsNil(o.BasicAuthenticationPassphraseProvider) {
		return true
	}

	return false
}

// SetBasicAuthenticationPassphraseProvider gets a reference to the given string and assigns it to the BasicAuthenticationPassphraseProvider field.
func (o *AddHttpProxyExternalServerRequest) SetBasicAuthenticationPassphraseProvider(v string) {
	o.BasicAuthenticationPassphraseProvider = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddHttpProxyExternalServerRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddHttpProxyExternalServerRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddHttpProxyExternalServerRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddHttpProxyExternalServerRequest) SetDescription(v string) {
	o.Description = &v
}

func (o AddHttpProxyExternalServerRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddHttpProxyExternalServerRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["serverName"] = o.ServerName
	toSerialize["schemas"] = o.Schemas
	toSerialize["serverHostName"] = o.ServerHostName
	toSerialize["serverPort"] = o.ServerPort
	if !IsNil(o.BasicAuthenticationUsername) {
		toSerialize["basicAuthenticationUsername"] = o.BasicAuthenticationUsername
	}
	if !IsNil(o.BasicAuthenticationPassphraseProvider) {
		toSerialize["basicAuthenticationPassphraseProvider"] = o.BasicAuthenticationPassphraseProvider
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return toSerialize, nil
}

type NullableAddHttpProxyExternalServerRequest struct {
	value *AddHttpProxyExternalServerRequest
	isSet bool
}

func (v NullableAddHttpProxyExternalServerRequest) Get() *AddHttpProxyExternalServerRequest {
	return v.value
}

func (v *NullableAddHttpProxyExternalServerRequest) Set(val *AddHttpProxyExternalServerRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddHttpProxyExternalServerRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddHttpProxyExternalServerRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddHttpProxyExternalServerRequest(val *AddHttpProxyExternalServerRequest) *NullableAddHttpProxyExternalServerRequest {
	return &NullableAddHttpProxyExternalServerRequest{value: val, isSet: true}
}

func (v NullableAddHttpProxyExternalServerRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddHttpProxyExternalServerRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}