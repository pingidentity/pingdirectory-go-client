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

// AddUncachedEntryCriteriaRequest - struct for AddUncachedEntryCriteriaRequest
type AddUncachedEntryCriteriaRequest struct {
	AddDefaultUncachedEntryCriteriaRequest        *AddDefaultUncachedEntryCriteriaRequest
	AddFilterBasedUncachedEntryCriteriaRequest    *AddFilterBasedUncachedEntryCriteriaRequest
	AddGroovyScriptedUncachedEntryCriteriaRequest *AddGroovyScriptedUncachedEntryCriteriaRequest
	AddLastAccessTimeUncachedEntryCriteriaRequest *AddLastAccessTimeUncachedEntryCriteriaRequest
	AddThirdPartyUncachedEntryCriteriaRequest     *AddThirdPartyUncachedEntryCriteriaRequest
}

// AddDefaultUncachedEntryCriteriaRequestAsAddUncachedEntryCriteriaRequest is a convenience function that returns AddDefaultUncachedEntryCriteriaRequest wrapped in AddUncachedEntryCriteriaRequest
func AddDefaultUncachedEntryCriteriaRequestAsAddUncachedEntryCriteriaRequest(v *AddDefaultUncachedEntryCriteriaRequest) AddUncachedEntryCriteriaRequest {
	return AddUncachedEntryCriteriaRequest{
		AddDefaultUncachedEntryCriteriaRequest: v,
	}
}

// AddFilterBasedUncachedEntryCriteriaRequestAsAddUncachedEntryCriteriaRequest is a convenience function that returns AddFilterBasedUncachedEntryCriteriaRequest wrapped in AddUncachedEntryCriteriaRequest
func AddFilterBasedUncachedEntryCriteriaRequestAsAddUncachedEntryCriteriaRequest(v *AddFilterBasedUncachedEntryCriteriaRequest) AddUncachedEntryCriteriaRequest {
	return AddUncachedEntryCriteriaRequest{
		AddFilterBasedUncachedEntryCriteriaRequest: v,
	}
}

// AddGroovyScriptedUncachedEntryCriteriaRequestAsAddUncachedEntryCriteriaRequest is a convenience function that returns AddGroovyScriptedUncachedEntryCriteriaRequest wrapped in AddUncachedEntryCriteriaRequest
func AddGroovyScriptedUncachedEntryCriteriaRequestAsAddUncachedEntryCriteriaRequest(v *AddGroovyScriptedUncachedEntryCriteriaRequest) AddUncachedEntryCriteriaRequest {
	return AddUncachedEntryCriteriaRequest{
		AddGroovyScriptedUncachedEntryCriteriaRequest: v,
	}
}

// AddLastAccessTimeUncachedEntryCriteriaRequestAsAddUncachedEntryCriteriaRequest is a convenience function that returns AddLastAccessTimeUncachedEntryCriteriaRequest wrapped in AddUncachedEntryCriteriaRequest
func AddLastAccessTimeUncachedEntryCriteriaRequestAsAddUncachedEntryCriteriaRequest(v *AddLastAccessTimeUncachedEntryCriteriaRequest) AddUncachedEntryCriteriaRequest {
	return AddUncachedEntryCriteriaRequest{
		AddLastAccessTimeUncachedEntryCriteriaRequest: v,
	}
}

// AddThirdPartyUncachedEntryCriteriaRequestAsAddUncachedEntryCriteriaRequest is a convenience function that returns AddThirdPartyUncachedEntryCriteriaRequest wrapped in AddUncachedEntryCriteriaRequest
func AddThirdPartyUncachedEntryCriteriaRequestAsAddUncachedEntryCriteriaRequest(v *AddThirdPartyUncachedEntryCriteriaRequest) AddUncachedEntryCriteriaRequest {
	return AddUncachedEntryCriteriaRequest{
		AddThirdPartyUncachedEntryCriteriaRequest: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *AddUncachedEntryCriteriaRequest) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AddDefaultUncachedEntryCriteriaRequest
	err = newStrictDecoder(data).Decode(&dst.AddDefaultUncachedEntryCriteriaRequest)
	if err == nil {
		jsonAddDefaultUncachedEntryCriteriaRequest, _ := json.Marshal(dst.AddDefaultUncachedEntryCriteriaRequest)
		if string(jsonAddDefaultUncachedEntryCriteriaRequest) == "{}" { // empty struct
			dst.AddDefaultUncachedEntryCriteriaRequest = nil
		} else {
			match++
		}
	} else {
		dst.AddDefaultUncachedEntryCriteriaRequest = nil
	}

	// try to unmarshal data into AddFilterBasedUncachedEntryCriteriaRequest
	err = newStrictDecoder(data).Decode(&dst.AddFilterBasedUncachedEntryCriteriaRequest)
	if err == nil {
		jsonAddFilterBasedUncachedEntryCriteriaRequest, _ := json.Marshal(dst.AddFilterBasedUncachedEntryCriteriaRequest)
		if string(jsonAddFilterBasedUncachedEntryCriteriaRequest) == "{}" { // empty struct
			dst.AddFilterBasedUncachedEntryCriteriaRequest = nil
		} else {
			match++
		}
	} else {
		dst.AddFilterBasedUncachedEntryCriteriaRequest = nil
	}

	// try to unmarshal data into AddGroovyScriptedUncachedEntryCriteriaRequest
	err = newStrictDecoder(data).Decode(&dst.AddGroovyScriptedUncachedEntryCriteriaRequest)
	if err == nil {
		jsonAddGroovyScriptedUncachedEntryCriteriaRequest, _ := json.Marshal(dst.AddGroovyScriptedUncachedEntryCriteriaRequest)
		if string(jsonAddGroovyScriptedUncachedEntryCriteriaRequest) == "{}" { // empty struct
			dst.AddGroovyScriptedUncachedEntryCriteriaRequest = nil
		} else {
			match++
		}
	} else {
		dst.AddGroovyScriptedUncachedEntryCriteriaRequest = nil
	}

	// try to unmarshal data into AddLastAccessTimeUncachedEntryCriteriaRequest
	err = newStrictDecoder(data).Decode(&dst.AddLastAccessTimeUncachedEntryCriteriaRequest)
	if err == nil {
		jsonAddLastAccessTimeUncachedEntryCriteriaRequest, _ := json.Marshal(dst.AddLastAccessTimeUncachedEntryCriteriaRequest)
		if string(jsonAddLastAccessTimeUncachedEntryCriteriaRequest) == "{}" { // empty struct
			dst.AddLastAccessTimeUncachedEntryCriteriaRequest = nil
		} else {
			match++
		}
	} else {
		dst.AddLastAccessTimeUncachedEntryCriteriaRequest = nil
	}

	// try to unmarshal data into AddThirdPartyUncachedEntryCriteriaRequest
	err = newStrictDecoder(data).Decode(&dst.AddThirdPartyUncachedEntryCriteriaRequest)
	if err == nil {
		jsonAddThirdPartyUncachedEntryCriteriaRequest, _ := json.Marshal(dst.AddThirdPartyUncachedEntryCriteriaRequest)
		if string(jsonAddThirdPartyUncachedEntryCriteriaRequest) == "{}" { // empty struct
			dst.AddThirdPartyUncachedEntryCriteriaRequest = nil
		} else {
			match++
		}
	} else {
		dst.AddThirdPartyUncachedEntryCriteriaRequest = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.AddDefaultUncachedEntryCriteriaRequest = nil
		dst.AddFilterBasedUncachedEntryCriteriaRequest = nil
		dst.AddGroovyScriptedUncachedEntryCriteriaRequest = nil
		dst.AddLastAccessTimeUncachedEntryCriteriaRequest = nil
		dst.AddThirdPartyUncachedEntryCriteriaRequest = nil

		return fmt.Errorf("data matches more than one schema in oneOf(AddUncachedEntryCriteriaRequest)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(AddUncachedEntryCriteriaRequest)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src AddUncachedEntryCriteriaRequest) MarshalJSON() ([]byte, error) {
	if src.AddDefaultUncachedEntryCriteriaRequest != nil {
		return json.Marshal(&src.AddDefaultUncachedEntryCriteriaRequest)
	}

	if src.AddFilterBasedUncachedEntryCriteriaRequest != nil {
		return json.Marshal(&src.AddFilterBasedUncachedEntryCriteriaRequest)
	}

	if src.AddGroovyScriptedUncachedEntryCriteriaRequest != nil {
		return json.Marshal(&src.AddGroovyScriptedUncachedEntryCriteriaRequest)
	}

	if src.AddLastAccessTimeUncachedEntryCriteriaRequest != nil {
		return json.Marshal(&src.AddLastAccessTimeUncachedEntryCriteriaRequest)
	}

	if src.AddThirdPartyUncachedEntryCriteriaRequest != nil {
		return json.Marshal(&src.AddThirdPartyUncachedEntryCriteriaRequest)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *AddUncachedEntryCriteriaRequest) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.AddDefaultUncachedEntryCriteriaRequest != nil {
		return obj.AddDefaultUncachedEntryCriteriaRequest
	}

	if obj.AddFilterBasedUncachedEntryCriteriaRequest != nil {
		return obj.AddFilterBasedUncachedEntryCriteriaRequest
	}

	if obj.AddGroovyScriptedUncachedEntryCriteriaRequest != nil {
		return obj.AddGroovyScriptedUncachedEntryCriteriaRequest
	}

	if obj.AddLastAccessTimeUncachedEntryCriteriaRequest != nil {
		return obj.AddLastAccessTimeUncachedEntryCriteriaRequest
	}

	if obj.AddThirdPartyUncachedEntryCriteriaRequest != nil {
		return obj.AddThirdPartyUncachedEntryCriteriaRequest
	}

	// all schemas are nil
	return nil
}

type NullableAddUncachedEntryCriteriaRequest struct {
	value *AddUncachedEntryCriteriaRequest
	isSet bool
}

func (v NullableAddUncachedEntryCriteriaRequest) Get() *AddUncachedEntryCriteriaRequest {
	return v.value
}

func (v *NullableAddUncachedEntryCriteriaRequest) Set(val *AddUncachedEntryCriteriaRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddUncachedEntryCriteriaRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddUncachedEntryCriteriaRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddUncachedEntryCriteriaRequest(val *AddUncachedEntryCriteriaRequest) *NullableAddUncachedEntryCriteriaRequest {
	return &NullableAddUncachedEntryCriteriaRequest{value: val, isSet: true}
}

func (v NullableAddUncachedEntryCriteriaRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddUncachedEntryCriteriaRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
