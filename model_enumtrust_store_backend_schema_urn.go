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

// EnumtrustStoreBackendSchemaUrn the model 'EnumtrustStoreBackendSchemaUrn'
type EnumtrustStoreBackendSchemaUrn string

// List of Enumtrust-store-backendSchemaUrn
const (
	ENUMTRUSTSTOREBACKENDSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0BACKENDTRUST_STORE EnumtrustStoreBackendSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:backend:trust-store"
)

// All allowed values of EnumtrustStoreBackendSchemaUrn enum
var AllowedEnumtrustStoreBackendSchemaUrnEnumValues = []EnumtrustStoreBackendSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:backend:trust-store",
}

func (v *EnumtrustStoreBackendSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumtrustStoreBackendSchemaUrn(value)
	for _, existing := range AllowedEnumtrustStoreBackendSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumtrustStoreBackendSchemaUrn", value)
}

// NewEnumtrustStoreBackendSchemaUrnFromValue returns a pointer to a valid EnumtrustStoreBackendSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumtrustStoreBackendSchemaUrnFromValue(v string) (*EnumtrustStoreBackendSchemaUrn, error) {
	ev := EnumtrustStoreBackendSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumtrustStoreBackendSchemaUrn: valid values are %v", v, AllowedEnumtrustStoreBackendSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumtrustStoreBackendSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumtrustStoreBackendSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumtrust-store-backendSchemaUrn value
func (v EnumtrustStoreBackendSchemaUrn) Ptr() *EnumtrustStoreBackendSchemaUrn {
	return &v
}

type NullableEnumtrustStoreBackendSchemaUrn struct {
	value *EnumtrustStoreBackendSchemaUrn
	isSet bool
}

func (v NullableEnumtrustStoreBackendSchemaUrn) Get() *EnumtrustStoreBackendSchemaUrn {
	return v.value
}

func (v *NullableEnumtrustStoreBackendSchemaUrn) Set(val *EnumtrustStoreBackendSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumtrustStoreBackendSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumtrustStoreBackendSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumtrustStoreBackendSchemaUrn(val *EnumtrustStoreBackendSchemaUrn) *NullableEnumtrustStoreBackendSchemaUrn {
	return &NullableEnumtrustStoreBackendSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumtrustStoreBackendSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumtrustStoreBackendSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

