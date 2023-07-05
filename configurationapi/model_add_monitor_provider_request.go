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

// AddMonitorProviderRequest - struct for AddMonitorProviderRequest
type AddMonitorProviderRequest struct {
	AddEncryptionSettingsDatabaseAccessibilityMonitorProviderRequest *AddEncryptionSettingsDatabaseAccessibilityMonitorProviderRequest
	AddThirdPartyMonitorProviderRequest                              *AddThirdPartyMonitorProviderRequest
}

// AddEncryptionSettingsDatabaseAccessibilityMonitorProviderRequestAsAddMonitorProviderRequest is a convenience function that returns AddEncryptionSettingsDatabaseAccessibilityMonitorProviderRequest wrapped in AddMonitorProviderRequest
func AddEncryptionSettingsDatabaseAccessibilityMonitorProviderRequestAsAddMonitorProviderRequest(v *AddEncryptionSettingsDatabaseAccessibilityMonitorProviderRequest) AddMonitorProviderRequest {
	return AddMonitorProviderRequest{
		AddEncryptionSettingsDatabaseAccessibilityMonitorProviderRequest: v,
	}
}

// AddThirdPartyMonitorProviderRequestAsAddMonitorProviderRequest is a convenience function that returns AddThirdPartyMonitorProviderRequest wrapped in AddMonitorProviderRequest
func AddThirdPartyMonitorProviderRequestAsAddMonitorProviderRequest(v *AddThirdPartyMonitorProviderRequest) AddMonitorProviderRequest {
	return AddMonitorProviderRequest{
		AddThirdPartyMonitorProviderRequest: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *AddMonitorProviderRequest) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AddEncryptionSettingsDatabaseAccessibilityMonitorProviderRequest
	err = newStrictDecoder(data).Decode(&dst.AddEncryptionSettingsDatabaseAccessibilityMonitorProviderRequest)
	if err == nil {
		jsonAddEncryptionSettingsDatabaseAccessibilityMonitorProviderRequest, _ := json.Marshal(dst.AddEncryptionSettingsDatabaseAccessibilityMonitorProviderRequest)
		if string(jsonAddEncryptionSettingsDatabaseAccessibilityMonitorProviderRequest) == "{}" { // empty struct
			dst.AddEncryptionSettingsDatabaseAccessibilityMonitorProviderRequest = nil
		} else {
			match++
		}
	} else {
		dst.AddEncryptionSettingsDatabaseAccessibilityMonitorProviderRequest = nil
	}

	// try to unmarshal data into AddThirdPartyMonitorProviderRequest
	err = newStrictDecoder(data).Decode(&dst.AddThirdPartyMonitorProviderRequest)
	if err == nil {
		jsonAddThirdPartyMonitorProviderRequest, _ := json.Marshal(dst.AddThirdPartyMonitorProviderRequest)
		if string(jsonAddThirdPartyMonitorProviderRequest) == "{}" { // empty struct
			dst.AddThirdPartyMonitorProviderRequest = nil
		} else {
			match++
		}
	} else {
		dst.AddThirdPartyMonitorProviderRequest = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.AddEncryptionSettingsDatabaseAccessibilityMonitorProviderRequest = nil
		dst.AddThirdPartyMonitorProviderRequest = nil

		return fmt.Errorf("data matches more than one schema in oneOf(AddMonitorProviderRequest)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(AddMonitorProviderRequest)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src AddMonitorProviderRequest) MarshalJSON() ([]byte, error) {
	if src.AddEncryptionSettingsDatabaseAccessibilityMonitorProviderRequest != nil {
		return json.Marshal(&src.AddEncryptionSettingsDatabaseAccessibilityMonitorProviderRequest)
	}

	if src.AddThirdPartyMonitorProviderRequest != nil {
		return json.Marshal(&src.AddThirdPartyMonitorProviderRequest)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *AddMonitorProviderRequest) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.AddEncryptionSettingsDatabaseAccessibilityMonitorProviderRequest != nil {
		return obj.AddEncryptionSettingsDatabaseAccessibilityMonitorProviderRequest
	}

	if obj.AddThirdPartyMonitorProviderRequest != nil {
		return obj.AddThirdPartyMonitorProviderRequest
	}

	// all schemas are nil
	return nil
}

type NullableAddMonitorProviderRequest struct {
	value *AddMonitorProviderRequest
	isSet bool
}

func (v NullableAddMonitorProviderRequest) Get() *AddMonitorProviderRequest {
	return v.value
}

func (v *NullableAddMonitorProviderRequest) Set(val *AddMonitorProviderRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddMonitorProviderRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddMonitorProviderRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddMonitorProviderRequest(val *AddMonitorProviderRequest) *NullableAddMonitorProviderRequest {
	return &NullableAddMonitorProviderRequest{value: val, isSet: true}
}

func (v NullableAddMonitorProviderRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddMonitorProviderRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
