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

// EnumunboundidDeliveredOtpSaslMechanismHandlerSchemaUrn the model 'EnumunboundidDeliveredOtpSaslMechanismHandlerSchemaUrn'
type EnumunboundidDeliveredOtpSaslMechanismHandlerSchemaUrn string

// List of Enumunboundid-delivered-otp-sasl-mechanism-handlerSchemaUrn
const (
	URNPINGIDENTITYSCHEMASCONFIGURATION2_0SASL_MECHANISM_HANDLERUNBOUNDID_DELIVERED_OTP EnumunboundidDeliveredOtpSaslMechanismHandlerSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:sasl-mechanism-handler:unboundid-delivered-otp"
)

// All allowed values of EnumunboundidDeliveredOtpSaslMechanismHandlerSchemaUrn enum
var AllowedEnumunboundidDeliveredOtpSaslMechanismHandlerSchemaUrnEnumValues = []EnumunboundidDeliveredOtpSaslMechanismHandlerSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:sasl-mechanism-handler:unboundid-delivered-otp",
}

func (v *EnumunboundidDeliveredOtpSaslMechanismHandlerSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumunboundidDeliveredOtpSaslMechanismHandlerSchemaUrn(value)
	for _, existing := range AllowedEnumunboundidDeliveredOtpSaslMechanismHandlerSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumunboundidDeliveredOtpSaslMechanismHandlerSchemaUrn", value)
}

// NewEnumunboundidDeliveredOtpSaslMechanismHandlerSchemaUrnFromValue returns a pointer to a valid EnumunboundidDeliveredOtpSaslMechanismHandlerSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumunboundidDeliveredOtpSaslMechanismHandlerSchemaUrnFromValue(v string) (*EnumunboundidDeliveredOtpSaslMechanismHandlerSchemaUrn, error) {
	ev := EnumunboundidDeliveredOtpSaslMechanismHandlerSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumunboundidDeliveredOtpSaslMechanismHandlerSchemaUrn: valid values are %v", v, AllowedEnumunboundidDeliveredOtpSaslMechanismHandlerSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumunboundidDeliveredOtpSaslMechanismHandlerSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumunboundidDeliveredOtpSaslMechanismHandlerSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumunboundid-delivered-otp-sasl-mechanism-handlerSchemaUrn value
func (v EnumunboundidDeliveredOtpSaslMechanismHandlerSchemaUrn) Ptr() *EnumunboundidDeliveredOtpSaslMechanismHandlerSchemaUrn {
	return &v
}

type NullableEnumunboundidDeliveredOtpSaslMechanismHandlerSchemaUrn struct {
	value *EnumunboundidDeliveredOtpSaslMechanismHandlerSchemaUrn
	isSet bool
}

func (v NullableEnumunboundidDeliveredOtpSaslMechanismHandlerSchemaUrn) Get() *EnumunboundidDeliveredOtpSaslMechanismHandlerSchemaUrn {
	return v.value
}

func (v *NullableEnumunboundidDeliveredOtpSaslMechanismHandlerSchemaUrn) Set(val *EnumunboundidDeliveredOtpSaslMechanismHandlerSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumunboundidDeliveredOtpSaslMechanismHandlerSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumunboundidDeliveredOtpSaslMechanismHandlerSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumunboundidDeliveredOtpSaslMechanismHandlerSchemaUrn(val *EnumunboundidDeliveredOtpSaslMechanismHandlerSchemaUrn) *NullableEnumunboundidDeliveredOtpSaslMechanismHandlerSchemaUrn {
	return &NullableEnumunboundidDeliveredOtpSaslMechanismHandlerSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumunboundidDeliveredOtpSaslMechanismHandlerSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumunboundidDeliveredOtpSaslMechanismHandlerSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

