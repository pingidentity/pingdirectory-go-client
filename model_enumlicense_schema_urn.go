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

// EnumlicenseSchemaUrn the model 'EnumlicenseSchemaUrn'
type EnumlicenseSchemaUrn string

// List of EnumlicenseSchemaUrn
const (
	ENUMLICENSESCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0LICENSE EnumlicenseSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:license"
)

// All allowed values of EnumlicenseSchemaUrn enum
var AllowedEnumlicenseSchemaUrnEnumValues = []EnumlicenseSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:license",
}

func (v *EnumlicenseSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumlicenseSchemaUrn(value)
	for _, existing := range AllowedEnumlicenseSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumlicenseSchemaUrn", value)
}

// NewEnumlicenseSchemaUrnFromValue returns a pointer to a valid EnumlicenseSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumlicenseSchemaUrnFromValue(v string) (*EnumlicenseSchemaUrn, error) {
	ev := EnumlicenseSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumlicenseSchemaUrn: valid values are %v", v, AllowedEnumlicenseSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumlicenseSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumlicenseSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EnumlicenseSchemaUrn value
func (v EnumlicenseSchemaUrn) Ptr() *EnumlicenseSchemaUrn {
	return &v
}

type NullableEnumlicenseSchemaUrn struct {
	value *EnumlicenseSchemaUrn
	isSet bool
}

func (v NullableEnumlicenseSchemaUrn) Get() *EnumlicenseSchemaUrn {
	return v.value
}

func (v *NullableEnumlicenseSchemaUrn) Set(val *EnumlicenseSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumlicenseSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumlicenseSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumlicenseSchemaUrn(val *EnumlicenseSchemaUrn) *NullableEnumlicenseSchemaUrn {
	return &NullableEnumlicenseSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumlicenseSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumlicenseSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

