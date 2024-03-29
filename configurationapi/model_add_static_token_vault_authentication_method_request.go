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

// checks if the AddStaticTokenVaultAuthenticationMethodRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddStaticTokenVaultAuthenticationMethodRequest{}

// AddStaticTokenVaultAuthenticationMethodRequest struct for AddStaticTokenVaultAuthenticationMethodRequest
type AddStaticTokenVaultAuthenticationMethodRequest struct {
	Schemas []EnumstaticTokenVaultAuthenticationMethodSchemaUrn `json:"schemas"`
	// The static token used to authenticate to the Vault server.
	VaultAccessToken string `json:"vaultAccessToken"`
	// A description for this Vault Authentication Method
	Description *string `json:"description,omitempty"`
	// Name of the new Vault Authentication Method
	MethodName string `json:"methodName"`
}

// NewAddStaticTokenVaultAuthenticationMethodRequest instantiates a new AddStaticTokenVaultAuthenticationMethodRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddStaticTokenVaultAuthenticationMethodRequest(schemas []EnumstaticTokenVaultAuthenticationMethodSchemaUrn, vaultAccessToken string, methodName string) *AddStaticTokenVaultAuthenticationMethodRequest {
	this := AddStaticTokenVaultAuthenticationMethodRequest{}
	this.Schemas = schemas
	this.VaultAccessToken = vaultAccessToken
	this.MethodName = methodName
	return &this
}

// NewAddStaticTokenVaultAuthenticationMethodRequestWithDefaults instantiates a new AddStaticTokenVaultAuthenticationMethodRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddStaticTokenVaultAuthenticationMethodRequestWithDefaults() *AddStaticTokenVaultAuthenticationMethodRequest {
	this := AddStaticTokenVaultAuthenticationMethodRequest{}
	return &this
}

// GetSchemas returns the Schemas field value
func (o *AddStaticTokenVaultAuthenticationMethodRequest) GetSchemas() []EnumstaticTokenVaultAuthenticationMethodSchemaUrn {
	if o == nil {
		var ret []EnumstaticTokenVaultAuthenticationMethodSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddStaticTokenVaultAuthenticationMethodRequest) GetSchemasOk() ([]EnumstaticTokenVaultAuthenticationMethodSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddStaticTokenVaultAuthenticationMethodRequest) SetSchemas(v []EnumstaticTokenVaultAuthenticationMethodSchemaUrn) {
	o.Schemas = v
}

// GetVaultAccessToken returns the VaultAccessToken field value
func (o *AddStaticTokenVaultAuthenticationMethodRequest) GetVaultAccessToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VaultAccessToken
}

// GetVaultAccessTokenOk returns a tuple with the VaultAccessToken field value
// and a boolean to check if the value has been set.
func (o *AddStaticTokenVaultAuthenticationMethodRequest) GetVaultAccessTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VaultAccessToken, true
}

// SetVaultAccessToken sets field value
func (o *AddStaticTokenVaultAuthenticationMethodRequest) SetVaultAccessToken(v string) {
	o.VaultAccessToken = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddStaticTokenVaultAuthenticationMethodRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddStaticTokenVaultAuthenticationMethodRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddStaticTokenVaultAuthenticationMethodRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddStaticTokenVaultAuthenticationMethodRequest) SetDescription(v string) {
	o.Description = &v
}

// GetMethodName returns the MethodName field value
func (o *AddStaticTokenVaultAuthenticationMethodRequest) GetMethodName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MethodName
}

// GetMethodNameOk returns a tuple with the MethodName field value
// and a boolean to check if the value has been set.
func (o *AddStaticTokenVaultAuthenticationMethodRequest) GetMethodNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MethodName, true
}

// SetMethodName sets field value
func (o *AddStaticTokenVaultAuthenticationMethodRequest) SetMethodName(v string) {
	o.MethodName = v
}

func (o AddStaticTokenVaultAuthenticationMethodRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddStaticTokenVaultAuthenticationMethodRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["schemas"] = o.Schemas
	toSerialize["vaultAccessToken"] = o.VaultAccessToken
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["methodName"] = o.MethodName
	return toSerialize, nil
}

type NullableAddStaticTokenVaultAuthenticationMethodRequest struct {
	value *AddStaticTokenVaultAuthenticationMethodRequest
	isSet bool
}

func (v NullableAddStaticTokenVaultAuthenticationMethodRequest) Get() *AddStaticTokenVaultAuthenticationMethodRequest {
	return v.value
}

func (v *NullableAddStaticTokenVaultAuthenticationMethodRequest) Set(val *AddStaticTokenVaultAuthenticationMethodRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddStaticTokenVaultAuthenticationMethodRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddStaticTokenVaultAuthenticationMethodRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddStaticTokenVaultAuthenticationMethodRequest(val *AddStaticTokenVaultAuthenticationMethodRequest) *NullableAddStaticTokenVaultAuthenticationMethodRequest {
	return &NullableAddStaticTokenVaultAuthenticationMethodRequest{value: val, isSet: true}
}

func (v NullableAddStaticTokenVaultAuthenticationMethodRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddStaticTokenVaultAuthenticationMethodRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
