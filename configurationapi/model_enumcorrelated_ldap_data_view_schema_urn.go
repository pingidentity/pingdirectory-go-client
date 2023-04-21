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

// EnumcorrelatedLdapDataViewSchemaUrn the model 'EnumcorrelatedLdapDataViewSchemaUrn'
type EnumcorrelatedLdapDataViewSchemaUrn string

// List of Enumcorrelated-ldap-data-viewSchemaUrn
const (
	ENUMCORRELATEDLDAPDATAVIEWSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0CORRELATED_LDAP_DATA_VIEW EnumcorrelatedLdapDataViewSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:correlated-ldap-data-view"
)

// All allowed values of EnumcorrelatedLdapDataViewSchemaUrn enum
var AllowedEnumcorrelatedLdapDataViewSchemaUrnEnumValues = []EnumcorrelatedLdapDataViewSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:correlated-ldap-data-view",
}

func (v *EnumcorrelatedLdapDataViewSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumcorrelatedLdapDataViewSchemaUrn(value)
	for _, existing := range AllowedEnumcorrelatedLdapDataViewSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumcorrelatedLdapDataViewSchemaUrn", value)
}

// NewEnumcorrelatedLdapDataViewSchemaUrnFromValue returns a pointer to a valid EnumcorrelatedLdapDataViewSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumcorrelatedLdapDataViewSchemaUrnFromValue(v string) (*EnumcorrelatedLdapDataViewSchemaUrn, error) {
	ev := EnumcorrelatedLdapDataViewSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumcorrelatedLdapDataViewSchemaUrn: valid values are %v", v, AllowedEnumcorrelatedLdapDataViewSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumcorrelatedLdapDataViewSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumcorrelatedLdapDataViewSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumcorrelated-ldap-data-viewSchemaUrn value
func (v EnumcorrelatedLdapDataViewSchemaUrn) Ptr() *EnumcorrelatedLdapDataViewSchemaUrn {
	return &v
}

type NullableEnumcorrelatedLdapDataViewSchemaUrn struct {
	value *EnumcorrelatedLdapDataViewSchemaUrn
	isSet bool
}

func (v NullableEnumcorrelatedLdapDataViewSchemaUrn) Get() *EnumcorrelatedLdapDataViewSchemaUrn {
	return v.value
}

func (v *NullableEnumcorrelatedLdapDataViewSchemaUrn) Set(val *EnumcorrelatedLdapDataViewSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumcorrelatedLdapDataViewSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumcorrelatedLdapDataViewSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumcorrelatedLdapDataViewSchemaUrn(val *EnumcorrelatedLdapDataViewSchemaUrn) *NullableEnumcorrelatedLdapDataViewSchemaUrn {
	return &NullableEnumcorrelatedLdapDataViewSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumcorrelatedLdapDataViewSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumcorrelatedLdapDataViewSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}