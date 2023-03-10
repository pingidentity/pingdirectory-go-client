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

// checks if the GenericAttributeSyntaxResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GenericAttributeSyntaxResponse{}

// GenericAttributeSyntaxResponse struct for GenericAttributeSyntaxResponse
type GenericAttributeSyntaxResponse struct {
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
	Schemas                                       []EnumgenericAttributeSyntaxSchemaUrn              `json:"schemas"`
	// Name of the Attribute Syntax
	Id string `json:"id"`
	// Indicates whether the Attribute Syntax is enabled.
	Enabled bool `json:"enabled"`
	// Indicates whether values of this attribute are required to have a \"binary\" transfer option as described in RFC 4522. Attributes with this syntax will generally be referenced with names including \";binary\" (e.g., \"userCertificate;binary\").
	RequireBinaryTransfer *bool `json:"requireBinaryTransfer,omitempty"`
}

// NewGenericAttributeSyntaxResponse instantiates a new GenericAttributeSyntaxResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGenericAttributeSyntaxResponse(schemas []EnumgenericAttributeSyntaxSchemaUrn, id string, enabled bool) *GenericAttributeSyntaxResponse {
	this := GenericAttributeSyntaxResponse{}
	this.Schemas = schemas
	this.Id = id
	this.Enabled = enabled
	return &this
}

// NewGenericAttributeSyntaxResponseWithDefaults instantiates a new GenericAttributeSyntaxResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenericAttributeSyntaxResponseWithDefaults() *GenericAttributeSyntaxResponse {
	this := GenericAttributeSyntaxResponse{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *GenericAttributeSyntaxResponse) GetMeta() MetaMeta {
	if o == nil || IsNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericAttributeSyntaxResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *GenericAttributeSyntaxResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *GenericAttributeSyntaxResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *GenericAttributeSyntaxResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericAttributeSyntaxResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *GenericAttributeSyntaxResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *GenericAttributeSyntaxResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetSchemas returns the Schemas field value
func (o *GenericAttributeSyntaxResponse) GetSchemas() []EnumgenericAttributeSyntaxSchemaUrn {
	if o == nil {
		var ret []EnumgenericAttributeSyntaxSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *GenericAttributeSyntaxResponse) GetSchemasOk() ([]EnumgenericAttributeSyntaxSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *GenericAttributeSyntaxResponse) SetSchemas(v []EnumgenericAttributeSyntaxSchemaUrn) {
	o.Schemas = v
}

// GetId returns the Id field value
func (o *GenericAttributeSyntaxResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GenericAttributeSyntaxResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GenericAttributeSyntaxResponse) SetId(v string) {
	o.Id = v
}

// GetEnabled returns the Enabled field value
func (o *GenericAttributeSyntaxResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *GenericAttributeSyntaxResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *GenericAttributeSyntaxResponse) SetEnabled(v bool) {
	o.Enabled = v
}

// GetRequireBinaryTransfer returns the RequireBinaryTransfer field value if set, zero value otherwise.
func (o *GenericAttributeSyntaxResponse) GetRequireBinaryTransfer() bool {
	if o == nil || IsNil(o.RequireBinaryTransfer) {
		var ret bool
		return ret
	}
	return *o.RequireBinaryTransfer
}

// GetRequireBinaryTransferOk returns a tuple with the RequireBinaryTransfer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericAttributeSyntaxResponse) GetRequireBinaryTransferOk() (*bool, bool) {
	if o == nil || IsNil(o.RequireBinaryTransfer) {
		return nil, false
	}
	return o.RequireBinaryTransfer, true
}

// HasRequireBinaryTransfer returns a boolean if a field has been set.
func (o *GenericAttributeSyntaxResponse) HasRequireBinaryTransfer() bool {
	if o != nil && !IsNil(o.RequireBinaryTransfer) {
		return true
	}

	return false
}

// SetRequireBinaryTransfer gets a reference to the given bool and assigns it to the RequireBinaryTransfer field.
func (o *GenericAttributeSyntaxResponse) SetRequireBinaryTransfer(v bool) {
	o.RequireBinaryTransfer = &v
}

func (o GenericAttributeSyntaxResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GenericAttributeSyntaxResponse) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.RequireBinaryTransfer) {
		toSerialize["requireBinaryTransfer"] = o.RequireBinaryTransfer
	}
	return toSerialize, nil
}

type NullableGenericAttributeSyntaxResponse struct {
	value *GenericAttributeSyntaxResponse
	isSet bool
}

func (v NullableGenericAttributeSyntaxResponse) Get() *GenericAttributeSyntaxResponse {
	return v.value
}

func (v *NullableGenericAttributeSyntaxResponse) Set(val *GenericAttributeSyntaxResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGenericAttributeSyntaxResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGenericAttributeSyntaxResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGenericAttributeSyntaxResponse(val *GenericAttributeSyntaxResponse) *NullableGenericAttributeSyntaxResponse {
	return &NullableGenericAttributeSyntaxResponse{value: val, isSet: true}
}

func (v NullableGenericAttributeSyntaxResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGenericAttributeSyntaxResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
