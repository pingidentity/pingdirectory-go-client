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

// VaultPassphraseProviderResponse struct for VaultPassphraseProviderResponse
type VaultPassphraseProviderResponse struct {
	// Name of the Passphrase Provider
	Id string `json:"id"`
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

// NewVaultPassphraseProviderResponse instantiates a new VaultPassphraseProviderResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVaultPassphraseProviderResponse(id string, schemas []EnumvaultPassphraseProviderSchemaUrn, vaultExternalServer string, vaultSecretPath string, vaultSecretFieldName string, enabled bool) *VaultPassphraseProviderResponse {
	this := VaultPassphraseProviderResponse{}
	this.Id = id
	this.Schemas = schemas
	this.VaultExternalServer = vaultExternalServer
	this.VaultSecretPath = vaultSecretPath
	this.VaultSecretFieldName = vaultSecretFieldName
	this.Enabled = enabled
	return &this
}

// NewVaultPassphraseProviderResponseWithDefaults instantiates a new VaultPassphraseProviderResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVaultPassphraseProviderResponseWithDefaults() *VaultPassphraseProviderResponse {
	this := VaultPassphraseProviderResponse{}
	return &this
}

// GetId returns the Id field value
func (o *VaultPassphraseProviderResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *VaultPassphraseProviderResponse) GetIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *VaultPassphraseProviderResponse) SetId(v string) {
	o.Id = v
}

// GetSchemas returns the Schemas field value
func (o *VaultPassphraseProviderResponse) GetSchemas() []EnumvaultPassphraseProviderSchemaUrn {
	if o == nil {
		var ret []EnumvaultPassphraseProviderSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *VaultPassphraseProviderResponse) GetSchemasOk() ([]EnumvaultPassphraseProviderSchemaUrn, bool) {
	if o == nil {
    return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *VaultPassphraseProviderResponse) SetSchemas(v []EnumvaultPassphraseProviderSchemaUrn) {
	o.Schemas = v
}

// GetVaultExternalServer returns the VaultExternalServer field value
func (o *VaultPassphraseProviderResponse) GetVaultExternalServer() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VaultExternalServer
}

// GetVaultExternalServerOk returns a tuple with the VaultExternalServer field value
// and a boolean to check if the value has been set.
func (o *VaultPassphraseProviderResponse) GetVaultExternalServerOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.VaultExternalServer, true
}

// SetVaultExternalServer sets field value
func (o *VaultPassphraseProviderResponse) SetVaultExternalServer(v string) {
	o.VaultExternalServer = v
}

// GetVaultSecretPath returns the VaultSecretPath field value
func (o *VaultPassphraseProviderResponse) GetVaultSecretPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VaultSecretPath
}

// GetVaultSecretPathOk returns a tuple with the VaultSecretPath field value
// and a boolean to check if the value has been set.
func (o *VaultPassphraseProviderResponse) GetVaultSecretPathOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.VaultSecretPath, true
}

// SetVaultSecretPath sets field value
func (o *VaultPassphraseProviderResponse) SetVaultSecretPath(v string) {
	o.VaultSecretPath = v
}

// GetVaultSecretFieldName returns the VaultSecretFieldName field value
func (o *VaultPassphraseProviderResponse) GetVaultSecretFieldName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VaultSecretFieldName
}

// GetVaultSecretFieldNameOk returns a tuple with the VaultSecretFieldName field value
// and a boolean to check if the value has been set.
func (o *VaultPassphraseProviderResponse) GetVaultSecretFieldNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.VaultSecretFieldName, true
}

// SetVaultSecretFieldName sets field value
func (o *VaultPassphraseProviderResponse) SetVaultSecretFieldName(v string) {
	o.VaultSecretFieldName = v
}

// GetMaxCacheDuration returns the MaxCacheDuration field value if set, zero value otherwise.
func (o *VaultPassphraseProviderResponse) GetMaxCacheDuration() string {
	if o == nil || isNil(o.MaxCacheDuration) {
		var ret string
		return ret
	}
	return *o.MaxCacheDuration
}

// GetMaxCacheDurationOk returns a tuple with the MaxCacheDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VaultPassphraseProviderResponse) GetMaxCacheDurationOk() (*string, bool) {
	if o == nil || isNil(o.MaxCacheDuration) {
    return nil, false
	}
	return o.MaxCacheDuration, true
}

// HasMaxCacheDuration returns a boolean if a field has been set.
func (o *VaultPassphraseProviderResponse) HasMaxCacheDuration() bool {
	if o != nil && !isNil(o.MaxCacheDuration) {
		return true
	}

	return false
}

// SetMaxCacheDuration gets a reference to the given string and assigns it to the MaxCacheDuration field.
func (o *VaultPassphraseProviderResponse) SetMaxCacheDuration(v string) {
	o.MaxCacheDuration = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *VaultPassphraseProviderResponse) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VaultPassphraseProviderResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *VaultPassphraseProviderResponse) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *VaultPassphraseProviderResponse) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *VaultPassphraseProviderResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *VaultPassphraseProviderResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *VaultPassphraseProviderResponse) SetEnabled(v bool) {
	o.Enabled = v
}

func (o VaultPassphraseProviderResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
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

type NullableVaultPassphraseProviderResponse struct {
	value *VaultPassphraseProviderResponse
	isSet bool
}

func (v NullableVaultPassphraseProviderResponse) Get() *VaultPassphraseProviderResponse {
	return v.value
}

func (v *NullableVaultPassphraseProviderResponse) Set(val *VaultPassphraseProviderResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableVaultPassphraseProviderResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableVaultPassphraseProviderResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVaultPassphraseProviderResponse(val *VaultPassphraseProviderResponse) *NullableVaultPassphraseProviderResponse {
	return &NullableVaultPassphraseProviderResponse{value: val, isSet: true}
}

func (v NullableVaultPassphraseProviderResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVaultPassphraseProviderResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


