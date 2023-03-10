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

// EnumclientConnectionPolicySchemaUrn the model 'EnumclientConnectionPolicySchemaUrn'
type EnumclientConnectionPolicySchemaUrn string

// List of Enumclient-connection-policySchemaUrn
const (
	ENUMCLIENTCONNECTIONPOLICYSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0CLIENT_CONNECTION_POLICY EnumclientConnectionPolicySchemaUrn = "urn:pingidentity:schemas:configuration:2.0:client-connection-policy"
)

// All allowed values of EnumclientConnectionPolicySchemaUrn enum
var AllowedEnumclientConnectionPolicySchemaUrnEnumValues = []EnumclientConnectionPolicySchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:client-connection-policy",
}

func (v *EnumclientConnectionPolicySchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumclientConnectionPolicySchemaUrn(value)
	for _, existing := range AllowedEnumclientConnectionPolicySchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumclientConnectionPolicySchemaUrn", value)
}

// NewEnumclientConnectionPolicySchemaUrnFromValue returns a pointer to a valid EnumclientConnectionPolicySchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumclientConnectionPolicySchemaUrnFromValue(v string) (*EnumclientConnectionPolicySchemaUrn, error) {
	ev := EnumclientConnectionPolicySchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumclientConnectionPolicySchemaUrn: valid values are %v", v, AllowedEnumclientConnectionPolicySchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumclientConnectionPolicySchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumclientConnectionPolicySchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumclient-connection-policySchemaUrn value
func (v EnumclientConnectionPolicySchemaUrn) Ptr() *EnumclientConnectionPolicySchemaUrn {
	return &v
}

type NullableEnumclientConnectionPolicySchemaUrn struct {
	value *EnumclientConnectionPolicySchemaUrn
	isSet bool
}

func (v NullableEnumclientConnectionPolicySchemaUrn) Get() *EnumclientConnectionPolicySchemaUrn {
	return v.value
}

func (v *NullableEnumclientConnectionPolicySchemaUrn) Set(val *EnumclientConnectionPolicySchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumclientConnectionPolicySchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumclientConnectionPolicySchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumclientConnectionPolicySchemaUrn(val *EnumclientConnectionPolicySchemaUrn) *NullableEnumclientConnectionPolicySchemaUrn {
	return &NullableEnumclientConnectionPolicySchemaUrn{value: val, isSet: true}
}

func (v NullableEnumclientConnectionPolicySchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumclientConnectionPolicySchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
