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

// EnumexternalServerJdbcDriverTypeProp Specifies a supported database driver type. The driver class will be automatically selected based on this selection. We highly recommend using a JDBC 4 driver that is suitable for the current Java platform.
type EnumexternalServerJdbcDriverTypeProp string

// List of Enumexternal-server-jdbcDriverTypeProp
const (
	ENUMEXTERNALSERVERJDBCDRIVERTYPEPROP_ORACLETHIN EnumexternalServerJdbcDriverTypeProp = "oraclethin"
	ENUMEXTERNALSERVERJDBCDRIVERTYPEPROP_ORACLEOCI  EnumexternalServerJdbcDriverTypeProp = "oracleoci"
	ENUMEXTERNALSERVERJDBCDRIVERTYPEPROP_MYSQL      EnumexternalServerJdbcDriverTypeProp = "mysql"
	ENUMEXTERNALSERVERJDBCDRIVERTYPEPROP_DB2        EnumexternalServerJdbcDriverTypeProp = "db2"
	ENUMEXTERNALSERVERJDBCDRIVERTYPEPROP_SQLSERVER  EnumexternalServerJdbcDriverTypeProp = "sqlserver"
	ENUMEXTERNALSERVERJDBCDRIVERTYPEPROP_POSTGRES   EnumexternalServerJdbcDriverTypeProp = "postgres"
	ENUMEXTERNALSERVERJDBCDRIVERTYPEPROP_OTHER      EnumexternalServerJdbcDriverTypeProp = "other"
)

// All allowed values of EnumexternalServerJdbcDriverTypeProp enum
var AllowedEnumexternalServerJdbcDriverTypePropEnumValues = []EnumexternalServerJdbcDriverTypeProp{
	"oraclethin",
	"oracleoci",
	"mysql",
	"db2",
	"sqlserver",
	"postgres",
	"other",
}

func (v *EnumexternalServerJdbcDriverTypeProp) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumexternalServerJdbcDriverTypeProp(value)
	for _, existing := range AllowedEnumexternalServerJdbcDriverTypePropEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumexternalServerJdbcDriverTypeProp", value)
}

// NewEnumexternalServerJdbcDriverTypePropFromValue returns a pointer to a valid EnumexternalServerJdbcDriverTypeProp
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumexternalServerJdbcDriverTypePropFromValue(v string) (*EnumexternalServerJdbcDriverTypeProp, error) {
	ev := EnumexternalServerJdbcDriverTypeProp(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumexternalServerJdbcDriverTypeProp: valid values are %v", v, AllowedEnumexternalServerJdbcDriverTypePropEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumexternalServerJdbcDriverTypeProp) IsValid() bool {
	for _, existing := range AllowedEnumexternalServerJdbcDriverTypePropEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumexternal-server-jdbcDriverTypeProp value
func (v EnumexternalServerJdbcDriverTypeProp) Ptr() *EnumexternalServerJdbcDriverTypeProp {
	return &v
}

type NullableEnumexternalServerJdbcDriverTypeProp struct {
	value *EnumexternalServerJdbcDriverTypeProp
	isSet bool
}

func (v NullableEnumexternalServerJdbcDriverTypeProp) Get() *EnumexternalServerJdbcDriverTypeProp {
	return v.value
}

func (v *NullableEnumexternalServerJdbcDriverTypeProp) Set(val *EnumexternalServerJdbcDriverTypeProp) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumexternalServerJdbcDriverTypeProp) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumexternalServerJdbcDriverTypeProp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumexternalServerJdbcDriverTypeProp(val *EnumexternalServerJdbcDriverTypeProp) *NullableEnumexternalServerJdbcDriverTypeProp {
	return &NullableEnumexternalServerJdbcDriverTypeProp{value: val, isSet: true}
}

func (v NullableEnumexternalServerJdbcDriverTypeProp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumexternalServerJdbcDriverTypeProp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
