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

// AddGroovyScriptedHttpOperationLogPublisherRequest struct for AddGroovyScriptedHttpOperationLogPublisherRequest
type AddGroovyScriptedHttpOperationLogPublisherRequest struct {
	// Name of the new Log Publisher
	PublisherName string                                                 `json:"publisherName"`
	Schemas       []EnumgroovyScriptedHttpOperationLogPublisherSchemaUrn `json:"schemas"`
	// The fully-qualified name of the Groovy class providing the logic for the Groovy Scripted HTTP Operation Log Publisher.
	ScriptClass string `json:"scriptClass"`
	// The set of arguments used to customize the behavior for the Scripted HTTP Operation Log Publisher. Each configuration property should be given in the form 'name=value'.
	ScriptArgument []string `json:"scriptArgument,omitempty"`
	// A description for this Log Publisher
	Description *string `json:"description,omitempty"`
	// Indicates whether the Log Publisher is enabled for use.
	Enabled              bool                                      `json:"enabled"`
	LoggingErrorBehavior *EnumlogPublisherLoggingErrorBehaviorProp `json:"loggingErrorBehavior,omitempty"`
}

// NewAddGroovyScriptedHttpOperationLogPublisherRequest instantiates a new AddGroovyScriptedHttpOperationLogPublisherRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddGroovyScriptedHttpOperationLogPublisherRequest(publisherName string, schemas []EnumgroovyScriptedHttpOperationLogPublisherSchemaUrn, scriptClass string, enabled bool) *AddGroovyScriptedHttpOperationLogPublisherRequest {
	this := AddGroovyScriptedHttpOperationLogPublisherRequest{}
	this.PublisherName = publisherName
	this.Schemas = schemas
	this.ScriptClass = scriptClass
	this.Enabled = enabled
	return &this
}

// NewAddGroovyScriptedHttpOperationLogPublisherRequestWithDefaults instantiates a new AddGroovyScriptedHttpOperationLogPublisherRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddGroovyScriptedHttpOperationLogPublisherRequestWithDefaults() *AddGroovyScriptedHttpOperationLogPublisherRequest {
	this := AddGroovyScriptedHttpOperationLogPublisherRequest{}
	return &this
}

// GetPublisherName returns the PublisherName field value
func (o *AddGroovyScriptedHttpOperationLogPublisherRequest) GetPublisherName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PublisherName
}

// GetPublisherNameOk returns a tuple with the PublisherName field value
// and a boolean to check if the value has been set.
func (o *AddGroovyScriptedHttpOperationLogPublisherRequest) GetPublisherNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PublisherName, true
}

// SetPublisherName sets field value
func (o *AddGroovyScriptedHttpOperationLogPublisherRequest) SetPublisherName(v string) {
	o.PublisherName = v
}

// GetSchemas returns the Schemas field value
func (o *AddGroovyScriptedHttpOperationLogPublisherRequest) GetSchemas() []EnumgroovyScriptedHttpOperationLogPublisherSchemaUrn {
	if o == nil {
		var ret []EnumgroovyScriptedHttpOperationLogPublisherSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddGroovyScriptedHttpOperationLogPublisherRequest) GetSchemasOk() ([]EnumgroovyScriptedHttpOperationLogPublisherSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddGroovyScriptedHttpOperationLogPublisherRequest) SetSchemas(v []EnumgroovyScriptedHttpOperationLogPublisherSchemaUrn) {
	o.Schemas = v
}

// GetScriptClass returns the ScriptClass field value
func (o *AddGroovyScriptedHttpOperationLogPublisherRequest) GetScriptClass() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ScriptClass
}

// GetScriptClassOk returns a tuple with the ScriptClass field value
// and a boolean to check if the value has been set.
func (o *AddGroovyScriptedHttpOperationLogPublisherRequest) GetScriptClassOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScriptClass, true
}

// SetScriptClass sets field value
func (o *AddGroovyScriptedHttpOperationLogPublisherRequest) SetScriptClass(v string) {
	o.ScriptClass = v
}

// GetScriptArgument returns the ScriptArgument field value if set, zero value otherwise.
func (o *AddGroovyScriptedHttpOperationLogPublisherRequest) GetScriptArgument() []string {
	if o == nil || isNil(o.ScriptArgument) {
		var ret []string
		return ret
	}
	return o.ScriptArgument
}

// GetScriptArgumentOk returns a tuple with the ScriptArgument field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddGroovyScriptedHttpOperationLogPublisherRequest) GetScriptArgumentOk() ([]string, bool) {
	if o == nil || isNil(o.ScriptArgument) {
		return nil, false
	}
	return o.ScriptArgument, true
}

// HasScriptArgument returns a boolean if a field has been set.
func (o *AddGroovyScriptedHttpOperationLogPublisherRequest) HasScriptArgument() bool {
	if o != nil && !isNil(o.ScriptArgument) {
		return true
	}

	return false
}

// SetScriptArgument gets a reference to the given []string and assigns it to the ScriptArgument field.
func (o *AddGroovyScriptedHttpOperationLogPublisherRequest) SetScriptArgument(v []string) {
	o.ScriptArgument = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddGroovyScriptedHttpOperationLogPublisherRequest) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddGroovyScriptedHttpOperationLogPublisherRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddGroovyScriptedHttpOperationLogPublisherRequest) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddGroovyScriptedHttpOperationLogPublisherRequest) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *AddGroovyScriptedHttpOperationLogPublisherRequest) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AddGroovyScriptedHttpOperationLogPublisherRequest) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AddGroovyScriptedHttpOperationLogPublisherRequest) SetEnabled(v bool) {
	o.Enabled = v
}

// GetLoggingErrorBehavior returns the LoggingErrorBehavior field value if set, zero value otherwise.
func (o *AddGroovyScriptedHttpOperationLogPublisherRequest) GetLoggingErrorBehavior() EnumlogPublisherLoggingErrorBehaviorProp {
	if o == nil || isNil(o.LoggingErrorBehavior) {
		var ret EnumlogPublisherLoggingErrorBehaviorProp
		return ret
	}
	return *o.LoggingErrorBehavior
}

// GetLoggingErrorBehaviorOk returns a tuple with the LoggingErrorBehavior field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddGroovyScriptedHttpOperationLogPublisherRequest) GetLoggingErrorBehaviorOk() (*EnumlogPublisherLoggingErrorBehaviorProp, bool) {
	if o == nil || isNil(o.LoggingErrorBehavior) {
		return nil, false
	}
	return o.LoggingErrorBehavior, true
}

// HasLoggingErrorBehavior returns a boolean if a field has been set.
func (o *AddGroovyScriptedHttpOperationLogPublisherRequest) HasLoggingErrorBehavior() bool {
	if o != nil && !isNil(o.LoggingErrorBehavior) {
		return true
	}

	return false
}

// SetLoggingErrorBehavior gets a reference to the given EnumlogPublisherLoggingErrorBehaviorProp and assigns it to the LoggingErrorBehavior field.
func (o *AddGroovyScriptedHttpOperationLogPublisherRequest) SetLoggingErrorBehavior(v EnumlogPublisherLoggingErrorBehaviorProp) {
	o.LoggingErrorBehavior = &v
}

func (o AddGroovyScriptedHttpOperationLogPublisherRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["publisherName"] = o.PublisherName
	}
	if true {
		toSerialize["schemas"] = o.Schemas
	}
	if true {
		toSerialize["scriptClass"] = o.ScriptClass
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
	if !isNil(o.LoggingErrorBehavior) {
		toSerialize["loggingErrorBehavior"] = o.LoggingErrorBehavior
	}
	return json.Marshal(toSerialize)
}

type NullableAddGroovyScriptedHttpOperationLogPublisherRequest struct {
	value *AddGroovyScriptedHttpOperationLogPublisherRequest
	isSet bool
}

func (v NullableAddGroovyScriptedHttpOperationLogPublisherRequest) Get() *AddGroovyScriptedHttpOperationLogPublisherRequest {
	return v.value
}

func (v *NullableAddGroovyScriptedHttpOperationLogPublisherRequest) Set(val *AddGroovyScriptedHttpOperationLogPublisherRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddGroovyScriptedHttpOperationLogPublisherRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddGroovyScriptedHttpOperationLogPublisherRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddGroovyScriptedHttpOperationLogPublisherRequest(val *AddGroovyScriptedHttpOperationLogPublisherRequest) *NullableAddGroovyScriptedHttpOperationLogPublisherRequest {
	return &NullableAddGroovyScriptedHttpOperationLogPublisherRequest{value: val, isSet: true}
}

func (v NullableAddGroovyScriptedHttpOperationLogPublisherRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddGroovyScriptedHttpOperationLogPublisherRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
