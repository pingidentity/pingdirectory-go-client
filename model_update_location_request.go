/*
PingData Location Config - OpenAPI 3.0

This is the PingData configuration API for the Location config object

API version: 0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// UpdateLocationRequest struct for UpdateLocationRequest
type UpdateLocationRequest struct {
	Operations []Operation `json:"operations"`
}

// NewUpdateLocationRequest instantiates a new UpdateLocationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateLocationRequest(operations []Operation) *UpdateLocationRequest {
	this := UpdateLocationRequest{}
	this.Operations = operations
	return &this
}

// NewUpdateLocationRequestWithDefaults instantiates a new UpdateLocationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateLocationRequestWithDefaults() *UpdateLocationRequest {
	this := UpdateLocationRequest{}
	return &this
}

// GetOperations returns the Operations field value
func (o *UpdateLocationRequest) GetOperations() []Operation {
	if o == nil {
		var ret []Operation
		return ret
	}

	return o.Operations
}

// GetOperationsOk returns a tuple with the Operations field value
// and a boolean to check if the value has been set.
func (o *UpdateLocationRequest) GetOperationsOk() ([]Operation, bool) {
	if o == nil {
    return nil, false
	}
	return o.Operations, true
}

// SetOperations sets field value
func (o *UpdateLocationRequest) SetOperations(v []Operation) {
	o.Operations = v
}

func (o UpdateLocationRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["operations"] = o.Operations
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateLocationRequest struct {
	value *UpdateLocationRequest
	isSet bool
}

func (v NullableUpdateLocationRequest) Get() *UpdateLocationRequest {
	return v.value
}

func (v *NullableUpdateLocationRequest) Set(val *UpdateLocationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateLocationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateLocationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateLocationRequest(val *UpdateLocationRequest) *NullableUpdateLocationRequest {
	return &NullableUpdateLocationRequest{value: val, isSet: true}
}

func (v NullableUpdateLocationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateLocationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

