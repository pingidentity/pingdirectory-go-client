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

// checks if the Aes256PasswordStorageSchemeResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Aes256PasswordStorageSchemeResponse{}

// Aes256PasswordStorageSchemeResponse struct for Aes256PasswordStorageSchemeResponse
type Aes256PasswordStorageSchemeResponse struct {
	Schemas []Enumaes256PasswordStorageSchemeSchemaUrn `json:"schemas"`
	// The identifier for the encryption settings definition that should be used to derive the encryption key to use when encrypting new passwords. If this is not provided, the server's preferred encryption settings definition will be used.
	EncryptionSettingsDefinitionID *string `json:"encryptionSettingsDefinitionID,omitempty"`
	// A description for this Password Storage Scheme
	Description *string `json:"description,omitempty"`
	// Indicates whether the Password Storage Scheme is enabled for use.
	Enabled                                       bool                                               `json:"enabled"`
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
	// Name of the Password Storage Scheme
	Id string `json:"id"`
}

// NewAes256PasswordStorageSchemeResponse instantiates a new Aes256PasswordStorageSchemeResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAes256PasswordStorageSchemeResponse(schemas []Enumaes256PasswordStorageSchemeSchemaUrn, enabled bool, id string) *Aes256PasswordStorageSchemeResponse {
	this := Aes256PasswordStorageSchemeResponse{}
	this.Schemas = schemas
	this.Enabled = enabled
	this.Id = id
	return &this
}

// NewAes256PasswordStorageSchemeResponseWithDefaults instantiates a new Aes256PasswordStorageSchemeResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAes256PasswordStorageSchemeResponseWithDefaults() *Aes256PasswordStorageSchemeResponse {
	this := Aes256PasswordStorageSchemeResponse{}
	return &this
}

// GetSchemas returns the Schemas field value
func (o *Aes256PasswordStorageSchemeResponse) GetSchemas() []Enumaes256PasswordStorageSchemeSchemaUrn {
	if o == nil {
		var ret []Enumaes256PasswordStorageSchemeSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *Aes256PasswordStorageSchemeResponse) GetSchemasOk() ([]Enumaes256PasswordStorageSchemeSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *Aes256PasswordStorageSchemeResponse) SetSchemas(v []Enumaes256PasswordStorageSchemeSchemaUrn) {
	o.Schemas = v
}

// GetEncryptionSettingsDefinitionID returns the EncryptionSettingsDefinitionID field value if set, zero value otherwise.
func (o *Aes256PasswordStorageSchemeResponse) GetEncryptionSettingsDefinitionID() string {
	if o == nil || IsNil(o.EncryptionSettingsDefinitionID) {
		var ret string
		return ret
	}
	return *o.EncryptionSettingsDefinitionID
}

// GetEncryptionSettingsDefinitionIDOk returns a tuple with the EncryptionSettingsDefinitionID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Aes256PasswordStorageSchemeResponse) GetEncryptionSettingsDefinitionIDOk() (*string, bool) {
	if o == nil || IsNil(o.EncryptionSettingsDefinitionID) {
		return nil, false
	}
	return o.EncryptionSettingsDefinitionID, true
}

// HasEncryptionSettingsDefinitionID returns a boolean if a field has been set.
func (o *Aes256PasswordStorageSchemeResponse) HasEncryptionSettingsDefinitionID() bool {
	if o != nil && !IsNil(o.EncryptionSettingsDefinitionID) {
		return true
	}

	return false
}

// SetEncryptionSettingsDefinitionID gets a reference to the given string and assigns it to the EncryptionSettingsDefinitionID field.
func (o *Aes256PasswordStorageSchemeResponse) SetEncryptionSettingsDefinitionID(v string) {
	o.EncryptionSettingsDefinitionID = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Aes256PasswordStorageSchemeResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Aes256PasswordStorageSchemeResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Aes256PasswordStorageSchemeResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Aes256PasswordStorageSchemeResponse) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *Aes256PasswordStorageSchemeResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *Aes256PasswordStorageSchemeResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *Aes256PasswordStorageSchemeResponse) SetEnabled(v bool) {
	o.Enabled = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *Aes256PasswordStorageSchemeResponse) GetMeta() MetaMeta {
	if o == nil || IsNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Aes256PasswordStorageSchemeResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *Aes256PasswordStorageSchemeResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *Aes256PasswordStorageSchemeResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *Aes256PasswordStorageSchemeResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Aes256PasswordStorageSchemeResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *Aes256PasswordStorageSchemeResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *Aes256PasswordStorageSchemeResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetId returns the Id field value
func (o *Aes256PasswordStorageSchemeResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Aes256PasswordStorageSchemeResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Aes256PasswordStorageSchemeResponse) SetId(v string) {
	o.Id = v
}

func (o Aes256PasswordStorageSchemeResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Aes256PasswordStorageSchemeResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["schemas"] = o.Schemas
	if !IsNil(o.EncryptionSettingsDefinitionID) {
		toSerialize["encryptionSettingsDefinitionID"] = o.EncryptionSettingsDefinitionID
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["enabled"] = o.Enabled
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		toSerialize["urn:pingidentity:schemas:configuration:messages:2.0"] = o.Urnpingidentityschemasconfigurationmessages20
	}
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

type NullableAes256PasswordStorageSchemeResponse struct {
	value *Aes256PasswordStorageSchemeResponse
	isSet bool
}

func (v NullableAes256PasswordStorageSchemeResponse) Get() *Aes256PasswordStorageSchemeResponse {
	return v.value
}

func (v *NullableAes256PasswordStorageSchemeResponse) Set(val *Aes256PasswordStorageSchemeResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAes256PasswordStorageSchemeResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAes256PasswordStorageSchemeResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAes256PasswordStorageSchemeResponse(val *Aes256PasswordStorageSchemeResponse) *NullableAes256PasswordStorageSchemeResponse {
	return &NullableAes256PasswordStorageSchemeResponse{value: val, isSet: true}
}

func (v NullableAes256PasswordStorageSchemeResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAes256PasswordStorageSchemeResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
