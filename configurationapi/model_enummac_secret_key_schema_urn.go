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

// EnummacSecretKeySchemaUrn the model 'EnummacSecretKeySchemaUrn'
type EnummacSecretKeySchemaUrn string

// List of Enummac-secret-keySchemaUrn
const (
	ENUMMACSECRETKEYSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0MAC_SECRET_KEY EnummacSecretKeySchemaUrn = "urn:pingidentity:schemas:configuration:2.0:mac-secret-key"
)

// All allowed values of EnummacSecretKeySchemaUrn enum
var AllowedEnummacSecretKeySchemaUrnEnumValues = []EnummacSecretKeySchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:mac-secret-key",
}

func (v *EnummacSecretKeySchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnummacSecretKeySchemaUrn(value)
	for _, existing := range AllowedEnummacSecretKeySchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnummacSecretKeySchemaUrn", value)
}

// NewEnummacSecretKeySchemaUrnFromValue returns a pointer to a valid EnummacSecretKeySchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnummacSecretKeySchemaUrnFromValue(v string) (*EnummacSecretKeySchemaUrn, error) {
	ev := EnummacSecretKeySchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnummacSecretKeySchemaUrn: valid values are %v", v, AllowedEnummacSecretKeySchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnummacSecretKeySchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnummacSecretKeySchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enummac-secret-keySchemaUrn value
func (v EnummacSecretKeySchemaUrn) Ptr() *EnummacSecretKeySchemaUrn {
	return &v
}

type NullableEnummacSecretKeySchemaUrn struct {
	value *EnummacSecretKeySchemaUrn
	isSet bool
}

func (v NullableEnummacSecretKeySchemaUrn) Get() *EnummacSecretKeySchemaUrn {
	return v.value
}

func (v *NullableEnummacSecretKeySchemaUrn) Set(val *EnummacSecretKeySchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnummacSecretKeySchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnummacSecretKeySchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnummacSecretKeySchemaUrn(val *EnummacSecretKeySchemaUrn) *NullableEnummacSecretKeySchemaUrn {
	return &NullableEnummacSecretKeySchemaUrn{value: val, isSet: true}
}

func (v NullableEnummacSecretKeySchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnummacSecretKeySchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}