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

// AddGaugeDataSourceRequest - struct for AddGaugeDataSourceRequest
type AddGaugeDataSourceRequest struct {
	AddIndicatorGaugeDataSourceRequest *AddIndicatorGaugeDataSourceRequest
	AddNumericGaugeDataSourceRequest *AddNumericGaugeDataSourceRequest
}

// AddIndicatorGaugeDataSourceRequestAsAddGaugeDataSourceRequest is a convenience function that returns AddIndicatorGaugeDataSourceRequest wrapped in AddGaugeDataSourceRequest
func AddIndicatorGaugeDataSourceRequestAsAddGaugeDataSourceRequest(v *AddIndicatorGaugeDataSourceRequest) AddGaugeDataSourceRequest {
	return AddGaugeDataSourceRequest{
		AddIndicatorGaugeDataSourceRequest: v,
	}
}

// AddNumericGaugeDataSourceRequestAsAddGaugeDataSourceRequest is a convenience function that returns AddNumericGaugeDataSourceRequest wrapped in AddGaugeDataSourceRequest
func AddNumericGaugeDataSourceRequestAsAddGaugeDataSourceRequest(v *AddNumericGaugeDataSourceRequest) AddGaugeDataSourceRequest {
	return AddGaugeDataSourceRequest{
		AddNumericGaugeDataSourceRequest: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *AddGaugeDataSourceRequest) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AddIndicatorGaugeDataSourceRequest
	err = newStrictDecoder(data).Decode(&dst.AddIndicatorGaugeDataSourceRequest)
	if err == nil {
		jsonAddIndicatorGaugeDataSourceRequest, _ := json.Marshal(dst.AddIndicatorGaugeDataSourceRequest)
		if string(jsonAddIndicatorGaugeDataSourceRequest) == "{}" { // empty struct
			dst.AddIndicatorGaugeDataSourceRequest = nil
		} else {
			match++
		}
	} else {
		dst.AddIndicatorGaugeDataSourceRequest = nil
	}

	// try to unmarshal data into AddNumericGaugeDataSourceRequest
	err = newStrictDecoder(data).Decode(&dst.AddNumericGaugeDataSourceRequest)
	if err == nil {
		jsonAddNumericGaugeDataSourceRequest, _ := json.Marshal(dst.AddNumericGaugeDataSourceRequest)
		if string(jsonAddNumericGaugeDataSourceRequest) == "{}" { // empty struct
			dst.AddNumericGaugeDataSourceRequest = nil
		} else {
			match++
		}
	} else {
		dst.AddNumericGaugeDataSourceRequest = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.AddIndicatorGaugeDataSourceRequest = nil
		dst.AddNumericGaugeDataSourceRequest = nil

		return fmt.Errorf("data matches more than one schema in oneOf(AddGaugeDataSourceRequest)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(AddGaugeDataSourceRequest)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src AddGaugeDataSourceRequest) MarshalJSON() ([]byte, error) {
	if src.AddIndicatorGaugeDataSourceRequest != nil {
		return json.Marshal(&src.AddIndicatorGaugeDataSourceRequest)
	}

	if src.AddNumericGaugeDataSourceRequest != nil {
		return json.Marshal(&src.AddNumericGaugeDataSourceRequest)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *AddGaugeDataSourceRequest) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.AddIndicatorGaugeDataSourceRequest != nil {
		return obj.AddIndicatorGaugeDataSourceRequest
	}

	if obj.AddNumericGaugeDataSourceRequest != nil {
		return obj.AddNumericGaugeDataSourceRequest
	}

	// all schemas are nil
	return nil
}

type NullableAddGaugeDataSourceRequest struct {
	value *AddGaugeDataSourceRequest
	isSet bool
}

func (v NullableAddGaugeDataSourceRequest) Get() *AddGaugeDataSourceRequest {
	return v.value
}

func (v *NullableAddGaugeDataSourceRequest) Set(val *AddGaugeDataSourceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddGaugeDataSourceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddGaugeDataSourceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddGaugeDataSourceRequest(val *AddGaugeDataSourceRequest) *NullableAddGaugeDataSourceRequest {
	return &NullableAddGaugeDataSourceRequest{value: val, isSet: true}
}

func (v NullableAddGaugeDataSourceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddGaugeDataSourceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

