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

// IntegerAttributeSyntaxResponse struct for IntegerAttributeSyntaxResponse
type IntegerAttributeSyntaxResponse struct {
	Meta *MetaMeta `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
	Schemas []EnumintegerAttributeSyntaxSchemaUrn `json:"schemas"`
	// Name of the Attribute Syntax
	Id *string `json:"id,omitempty"`
	// Indicates whether values of attributes with this syntax should be compacted when stored in a local DB database.
	EnableCompaction *bool `json:"enableCompaction,omitempty"`
	// Specifies the specific attributes (which should be associated with this syntax) whose values should be compacted. If one or more include attributes are specified, then only those attributes will have their values compacted. If not set then all attributes will have their values compacted. The exclude-attribute-from-compaction property takes precedence over this property.
	IncludeAttributeInCompaction []string `json:"includeAttributeInCompaction,omitempty"`
	// Specifies the specific attributes (which should be associated with this syntax) whose values should not be compacted. If one or more exclude attributes are specified, then values of those attributes will not have their values compacted. This property takes precedence over the include-attribute-in-compaction property.
	ExcludeAttributeFromCompaction []string `json:"excludeAttributeFromCompaction,omitempty"`
	// Indicates whether the Attribute Syntax is enabled.
	Enabled bool `json:"enabled"`
	// Indicates whether values of this attribute are required to have a \"binary\" transfer option as described in RFC 4522. Attributes with this syntax will generally be referenced with names including \";binary\" (e.g., \"userCertificate;binary\").
	RequireBinaryTransfer *bool `json:"requireBinaryTransfer,omitempty"`
}

// NewIntegerAttributeSyntaxResponse instantiates a new IntegerAttributeSyntaxResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIntegerAttributeSyntaxResponse(schemas []EnumintegerAttributeSyntaxSchemaUrn, enabled bool) *IntegerAttributeSyntaxResponse {
	this := IntegerAttributeSyntaxResponse{}
	this.Schemas = schemas
	this.Enabled = enabled
	return &this
}

// NewIntegerAttributeSyntaxResponseWithDefaults instantiates a new IntegerAttributeSyntaxResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIntegerAttributeSyntaxResponseWithDefaults() *IntegerAttributeSyntaxResponse {
	this := IntegerAttributeSyntaxResponse{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *IntegerAttributeSyntaxResponse) GetMeta() MetaMeta {
	if o == nil || isNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegerAttributeSyntaxResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || isNil(o.Meta) {
    return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *IntegerAttributeSyntaxResponse) HasMeta() bool {
	if o != nil && !isNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *IntegerAttributeSyntaxResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *IntegerAttributeSyntaxResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegerAttributeSyntaxResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || isNil(o.Urnpingidentityschemasconfigurationmessages20) {
    return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *IntegerAttributeSyntaxResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *IntegerAttributeSyntaxResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetSchemas returns the Schemas field value
func (o *IntegerAttributeSyntaxResponse) GetSchemas() []EnumintegerAttributeSyntaxSchemaUrn {
	if o == nil {
		var ret []EnumintegerAttributeSyntaxSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *IntegerAttributeSyntaxResponse) GetSchemasOk() ([]EnumintegerAttributeSyntaxSchemaUrn, bool) {
	if o == nil {
    return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *IntegerAttributeSyntaxResponse) SetSchemas(v []EnumintegerAttributeSyntaxSchemaUrn) {
	o.Schemas = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *IntegerAttributeSyntaxResponse) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegerAttributeSyntaxResponse) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *IntegerAttributeSyntaxResponse) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *IntegerAttributeSyntaxResponse) SetId(v string) {
	o.Id = &v
}

// GetEnableCompaction returns the EnableCompaction field value if set, zero value otherwise.
func (o *IntegerAttributeSyntaxResponse) GetEnableCompaction() bool {
	if o == nil || isNil(o.EnableCompaction) {
		var ret bool
		return ret
	}
	return *o.EnableCompaction
}

// GetEnableCompactionOk returns a tuple with the EnableCompaction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegerAttributeSyntaxResponse) GetEnableCompactionOk() (*bool, bool) {
	if o == nil || isNil(o.EnableCompaction) {
    return nil, false
	}
	return o.EnableCompaction, true
}

// HasEnableCompaction returns a boolean if a field has been set.
func (o *IntegerAttributeSyntaxResponse) HasEnableCompaction() bool {
	if o != nil && !isNil(o.EnableCompaction) {
		return true
	}

	return false
}

// SetEnableCompaction gets a reference to the given bool and assigns it to the EnableCompaction field.
func (o *IntegerAttributeSyntaxResponse) SetEnableCompaction(v bool) {
	o.EnableCompaction = &v
}

// GetIncludeAttributeInCompaction returns the IncludeAttributeInCompaction field value if set, zero value otherwise.
func (o *IntegerAttributeSyntaxResponse) GetIncludeAttributeInCompaction() []string {
	if o == nil || isNil(o.IncludeAttributeInCompaction) {
		var ret []string
		return ret
	}
	return o.IncludeAttributeInCompaction
}

// GetIncludeAttributeInCompactionOk returns a tuple with the IncludeAttributeInCompaction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegerAttributeSyntaxResponse) GetIncludeAttributeInCompactionOk() ([]string, bool) {
	if o == nil || isNil(o.IncludeAttributeInCompaction) {
    return nil, false
	}
	return o.IncludeAttributeInCompaction, true
}

// HasIncludeAttributeInCompaction returns a boolean if a field has been set.
func (o *IntegerAttributeSyntaxResponse) HasIncludeAttributeInCompaction() bool {
	if o != nil && !isNil(o.IncludeAttributeInCompaction) {
		return true
	}

	return false
}

// SetIncludeAttributeInCompaction gets a reference to the given []string and assigns it to the IncludeAttributeInCompaction field.
func (o *IntegerAttributeSyntaxResponse) SetIncludeAttributeInCompaction(v []string) {
	o.IncludeAttributeInCompaction = v
}

// GetExcludeAttributeFromCompaction returns the ExcludeAttributeFromCompaction field value if set, zero value otherwise.
func (o *IntegerAttributeSyntaxResponse) GetExcludeAttributeFromCompaction() []string {
	if o == nil || isNil(o.ExcludeAttributeFromCompaction) {
		var ret []string
		return ret
	}
	return o.ExcludeAttributeFromCompaction
}

// GetExcludeAttributeFromCompactionOk returns a tuple with the ExcludeAttributeFromCompaction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegerAttributeSyntaxResponse) GetExcludeAttributeFromCompactionOk() ([]string, bool) {
	if o == nil || isNil(o.ExcludeAttributeFromCompaction) {
    return nil, false
	}
	return o.ExcludeAttributeFromCompaction, true
}

// HasExcludeAttributeFromCompaction returns a boolean if a field has been set.
func (o *IntegerAttributeSyntaxResponse) HasExcludeAttributeFromCompaction() bool {
	if o != nil && !isNil(o.ExcludeAttributeFromCompaction) {
		return true
	}

	return false
}

// SetExcludeAttributeFromCompaction gets a reference to the given []string and assigns it to the ExcludeAttributeFromCompaction field.
func (o *IntegerAttributeSyntaxResponse) SetExcludeAttributeFromCompaction(v []string) {
	o.ExcludeAttributeFromCompaction = v
}

// GetEnabled returns the Enabled field value
func (o *IntegerAttributeSyntaxResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *IntegerAttributeSyntaxResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *IntegerAttributeSyntaxResponse) SetEnabled(v bool) {
	o.Enabled = v
}

// GetRequireBinaryTransfer returns the RequireBinaryTransfer field value if set, zero value otherwise.
func (o *IntegerAttributeSyntaxResponse) GetRequireBinaryTransfer() bool {
	if o == nil || isNil(o.RequireBinaryTransfer) {
		var ret bool
		return ret
	}
	return *o.RequireBinaryTransfer
}

// GetRequireBinaryTransferOk returns a tuple with the RequireBinaryTransfer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegerAttributeSyntaxResponse) GetRequireBinaryTransferOk() (*bool, bool) {
	if o == nil || isNil(o.RequireBinaryTransfer) {
    return nil, false
	}
	return o.RequireBinaryTransfer, true
}

// HasRequireBinaryTransfer returns a boolean if a field has been set.
func (o *IntegerAttributeSyntaxResponse) HasRequireBinaryTransfer() bool {
	if o != nil && !isNil(o.RequireBinaryTransfer) {
		return true
	}

	return false
}

// SetRequireBinaryTransfer gets a reference to the given bool and assigns it to the RequireBinaryTransfer field.
func (o *IntegerAttributeSyntaxResponse) SetRequireBinaryTransfer(v bool) {
	o.RequireBinaryTransfer = &v
}

func (o IntegerAttributeSyntaxResponse) MarshalJSON() ([]byte, error) {
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
	if !isNil(o.EnableCompaction) {
		toSerialize["enableCompaction"] = o.EnableCompaction
	}
	if !isNil(o.IncludeAttributeInCompaction) {
		toSerialize["includeAttributeInCompaction"] = o.IncludeAttributeInCompaction
	}
	if !isNil(o.ExcludeAttributeFromCompaction) {
		toSerialize["excludeAttributeFromCompaction"] = o.ExcludeAttributeFromCompaction
	}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	if !isNil(o.RequireBinaryTransfer) {
		toSerialize["requireBinaryTransfer"] = o.RequireBinaryTransfer
	}
	return json.Marshal(toSerialize)
}

type NullableIntegerAttributeSyntaxResponse struct {
	value *IntegerAttributeSyntaxResponse
	isSet bool
}

func (v NullableIntegerAttributeSyntaxResponse) Get() *IntegerAttributeSyntaxResponse {
	return v.value
}

func (v *NullableIntegerAttributeSyntaxResponse) Set(val *IntegerAttributeSyntaxResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableIntegerAttributeSyntaxResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableIntegerAttributeSyntaxResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIntegerAttributeSyntaxResponse(val *IntegerAttributeSyntaxResponse) *NullableIntegerAttributeSyntaxResponse {
	return &NullableIntegerAttributeSyntaxResponse{value: val, isSet: true}
}

func (v NullableIntegerAttributeSyntaxResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIntegerAttributeSyntaxResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


