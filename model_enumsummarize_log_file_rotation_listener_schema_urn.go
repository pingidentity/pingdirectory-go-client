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

// EnumsummarizeLogFileRotationListenerSchemaUrn the model 'EnumsummarizeLogFileRotationListenerSchemaUrn'
type EnumsummarizeLogFileRotationListenerSchemaUrn string

// List of Enumsummarize-log-file-rotation-listenerSchemaUrn
const (
	ENUMSUMMARIZELOGFILEROTATIONLISTENERSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0LOG_FILE_ROTATION_LISTENERSUMMARIZE EnumsummarizeLogFileRotationListenerSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:log-file-rotation-listener:summarize"
)

// All allowed values of EnumsummarizeLogFileRotationListenerSchemaUrn enum
var AllowedEnumsummarizeLogFileRotationListenerSchemaUrnEnumValues = []EnumsummarizeLogFileRotationListenerSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:log-file-rotation-listener:summarize",
}

func (v *EnumsummarizeLogFileRotationListenerSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumsummarizeLogFileRotationListenerSchemaUrn(value)
	for _, existing := range AllowedEnumsummarizeLogFileRotationListenerSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumsummarizeLogFileRotationListenerSchemaUrn", value)
}

// NewEnumsummarizeLogFileRotationListenerSchemaUrnFromValue returns a pointer to a valid EnumsummarizeLogFileRotationListenerSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumsummarizeLogFileRotationListenerSchemaUrnFromValue(v string) (*EnumsummarizeLogFileRotationListenerSchemaUrn, error) {
	ev := EnumsummarizeLogFileRotationListenerSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumsummarizeLogFileRotationListenerSchemaUrn: valid values are %v", v, AllowedEnumsummarizeLogFileRotationListenerSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumsummarizeLogFileRotationListenerSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumsummarizeLogFileRotationListenerSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumsummarize-log-file-rotation-listenerSchemaUrn value
func (v EnumsummarizeLogFileRotationListenerSchemaUrn) Ptr() *EnumsummarizeLogFileRotationListenerSchemaUrn {
	return &v
}

type NullableEnumsummarizeLogFileRotationListenerSchemaUrn struct {
	value *EnumsummarizeLogFileRotationListenerSchemaUrn
	isSet bool
}

func (v NullableEnumsummarizeLogFileRotationListenerSchemaUrn) Get() *EnumsummarizeLogFileRotationListenerSchemaUrn {
	return v.value
}

func (v *NullableEnumsummarizeLogFileRotationListenerSchemaUrn) Set(val *EnumsummarizeLogFileRotationListenerSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumsummarizeLogFileRotationListenerSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumsummarizeLogFileRotationListenerSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumsummarizeLogFileRotationListenerSchemaUrn(val *EnumsummarizeLogFileRotationListenerSchemaUrn) *NullableEnumsummarizeLogFileRotationListenerSchemaUrn {
	return &NullableEnumsummarizeLogFileRotationListenerSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumsummarizeLogFileRotationListenerSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumsummarizeLogFileRotationListenerSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
