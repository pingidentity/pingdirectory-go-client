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

// UncachedAttributeCriteriaListResponseResourcesInner - struct for UncachedAttributeCriteriaListResponseResourcesInner
type UncachedAttributeCriteriaListResponseResourcesInner struct {
	DefaultUncachedAttributeCriteriaResponse        *DefaultUncachedAttributeCriteriaResponse
	GroovyScriptedUncachedAttributeCriteriaResponse *GroovyScriptedUncachedAttributeCriteriaResponse
	SimpleUncachedAttributeCriteriaResponse         *SimpleUncachedAttributeCriteriaResponse
	ThirdPartyUncachedAttributeCriteriaResponse     *ThirdPartyUncachedAttributeCriteriaResponse
}

// DefaultUncachedAttributeCriteriaResponseAsUncachedAttributeCriteriaListResponseResourcesInner is a convenience function that returns DefaultUncachedAttributeCriteriaResponse wrapped in UncachedAttributeCriteriaListResponseResourcesInner
func DefaultUncachedAttributeCriteriaResponseAsUncachedAttributeCriteriaListResponseResourcesInner(v *DefaultUncachedAttributeCriteriaResponse) UncachedAttributeCriteriaListResponseResourcesInner {
	return UncachedAttributeCriteriaListResponseResourcesInner{
		DefaultUncachedAttributeCriteriaResponse: v,
	}
}

// GroovyScriptedUncachedAttributeCriteriaResponseAsUncachedAttributeCriteriaListResponseResourcesInner is a convenience function that returns GroovyScriptedUncachedAttributeCriteriaResponse wrapped in UncachedAttributeCriteriaListResponseResourcesInner
func GroovyScriptedUncachedAttributeCriteriaResponseAsUncachedAttributeCriteriaListResponseResourcesInner(v *GroovyScriptedUncachedAttributeCriteriaResponse) UncachedAttributeCriteriaListResponseResourcesInner {
	return UncachedAttributeCriteriaListResponseResourcesInner{
		GroovyScriptedUncachedAttributeCriteriaResponse: v,
	}
}

// SimpleUncachedAttributeCriteriaResponseAsUncachedAttributeCriteriaListResponseResourcesInner is a convenience function that returns SimpleUncachedAttributeCriteriaResponse wrapped in UncachedAttributeCriteriaListResponseResourcesInner
func SimpleUncachedAttributeCriteriaResponseAsUncachedAttributeCriteriaListResponseResourcesInner(v *SimpleUncachedAttributeCriteriaResponse) UncachedAttributeCriteriaListResponseResourcesInner {
	return UncachedAttributeCriteriaListResponseResourcesInner{
		SimpleUncachedAttributeCriteriaResponse: v,
	}
}

// ThirdPartyUncachedAttributeCriteriaResponseAsUncachedAttributeCriteriaListResponseResourcesInner is a convenience function that returns ThirdPartyUncachedAttributeCriteriaResponse wrapped in UncachedAttributeCriteriaListResponseResourcesInner
func ThirdPartyUncachedAttributeCriteriaResponseAsUncachedAttributeCriteriaListResponseResourcesInner(v *ThirdPartyUncachedAttributeCriteriaResponse) UncachedAttributeCriteriaListResponseResourcesInner {
	return UncachedAttributeCriteriaListResponseResourcesInner{
		ThirdPartyUncachedAttributeCriteriaResponse: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *UncachedAttributeCriteriaListResponseResourcesInner) UnmarshalJSON(data []byte) error {
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

		return fmt.Errorf("data matches more than one schema in oneOf(UncachedAttributeCriteriaListResponseResourcesInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(UncachedAttributeCriteriaListResponseResourcesInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src UncachedAttributeCriteriaListResponseResourcesInner) MarshalJSON() ([]byte, error) {
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
func (obj *UncachedAttributeCriteriaListResponseResourcesInner) GetActualInstance() interface{} {
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

type NullableUncachedAttributeCriteriaListResponseResourcesInner struct {
	value *UncachedAttributeCriteriaListResponseResourcesInner
	isSet bool
}

func (v NullableUncachedAttributeCriteriaListResponseResourcesInner) Get() *UncachedAttributeCriteriaListResponseResourcesInner {
	return v.value
}

func (v *NullableUncachedAttributeCriteriaListResponseResourcesInner) Set(val *UncachedAttributeCriteriaListResponseResourcesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableUncachedAttributeCriteriaListResponseResourcesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableUncachedAttributeCriteriaListResponseResourcesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUncachedAttributeCriteriaListResponseResourcesInner(val *UncachedAttributeCriteriaListResponseResourcesInner) *NullableUncachedAttributeCriteriaListResponseResourcesInner {
	return &NullableUncachedAttributeCriteriaListResponseResourcesInner{value: val, isSet: true}
}

func (v NullableUncachedAttributeCriteriaListResponseResourcesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUncachedAttributeCriteriaListResponseResourcesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
