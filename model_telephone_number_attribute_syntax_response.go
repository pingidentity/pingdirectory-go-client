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

// TelephoneNumberAttributeSyntaxResponse struct for TelephoneNumberAttributeSyntaxResponse
type TelephoneNumberAttributeSyntaxResponse struct {
	Meta *MetaMeta `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
	Schemas []EnumtelephoneNumberAttributeSyntaxSchemaUrn `json:"schemas"`
	// Name of the Attribute Syntax
	Id string `json:"id"`
	// Indicates whether to require telephone number values to strictly comply with the standard definition for this syntax.
	StrictFormat *bool `json:"strictFormat,omitempty"`
	// Indicates whether the Attribute Syntax is enabled.
	Enabled bool `json:"enabled"`
	// Indicates whether values of this attribute are required to have a \"binary\" transfer option as described in RFC 4522. Attributes with this syntax will generally be referenced with names including \";binary\" (e.g., \"userCertificate;binary\").
	RequireBinaryTransfer *bool `json:"requireBinaryTransfer,omitempty"`
}

// NewTelephoneNumberAttributeSyntaxResponse instantiates a new TelephoneNumberAttributeSyntaxResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTelephoneNumberAttributeSyntaxResponse(schemas []EnumtelephoneNumberAttributeSyntaxSchemaUrn, id string, enabled bool) *TelephoneNumberAttributeSyntaxResponse {
	this := TelephoneNumberAttributeSyntaxResponse{}
	this.Schemas = schemas
	this.Id = id
	this.Enabled = enabled
	return &this
}

// NewTelephoneNumberAttributeSyntaxResponseWithDefaults instantiates a new TelephoneNumberAttributeSyntaxResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTelephoneNumberAttributeSyntaxResponseWithDefaults() *TelephoneNumberAttributeSyntaxResponse {
	this := TelephoneNumberAttributeSyntaxResponse{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *TelephoneNumberAttributeSyntaxResponse) GetMeta() MetaMeta {
	if o == nil || isNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelephoneNumberAttributeSyntaxResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || isNil(o.Meta) {
    return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *TelephoneNumberAttributeSyntaxResponse) HasMeta() bool {
	if o != nil && !isNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *TelephoneNumberAttributeSyntaxResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *TelephoneNumberAttributeSyntaxResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelephoneNumberAttributeSyntaxResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || isNil(o.Urnpingidentityschemasconfigurationmessages20) {
    return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *TelephoneNumberAttributeSyntaxResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *TelephoneNumberAttributeSyntaxResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetSchemas returns the Schemas field value
func (o *TelephoneNumberAttributeSyntaxResponse) GetSchemas() []EnumtelephoneNumberAttributeSyntaxSchemaUrn {
	if o == nil {
		var ret []EnumtelephoneNumberAttributeSyntaxSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *TelephoneNumberAttributeSyntaxResponse) GetSchemasOk() ([]EnumtelephoneNumberAttributeSyntaxSchemaUrn, bool) {
	if o == nil {
    return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *TelephoneNumberAttributeSyntaxResponse) SetSchemas(v []EnumtelephoneNumberAttributeSyntaxSchemaUrn) {
	o.Schemas = v
}

// GetId returns the Id field value
func (o *TelephoneNumberAttributeSyntaxResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TelephoneNumberAttributeSyntaxResponse) GetIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TelephoneNumberAttributeSyntaxResponse) SetId(v string) {
	o.Id = v
}

// GetStrictFormat returns the StrictFormat field value if set, zero value otherwise.
func (o *TelephoneNumberAttributeSyntaxResponse) GetStrictFormat() bool {
	if o == nil || isNil(o.StrictFormat) {
		var ret bool
		return ret
	}
	return *o.StrictFormat
}

// GetStrictFormatOk returns a tuple with the StrictFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelephoneNumberAttributeSyntaxResponse) GetStrictFormatOk() (*bool, bool) {
	if o == nil || isNil(o.StrictFormat) {
    return nil, false
	}
	return o.StrictFormat, true
}

// HasStrictFormat returns a boolean if a field has been set.
func (o *TelephoneNumberAttributeSyntaxResponse) HasStrictFormat() bool {
	if o != nil && !isNil(o.StrictFormat) {
		return true
	}

	return false
}

// SetStrictFormat gets a reference to the given bool and assigns it to the StrictFormat field.
func (o *TelephoneNumberAttributeSyntaxResponse) SetStrictFormat(v bool) {
	o.StrictFormat = &v
}

// GetEnabled returns the Enabled field value
func (o *TelephoneNumberAttributeSyntaxResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *TelephoneNumberAttributeSyntaxResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *TelephoneNumberAttributeSyntaxResponse) SetEnabled(v bool) {
	o.Enabled = v
}

// GetRequireBinaryTransfer returns the RequireBinaryTransfer field value if set, zero value otherwise.
func (o *TelephoneNumberAttributeSyntaxResponse) GetRequireBinaryTransfer() bool {
	if o == nil || isNil(o.RequireBinaryTransfer) {
		var ret bool
		return ret
	}
	return *o.RequireBinaryTransfer
}

// GetRequireBinaryTransferOk returns a tuple with the RequireBinaryTransfer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelephoneNumberAttributeSyntaxResponse) GetRequireBinaryTransferOk() (*bool, bool) {
	if o == nil || isNil(o.RequireBinaryTransfer) {
    return nil, false
	}
	return o.RequireBinaryTransfer, true
}

// HasRequireBinaryTransfer returns a boolean if a field has been set.
func (o *TelephoneNumberAttributeSyntaxResponse) HasRequireBinaryTransfer() bool {
	if o != nil && !isNil(o.RequireBinaryTransfer) {
		return true
	}

	return false
}

// SetRequireBinaryTransfer gets a reference to the given bool and assigns it to the RequireBinaryTransfer field.
func (o *TelephoneNumberAttributeSyntaxResponse) SetRequireBinaryTransfer(v bool) {
	o.RequireBinaryTransfer = &v
}

func (o TelephoneNumberAttributeSyntaxResponse) MarshalJSON() ([]byte, error) {
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
	if !isNil(o.StrictFormat) {
		toSerialize["strictFormat"] = o.StrictFormat
	}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	if !isNil(o.RequireBinaryTransfer) {
		toSerialize["requireBinaryTransfer"] = o.RequireBinaryTransfer
	}
	return json.Marshal(toSerialize)
}

type NullableTelephoneNumberAttributeSyntaxResponse struct {
	value *TelephoneNumberAttributeSyntaxResponse
	isSet bool
}

func (v NullableTelephoneNumberAttributeSyntaxResponse) Get() *TelephoneNumberAttributeSyntaxResponse {
	return v.value
}

func (v *NullableTelephoneNumberAttributeSyntaxResponse) Set(val *TelephoneNumberAttributeSyntaxResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTelephoneNumberAttributeSyntaxResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTelephoneNumberAttributeSyntaxResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTelephoneNumberAttributeSyntaxResponse(val *TelephoneNumberAttributeSyntaxResponse) *NullableTelephoneNumberAttributeSyntaxResponse {
	return &NullableTelephoneNumberAttributeSyntaxResponse{value: val, isSet: true}
}

func (v NullableTelephoneNumberAttributeSyntaxResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTelephoneNumberAttributeSyntaxResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


