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

// EnumunboundidMsChapV2SaslMechanismHandlerSchemaUrn the model 'EnumunboundidMsChapV2SaslMechanismHandlerSchemaUrn'
type EnumunboundidMsChapV2SaslMechanismHandlerSchemaUrn string

// List of Enumunboundid-ms-chap-v2-sasl-mechanism-handlerSchemaUrn
const (
	ENUMUNBOUNDIDMSCHAPV2SASLMECHANISMHANDLERSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0SASL_MECHANISM_HANDLERUNBOUNDID_MS_CHAP_V2 EnumunboundidMsChapV2SaslMechanismHandlerSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:sasl-mechanism-handler:unboundid-ms-chap-v2"
)

// All allowed values of EnumunboundidMsChapV2SaslMechanismHandlerSchemaUrn enum
var AllowedEnumunboundidMsChapV2SaslMechanismHandlerSchemaUrnEnumValues = []EnumunboundidMsChapV2SaslMechanismHandlerSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:sasl-mechanism-handler:unboundid-ms-chap-v2",
}

func (v *EnumunboundidMsChapV2SaslMechanismHandlerSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumunboundidMsChapV2SaslMechanismHandlerSchemaUrn(value)
	for _, existing := range AllowedEnumunboundidMsChapV2SaslMechanismHandlerSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumunboundidMsChapV2SaslMechanismHandlerSchemaUrn", value)
}

// NewEnumunboundidMsChapV2SaslMechanismHandlerSchemaUrnFromValue returns a pointer to a valid EnumunboundidMsChapV2SaslMechanismHandlerSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumunboundidMsChapV2SaslMechanismHandlerSchemaUrnFromValue(v string) (*EnumunboundidMsChapV2SaslMechanismHandlerSchemaUrn, error) {
	ev := EnumunboundidMsChapV2SaslMechanismHandlerSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumunboundidMsChapV2SaslMechanismHandlerSchemaUrn: valid values are %v", v, AllowedEnumunboundidMsChapV2SaslMechanismHandlerSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumunboundidMsChapV2SaslMechanismHandlerSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumunboundidMsChapV2SaslMechanismHandlerSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumunboundid-ms-chap-v2-sasl-mechanism-handlerSchemaUrn value
func (v EnumunboundidMsChapV2SaslMechanismHandlerSchemaUrn) Ptr() *EnumunboundidMsChapV2SaslMechanismHandlerSchemaUrn {
	return &v
}

type NullableEnumunboundidMsChapV2SaslMechanismHandlerSchemaUrn struct {
	value *EnumunboundidMsChapV2SaslMechanismHandlerSchemaUrn
	isSet bool
}

func (v NullableEnumunboundidMsChapV2SaslMechanismHandlerSchemaUrn) Get() *EnumunboundidMsChapV2SaslMechanismHandlerSchemaUrn {
	return v.value
}

func (v *NullableEnumunboundidMsChapV2SaslMechanismHandlerSchemaUrn) Set(val *EnumunboundidMsChapV2SaslMechanismHandlerSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumunboundidMsChapV2SaslMechanismHandlerSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumunboundidMsChapV2SaslMechanismHandlerSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumunboundidMsChapV2SaslMechanismHandlerSchemaUrn(val *EnumunboundidMsChapV2SaslMechanismHandlerSchemaUrn) *NullableEnumunboundidMsChapV2SaslMechanismHandlerSchemaUrn {
	return &NullableEnumunboundidMsChapV2SaslMechanismHandlerSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumunboundidMsChapV2SaslMechanismHandlerSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumunboundidMsChapV2SaslMechanismHandlerSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}