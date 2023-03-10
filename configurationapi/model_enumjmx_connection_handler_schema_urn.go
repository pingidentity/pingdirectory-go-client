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

// EnumjmxConnectionHandlerSchemaUrn the model 'EnumjmxConnectionHandlerSchemaUrn'
type EnumjmxConnectionHandlerSchemaUrn string

// List of Enumjmx-connection-handlerSchemaUrn
const (
	ENUMJMXCONNECTIONHANDLERSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0CONNECTION_HANDLERJMX EnumjmxConnectionHandlerSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:connection-handler:jmx"
)

// All allowed values of EnumjmxConnectionHandlerSchemaUrn enum
var AllowedEnumjmxConnectionHandlerSchemaUrnEnumValues = []EnumjmxConnectionHandlerSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:connection-handler:jmx",
}

func (v *EnumjmxConnectionHandlerSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumjmxConnectionHandlerSchemaUrn(value)
	for _, existing := range AllowedEnumjmxConnectionHandlerSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumjmxConnectionHandlerSchemaUrn", value)
}

// NewEnumjmxConnectionHandlerSchemaUrnFromValue returns a pointer to a valid EnumjmxConnectionHandlerSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumjmxConnectionHandlerSchemaUrnFromValue(v string) (*EnumjmxConnectionHandlerSchemaUrn, error) {
	ev := EnumjmxConnectionHandlerSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumjmxConnectionHandlerSchemaUrn: valid values are %v", v, AllowedEnumjmxConnectionHandlerSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumjmxConnectionHandlerSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumjmxConnectionHandlerSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumjmx-connection-handlerSchemaUrn value
func (v EnumjmxConnectionHandlerSchemaUrn) Ptr() *EnumjmxConnectionHandlerSchemaUrn {
	return &v
}

type NullableEnumjmxConnectionHandlerSchemaUrn struct {
	value *EnumjmxConnectionHandlerSchemaUrn
	isSet bool
}

func (v NullableEnumjmxConnectionHandlerSchemaUrn) Get() *EnumjmxConnectionHandlerSchemaUrn {
	return v.value
}

func (v *NullableEnumjmxConnectionHandlerSchemaUrn) Set(val *EnumjmxConnectionHandlerSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumjmxConnectionHandlerSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumjmxConnectionHandlerSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumjmxConnectionHandlerSchemaUrn(val *EnumjmxConnectionHandlerSchemaUrn) *NullableEnumjmxConnectionHandlerSchemaUrn {
	return &NullableEnumjmxConnectionHandlerSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumjmxConnectionHandlerSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumjmxConnectionHandlerSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
