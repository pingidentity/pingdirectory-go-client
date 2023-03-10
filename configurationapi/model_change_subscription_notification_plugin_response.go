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

// checks if the ChangeSubscriptionNotificationPluginResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChangeSubscriptionNotificationPluginResponse{}

// ChangeSubscriptionNotificationPluginResponse struct for ChangeSubscriptionNotificationPluginResponse
type ChangeSubscriptionNotificationPluginResponse struct {
	Meta                                          *MetaMeta                                           `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20  `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
	Schemas                                       []EnumchangeSubscriptionNotificationPluginSchemaUrn `json:"schemas"`
	// Name of the Plugin
	Id         string                     `json:"id"`
	PluginType []EnumpluginPluginTypeProp `json:"pluginType"`
	// A description for this Plugin
	Description *string `json:"description,omitempty"`
	// Indicates whether the plug-in is enabled for use.
	Enabled bool `json:"enabled"`
	// Indicates whether the plug-in should be invoked for internal operations.
	InvokeForInternalOperations *bool `json:"invokeForInternalOperations,omitempty"`
}

// NewChangeSubscriptionNotificationPluginResponse instantiates a new ChangeSubscriptionNotificationPluginResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChangeSubscriptionNotificationPluginResponse(schemas []EnumchangeSubscriptionNotificationPluginSchemaUrn, id string, pluginType []EnumpluginPluginTypeProp, enabled bool) *ChangeSubscriptionNotificationPluginResponse {
	this := ChangeSubscriptionNotificationPluginResponse{}
	this.Schemas = schemas
	this.Id = id
	this.PluginType = pluginType
	this.Enabled = enabled
	return &this
}

// NewChangeSubscriptionNotificationPluginResponseWithDefaults instantiates a new ChangeSubscriptionNotificationPluginResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChangeSubscriptionNotificationPluginResponseWithDefaults() *ChangeSubscriptionNotificationPluginResponse {
	this := ChangeSubscriptionNotificationPluginResponse{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ChangeSubscriptionNotificationPluginResponse) GetMeta() MetaMeta {
	if o == nil || IsNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeSubscriptionNotificationPluginResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ChangeSubscriptionNotificationPluginResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *ChangeSubscriptionNotificationPluginResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *ChangeSubscriptionNotificationPluginResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeSubscriptionNotificationPluginResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *ChangeSubscriptionNotificationPluginResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *ChangeSubscriptionNotificationPluginResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetSchemas returns the Schemas field value
func (o *ChangeSubscriptionNotificationPluginResponse) GetSchemas() []EnumchangeSubscriptionNotificationPluginSchemaUrn {
	if o == nil {
		var ret []EnumchangeSubscriptionNotificationPluginSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *ChangeSubscriptionNotificationPluginResponse) GetSchemasOk() ([]EnumchangeSubscriptionNotificationPluginSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *ChangeSubscriptionNotificationPluginResponse) SetSchemas(v []EnumchangeSubscriptionNotificationPluginSchemaUrn) {
	o.Schemas = v
}

// GetId returns the Id field value
func (o *ChangeSubscriptionNotificationPluginResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ChangeSubscriptionNotificationPluginResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ChangeSubscriptionNotificationPluginResponse) SetId(v string) {
	o.Id = v
}

// GetPluginType returns the PluginType field value
func (o *ChangeSubscriptionNotificationPluginResponse) GetPluginType() []EnumpluginPluginTypeProp {
	if o == nil {
		var ret []EnumpluginPluginTypeProp
		return ret
	}

	return o.PluginType
}

// GetPluginTypeOk returns a tuple with the PluginType field value
// and a boolean to check if the value has been set.
func (o *ChangeSubscriptionNotificationPluginResponse) GetPluginTypeOk() ([]EnumpluginPluginTypeProp, bool) {
	if o == nil {
		return nil, false
	}
	return o.PluginType, true
}

// SetPluginType sets field value
func (o *ChangeSubscriptionNotificationPluginResponse) SetPluginType(v []EnumpluginPluginTypeProp) {
	o.PluginType = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ChangeSubscriptionNotificationPluginResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeSubscriptionNotificationPluginResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ChangeSubscriptionNotificationPluginResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ChangeSubscriptionNotificationPluginResponse) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *ChangeSubscriptionNotificationPluginResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *ChangeSubscriptionNotificationPluginResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *ChangeSubscriptionNotificationPluginResponse) SetEnabled(v bool) {
	o.Enabled = v
}

// GetInvokeForInternalOperations returns the InvokeForInternalOperations field value if set, zero value otherwise.
func (o *ChangeSubscriptionNotificationPluginResponse) GetInvokeForInternalOperations() bool {
	if o == nil || IsNil(o.InvokeForInternalOperations) {
		var ret bool
		return ret
	}
	return *o.InvokeForInternalOperations
}

// GetInvokeForInternalOperationsOk returns a tuple with the InvokeForInternalOperations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeSubscriptionNotificationPluginResponse) GetInvokeForInternalOperationsOk() (*bool, bool) {
	if o == nil || IsNil(o.InvokeForInternalOperations) {
		return nil, false
	}
	return o.InvokeForInternalOperations, true
}

// HasInvokeForInternalOperations returns a boolean if a field has been set.
func (o *ChangeSubscriptionNotificationPluginResponse) HasInvokeForInternalOperations() bool {
	if o != nil && !IsNil(o.InvokeForInternalOperations) {
		return true
	}

	return false
}

// SetInvokeForInternalOperations gets a reference to the given bool and assigns it to the InvokeForInternalOperations field.
func (o *ChangeSubscriptionNotificationPluginResponse) SetInvokeForInternalOperations(v bool) {
	o.InvokeForInternalOperations = &v
}

func (o ChangeSubscriptionNotificationPluginResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChangeSubscriptionNotificationPluginResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		toSerialize["urn:pingidentity:schemas:configuration:messages:2.0"] = o.Urnpingidentityschemasconfigurationmessages20
	}
	toSerialize["schemas"] = o.Schemas
	toSerialize["id"] = o.Id
	toSerialize["pluginType"] = o.PluginType
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["enabled"] = o.Enabled
	if !IsNil(o.InvokeForInternalOperations) {
		toSerialize["invokeForInternalOperations"] = o.InvokeForInternalOperations
	}
	return toSerialize, nil
}

type NullableChangeSubscriptionNotificationPluginResponse struct {
	value *ChangeSubscriptionNotificationPluginResponse
	isSet bool
}

func (v NullableChangeSubscriptionNotificationPluginResponse) Get() *ChangeSubscriptionNotificationPluginResponse {
	return v.value
}

func (v *NullableChangeSubscriptionNotificationPluginResponse) Set(val *ChangeSubscriptionNotificationPluginResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableChangeSubscriptionNotificationPluginResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableChangeSubscriptionNotificationPluginResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChangeSubscriptionNotificationPluginResponse(val *ChangeSubscriptionNotificationPluginResponse) *NullableChangeSubscriptionNotificationPluginResponse {
	return &NullableChangeSubscriptionNotificationPluginResponse{value: val, isSet: true}
}

func (v NullableChangeSubscriptionNotificationPluginResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChangeSubscriptionNotificationPluginResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
