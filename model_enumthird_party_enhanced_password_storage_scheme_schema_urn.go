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

// EnumthirdPartyEnhancedPasswordStorageSchemeSchemaUrn the model 'EnumthirdPartyEnhancedPasswordStorageSchemeSchemaUrn'
type EnumthirdPartyEnhancedPasswordStorageSchemeSchemaUrn string

// List of Enumthird-party-enhanced-password-storage-schemeSchemaUrn
const (
	ENUMTHIRDPARTYENHANCEDPASSWORDSTORAGESCHEMESCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0PASSWORD_STORAGE_SCHEMETHIRD_PARTY_ENHANCED EnumthirdPartyEnhancedPasswordStorageSchemeSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:password-storage-scheme:third-party-enhanced"
)

// All allowed values of EnumthirdPartyEnhancedPasswordStorageSchemeSchemaUrn enum
var AllowedEnumthirdPartyEnhancedPasswordStorageSchemeSchemaUrnEnumValues = []EnumthirdPartyEnhancedPasswordStorageSchemeSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:password-storage-scheme:third-party-enhanced",
}

func (v *EnumthirdPartyEnhancedPasswordStorageSchemeSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumthirdPartyEnhancedPasswordStorageSchemeSchemaUrn(value)
	for _, existing := range AllowedEnumthirdPartyEnhancedPasswordStorageSchemeSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumthirdPartyEnhancedPasswordStorageSchemeSchemaUrn", value)
}

// NewEnumthirdPartyEnhancedPasswordStorageSchemeSchemaUrnFromValue returns a pointer to a valid EnumthirdPartyEnhancedPasswordStorageSchemeSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumthirdPartyEnhancedPasswordStorageSchemeSchemaUrnFromValue(v string) (*EnumthirdPartyEnhancedPasswordStorageSchemeSchemaUrn, error) {
	ev := EnumthirdPartyEnhancedPasswordStorageSchemeSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumthirdPartyEnhancedPasswordStorageSchemeSchemaUrn: valid values are %v", v, AllowedEnumthirdPartyEnhancedPasswordStorageSchemeSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumthirdPartyEnhancedPasswordStorageSchemeSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumthirdPartyEnhancedPasswordStorageSchemeSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumthird-party-enhanced-password-storage-schemeSchemaUrn value
func (v EnumthirdPartyEnhancedPasswordStorageSchemeSchemaUrn) Ptr() *EnumthirdPartyEnhancedPasswordStorageSchemeSchemaUrn {
	return &v
}

type NullableEnumthirdPartyEnhancedPasswordStorageSchemeSchemaUrn struct {
	value *EnumthirdPartyEnhancedPasswordStorageSchemeSchemaUrn
	isSet bool
}

func (v NullableEnumthirdPartyEnhancedPasswordStorageSchemeSchemaUrn) Get() *EnumthirdPartyEnhancedPasswordStorageSchemeSchemaUrn {
	return v.value
}

func (v *NullableEnumthirdPartyEnhancedPasswordStorageSchemeSchemaUrn) Set(val *EnumthirdPartyEnhancedPasswordStorageSchemeSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumthirdPartyEnhancedPasswordStorageSchemeSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumthirdPartyEnhancedPasswordStorageSchemeSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumthirdPartyEnhancedPasswordStorageSchemeSchemaUrn(val *EnumthirdPartyEnhancedPasswordStorageSchemeSchemaUrn) *NullableEnumthirdPartyEnhancedPasswordStorageSchemeSchemaUrn {
	return &NullableEnumthirdPartyEnhancedPasswordStorageSchemeSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumthirdPartyEnhancedPasswordStorageSchemeSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumthirdPartyEnhancedPasswordStorageSchemeSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

