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

// AddPassThroughAuthenticationHandler200Response - struct for AddPassThroughAuthenticationHandler200Response
type AddPassThroughAuthenticationHandler200Response struct {
	AggregatePassThroughAuthenticationHandlerResponse  *AggregatePassThroughAuthenticationHandlerResponse
	LdapPassThroughAuthenticationHandlerResponse       *LdapPassThroughAuthenticationHandlerResponse
	PingOnePassThroughAuthenticationHandlerResponse    *PingOnePassThroughAuthenticationHandlerResponse
	ThirdPartyPassThroughAuthenticationHandlerResponse *ThirdPartyPassThroughAuthenticationHandlerResponse
}

// AggregatePassThroughAuthenticationHandlerResponseAsAddPassThroughAuthenticationHandler200Response is a convenience function that returns AggregatePassThroughAuthenticationHandlerResponse wrapped in AddPassThroughAuthenticationHandler200Response
func AggregatePassThroughAuthenticationHandlerResponseAsAddPassThroughAuthenticationHandler200Response(v *AggregatePassThroughAuthenticationHandlerResponse) AddPassThroughAuthenticationHandler200Response {
	return AddPassThroughAuthenticationHandler200Response{
		AggregatePassThroughAuthenticationHandlerResponse: v,
	}
}

// LdapPassThroughAuthenticationHandlerResponseAsAddPassThroughAuthenticationHandler200Response is a convenience function that returns LdapPassThroughAuthenticationHandlerResponse wrapped in AddPassThroughAuthenticationHandler200Response
func LdapPassThroughAuthenticationHandlerResponseAsAddPassThroughAuthenticationHandler200Response(v *LdapPassThroughAuthenticationHandlerResponse) AddPassThroughAuthenticationHandler200Response {
	return AddPassThroughAuthenticationHandler200Response{
		LdapPassThroughAuthenticationHandlerResponse: v,
	}
}

// PingOnePassThroughAuthenticationHandlerResponseAsAddPassThroughAuthenticationHandler200Response is a convenience function that returns PingOnePassThroughAuthenticationHandlerResponse wrapped in AddPassThroughAuthenticationHandler200Response
func PingOnePassThroughAuthenticationHandlerResponseAsAddPassThroughAuthenticationHandler200Response(v *PingOnePassThroughAuthenticationHandlerResponse) AddPassThroughAuthenticationHandler200Response {
	return AddPassThroughAuthenticationHandler200Response{
		PingOnePassThroughAuthenticationHandlerResponse: v,
	}
}

// ThirdPartyPassThroughAuthenticationHandlerResponseAsAddPassThroughAuthenticationHandler200Response is a convenience function that returns ThirdPartyPassThroughAuthenticationHandlerResponse wrapped in AddPassThroughAuthenticationHandler200Response
func ThirdPartyPassThroughAuthenticationHandlerResponseAsAddPassThroughAuthenticationHandler200Response(v *ThirdPartyPassThroughAuthenticationHandlerResponse) AddPassThroughAuthenticationHandler200Response {
	return AddPassThroughAuthenticationHandler200Response{
		ThirdPartyPassThroughAuthenticationHandlerResponse: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *AddPassThroughAuthenticationHandler200Response) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AggregatePassThroughAuthenticationHandlerResponse
	err = newStrictDecoder(data).Decode(&dst.AggregatePassThroughAuthenticationHandlerResponse)
	if err == nil {
		jsonAggregatePassThroughAuthenticationHandlerResponse, _ := json.Marshal(dst.AggregatePassThroughAuthenticationHandlerResponse)
		if string(jsonAggregatePassThroughAuthenticationHandlerResponse) == "{}" { // empty struct
			dst.AggregatePassThroughAuthenticationHandlerResponse = nil
		} else {
			match++
		}
	} else {
		dst.AggregatePassThroughAuthenticationHandlerResponse = nil
	}

	// try to unmarshal data into LdapPassThroughAuthenticationHandlerResponse
	err = newStrictDecoder(data).Decode(&dst.LdapPassThroughAuthenticationHandlerResponse)
	if err == nil {
		jsonLdapPassThroughAuthenticationHandlerResponse, _ := json.Marshal(dst.LdapPassThroughAuthenticationHandlerResponse)
		if string(jsonLdapPassThroughAuthenticationHandlerResponse) == "{}" { // empty struct
			dst.LdapPassThroughAuthenticationHandlerResponse = nil
		} else {
			match++
		}
	} else {
		dst.LdapPassThroughAuthenticationHandlerResponse = nil
	}

	// try to unmarshal data into PingOnePassThroughAuthenticationHandlerResponse
	err = newStrictDecoder(data).Decode(&dst.PingOnePassThroughAuthenticationHandlerResponse)
	if err == nil {
		jsonPingOnePassThroughAuthenticationHandlerResponse, _ := json.Marshal(dst.PingOnePassThroughAuthenticationHandlerResponse)
		if string(jsonPingOnePassThroughAuthenticationHandlerResponse) == "{}" { // empty struct
			dst.PingOnePassThroughAuthenticationHandlerResponse = nil
		} else {
			match++
		}
	} else {
		dst.PingOnePassThroughAuthenticationHandlerResponse = nil
	}

	// try to unmarshal data into ThirdPartyPassThroughAuthenticationHandlerResponse
	err = newStrictDecoder(data).Decode(&dst.ThirdPartyPassThroughAuthenticationHandlerResponse)
	if err == nil {
		jsonThirdPartyPassThroughAuthenticationHandlerResponse, _ := json.Marshal(dst.ThirdPartyPassThroughAuthenticationHandlerResponse)
		if string(jsonThirdPartyPassThroughAuthenticationHandlerResponse) == "{}" { // empty struct
			dst.ThirdPartyPassThroughAuthenticationHandlerResponse = nil
		} else {
			match++
		}
	} else {
		dst.ThirdPartyPassThroughAuthenticationHandlerResponse = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.AggregatePassThroughAuthenticationHandlerResponse = nil
		dst.LdapPassThroughAuthenticationHandlerResponse = nil
		dst.PingOnePassThroughAuthenticationHandlerResponse = nil
		dst.ThirdPartyPassThroughAuthenticationHandlerResponse = nil

		return fmt.Errorf("data matches more than one schema in oneOf(AddPassThroughAuthenticationHandler200Response)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(AddPassThroughAuthenticationHandler200Response)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src AddPassThroughAuthenticationHandler200Response) MarshalJSON() ([]byte, error) {
	if src.AggregatePassThroughAuthenticationHandlerResponse != nil {
		return json.Marshal(&src.AggregatePassThroughAuthenticationHandlerResponse)
	}

	if src.LdapPassThroughAuthenticationHandlerResponse != nil {
		return json.Marshal(&src.LdapPassThroughAuthenticationHandlerResponse)
	}

	if src.PingOnePassThroughAuthenticationHandlerResponse != nil {
		return json.Marshal(&src.PingOnePassThroughAuthenticationHandlerResponse)
	}

	if src.ThirdPartyPassThroughAuthenticationHandlerResponse != nil {
		return json.Marshal(&src.ThirdPartyPassThroughAuthenticationHandlerResponse)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *AddPassThroughAuthenticationHandler200Response) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.AggregatePassThroughAuthenticationHandlerResponse != nil {
		return obj.AggregatePassThroughAuthenticationHandlerResponse
	}

	if obj.LdapPassThroughAuthenticationHandlerResponse != nil {
		return obj.LdapPassThroughAuthenticationHandlerResponse
	}

	if obj.PingOnePassThroughAuthenticationHandlerResponse != nil {
		return obj.PingOnePassThroughAuthenticationHandlerResponse
	}

	if obj.ThirdPartyPassThroughAuthenticationHandlerResponse != nil {
		return obj.ThirdPartyPassThroughAuthenticationHandlerResponse
	}

	// all schemas are nil
	return nil
}

type NullableAddPassThroughAuthenticationHandler200Response struct {
	value *AddPassThroughAuthenticationHandler200Response
	isSet bool
}

func (v NullableAddPassThroughAuthenticationHandler200Response) Get() *AddPassThroughAuthenticationHandler200Response {
	return v.value
}

func (v *NullableAddPassThroughAuthenticationHandler200Response) Set(val *AddPassThroughAuthenticationHandler200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableAddPassThroughAuthenticationHandler200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableAddPassThroughAuthenticationHandler200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddPassThroughAuthenticationHandler200Response(val *AddPassThroughAuthenticationHandler200Response) *NullableAddPassThroughAuthenticationHandler200Response {
	return &NullableAddPassThroughAuthenticationHandler200Response{value: val, isSet: true}
}

func (v NullableAddPassThroughAuthenticationHandler200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddPassThroughAuthenticationHandler200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
