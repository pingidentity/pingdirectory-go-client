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

// EnumthirdPartyOauthTokenHandlerSchemaUrn the model 'EnumthirdPartyOauthTokenHandlerSchemaUrn'
type EnumthirdPartyOauthTokenHandlerSchemaUrn string

// List of Enumthird-party-oauth-token-handlerSchemaUrn
const (
	ENUMTHIRDPARTYOAUTHTOKENHANDLERSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0OAUTH_TOKEN_HANDLERTHIRD_PARTY EnumthirdPartyOauthTokenHandlerSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:oauth-token-handler:third-party"
)

// All allowed values of EnumthirdPartyOauthTokenHandlerSchemaUrn enum
var AllowedEnumthirdPartyOauthTokenHandlerSchemaUrnEnumValues = []EnumthirdPartyOauthTokenHandlerSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:oauth-token-handler:third-party",
}

func (v *EnumthirdPartyOauthTokenHandlerSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumthirdPartyOauthTokenHandlerSchemaUrn(value)
	for _, existing := range AllowedEnumthirdPartyOauthTokenHandlerSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumthirdPartyOauthTokenHandlerSchemaUrn", value)
}

// NewEnumthirdPartyOauthTokenHandlerSchemaUrnFromValue returns a pointer to a valid EnumthirdPartyOauthTokenHandlerSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumthirdPartyOauthTokenHandlerSchemaUrnFromValue(v string) (*EnumthirdPartyOauthTokenHandlerSchemaUrn, error) {
	ev := EnumthirdPartyOauthTokenHandlerSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumthirdPartyOauthTokenHandlerSchemaUrn: valid values are %v", v, AllowedEnumthirdPartyOauthTokenHandlerSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumthirdPartyOauthTokenHandlerSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumthirdPartyOauthTokenHandlerSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumthird-party-oauth-token-handlerSchemaUrn value
func (v EnumthirdPartyOauthTokenHandlerSchemaUrn) Ptr() *EnumthirdPartyOauthTokenHandlerSchemaUrn {
	return &v
}

type NullableEnumthirdPartyOauthTokenHandlerSchemaUrn struct {
	value *EnumthirdPartyOauthTokenHandlerSchemaUrn
	isSet bool
}

func (v NullableEnumthirdPartyOauthTokenHandlerSchemaUrn) Get() *EnumthirdPartyOauthTokenHandlerSchemaUrn {
	return v.value
}

func (v *NullableEnumthirdPartyOauthTokenHandlerSchemaUrn) Set(val *EnumthirdPartyOauthTokenHandlerSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumthirdPartyOauthTokenHandlerSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumthirdPartyOauthTokenHandlerSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumthirdPartyOauthTokenHandlerSchemaUrn(val *EnumthirdPartyOauthTokenHandlerSchemaUrn) *NullableEnumthirdPartyOauthTokenHandlerSchemaUrn {
	return &NullableEnumthirdPartyOauthTokenHandlerSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumthirdPartyOauthTokenHandlerSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumthirdPartyOauthTokenHandlerSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
