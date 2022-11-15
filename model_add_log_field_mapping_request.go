/*
PingData Config - OpenAPI 3.0

This is the PingData Configuration API

API version: 0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// AddLogFieldMappingRequest - struct for AddLogFieldMappingRequest
type AddLogFieldMappingRequest struct {
	AddAccessLogFieldMappingRequest *AddAccessLogFieldMappingRequest
	AddErrorLogFieldMappingRequest *AddErrorLogFieldMappingRequest
}

// AddAccessLogFieldMappingRequestAsAddLogFieldMappingRequest is a convenience function that returns AddAccessLogFieldMappingRequest wrapped in AddLogFieldMappingRequest
func AddAccessLogFieldMappingRequestAsAddLogFieldMappingRequest(v *AddAccessLogFieldMappingRequest) AddLogFieldMappingRequest {
	return AddLogFieldMappingRequest{
		AddAccessLogFieldMappingRequest: v,
	}
}

// AddErrorLogFieldMappingRequestAsAddLogFieldMappingRequest is a convenience function that returns AddErrorLogFieldMappingRequest wrapped in AddLogFieldMappingRequest
func AddErrorLogFieldMappingRequestAsAddLogFieldMappingRequest(v *AddErrorLogFieldMappingRequest) AddLogFieldMappingRequest {
	return AddLogFieldMappingRequest{
		AddErrorLogFieldMappingRequest: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *AddLogFieldMappingRequest) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AddAccessLogFieldMappingRequest
	err = newStrictDecoder(data).Decode(&dst.AddAccessLogFieldMappingRequest)
	if err == nil {
		jsonAddAccessLogFieldMappingRequest, _ := json.Marshal(dst.AddAccessLogFieldMappingRequest)
		if string(jsonAddAccessLogFieldMappingRequest) == "{}" { // empty struct
			dst.AddAccessLogFieldMappingRequest = nil
		} else {
			match++
		}
	} else {
		dst.AddAccessLogFieldMappingRequest = nil
	}

	// try to unmarshal data into AddErrorLogFieldMappingRequest
	err = newStrictDecoder(data).Decode(&dst.AddErrorLogFieldMappingRequest)
	if err == nil {
		jsonAddErrorLogFieldMappingRequest, _ := json.Marshal(dst.AddErrorLogFieldMappingRequest)
		if string(jsonAddErrorLogFieldMappingRequest) == "{}" { // empty struct
			dst.AddErrorLogFieldMappingRequest = nil
		} else {
			match++
		}
	} else {
		dst.AddErrorLogFieldMappingRequest = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.AddAccessLogFieldMappingRequest = nil
		dst.AddErrorLogFieldMappingRequest = nil

		return fmt.Errorf("data matches more than one schema in oneOf(AddLogFieldMappingRequest)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(AddLogFieldMappingRequest)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src AddLogFieldMappingRequest) MarshalJSON() ([]byte, error) {
	if src.AddAccessLogFieldMappingRequest != nil {
		return json.Marshal(&src.AddAccessLogFieldMappingRequest)
	}

	if src.AddErrorLogFieldMappingRequest != nil {
		return json.Marshal(&src.AddErrorLogFieldMappingRequest)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *AddLogFieldMappingRequest) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.AddAccessLogFieldMappingRequest != nil {
		return obj.AddAccessLogFieldMappingRequest
	}

	if obj.AddErrorLogFieldMappingRequest != nil {
		return obj.AddErrorLogFieldMappingRequest
	}

	// all schemas are nil
	return nil
}

type NullableAddLogFieldMappingRequest struct {
	value *AddLogFieldMappingRequest
	isSet bool
}

func (v NullableAddLogFieldMappingRequest) Get() *AddLogFieldMappingRequest {
	return v.value
}

func (v *NullableAddLogFieldMappingRequest) Set(val *AddLogFieldMappingRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddLogFieldMappingRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddLogFieldMappingRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddLogFieldMappingRequest(val *AddLogFieldMappingRequest) *NullableAddLogFieldMappingRequest {
	return &NullableAddLogFieldMappingRequest{value: val, isSet: true}
}

func (v NullableAddLogFieldMappingRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddLogFieldMappingRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


