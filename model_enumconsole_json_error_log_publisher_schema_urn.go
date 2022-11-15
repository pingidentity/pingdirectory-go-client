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

// EnumconsoleJsonErrorLogPublisherSchemaUrn the model 'EnumconsoleJsonErrorLogPublisherSchemaUrn'
type EnumconsoleJsonErrorLogPublisherSchemaUrn string

// List of Enumconsole-json-error-log-publisherSchemaUrn
const (
	ENUMCONSOLEJSONERRORLOGPUBLISHERSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0LOG_PUBLISHERCONSOLE_JSON_ERROR EnumconsoleJsonErrorLogPublisherSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:log-publisher:console-json-error"
)

// All allowed values of EnumconsoleJsonErrorLogPublisherSchemaUrn enum
var AllowedEnumconsoleJsonErrorLogPublisherSchemaUrnEnumValues = []EnumconsoleJsonErrorLogPublisherSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:log-publisher:console-json-error",
}

func (v *EnumconsoleJsonErrorLogPublisherSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumconsoleJsonErrorLogPublisherSchemaUrn(value)
	for _, existing := range AllowedEnumconsoleJsonErrorLogPublisherSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumconsoleJsonErrorLogPublisherSchemaUrn", value)
}

// NewEnumconsoleJsonErrorLogPublisherSchemaUrnFromValue returns a pointer to a valid EnumconsoleJsonErrorLogPublisherSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumconsoleJsonErrorLogPublisherSchemaUrnFromValue(v string) (*EnumconsoleJsonErrorLogPublisherSchemaUrn, error) {
	ev := EnumconsoleJsonErrorLogPublisherSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumconsoleJsonErrorLogPublisherSchemaUrn: valid values are %v", v, AllowedEnumconsoleJsonErrorLogPublisherSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumconsoleJsonErrorLogPublisherSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumconsoleJsonErrorLogPublisherSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumconsole-json-error-log-publisherSchemaUrn value
func (v EnumconsoleJsonErrorLogPublisherSchemaUrn) Ptr() *EnumconsoleJsonErrorLogPublisherSchemaUrn {
	return &v
}

type NullableEnumconsoleJsonErrorLogPublisherSchemaUrn struct {
	value *EnumconsoleJsonErrorLogPublisherSchemaUrn
	isSet bool
}

func (v NullableEnumconsoleJsonErrorLogPublisherSchemaUrn) Get() *EnumconsoleJsonErrorLogPublisherSchemaUrn {
	return v.value
}

func (v *NullableEnumconsoleJsonErrorLogPublisherSchemaUrn) Set(val *EnumconsoleJsonErrorLogPublisherSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumconsoleJsonErrorLogPublisherSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumconsoleJsonErrorLogPublisherSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumconsoleJsonErrorLogPublisherSchemaUrn(val *EnumconsoleJsonErrorLogPublisherSchemaUrn) *NullableEnumconsoleJsonErrorLogPublisherSchemaUrn {
	return &NullableEnumconsoleJsonErrorLogPublisherSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumconsoleJsonErrorLogPublisherSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumconsoleJsonErrorLogPublisherSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

