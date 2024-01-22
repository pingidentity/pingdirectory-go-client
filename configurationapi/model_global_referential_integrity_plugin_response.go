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

// checks if the GlobalReferentialIntegrityPluginResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GlobalReferentialIntegrityPluginResponse{}

// GlobalReferentialIntegrityPluginResponse struct for GlobalReferentialIntegrityPluginResponse
type GlobalReferentialIntegrityPluginResponse struct {
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
	Schemas                                       []EnumglobalReferentialIntegrityPluginSchemaUrn    `json:"schemas"`
	// Name of the Plugin
	Id string `json:"id"`
	// The attribute type(s) for which to maintain referential integrity. The attribute must have a distinguished name or a name and optional UID syntax and must be indexed for equality searches in all subtree views for which referential integrity is to be maintained.
	AttributeType []string `json:"attributeType"`
	// The subtree view(s) for which to maintain referential integrity.
	SubtreeView []string `json:"subtreeView"`
	// A description for this Plugin
	Description *string `json:"description,omitempty"`
	// Indicates whether the plug-in is enabled for use.
	Enabled bool `json:"enabled"`
}

// NewGlobalReferentialIntegrityPluginResponse instantiates a new GlobalReferentialIntegrityPluginResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGlobalReferentialIntegrityPluginResponse(schemas []EnumglobalReferentialIntegrityPluginSchemaUrn, id string, attributeType []string, subtreeView []string, enabled bool) *GlobalReferentialIntegrityPluginResponse {
	this := GlobalReferentialIntegrityPluginResponse{}
	this.Schemas = schemas
	this.Id = id
	this.AttributeType = attributeType
	this.SubtreeView = subtreeView
	this.Enabled = enabled
	return &this
}

// NewGlobalReferentialIntegrityPluginResponseWithDefaults instantiates a new GlobalReferentialIntegrityPluginResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGlobalReferentialIntegrityPluginResponseWithDefaults() *GlobalReferentialIntegrityPluginResponse {
	this := GlobalReferentialIntegrityPluginResponse{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *GlobalReferentialIntegrityPluginResponse) GetMeta() MetaMeta {
	if o == nil || IsNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalReferentialIntegrityPluginResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *GlobalReferentialIntegrityPluginResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *GlobalReferentialIntegrityPluginResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *GlobalReferentialIntegrityPluginResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalReferentialIntegrityPluginResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *GlobalReferentialIntegrityPluginResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *GlobalReferentialIntegrityPluginResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetSchemas returns the Schemas field value
func (o *GlobalReferentialIntegrityPluginResponse) GetSchemas() []EnumglobalReferentialIntegrityPluginSchemaUrn {
	if o == nil {
		var ret []EnumglobalReferentialIntegrityPluginSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *GlobalReferentialIntegrityPluginResponse) GetSchemasOk() ([]EnumglobalReferentialIntegrityPluginSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *GlobalReferentialIntegrityPluginResponse) SetSchemas(v []EnumglobalReferentialIntegrityPluginSchemaUrn) {
	o.Schemas = v
}

// GetId returns the Id field value
func (o *GlobalReferentialIntegrityPluginResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GlobalReferentialIntegrityPluginResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GlobalReferentialIntegrityPluginResponse) SetId(v string) {
	o.Id = v
}

// GetAttributeType returns the AttributeType field value
func (o *GlobalReferentialIntegrityPluginResponse) GetAttributeType() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.AttributeType
}

// GetAttributeTypeOk returns a tuple with the AttributeType field value
// and a boolean to check if the value has been set.
func (o *GlobalReferentialIntegrityPluginResponse) GetAttributeTypeOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AttributeType, true
}

// SetAttributeType sets field value
func (o *GlobalReferentialIntegrityPluginResponse) SetAttributeType(v []string) {
	o.AttributeType = v
}

// GetSubtreeView returns the SubtreeView field value
func (o *GlobalReferentialIntegrityPluginResponse) GetSubtreeView() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.SubtreeView
}

// GetSubtreeViewOk returns a tuple with the SubtreeView field value
// and a boolean to check if the value has been set.
func (o *GlobalReferentialIntegrityPluginResponse) GetSubtreeViewOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SubtreeView, true
}

// SetSubtreeView sets field value
func (o *GlobalReferentialIntegrityPluginResponse) SetSubtreeView(v []string) {
	o.SubtreeView = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *GlobalReferentialIntegrityPluginResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalReferentialIntegrityPluginResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *GlobalReferentialIntegrityPluginResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *GlobalReferentialIntegrityPluginResponse) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *GlobalReferentialIntegrityPluginResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *GlobalReferentialIntegrityPluginResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *GlobalReferentialIntegrityPluginResponse) SetEnabled(v bool) {
	o.Enabled = v
}

func (o GlobalReferentialIntegrityPluginResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GlobalReferentialIntegrityPluginResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		toSerialize["urn:pingidentity:schemas:configuration:messages:2.0"] = o.Urnpingidentityschemasconfigurationmessages20
	}
	toSerialize["schemas"] = o.Schemas
	toSerialize["id"] = o.Id
	toSerialize["attributeType"] = o.AttributeType
	toSerialize["subtreeView"] = o.SubtreeView
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["enabled"] = o.Enabled
	return toSerialize, nil
}

type NullableGlobalReferentialIntegrityPluginResponse struct {
	value *GlobalReferentialIntegrityPluginResponse
	isSet bool
}

func (v NullableGlobalReferentialIntegrityPluginResponse) Get() *GlobalReferentialIntegrityPluginResponse {
	return v.value
}

func (v *NullableGlobalReferentialIntegrityPluginResponse) Set(val *GlobalReferentialIntegrityPluginResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGlobalReferentialIntegrityPluginResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGlobalReferentialIntegrityPluginResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGlobalReferentialIntegrityPluginResponse(val *GlobalReferentialIntegrityPluginResponse) *NullableGlobalReferentialIntegrityPluginResponse {
	return &NullableGlobalReferentialIntegrityPluginResponse{value: val, isSet: true}
}

func (v NullableGlobalReferentialIntegrityPluginResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGlobalReferentialIntegrityPluginResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}