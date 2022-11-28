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

// NeverRotateLogRotationPolicyResponse struct for NeverRotateLogRotationPolicyResponse
type NeverRotateLogRotationPolicyResponse struct {
	Meta *MetaMeta `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
	// Name of the Log Rotation Policy
	Id string `json:"id"`
	Schemas []EnumneverRotateLogRotationPolicySchemaUrn `json:"schemas"`
	// A description for this Log Rotation Policy
	Description *string `json:"description,omitempty"`
}

// NewNeverRotateLogRotationPolicyResponse instantiates a new NeverRotateLogRotationPolicyResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNeverRotateLogRotationPolicyResponse(id string, schemas []EnumneverRotateLogRotationPolicySchemaUrn) *NeverRotateLogRotationPolicyResponse {
	this := NeverRotateLogRotationPolicyResponse{}
	this.Id = id
	this.Schemas = schemas
	return &this
}

// NewNeverRotateLogRotationPolicyResponseWithDefaults instantiates a new NeverRotateLogRotationPolicyResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNeverRotateLogRotationPolicyResponseWithDefaults() *NeverRotateLogRotationPolicyResponse {
	this := NeverRotateLogRotationPolicyResponse{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *NeverRotateLogRotationPolicyResponse) GetMeta() MetaMeta {
	if o == nil || isNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NeverRotateLogRotationPolicyResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || isNil(o.Meta) {
    return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *NeverRotateLogRotationPolicyResponse) HasMeta() bool {
	if o != nil && !isNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *NeverRotateLogRotationPolicyResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *NeverRotateLogRotationPolicyResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NeverRotateLogRotationPolicyResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || isNil(o.Urnpingidentityschemasconfigurationmessages20) {
    return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *NeverRotateLogRotationPolicyResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *NeverRotateLogRotationPolicyResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetId returns the Id field value
func (o *NeverRotateLogRotationPolicyResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *NeverRotateLogRotationPolicyResponse) GetIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *NeverRotateLogRotationPolicyResponse) SetId(v string) {
	o.Id = v
}

// GetSchemas returns the Schemas field value
func (o *NeverRotateLogRotationPolicyResponse) GetSchemas() []EnumneverRotateLogRotationPolicySchemaUrn {
	if o == nil {
		var ret []EnumneverRotateLogRotationPolicySchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *NeverRotateLogRotationPolicyResponse) GetSchemasOk() ([]EnumneverRotateLogRotationPolicySchemaUrn, bool) {
	if o == nil {
    return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *NeverRotateLogRotationPolicyResponse) SetSchemas(v []EnumneverRotateLogRotationPolicySchemaUrn) {
	o.Schemas = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *NeverRotateLogRotationPolicyResponse) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NeverRotateLogRotationPolicyResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *NeverRotateLogRotationPolicyResponse) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *NeverRotateLogRotationPolicyResponse) SetDescription(v string) {
	o.Description = &v
}

func (o NeverRotateLogRotationPolicyResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		toSerialize["urn:pingidentity:schemas:configuration:messages:2.0"] = o.Urnpingidentityschemasconfigurationmessages20
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["schemas"] = o.Schemas
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullableNeverRotateLogRotationPolicyResponse struct {
	value *NeverRotateLogRotationPolicyResponse
	isSet bool
}

func (v NullableNeverRotateLogRotationPolicyResponse) Get() *NeverRotateLogRotationPolicyResponse {
	return v.value
}

func (v *NullableNeverRotateLogRotationPolicyResponse) Set(val *NeverRotateLogRotationPolicyResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableNeverRotateLogRotationPolicyResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableNeverRotateLogRotationPolicyResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNeverRotateLogRotationPolicyResponse(val *NeverRotateLogRotationPolicyResponse) *NullableNeverRotateLogRotationPolicyResponse {
	return &NullableNeverRotateLogRotationPolicyResponse{value: val, isSet: true}
}

func (v NullableNeverRotateLogRotationPolicyResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNeverRotateLogRotationPolicyResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


