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

// EnumfileBasedErrorLogPublisherSchemaUrn the model 'EnumfileBasedErrorLogPublisherSchemaUrn'
type EnumfileBasedErrorLogPublisherSchemaUrn string

// List of Enumfile-based-error-log-publisherSchemaUrn
const (
	ENUMFILEBASEDERRORLOGPUBLISHERSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0LOG_PUBLISHERFILE_BASED_ERROR EnumfileBasedErrorLogPublisherSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:log-publisher:file-based-error"
)

// All allowed values of EnumfileBasedErrorLogPublisherSchemaUrn enum
var AllowedEnumfileBasedErrorLogPublisherSchemaUrnEnumValues = []EnumfileBasedErrorLogPublisherSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:log-publisher:file-based-error",
}

func (v *EnumfileBasedErrorLogPublisherSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumfileBasedErrorLogPublisherSchemaUrn(value)
	for _, existing := range AllowedEnumfileBasedErrorLogPublisherSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumfileBasedErrorLogPublisherSchemaUrn", value)
}

// NewEnumfileBasedErrorLogPublisherSchemaUrnFromValue returns a pointer to a valid EnumfileBasedErrorLogPublisherSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumfileBasedErrorLogPublisherSchemaUrnFromValue(v string) (*EnumfileBasedErrorLogPublisherSchemaUrn, error) {
	ev := EnumfileBasedErrorLogPublisherSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumfileBasedErrorLogPublisherSchemaUrn: valid values are %v", v, AllowedEnumfileBasedErrorLogPublisherSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumfileBasedErrorLogPublisherSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumfileBasedErrorLogPublisherSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumfile-based-error-log-publisherSchemaUrn value
func (v EnumfileBasedErrorLogPublisherSchemaUrn) Ptr() *EnumfileBasedErrorLogPublisherSchemaUrn {
	return &v
}

type NullableEnumfileBasedErrorLogPublisherSchemaUrn struct {
	value *EnumfileBasedErrorLogPublisherSchemaUrn
	isSet bool
}

func (v NullableEnumfileBasedErrorLogPublisherSchemaUrn) Get() *EnumfileBasedErrorLogPublisherSchemaUrn {
	return v.value
}

func (v *NullableEnumfileBasedErrorLogPublisherSchemaUrn) Set(val *EnumfileBasedErrorLogPublisherSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumfileBasedErrorLogPublisherSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumfileBasedErrorLogPublisherSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumfileBasedErrorLogPublisherSchemaUrn(val *EnumfileBasedErrorLogPublisherSchemaUrn) *NullableEnumfileBasedErrorLogPublisherSchemaUrn {
	return &NullableEnumfileBasedErrorLogPublisherSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumfileBasedErrorLogPublisherSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumfileBasedErrorLogPublisherSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
