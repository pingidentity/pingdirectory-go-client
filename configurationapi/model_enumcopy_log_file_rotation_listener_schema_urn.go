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

// EnumcopyLogFileRotationListenerSchemaUrn the model 'EnumcopyLogFileRotationListenerSchemaUrn'
type EnumcopyLogFileRotationListenerSchemaUrn string

// List of Enumcopy-log-file-rotation-listenerSchemaUrn
const (
	ENUMCOPYLOGFILEROTATIONLISTENERSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0LOG_FILE_ROTATION_LISTENERCOPY EnumcopyLogFileRotationListenerSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:log-file-rotation-listener:copy"
)

// All allowed values of EnumcopyLogFileRotationListenerSchemaUrn enum
var AllowedEnumcopyLogFileRotationListenerSchemaUrnEnumValues = []EnumcopyLogFileRotationListenerSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:log-file-rotation-listener:copy",
}

func (v *EnumcopyLogFileRotationListenerSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumcopyLogFileRotationListenerSchemaUrn(value)
	for _, existing := range AllowedEnumcopyLogFileRotationListenerSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumcopyLogFileRotationListenerSchemaUrn", value)
}

// NewEnumcopyLogFileRotationListenerSchemaUrnFromValue returns a pointer to a valid EnumcopyLogFileRotationListenerSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumcopyLogFileRotationListenerSchemaUrnFromValue(v string) (*EnumcopyLogFileRotationListenerSchemaUrn, error) {
	ev := EnumcopyLogFileRotationListenerSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumcopyLogFileRotationListenerSchemaUrn: valid values are %v", v, AllowedEnumcopyLogFileRotationListenerSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumcopyLogFileRotationListenerSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumcopyLogFileRotationListenerSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumcopy-log-file-rotation-listenerSchemaUrn value
func (v EnumcopyLogFileRotationListenerSchemaUrn) Ptr() *EnumcopyLogFileRotationListenerSchemaUrn {
	return &v
}

type NullableEnumcopyLogFileRotationListenerSchemaUrn struct {
	value *EnumcopyLogFileRotationListenerSchemaUrn
	isSet bool
}

func (v NullableEnumcopyLogFileRotationListenerSchemaUrn) Get() *EnumcopyLogFileRotationListenerSchemaUrn {
	return v.value
}

func (v *NullableEnumcopyLogFileRotationListenerSchemaUrn) Set(val *EnumcopyLogFileRotationListenerSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumcopyLogFileRotationListenerSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumcopyLogFileRotationListenerSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumcopyLogFileRotationListenerSchemaUrn(val *EnumcopyLogFileRotationListenerSchemaUrn) *NullableEnumcopyLogFileRotationListenerSchemaUrn {
	return &NullableEnumcopyLogFileRotationListenerSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumcopyLogFileRotationListenerSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumcopyLogFileRotationListenerSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
