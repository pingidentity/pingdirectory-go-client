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

// GetSaslMechanismHandler200Response - struct for GetSaslMechanismHandler200Response
type GetSaslMechanismHandler200Response struct {
	AnonymousSaslMechanismHandlerResponse *AnonymousSaslMechanismHandlerResponse
	CramMd5SaslMechanismHandlerResponse *CramMd5SaslMechanismHandlerResponse
	DigestMd5SaslMechanismHandlerResponse *DigestMd5SaslMechanismHandlerResponse
	ExternalSaslMechanismHandlerResponse *ExternalSaslMechanismHandlerResponse
	GssapiSaslMechanismHandlerResponse *GssapiSaslMechanismHandlerResponse
	OauthBearerSaslMechanismHandlerResponse *OauthBearerSaslMechanismHandlerResponse
	PlainSaslMechanismHandlerResponse *PlainSaslMechanismHandlerResponse
	ThirdPartySaslMechanismHandlerResponse *ThirdPartySaslMechanismHandlerResponse
	UnboundidCertificatePlusPasswordSaslMechanismHandlerResponse *UnboundidCertificatePlusPasswordSaslMechanismHandlerResponse
	UnboundidDeliveredOtpSaslMechanismHandlerResponse *UnboundidDeliveredOtpSaslMechanismHandlerResponse
	UnboundidExternalAuthSaslMechanismHandlerResponse *UnboundidExternalAuthSaslMechanismHandlerResponse
	UnboundidMsChapV2SaslMechanismHandlerResponse *UnboundidMsChapV2SaslMechanismHandlerResponse
	UnboundidTotpSaslMechanismHandlerResponse *UnboundidTotpSaslMechanismHandlerResponse
	UnboundidYubikeyOtpSaslMechanismHandlerResponse *UnboundidYubikeyOtpSaslMechanismHandlerResponse
}

// AnonymousSaslMechanismHandlerResponseAsGetSaslMechanismHandler200Response is a convenience function that returns AnonymousSaslMechanismHandlerResponse wrapped in GetSaslMechanismHandler200Response
func AnonymousSaslMechanismHandlerResponseAsGetSaslMechanismHandler200Response(v *AnonymousSaslMechanismHandlerResponse) GetSaslMechanismHandler200Response {
	return GetSaslMechanismHandler200Response{
		AnonymousSaslMechanismHandlerResponse: v,
	}
}

// CramMd5SaslMechanismHandlerResponseAsGetSaslMechanismHandler200Response is a convenience function that returns CramMd5SaslMechanismHandlerResponse wrapped in GetSaslMechanismHandler200Response
func CramMd5SaslMechanismHandlerResponseAsGetSaslMechanismHandler200Response(v *CramMd5SaslMechanismHandlerResponse) GetSaslMechanismHandler200Response {
	return GetSaslMechanismHandler200Response{
		CramMd5SaslMechanismHandlerResponse: v,
	}
}

// DigestMd5SaslMechanismHandlerResponseAsGetSaslMechanismHandler200Response is a convenience function that returns DigestMd5SaslMechanismHandlerResponse wrapped in GetSaslMechanismHandler200Response
func DigestMd5SaslMechanismHandlerResponseAsGetSaslMechanismHandler200Response(v *DigestMd5SaslMechanismHandlerResponse) GetSaslMechanismHandler200Response {
	return GetSaslMechanismHandler200Response{
		DigestMd5SaslMechanismHandlerResponse: v,
	}
}

// ExternalSaslMechanismHandlerResponseAsGetSaslMechanismHandler200Response is a convenience function that returns ExternalSaslMechanismHandlerResponse wrapped in GetSaslMechanismHandler200Response
func ExternalSaslMechanismHandlerResponseAsGetSaslMechanismHandler200Response(v *ExternalSaslMechanismHandlerResponse) GetSaslMechanismHandler200Response {
	return GetSaslMechanismHandler200Response{
		ExternalSaslMechanismHandlerResponse: v,
	}
}

// GssapiSaslMechanismHandlerResponseAsGetSaslMechanismHandler200Response is a convenience function that returns GssapiSaslMechanismHandlerResponse wrapped in GetSaslMechanismHandler200Response
func GssapiSaslMechanismHandlerResponseAsGetSaslMechanismHandler200Response(v *GssapiSaslMechanismHandlerResponse) GetSaslMechanismHandler200Response {
	return GetSaslMechanismHandler200Response{
		GssapiSaslMechanismHandlerResponse: v,
	}
}

// OauthBearerSaslMechanismHandlerResponseAsGetSaslMechanismHandler200Response is a convenience function that returns OauthBearerSaslMechanismHandlerResponse wrapped in GetSaslMechanismHandler200Response
func OauthBearerSaslMechanismHandlerResponseAsGetSaslMechanismHandler200Response(v *OauthBearerSaslMechanismHandlerResponse) GetSaslMechanismHandler200Response {
	return GetSaslMechanismHandler200Response{
		OauthBearerSaslMechanismHandlerResponse: v,
	}
}

// PlainSaslMechanismHandlerResponseAsGetSaslMechanismHandler200Response is a convenience function that returns PlainSaslMechanismHandlerResponse wrapped in GetSaslMechanismHandler200Response
func PlainSaslMechanismHandlerResponseAsGetSaslMechanismHandler200Response(v *PlainSaslMechanismHandlerResponse) GetSaslMechanismHandler200Response {
	return GetSaslMechanismHandler200Response{
		PlainSaslMechanismHandlerResponse: v,
	}
}

// ThirdPartySaslMechanismHandlerResponseAsGetSaslMechanismHandler200Response is a convenience function that returns ThirdPartySaslMechanismHandlerResponse wrapped in GetSaslMechanismHandler200Response
func ThirdPartySaslMechanismHandlerResponseAsGetSaslMechanismHandler200Response(v *ThirdPartySaslMechanismHandlerResponse) GetSaslMechanismHandler200Response {
	return GetSaslMechanismHandler200Response{
		ThirdPartySaslMechanismHandlerResponse: v,
	}
}

// UnboundidCertificatePlusPasswordSaslMechanismHandlerResponseAsGetSaslMechanismHandler200Response is a convenience function that returns UnboundidCertificatePlusPasswordSaslMechanismHandlerResponse wrapped in GetSaslMechanismHandler200Response
func UnboundidCertificatePlusPasswordSaslMechanismHandlerResponseAsGetSaslMechanismHandler200Response(v *UnboundidCertificatePlusPasswordSaslMechanismHandlerResponse) GetSaslMechanismHandler200Response {
	return GetSaslMechanismHandler200Response{
		UnboundidCertificatePlusPasswordSaslMechanismHandlerResponse: v,
	}
}

// UnboundidDeliveredOtpSaslMechanismHandlerResponseAsGetSaslMechanismHandler200Response is a convenience function that returns UnboundidDeliveredOtpSaslMechanismHandlerResponse wrapped in GetSaslMechanismHandler200Response
func UnboundidDeliveredOtpSaslMechanismHandlerResponseAsGetSaslMechanismHandler200Response(v *UnboundidDeliveredOtpSaslMechanismHandlerResponse) GetSaslMechanismHandler200Response {
	return GetSaslMechanismHandler200Response{
		UnboundidDeliveredOtpSaslMechanismHandlerResponse: v,
	}
}

// UnboundidExternalAuthSaslMechanismHandlerResponseAsGetSaslMechanismHandler200Response is a convenience function that returns UnboundidExternalAuthSaslMechanismHandlerResponse wrapped in GetSaslMechanismHandler200Response
func UnboundidExternalAuthSaslMechanismHandlerResponseAsGetSaslMechanismHandler200Response(v *UnboundidExternalAuthSaslMechanismHandlerResponse) GetSaslMechanismHandler200Response {
	return GetSaslMechanismHandler200Response{
		UnboundidExternalAuthSaslMechanismHandlerResponse: v,
	}
}

// UnboundidMsChapV2SaslMechanismHandlerResponseAsGetSaslMechanismHandler200Response is a convenience function that returns UnboundidMsChapV2SaslMechanismHandlerResponse wrapped in GetSaslMechanismHandler200Response
func UnboundidMsChapV2SaslMechanismHandlerResponseAsGetSaslMechanismHandler200Response(v *UnboundidMsChapV2SaslMechanismHandlerResponse) GetSaslMechanismHandler200Response {
	return GetSaslMechanismHandler200Response{
		UnboundidMsChapV2SaslMechanismHandlerResponse: v,
	}
}

// UnboundidTotpSaslMechanismHandlerResponseAsGetSaslMechanismHandler200Response is a convenience function that returns UnboundidTotpSaslMechanismHandlerResponse wrapped in GetSaslMechanismHandler200Response
func UnboundidTotpSaslMechanismHandlerResponseAsGetSaslMechanismHandler200Response(v *UnboundidTotpSaslMechanismHandlerResponse) GetSaslMechanismHandler200Response {
	return GetSaslMechanismHandler200Response{
		UnboundidTotpSaslMechanismHandlerResponse: v,
	}
}

// UnboundidYubikeyOtpSaslMechanismHandlerResponseAsGetSaslMechanismHandler200Response is a convenience function that returns UnboundidYubikeyOtpSaslMechanismHandlerResponse wrapped in GetSaslMechanismHandler200Response
func UnboundidYubikeyOtpSaslMechanismHandlerResponseAsGetSaslMechanismHandler200Response(v *UnboundidYubikeyOtpSaslMechanismHandlerResponse) GetSaslMechanismHandler200Response {
	return GetSaslMechanismHandler200Response{
		UnboundidYubikeyOtpSaslMechanismHandlerResponse: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *GetSaslMechanismHandler200Response) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AnonymousSaslMechanismHandlerResponse
	err = newStrictDecoder(data).Decode(&dst.AnonymousSaslMechanismHandlerResponse)
	if err == nil {
		jsonAnonymousSaslMechanismHandlerResponse, _ := json.Marshal(dst.AnonymousSaslMechanismHandlerResponse)
		if string(jsonAnonymousSaslMechanismHandlerResponse) == "{}" { // empty struct
			dst.AnonymousSaslMechanismHandlerResponse = nil
		} else {
			match++
		}
	} else {
		dst.AnonymousSaslMechanismHandlerResponse = nil
	}

	// try to unmarshal data into CramMd5SaslMechanismHandlerResponse
	err = newStrictDecoder(data).Decode(&dst.CramMd5SaslMechanismHandlerResponse)
	if err == nil {
		jsonCramMd5SaslMechanismHandlerResponse, _ := json.Marshal(dst.CramMd5SaslMechanismHandlerResponse)
		if string(jsonCramMd5SaslMechanismHandlerResponse) == "{}" { // empty struct
			dst.CramMd5SaslMechanismHandlerResponse = nil
		} else {
			match++
		}
	} else {
		dst.CramMd5SaslMechanismHandlerResponse = nil
	}

	// try to unmarshal data into DigestMd5SaslMechanismHandlerResponse
	err = newStrictDecoder(data).Decode(&dst.DigestMd5SaslMechanismHandlerResponse)
	if err == nil {
		jsonDigestMd5SaslMechanismHandlerResponse, _ := json.Marshal(dst.DigestMd5SaslMechanismHandlerResponse)
		if string(jsonDigestMd5SaslMechanismHandlerResponse) == "{}" { // empty struct
			dst.DigestMd5SaslMechanismHandlerResponse = nil
		} else {
			match++
		}
	} else {
		dst.DigestMd5SaslMechanismHandlerResponse = nil
	}

	// try to unmarshal data into ExternalSaslMechanismHandlerResponse
	err = newStrictDecoder(data).Decode(&dst.ExternalSaslMechanismHandlerResponse)
	if err == nil {
		jsonExternalSaslMechanismHandlerResponse, _ := json.Marshal(dst.ExternalSaslMechanismHandlerResponse)
		if string(jsonExternalSaslMechanismHandlerResponse) == "{}" { // empty struct
			dst.ExternalSaslMechanismHandlerResponse = nil
		} else {
			match++
		}
	} else {
		dst.ExternalSaslMechanismHandlerResponse = nil
	}

	// try to unmarshal data into GssapiSaslMechanismHandlerResponse
	err = newStrictDecoder(data).Decode(&dst.GssapiSaslMechanismHandlerResponse)
	if err == nil {
		jsonGssapiSaslMechanismHandlerResponse, _ := json.Marshal(dst.GssapiSaslMechanismHandlerResponse)
		if string(jsonGssapiSaslMechanismHandlerResponse) == "{}" { // empty struct
			dst.GssapiSaslMechanismHandlerResponse = nil
		} else {
			match++
		}
	} else {
		dst.GssapiSaslMechanismHandlerResponse = nil
	}

	// try to unmarshal data into OauthBearerSaslMechanismHandlerResponse
	err = newStrictDecoder(data).Decode(&dst.OauthBearerSaslMechanismHandlerResponse)
	if err == nil {
		jsonOauthBearerSaslMechanismHandlerResponse, _ := json.Marshal(dst.OauthBearerSaslMechanismHandlerResponse)
		if string(jsonOauthBearerSaslMechanismHandlerResponse) == "{}" { // empty struct
			dst.OauthBearerSaslMechanismHandlerResponse = nil
		} else {
			match++
		}
	} else {
		dst.OauthBearerSaslMechanismHandlerResponse = nil
	}

	// try to unmarshal data into PlainSaslMechanismHandlerResponse
	err = newStrictDecoder(data).Decode(&dst.PlainSaslMechanismHandlerResponse)
	if err == nil {
		jsonPlainSaslMechanismHandlerResponse, _ := json.Marshal(dst.PlainSaslMechanismHandlerResponse)
		if string(jsonPlainSaslMechanismHandlerResponse) == "{}" { // empty struct
			dst.PlainSaslMechanismHandlerResponse = nil
		} else {
			match++
		}
	} else {
		dst.PlainSaslMechanismHandlerResponse = nil
	}

	// try to unmarshal data into ThirdPartySaslMechanismHandlerResponse
	err = newStrictDecoder(data).Decode(&dst.ThirdPartySaslMechanismHandlerResponse)
	if err == nil {
		jsonThirdPartySaslMechanismHandlerResponse, _ := json.Marshal(dst.ThirdPartySaslMechanismHandlerResponse)
		if string(jsonThirdPartySaslMechanismHandlerResponse) == "{}" { // empty struct
			dst.ThirdPartySaslMechanismHandlerResponse = nil
		} else {
			match++
		}
	} else {
		dst.ThirdPartySaslMechanismHandlerResponse = nil
	}

	// try to unmarshal data into UnboundidCertificatePlusPasswordSaslMechanismHandlerResponse
	err = newStrictDecoder(data).Decode(&dst.UnboundidCertificatePlusPasswordSaslMechanismHandlerResponse)
	if err == nil {
		jsonUnboundidCertificatePlusPasswordSaslMechanismHandlerResponse, _ := json.Marshal(dst.UnboundidCertificatePlusPasswordSaslMechanismHandlerResponse)
		if string(jsonUnboundidCertificatePlusPasswordSaslMechanismHandlerResponse) == "{}" { // empty struct
			dst.UnboundidCertificatePlusPasswordSaslMechanismHandlerResponse = nil
		} else {
			match++
		}
	} else {
		dst.UnboundidCertificatePlusPasswordSaslMechanismHandlerResponse = nil
	}

	// try to unmarshal data into UnboundidDeliveredOtpSaslMechanismHandlerResponse
	err = newStrictDecoder(data).Decode(&dst.UnboundidDeliveredOtpSaslMechanismHandlerResponse)
	if err == nil {
		jsonUnboundidDeliveredOtpSaslMechanismHandlerResponse, _ := json.Marshal(dst.UnboundidDeliveredOtpSaslMechanismHandlerResponse)
		if string(jsonUnboundidDeliveredOtpSaslMechanismHandlerResponse) == "{}" { // empty struct
			dst.UnboundidDeliveredOtpSaslMechanismHandlerResponse = nil
		} else {
			match++
		}
	} else {
		dst.UnboundidDeliveredOtpSaslMechanismHandlerResponse = nil
	}

	// try to unmarshal data into UnboundidExternalAuthSaslMechanismHandlerResponse
	err = newStrictDecoder(data).Decode(&dst.UnboundidExternalAuthSaslMechanismHandlerResponse)
	if err == nil {
		jsonUnboundidExternalAuthSaslMechanismHandlerResponse, _ := json.Marshal(dst.UnboundidExternalAuthSaslMechanismHandlerResponse)
		if string(jsonUnboundidExternalAuthSaslMechanismHandlerResponse) == "{}" { // empty struct
			dst.UnboundidExternalAuthSaslMechanismHandlerResponse = nil
		} else {
			match++
		}
	} else {
		dst.UnboundidExternalAuthSaslMechanismHandlerResponse = nil
	}

	// try to unmarshal data into UnboundidMsChapV2SaslMechanismHandlerResponse
	err = newStrictDecoder(data).Decode(&dst.UnboundidMsChapV2SaslMechanismHandlerResponse)
	if err == nil {
		jsonUnboundidMsChapV2SaslMechanismHandlerResponse, _ := json.Marshal(dst.UnboundidMsChapV2SaslMechanismHandlerResponse)
		if string(jsonUnboundidMsChapV2SaslMechanismHandlerResponse) == "{}" { // empty struct
			dst.UnboundidMsChapV2SaslMechanismHandlerResponse = nil
		} else {
			match++
		}
	} else {
		dst.UnboundidMsChapV2SaslMechanismHandlerResponse = nil
	}

	// try to unmarshal data into UnboundidTotpSaslMechanismHandlerResponse
	err = newStrictDecoder(data).Decode(&dst.UnboundidTotpSaslMechanismHandlerResponse)
	if err == nil {
		jsonUnboundidTotpSaslMechanismHandlerResponse, _ := json.Marshal(dst.UnboundidTotpSaslMechanismHandlerResponse)
		if string(jsonUnboundidTotpSaslMechanismHandlerResponse) == "{}" { // empty struct
			dst.UnboundidTotpSaslMechanismHandlerResponse = nil
		} else {
			match++
		}
	} else {
		dst.UnboundidTotpSaslMechanismHandlerResponse = nil
	}

	// try to unmarshal data into UnboundidYubikeyOtpSaslMechanismHandlerResponse
	err = newStrictDecoder(data).Decode(&dst.UnboundidYubikeyOtpSaslMechanismHandlerResponse)
	if err == nil {
		jsonUnboundidYubikeyOtpSaslMechanismHandlerResponse, _ := json.Marshal(dst.UnboundidYubikeyOtpSaslMechanismHandlerResponse)
		if string(jsonUnboundidYubikeyOtpSaslMechanismHandlerResponse) == "{}" { // empty struct
			dst.UnboundidYubikeyOtpSaslMechanismHandlerResponse = nil
		} else {
			match++
		}
	} else {
		dst.UnboundidYubikeyOtpSaslMechanismHandlerResponse = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.AnonymousSaslMechanismHandlerResponse = nil
		dst.CramMd5SaslMechanismHandlerResponse = nil
		dst.DigestMd5SaslMechanismHandlerResponse = nil
		dst.ExternalSaslMechanismHandlerResponse = nil
		dst.GssapiSaslMechanismHandlerResponse = nil
		dst.OauthBearerSaslMechanismHandlerResponse = nil
		dst.PlainSaslMechanismHandlerResponse = nil
		dst.ThirdPartySaslMechanismHandlerResponse = nil
		dst.UnboundidCertificatePlusPasswordSaslMechanismHandlerResponse = nil
		dst.UnboundidDeliveredOtpSaslMechanismHandlerResponse = nil
		dst.UnboundidExternalAuthSaslMechanismHandlerResponse = nil
		dst.UnboundidMsChapV2SaslMechanismHandlerResponse = nil
		dst.UnboundidTotpSaslMechanismHandlerResponse = nil
		dst.UnboundidYubikeyOtpSaslMechanismHandlerResponse = nil

		return fmt.Errorf("data matches more than one schema in oneOf(GetSaslMechanismHandler200Response)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(GetSaslMechanismHandler200Response)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GetSaslMechanismHandler200Response) MarshalJSON() ([]byte, error) {
	if src.AnonymousSaslMechanismHandlerResponse != nil {
		return json.Marshal(&src.AnonymousSaslMechanismHandlerResponse)
	}

	if src.CramMd5SaslMechanismHandlerResponse != nil {
		return json.Marshal(&src.CramMd5SaslMechanismHandlerResponse)
	}

	if src.DigestMd5SaslMechanismHandlerResponse != nil {
		return json.Marshal(&src.DigestMd5SaslMechanismHandlerResponse)
	}

	if src.ExternalSaslMechanismHandlerResponse != nil {
		return json.Marshal(&src.ExternalSaslMechanismHandlerResponse)
	}

	if src.GssapiSaslMechanismHandlerResponse != nil {
		return json.Marshal(&src.GssapiSaslMechanismHandlerResponse)
	}

	if src.OauthBearerSaslMechanismHandlerResponse != nil {
		return json.Marshal(&src.OauthBearerSaslMechanismHandlerResponse)
	}

	if src.PlainSaslMechanismHandlerResponse != nil {
		return json.Marshal(&src.PlainSaslMechanismHandlerResponse)
	}

	if src.ThirdPartySaslMechanismHandlerResponse != nil {
		return json.Marshal(&src.ThirdPartySaslMechanismHandlerResponse)
	}

	if src.UnboundidCertificatePlusPasswordSaslMechanismHandlerResponse != nil {
		return json.Marshal(&src.UnboundidCertificatePlusPasswordSaslMechanismHandlerResponse)
	}

	if src.UnboundidDeliveredOtpSaslMechanismHandlerResponse != nil {
		return json.Marshal(&src.UnboundidDeliveredOtpSaslMechanismHandlerResponse)
	}

	if src.UnboundidExternalAuthSaslMechanismHandlerResponse != nil {
		return json.Marshal(&src.UnboundidExternalAuthSaslMechanismHandlerResponse)
	}

	if src.UnboundidMsChapV2SaslMechanismHandlerResponse != nil {
		return json.Marshal(&src.UnboundidMsChapV2SaslMechanismHandlerResponse)
	}

	if src.UnboundidTotpSaslMechanismHandlerResponse != nil {
		return json.Marshal(&src.UnboundidTotpSaslMechanismHandlerResponse)
	}

	if src.UnboundidYubikeyOtpSaslMechanismHandlerResponse != nil {
		return json.Marshal(&src.UnboundidYubikeyOtpSaslMechanismHandlerResponse)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GetSaslMechanismHandler200Response) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.AnonymousSaslMechanismHandlerResponse != nil {
		return obj.AnonymousSaslMechanismHandlerResponse
	}

	if obj.CramMd5SaslMechanismHandlerResponse != nil {
		return obj.CramMd5SaslMechanismHandlerResponse
	}

	if obj.DigestMd5SaslMechanismHandlerResponse != nil {
		return obj.DigestMd5SaslMechanismHandlerResponse
	}

	if obj.ExternalSaslMechanismHandlerResponse != nil {
		return obj.ExternalSaslMechanismHandlerResponse
	}

	if obj.GssapiSaslMechanismHandlerResponse != nil {
		return obj.GssapiSaslMechanismHandlerResponse
	}

	if obj.OauthBearerSaslMechanismHandlerResponse != nil {
		return obj.OauthBearerSaslMechanismHandlerResponse
	}

	if obj.PlainSaslMechanismHandlerResponse != nil {
		return obj.PlainSaslMechanismHandlerResponse
	}

	if obj.ThirdPartySaslMechanismHandlerResponse != nil {
		return obj.ThirdPartySaslMechanismHandlerResponse
	}

	if obj.UnboundidCertificatePlusPasswordSaslMechanismHandlerResponse != nil {
		return obj.UnboundidCertificatePlusPasswordSaslMechanismHandlerResponse
	}

	if obj.UnboundidDeliveredOtpSaslMechanismHandlerResponse != nil {
		return obj.UnboundidDeliveredOtpSaslMechanismHandlerResponse
	}

	if obj.UnboundidExternalAuthSaslMechanismHandlerResponse != nil {
		return obj.UnboundidExternalAuthSaslMechanismHandlerResponse
	}

	if obj.UnboundidMsChapV2SaslMechanismHandlerResponse != nil {
		return obj.UnboundidMsChapV2SaslMechanismHandlerResponse
	}

	if obj.UnboundidTotpSaslMechanismHandlerResponse != nil {
		return obj.UnboundidTotpSaslMechanismHandlerResponse
	}

	if obj.UnboundidYubikeyOtpSaslMechanismHandlerResponse != nil {
		return obj.UnboundidYubikeyOtpSaslMechanismHandlerResponse
	}

	// all schemas are nil
	return nil
}

type NullableGetSaslMechanismHandler200Response struct {
	value *GetSaslMechanismHandler200Response
	isSet bool
}

func (v NullableGetSaslMechanismHandler200Response) Get() *GetSaslMechanismHandler200Response {
	return v.value
}

func (v *NullableGetSaslMechanismHandler200Response) Set(val *GetSaslMechanismHandler200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSaslMechanismHandler200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSaslMechanismHandler200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSaslMechanismHandler200Response(val *GetSaslMechanismHandler200Response) *NullableGetSaslMechanismHandler200Response {
	return &NullableGetSaslMechanismHandler200Response{value: val, isSet: true}
}

func (v NullableGetSaslMechanismHandler200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSaslMechanismHandler200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


