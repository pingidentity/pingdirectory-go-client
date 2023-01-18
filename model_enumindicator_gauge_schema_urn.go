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

// EnumindicatorGaugeSchemaUrn the model 'EnumindicatorGaugeSchemaUrn'
type EnumindicatorGaugeSchemaUrn string

// List of Enumindicator-gaugeSchemaUrn
const (
	ENUMINDICATORGAUGESCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0GAUGEINDICATOR EnumindicatorGaugeSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:gauge:indicator"
)

// All allowed values of EnumindicatorGaugeSchemaUrn enum
var AllowedEnumindicatorGaugeSchemaUrnEnumValues = []EnumindicatorGaugeSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:gauge:indicator",
}

func (v *EnumindicatorGaugeSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumindicatorGaugeSchemaUrn(value)
	for _, existing := range AllowedEnumindicatorGaugeSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumindicatorGaugeSchemaUrn", value)
}

// NewEnumindicatorGaugeSchemaUrnFromValue returns a pointer to a valid EnumindicatorGaugeSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumindicatorGaugeSchemaUrnFromValue(v string) (*EnumindicatorGaugeSchemaUrn, error) {
	ev := EnumindicatorGaugeSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumindicatorGaugeSchemaUrn: valid values are %v", v, AllowedEnumindicatorGaugeSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumindicatorGaugeSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumindicatorGaugeSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumindicator-gaugeSchemaUrn value
func (v EnumindicatorGaugeSchemaUrn) Ptr() *EnumindicatorGaugeSchemaUrn {
	return &v
}

type NullableEnumindicatorGaugeSchemaUrn struct {
	value *EnumindicatorGaugeSchemaUrn
	isSet bool
}

func (v NullableEnumindicatorGaugeSchemaUrn) Get() *EnumindicatorGaugeSchemaUrn {
	return v.value
}

func (v *NullableEnumindicatorGaugeSchemaUrn) Set(val *EnumindicatorGaugeSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumindicatorGaugeSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumindicatorGaugeSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumindicatorGaugeSchemaUrn(val *EnumindicatorGaugeSchemaUrn) *NullableEnumindicatorGaugeSchemaUrn {
	return &NullableEnumindicatorGaugeSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumindicatorGaugeSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumindicatorGaugeSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
