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

// checks if the CannedResponseWorkQueueResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CannedResponseWorkQueueResponse{}

// CannedResponseWorkQueueResponse struct for CannedResponseWorkQueueResponse
type CannedResponseWorkQueueResponse struct {
	Schemas                                       []EnumcannedResponseWorkQueueSchemaUrn             `json:"schemas"`
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
}

// NewCannedResponseWorkQueueResponse instantiates a new CannedResponseWorkQueueResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCannedResponseWorkQueueResponse(schemas []EnumcannedResponseWorkQueueSchemaUrn) *CannedResponseWorkQueueResponse {
	this := CannedResponseWorkQueueResponse{}
	this.Schemas = schemas
	return &this
}

// NewCannedResponseWorkQueueResponseWithDefaults instantiates a new CannedResponseWorkQueueResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCannedResponseWorkQueueResponseWithDefaults() *CannedResponseWorkQueueResponse {
	this := CannedResponseWorkQueueResponse{}
	return &this
}

// GetSchemas returns the Schemas field value
func (o *CannedResponseWorkQueueResponse) GetSchemas() []EnumcannedResponseWorkQueueSchemaUrn {
	if o == nil {
		var ret []EnumcannedResponseWorkQueueSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *CannedResponseWorkQueueResponse) GetSchemasOk() ([]EnumcannedResponseWorkQueueSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *CannedResponseWorkQueueResponse) SetSchemas(v []EnumcannedResponseWorkQueueSchemaUrn) {
	o.Schemas = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *CannedResponseWorkQueueResponse) GetMeta() MetaMeta {
	if o == nil || IsNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CannedResponseWorkQueueResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *CannedResponseWorkQueueResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *CannedResponseWorkQueueResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *CannedResponseWorkQueueResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CannedResponseWorkQueueResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *CannedResponseWorkQueueResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *CannedResponseWorkQueueResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

func (o CannedResponseWorkQueueResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CannedResponseWorkQueueResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["schemas"] = o.Schemas
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		toSerialize["urn:pingidentity:schemas:configuration:messages:2.0"] = o.Urnpingidentityschemasconfigurationmessages20
	}
	return toSerialize, nil
}

type NullableCannedResponseWorkQueueResponse struct {
	value *CannedResponseWorkQueueResponse
	isSet bool
}

func (v NullableCannedResponseWorkQueueResponse) Get() *CannedResponseWorkQueueResponse {
	return v.value
}

func (v *NullableCannedResponseWorkQueueResponse) Set(val *CannedResponseWorkQueueResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCannedResponseWorkQueueResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCannedResponseWorkQueueResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCannedResponseWorkQueueResponse(val *CannedResponseWorkQueueResponse) *NullableCannedResponseWorkQueueResponse {
	return &NullableCannedResponseWorkQueueResponse{value: val, isSet: true}
}

func (v NullableCannedResponseWorkQueueResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCannedResponseWorkQueueResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
