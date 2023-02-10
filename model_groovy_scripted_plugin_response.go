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

// GroovyScriptedPluginResponse struct for GroovyScriptedPluginResponse
type GroovyScriptedPluginResponse struct {
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
	// Name of the Plugin
	Id      string                              `json:"id"`
	Schemas []EnumgroovyScriptedPluginSchemaUrn `json:"schemas"`
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

// NewGroovyScriptedPluginResponse instantiates a new GroovyScriptedPluginResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroovyScriptedPluginResponse(id string, schemas []EnumgroovyScriptedPluginSchemaUrn, scriptClass string, enabled bool, pluginType []EnumpluginPluginTypeProp) *GroovyScriptedPluginResponse {
	this := GroovyScriptedPluginResponse{}
	this.Id = id
	this.Schemas = schemas
	this.ScriptClass = scriptClass
	this.Enabled = enabled
	this.PluginType = pluginType
	return &this
}

// NewGroovyScriptedPluginResponseWithDefaults instantiates a new GroovyScriptedPluginResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroovyScriptedPluginResponseWithDefaults() *GroovyScriptedPluginResponse {
	this := GroovyScriptedPluginResponse{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *GroovyScriptedPluginResponse) GetMeta() MetaMeta {
	if o == nil || isNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroovyScriptedPluginResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || isNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *GroovyScriptedPluginResponse) HasMeta() bool {
	if o != nil && !isNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *GroovyScriptedPluginResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *GroovyScriptedPluginResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroovyScriptedPluginResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *GroovyScriptedPluginResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *GroovyScriptedPluginResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetId returns the Id field value
func (o *GroovyScriptedPluginResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GroovyScriptedPluginResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GroovyScriptedPluginResponse) SetId(v string) {
	o.Id = v
}

// GetSchemas returns the Schemas field value
func (o *GroovyScriptedPluginResponse) GetSchemas() []EnumgroovyScriptedPluginSchemaUrn {
	if o == nil {
		var ret []EnumgroovyScriptedPluginSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *GroovyScriptedPluginResponse) GetSchemasOk() ([]EnumgroovyScriptedPluginSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *GroovyScriptedPluginResponse) SetSchemas(v []EnumgroovyScriptedPluginSchemaUrn) {
	o.Schemas = v
}

// GetScriptClass returns the ScriptClass field value
func (o *GroovyScriptedPluginResponse) GetScriptClass() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ScriptClass
}

// GetScriptClassOk returns a tuple with the ScriptClass field value
// and a boolean to check if the value has been set.
func (o *GroovyScriptedPluginResponse) GetScriptClassOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScriptClass, true
}

// SetScriptClass sets field value
func (o *GroovyScriptedPluginResponse) SetScriptClass(v string) {
	o.ScriptClass = v
}

// GetRequestCriteria returns the RequestCriteria field value if set, zero value otherwise.
func (o *GroovyScriptedPluginResponse) GetRequestCriteria() string {
	if o == nil || isNil(o.RequestCriteria) {
		var ret string
		return ret
	}
	return *o.RequestCriteria
}

// GetRequestCriteriaOk returns a tuple with the RequestCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroovyScriptedPluginResponse) GetRequestCriteriaOk() (*string, bool) {
	if o == nil || isNil(o.RequestCriteria) {
		return nil, false
	}
	return o.RequestCriteria, true
}

// HasRequestCriteria returns a boolean if a field has been set.
func (o *GroovyScriptedPluginResponse) HasRequestCriteria() bool {
	if o != nil && !isNil(o.RequestCriteria) {
		return true
	}

	return false
}

// SetRequestCriteria gets a reference to the given string and assigns it to the RequestCriteria field.
func (o *GroovyScriptedPluginResponse) SetRequestCriteria(v string) {
	o.RequestCriteria = &v
}

// GetScriptArgument returns the ScriptArgument field value if set, zero value otherwise.
func (o *GroovyScriptedPluginResponse) GetScriptArgument() []string {
	if o == nil || isNil(o.ScriptArgument) {
		var ret []string
		return ret
	}
	return o.ScriptArgument
}

// GetScriptArgumentOk returns a tuple with the ScriptArgument field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroovyScriptedPluginResponse) GetScriptArgumentOk() ([]string, bool) {
	if o == nil || isNil(o.ScriptArgument) {
		return nil, false
	}
	return o.ScriptArgument, true
}

// HasScriptArgument returns a boolean if a field has been set.
func (o *GroovyScriptedPluginResponse) HasScriptArgument() bool {
	if o != nil && !isNil(o.ScriptArgument) {
		return true
	}

	return false
}

// SetScriptArgument gets a reference to the given []string and assigns it to the ScriptArgument field.
func (o *GroovyScriptedPluginResponse) SetScriptArgument(v []string) {
	o.ScriptArgument = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *GroovyScriptedPluginResponse) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroovyScriptedPluginResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *GroovyScriptedPluginResponse) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *GroovyScriptedPluginResponse) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *GroovyScriptedPluginResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *GroovyScriptedPluginResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *GroovyScriptedPluginResponse) SetEnabled(v bool) {
	o.Enabled = v
}

// GetPluginType returns the PluginType field value
func (o *GroovyScriptedPluginResponse) GetPluginType() []EnumpluginPluginTypeProp {
	if o == nil {
		var ret []EnumpluginPluginTypeProp
		return ret
	}

	return o.PluginType
}

// GetPluginTypeOk returns a tuple with the PluginType field value
// and a boolean to check if the value has been set.
func (o *GroovyScriptedPluginResponse) GetPluginTypeOk() ([]EnumpluginPluginTypeProp, bool) {
	if o == nil {
		return nil, false
	}
	return o.PluginType, true
}

// SetPluginType sets field value
func (o *GroovyScriptedPluginResponse) SetPluginType(v []EnumpluginPluginTypeProp) {
	o.PluginType = v
}

// GetInvokeForInternalOperations returns the InvokeForInternalOperations field value if set, zero value otherwise.
func (o *GroovyScriptedPluginResponse) GetInvokeForInternalOperations() bool {
	if o == nil || isNil(o.InvokeForInternalOperations) {
		var ret bool
		return ret
	}
	return *o.InvokeForInternalOperations
}

// GetInvokeForInternalOperationsOk returns a tuple with the InvokeForInternalOperations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroovyScriptedPluginResponse) GetInvokeForInternalOperationsOk() (*bool, bool) {
	if o == nil || isNil(o.InvokeForInternalOperations) {
		return nil, false
	}
	return o.InvokeForInternalOperations, true
}

// HasInvokeForInternalOperations returns a boolean if a field has been set.
func (o *GroovyScriptedPluginResponse) HasInvokeForInternalOperations() bool {
	if o != nil && !isNil(o.InvokeForInternalOperations) {
		return true
	}

	return false
}

// SetInvokeForInternalOperations gets a reference to the given bool and assigns it to the InvokeForInternalOperations field.
func (o *GroovyScriptedPluginResponse) SetInvokeForInternalOperations(v bool) {
	o.InvokeForInternalOperations = &v
}

func (o GroovyScriptedPluginResponse) MarshalJSON() ([]byte, error) {
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

type NullableGroovyScriptedPluginResponse struct {
	value *GroovyScriptedPluginResponse
	isSet bool
}

func (v NullableGroovyScriptedPluginResponse) Get() *GroovyScriptedPluginResponse {
	return v.value
}

func (v *NullableGroovyScriptedPluginResponse) Set(val *GroovyScriptedPluginResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGroovyScriptedPluginResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGroovyScriptedPluginResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroovyScriptedPluginResponse(val *GroovyScriptedPluginResponse) *NullableGroovyScriptedPluginResponse {
	return &NullableGroovyScriptedPluginResponse{value: val, isSet: true}
}

func (v NullableGroovyScriptedPluginResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroovyScriptedPluginResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
