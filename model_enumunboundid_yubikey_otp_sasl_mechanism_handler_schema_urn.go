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

// EnumunboundidYubikeyOtpSaslMechanismHandlerSchemaUrn the model 'EnumunboundidYubikeyOtpSaslMechanismHandlerSchemaUrn'
type EnumunboundidYubikeyOtpSaslMechanismHandlerSchemaUrn string

// List of Enumunboundid-yubikey-otp-sasl-mechanism-handlerSchemaUrn
const (
	ENUMUNBOUNDIDYUBIKEYOTPSASLMECHANISMHANDLERSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0SASL_MECHANISM_HANDLERUNBOUNDID_YUBIKEY_OTP EnumunboundidYubikeyOtpSaslMechanismHandlerSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:sasl-mechanism-handler:unboundid-yubikey-otp"
)

// All allowed values of EnumunboundidYubikeyOtpSaslMechanismHandlerSchemaUrn enum
var AllowedEnumunboundidYubikeyOtpSaslMechanismHandlerSchemaUrnEnumValues = []EnumunboundidYubikeyOtpSaslMechanismHandlerSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:sasl-mechanism-handler:unboundid-yubikey-otp",
}

func (v *EnumunboundidYubikeyOtpSaslMechanismHandlerSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumunboundidYubikeyOtpSaslMechanismHandlerSchemaUrn(value)
	for _, existing := range AllowedEnumunboundidYubikeyOtpSaslMechanismHandlerSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumunboundidYubikeyOtpSaslMechanismHandlerSchemaUrn", value)
}

// NewEnumunboundidYubikeyOtpSaslMechanismHandlerSchemaUrnFromValue returns a pointer to a valid EnumunboundidYubikeyOtpSaslMechanismHandlerSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumunboundidYubikeyOtpSaslMechanismHandlerSchemaUrnFromValue(v string) (*EnumunboundidYubikeyOtpSaslMechanismHandlerSchemaUrn, error) {
	ev := EnumunboundidYubikeyOtpSaslMechanismHandlerSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumunboundidYubikeyOtpSaslMechanismHandlerSchemaUrn: valid values are %v", v, AllowedEnumunboundidYubikeyOtpSaslMechanismHandlerSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumunboundidYubikeyOtpSaslMechanismHandlerSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumunboundidYubikeyOtpSaslMechanismHandlerSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumunboundid-yubikey-otp-sasl-mechanism-handlerSchemaUrn value
func (v EnumunboundidYubikeyOtpSaslMechanismHandlerSchemaUrn) Ptr() *EnumunboundidYubikeyOtpSaslMechanismHandlerSchemaUrn {
	return &v
}

type NullableEnumunboundidYubikeyOtpSaslMechanismHandlerSchemaUrn struct {
	value *EnumunboundidYubikeyOtpSaslMechanismHandlerSchemaUrn
	isSet bool
}

func (v NullableEnumunboundidYubikeyOtpSaslMechanismHandlerSchemaUrn) Get() *EnumunboundidYubikeyOtpSaslMechanismHandlerSchemaUrn {
	return v.value
}

func (v *NullableEnumunboundidYubikeyOtpSaslMechanismHandlerSchemaUrn) Set(val *EnumunboundidYubikeyOtpSaslMechanismHandlerSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumunboundidYubikeyOtpSaslMechanismHandlerSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumunboundidYubikeyOtpSaslMechanismHandlerSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumunboundidYubikeyOtpSaslMechanismHandlerSchemaUrn(val *EnumunboundidYubikeyOtpSaslMechanismHandlerSchemaUrn) *NullableEnumunboundidYubikeyOtpSaslMechanismHandlerSchemaUrn {
	return &NullableEnumunboundidYubikeyOtpSaslMechanismHandlerSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumunboundidYubikeyOtpSaslMechanismHandlerSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumunboundidYubikeyOtpSaslMechanismHandlerSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
