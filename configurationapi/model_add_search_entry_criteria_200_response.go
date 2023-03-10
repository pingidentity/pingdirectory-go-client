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

// AddSearchEntryCriteria200Response - struct for AddSearchEntryCriteria200Response
type AddSearchEntryCriteria200Response struct {
	AggregateSearchEntryCriteriaResponse  *AggregateSearchEntryCriteriaResponse
	SimpleSearchEntryCriteriaResponse     *SimpleSearchEntryCriteriaResponse
	ThirdPartySearchEntryCriteriaResponse *ThirdPartySearchEntryCriteriaResponse
}

// AggregateSearchEntryCriteriaResponseAsAddSearchEntryCriteria200Response is a convenience function that returns AggregateSearchEntryCriteriaResponse wrapped in AddSearchEntryCriteria200Response
func AggregateSearchEntryCriteriaResponseAsAddSearchEntryCriteria200Response(v *AggregateSearchEntryCriteriaResponse) AddSearchEntryCriteria200Response {
	return AddSearchEntryCriteria200Response{
		AggregateSearchEntryCriteriaResponse: v,
	}
}

// SimpleSearchEntryCriteriaResponseAsAddSearchEntryCriteria200Response is a convenience function that returns SimpleSearchEntryCriteriaResponse wrapped in AddSearchEntryCriteria200Response
func SimpleSearchEntryCriteriaResponseAsAddSearchEntryCriteria200Response(v *SimpleSearchEntryCriteriaResponse) AddSearchEntryCriteria200Response {
	return AddSearchEntryCriteria200Response{
		SimpleSearchEntryCriteriaResponse: v,
	}
}

// ThirdPartySearchEntryCriteriaResponseAsAddSearchEntryCriteria200Response is a convenience function that returns ThirdPartySearchEntryCriteriaResponse wrapped in AddSearchEntryCriteria200Response
func ThirdPartySearchEntryCriteriaResponseAsAddSearchEntryCriteria200Response(v *ThirdPartySearchEntryCriteriaResponse) AddSearchEntryCriteria200Response {
	return AddSearchEntryCriteria200Response{
		ThirdPartySearchEntryCriteriaResponse: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *AddSearchEntryCriteria200Response) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AggregateSearchEntryCriteriaResponse
	err = newStrictDecoder(data).Decode(&dst.AggregateSearchEntryCriteriaResponse)
	if err == nil {
		jsonAggregateSearchEntryCriteriaResponse, _ := json.Marshal(dst.AggregateSearchEntryCriteriaResponse)
		if string(jsonAggregateSearchEntryCriteriaResponse) == "{}" { // empty struct
			dst.AggregateSearchEntryCriteriaResponse = nil
		} else {
			match++
		}
	} else {
		dst.AggregateSearchEntryCriteriaResponse = nil
	}

	// try to unmarshal data into SimpleSearchEntryCriteriaResponse
	err = newStrictDecoder(data).Decode(&dst.SimpleSearchEntryCriteriaResponse)
	if err == nil {
		jsonSimpleSearchEntryCriteriaResponse, _ := json.Marshal(dst.SimpleSearchEntryCriteriaResponse)
		if string(jsonSimpleSearchEntryCriteriaResponse) == "{}" { // empty struct
			dst.SimpleSearchEntryCriteriaResponse = nil
		} else {
			match++
		}
	} else {
		dst.SimpleSearchEntryCriteriaResponse = nil
	}

	// try to unmarshal data into ThirdPartySearchEntryCriteriaResponse
	err = newStrictDecoder(data).Decode(&dst.ThirdPartySearchEntryCriteriaResponse)
	if err == nil {
		jsonThirdPartySearchEntryCriteriaResponse, _ := json.Marshal(dst.ThirdPartySearchEntryCriteriaResponse)
		if string(jsonThirdPartySearchEntryCriteriaResponse) == "{}" { // empty struct
			dst.ThirdPartySearchEntryCriteriaResponse = nil
		} else {
			match++
		}
	} else {
		dst.ThirdPartySearchEntryCriteriaResponse = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.AggregateSearchEntryCriteriaResponse = nil
		dst.SimpleSearchEntryCriteriaResponse = nil
		dst.ThirdPartySearchEntryCriteriaResponse = nil

		return fmt.Errorf("data matches more than one schema in oneOf(AddSearchEntryCriteria200Response)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(AddSearchEntryCriteria200Response)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src AddSearchEntryCriteria200Response) MarshalJSON() ([]byte, error) {
	if src.AggregateSearchEntryCriteriaResponse != nil {
		return json.Marshal(&src.AggregateSearchEntryCriteriaResponse)
	}

	if src.SimpleSearchEntryCriteriaResponse != nil {
		return json.Marshal(&src.SimpleSearchEntryCriteriaResponse)
	}

	if src.ThirdPartySearchEntryCriteriaResponse != nil {
		return json.Marshal(&src.ThirdPartySearchEntryCriteriaResponse)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *AddSearchEntryCriteria200Response) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.AggregateSearchEntryCriteriaResponse != nil {
		return obj.AggregateSearchEntryCriteriaResponse
	}

	if obj.SimpleSearchEntryCriteriaResponse != nil {
		return obj.SimpleSearchEntryCriteriaResponse
	}

	if obj.ThirdPartySearchEntryCriteriaResponse != nil {
		return obj.ThirdPartySearchEntryCriteriaResponse
	}

	// all schemas are nil
	return nil
}

type NullableAddSearchEntryCriteria200Response struct {
	value *AddSearchEntryCriteria200Response
	isSet bool
}

func (v NullableAddSearchEntryCriteria200Response) Get() *AddSearchEntryCriteria200Response {
	return v.value
}

func (v *NullableAddSearchEntryCriteria200Response) Set(val *AddSearchEntryCriteria200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableAddSearchEntryCriteria200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableAddSearchEntryCriteria200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddSearchEntryCriteria200Response(val *AddSearchEntryCriteria200Response) *NullableAddSearchEntryCriteria200Response {
	return &NullableAddSearchEntryCriteria200Response{value: val, isSet: true}
}

func (v NullableAddSearchEntryCriteria200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddSearchEntryCriteria200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
