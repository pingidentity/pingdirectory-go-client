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

// AddAzureAuthenticationMethodRequest - struct for AddAzureAuthenticationMethodRequest
type AddAzureAuthenticationMethodRequest struct {
	AddClientSecretAzureAuthenticationMethodRequest     *AddClientSecretAzureAuthenticationMethodRequest
	AddDefaultAzureAuthenticationMethodRequest          *AddDefaultAzureAuthenticationMethodRequest
	AddUsernamePasswordAzureAuthenticationMethodRequest *AddUsernamePasswordAzureAuthenticationMethodRequest
}

// AddClientSecretAzureAuthenticationMethodRequestAsAddAzureAuthenticationMethodRequest is a convenience function that returns AddClientSecretAzureAuthenticationMethodRequest wrapped in AddAzureAuthenticationMethodRequest
func AddClientSecretAzureAuthenticationMethodRequestAsAddAzureAuthenticationMethodRequest(v *AddClientSecretAzureAuthenticationMethodRequest) AddAzureAuthenticationMethodRequest {
	return AddAzureAuthenticationMethodRequest{
		AddClientSecretAzureAuthenticationMethodRequest: v,
	}
}

// AddDefaultAzureAuthenticationMethodRequestAsAddAzureAuthenticationMethodRequest is a convenience function that returns AddDefaultAzureAuthenticationMethodRequest wrapped in AddAzureAuthenticationMethodRequest
func AddDefaultAzureAuthenticationMethodRequestAsAddAzureAuthenticationMethodRequest(v *AddDefaultAzureAuthenticationMethodRequest) AddAzureAuthenticationMethodRequest {
	return AddAzureAuthenticationMethodRequest{
		AddDefaultAzureAuthenticationMethodRequest: v,
	}
}

// AddUsernamePasswordAzureAuthenticationMethodRequestAsAddAzureAuthenticationMethodRequest is a convenience function that returns AddUsernamePasswordAzureAuthenticationMethodRequest wrapped in AddAzureAuthenticationMethodRequest
func AddUsernamePasswordAzureAuthenticationMethodRequestAsAddAzureAuthenticationMethodRequest(v *AddUsernamePasswordAzureAuthenticationMethodRequest) AddAzureAuthenticationMethodRequest {
	return AddAzureAuthenticationMethodRequest{
		AddUsernamePasswordAzureAuthenticationMethodRequest: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *AddAzureAuthenticationMethodRequest) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AddClientSecretAzureAuthenticationMethodRequest
	err = newStrictDecoder(data).Decode(&dst.AddClientSecretAzureAuthenticationMethodRequest)
	if err == nil {
		jsonAddClientSecretAzureAuthenticationMethodRequest, _ := json.Marshal(dst.AddClientSecretAzureAuthenticationMethodRequest)
		if string(jsonAddClientSecretAzureAuthenticationMethodRequest) == "{}" { // empty struct
			dst.AddClientSecretAzureAuthenticationMethodRequest = nil
		} else {
			match++
		}
	} else {
		dst.AddClientSecretAzureAuthenticationMethodRequest = nil
	}

	// try to unmarshal data into AddDefaultAzureAuthenticationMethodRequest
	err = newStrictDecoder(data).Decode(&dst.AddDefaultAzureAuthenticationMethodRequest)
	if err == nil {
		jsonAddDefaultAzureAuthenticationMethodRequest, _ := json.Marshal(dst.AddDefaultAzureAuthenticationMethodRequest)
		if string(jsonAddDefaultAzureAuthenticationMethodRequest) == "{}" { // empty struct
			dst.AddDefaultAzureAuthenticationMethodRequest = nil
		} else {
			match++
		}
	} else {
		dst.AddDefaultAzureAuthenticationMethodRequest = nil
	}

	// try to unmarshal data into AddUsernamePasswordAzureAuthenticationMethodRequest
	err = newStrictDecoder(data).Decode(&dst.AddUsernamePasswordAzureAuthenticationMethodRequest)
	if err == nil {
		jsonAddUsernamePasswordAzureAuthenticationMethodRequest, _ := json.Marshal(dst.AddUsernamePasswordAzureAuthenticationMethodRequest)
		if string(jsonAddUsernamePasswordAzureAuthenticationMethodRequest) == "{}" { // empty struct
			dst.AddUsernamePasswordAzureAuthenticationMethodRequest = nil
		} else {
			match++
		}
	} else {
		dst.AddUsernamePasswordAzureAuthenticationMethodRequest = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.AddClientSecretAzureAuthenticationMethodRequest = nil
		dst.AddDefaultAzureAuthenticationMethodRequest = nil
		dst.AddUsernamePasswordAzureAuthenticationMethodRequest = nil

		return fmt.Errorf("data matches more than one schema in oneOf(AddAzureAuthenticationMethodRequest)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(AddAzureAuthenticationMethodRequest)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src AddAzureAuthenticationMethodRequest) MarshalJSON() ([]byte, error) {
	if src.AddClientSecretAzureAuthenticationMethodRequest != nil {
		return json.Marshal(&src.AddClientSecretAzureAuthenticationMethodRequest)
	}

	if src.AddDefaultAzureAuthenticationMethodRequest != nil {
		return json.Marshal(&src.AddDefaultAzureAuthenticationMethodRequest)
	}

	if src.AddUsernamePasswordAzureAuthenticationMethodRequest != nil {
		return json.Marshal(&src.AddUsernamePasswordAzureAuthenticationMethodRequest)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *AddAzureAuthenticationMethodRequest) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.AddClientSecretAzureAuthenticationMethodRequest != nil {
		return obj.AddClientSecretAzureAuthenticationMethodRequest
	}

	if obj.AddDefaultAzureAuthenticationMethodRequest != nil {
		return obj.AddDefaultAzureAuthenticationMethodRequest
	}

	if obj.AddUsernamePasswordAzureAuthenticationMethodRequest != nil {
		return obj.AddUsernamePasswordAzureAuthenticationMethodRequest
	}

	// all schemas are nil
	return nil
}

type NullableAddAzureAuthenticationMethodRequest struct {
	value *AddAzureAuthenticationMethodRequest
	isSet bool
}

func (v NullableAddAzureAuthenticationMethodRequest) Get() *AddAzureAuthenticationMethodRequest {
	return v.value
}

func (v *NullableAddAzureAuthenticationMethodRequest) Set(val *AddAzureAuthenticationMethodRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddAzureAuthenticationMethodRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddAzureAuthenticationMethodRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddAzureAuthenticationMethodRequest(val *AddAzureAuthenticationMethodRequest) *NullableAddAzureAuthenticationMethodRequest {
	return &NullableAddAzureAuthenticationMethodRequest{value: val, isSet: true}
}

func (v NullableAddAzureAuthenticationMethodRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddAzureAuthenticationMethodRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
