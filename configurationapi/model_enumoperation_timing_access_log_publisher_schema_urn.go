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

// EnumoperationTimingAccessLogPublisherSchemaUrn the model 'EnumoperationTimingAccessLogPublisherSchemaUrn'
type EnumoperationTimingAccessLogPublisherSchemaUrn string

// List of Enumoperation-timing-access-log-publisherSchemaUrn
const (
	ENUMOPERATIONTIMINGACCESSLOGPUBLISHERSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0LOG_PUBLISHEROPERATION_TIMING_ACCESS EnumoperationTimingAccessLogPublisherSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:log-publisher:operation-timing-access"
)

// All allowed values of EnumoperationTimingAccessLogPublisherSchemaUrn enum
var AllowedEnumoperationTimingAccessLogPublisherSchemaUrnEnumValues = []EnumoperationTimingAccessLogPublisherSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:log-publisher:operation-timing-access",
}

func (v *EnumoperationTimingAccessLogPublisherSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumoperationTimingAccessLogPublisherSchemaUrn(value)
	for _, existing := range AllowedEnumoperationTimingAccessLogPublisherSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumoperationTimingAccessLogPublisherSchemaUrn", value)
}

// NewEnumoperationTimingAccessLogPublisherSchemaUrnFromValue returns a pointer to a valid EnumoperationTimingAccessLogPublisherSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumoperationTimingAccessLogPublisherSchemaUrnFromValue(v string) (*EnumoperationTimingAccessLogPublisherSchemaUrn, error) {
	ev := EnumoperationTimingAccessLogPublisherSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumoperationTimingAccessLogPublisherSchemaUrn: valid values are %v", v, AllowedEnumoperationTimingAccessLogPublisherSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumoperationTimingAccessLogPublisherSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumoperationTimingAccessLogPublisherSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumoperation-timing-access-log-publisherSchemaUrn value
func (v EnumoperationTimingAccessLogPublisherSchemaUrn) Ptr() *EnumoperationTimingAccessLogPublisherSchemaUrn {
	return &v
}

type NullableEnumoperationTimingAccessLogPublisherSchemaUrn struct {
	value *EnumoperationTimingAccessLogPublisherSchemaUrn
	isSet bool
}

func (v NullableEnumoperationTimingAccessLogPublisherSchemaUrn) Get() *EnumoperationTimingAccessLogPublisherSchemaUrn {
	return v.value
}

func (v *NullableEnumoperationTimingAccessLogPublisherSchemaUrn) Set(val *EnumoperationTimingAccessLogPublisherSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumoperationTimingAccessLogPublisherSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumoperationTimingAccessLogPublisherSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumoperationTimingAccessLogPublisherSchemaUrn(val *EnumoperationTimingAccessLogPublisherSchemaUrn) *NullableEnumoperationTimingAccessLogPublisherSchemaUrn {
	return &NullableEnumoperationTimingAccessLogPublisherSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumoperationTimingAccessLogPublisherSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumoperationTimingAccessLogPublisherSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
