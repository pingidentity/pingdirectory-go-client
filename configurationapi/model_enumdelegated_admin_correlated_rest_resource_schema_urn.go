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

// EnumdelegatedAdminCorrelatedRestResourceSchemaUrn the model 'EnumdelegatedAdminCorrelatedRestResourceSchemaUrn'
type EnumdelegatedAdminCorrelatedRestResourceSchemaUrn string

// List of Enumdelegated-admin-correlated-rest-resourceSchemaUrn
const (
	ENUMDELEGATEDADMINCORRELATEDRESTRESOURCESCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0DELEGATED_ADMIN_CORRELATED_REST_RESOURCE EnumdelegatedAdminCorrelatedRestResourceSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:delegated-admin-correlated-rest-resource"
)

// All allowed values of EnumdelegatedAdminCorrelatedRestResourceSchemaUrn enum
var AllowedEnumdelegatedAdminCorrelatedRestResourceSchemaUrnEnumValues = []EnumdelegatedAdminCorrelatedRestResourceSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:delegated-admin-correlated-rest-resource",
}

func (v *EnumdelegatedAdminCorrelatedRestResourceSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumdelegatedAdminCorrelatedRestResourceSchemaUrn(value)
	for _, existing := range AllowedEnumdelegatedAdminCorrelatedRestResourceSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumdelegatedAdminCorrelatedRestResourceSchemaUrn", value)
}

// NewEnumdelegatedAdminCorrelatedRestResourceSchemaUrnFromValue returns a pointer to a valid EnumdelegatedAdminCorrelatedRestResourceSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumdelegatedAdminCorrelatedRestResourceSchemaUrnFromValue(v string) (*EnumdelegatedAdminCorrelatedRestResourceSchemaUrn, error) {
	ev := EnumdelegatedAdminCorrelatedRestResourceSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumdelegatedAdminCorrelatedRestResourceSchemaUrn: valid values are %v", v, AllowedEnumdelegatedAdminCorrelatedRestResourceSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumdelegatedAdminCorrelatedRestResourceSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumdelegatedAdminCorrelatedRestResourceSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumdelegated-admin-correlated-rest-resourceSchemaUrn value
func (v EnumdelegatedAdminCorrelatedRestResourceSchemaUrn) Ptr() *EnumdelegatedAdminCorrelatedRestResourceSchemaUrn {
	return &v
}

type NullableEnumdelegatedAdminCorrelatedRestResourceSchemaUrn struct {
	value *EnumdelegatedAdminCorrelatedRestResourceSchemaUrn
	isSet bool
}

func (v NullableEnumdelegatedAdminCorrelatedRestResourceSchemaUrn) Get() *EnumdelegatedAdminCorrelatedRestResourceSchemaUrn {
	return v.value
}

func (v *NullableEnumdelegatedAdminCorrelatedRestResourceSchemaUrn) Set(val *EnumdelegatedAdminCorrelatedRestResourceSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumdelegatedAdminCorrelatedRestResourceSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumdelegatedAdminCorrelatedRestResourceSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumdelegatedAdminCorrelatedRestResourceSchemaUrn(val *EnumdelegatedAdminCorrelatedRestResourceSchemaUrn) *NullableEnumdelegatedAdminCorrelatedRestResourceSchemaUrn {
	return &NullableEnumdelegatedAdminCorrelatedRestResourceSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumdelegatedAdminCorrelatedRestResourceSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumdelegatedAdminCorrelatedRestResourceSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
