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

// TopologyAdminUserResponse struct for TopologyAdminUserResponse
type TopologyAdminUserResponse struct {
	// Name of the Topology Admin User
	Id string `json:"id"`
	Schemas []EnumtopologyAdminUserSchemaUrn `json:"schemas,omitempty"`
	Meta *MetaMeta `json:"meta,omitempty"`
}

// NewTopologyAdminUserResponse instantiates a new TopologyAdminUserResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTopologyAdminUserResponse(id string) *TopologyAdminUserResponse {
	this := TopologyAdminUserResponse{}
	this.Id = id
	return &this
}

// NewTopologyAdminUserResponseWithDefaults instantiates a new TopologyAdminUserResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTopologyAdminUserResponseWithDefaults() *TopologyAdminUserResponse {
	this := TopologyAdminUserResponse{}
	return &this
}

// GetId returns the Id field value
func (o *TopologyAdminUserResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TopologyAdminUserResponse) GetIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TopologyAdminUserResponse) SetId(v string) {
	o.Id = v
}

// GetSchemas returns the Schemas field value if set, zero value otherwise.
func (o *TopologyAdminUserResponse) GetSchemas() []EnumtopologyAdminUserSchemaUrn {
	if o == nil || isNil(o.Schemas) {
		var ret []EnumtopologyAdminUserSchemaUrn
		return ret
	}
	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TopologyAdminUserResponse) GetSchemasOk() ([]EnumtopologyAdminUserSchemaUrn, bool) {
	if o == nil || isNil(o.Schemas) {
    return nil, false
	}
	return o.Schemas, true
}

// HasSchemas returns a boolean if a field has been set.
func (o *TopologyAdminUserResponse) HasSchemas() bool {
	if o != nil && !isNil(o.Schemas) {
		return true
	}

	return false
}

// SetSchemas gets a reference to the given []EnumtopologyAdminUserSchemaUrn and assigns it to the Schemas field.
func (o *TopologyAdminUserResponse) SetSchemas(v []EnumtopologyAdminUserSchemaUrn) {
	o.Schemas = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *TopologyAdminUserResponse) GetMeta() MetaMeta {
	if o == nil || isNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TopologyAdminUserResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || isNil(o.Meta) {
    return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *TopologyAdminUserResponse) HasMeta() bool {
	if o != nil && !isNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *TopologyAdminUserResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

func (o TopologyAdminUserResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Schemas) {
		toSerialize["schemas"] = o.Schemas
	}
	if !isNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return json.Marshal(toSerialize)
}

type NullableTopologyAdminUserResponse struct {
	value *TopologyAdminUserResponse
	isSet bool
}

func (v NullableTopologyAdminUserResponse) Get() *TopologyAdminUserResponse {
	return v.value
}

func (v *NullableTopologyAdminUserResponse) Set(val *TopologyAdminUserResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTopologyAdminUserResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTopologyAdminUserResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTopologyAdminUserResponse(val *TopologyAdminUserResponse) *NullableTopologyAdminUserResponse {
	return &NullableTopologyAdminUserResponse{value: val, isSet: true}
}

func (v NullableTopologyAdminUserResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTopologyAdminUserResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


