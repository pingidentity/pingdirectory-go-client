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

// GetExtendedOperationHandler200Response - struct for GetExtendedOperationHandler200Response
type GetExtendedOperationHandler200Response struct {
	BatchedTransactionsExtendedOperationHandlerResponse               *BatchedTransactionsExtendedOperationHandlerResponse
	CancelExtendedOperationHandlerResponse                            *CancelExtendedOperationHandlerResponse
	CollectSupportDataExtendedOperationHandlerResponse                *CollectSupportDataExtendedOperationHandlerResponse
	CustomExtendedOperationHandlerResponse                            *CustomExtendedOperationHandlerResponse
	DeliverOtpExtendedOperationHandlerResponse                        *DeliverOtpExtendedOperationHandlerResponse
	DeliverPasswordResetTokenExtendedOperationHandlerResponse         *DeliverPasswordResetTokenExtendedOperationHandlerResponse
	ExportReversiblePasswordsExtendedOperationHandlerResponse         *ExportReversiblePasswordsExtendedOperationHandlerResponse
	GeneratePasswordExtendedOperationHandlerResponse                  *GeneratePasswordExtendedOperationHandlerResponse
	GetChangelogBatchExtendedOperationHandlerResponse                 *GetChangelogBatchExtendedOperationHandlerResponse
	GetConnectionIdExtendedOperationHandlerResponse                   *GetConnectionIdExtendedOperationHandlerResponse
	GetPasswordQualityRequirementsExtendedOperationHandlerResponse    *GetPasswordQualityRequirementsExtendedOperationHandlerResponse
	GetSupportedOtpDeliveryMechanismsExtendedOperationHandlerResponse *GetSupportedOtpDeliveryMechanismsExtendedOperationHandlerResponse
	MultiUpdateExtendedOperationHandlerResponse                       *MultiUpdateExtendedOperationHandlerResponse
	NotificationSubscriptionExtendedOperationHandlerResponse          *NotificationSubscriptionExtendedOperationHandlerResponse
	PasswordModifyExtendedOperationHandlerResponse                    *PasswordModifyExtendedOperationHandlerResponse
	PasswordPolicyStateExtendedOperationHandlerResponse               *PasswordPolicyStateExtendedOperationHandlerResponse
	ReplaceCertificateExtendedOperationHandlerResponse                *ReplaceCertificateExtendedOperationHandlerResponse
	SingleUseTokensExtendedOperationHandlerResponse                   *SingleUseTokensExtendedOperationHandlerResponse
	StartTlsExtendedOperationHandlerResponse                          *StartTlsExtendedOperationHandlerResponse
	ThirdPartyExtendedOperationHandlerResponse                        *ThirdPartyExtendedOperationHandlerResponse
	ValidateTotpPasswordExtendedOperationHandlerResponse              *ValidateTotpPasswordExtendedOperationHandlerResponse
	WhoAmIExtendedOperationHandlerResponse                            *WhoAmIExtendedOperationHandlerResponse
}

// BatchedTransactionsExtendedOperationHandlerResponseAsGetExtendedOperationHandler200Response is a convenience function that returns BatchedTransactionsExtendedOperationHandlerResponse wrapped in GetExtendedOperationHandler200Response
func BatchedTransactionsExtendedOperationHandlerResponseAsGetExtendedOperationHandler200Response(v *BatchedTransactionsExtendedOperationHandlerResponse) GetExtendedOperationHandler200Response {
	return GetExtendedOperationHandler200Response{
		BatchedTransactionsExtendedOperationHandlerResponse: v,
	}
}

// CancelExtendedOperationHandlerResponseAsGetExtendedOperationHandler200Response is a convenience function that returns CancelExtendedOperationHandlerResponse wrapped in GetExtendedOperationHandler200Response
func CancelExtendedOperationHandlerResponseAsGetExtendedOperationHandler200Response(v *CancelExtendedOperationHandlerResponse) GetExtendedOperationHandler200Response {
	return GetExtendedOperationHandler200Response{
		CancelExtendedOperationHandlerResponse: v,
	}
}

// CollectSupportDataExtendedOperationHandlerResponseAsGetExtendedOperationHandler200Response is a convenience function that returns CollectSupportDataExtendedOperationHandlerResponse wrapped in GetExtendedOperationHandler200Response
func CollectSupportDataExtendedOperationHandlerResponseAsGetExtendedOperationHandler200Response(v *CollectSupportDataExtendedOperationHandlerResponse) GetExtendedOperationHandler200Response {
	return GetExtendedOperationHandler200Response{
		CollectSupportDataExtendedOperationHandlerResponse: v,
	}
}

// CustomExtendedOperationHandlerResponseAsGetExtendedOperationHandler200Response is a convenience function that returns CustomExtendedOperationHandlerResponse wrapped in GetExtendedOperationHandler200Response
func CustomExtendedOperationHandlerResponseAsGetExtendedOperationHandler200Response(v *CustomExtendedOperationHandlerResponse) GetExtendedOperationHandler200Response {
	return GetExtendedOperationHandler200Response{
		CustomExtendedOperationHandlerResponse: v,
	}
}

// DeliverOtpExtendedOperationHandlerResponseAsGetExtendedOperationHandler200Response is a convenience function that returns DeliverOtpExtendedOperationHandlerResponse wrapped in GetExtendedOperationHandler200Response
func DeliverOtpExtendedOperationHandlerResponseAsGetExtendedOperationHandler200Response(v *DeliverOtpExtendedOperationHandlerResponse) GetExtendedOperationHandler200Response {
	return GetExtendedOperationHandler200Response{
		DeliverOtpExtendedOperationHandlerResponse: v,
	}
}

// DeliverPasswordResetTokenExtendedOperationHandlerResponseAsGetExtendedOperationHandler200Response is a convenience function that returns DeliverPasswordResetTokenExtendedOperationHandlerResponse wrapped in GetExtendedOperationHandler200Response
func DeliverPasswordResetTokenExtendedOperationHandlerResponseAsGetExtendedOperationHandler200Response(v *DeliverPasswordResetTokenExtendedOperationHandlerResponse) GetExtendedOperationHandler200Response {
	return GetExtendedOperationHandler200Response{
		DeliverPasswordResetTokenExtendedOperationHandlerResponse: v,
	}
}

// ExportReversiblePasswordsExtendedOperationHandlerResponseAsGetExtendedOperationHandler200Response is a convenience function that returns ExportReversiblePasswordsExtendedOperationHandlerResponse wrapped in GetExtendedOperationHandler200Response
func ExportReversiblePasswordsExtendedOperationHandlerResponseAsGetExtendedOperationHandler200Response(v *ExportReversiblePasswordsExtendedOperationHandlerResponse) GetExtendedOperationHandler200Response {
	return GetExtendedOperationHandler200Response{
		ExportReversiblePasswordsExtendedOperationHandlerResponse: v,
	}
}

// GeneratePasswordExtendedOperationHandlerResponseAsGetExtendedOperationHandler200Response is a convenience function that returns GeneratePasswordExtendedOperationHandlerResponse wrapped in GetExtendedOperationHandler200Response
func GeneratePasswordExtendedOperationHandlerResponseAsGetExtendedOperationHandler200Response(v *GeneratePasswordExtendedOperationHandlerResponse) GetExtendedOperationHandler200Response {
	return GetExtendedOperationHandler200Response{
		GeneratePasswordExtendedOperationHandlerResponse: v,
	}
}

// GetChangelogBatchExtendedOperationHandlerResponseAsGetExtendedOperationHandler200Response is a convenience function that returns GetChangelogBatchExtendedOperationHandlerResponse wrapped in GetExtendedOperationHandler200Response
func GetChangelogBatchExtendedOperationHandlerResponseAsGetExtendedOperationHandler200Response(v *GetChangelogBatchExtendedOperationHandlerResponse) GetExtendedOperationHandler200Response {
	return GetExtendedOperationHandler200Response{
		GetChangelogBatchExtendedOperationHandlerResponse: v,
	}
}

// GetConnectionIdExtendedOperationHandlerResponseAsGetExtendedOperationHandler200Response is a convenience function that returns GetConnectionIdExtendedOperationHandlerResponse wrapped in GetExtendedOperationHandler200Response
func GetConnectionIdExtendedOperationHandlerResponseAsGetExtendedOperationHandler200Response(v *GetConnectionIdExtendedOperationHandlerResponse) GetExtendedOperationHandler200Response {
	return GetExtendedOperationHandler200Response{
		GetConnectionIdExtendedOperationHandlerResponse: v,
	}
}

// GetPasswordQualityRequirementsExtendedOperationHandlerResponseAsGetExtendedOperationHandler200Response is a convenience function that returns GetPasswordQualityRequirementsExtendedOperationHandlerResponse wrapped in GetExtendedOperationHandler200Response
func GetPasswordQualityRequirementsExtendedOperationHandlerResponseAsGetExtendedOperationHandler200Response(v *GetPasswordQualityRequirementsExtendedOperationHandlerResponse) GetExtendedOperationHandler200Response {
	return GetExtendedOperationHandler200Response{
		GetPasswordQualityRequirementsExtendedOperationHandlerResponse: v,
	}
}

// GetSupportedOtpDeliveryMechanismsExtendedOperationHandlerResponseAsGetExtendedOperationHandler200Response is a convenience function that returns GetSupportedOtpDeliveryMechanismsExtendedOperationHandlerResponse wrapped in GetExtendedOperationHandler200Response
func GetSupportedOtpDeliveryMechanismsExtendedOperationHandlerResponseAsGetExtendedOperationHandler200Response(v *GetSupportedOtpDeliveryMechanismsExtendedOperationHandlerResponse) GetExtendedOperationHandler200Response {
	return GetExtendedOperationHandler200Response{
		GetSupportedOtpDeliveryMechanismsExtendedOperationHandlerResponse: v,
	}
}

// MultiUpdateExtendedOperationHandlerResponseAsGetExtendedOperationHandler200Response is a convenience function that returns MultiUpdateExtendedOperationHandlerResponse wrapped in GetExtendedOperationHandler200Response
func MultiUpdateExtendedOperationHandlerResponseAsGetExtendedOperationHandler200Response(v *MultiUpdateExtendedOperationHandlerResponse) GetExtendedOperationHandler200Response {
	return GetExtendedOperationHandler200Response{
		MultiUpdateExtendedOperationHandlerResponse: v,
	}
}

// NotificationSubscriptionExtendedOperationHandlerResponseAsGetExtendedOperationHandler200Response is a convenience function that returns NotificationSubscriptionExtendedOperationHandlerResponse wrapped in GetExtendedOperationHandler200Response
func NotificationSubscriptionExtendedOperationHandlerResponseAsGetExtendedOperationHandler200Response(v *NotificationSubscriptionExtendedOperationHandlerResponse) GetExtendedOperationHandler200Response {
	return GetExtendedOperationHandler200Response{
		NotificationSubscriptionExtendedOperationHandlerResponse: v,
	}
}

// PasswordModifyExtendedOperationHandlerResponseAsGetExtendedOperationHandler200Response is a convenience function that returns PasswordModifyExtendedOperationHandlerResponse wrapped in GetExtendedOperationHandler200Response
func PasswordModifyExtendedOperationHandlerResponseAsGetExtendedOperationHandler200Response(v *PasswordModifyExtendedOperationHandlerResponse) GetExtendedOperationHandler200Response {
	return GetExtendedOperationHandler200Response{
		PasswordModifyExtendedOperationHandlerResponse: v,
	}
}

// PasswordPolicyStateExtendedOperationHandlerResponseAsGetExtendedOperationHandler200Response is a convenience function that returns PasswordPolicyStateExtendedOperationHandlerResponse wrapped in GetExtendedOperationHandler200Response
func PasswordPolicyStateExtendedOperationHandlerResponseAsGetExtendedOperationHandler200Response(v *PasswordPolicyStateExtendedOperationHandlerResponse) GetExtendedOperationHandler200Response {
	return GetExtendedOperationHandler200Response{
		PasswordPolicyStateExtendedOperationHandlerResponse: v,
	}
}

// ReplaceCertificateExtendedOperationHandlerResponseAsGetExtendedOperationHandler200Response is a convenience function that returns ReplaceCertificateExtendedOperationHandlerResponse wrapped in GetExtendedOperationHandler200Response
func ReplaceCertificateExtendedOperationHandlerResponseAsGetExtendedOperationHandler200Response(v *ReplaceCertificateExtendedOperationHandlerResponse) GetExtendedOperationHandler200Response {
	return GetExtendedOperationHandler200Response{
		ReplaceCertificateExtendedOperationHandlerResponse: v,
	}
}

// SingleUseTokensExtendedOperationHandlerResponseAsGetExtendedOperationHandler200Response is a convenience function that returns SingleUseTokensExtendedOperationHandlerResponse wrapped in GetExtendedOperationHandler200Response
func SingleUseTokensExtendedOperationHandlerResponseAsGetExtendedOperationHandler200Response(v *SingleUseTokensExtendedOperationHandlerResponse) GetExtendedOperationHandler200Response {
	return GetExtendedOperationHandler200Response{
		SingleUseTokensExtendedOperationHandlerResponse: v,
	}
}

// StartTlsExtendedOperationHandlerResponseAsGetExtendedOperationHandler200Response is a convenience function that returns StartTlsExtendedOperationHandlerResponse wrapped in GetExtendedOperationHandler200Response
func StartTlsExtendedOperationHandlerResponseAsGetExtendedOperationHandler200Response(v *StartTlsExtendedOperationHandlerResponse) GetExtendedOperationHandler200Response {
	return GetExtendedOperationHandler200Response{
		StartTlsExtendedOperationHandlerResponse: v,
	}
}

// ThirdPartyExtendedOperationHandlerResponseAsGetExtendedOperationHandler200Response is a convenience function that returns ThirdPartyExtendedOperationHandlerResponse wrapped in GetExtendedOperationHandler200Response
func ThirdPartyExtendedOperationHandlerResponseAsGetExtendedOperationHandler200Response(v *ThirdPartyExtendedOperationHandlerResponse) GetExtendedOperationHandler200Response {
	return GetExtendedOperationHandler200Response{
		ThirdPartyExtendedOperationHandlerResponse: v,
	}
}

// ValidateTotpPasswordExtendedOperationHandlerResponseAsGetExtendedOperationHandler200Response is a convenience function that returns ValidateTotpPasswordExtendedOperationHandlerResponse wrapped in GetExtendedOperationHandler200Response
func ValidateTotpPasswordExtendedOperationHandlerResponseAsGetExtendedOperationHandler200Response(v *ValidateTotpPasswordExtendedOperationHandlerResponse) GetExtendedOperationHandler200Response {
	return GetExtendedOperationHandler200Response{
		ValidateTotpPasswordExtendedOperationHandlerResponse: v,
	}
}

// WhoAmIExtendedOperationHandlerResponseAsGetExtendedOperationHandler200Response is a convenience function that returns WhoAmIExtendedOperationHandlerResponse wrapped in GetExtendedOperationHandler200Response
func WhoAmIExtendedOperationHandlerResponseAsGetExtendedOperationHandler200Response(v *WhoAmIExtendedOperationHandlerResponse) GetExtendedOperationHandler200Response {
	return GetExtendedOperationHandler200Response{
		WhoAmIExtendedOperationHandlerResponse: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *GetExtendedOperationHandler200Response) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into BatchedTransactionsExtendedOperationHandlerResponse
	err = newStrictDecoder(data).Decode(&dst.BatchedTransactionsExtendedOperationHandlerResponse)
	if err == nil {
		jsonBatchedTransactionsExtendedOperationHandlerResponse, _ := json.Marshal(dst.BatchedTransactionsExtendedOperationHandlerResponse)
		if string(jsonBatchedTransactionsExtendedOperationHandlerResponse) == "{}" { // empty struct
			dst.BatchedTransactionsExtendedOperationHandlerResponse = nil
		} else {
			match++
		}
	} else {
		dst.BatchedTransactionsExtendedOperationHandlerResponse = nil
	}

	// try to unmarshal data into CancelExtendedOperationHandlerResponse
	err = newStrictDecoder(data).Decode(&dst.CancelExtendedOperationHandlerResponse)
	if err == nil {
		jsonCancelExtendedOperationHandlerResponse, _ := json.Marshal(dst.CancelExtendedOperationHandlerResponse)
		if string(jsonCancelExtendedOperationHandlerResponse) == "{}" { // empty struct
			dst.CancelExtendedOperationHandlerResponse = nil
		} else {
			match++
		}
	} else {
		dst.CancelExtendedOperationHandlerResponse = nil
	}

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

	// try to unmarshal data into CustomExtendedOperationHandlerResponse
	err = newStrictDecoder(data).Decode(&dst.CustomExtendedOperationHandlerResponse)
	if err == nil {
		jsonCustomExtendedOperationHandlerResponse, _ := json.Marshal(dst.CustomExtendedOperationHandlerResponse)
		if string(jsonCustomExtendedOperationHandlerResponse) == "{}" { // empty struct
			dst.CustomExtendedOperationHandlerResponse = nil
		} else {
			match++
		}
	} else {
		dst.CustomExtendedOperationHandlerResponse = nil
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

	// try to unmarshal data into GeneratePasswordExtendedOperationHandlerResponse
	err = newStrictDecoder(data).Decode(&dst.GeneratePasswordExtendedOperationHandlerResponse)
	if err == nil {
		jsonGeneratePasswordExtendedOperationHandlerResponse, _ := json.Marshal(dst.GeneratePasswordExtendedOperationHandlerResponse)
		if string(jsonGeneratePasswordExtendedOperationHandlerResponse) == "{}" { // empty struct
			dst.GeneratePasswordExtendedOperationHandlerResponse = nil
		} else {
			match++
		}
	} else {
		dst.GeneratePasswordExtendedOperationHandlerResponse = nil
	}

	// try to unmarshal data into GetChangelogBatchExtendedOperationHandlerResponse
	err = newStrictDecoder(data).Decode(&dst.GetChangelogBatchExtendedOperationHandlerResponse)
	if err == nil {
		jsonGetChangelogBatchExtendedOperationHandlerResponse, _ := json.Marshal(dst.GetChangelogBatchExtendedOperationHandlerResponse)
		if string(jsonGetChangelogBatchExtendedOperationHandlerResponse) == "{}" { // empty struct
			dst.GetChangelogBatchExtendedOperationHandlerResponse = nil
		} else {
			match++
		}
	} else {
		dst.GetChangelogBatchExtendedOperationHandlerResponse = nil
	}

	// try to unmarshal data into GetConnectionIdExtendedOperationHandlerResponse
	err = newStrictDecoder(data).Decode(&dst.GetConnectionIdExtendedOperationHandlerResponse)
	if err == nil {
		jsonGetConnectionIdExtendedOperationHandlerResponse, _ := json.Marshal(dst.GetConnectionIdExtendedOperationHandlerResponse)
		if string(jsonGetConnectionIdExtendedOperationHandlerResponse) == "{}" { // empty struct
			dst.GetConnectionIdExtendedOperationHandlerResponse = nil
		} else {
			match++
		}
	} else {
		dst.GetConnectionIdExtendedOperationHandlerResponse = nil
	}

	// try to unmarshal data into GetPasswordQualityRequirementsExtendedOperationHandlerResponse
	err = newStrictDecoder(data).Decode(&dst.GetPasswordQualityRequirementsExtendedOperationHandlerResponse)
	if err == nil {
		jsonGetPasswordQualityRequirementsExtendedOperationHandlerResponse, _ := json.Marshal(dst.GetPasswordQualityRequirementsExtendedOperationHandlerResponse)
		if string(jsonGetPasswordQualityRequirementsExtendedOperationHandlerResponse) == "{}" { // empty struct
			dst.GetPasswordQualityRequirementsExtendedOperationHandlerResponse = nil
		} else {
			match++
		}
	} else {
		dst.GetPasswordQualityRequirementsExtendedOperationHandlerResponse = nil
	}

	// try to unmarshal data into GetSupportedOtpDeliveryMechanismsExtendedOperationHandlerResponse
	err = newStrictDecoder(data).Decode(&dst.GetSupportedOtpDeliveryMechanismsExtendedOperationHandlerResponse)
	if err == nil {
		jsonGetSupportedOtpDeliveryMechanismsExtendedOperationHandlerResponse, _ := json.Marshal(dst.GetSupportedOtpDeliveryMechanismsExtendedOperationHandlerResponse)
		if string(jsonGetSupportedOtpDeliveryMechanismsExtendedOperationHandlerResponse) == "{}" { // empty struct
			dst.GetSupportedOtpDeliveryMechanismsExtendedOperationHandlerResponse = nil
		} else {
			match++
		}
	} else {
		dst.GetSupportedOtpDeliveryMechanismsExtendedOperationHandlerResponse = nil
	}

	// try to unmarshal data into MultiUpdateExtendedOperationHandlerResponse
	err = newStrictDecoder(data).Decode(&dst.MultiUpdateExtendedOperationHandlerResponse)
	if err == nil {
		jsonMultiUpdateExtendedOperationHandlerResponse, _ := json.Marshal(dst.MultiUpdateExtendedOperationHandlerResponse)
		if string(jsonMultiUpdateExtendedOperationHandlerResponse) == "{}" { // empty struct
			dst.MultiUpdateExtendedOperationHandlerResponse = nil
		} else {
			match++
		}
	} else {
		dst.MultiUpdateExtendedOperationHandlerResponse = nil
	}

	// try to unmarshal data into NotificationSubscriptionExtendedOperationHandlerResponse
	err = newStrictDecoder(data).Decode(&dst.NotificationSubscriptionExtendedOperationHandlerResponse)
	if err == nil {
		jsonNotificationSubscriptionExtendedOperationHandlerResponse, _ := json.Marshal(dst.NotificationSubscriptionExtendedOperationHandlerResponse)
		if string(jsonNotificationSubscriptionExtendedOperationHandlerResponse) == "{}" { // empty struct
			dst.NotificationSubscriptionExtendedOperationHandlerResponse = nil
		} else {
			match++
		}
	} else {
		dst.NotificationSubscriptionExtendedOperationHandlerResponse = nil
	}

	// try to unmarshal data into PasswordModifyExtendedOperationHandlerResponse
	err = newStrictDecoder(data).Decode(&dst.PasswordModifyExtendedOperationHandlerResponse)
	if err == nil {
		jsonPasswordModifyExtendedOperationHandlerResponse, _ := json.Marshal(dst.PasswordModifyExtendedOperationHandlerResponse)
		if string(jsonPasswordModifyExtendedOperationHandlerResponse) == "{}" { // empty struct
			dst.PasswordModifyExtendedOperationHandlerResponse = nil
		} else {
			match++
		}
	} else {
		dst.PasswordModifyExtendedOperationHandlerResponse = nil
	}

	// try to unmarshal data into PasswordPolicyStateExtendedOperationHandlerResponse
	err = newStrictDecoder(data).Decode(&dst.PasswordPolicyStateExtendedOperationHandlerResponse)
	if err == nil {
		jsonPasswordPolicyStateExtendedOperationHandlerResponse, _ := json.Marshal(dst.PasswordPolicyStateExtendedOperationHandlerResponse)
		if string(jsonPasswordPolicyStateExtendedOperationHandlerResponse) == "{}" { // empty struct
			dst.PasswordPolicyStateExtendedOperationHandlerResponse = nil
		} else {
			match++
		}
	} else {
		dst.PasswordPolicyStateExtendedOperationHandlerResponse = nil
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

	// try to unmarshal data into StartTlsExtendedOperationHandlerResponse
	err = newStrictDecoder(data).Decode(&dst.StartTlsExtendedOperationHandlerResponse)
	if err == nil {
		jsonStartTlsExtendedOperationHandlerResponse, _ := json.Marshal(dst.StartTlsExtendedOperationHandlerResponse)
		if string(jsonStartTlsExtendedOperationHandlerResponse) == "{}" { // empty struct
			dst.StartTlsExtendedOperationHandlerResponse = nil
		} else {
			match++
		}
	} else {
		dst.StartTlsExtendedOperationHandlerResponse = nil
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

	// try to unmarshal data into WhoAmIExtendedOperationHandlerResponse
	err = newStrictDecoder(data).Decode(&dst.WhoAmIExtendedOperationHandlerResponse)
	if err == nil {
		jsonWhoAmIExtendedOperationHandlerResponse, _ := json.Marshal(dst.WhoAmIExtendedOperationHandlerResponse)
		if string(jsonWhoAmIExtendedOperationHandlerResponse) == "{}" { // empty struct
			dst.WhoAmIExtendedOperationHandlerResponse = nil
		} else {
			match++
		}
	} else {
		dst.WhoAmIExtendedOperationHandlerResponse = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.BatchedTransactionsExtendedOperationHandlerResponse = nil
		dst.CancelExtendedOperationHandlerResponse = nil
		dst.CollectSupportDataExtendedOperationHandlerResponse = nil
		dst.CustomExtendedOperationHandlerResponse = nil
		dst.DeliverOtpExtendedOperationHandlerResponse = nil
		dst.DeliverPasswordResetTokenExtendedOperationHandlerResponse = nil
		dst.ExportReversiblePasswordsExtendedOperationHandlerResponse = nil
		dst.GeneratePasswordExtendedOperationHandlerResponse = nil
		dst.GetChangelogBatchExtendedOperationHandlerResponse = nil
		dst.GetConnectionIdExtendedOperationHandlerResponse = nil
		dst.GetPasswordQualityRequirementsExtendedOperationHandlerResponse = nil
		dst.GetSupportedOtpDeliveryMechanismsExtendedOperationHandlerResponse = nil
		dst.MultiUpdateExtendedOperationHandlerResponse = nil
		dst.NotificationSubscriptionExtendedOperationHandlerResponse = nil
		dst.PasswordModifyExtendedOperationHandlerResponse = nil
		dst.PasswordPolicyStateExtendedOperationHandlerResponse = nil
		dst.ReplaceCertificateExtendedOperationHandlerResponse = nil
		dst.SingleUseTokensExtendedOperationHandlerResponse = nil
		dst.StartTlsExtendedOperationHandlerResponse = nil
		dst.ThirdPartyExtendedOperationHandlerResponse = nil
		dst.ValidateTotpPasswordExtendedOperationHandlerResponse = nil
		dst.WhoAmIExtendedOperationHandlerResponse = nil

		return fmt.Errorf("data matches more than one schema in oneOf(GetExtendedOperationHandler200Response)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(GetExtendedOperationHandler200Response)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GetExtendedOperationHandler200Response) MarshalJSON() ([]byte, error) {
	if src.BatchedTransactionsExtendedOperationHandlerResponse != nil {
		return json.Marshal(&src.BatchedTransactionsExtendedOperationHandlerResponse)
	}

	if src.CancelExtendedOperationHandlerResponse != nil {
		return json.Marshal(&src.CancelExtendedOperationHandlerResponse)
	}

	if src.CollectSupportDataExtendedOperationHandlerResponse != nil {
		return json.Marshal(&src.CollectSupportDataExtendedOperationHandlerResponse)
	}

	if src.CustomExtendedOperationHandlerResponse != nil {
		return json.Marshal(&src.CustomExtendedOperationHandlerResponse)
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

	if src.GeneratePasswordExtendedOperationHandlerResponse != nil {
		return json.Marshal(&src.GeneratePasswordExtendedOperationHandlerResponse)
	}

	if src.GetChangelogBatchExtendedOperationHandlerResponse != nil {
		return json.Marshal(&src.GetChangelogBatchExtendedOperationHandlerResponse)
	}

	if src.GetConnectionIdExtendedOperationHandlerResponse != nil {
		return json.Marshal(&src.GetConnectionIdExtendedOperationHandlerResponse)
	}

	if src.GetPasswordQualityRequirementsExtendedOperationHandlerResponse != nil {
		return json.Marshal(&src.GetPasswordQualityRequirementsExtendedOperationHandlerResponse)
	}

	if src.GetSupportedOtpDeliveryMechanismsExtendedOperationHandlerResponse != nil {
		return json.Marshal(&src.GetSupportedOtpDeliveryMechanismsExtendedOperationHandlerResponse)
	}

	if src.MultiUpdateExtendedOperationHandlerResponse != nil {
		return json.Marshal(&src.MultiUpdateExtendedOperationHandlerResponse)
	}

	if src.NotificationSubscriptionExtendedOperationHandlerResponse != nil {
		return json.Marshal(&src.NotificationSubscriptionExtendedOperationHandlerResponse)
	}

	if src.PasswordModifyExtendedOperationHandlerResponse != nil {
		return json.Marshal(&src.PasswordModifyExtendedOperationHandlerResponse)
	}

	if src.PasswordPolicyStateExtendedOperationHandlerResponse != nil {
		return json.Marshal(&src.PasswordPolicyStateExtendedOperationHandlerResponse)
	}

	if src.ReplaceCertificateExtendedOperationHandlerResponse != nil {
		return json.Marshal(&src.ReplaceCertificateExtendedOperationHandlerResponse)
	}

	if src.SingleUseTokensExtendedOperationHandlerResponse != nil {
		return json.Marshal(&src.SingleUseTokensExtendedOperationHandlerResponse)
	}

	if src.StartTlsExtendedOperationHandlerResponse != nil {
		return json.Marshal(&src.StartTlsExtendedOperationHandlerResponse)
	}

	if src.ThirdPartyExtendedOperationHandlerResponse != nil {
		return json.Marshal(&src.ThirdPartyExtendedOperationHandlerResponse)
	}

	if src.ValidateTotpPasswordExtendedOperationHandlerResponse != nil {
		return json.Marshal(&src.ValidateTotpPasswordExtendedOperationHandlerResponse)
	}

	if src.WhoAmIExtendedOperationHandlerResponse != nil {
		return json.Marshal(&src.WhoAmIExtendedOperationHandlerResponse)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GetExtendedOperationHandler200Response) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.BatchedTransactionsExtendedOperationHandlerResponse != nil {
		return obj.BatchedTransactionsExtendedOperationHandlerResponse
	}

	if obj.CancelExtendedOperationHandlerResponse != nil {
		return obj.CancelExtendedOperationHandlerResponse
	}

	if obj.CollectSupportDataExtendedOperationHandlerResponse != nil {
		return obj.CollectSupportDataExtendedOperationHandlerResponse
	}

	if obj.CustomExtendedOperationHandlerResponse != nil {
		return obj.CustomExtendedOperationHandlerResponse
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

	if obj.GeneratePasswordExtendedOperationHandlerResponse != nil {
		return obj.GeneratePasswordExtendedOperationHandlerResponse
	}

	if obj.GetChangelogBatchExtendedOperationHandlerResponse != nil {
		return obj.GetChangelogBatchExtendedOperationHandlerResponse
	}

	if obj.GetConnectionIdExtendedOperationHandlerResponse != nil {
		return obj.GetConnectionIdExtendedOperationHandlerResponse
	}

	if obj.GetPasswordQualityRequirementsExtendedOperationHandlerResponse != nil {
		return obj.GetPasswordQualityRequirementsExtendedOperationHandlerResponse
	}

	if obj.GetSupportedOtpDeliveryMechanismsExtendedOperationHandlerResponse != nil {
		return obj.GetSupportedOtpDeliveryMechanismsExtendedOperationHandlerResponse
	}

	if obj.MultiUpdateExtendedOperationHandlerResponse != nil {
		return obj.MultiUpdateExtendedOperationHandlerResponse
	}

	if obj.NotificationSubscriptionExtendedOperationHandlerResponse != nil {
		return obj.NotificationSubscriptionExtendedOperationHandlerResponse
	}

	if obj.PasswordModifyExtendedOperationHandlerResponse != nil {
		return obj.PasswordModifyExtendedOperationHandlerResponse
	}

	if obj.PasswordPolicyStateExtendedOperationHandlerResponse != nil {
		return obj.PasswordPolicyStateExtendedOperationHandlerResponse
	}

	if obj.ReplaceCertificateExtendedOperationHandlerResponse != nil {
		return obj.ReplaceCertificateExtendedOperationHandlerResponse
	}

	if obj.SingleUseTokensExtendedOperationHandlerResponse != nil {
		return obj.SingleUseTokensExtendedOperationHandlerResponse
	}

	if obj.StartTlsExtendedOperationHandlerResponse != nil {
		return obj.StartTlsExtendedOperationHandlerResponse
	}

	if obj.ThirdPartyExtendedOperationHandlerResponse != nil {
		return obj.ThirdPartyExtendedOperationHandlerResponse
	}

	if obj.ValidateTotpPasswordExtendedOperationHandlerResponse != nil {
		return obj.ValidateTotpPasswordExtendedOperationHandlerResponse
	}

	if obj.WhoAmIExtendedOperationHandlerResponse != nil {
		return obj.WhoAmIExtendedOperationHandlerResponse
	}

	// all schemas are nil
	return nil
}

type NullableGetExtendedOperationHandler200Response struct {
	value *GetExtendedOperationHandler200Response
	isSet bool
}

func (v NullableGetExtendedOperationHandler200Response) Get() *GetExtendedOperationHandler200Response {
	return v.value
}

func (v *NullableGetExtendedOperationHandler200Response) Set(val *GetExtendedOperationHandler200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetExtendedOperationHandler200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetExtendedOperationHandler200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetExtendedOperationHandler200Response(val *GetExtendedOperationHandler200Response) *NullableGetExtendedOperationHandler200Response {
	return &NullableGetExtendedOperationHandler200Response{value: val, isSet: true}
}

func (v NullableGetExtendedOperationHandler200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetExtendedOperationHandler200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
