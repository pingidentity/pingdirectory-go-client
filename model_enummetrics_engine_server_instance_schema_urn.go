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

// EnummetricsEngineServerInstanceSchemaUrn the model 'EnummetricsEngineServerInstanceSchemaUrn'
type EnummetricsEngineServerInstanceSchemaUrn string

// List of Enummetrics-engine-server-instanceSchemaUrn
const (
	URNPINGIDENTITYSCHEMASCONFIGURATION2_0SERVER_INSTANCEMETRICS_ENGINE EnummetricsEngineServerInstanceSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:server-instance:metrics-engine"
)

// All allowed values of EnummetricsEngineServerInstanceSchemaUrn enum
var AllowedEnummetricsEngineServerInstanceSchemaUrnEnumValues = []EnummetricsEngineServerInstanceSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:server-instance:metrics-engine",
}

func (v *EnummetricsEngineServerInstanceSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnummetricsEngineServerInstanceSchemaUrn(value)
	for _, existing := range AllowedEnummetricsEngineServerInstanceSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnummetricsEngineServerInstanceSchemaUrn", value)
}

// NewEnummetricsEngineServerInstanceSchemaUrnFromValue returns a pointer to a valid EnummetricsEngineServerInstanceSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnummetricsEngineServerInstanceSchemaUrnFromValue(v string) (*EnummetricsEngineServerInstanceSchemaUrn, error) {
	ev := EnummetricsEngineServerInstanceSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnummetricsEngineServerInstanceSchemaUrn: valid values are %v", v, AllowedEnummetricsEngineServerInstanceSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnummetricsEngineServerInstanceSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnummetricsEngineServerInstanceSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enummetrics-engine-server-instanceSchemaUrn value
func (v EnummetricsEngineServerInstanceSchemaUrn) Ptr() *EnummetricsEngineServerInstanceSchemaUrn {
	return &v
}

type NullableEnummetricsEngineServerInstanceSchemaUrn struct {
	value *EnummetricsEngineServerInstanceSchemaUrn
	isSet bool
}

func (v NullableEnummetricsEngineServerInstanceSchemaUrn) Get() *EnummetricsEngineServerInstanceSchemaUrn {
	return v.value
}

func (v *NullableEnummetricsEngineServerInstanceSchemaUrn) Set(val *EnummetricsEngineServerInstanceSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnummetricsEngineServerInstanceSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnummetricsEngineServerInstanceSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnummetricsEngineServerInstanceSchemaUrn(val *EnummetricsEngineServerInstanceSchemaUrn) *NullableEnummetricsEngineServerInstanceSchemaUrn {
	return &NullableEnummetricsEngineServerInstanceSchemaUrn{value: val, isSet: true}
}

func (v NullableEnummetricsEngineServerInstanceSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnummetricsEngineServerInstanceSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

