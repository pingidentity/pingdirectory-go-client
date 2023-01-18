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

// GroovyScriptedAlertHandlerResponse struct for GroovyScriptedAlertHandlerResponse
type GroovyScriptedAlertHandlerResponse struct {
	Meta *MetaMeta `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
	// Name of the Alert Handler
	Id string `json:"id"`
	Schemas []EnumgroovyScriptedAlertHandlerSchemaUrn `json:"schemas"`
	// The fully-qualified name of the Groovy class providing the logic for the Groovy Scripted Alert Handler.
	ScriptClass string `json:"scriptClass"`
	// The set of arguments used to customize the behavior for the Scripted Alert Handler. Each configuration property should be given in the form 'name=value'.
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

// NewGroovyScriptedAlertHandlerResponse instantiates a new GroovyScriptedAlertHandlerResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroovyScriptedAlertHandlerResponse(id string, schemas []EnumgroovyScriptedAlertHandlerSchemaUrn, scriptClass string, enabled bool) *GroovyScriptedAlertHandlerResponse {
	this := GroovyScriptedAlertHandlerResponse{}
	this.Id = id
	this.Schemas = schemas
	this.ScriptClass = scriptClass
	this.Enabled = enabled
	return &this
}

// NewGroovyScriptedAlertHandlerResponseWithDefaults instantiates a new GroovyScriptedAlertHandlerResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroovyScriptedAlertHandlerResponseWithDefaults() *GroovyScriptedAlertHandlerResponse {
	this := GroovyScriptedAlertHandlerResponse{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *GroovyScriptedAlertHandlerResponse) GetMeta() MetaMeta {
	if o == nil || isNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroovyScriptedAlertHandlerResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || isNil(o.Meta) {
    return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *GroovyScriptedAlertHandlerResponse) HasMeta() bool {
	if o != nil && !isNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *GroovyScriptedAlertHandlerResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *GroovyScriptedAlertHandlerResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroovyScriptedAlertHandlerResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || isNil(o.Urnpingidentityschemasconfigurationmessages20) {
    return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *GroovyScriptedAlertHandlerResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *GroovyScriptedAlertHandlerResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetId returns the Id field value
func (o *GroovyScriptedAlertHandlerResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GroovyScriptedAlertHandlerResponse) GetIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GroovyScriptedAlertHandlerResponse) SetId(v string) {
	o.Id = v
}

// GetSchemas returns the Schemas field value
func (o *GroovyScriptedAlertHandlerResponse) GetSchemas() []EnumgroovyScriptedAlertHandlerSchemaUrn {
	if o == nil {
		var ret []EnumgroovyScriptedAlertHandlerSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *GroovyScriptedAlertHandlerResponse) GetSchemasOk() ([]EnumgroovyScriptedAlertHandlerSchemaUrn, bool) {
	if o == nil {
    return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *GroovyScriptedAlertHandlerResponse) SetSchemas(v []EnumgroovyScriptedAlertHandlerSchemaUrn) {
	o.Schemas = v
}

// GetScriptClass returns the ScriptClass field value
func (o *GroovyScriptedAlertHandlerResponse) GetScriptClass() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ScriptClass
}

// GetScriptClassOk returns a tuple with the ScriptClass field value
// and a boolean to check if the value has been set.
func (o *GroovyScriptedAlertHandlerResponse) GetScriptClassOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ScriptClass, true
}

// SetScriptClass sets field value
func (o *GroovyScriptedAlertHandlerResponse) SetScriptClass(v string) {
	o.ScriptClass = v
}

// GetScriptArgument returns the ScriptArgument field value if set, zero value otherwise.
func (o *GroovyScriptedAlertHandlerResponse) GetScriptArgument() []string {
	if o == nil || isNil(o.ScriptArgument) {
		var ret []string
		return ret
	}
	return o.ScriptArgument
}

// GetScriptArgumentOk returns a tuple with the ScriptArgument field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroovyScriptedAlertHandlerResponse) GetScriptArgumentOk() ([]string, bool) {
	if o == nil || isNil(o.ScriptArgument) {
    return nil, false
	}
	return o.ScriptArgument, true
}

// HasScriptArgument returns a boolean if a field has been set.
func (o *GroovyScriptedAlertHandlerResponse) HasScriptArgument() bool {
	if o != nil && !isNil(o.ScriptArgument) {
		return true
	}

	return false
}

// SetScriptArgument gets a reference to the given []string and assigns it to the ScriptArgument field.
func (o *GroovyScriptedAlertHandlerResponse) SetScriptArgument(v []string) {
	o.ScriptArgument = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *GroovyScriptedAlertHandlerResponse) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroovyScriptedAlertHandlerResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *GroovyScriptedAlertHandlerResponse) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *GroovyScriptedAlertHandlerResponse) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *GroovyScriptedAlertHandlerResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *GroovyScriptedAlertHandlerResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *GroovyScriptedAlertHandlerResponse) SetEnabled(v bool) {
	o.Enabled = v
}

// GetAsynchronous returns the Asynchronous field value if set, zero value otherwise.
func (o *GroovyScriptedAlertHandlerResponse) GetAsynchronous() bool {
	if o == nil || isNil(o.Asynchronous) {
		var ret bool
		return ret
	}
	return *o.Asynchronous
}

// GetAsynchronousOk returns a tuple with the Asynchronous field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroovyScriptedAlertHandlerResponse) GetAsynchronousOk() (*bool, bool) {
	if o == nil || isNil(o.Asynchronous) {
    return nil, false
	}
	return o.Asynchronous, true
}

// HasAsynchronous returns a boolean if a field has been set.
func (o *GroovyScriptedAlertHandlerResponse) HasAsynchronous() bool {
	if o != nil && !isNil(o.Asynchronous) {
		return true
	}

	return false
}

// SetAsynchronous gets a reference to the given bool and assigns it to the Asynchronous field.
func (o *GroovyScriptedAlertHandlerResponse) SetAsynchronous(v bool) {
	o.Asynchronous = &v
}

// GetEnabledAlertSeverity returns the EnabledAlertSeverity field value if set, zero value otherwise.
func (o *GroovyScriptedAlertHandlerResponse) GetEnabledAlertSeverity() []EnumalertHandlerEnabledAlertSeverityProp {
	if o == nil || isNil(o.EnabledAlertSeverity) {
		var ret []EnumalertHandlerEnabledAlertSeverityProp
		return ret
	}
	return o.EnabledAlertSeverity
}

// GetEnabledAlertSeverityOk returns a tuple with the EnabledAlertSeverity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroovyScriptedAlertHandlerResponse) GetEnabledAlertSeverityOk() ([]EnumalertHandlerEnabledAlertSeverityProp, bool) {
	if o == nil || isNil(o.EnabledAlertSeverity) {
    return nil, false
	}
	return o.EnabledAlertSeverity, true
}

// HasEnabledAlertSeverity returns a boolean if a field has been set.
func (o *GroovyScriptedAlertHandlerResponse) HasEnabledAlertSeverity() bool {
	if o != nil && !isNil(o.EnabledAlertSeverity) {
		return true
	}

	return false
}

// SetEnabledAlertSeverity gets a reference to the given []EnumalertHandlerEnabledAlertSeverityProp and assigns it to the EnabledAlertSeverity field.
func (o *GroovyScriptedAlertHandlerResponse) SetEnabledAlertSeverity(v []EnumalertHandlerEnabledAlertSeverityProp) {
	o.EnabledAlertSeverity = v
}

// GetEnabledAlertType returns the EnabledAlertType field value if set, zero value otherwise.
func (o *GroovyScriptedAlertHandlerResponse) GetEnabledAlertType() []EnumalertHandlerEnabledAlertTypeProp {
	if o == nil || isNil(o.EnabledAlertType) {
		var ret []EnumalertHandlerEnabledAlertTypeProp
		return ret
	}
	return o.EnabledAlertType
}

// GetEnabledAlertTypeOk returns a tuple with the EnabledAlertType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroovyScriptedAlertHandlerResponse) GetEnabledAlertTypeOk() ([]EnumalertHandlerEnabledAlertTypeProp, bool) {
	if o == nil || isNil(o.EnabledAlertType) {
    return nil, false
	}
	return o.EnabledAlertType, true
}

// HasEnabledAlertType returns a boolean if a field has been set.
func (o *GroovyScriptedAlertHandlerResponse) HasEnabledAlertType() bool {
	if o != nil && !isNil(o.EnabledAlertType) {
		return true
	}

	return false
}

// SetEnabledAlertType gets a reference to the given []EnumalertHandlerEnabledAlertTypeProp and assigns it to the EnabledAlertType field.
func (o *GroovyScriptedAlertHandlerResponse) SetEnabledAlertType(v []EnumalertHandlerEnabledAlertTypeProp) {
	o.EnabledAlertType = v
}

// GetDisabledAlertType returns the DisabledAlertType field value if set, zero value otherwise.
func (o *GroovyScriptedAlertHandlerResponse) GetDisabledAlertType() []EnumalertHandlerDisabledAlertTypeProp {
	if o == nil || isNil(o.DisabledAlertType) {
		var ret []EnumalertHandlerDisabledAlertTypeProp
		return ret
	}
	return o.DisabledAlertType
}

// GetDisabledAlertTypeOk returns a tuple with the DisabledAlertType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroovyScriptedAlertHandlerResponse) GetDisabledAlertTypeOk() ([]EnumalertHandlerDisabledAlertTypeProp, bool) {
	if o == nil || isNil(o.DisabledAlertType) {
    return nil, false
	}
	return o.DisabledAlertType, true
}

// HasDisabledAlertType returns a boolean if a field has been set.
func (o *GroovyScriptedAlertHandlerResponse) HasDisabledAlertType() bool {
	if o != nil && !isNil(o.DisabledAlertType) {
		return true
	}

	return false
}

// SetDisabledAlertType gets a reference to the given []EnumalertHandlerDisabledAlertTypeProp and assigns it to the DisabledAlertType field.
func (o *GroovyScriptedAlertHandlerResponse) SetDisabledAlertType(v []EnumalertHandlerDisabledAlertTypeProp) {
	o.DisabledAlertType = v
}

func (o GroovyScriptedAlertHandlerResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		toSerialize["urn:pingidentity:schemas:configuration:messages:2.0"] = o.Urnpingidentityschemasconfigurationmessages20
	}
	if true {
		toSerialize["id"] = o.Id
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

type NullableGroovyScriptedAlertHandlerResponse struct {
	value *GroovyScriptedAlertHandlerResponse
	isSet bool
}

func (v NullableGroovyScriptedAlertHandlerResponse) Get() *GroovyScriptedAlertHandlerResponse {
	return v.value
}

func (v *NullableGroovyScriptedAlertHandlerResponse) Set(val *GroovyScriptedAlertHandlerResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGroovyScriptedAlertHandlerResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGroovyScriptedAlertHandlerResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroovyScriptedAlertHandlerResponse(val *GroovyScriptedAlertHandlerResponse) *NullableGroovyScriptedAlertHandlerResponse {
	return &NullableGroovyScriptedAlertHandlerResponse{value: val, isSet: true}
}

func (v NullableGroovyScriptedAlertHandlerResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroovyScriptedAlertHandlerResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

