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

// checks if the AttributeTypeDescriptionAttributeSyntaxResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AttributeTypeDescriptionAttributeSyntaxResponse{}

// AttributeTypeDescriptionAttributeSyntaxResponse struct for AttributeTypeDescriptionAttributeSyntaxResponse
type AttributeTypeDescriptionAttributeSyntaxResponse struct {
	Meta                                          *MetaMeta                                              `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20     `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
	Schemas                                       []EnumattributeTypeDescriptionAttributeSyntaxSchemaUrn `json:"schemas"`
	// Name of the Attribute Syntax
	Id string `json:"id"`
	// Indicates whether the suggested minimum upper bound appended to an attribute's syntax OID in its schema definition Attribute Type Description should be stripped.
	StripSyntaxMinUpperBound *bool `json:"stripSyntaxMinUpperBound,omitempty"`
	// Indicates whether the Attribute Syntax is enabled.
	Enabled bool `json:"enabled"`
	// Indicates whether values of this attribute are required to have a \"binary\" transfer option as described in RFC 4522. Attributes with this syntax will generally be referenced with names including \";binary\" (e.g., \"userCertificate;binary\").
	RequireBinaryTransfer *bool `json:"requireBinaryTransfer,omitempty"`
}

// NewAttributeTypeDescriptionAttributeSyntaxResponse instantiates a new AttributeTypeDescriptionAttributeSyntaxResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAttributeTypeDescriptionAttributeSyntaxResponse(schemas []EnumattributeTypeDescriptionAttributeSyntaxSchemaUrn, id string, enabled bool) *AttributeTypeDescriptionAttributeSyntaxResponse {
	this := AttributeTypeDescriptionAttributeSyntaxResponse{}
	this.Schemas = schemas
	this.Id = id
	this.Enabled = enabled
	return &this
}

// NewAttributeTypeDescriptionAttributeSyntaxResponseWithDefaults instantiates a new AttributeTypeDescriptionAttributeSyntaxResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAttributeTypeDescriptionAttributeSyntaxResponseWithDefaults() *AttributeTypeDescriptionAttributeSyntaxResponse {
	this := AttributeTypeDescriptionAttributeSyntaxResponse{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *AttributeTypeDescriptionAttributeSyntaxResponse) GetMeta() MetaMeta {
	if o == nil || IsNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttributeTypeDescriptionAttributeSyntaxResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *AttributeTypeDescriptionAttributeSyntaxResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *AttributeTypeDescriptionAttributeSyntaxResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *AttributeTypeDescriptionAttributeSyntaxResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttributeTypeDescriptionAttributeSyntaxResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *AttributeTypeDescriptionAttributeSyntaxResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *AttributeTypeDescriptionAttributeSyntaxResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetSchemas returns the Schemas field value
func (o *AttributeTypeDescriptionAttributeSyntaxResponse) GetSchemas() []EnumattributeTypeDescriptionAttributeSyntaxSchemaUrn {
	if o == nil {
		var ret []EnumattributeTypeDescriptionAttributeSyntaxSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AttributeTypeDescriptionAttributeSyntaxResponse) GetSchemasOk() ([]EnumattributeTypeDescriptionAttributeSyntaxSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AttributeTypeDescriptionAttributeSyntaxResponse) SetSchemas(v []EnumattributeTypeDescriptionAttributeSyntaxSchemaUrn) {
	o.Schemas = v
}

// GetId returns the Id field value
func (o *AttributeTypeDescriptionAttributeSyntaxResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AttributeTypeDescriptionAttributeSyntaxResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AttributeTypeDescriptionAttributeSyntaxResponse) SetId(v string) {
	o.Id = v
}

// GetStripSyntaxMinUpperBound returns the StripSyntaxMinUpperBound field value if set, zero value otherwise.
func (o *AttributeTypeDescriptionAttributeSyntaxResponse) GetStripSyntaxMinUpperBound() bool {
	if o == nil || IsNil(o.StripSyntaxMinUpperBound) {
		var ret bool
		return ret
	}
	return *o.StripSyntaxMinUpperBound
}

// GetStripSyntaxMinUpperBoundOk returns a tuple with the StripSyntaxMinUpperBound field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttributeTypeDescriptionAttributeSyntaxResponse) GetStripSyntaxMinUpperBoundOk() (*bool, bool) {
	if o == nil || IsNil(o.StripSyntaxMinUpperBound) {
		return nil, false
	}
	return o.StripSyntaxMinUpperBound, true
}

// HasStripSyntaxMinUpperBound returns a boolean if a field has been set.
func (o *AttributeTypeDescriptionAttributeSyntaxResponse) HasStripSyntaxMinUpperBound() bool {
	if o != nil && !IsNil(o.StripSyntaxMinUpperBound) {
		return true
	}

	return false
}

// SetStripSyntaxMinUpperBound gets a reference to the given bool and assigns it to the StripSyntaxMinUpperBound field.
func (o *AttributeTypeDescriptionAttributeSyntaxResponse) SetStripSyntaxMinUpperBound(v bool) {
	o.StripSyntaxMinUpperBound = &v
}

// GetEnabled returns the Enabled field value
func (o *AttributeTypeDescriptionAttributeSyntaxResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AttributeTypeDescriptionAttributeSyntaxResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AttributeTypeDescriptionAttributeSyntaxResponse) SetEnabled(v bool) {
	o.Enabled = v
}

// GetRequireBinaryTransfer returns the RequireBinaryTransfer field value if set, zero value otherwise.
func (o *AttributeTypeDescriptionAttributeSyntaxResponse) GetRequireBinaryTransfer() bool {
	if o == nil || IsNil(o.RequireBinaryTransfer) {
		var ret bool
		return ret
	}
	return *o.RequireBinaryTransfer
}

// GetRequireBinaryTransferOk returns a tuple with the RequireBinaryTransfer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttributeTypeDescriptionAttributeSyntaxResponse) GetRequireBinaryTransferOk() (*bool, bool) {
	if o == nil || IsNil(o.RequireBinaryTransfer) {
		return nil, false
	}
	return o.RequireBinaryTransfer, true
}

// HasRequireBinaryTransfer returns a boolean if a field has been set.
func (o *AttributeTypeDescriptionAttributeSyntaxResponse) HasRequireBinaryTransfer() bool {
	if o != nil && !IsNil(o.RequireBinaryTransfer) {
		return true
	}

	return false
}

// SetRequireBinaryTransfer gets a reference to the given bool and assigns it to the RequireBinaryTransfer field.
func (o *AttributeTypeDescriptionAttributeSyntaxResponse) SetRequireBinaryTransfer(v bool) {
	o.RequireBinaryTransfer = &v
}

func (o AttributeTypeDescriptionAttributeSyntaxResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AttributeTypeDescriptionAttributeSyntaxResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		toSerialize["urn:pingidentity:schemas:configuration:messages:2.0"] = o.Urnpingidentityschemasconfigurationmessages20
	}
	toSerialize["schemas"] = o.Schemas
	toSerialize["id"] = o.Id
	if !IsNil(o.StripSyntaxMinUpperBound) {
		toSerialize["stripSyntaxMinUpperBound"] = o.StripSyntaxMinUpperBound
	}
	toSerialize["enabled"] = o.Enabled
	if !IsNil(o.RequireBinaryTransfer) {
		toSerialize["requireBinaryTransfer"] = o.RequireBinaryTransfer
	}
	return toSerialize, nil
}

type NullableAttributeTypeDescriptionAttributeSyntaxResponse struct {
	value *AttributeTypeDescriptionAttributeSyntaxResponse
	isSet bool
}

func (v NullableAttributeTypeDescriptionAttributeSyntaxResponse) Get() *AttributeTypeDescriptionAttributeSyntaxResponse {
	return v.value
}

func (v *NullableAttributeTypeDescriptionAttributeSyntaxResponse) Set(val *AttributeTypeDescriptionAttributeSyntaxResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAttributeTypeDescriptionAttributeSyntaxResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAttributeTypeDescriptionAttributeSyntaxResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAttributeTypeDescriptionAttributeSyntaxResponse(val *AttributeTypeDescriptionAttributeSyntaxResponse) *NullableAttributeTypeDescriptionAttributeSyntaxResponse {
	return &NullableAttributeTypeDescriptionAttributeSyntaxResponse{value: val, isSet: true}
}

func (v NullableAttributeTypeDescriptionAttributeSyntaxResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAttributeTypeDescriptionAttributeSyntaxResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
