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

// AddCertificateMapper200Response - struct for AddCertificateMapper200Response
type AddCertificateMapper200Response struct {
	FingerprintCertificateMapperResponse *FingerprintCertificateMapperResponse
	GroovyScriptedCertificateMapperResponse *GroovyScriptedCertificateMapperResponse
	SubjectAttributeToUserAttributeCertificateMapperResponse *SubjectAttributeToUserAttributeCertificateMapperResponse
	SubjectDnToUserAttributeCertificateMapperResponse *SubjectDnToUserAttributeCertificateMapperResponse
	SubjectEqualsDnCertificateMapperResponse *SubjectEqualsDnCertificateMapperResponse
	ThirdPartyCertificateMapperResponse *ThirdPartyCertificateMapperResponse
}

// FingerprintCertificateMapperResponseAsAddCertificateMapper200Response is a convenience function that returns FingerprintCertificateMapperResponse wrapped in AddCertificateMapper200Response
func FingerprintCertificateMapperResponseAsAddCertificateMapper200Response(v *FingerprintCertificateMapperResponse) AddCertificateMapper200Response {
	return AddCertificateMapper200Response{
		FingerprintCertificateMapperResponse: v,
	}
}

// GroovyScriptedCertificateMapperResponseAsAddCertificateMapper200Response is a convenience function that returns GroovyScriptedCertificateMapperResponse wrapped in AddCertificateMapper200Response
func GroovyScriptedCertificateMapperResponseAsAddCertificateMapper200Response(v *GroovyScriptedCertificateMapperResponse) AddCertificateMapper200Response {
	return AddCertificateMapper200Response{
		GroovyScriptedCertificateMapperResponse: v,
	}
}

// SubjectAttributeToUserAttributeCertificateMapperResponseAsAddCertificateMapper200Response is a convenience function that returns SubjectAttributeToUserAttributeCertificateMapperResponse wrapped in AddCertificateMapper200Response
func SubjectAttributeToUserAttributeCertificateMapperResponseAsAddCertificateMapper200Response(v *SubjectAttributeToUserAttributeCertificateMapperResponse) AddCertificateMapper200Response {
	return AddCertificateMapper200Response{
		SubjectAttributeToUserAttributeCertificateMapperResponse: v,
	}
}

// SubjectDnToUserAttributeCertificateMapperResponseAsAddCertificateMapper200Response is a convenience function that returns SubjectDnToUserAttributeCertificateMapperResponse wrapped in AddCertificateMapper200Response
func SubjectDnToUserAttributeCertificateMapperResponseAsAddCertificateMapper200Response(v *SubjectDnToUserAttributeCertificateMapperResponse) AddCertificateMapper200Response {
	return AddCertificateMapper200Response{
		SubjectDnToUserAttributeCertificateMapperResponse: v,
	}
}

// SubjectEqualsDnCertificateMapperResponseAsAddCertificateMapper200Response is a convenience function that returns SubjectEqualsDnCertificateMapperResponse wrapped in AddCertificateMapper200Response
func SubjectEqualsDnCertificateMapperResponseAsAddCertificateMapper200Response(v *SubjectEqualsDnCertificateMapperResponse) AddCertificateMapper200Response {
	return AddCertificateMapper200Response{
		SubjectEqualsDnCertificateMapperResponse: v,
	}
}

// ThirdPartyCertificateMapperResponseAsAddCertificateMapper200Response is a convenience function that returns ThirdPartyCertificateMapperResponse wrapped in AddCertificateMapper200Response
func ThirdPartyCertificateMapperResponseAsAddCertificateMapper200Response(v *ThirdPartyCertificateMapperResponse) AddCertificateMapper200Response {
	return AddCertificateMapper200Response{
		ThirdPartyCertificateMapperResponse: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *AddCertificateMapper200Response) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into FingerprintCertificateMapperResponse
	err = newStrictDecoder(data).Decode(&dst.FingerprintCertificateMapperResponse)
	if err == nil {
		jsonFingerprintCertificateMapperResponse, _ := json.Marshal(dst.FingerprintCertificateMapperResponse)
		if string(jsonFingerprintCertificateMapperResponse) == "{}" { // empty struct
			dst.FingerprintCertificateMapperResponse = nil
		} else {
			match++
		}
	} else {
		dst.FingerprintCertificateMapperResponse = nil
	}

	// try to unmarshal data into GroovyScriptedCertificateMapperResponse
	err = newStrictDecoder(data).Decode(&dst.GroovyScriptedCertificateMapperResponse)
	if err == nil {
		jsonGroovyScriptedCertificateMapperResponse, _ := json.Marshal(dst.GroovyScriptedCertificateMapperResponse)
		if string(jsonGroovyScriptedCertificateMapperResponse) == "{}" { // empty struct
			dst.GroovyScriptedCertificateMapperResponse = nil
		} else {
			match++
		}
	} else {
		dst.GroovyScriptedCertificateMapperResponse = nil
	}

	// try to unmarshal data into SubjectAttributeToUserAttributeCertificateMapperResponse
	err = newStrictDecoder(data).Decode(&dst.SubjectAttributeToUserAttributeCertificateMapperResponse)
	if err == nil {
		jsonSubjectAttributeToUserAttributeCertificateMapperResponse, _ := json.Marshal(dst.SubjectAttributeToUserAttributeCertificateMapperResponse)
		if string(jsonSubjectAttributeToUserAttributeCertificateMapperResponse) == "{}" { // empty struct
			dst.SubjectAttributeToUserAttributeCertificateMapperResponse = nil
		} else {
			match++
		}
	} else {
		dst.SubjectAttributeToUserAttributeCertificateMapperResponse = nil
	}

	// try to unmarshal data into SubjectDnToUserAttributeCertificateMapperResponse
	err = newStrictDecoder(data).Decode(&dst.SubjectDnToUserAttributeCertificateMapperResponse)
	if err == nil {
		jsonSubjectDnToUserAttributeCertificateMapperResponse, _ := json.Marshal(dst.SubjectDnToUserAttributeCertificateMapperResponse)
		if string(jsonSubjectDnToUserAttributeCertificateMapperResponse) == "{}" { // empty struct
			dst.SubjectDnToUserAttributeCertificateMapperResponse = nil
		} else {
			match++
		}
	} else {
		dst.SubjectDnToUserAttributeCertificateMapperResponse = nil
	}

	// try to unmarshal data into SubjectEqualsDnCertificateMapperResponse
	err = newStrictDecoder(data).Decode(&dst.SubjectEqualsDnCertificateMapperResponse)
	if err == nil {
		jsonSubjectEqualsDnCertificateMapperResponse, _ := json.Marshal(dst.SubjectEqualsDnCertificateMapperResponse)
		if string(jsonSubjectEqualsDnCertificateMapperResponse) == "{}" { // empty struct
			dst.SubjectEqualsDnCertificateMapperResponse = nil
		} else {
			match++
		}
	} else {
		dst.SubjectEqualsDnCertificateMapperResponse = nil
	}

	// try to unmarshal data into ThirdPartyCertificateMapperResponse
	err = newStrictDecoder(data).Decode(&dst.ThirdPartyCertificateMapperResponse)
	if err == nil {
		jsonThirdPartyCertificateMapperResponse, _ := json.Marshal(dst.ThirdPartyCertificateMapperResponse)
		if string(jsonThirdPartyCertificateMapperResponse) == "{}" { // empty struct
			dst.ThirdPartyCertificateMapperResponse = nil
		} else {
			match++
		}
	} else {
		dst.ThirdPartyCertificateMapperResponse = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.FingerprintCertificateMapperResponse = nil
		dst.GroovyScriptedCertificateMapperResponse = nil
		dst.SubjectAttributeToUserAttributeCertificateMapperResponse = nil
		dst.SubjectDnToUserAttributeCertificateMapperResponse = nil
		dst.SubjectEqualsDnCertificateMapperResponse = nil
		dst.ThirdPartyCertificateMapperResponse = nil

		return fmt.Errorf("data matches more than one schema in oneOf(AddCertificateMapper200Response)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(AddCertificateMapper200Response)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src AddCertificateMapper200Response) MarshalJSON() ([]byte, error) {
	if src.FingerprintCertificateMapperResponse != nil {
		return json.Marshal(&src.FingerprintCertificateMapperResponse)
	}

	if src.GroovyScriptedCertificateMapperResponse != nil {
		return json.Marshal(&src.GroovyScriptedCertificateMapperResponse)
	}

	if src.SubjectAttributeToUserAttributeCertificateMapperResponse != nil {
		return json.Marshal(&src.SubjectAttributeToUserAttributeCertificateMapperResponse)
	}

	if src.SubjectDnToUserAttributeCertificateMapperResponse != nil {
		return json.Marshal(&src.SubjectDnToUserAttributeCertificateMapperResponse)
	}

	if src.SubjectEqualsDnCertificateMapperResponse != nil {
		return json.Marshal(&src.SubjectEqualsDnCertificateMapperResponse)
	}

	if src.ThirdPartyCertificateMapperResponse != nil {
		return json.Marshal(&src.ThirdPartyCertificateMapperResponse)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *AddCertificateMapper200Response) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.FingerprintCertificateMapperResponse != nil {
		return obj.FingerprintCertificateMapperResponse
	}

	if obj.GroovyScriptedCertificateMapperResponse != nil {
		return obj.GroovyScriptedCertificateMapperResponse
	}

	if obj.SubjectAttributeToUserAttributeCertificateMapperResponse != nil {
		return obj.SubjectAttributeToUserAttributeCertificateMapperResponse
	}

	if obj.SubjectDnToUserAttributeCertificateMapperResponse != nil {
		return obj.SubjectDnToUserAttributeCertificateMapperResponse
	}

	if obj.SubjectEqualsDnCertificateMapperResponse != nil {
		return obj.SubjectEqualsDnCertificateMapperResponse
	}

	if obj.ThirdPartyCertificateMapperResponse != nil {
		return obj.ThirdPartyCertificateMapperResponse
	}

	// all schemas are nil
	return nil
}

type NullableAddCertificateMapper200Response struct {
	value *AddCertificateMapper200Response
	isSet bool
}

func (v NullableAddCertificateMapper200Response) Get() *AddCertificateMapper200Response {
	return v.value
}

func (v *NullableAddCertificateMapper200Response) Set(val *AddCertificateMapper200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableAddCertificateMapper200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableAddCertificateMapper200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddCertificateMapper200Response(val *AddCertificateMapper200Response) *NullableAddCertificateMapper200Response {
	return &NullableAddCertificateMapper200Response{value: val, isSet: true}
}

func (v NullableAddCertificateMapper200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddCertificateMapper200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


