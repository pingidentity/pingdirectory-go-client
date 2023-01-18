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

// GetMatchingRule200Response - struct for GetMatchingRule200Response
type GetMatchingRule200Response struct {
	ApproximateMatchingRuleResponse *ApproximateMatchingRuleResponse
	EqualityMatchingRuleResponse *EqualityMatchingRuleResponse
	OrderingMatchingRuleResponse *OrderingMatchingRuleResponse
	SubstringMatchingRuleResponse *SubstringMatchingRuleResponse
}

// ApproximateMatchingRuleResponseAsGetMatchingRule200Response is a convenience function that returns ApproximateMatchingRuleResponse wrapped in GetMatchingRule200Response
func ApproximateMatchingRuleResponseAsGetMatchingRule200Response(v *ApproximateMatchingRuleResponse) GetMatchingRule200Response {
	return GetMatchingRule200Response{
		ApproximateMatchingRuleResponse: v,
	}
}

// EqualityMatchingRuleResponseAsGetMatchingRule200Response is a convenience function that returns EqualityMatchingRuleResponse wrapped in GetMatchingRule200Response
func EqualityMatchingRuleResponseAsGetMatchingRule200Response(v *EqualityMatchingRuleResponse) GetMatchingRule200Response {
	return GetMatchingRule200Response{
		EqualityMatchingRuleResponse: v,
	}
}

// OrderingMatchingRuleResponseAsGetMatchingRule200Response is a convenience function that returns OrderingMatchingRuleResponse wrapped in GetMatchingRule200Response
func OrderingMatchingRuleResponseAsGetMatchingRule200Response(v *OrderingMatchingRuleResponse) GetMatchingRule200Response {
	return GetMatchingRule200Response{
		OrderingMatchingRuleResponse: v,
	}
}

// SubstringMatchingRuleResponseAsGetMatchingRule200Response is a convenience function that returns SubstringMatchingRuleResponse wrapped in GetMatchingRule200Response
func SubstringMatchingRuleResponseAsGetMatchingRule200Response(v *SubstringMatchingRuleResponse) GetMatchingRule200Response {
	return GetMatchingRule200Response{
		SubstringMatchingRuleResponse: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *GetMatchingRule200Response) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ApproximateMatchingRuleResponse
	err = newStrictDecoder(data).Decode(&dst.ApproximateMatchingRuleResponse)
	if err == nil {
		jsonApproximateMatchingRuleResponse, _ := json.Marshal(dst.ApproximateMatchingRuleResponse)
		if string(jsonApproximateMatchingRuleResponse) == "{}" { // empty struct
			dst.ApproximateMatchingRuleResponse = nil
		} else {
			match++
		}
	} else {
		dst.ApproximateMatchingRuleResponse = nil
	}

	// try to unmarshal data into EqualityMatchingRuleResponse
	err = newStrictDecoder(data).Decode(&dst.EqualityMatchingRuleResponse)
	if err == nil {
		jsonEqualityMatchingRuleResponse, _ := json.Marshal(dst.EqualityMatchingRuleResponse)
		if string(jsonEqualityMatchingRuleResponse) == "{}" { // empty struct
			dst.EqualityMatchingRuleResponse = nil
		} else {
			match++
		}
	} else {
		dst.EqualityMatchingRuleResponse = nil
	}

	// try to unmarshal data into OrderingMatchingRuleResponse
	err = newStrictDecoder(data).Decode(&dst.OrderingMatchingRuleResponse)
	if err == nil {
		jsonOrderingMatchingRuleResponse, _ := json.Marshal(dst.OrderingMatchingRuleResponse)
		if string(jsonOrderingMatchingRuleResponse) == "{}" { // empty struct
			dst.OrderingMatchingRuleResponse = nil
		} else {
			match++
		}
	} else {
		dst.OrderingMatchingRuleResponse = nil
	}

	// try to unmarshal data into SubstringMatchingRuleResponse
	err = newStrictDecoder(data).Decode(&dst.SubstringMatchingRuleResponse)
	if err == nil {
		jsonSubstringMatchingRuleResponse, _ := json.Marshal(dst.SubstringMatchingRuleResponse)
		if string(jsonSubstringMatchingRuleResponse) == "{}" { // empty struct
			dst.SubstringMatchingRuleResponse = nil
		} else {
			match++
		}
	} else {
		dst.SubstringMatchingRuleResponse = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ApproximateMatchingRuleResponse = nil
		dst.EqualityMatchingRuleResponse = nil
		dst.OrderingMatchingRuleResponse = nil
		dst.SubstringMatchingRuleResponse = nil

		return fmt.Errorf("data matches more than one schema in oneOf(GetMatchingRule200Response)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(GetMatchingRule200Response)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GetMatchingRule200Response) MarshalJSON() ([]byte, error) {
	if src.ApproximateMatchingRuleResponse != nil {
		return json.Marshal(&src.ApproximateMatchingRuleResponse)
	}

	if src.EqualityMatchingRuleResponse != nil {
		return json.Marshal(&src.EqualityMatchingRuleResponse)
	}

	if src.OrderingMatchingRuleResponse != nil {
		return json.Marshal(&src.OrderingMatchingRuleResponse)
	}

	if src.SubstringMatchingRuleResponse != nil {
		return json.Marshal(&src.SubstringMatchingRuleResponse)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GetMatchingRule200Response) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.ApproximateMatchingRuleResponse != nil {
		return obj.ApproximateMatchingRuleResponse
	}

	if obj.EqualityMatchingRuleResponse != nil {
		return obj.EqualityMatchingRuleResponse
	}

	if obj.OrderingMatchingRuleResponse != nil {
		return obj.OrderingMatchingRuleResponse
	}

	if obj.SubstringMatchingRuleResponse != nil {
		return obj.SubstringMatchingRuleResponse
	}

	// all schemas are nil
	return nil
}

type NullableGetMatchingRule200Response struct {
	value *GetMatchingRule200Response
	isSet bool
}

func (v NullableGetMatchingRule200Response) Get() *GetMatchingRule200Response {
	return v.value
}

func (v *NullableGetMatchingRule200Response) Set(val *GetMatchingRule200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMatchingRule200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMatchingRule200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMatchingRule200Response(val *GetMatchingRule200Response) *NullableGetMatchingRule200Response {
	return &NullableGetMatchingRule200Response{value: val, isSet: true}
}

func (v NullableGetMatchingRule200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetMatchingRule200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

