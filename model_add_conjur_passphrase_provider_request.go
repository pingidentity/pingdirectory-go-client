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

// AddConjurPassphraseProviderRequest struct for AddConjurPassphraseProviderRequest
type AddConjurPassphraseProviderRequest struct {
	// Name of the new Passphrase Provider
	ProviderName string                                  `json:"providerName"`
	Schemas      []EnumconjurPassphraseProviderSchemaUrn `json:"schemas"`
	// An external server definition with information needed to connect and authenticate to the Conjur instance containing the passphrase.
	ConjurExternalServer string `json:"conjurExternalServer"`
	// The portion of the path that follows the account name in the URI needed to obtain the desired secret. Any special characters in the path must be URL-encoded.
	ConjurSecretRelativePath string `json:"conjurSecretRelativePath"`
	// The maximum length of time that the passphrase provider may cache the passphrase that has been read from Conjur. A value of zero seconds indicates that the provider should always attempt to read the passphrase from Conjur.
	MaxCacheDuration *string `json:"maxCacheDuration,omitempty"`
	// A description for this Passphrase Provider
	Description *string `json:"description,omitempty"`
	// Indicates whether this Passphrase Provider is enabled for use in the server.
	Enabled bool `json:"enabled"`
}

// NewAddConjurPassphraseProviderRequest instantiates a new AddConjurPassphraseProviderRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddConjurPassphraseProviderRequest(providerName string, schemas []EnumconjurPassphraseProviderSchemaUrn, conjurExternalServer string, conjurSecretRelativePath string, enabled bool) *AddConjurPassphraseProviderRequest {
	this := AddConjurPassphraseProviderRequest{}
	this.ProviderName = providerName
	this.Schemas = schemas
	this.ConjurExternalServer = conjurExternalServer
	this.ConjurSecretRelativePath = conjurSecretRelativePath
	this.Enabled = enabled
	return &this
}

// NewAddConjurPassphraseProviderRequestWithDefaults instantiates a new AddConjurPassphraseProviderRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddConjurPassphraseProviderRequestWithDefaults() *AddConjurPassphraseProviderRequest {
	this := AddConjurPassphraseProviderRequest{}
	return &this
}

// GetProviderName returns the ProviderName field value
func (o *AddConjurPassphraseProviderRequest) GetProviderName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProviderName
}

// GetProviderNameOk returns a tuple with the ProviderName field value
// and a boolean to check if the value has been set.
func (o *AddConjurPassphraseProviderRequest) GetProviderNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProviderName, true
}

// SetProviderName sets field value
func (o *AddConjurPassphraseProviderRequest) SetProviderName(v string) {
	o.ProviderName = v
}

// GetSchemas returns the Schemas field value
func (o *AddConjurPassphraseProviderRequest) GetSchemas() []EnumconjurPassphraseProviderSchemaUrn {
	if o == nil {
		var ret []EnumconjurPassphraseProviderSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddConjurPassphraseProviderRequest) GetSchemasOk() ([]EnumconjurPassphraseProviderSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddConjurPassphraseProviderRequest) SetSchemas(v []EnumconjurPassphraseProviderSchemaUrn) {
	o.Schemas = v
}

// GetConjurExternalServer returns the ConjurExternalServer field value
func (o *AddConjurPassphraseProviderRequest) GetConjurExternalServer() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConjurExternalServer
}

// GetConjurExternalServerOk returns a tuple with the ConjurExternalServer field value
// and a boolean to check if the value has been set.
func (o *AddConjurPassphraseProviderRequest) GetConjurExternalServerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConjurExternalServer, true
}

// SetConjurExternalServer sets field value
func (o *AddConjurPassphraseProviderRequest) SetConjurExternalServer(v string) {
	o.ConjurExternalServer = v
}

// GetConjurSecretRelativePath returns the ConjurSecretRelativePath field value
func (o *AddConjurPassphraseProviderRequest) GetConjurSecretRelativePath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConjurSecretRelativePath
}

// GetConjurSecretRelativePathOk returns a tuple with the ConjurSecretRelativePath field value
// and a boolean to check if the value has been set.
func (o *AddConjurPassphraseProviderRequest) GetConjurSecretRelativePathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConjurSecretRelativePath, true
}

// SetConjurSecretRelativePath sets field value
func (o *AddConjurPassphraseProviderRequest) SetConjurSecretRelativePath(v string) {
	o.ConjurSecretRelativePath = v
}

// GetMaxCacheDuration returns the MaxCacheDuration field value if set, zero value otherwise.
func (o *AddConjurPassphraseProviderRequest) GetMaxCacheDuration() string {
	if o == nil || isNil(o.MaxCacheDuration) {
		var ret string
		return ret
	}
	return *o.MaxCacheDuration
}

// GetMaxCacheDurationOk returns a tuple with the MaxCacheDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddConjurPassphraseProviderRequest) GetMaxCacheDurationOk() (*string, bool) {
	if o == nil || isNil(o.MaxCacheDuration) {
		return nil, false
	}
	return o.MaxCacheDuration, true
}

// HasMaxCacheDuration returns a boolean if a field has been set.
func (o *AddConjurPassphraseProviderRequest) HasMaxCacheDuration() bool {
	if o != nil && !isNil(o.MaxCacheDuration) {
		return true
	}

	return false
}

// SetMaxCacheDuration gets a reference to the given string and assigns it to the MaxCacheDuration field.
func (o *AddConjurPassphraseProviderRequest) SetMaxCacheDuration(v string) {
	o.MaxCacheDuration = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddConjurPassphraseProviderRequest) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddConjurPassphraseProviderRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddConjurPassphraseProviderRequest) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddConjurPassphraseProviderRequest) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *AddConjurPassphraseProviderRequest) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AddConjurPassphraseProviderRequest) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AddConjurPassphraseProviderRequest) SetEnabled(v bool) {
	o.Enabled = v
}

func (o AddConjurPassphraseProviderRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["providerName"] = o.ProviderName
	}
	if true {
		toSerialize["schemas"] = o.Schemas
	}
	if true {
		toSerialize["conjurExternalServer"] = o.ConjurExternalServer
	}
	if true {
		toSerialize["conjurSecretRelativePath"] = o.ConjurSecretRelativePath
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

type NullableAddConjurPassphraseProviderRequest struct {
	value *AddConjurPassphraseProviderRequest
	isSet bool
}

func (v NullableAddConjurPassphraseProviderRequest) Get() *AddConjurPassphraseProviderRequest {
	return v.value
}

func (v *NullableAddConjurPassphraseProviderRequest) Set(val *AddConjurPassphraseProviderRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddConjurPassphraseProviderRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddConjurPassphraseProviderRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddConjurPassphraseProviderRequest(val *AddConjurPassphraseProviderRequest) *NullableAddConjurPassphraseProviderRequest {
	return &NullableAddConjurPassphraseProviderRequest{value: val, isSet: true}
}

func (v NullableAddConjurPassphraseProviderRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddConjurPassphraseProviderRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
