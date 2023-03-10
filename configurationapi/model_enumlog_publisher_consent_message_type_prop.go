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

// EnumlogPublisherConsentMessageTypeProp Specifies the consent message types that can be logged.
type EnumlogPublisherConsentMessageTypeProp string

// List of Enumlog-publisher-consentMessageTypeProp
const (
	ENUMLOGPUBLISHERCONSENTMESSAGETYPEPROP_CONSENT_CREATED        EnumlogPublisherConsentMessageTypeProp = "consent-created"
	ENUMLOGPUBLISHERCONSENTMESSAGETYPEPROP_CONSENT_UPDATED        EnumlogPublisherConsentMessageTypeProp = "consent-updated"
	ENUMLOGPUBLISHERCONSENTMESSAGETYPEPROP_CONSENT_DELETED        EnumlogPublisherConsentMessageTypeProp = "consent-deleted"
	ENUMLOGPUBLISHERCONSENTMESSAGETYPEPROP_CONSENT_RETRIEVED      EnumlogPublisherConsentMessageTypeProp = "consent-retrieved"
	ENUMLOGPUBLISHERCONSENTMESSAGETYPEPROP_CONSENT_SEARCH         EnumlogPublisherConsentMessageTypeProp = "consent-search"
	ENUMLOGPUBLISHERCONSENTMESSAGETYPEPROP_DEFINITION_CREATED     EnumlogPublisherConsentMessageTypeProp = "definition-created"
	ENUMLOGPUBLISHERCONSENTMESSAGETYPEPROP_DEFINITION_UPDATED     EnumlogPublisherConsentMessageTypeProp = "definition-updated"
	ENUMLOGPUBLISHERCONSENTMESSAGETYPEPROP_DEFINITION_DELETED     EnumlogPublisherConsentMessageTypeProp = "definition-deleted"
	ENUMLOGPUBLISHERCONSENTMESSAGETYPEPROP_DEFINITION_RETRIEVED   EnumlogPublisherConsentMessageTypeProp = "definition-retrieved"
	ENUMLOGPUBLISHERCONSENTMESSAGETYPEPROP_DEFINITION_SEARCH      EnumlogPublisherConsentMessageTypeProp = "definition-search"
	ENUMLOGPUBLISHERCONSENTMESSAGETYPEPROP_LOCALIZATION_CREATED   EnumlogPublisherConsentMessageTypeProp = "localization-created"
	ENUMLOGPUBLISHERCONSENTMESSAGETYPEPROP_LOCALIZATION_UPDATED   EnumlogPublisherConsentMessageTypeProp = "localization-updated"
	ENUMLOGPUBLISHERCONSENTMESSAGETYPEPROP_LOCALIZATION_DELETED   EnumlogPublisherConsentMessageTypeProp = "localization-deleted"
	ENUMLOGPUBLISHERCONSENTMESSAGETYPEPROP_LOCALIZATION_RETRIEVED EnumlogPublisherConsentMessageTypeProp = "localization-retrieved"
	ENUMLOGPUBLISHERCONSENTMESSAGETYPEPROP_LOCALIZATION_SEARCH    EnumlogPublisherConsentMessageTypeProp = "localization-search"
	ENUMLOGPUBLISHERCONSENTMESSAGETYPEPROP_ERROR                  EnumlogPublisherConsentMessageTypeProp = "error"
	ENUMLOGPUBLISHERCONSENTMESSAGETYPEPROP_AUDIT                  EnumlogPublisherConsentMessageTypeProp = "audit"
)

// All allowed values of EnumlogPublisherConsentMessageTypeProp enum
var AllowedEnumlogPublisherConsentMessageTypePropEnumValues = []EnumlogPublisherConsentMessageTypeProp{
	"consent-created",
	"consent-updated",
	"consent-deleted",
	"consent-retrieved",
	"consent-search",
	"definition-created",
	"definition-updated",
	"definition-deleted",
	"definition-retrieved",
	"definition-search",
	"localization-created",
	"localization-updated",
	"localization-deleted",
	"localization-retrieved",
	"localization-search",
	"error",
	"audit",
}

func (v *EnumlogPublisherConsentMessageTypeProp) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumlogPublisherConsentMessageTypeProp(value)
	for _, existing := range AllowedEnumlogPublisherConsentMessageTypePropEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumlogPublisherConsentMessageTypeProp", value)
}

// NewEnumlogPublisherConsentMessageTypePropFromValue returns a pointer to a valid EnumlogPublisherConsentMessageTypeProp
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumlogPublisherConsentMessageTypePropFromValue(v string) (*EnumlogPublisherConsentMessageTypeProp, error) {
	ev := EnumlogPublisherConsentMessageTypeProp(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumlogPublisherConsentMessageTypeProp: valid values are %v", v, AllowedEnumlogPublisherConsentMessageTypePropEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumlogPublisherConsentMessageTypeProp) IsValid() bool {
	for _, existing := range AllowedEnumlogPublisherConsentMessageTypePropEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumlog-publisher-consentMessageTypeProp value
func (v EnumlogPublisherConsentMessageTypeProp) Ptr() *EnumlogPublisherConsentMessageTypeProp {
	return &v
}

type NullableEnumlogPublisherConsentMessageTypeProp struct {
	value *EnumlogPublisherConsentMessageTypeProp
	isSet bool
}

func (v NullableEnumlogPublisherConsentMessageTypeProp) Get() *EnumlogPublisherConsentMessageTypeProp {
	return v.value
}

func (v *NullableEnumlogPublisherConsentMessageTypeProp) Set(val *EnumlogPublisherConsentMessageTypeProp) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumlogPublisherConsentMessageTypeProp) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumlogPublisherConsentMessageTypeProp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumlogPublisherConsentMessageTypeProp(val *EnumlogPublisherConsentMessageTypeProp) *NullableEnumlogPublisherConsentMessageTypeProp {
	return &NullableEnumlogPublisherConsentMessageTypeProp{value: val, isSet: true}
}

func (v NullableEnumlogPublisherConsentMessageTypeProp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumlogPublisherConsentMessageTypeProp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
