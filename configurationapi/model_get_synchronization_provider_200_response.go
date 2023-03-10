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

// GetSynchronizationProvider200Response - struct for GetSynchronizationProvider200Response
type GetSynchronizationProvider200Response struct {
	CustomSynchronizationProviderResponse      *CustomSynchronizationProviderResponse
	ReplicationSynchronizationProviderResponse *ReplicationSynchronizationProviderResponse
}

// CustomSynchronizationProviderResponseAsGetSynchronizationProvider200Response is a convenience function that returns CustomSynchronizationProviderResponse wrapped in GetSynchronizationProvider200Response
func CustomSynchronizationProviderResponseAsGetSynchronizationProvider200Response(v *CustomSynchronizationProviderResponse) GetSynchronizationProvider200Response {
	return GetSynchronizationProvider200Response{
		CustomSynchronizationProviderResponse: v,
	}
}

// ReplicationSynchronizationProviderResponseAsGetSynchronizationProvider200Response is a convenience function that returns ReplicationSynchronizationProviderResponse wrapped in GetSynchronizationProvider200Response
func ReplicationSynchronizationProviderResponseAsGetSynchronizationProvider200Response(v *ReplicationSynchronizationProviderResponse) GetSynchronizationProvider200Response {
	return GetSynchronizationProvider200Response{
		ReplicationSynchronizationProviderResponse: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *GetSynchronizationProvider200Response) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into CustomSynchronizationProviderResponse
	err = newStrictDecoder(data).Decode(&dst.CustomSynchronizationProviderResponse)
	if err == nil {
		jsonCustomSynchronizationProviderResponse, _ := json.Marshal(dst.CustomSynchronizationProviderResponse)
		if string(jsonCustomSynchronizationProviderResponse) == "{}" { // empty struct
			dst.CustomSynchronizationProviderResponse = nil
		} else {
			match++
		}
	} else {
		dst.CustomSynchronizationProviderResponse = nil
	}

	// try to unmarshal data into ReplicationSynchronizationProviderResponse
	err = newStrictDecoder(data).Decode(&dst.ReplicationSynchronizationProviderResponse)
	if err == nil {
		jsonReplicationSynchronizationProviderResponse, _ := json.Marshal(dst.ReplicationSynchronizationProviderResponse)
		if string(jsonReplicationSynchronizationProviderResponse) == "{}" { // empty struct
			dst.ReplicationSynchronizationProviderResponse = nil
		} else {
			match++
		}
	} else {
		dst.ReplicationSynchronizationProviderResponse = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.CustomSynchronizationProviderResponse = nil
		dst.ReplicationSynchronizationProviderResponse = nil

		return fmt.Errorf("data matches more than one schema in oneOf(GetSynchronizationProvider200Response)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(GetSynchronizationProvider200Response)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GetSynchronizationProvider200Response) MarshalJSON() ([]byte, error) {
	if src.CustomSynchronizationProviderResponse != nil {
		return json.Marshal(&src.CustomSynchronizationProviderResponse)
	}

	if src.ReplicationSynchronizationProviderResponse != nil {
		return json.Marshal(&src.ReplicationSynchronizationProviderResponse)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GetSynchronizationProvider200Response) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.CustomSynchronizationProviderResponse != nil {
		return obj.CustomSynchronizationProviderResponse
	}

	if obj.ReplicationSynchronizationProviderResponse != nil {
		return obj.ReplicationSynchronizationProviderResponse
	}

	// all schemas are nil
	return nil
}

type NullableGetSynchronizationProvider200Response struct {
	value *GetSynchronizationProvider200Response
	isSet bool
}

func (v NullableGetSynchronizationProvider200Response) Get() *GetSynchronizationProvider200Response {
	return v.value
}

func (v *NullableGetSynchronizationProvider200Response) Set(val *GetSynchronizationProvider200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSynchronizationProvider200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSynchronizationProvider200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSynchronizationProvider200Response(val *GetSynchronizationProvider200Response) *NullableGetSynchronizationProvider200Response {
	return &NullableGetSynchronizationProvider200Response{value: val, isSet: true}
}

func (v NullableGetSynchronizationProvider200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSynchronizationProvider200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
