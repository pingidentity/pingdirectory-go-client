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

// EnumthirdPartyOtpDeliveryMechanismSchemaUrn the model 'EnumthirdPartyOtpDeliveryMechanismSchemaUrn'
type EnumthirdPartyOtpDeliveryMechanismSchemaUrn string

// List of Enumthird-party-otp-delivery-mechanismSchemaUrn
const (
	ENUMTHIRDPARTYOTPDELIVERYMECHANISMSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0OTP_DELIVERY_MECHANISMTHIRD_PARTY EnumthirdPartyOtpDeliveryMechanismSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:otp-delivery-mechanism:third-party"
)

// All allowed values of EnumthirdPartyOtpDeliveryMechanismSchemaUrn enum
var AllowedEnumthirdPartyOtpDeliveryMechanismSchemaUrnEnumValues = []EnumthirdPartyOtpDeliveryMechanismSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:otp-delivery-mechanism:third-party",
}

func (v *EnumthirdPartyOtpDeliveryMechanismSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumthirdPartyOtpDeliveryMechanismSchemaUrn(value)
	for _, existing := range AllowedEnumthirdPartyOtpDeliveryMechanismSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumthirdPartyOtpDeliveryMechanismSchemaUrn", value)
}

// NewEnumthirdPartyOtpDeliveryMechanismSchemaUrnFromValue returns a pointer to a valid EnumthirdPartyOtpDeliveryMechanismSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumthirdPartyOtpDeliveryMechanismSchemaUrnFromValue(v string) (*EnumthirdPartyOtpDeliveryMechanismSchemaUrn, error) {
	ev := EnumthirdPartyOtpDeliveryMechanismSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumthirdPartyOtpDeliveryMechanismSchemaUrn: valid values are %v", v, AllowedEnumthirdPartyOtpDeliveryMechanismSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumthirdPartyOtpDeliveryMechanismSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumthirdPartyOtpDeliveryMechanismSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumthird-party-otp-delivery-mechanismSchemaUrn value
func (v EnumthirdPartyOtpDeliveryMechanismSchemaUrn) Ptr() *EnumthirdPartyOtpDeliveryMechanismSchemaUrn {
	return &v
}

type NullableEnumthirdPartyOtpDeliveryMechanismSchemaUrn struct {
	value *EnumthirdPartyOtpDeliveryMechanismSchemaUrn
	isSet bool
}

func (v NullableEnumthirdPartyOtpDeliveryMechanismSchemaUrn) Get() *EnumthirdPartyOtpDeliveryMechanismSchemaUrn {
	return v.value
}

func (v *NullableEnumthirdPartyOtpDeliveryMechanismSchemaUrn) Set(val *EnumthirdPartyOtpDeliveryMechanismSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumthirdPartyOtpDeliveryMechanismSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumthirdPartyOtpDeliveryMechanismSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumthirdPartyOtpDeliveryMechanismSchemaUrn(val *EnumthirdPartyOtpDeliveryMechanismSchemaUrn) *NullableEnumthirdPartyOtpDeliveryMechanismSchemaUrn {
	return &NullableEnumthirdPartyOtpDeliveryMechanismSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumthirdPartyOtpDeliveryMechanismSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumthirdPartyOtpDeliveryMechanismSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
