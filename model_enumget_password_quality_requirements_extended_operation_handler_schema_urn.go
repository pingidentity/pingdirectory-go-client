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

// EnumgetPasswordQualityRequirementsExtendedOperationHandlerSchemaUrn the model 'EnumgetPasswordQualityRequirementsExtendedOperationHandlerSchemaUrn'
type EnumgetPasswordQualityRequirementsExtendedOperationHandlerSchemaUrn string

// List of Enumget-password-quality-requirements-extended-operation-handlerSchemaUrn
const (
	ENUMGETPASSWORDQUALITYREQUIREMENTSEXTENDEDOPERATIONHANDLERSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0EXTENDED_OPERATION_HANDLERGET_PASSWORD_QUALITY_REQUIREMENTS EnumgetPasswordQualityRequirementsExtendedOperationHandlerSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:extended-operation-handler:get-password-quality-requirements"
)

// All allowed values of EnumgetPasswordQualityRequirementsExtendedOperationHandlerSchemaUrn enum
var AllowedEnumgetPasswordQualityRequirementsExtendedOperationHandlerSchemaUrnEnumValues = []EnumgetPasswordQualityRequirementsExtendedOperationHandlerSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:extended-operation-handler:get-password-quality-requirements",
}

func (v *EnumgetPasswordQualityRequirementsExtendedOperationHandlerSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumgetPasswordQualityRequirementsExtendedOperationHandlerSchemaUrn(value)
	for _, existing := range AllowedEnumgetPasswordQualityRequirementsExtendedOperationHandlerSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumgetPasswordQualityRequirementsExtendedOperationHandlerSchemaUrn", value)
}

// NewEnumgetPasswordQualityRequirementsExtendedOperationHandlerSchemaUrnFromValue returns a pointer to a valid EnumgetPasswordQualityRequirementsExtendedOperationHandlerSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumgetPasswordQualityRequirementsExtendedOperationHandlerSchemaUrnFromValue(v string) (*EnumgetPasswordQualityRequirementsExtendedOperationHandlerSchemaUrn, error) {
	ev := EnumgetPasswordQualityRequirementsExtendedOperationHandlerSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumgetPasswordQualityRequirementsExtendedOperationHandlerSchemaUrn: valid values are %v", v, AllowedEnumgetPasswordQualityRequirementsExtendedOperationHandlerSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumgetPasswordQualityRequirementsExtendedOperationHandlerSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumgetPasswordQualityRequirementsExtendedOperationHandlerSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumget-password-quality-requirements-extended-operation-handlerSchemaUrn value
func (v EnumgetPasswordQualityRequirementsExtendedOperationHandlerSchemaUrn) Ptr() *EnumgetPasswordQualityRequirementsExtendedOperationHandlerSchemaUrn {
	return &v
}

type NullableEnumgetPasswordQualityRequirementsExtendedOperationHandlerSchemaUrn struct {
	value *EnumgetPasswordQualityRequirementsExtendedOperationHandlerSchemaUrn
	isSet bool
}

func (v NullableEnumgetPasswordQualityRequirementsExtendedOperationHandlerSchemaUrn) Get() *EnumgetPasswordQualityRequirementsExtendedOperationHandlerSchemaUrn {
	return v.value
}

func (v *NullableEnumgetPasswordQualityRequirementsExtendedOperationHandlerSchemaUrn) Set(val *EnumgetPasswordQualityRequirementsExtendedOperationHandlerSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumgetPasswordQualityRequirementsExtendedOperationHandlerSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumgetPasswordQualityRequirementsExtendedOperationHandlerSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumgetPasswordQualityRequirementsExtendedOperationHandlerSchemaUrn(val *EnumgetPasswordQualityRequirementsExtendedOperationHandlerSchemaUrn) *NullableEnumgetPasswordQualityRequirementsExtendedOperationHandlerSchemaUrn {
	return &NullableEnumgetPasswordQualityRequirementsExtendedOperationHandlerSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumgetPasswordQualityRequirementsExtendedOperationHandlerSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumgetPasswordQualityRequirementsExtendedOperationHandlerSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

