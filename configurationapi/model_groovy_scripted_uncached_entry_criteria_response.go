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

// checks if the GroovyScriptedUncachedEntryCriteriaResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroovyScriptedUncachedEntryCriteriaResponse{}

// GroovyScriptedUncachedEntryCriteriaResponse struct for GroovyScriptedUncachedEntryCriteriaResponse
type GroovyScriptedUncachedEntryCriteriaResponse struct {
	Schemas []EnumgroovyScriptedUncachedEntryCriteriaSchemaUrn `json:"schemas"`
	// The fully-qualified name of the Groovy class providing the logic for the Groovy Scripted Uncached Entry Criteria.
	ScriptClass string `json:"scriptClass"`
	// The set of arguments used to customize the behavior for the Scripted Uncached Entry Criteria. Each configuration property should be given in the form 'name=value'.
	ScriptArgument []string `json:"scriptArgument,omitempty"`
	// A description for this Uncached Entry Criteria
	Description *string `json:"description,omitempty"`
	// Indicates whether this Uncached Entry Criteria is enabled for use in the server.
	Enabled                                       bool                                               `json:"enabled"`
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
	// Name of the Uncached Entry Criteria
	Id string `json:"id"`
}

// NewGroovyScriptedUncachedEntryCriteriaResponse instantiates a new GroovyScriptedUncachedEntryCriteriaResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroovyScriptedUncachedEntryCriteriaResponse(schemas []EnumgroovyScriptedUncachedEntryCriteriaSchemaUrn, scriptClass string, enabled bool, id string) *GroovyScriptedUncachedEntryCriteriaResponse {
	this := GroovyScriptedUncachedEntryCriteriaResponse{}
	this.Schemas = schemas
	this.ScriptClass = scriptClass
	this.Enabled = enabled
	this.Id = id
	return &this
}

// NewGroovyScriptedUncachedEntryCriteriaResponseWithDefaults instantiates a new GroovyScriptedUncachedEntryCriteriaResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroovyScriptedUncachedEntryCriteriaResponseWithDefaults() *GroovyScriptedUncachedEntryCriteriaResponse {
	this := GroovyScriptedUncachedEntryCriteriaResponse{}
	return &this
}

// GetSchemas returns the Schemas field value
func (o *GroovyScriptedUncachedEntryCriteriaResponse) GetSchemas() []EnumgroovyScriptedUncachedEntryCriteriaSchemaUrn {
	if o == nil {
		var ret []EnumgroovyScriptedUncachedEntryCriteriaSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *GroovyScriptedUncachedEntryCriteriaResponse) GetSchemasOk() ([]EnumgroovyScriptedUncachedEntryCriteriaSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *GroovyScriptedUncachedEntryCriteriaResponse) SetSchemas(v []EnumgroovyScriptedUncachedEntryCriteriaSchemaUrn) {
	o.Schemas = v
}

// GetScriptClass returns the ScriptClass field value
func (o *GroovyScriptedUncachedEntryCriteriaResponse) GetScriptClass() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ScriptClass
}

// GetScriptClassOk returns a tuple with the ScriptClass field value
// and a boolean to check if the value has been set.
func (o *GroovyScriptedUncachedEntryCriteriaResponse) GetScriptClassOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScriptClass, true
}

// SetScriptClass sets field value
func (o *GroovyScriptedUncachedEntryCriteriaResponse) SetScriptClass(v string) {
	o.ScriptClass = v
}

// GetScriptArgument returns the ScriptArgument field value if set, zero value otherwise.
func (o *GroovyScriptedUncachedEntryCriteriaResponse) GetScriptArgument() []string {
	if o == nil || IsNil(o.ScriptArgument) {
		var ret []string
		return ret
	}
	return o.ScriptArgument
}

// GetScriptArgumentOk returns a tuple with the ScriptArgument field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroovyScriptedUncachedEntryCriteriaResponse) GetScriptArgumentOk() ([]string, bool) {
	if o == nil || IsNil(o.ScriptArgument) {
		return nil, false
	}
	return o.ScriptArgument, true
}

// HasScriptArgument returns a boolean if a field has been set.
func (o *GroovyScriptedUncachedEntryCriteriaResponse) HasScriptArgument() bool {
	if o != nil && !IsNil(o.ScriptArgument) {
		return true
	}

	return false
}

// SetScriptArgument gets a reference to the given []string and assigns it to the ScriptArgument field.
func (o *GroovyScriptedUncachedEntryCriteriaResponse) SetScriptArgument(v []string) {
	o.ScriptArgument = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *GroovyScriptedUncachedEntryCriteriaResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroovyScriptedUncachedEntryCriteriaResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *GroovyScriptedUncachedEntryCriteriaResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *GroovyScriptedUncachedEntryCriteriaResponse) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *GroovyScriptedUncachedEntryCriteriaResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *GroovyScriptedUncachedEntryCriteriaResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *GroovyScriptedUncachedEntryCriteriaResponse) SetEnabled(v bool) {
	o.Enabled = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *GroovyScriptedUncachedEntryCriteriaResponse) GetMeta() MetaMeta {
	if o == nil || IsNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroovyScriptedUncachedEntryCriteriaResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *GroovyScriptedUncachedEntryCriteriaResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *GroovyScriptedUncachedEntryCriteriaResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *GroovyScriptedUncachedEntryCriteriaResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroovyScriptedUncachedEntryCriteriaResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *GroovyScriptedUncachedEntryCriteriaResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *GroovyScriptedUncachedEntryCriteriaResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetId returns the Id field value
func (o *GroovyScriptedUncachedEntryCriteriaResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GroovyScriptedUncachedEntryCriteriaResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GroovyScriptedUncachedEntryCriteriaResponse) SetId(v string) {
	o.Id = v
}

func (o GroovyScriptedUncachedEntryCriteriaResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroovyScriptedUncachedEntryCriteriaResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["schemas"] = o.Schemas
	toSerialize["scriptClass"] = o.ScriptClass
	if !IsNil(o.ScriptArgument) {
		toSerialize["scriptArgument"] = o.ScriptArgument
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["enabled"] = o.Enabled
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		toSerialize["urn:pingidentity:schemas:configuration:messages:2.0"] = o.Urnpingidentityschemasconfigurationmessages20
	}
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

type NullableGroovyScriptedUncachedEntryCriteriaResponse struct {
	value *GroovyScriptedUncachedEntryCriteriaResponse
	isSet bool
}

func (v NullableGroovyScriptedUncachedEntryCriteriaResponse) Get() *GroovyScriptedUncachedEntryCriteriaResponse {
	return v.value
}

func (v *NullableGroovyScriptedUncachedEntryCriteriaResponse) Set(val *GroovyScriptedUncachedEntryCriteriaResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGroovyScriptedUncachedEntryCriteriaResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGroovyScriptedUncachedEntryCriteriaResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroovyScriptedUncachedEntryCriteriaResponse(val *GroovyScriptedUncachedEntryCriteriaResponse) *NullableGroovyScriptedUncachedEntryCriteriaResponse {
	return &NullableGroovyScriptedUncachedEntryCriteriaResponse{value: val, isSet: true}
}

func (v NullableGroovyScriptedUncachedEntryCriteriaResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroovyScriptedUncachedEntryCriteriaResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
