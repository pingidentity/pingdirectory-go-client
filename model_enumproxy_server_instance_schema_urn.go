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

// EnumproxyServerInstanceSchemaUrn the model 'EnumproxyServerInstanceSchemaUrn'
type EnumproxyServerInstanceSchemaUrn string

// List of Enumproxy-server-instanceSchemaUrn
const (
	ENUMPROXYSERVERINSTANCESCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0SERVER_INSTANCEPROXY EnumproxyServerInstanceSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:server-instance:proxy"
)

// All allowed values of EnumproxyServerInstanceSchemaUrn enum
var AllowedEnumproxyServerInstanceSchemaUrnEnumValues = []EnumproxyServerInstanceSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:server-instance:proxy",
}

func (v *EnumproxyServerInstanceSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumproxyServerInstanceSchemaUrn(value)
	for _, existing := range AllowedEnumproxyServerInstanceSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumproxyServerInstanceSchemaUrn", value)
}

// NewEnumproxyServerInstanceSchemaUrnFromValue returns a pointer to a valid EnumproxyServerInstanceSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumproxyServerInstanceSchemaUrnFromValue(v string) (*EnumproxyServerInstanceSchemaUrn, error) {
	ev := EnumproxyServerInstanceSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumproxyServerInstanceSchemaUrn: valid values are %v", v, AllowedEnumproxyServerInstanceSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumproxyServerInstanceSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumproxyServerInstanceSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumproxy-server-instanceSchemaUrn value
func (v EnumproxyServerInstanceSchemaUrn) Ptr() *EnumproxyServerInstanceSchemaUrn {
	return &v
}

type NullableEnumproxyServerInstanceSchemaUrn struct {
	value *EnumproxyServerInstanceSchemaUrn
	isSet bool
}

func (v NullableEnumproxyServerInstanceSchemaUrn) Get() *EnumproxyServerInstanceSchemaUrn {
	return v.value
}

func (v *NullableEnumproxyServerInstanceSchemaUrn) Set(val *EnumproxyServerInstanceSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumproxyServerInstanceSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumproxyServerInstanceSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumproxyServerInstanceSchemaUrn(val *EnumproxyServerInstanceSchemaUrn) *NullableEnumproxyServerInstanceSchemaUrn {
	return &NullableEnumproxyServerInstanceSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumproxyServerInstanceSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumproxyServerInstanceSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

