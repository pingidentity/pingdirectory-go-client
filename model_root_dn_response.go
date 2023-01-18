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

// RootDnResponse struct for RootDnResponse
type RootDnResponse struct {
	Meta *MetaMeta `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
	Schemas []EnumrootDnSchemaUrn `json:"schemas,omitempty"`
	DefaultRootPrivilegeName []EnumrootDnDefaultRootPrivilegeNameProp `json:"defaultRootPrivilegeName,omitempty"`
}

// NewRootDnResponse instantiates a new RootDnResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRootDnResponse() *RootDnResponse {
	this := RootDnResponse{}
	return &this
}

// NewRootDnResponseWithDefaults instantiates a new RootDnResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRootDnResponseWithDefaults() *RootDnResponse {
	this := RootDnResponse{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *RootDnResponse) GetMeta() MetaMeta {
	if o == nil || isNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RootDnResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || isNil(o.Meta) {
    return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *RootDnResponse) HasMeta() bool {
	if o != nil && !isNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *RootDnResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *RootDnResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RootDnResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || isNil(o.Urnpingidentityschemasconfigurationmessages20) {
    return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *RootDnResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *RootDnResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetSchemas returns the Schemas field value if set, zero value otherwise.
func (o *RootDnResponse) GetSchemas() []EnumrootDnSchemaUrn {
	if o == nil || isNil(o.Schemas) {
		var ret []EnumrootDnSchemaUrn
		return ret
	}
	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RootDnResponse) GetSchemasOk() ([]EnumrootDnSchemaUrn, bool) {
	if o == nil || isNil(o.Schemas) {
    return nil, false
	}
	return o.Schemas, true
}

// HasSchemas returns a boolean if a field has been set.
func (o *RootDnResponse) HasSchemas() bool {
	if o != nil && !isNil(o.Schemas) {
		return true
	}

	return false
}

// SetSchemas gets a reference to the given []EnumrootDnSchemaUrn and assigns it to the Schemas field.
func (o *RootDnResponse) SetSchemas(v []EnumrootDnSchemaUrn) {
	o.Schemas = v
}

// GetDefaultRootPrivilegeName returns the DefaultRootPrivilegeName field value if set, zero value otherwise.
func (o *RootDnResponse) GetDefaultRootPrivilegeName() []EnumrootDnDefaultRootPrivilegeNameProp {
	if o == nil || isNil(o.DefaultRootPrivilegeName) {
		var ret []EnumrootDnDefaultRootPrivilegeNameProp
		return ret
	}
	return o.DefaultRootPrivilegeName
}

// GetDefaultRootPrivilegeNameOk returns a tuple with the DefaultRootPrivilegeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RootDnResponse) GetDefaultRootPrivilegeNameOk() ([]EnumrootDnDefaultRootPrivilegeNameProp, bool) {
	if o == nil || isNil(o.DefaultRootPrivilegeName) {
    return nil, false
	}
	return o.DefaultRootPrivilegeName, true
}

// HasDefaultRootPrivilegeName returns a boolean if a field has been set.
func (o *RootDnResponse) HasDefaultRootPrivilegeName() bool {
	if o != nil && !isNil(o.DefaultRootPrivilegeName) {
		return true
	}

	return false
}

// SetDefaultRootPrivilegeName gets a reference to the given []EnumrootDnDefaultRootPrivilegeNameProp and assigns it to the DefaultRootPrivilegeName field.
func (o *RootDnResponse) SetDefaultRootPrivilegeName(v []EnumrootDnDefaultRootPrivilegeNameProp) {
	o.DefaultRootPrivilegeName = v
}

func (o RootDnResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		toSerialize["urn:pingidentity:schemas:configuration:messages:2.0"] = o.Urnpingidentityschemasconfigurationmessages20
	}
	if !isNil(o.Schemas) {
		toSerialize["schemas"] = o.Schemas
	}
	if !isNil(o.DefaultRootPrivilegeName) {
		toSerialize["defaultRootPrivilegeName"] = o.DefaultRootPrivilegeName
	}
	return json.Marshal(toSerialize)
}

type NullableRootDnResponse struct {
	value *RootDnResponse
	isSet bool
}

func (v NullableRootDnResponse) Get() *RootDnResponse {
	return v.value
}

func (v *NullableRootDnResponse) Set(val *RootDnResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRootDnResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRootDnResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRootDnResponse(val *RootDnResponse) *NullableRootDnResponse {
	return &NullableRootDnResponse{value: val, isSet: true}
}

func (v NullableRootDnResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRootDnResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

