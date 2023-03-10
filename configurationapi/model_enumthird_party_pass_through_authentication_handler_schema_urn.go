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

// EnumthirdPartyPassThroughAuthenticationHandlerSchemaUrn the model 'EnumthirdPartyPassThroughAuthenticationHandlerSchemaUrn'
type EnumthirdPartyPassThroughAuthenticationHandlerSchemaUrn string

// List of Enumthird-party-pass-through-authentication-handlerSchemaUrn
const (
	ENUMTHIRDPARTYPASSTHROUGHAUTHENTICATIONHANDLERSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0PASS_THROUGH_AUTHENTICATION_HANDLERTHIRD_PARTY EnumthirdPartyPassThroughAuthenticationHandlerSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:pass-through-authentication-handler:third-party"
)

// All allowed values of EnumthirdPartyPassThroughAuthenticationHandlerSchemaUrn enum
var AllowedEnumthirdPartyPassThroughAuthenticationHandlerSchemaUrnEnumValues = []EnumthirdPartyPassThroughAuthenticationHandlerSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:pass-through-authentication-handler:third-party",
}

func (v *EnumthirdPartyPassThroughAuthenticationHandlerSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumthirdPartyPassThroughAuthenticationHandlerSchemaUrn(value)
	for _, existing := range AllowedEnumthirdPartyPassThroughAuthenticationHandlerSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumthirdPartyPassThroughAuthenticationHandlerSchemaUrn", value)
}

// NewEnumthirdPartyPassThroughAuthenticationHandlerSchemaUrnFromValue returns a pointer to a valid EnumthirdPartyPassThroughAuthenticationHandlerSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumthirdPartyPassThroughAuthenticationHandlerSchemaUrnFromValue(v string) (*EnumthirdPartyPassThroughAuthenticationHandlerSchemaUrn, error) {
	ev := EnumthirdPartyPassThroughAuthenticationHandlerSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumthirdPartyPassThroughAuthenticationHandlerSchemaUrn: valid values are %v", v, AllowedEnumthirdPartyPassThroughAuthenticationHandlerSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumthirdPartyPassThroughAuthenticationHandlerSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumthirdPartyPassThroughAuthenticationHandlerSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumthird-party-pass-through-authentication-handlerSchemaUrn value
func (v EnumthirdPartyPassThroughAuthenticationHandlerSchemaUrn) Ptr() *EnumthirdPartyPassThroughAuthenticationHandlerSchemaUrn {
	return &v
}

type NullableEnumthirdPartyPassThroughAuthenticationHandlerSchemaUrn struct {
	value *EnumthirdPartyPassThroughAuthenticationHandlerSchemaUrn
	isSet bool
}

func (v NullableEnumthirdPartyPassThroughAuthenticationHandlerSchemaUrn) Get() *EnumthirdPartyPassThroughAuthenticationHandlerSchemaUrn {
	return v.value
}

func (v *NullableEnumthirdPartyPassThroughAuthenticationHandlerSchemaUrn) Set(val *EnumthirdPartyPassThroughAuthenticationHandlerSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumthirdPartyPassThroughAuthenticationHandlerSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumthirdPartyPassThroughAuthenticationHandlerSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumthirdPartyPassThroughAuthenticationHandlerSchemaUrn(val *EnumthirdPartyPassThroughAuthenticationHandlerSchemaUrn) *NullableEnumthirdPartyPassThroughAuthenticationHandlerSchemaUrn {
	return &NullableEnumthirdPartyPassThroughAuthenticationHandlerSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumthirdPartyPassThroughAuthenticationHandlerSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumthirdPartyPassThroughAuthenticationHandlerSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
