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

// AddGaugeDataSource200Response - struct for AddGaugeDataSource200Response
type AddGaugeDataSource200Response struct {
	IndicatorGaugeDataSourceResponse *IndicatorGaugeDataSourceResponse
	NumericGaugeDataSourceResponse   *NumericGaugeDataSourceResponse
}

// IndicatorGaugeDataSourceResponseAsAddGaugeDataSource200Response is a convenience function that returns IndicatorGaugeDataSourceResponse wrapped in AddGaugeDataSource200Response
func IndicatorGaugeDataSourceResponseAsAddGaugeDataSource200Response(v *IndicatorGaugeDataSourceResponse) AddGaugeDataSource200Response {
	return AddGaugeDataSource200Response{
		IndicatorGaugeDataSourceResponse: v,
	}
}

// NumericGaugeDataSourceResponseAsAddGaugeDataSource200Response is a convenience function that returns NumericGaugeDataSourceResponse wrapped in AddGaugeDataSource200Response
func NumericGaugeDataSourceResponseAsAddGaugeDataSource200Response(v *NumericGaugeDataSourceResponse) AddGaugeDataSource200Response {
	return AddGaugeDataSource200Response{
		NumericGaugeDataSourceResponse: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *AddGaugeDataSource200Response) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into IndicatorGaugeDataSourceResponse
	err = newStrictDecoder(data).Decode(&dst.IndicatorGaugeDataSourceResponse)
	if err == nil {
		jsonIndicatorGaugeDataSourceResponse, _ := json.Marshal(dst.IndicatorGaugeDataSourceResponse)
		if string(jsonIndicatorGaugeDataSourceResponse) == "{}" { // empty struct
			dst.IndicatorGaugeDataSourceResponse = nil
		} else {
			match++
		}
	} else {
		dst.IndicatorGaugeDataSourceResponse = nil
	}

	// try to unmarshal data into NumericGaugeDataSourceResponse
	err = newStrictDecoder(data).Decode(&dst.NumericGaugeDataSourceResponse)
	if err == nil {
		jsonNumericGaugeDataSourceResponse, _ := json.Marshal(dst.NumericGaugeDataSourceResponse)
		if string(jsonNumericGaugeDataSourceResponse) == "{}" { // empty struct
			dst.NumericGaugeDataSourceResponse = nil
		} else {
			match++
		}
	} else {
		dst.NumericGaugeDataSourceResponse = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.IndicatorGaugeDataSourceResponse = nil
		dst.NumericGaugeDataSourceResponse = nil

		return fmt.Errorf("data matches more than one schema in oneOf(AddGaugeDataSource200Response)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(AddGaugeDataSource200Response)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src AddGaugeDataSource200Response) MarshalJSON() ([]byte, error) {
	if src.IndicatorGaugeDataSourceResponse != nil {
		return json.Marshal(&src.IndicatorGaugeDataSourceResponse)
	}

	if src.NumericGaugeDataSourceResponse != nil {
		return json.Marshal(&src.NumericGaugeDataSourceResponse)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *AddGaugeDataSource200Response) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.IndicatorGaugeDataSourceResponse != nil {
		return obj.IndicatorGaugeDataSourceResponse
	}

	if obj.NumericGaugeDataSourceResponse != nil {
		return obj.NumericGaugeDataSourceResponse
	}

	// all schemas are nil
	return nil
}

type NullableAddGaugeDataSource200Response struct {
	value *AddGaugeDataSource200Response
	isSet bool
}

func (v NullableAddGaugeDataSource200Response) Get() *AddGaugeDataSource200Response {
	return v.value
}

func (v *NullableAddGaugeDataSource200Response) Set(val *AddGaugeDataSource200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableAddGaugeDataSource200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableAddGaugeDataSource200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddGaugeDataSource200Response(val *AddGaugeDataSource200Response) *NullableAddGaugeDataSource200Response {
	return &NullableAddGaugeDataSource200Response{value: val, isSet: true}
}

func (v NullableAddGaugeDataSource200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddGaugeDataSource200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
