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

// JsonObjectAttributeSyntaxResponse struct for JsonObjectAttributeSyntaxResponse
type JsonObjectAttributeSyntaxResponse struct {
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
	Schemas                                       []EnumjsonObjectAttributeSyntaxSchemaUrn           `json:"schemas"`
	// Name of the Attribute Syntax
	Id string `json:"id"`
	// Indicates whether the Attribute Syntax is enabled.
	Enabled bool `json:"enabled"`
	// Indicates whether values of this attribute are required to have a \"binary\" transfer option as described in RFC 4522. Attributes with this syntax will generally be referenced with names including \";binary\" (e.g., \"userCertificate;binary\").
	RequireBinaryTransfer *bool `json:"requireBinaryTransfer,omitempty"`
}

// NewJsonObjectAttributeSyntaxResponse instantiates a new JsonObjectAttributeSyntaxResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJsonObjectAttributeSyntaxResponse(schemas []EnumjsonObjectAttributeSyntaxSchemaUrn, id string, enabled bool) *JsonObjectAttributeSyntaxResponse {
	this := JsonObjectAttributeSyntaxResponse{}
	this.Schemas = schemas
	this.Id = id
	this.Enabled = enabled
	return &this
}

// NewJsonObjectAttributeSyntaxResponseWithDefaults instantiates a new JsonObjectAttributeSyntaxResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJsonObjectAttributeSyntaxResponseWithDefaults() *JsonObjectAttributeSyntaxResponse {
	this := JsonObjectAttributeSyntaxResponse{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *JsonObjectAttributeSyntaxResponse) GetMeta() MetaMeta {
	if o == nil || isNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonObjectAttributeSyntaxResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || isNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *JsonObjectAttributeSyntaxResponse) HasMeta() bool {
	if o != nil && !isNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *JsonObjectAttributeSyntaxResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *JsonObjectAttributeSyntaxResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonObjectAttributeSyntaxResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *JsonObjectAttributeSyntaxResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *JsonObjectAttributeSyntaxResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetSchemas returns the Schemas field value
func (o *JsonObjectAttributeSyntaxResponse) GetSchemas() []EnumjsonObjectAttributeSyntaxSchemaUrn {
	if o == nil {
		var ret []EnumjsonObjectAttributeSyntaxSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *JsonObjectAttributeSyntaxResponse) GetSchemasOk() ([]EnumjsonObjectAttributeSyntaxSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *JsonObjectAttributeSyntaxResponse) SetSchemas(v []EnumjsonObjectAttributeSyntaxSchemaUrn) {
	o.Schemas = v
}

// GetId returns the Id field value
func (o *JsonObjectAttributeSyntaxResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *JsonObjectAttributeSyntaxResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *JsonObjectAttributeSyntaxResponse) SetId(v string) {
	o.Id = v
}

// GetEnabled returns the Enabled field value
func (o *JsonObjectAttributeSyntaxResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *JsonObjectAttributeSyntaxResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *JsonObjectAttributeSyntaxResponse) SetEnabled(v bool) {
	o.Enabled = v
}

// GetRequireBinaryTransfer returns the RequireBinaryTransfer field value if set, zero value otherwise.
func (o *JsonObjectAttributeSyntaxResponse) GetRequireBinaryTransfer() bool {
	if o == nil || isNil(o.RequireBinaryTransfer) {
		var ret bool
		return ret
	}
	return *o.RequireBinaryTransfer
}

// GetRequireBinaryTransferOk returns a tuple with the RequireBinaryTransfer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonObjectAttributeSyntaxResponse) GetRequireBinaryTransferOk() (*bool, bool) {
	if o == nil || isNil(o.RequireBinaryTransfer) {
		return nil, false
	}
	return o.RequireBinaryTransfer, true
}

// HasRequireBinaryTransfer returns a boolean if a field has been set.
func (o *JsonObjectAttributeSyntaxResponse) HasRequireBinaryTransfer() bool {
	if o != nil && !isNil(o.RequireBinaryTransfer) {
		return true
	}

	return false
}

// SetRequireBinaryTransfer gets a reference to the given bool and assigns it to the RequireBinaryTransfer field.
func (o *JsonObjectAttributeSyntaxResponse) SetRequireBinaryTransfer(v bool) {
	o.RequireBinaryTransfer = &v
}

func (o JsonObjectAttributeSyntaxResponse) MarshalJSON() ([]byte, error) {
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
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	if !isNil(o.RequireBinaryTransfer) {
		toSerialize["requireBinaryTransfer"] = o.RequireBinaryTransfer
	}
	return json.Marshal(toSerialize)
}

type NullableJsonObjectAttributeSyntaxResponse struct {
	value *JsonObjectAttributeSyntaxResponse
	isSet bool
}

func (v NullableJsonObjectAttributeSyntaxResponse) Get() *JsonObjectAttributeSyntaxResponse {
	return v.value
}

func (v *NullableJsonObjectAttributeSyntaxResponse) Set(val *JsonObjectAttributeSyntaxResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableJsonObjectAttributeSyntaxResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableJsonObjectAttributeSyntaxResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJsonObjectAttributeSyntaxResponse(val *JsonObjectAttributeSyntaxResponse) *NullableJsonObjectAttributeSyntaxResponse {
	return &NullableJsonObjectAttributeSyntaxResponse{value: val, isSet: true}
}

func (v NullableJsonObjectAttributeSyntaxResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJsonObjectAttributeSyntaxResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}