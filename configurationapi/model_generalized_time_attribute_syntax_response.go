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

// checks if the GeneralizedTimeAttributeSyntaxResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GeneralizedTimeAttributeSyntaxResponse{}

// GeneralizedTimeAttributeSyntaxResponse struct for GeneralizedTimeAttributeSyntaxResponse
type GeneralizedTimeAttributeSyntaxResponse struct {
	Schemas []EnumgeneralizedTimeAttributeSyntaxSchemaUrn `json:"schemas"`
	// Name of the Attribute Syntax
	Id string `json:"id"`
	// Indicates whether values of attributes with this syntax should be compacted when stored in a local DB database.
	EnableCompaction *bool `json:"enableCompaction,omitempty"`
	// Specifies the specific attributes (which should be associated with this syntax) whose values should be compacted. If one or more include attributes are specified, then only those attributes will have their values compacted. If not set then all attributes will have their values compacted. The exclude-attribute-from-compaction property takes precedence over this property.
	IncludeAttributeInCompaction []string `json:"includeAttributeInCompaction,omitempty"`
	// Specifies the specific attributes (which should be associated with this syntax) whose values should not be compacted. If one or more exclude attributes are specified, then values of those attributes will not have their values compacted. This property takes precedence over the include-attribute-in-compaction property.
	ExcludeAttributeFromCompaction []string `json:"excludeAttributeFromCompaction,omitempty"`
	// Indicates whether the Attribute Syntax is enabled.
	Enabled bool `json:"enabled"`
	// Indicates whether values of this attribute are required to have a \"binary\" transfer option as described in RFC 4522. Attributes with this syntax will generally be referenced with names including \";binary\" (e.g., \"userCertificate;binary\").
	RequireBinaryTransfer                         *bool                                              `json:"requireBinaryTransfer,omitempty"`
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
}

// NewGeneralizedTimeAttributeSyntaxResponse instantiates a new GeneralizedTimeAttributeSyntaxResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGeneralizedTimeAttributeSyntaxResponse(schemas []EnumgeneralizedTimeAttributeSyntaxSchemaUrn, id string, enabled bool) *GeneralizedTimeAttributeSyntaxResponse {
	this := GeneralizedTimeAttributeSyntaxResponse{}
	this.Schemas = schemas
	this.Id = id
	this.Enabled = enabled
	return &this
}

// NewGeneralizedTimeAttributeSyntaxResponseWithDefaults instantiates a new GeneralizedTimeAttributeSyntaxResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGeneralizedTimeAttributeSyntaxResponseWithDefaults() *GeneralizedTimeAttributeSyntaxResponse {
	this := GeneralizedTimeAttributeSyntaxResponse{}
	return &this
}

// GetSchemas returns the Schemas field value
func (o *GeneralizedTimeAttributeSyntaxResponse) GetSchemas() []EnumgeneralizedTimeAttributeSyntaxSchemaUrn {
	if o == nil {
		var ret []EnumgeneralizedTimeAttributeSyntaxSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *GeneralizedTimeAttributeSyntaxResponse) GetSchemasOk() ([]EnumgeneralizedTimeAttributeSyntaxSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *GeneralizedTimeAttributeSyntaxResponse) SetSchemas(v []EnumgeneralizedTimeAttributeSyntaxSchemaUrn) {
	o.Schemas = v
}

// GetId returns the Id field value
func (o *GeneralizedTimeAttributeSyntaxResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GeneralizedTimeAttributeSyntaxResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GeneralizedTimeAttributeSyntaxResponse) SetId(v string) {
	o.Id = v
}

// GetEnableCompaction returns the EnableCompaction field value if set, zero value otherwise.
func (o *GeneralizedTimeAttributeSyntaxResponse) GetEnableCompaction() bool {
	if o == nil || IsNil(o.EnableCompaction) {
		var ret bool
		return ret
	}
	return *o.EnableCompaction
}

// GetEnableCompactionOk returns a tuple with the EnableCompaction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeneralizedTimeAttributeSyntaxResponse) GetEnableCompactionOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableCompaction) {
		return nil, false
	}
	return o.EnableCompaction, true
}

// HasEnableCompaction returns a boolean if a field has been set.
func (o *GeneralizedTimeAttributeSyntaxResponse) HasEnableCompaction() bool {
	if o != nil && !IsNil(o.EnableCompaction) {
		return true
	}

	return false
}

// SetEnableCompaction gets a reference to the given bool and assigns it to the EnableCompaction field.
func (o *GeneralizedTimeAttributeSyntaxResponse) SetEnableCompaction(v bool) {
	o.EnableCompaction = &v
}

// GetIncludeAttributeInCompaction returns the IncludeAttributeInCompaction field value if set, zero value otherwise.
func (o *GeneralizedTimeAttributeSyntaxResponse) GetIncludeAttributeInCompaction() []string {
	if o == nil || IsNil(o.IncludeAttributeInCompaction) {
		var ret []string
		return ret
	}
	return o.IncludeAttributeInCompaction
}

// GetIncludeAttributeInCompactionOk returns a tuple with the IncludeAttributeInCompaction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeneralizedTimeAttributeSyntaxResponse) GetIncludeAttributeInCompactionOk() ([]string, bool) {
	if o == nil || IsNil(o.IncludeAttributeInCompaction) {
		return nil, false
	}
	return o.IncludeAttributeInCompaction, true
}

// HasIncludeAttributeInCompaction returns a boolean if a field has been set.
func (o *GeneralizedTimeAttributeSyntaxResponse) HasIncludeAttributeInCompaction() bool {
	if o != nil && !IsNil(o.IncludeAttributeInCompaction) {
		return true
	}

	return false
}

// SetIncludeAttributeInCompaction gets a reference to the given []string and assigns it to the IncludeAttributeInCompaction field.
func (o *GeneralizedTimeAttributeSyntaxResponse) SetIncludeAttributeInCompaction(v []string) {
	o.IncludeAttributeInCompaction = v
}

// GetExcludeAttributeFromCompaction returns the ExcludeAttributeFromCompaction field value if set, zero value otherwise.
func (o *GeneralizedTimeAttributeSyntaxResponse) GetExcludeAttributeFromCompaction() []string {
	if o == nil || IsNil(o.ExcludeAttributeFromCompaction) {
		var ret []string
		return ret
	}
	return o.ExcludeAttributeFromCompaction
}

// GetExcludeAttributeFromCompactionOk returns a tuple with the ExcludeAttributeFromCompaction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeneralizedTimeAttributeSyntaxResponse) GetExcludeAttributeFromCompactionOk() ([]string, bool) {
	if o == nil || IsNil(o.ExcludeAttributeFromCompaction) {
		return nil, false
	}
	return o.ExcludeAttributeFromCompaction, true
}

// HasExcludeAttributeFromCompaction returns a boolean if a field has been set.
func (o *GeneralizedTimeAttributeSyntaxResponse) HasExcludeAttributeFromCompaction() bool {
	if o != nil && !IsNil(o.ExcludeAttributeFromCompaction) {
		return true
	}

	return false
}

// SetExcludeAttributeFromCompaction gets a reference to the given []string and assigns it to the ExcludeAttributeFromCompaction field.
func (o *GeneralizedTimeAttributeSyntaxResponse) SetExcludeAttributeFromCompaction(v []string) {
	o.ExcludeAttributeFromCompaction = v
}

// GetEnabled returns the Enabled field value
func (o *GeneralizedTimeAttributeSyntaxResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *GeneralizedTimeAttributeSyntaxResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *GeneralizedTimeAttributeSyntaxResponse) SetEnabled(v bool) {
	o.Enabled = v
}

// GetRequireBinaryTransfer returns the RequireBinaryTransfer field value if set, zero value otherwise.
func (o *GeneralizedTimeAttributeSyntaxResponse) GetRequireBinaryTransfer() bool {
	if o == nil || IsNil(o.RequireBinaryTransfer) {
		var ret bool
		return ret
	}
	return *o.RequireBinaryTransfer
}

// GetRequireBinaryTransferOk returns a tuple with the RequireBinaryTransfer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeneralizedTimeAttributeSyntaxResponse) GetRequireBinaryTransferOk() (*bool, bool) {
	if o == nil || IsNil(o.RequireBinaryTransfer) {
		return nil, false
	}
	return o.RequireBinaryTransfer, true
}

// HasRequireBinaryTransfer returns a boolean if a field has been set.
func (o *GeneralizedTimeAttributeSyntaxResponse) HasRequireBinaryTransfer() bool {
	if o != nil && !IsNil(o.RequireBinaryTransfer) {
		return true
	}

	return false
}

// SetRequireBinaryTransfer gets a reference to the given bool and assigns it to the RequireBinaryTransfer field.
func (o *GeneralizedTimeAttributeSyntaxResponse) SetRequireBinaryTransfer(v bool) {
	o.RequireBinaryTransfer = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *GeneralizedTimeAttributeSyntaxResponse) GetMeta() MetaMeta {
	if o == nil || IsNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeneralizedTimeAttributeSyntaxResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *GeneralizedTimeAttributeSyntaxResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *GeneralizedTimeAttributeSyntaxResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *GeneralizedTimeAttributeSyntaxResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeneralizedTimeAttributeSyntaxResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *GeneralizedTimeAttributeSyntaxResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *GeneralizedTimeAttributeSyntaxResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

func (o GeneralizedTimeAttributeSyntaxResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GeneralizedTimeAttributeSyntaxResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["schemas"] = o.Schemas
	toSerialize["id"] = o.Id
	if !IsNil(o.EnableCompaction) {
		toSerialize["enableCompaction"] = o.EnableCompaction
	}
	if !IsNil(o.IncludeAttributeInCompaction) {
		toSerialize["includeAttributeInCompaction"] = o.IncludeAttributeInCompaction
	}
	if !IsNil(o.ExcludeAttributeFromCompaction) {
		toSerialize["excludeAttributeFromCompaction"] = o.ExcludeAttributeFromCompaction
	}
	toSerialize["enabled"] = o.Enabled
	if !IsNil(o.RequireBinaryTransfer) {
		toSerialize["requireBinaryTransfer"] = o.RequireBinaryTransfer
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		toSerialize["urn:pingidentity:schemas:configuration:messages:2.0"] = o.Urnpingidentityschemasconfigurationmessages20
	}
	return toSerialize, nil
}

type NullableGeneralizedTimeAttributeSyntaxResponse struct {
	value *GeneralizedTimeAttributeSyntaxResponse
	isSet bool
}

func (v NullableGeneralizedTimeAttributeSyntaxResponse) Get() *GeneralizedTimeAttributeSyntaxResponse {
	return v.value
}

func (v *NullableGeneralizedTimeAttributeSyntaxResponse) Set(val *GeneralizedTimeAttributeSyntaxResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGeneralizedTimeAttributeSyntaxResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGeneralizedTimeAttributeSyntaxResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGeneralizedTimeAttributeSyntaxResponse(val *GeneralizedTimeAttributeSyntaxResponse) *NullableGeneralizedTimeAttributeSyntaxResponse {
	return &NullableGeneralizedTimeAttributeSyntaxResponse{value: val, isSet: true}
}

func (v NullableGeneralizedTimeAttributeSyntaxResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGeneralizedTimeAttributeSyntaxResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
