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

// EnumanonymousSaslMechanismHandlerSchemaUrn the model 'EnumanonymousSaslMechanismHandlerSchemaUrn'
type EnumanonymousSaslMechanismHandlerSchemaUrn string

// List of Enumanonymous-sasl-mechanism-handlerSchemaUrn
const (
	URNPINGIDENTITYSCHEMASCONFIGURATION2_0SASL_MECHANISM_HANDLERANONYMOUS EnumanonymousSaslMechanismHandlerSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:sasl-mechanism-handler:anonymous"
)

// All allowed values of EnumanonymousSaslMechanismHandlerSchemaUrn enum
var AllowedEnumanonymousSaslMechanismHandlerSchemaUrnEnumValues = []EnumanonymousSaslMechanismHandlerSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:sasl-mechanism-handler:anonymous",
}

func (v *EnumanonymousSaslMechanismHandlerSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumanonymousSaslMechanismHandlerSchemaUrn(value)
	for _, existing := range AllowedEnumanonymousSaslMechanismHandlerSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumanonymousSaslMechanismHandlerSchemaUrn", value)
}

// NewEnumanonymousSaslMechanismHandlerSchemaUrnFromValue returns a pointer to a valid EnumanonymousSaslMechanismHandlerSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumanonymousSaslMechanismHandlerSchemaUrnFromValue(v string) (*EnumanonymousSaslMechanismHandlerSchemaUrn, error) {
	ev := EnumanonymousSaslMechanismHandlerSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumanonymousSaslMechanismHandlerSchemaUrn: valid values are %v", v, AllowedEnumanonymousSaslMechanismHandlerSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumanonymousSaslMechanismHandlerSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumanonymousSaslMechanismHandlerSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumanonymous-sasl-mechanism-handlerSchemaUrn value
func (v EnumanonymousSaslMechanismHandlerSchemaUrn) Ptr() *EnumanonymousSaslMechanismHandlerSchemaUrn {
	return &v
}

type NullableEnumanonymousSaslMechanismHandlerSchemaUrn struct {
	value *EnumanonymousSaslMechanismHandlerSchemaUrn
	isSet bool
}

func (v NullableEnumanonymousSaslMechanismHandlerSchemaUrn) Get() *EnumanonymousSaslMechanismHandlerSchemaUrn {
	return v.value
}

func (v *NullableEnumanonymousSaslMechanismHandlerSchemaUrn) Set(val *EnumanonymousSaslMechanismHandlerSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumanonymousSaslMechanismHandlerSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumanonymousSaslMechanismHandlerSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumanonymousSaslMechanismHandlerSchemaUrn(val *EnumanonymousSaslMechanismHandlerSchemaUrn) *NullableEnumanonymousSaslMechanismHandlerSchemaUrn {
	return &NullableEnumanonymousSaslMechanismHandlerSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumanonymousSaslMechanismHandlerSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumanonymousSaslMechanismHandlerSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

