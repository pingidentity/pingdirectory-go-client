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

// AddSearchEntryCriteriaRequest - struct for AddSearchEntryCriteriaRequest
type AddSearchEntryCriteriaRequest struct {
	AddAggregateSearchEntryCriteriaRequest *AddAggregateSearchEntryCriteriaRequest
	AddSimpleSearchEntryCriteriaRequest *AddSimpleSearchEntryCriteriaRequest
	AddThirdPartySearchEntryCriteriaRequest *AddThirdPartySearchEntryCriteriaRequest
}

// AddAggregateSearchEntryCriteriaRequestAsAddSearchEntryCriteriaRequest is a convenience function that returns AddAggregateSearchEntryCriteriaRequest wrapped in AddSearchEntryCriteriaRequest
func AddAggregateSearchEntryCriteriaRequestAsAddSearchEntryCriteriaRequest(v *AddAggregateSearchEntryCriteriaRequest) AddSearchEntryCriteriaRequest {
	return AddSearchEntryCriteriaRequest{
		AddAggregateSearchEntryCriteriaRequest: v,
	}
}

// AddSimpleSearchEntryCriteriaRequestAsAddSearchEntryCriteriaRequest is a convenience function that returns AddSimpleSearchEntryCriteriaRequest wrapped in AddSearchEntryCriteriaRequest
func AddSimpleSearchEntryCriteriaRequestAsAddSearchEntryCriteriaRequest(v *AddSimpleSearchEntryCriteriaRequest) AddSearchEntryCriteriaRequest {
	return AddSearchEntryCriteriaRequest{
		AddSimpleSearchEntryCriteriaRequest: v,
	}
}

// AddThirdPartySearchEntryCriteriaRequestAsAddSearchEntryCriteriaRequest is a convenience function that returns AddThirdPartySearchEntryCriteriaRequest wrapped in AddSearchEntryCriteriaRequest
func AddThirdPartySearchEntryCriteriaRequestAsAddSearchEntryCriteriaRequest(v *AddThirdPartySearchEntryCriteriaRequest) AddSearchEntryCriteriaRequest {
	return AddSearchEntryCriteriaRequest{
		AddThirdPartySearchEntryCriteriaRequest: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *AddSearchEntryCriteriaRequest) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AddAggregateSearchEntryCriteriaRequest
	err = newStrictDecoder(data).Decode(&dst.AddAggregateSearchEntryCriteriaRequest)
	if err == nil {
		jsonAddAggregateSearchEntryCriteriaRequest, _ := json.Marshal(dst.AddAggregateSearchEntryCriteriaRequest)
		if string(jsonAddAggregateSearchEntryCriteriaRequest) == "{}" { // empty struct
			dst.AddAggregateSearchEntryCriteriaRequest = nil
		} else {
			match++
		}
	} else {
		dst.AddAggregateSearchEntryCriteriaRequest = nil
	}

	// try to unmarshal data into AddSimpleSearchEntryCriteriaRequest
	err = newStrictDecoder(data).Decode(&dst.AddSimpleSearchEntryCriteriaRequest)
	if err == nil {
		jsonAddSimpleSearchEntryCriteriaRequest, _ := json.Marshal(dst.AddSimpleSearchEntryCriteriaRequest)
		if string(jsonAddSimpleSearchEntryCriteriaRequest) == "{}" { // empty struct
			dst.AddSimpleSearchEntryCriteriaRequest = nil
		} else {
			match++
		}
	} else {
		dst.AddSimpleSearchEntryCriteriaRequest = nil
	}

	// try to unmarshal data into AddThirdPartySearchEntryCriteriaRequest
	err = newStrictDecoder(data).Decode(&dst.AddThirdPartySearchEntryCriteriaRequest)
	if err == nil {
		jsonAddThirdPartySearchEntryCriteriaRequest, _ := json.Marshal(dst.AddThirdPartySearchEntryCriteriaRequest)
		if string(jsonAddThirdPartySearchEntryCriteriaRequest) == "{}" { // empty struct
			dst.AddThirdPartySearchEntryCriteriaRequest = nil
		} else {
			match++
		}
	} else {
		dst.AddThirdPartySearchEntryCriteriaRequest = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.AddAggregateSearchEntryCriteriaRequest = nil
		dst.AddSimpleSearchEntryCriteriaRequest = nil
		dst.AddThirdPartySearchEntryCriteriaRequest = nil

		return fmt.Errorf("data matches more than one schema in oneOf(AddSearchEntryCriteriaRequest)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(AddSearchEntryCriteriaRequest)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src AddSearchEntryCriteriaRequest) MarshalJSON() ([]byte, error) {
	if src.AddAggregateSearchEntryCriteriaRequest != nil {
		return json.Marshal(&src.AddAggregateSearchEntryCriteriaRequest)
	}

	if src.AddSimpleSearchEntryCriteriaRequest != nil {
		return json.Marshal(&src.AddSimpleSearchEntryCriteriaRequest)
	}

	if src.AddThirdPartySearchEntryCriteriaRequest != nil {
		return json.Marshal(&src.AddThirdPartySearchEntryCriteriaRequest)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *AddSearchEntryCriteriaRequest) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.AddAggregateSearchEntryCriteriaRequest != nil {
		return obj.AddAggregateSearchEntryCriteriaRequest
	}

	if obj.AddSimpleSearchEntryCriteriaRequest != nil {
		return obj.AddSimpleSearchEntryCriteriaRequest
	}

	if obj.AddThirdPartySearchEntryCriteriaRequest != nil {
		return obj.AddThirdPartySearchEntryCriteriaRequest
	}

	// all schemas are nil
	return nil
}

type NullableAddSearchEntryCriteriaRequest struct {
	value *AddSearchEntryCriteriaRequest
	isSet bool
}

func (v NullableAddSearchEntryCriteriaRequest) Get() *AddSearchEntryCriteriaRequest {
	return v.value
}

func (v *NullableAddSearchEntryCriteriaRequest) Set(val *AddSearchEntryCriteriaRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddSearchEntryCriteriaRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddSearchEntryCriteriaRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddSearchEntryCriteriaRequest(val *AddSearchEntryCriteriaRequest) *NullableAddSearchEntryCriteriaRequest {
	return &NullableAddSearchEntryCriteriaRequest{value: val, isSet: true}
}

func (v NullableAddSearchEntryCriteriaRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddSearchEntryCriteriaRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


