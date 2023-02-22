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

// AddGroovyScriptedPluginRequest struct for AddGroovyScriptedPluginRequest
type AddGroovyScriptedPluginRequest struct {
	// Name of the new Plugin
	PluginName string                              `json:"pluginName"`
	Schemas    []EnumgroovyScriptedPluginSchemaUrn `json:"schemas"`
	// The fully-qualified name of the Groovy class providing the logic for the Groovy Scripted Plugin.
	ScriptClass string `json:"scriptClass"`
	// Specifies a set of request criteria that may be used to indicate that this Groovy Scripted Plugin should only be invoked for operations in which the associated request matches this criteria.
	RequestCriteria *string `json:"requestCriteria,omitempty"`
	// The set of arguments used to customize the behavior for the Scripted Plugin. Each configuration property should be given in the form 'name=value'.
	ScriptArgument []string `json:"scriptArgument,omitempty"`
	// A description for this Plugin
	Description *string `json:"description,omitempty"`
	// Indicates whether the plug-in is enabled for use.
	Enabled    bool                       `json:"enabled"`
	PluginType []EnumpluginPluginTypeProp `json:"pluginType"`
	// Indicates whether the plug-in should be invoked for internal operations.
	InvokeForInternalOperations *bool `json:"invokeForInternalOperations,omitempty"`
}

// NewAddGroovyScriptedPluginRequest instantiates a new AddGroovyScriptedPluginRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddGroovyScriptedPluginRequest(pluginName string, schemas []EnumgroovyScriptedPluginSchemaUrn, scriptClass string, enabled bool, pluginType []EnumpluginPluginTypeProp) *AddGroovyScriptedPluginRequest {
	this := AddGroovyScriptedPluginRequest{}
	this.PluginName = pluginName
	this.Schemas = schemas
	this.ScriptClass = scriptClass
	this.Enabled = enabled
	this.PluginType = pluginType
	return &this
}

// NewAddGroovyScriptedPluginRequestWithDefaults instantiates a new AddGroovyScriptedPluginRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddGroovyScriptedPluginRequestWithDefaults() *AddGroovyScriptedPluginRequest {
	this := AddGroovyScriptedPluginRequest{}
	return &this
}

// GetPluginName returns the PluginName field value
func (o *AddGroovyScriptedPluginRequest) GetPluginName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PluginName
}

// GetPluginNameOk returns a tuple with the PluginName field value
// and a boolean to check if the value has been set.
func (o *AddGroovyScriptedPluginRequest) GetPluginNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PluginName, true
}

// SetPluginName sets field value
func (o *AddGroovyScriptedPluginRequest) SetPluginName(v string) {
	o.PluginName = v
}

// GetSchemas returns the Schemas field value
func (o *AddGroovyScriptedPluginRequest) GetSchemas() []EnumgroovyScriptedPluginSchemaUrn {
	if o == nil {
		var ret []EnumgroovyScriptedPluginSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddGroovyScriptedPluginRequest) GetSchemasOk() ([]EnumgroovyScriptedPluginSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddGroovyScriptedPluginRequest) SetSchemas(v []EnumgroovyScriptedPluginSchemaUrn) {
	o.Schemas = v
}

// GetScriptClass returns the ScriptClass field value
func (o *AddGroovyScriptedPluginRequest) GetScriptClass() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ScriptClass
}

// GetScriptClassOk returns a tuple with the ScriptClass field value
// and a boolean to check if the value has been set.
func (o *AddGroovyScriptedPluginRequest) GetScriptClassOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScriptClass, true
}

// SetScriptClass sets field value
func (o *AddGroovyScriptedPluginRequest) SetScriptClass(v string) {
	o.ScriptClass = v
}

// GetRequestCriteria returns the RequestCriteria field value if set, zero value otherwise.
func (o *AddGroovyScriptedPluginRequest) GetRequestCriteria() string {
	if o == nil || isNil(o.RequestCriteria) {
		var ret string
		return ret
	}
	return *o.RequestCriteria
}

// GetRequestCriteriaOk returns a tuple with the RequestCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddGroovyScriptedPluginRequest) GetRequestCriteriaOk() (*string, bool) {
	if o == nil || isNil(o.RequestCriteria) {
		return nil, false
	}
	return o.RequestCriteria, true
}

// HasRequestCriteria returns a boolean if a field has been set.
func (o *AddGroovyScriptedPluginRequest) HasRequestCriteria() bool {
	if o != nil && !isNil(o.RequestCriteria) {
		return true
	}

	return false
}

// SetRequestCriteria gets a reference to the given string and assigns it to the RequestCriteria field.
func (o *AddGroovyScriptedPluginRequest) SetRequestCriteria(v string) {
	o.RequestCriteria = &v
}

// GetScriptArgument returns the ScriptArgument field value if set, zero value otherwise.
func (o *AddGroovyScriptedPluginRequest) GetScriptArgument() []string {
	if o == nil || isNil(o.ScriptArgument) {
		var ret []string
		return ret
	}
	return o.ScriptArgument
}

// GetScriptArgumentOk returns a tuple with the ScriptArgument field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddGroovyScriptedPluginRequest) GetScriptArgumentOk() ([]string, bool) {
	if o == nil || isNil(o.ScriptArgument) {
		return nil, false
	}
	return o.ScriptArgument, true
}

// HasScriptArgument returns a boolean if a field has been set.
func (o *AddGroovyScriptedPluginRequest) HasScriptArgument() bool {
	if o != nil && !isNil(o.ScriptArgument) {
		return true
	}

	return false
}

// SetScriptArgument gets a reference to the given []string and assigns it to the ScriptArgument field.
func (o *AddGroovyScriptedPluginRequest) SetScriptArgument(v []string) {
	o.ScriptArgument = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddGroovyScriptedPluginRequest) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddGroovyScriptedPluginRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddGroovyScriptedPluginRequest) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddGroovyScriptedPluginRequest) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *AddGroovyScriptedPluginRequest) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AddGroovyScriptedPluginRequest) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AddGroovyScriptedPluginRequest) SetEnabled(v bool) {
	o.Enabled = v
}

// GetPluginType returns the PluginType field value
func (o *AddGroovyScriptedPluginRequest) GetPluginType() []EnumpluginPluginTypeProp {
	if o == nil {
		var ret []EnumpluginPluginTypeProp
		return ret
	}

	return o.PluginType
}

// GetPluginTypeOk returns a tuple with the PluginType field value
// and a boolean to check if the value has been set.
func (o *AddGroovyScriptedPluginRequest) GetPluginTypeOk() ([]EnumpluginPluginTypeProp, bool) {
	if o == nil {
		return nil, false
	}
	return o.PluginType, true
}

// SetPluginType sets field value
func (o *AddGroovyScriptedPluginRequest) SetPluginType(v []EnumpluginPluginTypeProp) {
	o.PluginType = v
}

// GetInvokeForInternalOperations returns the InvokeForInternalOperations field value if set, zero value otherwise.
func (o *AddGroovyScriptedPluginRequest) GetInvokeForInternalOperations() bool {
	if o == nil || isNil(o.InvokeForInternalOperations) {
		var ret bool
		return ret
	}
	return *o.InvokeForInternalOperations
}

// GetInvokeForInternalOperationsOk returns a tuple with the InvokeForInternalOperations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddGroovyScriptedPluginRequest) GetInvokeForInternalOperationsOk() (*bool, bool) {
	if o == nil || isNil(o.InvokeForInternalOperations) {
		return nil, false
	}
	return o.InvokeForInternalOperations, true
}

// HasInvokeForInternalOperations returns a boolean if a field has been set.
func (o *AddGroovyScriptedPluginRequest) HasInvokeForInternalOperations() bool {
	if o != nil && !isNil(o.InvokeForInternalOperations) {
		return true
	}

	return false
}

// SetInvokeForInternalOperations gets a reference to the given bool and assigns it to the InvokeForInternalOperations field.
func (o *AddGroovyScriptedPluginRequest) SetInvokeForInternalOperations(v bool) {
	o.InvokeForInternalOperations = &v
}

func (o AddGroovyScriptedPluginRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pluginName"] = o.PluginName
	}
	if true {
		toSerialize["schemas"] = o.Schemas
	}
	if true {
		toSerialize["scriptClass"] = o.ScriptClass
	}
	if !isNil(o.RequestCriteria) {
		toSerialize["requestCriteria"] = o.RequestCriteria
	}
	if !isNil(o.ScriptArgument) {
		toSerialize["scriptArgument"] = o.ScriptArgument
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	if true {
		toSerialize["pluginType"] = o.PluginType
	}
	if !isNil(o.InvokeForInternalOperations) {
		toSerialize["invokeForInternalOperations"] = o.InvokeForInternalOperations
	}
	return json.Marshal(toSerialize)
}

type NullableAddGroovyScriptedPluginRequest struct {
	value *AddGroovyScriptedPluginRequest
	isSet bool
}

func (v NullableAddGroovyScriptedPluginRequest) Get() *AddGroovyScriptedPluginRequest {
	return v.value
}

func (v *NullableAddGroovyScriptedPluginRequest) Set(val *AddGroovyScriptedPluginRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddGroovyScriptedPluginRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddGroovyScriptedPluginRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddGroovyScriptedPluginRequest(val *AddGroovyScriptedPluginRequest) *NullableAddGroovyScriptedPluginRequest {
	return &NullableAddGroovyScriptedPluginRequest{value: val, isSet: true}
}

func (v NullableAddGroovyScriptedPluginRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddGroovyScriptedPluginRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}