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

// checks if the SaltedSha1PasswordStorageSchemeResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SaltedSha1PasswordStorageSchemeResponse{}

// SaltedSha1PasswordStorageSchemeResponse struct for SaltedSha1PasswordStorageSchemeResponse
type SaltedSha1PasswordStorageSchemeResponse struct {
	Schemas []EnumsaltedSha1PasswordStorageSchemeSchemaUrn `json:"schemas"`
	// Name of the Password Storage Scheme
	Id string `json:"id"`
	// Indicates whether the Salted SHA1 Password Storage Scheme is enabled for use.
	Enabled bool `json:"enabled"`
	// Specifies the number of bytes to use for the generated salt.
	SaltLengthBytes *int64 `json:"saltLengthBytes,omitempty"`
	// A description for this Password Storage Scheme
	Description                                   *string                                            `json:"description,omitempty"`
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
}

// NewSaltedSha1PasswordStorageSchemeResponse instantiates a new SaltedSha1PasswordStorageSchemeResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSaltedSha1PasswordStorageSchemeResponse(schemas []EnumsaltedSha1PasswordStorageSchemeSchemaUrn, id string, enabled bool) *SaltedSha1PasswordStorageSchemeResponse {
	this := SaltedSha1PasswordStorageSchemeResponse{}
	this.Schemas = schemas
	this.Id = id
	this.Enabled = enabled
	return &this
}

// NewSaltedSha1PasswordStorageSchemeResponseWithDefaults instantiates a new SaltedSha1PasswordStorageSchemeResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSaltedSha1PasswordStorageSchemeResponseWithDefaults() *SaltedSha1PasswordStorageSchemeResponse {
	this := SaltedSha1PasswordStorageSchemeResponse{}
	return &this
}

// GetSchemas returns the Schemas field value
func (o *SaltedSha1PasswordStorageSchemeResponse) GetSchemas() []EnumsaltedSha1PasswordStorageSchemeSchemaUrn {
	if o == nil {
		var ret []EnumsaltedSha1PasswordStorageSchemeSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *SaltedSha1PasswordStorageSchemeResponse) GetSchemasOk() ([]EnumsaltedSha1PasswordStorageSchemeSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *SaltedSha1PasswordStorageSchemeResponse) SetSchemas(v []EnumsaltedSha1PasswordStorageSchemeSchemaUrn) {
	o.Schemas = v
}

// GetId returns the Id field value
func (o *SaltedSha1PasswordStorageSchemeResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SaltedSha1PasswordStorageSchemeResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SaltedSha1PasswordStorageSchemeResponse) SetId(v string) {
	o.Id = v
}

// GetEnabled returns the Enabled field value
func (o *SaltedSha1PasswordStorageSchemeResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *SaltedSha1PasswordStorageSchemeResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *SaltedSha1PasswordStorageSchemeResponse) SetEnabled(v bool) {
	o.Enabled = v
}

// GetSaltLengthBytes returns the SaltLengthBytes field value if set, zero value otherwise.
func (o *SaltedSha1PasswordStorageSchemeResponse) GetSaltLengthBytes() int64 {
	if o == nil || IsNil(o.SaltLengthBytes) {
		var ret int64
		return ret
	}
	return *o.SaltLengthBytes
}

// GetSaltLengthBytesOk returns a tuple with the SaltLengthBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SaltedSha1PasswordStorageSchemeResponse) GetSaltLengthBytesOk() (*int64, bool) {
	if o == nil || IsNil(o.SaltLengthBytes) {
		return nil, false
	}
	return o.SaltLengthBytes, true
}

// HasSaltLengthBytes returns a boolean if a field has been set.
func (o *SaltedSha1PasswordStorageSchemeResponse) HasSaltLengthBytes() bool {
	if o != nil && !IsNil(o.SaltLengthBytes) {
		return true
	}

	return false
}

// SetSaltLengthBytes gets a reference to the given int64 and assigns it to the SaltLengthBytes field.
func (o *SaltedSha1PasswordStorageSchemeResponse) SetSaltLengthBytes(v int64) {
	o.SaltLengthBytes = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SaltedSha1PasswordStorageSchemeResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SaltedSha1PasswordStorageSchemeResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SaltedSha1PasswordStorageSchemeResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SaltedSha1PasswordStorageSchemeResponse) SetDescription(v string) {
	o.Description = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *SaltedSha1PasswordStorageSchemeResponse) GetMeta() MetaMeta {
	if o == nil || IsNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SaltedSha1PasswordStorageSchemeResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *SaltedSha1PasswordStorageSchemeResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *SaltedSha1PasswordStorageSchemeResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *SaltedSha1PasswordStorageSchemeResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SaltedSha1PasswordStorageSchemeResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *SaltedSha1PasswordStorageSchemeResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *SaltedSha1PasswordStorageSchemeResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

func (o SaltedSha1PasswordStorageSchemeResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SaltedSha1PasswordStorageSchemeResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["schemas"] = o.Schemas
	toSerialize["id"] = o.Id
	toSerialize["enabled"] = o.Enabled
	if !IsNil(o.SaltLengthBytes) {
		toSerialize["saltLengthBytes"] = o.SaltLengthBytes
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		toSerialize["urn:pingidentity:schemas:configuration:messages:2.0"] = o.Urnpingidentityschemasconfigurationmessages20
	}
	return toSerialize, nil
}

type NullableSaltedSha1PasswordStorageSchemeResponse struct {
	value *SaltedSha1PasswordStorageSchemeResponse
	isSet bool
}

func (v NullableSaltedSha1PasswordStorageSchemeResponse) Get() *SaltedSha1PasswordStorageSchemeResponse {
	return v.value
}

func (v *NullableSaltedSha1PasswordStorageSchemeResponse) Set(val *SaltedSha1PasswordStorageSchemeResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSaltedSha1PasswordStorageSchemeResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSaltedSha1PasswordStorageSchemeResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSaltedSha1PasswordStorageSchemeResponse(val *SaltedSha1PasswordStorageSchemeResponse) *NullableSaltedSha1PasswordStorageSchemeResponse {
	return &NullableSaltedSha1PasswordStorageSchemeResponse{value: val, isSet: true}
}

func (v NullableSaltedSha1PasswordStorageSchemeResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSaltedSha1PasswordStorageSchemeResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
