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

// GetLogFieldSyntax200Response - struct for GetLogFieldSyntax200Response
type GetLogFieldSyntax200Response struct {
	AttributeBasedLogFieldSyntaxResponse *AttributeBasedLogFieldSyntaxResponse
	JsonLogFieldSyntaxResponse           *JsonLogFieldSyntaxResponse
}

// AttributeBasedLogFieldSyntaxResponseAsGetLogFieldSyntax200Response is a convenience function that returns AttributeBasedLogFieldSyntaxResponse wrapped in GetLogFieldSyntax200Response
func AttributeBasedLogFieldSyntaxResponseAsGetLogFieldSyntax200Response(v *AttributeBasedLogFieldSyntaxResponse) GetLogFieldSyntax200Response {
	return GetLogFieldSyntax200Response{
		AttributeBasedLogFieldSyntaxResponse: v,
	}
}

// JsonLogFieldSyntaxResponseAsGetLogFieldSyntax200Response is a convenience function that returns JsonLogFieldSyntaxResponse wrapped in GetLogFieldSyntax200Response
func JsonLogFieldSyntaxResponseAsGetLogFieldSyntax200Response(v *JsonLogFieldSyntaxResponse) GetLogFieldSyntax200Response {
	return GetLogFieldSyntax200Response{
		JsonLogFieldSyntaxResponse: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *GetLogFieldSyntax200Response) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AttributeBasedLogFieldSyntaxResponse
	err = newStrictDecoder(data).Decode(&dst.AttributeBasedLogFieldSyntaxResponse)
	if err == nil {
		jsonAttributeBasedLogFieldSyntaxResponse, _ := json.Marshal(dst.AttributeBasedLogFieldSyntaxResponse)
		if string(jsonAttributeBasedLogFieldSyntaxResponse) == "{}" { // empty struct
			dst.AttributeBasedLogFieldSyntaxResponse = nil
		} else {
			match++
		}
	} else {
		dst.AttributeBasedLogFieldSyntaxResponse = nil
	}

	// try to unmarshal data into JsonLogFieldSyntaxResponse
	err = newStrictDecoder(data).Decode(&dst.JsonLogFieldSyntaxResponse)
	if err == nil {
		jsonJsonLogFieldSyntaxResponse, _ := json.Marshal(dst.JsonLogFieldSyntaxResponse)
		if string(jsonJsonLogFieldSyntaxResponse) == "{}" { // empty struct
			dst.JsonLogFieldSyntaxResponse = nil
		} else {
			match++
		}
	} else {
		dst.JsonLogFieldSyntaxResponse = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.AttributeBasedLogFieldSyntaxResponse = nil
		dst.JsonLogFieldSyntaxResponse = nil

		return fmt.Errorf("data matches more than one schema in oneOf(GetLogFieldSyntax200Response)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(GetLogFieldSyntax200Response)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GetLogFieldSyntax200Response) MarshalJSON() ([]byte, error) {
	if src.AttributeBasedLogFieldSyntaxResponse != nil {
		return json.Marshal(&src.AttributeBasedLogFieldSyntaxResponse)
	}

	if src.JsonLogFieldSyntaxResponse != nil {
		return json.Marshal(&src.JsonLogFieldSyntaxResponse)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GetLogFieldSyntax200Response) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.AttributeBasedLogFieldSyntaxResponse != nil {
		return obj.AttributeBasedLogFieldSyntaxResponse
	}

	if obj.JsonLogFieldSyntaxResponse != nil {
		return obj.JsonLogFieldSyntaxResponse
	}

	// all schemas are nil
	return nil
}

type NullableGetLogFieldSyntax200Response struct {
	value *GetLogFieldSyntax200Response
	isSet bool
}

func (v NullableGetLogFieldSyntax200Response) Get() *GetLogFieldSyntax200Response {
	return v.value
}

func (v *NullableGetLogFieldSyntax200Response) Set(val *GetLogFieldSyntax200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetLogFieldSyntax200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetLogFieldSyntax200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetLogFieldSyntax200Response(val *GetLogFieldSyntax200Response) *NullableGetLogFieldSyntax200Response {
	return &NullableGetLogFieldSyntax200Response{value: val, isSet: true}
}

func (v NullableGetLogFieldSyntax200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetLogFieldSyntax200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
