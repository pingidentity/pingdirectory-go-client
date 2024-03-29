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

// checks if the SaltedSha256PasswordStorageSchemeResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SaltedSha256PasswordStorageSchemeResponse{}

// SaltedSha256PasswordStorageSchemeResponse struct for SaltedSha256PasswordStorageSchemeResponse
type SaltedSha256PasswordStorageSchemeResponse struct {
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
	Schemas                                       []EnumsaltedSha256PasswordStorageSchemeSchemaUrn   `json:"schemas"`
	// Name of the Password Storage Scheme
	Id string `json:"id"`
	// Specifies the number of bytes to use for the generated salt.
	SaltLengthBytes *int64 `json:"saltLengthBytes,omitempty"`
	// A description for this Password Storage Scheme
	Description *string `json:"description,omitempty"`
	// Indicates whether the Password Storage Scheme is enabled for use.
	Enabled bool `json:"enabled"`
}

// NewSaltedSha256PasswordStorageSchemeResponse instantiates a new SaltedSha256PasswordStorageSchemeResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSaltedSha256PasswordStorageSchemeResponse(schemas []EnumsaltedSha256PasswordStorageSchemeSchemaUrn, id string, enabled bool) *SaltedSha256PasswordStorageSchemeResponse {
	this := SaltedSha256PasswordStorageSchemeResponse{}
	this.Schemas = schemas
	this.Id = id
	this.Enabled = enabled
	return &this
}

// NewSaltedSha256PasswordStorageSchemeResponseWithDefaults instantiates a new SaltedSha256PasswordStorageSchemeResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSaltedSha256PasswordStorageSchemeResponseWithDefaults() *SaltedSha256PasswordStorageSchemeResponse {
	this := SaltedSha256PasswordStorageSchemeResponse{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *SaltedSha256PasswordStorageSchemeResponse) GetMeta() MetaMeta {
	if o == nil || IsNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SaltedSha256PasswordStorageSchemeResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *SaltedSha256PasswordStorageSchemeResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *SaltedSha256PasswordStorageSchemeResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *SaltedSha256PasswordStorageSchemeResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SaltedSha256PasswordStorageSchemeResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *SaltedSha256PasswordStorageSchemeResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *SaltedSha256PasswordStorageSchemeResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetSchemas returns the Schemas field value
func (o *SaltedSha256PasswordStorageSchemeResponse) GetSchemas() []EnumsaltedSha256PasswordStorageSchemeSchemaUrn {
	if o == nil {
		var ret []EnumsaltedSha256PasswordStorageSchemeSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *SaltedSha256PasswordStorageSchemeResponse) GetSchemasOk() ([]EnumsaltedSha256PasswordStorageSchemeSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *SaltedSha256PasswordStorageSchemeResponse) SetSchemas(v []EnumsaltedSha256PasswordStorageSchemeSchemaUrn) {
	o.Schemas = v
}

// GetId returns the Id field value
func (o *SaltedSha256PasswordStorageSchemeResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SaltedSha256PasswordStorageSchemeResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SaltedSha256PasswordStorageSchemeResponse) SetId(v string) {
	o.Id = v
}

// GetSaltLengthBytes returns the SaltLengthBytes field value if set, zero value otherwise.
func (o *SaltedSha256PasswordStorageSchemeResponse) GetSaltLengthBytes() int64 {
	if o == nil || IsNil(o.SaltLengthBytes) {
		var ret int64
		return ret
	}
	return *o.SaltLengthBytes
}

// GetSaltLengthBytesOk returns a tuple with the SaltLengthBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SaltedSha256PasswordStorageSchemeResponse) GetSaltLengthBytesOk() (*int64, bool) {
	if o == nil || IsNil(o.SaltLengthBytes) {
		return nil, false
	}
	return o.SaltLengthBytes, true
}

// HasSaltLengthBytes returns a boolean if a field has been set.
func (o *SaltedSha256PasswordStorageSchemeResponse) HasSaltLengthBytes() bool {
	if o != nil && !IsNil(o.SaltLengthBytes) {
		return true
	}

	return false
}

// SetSaltLengthBytes gets a reference to the given int64 and assigns it to the SaltLengthBytes field.
func (o *SaltedSha256PasswordStorageSchemeResponse) SetSaltLengthBytes(v int64) {
	o.SaltLengthBytes = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SaltedSha256PasswordStorageSchemeResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SaltedSha256PasswordStorageSchemeResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SaltedSha256PasswordStorageSchemeResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SaltedSha256PasswordStorageSchemeResponse) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *SaltedSha256PasswordStorageSchemeResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *SaltedSha256PasswordStorageSchemeResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *SaltedSha256PasswordStorageSchemeResponse) SetEnabled(v bool) {
	o.Enabled = v
}

func (o SaltedSha256PasswordStorageSchemeResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SaltedSha256PasswordStorageSchemeResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		toSerialize["urn:pingidentity:schemas:configuration:messages:2.0"] = o.Urnpingidentityschemasconfigurationmessages20
	}
	toSerialize["schemas"] = o.Schemas
	toSerialize["id"] = o.Id
	if !IsNil(o.SaltLengthBytes) {
		toSerialize["saltLengthBytes"] = o.SaltLengthBytes
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["enabled"] = o.Enabled
	return toSerialize, nil
}

type NullableSaltedSha256PasswordStorageSchemeResponse struct {
	value *SaltedSha256PasswordStorageSchemeResponse
	isSet bool
}

func (v NullableSaltedSha256PasswordStorageSchemeResponse) Get() *SaltedSha256PasswordStorageSchemeResponse {
	return v.value
}

func (v *NullableSaltedSha256PasswordStorageSchemeResponse) Set(val *SaltedSha256PasswordStorageSchemeResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSaltedSha256PasswordStorageSchemeResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSaltedSha256PasswordStorageSchemeResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSaltedSha256PasswordStorageSchemeResponse(val *SaltedSha256PasswordStorageSchemeResponse) *NullableSaltedSha256PasswordStorageSchemeResponse {
	return &NullableSaltedSha256PasswordStorageSchemeResponse{value: val, isSet: true}
}

func (v NullableSaltedSha256PasswordStorageSchemeResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSaltedSha256PasswordStorageSchemeResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
