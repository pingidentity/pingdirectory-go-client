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

// checks if the GroovyScriptedChangeSubscriptionHandlerResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroovyScriptedChangeSubscriptionHandlerResponse{}

// GroovyScriptedChangeSubscriptionHandlerResponse struct for GroovyScriptedChangeSubscriptionHandlerResponse
type GroovyScriptedChangeSubscriptionHandlerResponse struct {
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
	// Name of the Change Subscription Handler
	Id      string                                                 `json:"id"`
	Schemas []EnumgroovyScriptedChangeSubscriptionHandlerSchemaUrn `json:"schemas"`
	// The fully-qualified name of the Groovy class providing the logic for the Groovy Scripted Change Subscription Handler.
	ScriptClass string `json:"scriptClass"`
	// The set of arguments used to customize the behavior for the Scripted Change Subscription Handler. Each configuration property should be given in the form 'name=value'.
	ScriptArgument []string `json:"scriptArgument,omitempty"`
	// A description for this Change Subscription Handler
	Description *string `json:"description,omitempty"`
	// Indicates whether this change subscription handler is enabled within the server.
	Enabled bool `json:"enabled"`
	// The set of change subscriptions for which this change subscription handler should be notified. If no values are provided then it will be notified for all change subscriptions defined in the server.
	ChangeSubscription []string `json:"changeSubscription,omitempty"`
}

// NewGroovyScriptedChangeSubscriptionHandlerResponse instantiates a new GroovyScriptedChangeSubscriptionHandlerResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroovyScriptedChangeSubscriptionHandlerResponse(id string, schemas []EnumgroovyScriptedChangeSubscriptionHandlerSchemaUrn, scriptClass string, enabled bool) *GroovyScriptedChangeSubscriptionHandlerResponse {
	this := GroovyScriptedChangeSubscriptionHandlerResponse{}
	this.Id = id
	this.Schemas = schemas
	this.ScriptClass = scriptClass
	this.Enabled = enabled
	return &this
}

// NewGroovyScriptedChangeSubscriptionHandlerResponseWithDefaults instantiates a new GroovyScriptedChangeSubscriptionHandlerResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroovyScriptedChangeSubscriptionHandlerResponseWithDefaults() *GroovyScriptedChangeSubscriptionHandlerResponse {
	this := GroovyScriptedChangeSubscriptionHandlerResponse{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *GroovyScriptedChangeSubscriptionHandlerResponse) GetMeta() MetaMeta {
	if o == nil || IsNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroovyScriptedChangeSubscriptionHandlerResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *GroovyScriptedChangeSubscriptionHandlerResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *GroovyScriptedChangeSubscriptionHandlerResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *GroovyScriptedChangeSubscriptionHandlerResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroovyScriptedChangeSubscriptionHandlerResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *GroovyScriptedChangeSubscriptionHandlerResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *GroovyScriptedChangeSubscriptionHandlerResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetId returns the Id field value
func (o *GroovyScriptedChangeSubscriptionHandlerResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GroovyScriptedChangeSubscriptionHandlerResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GroovyScriptedChangeSubscriptionHandlerResponse) SetId(v string) {
	o.Id = v
}

// GetSchemas returns the Schemas field value
func (o *GroovyScriptedChangeSubscriptionHandlerResponse) GetSchemas() []EnumgroovyScriptedChangeSubscriptionHandlerSchemaUrn {
	if o == nil {
		var ret []EnumgroovyScriptedChangeSubscriptionHandlerSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *GroovyScriptedChangeSubscriptionHandlerResponse) GetSchemasOk() ([]EnumgroovyScriptedChangeSubscriptionHandlerSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *GroovyScriptedChangeSubscriptionHandlerResponse) SetSchemas(v []EnumgroovyScriptedChangeSubscriptionHandlerSchemaUrn) {
	o.Schemas = v
}

// GetScriptClass returns the ScriptClass field value
func (o *GroovyScriptedChangeSubscriptionHandlerResponse) GetScriptClass() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ScriptClass
}

// GetScriptClassOk returns a tuple with the ScriptClass field value
// and a boolean to check if the value has been set.
func (o *GroovyScriptedChangeSubscriptionHandlerResponse) GetScriptClassOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScriptClass, true
}

// SetScriptClass sets field value
func (o *GroovyScriptedChangeSubscriptionHandlerResponse) SetScriptClass(v string) {
	o.ScriptClass = v
}

// GetScriptArgument returns the ScriptArgument field value if set, zero value otherwise.
func (o *GroovyScriptedChangeSubscriptionHandlerResponse) GetScriptArgument() []string {
	if o == nil || IsNil(o.ScriptArgument) {
		var ret []string
		return ret
	}
	return o.ScriptArgument
}

// GetScriptArgumentOk returns a tuple with the ScriptArgument field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroovyScriptedChangeSubscriptionHandlerResponse) GetScriptArgumentOk() ([]string, bool) {
	if o == nil || IsNil(o.ScriptArgument) {
		return nil, false
	}
	return o.ScriptArgument, true
}

// HasScriptArgument returns a boolean if a field has been set.
func (o *GroovyScriptedChangeSubscriptionHandlerResponse) HasScriptArgument() bool {
	if o != nil && !IsNil(o.ScriptArgument) {
		return true
	}

	return false
}

// SetScriptArgument gets a reference to the given []string and assigns it to the ScriptArgument field.
func (o *GroovyScriptedChangeSubscriptionHandlerResponse) SetScriptArgument(v []string) {
	o.ScriptArgument = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *GroovyScriptedChangeSubscriptionHandlerResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroovyScriptedChangeSubscriptionHandlerResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *GroovyScriptedChangeSubscriptionHandlerResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *GroovyScriptedChangeSubscriptionHandlerResponse) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *GroovyScriptedChangeSubscriptionHandlerResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *GroovyScriptedChangeSubscriptionHandlerResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *GroovyScriptedChangeSubscriptionHandlerResponse) SetEnabled(v bool) {
	o.Enabled = v
}

// GetChangeSubscription returns the ChangeSubscription field value if set, zero value otherwise.
func (o *GroovyScriptedChangeSubscriptionHandlerResponse) GetChangeSubscription() []string {
	if o == nil || IsNil(o.ChangeSubscription) {
		var ret []string
		return ret
	}
	return o.ChangeSubscription
}

// GetChangeSubscriptionOk returns a tuple with the ChangeSubscription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroovyScriptedChangeSubscriptionHandlerResponse) GetChangeSubscriptionOk() ([]string, bool) {
	if o == nil || IsNil(o.ChangeSubscription) {
		return nil, false
	}
	return o.ChangeSubscription, true
}

// HasChangeSubscription returns a boolean if a field has been set.
func (o *GroovyScriptedChangeSubscriptionHandlerResponse) HasChangeSubscription() bool {
	if o != nil && !IsNil(o.ChangeSubscription) {
		return true
	}

	return false
}

// SetChangeSubscription gets a reference to the given []string and assigns it to the ChangeSubscription field.
func (o *GroovyScriptedChangeSubscriptionHandlerResponse) SetChangeSubscription(v []string) {
	o.ChangeSubscription = v
}

func (o GroovyScriptedChangeSubscriptionHandlerResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroovyScriptedChangeSubscriptionHandlerResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		toSerialize["urn:pingidentity:schemas:configuration:messages:2.0"] = o.Urnpingidentityschemasconfigurationmessages20
	}
	toSerialize["id"] = o.Id
	toSerialize["schemas"] = o.Schemas
	toSerialize["scriptClass"] = o.ScriptClass
	if !IsNil(o.ScriptArgument) {
		toSerialize["scriptArgument"] = o.ScriptArgument
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["enabled"] = o.Enabled
	if !IsNil(o.ChangeSubscription) {
		toSerialize["changeSubscription"] = o.ChangeSubscription
	}
	return toSerialize, nil
}

type NullableGroovyScriptedChangeSubscriptionHandlerResponse struct {
	value *GroovyScriptedChangeSubscriptionHandlerResponse
	isSet bool
}

func (v NullableGroovyScriptedChangeSubscriptionHandlerResponse) Get() *GroovyScriptedChangeSubscriptionHandlerResponse {
	return v.value
}

func (v *NullableGroovyScriptedChangeSubscriptionHandlerResponse) Set(val *GroovyScriptedChangeSubscriptionHandlerResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGroovyScriptedChangeSubscriptionHandlerResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGroovyScriptedChangeSubscriptionHandlerResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroovyScriptedChangeSubscriptionHandlerResponse(val *GroovyScriptedChangeSubscriptionHandlerResponse) *NullableGroovyScriptedChangeSubscriptionHandlerResponse {
	return &NullableGroovyScriptedChangeSubscriptionHandlerResponse{value: val, isSet: true}
}

func (v NullableGroovyScriptedChangeSubscriptionHandlerResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroovyScriptedChangeSubscriptionHandlerResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
