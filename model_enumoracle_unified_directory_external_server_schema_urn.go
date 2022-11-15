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

// EnumoracleUnifiedDirectoryExternalServerSchemaUrn the model 'EnumoracleUnifiedDirectoryExternalServerSchemaUrn'
type EnumoracleUnifiedDirectoryExternalServerSchemaUrn string

// List of Enumoracle-unified-directory-external-serverSchemaUrn
const (
	ENUMORACLEUNIFIEDDIRECTORYEXTERNALSERVERSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0EXTERNAL_SERVERORACLE_UNIFIED_DIRECTORY EnumoracleUnifiedDirectoryExternalServerSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:external-server:oracle-unified-directory"
)

// All allowed values of EnumoracleUnifiedDirectoryExternalServerSchemaUrn enum
var AllowedEnumoracleUnifiedDirectoryExternalServerSchemaUrnEnumValues = []EnumoracleUnifiedDirectoryExternalServerSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:external-server:oracle-unified-directory",
}

func (v *EnumoracleUnifiedDirectoryExternalServerSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumoracleUnifiedDirectoryExternalServerSchemaUrn(value)
	for _, existing := range AllowedEnumoracleUnifiedDirectoryExternalServerSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumoracleUnifiedDirectoryExternalServerSchemaUrn", value)
}

// NewEnumoracleUnifiedDirectoryExternalServerSchemaUrnFromValue returns a pointer to a valid EnumoracleUnifiedDirectoryExternalServerSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumoracleUnifiedDirectoryExternalServerSchemaUrnFromValue(v string) (*EnumoracleUnifiedDirectoryExternalServerSchemaUrn, error) {
	ev := EnumoracleUnifiedDirectoryExternalServerSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumoracleUnifiedDirectoryExternalServerSchemaUrn: valid values are %v", v, AllowedEnumoracleUnifiedDirectoryExternalServerSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumoracleUnifiedDirectoryExternalServerSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumoracleUnifiedDirectoryExternalServerSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumoracle-unified-directory-external-serverSchemaUrn value
func (v EnumoracleUnifiedDirectoryExternalServerSchemaUrn) Ptr() *EnumoracleUnifiedDirectoryExternalServerSchemaUrn {
	return &v
}

type NullableEnumoracleUnifiedDirectoryExternalServerSchemaUrn struct {
	value *EnumoracleUnifiedDirectoryExternalServerSchemaUrn
	isSet bool
}

func (v NullableEnumoracleUnifiedDirectoryExternalServerSchemaUrn) Get() *EnumoracleUnifiedDirectoryExternalServerSchemaUrn {
	return v.value
}

func (v *NullableEnumoracleUnifiedDirectoryExternalServerSchemaUrn) Set(val *EnumoracleUnifiedDirectoryExternalServerSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumoracleUnifiedDirectoryExternalServerSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumoracleUnifiedDirectoryExternalServerSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumoracleUnifiedDirectoryExternalServerSchemaUrn(val *EnumoracleUnifiedDirectoryExternalServerSchemaUrn) *NullableEnumoracleUnifiedDirectoryExternalServerSchemaUrn {
	return &NullableEnumoracleUnifiedDirectoryExternalServerSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumoracleUnifiedDirectoryExternalServerSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumoracleUnifiedDirectoryExternalServerSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

