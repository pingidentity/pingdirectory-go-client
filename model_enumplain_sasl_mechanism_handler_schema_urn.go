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

// EnumplainSaslMechanismHandlerSchemaUrn the model 'EnumplainSaslMechanismHandlerSchemaUrn'
type EnumplainSaslMechanismHandlerSchemaUrn string

// List of Enumplain-sasl-mechanism-handlerSchemaUrn
const (
	ENUMPLAINSASLMECHANISMHANDLERSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0SASL_MECHANISM_HANDLERPLAIN EnumplainSaslMechanismHandlerSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:sasl-mechanism-handler:plain"
)

// All allowed values of EnumplainSaslMechanismHandlerSchemaUrn enum
var AllowedEnumplainSaslMechanismHandlerSchemaUrnEnumValues = []EnumplainSaslMechanismHandlerSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:sasl-mechanism-handler:plain",
}

func (v *EnumplainSaslMechanismHandlerSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumplainSaslMechanismHandlerSchemaUrn(value)
	for _, existing := range AllowedEnumplainSaslMechanismHandlerSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumplainSaslMechanismHandlerSchemaUrn", value)
}

// NewEnumplainSaslMechanismHandlerSchemaUrnFromValue returns a pointer to a valid EnumplainSaslMechanismHandlerSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumplainSaslMechanismHandlerSchemaUrnFromValue(v string) (*EnumplainSaslMechanismHandlerSchemaUrn, error) {
	ev := EnumplainSaslMechanismHandlerSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumplainSaslMechanismHandlerSchemaUrn: valid values are %v", v, AllowedEnumplainSaslMechanismHandlerSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumplainSaslMechanismHandlerSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumplainSaslMechanismHandlerSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumplain-sasl-mechanism-handlerSchemaUrn value
func (v EnumplainSaslMechanismHandlerSchemaUrn) Ptr() *EnumplainSaslMechanismHandlerSchemaUrn {
	return &v
}

type NullableEnumplainSaslMechanismHandlerSchemaUrn struct {
	value *EnumplainSaslMechanismHandlerSchemaUrn
	isSet bool
}

func (v NullableEnumplainSaslMechanismHandlerSchemaUrn) Get() *EnumplainSaslMechanismHandlerSchemaUrn {
	return v.value
}

func (v *NullableEnumplainSaslMechanismHandlerSchemaUrn) Set(val *EnumplainSaslMechanismHandlerSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumplainSaslMechanismHandlerSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumplainSaslMechanismHandlerSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumplainSaslMechanismHandlerSchemaUrn(val *EnumplainSaslMechanismHandlerSchemaUrn) *NullableEnumplainSaslMechanismHandlerSchemaUrn {
	return &NullableEnumplainSaslMechanismHandlerSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumplainSaslMechanismHandlerSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumplainSaslMechanismHandlerSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

