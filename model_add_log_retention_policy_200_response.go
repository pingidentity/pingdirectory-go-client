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

// AddLogRetentionPolicy200Response - struct for AddLogRetentionPolicy200Response
type AddLogRetentionPolicy200Response struct {
	FileCountLogRetentionPolicyResponse     *FileCountLogRetentionPolicyResponse
	FreeDiskSpaceLogRetentionPolicyResponse *FreeDiskSpaceLogRetentionPolicyResponse
	NeverDeleteLogRetentionPolicyResponse   *NeverDeleteLogRetentionPolicyResponse
	SizeLimitLogRetentionPolicyResponse     *SizeLimitLogRetentionPolicyResponse
	TimeLimitLogRetentionPolicyResponse     *TimeLimitLogRetentionPolicyResponse
}

// FileCountLogRetentionPolicyResponseAsAddLogRetentionPolicy200Response is a convenience function that returns FileCountLogRetentionPolicyResponse wrapped in AddLogRetentionPolicy200Response
func FileCountLogRetentionPolicyResponseAsAddLogRetentionPolicy200Response(v *FileCountLogRetentionPolicyResponse) AddLogRetentionPolicy200Response {
	return AddLogRetentionPolicy200Response{
		FileCountLogRetentionPolicyResponse: v,
	}
}

// FreeDiskSpaceLogRetentionPolicyResponseAsAddLogRetentionPolicy200Response is a convenience function that returns FreeDiskSpaceLogRetentionPolicyResponse wrapped in AddLogRetentionPolicy200Response
func FreeDiskSpaceLogRetentionPolicyResponseAsAddLogRetentionPolicy200Response(v *FreeDiskSpaceLogRetentionPolicyResponse) AddLogRetentionPolicy200Response {
	return AddLogRetentionPolicy200Response{
		FreeDiskSpaceLogRetentionPolicyResponse: v,
	}
}

// NeverDeleteLogRetentionPolicyResponseAsAddLogRetentionPolicy200Response is a convenience function that returns NeverDeleteLogRetentionPolicyResponse wrapped in AddLogRetentionPolicy200Response
func NeverDeleteLogRetentionPolicyResponseAsAddLogRetentionPolicy200Response(v *NeverDeleteLogRetentionPolicyResponse) AddLogRetentionPolicy200Response {
	return AddLogRetentionPolicy200Response{
		NeverDeleteLogRetentionPolicyResponse: v,
	}
}

// SizeLimitLogRetentionPolicyResponseAsAddLogRetentionPolicy200Response is a convenience function that returns SizeLimitLogRetentionPolicyResponse wrapped in AddLogRetentionPolicy200Response
func SizeLimitLogRetentionPolicyResponseAsAddLogRetentionPolicy200Response(v *SizeLimitLogRetentionPolicyResponse) AddLogRetentionPolicy200Response {
	return AddLogRetentionPolicy200Response{
		SizeLimitLogRetentionPolicyResponse: v,
	}
}

// TimeLimitLogRetentionPolicyResponseAsAddLogRetentionPolicy200Response is a convenience function that returns TimeLimitLogRetentionPolicyResponse wrapped in AddLogRetentionPolicy200Response
func TimeLimitLogRetentionPolicyResponseAsAddLogRetentionPolicy200Response(v *TimeLimitLogRetentionPolicyResponse) AddLogRetentionPolicy200Response {
	return AddLogRetentionPolicy200Response{
		TimeLimitLogRetentionPolicyResponse: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *AddLogRetentionPolicy200Response) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into FileCountLogRetentionPolicyResponse
	err = newStrictDecoder(data).Decode(&dst.FileCountLogRetentionPolicyResponse)
	if err == nil {
		jsonFileCountLogRetentionPolicyResponse, _ := json.Marshal(dst.FileCountLogRetentionPolicyResponse)
		if string(jsonFileCountLogRetentionPolicyResponse) == "{}" { // empty struct
			dst.FileCountLogRetentionPolicyResponse = nil
		} else {
			match++
		}
	} else {
		dst.FileCountLogRetentionPolicyResponse = nil
	}

	// try to unmarshal data into FreeDiskSpaceLogRetentionPolicyResponse
	err = newStrictDecoder(data).Decode(&dst.FreeDiskSpaceLogRetentionPolicyResponse)
	if err == nil {
		jsonFreeDiskSpaceLogRetentionPolicyResponse, _ := json.Marshal(dst.FreeDiskSpaceLogRetentionPolicyResponse)
		if string(jsonFreeDiskSpaceLogRetentionPolicyResponse) == "{}" { // empty struct
			dst.FreeDiskSpaceLogRetentionPolicyResponse = nil
		} else {
			match++
		}
	} else {
		dst.FreeDiskSpaceLogRetentionPolicyResponse = nil
	}

	// try to unmarshal data into NeverDeleteLogRetentionPolicyResponse
	err = newStrictDecoder(data).Decode(&dst.NeverDeleteLogRetentionPolicyResponse)
	if err == nil {
		jsonNeverDeleteLogRetentionPolicyResponse, _ := json.Marshal(dst.NeverDeleteLogRetentionPolicyResponse)
		if string(jsonNeverDeleteLogRetentionPolicyResponse) == "{}" { // empty struct
			dst.NeverDeleteLogRetentionPolicyResponse = nil
		} else {
			match++
		}
	} else {
		dst.NeverDeleteLogRetentionPolicyResponse = nil
	}

	// try to unmarshal data into SizeLimitLogRetentionPolicyResponse
	err = newStrictDecoder(data).Decode(&dst.SizeLimitLogRetentionPolicyResponse)
	if err == nil {
		jsonSizeLimitLogRetentionPolicyResponse, _ := json.Marshal(dst.SizeLimitLogRetentionPolicyResponse)
		if string(jsonSizeLimitLogRetentionPolicyResponse) == "{}" { // empty struct
			dst.SizeLimitLogRetentionPolicyResponse = nil
		} else {
			match++
		}
	} else {
		dst.SizeLimitLogRetentionPolicyResponse = nil
	}

	// try to unmarshal data into TimeLimitLogRetentionPolicyResponse
	err = newStrictDecoder(data).Decode(&dst.TimeLimitLogRetentionPolicyResponse)
	if err == nil {
		jsonTimeLimitLogRetentionPolicyResponse, _ := json.Marshal(dst.TimeLimitLogRetentionPolicyResponse)
		if string(jsonTimeLimitLogRetentionPolicyResponse) == "{}" { // empty struct
			dst.TimeLimitLogRetentionPolicyResponse = nil
		} else {
			match++
		}
	} else {
		dst.TimeLimitLogRetentionPolicyResponse = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.FileCountLogRetentionPolicyResponse = nil
		dst.FreeDiskSpaceLogRetentionPolicyResponse = nil
		dst.NeverDeleteLogRetentionPolicyResponse = nil
		dst.SizeLimitLogRetentionPolicyResponse = nil
		dst.TimeLimitLogRetentionPolicyResponse = nil

		return fmt.Errorf("data matches more than one schema in oneOf(AddLogRetentionPolicy200Response)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(AddLogRetentionPolicy200Response)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src AddLogRetentionPolicy200Response) MarshalJSON() ([]byte, error) {
	if src.FileCountLogRetentionPolicyResponse != nil {
		return json.Marshal(&src.FileCountLogRetentionPolicyResponse)
	}

	if src.FreeDiskSpaceLogRetentionPolicyResponse != nil {
		return json.Marshal(&src.FreeDiskSpaceLogRetentionPolicyResponse)
	}

	if src.NeverDeleteLogRetentionPolicyResponse != nil {
		return json.Marshal(&src.NeverDeleteLogRetentionPolicyResponse)
	}

	if src.SizeLimitLogRetentionPolicyResponse != nil {
		return json.Marshal(&src.SizeLimitLogRetentionPolicyResponse)
	}

	if src.TimeLimitLogRetentionPolicyResponse != nil {
		return json.Marshal(&src.TimeLimitLogRetentionPolicyResponse)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *AddLogRetentionPolicy200Response) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.FileCountLogRetentionPolicyResponse != nil {
		return obj.FileCountLogRetentionPolicyResponse
	}

	if obj.FreeDiskSpaceLogRetentionPolicyResponse != nil {
		return obj.FreeDiskSpaceLogRetentionPolicyResponse
	}

	if obj.NeverDeleteLogRetentionPolicyResponse != nil {
		return obj.NeverDeleteLogRetentionPolicyResponse
	}

	if obj.SizeLimitLogRetentionPolicyResponse != nil {
		return obj.SizeLimitLogRetentionPolicyResponse
	}

	if obj.TimeLimitLogRetentionPolicyResponse != nil {
		return obj.TimeLimitLogRetentionPolicyResponse
	}

	// all schemas are nil
	return nil
}

type NullableAddLogRetentionPolicy200Response struct {
	value *AddLogRetentionPolicy200Response
	isSet bool
}

func (v NullableAddLogRetentionPolicy200Response) Get() *AddLogRetentionPolicy200Response {
	return v.value
}

func (v *NullableAddLogRetentionPolicy200Response) Set(val *AddLogRetentionPolicy200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableAddLogRetentionPolicy200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableAddLogRetentionPolicy200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddLogRetentionPolicy200Response(val *AddLogRetentionPolicy200Response) *NullableAddLogRetentionPolicy200Response {
	return &NullableAddLogRetentionPolicy200Response{value: val, isSet: true}
}

func (v NullableAddLogRetentionPolicy200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddLogRetentionPolicy200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
