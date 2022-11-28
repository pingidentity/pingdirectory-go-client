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

// TripleDesPasswordStorageSchemeResponse struct for TripleDesPasswordStorageSchemeResponse
type TripleDesPasswordStorageSchemeResponse struct {
	Meta *MetaMeta `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
	Schemas []EnumtripleDesPasswordStorageSchemeSchemaUrn `json:"schemas"`
	// Name of the Password Storage Scheme
	Id *string `json:"id,omitempty"`
	// Indicates whether the Triple DES Password Storage Scheme is enabled for use.
	Enabled bool `json:"enabled"`
	// A description for this Password Storage Scheme
	Description *string `json:"description,omitempty"`
}

// NewTripleDesPasswordStorageSchemeResponse instantiates a new TripleDesPasswordStorageSchemeResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTripleDesPasswordStorageSchemeResponse(schemas []EnumtripleDesPasswordStorageSchemeSchemaUrn, enabled bool) *TripleDesPasswordStorageSchemeResponse {
	this := TripleDesPasswordStorageSchemeResponse{}
	this.Schemas = schemas
	this.Enabled = enabled
	return &this
}

// NewTripleDesPasswordStorageSchemeResponseWithDefaults instantiates a new TripleDesPasswordStorageSchemeResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTripleDesPasswordStorageSchemeResponseWithDefaults() *TripleDesPasswordStorageSchemeResponse {
	this := TripleDesPasswordStorageSchemeResponse{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *TripleDesPasswordStorageSchemeResponse) GetMeta() MetaMeta {
	if o == nil || isNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TripleDesPasswordStorageSchemeResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || isNil(o.Meta) {
    return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *TripleDesPasswordStorageSchemeResponse) HasMeta() bool {
	if o != nil && !isNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *TripleDesPasswordStorageSchemeResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *TripleDesPasswordStorageSchemeResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TripleDesPasswordStorageSchemeResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || isNil(o.Urnpingidentityschemasconfigurationmessages20) {
    return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *TripleDesPasswordStorageSchemeResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *TripleDesPasswordStorageSchemeResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetSchemas returns the Schemas field value
func (o *TripleDesPasswordStorageSchemeResponse) GetSchemas() []EnumtripleDesPasswordStorageSchemeSchemaUrn {
	if o == nil {
		var ret []EnumtripleDesPasswordStorageSchemeSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *TripleDesPasswordStorageSchemeResponse) GetSchemasOk() ([]EnumtripleDesPasswordStorageSchemeSchemaUrn, bool) {
	if o == nil {
    return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *TripleDesPasswordStorageSchemeResponse) SetSchemas(v []EnumtripleDesPasswordStorageSchemeSchemaUrn) {
	o.Schemas = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TripleDesPasswordStorageSchemeResponse) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TripleDesPasswordStorageSchemeResponse) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TripleDesPasswordStorageSchemeResponse) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TripleDesPasswordStorageSchemeResponse) SetId(v string) {
	o.Id = &v
}

// GetEnabled returns the Enabled field value
func (o *TripleDesPasswordStorageSchemeResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *TripleDesPasswordStorageSchemeResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *TripleDesPasswordStorageSchemeResponse) SetEnabled(v bool) {
	o.Enabled = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *TripleDesPasswordStorageSchemeResponse) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TripleDesPasswordStorageSchemeResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *TripleDesPasswordStorageSchemeResponse) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *TripleDesPasswordStorageSchemeResponse) SetDescription(v string) {
	o.Description = &v
}

func (o TripleDesPasswordStorageSchemeResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		toSerialize["urn:pingidentity:schemas:configuration:messages:2.0"] = o.Urnpingidentityschemasconfigurationmessages20
	}
	if true {
		toSerialize["schemas"] = o.Schemas
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullableTripleDesPasswordStorageSchemeResponse struct {
	value *TripleDesPasswordStorageSchemeResponse
	isSet bool
}

func (v NullableTripleDesPasswordStorageSchemeResponse) Get() *TripleDesPasswordStorageSchemeResponse {
	return v.value
}

func (v *NullableTripleDesPasswordStorageSchemeResponse) Set(val *TripleDesPasswordStorageSchemeResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTripleDesPasswordStorageSchemeResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTripleDesPasswordStorageSchemeResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTripleDesPasswordStorageSchemeResponse(val *TripleDesPasswordStorageSchemeResponse) *NullableTripleDesPasswordStorageSchemeResponse {
	return &NullableTripleDesPasswordStorageSchemeResponse{value: val, isSet: true}
}

func (v NullableTripleDesPasswordStorageSchemeResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTripleDesPasswordStorageSchemeResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


