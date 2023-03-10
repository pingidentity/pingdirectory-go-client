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

// AddConnectionHandlerRequest - struct for AddConnectionHandlerRequest
type AddConnectionHandlerRequest struct {
	AddHttpConnectionHandlerRequest *AddHttpConnectionHandlerRequest
	AddJmxConnectionHandlerRequest  *AddJmxConnectionHandlerRequest
	AddLdapConnectionHandlerRequest *AddLdapConnectionHandlerRequest
	AddLdifConnectionHandlerRequest *AddLdifConnectionHandlerRequest
}

// AddHttpConnectionHandlerRequestAsAddConnectionHandlerRequest is a convenience function that returns AddHttpConnectionHandlerRequest wrapped in AddConnectionHandlerRequest
func AddHttpConnectionHandlerRequestAsAddConnectionHandlerRequest(v *AddHttpConnectionHandlerRequest) AddConnectionHandlerRequest {
	return AddConnectionHandlerRequest{
		AddHttpConnectionHandlerRequest: v,
	}
}

// AddJmxConnectionHandlerRequestAsAddConnectionHandlerRequest is a convenience function that returns AddJmxConnectionHandlerRequest wrapped in AddConnectionHandlerRequest
func AddJmxConnectionHandlerRequestAsAddConnectionHandlerRequest(v *AddJmxConnectionHandlerRequest) AddConnectionHandlerRequest {
	return AddConnectionHandlerRequest{
		AddJmxConnectionHandlerRequest: v,
	}
}

// AddLdapConnectionHandlerRequestAsAddConnectionHandlerRequest is a convenience function that returns AddLdapConnectionHandlerRequest wrapped in AddConnectionHandlerRequest
func AddLdapConnectionHandlerRequestAsAddConnectionHandlerRequest(v *AddLdapConnectionHandlerRequest) AddConnectionHandlerRequest {
	return AddConnectionHandlerRequest{
		AddLdapConnectionHandlerRequest: v,
	}
}

// AddLdifConnectionHandlerRequestAsAddConnectionHandlerRequest is a convenience function that returns AddLdifConnectionHandlerRequest wrapped in AddConnectionHandlerRequest
func AddLdifConnectionHandlerRequestAsAddConnectionHandlerRequest(v *AddLdifConnectionHandlerRequest) AddConnectionHandlerRequest {
	return AddConnectionHandlerRequest{
		AddLdifConnectionHandlerRequest: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *AddConnectionHandlerRequest) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AddHttpConnectionHandlerRequest
	err = newStrictDecoder(data).Decode(&dst.AddHttpConnectionHandlerRequest)
	if err == nil {
		jsonAddHttpConnectionHandlerRequest, _ := json.Marshal(dst.AddHttpConnectionHandlerRequest)
		if string(jsonAddHttpConnectionHandlerRequest) == "{}" { // empty struct
			dst.AddHttpConnectionHandlerRequest = nil
		} else {
			match++
		}
	} else {
		dst.AddHttpConnectionHandlerRequest = nil
	}

	// try to unmarshal data into AddJmxConnectionHandlerRequest
	err = newStrictDecoder(data).Decode(&dst.AddJmxConnectionHandlerRequest)
	if err == nil {
		jsonAddJmxConnectionHandlerRequest, _ := json.Marshal(dst.AddJmxConnectionHandlerRequest)
		if string(jsonAddJmxConnectionHandlerRequest) == "{}" { // empty struct
			dst.AddJmxConnectionHandlerRequest = nil
		} else {
			match++
		}
	} else {
		dst.AddJmxConnectionHandlerRequest = nil
	}

	// try to unmarshal data into AddLdapConnectionHandlerRequest
	err = newStrictDecoder(data).Decode(&dst.AddLdapConnectionHandlerRequest)
	if err == nil {
		jsonAddLdapConnectionHandlerRequest, _ := json.Marshal(dst.AddLdapConnectionHandlerRequest)
		if string(jsonAddLdapConnectionHandlerRequest) == "{}" { // empty struct
			dst.AddLdapConnectionHandlerRequest = nil
		} else {
			match++
		}
	} else {
		dst.AddLdapConnectionHandlerRequest = nil
	}

	// try to unmarshal data into AddLdifConnectionHandlerRequest
	err = newStrictDecoder(data).Decode(&dst.AddLdifConnectionHandlerRequest)
	if err == nil {
		jsonAddLdifConnectionHandlerRequest, _ := json.Marshal(dst.AddLdifConnectionHandlerRequest)
		if string(jsonAddLdifConnectionHandlerRequest) == "{}" { // empty struct
			dst.AddLdifConnectionHandlerRequest = nil
		} else {
			match++
		}
	} else {
		dst.AddLdifConnectionHandlerRequest = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.AddHttpConnectionHandlerRequest = nil
		dst.AddJmxConnectionHandlerRequest = nil
		dst.AddLdapConnectionHandlerRequest = nil
		dst.AddLdifConnectionHandlerRequest = nil

		return fmt.Errorf("data matches more than one schema in oneOf(AddConnectionHandlerRequest)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(AddConnectionHandlerRequest)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src AddConnectionHandlerRequest) MarshalJSON() ([]byte, error) {
	if src.AddHttpConnectionHandlerRequest != nil {
		return json.Marshal(&src.AddHttpConnectionHandlerRequest)
	}

	if src.AddJmxConnectionHandlerRequest != nil {
		return json.Marshal(&src.AddJmxConnectionHandlerRequest)
	}

	if src.AddLdapConnectionHandlerRequest != nil {
		return json.Marshal(&src.AddLdapConnectionHandlerRequest)
	}

	if src.AddLdifConnectionHandlerRequest != nil {
		return json.Marshal(&src.AddLdifConnectionHandlerRequest)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *AddConnectionHandlerRequest) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.AddHttpConnectionHandlerRequest != nil {
		return obj.AddHttpConnectionHandlerRequest
	}

	if obj.AddJmxConnectionHandlerRequest != nil {
		return obj.AddJmxConnectionHandlerRequest
	}

	if obj.AddLdapConnectionHandlerRequest != nil {
		return obj.AddLdapConnectionHandlerRequest
	}

	if obj.AddLdifConnectionHandlerRequest != nil {
		return obj.AddLdifConnectionHandlerRequest
	}

	// all schemas are nil
	return nil
}

type NullableAddConnectionHandlerRequest struct {
	value *AddConnectionHandlerRequest
	isSet bool
}

func (v NullableAddConnectionHandlerRequest) Get() *AddConnectionHandlerRequest {
	return v.value
}

func (v *NullableAddConnectionHandlerRequest) Set(val *AddConnectionHandlerRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddConnectionHandlerRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddConnectionHandlerRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddConnectionHandlerRequest(val *AddConnectionHandlerRequest) *NullableAddConnectionHandlerRequest {
	return &NullableAddConnectionHandlerRequest{value: val, isSet: true}
}

func (v NullableAddConnectionHandlerRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddConnectionHandlerRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
