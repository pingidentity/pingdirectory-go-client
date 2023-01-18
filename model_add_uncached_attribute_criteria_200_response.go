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

// AddUncachedAttributeCriteria200Response - struct for AddUncachedAttributeCriteria200Response
type AddUncachedAttributeCriteria200Response struct {
	DefaultUncachedAttributeCriteriaResponse *DefaultUncachedAttributeCriteriaResponse
	GroovyScriptedUncachedAttributeCriteriaResponse *GroovyScriptedUncachedAttributeCriteriaResponse
	SimpleUncachedAttributeCriteriaResponse *SimpleUncachedAttributeCriteriaResponse
	ThirdPartyUncachedAttributeCriteriaResponse *ThirdPartyUncachedAttributeCriteriaResponse
}

// DefaultUncachedAttributeCriteriaResponseAsAddUncachedAttributeCriteria200Response is a convenience function that returns DefaultUncachedAttributeCriteriaResponse wrapped in AddUncachedAttributeCriteria200Response
func DefaultUncachedAttributeCriteriaResponseAsAddUncachedAttributeCriteria200Response(v *DefaultUncachedAttributeCriteriaResponse) AddUncachedAttributeCriteria200Response {
	return AddUncachedAttributeCriteria200Response{
		DefaultUncachedAttributeCriteriaResponse: v,
	}
}

// GroovyScriptedUncachedAttributeCriteriaResponseAsAddUncachedAttributeCriteria200Response is a convenience function that returns GroovyScriptedUncachedAttributeCriteriaResponse wrapped in AddUncachedAttributeCriteria200Response
func GroovyScriptedUncachedAttributeCriteriaResponseAsAddUncachedAttributeCriteria200Response(v *GroovyScriptedUncachedAttributeCriteriaResponse) AddUncachedAttributeCriteria200Response {
	return AddUncachedAttributeCriteria200Response{
		GroovyScriptedUncachedAttributeCriteriaResponse: v,
	}
}

// SimpleUncachedAttributeCriteriaResponseAsAddUncachedAttributeCriteria200Response is a convenience function that returns SimpleUncachedAttributeCriteriaResponse wrapped in AddUncachedAttributeCriteria200Response
func SimpleUncachedAttributeCriteriaResponseAsAddUncachedAttributeCriteria200Response(v *SimpleUncachedAttributeCriteriaResponse) AddUncachedAttributeCriteria200Response {
	return AddUncachedAttributeCriteria200Response{
		SimpleUncachedAttributeCriteriaResponse: v,
	}
}

// ThirdPartyUncachedAttributeCriteriaResponseAsAddUncachedAttributeCriteria200Response is a convenience function that returns ThirdPartyUncachedAttributeCriteriaResponse wrapped in AddUncachedAttributeCriteria200Response
func ThirdPartyUncachedAttributeCriteriaResponseAsAddUncachedAttributeCriteria200Response(v *ThirdPartyUncachedAttributeCriteriaResponse) AddUncachedAttributeCriteria200Response {
	return AddUncachedAttributeCriteria200Response{
		ThirdPartyUncachedAttributeCriteriaResponse: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *AddUncachedAttributeCriteria200Response) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into DefaultUncachedAttributeCriteriaResponse
	err = newStrictDecoder(data).Decode(&dst.DefaultUncachedAttributeCriteriaResponse)
	if err == nil {
		jsonDefaultUncachedAttributeCriteriaResponse, _ := json.Marshal(dst.DefaultUncachedAttributeCriteriaResponse)
		if string(jsonDefaultUncachedAttributeCriteriaResponse) == "{}" { // empty struct
			dst.DefaultUncachedAttributeCriteriaResponse = nil
		} else {
			match++
		}
	} else {
		dst.DefaultUncachedAttributeCriteriaResponse = nil
	}

	// try to unmarshal data into GroovyScriptedUncachedAttributeCriteriaResponse
	err = newStrictDecoder(data).Decode(&dst.GroovyScriptedUncachedAttributeCriteriaResponse)
	if err == nil {
		jsonGroovyScriptedUncachedAttributeCriteriaResponse, _ := json.Marshal(dst.GroovyScriptedUncachedAttributeCriteriaResponse)
		if string(jsonGroovyScriptedUncachedAttributeCriteriaResponse) == "{}" { // empty struct
			dst.GroovyScriptedUncachedAttributeCriteriaResponse = nil
		} else {
			match++
		}
	} else {
		dst.GroovyScriptedUncachedAttributeCriteriaResponse = nil
	}

	// try to unmarshal data into SimpleUncachedAttributeCriteriaResponse
	err = newStrictDecoder(data).Decode(&dst.SimpleUncachedAttributeCriteriaResponse)
	if err == nil {
		jsonSimpleUncachedAttributeCriteriaResponse, _ := json.Marshal(dst.SimpleUncachedAttributeCriteriaResponse)
		if string(jsonSimpleUncachedAttributeCriteriaResponse) == "{}" { // empty struct
			dst.SimpleUncachedAttributeCriteriaResponse = nil
		} else {
			match++
		}
	} else {
		dst.SimpleUncachedAttributeCriteriaResponse = nil
	}

	// try to unmarshal data into ThirdPartyUncachedAttributeCriteriaResponse
	err = newStrictDecoder(data).Decode(&dst.ThirdPartyUncachedAttributeCriteriaResponse)
	if err == nil {
		jsonThirdPartyUncachedAttributeCriteriaResponse, _ := json.Marshal(dst.ThirdPartyUncachedAttributeCriteriaResponse)
		if string(jsonThirdPartyUncachedAttributeCriteriaResponse) == "{}" { // empty struct
			dst.ThirdPartyUncachedAttributeCriteriaResponse = nil
		} else {
			match++
		}
	} else {
		dst.ThirdPartyUncachedAttributeCriteriaResponse = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.DefaultUncachedAttributeCriteriaResponse = nil
		dst.GroovyScriptedUncachedAttributeCriteriaResponse = nil
		dst.SimpleUncachedAttributeCriteriaResponse = nil
		dst.ThirdPartyUncachedAttributeCriteriaResponse = nil

		return fmt.Errorf("data matches more than one schema in oneOf(AddUncachedAttributeCriteria200Response)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(AddUncachedAttributeCriteria200Response)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src AddUncachedAttributeCriteria200Response) MarshalJSON() ([]byte, error) {
	if src.DefaultUncachedAttributeCriteriaResponse != nil {
		return json.Marshal(&src.DefaultUncachedAttributeCriteriaResponse)
	}

	if src.GroovyScriptedUncachedAttributeCriteriaResponse != nil {
		return json.Marshal(&src.GroovyScriptedUncachedAttributeCriteriaResponse)
	}

	if src.SimpleUncachedAttributeCriteriaResponse != nil {
		return json.Marshal(&src.SimpleUncachedAttributeCriteriaResponse)
	}

	if src.ThirdPartyUncachedAttributeCriteriaResponse != nil {
		return json.Marshal(&src.ThirdPartyUncachedAttributeCriteriaResponse)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *AddUncachedAttributeCriteria200Response) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.DefaultUncachedAttributeCriteriaResponse != nil {
		return obj.DefaultUncachedAttributeCriteriaResponse
	}

	if obj.GroovyScriptedUncachedAttributeCriteriaResponse != nil {
		return obj.GroovyScriptedUncachedAttributeCriteriaResponse
	}

	if obj.SimpleUncachedAttributeCriteriaResponse != nil {
		return obj.SimpleUncachedAttributeCriteriaResponse
	}

	if obj.ThirdPartyUncachedAttributeCriteriaResponse != nil {
		return obj.ThirdPartyUncachedAttributeCriteriaResponse
	}

	// all schemas are nil
	return nil
}

type NullableAddUncachedAttributeCriteria200Response struct {
	value *AddUncachedAttributeCriteria200Response
	isSet bool
}

func (v NullableAddUncachedAttributeCriteria200Response) Get() *AddUncachedAttributeCriteria200Response {
	return v.value
}

func (v *NullableAddUncachedAttributeCriteria200Response) Set(val *AddUncachedAttributeCriteria200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableAddUncachedAttributeCriteria200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableAddUncachedAttributeCriteria200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddUncachedAttributeCriteria200Response(val *AddUncachedAttributeCriteria200Response) *NullableAddUncachedAttributeCriteria200Response {
	return &NullableAddUncachedAttributeCriteria200Response{value: val, isSet: true}
}

func (v NullableAddUncachedAttributeCriteria200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddUncachedAttributeCriteria200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

