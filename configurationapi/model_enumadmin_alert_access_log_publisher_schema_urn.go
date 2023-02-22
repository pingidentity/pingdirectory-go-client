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

// EnumadminAlertAccessLogPublisherSchemaUrn the model 'EnumadminAlertAccessLogPublisherSchemaUrn'
type EnumadminAlertAccessLogPublisherSchemaUrn string

// List of Enumadmin-alert-access-log-publisherSchemaUrn
const (
	ENUMADMINALERTACCESSLOGPUBLISHERSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0LOG_PUBLISHERADMIN_ALERT_ACCESS EnumadminAlertAccessLogPublisherSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:log-publisher:admin-alert-access"
)

// All allowed values of EnumadminAlertAccessLogPublisherSchemaUrn enum
var AllowedEnumadminAlertAccessLogPublisherSchemaUrnEnumValues = []EnumadminAlertAccessLogPublisherSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:log-publisher:admin-alert-access",
}

func (v *EnumadminAlertAccessLogPublisherSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumadminAlertAccessLogPublisherSchemaUrn(value)
	for _, existing := range AllowedEnumadminAlertAccessLogPublisherSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumadminAlertAccessLogPublisherSchemaUrn", value)
}

// NewEnumadminAlertAccessLogPublisherSchemaUrnFromValue returns a pointer to a valid EnumadminAlertAccessLogPublisherSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumadminAlertAccessLogPublisherSchemaUrnFromValue(v string) (*EnumadminAlertAccessLogPublisherSchemaUrn, error) {
	ev := EnumadminAlertAccessLogPublisherSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumadminAlertAccessLogPublisherSchemaUrn: valid values are %v", v, AllowedEnumadminAlertAccessLogPublisherSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumadminAlertAccessLogPublisherSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumadminAlertAccessLogPublisherSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumadmin-alert-access-log-publisherSchemaUrn value
func (v EnumadminAlertAccessLogPublisherSchemaUrn) Ptr() *EnumadminAlertAccessLogPublisherSchemaUrn {
	return &v
}

type NullableEnumadminAlertAccessLogPublisherSchemaUrn struct {
	value *EnumadminAlertAccessLogPublisherSchemaUrn
	isSet bool
}

func (v NullableEnumadminAlertAccessLogPublisherSchemaUrn) Get() *EnumadminAlertAccessLogPublisherSchemaUrn {
	return v.value
}

func (v *NullableEnumadminAlertAccessLogPublisherSchemaUrn) Set(val *EnumadminAlertAccessLogPublisherSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumadminAlertAccessLogPublisherSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumadminAlertAccessLogPublisherSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumadminAlertAccessLogPublisherSchemaUrn(val *EnumadminAlertAccessLogPublisherSchemaUrn) *NullableEnumadminAlertAccessLogPublisherSchemaUrn {
	return &NullableEnumadminAlertAccessLogPublisherSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumadminAlertAccessLogPublisherSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumadminAlertAccessLogPublisherSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}