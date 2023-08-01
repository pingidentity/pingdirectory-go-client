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

// EnumopenBankingHttpServletExtensionSchemaUrn the model 'EnumopenBankingHttpServletExtensionSchemaUrn'
type EnumopenBankingHttpServletExtensionSchemaUrn string

// List of Enumopen-banking-http-servlet-extensionSchemaUrn
const (
	ENUMOPENBANKINGHTTPSERVLETEXTENSIONSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0HTTP_SERVLET_EXTENSIONOPEN_BANKING EnumopenBankingHttpServletExtensionSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:http-servlet-extension:open-banking"
)

// All allowed values of EnumopenBankingHttpServletExtensionSchemaUrn enum
var AllowedEnumopenBankingHttpServletExtensionSchemaUrnEnumValues = []EnumopenBankingHttpServletExtensionSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:http-servlet-extension:open-banking",
}

func (v *EnumopenBankingHttpServletExtensionSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumopenBankingHttpServletExtensionSchemaUrn(value)
	for _, existing := range AllowedEnumopenBankingHttpServletExtensionSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumopenBankingHttpServletExtensionSchemaUrn", value)
}

// NewEnumopenBankingHttpServletExtensionSchemaUrnFromValue returns a pointer to a valid EnumopenBankingHttpServletExtensionSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumopenBankingHttpServletExtensionSchemaUrnFromValue(v string) (*EnumopenBankingHttpServletExtensionSchemaUrn, error) {
	ev := EnumopenBankingHttpServletExtensionSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumopenBankingHttpServletExtensionSchemaUrn: valid values are %v", v, AllowedEnumopenBankingHttpServletExtensionSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumopenBankingHttpServletExtensionSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumopenBankingHttpServletExtensionSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumopen-banking-http-servlet-extensionSchemaUrn value
func (v EnumopenBankingHttpServletExtensionSchemaUrn) Ptr() *EnumopenBankingHttpServletExtensionSchemaUrn {
	return &v
}

type NullableEnumopenBankingHttpServletExtensionSchemaUrn struct {
	value *EnumopenBankingHttpServletExtensionSchemaUrn
	isSet bool
}

func (v NullableEnumopenBankingHttpServletExtensionSchemaUrn) Get() *EnumopenBankingHttpServletExtensionSchemaUrn {
	return v.value
}

func (v *NullableEnumopenBankingHttpServletExtensionSchemaUrn) Set(val *EnumopenBankingHttpServletExtensionSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumopenBankingHttpServletExtensionSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumopenBankingHttpServletExtensionSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumopenBankingHttpServletExtensionSchemaUrn(val *EnumopenBankingHttpServletExtensionSchemaUrn) *NullableEnumopenBankingHttpServletExtensionSchemaUrn {
	return &NullableEnumopenBankingHttpServletExtensionSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumopenBankingHttpServletExtensionSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumopenBankingHttpServletExtensionSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
