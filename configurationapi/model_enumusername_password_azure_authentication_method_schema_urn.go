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

// EnumusernamePasswordAzureAuthenticationMethodSchemaUrn the model 'EnumusernamePasswordAzureAuthenticationMethodSchemaUrn'
type EnumusernamePasswordAzureAuthenticationMethodSchemaUrn string

// List of Enumusername-password-azure-authentication-methodSchemaUrn
const (
	ENUMUSERNAMEPASSWORDAZUREAUTHENTICATIONMETHODSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0AZURE_AUTHENTICATION_METHODUSERNAME_PASSWORD EnumusernamePasswordAzureAuthenticationMethodSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:azure-authentication-method:username-password"
)

// All allowed values of EnumusernamePasswordAzureAuthenticationMethodSchemaUrn enum
var AllowedEnumusernamePasswordAzureAuthenticationMethodSchemaUrnEnumValues = []EnumusernamePasswordAzureAuthenticationMethodSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:azure-authentication-method:username-password",
}

func (v *EnumusernamePasswordAzureAuthenticationMethodSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumusernamePasswordAzureAuthenticationMethodSchemaUrn(value)
	for _, existing := range AllowedEnumusernamePasswordAzureAuthenticationMethodSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumusernamePasswordAzureAuthenticationMethodSchemaUrn", value)
}

// NewEnumusernamePasswordAzureAuthenticationMethodSchemaUrnFromValue returns a pointer to a valid EnumusernamePasswordAzureAuthenticationMethodSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumusernamePasswordAzureAuthenticationMethodSchemaUrnFromValue(v string) (*EnumusernamePasswordAzureAuthenticationMethodSchemaUrn, error) {
	ev := EnumusernamePasswordAzureAuthenticationMethodSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumusernamePasswordAzureAuthenticationMethodSchemaUrn: valid values are %v", v, AllowedEnumusernamePasswordAzureAuthenticationMethodSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumusernamePasswordAzureAuthenticationMethodSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumusernamePasswordAzureAuthenticationMethodSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumusername-password-azure-authentication-methodSchemaUrn value
func (v EnumusernamePasswordAzureAuthenticationMethodSchemaUrn) Ptr() *EnumusernamePasswordAzureAuthenticationMethodSchemaUrn {
	return &v
}

type NullableEnumusernamePasswordAzureAuthenticationMethodSchemaUrn struct {
	value *EnumusernamePasswordAzureAuthenticationMethodSchemaUrn
	isSet bool
}

func (v NullableEnumusernamePasswordAzureAuthenticationMethodSchemaUrn) Get() *EnumusernamePasswordAzureAuthenticationMethodSchemaUrn {
	return v.value
}

func (v *NullableEnumusernamePasswordAzureAuthenticationMethodSchemaUrn) Set(val *EnumusernamePasswordAzureAuthenticationMethodSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumusernamePasswordAzureAuthenticationMethodSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumusernamePasswordAzureAuthenticationMethodSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumusernamePasswordAzureAuthenticationMethodSchemaUrn(val *EnumusernamePasswordAzureAuthenticationMethodSchemaUrn) *NullableEnumusernamePasswordAzureAuthenticationMethodSchemaUrn {
	return &NullableEnumusernamePasswordAzureAuthenticationMethodSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumusernamePasswordAzureAuthenticationMethodSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumusernamePasswordAzureAuthenticationMethodSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
