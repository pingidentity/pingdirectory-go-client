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

// AddSubOperationTimingPluginRequest struct for AddSubOperationTimingPluginRequest
type AddSubOperationTimingPluginRequest struct {
	// Name of the new Plugin
	PluginName string                                  `json:"pluginName"`
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

// NewAddSubOperationTimingPluginRequest instantiates a new AddSubOperationTimingPluginRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddSubOperationTimingPluginRequest(pluginName string, schemas []EnumsubOperationTimingPluginSchemaUrn, pluginType []EnumpluginPluginTypeProp, numMostExpensivePhasesShown int32, enabled bool) *AddSubOperationTimingPluginRequest {
	this := AddSubOperationTimingPluginRequest{}
	this.PluginName = pluginName
	this.Schemas = schemas
	this.PluginType = pluginType
	this.NumMostExpensivePhasesShown = numMostExpensivePhasesShown
	this.Enabled = enabled
	return &this
}

// NewAddSubOperationTimingPluginRequestWithDefaults instantiates a new AddSubOperationTimingPluginRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddSubOperationTimingPluginRequestWithDefaults() *AddSubOperationTimingPluginRequest {
	this := AddSubOperationTimingPluginRequest{}
	return &this
}

// GetPluginName returns the PluginName field value
func (o *AddSubOperationTimingPluginRequest) GetPluginName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PluginName
}

// GetPluginNameOk returns a tuple with the PluginName field value
// and a boolean to check if the value has been set.
func (o *AddSubOperationTimingPluginRequest) GetPluginNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PluginName, true
}

// SetPluginName sets field value
func (o *AddSubOperationTimingPluginRequest) SetPluginName(v string) {
	o.PluginName = v
}

// GetSchemas returns the Schemas field value
func (o *AddSubOperationTimingPluginRequest) GetSchemas() []EnumsubOperationTimingPluginSchemaUrn {
	if o == nil {
		var ret []EnumsubOperationTimingPluginSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddSubOperationTimingPluginRequest) GetSchemasOk() ([]EnumsubOperationTimingPluginSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddSubOperationTimingPluginRequest) SetSchemas(v []EnumsubOperationTimingPluginSchemaUrn) {
	o.Schemas = v
}

// GetPluginType returns the PluginType field value
func (o *AddSubOperationTimingPluginRequest) GetPluginType() []EnumpluginPluginTypeProp {
	if o == nil {
		var ret []EnumpluginPluginTypeProp
		return ret
	}

	return o.PluginType
}

// GetPluginTypeOk returns a tuple with the PluginType field value
// and a boolean to check if the value has been set.
func (o *AddSubOperationTimingPluginRequest) GetPluginTypeOk() ([]EnumpluginPluginTypeProp, bool) {
	if o == nil {
		return nil, false
	}
	return o.PluginType, true
}

// SetPluginType sets field value
func (o *AddSubOperationTimingPluginRequest) SetPluginType(v []EnumpluginPluginTypeProp) {
	o.PluginType = v
}

// GetRequestCriteria returns the RequestCriteria field value if set, zero value otherwise.
func (o *AddSubOperationTimingPluginRequest) GetRequestCriteria() string {
	if o == nil || isNil(o.RequestCriteria) {
		var ret string
		return ret
	}
	return *o.RequestCriteria
}

// GetRequestCriteriaOk returns a tuple with the RequestCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddSubOperationTimingPluginRequest) GetRequestCriteriaOk() (*string, bool) {
	if o == nil || isNil(o.RequestCriteria) {
		return nil, false
	}
	return o.RequestCriteria, true
}

// HasRequestCriteria returns a boolean if a field has been set.
func (o *AddSubOperationTimingPluginRequest) HasRequestCriteria() bool {
	if o != nil && !isNil(o.RequestCriteria) {
		return true
	}

	return false
}

// SetRequestCriteria gets a reference to the given string and assigns it to the RequestCriteria field.
func (o *AddSubOperationTimingPluginRequest) SetRequestCriteria(v string) {
	o.RequestCriteria = &v
}

// GetNumMostExpensivePhasesShown returns the NumMostExpensivePhasesShown field value
func (o *AddSubOperationTimingPluginRequest) GetNumMostExpensivePhasesShown() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NumMostExpensivePhasesShown
}

// GetNumMostExpensivePhasesShownOk returns a tuple with the NumMostExpensivePhasesShown field value
// and a boolean to check if the value has been set.
func (o *AddSubOperationTimingPluginRequest) GetNumMostExpensivePhasesShownOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NumMostExpensivePhasesShown, true
}

// SetNumMostExpensivePhasesShown sets field value
func (o *AddSubOperationTimingPluginRequest) SetNumMostExpensivePhasesShown(v int32) {
	o.NumMostExpensivePhasesShown = v
}

// GetInvokeForInternalOperations returns the InvokeForInternalOperations field value if set, zero value otherwise.
func (o *AddSubOperationTimingPluginRequest) GetInvokeForInternalOperations() bool {
	if o == nil || isNil(o.InvokeForInternalOperations) {
		var ret bool
		return ret
	}
	return *o.InvokeForInternalOperations
}

// GetInvokeForInternalOperationsOk returns a tuple with the InvokeForInternalOperations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddSubOperationTimingPluginRequest) GetInvokeForInternalOperationsOk() (*bool, bool) {
	if o == nil || isNil(o.InvokeForInternalOperations) {
		return nil, false
	}
	return o.InvokeForInternalOperations, true
}

// HasInvokeForInternalOperations returns a boolean if a field has been set.
func (o *AddSubOperationTimingPluginRequest) HasInvokeForInternalOperations() bool {
	if o != nil && !isNil(o.InvokeForInternalOperations) {
		return true
	}

	return false
}

// SetInvokeForInternalOperations gets a reference to the given bool and assigns it to the InvokeForInternalOperations field.
func (o *AddSubOperationTimingPluginRequest) SetInvokeForInternalOperations(v bool) {
	o.InvokeForInternalOperations = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddSubOperationTimingPluginRequest) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddSubOperationTimingPluginRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddSubOperationTimingPluginRequest) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddSubOperationTimingPluginRequest) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *AddSubOperationTimingPluginRequest) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AddSubOperationTimingPluginRequest) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AddSubOperationTimingPluginRequest) SetEnabled(v bool) {
	o.Enabled = v
}

func (o AddSubOperationTimingPluginRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pluginName"] = o.PluginName
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

type NullableAddSubOperationTimingPluginRequest struct {
	value *AddSubOperationTimingPluginRequest
	isSet bool
}

func (v NullableAddSubOperationTimingPluginRequest) Get() *AddSubOperationTimingPluginRequest {
	return v.value
}

func (v *NullableAddSubOperationTimingPluginRequest) Set(val *AddSubOperationTimingPluginRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddSubOperationTimingPluginRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddSubOperationTimingPluginRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddSubOperationTimingPluginRequest(val *AddSubOperationTimingPluginRequest) *NullableAddSubOperationTimingPluginRequest {
	return &NullableAddSubOperationTimingPluginRequest{value: val, isSet: true}
}

func (v NullableAddSubOperationTimingPluginRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddSubOperationTimingPluginRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
