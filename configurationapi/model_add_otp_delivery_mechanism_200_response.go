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

// AddOtpDeliveryMechanism200Response - struct for AddOtpDeliveryMechanism200Response
type AddOtpDeliveryMechanism200Response struct {
	EmailOtpDeliveryMechanismResponse      *EmailOtpDeliveryMechanismResponse
	ThirdPartyOtpDeliveryMechanismResponse *ThirdPartyOtpDeliveryMechanismResponse
	TwilioOtpDeliveryMechanismResponse     *TwilioOtpDeliveryMechanismResponse
}

// EmailOtpDeliveryMechanismResponseAsAddOtpDeliveryMechanism200Response is a convenience function that returns EmailOtpDeliveryMechanismResponse wrapped in AddOtpDeliveryMechanism200Response
func EmailOtpDeliveryMechanismResponseAsAddOtpDeliveryMechanism200Response(v *EmailOtpDeliveryMechanismResponse) AddOtpDeliveryMechanism200Response {
	return AddOtpDeliveryMechanism200Response{
		EmailOtpDeliveryMechanismResponse: v,
	}
}

// ThirdPartyOtpDeliveryMechanismResponseAsAddOtpDeliveryMechanism200Response is a convenience function that returns ThirdPartyOtpDeliveryMechanismResponse wrapped in AddOtpDeliveryMechanism200Response
func ThirdPartyOtpDeliveryMechanismResponseAsAddOtpDeliveryMechanism200Response(v *ThirdPartyOtpDeliveryMechanismResponse) AddOtpDeliveryMechanism200Response {
	return AddOtpDeliveryMechanism200Response{
		ThirdPartyOtpDeliveryMechanismResponse: v,
	}
}

// TwilioOtpDeliveryMechanismResponseAsAddOtpDeliveryMechanism200Response is a convenience function that returns TwilioOtpDeliveryMechanismResponse wrapped in AddOtpDeliveryMechanism200Response
func TwilioOtpDeliveryMechanismResponseAsAddOtpDeliveryMechanism200Response(v *TwilioOtpDeliveryMechanismResponse) AddOtpDeliveryMechanism200Response {
	return AddOtpDeliveryMechanism200Response{
		TwilioOtpDeliveryMechanismResponse: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *AddOtpDeliveryMechanism200Response) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into EmailOtpDeliveryMechanismResponse
	err = newStrictDecoder(data).Decode(&dst.EmailOtpDeliveryMechanismResponse)
	if err == nil {
		jsonEmailOtpDeliveryMechanismResponse, _ := json.Marshal(dst.EmailOtpDeliveryMechanismResponse)
		if string(jsonEmailOtpDeliveryMechanismResponse) == "{}" { // empty struct
			dst.EmailOtpDeliveryMechanismResponse = nil
		} else {
			match++
		}
	} else {
		dst.EmailOtpDeliveryMechanismResponse = nil
	}

	// try to unmarshal data into ThirdPartyOtpDeliveryMechanismResponse
	err = newStrictDecoder(data).Decode(&dst.ThirdPartyOtpDeliveryMechanismResponse)
	if err == nil {
		jsonThirdPartyOtpDeliveryMechanismResponse, _ := json.Marshal(dst.ThirdPartyOtpDeliveryMechanismResponse)
		if string(jsonThirdPartyOtpDeliveryMechanismResponse) == "{}" { // empty struct
			dst.ThirdPartyOtpDeliveryMechanismResponse = nil
		} else {
			match++
		}
	} else {
		dst.ThirdPartyOtpDeliveryMechanismResponse = nil
	}

	// try to unmarshal data into TwilioOtpDeliveryMechanismResponse
	err = newStrictDecoder(data).Decode(&dst.TwilioOtpDeliveryMechanismResponse)
	if err == nil {
		jsonTwilioOtpDeliveryMechanismResponse, _ := json.Marshal(dst.TwilioOtpDeliveryMechanismResponse)
		if string(jsonTwilioOtpDeliveryMechanismResponse) == "{}" { // empty struct
			dst.TwilioOtpDeliveryMechanismResponse = nil
		} else {
			match++
		}
	} else {
		dst.TwilioOtpDeliveryMechanismResponse = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.EmailOtpDeliveryMechanismResponse = nil
		dst.ThirdPartyOtpDeliveryMechanismResponse = nil
		dst.TwilioOtpDeliveryMechanismResponse = nil

		return fmt.Errorf("data matches more than one schema in oneOf(AddOtpDeliveryMechanism200Response)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(AddOtpDeliveryMechanism200Response)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src AddOtpDeliveryMechanism200Response) MarshalJSON() ([]byte, error) {
	if src.EmailOtpDeliveryMechanismResponse != nil {
		return json.Marshal(&src.EmailOtpDeliveryMechanismResponse)
	}

	if src.ThirdPartyOtpDeliveryMechanismResponse != nil {
		return json.Marshal(&src.ThirdPartyOtpDeliveryMechanismResponse)
	}

	if src.TwilioOtpDeliveryMechanismResponse != nil {
		return json.Marshal(&src.TwilioOtpDeliveryMechanismResponse)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *AddOtpDeliveryMechanism200Response) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.EmailOtpDeliveryMechanismResponse != nil {
		return obj.EmailOtpDeliveryMechanismResponse
	}

	if obj.ThirdPartyOtpDeliveryMechanismResponse != nil {
		return obj.ThirdPartyOtpDeliveryMechanismResponse
	}

	if obj.TwilioOtpDeliveryMechanismResponse != nil {
		return obj.TwilioOtpDeliveryMechanismResponse
	}

	// all schemas are nil
	return nil
}

type NullableAddOtpDeliveryMechanism200Response struct {
	value *AddOtpDeliveryMechanism200Response
	isSet bool
}

func (v NullableAddOtpDeliveryMechanism200Response) Get() *AddOtpDeliveryMechanism200Response {
	return v.value
}

func (v *NullableAddOtpDeliveryMechanism200Response) Set(val *AddOtpDeliveryMechanism200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableAddOtpDeliveryMechanism200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableAddOtpDeliveryMechanism200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddOtpDeliveryMechanism200Response(val *AddOtpDeliveryMechanism200Response) *NullableAddOtpDeliveryMechanism200Response {
	return &NullableAddOtpDeliveryMechanism200Response{value: val, isSet: true}
}

func (v NullableAddOtpDeliveryMechanism200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddOtpDeliveryMechanism200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
