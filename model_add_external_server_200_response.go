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

// AddExternalServer200Response - struct for AddExternalServer200Response
type AddExternalServer200Response struct {
	ActiveDirectoryExternalServerResponse         *ActiveDirectoryExternalServerResponse
	AmazonAwsExternalServerResponse               *AmazonAwsExternalServerResponse
	ConjurExternalServerResponse                  *ConjurExternalServerResponse
	HttpExternalServerResponse                    *HttpExternalServerResponse
	JdbcExternalServerResponse                    *JdbcExternalServerResponse
	LdapExternalServerResponse                    *LdapExternalServerResponse
	NokiaDsExternalServerResponse                 *NokiaDsExternalServerResponse
	NokiaProxyServerExternalServerResponse        *NokiaProxyServerExternalServerResponse
	OpendjExternalServerResponse                  *OpendjExternalServerResponse
	OracleUnifiedDirectoryExternalServerResponse  *OracleUnifiedDirectoryExternalServerResponse
	PingIdentityDsExternalServerResponse          *PingIdentityDsExternalServerResponse
	PingIdentityProxyServerExternalServerResponse *PingIdentityProxyServerExternalServerResponse
	PingOneHttpExternalServerResponse             *PingOneHttpExternalServerResponse
	SmtpExternalServerResponse                    *SmtpExternalServerResponse
	SyslogExternalServerResponse                  *SyslogExternalServerResponse
	VaultExternalServerResponse                   *VaultExternalServerResponse
}

// ActiveDirectoryExternalServerResponseAsAddExternalServer200Response is a convenience function that returns ActiveDirectoryExternalServerResponse wrapped in AddExternalServer200Response
func ActiveDirectoryExternalServerResponseAsAddExternalServer200Response(v *ActiveDirectoryExternalServerResponse) AddExternalServer200Response {
	return AddExternalServer200Response{
		ActiveDirectoryExternalServerResponse: v,
	}
}

// AmazonAwsExternalServerResponseAsAddExternalServer200Response is a convenience function that returns AmazonAwsExternalServerResponse wrapped in AddExternalServer200Response
func AmazonAwsExternalServerResponseAsAddExternalServer200Response(v *AmazonAwsExternalServerResponse) AddExternalServer200Response {
	return AddExternalServer200Response{
		AmazonAwsExternalServerResponse: v,
	}
}

// ConjurExternalServerResponseAsAddExternalServer200Response is a convenience function that returns ConjurExternalServerResponse wrapped in AddExternalServer200Response
func ConjurExternalServerResponseAsAddExternalServer200Response(v *ConjurExternalServerResponse) AddExternalServer200Response {
	return AddExternalServer200Response{
		ConjurExternalServerResponse: v,
	}
}

// HttpExternalServerResponseAsAddExternalServer200Response is a convenience function that returns HttpExternalServerResponse wrapped in AddExternalServer200Response
func HttpExternalServerResponseAsAddExternalServer200Response(v *HttpExternalServerResponse) AddExternalServer200Response {
	return AddExternalServer200Response{
		HttpExternalServerResponse: v,
	}
}

// JdbcExternalServerResponseAsAddExternalServer200Response is a convenience function that returns JdbcExternalServerResponse wrapped in AddExternalServer200Response
func JdbcExternalServerResponseAsAddExternalServer200Response(v *JdbcExternalServerResponse) AddExternalServer200Response {
	return AddExternalServer200Response{
		JdbcExternalServerResponse: v,
	}
}

// LdapExternalServerResponseAsAddExternalServer200Response is a convenience function that returns LdapExternalServerResponse wrapped in AddExternalServer200Response
func LdapExternalServerResponseAsAddExternalServer200Response(v *LdapExternalServerResponse) AddExternalServer200Response {
	return AddExternalServer200Response{
		LdapExternalServerResponse: v,
	}
}

// NokiaDsExternalServerResponseAsAddExternalServer200Response is a convenience function that returns NokiaDsExternalServerResponse wrapped in AddExternalServer200Response
func NokiaDsExternalServerResponseAsAddExternalServer200Response(v *NokiaDsExternalServerResponse) AddExternalServer200Response {
	return AddExternalServer200Response{
		NokiaDsExternalServerResponse: v,
	}
}

// NokiaProxyServerExternalServerResponseAsAddExternalServer200Response is a convenience function that returns NokiaProxyServerExternalServerResponse wrapped in AddExternalServer200Response
func NokiaProxyServerExternalServerResponseAsAddExternalServer200Response(v *NokiaProxyServerExternalServerResponse) AddExternalServer200Response {
	return AddExternalServer200Response{
		NokiaProxyServerExternalServerResponse: v,
	}
}

// OpendjExternalServerResponseAsAddExternalServer200Response is a convenience function that returns OpendjExternalServerResponse wrapped in AddExternalServer200Response
func OpendjExternalServerResponseAsAddExternalServer200Response(v *OpendjExternalServerResponse) AddExternalServer200Response {
	return AddExternalServer200Response{
		OpendjExternalServerResponse: v,
	}
}

// OracleUnifiedDirectoryExternalServerResponseAsAddExternalServer200Response is a convenience function that returns OracleUnifiedDirectoryExternalServerResponse wrapped in AddExternalServer200Response
func OracleUnifiedDirectoryExternalServerResponseAsAddExternalServer200Response(v *OracleUnifiedDirectoryExternalServerResponse) AddExternalServer200Response {
	return AddExternalServer200Response{
		OracleUnifiedDirectoryExternalServerResponse: v,
	}
}

// PingIdentityDsExternalServerResponseAsAddExternalServer200Response is a convenience function that returns PingIdentityDsExternalServerResponse wrapped in AddExternalServer200Response
func PingIdentityDsExternalServerResponseAsAddExternalServer200Response(v *PingIdentityDsExternalServerResponse) AddExternalServer200Response {
	return AddExternalServer200Response{
		PingIdentityDsExternalServerResponse: v,
	}
}

// PingIdentityProxyServerExternalServerResponseAsAddExternalServer200Response is a convenience function that returns PingIdentityProxyServerExternalServerResponse wrapped in AddExternalServer200Response
func PingIdentityProxyServerExternalServerResponseAsAddExternalServer200Response(v *PingIdentityProxyServerExternalServerResponse) AddExternalServer200Response {
	return AddExternalServer200Response{
		PingIdentityProxyServerExternalServerResponse: v,
	}
}

// PingOneHttpExternalServerResponseAsAddExternalServer200Response is a convenience function that returns PingOneHttpExternalServerResponse wrapped in AddExternalServer200Response
func PingOneHttpExternalServerResponseAsAddExternalServer200Response(v *PingOneHttpExternalServerResponse) AddExternalServer200Response {
	return AddExternalServer200Response{
		PingOneHttpExternalServerResponse: v,
	}
}

// SmtpExternalServerResponseAsAddExternalServer200Response is a convenience function that returns SmtpExternalServerResponse wrapped in AddExternalServer200Response
func SmtpExternalServerResponseAsAddExternalServer200Response(v *SmtpExternalServerResponse) AddExternalServer200Response {
	return AddExternalServer200Response{
		SmtpExternalServerResponse: v,
	}
}

// SyslogExternalServerResponseAsAddExternalServer200Response is a convenience function that returns SyslogExternalServerResponse wrapped in AddExternalServer200Response
func SyslogExternalServerResponseAsAddExternalServer200Response(v *SyslogExternalServerResponse) AddExternalServer200Response {
	return AddExternalServer200Response{
		SyslogExternalServerResponse: v,
	}
}

// VaultExternalServerResponseAsAddExternalServer200Response is a convenience function that returns VaultExternalServerResponse wrapped in AddExternalServer200Response
func VaultExternalServerResponseAsAddExternalServer200Response(v *VaultExternalServerResponse) AddExternalServer200Response {
	return AddExternalServer200Response{
		VaultExternalServerResponse: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *AddExternalServer200Response) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ActiveDirectoryExternalServerResponse
	err = newStrictDecoder(data).Decode(&dst.ActiveDirectoryExternalServerResponse)
	if err == nil {
		jsonActiveDirectoryExternalServerResponse, _ := json.Marshal(dst.ActiveDirectoryExternalServerResponse)
		if string(jsonActiveDirectoryExternalServerResponse) == "{}" { // empty struct
			dst.ActiveDirectoryExternalServerResponse = nil
		} else {
			match++
		}
	} else {
		dst.ActiveDirectoryExternalServerResponse = nil
	}

	// try to unmarshal data into AmazonAwsExternalServerResponse
	err = newStrictDecoder(data).Decode(&dst.AmazonAwsExternalServerResponse)
	if err == nil {
		jsonAmazonAwsExternalServerResponse, _ := json.Marshal(dst.AmazonAwsExternalServerResponse)
		if string(jsonAmazonAwsExternalServerResponse) == "{}" { // empty struct
			dst.AmazonAwsExternalServerResponse = nil
		} else {
			match++
		}
	} else {
		dst.AmazonAwsExternalServerResponse = nil
	}

	// try to unmarshal data into ConjurExternalServerResponse
	err = newStrictDecoder(data).Decode(&dst.ConjurExternalServerResponse)
	if err == nil {
		jsonConjurExternalServerResponse, _ := json.Marshal(dst.ConjurExternalServerResponse)
		if string(jsonConjurExternalServerResponse) == "{}" { // empty struct
			dst.ConjurExternalServerResponse = nil
		} else {
			match++
		}
	} else {
		dst.ConjurExternalServerResponse = nil
	}

	// try to unmarshal data into HttpExternalServerResponse
	err = newStrictDecoder(data).Decode(&dst.HttpExternalServerResponse)
	if err == nil {
		jsonHttpExternalServerResponse, _ := json.Marshal(dst.HttpExternalServerResponse)
		if string(jsonHttpExternalServerResponse) == "{}" { // empty struct
			dst.HttpExternalServerResponse = nil
		} else {
			match++
		}
	} else {
		dst.HttpExternalServerResponse = nil
	}

	// try to unmarshal data into JdbcExternalServerResponse
	err = newStrictDecoder(data).Decode(&dst.JdbcExternalServerResponse)
	if err == nil {
		jsonJdbcExternalServerResponse, _ := json.Marshal(dst.JdbcExternalServerResponse)
		if string(jsonJdbcExternalServerResponse) == "{}" { // empty struct
			dst.JdbcExternalServerResponse = nil
		} else {
			match++
		}
	} else {
		dst.JdbcExternalServerResponse = nil
	}

	// try to unmarshal data into LdapExternalServerResponse
	err = newStrictDecoder(data).Decode(&dst.LdapExternalServerResponse)
	if err == nil {
		jsonLdapExternalServerResponse, _ := json.Marshal(dst.LdapExternalServerResponse)
		if string(jsonLdapExternalServerResponse) == "{}" { // empty struct
			dst.LdapExternalServerResponse = nil
		} else {
			match++
		}
	} else {
		dst.LdapExternalServerResponse = nil
	}

	// try to unmarshal data into NokiaDsExternalServerResponse
	err = newStrictDecoder(data).Decode(&dst.NokiaDsExternalServerResponse)
	if err == nil {
		jsonNokiaDsExternalServerResponse, _ := json.Marshal(dst.NokiaDsExternalServerResponse)
		if string(jsonNokiaDsExternalServerResponse) == "{}" { // empty struct
			dst.NokiaDsExternalServerResponse = nil
		} else {
			match++
		}
	} else {
		dst.NokiaDsExternalServerResponse = nil
	}

	// try to unmarshal data into NokiaProxyServerExternalServerResponse
	err = newStrictDecoder(data).Decode(&dst.NokiaProxyServerExternalServerResponse)
	if err == nil {
		jsonNokiaProxyServerExternalServerResponse, _ := json.Marshal(dst.NokiaProxyServerExternalServerResponse)
		if string(jsonNokiaProxyServerExternalServerResponse) == "{}" { // empty struct
			dst.NokiaProxyServerExternalServerResponse = nil
		} else {
			match++
		}
	} else {
		dst.NokiaProxyServerExternalServerResponse = nil
	}

	// try to unmarshal data into OpendjExternalServerResponse
	err = newStrictDecoder(data).Decode(&dst.OpendjExternalServerResponse)
	if err == nil {
		jsonOpendjExternalServerResponse, _ := json.Marshal(dst.OpendjExternalServerResponse)
		if string(jsonOpendjExternalServerResponse) == "{}" { // empty struct
			dst.OpendjExternalServerResponse = nil
		} else {
			match++
		}
	} else {
		dst.OpendjExternalServerResponse = nil
	}

	// try to unmarshal data into OracleUnifiedDirectoryExternalServerResponse
	err = newStrictDecoder(data).Decode(&dst.OracleUnifiedDirectoryExternalServerResponse)
	if err == nil {
		jsonOracleUnifiedDirectoryExternalServerResponse, _ := json.Marshal(dst.OracleUnifiedDirectoryExternalServerResponse)
		if string(jsonOracleUnifiedDirectoryExternalServerResponse) == "{}" { // empty struct
			dst.OracleUnifiedDirectoryExternalServerResponse = nil
		} else {
			match++
		}
	} else {
		dst.OracleUnifiedDirectoryExternalServerResponse = nil
	}

	// try to unmarshal data into PingIdentityDsExternalServerResponse
	err = newStrictDecoder(data).Decode(&dst.PingIdentityDsExternalServerResponse)
	if err == nil {
		jsonPingIdentityDsExternalServerResponse, _ := json.Marshal(dst.PingIdentityDsExternalServerResponse)
		if string(jsonPingIdentityDsExternalServerResponse) == "{}" { // empty struct
			dst.PingIdentityDsExternalServerResponse = nil
		} else {
			match++
		}
	} else {
		dst.PingIdentityDsExternalServerResponse = nil
	}

	// try to unmarshal data into PingIdentityProxyServerExternalServerResponse
	err = newStrictDecoder(data).Decode(&dst.PingIdentityProxyServerExternalServerResponse)
	if err == nil {
		jsonPingIdentityProxyServerExternalServerResponse, _ := json.Marshal(dst.PingIdentityProxyServerExternalServerResponse)
		if string(jsonPingIdentityProxyServerExternalServerResponse) == "{}" { // empty struct
			dst.PingIdentityProxyServerExternalServerResponse = nil
		} else {
			match++
		}
	} else {
		dst.PingIdentityProxyServerExternalServerResponse = nil
	}

	// try to unmarshal data into PingOneHttpExternalServerResponse
	err = newStrictDecoder(data).Decode(&dst.PingOneHttpExternalServerResponse)
	if err == nil {
		jsonPingOneHttpExternalServerResponse, _ := json.Marshal(dst.PingOneHttpExternalServerResponse)
		if string(jsonPingOneHttpExternalServerResponse) == "{}" { // empty struct
			dst.PingOneHttpExternalServerResponse = nil
		} else {
			match++
		}
	} else {
		dst.PingOneHttpExternalServerResponse = nil
	}

	// try to unmarshal data into SmtpExternalServerResponse
	err = newStrictDecoder(data).Decode(&dst.SmtpExternalServerResponse)
	if err == nil {
		jsonSmtpExternalServerResponse, _ := json.Marshal(dst.SmtpExternalServerResponse)
		if string(jsonSmtpExternalServerResponse) == "{}" { // empty struct
			dst.SmtpExternalServerResponse = nil
		} else {
			match++
		}
	} else {
		dst.SmtpExternalServerResponse = nil
	}

	// try to unmarshal data into SyslogExternalServerResponse
	err = newStrictDecoder(data).Decode(&dst.SyslogExternalServerResponse)
	if err == nil {
		jsonSyslogExternalServerResponse, _ := json.Marshal(dst.SyslogExternalServerResponse)
		if string(jsonSyslogExternalServerResponse) == "{}" { // empty struct
			dst.SyslogExternalServerResponse = nil
		} else {
			match++
		}
	} else {
		dst.SyslogExternalServerResponse = nil
	}

	// try to unmarshal data into VaultExternalServerResponse
	err = newStrictDecoder(data).Decode(&dst.VaultExternalServerResponse)
	if err == nil {
		jsonVaultExternalServerResponse, _ := json.Marshal(dst.VaultExternalServerResponse)
		if string(jsonVaultExternalServerResponse) == "{}" { // empty struct
			dst.VaultExternalServerResponse = nil
		} else {
			match++
		}
	} else {
		dst.VaultExternalServerResponse = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ActiveDirectoryExternalServerResponse = nil
		dst.AmazonAwsExternalServerResponse = nil
		dst.ConjurExternalServerResponse = nil
		dst.HttpExternalServerResponse = nil
		dst.JdbcExternalServerResponse = nil
		dst.LdapExternalServerResponse = nil
		dst.NokiaDsExternalServerResponse = nil
		dst.NokiaProxyServerExternalServerResponse = nil
		dst.OpendjExternalServerResponse = nil
		dst.OracleUnifiedDirectoryExternalServerResponse = nil
		dst.PingIdentityDsExternalServerResponse = nil
		dst.PingIdentityProxyServerExternalServerResponse = nil
		dst.PingOneHttpExternalServerResponse = nil
		dst.SmtpExternalServerResponse = nil
		dst.SyslogExternalServerResponse = nil
		dst.VaultExternalServerResponse = nil

		return fmt.Errorf("data matches more than one schema in oneOf(AddExternalServer200Response)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(AddExternalServer200Response)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src AddExternalServer200Response) MarshalJSON() ([]byte, error) {
	if src.ActiveDirectoryExternalServerResponse != nil {
		return json.Marshal(&src.ActiveDirectoryExternalServerResponse)
	}

	if src.AmazonAwsExternalServerResponse != nil {
		return json.Marshal(&src.AmazonAwsExternalServerResponse)
	}

	if src.ConjurExternalServerResponse != nil {
		return json.Marshal(&src.ConjurExternalServerResponse)
	}

	if src.HttpExternalServerResponse != nil {
		return json.Marshal(&src.HttpExternalServerResponse)
	}

	if src.JdbcExternalServerResponse != nil {
		return json.Marshal(&src.JdbcExternalServerResponse)
	}

	if src.LdapExternalServerResponse != nil {
		return json.Marshal(&src.LdapExternalServerResponse)
	}

	if src.NokiaDsExternalServerResponse != nil {
		return json.Marshal(&src.NokiaDsExternalServerResponse)
	}

	if src.NokiaProxyServerExternalServerResponse != nil {
		return json.Marshal(&src.NokiaProxyServerExternalServerResponse)
	}

	if src.OpendjExternalServerResponse != nil {
		return json.Marshal(&src.OpendjExternalServerResponse)
	}

	if src.OracleUnifiedDirectoryExternalServerResponse != nil {
		return json.Marshal(&src.OracleUnifiedDirectoryExternalServerResponse)
	}

	if src.PingIdentityDsExternalServerResponse != nil {
		return json.Marshal(&src.PingIdentityDsExternalServerResponse)
	}

	if src.PingIdentityProxyServerExternalServerResponse != nil {
		return json.Marshal(&src.PingIdentityProxyServerExternalServerResponse)
	}

	if src.PingOneHttpExternalServerResponse != nil {
		return json.Marshal(&src.PingOneHttpExternalServerResponse)
	}

	if src.SmtpExternalServerResponse != nil {
		return json.Marshal(&src.SmtpExternalServerResponse)
	}

	if src.SyslogExternalServerResponse != nil {
		return json.Marshal(&src.SyslogExternalServerResponse)
	}

	if src.VaultExternalServerResponse != nil {
		return json.Marshal(&src.VaultExternalServerResponse)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *AddExternalServer200Response) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.ActiveDirectoryExternalServerResponse != nil {
		return obj.ActiveDirectoryExternalServerResponse
	}

	if obj.AmazonAwsExternalServerResponse != nil {
		return obj.AmazonAwsExternalServerResponse
	}

	if obj.ConjurExternalServerResponse != nil {
		return obj.ConjurExternalServerResponse
	}

	if obj.HttpExternalServerResponse != nil {
		return obj.HttpExternalServerResponse
	}

	if obj.JdbcExternalServerResponse != nil {
		return obj.JdbcExternalServerResponse
	}

	if obj.LdapExternalServerResponse != nil {
		return obj.LdapExternalServerResponse
	}

	if obj.NokiaDsExternalServerResponse != nil {
		return obj.NokiaDsExternalServerResponse
	}

	if obj.NokiaProxyServerExternalServerResponse != nil {
		return obj.NokiaProxyServerExternalServerResponse
	}

	if obj.OpendjExternalServerResponse != nil {
		return obj.OpendjExternalServerResponse
	}

	if obj.OracleUnifiedDirectoryExternalServerResponse != nil {
		return obj.OracleUnifiedDirectoryExternalServerResponse
	}

	if obj.PingIdentityDsExternalServerResponse != nil {
		return obj.PingIdentityDsExternalServerResponse
	}

	if obj.PingIdentityProxyServerExternalServerResponse != nil {
		return obj.PingIdentityProxyServerExternalServerResponse
	}

	if obj.PingOneHttpExternalServerResponse != nil {
		return obj.PingOneHttpExternalServerResponse
	}

	if obj.SmtpExternalServerResponse != nil {
		return obj.SmtpExternalServerResponse
	}

	if obj.SyslogExternalServerResponse != nil {
		return obj.SyslogExternalServerResponse
	}

	if obj.VaultExternalServerResponse != nil {
		return obj.VaultExternalServerResponse
	}

	// all schemas are nil
	return nil
}

type NullableAddExternalServer200Response struct {
	value *AddExternalServer200Response
	isSet bool
}

func (v NullableAddExternalServer200Response) Get() *AddExternalServer200Response {
	return v.value
}

func (v *NullableAddExternalServer200Response) Set(val *AddExternalServer200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableAddExternalServer200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableAddExternalServer200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddExternalServer200Response(val *AddExternalServer200Response) *NullableAddExternalServer200Response {
	return &NullableAddExternalServer200Response{value: val, isSet: true}
}

func (v NullableAddExternalServer200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddExternalServer200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
