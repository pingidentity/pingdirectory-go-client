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

// checks if the AddLockAccountFailureLockoutActionRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddLockAccountFailureLockoutActionRequest{}

// AddLockAccountFailureLockoutActionRequest struct for AddLockAccountFailureLockoutActionRequest
type AddLockAccountFailureLockoutActionRequest struct {
	Schemas []EnumlockAccountFailureLockoutActionSchemaUrn `json:"schemas"`
	// A description for this Failure Lockout Action
	Description *string `json:"description,omitempty"`
	// Name of the new Failure Lockout Action
	ActionName string `json:"actionName"`
}

// NewAddLockAccountFailureLockoutActionRequest instantiates a new AddLockAccountFailureLockoutActionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddLockAccountFailureLockoutActionRequest(schemas []EnumlockAccountFailureLockoutActionSchemaUrn, actionName string) *AddLockAccountFailureLockoutActionRequest {
	this := AddLockAccountFailureLockoutActionRequest{}
	this.Schemas = schemas
	this.ActionName = actionName
	return &this
}

// NewAddLockAccountFailureLockoutActionRequestWithDefaults instantiates a new AddLockAccountFailureLockoutActionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddLockAccountFailureLockoutActionRequestWithDefaults() *AddLockAccountFailureLockoutActionRequest {
	this := AddLockAccountFailureLockoutActionRequest{}
	return &this
}

// GetSchemas returns the Schemas field value
func (o *AddLockAccountFailureLockoutActionRequest) GetSchemas() []EnumlockAccountFailureLockoutActionSchemaUrn {
	if o == nil {
		var ret []EnumlockAccountFailureLockoutActionSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddLockAccountFailureLockoutActionRequest) GetSchemasOk() ([]EnumlockAccountFailureLockoutActionSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddLockAccountFailureLockoutActionRequest) SetSchemas(v []EnumlockAccountFailureLockoutActionSchemaUrn) {
	o.Schemas = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddLockAccountFailureLockoutActionRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddLockAccountFailureLockoutActionRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddLockAccountFailureLockoutActionRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddLockAccountFailureLockoutActionRequest) SetDescription(v string) {
	o.Description = &v
}

// GetActionName returns the ActionName field value
func (o *AddLockAccountFailureLockoutActionRequest) GetActionName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ActionName
}

// GetActionNameOk returns a tuple with the ActionName field value
// and a boolean to check if the value has been set.
func (o *AddLockAccountFailureLockoutActionRequest) GetActionNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActionName, true
}

// SetActionName sets field value
func (o *AddLockAccountFailureLockoutActionRequest) SetActionName(v string) {
	o.ActionName = v
}

func (o AddLockAccountFailureLockoutActionRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddLockAccountFailureLockoutActionRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["schemas"] = o.Schemas
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["actionName"] = o.ActionName
	return toSerialize, nil
}

type NullableAddLockAccountFailureLockoutActionRequest struct {
	value *AddLockAccountFailureLockoutActionRequest
	isSet bool
}

func (v NullableAddLockAccountFailureLockoutActionRequest) Get() *AddLockAccountFailureLockoutActionRequest {
	return v.value
}

func (v *NullableAddLockAccountFailureLockoutActionRequest) Set(val *AddLockAccountFailureLockoutActionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddLockAccountFailureLockoutActionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddLockAccountFailureLockoutActionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddLockAccountFailureLockoutActionRequest(val *AddLockAccountFailureLockoutActionRequest) *NullableAddLockAccountFailureLockoutActionRequest {
	return &NullableAddLockAccountFailureLockoutActionRequest{value: val, isSet: true}
}

func (v NullableAddLockAccountFailureLockoutActionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddLockAccountFailureLockoutActionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
