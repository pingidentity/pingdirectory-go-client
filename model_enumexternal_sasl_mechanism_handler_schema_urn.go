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

// EnumexternalSaslMechanismHandlerSchemaUrn the model 'EnumexternalSaslMechanismHandlerSchemaUrn'
type EnumexternalSaslMechanismHandlerSchemaUrn string

// List of Enumexternal-sasl-mechanism-handlerSchemaUrn
const (
	ENUMEXTERNALSASLMECHANISMHANDLERSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0SASL_MECHANISM_HANDLEREXTERNAL EnumexternalSaslMechanismHandlerSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:sasl-mechanism-handler:external"
)

// All allowed values of EnumexternalSaslMechanismHandlerSchemaUrn enum
var AllowedEnumexternalSaslMechanismHandlerSchemaUrnEnumValues = []EnumexternalSaslMechanismHandlerSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:sasl-mechanism-handler:external",
}

func (v *EnumexternalSaslMechanismHandlerSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumexternalSaslMechanismHandlerSchemaUrn(value)
	for _, existing := range AllowedEnumexternalSaslMechanismHandlerSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumexternalSaslMechanismHandlerSchemaUrn", value)
}

// NewEnumexternalSaslMechanismHandlerSchemaUrnFromValue returns a pointer to a valid EnumexternalSaslMechanismHandlerSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumexternalSaslMechanismHandlerSchemaUrnFromValue(v string) (*EnumexternalSaslMechanismHandlerSchemaUrn, error) {
	ev := EnumexternalSaslMechanismHandlerSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumexternalSaslMechanismHandlerSchemaUrn: valid values are %v", v, AllowedEnumexternalSaslMechanismHandlerSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumexternalSaslMechanismHandlerSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumexternalSaslMechanismHandlerSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumexternal-sasl-mechanism-handlerSchemaUrn value
func (v EnumexternalSaslMechanismHandlerSchemaUrn) Ptr() *EnumexternalSaslMechanismHandlerSchemaUrn {
	return &v
}

type NullableEnumexternalSaslMechanismHandlerSchemaUrn struct {
	value *EnumexternalSaslMechanismHandlerSchemaUrn
	isSet bool
}

func (v NullableEnumexternalSaslMechanismHandlerSchemaUrn) Get() *EnumexternalSaslMechanismHandlerSchemaUrn {
	return v.value
}

func (v *NullableEnumexternalSaslMechanismHandlerSchemaUrn) Set(val *EnumexternalSaslMechanismHandlerSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumexternalSaslMechanismHandlerSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumexternalSaslMechanismHandlerSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumexternalSaslMechanismHandlerSchemaUrn(val *EnumexternalSaslMechanismHandlerSchemaUrn) *NullableEnumexternalSaslMechanismHandlerSchemaUrn {
	return &NullableEnumexternalSaslMechanismHandlerSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumexternalSaslMechanismHandlerSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumexternalSaslMechanismHandlerSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
