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

// checks if the AddAes256PasswordStorageSchemeRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddAes256PasswordStorageSchemeRequest{}

// AddAes256PasswordStorageSchemeRequest struct for AddAes256PasswordStorageSchemeRequest
type AddAes256PasswordStorageSchemeRequest struct {
	Schemas []Enumaes256PasswordStorageSchemeSchemaUrn `json:"schemas"`
	// The identifier for the encryption settings definition that should be used to derive the encryption key to use when encrypting new passwords. If this is not provided, the server's preferred encryption settings definition will be used.
	EncryptionSettingsDefinitionID *string `json:"encryptionSettingsDefinitionID,omitempty"`
	// A description for this Password Storage Scheme
	Description *string `json:"description,omitempty"`
	// Indicates whether the Password Storage Scheme is enabled for use.
	Enabled bool `json:"enabled"`
	// Name of the new Password Storage Scheme
	SchemeName string `json:"schemeName"`
}

// NewAddAes256PasswordStorageSchemeRequest instantiates a new AddAes256PasswordStorageSchemeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddAes256PasswordStorageSchemeRequest(schemas []Enumaes256PasswordStorageSchemeSchemaUrn, enabled bool, schemeName string) *AddAes256PasswordStorageSchemeRequest {
	this := AddAes256PasswordStorageSchemeRequest{}
	this.Schemas = schemas
	this.Enabled = enabled
	this.SchemeName = schemeName
	return &this
}

// NewAddAes256PasswordStorageSchemeRequestWithDefaults instantiates a new AddAes256PasswordStorageSchemeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddAes256PasswordStorageSchemeRequestWithDefaults() *AddAes256PasswordStorageSchemeRequest {
	this := AddAes256PasswordStorageSchemeRequest{}
	return &this
}

// GetSchemas returns the Schemas field value
func (o *AddAes256PasswordStorageSchemeRequest) GetSchemas() []Enumaes256PasswordStorageSchemeSchemaUrn {
	if o == nil {
		var ret []Enumaes256PasswordStorageSchemeSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddAes256PasswordStorageSchemeRequest) GetSchemasOk() ([]Enumaes256PasswordStorageSchemeSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddAes256PasswordStorageSchemeRequest) SetSchemas(v []Enumaes256PasswordStorageSchemeSchemaUrn) {
	o.Schemas = v
}

// GetEncryptionSettingsDefinitionID returns the EncryptionSettingsDefinitionID field value if set, zero value otherwise.
func (o *AddAes256PasswordStorageSchemeRequest) GetEncryptionSettingsDefinitionID() string {
	if o == nil || IsNil(o.EncryptionSettingsDefinitionID) {
		var ret string
		return ret
	}
	return *o.EncryptionSettingsDefinitionID
}

// GetEncryptionSettingsDefinitionIDOk returns a tuple with the EncryptionSettingsDefinitionID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddAes256PasswordStorageSchemeRequest) GetEncryptionSettingsDefinitionIDOk() (*string, bool) {
	if o == nil || IsNil(o.EncryptionSettingsDefinitionID) {
		return nil, false
	}
	return o.EncryptionSettingsDefinitionID, true
}

// HasEncryptionSettingsDefinitionID returns a boolean if a field has been set.
func (o *AddAes256PasswordStorageSchemeRequest) HasEncryptionSettingsDefinitionID() bool {
	if o != nil && !IsNil(o.EncryptionSettingsDefinitionID) {
		return true
	}

	return false
}

// SetEncryptionSettingsDefinitionID gets a reference to the given string and assigns it to the EncryptionSettingsDefinitionID field.
func (o *AddAes256PasswordStorageSchemeRequest) SetEncryptionSettingsDefinitionID(v string) {
	o.EncryptionSettingsDefinitionID = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddAes256PasswordStorageSchemeRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddAes256PasswordStorageSchemeRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddAes256PasswordStorageSchemeRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddAes256PasswordStorageSchemeRequest) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *AddAes256PasswordStorageSchemeRequest) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AddAes256PasswordStorageSchemeRequest) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AddAes256PasswordStorageSchemeRequest) SetEnabled(v bool) {
	o.Enabled = v
}

// GetSchemeName returns the SchemeName field value
func (o *AddAes256PasswordStorageSchemeRequest) GetSchemeName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SchemeName
}

// GetSchemeNameOk returns a tuple with the SchemeName field value
// and a boolean to check if the value has been set.
func (o *AddAes256PasswordStorageSchemeRequest) GetSchemeNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SchemeName, true
}

// SetSchemeName sets field value
func (o *AddAes256PasswordStorageSchemeRequest) SetSchemeName(v string) {
	o.SchemeName = v
}

func (o AddAes256PasswordStorageSchemeRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddAes256PasswordStorageSchemeRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["schemas"] = o.Schemas
	if !IsNil(o.EncryptionSettingsDefinitionID) {
		toSerialize["encryptionSettingsDefinitionID"] = o.EncryptionSettingsDefinitionID
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["enabled"] = o.Enabled
	toSerialize["schemeName"] = o.SchemeName
	return toSerialize, nil
}

type NullableAddAes256PasswordStorageSchemeRequest struct {
	value *AddAes256PasswordStorageSchemeRequest
	isSet bool
}

func (v NullableAddAes256PasswordStorageSchemeRequest) Get() *AddAes256PasswordStorageSchemeRequest {
	return v.value
}

func (v *NullableAddAes256PasswordStorageSchemeRequest) Set(val *AddAes256PasswordStorageSchemeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddAes256PasswordStorageSchemeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddAes256PasswordStorageSchemeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddAes256PasswordStorageSchemeRequest(val *AddAes256PasswordStorageSchemeRequest) *NullableAddAes256PasswordStorageSchemeRequest {
	return &NullableAddAes256PasswordStorageSchemeRequest{value: val, isSet: true}
}

func (v NullableAddAes256PasswordStorageSchemeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddAes256PasswordStorageSchemeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
