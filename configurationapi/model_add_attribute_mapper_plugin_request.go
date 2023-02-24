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

// checks if the AddAttributeMapperPluginRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddAttributeMapperPluginRequest{}

// AddAttributeMapperPluginRequest struct for AddAttributeMapperPluginRequest
type AddAttributeMapperPluginRequest struct {
	// Name of the new Plugin
	PluginName string                               `json:"pluginName"`
	Schemas    []EnumattributeMapperPluginSchemaUrn `json:"schemas"`
	PluginType []EnumpluginPluginTypeProp           `json:"pluginType,omitempty"`
	// Specifies the source attribute type that may appear in client requests which should be remapped to the target attribute. Note that the source attribute type must be defined in the server schema and must not be equal to the target attribute type.
	SourceAttribute string `json:"sourceAttribute"`
	// Specifies the target attribute type to which the source attribute type should be mapped. Note that the target attribute type must be defined in the server schema and must not be equal to the source attribute type.
	TargetAttribute string `json:"targetAttribute"`
	// Indicates whether mapping should be applied to attribute types that may be present in specific controls. If enabled, attribute mapping will only be applied for control types which are specifically supported by the attribute mapper plugin.
	EnableControlMapping *bool `json:"enableControlMapping,omitempty"`
	// Indicates whether the target attribute in response messages should always be remapped back to the source attribute. If this is \"false\", then the mapping will be performed for a response message only if one or more elements of the associated request are mapped. Otherwise, the mapping will be performed for all responses regardless of whether the mapping was applied to the request.
	AlwaysMapResponses *bool `json:"alwaysMapResponses,omitempty"`
	// A description for this Plugin
	Description *string `json:"description,omitempty"`
	// Indicates whether the plug-in is enabled for use.
	Enabled bool `json:"enabled"`
	// Indicates whether the plug-in should be invoked for internal operations.
	InvokeForInternalOperations *bool `json:"invokeForInternalOperations,omitempty"`
}

// NewAddAttributeMapperPluginRequest instantiates a new AddAttributeMapperPluginRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddAttributeMapperPluginRequest(pluginName string, schemas []EnumattributeMapperPluginSchemaUrn, sourceAttribute string, targetAttribute string, enabled bool) *AddAttributeMapperPluginRequest {
	this := AddAttributeMapperPluginRequest{}
	this.PluginName = pluginName
	this.Schemas = schemas
	this.SourceAttribute = sourceAttribute
	this.TargetAttribute = targetAttribute
	this.Enabled = enabled
	return &this
}

// NewAddAttributeMapperPluginRequestWithDefaults instantiates a new AddAttributeMapperPluginRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddAttributeMapperPluginRequestWithDefaults() *AddAttributeMapperPluginRequest {
	this := AddAttributeMapperPluginRequest{}
	return &this
}

// GetPluginName returns the PluginName field value
func (o *AddAttributeMapperPluginRequest) GetPluginName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PluginName
}

// GetPluginNameOk returns a tuple with the PluginName field value
// and a boolean to check if the value has been set.
func (o *AddAttributeMapperPluginRequest) GetPluginNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PluginName, true
}

// SetPluginName sets field value
func (o *AddAttributeMapperPluginRequest) SetPluginName(v string) {
	o.PluginName = v
}

// GetSchemas returns the Schemas field value
func (o *AddAttributeMapperPluginRequest) GetSchemas() []EnumattributeMapperPluginSchemaUrn {
	if o == nil {
		var ret []EnumattributeMapperPluginSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddAttributeMapperPluginRequest) GetSchemasOk() ([]EnumattributeMapperPluginSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddAttributeMapperPluginRequest) SetSchemas(v []EnumattributeMapperPluginSchemaUrn) {
	o.Schemas = v
}

// GetPluginType returns the PluginType field value if set, zero value otherwise.
func (o *AddAttributeMapperPluginRequest) GetPluginType() []EnumpluginPluginTypeProp {
	if o == nil || IsNil(o.PluginType) {
		var ret []EnumpluginPluginTypeProp
		return ret
	}
	return o.PluginType
}

// GetPluginTypeOk returns a tuple with the PluginType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddAttributeMapperPluginRequest) GetPluginTypeOk() ([]EnumpluginPluginTypeProp, bool) {
	if o == nil || IsNil(o.PluginType) {
		return nil, false
	}
	return o.PluginType, true
}

// HasPluginType returns a boolean if a field has been set.
func (o *AddAttributeMapperPluginRequest) HasPluginType() bool {
	if o != nil && !IsNil(o.PluginType) {
		return true
	}

	return false
}

// SetPluginType gets a reference to the given []EnumpluginPluginTypeProp and assigns it to the PluginType field.
func (o *AddAttributeMapperPluginRequest) SetPluginType(v []EnumpluginPluginTypeProp) {
	o.PluginType = v
}

// GetSourceAttribute returns the SourceAttribute field value
func (o *AddAttributeMapperPluginRequest) GetSourceAttribute() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceAttribute
}

// GetSourceAttributeOk returns a tuple with the SourceAttribute field value
// and a boolean to check if the value has been set.
func (o *AddAttributeMapperPluginRequest) GetSourceAttributeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceAttribute, true
}

// SetSourceAttribute sets field value
func (o *AddAttributeMapperPluginRequest) SetSourceAttribute(v string) {
	o.SourceAttribute = v
}

// GetTargetAttribute returns the TargetAttribute field value
func (o *AddAttributeMapperPluginRequest) GetTargetAttribute() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TargetAttribute
}

// GetTargetAttributeOk returns a tuple with the TargetAttribute field value
// and a boolean to check if the value has been set.
func (o *AddAttributeMapperPluginRequest) GetTargetAttributeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetAttribute, true
}

// SetTargetAttribute sets field value
func (o *AddAttributeMapperPluginRequest) SetTargetAttribute(v string) {
	o.TargetAttribute = v
}

// GetEnableControlMapping returns the EnableControlMapping field value if set, zero value otherwise.
func (o *AddAttributeMapperPluginRequest) GetEnableControlMapping() bool {
	if o == nil || IsNil(o.EnableControlMapping) {
		var ret bool
		return ret
	}
	return *o.EnableControlMapping
}

// GetEnableControlMappingOk returns a tuple with the EnableControlMapping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddAttributeMapperPluginRequest) GetEnableControlMappingOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableControlMapping) {
		return nil, false
	}
	return o.EnableControlMapping, true
}

// HasEnableControlMapping returns a boolean if a field has been set.
func (o *AddAttributeMapperPluginRequest) HasEnableControlMapping() bool {
	if o != nil && !IsNil(o.EnableControlMapping) {
		return true
	}

	return false
}

// SetEnableControlMapping gets a reference to the given bool and assigns it to the EnableControlMapping field.
func (o *AddAttributeMapperPluginRequest) SetEnableControlMapping(v bool) {
	o.EnableControlMapping = &v
}

// GetAlwaysMapResponses returns the AlwaysMapResponses field value if set, zero value otherwise.
func (o *AddAttributeMapperPluginRequest) GetAlwaysMapResponses() bool {
	if o == nil || IsNil(o.AlwaysMapResponses) {
		var ret bool
		return ret
	}
	return *o.AlwaysMapResponses
}

// GetAlwaysMapResponsesOk returns a tuple with the AlwaysMapResponses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddAttributeMapperPluginRequest) GetAlwaysMapResponsesOk() (*bool, bool) {
	if o == nil || IsNil(o.AlwaysMapResponses) {
		return nil, false
	}
	return o.AlwaysMapResponses, true
}

// HasAlwaysMapResponses returns a boolean if a field has been set.
func (o *AddAttributeMapperPluginRequest) HasAlwaysMapResponses() bool {
	if o != nil && !IsNil(o.AlwaysMapResponses) {
		return true
	}

	return false
}

// SetAlwaysMapResponses gets a reference to the given bool and assigns it to the AlwaysMapResponses field.
func (o *AddAttributeMapperPluginRequest) SetAlwaysMapResponses(v bool) {
	o.AlwaysMapResponses = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddAttributeMapperPluginRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddAttributeMapperPluginRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddAttributeMapperPluginRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddAttributeMapperPluginRequest) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *AddAttributeMapperPluginRequest) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AddAttributeMapperPluginRequest) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AddAttributeMapperPluginRequest) SetEnabled(v bool) {
	o.Enabled = v
}

// GetInvokeForInternalOperations returns the InvokeForInternalOperations field value if set, zero value otherwise.
func (o *AddAttributeMapperPluginRequest) GetInvokeForInternalOperations() bool {
	if o == nil || IsNil(o.InvokeForInternalOperations) {
		var ret bool
		return ret
	}
	return *o.InvokeForInternalOperations
}

// GetInvokeForInternalOperationsOk returns a tuple with the InvokeForInternalOperations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddAttributeMapperPluginRequest) GetInvokeForInternalOperationsOk() (*bool, bool) {
	if o == nil || IsNil(o.InvokeForInternalOperations) {
		return nil, false
	}
	return o.InvokeForInternalOperations, true
}

// HasInvokeForInternalOperations returns a boolean if a field has been set.
func (o *AddAttributeMapperPluginRequest) HasInvokeForInternalOperations() bool {
	if o != nil && !IsNil(o.InvokeForInternalOperations) {
		return true
	}

	return false
}

// SetInvokeForInternalOperations gets a reference to the given bool and assigns it to the InvokeForInternalOperations field.
func (o *AddAttributeMapperPluginRequest) SetInvokeForInternalOperations(v bool) {
	o.InvokeForInternalOperations = &v
}

func (o AddAttributeMapperPluginRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddAttributeMapperPluginRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pluginName"] = o.PluginName
	toSerialize["schemas"] = o.Schemas
	if !IsNil(o.PluginType) {
		toSerialize["pluginType"] = o.PluginType
	}
	toSerialize["sourceAttribute"] = o.SourceAttribute
	toSerialize["targetAttribute"] = o.TargetAttribute
	if !IsNil(o.EnableControlMapping) {
		toSerialize["enableControlMapping"] = o.EnableControlMapping
	}
	if !IsNil(o.AlwaysMapResponses) {
		toSerialize["alwaysMapResponses"] = o.AlwaysMapResponses
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["enabled"] = o.Enabled
	if !IsNil(o.InvokeForInternalOperations) {
		toSerialize["invokeForInternalOperations"] = o.InvokeForInternalOperations
	}
	return toSerialize, nil
}

type NullableAddAttributeMapperPluginRequest struct {
	value *AddAttributeMapperPluginRequest
	isSet bool
}

func (v NullableAddAttributeMapperPluginRequest) Get() *AddAttributeMapperPluginRequest {
	return v.value
}

func (v *NullableAddAttributeMapperPluginRequest) Set(val *AddAttributeMapperPluginRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddAttributeMapperPluginRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddAttributeMapperPluginRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddAttributeMapperPluginRequest(val *AddAttributeMapperPluginRequest) *NullableAddAttributeMapperPluginRequest {
	return &NullableAddAttributeMapperPluginRequest{value: val, isSet: true}
}

func (v NullableAddAttributeMapperPluginRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddAttributeMapperPluginRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
