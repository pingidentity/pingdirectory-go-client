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

// AddConnectionHandler200Response - struct for AddConnectionHandler200Response
type AddConnectionHandler200Response struct {
	HttpConnectionHandlerResponse *HttpConnectionHandlerResponse
	JmxConnectionHandlerResponse *JmxConnectionHandlerResponse
	LdapConnectionHandlerResponse *LdapConnectionHandlerResponse
	LdifConnectionHandlerResponse *LdifConnectionHandlerResponse
}

// HttpConnectionHandlerResponseAsAddConnectionHandler200Response is a convenience function that returns HttpConnectionHandlerResponse wrapped in AddConnectionHandler200Response
func HttpConnectionHandlerResponseAsAddConnectionHandler200Response(v *HttpConnectionHandlerResponse) AddConnectionHandler200Response {
	return AddConnectionHandler200Response{
		HttpConnectionHandlerResponse: v,
	}
}

// JmxConnectionHandlerResponseAsAddConnectionHandler200Response is a convenience function that returns JmxConnectionHandlerResponse wrapped in AddConnectionHandler200Response
func JmxConnectionHandlerResponseAsAddConnectionHandler200Response(v *JmxConnectionHandlerResponse) AddConnectionHandler200Response {
	return AddConnectionHandler200Response{
		JmxConnectionHandlerResponse: v,
	}
}

// LdapConnectionHandlerResponseAsAddConnectionHandler200Response is a convenience function that returns LdapConnectionHandlerResponse wrapped in AddConnectionHandler200Response
func LdapConnectionHandlerResponseAsAddConnectionHandler200Response(v *LdapConnectionHandlerResponse) AddConnectionHandler200Response {
	return AddConnectionHandler200Response{
		LdapConnectionHandlerResponse: v,
	}
}

// LdifConnectionHandlerResponseAsAddConnectionHandler200Response is a convenience function that returns LdifConnectionHandlerResponse wrapped in AddConnectionHandler200Response
func LdifConnectionHandlerResponseAsAddConnectionHandler200Response(v *LdifConnectionHandlerResponse) AddConnectionHandler200Response {
	return AddConnectionHandler200Response{
		LdifConnectionHandlerResponse: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *AddConnectionHandler200Response) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into HttpConnectionHandlerResponse
	err = newStrictDecoder(data).Decode(&dst.HttpConnectionHandlerResponse)
	if err == nil {
		jsonHttpConnectionHandlerResponse, _ := json.Marshal(dst.HttpConnectionHandlerResponse)
		if string(jsonHttpConnectionHandlerResponse) == "{}" { // empty struct
			dst.HttpConnectionHandlerResponse = nil
		} else {
			match++
		}
	} else {
		dst.HttpConnectionHandlerResponse = nil
	}

	// try to unmarshal data into JmxConnectionHandlerResponse
	err = newStrictDecoder(data).Decode(&dst.JmxConnectionHandlerResponse)
	if err == nil {
		jsonJmxConnectionHandlerResponse, _ := json.Marshal(dst.JmxConnectionHandlerResponse)
		if string(jsonJmxConnectionHandlerResponse) == "{}" { // empty struct
			dst.JmxConnectionHandlerResponse = nil
		} else {
			match++
		}
	} else {
		dst.JmxConnectionHandlerResponse = nil
	}

	// try to unmarshal data into LdapConnectionHandlerResponse
	err = newStrictDecoder(data).Decode(&dst.LdapConnectionHandlerResponse)
	if err == nil {
		jsonLdapConnectionHandlerResponse, _ := json.Marshal(dst.LdapConnectionHandlerResponse)
		if string(jsonLdapConnectionHandlerResponse) == "{}" { // empty struct
			dst.LdapConnectionHandlerResponse = nil
		} else {
			match++
		}
	} else {
		dst.LdapConnectionHandlerResponse = nil
	}

	// try to unmarshal data into LdifConnectionHandlerResponse
	err = newStrictDecoder(data).Decode(&dst.LdifConnectionHandlerResponse)
	if err == nil {
		jsonLdifConnectionHandlerResponse, _ := json.Marshal(dst.LdifConnectionHandlerResponse)
		if string(jsonLdifConnectionHandlerResponse) == "{}" { // empty struct
			dst.LdifConnectionHandlerResponse = nil
		} else {
			match++
		}
	} else {
		dst.LdifConnectionHandlerResponse = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.HttpConnectionHandlerResponse = nil
		dst.JmxConnectionHandlerResponse = nil
		dst.LdapConnectionHandlerResponse = nil
		dst.LdifConnectionHandlerResponse = nil

		return fmt.Errorf("data matches more than one schema in oneOf(AddConnectionHandler200Response)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(AddConnectionHandler200Response)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src AddConnectionHandler200Response) MarshalJSON() ([]byte, error) {
	if src.HttpConnectionHandlerResponse != nil {
		return json.Marshal(&src.HttpConnectionHandlerResponse)
	}

	if src.JmxConnectionHandlerResponse != nil {
		return json.Marshal(&src.JmxConnectionHandlerResponse)
	}

	if src.LdapConnectionHandlerResponse != nil {
		return json.Marshal(&src.LdapConnectionHandlerResponse)
	}

	if src.LdifConnectionHandlerResponse != nil {
		return json.Marshal(&src.LdifConnectionHandlerResponse)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *AddConnectionHandler200Response) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.HttpConnectionHandlerResponse != nil {
		return obj.HttpConnectionHandlerResponse
	}

	if obj.JmxConnectionHandlerResponse != nil {
		return obj.JmxConnectionHandlerResponse
	}

	if obj.LdapConnectionHandlerResponse != nil {
		return obj.LdapConnectionHandlerResponse
	}

	if obj.LdifConnectionHandlerResponse != nil {
		return obj.LdifConnectionHandlerResponse
	}

	// all schemas are nil
	return nil
}

type NullableAddConnectionHandler200Response struct {
	value *AddConnectionHandler200Response
	isSet bool
}

func (v NullableAddConnectionHandler200Response) Get() *AddConnectionHandler200Response {
	return v.value
}

func (v *NullableAddConnectionHandler200Response) Set(val *AddConnectionHandler200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableAddConnectionHandler200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableAddConnectionHandler200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddConnectionHandler200Response(val *AddConnectionHandler200Response) *NullableAddConnectionHandler200Response {
	return &NullableAddConnectionHandler200Response{value: val, isSet: true}
}

func (v NullableAddConnectionHandler200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddConnectionHandler200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


