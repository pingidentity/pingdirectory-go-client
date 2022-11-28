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

// JsonAttributeConstraintsResponse struct for JsonAttributeConstraintsResponse
type JsonAttributeConstraintsResponse struct {
	Meta *MetaMeta `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
	// Name of the JSON Attribute Constraints
	Id string `json:"id"`
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

// NewJsonAttributeConstraintsResponse instantiates a new JsonAttributeConstraintsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJsonAttributeConstraintsResponse(id string, attributeType string) *JsonAttributeConstraintsResponse {
	this := JsonAttributeConstraintsResponse{}
	this.Id = id
	this.AttributeType = attributeType
	return &this
}

// NewJsonAttributeConstraintsResponseWithDefaults instantiates a new JsonAttributeConstraintsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJsonAttributeConstraintsResponseWithDefaults() *JsonAttributeConstraintsResponse {
	this := JsonAttributeConstraintsResponse{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *JsonAttributeConstraintsResponse) GetMeta() MetaMeta {
	if o == nil || isNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonAttributeConstraintsResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || isNil(o.Meta) {
    return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *JsonAttributeConstraintsResponse) HasMeta() bool {
	if o != nil && !isNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *JsonAttributeConstraintsResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *JsonAttributeConstraintsResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonAttributeConstraintsResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || isNil(o.Urnpingidentityschemasconfigurationmessages20) {
    return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *JsonAttributeConstraintsResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *JsonAttributeConstraintsResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetId returns the Id field value
func (o *JsonAttributeConstraintsResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *JsonAttributeConstraintsResponse) GetIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *JsonAttributeConstraintsResponse) SetId(v string) {
	o.Id = v
}

// GetSchemas returns the Schemas field value if set, zero value otherwise.
func (o *JsonAttributeConstraintsResponse) GetSchemas() []EnumjsonAttributeConstraintsSchemaUrn {
	if o == nil || isNil(o.Schemas) {
		var ret []EnumjsonAttributeConstraintsSchemaUrn
		return ret
	}
	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonAttributeConstraintsResponse) GetSchemasOk() ([]EnumjsonAttributeConstraintsSchemaUrn, bool) {
	if o == nil || isNil(o.Schemas) {
    return nil, false
	}
	return o.Schemas, true
}

// HasSchemas returns a boolean if a field has been set.
func (o *JsonAttributeConstraintsResponse) HasSchemas() bool {
	if o != nil && !isNil(o.Schemas) {
		return true
	}

	return false
}

// SetSchemas gets a reference to the given []EnumjsonAttributeConstraintsSchemaUrn and assigns it to the Schemas field.
func (o *JsonAttributeConstraintsResponse) SetSchemas(v []EnumjsonAttributeConstraintsSchemaUrn) {
	o.Schemas = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *JsonAttributeConstraintsResponse) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonAttributeConstraintsResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *JsonAttributeConstraintsResponse) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *JsonAttributeConstraintsResponse) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *JsonAttributeConstraintsResponse) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonAttributeConstraintsResponse) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *JsonAttributeConstraintsResponse) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *JsonAttributeConstraintsResponse) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetAttributeType returns the AttributeType field value
func (o *JsonAttributeConstraintsResponse) GetAttributeType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AttributeType
}

// GetAttributeTypeOk returns a tuple with the AttributeType field value
// and a boolean to check if the value has been set.
func (o *JsonAttributeConstraintsResponse) GetAttributeTypeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.AttributeType, true
}

// SetAttributeType sets field value
func (o *JsonAttributeConstraintsResponse) SetAttributeType(v string) {
	o.AttributeType = v
}

// GetAllowUnnamedFields returns the AllowUnnamedFields field value if set, zero value otherwise.
func (o *JsonAttributeConstraintsResponse) GetAllowUnnamedFields() bool {
	if o == nil || isNil(o.AllowUnnamedFields) {
		var ret bool
		return ret
	}
	return *o.AllowUnnamedFields
}

// GetAllowUnnamedFieldsOk returns a tuple with the AllowUnnamedFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonAttributeConstraintsResponse) GetAllowUnnamedFieldsOk() (*bool, bool) {
	if o == nil || isNil(o.AllowUnnamedFields) {
    return nil, false
	}
	return o.AllowUnnamedFields, true
}

// HasAllowUnnamedFields returns a boolean if a field has been set.
func (o *JsonAttributeConstraintsResponse) HasAllowUnnamedFields() bool {
	if o != nil && !isNil(o.AllowUnnamedFields) {
		return true
	}

	return false
}

// SetAllowUnnamedFields gets a reference to the given bool and assigns it to the AllowUnnamedFields field.
func (o *JsonAttributeConstraintsResponse) SetAllowUnnamedFields(v bool) {
	o.AllowUnnamedFields = &v
}

func (o JsonAttributeConstraintsResponse) MarshalJSON() ([]byte, error) {
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

type NullableJsonAttributeConstraintsResponse struct {
	value *JsonAttributeConstraintsResponse
	isSet bool
}

func (v NullableJsonAttributeConstraintsResponse) Get() *JsonAttributeConstraintsResponse {
	return v.value
}

func (v *NullableJsonAttributeConstraintsResponse) Set(val *JsonAttributeConstraintsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableJsonAttributeConstraintsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableJsonAttributeConstraintsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJsonAttributeConstraintsResponse(val *JsonAttributeConstraintsResponse) *NullableJsonAttributeConstraintsResponse {
	return &NullableJsonAttributeConstraintsResponse{value: val, isSet: true}
}

func (v NullableJsonAttributeConstraintsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJsonAttributeConstraintsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


