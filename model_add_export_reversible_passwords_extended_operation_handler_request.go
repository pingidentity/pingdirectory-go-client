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

// AddExportReversiblePasswordsExtendedOperationHandlerRequest struct for AddExportReversiblePasswordsExtendedOperationHandlerRequest
type AddExportReversiblePasswordsExtendedOperationHandlerRequest struct {
	// Name of the new Extended Operation Handler
	HandlerName string `json:"handlerName"`
	Schemas []EnumexportReversiblePasswordsExtendedOperationHandlerSchemaUrn `json:"schemas"`
	// A description for this Extended Operation Handler
	Description *string `json:"description,omitempty"`
	// Indicates whether the Extended Operation Handler is enabled (that is, whether the types of extended operations are allowed in the server).
	Enabled bool `json:"enabled"`
}

// NewAddExportReversiblePasswordsExtendedOperationHandlerRequest instantiates a new AddExportReversiblePasswordsExtendedOperationHandlerRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddExportReversiblePasswordsExtendedOperationHandlerRequest(handlerName string, schemas []EnumexportReversiblePasswordsExtendedOperationHandlerSchemaUrn, enabled bool) *AddExportReversiblePasswordsExtendedOperationHandlerRequest {
	this := AddExportReversiblePasswordsExtendedOperationHandlerRequest{}
	this.HandlerName = handlerName
	this.Schemas = schemas
	this.Enabled = enabled
	return &this
}

// NewAddExportReversiblePasswordsExtendedOperationHandlerRequestWithDefaults instantiates a new AddExportReversiblePasswordsExtendedOperationHandlerRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddExportReversiblePasswordsExtendedOperationHandlerRequestWithDefaults() *AddExportReversiblePasswordsExtendedOperationHandlerRequest {
	this := AddExportReversiblePasswordsExtendedOperationHandlerRequest{}
	return &this
}

// GetHandlerName returns the HandlerName field value
func (o *AddExportReversiblePasswordsExtendedOperationHandlerRequest) GetHandlerName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HandlerName
}

// GetHandlerNameOk returns a tuple with the HandlerName field value
// and a boolean to check if the value has been set.
func (o *AddExportReversiblePasswordsExtendedOperationHandlerRequest) GetHandlerNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.HandlerName, true
}

// SetHandlerName sets field value
func (o *AddExportReversiblePasswordsExtendedOperationHandlerRequest) SetHandlerName(v string) {
	o.HandlerName = v
}

// GetSchemas returns the Schemas field value
func (o *AddExportReversiblePasswordsExtendedOperationHandlerRequest) GetSchemas() []EnumexportReversiblePasswordsExtendedOperationHandlerSchemaUrn {
	if o == nil {
		var ret []EnumexportReversiblePasswordsExtendedOperationHandlerSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddExportReversiblePasswordsExtendedOperationHandlerRequest) GetSchemasOk() ([]EnumexportReversiblePasswordsExtendedOperationHandlerSchemaUrn, bool) {
	if o == nil {
    return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddExportReversiblePasswordsExtendedOperationHandlerRequest) SetSchemas(v []EnumexportReversiblePasswordsExtendedOperationHandlerSchemaUrn) {
	o.Schemas = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddExportReversiblePasswordsExtendedOperationHandlerRequest) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddExportReversiblePasswordsExtendedOperationHandlerRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddExportReversiblePasswordsExtendedOperationHandlerRequest) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddExportReversiblePasswordsExtendedOperationHandlerRequest) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *AddExportReversiblePasswordsExtendedOperationHandlerRequest) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AddExportReversiblePasswordsExtendedOperationHandlerRequest) GetEnabledOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AddExportReversiblePasswordsExtendedOperationHandlerRequest) SetEnabled(v bool) {
	o.Enabled = v
}

func (o AddExportReversiblePasswordsExtendedOperationHandlerRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["handlerName"] = o.HandlerName
	}
	if true {
		toSerialize["schemas"] = o.Schemas
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableAddExportReversiblePasswordsExtendedOperationHandlerRequest struct {
	value *AddExportReversiblePasswordsExtendedOperationHandlerRequest
	isSet bool
}

func (v NullableAddExportReversiblePasswordsExtendedOperationHandlerRequest) Get() *AddExportReversiblePasswordsExtendedOperationHandlerRequest {
	return v.value
}

func (v *NullableAddExportReversiblePasswordsExtendedOperationHandlerRequest) Set(val *AddExportReversiblePasswordsExtendedOperationHandlerRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddExportReversiblePasswordsExtendedOperationHandlerRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddExportReversiblePasswordsExtendedOperationHandlerRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddExportReversiblePasswordsExtendedOperationHandlerRequest(val *AddExportReversiblePasswordsExtendedOperationHandlerRequest) *NullableAddExportReversiblePasswordsExtendedOperationHandlerRequest {
	return &NullableAddExportReversiblePasswordsExtendedOperationHandlerRequest{value: val, isSet: true}
}

func (v NullableAddExportReversiblePasswordsExtendedOperationHandlerRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddExportReversiblePasswordsExtendedOperationHandlerRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

