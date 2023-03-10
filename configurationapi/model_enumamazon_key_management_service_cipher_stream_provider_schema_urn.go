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

// EnumamazonKeyManagementServiceCipherStreamProviderSchemaUrn the model 'EnumamazonKeyManagementServiceCipherStreamProviderSchemaUrn'
type EnumamazonKeyManagementServiceCipherStreamProviderSchemaUrn string

// List of Enumamazon-key-management-service-cipher-stream-providerSchemaUrn
const (
	ENUMAMAZONKEYMANAGEMENTSERVICECIPHERSTREAMPROVIDERSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0CIPHER_STREAM_PROVIDERAMAZON_KEY_MANAGEMENT_SERVICE EnumamazonKeyManagementServiceCipherStreamProviderSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:cipher-stream-provider:amazon-key-management-service"
)

// All allowed values of EnumamazonKeyManagementServiceCipherStreamProviderSchemaUrn enum
var AllowedEnumamazonKeyManagementServiceCipherStreamProviderSchemaUrnEnumValues = []EnumamazonKeyManagementServiceCipherStreamProviderSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:cipher-stream-provider:amazon-key-management-service",
}

func (v *EnumamazonKeyManagementServiceCipherStreamProviderSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumamazonKeyManagementServiceCipherStreamProviderSchemaUrn(value)
	for _, existing := range AllowedEnumamazonKeyManagementServiceCipherStreamProviderSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumamazonKeyManagementServiceCipherStreamProviderSchemaUrn", value)
}

// NewEnumamazonKeyManagementServiceCipherStreamProviderSchemaUrnFromValue returns a pointer to a valid EnumamazonKeyManagementServiceCipherStreamProviderSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumamazonKeyManagementServiceCipherStreamProviderSchemaUrnFromValue(v string) (*EnumamazonKeyManagementServiceCipherStreamProviderSchemaUrn, error) {
	ev := EnumamazonKeyManagementServiceCipherStreamProviderSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumamazonKeyManagementServiceCipherStreamProviderSchemaUrn: valid values are %v", v, AllowedEnumamazonKeyManagementServiceCipherStreamProviderSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumamazonKeyManagementServiceCipherStreamProviderSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumamazonKeyManagementServiceCipherStreamProviderSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumamazon-key-management-service-cipher-stream-providerSchemaUrn value
func (v EnumamazonKeyManagementServiceCipherStreamProviderSchemaUrn) Ptr() *EnumamazonKeyManagementServiceCipherStreamProviderSchemaUrn {
	return &v
}

type NullableEnumamazonKeyManagementServiceCipherStreamProviderSchemaUrn struct {
	value *EnumamazonKeyManagementServiceCipherStreamProviderSchemaUrn
	isSet bool
}

func (v NullableEnumamazonKeyManagementServiceCipherStreamProviderSchemaUrn) Get() *EnumamazonKeyManagementServiceCipherStreamProviderSchemaUrn {
	return v.value
}

func (v *NullableEnumamazonKeyManagementServiceCipherStreamProviderSchemaUrn) Set(val *EnumamazonKeyManagementServiceCipherStreamProviderSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumamazonKeyManagementServiceCipherStreamProviderSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumamazonKeyManagementServiceCipherStreamProviderSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumamazonKeyManagementServiceCipherStreamProviderSchemaUrn(val *EnumamazonKeyManagementServiceCipherStreamProviderSchemaUrn) *NullableEnumamazonKeyManagementServiceCipherStreamProviderSchemaUrn {
	return &NullableEnumamazonKeyManagementServiceCipherStreamProviderSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumamazonKeyManagementServiceCipherStreamProviderSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumamazonKeyManagementServiceCipherStreamProviderSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
