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

// EnumoauthBearerSaslMechanismHandlerSchemaUrn the model 'EnumoauthBearerSaslMechanismHandlerSchemaUrn'
type EnumoauthBearerSaslMechanismHandlerSchemaUrn string

// List of Enumoauth-bearer-sasl-mechanism-handlerSchemaUrn
const (
	ENUMOAUTHBEARERSASLMECHANISMHANDLERSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0SASL_MECHANISM_HANDLEROAUTH_BEARER EnumoauthBearerSaslMechanismHandlerSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:sasl-mechanism-handler:oauth-bearer"
)

// All allowed values of EnumoauthBearerSaslMechanismHandlerSchemaUrn enum
var AllowedEnumoauthBearerSaslMechanismHandlerSchemaUrnEnumValues = []EnumoauthBearerSaslMechanismHandlerSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:sasl-mechanism-handler:oauth-bearer",
}

func (v *EnumoauthBearerSaslMechanismHandlerSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumoauthBearerSaslMechanismHandlerSchemaUrn(value)
	for _, existing := range AllowedEnumoauthBearerSaslMechanismHandlerSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumoauthBearerSaslMechanismHandlerSchemaUrn", value)
}

// NewEnumoauthBearerSaslMechanismHandlerSchemaUrnFromValue returns a pointer to a valid EnumoauthBearerSaslMechanismHandlerSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumoauthBearerSaslMechanismHandlerSchemaUrnFromValue(v string) (*EnumoauthBearerSaslMechanismHandlerSchemaUrn, error) {
	ev := EnumoauthBearerSaslMechanismHandlerSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumoauthBearerSaslMechanismHandlerSchemaUrn: valid values are %v", v, AllowedEnumoauthBearerSaslMechanismHandlerSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumoauthBearerSaslMechanismHandlerSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumoauthBearerSaslMechanismHandlerSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumoauth-bearer-sasl-mechanism-handlerSchemaUrn value
func (v EnumoauthBearerSaslMechanismHandlerSchemaUrn) Ptr() *EnumoauthBearerSaslMechanismHandlerSchemaUrn {
	return &v
}

type NullableEnumoauthBearerSaslMechanismHandlerSchemaUrn struct {
	value *EnumoauthBearerSaslMechanismHandlerSchemaUrn
	isSet bool
}

func (v NullableEnumoauthBearerSaslMechanismHandlerSchemaUrn) Get() *EnumoauthBearerSaslMechanismHandlerSchemaUrn {
	return v.value
}

func (v *NullableEnumoauthBearerSaslMechanismHandlerSchemaUrn) Set(val *EnumoauthBearerSaslMechanismHandlerSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumoauthBearerSaslMechanismHandlerSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumoauthBearerSaslMechanismHandlerSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumoauthBearerSaslMechanismHandlerSchemaUrn(val *EnumoauthBearerSaslMechanismHandlerSchemaUrn) *NullableEnumoauthBearerSaslMechanismHandlerSchemaUrn {
	return &NullableEnumoauthBearerSaslMechanismHandlerSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumoauthBearerSaslMechanismHandlerSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumoauthBearerSaslMechanismHandlerSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

