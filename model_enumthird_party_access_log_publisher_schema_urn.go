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

// EnumthirdPartyAccessLogPublisherSchemaUrn the model 'EnumthirdPartyAccessLogPublisherSchemaUrn'
type EnumthirdPartyAccessLogPublisherSchemaUrn string

// List of Enumthird-party-access-log-publisherSchemaUrn
const (
	ENUMTHIRDPARTYACCESSLOGPUBLISHERSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0LOG_PUBLISHERTHIRD_PARTY_ACCESS EnumthirdPartyAccessLogPublisherSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:log-publisher:third-party-access"
)

// All allowed values of EnumthirdPartyAccessLogPublisherSchemaUrn enum
var AllowedEnumthirdPartyAccessLogPublisherSchemaUrnEnumValues = []EnumthirdPartyAccessLogPublisherSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:log-publisher:third-party-access",
}

func (v *EnumthirdPartyAccessLogPublisherSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumthirdPartyAccessLogPublisherSchemaUrn(value)
	for _, existing := range AllowedEnumthirdPartyAccessLogPublisherSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumthirdPartyAccessLogPublisherSchemaUrn", value)
}

// NewEnumthirdPartyAccessLogPublisherSchemaUrnFromValue returns a pointer to a valid EnumthirdPartyAccessLogPublisherSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumthirdPartyAccessLogPublisherSchemaUrnFromValue(v string) (*EnumthirdPartyAccessLogPublisherSchemaUrn, error) {
	ev := EnumthirdPartyAccessLogPublisherSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumthirdPartyAccessLogPublisherSchemaUrn: valid values are %v", v, AllowedEnumthirdPartyAccessLogPublisherSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumthirdPartyAccessLogPublisherSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumthirdPartyAccessLogPublisherSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumthird-party-access-log-publisherSchemaUrn value
func (v EnumthirdPartyAccessLogPublisherSchemaUrn) Ptr() *EnumthirdPartyAccessLogPublisherSchemaUrn {
	return &v
}

type NullableEnumthirdPartyAccessLogPublisherSchemaUrn struct {
	value *EnumthirdPartyAccessLogPublisherSchemaUrn
	isSet bool
}

func (v NullableEnumthirdPartyAccessLogPublisherSchemaUrn) Get() *EnumthirdPartyAccessLogPublisherSchemaUrn {
	return v.value
}

func (v *NullableEnumthirdPartyAccessLogPublisherSchemaUrn) Set(val *EnumthirdPartyAccessLogPublisherSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumthirdPartyAccessLogPublisherSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumthirdPartyAccessLogPublisherSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumthirdPartyAccessLogPublisherSchemaUrn(val *EnumthirdPartyAccessLogPublisherSchemaUrn) *NullableEnumthirdPartyAccessLogPublisherSchemaUrn {
	return &NullableEnumthirdPartyAccessLogPublisherSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumthirdPartyAccessLogPublisherSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumthirdPartyAccessLogPublisherSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

