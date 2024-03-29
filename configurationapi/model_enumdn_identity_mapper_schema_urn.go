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

// EnumdnIdentityMapperSchemaUrn the model 'EnumdnIdentityMapperSchemaUrn'
type EnumdnIdentityMapperSchemaUrn string

// List of Enumdn-identity-mapperSchemaUrn
const (
	ENUMDNIDENTITYMAPPERSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0IDENTITY_MAPPERDN EnumdnIdentityMapperSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:identity-mapper:dn"
)

// All allowed values of EnumdnIdentityMapperSchemaUrn enum
var AllowedEnumdnIdentityMapperSchemaUrnEnumValues = []EnumdnIdentityMapperSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:identity-mapper:dn",
}

func (v *EnumdnIdentityMapperSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumdnIdentityMapperSchemaUrn(value)
	for _, existing := range AllowedEnumdnIdentityMapperSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumdnIdentityMapperSchemaUrn", value)
}

// NewEnumdnIdentityMapperSchemaUrnFromValue returns a pointer to a valid EnumdnIdentityMapperSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumdnIdentityMapperSchemaUrnFromValue(v string) (*EnumdnIdentityMapperSchemaUrn, error) {
	ev := EnumdnIdentityMapperSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumdnIdentityMapperSchemaUrn: valid values are %v", v, AllowedEnumdnIdentityMapperSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumdnIdentityMapperSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumdnIdentityMapperSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumdn-identity-mapperSchemaUrn value
func (v EnumdnIdentityMapperSchemaUrn) Ptr() *EnumdnIdentityMapperSchemaUrn {
	return &v
}

type NullableEnumdnIdentityMapperSchemaUrn struct {
	value *EnumdnIdentityMapperSchemaUrn
	isSet bool
}

func (v NullableEnumdnIdentityMapperSchemaUrn) Get() *EnumdnIdentityMapperSchemaUrn {
	return v.value
}

func (v *NullableEnumdnIdentityMapperSchemaUrn) Set(val *EnumdnIdentityMapperSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumdnIdentityMapperSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumdnIdentityMapperSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumdnIdentityMapperSchemaUrn(val *EnumdnIdentityMapperSchemaUrn) *NullableEnumdnIdentityMapperSchemaUrn {
	return &NullableEnumdnIdentityMapperSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumdnIdentityMapperSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumdnIdentityMapperSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
