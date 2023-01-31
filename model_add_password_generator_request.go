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

// AddPasswordGeneratorRequest - struct for AddPasswordGeneratorRequest
type AddPasswordGeneratorRequest struct {
	AddGroovyScriptedPasswordGeneratorRequest *AddGroovyScriptedPasswordGeneratorRequest
	AddPassphrasePasswordGeneratorRequest     *AddPassphrasePasswordGeneratorRequest
	AddRandomPasswordGeneratorRequest         *AddRandomPasswordGeneratorRequest
	AddThirdPartyPasswordGeneratorRequest     *AddThirdPartyPasswordGeneratorRequest
}

// AddGroovyScriptedPasswordGeneratorRequestAsAddPasswordGeneratorRequest is a convenience function that returns AddGroovyScriptedPasswordGeneratorRequest wrapped in AddPasswordGeneratorRequest
func AddGroovyScriptedPasswordGeneratorRequestAsAddPasswordGeneratorRequest(v *AddGroovyScriptedPasswordGeneratorRequest) AddPasswordGeneratorRequest {
	return AddPasswordGeneratorRequest{
		AddGroovyScriptedPasswordGeneratorRequest: v,
	}
}

// AddPassphrasePasswordGeneratorRequestAsAddPasswordGeneratorRequest is a convenience function that returns AddPassphrasePasswordGeneratorRequest wrapped in AddPasswordGeneratorRequest
func AddPassphrasePasswordGeneratorRequestAsAddPasswordGeneratorRequest(v *AddPassphrasePasswordGeneratorRequest) AddPasswordGeneratorRequest {
	return AddPasswordGeneratorRequest{
		AddPassphrasePasswordGeneratorRequest: v,
	}
}

// AddRandomPasswordGeneratorRequestAsAddPasswordGeneratorRequest is a convenience function that returns AddRandomPasswordGeneratorRequest wrapped in AddPasswordGeneratorRequest
func AddRandomPasswordGeneratorRequestAsAddPasswordGeneratorRequest(v *AddRandomPasswordGeneratorRequest) AddPasswordGeneratorRequest {
	return AddPasswordGeneratorRequest{
		AddRandomPasswordGeneratorRequest: v,
	}
}

// AddThirdPartyPasswordGeneratorRequestAsAddPasswordGeneratorRequest is a convenience function that returns AddThirdPartyPasswordGeneratorRequest wrapped in AddPasswordGeneratorRequest
func AddThirdPartyPasswordGeneratorRequestAsAddPasswordGeneratorRequest(v *AddThirdPartyPasswordGeneratorRequest) AddPasswordGeneratorRequest {
	return AddPasswordGeneratorRequest{
		AddThirdPartyPasswordGeneratorRequest: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *AddPasswordGeneratorRequest) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AddGroovyScriptedPasswordGeneratorRequest
	err = newStrictDecoder(data).Decode(&dst.AddGroovyScriptedPasswordGeneratorRequest)
	if err == nil {
		jsonAddGroovyScriptedPasswordGeneratorRequest, _ := json.Marshal(dst.AddGroovyScriptedPasswordGeneratorRequest)
		if string(jsonAddGroovyScriptedPasswordGeneratorRequest) == "{}" { // empty struct
			dst.AddGroovyScriptedPasswordGeneratorRequest = nil
		} else {
			match++
		}
	} else {
		dst.AddGroovyScriptedPasswordGeneratorRequest = nil
	}

	// try to unmarshal data into AddPassphrasePasswordGeneratorRequest
	err = newStrictDecoder(data).Decode(&dst.AddPassphrasePasswordGeneratorRequest)
	if err == nil {
		jsonAddPassphrasePasswordGeneratorRequest, _ := json.Marshal(dst.AddPassphrasePasswordGeneratorRequest)
		if string(jsonAddPassphrasePasswordGeneratorRequest) == "{}" { // empty struct
			dst.AddPassphrasePasswordGeneratorRequest = nil
		} else {
			match++
		}
	} else {
		dst.AddPassphrasePasswordGeneratorRequest = nil
	}

	// try to unmarshal data into AddRandomPasswordGeneratorRequest
	err = newStrictDecoder(data).Decode(&dst.AddRandomPasswordGeneratorRequest)
	if err == nil {
		jsonAddRandomPasswordGeneratorRequest, _ := json.Marshal(dst.AddRandomPasswordGeneratorRequest)
		if string(jsonAddRandomPasswordGeneratorRequest) == "{}" { // empty struct
			dst.AddRandomPasswordGeneratorRequest = nil
		} else {
			match++
		}
	} else {
		dst.AddRandomPasswordGeneratorRequest = nil
	}

	// try to unmarshal data into AddThirdPartyPasswordGeneratorRequest
	err = newStrictDecoder(data).Decode(&dst.AddThirdPartyPasswordGeneratorRequest)
	if err == nil {
		jsonAddThirdPartyPasswordGeneratorRequest, _ := json.Marshal(dst.AddThirdPartyPasswordGeneratorRequest)
		if string(jsonAddThirdPartyPasswordGeneratorRequest) == "{}" { // empty struct
			dst.AddThirdPartyPasswordGeneratorRequest = nil
		} else {
			match++
		}
	} else {
		dst.AddThirdPartyPasswordGeneratorRequest = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.AddGroovyScriptedPasswordGeneratorRequest = nil
		dst.AddPassphrasePasswordGeneratorRequest = nil
		dst.AddRandomPasswordGeneratorRequest = nil
		dst.AddThirdPartyPasswordGeneratorRequest = nil

		return fmt.Errorf("data matches more than one schema in oneOf(AddPasswordGeneratorRequest)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(AddPasswordGeneratorRequest)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src AddPasswordGeneratorRequest) MarshalJSON() ([]byte, error) {
	if src.AddGroovyScriptedPasswordGeneratorRequest != nil {
		return json.Marshal(&src.AddGroovyScriptedPasswordGeneratorRequest)
	}

	if src.AddPassphrasePasswordGeneratorRequest != nil {
		return json.Marshal(&src.AddPassphrasePasswordGeneratorRequest)
	}

	if src.AddRandomPasswordGeneratorRequest != nil {
		return json.Marshal(&src.AddRandomPasswordGeneratorRequest)
	}

	if src.AddThirdPartyPasswordGeneratorRequest != nil {
		return json.Marshal(&src.AddThirdPartyPasswordGeneratorRequest)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *AddPasswordGeneratorRequest) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.AddGroovyScriptedPasswordGeneratorRequest != nil {
		return obj.AddGroovyScriptedPasswordGeneratorRequest
	}

	if obj.AddPassphrasePasswordGeneratorRequest != nil {
		return obj.AddPassphrasePasswordGeneratorRequest
	}

	if obj.AddRandomPasswordGeneratorRequest != nil {
		return obj.AddRandomPasswordGeneratorRequest
	}

	if obj.AddThirdPartyPasswordGeneratorRequest != nil {
		return obj.AddThirdPartyPasswordGeneratorRequest
	}

	// all schemas are nil
	return nil
}

type NullableAddPasswordGeneratorRequest struct {
	value *AddPasswordGeneratorRequest
	isSet bool
}

func (v NullableAddPasswordGeneratorRequest) Get() *AddPasswordGeneratorRequest {
	return v.value
}

func (v *NullableAddPasswordGeneratorRequest) Set(val *AddPasswordGeneratorRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddPasswordGeneratorRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddPasswordGeneratorRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddPasswordGeneratorRequest(val *AddPasswordGeneratorRequest) *NullableAddPasswordGeneratorRequest {
	return &NullableAddPasswordGeneratorRequest{value: val, isSet: true}
}

func (v NullableAddPasswordGeneratorRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddPasswordGeneratorRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
