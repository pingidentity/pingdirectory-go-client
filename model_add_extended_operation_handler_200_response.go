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

// AddExtendedOperationHandler200Response - struct for AddExtendedOperationHandler200Response
type AddExtendedOperationHandler200Response struct {
	CollectSupportDataExtendedOperationHandlerResponse        *CollectSupportDataExtendedOperationHandlerResponse
	DeliverOtpExtendedOperationHandlerResponse                *DeliverOtpExtendedOperationHandlerResponse
	DeliverPasswordResetTokenExtendedOperationHandlerResponse *DeliverPasswordResetTokenExtendedOperationHandlerResponse
	ExportReversiblePasswordsExtendedOperationHandlerResponse *ExportReversiblePasswordsExtendedOperationHandlerResponse
	ReplaceCertificateExtendedOperationHandlerResponse        *ReplaceCertificateExtendedOperationHandlerResponse
	SingleUseTokensExtendedOperationHandlerResponse           *SingleUseTokensExtendedOperationHandlerResponse
	ThirdPartyExtendedOperationHandlerResponse                *ThirdPartyExtendedOperationHandlerResponse
	ValidateTotpPasswordExtendedOperationHandlerResponse      *ValidateTotpPasswordExtendedOperationHandlerResponse
}

// CollectSupportDataExtendedOperationHandlerResponseAsAddExtendedOperationHandler200Response is a convenience function that returns CollectSupportDataExtendedOperationHandlerResponse wrapped in AddExtendedOperationHandler200Response
func CollectSupportDataExtendedOperationHandlerResponseAsAddExtendedOperationHandler200Response(v *CollectSupportDataExtendedOperationHandlerResponse) AddExtendedOperationHandler200Response {
	return AddExtendedOperationHandler200Response{
		CollectSupportDataExtendedOperationHandlerResponse: v,
	}
}

// DeliverOtpExtendedOperationHandlerResponseAsAddExtendedOperationHandler200Response is a convenience function that returns DeliverOtpExtendedOperationHandlerResponse wrapped in AddExtendedOperationHandler200Response
func DeliverOtpExtendedOperationHandlerResponseAsAddExtendedOperationHandler200Response(v *DeliverOtpExtendedOperationHandlerResponse) AddExtendedOperationHandler200Response {
	return AddExtendedOperationHandler200Response{
		DeliverOtpExtendedOperationHandlerResponse: v,
	}
}

// DeliverPasswordResetTokenExtendedOperationHandlerResponseAsAddExtendedOperationHandler200Response is a convenience function that returns DeliverPasswordResetTokenExtendedOperationHandlerResponse wrapped in AddExtendedOperationHandler200Response
func DeliverPasswordResetTokenExtendedOperationHandlerResponseAsAddExtendedOperationHandler200Response(v *DeliverPasswordResetTokenExtendedOperationHandlerResponse) AddExtendedOperationHandler200Response {
	return AddExtendedOperationHandler200Response{
		DeliverPasswordResetTokenExtendedOperationHandlerResponse: v,
	}
}

// ExportReversiblePasswordsExtendedOperationHandlerResponseAsAddExtendedOperationHandler200Response is a convenience function that returns ExportReversiblePasswordsExtendedOperationHandlerResponse wrapped in AddExtendedOperationHandler200Response
func ExportReversiblePasswordsExtendedOperationHandlerResponseAsAddExtendedOperationHandler200Response(v *ExportReversiblePasswordsExtendedOperationHandlerResponse) AddExtendedOperationHandler200Response {
	return AddExtendedOperationHandler200Response{
		ExportReversiblePasswordsExtendedOperationHandlerResponse: v,
	}
}

// ReplaceCertificateExtendedOperationHandlerResponseAsAddExtendedOperationHandler200Response is a convenience function that returns ReplaceCertificateExtendedOperationHandlerResponse wrapped in AddExtendedOperationHandler200Response
func ReplaceCertificateExtendedOperationHandlerResponseAsAddExtendedOperationHandler200Response(v *ReplaceCertificateExtendedOperationHandlerResponse) AddExtendedOperationHandler200Response {
	return AddExtendedOperationHandler200Response{
		ReplaceCertificateExtendedOperationHandlerResponse: v,
	}
}

// SingleUseTokensExtendedOperationHandlerResponseAsAddExtendedOperationHandler200Response is a convenience function that returns SingleUseTokensExtendedOperationHandlerResponse wrapped in AddExtendedOperationHandler200Response
func SingleUseTokensExtendedOperationHandlerResponseAsAddExtendedOperationHandler200Response(v *SingleUseTokensExtendedOperationHandlerResponse) AddExtendedOperationHandler200Response {
	return AddExtendedOperationHandler200Response{
		SingleUseTokensExtendedOperationHandlerResponse: v,
	}
}

// ThirdPartyExtendedOperationHandlerResponseAsAddExtendedOperationHandler200Response is a convenience function that returns ThirdPartyExtendedOperationHandlerResponse wrapped in AddExtendedOperationHandler200Response
func ThirdPartyExtendedOperationHandlerResponseAsAddExtendedOperationHandler200Response(v *ThirdPartyExtendedOperationHandlerResponse) AddExtendedOperationHandler200Response {
	return AddExtendedOperationHandler200Response{
		ThirdPartyExtendedOperationHandlerResponse: v,
	}
}

// ValidateTotpPasswordExtendedOperationHandlerResponseAsAddExtendedOperationHandler200Response is a convenience function that returns ValidateTotpPasswordExtendedOperationHandlerResponse wrapped in AddExtendedOperationHandler200Response
func ValidateTotpPasswordExtendedOperationHandlerResponseAsAddExtendedOperationHandler200Response(v *ValidateTotpPasswordExtendedOperationHandlerResponse) AddExtendedOperationHandler200Response {
	return AddExtendedOperationHandler200Response{
		ValidateTotpPasswordExtendedOperationHandlerResponse: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *AddExtendedOperationHandler200Response) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into CollectSupportDataExtendedOperationHandlerResponse
	err = newStrictDecoder(data).Decode(&dst.CollectSupportDataExtendedOperationHandlerResponse)
	if err == nil {
		jsonCollectSupportDataExtendedOperationHandlerResponse, _ := json.Marshal(dst.CollectSupportDataExtendedOperationHandlerResponse)
		if string(jsonCollectSupportDataExtendedOperationHandlerResponse) == "{}" { // empty struct
			dst.CollectSupportDataExtendedOperationHandlerResponse = nil
		} else {
			match++
		}
	} else {
		dst.CollectSupportDataExtendedOperationHandlerResponse = nil
	}

	// try to unmarshal data into DeliverOtpExtendedOperationHandlerResponse
	err = newStrictDecoder(data).Decode(&dst.DeliverOtpExtendedOperationHandlerResponse)
	if err == nil {
		jsonDeliverOtpExtendedOperationHandlerResponse, _ := json.Marshal(dst.DeliverOtpExtendedOperationHandlerResponse)
		if string(jsonDeliverOtpExtendedOperationHandlerResponse) == "{}" { // empty struct
			dst.DeliverOtpExtendedOperationHandlerResponse = nil
		} else {
			match++
		}
	} else {
		dst.DeliverOtpExtendedOperationHandlerResponse = nil
	}

	// try to unmarshal data into DeliverPasswordResetTokenExtendedOperationHandlerResponse
	err = newStrictDecoder(data).Decode(&dst.DeliverPasswordResetTokenExtendedOperationHandlerResponse)
	if err == nil {
		jsonDeliverPasswordResetTokenExtendedOperationHandlerResponse, _ := json.Marshal(dst.DeliverPasswordResetTokenExtendedOperationHandlerResponse)
		if string(jsonDeliverPasswordResetTokenExtendedOperationHandlerResponse) == "{}" { // empty struct
			dst.DeliverPasswordResetTokenExtendedOperationHandlerResponse = nil
		} else {
			match++
		}
	} else {
		dst.DeliverPasswordResetTokenExtendedOperationHandlerResponse = nil
	}

	// try to unmarshal data into ExportReversiblePasswordsExtendedOperationHandlerResponse
	err = newStrictDecoder(data).Decode(&dst.ExportReversiblePasswordsExtendedOperationHandlerResponse)
	if err == nil {
		jsonExportReversiblePasswordsExtendedOperationHandlerResponse, _ := json.Marshal(dst.ExportReversiblePasswordsExtendedOperationHandlerResponse)
		if string(jsonExportReversiblePasswordsExtendedOperationHandlerResponse) == "{}" { // empty struct
			dst.ExportReversiblePasswordsExtendedOperationHandlerResponse = nil
		} else {
			match++
		}
	} else {
		dst.ExportReversiblePasswordsExtendedOperationHandlerResponse = nil
	}

	// try to unmarshal data into ReplaceCertificateExtendedOperationHandlerResponse
	err = newStrictDecoder(data).Decode(&dst.ReplaceCertificateExtendedOperationHandlerResponse)
	if err == nil {
		jsonReplaceCertificateExtendedOperationHandlerResponse, _ := json.Marshal(dst.ReplaceCertificateExtendedOperationHandlerResponse)
		if string(jsonReplaceCertificateExtendedOperationHandlerResponse) == "{}" { // empty struct
			dst.ReplaceCertificateExtendedOperationHandlerResponse = nil
		} else {
			match++
		}
	} else {
		dst.ReplaceCertificateExtendedOperationHandlerResponse = nil
	}

	// try to unmarshal data into SingleUseTokensExtendedOperationHandlerResponse
	err = newStrictDecoder(data).Decode(&dst.SingleUseTokensExtendedOperationHandlerResponse)
	if err == nil {
		jsonSingleUseTokensExtendedOperationHandlerResponse, _ := json.Marshal(dst.SingleUseTokensExtendedOperationHandlerResponse)
		if string(jsonSingleUseTokensExtendedOperationHandlerResponse) == "{}" { // empty struct
			dst.SingleUseTokensExtendedOperationHandlerResponse = nil
		} else {
			match++
		}
	} else {
		dst.SingleUseTokensExtendedOperationHandlerResponse = nil
	}

	// try to unmarshal data into ThirdPartyExtendedOperationHandlerResponse
	err = newStrictDecoder(data).Decode(&dst.ThirdPartyExtendedOperationHandlerResponse)
	if err == nil {
		jsonThirdPartyExtendedOperationHandlerResponse, _ := json.Marshal(dst.ThirdPartyExtendedOperationHandlerResponse)
		if string(jsonThirdPartyExtendedOperationHandlerResponse) == "{}" { // empty struct
			dst.ThirdPartyExtendedOperationHandlerResponse = nil
		} else {
			match++
		}
	} else {
		dst.ThirdPartyExtendedOperationHandlerResponse = nil
	}

	// try to unmarshal data into ValidateTotpPasswordExtendedOperationHandlerResponse
	err = newStrictDecoder(data).Decode(&dst.ValidateTotpPasswordExtendedOperationHandlerResponse)
	if err == nil {
		jsonValidateTotpPasswordExtendedOperationHandlerResponse, _ := json.Marshal(dst.ValidateTotpPasswordExtendedOperationHandlerResponse)
		if string(jsonValidateTotpPasswordExtendedOperationHandlerResponse) == "{}" { // empty struct
			dst.ValidateTotpPasswordExtendedOperationHandlerResponse = nil
		} else {
			match++
		}
	} else {
		dst.ValidateTotpPasswordExtendedOperationHandlerResponse = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.CollectSupportDataExtendedOperationHandlerResponse = nil
		dst.DeliverOtpExtendedOperationHandlerResponse = nil
		dst.DeliverPasswordResetTokenExtendedOperationHandlerResponse = nil
		dst.ExportReversiblePasswordsExtendedOperationHandlerResponse = nil
		dst.ReplaceCertificateExtendedOperationHandlerResponse = nil
		dst.SingleUseTokensExtendedOperationHandlerResponse = nil
		dst.ThirdPartyExtendedOperationHandlerResponse = nil
		dst.ValidateTotpPasswordExtendedOperationHandlerResponse = nil

		return fmt.Errorf("data matches more than one schema in oneOf(AddExtendedOperationHandler200Response)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(AddExtendedOperationHandler200Response)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src AddExtendedOperationHandler200Response) MarshalJSON() ([]byte, error) {
	if src.CollectSupportDataExtendedOperationHandlerResponse != nil {
		return json.Marshal(&src.CollectSupportDataExtendedOperationHandlerResponse)
	}

	if src.DeliverOtpExtendedOperationHandlerResponse != nil {
		return json.Marshal(&src.DeliverOtpExtendedOperationHandlerResponse)
	}

	if src.DeliverPasswordResetTokenExtendedOperationHandlerResponse != nil {
		return json.Marshal(&src.DeliverPasswordResetTokenExtendedOperationHandlerResponse)
	}

	if src.ExportReversiblePasswordsExtendedOperationHandlerResponse != nil {
		return json.Marshal(&src.ExportReversiblePasswordsExtendedOperationHandlerResponse)
	}

	if src.ReplaceCertificateExtendedOperationHandlerResponse != nil {
		return json.Marshal(&src.ReplaceCertificateExtendedOperationHandlerResponse)
	}

	if src.SingleUseTokensExtendedOperationHandlerResponse != nil {
		return json.Marshal(&src.SingleUseTokensExtendedOperationHandlerResponse)
	}

	if src.ThirdPartyExtendedOperationHandlerResponse != nil {
		return json.Marshal(&src.ThirdPartyExtendedOperationHandlerResponse)
	}

	if src.ValidateTotpPasswordExtendedOperationHandlerResponse != nil {
		return json.Marshal(&src.ValidateTotpPasswordExtendedOperationHandlerResponse)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *AddExtendedOperationHandler200Response) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.CollectSupportDataExtendedOperationHandlerResponse != nil {
		return obj.CollectSupportDataExtendedOperationHandlerResponse
	}

	if obj.DeliverOtpExtendedOperationHandlerResponse != nil {
		return obj.DeliverOtpExtendedOperationHandlerResponse
	}

	if obj.DeliverPasswordResetTokenExtendedOperationHandlerResponse != nil {
		return obj.DeliverPasswordResetTokenExtendedOperationHandlerResponse
	}

	if obj.ExportReversiblePasswordsExtendedOperationHandlerResponse != nil {
		return obj.ExportReversiblePasswordsExtendedOperationHandlerResponse
	}

	if obj.ReplaceCertificateExtendedOperationHandlerResponse != nil {
		return obj.ReplaceCertificateExtendedOperationHandlerResponse
	}

	if obj.SingleUseTokensExtendedOperationHandlerResponse != nil {
		return obj.SingleUseTokensExtendedOperationHandlerResponse
	}

	if obj.ThirdPartyExtendedOperationHandlerResponse != nil {
		return obj.ThirdPartyExtendedOperationHandlerResponse
	}

	if obj.ValidateTotpPasswordExtendedOperationHandlerResponse != nil {
		return obj.ValidateTotpPasswordExtendedOperationHandlerResponse
	}

	// all schemas are nil
	return nil
}

type NullableAddExtendedOperationHandler200Response struct {
	value *AddExtendedOperationHandler200Response
	isSet bool
}

func (v NullableAddExtendedOperationHandler200Response) Get() *AddExtendedOperationHandler200Response {
	return v.value
}

func (v *NullableAddExtendedOperationHandler200Response) Set(val *AddExtendedOperationHandler200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableAddExtendedOperationHandler200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableAddExtendedOperationHandler200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddExtendedOperationHandler200Response(val *AddExtendedOperationHandler200Response) *NullableAddExtendedOperationHandler200Response {
	return &NullableAddExtendedOperationHandler200Response{value: val, isSet: true}
}

func (v NullableAddExtendedOperationHandler200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddExtendedOperationHandler200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
