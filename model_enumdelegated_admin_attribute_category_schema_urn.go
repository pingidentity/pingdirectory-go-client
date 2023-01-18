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

// EnumdelegatedAdminAttributeCategorySchemaUrn the model 'EnumdelegatedAdminAttributeCategorySchemaUrn'
type EnumdelegatedAdminAttributeCategorySchemaUrn string

// List of Enumdelegated-admin-attribute-categorySchemaUrn
const (
	ENUMDELEGATEDADMINATTRIBUTECATEGORYSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0DELEGATED_ADMIN_ATTRIBUTE_CATEGORY EnumdelegatedAdminAttributeCategorySchemaUrn = "urn:pingidentity:schemas:configuration:2.0:delegated-admin-attribute-category"
)

// All allowed values of EnumdelegatedAdminAttributeCategorySchemaUrn enum
var AllowedEnumdelegatedAdminAttributeCategorySchemaUrnEnumValues = []EnumdelegatedAdminAttributeCategorySchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:delegated-admin-attribute-category",
}

func (v *EnumdelegatedAdminAttributeCategorySchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumdelegatedAdminAttributeCategorySchemaUrn(value)
	for _, existing := range AllowedEnumdelegatedAdminAttributeCategorySchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumdelegatedAdminAttributeCategorySchemaUrn", value)
}

// NewEnumdelegatedAdminAttributeCategorySchemaUrnFromValue returns a pointer to a valid EnumdelegatedAdminAttributeCategorySchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumdelegatedAdminAttributeCategorySchemaUrnFromValue(v string) (*EnumdelegatedAdminAttributeCategorySchemaUrn, error) {
	ev := EnumdelegatedAdminAttributeCategorySchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumdelegatedAdminAttributeCategorySchemaUrn: valid values are %v", v, AllowedEnumdelegatedAdminAttributeCategorySchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumdelegatedAdminAttributeCategorySchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumdelegatedAdminAttributeCategorySchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumdelegated-admin-attribute-categorySchemaUrn value
func (v EnumdelegatedAdminAttributeCategorySchemaUrn) Ptr() *EnumdelegatedAdminAttributeCategorySchemaUrn {
	return &v
}

type NullableEnumdelegatedAdminAttributeCategorySchemaUrn struct {
	value *EnumdelegatedAdminAttributeCategorySchemaUrn
	isSet bool
}

func (v NullableEnumdelegatedAdminAttributeCategorySchemaUrn) Get() *EnumdelegatedAdminAttributeCategorySchemaUrn {
	return v.value
}

func (v *NullableEnumdelegatedAdminAttributeCategorySchemaUrn) Set(val *EnumdelegatedAdminAttributeCategorySchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumdelegatedAdminAttributeCategorySchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumdelegatedAdminAttributeCategorySchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumdelegatedAdminAttributeCategorySchemaUrn(val *EnumdelegatedAdminAttributeCategorySchemaUrn) *NullableEnumdelegatedAdminAttributeCategorySchemaUrn {
	return &NullableEnumdelegatedAdminAttributeCategorySchemaUrn{value: val, isSet: true}
}

func (v NullableEnumdelegatedAdminAttributeCategorySchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumdelegatedAdminAttributeCategorySchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
