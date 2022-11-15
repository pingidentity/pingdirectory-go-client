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

// GroovyScriptedAlertHandlerShared struct for GroovyScriptedAlertHandlerShared
type GroovyScriptedAlertHandlerShared struct {
	Schemas []EnumgroovyScriptedAlertHandlerSchemaUrn `json:"schemas"`
	// The fully-qualified name of the Groovy class providing the logic for the Groovy Scripted Alert Handler.
	ScriptClass string `json:"scriptClass"`
	ScriptArgument []string `json:"scriptArgument,omitempty"`
	// A description for this Alert Handler
	Description *string `json:"description,omitempty"`
	// Indicates whether the Alert Handler is enabled.
	Enabled bool `json:"enabled"`
	// Indicates whether the server should attempt to invoke this Alert Handler in a background thread so that any potentially-expensive processing (e.g., performing network communication to deliver the alert notification) will not delay whatever processing the server was performing when the alert was generated.
	Asynchronous *bool `json:"asynchronous,omitempty"`
	EnabledAlertSeverity []EnumalertHandlerEnabledAlertSeverityProp `json:"enabledAlertSeverity,omitempty"`
	EnabledAlertType []EnumalertHandlerEnabledAlertTypeProp `json:"enabledAlertType,omitempty"`
	DisabledAlertType []EnumalertHandlerDisabledAlertTypeProp `json:"disabledAlertType,omitempty"`
}

// NewGroovyScriptedAlertHandlerShared instantiates a new GroovyScriptedAlertHandlerShared object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroovyScriptedAlertHandlerShared(schemas []EnumgroovyScriptedAlertHandlerSchemaUrn, scriptClass string, enabled bool) *GroovyScriptedAlertHandlerShared {
	this := GroovyScriptedAlertHandlerShared{}
	this.Schemas = schemas
	this.ScriptClass = scriptClass
	this.Enabled = enabled
	return &this
}

// NewGroovyScriptedAlertHandlerSharedWithDefaults instantiates a new GroovyScriptedAlertHandlerShared object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroovyScriptedAlertHandlerSharedWithDefaults() *GroovyScriptedAlertHandlerShared {
	this := GroovyScriptedAlertHandlerShared{}
	return &this
}

// GetSchemas returns the Schemas field value
func (o *GroovyScriptedAlertHandlerShared) GetSchemas() []EnumgroovyScriptedAlertHandlerSchemaUrn {
	if o == nil {
		var ret []EnumgroovyScriptedAlertHandlerSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *GroovyScriptedAlertHandlerShared) GetSchemasOk() ([]EnumgroovyScriptedAlertHandlerSchemaUrn, bool) {
	if o == nil {
    return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *GroovyScriptedAlertHandlerShared) SetSchemas(v []EnumgroovyScriptedAlertHandlerSchemaUrn) {
	o.Schemas = v
}

// GetScriptClass returns the ScriptClass field value
func (o *GroovyScriptedAlertHandlerShared) GetScriptClass() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ScriptClass
}

// GetScriptClassOk returns a tuple with the ScriptClass field value
// and a boolean to check if the value has been set.
func (o *GroovyScriptedAlertHandlerShared) GetScriptClassOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ScriptClass, true
}

// SetScriptClass sets field value
func (o *GroovyScriptedAlertHandlerShared) SetScriptClass(v string) {
	o.ScriptClass = v
}

// GetScriptArgument returns the ScriptArgument field value if set, zero value otherwise.
func (o *GroovyScriptedAlertHandlerShared) GetScriptArgument() []string {
	if o == nil || isNil(o.ScriptArgument) {
		var ret []string
		return ret
	}
	return o.ScriptArgument
}

// GetScriptArgumentOk returns a tuple with the ScriptArgument field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroovyScriptedAlertHandlerShared) GetScriptArgumentOk() ([]string, bool) {
	if o == nil || isNil(o.ScriptArgument) {
    return nil, false
	}
	return o.ScriptArgument, true
}

// HasScriptArgument returns a boolean if a field has been set.
func (o *GroovyScriptedAlertHandlerShared) HasScriptArgument() bool {
	if o != nil && !isNil(o.ScriptArgument) {
		return true
	}

	return false
}

// SetScriptArgument gets a reference to the given []string and assigns it to the ScriptArgument field.
func (o *GroovyScriptedAlertHandlerShared) SetScriptArgument(v []string) {
	o.ScriptArgument = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *GroovyScriptedAlertHandlerShared) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroovyScriptedAlertHandlerShared) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *GroovyScriptedAlertHandlerShared) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *GroovyScriptedAlertHandlerShared) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *GroovyScriptedAlertHandlerShared) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *GroovyScriptedAlertHandlerShared) GetEnabledOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *GroovyScriptedAlertHandlerShared) SetEnabled(v bool) {
	o.Enabled = v
}

// GetAsynchronous returns the Asynchronous field value if set, zero value otherwise.
func (o *GroovyScriptedAlertHandlerShared) GetAsynchronous() bool {
	if o == nil || isNil(o.Asynchronous) {
		var ret bool
		return ret
	}
	return *o.Asynchronous
}

// GetAsynchronousOk returns a tuple with the Asynchronous field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroovyScriptedAlertHandlerShared) GetAsynchronousOk() (*bool, bool) {
	if o == nil || isNil(o.Asynchronous) {
    return nil, false
	}
	return o.Asynchronous, true
}

// HasAsynchronous returns a boolean if a field has been set.
func (o *GroovyScriptedAlertHandlerShared) HasAsynchronous() bool {
	if o != nil && !isNil(o.Asynchronous) {
		return true
	}

	return false
}

// SetAsynchronous gets a reference to the given bool and assigns it to the Asynchronous field.
func (o *GroovyScriptedAlertHandlerShared) SetAsynchronous(v bool) {
	o.Asynchronous = &v
}

// GetEnabledAlertSeverity returns the EnabledAlertSeverity field value if set, zero value otherwise.
func (o *GroovyScriptedAlertHandlerShared) GetEnabledAlertSeverity() []EnumalertHandlerEnabledAlertSeverityProp {
	if o == nil || isNil(o.EnabledAlertSeverity) {
		var ret []EnumalertHandlerEnabledAlertSeverityProp
		return ret
	}
	return o.EnabledAlertSeverity
}

// GetEnabledAlertSeverityOk returns a tuple with the EnabledAlertSeverity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroovyScriptedAlertHandlerShared) GetEnabledAlertSeverityOk() ([]EnumalertHandlerEnabledAlertSeverityProp, bool) {
	if o == nil || isNil(o.EnabledAlertSeverity) {
    return nil, false
	}
	return o.EnabledAlertSeverity, true
}

// HasEnabledAlertSeverity returns a boolean if a field has been set.
func (o *GroovyScriptedAlertHandlerShared) HasEnabledAlertSeverity() bool {
	if o != nil && !isNil(o.EnabledAlertSeverity) {
		return true
	}

	return false
}

// SetEnabledAlertSeverity gets a reference to the given []EnumalertHandlerEnabledAlertSeverityProp and assigns it to the EnabledAlertSeverity field.
func (o *GroovyScriptedAlertHandlerShared) SetEnabledAlertSeverity(v []EnumalertHandlerEnabledAlertSeverityProp) {
	o.EnabledAlertSeverity = v
}

// GetEnabledAlertType returns the EnabledAlertType field value if set, zero value otherwise.
func (o *GroovyScriptedAlertHandlerShared) GetEnabledAlertType() []EnumalertHandlerEnabledAlertTypeProp {
	if o == nil || isNil(o.EnabledAlertType) {
		var ret []EnumalertHandlerEnabledAlertTypeProp
		return ret
	}
	return o.EnabledAlertType
}

// GetEnabledAlertTypeOk returns a tuple with the EnabledAlertType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroovyScriptedAlertHandlerShared) GetEnabledAlertTypeOk() ([]EnumalertHandlerEnabledAlertTypeProp, bool) {
	if o == nil || isNil(o.EnabledAlertType) {
    return nil, false
	}
	return o.EnabledAlertType, true
}

// HasEnabledAlertType returns a boolean if a field has been set.
func (o *GroovyScriptedAlertHandlerShared) HasEnabledAlertType() bool {
	if o != nil && !isNil(o.EnabledAlertType) {
		return true
	}

	return false
}

// SetEnabledAlertType gets a reference to the given []EnumalertHandlerEnabledAlertTypeProp and assigns it to the EnabledAlertType field.
func (o *GroovyScriptedAlertHandlerShared) SetEnabledAlertType(v []EnumalertHandlerEnabledAlertTypeProp) {
	o.EnabledAlertType = v
}

// GetDisabledAlertType returns the DisabledAlertType field value if set, zero value otherwise.
func (o *GroovyScriptedAlertHandlerShared) GetDisabledAlertType() []EnumalertHandlerDisabledAlertTypeProp {
	if o == nil || isNil(o.DisabledAlertType) {
		var ret []EnumalertHandlerDisabledAlertTypeProp
		return ret
	}
	return o.DisabledAlertType
}

// GetDisabledAlertTypeOk returns a tuple with the DisabledAlertType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroovyScriptedAlertHandlerShared) GetDisabledAlertTypeOk() ([]EnumalertHandlerDisabledAlertTypeProp, bool) {
	if o == nil || isNil(o.DisabledAlertType) {
    return nil, false
	}
	return o.DisabledAlertType, true
}

// HasDisabledAlertType returns a boolean if a field has been set.
func (o *GroovyScriptedAlertHandlerShared) HasDisabledAlertType() bool {
	if o != nil && !isNil(o.DisabledAlertType) {
		return true
	}

	return false
}

// SetDisabledAlertType gets a reference to the given []EnumalertHandlerDisabledAlertTypeProp and assigns it to the DisabledAlertType field.
func (o *GroovyScriptedAlertHandlerShared) SetDisabledAlertType(v []EnumalertHandlerDisabledAlertTypeProp) {
	o.DisabledAlertType = v
}

func (o GroovyScriptedAlertHandlerShared) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
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
	if !isNil(o.Asynchronous) {
		toSerialize["asynchronous"] = o.Asynchronous
	}
	if !isNil(o.EnabledAlertSeverity) {
		toSerialize["enabledAlertSeverity"] = o.EnabledAlertSeverity
	}
	if !isNil(o.EnabledAlertType) {
		toSerialize["enabledAlertType"] = o.EnabledAlertType
	}
	if !isNil(o.DisabledAlertType) {
		toSerialize["disabledAlertType"] = o.DisabledAlertType
	}
	return json.Marshal(toSerialize)
}

type NullableGroovyScriptedAlertHandlerShared struct {
	value *GroovyScriptedAlertHandlerShared
	isSet bool
}

func (v NullableGroovyScriptedAlertHandlerShared) Get() *GroovyScriptedAlertHandlerShared {
	return v.value
}

func (v *NullableGroovyScriptedAlertHandlerShared) Set(val *GroovyScriptedAlertHandlerShared) {
	v.value = val
	v.isSet = true
}

func (v NullableGroovyScriptedAlertHandlerShared) IsSet() bool {
	return v.isSet
}

func (v *NullableGroovyScriptedAlertHandlerShared) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroovyScriptedAlertHandlerShared(val *GroovyScriptedAlertHandlerShared) *NullableGroovyScriptedAlertHandlerShared {
	return &NullableGroovyScriptedAlertHandlerShared{value: val, isSet: true}
}

func (v NullableGroovyScriptedAlertHandlerShared) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroovyScriptedAlertHandlerShared) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


