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

// AddDnMapperPluginRequest struct for AddDnMapperPluginRequest
type AddDnMapperPluginRequest struct {
	// Name of the new Plugin
	PluginName string                        `json:"pluginName"`
	Schemas    []EnumdnMapperPluginSchemaUrn `json:"schemas"`
	PluginType []EnumpluginPluginTypeProp    `json:"pluginType"`
	// Specifies the source DN that may appear in client requests which should be remapped to the target DN. Note that the source DN must not be equal to the target DN.
	SourceDN string `json:"sourceDN"`
	// Specifies the DN to which the source DN should be mapped. Note that the target DN must not be equal to the source DN.
	TargetDN string `json:"targetDN"`
	// Indicates whether DN mapping should be applied to the values of attributes with appropriate syntaxes.
	EnableAttributeMapping bool `json:"enableAttributeMapping"`
	// Specifies a set of specific attributes for which DN mapping should be applied. This will only be applicable if the enable-attribute-mapping property has a value of \"true\". Any attributes listed must be defined in the server schema with either the distinguished name syntax or the name and optional UID syntax.
	MapAttribute []string `json:"mapAttribute,omitempty"`
	// Indicates whether DN mapping should be applied to DNs that may be present in specific controls. DN mapping will only be applied for control types which are specifically supported by the DN mapper plugin.
	EnableControlMapping bool `json:"enableControlMapping"`
	// Indicates whether DNs in response messages containing the target DN should always be remapped back to the source DN. If this is \"false\", then mapping will be performed for a response message only if one or more elements of the associated request are mapped. Otherwise, the mapping will be performed for all responses regardless of whether the mapping was applied to the request.
	AlwaysMapResponses bool `json:"alwaysMapResponses"`
	// A description for this Plugin
	Description *string `json:"description,omitempty"`
	// Indicates whether the plug-in is enabled for use.
	Enabled bool `json:"enabled"`
	// Indicates whether the plug-in should be invoked for internal operations.
	InvokeForInternalOperations *bool `json:"invokeForInternalOperations,omitempty"`
}

// NewAddDnMapperPluginRequest instantiates a new AddDnMapperPluginRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddDnMapperPluginRequest(pluginName string, schemas []EnumdnMapperPluginSchemaUrn, pluginType []EnumpluginPluginTypeProp, sourceDN string, targetDN string, enableAttributeMapping bool, enableControlMapping bool, alwaysMapResponses bool, enabled bool) *AddDnMapperPluginRequest {
	this := AddDnMapperPluginRequest{}
	this.PluginName = pluginName
	this.Schemas = schemas
	this.PluginType = pluginType
	this.SourceDN = sourceDN
	this.TargetDN = targetDN
	this.EnableAttributeMapping = enableAttributeMapping
	this.EnableControlMapping = enableControlMapping
	this.AlwaysMapResponses = alwaysMapResponses
	this.Enabled = enabled
	return &this
}

// NewAddDnMapperPluginRequestWithDefaults instantiates a new AddDnMapperPluginRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddDnMapperPluginRequestWithDefaults() *AddDnMapperPluginRequest {
	this := AddDnMapperPluginRequest{}
	return &this
}

// GetPluginName returns the PluginName field value
func (o *AddDnMapperPluginRequest) GetPluginName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PluginName
}

// GetPluginNameOk returns a tuple with the PluginName field value
// and a boolean to check if the value has been set.
func (o *AddDnMapperPluginRequest) GetPluginNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PluginName, true
}

// SetPluginName sets field value
func (o *AddDnMapperPluginRequest) SetPluginName(v string) {
	o.PluginName = v
}

// GetSchemas returns the Schemas field value
func (o *AddDnMapperPluginRequest) GetSchemas() []EnumdnMapperPluginSchemaUrn {
	if o == nil {
		var ret []EnumdnMapperPluginSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddDnMapperPluginRequest) GetSchemasOk() ([]EnumdnMapperPluginSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddDnMapperPluginRequest) SetSchemas(v []EnumdnMapperPluginSchemaUrn) {
	o.Schemas = v
}

// GetPluginType returns the PluginType field value
func (o *AddDnMapperPluginRequest) GetPluginType() []EnumpluginPluginTypeProp {
	if o == nil {
		var ret []EnumpluginPluginTypeProp
		return ret
	}

	return o.PluginType
}

// GetPluginTypeOk returns a tuple with the PluginType field value
// and a boolean to check if the value has been set.
func (o *AddDnMapperPluginRequest) GetPluginTypeOk() ([]EnumpluginPluginTypeProp, bool) {
	if o == nil {
		return nil, false
	}
	return o.PluginType, true
}

// SetPluginType sets field value
func (o *AddDnMapperPluginRequest) SetPluginType(v []EnumpluginPluginTypeProp) {
	o.PluginType = v
}

// GetSourceDN returns the SourceDN field value
func (o *AddDnMapperPluginRequest) GetSourceDN() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceDN
}

// GetSourceDNOk returns a tuple with the SourceDN field value
// and a boolean to check if the value has been set.
func (o *AddDnMapperPluginRequest) GetSourceDNOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceDN, true
}

// SetSourceDN sets field value
func (o *AddDnMapperPluginRequest) SetSourceDN(v string) {
	o.SourceDN = v
}

// GetTargetDN returns the TargetDN field value
func (o *AddDnMapperPluginRequest) GetTargetDN() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TargetDN
}

// GetTargetDNOk returns a tuple with the TargetDN field value
// and a boolean to check if the value has been set.
func (o *AddDnMapperPluginRequest) GetTargetDNOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetDN, true
}

// SetTargetDN sets field value
func (o *AddDnMapperPluginRequest) SetTargetDN(v string) {
	o.TargetDN = v
}

// GetEnableAttributeMapping returns the EnableAttributeMapping field value
func (o *AddDnMapperPluginRequest) GetEnableAttributeMapping() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.EnableAttributeMapping
}

// GetEnableAttributeMappingOk returns a tuple with the EnableAttributeMapping field value
// and a boolean to check if the value has been set.
func (o *AddDnMapperPluginRequest) GetEnableAttributeMappingOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnableAttributeMapping, true
}

// SetEnableAttributeMapping sets field value
func (o *AddDnMapperPluginRequest) SetEnableAttributeMapping(v bool) {
	o.EnableAttributeMapping = v
}

// GetMapAttribute returns the MapAttribute field value if set, zero value otherwise.
func (o *AddDnMapperPluginRequest) GetMapAttribute() []string {
	if o == nil || isNil(o.MapAttribute) {
		var ret []string
		return ret
	}
	return o.MapAttribute
}

// GetMapAttributeOk returns a tuple with the MapAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddDnMapperPluginRequest) GetMapAttributeOk() ([]string, bool) {
	if o == nil || isNil(o.MapAttribute) {
		return nil, false
	}
	return o.MapAttribute, true
}

// HasMapAttribute returns a boolean if a field has been set.
func (o *AddDnMapperPluginRequest) HasMapAttribute() bool {
	if o != nil && !isNil(o.MapAttribute) {
		return true
	}

	return false
}

// SetMapAttribute gets a reference to the given []string and assigns it to the MapAttribute field.
func (o *AddDnMapperPluginRequest) SetMapAttribute(v []string) {
	o.MapAttribute = v
}

// GetEnableControlMapping returns the EnableControlMapping field value
func (o *AddDnMapperPluginRequest) GetEnableControlMapping() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.EnableControlMapping
}

// GetEnableControlMappingOk returns a tuple with the EnableControlMapping field value
// and a boolean to check if the value has been set.
func (o *AddDnMapperPluginRequest) GetEnableControlMappingOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnableControlMapping, true
}

// SetEnableControlMapping sets field value
func (o *AddDnMapperPluginRequest) SetEnableControlMapping(v bool) {
	o.EnableControlMapping = v
}

// GetAlwaysMapResponses returns the AlwaysMapResponses field value
func (o *AddDnMapperPluginRequest) GetAlwaysMapResponses() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AlwaysMapResponses
}

// GetAlwaysMapResponsesOk returns a tuple with the AlwaysMapResponses field value
// and a boolean to check if the value has been set.
func (o *AddDnMapperPluginRequest) GetAlwaysMapResponsesOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AlwaysMapResponses, true
}

// SetAlwaysMapResponses sets field value
func (o *AddDnMapperPluginRequest) SetAlwaysMapResponses(v bool) {
	o.AlwaysMapResponses = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddDnMapperPluginRequest) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddDnMapperPluginRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddDnMapperPluginRequest) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddDnMapperPluginRequest) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *AddDnMapperPluginRequest) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AddDnMapperPluginRequest) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AddDnMapperPluginRequest) SetEnabled(v bool) {
	o.Enabled = v
}

// GetInvokeForInternalOperations returns the InvokeForInternalOperations field value if set, zero value otherwise.
func (o *AddDnMapperPluginRequest) GetInvokeForInternalOperations() bool {
	if o == nil || isNil(o.InvokeForInternalOperations) {
		var ret bool
		return ret
	}
	return *o.InvokeForInternalOperations
}

// GetInvokeForInternalOperationsOk returns a tuple with the InvokeForInternalOperations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddDnMapperPluginRequest) GetInvokeForInternalOperationsOk() (*bool, bool) {
	if o == nil || isNil(o.InvokeForInternalOperations) {
		return nil, false
	}
	return o.InvokeForInternalOperations, true
}

// HasInvokeForInternalOperations returns a boolean if a field has been set.
func (o *AddDnMapperPluginRequest) HasInvokeForInternalOperations() bool {
	if o != nil && !isNil(o.InvokeForInternalOperations) {
		return true
	}

	return false
}

// SetInvokeForInternalOperations gets a reference to the given bool and assigns it to the InvokeForInternalOperations field.
func (o *AddDnMapperPluginRequest) SetInvokeForInternalOperations(v bool) {
	o.InvokeForInternalOperations = &v
}

func (o AddDnMapperPluginRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pluginName"] = o.PluginName
	}
	if true {
		toSerialize["schemas"] = o.Schemas
	}
	if true {
		toSerialize["pluginType"] = o.PluginType
	}
	if true {
		toSerialize["sourceDN"] = o.SourceDN
	}
	if true {
		toSerialize["targetDN"] = o.TargetDN
	}
	if true {
		toSerialize["enableAttributeMapping"] = o.EnableAttributeMapping
	}
	if !isNil(o.MapAttribute) {
		toSerialize["mapAttribute"] = o.MapAttribute
	}
	if true {
		toSerialize["enableControlMapping"] = o.EnableControlMapping
	}
	if true {
		toSerialize["alwaysMapResponses"] = o.AlwaysMapResponses
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	if !isNil(o.InvokeForInternalOperations) {
		toSerialize["invokeForInternalOperations"] = o.InvokeForInternalOperations
	}
	return json.Marshal(toSerialize)
}

type NullableAddDnMapperPluginRequest struct {
	value *AddDnMapperPluginRequest
	isSet bool
}

func (v NullableAddDnMapperPluginRequest) Get() *AddDnMapperPluginRequest {
	return v.value
}

func (v *NullableAddDnMapperPluginRequest) Set(val *AddDnMapperPluginRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddDnMapperPluginRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddDnMapperPluginRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddDnMapperPluginRequest(val *AddDnMapperPluginRequest) *NullableAddDnMapperPluginRequest {
	return &NullableAddDnMapperPluginRequest{value: val, isSet: true}
}

func (v NullableAddDnMapperPluginRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddDnMapperPluginRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
