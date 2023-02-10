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

// EnumthirdPartyVelocityContextProviderSchemaUrn the model 'EnumthirdPartyVelocityContextProviderSchemaUrn'
type EnumthirdPartyVelocityContextProviderSchemaUrn string

// List of Enumthird-party-velocity-context-providerSchemaUrn
const (
	ENUMTHIRDPARTYVELOCITYCONTEXTPROVIDERSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0VELOCITY_CONTEXT_PROVIDERTHIRD_PARTY EnumthirdPartyVelocityContextProviderSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:velocity-context-provider:third-party"
)

// All allowed values of EnumthirdPartyVelocityContextProviderSchemaUrn enum
var AllowedEnumthirdPartyVelocityContextProviderSchemaUrnEnumValues = []EnumthirdPartyVelocityContextProviderSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:velocity-context-provider:third-party",
}

func (v *EnumthirdPartyVelocityContextProviderSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumthirdPartyVelocityContextProviderSchemaUrn(value)
	for _, existing := range AllowedEnumthirdPartyVelocityContextProviderSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumthirdPartyVelocityContextProviderSchemaUrn", value)
}

// NewEnumthirdPartyVelocityContextProviderSchemaUrnFromValue returns a pointer to a valid EnumthirdPartyVelocityContextProviderSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumthirdPartyVelocityContextProviderSchemaUrnFromValue(v string) (*EnumthirdPartyVelocityContextProviderSchemaUrn, error) {
	ev := EnumthirdPartyVelocityContextProviderSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumthirdPartyVelocityContextProviderSchemaUrn: valid values are %v", v, AllowedEnumthirdPartyVelocityContextProviderSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumthirdPartyVelocityContextProviderSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumthirdPartyVelocityContextProviderSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumthird-party-velocity-context-providerSchemaUrn value
func (v EnumthirdPartyVelocityContextProviderSchemaUrn) Ptr() *EnumthirdPartyVelocityContextProviderSchemaUrn {
	return &v
}

type NullableEnumthirdPartyVelocityContextProviderSchemaUrn struct {
	value *EnumthirdPartyVelocityContextProviderSchemaUrn
	isSet bool
}

func (v NullableEnumthirdPartyVelocityContextProviderSchemaUrn) Get() *EnumthirdPartyVelocityContextProviderSchemaUrn {
	return v.value
}

func (v *NullableEnumthirdPartyVelocityContextProviderSchemaUrn) Set(val *EnumthirdPartyVelocityContextProviderSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumthirdPartyVelocityContextProviderSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumthirdPartyVelocityContextProviderSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumthirdPartyVelocityContextProviderSchemaUrn(val *EnumthirdPartyVelocityContextProviderSchemaUrn) *NullableEnumthirdPartyVelocityContextProviderSchemaUrn {
	return &NullableEnumthirdPartyVelocityContextProviderSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumthirdPartyVelocityContextProviderSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumthirdPartyVelocityContextProviderSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
