/*
PingData Config - OpenAPI 3.0

This is the PingData Configuration API

API version: 0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// AddAppRoleVaultAuthenticationMethodRequest struct for AddAppRoleVaultAuthenticationMethodRequest
type AddAppRoleVaultAuthenticationMethodRequest struct {
	// Name of the new Vault Authentication Method
	MethodName string `json:"methodName"`
	Schemas []EnumappRoleVaultAuthenticationMethodSchemaUrn `json:"schemas"`
	// The role ID for the AppRole to authenticate.
	VaultRoleID string `json:"vaultRoleID"`
	// The secret ID for the AppRole to authenticate.
	VaultSecretID string `json:"vaultSecretID"`
	// The name used when enabling the desired AppRole authentication mechanism in the Vault server.
	LoginMechanismName *string `json:"loginMechanismName,omitempty"`
	// A description for this Vault Authentication Method
	Description *string `json:"description,omitempty"`
}

// NewAddAppRoleVaultAuthenticationMethodRequest instantiates a new AddAppRoleVaultAuthenticationMethodRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddAppRoleVaultAuthenticationMethodRequest(methodName string, schemas []EnumappRoleVaultAuthenticationMethodSchemaUrn, vaultRoleID string, vaultSecretID string) *AddAppRoleVaultAuthenticationMethodRequest {
	this := AddAppRoleVaultAuthenticationMethodRequest{}
	this.MethodName = methodName
	this.Schemas = schemas
	this.VaultRoleID = vaultRoleID
	this.VaultSecretID = vaultSecretID
	return &this
}

// NewAddAppRoleVaultAuthenticationMethodRequestWithDefaults instantiates a new AddAppRoleVaultAuthenticationMethodRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddAppRoleVaultAuthenticationMethodRequestWithDefaults() *AddAppRoleVaultAuthenticationMethodRequest {
	this := AddAppRoleVaultAuthenticationMethodRequest{}
	return &this
}

// GetMethodName returns the MethodName field value
func (o *AddAppRoleVaultAuthenticationMethodRequest) GetMethodName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MethodName
}

// GetMethodNameOk returns a tuple with the MethodName field value
// and a boolean to check if the value has been set.
func (o *AddAppRoleVaultAuthenticationMethodRequest) GetMethodNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.MethodName, true
}

// SetMethodName sets field value
func (o *AddAppRoleVaultAuthenticationMethodRequest) SetMethodName(v string) {
	o.MethodName = v
}

// GetSchemas returns the Schemas field value
func (o *AddAppRoleVaultAuthenticationMethodRequest) GetSchemas() []EnumappRoleVaultAuthenticationMethodSchemaUrn {
	if o == nil {
		var ret []EnumappRoleVaultAuthenticationMethodSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddAppRoleVaultAuthenticationMethodRequest) GetSchemasOk() ([]EnumappRoleVaultAuthenticationMethodSchemaUrn, bool) {
	if o == nil {
    return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddAppRoleVaultAuthenticationMethodRequest) SetSchemas(v []EnumappRoleVaultAuthenticationMethodSchemaUrn) {
	o.Schemas = v
}

// GetVaultRoleID returns the VaultRoleID field value
func (o *AddAppRoleVaultAuthenticationMethodRequest) GetVaultRoleID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VaultRoleID
}

// GetVaultRoleIDOk returns a tuple with the VaultRoleID field value
// and a boolean to check if the value has been set.
func (o *AddAppRoleVaultAuthenticationMethodRequest) GetVaultRoleIDOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.VaultRoleID, true
}

// SetVaultRoleID sets field value
func (o *AddAppRoleVaultAuthenticationMethodRequest) SetVaultRoleID(v string) {
	o.VaultRoleID = v
}

// GetVaultSecretID returns the VaultSecretID field value
func (o *AddAppRoleVaultAuthenticationMethodRequest) GetVaultSecretID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VaultSecretID
}

// GetVaultSecretIDOk returns a tuple with the VaultSecretID field value
// and a boolean to check if the value has been set.
func (o *AddAppRoleVaultAuthenticationMethodRequest) GetVaultSecretIDOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.VaultSecretID, true
}

// SetVaultSecretID sets field value
func (o *AddAppRoleVaultAuthenticationMethodRequest) SetVaultSecretID(v string) {
	o.VaultSecretID = v
}

// GetLoginMechanismName returns the LoginMechanismName field value if set, zero value otherwise.
func (o *AddAppRoleVaultAuthenticationMethodRequest) GetLoginMechanismName() string {
	if o == nil || isNil(o.LoginMechanismName) {
		var ret string
		return ret
	}
	return *o.LoginMechanismName
}

// GetLoginMechanismNameOk returns a tuple with the LoginMechanismName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddAppRoleVaultAuthenticationMethodRequest) GetLoginMechanismNameOk() (*string, bool) {
	if o == nil || isNil(o.LoginMechanismName) {
    return nil, false
	}
	return o.LoginMechanismName, true
}

// HasLoginMechanismName returns a boolean if a field has been set.
func (o *AddAppRoleVaultAuthenticationMethodRequest) HasLoginMechanismName() bool {
	if o != nil && !isNil(o.LoginMechanismName) {
		return true
	}

	return false
}

// SetLoginMechanismName gets a reference to the given string and assigns it to the LoginMechanismName field.
func (o *AddAppRoleVaultAuthenticationMethodRequest) SetLoginMechanismName(v string) {
	o.LoginMechanismName = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddAppRoleVaultAuthenticationMethodRequest) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddAppRoleVaultAuthenticationMethodRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddAppRoleVaultAuthenticationMethodRequest) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddAppRoleVaultAuthenticationMethodRequest) SetDescription(v string) {
	o.Description = &v
}

func (o AddAppRoleVaultAuthenticationMethodRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["methodName"] = o.MethodName
	}
	if true {
		toSerialize["schemas"] = o.Schemas
	}
	if true {
		toSerialize["vaultRoleID"] = o.VaultRoleID
	}
	if true {
		toSerialize["vaultSecretID"] = o.VaultSecretID
	}
	if !isNil(o.LoginMechanismName) {
		toSerialize["loginMechanismName"] = o.LoginMechanismName
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullableAddAppRoleVaultAuthenticationMethodRequest struct {
	value *AddAppRoleVaultAuthenticationMethodRequest
	isSet bool
}

func (v NullableAddAppRoleVaultAuthenticationMethodRequest) Get() *AddAppRoleVaultAuthenticationMethodRequest {
	return v.value
}

func (v *NullableAddAppRoleVaultAuthenticationMethodRequest) Set(val *AddAppRoleVaultAuthenticationMethodRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddAppRoleVaultAuthenticationMethodRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddAppRoleVaultAuthenticationMethodRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddAppRoleVaultAuthenticationMethodRequest(val *AddAppRoleVaultAuthenticationMethodRequest) *NullableAddAppRoleVaultAuthenticationMethodRequest {
	return &NullableAddAppRoleVaultAuthenticationMethodRequest{value: val, isSet: true}
}

func (v NullableAddAppRoleVaultAuthenticationMethodRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddAppRoleVaultAuthenticationMethodRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

