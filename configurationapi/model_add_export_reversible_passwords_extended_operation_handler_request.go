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

// checks if the AddExportReversiblePasswordsExtendedOperationHandlerRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddExportReversiblePasswordsExtendedOperationHandlerRequest{}

// AddExportReversiblePasswordsExtendedOperationHandlerRequest struct for AddExportReversiblePasswordsExtendedOperationHandlerRequest
type AddExportReversiblePasswordsExtendedOperationHandlerRequest struct {
	Schemas []EnumexportReversiblePasswordsExtendedOperationHandlerSchemaUrn `json:"schemas"`
	// A description for this Extended Operation Handler
	Description *string `json:"description,omitempty"`
	// Indicates whether the Extended Operation Handler is enabled (that is, whether the types of extended operations are allowed in the server).
	Enabled bool `json:"enabled"`
	// Name of the new Extended Operation Handler
	HandlerName string `json:"handlerName"`
}

// NewAddExportReversiblePasswordsExtendedOperationHandlerRequest instantiates a new AddExportReversiblePasswordsExtendedOperationHandlerRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddExportReversiblePasswordsExtendedOperationHandlerRequest(schemas []EnumexportReversiblePasswordsExtendedOperationHandlerSchemaUrn, enabled bool, handlerName string) *AddExportReversiblePasswordsExtendedOperationHandlerRequest {
	this := AddExportReversiblePasswordsExtendedOperationHandlerRequest{}
	this.Schemas = schemas
	this.Enabled = enabled
	this.HandlerName = handlerName
	return &this
}

// NewAddExportReversiblePasswordsExtendedOperationHandlerRequestWithDefaults instantiates a new AddExportReversiblePasswordsExtendedOperationHandlerRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddExportReversiblePasswordsExtendedOperationHandlerRequestWithDefaults() *AddExportReversiblePasswordsExtendedOperationHandlerRequest {
	this := AddExportReversiblePasswordsExtendedOperationHandlerRequest{}
	return &this
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
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddExportReversiblePasswordsExtendedOperationHandlerRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddExportReversiblePasswordsExtendedOperationHandlerRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
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

func (o AddExportReversiblePasswordsExtendedOperationHandlerRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddExportReversiblePasswordsExtendedOperationHandlerRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["schemas"] = o.Schemas
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["enabled"] = o.Enabled
	toSerialize["handlerName"] = o.HandlerName
	return toSerialize, nil
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
