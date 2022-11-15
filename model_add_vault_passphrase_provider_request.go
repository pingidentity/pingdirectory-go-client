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

// AddVaultPassphraseProviderRequest struct for AddVaultPassphraseProviderRequest
type AddVaultPassphraseProviderRequest struct {
	// Name of the new Passphrase Provider
	ProviderName string `json:"providerName"`
	Schemas []EnumvaultPassphraseProviderSchemaUrn `json:"schemas"`
	// An external server definition with information needed to connect and authenticate to the Vault instance containing the passphrase.
	VaultExternalServer string `json:"vaultExternalServer"`
	// The path to the desired secret in the Vault service. This will be appended to the value of the base-url property for the associated Vault external server.
	VaultSecretPath string `json:"vaultSecretPath"`
	// The name of the field in the Vault secret record that contains the passphrase to use to generate the encryption key.
	VaultSecretFieldName string `json:"vaultSecretFieldName"`
	// The maximum length of time that the passphrase provider may cache the passphrase that has been read from Vault. A value of zero seconds indicates that the provider should always attempt to read the passphrase from Vault.
	MaxCacheDuration *string `json:"maxCacheDuration,omitempty"`
	// A description for this Passphrase Provider
	Description *string `json:"description,omitempty"`
	// Indicates whether this Passphrase Provider is enabled for use in the server.
	Enabled bool `json:"enabled"`
}

// NewAddVaultPassphraseProviderRequest instantiates a new AddVaultPassphraseProviderRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddVaultPassphraseProviderRequest(providerName string, schemas []EnumvaultPassphraseProviderSchemaUrn, vaultExternalServer string, vaultSecretPath string, vaultSecretFieldName string, enabled bool) *AddVaultPassphraseProviderRequest {
	this := AddVaultPassphraseProviderRequest{}
	this.ProviderName = providerName
	this.Schemas = schemas
	this.VaultExternalServer = vaultExternalServer
	this.VaultSecretPath = vaultSecretPath
	this.VaultSecretFieldName = vaultSecretFieldName
	this.Enabled = enabled
	return &this
}

// NewAddVaultPassphraseProviderRequestWithDefaults instantiates a new AddVaultPassphraseProviderRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddVaultPassphraseProviderRequestWithDefaults() *AddVaultPassphraseProviderRequest {
	this := AddVaultPassphraseProviderRequest{}
	return &this
}

// GetProviderName returns the ProviderName field value
func (o *AddVaultPassphraseProviderRequest) GetProviderName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProviderName
}

// GetProviderNameOk returns a tuple with the ProviderName field value
// and a boolean to check if the value has been set.
func (o *AddVaultPassphraseProviderRequest) GetProviderNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ProviderName, true
}

// SetProviderName sets field value
func (o *AddVaultPassphraseProviderRequest) SetProviderName(v string) {
	o.ProviderName = v
}

// GetSchemas returns the Schemas field value
func (o *AddVaultPassphraseProviderRequest) GetSchemas() []EnumvaultPassphraseProviderSchemaUrn {
	if o == nil {
		var ret []EnumvaultPassphraseProviderSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddVaultPassphraseProviderRequest) GetSchemasOk() ([]EnumvaultPassphraseProviderSchemaUrn, bool) {
	if o == nil {
    return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddVaultPassphraseProviderRequest) SetSchemas(v []EnumvaultPassphraseProviderSchemaUrn) {
	o.Schemas = v
}

// GetVaultExternalServer returns the VaultExternalServer field value
func (o *AddVaultPassphraseProviderRequest) GetVaultExternalServer() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VaultExternalServer
}

// GetVaultExternalServerOk returns a tuple with the VaultExternalServer field value
// and a boolean to check if the value has been set.
func (o *AddVaultPassphraseProviderRequest) GetVaultExternalServerOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.VaultExternalServer, true
}

// SetVaultExternalServer sets field value
func (o *AddVaultPassphraseProviderRequest) SetVaultExternalServer(v string) {
	o.VaultExternalServer = v
}

// GetVaultSecretPath returns the VaultSecretPath field value
func (o *AddVaultPassphraseProviderRequest) GetVaultSecretPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VaultSecretPath
}

// GetVaultSecretPathOk returns a tuple with the VaultSecretPath field value
// and a boolean to check if the value has been set.
func (o *AddVaultPassphraseProviderRequest) GetVaultSecretPathOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.VaultSecretPath, true
}

// SetVaultSecretPath sets field value
func (o *AddVaultPassphraseProviderRequest) SetVaultSecretPath(v string) {
	o.VaultSecretPath = v
}

// GetVaultSecretFieldName returns the VaultSecretFieldName field value
func (o *AddVaultPassphraseProviderRequest) GetVaultSecretFieldName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VaultSecretFieldName
}

// GetVaultSecretFieldNameOk returns a tuple with the VaultSecretFieldName field value
// and a boolean to check if the value has been set.
func (o *AddVaultPassphraseProviderRequest) GetVaultSecretFieldNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.VaultSecretFieldName, true
}

// SetVaultSecretFieldName sets field value
func (o *AddVaultPassphraseProviderRequest) SetVaultSecretFieldName(v string) {
	o.VaultSecretFieldName = v
}

// GetMaxCacheDuration returns the MaxCacheDuration field value if set, zero value otherwise.
func (o *AddVaultPassphraseProviderRequest) GetMaxCacheDuration() string {
	if o == nil || isNil(o.MaxCacheDuration) {
		var ret string
		return ret
	}
	return *o.MaxCacheDuration
}

// GetMaxCacheDurationOk returns a tuple with the MaxCacheDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddVaultPassphraseProviderRequest) GetMaxCacheDurationOk() (*string, bool) {
	if o == nil || isNil(o.MaxCacheDuration) {
    return nil, false
	}
	return o.MaxCacheDuration, true
}

// HasMaxCacheDuration returns a boolean if a field has been set.
func (o *AddVaultPassphraseProviderRequest) HasMaxCacheDuration() bool {
	if o != nil && !isNil(o.MaxCacheDuration) {
		return true
	}

	return false
}

// SetMaxCacheDuration gets a reference to the given string and assigns it to the MaxCacheDuration field.
func (o *AddVaultPassphraseProviderRequest) SetMaxCacheDuration(v string) {
	o.MaxCacheDuration = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddVaultPassphraseProviderRequest) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddVaultPassphraseProviderRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddVaultPassphraseProviderRequest) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddVaultPassphraseProviderRequest) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *AddVaultPassphraseProviderRequest) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AddVaultPassphraseProviderRequest) GetEnabledOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AddVaultPassphraseProviderRequest) SetEnabled(v bool) {
	o.Enabled = v
}

func (o AddVaultPassphraseProviderRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["providerName"] = o.ProviderName
	}
	if true {
		toSerialize["schemas"] = o.Schemas
	}
	if true {
		toSerialize["vaultExternalServer"] = o.VaultExternalServer
	}
	if true {
		toSerialize["vaultSecretPath"] = o.VaultSecretPath
	}
	if true {
		toSerialize["vaultSecretFieldName"] = o.VaultSecretFieldName
	}
	if !isNil(o.MaxCacheDuration) {
		toSerialize["maxCacheDuration"] = o.MaxCacheDuration
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableAddVaultPassphraseProviderRequest struct {
	value *AddVaultPassphraseProviderRequest
	isSet bool
}

func (v NullableAddVaultPassphraseProviderRequest) Get() *AddVaultPassphraseProviderRequest {
	return v.value
}

func (v *NullableAddVaultPassphraseProviderRequest) Set(val *AddVaultPassphraseProviderRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddVaultPassphraseProviderRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddVaultPassphraseProviderRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddVaultPassphraseProviderRequest(val *AddVaultPassphraseProviderRequest) *NullableAddVaultPassphraseProviderRequest {
	return &NullableAddVaultPassphraseProviderRequest{value: val, isSet: true}
}

func (v NullableAddVaultPassphraseProviderRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddVaultPassphraseProviderRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


