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

// EnumlastAccessTimePluginSchemaUrn the model 'EnumlastAccessTimePluginSchemaUrn'
type EnumlastAccessTimePluginSchemaUrn string

// List of Enumlast-access-time-pluginSchemaUrn
const (
	ENUMLASTACCESSTIMEPLUGINSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0PLUGINLAST_ACCESS_TIME EnumlastAccessTimePluginSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:plugin:last-access-time"
)

// All allowed values of EnumlastAccessTimePluginSchemaUrn enum
var AllowedEnumlastAccessTimePluginSchemaUrnEnumValues = []EnumlastAccessTimePluginSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:plugin:last-access-time",
}

func (v *EnumlastAccessTimePluginSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumlastAccessTimePluginSchemaUrn(value)
	for _, existing := range AllowedEnumlastAccessTimePluginSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumlastAccessTimePluginSchemaUrn", value)
}

// NewEnumlastAccessTimePluginSchemaUrnFromValue returns a pointer to a valid EnumlastAccessTimePluginSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumlastAccessTimePluginSchemaUrnFromValue(v string) (*EnumlastAccessTimePluginSchemaUrn, error) {
	ev := EnumlastAccessTimePluginSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumlastAccessTimePluginSchemaUrn: valid values are %v", v, AllowedEnumlastAccessTimePluginSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumlastAccessTimePluginSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumlastAccessTimePluginSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumlast-access-time-pluginSchemaUrn value
func (v EnumlastAccessTimePluginSchemaUrn) Ptr() *EnumlastAccessTimePluginSchemaUrn {
	return &v
}

type NullableEnumlastAccessTimePluginSchemaUrn struct {
	value *EnumlastAccessTimePluginSchemaUrn
	isSet bool
}

func (v NullableEnumlastAccessTimePluginSchemaUrn) Get() *EnumlastAccessTimePluginSchemaUrn {
	return v.value
}

func (v *NullableEnumlastAccessTimePluginSchemaUrn) Set(val *EnumlastAccessTimePluginSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumlastAccessTimePluginSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumlastAccessTimePluginSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumlastAccessTimePluginSchemaUrn(val *EnumlastAccessTimePluginSchemaUrn) *NullableEnumlastAccessTimePluginSchemaUrn {
	return &NullableEnumlastAccessTimePluginSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumlastAccessTimePluginSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumlastAccessTimePluginSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}