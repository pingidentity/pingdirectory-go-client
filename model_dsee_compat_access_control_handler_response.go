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

// DseeCompatAccessControlHandlerResponse struct for DseeCompatAccessControlHandlerResponse
type DseeCompatAccessControlHandlerResponse struct {
	Meta *MetaMeta `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
	Schemas []EnumdseeCompatAccessControlHandlerSchemaUrn `json:"schemas,omitempty"`
	// Defines global access control rules.
	GlobalACI []string `json:"globalACI,omitempty"`
	AllowedBindControl []EnumaccessControlHandlerAllowedBindControlProp `json:"allowedBindControl,omitempty"`
	// Specifies the OIDs of any additional controls (not covered by the allowed-bind-control property) that should be permitted in bind requests.
	AllowedBindControlOID []string `json:"allowedBindControlOID,omitempty"`
	// Indicates whether this Access Control Handler is enabled. If set to FALSE, then no access control is enforced, and any client (including unauthenticated or anonymous clients) could be allowed to perform any operation if not subject to other restrictions, such as those enforced by the privilege subsystem.
	Enabled bool `json:"enabled"`
}

// NewDseeCompatAccessControlHandlerResponse instantiates a new DseeCompatAccessControlHandlerResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDseeCompatAccessControlHandlerResponse(enabled bool) *DseeCompatAccessControlHandlerResponse {
	this := DseeCompatAccessControlHandlerResponse{}
	this.Enabled = enabled
	return &this
}

// NewDseeCompatAccessControlHandlerResponseWithDefaults instantiates a new DseeCompatAccessControlHandlerResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDseeCompatAccessControlHandlerResponseWithDefaults() *DseeCompatAccessControlHandlerResponse {
	this := DseeCompatAccessControlHandlerResponse{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *DseeCompatAccessControlHandlerResponse) GetMeta() MetaMeta {
	if o == nil || isNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DseeCompatAccessControlHandlerResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || isNil(o.Meta) {
    return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *DseeCompatAccessControlHandlerResponse) HasMeta() bool {
	if o != nil && !isNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *DseeCompatAccessControlHandlerResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *DseeCompatAccessControlHandlerResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DseeCompatAccessControlHandlerResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || isNil(o.Urnpingidentityschemasconfigurationmessages20) {
    return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *DseeCompatAccessControlHandlerResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *DseeCompatAccessControlHandlerResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetSchemas returns the Schemas field value if set, zero value otherwise.
func (o *DseeCompatAccessControlHandlerResponse) GetSchemas() []EnumdseeCompatAccessControlHandlerSchemaUrn {
	if o == nil || isNil(o.Schemas) {
		var ret []EnumdseeCompatAccessControlHandlerSchemaUrn
		return ret
	}
	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DseeCompatAccessControlHandlerResponse) GetSchemasOk() ([]EnumdseeCompatAccessControlHandlerSchemaUrn, bool) {
	if o == nil || isNil(o.Schemas) {
    return nil, false
	}
	return o.Schemas, true
}

// HasSchemas returns a boolean if a field has been set.
func (o *DseeCompatAccessControlHandlerResponse) HasSchemas() bool {
	if o != nil && !isNil(o.Schemas) {
		return true
	}

	return false
}

// SetSchemas gets a reference to the given []EnumdseeCompatAccessControlHandlerSchemaUrn and assigns it to the Schemas field.
func (o *DseeCompatAccessControlHandlerResponse) SetSchemas(v []EnumdseeCompatAccessControlHandlerSchemaUrn) {
	o.Schemas = v
}

// GetGlobalACI returns the GlobalACI field value if set, zero value otherwise.
func (o *DseeCompatAccessControlHandlerResponse) GetGlobalACI() []string {
	if o == nil || isNil(o.GlobalACI) {
		var ret []string
		return ret
	}
	return o.GlobalACI
}

// GetGlobalACIOk returns a tuple with the GlobalACI field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DseeCompatAccessControlHandlerResponse) GetGlobalACIOk() ([]string, bool) {
	if o == nil || isNil(o.GlobalACI) {
    return nil, false
	}
	return o.GlobalACI, true
}

// HasGlobalACI returns a boolean if a field has been set.
func (o *DseeCompatAccessControlHandlerResponse) HasGlobalACI() bool {
	if o != nil && !isNil(o.GlobalACI) {
		return true
	}

	return false
}

// SetGlobalACI gets a reference to the given []string and assigns it to the GlobalACI field.
func (o *DseeCompatAccessControlHandlerResponse) SetGlobalACI(v []string) {
	o.GlobalACI = v
}

// GetAllowedBindControl returns the AllowedBindControl field value if set, zero value otherwise.
func (o *DseeCompatAccessControlHandlerResponse) GetAllowedBindControl() []EnumaccessControlHandlerAllowedBindControlProp {
	if o == nil || isNil(o.AllowedBindControl) {
		var ret []EnumaccessControlHandlerAllowedBindControlProp
		return ret
	}
	return o.AllowedBindControl
}

// GetAllowedBindControlOk returns a tuple with the AllowedBindControl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DseeCompatAccessControlHandlerResponse) GetAllowedBindControlOk() ([]EnumaccessControlHandlerAllowedBindControlProp, bool) {
	if o == nil || isNil(o.AllowedBindControl) {
    return nil, false
	}
	return o.AllowedBindControl, true
}

// HasAllowedBindControl returns a boolean if a field has been set.
func (o *DseeCompatAccessControlHandlerResponse) HasAllowedBindControl() bool {
	if o != nil && !isNil(o.AllowedBindControl) {
		return true
	}

	return false
}

// SetAllowedBindControl gets a reference to the given []EnumaccessControlHandlerAllowedBindControlProp and assigns it to the AllowedBindControl field.
func (o *DseeCompatAccessControlHandlerResponse) SetAllowedBindControl(v []EnumaccessControlHandlerAllowedBindControlProp) {
	o.AllowedBindControl = v
}

// GetAllowedBindControlOID returns the AllowedBindControlOID field value if set, zero value otherwise.
func (o *DseeCompatAccessControlHandlerResponse) GetAllowedBindControlOID() []string {
	if o == nil || isNil(o.AllowedBindControlOID) {
		var ret []string
		return ret
	}
	return o.AllowedBindControlOID
}

// GetAllowedBindControlOIDOk returns a tuple with the AllowedBindControlOID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DseeCompatAccessControlHandlerResponse) GetAllowedBindControlOIDOk() ([]string, bool) {
	if o == nil || isNil(o.AllowedBindControlOID) {
    return nil, false
	}
	return o.AllowedBindControlOID, true
}

// HasAllowedBindControlOID returns a boolean if a field has been set.
func (o *DseeCompatAccessControlHandlerResponse) HasAllowedBindControlOID() bool {
	if o != nil && !isNil(o.AllowedBindControlOID) {
		return true
	}

	return false
}

// SetAllowedBindControlOID gets a reference to the given []string and assigns it to the AllowedBindControlOID field.
func (o *DseeCompatAccessControlHandlerResponse) SetAllowedBindControlOID(v []string) {
	o.AllowedBindControlOID = v
}

// GetEnabled returns the Enabled field value
func (o *DseeCompatAccessControlHandlerResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *DseeCompatAccessControlHandlerResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *DseeCompatAccessControlHandlerResponse) SetEnabled(v bool) {
	o.Enabled = v
}

func (o DseeCompatAccessControlHandlerResponse) MarshalJSON() ([]byte, error) {
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
	if !isNil(o.GlobalACI) {
		toSerialize["globalACI"] = o.GlobalACI
	}
	if !isNil(o.AllowedBindControl) {
		toSerialize["allowedBindControl"] = o.AllowedBindControl
	}
	if !isNil(o.AllowedBindControlOID) {
		toSerialize["allowedBindControlOID"] = o.AllowedBindControlOID
	}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableDseeCompatAccessControlHandlerResponse struct {
	value *DseeCompatAccessControlHandlerResponse
	isSet bool
}

func (v NullableDseeCompatAccessControlHandlerResponse) Get() *DseeCompatAccessControlHandlerResponse {
	return v.value
}

func (v *NullableDseeCompatAccessControlHandlerResponse) Set(val *DseeCompatAccessControlHandlerResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDseeCompatAccessControlHandlerResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDseeCompatAccessControlHandlerResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDseeCompatAccessControlHandlerResponse(val *DseeCompatAccessControlHandlerResponse) *NullableDseeCompatAccessControlHandlerResponse {
	return &NullableDseeCompatAccessControlHandlerResponse{value: val, isSet: true}
}

func (v NullableDseeCompatAccessControlHandlerResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDseeCompatAccessControlHandlerResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

