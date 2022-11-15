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

// AddSearchReferenceCriteria200Response - struct for AddSearchReferenceCriteria200Response
type AddSearchReferenceCriteria200Response struct {
	AggregateSearchReferenceCriteriaResponse *AggregateSearchReferenceCriteriaResponse
	SimpleSearchReferenceCriteriaResponse *SimpleSearchReferenceCriteriaResponse
	ThirdPartySearchReferenceCriteriaResponse *ThirdPartySearchReferenceCriteriaResponse
}

// AggregateSearchReferenceCriteriaResponseAsAddSearchReferenceCriteria200Response is a convenience function that returns AggregateSearchReferenceCriteriaResponse wrapped in AddSearchReferenceCriteria200Response
func AggregateSearchReferenceCriteriaResponseAsAddSearchReferenceCriteria200Response(v *AggregateSearchReferenceCriteriaResponse) AddSearchReferenceCriteria200Response {
	return AddSearchReferenceCriteria200Response{
		AggregateSearchReferenceCriteriaResponse: v,
	}
}

// SimpleSearchReferenceCriteriaResponseAsAddSearchReferenceCriteria200Response is a convenience function that returns SimpleSearchReferenceCriteriaResponse wrapped in AddSearchReferenceCriteria200Response
func SimpleSearchReferenceCriteriaResponseAsAddSearchReferenceCriteria200Response(v *SimpleSearchReferenceCriteriaResponse) AddSearchReferenceCriteria200Response {
	return AddSearchReferenceCriteria200Response{
		SimpleSearchReferenceCriteriaResponse: v,
	}
}

// ThirdPartySearchReferenceCriteriaResponseAsAddSearchReferenceCriteria200Response is a convenience function that returns ThirdPartySearchReferenceCriteriaResponse wrapped in AddSearchReferenceCriteria200Response
func ThirdPartySearchReferenceCriteriaResponseAsAddSearchReferenceCriteria200Response(v *ThirdPartySearchReferenceCriteriaResponse) AddSearchReferenceCriteria200Response {
	return AddSearchReferenceCriteria200Response{
		ThirdPartySearchReferenceCriteriaResponse: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *AddSearchReferenceCriteria200Response) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AggregateSearchReferenceCriteriaResponse
	err = newStrictDecoder(data).Decode(&dst.AggregateSearchReferenceCriteriaResponse)
	if err == nil {
		jsonAggregateSearchReferenceCriteriaResponse, _ := json.Marshal(dst.AggregateSearchReferenceCriteriaResponse)
		if string(jsonAggregateSearchReferenceCriteriaResponse) == "{}" { // empty struct
			dst.AggregateSearchReferenceCriteriaResponse = nil
		} else {
			match++
		}
	} else {
		dst.AggregateSearchReferenceCriteriaResponse = nil
	}

	// try to unmarshal data into SimpleSearchReferenceCriteriaResponse
	err = newStrictDecoder(data).Decode(&dst.SimpleSearchReferenceCriteriaResponse)
	if err == nil {
		jsonSimpleSearchReferenceCriteriaResponse, _ := json.Marshal(dst.SimpleSearchReferenceCriteriaResponse)
		if string(jsonSimpleSearchReferenceCriteriaResponse) == "{}" { // empty struct
			dst.SimpleSearchReferenceCriteriaResponse = nil
		} else {
			match++
		}
	} else {
		dst.SimpleSearchReferenceCriteriaResponse = nil
	}

	// try to unmarshal data into ThirdPartySearchReferenceCriteriaResponse
	err = newStrictDecoder(data).Decode(&dst.ThirdPartySearchReferenceCriteriaResponse)
	if err == nil {
		jsonThirdPartySearchReferenceCriteriaResponse, _ := json.Marshal(dst.ThirdPartySearchReferenceCriteriaResponse)
		if string(jsonThirdPartySearchReferenceCriteriaResponse) == "{}" { // empty struct
			dst.ThirdPartySearchReferenceCriteriaResponse = nil
		} else {
			match++
		}
	} else {
		dst.ThirdPartySearchReferenceCriteriaResponse = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.AggregateSearchReferenceCriteriaResponse = nil
		dst.SimpleSearchReferenceCriteriaResponse = nil
		dst.ThirdPartySearchReferenceCriteriaResponse = nil

		return fmt.Errorf("data matches more than one schema in oneOf(AddSearchReferenceCriteria200Response)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(AddSearchReferenceCriteria200Response)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src AddSearchReferenceCriteria200Response) MarshalJSON() ([]byte, error) {
	if src.AggregateSearchReferenceCriteriaResponse != nil {
		return json.Marshal(&src.AggregateSearchReferenceCriteriaResponse)
	}

	if src.SimpleSearchReferenceCriteriaResponse != nil {
		return json.Marshal(&src.SimpleSearchReferenceCriteriaResponse)
	}

	if src.ThirdPartySearchReferenceCriteriaResponse != nil {
		return json.Marshal(&src.ThirdPartySearchReferenceCriteriaResponse)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *AddSearchReferenceCriteria200Response) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.AggregateSearchReferenceCriteriaResponse != nil {
		return obj.AggregateSearchReferenceCriteriaResponse
	}

	if obj.SimpleSearchReferenceCriteriaResponse != nil {
		return obj.SimpleSearchReferenceCriteriaResponse
	}

	if obj.ThirdPartySearchReferenceCriteriaResponse != nil {
		return obj.ThirdPartySearchReferenceCriteriaResponse
	}

	// all schemas are nil
	return nil
}

type NullableAddSearchReferenceCriteria200Response struct {
	value *AddSearchReferenceCriteria200Response
	isSet bool
}

func (v NullableAddSearchReferenceCriteria200Response) Get() *AddSearchReferenceCriteria200Response {
	return v.value
}

func (v *NullableAddSearchReferenceCriteria200Response) Set(val *AddSearchReferenceCriteria200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableAddSearchReferenceCriteria200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableAddSearchReferenceCriteria200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddSearchReferenceCriteria200Response(val *AddSearchReferenceCriteria200Response) *NullableAddSearchReferenceCriteria200Response {
	return &NullableAddSearchReferenceCriteria200Response{value: val, isSet: true}
}

func (v NullableAddSearchReferenceCriteria200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddSearchReferenceCriteria200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


