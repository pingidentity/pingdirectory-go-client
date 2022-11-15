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

// AddAlertHandler200Response - struct for AddAlertHandler200Response
type AddAlertHandler200Response struct {
	ErrorLogAlertHandlerResponse *ErrorLogAlertHandlerResponse
	ExecAlertHandlerResponse *ExecAlertHandlerResponse
	GroovyScriptedAlertHandlerResponse *GroovyScriptedAlertHandlerResponse
	JmxAlertHandlerResponse *JmxAlertHandlerResponse
	SmtpAlertHandlerResponse *SmtpAlertHandlerResponse
	SnmpAlertHandlerResponse *SnmpAlertHandlerResponse
	SnmpSubAgentAlertHandlerResponse *SnmpSubAgentAlertHandlerResponse
	ThirdPartyAlertHandlerResponse *ThirdPartyAlertHandlerResponse
	TwilioAlertHandlerResponse *TwilioAlertHandlerResponse
}

// ErrorLogAlertHandlerResponseAsAddAlertHandler200Response is a convenience function that returns ErrorLogAlertHandlerResponse wrapped in AddAlertHandler200Response
func ErrorLogAlertHandlerResponseAsAddAlertHandler200Response(v *ErrorLogAlertHandlerResponse) AddAlertHandler200Response {
	return AddAlertHandler200Response{
		ErrorLogAlertHandlerResponse: v,
	}
}

// ExecAlertHandlerResponseAsAddAlertHandler200Response is a convenience function that returns ExecAlertHandlerResponse wrapped in AddAlertHandler200Response
func ExecAlertHandlerResponseAsAddAlertHandler200Response(v *ExecAlertHandlerResponse) AddAlertHandler200Response {
	return AddAlertHandler200Response{
		ExecAlertHandlerResponse: v,
	}
}

// GroovyScriptedAlertHandlerResponseAsAddAlertHandler200Response is a convenience function that returns GroovyScriptedAlertHandlerResponse wrapped in AddAlertHandler200Response
func GroovyScriptedAlertHandlerResponseAsAddAlertHandler200Response(v *GroovyScriptedAlertHandlerResponse) AddAlertHandler200Response {
	return AddAlertHandler200Response{
		GroovyScriptedAlertHandlerResponse: v,
	}
}

// JmxAlertHandlerResponseAsAddAlertHandler200Response is a convenience function that returns JmxAlertHandlerResponse wrapped in AddAlertHandler200Response
func JmxAlertHandlerResponseAsAddAlertHandler200Response(v *JmxAlertHandlerResponse) AddAlertHandler200Response {
	return AddAlertHandler200Response{
		JmxAlertHandlerResponse: v,
	}
}

// SmtpAlertHandlerResponseAsAddAlertHandler200Response is a convenience function that returns SmtpAlertHandlerResponse wrapped in AddAlertHandler200Response
func SmtpAlertHandlerResponseAsAddAlertHandler200Response(v *SmtpAlertHandlerResponse) AddAlertHandler200Response {
	return AddAlertHandler200Response{
		SmtpAlertHandlerResponse: v,
	}
}

// SnmpAlertHandlerResponseAsAddAlertHandler200Response is a convenience function that returns SnmpAlertHandlerResponse wrapped in AddAlertHandler200Response
func SnmpAlertHandlerResponseAsAddAlertHandler200Response(v *SnmpAlertHandlerResponse) AddAlertHandler200Response {
	return AddAlertHandler200Response{
		SnmpAlertHandlerResponse: v,
	}
}

// SnmpSubAgentAlertHandlerResponseAsAddAlertHandler200Response is a convenience function that returns SnmpSubAgentAlertHandlerResponse wrapped in AddAlertHandler200Response
func SnmpSubAgentAlertHandlerResponseAsAddAlertHandler200Response(v *SnmpSubAgentAlertHandlerResponse) AddAlertHandler200Response {
	return AddAlertHandler200Response{
		SnmpSubAgentAlertHandlerResponse: v,
	}
}

// ThirdPartyAlertHandlerResponseAsAddAlertHandler200Response is a convenience function that returns ThirdPartyAlertHandlerResponse wrapped in AddAlertHandler200Response
func ThirdPartyAlertHandlerResponseAsAddAlertHandler200Response(v *ThirdPartyAlertHandlerResponse) AddAlertHandler200Response {
	return AddAlertHandler200Response{
		ThirdPartyAlertHandlerResponse: v,
	}
}

// TwilioAlertHandlerResponseAsAddAlertHandler200Response is a convenience function that returns TwilioAlertHandlerResponse wrapped in AddAlertHandler200Response
func TwilioAlertHandlerResponseAsAddAlertHandler200Response(v *TwilioAlertHandlerResponse) AddAlertHandler200Response {
	return AddAlertHandler200Response{
		TwilioAlertHandlerResponse: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *AddAlertHandler200Response) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ErrorLogAlertHandlerResponse
	err = newStrictDecoder(data).Decode(&dst.ErrorLogAlertHandlerResponse)
	if err == nil {
		jsonErrorLogAlertHandlerResponse, _ := json.Marshal(dst.ErrorLogAlertHandlerResponse)
		if string(jsonErrorLogAlertHandlerResponse) == "{}" { // empty struct
			dst.ErrorLogAlertHandlerResponse = nil
		} else {
			match++
		}
	} else {
		dst.ErrorLogAlertHandlerResponse = nil
	}

	// try to unmarshal data into ExecAlertHandlerResponse
	err = newStrictDecoder(data).Decode(&dst.ExecAlertHandlerResponse)
	if err == nil {
		jsonExecAlertHandlerResponse, _ := json.Marshal(dst.ExecAlertHandlerResponse)
		if string(jsonExecAlertHandlerResponse) == "{}" { // empty struct
			dst.ExecAlertHandlerResponse = nil
		} else {
			match++
		}
	} else {
		dst.ExecAlertHandlerResponse = nil
	}

	// try to unmarshal data into GroovyScriptedAlertHandlerResponse
	err = newStrictDecoder(data).Decode(&dst.GroovyScriptedAlertHandlerResponse)
	if err == nil {
		jsonGroovyScriptedAlertHandlerResponse, _ := json.Marshal(dst.GroovyScriptedAlertHandlerResponse)
		if string(jsonGroovyScriptedAlertHandlerResponse) == "{}" { // empty struct
			dst.GroovyScriptedAlertHandlerResponse = nil
		} else {
			match++
		}
	} else {
		dst.GroovyScriptedAlertHandlerResponse = nil
	}

	// try to unmarshal data into JmxAlertHandlerResponse
	err = newStrictDecoder(data).Decode(&dst.JmxAlertHandlerResponse)
	if err == nil {
		jsonJmxAlertHandlerResponse, _ := json.Marshal(dst.JmxAlertHandlerResponse)
		if string(jsonJmxAlertHandlerResponse) == "{}" { // empty struct
			dst.JmxAlertHandlerResponse = nil
		} else {
			match++
		}
	} else {
		dst.JmxAlertHandlerResponse = nil
	}

	// try to unmarshal data into SmtpAlertHandlerResponse
	err = newStrictDecoder(data).Decode(&dst.SmtpAlertHandlerResponse)
	if err == nil {
		jsonSmtpAlertHandlerResponse, _ := json.Marshal(dst.SmtpAlertHandlerResponse)
		if string(jsonSmtpAlertHandlerResponse) == "{}" { // empty struct
			dst.SmtpAlertHandlerResponse = nil
		} else {
			match++
		}
	} else {
		dst.SmtpAlertHandlerResponse = nil
	}

	// try to unmarshal data into SnmpAlertHandlerResponse
	err = newStrictDecoder(data).Decode(&dst.SnmpAlertHandlerResponse)
	if err == nil {
		jsonSnmpAlertHandlerResponse, _ := json.Marshal(dst.SnmpAlertHandlerResponse)
		if string(jsonSnmpAlertHandlerResponse) == "{}" { // empty struct
			dst.SnmpAlertHandlerResponse = nil
		} else {
			match++
		}
	} else {
		dst.SnmpAlertHandlerResponse = nil
	}

	// try to unmarshal data into SnmpSubAgentAlertHandlerResponse
	err = newStrictDecoder(data).Decode(&dst.SnmpSubAgentAlertHandlerResponse)
	if err == nil {
		jsonSnmpSubAgentAlertHandlerResponse, _ := json.Marshal(dst.SnmpSubAgentAlertHandlerResponse)
		if string(jsonSnmpSubAgentAlertHandlerResponse) == "{}" { // empty struct
			dst.SnmpSubAgentAlertHandlerResponse = nil
		} else {
			match++
		}
	} else {
		dst.SnmpSubAgentAlertHandlerResponse = nil
	}

	// try to unmarshal data into ThirdPartyAlertHandlerResponse
	err = newStrictDecoder(data).Decode(&dst.ThirdPartyAlertHandlerResponse)
	if err == nil {
		jsonThirdPartyAlertHandlerResponse, _ := json.Marshal(dst.ThirdPartyAlertHandlerResponse)
		if string(jsonThirdPartyAlertHandlerResponse) == "{}" { // empty struct
			dst.ThirdPartyAlertHandlerResponse = nil
		} else {
			match++
		}
	} else {
		dst.ThirdPartyAlertHandlerResponse = nil
	}

	// try to unmarshal data into TwilioAlertHandlerResponse
	err = newStrictDecoder(data).Decode(&dst.TwilioAlertHandlerResponse)
	if err == nil {
		jsonTwilioAlertHandlerResponse, _ := json.Marshal(dst.TwilioAlertHandlerResponse)
		if string(jsonTwilioAlertHandlerResponse) == "{}" { // empty struct
			dst.TwilioAlertHandlerResponse = nil
		} else {
			match++
		}
	} else {
		dst.TwilioAlertHandlerResponse = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ErrorLogAlertHandlerResponse = nil
		dst.ExecAlertHandlerResponse = nil
		dst.GroovyScriptedAlertHandlerResponse = nil
		dst.JmxAlertHandlerResponse = nil
		dst.SmtpAlertHandlerResponse = nil
		dst.SnmpAlertHandlerResponse = nil
		dst.SnmpSubAgentAlertHandlerResponse = nil
		dst.ThirdPartyAlertHandlerResponse = nil
		dst.TwilioAlertHandlerResponse = nil

		return fmt.Errorf("data matches more than one schema in oneOf(AddAlertHandler200Response)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(AddAlertHandler200Response)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src AddAlertHandler200Response) MarshalJSON() ([]byte, error) {
	if src.ErrorLogAlertHandlerResponse != nil {
		return json.Marshal(&src.ErrorLogAlertHandlerResponse)
	}

	if src.ExecAlertHandlerResponse != nil {
		return json.Marshal(&src.ExecAlertHandlerResponse)
	}

	if src.GroovyScriptedAlertHandlerResponse != nil {
		return json.Marshal(&src.GroovyScriptedAlertHandlerResponse)
	}

	if src.JmxAlertHandlerResponse != nil {
		return json.Marshal(&src.JmxAlertHandlerResponse)
	}

	if src.SmtpAlertHandlerResponse != nil {
		return json.Marshal(&src.SmtpAlertHandlerResponse)
	}

	if src.SnmpAlertHandlerResponse != nil {
		return json.Marshal(&src.SnmpAlertHandlerResponse)
	}

	if src.SnmpSubAgentAlertHandlerResponse != nil {
		return json.Marshal(&src.SnmpSubAgentAlertHandlerResponse)
	}

	if src.ThirdPartyAlertHandlerResponse != nil {
		return json.Marshal(&src.ThirdPartyAlertHandlerResponse)
	}

	if src.TwilioAlertHandlerResponse != nil {
		return json.Marshal(&src.TwilioAlertHandlerResponse)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *AddAlertHandler200Response) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.ErrorLogAlertHandlerResponse != nil {
		return obj.ErrorLogAlertHandlerResponse
	}

	if obj.ExecAlertHandlerResponse != nil {
		return obj.ExecAlertHandlerResponse
	}

	if obj.GroovyScriptedAlertHandlerResponse != nil {
		return obj.GroovyScriptedAlertHandlerResponse
	}

	if obj.JmxAlertHandlerResponse != nil {
		return obj.JmxAlertHandlerResponse
	}

	if obj.SmtpAlertHandlerResponse != nil {
		return obj.SmtpAlertHandlerResponse
	}

	if obj.SnmpAlertHandlerResponse != nil {
		return obj.SnmpAlertHandlerResponse
	}

	if obj.SnmpSubAgentAlertHandlerResponse != nil {
		return obj.SnmpSubAgentAlertHandlerResponse
	}

	if obj.ThirdPartyAlertHandlerResponse != nil {
		return obj.ThirdPartyAlertHandlerResponse
	}

	if obj.TwilioAlertHandlerResponse != nil {
		return obj.TwilioAlertHandlerResponse
	}

	// all schemas are nil
	return nil
}

type NullableAddAlertHandler200Response struct {
	value *AddAlertHandler200Response
	isSet bool
}

func (v NullableAddAlertHandler200Response) Get() *AddAlertHandler200Response {
	return v.value
}

func (v *NullableAddAlertHandler200Response) Set(val *AddAlertHandler200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableAddAlertHandler200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableAddAlertHandler200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddAlertHandler200Response(val *AddAlertHandler200Response) *NullableAddAlertHandler200Response {
	return &NullableAddAlertHandler200Response{value: val, isSet: true}
}

func (v NullableAddAlertHandler200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddAlertHandler200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


