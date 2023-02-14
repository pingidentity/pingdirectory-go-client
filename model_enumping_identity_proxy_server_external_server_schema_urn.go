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

// EnumpingIdentityProxyServerExternalServerSchemaUrn the model 'EnumpingIdentityProxyServerExternalServerSchemaUrn'
type EnumpingIdentityProxyServerExternalServerSchemaUrn string

// List of Enumping-identity-proxy-server-external-serverSchemaUrn
const (
	ENUMPINGIDENTITYPROXYSERVEREXTERNALSERVERSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0EXTERNAL_SERVERPING_IDENTITY_PROXY_SERVER EnumpingIdentityProxyServerExternalServerSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:external-server:ping-identity-proxy-server"
)

// All allowed values of EnumpingIdentityProxyServerExternalServerSchemaUrn enum
var AllowedEnumpingIdentityProxyServerExternalServerSchemaUrnEnumValues = []EnumpingIdentityProxyServerExternalServerSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:external-server:ping-identity-proxy-server",
}

func (v *EnumpingIdentityProxyServerExternalServerSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumpingIdentityProxyServerExternalServerSchemaUrn(value)
	for _, existing := range AllowedEnumpingIdentityProxyServerExternalServerSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumpingIdentityProxyServerExternalServerSchemaUrn", value)
}

// NewEnumpingIdentityProxyServerExternalServerSchemaUrnFromValue returns a pointer to a valid EnumpingIdentityProxyServerExternalServerSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumpingIdentityProxyServerExternalServerSchemaUrnFromValue(v string) (*EnumpingIdentityProxyServerExternalServerSchemaUrn, error) {
	ev := EnumpingIdentityProxyServerExternalServerSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumpingIdentityProxyServerExternalServerSchemaUrn: valid values are %v", v, AllowedEnumpingIdentityProxyServerExternalServerSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumpingIdentityProxyServerExternalServerSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumpingIdentityProxyServerExternalServerSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumping-identity-proxy-server-external-serverSchemaUrn value
func (v EnumpingIdentityProxyServerExternalServerSchemaUrn) Ptr() *EnumpingIdentityProxyServerExternalServerSchemaUrn {
	return &v
}

type NullableEnumpingIdentityProxyServerExternalServerSchemaUrn struct {
	value *EnumpingIdentityProxyServerExternalServerSchemaUrn
	isSet bool
}

func (v NullableEnumpingIdentityProxyServerExternalServerSchemaUrn) Get() *EnumpingIdentityProxyServerExternalServerSchemaUrn {
	return v.value
}

func (v *NullableEnumpingIdentityProxyServerExternalServerSchemaUrn) Set(val *EnumpingIdentityProxyServerExternalServerSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumpingIdentityProxyServerExternalServerSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumpingIdentityProxyServerExternalServerSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumpingIdentityProxyServerExternalServerSchemaUrn(val *EnumpingIdentityProxyServerExternalServerSchemaUrn) *NullableEnumpingIdentityProxyServerExternalServerSchemaUrn {
	return &NullableEnumpingIdentityProxyServerExternalServerSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumpingIdentityProxyServerExternalServerSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumpingIdentityProxyServerExternalServerSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
