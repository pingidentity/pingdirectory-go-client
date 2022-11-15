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

// EnumsaslMechanismHandlerCertificateValidationPolicyProp Indicates whether to attempt to validate the peer certificate against a certificate held in the user's entry.
type EnumsaslMechanismHandlerCertificateValidationPolicyProp string

// List of Enumsasl-mechanism-handler-certificateValidationPolicyProp
const (
	ENUMSASLMECHANISMHANDLERCERTIFICATEVALIDATIONPOLICYPROP_ALWAYS EnumsaslMechanismHandlerCertificateValidationPolicyProp = "always"
	ENUMSASLMECHANISMHANDLERCERTIFICATEVALIDATIONPOLICYPROP_IFPRESENT EnumsaslMechanismHandlerCertificateValidationPolicyProp = "ifpresent"
	ENUMSASLMECHANISMHANDLERCERTIFICATEVALIDATIONPOLICYPROP_NEVER EnumsaslMechanismHandlerCertificateValidationPolicyProp = "never"
)

// All allowed values of EnumsaslMechanismHandlerCertificateValidationPolicyProp enum
var AllowedEnumsaslMechanismHandlerCertificateValidationPolicyPropEnumValues = []EnumsaslMechanismHandlerCertificateValidationPolicyProp{
	"always",
	"ifpresent",
	"never",
}

func (v *EnumsaslMechanismHandlerCertificateValidationPolicyProp) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumsaslMechanismHandlerCertificateValidationPolicyProp(value)
	for _, existing := range AllowedEnumsaslMechanismHandlerCertificateValidationPolicyPropEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumsaslMechanismHandlerCertificateValidationPolicyProp", value)
}

// NewEnumsaslMechanismHandlerCertificateValidationPolicyPropFromValue returns a pointer to a valid EnumsaslMechanismHandlerCertificateValidationPolicyProp
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumsaslMechanismHandlerCertificateValidationPolicyPropFromValue(v string) (*EnumsaslMechanismHandlerCertificateValidationPolicyProp, error) {
	ev := EnumsaslMechanismHandlerCertificateValidationPolicyProp(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumsaslMechanismHandlerCertificateValidationPolicyProp: valid values are %v", v, AllowedEnumsaslMechanismHandlerCertificateValidationPolicyPropEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumsaslMechanismHandlerCertificateValidationPolicyProp) IsValid() bool {
	for _, existing := range AllowedEnumsaslMechanismHandlerCertificateValidationPolicyPropEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumsasl-mechanism-handler-certificateValidationPolicyProp value
func (v EnumsaslMechanismHandlerCertificateValidationPolicyProp) Ptr() *EnumsaslMechanismHandlerCertificateValidationPolicyProp {
	return &v
}

type NullableEnumsaslMechanismHandlerCertificateValidationPolicyProp struct {
	value *EnumsaslMechanismHandlerCertificateValidationPolicyProp
	isSet bool
}

func (v NullableEnumsaslMechanismHandlerCertificateValidationPolicyProp) Get() *EnumsaslMechanismHandlerCertificateValidationPolicyProp {
	return v.value
}

func (v *NullableEnumsaslMechanismHandlerCertificateValidationPolicyProp) Set(val *EnumsaslMechanismHandlerCertificateValidationPolicyProp) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumsaslMechanismHandlerCertificateValidationPolicyProp) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumsaslMechanismHandlerCertificateValidationPolicyProp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumsaslMechanismHandlerCertificateValidationPolicyProp(val *EnumsaslMechanismHandlerCertificateValidationPolicyProp) *NullableEnumsaslMechanismHandlerCertificateValidationPolicyProp {
	return &NullableEnumsaslMechanismHandlerCertificateValidationPolicyProp{value: val, isSet: true}
}

func (v NullableEnumsaslMechanismHandlerCertificateValidationPolicyProp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumsaslMechanismHandlerCertificateValidationPolicyProp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

