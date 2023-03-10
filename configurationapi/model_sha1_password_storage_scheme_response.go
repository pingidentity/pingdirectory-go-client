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

// checks if the Sha1PasswordStorageSchemeResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Sha1PasswordStorageSchemeResponse{}

// Sha1PasswordStorageSchemeResponse struct for Sha1PasswordStorageSchemeResponse
type Sha1PasswordStorageSchemeResponse struct {
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
	Schemas                                       []Enumsha1PasswordStorageSchemeSchemaUrn           `json:"schemas"`
	// Name of the Password Storage Scheme
	Id string `json:"id"`
	// Indicates whether the SHA1 Password Storage Scheme is enabled for use.
	Enabled bool `json:"enabled"`
	// A description for this Password Storage Scheme
	Description *string `json:"description,omitempty"`
}

// NewSha1PasswordStorageSchemeResponse instantiates a new Sha1PasswordStorageSchemeResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSha1PasswordStorageSchemeResponse(schemas []Enumsha1PasswordStorageSchemeSchemaUrn, id string, enabled bool) *Sha1PasswordStorageSchemeResponse {
	this := Sha1PasswordStorageSchemeResponse{}
	this.Schemas = schemas
	this.Id = id
	this.Enabled = enabled
	return &this
}

// NewSha1PasswordStorageSchemeResponseWithDefaults instantiates a new Sha1PasswordStorageSchemeResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSha1PasswordStorageSchemeResponseWithDefaults() *Sha1PasswordStorageSchemeResponse {
	this := Sha1PasswordStorageSchemeResponse{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *Sha1PasswordStorageSchemeResponse) GetMeta() MetaMeta {
	if o == nil || IsNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Sha1PasswordStorageSchemeResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *Sha1PasswordStorageSchemeResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *Sha1PasswordStorageSchemeResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *Sha1PasswordStorageSchemeResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Sha1PasswordStorageSchemeResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *Sha1PasswordStorageSchemeResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *Sha1PasswordStorageSchemeResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetSchemas returns the Schemas field value
func (o *Sha1PasswordStorageSchemeResponse) GetSchemas() []Enumsha1PasswordStorageSchemeSchemaUrn {
	if o == nil {
		var ret []Enumsha1PasswordStorageSchemeSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *Sha1PasswordStorageSchemeResponse) GetSchemasOk() ([]Enumsha1PasswordStorageSchemeSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *Sha1PasswordStorageSchemeResponse) SetSchemas(v []Enumsha1PasswordStorageSchemeSchemaUrn) {
	o.Schemas = v
}

// GetId returns the Id field value
func (o *Sha1PasswordStorageSchemeResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Sha1PasswordStorageSchemeResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Sha1PasswordStorageSchemeResponse) SetId(v string) {
	o.Id = v
}

// GetEnabled returns the Enabled field value
func (o *Sha1PasswordStorageSchemeResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *Sha1PasswordStorageSchemeResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *Sha1PasswordStorageSchemeResponse) SetEnabled(v bool) {
	o.Enabled = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Sha1PasswordStorageSchemeResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Sha1PasswordStorageSchemeResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Sha1PasswordStorageSchemeResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Sha1PasswordStorageSchemeResponse) SetDescription(v string) {
	o.Description = &v
}

func (o Sha1PasswordStorageSchemeResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Sha1PasswordStorageSchemeResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		toSerialize["urn:pingidentity:schemas:configuration:messages:2.0"] = o.Urnpingidentityschemasconfigurationmessages20
	}
	toSerialize["schemas"] = o.Schemas
	toSerialize["id"] = o.Id
	toSerialize["enabled"] = o.Enabled
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return toSerialize, nil
}

type NullableSha1PasswordStorageSchemeResponse struct {
	value *Sha1PasswordStorageSchemeResponse
	isSet bool
}

func (v NullableSha1PasswordStorageSchemeResponse) Get() *Sha1PasswordStorageSchemeResponse {
	return v.value
}

func (v *NullableSha1PasswordStorageSchemeResponse) Set(val *Sha1PasswordStorageSchemeResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSha1PasswordStorageSchemeResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSha1PasswordStorageSchemeResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSha1PasswordStorageSchemeResponse(val *Sha1PasswordStorageSchemeResponse) *NullableSha1PasswordStorageSchemeResponse {
	return &NullableSha1PasswordStorageSchemeResponse{value: val, isSet: true}
}

func (v NullableSha1PasswordStorageSchemeResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSha1PasswordStorageSchemeResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
