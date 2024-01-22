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

// EnumlogPublisherDebugMessageTypeProp Specifies the debug message types which can be logged. Note that enabling these may result in sensitive information being logged.
type EnumlogPublisherDebugMessageTypeProp string

// List of Enumlog-publisher-debugMessageTypeProp
const (
	ENUMLOGPUBLISHERDEBUGMESSAGETYPEPROP_HTTP_FULL_REQUEST_AND_RESPONSE              EnumlogPublisherDebugMessageTypeProp = "http-full-request-and-response"
	ENUMLOGPUBLISHERDEBUGMESSAGETYPEPROP_LDAP_EXTERNAL_SERVER_REQUEST                EnumlogPublisherDebugMessageTypeProp = "ldap-external-server-request"
	ENUMLOGPUBLISHERDEBUGMESSAGETYPEPROP_SERVER_SDK_EXTENSION                        EnumlogPublisherDebugMessageTypeProp = "server-sdk-extension"
	ENUMLOGPUBLISHERDEBUGMESSAGETYPEPROP_STORE_ADAPTER_MAPPING                       EnumlogPublisherDebugMessageTypeProp = "store-adapter-mapping"
	ENUMLOGPUBLISHERDEBUGMESSAGETYPEPROP_STORE_ADAPTER_PROCESSING                    EnumlogPublisherDebugMessageTypeProp = "store-adapter-processing"
	ENUMLOGPUBLISHERDEBUGMESSAGETYPEPROP_ACCESS_TOKEN_VALIDATOR_REQUEST_AND_RESPONSE EnumlogPublisherDebugMessageTypeProp = "access-token-validator-request-and-response"
	ENUMLOGPUBLISHERDEBUGMESSAGETYPEPROP_ACCESS_TOKEN_VALIDATOR_PROCESSING           EnumlogPublisherDebugMessageTypeProp = "access-token-validator-processing"
	ENUMLOGPUBLISHERDEBUGMESSAGETYPEPROP_ID_TOKEN_VALIDATOR_REQUEST_AND_RESPONSE     EnumlogPublisherDebugMessageTypeProp = "id-token-validator-request-and-response"
	ENUMLOGPUBLISHERDEBUGMESSAGETYPEPROP_ID_TOKEN_VALIDATOR_PROCESSING               EnumlogPublisherDebugMessageTypeProp = "id-token-validator-processing"
	ENUMLOGPUBLISHERDEBUGMESSAGETYPEPROP_CONSENT_SERVICE_REQUEST_AND_RESPONSE        EnumlogPublisherDebugMessageTypeProp = "consent-service-request-and-response"
	ENUMLOGPUBLISHERDEBUGMESSAGETYPEPROP_GATEWAY_REQUEST_AND_RESPONSE                EnumlogPublisherDebugMessageTypeProp = "gateway-request-and-response"
	ENUMLOGPUBLISHERDEBUGMESSAGETYPEPROP_POLICY_REQUEST_AND_RESPONSE                 EnumlogPublisherDebugMessageTypeProp = "policy-request-and-response"
	ENUMLOGPUBLISHERDEBUGMESSAGETYPEPROP_POLICY_QUERY_REQUEST_AND_RESPONSE           EnumlogPublisherDebugMessageTypeProp = "policy-query-request-and-response"
)

// All allowed values of EnumlogPublisherDebugMessageTypeProp enum
var AllowedEnumlogPublisherDebugMessageTypePropEnumValues = []EnumlogPublisherDebugMessageTypeProp{
	"http-full-request-and-response",
	"ldap-external-server-request",
	"server-sdk-extension",
	"store-adapter-mapping",
	"store-adapter-processing",
	"access-token-validator-request-and-response",
	"access-token-validator-processing",
	"id-token-validator-request-and-response",
	"id-token-validator-processing",
	"consent-service-request-and-response",
	"gateway-request-and-response",
	"policy-request-and-response",
	"policy-query-request-and-response",
}

func (v *EnumlogPublisherDebugMessageTypeProp) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumlogPublisherDebugMessageTypeProp(value)
	for _, existing := range AllowedEnumlogPublisherDebugMessageTypePropEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumlogPublisherDebugMessageTypeProp", value)
}

// NewEnumlogPublisherDebugMessageTypePropFromValue returns a pointer to a valid EnumlogPublisherDebugMessageTypeProp
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumlogPublisherDebugMessageTypePropFromValue(v string) (*EnumlogPublisherDebugMessageTypeProp, error) {
	ev := EnumlogPublisherDebugMessageTypeProp(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumlogPublisherDebugMessageTypeProp: valid values are %v", v, AllowedEnumlogPublisherDebugMessageTypePropEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumlogPublisherDebugMessageTypeProp) IsValid() bool {
	for _, existing := range AllowedEnumlogPublisherDebugMessageTypePropEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumlog-publisher-debugMessageTypeProp value
func (v EnumlogPublisherDebugMessageTypeProp) Ptr() *EnumlogPublisherDebugMessageTypeProp {
	return &v
}

type NullableEnumlogPublisherDebugMessageTypeProp struct {
	value *EnumlogPublisherDebugMessageTypeProp
	isSet bool
}

func (v NullableEnumlogPublisherDebugMessageTypeProp) Get() *EnumlogPublisherDebugMessageTypeProp {
	return v.value
}

func (v *NullableEnumlogPublisherDebugMessageTypeProp) Set(val *EnumlogPublisherDebugMessageTypeProp) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumlogPublisherDebugMessageTypeProp) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumlogPublisherDebugMessageTypeProp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumlogPublisherDebugMessageTypeProp(val *EnumlogPublisherDebugMessageTypeProp) *NullableEnumlogPublisherDebugMessageTypeProp {
	return &NullableEnumlogPublisherDebugMessageTypeProp{value: val, isSet: true}
}

func (v NullableEnumlogPublisherDebugMessageTypeProp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumlogPublisherDebugMessageTypeProp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
