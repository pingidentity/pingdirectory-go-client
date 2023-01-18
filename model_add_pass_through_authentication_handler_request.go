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

// AddPassThroughAuthenticationHandlerRequest - struct for AddPassThroughAuthenticationHandlerRequest
type AddPassThroughAuthenticationHandlerRequest struct {
	AddLdapPassThroughAuthenticationHandlerRequest *AddLdapPassThroughAuthenticationHandlerRequest
	AddThirdPartyPassThroughAuthenticationHandlerRequest *AddThirdPartyPassThroughAuthenticationHandlerRequest
}

// AddLdapPassThroughAuthenticationHandlerRequestAsAddPassThroughAuthenticationHandlerRequest is a convenience function that returns AddLdapPassThroughAuthenticationHandlerRequest wrapped in AddPassThroughAuthenticationHandlerRequest
func AddLdapPassThroughAuthenticationHandlerRequestAsAddPassThroughAuthenticationHandlerRequest(v *AddLdapPassThroughAuthenticationHandlerRequest) AddPassThroughAuthenticationHandlerRequest {
	return AddPassThroughAuthenticationHandlerRequest{
		AddLdapPassThroughAuthenticationHandlerRequest: v,
	}
}

// AddThirdPartyPassThroughAuthenticationHandlerRequestAsAddPassThroughAuthenticationHandlerRequest is a convenience function that returns AddThirdPartyPassThroughAuthenticationHandlerRequest wrapped in AddPassThroughAuthenticationHandlerRequest
func AddThirdPartyPassThroughAuthenticationHandlerRequestAsAddPassThroughAuthenticationHandlerRequest(v *AddThirdPartyPassThroughAuthenticationHandlerRequest) AddPassThroughAuthenticationHandlerRequest {
	return AddPassThroughAuthenticationHandlerRequest{
		AddThirdPartyPassThroughAuthenticationHandlerRequest: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *AddPassThroughAuthenticationHandlerRequest) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AddLdapPassThroughAuthenticationHandlerRequest
	err = newStrictDecoder(data).Decode(&dst.AddLdapPassThroughAuthenticationHandlerRequest)
	if err == nil {
		jsonAddLdapPassThroughAuthenticationHandlerRequest, _ := json.Marshal(dst.AddLdapPassThroughAuthenticationHandlerRequest)
		if string(jsonAddLdapPassThroughAuthenticationHandlerRequest) == "{}" { // empty struct
			dst.AddLdapPassThroughAuthenticationHandlerRequest = nil
		} else {
			match++
		}
	} else {
		dst.AddLdapPassThroughAuthenticationHandlerRequest = nil
	}

	// try to unmarshal data into AddThirdPartyPassThroughAuthenticationHandlerRequest
	err = newStrictDecoder(data).Decode(&dst.AddThirdPartyPassThroughAuthenticationHandlerRequest)
	if err == nil {
		jsonAddThirdPartyPassThroughAuthenticationHandlerRequest, _ := json.Marshal(dst.AddThirdPartyPassThroughAuthenticationHandlerRequest)
		if string(jsonAddThirdPartyPassThroughAuthenticationHandlerRequest) == "{}" { // empty struct
			dst.AddThirdPartyPassThroughAuthenticationHandlerRequest = nil
		} else {
			match++
		}
	} else {
		dst.AddThirdPartyPassThroughAuthenticationHandlerRequest = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.AddLdapPassThroughAuthenticationHandlerRequest = nil
		dst.AddThirdPartyPassThroughAuthenticationHandlerRequest = nil

		return fmt.Errorf("data matches more than one schema in oneOf(AddPassThroughAuthenticationHandlerRequest)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(AddPassThroughAuthenticationHandlerRequest)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src AddPassThroughAuthenticationHandlerRequest) MarshalJSON() ([]byte, error) {
	if src.AddLdapPassThroughAuthenticationHandlerRequest != nil {
		return json.Marshal(&src.AddLdapPassThroughAuthenticationHandlerRequest)
	}

	if src.AddThirdPartyPassThroughAuthenticationHandlerRequest != nil {
		return json.Marshal(&src.AddThirdPartyPassThroughAuthenticationHandlerRequest)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *AddPassThroughAuthenticationHandlerRequest) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.AddLdapPassThroughAuthenticationHandlerRequest != nil {
		return obj.AddLdapPassThroughAuthenticationHandlerRequest
	}

	if obj.AddThirdPartyPassThroughAuthenticationHandlerRequest != nil {
		return obj.AddThirdPartyPassThroughAuthenticationHandlerRequest
	}

	// all schemas are nil
	return nil
}

type NullableAddPassThroughAuthenticationHandlerRequest struct {
	value *AddPassThroughAuthenticationHandlerRequest
	isSet bool
}

func (v NullableAddPassThroughAuthenticationHandlerRequest) Get() *AddPassThroughAuthenticationHandlerRequest {
	return v.value
}

func (v *NullableAddPassThroughAuthenticationHandlerRequest) Set(val *AddPassThroughAuthenticationHandlerRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddPassThroughAuthenticationHandlerRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddPassThroughAuthenticationHandlerRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddPassThroughAuthenticationHandlerRequest(val *AddPassThroughAuthenticationHandlerRequest) *NullableAddPassThroughAuthenticationHandlerRequest {
	return &NullableAddPassThroughAuthenticationHandlerRequest{value: val, isSet: true}
}

func (v NullableAddPassThroughAuthenticationHandlerRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddPassThroughAuthenticationHandlerRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

