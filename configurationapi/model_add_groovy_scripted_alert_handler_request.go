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

// checks if the AddGroovyScriptedAlertHandlerRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddGroovyScriptedAlertHandlerRequest{}

// AddGroovyScriptedAlertHandlerRequest struct for AddGroovyScriptedAlertHandlerRequest
type AddGroovyScriptedAlertHandlerRequest struct {
	// Name of the new Alert Handler
	HandlerName string                                    `json:"handlerName"`
	Schemas     []EnumgroovyScriptedAlertHandlerSchemaUrn `json:"schemas"`
	// The fully-qualified name of the Groovy class providing the logic for the Groovy Scripted Alert Handler.
	ScriptClass string `json:"scriptClass"`
	// The set of arguments used to customize the behavior for the Scripted Alert Handler. Each configuration property should be given in the form 'name=value'.
	ScriptArgument []string `json:"scriptArgument,omitempty"`
	// A description for this Alert Handler
	Description *string `json:"description,omitempty"`
	// Indicates whether the Alert Handler is enabled.
	Enabled bool `json:"enabled"`
	// Indicates whether the server should attempt to invoke this Alert Handler in a background thread so that any potentially-expensive processing (e.g., performing network communication to deliver the alert notification) will not delay whatever processing the server was performing when the alert was generated.
	Asynchronous         *bool                                      `json:"asynchronous,omitempty"`
	EnabledAlertSeverity []EnumalertHandlerEnabledAlertSeverityProp `json:"enabledAlertSeverity,omitempty"`
	EnabledAlertType     []EnumalertHandlerEnabledAlertTypeProp     `json:"enabledAlertType,omitempty"`
	DisabledAlertType    []EnumalertHandlerDisabledAlertTypeProp    `json:"disabledAlertType,omitempty"`
}

// NewAddGroovyScriptedAlertHandlerRequest instantiates a new AddGroovyScriptedAlertHandlerRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddGroovyScriptedAlertHandlerRequest(handlerName string, schemas []EnumgroovyScriptedAlertHandlerSchemaUrn, scriptClass string, enabled bool) *AddGroovyScriptedAlertHandlerRequest {
	this := AddGroovyScriptedAlertHandlerRequest{}
	this.HandlerName = handlerName
	this.Schemas = schemas
	this.ScriptClass = scriptClass
	this.Enabled = enabled
	return &this
}

// NewAddGroovyScriptedAlertHandlerRequestWithDefaults instantiates a new AddGroovyScriptedAlertHandlerRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddGroovyScriptedAlertHandlerRequestWithDefaults() *AddGroovyScriptedAlertHandlerRequest {
	this := AddGroovyScriptedAlertHandlerRequest{}
	return &this
}

// GetHandlerName returns the HandlerName field value
func (o *AddGroovyScriptedAlertHandlerRequest) GetHandlerName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HandlerName
}

// GetHandlerNameOk returns a tuple with the HandlerName field value
// and a boolean to check if the value has been set.
func (o *AddGroovyScriptedAlertHandlerRequest) GetHandlerNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HandlerName, true
}

// SetHandlerName sets field value
func (o *AddGroovyScriptedAlertHandlerRequest) SetHandlerName(v string) {
	o.HandlerName = v
}

// GetSchemas returns the Schemas field value
func (o *AddGroovyScriptedAlertHandlerRequest) GetSchemas() []EnumgroovyScriptedAlertHandlerSchemaUrn {
	if o == nil {
		var ret []EnumgroovyScriptedAlertHandlerSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddGroovyScriptedAlertHandlerRequest) GetSchemasOk() ([]EnumgroovyScriptedAlertHandlerSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddGroovyScriptedAlertHandlerRequest) SetSchemas(v []EnumgroovyScriptedAlertHandlerSchemaUrn) {
	o.Schemas = v
}

// GetScriptClass returns the ScriptClass field value
func (o *AddGroovyScriptedAlertHandlerRequest) GetScriptClass() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ScriptClass
}

// GetScriptClassOk returns a tuple with the ScriptClass field value
// and a boolean to check if the value has been set.
func (o *AddGroovyScriptedAlertHandlerRequest) GetScriptClassOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScriptClass, true
}

// SetScriptClass sets field value
func (o *AddGroovyScriptedAlertHandlerRequest) SetScriptClass(v string) {
	o.ScriptClass = v
}

// GetScriptArgument returns the ScriptArgument field value if set, zero value otherwise.
func (o *AddGroovyScriptedAlertHandlerRequest) GetScriptArgument() []string {
	if o == nil || IsNil(o.ScriptArgument) {
		var ret []string
		return ret
	}
	return o.ScriptArgument
}

// GetScriptArgumentOk returns a tuple with the ScriptArgument field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddGroovyScriptedAlertHandlerRequest) GetScriptArgumentOk() ([]string, bool) {
	if o == nil || IsNil(o.ScriptArgument) {
		return nil, false
	}
	return o.ScriptArgument, true
}

// HasScriptArgument returns a boolean if a field has been set.
func (o *AddGroovyScriptedAlertHandlerRequest) HasScriptArgument() bool {
	if o != nil && !IsNil(o.ScriptArgument) {
		return true
	}

	return false
}

// SetScriptArgument gets a reference to the given []string and assigns it to the ScriptArgument field.
func (o *AddGroovyScriptedAlertHandlerRequest) SetScriptArgument(v []string) {
	o.ScriptArgument = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddGroovyScriptedAlertHandlerRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddGroovyScriptedAlertHandlerRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddGroovyScriptedAlertHandlerRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddGroovyScriptedAlertHandlerRequest) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *AddGroovyScriptedAlertHandlerRequest) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AddGroovyScriptedAlertHandlerRequest) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AddGroovyScriptedAlertHandlerRequest) SetEnabled(v bool) {
	o.Enabled = v
}

// GetAsynchronous returns the Asynchronous field value if set, zero value otherwise.
func (o *AddGroovyScriptedAlertHandlerRequest) GetAsynchronous() bool {
	if o == nil || IsNil(o.Asynchronous) {
		var ret bool
		return ret
	}
	return *o.Asynchronous
}

// GetAsynchronousOk returns a tuple with the Asynchronous field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddGroovyScriptedAlertHandlerRequest) GetAsynchronousOk() (*bool, bool) {
	if o == nil || IsNil(o.Asynchronous) {
		return nil, false
	}
	return o.Asynchronous, true
}

// HasAsynchronous returns a boolean if a field has been set.
func (o *AddGroovyScriptedAlertHandlerRequest) HasAsynchronous() bool {
	if o != nil && !IsNil(o.Asynchronous) {
		return true
	}

	return false
}

// SetAsynchronous gets a reference to the given bool and assigns it to the Asynchronous field.
func (o *AddGroovyScriptedAlertHandlerRequest) SetAsynchronous(v bool) {
	o.Asynchronous = &v
}

// GetEnabledAlertSeverity returns the EnabledAlertSeverity field value if set, zero value otherwise.
func (o *AddGroovyScriptedAlertHandlerRequest) GetEnabledAlertSeverity() []EnumalertHandlerEnabledAlertSeverityProp {
	if o == nil || IsNil(o.EnabledAlertSeverity) {
		var ret []EnumalertHandlerEnabledAlertSeverityProp
		return ret
	}
	return o.EnabledAlertSeverity
}

// GetEnabledAlertSeverityOk returns a tuple with the EnabledAlertSeverity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddGroovyScriptedAlertHandlerRequest) GetEnabledAlertSeverityOk() ([]EnumalertHandlerEnabledAlertSeverityProp, bool) {
	if o == nil || IsNil(o.EnabledAlertSeverity) {
		return nil, false
	}
	return o.EnabledAlertSeverity, true
}

// HasEnabledAlertSeverity returns a boolean if a field has been set.
func (o *AddGroovyScriptedAlertHandlerRequest) HasEnabledAlertSeverity() bool {
	if o != nil && !IsNil(o.EnabledAlertSeverity) {
		return true
	}

	return false
}

// SetEnabledAlertSeverity gets a reference to the given []EnumalertHandlerEnabledAlertSeverityProp and assigns it to the EnabledAlertSeverity field.
func (o *AddGroovyScriptedAlertHandlerRequest) SetEnabledAlertSeverity(v []EnumalertHandlerEnabledAlertSeverityProp) {
	o.EnabledAlertSeverity = v
}

// GetEnabledAlertType returns the EnabledAlertType field value if set, zero value otherwise.
func (o *AddGroovyScriptedAlertHandlerRequest) GetEnabledAlertType() []EnumalertHandlerEnabledAlertTypeProp {
	if o == nil || IsNil(o.EnabledAlertType) {
		var ret []EnumalertHandlerEnabledAlertTypeProp
		return ret
	}
	return o.EnabledAlertType
}

// GetEnabledAlertTypeOk returns a tuple with the EnabledAlertType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddGroovyScriptedAlertHandlerRequest) GetEnabledAlertTypeOk() ([]EnumalertHandlerEnabledAlertTypeProp, bool) {
	if o == nil || IsNil(o.EnabledAlertType) {
		return nil, false
	}
	return o.EnabledAlertType, true
}

// HasEnabledAlertType returns a boolean if a field has been set.
func (o *AddGroovyScriptedAlertHandlerRequest) HasEnabledAlertType() bool {
	if o != nil && !IsNil(o.EnabledAlertType) {
		return true
	}

	return false
}

// SetEnabledAlertType gets a reference to the given []EnumalertHandlerEnabledAlertTypeProp and assigns it to the EnabledAlertType field.
func (o *AddGroovyScriptedAlertHandlerRequest) SetEnabledAlertType(v []EnumalertHandlerEnabledAlertTypeProp) {
	o.EnabledAlertType = v
}

// GetDisabledAlertType returns the DisabledAlertType field value if set, zero value otherwise.
func (o *AddGroovyScriptedAlertHandlerRequest) GetDisabledAlertType() []EnumalertHandlerDisabledAlertTypeProp {
	if o == nil || IsNil(o.DisabledAlertType) {
		var ret []EnumalertHandlerDisabledAlertTypeProp
		return ret
	}
	return o.DisabledAlertType
}

// GetDisabledAlertTypeOk returns a tuple with the DisabledAlertType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddGroovyScriptedAlertHandlerRequest) GetDisabledAlertTypeOk() ([]EnumalertHandlerDisabledAlertTypeProp, bool) {
	if o == nil || IsNil(o.DisabledAlertType) {
		return nil, false
	}
	return o.DisabledAlertType, true
}

// HasDisabledAlertType returns a boolean if a field has been set.
func (o *AddGroovyScriptedAlertHandlerRequest) HasDisabledAlertType() bool {
	if o != nil && !IsNil(o.DisabledAlertType) {
		return true
	}

	return false
}

// SetDisabledAlertType gets a reference to the given []EnumalertHandlerDisabledAlertTypeProp and assigns it to the DisabledAlertType field.
func (o *AddGroovyScriptedAlertHandlerRequest) SetDisabledAlertType(v []EnumalertHandlerDisabledAlertTypeProp) {
	o.DisabledAlertType = v
}

func (o AddGroovyScriptedAlertHandlerRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddGroovyScriptedAlertHandlerRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["handlerName"] = o.HandlerName
	toSerialize["schemas"] = o.Schemas
	toSerialize["scriptClass"] = o.ScriptClass
	if !IsNil(o.ScriptArgument) {
		toSerialize["scriptArgument"] = o.ScriptArgument
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["enabled"] = o.Enabled
	if !IsNil(o.Asynchronous) {
		toSerialize["asynchronous"] = o.Asynchronous
	}
	if !IsNil(o.EnabledAlertSeverity) {
		toSerialize["enabledAlertSeverity"] = o.EnabledAlertSeverity
	}
	if !IsNil(o.EnabledAlertType) {
		toSerialize["enabledAlertType"] = o.EnabledAlertType
	}
	if !IsNil(o.DisabledAlertType) {
		toSerialize["disabledAlertType"] = o.DisabledAlertType
	}
	return toSerialize, nil
}

type NullableAddGroovyScriptedAlertHandlerRequest struct {
	value *AddGroovyScriptedAlertHandlerRequest
	isSet bool
}

func (v NullableAddGroovyScriptedAlertHandlerRequest) Get() *AddGroovyScriptedAlertHandlerRequest {
	return v.value
}

func (v *NullableAddGroovyScriptedAlertHandlerRequest) Set(val *AddGroovyScriptedAlertHandlerRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddGroovyScriptedAlertHandlerRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddGroovyScriptedAlertHandlerRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddGroovyScriptedAlertHandlerRequest(val *AddGroovyScriptedAlertHandlerRequest) *NullableAddGroovyScriptedAlertHandlerRequest {
	return &NullableAddGroovyScriptedAlertHandlerRequest{value: val, isSet: true}
}

func (v NullableAddGroovyScriptedAlertHandlerRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddGroovyScriptedAlertHandlerRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
