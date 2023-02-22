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

// SubOperationTimingPluginResponse struct for SubOperationTimingPluginResponse
type SubOperationTimingPluginResponse struct {
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
	// Name of the Plugin
	Id         string                                  `json:"id"`
	Schemas    []EnumsubOperationTimingPluginSchemaUrn `json:"schemas"`
	PluginType []EnumpluginPluginTypeProp              `json:"pluginType"`
	// Specifies a set of request criteria used to indicate that only operations for requests matching this criteria should be counted when aggregating timing data.
	RequestCriteria *string `json:"requestCriteria,omitempty"`
	// This controls how many of the most expensive phases are included per operation type in the monitor entry.
	NumMostExpensivePhasesShown int32 `json:"numMostExpensivePhasesShown"`
	// Indicates whether the plug-in should be invoked for internal operations.
	InvokeForInternalOperations *bool `json:"invokeForInternalOperations,omitempty"`
	// A description for this Plugin
	Description *string `json:"description,omitempty"`
	// Indicates whether the plug-in is enabled for use.
	Enabled bool `json:"enabled"`
}

// NewSubOperationTimingPluginResponse instantiates a new SubOperationTimingPluginResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubOperationTimingPluginResponse(id string, schemas []EnumsubOperationTimingPluginSchemaUrn, pluginType []EnumpluginPluginTypeProp, numMostExpensivePhasesShown int32, enabled bool) *SubOperationTimingPluginResponse {
	this := SubOperationTimingPluginResponse{}
	this.Id = id
	this.Schemas = schemas
	this.PluginType = pluginType
	this.NumMostExpensivePhasesShown = numMostExpensivePhasesShown
	this.Enabled = enabled
	return &this
}

// NewSubOperationTimingPluginResponseWithDefaults instantiates a new SubOperationTimingPluginResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubOperationTimingPluginResponseWithDefaults() *SubOperationTimingPluginResponse {
	this := SubOperationTimingPluginResponse{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *SubOperationTimingPluginResponse) GetMeta() MetaMeta {
	if o == nil || isNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubOperationTimingPluginResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || isNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *SubOperationTimingPluginResponse) HasMeta() bool {
	if o != nil && !isNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *SubOperationTimingPluginResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *SubOperationTimingPluginResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubOperationTimingPluginResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *SubOperationTimingPluginResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *SubOperationTimingPluginResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetId returns the Id field value
func (o *SubOperationTimingPluginResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SubOperationTimingPluginResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SubOperationTimingPluginResponse) SetId(v string) {
	o.Id = v
}

// GetSchemas returns the Schemas field value
func (o *SubOperationTimingPluginResponse) GetSchemas() []EnumsubOperationTimingPluginSchemaUrn {
	if o == nil {
		var ret []EnumsubOperationTimingPluginSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *SubOperationTimingPluginResponse) GetSchemasOk() ([]EnumsubOperationTimingPluginSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *SubOperationTimingPluginResponse) SetSchemas(v []EnumsubOperationTimingPluginSchemaUrn) {
	o.Schemas = v
}

// GetPluginType returns the PluginType field value
func (o *SubOperationTimingPluginResponse) GetPluginType() []EnumpluginPluginTypeProp {
	if o == nil {
		var ret []EnumpluginPluginTypeProp
		return ret
	}

	return o.PluginType
}

// GetPluginTypeOk returns a tuple with the PluginType field value
// and a boolean to check if the value has been set.
func (o *SubOperationTimingPluginResponse) GetPluginTypeOk() ([]EnumpluginPluginTypeProp, bool) {
	if o == nil {
		return nil, false
	}
	return o.PluginType, true
}

// SetPluginType sets field value
func (o *SubOperationTimingPluginResponse) SetPluginType(v []EnumpluginPluginTypeProp) {
	o.PluginType = v
}

// GetRequestCriteria returns the RequestCriteria field value if set, zero value otherwise.
func (o *SubOperationTimingPluginResponse) GetRequestCriteria() string {
	if o == nil || isNil(o.RequestCriteria) {
		var ret string
		return ret
	}
	return *o.RequestCriteria
}

// GetRequestCriteriaOk returns a tuple with the RequestCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubOperationTimingPluginResponse) GetRequestCriteriaOk() (*string, bool) {
	if o == nil || isNil(o.RequestCriteria) {
		return nil, false
	}
	return o.RequestCriteria, true
}

// HasRequestCriteria returns a boolean if a field has been set.
func (o *SubOperationTimingPluginResponse) HasRequestCriteria() bool {
	if o != nil && !isNil(o.RequestCriteria) {
		return true
	}

	return false
}

// SetRequestCriteria gets a reference to the given string and assigns it to the RequestCriteria field.
func (o *SubOperationTimingPluginResponse) SetRequestCriteria(v string) {
	o.RequestCriteria = &v
}

// GetNumMostExpensivePhasesShown returns the NumMostExpensivePhasesShown field value
func (o *SubOperationTimingPluginResponse) GetNumMostExpensivePhasesShown() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NumMostExpensivePhasesShown
}

// GetNumMostExpensivePhasesShownOk returns a tuple with the NumMostExpensivePhasesShown field value
// and a boolean to check if the value has been set.
func (o *SubOperationTimingPluginResponse) GetNumMostExpensivePhasesShownOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NumMostExpensivePhasesShown, true
}

// SetNumMostExpensivePhasesShown sets field value
func (o *SubOperationTimingPluginResponse) SetNumMostExpensivePhasesShown(v int32) {
	o.NumMostExpensivePhasesShown = v
}

// GetInvokeForInternalOperations returns the InvokeForInternalOperations field value if set, zero value otherwise.
func (o *SubOperationTimingPluginResponse) GetInvokeForInternalOperations() bool {
	if o == nil || isNil(o.InvokeForInternalOperations) {
		var ret bool
		return ret
	}
	return *o.InvokeForInternalOperations
}

// GetInvokeForInternalOperationsOk returns a tuple with the InvokeForInternalOperations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubOperationTimingPluginResponse) GetInvokeForInternalOperationsOk() (*bool, bool) {
	if o == nil || isNil(o.InvokeForInternalOperations) {
		return nil, false
	}
	return o.InvokeForInternalOperations, true
}

// HasInvokeForInternalOperations returns a boolean if a field has been set.
func (o *SubOperationTimingPluginResponse) HasInvokeForInternalOperations() bool {
	if o != nil && !isNil(o.InvokeForInternalOperations) {
		return true
	}

	return false
}

// SetInvokeForInternalOperations gets a reference to the given bool and assigns it to the InvokeForInternalOperations field.
func (o *SubOperationTimingPluginResponse) SetInvokeForInternalOperations(v bool) {
	o.InvokeForInternalOperations = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SubOperationTimingPluginResponse) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubOperationTimingPluginResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SubOperationTimingPluginResponse) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SubOperationTimingPluginResponse) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *SubOperationTimingPluginResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *SubOperationTimingPluginResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *SubOperationTimingPluginResponse) SetEnabled(v bool) {
	o.Enabled = v
}

func (o SubOperationTimingPluginResponse) MarshalJSON() ([]byte, error) {
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
		toSerialize["pluginType"] = o.PluginType
	}
	if !isNil(o.RequestCriteria) {
		toSerialize["requestCriteria"] = o.RequestCriteria
	}
	if true {
		toSerialize["numMostExpensivePhasesShown"] = o.NumMostExpensivePhasesShown
	}
	if !isNil(o.InvokeForInternalOperations) {
		toSerialize["invokeForInternalOperations"] = o.InvokeForInternalOperations
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableSubOperationTimingPluginResponse struct {
	value *SubOperationTimingPluginResponse
	isSet bool
}

func (v NullableSubOperationTimingPluginResponse) Get() *SubOperationTimingPluginResponse {
	return v.value
}

func (v *NullableSubOperationTimingPluginResponse) Set(val *SubOperationTimingPluginResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSubOperationTimingPluginResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSubOperationTimingPluginResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubOperationTimingPluginResponse(val *SubOperationTimingPluginResponse) *NullableSubOperationTimingPluginResponse {
	return &NullableSubOperationTimingPluginResponse{value: val, isSet: true}
}

func (v NullableSubOperationTimingPluginResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubOperationTimingPluginResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}