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

// checks if the LdapUrlAttributeSyntaxResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LdapUrlAttributeSyntaxResponse{}

// LdapUrlAttributeSyntaxResponse struct for LdapUrlAttributeSyntaxResponse
type LdapUrlAttributeSyntaxResponse struct {
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
	Schemas                                       []EnumldapUrlAttributeSyntaxSchemaUrn              `json:"schemas"`
	// Name of the Attribute Syntax
	Id string `json:"id"`
	// Indicates whether values for attributes with this syntax will be required to be in the valid LDAP URL format. If this is set to false, then arbitrary strings will be allowed.
	StrictFormat *bool `json:"strictFormat,omitempty"`
	// Indicates whether the Attribute Syntax is enabled.
	Enabled bool `json:"enabled"`
	// Indicates whether values of this attribute are required to have a \"binary\" transfer option as described in RFC 4522. Attributes with this syntax will generally be referenced with names including \";binary\" (e.g., \"userCertificate;binary\").
	RequireBinaryTransfer *bool `json:"requireBinaryTransfer,omitempty"`
}

// NewLdapUrlAttributeSyntaxResponse instantiates a new LdapUrlAttributeSyntaxResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLdapUrlAttributeSyntaxResponse(schemas []EnumldapUrlAttributeSyntaxSchemaUrn, id string, enabled bool) *LdapUrlAttributeSyntaxResponse {
	this := LdapUrlAttributeSyntaxResponse{}
	this.Schemas = schemas
	this.Id = id
	this.Enabled = enabled
	return &this
}

// NewLdapUrlAttributeSyntaxResponseWithDefaults instantiates a new LdapUrlAttributeSyntaxResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLdapUrlAttributeSyntaxResponseWithDefaults() *LdapUrlAttributeSyntaxResponse {
	this := LdapUrlAttributeSyntaxResponse{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *LdapUrlAttributeSyntaxResponse) GetMeta() MetaMeta {
	if o == nil || IsNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapUrlAttributeSyntaxResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *LdapUrlAttributeSyntaxResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *LdapUrlAttributeSyntaxResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *LdapUrlAttributeSyntaxResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapUrlAttributeSyntaxResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *LdapUrlAttributeSyntaxResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *LdapUrlAttributeSyntaxResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetSchemas returns the Schemas field value
func (o *LdapUrlAttributeSyntaxResponse) GetSchemas() []EnumldapUrlAttributeSyntaxSchemaUrn {
	if o == nil {
		var ret []EnumldapUrlAttributeSyntaxSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *LdapUrlAttributeSyntaxResponse) GetSchemasOk() ([]EnumldapUrlAttributeSyntaxSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *LdapUrlAttributeSyntaxResponse) SetSchemas(v []EnumldapUrlAttributeSyntaxSchemaUrn) {
	o.Schemas = v
}

// GetId returns the Id field value
func (o *LdapUrlAttributeSyntaxResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *LdapUrlAttributeSyntaxResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *LdapUrlAttributeSyntaxResponse) SetId(v string) {
	o.Id = v
}

// GetStrictFormat returns the StrictFormat field value if set, zero value otherwise.
func (o *LdapUrlAttributeSyntaxResponse) GetStrictFormat() bool {
	if o == nil || IsNil(o.StrictFormat) {
		var ret bool
		return ret
	}
	return *o.StrictFormat
}

// GetStrictFormatOk returns a tuple with the StrictFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapUrlAttributeSyntaxResponse) GetStrictFormatOk() (*bool, bool) {
	if o == nil || IsNil(o.StrictFormat) {
		return nil, false
	}
	return o.StrictFormat, true
}

// HasStrictFormat returns a boolean if a field has been set.
func (o *LdapUrlAttributeSyntaxResponse) HasStrictFormat() bool {
	if o != nil && !IsNil(o.StrictFormat) {
		return true
	}

	return false
}

// SetStrictFormat gets a reference to the given bool and assigns it to the StrictFormat field.
func (o *LdapUrlAttributeSyntaxResponse) SetStrictFormat(v bool) {
	o.StrictFormat = &v
}

// GetEnabled returns the Enabled field value
func (o *LdapUrlAttributeSyntaxResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *LdapUrlAttributeSyntaxResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *LdapUrlAttributeSyntaxResponse) SetEnabled(v bool) {
	o.Enabled = v
}

// GetRequireBinaryTransfer returns the RequireBinaryTransfer field value if set, zero value otherwise.
func (o *LdapUrlAttributeSyntaxResponse) GetRequireBinaryTransfer() bool {
	if o == nil || IsNil(o.RequireBinaryTransfer) {
		var ret bool
		return ret
	}
	return *o.RequireBinaryTransfer
}

// GetRequireBinaryTransferOk returns a tuple with the RequireBinaryTransfer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapUrlAttributeSyntaxResponse) GetRequireBinaryTransferOk() (*bool, bool) {
	if o == nil || IsNil(o.RequireBinaryTransfer) {
		return nil, false
	}
	return o.RequireBinaryTransfer, true
}

// HasRequireBinaryTransfer returns a boolean if a field has been set.
func (o *LdapUrlAttributeSyntaxResponse) HasRequireBinaryTransfer() bool {
	if o != nil && !IsNil(o.RequireBinaryTransfer) {
		return true
	}

	return false
}

// SetRequireBinaryTransfer gets a reference to the given bool and assigns it to the RequireBinaryTransfer field.
func (o *LdapUrlAttributeSyntaxResponse) SetRequireBinaryTransfer(v bool) {
	o.RequireBinaryTransfer = &v
}

func (o LdapUrlAttributeSyntaxResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LdapUrlAttributeSyntaxResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		toSerialize["urn:pingidentity:schemas:configuration:messages:2.0"] = o.Urnpingidentityschemasconfigurationmessages20
	}
	toSerialize["schemas"] = o.Schemas
	toSerialize["id"] = o.Id
	if !IsNil(o.StrictFormat) {
		toSerialize["strictFormat"] = o.StrictFormat
	}
	toSerialize["enabled"] = o.Enabled
	if !IsNil(o.RequireBinaryTransfer) {
		toSerialize["requireBinaryTransfer"] = o.RequireBinaryTransfer
	}
	return toSerialize, nil
}

type NullableLdapUrlAttributeSyntaxResponse struct {
	value *LdapUrlAttributeSyntaxResponse
	isSet bool
}

func (v NullableLdapUrlAttributeSyntaxResponse) Get() *LdapUrlAttributeSyntaxResponse {
	return v.value
}

func (v *NullableLdapUrlAttributeSyntaxResponse) Set(val *LdapUrlAttributeSyntaxResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableLdapUrlAttributeSyntaxResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableLdapUrlAttributeSyntaxResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLdapUrlAttributeSyntaxResponse(val *LdapUrlAttributeSyntaxResponse) *NullableLdapUrlAttributeSyntaxResponse {
	return &NullableLdapUrlAttributeSyntaxResponse{value: val, isSet: true}
}

func (v NullableLdapUrlAttributeSyntaxResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLdapUrlAttributeSyntaxResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
