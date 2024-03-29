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

// EntryCacheListResponseResourcesInner - struct for EntryCacheListResponseResourcesInner
type EntryCacheListResponseResourcesInner struct {
	FifoEntryCacheResponse          *FifoEntryCacheResponse
	FileSystemEntryCacheResponse    *FileSystemEntryCacheResponse
	SoftReferenceEntryCacheResponse *SoftReferenceEntryCacheResponse
}

// FifoEntryCacheResponseAsEntryCacheListResponseResourcesInner is a convenience function that returns FifoEntryCacheResponse wrapped in EntryCacheListResponseResourcesInner
func FifoEntryCacheResponseAsEntryCacheListResponseResourcesInner(v *FifoEntryCacheResponse) EntryCacheListResponseResourcesInner {
	return EntryCacheListResponseResourcesInner{
		FifoEntryCacheResponse: v,
	}
}

// FileSystemEntryCacheResponseAsEntryCacheListResponseResourcesInner is a convenience function that returns FileSystemEntryCacheResponse wrapped in EntryCacheListResponseResourcesInner
func FileSystemEntryCacheResponseAsEntryCacheListResponseResourcesInner(v *FileSystemEntryCacheResponse) EntryCacheListResponseResourcesInner {
	return EntryCacheListResponseResourcesInner{
		FileSystemEntryCacheResponse: v,
	}
}

// SoftReferenceEntryCacheResponseAsEntryCacheListResponseResourcesInner is a convenience function that returns SoftReferenceEntryCacheResponse wrapped in EntryCacheListResponseResourcesInner
func SoftReferenceEntryCacheResponseAsEntryCacheListResponseResourcesInner(v *SoftReferenceEntryCacheResponse) EntryCacheListResponseResourcesInner {
	return EntryCacheListResponseResourcesInner{
		SoftReferenceEntryCacheResponse: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *EntryCacheListResponseResourcesInner) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into FifoEntryCacheResponse
	err = newStrictDecoder(data).Decode(&dst.FifoEntryCacheResponse)
	if err == nil {
		jsonFifoEntryCacheResponse, _ := json.Marshal(dst.FifoEntryCacheResponse)
		if string(jsonFifoEntryCacheResponse) == "{}" { // empty struct
			dst.FifoEntryCacheResponse = nil
		} else {
			match++
		}
	} else {
		dst.FifoEntryCacheResponse = nil
	}

	// try to unmarshal data into FileSystemEntryCacheResponse
	err = newStrictDecoder(data).Decode(&dst.FileSystemEntryCacheResponse)
	if err == nil {
		jsonFileSystemEntryCacheResponse, _ := json.Marshal(dst.FileSystemEntryCacheResponse)
		if string(jsonFileSystemEntryCacheResponse) == "{}" { // empty struct
			dst.FileSystemEntryCacheResponse = nil
		} else {
			match++
		}
	} else {
		dst.FileSystemEntryCacheResponse = nil
	}

	// try to unmarshal data into SoftReferenceEntryCacheResponse
	err = newStrictDecoder(data).Decode(&dst.SoftReferenceEntryCacheResponse)
	if err == nil {
		jsonSoftReferenceEntryCacheResponse, _ := json.Marshal(dst.SoftReferenceEntryCacheResponse)
		if string(jsonSoftReferenceEntryCacheResponse) == "{}" { // empty struct
			dst.SoftReferenceEntryCacheResponse = nil
		} else {
			match++
		}
	} else {
		dst.SoftReferenceEntryCacheResponse = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.FifoEntryCacheResponse = nil
		dst.FileSystemEntryCacheResponse = nil
		dst.SoftReferenceEntryCacheResponse = nil

		return fmt.Errorf("data matches more than one schema in oneOf(EntryCacheListResponseResourcesInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(EntryCacheListResponseResourcesInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src EntryCacheListResponseResourcesInner) MarshalJSON() ([]byte, error) {
	if src.FifoEntryCacheResponse != nil {
		return json.Marshal(&src.FifoEntryCacheResponse)
	}

	if src.FileSystemEntryCacheResponse != nil {
		return json.Marshal(&src.FileSystemEntryCacheResponse)
	}

	if src.SoftReferenceEntryCacheResponse != nil {
		return json.Marshal(&src.SoftReferenceEntryCacheResponse)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *EntryCacheListResponseResourcesInner) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.FifoEntryCacheResponse != nil {
		return obj.FifoEntryCacheResponse
	}

	if obj.FileSystemEntryCacheResponse != nil {
		return obj.FileSystemEntryCacheResponse
	}

	if obj.SoftReferenceEntryCacheResponse != nil {
		return obj.SoftReferenceEntryCacheResponse
	}

	// all schemas are nil
	return nil
}

type NullableEntryCacheListResponseResourcesInner struct {
	value *EntryCacheListResponseResourcesInner
	isSet bool
}

func (v NullableEntryCacheListResponseResourcesInner) Get() *EntryCacheListResponseResourcesInner {
	return v.value
}

func (v *NullableEntryCacheListResponseResourcesInner) Set(val *EntryCacheListResponseResourcesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableEntryCacheListResponseResourcesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableEntryCacheListResponseResourcesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntryCacheListResponseResourcesInner(val *EntryCacheListResponseResourcesInner) *NullableEntryCacheListResponseResourcesInner {
	return &NullableEntryCacheListResponseResourcesInner{value: val, isSet: true}
}

func (v NullableEntryCacheListResponseResourcesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntryCacheListResponseResourcesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
