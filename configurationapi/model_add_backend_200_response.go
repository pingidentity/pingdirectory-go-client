/*
PingData Config - OpenAPI 3.0

This is the PingData Configuration API

API version: 0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configurationapi

import (
	"encoding/json"
	"fmt"
)

// AddBackend200Response - struct for AddBackend200Response
type AddBackend200Response struct {
	LocalDbBackendResponse *LocalDbBackendResponse
}

// LocalDbBackendResponseAsAddBackend200Response is a convenience function that returns LocalDbBackendResponse wrapped in AddBackend200Response
func LocalDbBackendResponseAsAddBackend200Response(v *LocalDbBackendResponse) AddBackend200Response {
	return AddBackend200Response{
		LocalDbBackendResponse: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *AddBackend200Response) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into LocalDbBackendResponse
	err = newStrictDecoder(data).Decode(&dst.LocalDbBackendResponse)
	if err == nil {
		jsonLocalDbBackendResponse, _ := json.Marshal(dst.LocalDbBackendResponse)
		if string(jsonLocalDbBackendResponse) == "{}" { // empty struct
			dst.LocalDbBackendResponse = nil
		} else {
			match++
		}
	} else {
		dst.LocalDbBackendResponse = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.LocalDbBackendResponse = nil

		return fmt.Errorf("data matches more than one schema in oneOf(AddBackend200Response)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(AddBackend200Response)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src AddBackend200Response) MarshalJSON() ([]byte, error) {
	if src.LocalDbBackendResponse != nil {
		return json.Marshal(&src.LocalDbBackendResponse)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *AddBackend200Response) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.LocalDbBackendResponse != nil {
		return obj.LocalDbBackendResponse
	}

	// all schemas are nil
	return nil
}

type NullableAddBackend200Response struct {
	value *AddBackend200Response
	isSet bool
}

func (v NullableAddBackend200Response) Get() *AddBackend200Response {
	return v.value
}

func (v *NullableAddBackend200Response) Set(val *AddBackend200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableAddBackend200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableAddBackend200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddBackend200Response(val *AddBackend200Response) *NullableAddBackend200Response {
	return &NullableAddBackend200Response{value: val, isSet: true}
}

func (v NullableAddBackend200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddBackend200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
