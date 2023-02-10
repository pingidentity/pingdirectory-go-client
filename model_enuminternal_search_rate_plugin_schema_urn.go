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

// EnuminternalSearchRatePluginSchemaUrn the model 'EnuminternalSearchRatePluginSchemaUrn'
type EnuminternalSearchRatePluginSchemaUrn string

// List of Enuminternal-search-rate-pluginSchemaUrn
const (
	ENUMINTERNALSEARCHRATEPLUGINSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0PLUGININTERNAL_SEARCH_RATE EnuminternalSearchRatePluginSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:plugin:internal-search-rate"
)

// All allowed values of EnuminternalSearchRatePluginSchemaUrn enum
var AllowedEnuminternalSearchRatePluginSchemaUrnEnumValues = []EnuminternalSearchRatePluginSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:plugin:internal-search-rate",
}

func (v *EnuminternalSearchRatePluginSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnuminternalSearchRatePluginSchemaUrn(value)
	for _, existing := range AllowedEnuminternalSearchRatePluginSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnuminternalSearchRatePluginSchemaUrn", value)
}

// NewEnuminternalSearchRatePluginSchemaUrnFromValue returns a pointer to a valid EnuminternalSearchRatePluginSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnuminternalSearchRatePluginSchemaUrnFromValue(v string) (*EnuminternalSearchRatePluginSchemaUrn, error) {
	ev := EnuminternalSearchRatePluginSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnuminternalSearchRatePluginSchemaUrn: valid values are %v", v, AllowedEnuminternalSearchRatePluginSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnuminternalSearchRatePluginSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnuminternalSearchRatePluginSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enuminternal-search-rate-pluginSchemaUrn value
func (v EnuminternalSearchRatePluginSchemaUrn) Ptr() *EnuminternalSearchRatePluginSchemaUrn {
	return &v
}

type NullableEnuminternalSearchRatePluginSchemaUrn struct {
	value *EnuminternalSearchRatePluginSchemaUrn
	isSet bool
}

func (v NullableEnuminternalSearchRatePluginSchemaUrn) Get() *EnuminternalSearchRatePluginSchemaUrn {
	return v.value
}

func (v *NullableEnuminternalSearchRatePluginSchemaUrn) Set(val *EnuminternalSearchRatePluginSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnuminternalSearchRatePluginSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnuminternalSearchRatePluginSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnuminternalSearchRatePluginSchemaUrn(val *EnuminternalSearchRatePluginSchemaUrn) *NullableEnuminternalSearchRatePluginSchemaUrn {
	return &NullableEnuminternalSearchRatePluginSchemaUrn{value: val, isSet: true}
}

func (v NullableEnuminternalSearchRatePluginSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnuminternalSearchRatePluginSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
