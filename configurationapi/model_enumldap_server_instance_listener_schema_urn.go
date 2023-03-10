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

// EnumldapServerInstanceListenerSchemaUrn the model 'EnumldapServerInstanceListenerSchemaUrn'
type EnumldapServerInstanceListenerSchemaUrn string

// List of Enumldap-server-instance-listenerSchemaUrn
const (
	ENUMLDAPSERVERINSTANCELISTENERSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0SERVER_INSTANCE_LISTENERLDAP EnumldapServerInstanceListenerSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:server-instance-listener:ldap"
)

// All allowed values of EnumldapServerInstanceListenerSchemaUrn enum
var AllowedEnumldapServerInstanceListenerSchemaUrnEnumValues = []EnumldapServerInstanceListenerSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:server-instance-listener:ldap",
}

func (v *EnumldapServerInstanceListenerSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumldapServerInstanceListenerSchemaUrn(value)
	for _, existing := range AllowedEnumldapServerInstanceListenerSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumldapServerInstanceListenerSchemaUrn", value)
}

// NewEnumldapServerInstanceListenerSchemaUrnFromValue returns a pointer to a valid EnumldapServerInstanceListenerSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumldapServerInstanceListenerSchemaUrnFromValue(v string) (*EnumldapServerInstanceListenerSchemaUrn, error) {
	ev := EnumldapServerInstanceListenerSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumldapServerInstanceListenerSchemaUrn: valid values are %v", v, AllowedEnumldapServerInstanceListenerSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumldapServerInstanceListenerSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumldapServerInstanceListenerSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumldap-server-instance-listenerSchemaUrn value
func (v EnumldapServerInstanceListenerSchemaUrn) Ptr() *EnumldapServerInstanceListenerSchemaUrn {
	return &v
}

type NullableEnumldapServerInstanceListenerSchemaUrn struct {
	value *EnumldapServerInstanceListenerSchemaUrn
	isSet bool
}

func (v NullableEnumldapServerInstanceListenerSchemaUrn) Get() *EnumldapServerInstanceListenerSchemaUrn {
	return v.value
}

func (v *NullableEnumldapServerInstanceListenerSchemaUrn) Set(val *EnumldapServerInstanceListenerSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumldapServerInstanceListenerSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumldapServerInstanceListenerSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumldapServerInstanceListenerSchemaUrn(val *EnumldapServerInstanceListenerSchemaUrn) *NullableEnumldapServerInstanceListenerSchemaUrn {
	return &NullableEnumldapServerInstanceListenerSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumldapServerInstanceListenerSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumldapServerInstanceListenerSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
