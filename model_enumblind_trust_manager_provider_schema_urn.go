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

// EnumblindTrustManagerProviderSchemaUrn the model 'EnumblindTrustManagerProviderSchemaUrn'
type EnumblindTrustManagerProviderSchemaUrn string

// List of Enumblind-trust-manager-providerSchemaUrn
const (
	ENUMBLINDTRUSTMANAGERPROVIDERSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0TRUST_MANAGER_PROVIDERBLIND EnumblindTrustManagerProviderSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:trust-manager-provider:blind"
)

// All allowed values of EnumblindTrustManagerProviderSchemaUrn enum
var AllowedEnumblindTrustManagerProviderSchemaUrnEnumValues = []EnumblindTrustManagerProviderSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:trust-manager-provider:blind",
}

func (v *EnumblindTrustManagerProviderSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumblindTrustManagerProviderSchemaUrn(value)
	for _, existing := range AllowedEnumblindTrustManagerProviderSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumblindTrustManagerProviderSchemaUrn", value)
}

// NewEnumblindTrustManagerProviderSchemaUrnFromValue returns a pointer to a valid EnumblindTrustManagerProviderSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumblindTrustManagerProviderSchemaUrnFromValue(v string) (*EnumblindTrustManagerProviderSchemaUrn, error) {
	ev := EnumblindTrustManagerProviderSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumblindTrustManagerProviderSchemaUrn: valid values are %v", v, AllowedEnumblindTrustManagerProviderSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumblindTrustManagerProviderSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumblindTrustManagerProviderSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumblind-trust-manager-providerSchemaUrn value
func (v EnumblindTrustManagerProviderSchemaUrn) Ptr() *EnumblindTrustManagerProviderSchemaUrn {
	return &v
}

type NullableEnumblindTrustManagerProviderSchemaUrn struct {
	value *EnumblindTrustManagerProviderSchemaUrn
	isSet bool
}

func (v NullableEnumblindTrustManagerProviderSchemaUrn) Get() *EnumblindTrustManagerProviderSchemaUrn {
	return v.value
}

func (v *NullableEnumblindTrustManagerProviderSchemaUrn) Set(val *EnumblindTrustManagerProviderSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumblindTrustManagerProviderSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumblindTrustManagerProviderSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumblindTrustManagerProviderSchemaUrn(val *EnumblindTrustManagerProviderSchemaUrn) *NullableEnumblindTrustManagerProviderSchemaUrn {
	return &NullableEnumblindTrustManagerProviderSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumblindTrustManagerProviderSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumblindTrustManagerProviderSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
