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

// checks if the AttributeMapperPluginResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AttributeMapperPluginResponse{}

// AttributeMapperPluginResponse struct for AttributeMapperPluginResponse
type AttributeMapperPluginResponse struct {
	Schemas    []EnumattributeMapperPluginSchemaUrn `json:"schemas"`
	PluginType []EnumpluginPluginTypeProp           `json:"pluginType"`
	// Specifies the source attribute type that may appear in client requests which should be remapped to the target attribute. Note that the source attribute type must be defined in the server schema and must not be equal to the target attribute type.
	SourceAttribute string `json:"sourceAttribute"`
	// Specifies the target attribute type to which the source attribute type should be mapped. Note that the target attribute type must be defined in the server schema and must not be equal to the source attribute type.
	TargetAttribute string `json:"targetAttribute"`
	// Indicates whether mapping should be applied to attribute types that may be present in specific controls. If enabled, attribute mapping will only be applied for control types which are specifically supported by the attribute mapper plugin.
	EnableControlMapping bool `json:"enableControlMapping"`
	// Indicates whether the target attribute in response messages should always be remapped back to the source attribute. If this is \"false\", then the mapping will be performed for a response message only if one or more elements of the associated request are mapped. Otherwise, the mapping will be performed for all responses regardless of whether the mapping was applied to the request.
	AlwaysMapResponses bool `json:"alwaysMapResponses"`
	// A description for this Plugin
	Description *string `json:"description,omitempty"`
	// Indicates whether the plug-in is enabled for use.
	Enabled bool `json:"enabled"`
	// Indicates whether the plug-in should be invoked for internal operations.
	InvokeForInternalOperations                   *bool                                              `json:"invokeForInternalOperations,omitempty"`
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
	// Name of the Plugin
	Id string `json:"id"`
}

// NewAttributeMapperPluginResponse instantiates a new AttributeMapperPluginResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAttributeMapperPluginResponse(schemas []EnumattributeMapperPluginSchemaUrn, pluginType []EnumpluginPluginTypeProp, sourceAttribute string, targetAttribute string, enableControlMapping bool, alwaysMapResponses bool, enabled bool, id string) *AttributeMapperPluginResponse {
	this := AttributeMapperPluginResponse{}
	this.Schemas = schemas
	this.PluginType = pluginType
	this.SourceAttribute = sourceAttribute
	this.TargetAttribute = targetAttribute
	this.EnableControlMapping = enableControlMapping
	this.AlwaysMapResponses = alwaysMapResponses
	this.Enabled = enabled
	this.Id = id
	return &this
}

// NewAttributeMapperPluginResponseWithDefaults instantiates a new AttributeMapperPluginResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAttributeMapperPluginResponseWithDefaults() *AttributeMapperPluginResponse {
	this := AttributeMapperPluginResponse{}
	return &this
}

// GetSchemas returns the Schemas field value
func (o *AttributeMapperPluginResponse) GetSchemas() []EnumattributeMapperPluginSchemaUrn {
	if o == nil {
		var ret []EnumattributeMapperPluginSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AttributeMapperPluginResponse) GetSchemasOk() ([]EnumattributeMapperPluginSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AttributeMapperPluginResponse) SetSchemas(v []EnumattributeMapperPluginSchemaUrn) {
	o.Schemas = v
}

// GetPluginType returns the PluginType field value
func (o *AttributeMapperPluginResponse) GetPluginType() []EnumpluginPluginTypeProp {
	if o == nil {
		var ret []EnumpluginPluginTypeProp
		return ret
	}

	return o.PluginType
}

// GetPluginTypeOk returns a tuple with the PluginType field value
// and a boolean to check if the value has been set.
func (o *AttributeMapperPluginResponse) GetPluginTypeOk() ([]EnumpluginPluginTypeProp, bool) {
	if o == nil {
		return nil, false
	}
	return o.PluginType, true
}

// SetPluginType sets field value
func (o *AttributeMapperPluginResponse) SetPluginType(v []EnumpluginPluginTypeProp) {
	o.PluginType = v
}

// GetSourceAttribute returns the SourceAttribute field value
func (o *AttributeMapperPluginResponse) GetSourceAttribute() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceAttribute
}

// GetSourceAttributeOk returns a tuple with the SourceAttribute field value
// and a boolean to check if the value has been set.
func (o *AttributeMapperPluginResponse) GetSourceAttributeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceAttribute, true
}

// SetSourceAttribute sets field value
func (o *AttributeMapperPluginResponse) SetSourceAttribute(v string) {
	o.SourceAttribute = v
}

// GetTargetAttribute returns the TargetAttribute field value
func (o *AttributeMapperPluginResponse) GetTargetAttribute() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TargetAttribute
}

// GetTargetAttributeOk returns a tuple with the TargetAttribute field value
// and a boolean to check if the value has been set.
func (o *AttributeMapperPluginResponse) GetTargetAttributeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetAttribute, true
}

// SetTargetAttribute sets field value
func (o *AttributeMapperPluginResponse) SetTargetAttribute(v string) {
	o.TargetAttribute = v
}

// GetEnableControlMapping returns the EnableControlMapping field value
func (o *AttributeMapperPluginResponse) GetEnableControlMapping() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.EnableControlMapping
}

// GetEnableControlMappingOk returns a tuple with the EnableControlMapping field value
// and a boolean to check if the value has been set.
func (o *AttributeMapperPluginResponse) GetEnableControlMappingOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnableControlMapping, true
}

// SetEnableControlMapping sets field value
func (o *AttributeMapperPluginResponse) SetEnableControlMapping(v bool) {
	o.EnableControlMapping = v
}

// GetAlwaysMapResponses returns the AlwaysMapResponses field value
func (o *AttributeMapperPluginResponse) GetAlwaysMapResponses() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AlwaysMapResponses
}

// GetAlwaysMapResponsesOk returns a tuple with the AlwaysMapResponses field value
// and a boolean to check if the value has been set.
func (o *AttributeMapperPluginResponse) GetAlwaysMapResponsesOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AlwaysMapResponses, true
}

// SetAlwaysMapResponses sets field value
func (o *AttributeMapperPluginResponse) SetAlwaysMapResponses(v bool) {
	o.AlwaysMapResponses = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AttributeMapperPluginResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttributeMapperPluginResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AttributeMapperPluginResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AttributeMapperPluginResponse) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *AttributeMapperPluginResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AttributeMapperPluginResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AttributeMapperPluginResponse) SetEnabled(v bool) {
	o.Enabled = v
}

// GetInvokeForInternalOperations returns the InvokeForInternalOperations field value if set, zero value otherwise.
func (o *AttributeMapperPluginResponse) GetInvokeForInternalOperations() bool {
	if o == nil || IsNil(o.InvokeForInternalOperations) {
		var ret bool
		return ret
	}
	return *o.InvokeForInternalOperations
}

// GetInvokeForInternalOperationsOk returns a tuple with the InvokeForInternalOperations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttributeMapperPluginResponse) GetInvokeForInternalOperationsOk() (*bool, bool) {
	if o == nil || IsNil(o.InvokeForInternalOperations) {
		return nil, false
	}
	return o.InvokeForInternalOperations, true
}

// HasInvokeForInternalOperations returns a boolean if a field has been set.
func (o *AttributeMapperPluginResponse) HasInvokeForInternalOperations() bool {
	if o != nil && !IsNil(o.InvokeForInternalOperations) {
		return true
	}

	return false
}

// SetInvokeForInternalOperations gets a reference to the given bool and assigns it to the InvokeForInternalOperations field.
func (o *AttributeMapperPluginResponse) SetInvokeForInternalOperations(v bool) {
	o.InvokeForInternalOperations = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *AttributeMapperPluginResponse) GetMeta() MetaMeta {
	if o == nil || IsNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttributeMapperPluginResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *AttributeMapperPluginResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *AttributeMapperPluginResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *AttributeMapperPluginResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttributeMapperPluginResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *AttributeMapperPluginResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *AttributeMapperPluginResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetId returns the Id field value
func (o *AttributeMapperPluginResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AttributeMapperPluginResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AttributeMapperPluginResponse) SetId(v string) {
	o.Id = v
}

func (o AttributeMapperPluginResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AttributeMapperPluginResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["schemas"] = o.Schemas
	toSerialize["pluginType"] = o.PluginType
	toSerialize["sourceAttribute"] = o.SourceAttribute
	toSerialize["targetAttribute"] = o.TargetAttribute
	toSerialize["enableControlMapping"] = o.EnableControlMapping
	toSerialize["alwaysMapResponses"] = o.AlwaysMapResponses
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["enabled"] = o.Enabled
	if !IsNil(o.InvokeForInternalOperations) {
		toSerialize["invokeForInternalOperations"] = o.InvokeForInternalOperations
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		toSerialize["urn:pingidentity:schemas:configuration:messages:2.0"] = o.Urnpingidentityschemasconfigurationmessages20
	}
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

type NullableAttributeMapperPluginResponse struct {
	value *AttributeMapperPluginResponse
	isSet bool
}

func (v NullableAttributeMapperPluginResponse) Get() *AttributeMapperPluginResponse {
	return v.value
}

func (v *NullableAttributeMapperPluginResponse) Set(val *AttributeMapperPluginResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAttributeMapperPluginResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAttributeMapperPluginResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAttributeMapperPluginResponse(val *AttributeMapperPluginResponse) *NullableAttributeMapperPluginResponse {
	return &NullableAttributeMapperPluginResponse{value: val, isSet: true}
}

func (v NullableAttributeMapperPluginResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAttributeMapperPluginResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
