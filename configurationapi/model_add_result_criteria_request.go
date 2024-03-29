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

// AddResultCriteriaRequest - struct for AddResultCriteriaRequest
type AddResultCriteriaRequest struct {
	AddAggregateResultCriteriaRequest            *AddAggregateResultCriteriaRequest
	AddReplicationAssuranceResultCriteriaRequest *AddReplicationAssuranceResultCriteriaRequest
	AddSimpleResultCriteriaRequest               *AddSimpleResultCriteriaRequest
	AddSuccessfulBindResultCriteriaRequest       *AddSuccessfulBindResultCriteriaRequest
	AddThirdPartyResultCriteriaRequest           *AddThirdPartyResultCriteriaRequest
}

// AddAggregateResultCriteriaRequestAsAddResultCriteriaRequest is a convenience function that returns AddAggregateResultCriteriaRequest wrapped in AddResultCriteriaRequest
func AddAggregateResultCriteriaRequestAsAddResultCriteriaRequest(v *AddAggregateResultCriteriaRequest) AddResultCriteriaRequest {
	return AddResultCriteriaRequest{
		AddAggregateResultCriteriaRequest: v,
	}
}

// AddReplicationAssuranceResultCriteriaRequestAsAddResultCriteriaRequest is a convenience function that returns AddReplicationAssuranceResultCriteriaRequest wrapped in AddResultCriteriaRequest
func AddReplicationAssuranceResultCriteriaRequestAsAddResultCriteriaRequest(v *AddReplicationAssuranceResultCriteriaRequest) AddResultCriteriaRequest {
	return AddResultCriteriaRequest{
		AddReplicationAssuranceResultCriteriaRequest: v,
	}
}

// AddSimpleResultCriteriaRequestAsAddResultCriteriaRequest is a convenience function that returns AddSimpleResultCriteriaRequest wrapped in AddResultCriteriaRequest
func AddSimpleResultCriteriaRequestAsAddResultCriteriaRequest(v *AddSimpleResultCriteriaRequest) AddResultCriteriaRequest {
	return AddResultCriteriaRequest{
		AddSimpleResultCriteriaRequest: v,
	}
}

// AddSuccessfulBindResultCriteriaRequestAsAddResultCriteriaRequest is a convenience function that returns AddSuccessfulBindResultCriteriaRequest wrapped in AddResultCriteriaRequest
func AddSuccessfulBindResultCriteriaRequestAsAddResultCriteriaRequest(v *AddSuccessfulBindResultCriteriaRequest) AddResultCriteriaRequest {
	return AddResultCriteriaRequest{
		AddSuccessfulBindResultCriteriaRequest: v,
	}
}

// AddThirdPartyResultCriteriaRequestAsAddResultCriteriaRequest is a convenience function that returns AddThirdPartyResultCriteriaRequest wrapped in AddResultCriteriaRequest
func AddThirdPartyResultCriteriaRequestAsAddResultCriteriaRequest(v *AddThirdPartyResultCriteriaRequest) AddResultCriteriaRequest {
	return AddResultCriteriaRequest{
		AddThirdPartyResultCriteriaRequest: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *AddResultCriteriaRequest) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AddAggregateResultCriteriaRequest
	err = newStrictDecoder(data).Decode(&dst.AddAggregateResultCriteriaRequest)
	if err == nil {
		jsonAddAggregateResultCriteriaRequest, _ := json.Marshal(dst.AddAggregateResultCriteriaRequest)
		if string(jsonAddAggregateResultCriteriaRequest) == "{}" { // empty struct
			dst.AddAggregateResultCriteriaRequest = nil
		} else {
			match++
		}
	} else {
		dst.AddAggregateResultCriteriaRequest = nil
	}

	// try to unmarshal data into AddReplicationAssuranceResultCriteriaRequest
	err = newStrictDecoder(data).Decode(&dst.AddReplicationAssuranceResultCriteriaRequest)
	if err == nil {
		jsonAddReplicationAssuranceResultCriteriaRequest, _ := json.Marshal(dst.AddReplicationAssuranceResultCriteriaRequest)
		if string(jsonAddReplicationAssuranceResultCriteriaRequest) == "{}" { // empty struct
			dst.AddReplicationAssuranceResultCriteriaRequest = nil
		} else {
			match++
		}
	} else {
		dst.AddReplicationAssuranceResultCriteriaRequest = nil
	}

	// try to unmarshal data into AddSimpleResultCriteriaRequest
	err = newStrictDecoder(data).Decode(&dst.AddSimpleResultCriteriaRequest)
	if err == nil {
		jsonAddSimpleResultCriteriaRequest, _ := json.Marshal(dst.AddSimpleResultCriteriaRequest)
		if string(jsonAddSimpleResultCriteriaRequest) == "{}" { // empty struct
			dst.AddSimpleResultCriteriaRequest = nil
		} else {
			match++
		}
	} else {
		dst.AddSimpleResultCriteriaRequest = nil
	}

	// try to unmarshal data into AddSuccessfulBindResultCriteriaRequest
	err = newStrictDecoder(data).Decode(&dst.AddSuccessfulBindResultCriteriaRequest)
	if err == nil {
		jsonAddSuccessfulBindResultCriteriaRequest, _ := json.Marshal(dst.AddSuccessfulBindResultCriteriaRequest)
		if string(jsonAddSuccessfulBindResultCriteriaRequest) == "{}" { // empty struct
			dst.AddSuccessfulBindResultCriteriaRequest = nil
		} else {
			match++
		}
	} else {
		dst.AddSuccessfulBindResultCriteriaRequest = nil
	}

	// try to unmarshal data into AddThirdPartyResultCriteriaRequest
	err = newStrictDecoder(data).Decode(&dst.AddThirdPartyResultCriteriaRequest)
	if err == nil {
		jsonAddThirdPartyResultCriteriaRequest, _ := json.Marshal(dst.AddThirdPartyResultCriteriaRequest)
		if string(jsonAddThirdPartyResultCriteriaRequest) == "{}" { // empty struct
			dst.AddThirdPartyResultCriteriaRequest = nil
		} else {
			match++
		}
	} else {
		dst.AddThirdPartyResultCriteriaRequest = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.AddAggregateResultCriteriaRequest = nil
		dst.AddReplicationAssuranceResultCriteriaRequest = nil
		dst.AddSimpleResultCriteriaRequest = nil
		dst.AddSuccessfulBindResultCriteriaRequest = nil
		dst.AddThirdPartyResultCriteriaRequest = nil

		return fmt.Errorf("data matches more than one schema in oneOf(AddResultCriteriaRequest)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(AddResultCriteriaRequest)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src AddResultCriteriaRequest) MarshalJSON() ([]byte, error) {
	if src.AddAggregateResultCriteriaRequest != nil {
		return json.Marshal(&src.AddAggregateResultCriteriaRequest)
	}

	if src.AddReplicationAssuranceResultCriteriaRequest != nil {
		return json.Marshal(&src.AddReplicationAssuranceResultCriteriaRequest)
	}

	if src.AddSimpleResultCriteriaRequest != nil {
		return json.Marshal(&src.AddSimpleResultCriteriaRequest)
	}

	if src.AddSuccessfulBindResultCriteriaRequest != nil {
		return json.Marshal(&src.AddSuccessfulBindResultCriteriaRequest)
	}

	if src.AddThirdPartyResultCriteriaRequest != nil {
		return json.Marshal(&src.AddThirdPartyResultCriteriaRequest)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *AddResultCriteriaRequest) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.AddAggregateResultCriteriaRequest != nil {
		return obj.AddAggregateResultCriteriaRequest
	}

	if obj.AddReplicationAssuranceResultCriteriaRequest != nil {
		return obj.AddReplicationAssuranceResultCriteriaRequest
	}

	if obj.AddSimpleResultCriteriaRequest != nil {
		return obj.AddSimpleResultCriteriaRequest
	}

	if obj.AddSuccessfulBindResultCriteriaRequest != nil {
		return obj.AddSuccessfulBindResultCriteriaRequest
	}

	if obj.AddThirdPartyResultCriteriaRequest != nil {
		return obj.AddThirdPartyResultCriteriaRequest
	}

	// all schemas are nil
	return nil
}

type NullableAddResultCriteriaRequest struct {
	value *AddResultCriteriaRequest
	isSet bool
}

func (v NullableAddResultCriteriaRequest) Get() *AddResultCriteriaRequest {
	return v.value
}

func (v *NullableAddResultCriteriaRequest) Set(val *AddResultCriteriaRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddResultCriteriaRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddResultCriteriaRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddResultCriteriaRequest(val *AddResultCriteriaRequest) *NullableAddResultCriteriaRequest {
	return &NullableAddResultCriteriaRequest{value: val, isSet: true}
}

func (v NullableAddResultCriteriaRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddResultCriteriaRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
