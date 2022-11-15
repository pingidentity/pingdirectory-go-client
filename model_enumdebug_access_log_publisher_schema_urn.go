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

// EnumdebugAccessLogPublisherSchemaUrn the model 'EnumdebugAccessLogPublisherSchemaUrn'
type EnumdebugAccessLogPublisherSchemaUrn string

// List of Enumdebug-access-log-publisherSchemaUrn
const (
	URNPINGIDENTITYSCHEMASCONFIGURATION2_0LOG_PUBLISHERDEBUG_ACCESS EnumdebugAccessLogPublisherSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:log-publisher:debug-access"
)

// All allowed values of EnumdebugAccessLogPublisherSchemaUrn enum
var AllowedEnumdebugAccessLogPublisherSchemaUrnEnumValues = []EnumdebugAccessLogPublisherSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:log-publisher:debug-access",
}

func (v *EnumdebugAccessLogPublisherSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumdebugAccessLogPublisherSchemaUrn(value)
	for _, existing := range AllowedEnumdebugAccessLogPublisherSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumdebugAccessLogPublisherSchemaUrn", value)
}

// NewEnumdebugAccessLogPublisherSchemaUrnFromValue returns a pointer to a valid EnumdebugAccessLogPublisherSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumdebugAccessLogPublisherSchemaUrnFromValue(v string) (*EnumdebugAccessLogPublisherSchemaUrn, error) {
	ev := EnumdebugAccessLogPublisherSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumdebugAccessLogPublisherSchemaUrn: valid values are %v", v, AllowedEnumdebugAccessLogPublisherSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumdebugAccessLogPublisherSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumdebugAccessLogPublisherSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumdebug-access-log-publisherSchemaUrn value
func (v EnumdebugAccessLogPublisherSchemaUrn) Ptr() *EnumdebugAccessLogPublisherSchemaUrn {
	return &v
}

type NullableEnumdebugAccessLogPublisherSchemaUrn struct {
	value *EnumdebugAccessLogPublisherSchemaUrn
	isSet bool
}

func (v NullableEnumdebugAccessLogPublisherSchemaUrn) Get() *EnumdebugAccessLogPublisherSchemaUrn {
	return v.value
}

func (v *NullableEnumdebugAccessLogPublisherSchemaUrn) Set(val *EnumdebugAccessLogPublisherSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumdebugAccessLogPublisherSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumdebugAccessLogPublisherSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumdebugAccessLogPublisherSchemaUrn(val *EnumdebugAccessLogPublisherSchemaUrn) *NullableEnumdebugAccessLogPublisherSchemaUrn {
	return &NullableEnumdebugAccessLogPublisherSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumdebugAccessLogPublisherSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumdebugAccessLogPublisherSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

