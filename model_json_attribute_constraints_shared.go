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

// JsonAttributeConstraintsShared struct for JsonAttributeConstraintsShared
type JsonAttributeConstraintsShared struct {
	Schemas []EnumjsonAttributeConstraintsSchemaUrn `json:"schemas,omitempty"`
	// A description for this JSON Attribute Constraints
	Description *string `json:"description,omitempty"`
	// Indicates whether this JSON Attribute Constraints is enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// The name or OID of the LDAP attribute type whose values will be subject to the associated field constraints. This attribute type must be defined in the server schema, and it must have a \"JSON object\" syntax.
	AttributeType string `json:"attributeType"`
	// Indicates whether JSON objects stored as values of attributes with the associated attribute-type will be permitted to include fields for which there is no subordinate json-field-constraints definition. If unnamed fields are allowed, then no constraints will be imposed on the values of those fields. However, if unnamed fields are not allowed, then the server will reject any attempt to store a JSON object with a field for which there is no corresponding json-fields-constraints definition.
	AllowUnnamedFields *bool `json:"allowUnnamedFields,omitempty"`
}

// NewJsonAttributeConstraintsShared instantiates a new JsonAttributeConstraintsShared object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJsonAttributeConstraintsShared(attributeType string) *JsonAttributeConstraintsShared {
	this := JsonAttributeConstraintsShared{}
	this.AttributeType = attributeType
	return &this
}

// NewJsonAttributeConstraintsSharedWithDefaults instantiates a new JsonAttributeConstraintsShared object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJsonAttributeConstraintsSharedWithDefaults() *JsonAttributeConstraintsShared {
	this := JsonAttributeConstraintsShared{}
	return &this
}

// GetSchemas returns the Schemas field value if set, zero value otherwise.
func (o *JsonAttributeConstraintsShared) GetSchemas() []EnumjsonAttributeConstraintsSchemaUrn {
	if o == nil || isNil(o.Schemas) {
		var ret []EnumjsonAttributeConstraintsSchemaUrn
		return ret
	}
	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonAttributeConstraintsShared) GetSchemasOk() ([]EnumjsonAttributeConstraintsSchemaUrn, bool) {
	if o == nil || isNil(o.Schemas) {
    return nil, false
	}
	return o.Schemas, true
}

// HasSchemas returns a boolean if a field has been set.
func (o *JsonAttributeConstraintsShared) HasSchemas() bool {
	if o != nil && !isNil(o.Schemas) {
		return true
	}

	return false
}

// SetSchemas gets a reference to the given []EnumjsonAttributeConstraintsSchemaUrn and assigns it to the Schemas field.
func (o *JsonAttributeConstraintsShared) SetSchemas(v []EnumjsonAttributeConstraintsSchemaUrn) {
	o.Schemas = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *JsonAttributeConstraintsShared) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonAttributeConstraintsShared) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *JsonAttributeConstraintsShared) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *JsonAttributeConstraintsShared) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *JsonAttributeConstraintsShared) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonAttributeConstraintsShared) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *JsonAttributeConstraintsShared) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *JsonAttributeConstraintsShared) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetAttributeType returns the AttributeType field value
func (o *JsonAttributeConstraintsShared) GetAttributeType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AttributeType
}

// GetAttributeTypeOk returns a tuple with the AttributeType field value
// and a boolean to check if the value has been set.
func (o *JsonAttributeConstraintsShared) GetAttributeTypeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.AttributeType, true
}

// SetAttributeType sets field value
func (o *JsonAttributeConstraintsShared) SetAttributeType(v string) {
	o.AttributeType = v
}

// GetAllowUnnamedFields returns the AllowUnnamedFields field value if set, zero value otherwise.
func (o *JsonAttributeConstraintsShared) GetAllowUnnamedFields() bool {
	if o == nil || isNil(o.AllowUnnamedFields) {
		var ret bool
		return ret
	}
	return *o.AllowUnnamedFields
}

// GetAllowUnnamedFieldsOk returns a tuple with the AllowUnnamedFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonAttributeConstraintsShared) GetAllowUnnamedFieldsOk() (*bool, bool) {
	if o == nil || isNil(o.AllowUnnamedFields) {
    return nil, false
	}
	return o.AllowUnnamedFields, true
}

// HasAllowUnnamedFields returns a boolean if a field has been set.
func (o *JsonAttributeConstraintsShared) HasAllowUnnamedFields() bool {
	if o != nil && !isNil(o.AllowUnnamedFields) {
		return true
	}

	return false
}

// SetAllowUnnamedFields gets a reference to the given bool and assigns it to the AllowUnnamedFields field.
func (o *JsonAttributeConstraintsShared) SetAllowUnnamedFields(v bool) {
	o.AllowUnnamedFields = &v
}

func (o JsonAttributeConstraintsShared) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Schemas) {
		toSerialize["schemas"] = o.Schemas
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if true {
		toSerialize["attributeType"] = o.AttributeType
	}
	if !isNil(o.AllowUnnamedFields) {
		toSerialize["allowUnnamedFields"] = o.AllowUnnamedFields
	}
	return json.Marshal(toSerialize)
}

type NullableJsonAttributeConstraintsShared struct {
	value *JsonAttributeConstraintsShared
	isSet bool
}

func (v NullableJsonAttributeConstraintsShared) Get() *JsonAttributeConstraintsShared {
	return v.value
}

func (v *NullableJsonAttributeConstraintsShared) Set(val *JsonAttributeConstraintsShared) {
	v.value = val
	v.isSet = true
}

func (v NullableJsonAttributeConstraintsShared) IsSet() bool {
	return v.isSet
}

func (v *NullableJsonAttributeConstraintsShared) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJsonAttributeConstraintsShared(val *JsonAttributeConstraintsShared) *NullableJsonAttributeConstraintsShared {
	return &NullableJsonAttributeConstraintsShared{value: val, isSet: true}
}

func (v NullableJsonAttributeConstraintsShared) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJsonAttributeConstraintsShared) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


