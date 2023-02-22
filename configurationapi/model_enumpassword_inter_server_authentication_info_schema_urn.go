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

// EnumpasswordInterServerAuthenticationInfoSchemaUrn the model 'EnumpasswordInterServerAuthenticationInfoSchemaUrn'
type EnumpasswordInterServerAuthenticationInfoSchemaUrn string

// List of Enumpassword-inter-server-authentication-infoSchemaUrn
const (
	ENUMPASSWORDINTERSERVERAUTHENTICATIONINFOSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0INTER_SERVER_AUTHENTICATION_INFOPASSWORD EnumpasswordInterServerAuthenticationInfoSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:inter-server-authentication-info:password"
)

// All allowed values of EnumpasswordInterServerAuthenticationInfoSchemaUrn enum
var AllowedEnumpasswordInterServerAuthenticationInfoSchemaUrnEnumValues = []EnumpasswordInterServerAuthenticationInfoSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:inter-server-authentication-info:password",
}

func (v *EnumpasswordInterServerAuthenticationInfoSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumpasswordInterServerAuthenticationInfoSchemaUrn(value)
	for _, existing := range AllowedEnumpasswordInterServerAuthenticationInfoSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumpasswordInterServerAuthenticationInfoSchemaUrn", value)
}

// NewEnumpasswordInterServerAuthenticationInfoSchemaUrnFromValue returns a pointer to a valid EnumpasswordInterServerAuthenticationInfoSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumpasswordInterServerAuthenticationInfoSchemaUrnFromValue(v string) (*EnumpasswordInterServerAuthenticationInfoSchemaUrn, error) {
	ev := EnumpasswordInterServerAuthenticationInfoSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumpasswordInterServerAuthenticationInfoSchemaUrn: valid values are %v", v, AllowedEnumpasswordInterServerAuthenticationInfoSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumpasswordInterServerAuthenticationInfoSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumpasswordInterServerAuthenticationInfoSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumpassword-inter-server-authentication-infoSchemaUrn value
func (v EnumpasswordInterServerAuthenticationInfoSchemaUrn) Ptr() *EnumpasswordInterServerAuthenticationInfoSchemaUrn {
	return &v
}

type NullableEnumpasswordInterServerAuthenticationInfoSchemaUrn struct {
	value *EnumpasswordInterServerAuthenticationInfoSchemaUrn
	isSet bool
}

func (v NullableEnumpasswordInterServerAuthenticationInfoSchemaUrn) Get() *EnumpasswordInterServerAuthenticationInfoSchemaUrn {
	return v.value
}

func (v *NullableEnumpasswordInterServerAuthenticationInfoSchemaUrn) Set(val *EnumpasswordInterServerAuthenticationInfoSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumpasswordInterServerAuthenticationInfoSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumpasswordInterServerAuthenticationInfoSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumpasswordInterServerAuthenticationInfoSchemaUrn(val *EnumpasswordInterServerAuthenticationInfoSchemaUrn) *NullableEnumpasswordInterServerAuthenticationInfoSchemaUrn {
	return &NullableEnumpasswordInterServerAuthenticationInfoSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumpasswordInterServerAuthenticationInfoSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumpasswordInterServerAuthenticationInfoSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}