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

// GroovyScriptedUncachedAttributeCriteriaResponse struct for GroovyScriptedUncachedAttributeCriteriaResponse
type GroovyScriptedUncachedAttributeCriteriaResponse struct {
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
	// Name of the Uncached Attribute Criteria
	Id      string                                                 `json:"id"`
	Schemas []EnumgroovyScriptedUncachedAttributeCriteriaSchemaUrn `json:"schemas"`
	// The fully-qualified name of the Groovy class providing the logic for the Groovy Scripted Uncached Attribute Criteria.
	ScriptClass string `json:"scriptClass"`
	// The set of arguments used to customize the behavior for the Scripted Uncached Attribute Criteria. Each configuration property should be given in the form 'name=value'.
	ScriptArgument []string `json:"scriptArgument,omitempty"`
	// A description for this Uncached Attribute Criteria
	Description *string `json:"description,omitempty"`
	// Indicates whether this Uncached Attribute Criteria is enabled for use in the server.
	Enabled bool `json:"enabled"`
}

// NewGroovyScriptedUncachedAttributeCriteriaResponse instantiates a new GroovyScriptedUncachedAttributeCriteriaResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroovyScriptedUncachedAttributeCriteriaResponse(id string, schemas []EnumgroovyScriptedUncachedAttributeCriteriaSchemaUrn, scriptClass string, enabled bool) *GroovyScriptedUncachedAttributeCriteriaResponse {
	this := GroovyScriptedUncachedAttributeCriteriaResponse{}
	this.Id = id
	this.Schemas = schemas
	this.ScriptClass = scriptClass
	this.Enabled = enabled
	return &this
}

// NewGroovyScriptedUncachedAttributeCriteriaResponseWithDefaults instantiates a new GroovyScriptedUncachedAttributeCriteriaResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroovyScriptedUncachedAttributeCriteriaResponseWithDefaults() *GroovyScriptedUncachedAttributeCriteriaResponse {
	this := GroovyScriptedUncachedAttributeCriteriaResponse{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *GroovyScriptedUncachedAttributeCriteriaResponse) GetMeta() MetaMeta {
	if o == nil || isNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroovyScriptedUncachedAttributeCriteriaResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || isNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *GroovyScriptedUncachedAttributeCriteriaResponse) HasMeta() bool {
	if o != nil && !isNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *GroovyScriptedUncachedAttributeCriteriaResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *GroovyScriptedUncachedAttributeCriteriaResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroovyScriptedUncachedAttributeCriteriaResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *GroovyScriptedUncachedAttributeCriteriaResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *GroovyScriptedUncachedAttributeCriteriaResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetId returns the Id field value
func (o *GroovyScriptedUncachedAttributeCriteriaResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GroovyScriptedUncachedAttributeCriteriaResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GroovyScriptedUncachedAttributeCriteriaResponse) SetId(v string) {
	o.Id = v
}

// GetSchemas returns the Schemas field value
func (o *GroovyScriptedUncachedAttributeCriteriaResponse) GetSchemas() []EnumgroovyScriptedUncachedAttributeCriteriaSchemaUrn {
	if o == nil {
		var ret []EnumgroovyScriptedUncachedAttributeCriteriaSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *GroovyScriptedUncachedAttributeCriteriaResponse) GetSchemasOk() ([]EnumgroovyScriptedUncachedAttributeCriteriaSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *GroovyScriptedUncachedAttributeCriteriaResponse) SetSchemas(v []EnumgroovyScriptedUncachedAttributeCriteriaSchemaUrn) {
	o.Schemas = v
}

// GetScriptClass returns the ScriptClass field value
func (o *GroovyScriptedUncachedAttributeCriteriaResponse) GetScriptClass() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ScriptClass
}

// GetScriptClassOk returns a tuple with the ScriptClass field value
// and a boolean to check if the value has been set.
func (o *GroovyScriptedUncachedAttributeCriteriaResponse) GetScriptClassOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScriptClass, true
}

// SetScriptClass sets field value
func (o *GroovyScriptedUncachedAttributeCriteriaResponse) SetScriptClass(v string) {
	o.ScriptClass = v
}

// GetScriptArgument returns the ScriptArgument field value if set, zero value otherwise.
func (o *GroovyScriptedUncachedAttributeCriteriaResponse) GetScriptArgument() []string {
	if o == nil || isNil(o.ScriptArgument) {
		var ret []string
		return ret
	}
	return o.ScriptArgument
}

// GetScriptArgumentOk returns a tuple with the ScriptArgument field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroovyScriptedUncachedAttributeCriteriaResponse) GetScriptArgumentOk() ([]string, bool) {
	if o == nil || isNil(o.ScriptArgument) {
		return nil, false
	}
	return o.ScriptArgument, true
}

// HasScriptArgument returns a boolean if a field has been set.
func (o *GroovyScriptedUncachedAttributeCriteriaResponse) HasScriptArgument() bool {
	if o != nil && !isNil(o.ScriptArgument) {
		return true
	}

	return false
}

// SetScriptArgument gets a reference to the given []string and assigns it to the ScriptArgument field.
func (o *GroovyScriptedUncachedAttributeCriteriaResponse) SetScriptArgument(v []string) {
	o.ScriptArgument = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *GroovyScriptedUncachedAttributeCriteriaResponse) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroovyScriptedUncachedAttributeCriteriaResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *GroovyScriptedUncachedAttributeCriteriaResponse) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *GroovyScriptedUncachedAttributeCriteriaResponse) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *GroovyScriptedUncachedAttributeCriteriaResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *GroovyScriptedUncachedAttributeCriteriaResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *GroovyScriptedUncachedAttributeCriteriaResponse) SetEnabled(v bool) {
	o.Enabled = v
}

func (o GroovyScriptedUncachedAttributeCriteriaResponse) MarshalJSON() ([]byte, error) {
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
	return json.Marshal(toSerialize)
}

type NullableGroovyScriptedUncachedAttributeCriteriaResponse struct {
	value *GroovyScriptedUncachedAttributeCriteriaResponse
	isSet bool
}

func (v NullableGroovyScriptedUncachedAttributeCriteriaResponse) Get() *GroovyScriptedUncachedAttributeCriteriaResponse {
	return v.value
}

func (v *NullableGroovyScriptedUncachedAttributeCriteriaResponse) Set(val *GroovyScriptedUncachedAttributeCriteriaResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGroovyScriptedUncachedAttributeCriteriaResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGroovyScriptedUncachedAttributeCriteriaResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroovyScriptedUncachedAttributeCriteriaResponse(val *GroovyScriptedUncachedAttributeCriteriaResponse) *NullableGroovyScriptedUncachedAttributeCriteriaResponse {
	return &NullableGroovyScriptedUncachedAttributeCriteriaResponse{value: val, isSet: true}
}

func (v NullableGroovyScriptedUncachedAttributeCriteriaResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroovyScriptedUncachedAttributeCriteriaResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
