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

// EnummultiPartEmailAccountStatusNotificationHandlerSchemaUrn the model 'EnummultiPartEmailAccountStatusNotificationHandlerSchemaUrn'
type EnummultiPartEmailAccountStatusNotificationHandlerSchemaUrn string

// List of Enummulti-part-email-account-status-notification-handlerSchemaUrn
const (
	ENUMMULTIPARTEMAILACCOUNTSTATUSNOTIFICATIONHANDLERSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0ACCOUNT_STATUS_NOTIFICATION_HANDLERMULTI_PART_EMAIL EnummultiPartEmailAccountStatusNotificationHandlerSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:account-status-notification-handler:multi-part-email"
)

// All allowed values of EnummultiPartEmailAccountStatusNotificationHandlerSchemaUrn enum
var AllowedEnummultiPartEmailAccountStatusNotificationHandlerSchemaUrnEnumValues = []EnummultiPartEmailAccountStatusNotificationHandlerSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:account-status-notification-handler:multi-part-email",
}

func (v *EnummultiPartEmailAccountStatusNotificationHandlerSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnummultiPartEmailAccountStatusNotificationHandlerSchemaUrn(value)
	for _, existing := range AllowedEnummultiPartEmailAccountStatusNotificationHandlerSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnummultiPartEmailAccountStatusNotificationHandlerSchemaUrn", value)
}

// NewEnummultiPartEmailAccountStatusNotificationHandlerSchemaUrnFromValue returns a pointer to a valid EnummultiPartEmailAccountStatusNotificationHandlerSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnummultiPartEmailAccountStatusNotificationHandlerSchemaUrnFromValue(v string) (*EnummultiPartEmailAccountStatusNotificationHandlerSchemaUrn, error) {
	ev := EnummultiPartEmailAccountStatusNotificationHandlerSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnummultiPartEmailAccountStatusNotificationHandlerSchemaUrn: valid values are %v", v, AllowedEnummultiPartEmailAccountStatusNotificationHandlerSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnummultiPartEmailAccountStatusNotificationHandlerSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnummultiPartEmailAccountStatusNotificationHandlerSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enummulti-part-email-account-status-notification-handlerSchemaUrn value
func (v EnummultiPartEmailAccountStatusNotificationHandlerSchemaUrn) Ptr() *EnummultiPartEmailAccountStatusNotificationHandlerSchemaUrn {
	return &v
}

type NullableEnummultiPartEmailAccountStatusNotificationHandlerSchemaUrn struct {
	value *EnummultiPartEmailAccountStatusNotificationHandlerSchemaUrn
	isSet bool
}

func (v NullableEnummultiPartEmailAccountStatusNotificationHandlerSchemaUrn) Get() *EnummultiPartEmailAccountStatusNotificationHandlerSchemaUrn {
	return v.value
}

func (v *NullableEnummultiPartEmailAccountStatusNotificationHandlerSchemaUrn) Set(val *EnummultiPartEmailAccountStatusNotificationHandlerSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnummultiPartEmailAccountStatusNotificationHandlerSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnummultiPartEmailAccountStatusNotificationHandlerSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnummultiPartEmailAccountStatusNotificationHandlerSchemaUrn(val *EnummultiPartEmailAccountStatusNotificationHandlerSchemaUrn) *NullableEnummultiPartEmailAccountStatusNotificationHandlerSchemaUrn {
	return &NullableEnummultiPartEmailAccountStatusNotificationHandlerSchemaUrn{value: val, isSet: true}
}

func (v NullableEnummultiPartEmailAccountStatusNotificationHandlerSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnummultiPartEmailAccountStatusNotificationHandlerSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}