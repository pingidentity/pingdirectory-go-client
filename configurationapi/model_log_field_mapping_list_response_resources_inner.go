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

// LogFieldMappingListResponseResourcesInner - struct for LogFieldMappingListResponseResourcesInner
type LogFieldMappingListResponseResourcesInner struct {
	AccessLogFieldMappingResponse *AccessLogFieldMappingResponse
	ErrorLogFieldMappingResponse  *ErrorLogFieldMappingResponse
}

// AccessLogFieldMappingResponseAsLogFieldMappingListResponseResourcesInner is a convenience function that returns AccessLogFieldMappingResponse wrapped in LogFieldMappingListResponseResourcesInner
func AccessLogFieldMappingResponseAsLogFieldMappingListResponseResourcesInner(v *AccessLogFieldMappingResponse) LogFieldMappingListResponseResourcesInner {
	return LogFieldMappingListResponseResourcesInner{
		AccessLogFieldMappingResponse: v,
	}
}

// ErrorLogFieldMappingResponseAsLogFieldMappingListResponseResourcesInner is a convenience function that returns ErrorLogFieldMappingResponse wrapped in LogFieldMappingListResponseResourcesInner
func ErrorLogFieldMappingResponseAsLogFieldMappingListResponseResourcesInner(v *ErrorLogFieldMappingResponse) LogFieldMappingListResponseResourcesInner {
	return LogFieldMappingListResponseResourcesInner{
		ErrorLogFieldMappingResponse: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *LogFieldMappingListResponseResourcesInner) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AccessLogFieldMappingResponse
	err = newStrictDecoder(data).Decode(&dst.AccessLogFieldMappingResponse)
	if err == nil {
		jsonAccessLogFieldMappingResponse, _ := json.Marshal(dst.AccessLogFieldMappingResponse)
		if string(jsonAccessLogFieldMappingResponse) == "{}" { // empty struct
			dst.AccessLogFieldMappingResponse = nil
		} else {
			match++
		}
	} else {
		dst.AccessLogFieldMappingResponse = nil
	}

	// try to unmarshal data into ErrorLogFieldMappingResponse
	err = newStrictDecoder(data).Decode(&dst.ErrorLogFieldMappingResponse)
	if err == nil {
		jsonErrorLogFieldMappingResponse, _ := json.Marshal(dst.ErrorLogFieldMappingResponse)
		if string(jsonErrorLogFieldMappingResponse) == "{}" { // empty struct
			dst.ErrorLogFieldMappingResponse = nil
		} else {
			match++
		}
	} else {
		dst.ErrorLogFieldMappingResponse = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.AccessLogFieldMappingResponse = nil
		dst.ErrorLogFieldMappingResponse = nil

		return fmt.Errorf("data matches more than one schema in oneOf(LogFieldMappingListResponseResourcesInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(LogFieldMappingListResponseResourcesInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src LogFieldMappingListResponseResourcesInner) MarshalJSON() ([]byte, error) {
	if src.AccessLogFieldMappingResponse != nil {
		return json.Marshal(&src.AccessLogFieldMappingResponse)
	}

	if src.ErrorLogFieldMappingResponse != nil {
		return json.Marshal(&src.ErrorLogFieldMappingResponse)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *LogFieldMappingListResponseResourcesInner) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.AccessLogFieldMappingResponse != nil {
		return obj.AccessLogFieldMappingResponse
	}

	if obj.ErrorLogFieldMappingResponse != nil {
		return obj.ErrorLogFieldMappingResponse
	}

	// all schemas are nil
	return nil
}

type NullableLogFieldMappingListResponseResourcesInner struct {
	value *LogFieldMappingListResponseResourcesInner
	isSet bool
}

func (v NullableLogFieldMappingListResponseResourcesInner) Get() *LogFieldMappingListResponseResourcesInner {
	return v.value
}

func (v *NullableLogFieldMappingListResponseResourcesInner) Set(val *LogFieldMappingListResponseResourcesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableLogFieldMappingListResponseResourcesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableLogFieldMappingListResponseResourcesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogFieldMappingListResponseResourcesInner(val *LogFieldMappingListResponseResourcesInner) *NullableLogFieldMappingListResponseResourcesInner {
	return &NullableLogFieldMappingListResponseResourcesInner{value: val, isSet: true}
}

func (v NullableLogFieldMappingListResponseResourcesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogFieldMappingListResponseResourcesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
