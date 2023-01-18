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

// AddAccessTokenValidatorRequest - struct for AddAccessTokenValidatorRequest
type AddAccessTokenValidatorRequest struct {
	AddJwtAccessTokenValidatorRequest *AddJwtAccessTokenValidatorRequest
	AddMockAccessTokenValidatorRequest *AddMockAccessTokenValidatorRequest
	AddPingFederateAccessTokenValidatorRequest *AddPingFederateAccessTokenValidatorRequest
	AddThirdPartyAccessTokenValidatorRequest *AddThirdPartyAccessTokenValidatorRequest
}

// AddJwtAccessTokenValidatorRequestAsAddAccessTokenValidatorRequest is a convenience function that returns AddJwtAccessTokenValidatorRequest wrapped in AddAccessTokenValidatorRequest
func AddJwtAccessTokenValidatorRequestAsAddAccessTokenValidatorRequest(v *AddJwtAccessTokenValidatorRequest) AddAccessTokenValidatorRequest {
	return AddAccessTokenValidatorRequest{
		AddJwtAccessTokenValidatorRequest: v,
	}
}

// AddMockAccessTokenValidatorRequestAsAddAccessTokenValidatorRequest is a convenience function that returns AddMockAccessTokenValidatorRequest wrapped in AddAccessTokenValidatorRequest
func AddMockAccessTokenValidatorRequestAsAddAccessTokenValidatorRequest(v *AddMockAccessTokenValidatorRequest) AddAccessTokenValidatorRequest {
	return AddAccessTokenValidatorRequest{
		AddMockAccessTokenValidatorRequest: v,
	}
}

// AddPingFederateAccessTokenValidatorRequestAsAddAccessTokenValidatorRequest is a convenience function that returns AddPingFederateAccessTokenValidatorRequest wrapped in AddAccessTokenValidatorRequest
func AddPingFederateAccessTokenValidatorRequestAsAddAccessTokenValidatorRequest(v *AddPingFederateAccessTokenValidatorRequest) AddAccessTokenValidatorRequest {
	return AddAccessTokenValidatorRequest{
		AddPingFederateAccessTokenValidatorRequest: v,
	}
}

// AddThirdPartyAccessTokenValidatorRequestAsAddAccessTokenValidatorRequest is a convenience function that returns AddThirdPartyAccessTokenValidatorRequest wrapped in AddAccessTokenValidatorRequest
func AddThirdPartyAccessTokenValidatorRequestAsAddAccessTokenValidatorRequest(v *AddThirdPartyAccessTokenValidatorRequest) AddAccessTokenValidatorRequest {
	return AddAccessTokenValidatorRequest{
		AddThirdPartyAccessTokenValidatorRequest: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *AddAccessTokenValidatorRequest) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AddJwtAccessTokenValidatorRequest
	err = newStrictDecoder(data).Decode(&dst.AddJwtAccessTokenValidatorRequest)
	if err == nil {
		jsonAddJwtAccessTokenValidatorRequest, _ := json.Marshal(dst.AddJwtAccessTokenValidatorRequest)
		if string(jsonAddJwtAccessTokenValidatorRequest) == "{}" { // empty struct
			dst.AddJwtAccessTokenValidatorRequest = nil
		} else {
			match++
		}
	} else {
		dst.AddJwtAccessTokenValidatorRequest = nil
	}

	// try to unmarshal data into AddMockAccessTokenValidatorRequest
	err = newStrictDecoder(data).Decode(&dst.AddMockAccessTokenValidatorRequest)
	if err == nil {
		jsonAddMockAccessTokenValidatorRequest, _ := json.Marshal(dst.AddMockAccessTokenValidatorRequest)
		if string(jsonAddMockAccessTokenValidatorRequest) == "{}" { // empty struct
			dst.AddMockAccessTokenValidatorRequest = nil
		} else {
			match++
		}
	} else {
		dst.AddMockAccessTokenValidatorRequest = nil
	}

	// try to unmarshal data into AddPingFederateAccessTokenValidatorRequest
	err = newStrictDecoder(data).Decode(&dst.AddPingFederateAccessTokenValidatorRequest)
	if err == nil {
		jsonAddPingFederateAccessTokenValidatorRequest, _ := json.Marshal(dst.AddPingFederateAccessTokenValidatorRequest)
		if string(jsonAddPingFederateAccessTokenValidatorRequest) == "{}" { // empty struct
			dst.AddPingFederateAccessTokenValidatorRequest = nil
		} else {
			match++
		}
	} else {
		dst.AddPingFederateAccessTokenValidatorRequest = nil
	}

	// try to unmarshal data into AddThirdPartyAccessTokenValidatorRequest
	err = newStrictDecoder(data).Decode(&dst.AddThirdPartyAccessTokenValidatorRequest)
	if err == nil {
		jsonAddThirdPartyAccessTokenValidatorRequest, _ := json.Marshal(dst.AddThirdPartyAccessTokenValidatorRequest)
		if string(jsonAddThirdPartyAccessTokenValidatorRequest) == "{}" { // empty struct
			dst.AddThirdPartyAccessTokenValidatorRequest = nil
		} else {
			match++
		}
	} else {
		dst.AddThirdPartyAccessTokenValidatorRequest = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.AddJwtAccessTokenValidatorRequest = nil
		dst.AddMockAccessTokenValidatorRequest = nil
		dst.AddPingFederateAccessTokenValidatorRequest = nil
		dst.AddThirdPartyAccessTokenValidatorRequest = nil

		return fmt.Errorf("data matches more than one schema in oneOf(AddAccessTokenValidatorRequest)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(AddAccessTokenValidatorRequest)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src AddAccessTokenValidatorRequest) MarshalJSON() ([]byte, error) {
	if src.AddJwtAccessTokenValidatorRequest != nil {
		return json.Marshal(&src.AddJwtAccessTokenValidatorRequest)
	}

	if src.AddMockAccessTokenValidatorRequest != nil {
		return json.Marshal(&src.AddMockAccessTokenValidatorRequest)
	}

	if src.AddPingFederateAccessTokenValidatorRequest != nil {
		return json.Marshal(&src.AddPingFederateAccessTokenValidatorRequest)
	}

	if src.AddThirdPartyAccessTokenValidatorRequest != nil {
		return json.Marshal(&src.AddThirdPartyAccessTokenValidatorRequest)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *AddAccessTokenValidatorRequest) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.AddJwtAccessTokenValidatorRequest != nil {
		return obj.AddJwtAccessTokenValidatorRequest
	}

	if obj.AddMockAccessTokenValidatorRequest != nil {
		return obj.AddMockAccessTokenValidatorRequest
	}

	if obj.AddPingFederateAccessTokenValidatorRequest != nil {
		return obj.AddPingFederateAccessTokenValidatorRequest
	}

	if obj.AddThirdPartyAccessTokenValidatorRequest != nil {
		return obj.AddThirdPartyAccessTokenValidatorRequest
	}

	// all schemas are nil
	return nil
}

type NullableAddAccessTokenValidatorRequest struct {
	value *AddAccessTokenValidatorRequest
	isSet bool
}

func (v NullableAddAccessTokenValidatorRequest) Get() *AddAccessTokenValidatorRequest {
	return v.value
}

func (v *NullableAddAccessTokenValidatorRequest) Set(val *AddAccessTokenValidatorRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddAccessTokenValidatorRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddAccessTokenValidatorRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddAccessTokenValidatorRequest(val *AddAccessTokenValidatorRequest) *NullableAddAccessTokenValidatorRequest {
	return &NullableAddAccessTokenValidatorRequest{value: val, isSet: true}
}

func (v NullableAddAccessTokenValidatorRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddAccessTokenValidatorRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

