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

// GetWebApplicationExtension200Response - struct for GetWebApplicationExtension200Response
type GetWebApplicationExtension200Response struct {
	ConsoleWebApplicationExtensionResponse *ConsoleWebApplicationExtensionResponse
	GenericWebApplicationExtensionResponse *GenericWebApplicationExtensionResponse
}

// ConsoleWebApplicationExtensionResponseAsGetWebApplicationExtension200Response is a convenience function that returns ConsoleWebApplicationExtensionResponse wrapped in GetWebApplicationExtension200Response
func ConsoleWebApplicationExtensionResponseAsGetWebApplicationExtension200Response(v *ConsoleWebApplicationExtensionResponse) GetWebApplicationExtension200Response {
	return GetWebApplicationExtension200Response{
		ConsoleWebApplicationExtensionResponse: v,
	}
}

// GenericWebApplicationExtensionResponseAsGetWebApplicationExtension200Response is a convenience function that returns GenericWebApplicationExtensionResponse wrapped in GetWebApplicationExtension200Response
func GenericWebApplicationExtensionResponseAsGetWebApplicationExtension200Response(v *GenericWebApplicationExtensionResponse) GetWebApplicationExtension200Response {
	return GetWebApplicationExtension200Response{
		GenericWebApplicationExtensionResponse: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *GetWebApplicationExtension200Response) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ConsoleWebApplicationExtensionResponse
	err = newStrictDecoder(data).Decode(&dst.ConsoleWebApplicationExtensionResponse)
	if err == nil {
		jsonConsoleWebApplicationExtensionResponse, _ := json.Marshal(dst.ConsoleWebApplicationExtensionResponse)
		if string(jsonConsoleWebApplicationExtensionResponse) == "{}" { // empty struct
			dst.ConsoleWebApplicationExtensionResponse = nil
		} else {
			match++
		}
	} else {
		dst.ConsoleWebApplicationExtensionResponse = nil
	}

	// try to unmarshal data into GenericWebApplicationExtensionResponse
	err = newStrictDecoder(data).Decode(&dst.GenericWebApplicationExtensionResponse)
	if err == nil {
		jsonGenericWebApplicationExtensionResponse, _ := json.Marshal(dst.GenericWebApplicationExtensionResponse)
		if string(jsonGenericWebApplicationExtensionResponse) == "{}" { // empty struct
			dst.GenericWebApplicationExtensionResponse = nil
		} else {
			match++
		}
	} else {
		dst.GenericWebApplicationExtensionResponse = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ConsoleWebApplicationExtensionResponse = nil
		dst.GenericWebApplicationExtensionResponse = nil

		return fmt.Errorf("data matches more than one schema in oneOf(GetWebApplicationExtension200Response)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(GetWebApplicationExtension200Response)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GetWebApplicationExtension200Response) MarshalJSON() ([]byte, error) {
	if src.ConsoleWebApplicationExtensionResponse != nil {
		return json.Marshal(&src.ConsoleWebApplicationExtensionResponse)
	}

	if src.GenericWebApplicationExtensionResponse != nil {
		return json.Marshal(&src.GenericWebApplicationExtensionResponse)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GetWebApplicationExtension200Response) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.ConsoleWebApplicationExtensionResponse != nil {
		return obj.ConsoleWebApplicationExtensionResponse
	}

	if obj.GenericWebApplicationExtensionResponse != nil {
		return obj.GenericWebApplicationExtensionResponse
	}

	// all schemas are nil
	return nil
}

type NullableGetWebApplicationExtension200Response struct {
	value *GetWebApplicationExtension200Response
	isSet bool
}

func (v NullableGetWebApplicationExtension200Response) Get() *GetWebApplicationExtension200Response {
	return v.value
}

func (v *NullableGetWebApplicationExtension200Response) Set(val *GetWebApplicationExtension200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetWebApplicationExtension200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetWebApplicationExtension200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetWebApplicationExtension200Response(val *GetWebApplicationExtension200Response) *NullableGetWebApplicationExtension200Response {
	return &NullableGetWebApplicationExtension200Response{value: val, isSet: true}
}

func (v NullableGetWebApplicationExtension200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetWebApplicationExtension200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}