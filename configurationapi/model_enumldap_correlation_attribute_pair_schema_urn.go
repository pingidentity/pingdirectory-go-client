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

// EnumldapCorrelationAttributePairSchemaUrn the model 'EnumldapCorrelationAttributePairSchemaUrn'
type EnumldapCorrelationAttributePairSchemaUrn string

// List of Enumldap-correlation-attribute-pairSchemaUrn
const (
	ENUMLDAPCORRELATIONATTRIBUTEPAIRSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0LDAP_CORRELATION_ATTRIBUTE_PAIR EnumldapCorrelationAttributePairSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:ldap-correlation-attribute-pair"
)

// All allowed values of EnumldapCorrelationAttributePairSchemaUrn enum
var AllowedEnumldapCorrelationAttributePairSchemaUrnEnumValues = []EnumldapCorrelationAttributePairSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:ldap-correlation-attribute-pair",
}

func (v *EnumldapCorrelationAttributePairSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumldapCorrelationAttributePairSchemaUrn(value)
	for _, existing := range AllowedEnumldapCorrelationAttributePairSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumldapCorrelationAttributePairSchemaUrn", value)
}

// NewEnumldapCorrelationAttributePairSchemaUrnFromValue returns a pointer to a valid EnumldapCorrelationAttributePairSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumldapCorrelationAttributePairSchemaUrnFromValue(v string) (*EnumldapCorrelationAttributePairSchemaUrn, error) {
	ev := EnumldapCorrelationAttributePairSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumldapCorrelationAttributePairSchemaUrn: valid values are %v", v, AllowedEnumldapCorrelationAttributePairSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumldapCorrelationAttributePairSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumldapCorrelationAttributePairSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumldap-correlation-attribute-pairSchemaUrn value
func (v EnumldapCorrelationAttributePairSchemaUrn) Ptr() *EnumldapCorrelationAttributePairSchemaUrn {
	return &v
}

type NullableEnumldapCorrelationAttributePairSchemaUrn struct {
	value *EnumldapCorrelationAttributePairSchemaUrn
	isSet bool
}

func (v NullableEnumldapCorrelationAttributePairSchemaUrn) Get() *EnumldapCorrelationAttributePairSchemaUrn {
	return v.value
}

func (v *NullableEnumldapCorrelationAttributePairSchemaUrn) Set(val *EnumldapCorrelationAttributePairSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumldapCorrelationAttributePairSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumldapCorrelationAttributePairSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumldapCorrelationAttributePairSchemaUrn(val *EnumldapCorrelationAttributePairSchemaUrn) *NullableEnumldapCorrelationAttributePairSchemaUrn {
	return &NullableEnumldapCorrelationAttributePairSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumldapCorrelationAttributePairSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumldapCorrelationAttributePairSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
