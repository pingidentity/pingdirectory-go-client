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

// ResultCriteriaListResponseResourcesInner - struct for ResultCriteriaListResponseResourcesInner
type ResultCriteriaListResponseResourcesInner struct {
	AggregateResultCriteriaResponse            *AggregateResultCriteriaResponse
	ReplicationAssuranceResultCriteriaResponse *ReplicationAssuranceResultCriteriaResponse
	SimpleResultCriteriaResponse               *SimpleResultCriteriaResponse
	SuccessfulBindResultCriteriaResponse       *SuccessfulBindResultCriteriaResponse
	ThirdPartyResultCriteriaResponse           *ThirdPartyResultCriteriaResponse
}

// AggregateResultCriteriaResponseAsResultCriteriaListResponseResourcesInner is a convenience function that returns AggregateResultCriteriaResponse wrapped in ResultCriteriaListResponseResourcesInner
func AggregateResultCriteriaResponseAsResultCriteriaListResponseResourcesInner(v *AggregateResultCriteriaResponse) ResultCriteriaListResponseResourcesInner {
	return ResultCriteriaListResponseResourcesInner{
		AggregateResultCriteriaResponse: v,
	}
}

// ReplicationAssuranceResultCriteriaResponseAsResultCriteriaListResponseResourcesInner is a convenience function that returns ReplicationAssuranceResultCriteriaResponse wrapped in ResultCriteriaListResponseResourcesInner
func ReplicationAssuranceResultCriteriaResponseAsResultCriteriaListResponseResourcesInner(v *ReplicationAssuranceResultCriteriaResponse) ResultCriteriaListResponseResourcesInner {
	return ResultCriteriaListResponseResourcesInner{
		ReplicationAssuranceResultCriteriaResponse: v,
	}
}

// SimpleResultCriteriaResponseAsResultCriteriaListResponseResourcesInner is a convenience function that returns SimpleResultCriteriaResponse wrapped in ResultCriteriaListResponseResourcesInner
func SimpleResultCriteriaResponseAsResultCriteriaListResponseResourcesInner(v *SimpleResultCriteriaResponse) ResultCriteriaListResponseResourcesInner {
	return ResultCriteriaListResponseResourcesInner{
		SimpleResultCriteriaResponse: v,
	}
}

// SuccessfulBindResultCriteriaResponseAsResultCriteriaListResponseResourcesInner is a convenience function that returns SuccessfulBindResultCriteriaResponse wrapped in ResultCriteriaListResponseResourcesInner
func SuccessfulBindResultCriteriaResponseAsResultCriteriaListResponseResourcesInner(v *SuccessfulBindResultCriteriaResponse) ResultCriteriaListResponseResourcesInner {
	return ResultCriteriaListResponseResourcesInner{
		SuccessfulBindResultCriteriaResponse: v,
	}
}

// ThirdPartyResultCriteriaResponseAsResultCriteriaListResponseResourcesInner is a convenience function that returns ThirdPartyResultCriteriaResponse wrapped in ResultCriteriaListResponseResourcesInner
func ThirdPartyResultCriteriaResponseAsResultCriteriaListResponseResourcesInner(v *ThirdPartyResultCriteriaResponse) ResultCriteriaListResponseResourcesInner {
	return ResultCriteriaListResponseResourcesInner{
		ThirdPartyResultCriteriaResponse: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *ResultCriteriaListResponseResourcesInner) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AggregateResultCriteriaResponse
	err = newStrictDecoder(data).Decode(&dst.AggregateResultCriteriaResponse)
	if err == nil {
		jsonAggregateResultCriteriaResponse, _ := json.Marshal(dst.AggregateResultCriteriaResponse)
		if string(jsonAggregateResultCriteriaResponse) == "{}" { // empty struct
			dst.AggregateResultCriteriaResponse = nil
		} else {
			match++
		}
	} else {
		dst.AggregateResultCriteriaResponse = nil
	}

	// try to unmarshal data into ReplicationAssuranceResultCriteriaResponse
	err = newStrictDecoder(data).Decode(&dst.ReplicationAssuranceResultCriteriaResponse)
	if err == nil {
		jsonReplicationAssuranceResultCriteriaResponse, _ := json.Marshal(dst.ReplicationAssuranceResultCriteriaResponse)
		if string(jsonReplicationAssuranceResultCriteriaResponse) == "{}" { // empty struct
			dst.ReplicationAssuranceResultCriteriaResponse = nil
		} else {
			match++
		}
	} else {
		dst.ReplicationAssuranceResultCriteriaResponse = nil
	}

	// try to unmarshal data into SimpleResultCriteriaResponse
	err = newStrictDecoder(data).Decode(&dst.SimpleResultCriteriaResponse)
	if err == nil {
		jsonSimpleResultCriteriaResponse, _ := json.Marshal(dst.SimpleResultCriteriaResponse)
		if string(jsonSimpleResultCriteriaResponse) == "{}" { // empty struct
			dst.SimpleResultCriteriaResponse = nil
		} else {
			match++
		}
	} else {
		dst.SimpleResultCriteriaResponse = nil
	}

	// try to unmarshal data into SuccessfulBindResultCriteriaResponse
	err = newStrictDecoder(data).Decode(&dst.SuccessfulBindResultCriteriaResponse)
	if err == nil {
		jsonSuccessfulBindResultCriteriaResponse, _ := json.Marshal(dst.SuccessfulBindResultCriteriaResponse)
		if string(jsonSuccessfulBindResultCriteriaResponse) == "{}" { // empty struct
			dst.SuccessfulBindResultCriteriaResponse = nil
		} else {
			match++
		}
	} else {
		dst.SuccessfulBindResultCriteriaResponse = nil
	}

	// try to unmarshal data into ThirdPartyResultCriteriaResponse
	err = newStrictDecoder(data).Decode(&dst.ThirdPartyResultCriteriaResponse)
	if err == nil {
		jsonThirdPartyResultCriteriaResponse, _ := json.Marshal(dst.ThirdPartyResultCriteriaResponse)
		if string(jsonThirdPartyResultCriteriaResponse) == "{}" { // empty struct
			dst.ThirdPartyResultCriteriaResponse = nil
		} else {
			match++
		}
	} else {
		dst.ThirdPartyResultCriteriaResponse = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.AggregateResultCriteriaResponse = nil
		dst.ReplicationAssuranceResultCriteriaResponse = nil
		dst.SimpleResultCriteriaResponse = nil
		dst.SuccessfulBindResultCriteriaResponse = nil
		dst.ThirdPartyResultCriteriaResponse = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ResultCriteriaListResponseResourcesInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ResultCriteriaListResponseResourcesInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ResultCriteriaListResponseResourcesInner) MarshalJSON() ([]byte, error) {
	if src.AggregateResultCriteriaResponse != nil {
		return json.Marshal(&src.AggregateResultCriteriaResponse)
	}

	if src.ReplicationAssuranceResultCriteriaResponse != nil {
		return json.Marshal(&src.ReplicationAssuranceResultCriteriaResponse)
	}

	if src.SimpleResultCriteriaResponse != nil {
		return json.Marshal(&src.SimpleResultCriteriaResponse)
	}

	if src.SuccessfulBindResultCriteriaResponse != nil {
		return json.Marshal(&src.SuccessfulBindResultCriteriaResponse)
	}

	if src.ThirdPartyResultCriteriaResponse != nil {
		return json.Marshal(&src.ThirdPartyResultCriteriaResponse)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ResultCriteriaListResponseResourcesInner) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.AggregateResultCriteriaResponse != nil {
		return obj.AggregateResultCriteriaResponse
	}

	if obj.ReplicationAssuranceResultCriteriaResponse != nil {
		return obj.ReplicationAssuranceResultCriteriaResponse
	}

	if obj.SimpleResultCriteriaResponse != nil {
		return obj.SimpleResultCriteriaResponse
	}

	if obj.SuccessfulBindResultCriteriaResponse != nil {
		return obj.SuccessfulBindResultCriteriaResponse
	}

	if obj.ThirdPartyResultCriteriaResponse != nil {
		return obj.ThirdPartyResultCriteriaResponse
	}

	// all schemas are nil
	return nil
}

type NullableResultCriteriaListResponseResourcesInner struct {
	value *ResultCriteriaListResponseResourcesInner
	isSet bool
}

func (v NullableResultCriteriaListResponseResourcesInner) Get() *ResultCriteriaListResponseResourcesInner {
	return v.value
}

func (v *NullableResultCriteriaListResponseResourcesInner) Set(val *ResultCriteriaListResponseResourcesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableResultCriteriaListResponseResourcesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableResultCriteriaListResponseResourcesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResultCriteriaListResponseResourcesInner(val *ResultCriteriaListResponseResourcesInner) *NullableResultCriteriaListResponseResourcesInner {
	return &NullableResultCriteriaListResponseResourcesInner{value: val, isSet: true}
}

func (v NullableResultCriteriaListResponseResourcesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResultCriteriaListResponseResourcesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
