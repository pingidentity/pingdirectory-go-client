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

// AddCertificateMapperRequest - struct for AddCertificateMapperRequest
type AddCertificateMapperRequest struct {
	AddFingerprintCertificateMapperRequest *AddFingerprintCertificateMapperRequest
	AddGroovyScriptedCertificateMapperRequest *AddGroovyScriptedCertificateMapperRequest
	AddSubjectAttributeToUserAttributeCertificateMapperRequest *AddSubjectAttributeToUserAttributeCertificateMapperRequest
	AddSubjectDnToUserAttributeCertificateMapperRequest *AddSubjectDnToUserAttributeCertificateMapperRequest
	AddSubjectEqualsDnCertificateMapperRequest *AddSubjectEqualsDnCertificateMapperRequest
	AddThirdPartyCertificateMapperRequest *AddThirdPartyCertificateMapperRequest
}

// AddFingerprintCertificateMapperRequestAsAddCertificateMapperRequest is a convenience function that returns AddFingerprintCertificateMapperRequest wrapped in AddCertificateMapperRequest
func AddFingerprintCertificateMapperRequestAsAddCertificateMapperRequest(v *AddFingerprintCertificateMapperRequest) AddCertificateMapperRequest {
	return AddCertificateMapperRequest{
		AddFingerprintCertificateMapperRequest: v,
	}
}

// AddGroovyScriptedCertificateMapperRequestAsAddCertificateMapperRequest is a convenience function that returns AddGroovyScriptedCertificateMapperRequest wrapped in AddCertificateMapperRequest
func AddGroovyScriptedCertificateMapperRequestAsAddCertificateMapperRequest(v *AddGroovyScriptedCertificateMapperRequest) AddCertificateMapperRequest {
	return AddCertificateMapperRequest{
		AddGroovyScriptedCertificateMapperRequest: v,
	}
}

// AddSubjectAttributeToUserAttributeCertificateMapperRequestAsAddCertificateMapperRequest is a convenience function that returns AddSubjectAttributeToUserAttributeCertificateMapperRequest wrapped in AddCertificateMapperRequest
func AddSubjectAttributeToUserAttributeCertificateMapperRequestAsAddCertificateMapperRequest(v *AddSubjectAttributeToUserAttributeCertificateMapperRequest) AddCertificateMapperRequest {
	return AddCertificateMapperRequest{
		AddSubjectAttributeToUserAttributeCertificateMapperRequest: v,
	}
}

// AddSubjectDnToUserAttributeCertificateMapperRequestAsAddCertificateMapperRequest is a convenience function that returns AddSubjectDnToUserAttributeCertificateMapperRequest wrapped in AddCertificateMapperRequest
func AddSubjectDnToUserAttributeCertificateMapperRequestAsAddCertificateMapperRequest(v *AddSubjectDnToUserAttributeCertificateMapperRequest) AddCertificateMapperRequest {
	return AddCertificateMapperRequest{
		AddSubjectDnToUserAttributeCertificateMapperRequest: v,
	}
}

// AddSubjectEqualsDnCertificateMapperRequestAsAddCertificateMapperRequest is a convenience function that returns AddSubjectEqualsDnCertificateMapperRequest wrapped in AddCertificateMapperRequest
func AddSubjectEqualsDnCertificateMapperRequestAsAddCertificateMapperRequest(v *AddSubjectEqualsDnCertificateMapperRequest) AddCertificateMapperRequest {
	return AddCertificateMapperRequest{
		AddSubjectEqualsDnCertificateMapperRequest: v,
	}
}

// AddThirdPartyCertificateMapperRequestAsAddCertificateMapperRequest is a convenience function that returns AddThirdPartyCertificateMapperRequest wrapped in AddCertificateMapperRequest
func AddThirdPartyCertificateMapperRequestAsAddCertificateMapperRequest(v *AddThirdPartyCertificateMapperRequest) AddCertificateMapperRequest {
	return AddCertificateMapperRequest{
		AddThirdPartyCertificateMapperRequest: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *AddCertificateMapperRequest) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AddFingerprintCertificateMapperRequest
	err = newStrictDecoder(data).Decode(&dst.AddFingerprintCertificateMapperRequest)
	if err == nil {
		jsonAddFingerprintCertificateMapperRequest, _ := json.Marshal(dst.AddFingerprintCertificateMapperRequest)
		if string(jsonAddFingerprintCertificateMapperRequest) == "{}" { // empty struct
			dst.AddFingerprintCertificateMapperRequest = nil
		} else {
			match++
		}
	} else {
		dst.AddFingerprintCertificateMapperRequest = nil
	}

	// try to unmarshal data into AddGroovyScriptedCertificateMapperRequest
	err = newStrictDecoder(data).Decode(&dst.AddGroovyScriptedCertificateMapperRequest)
	if err == nil {
		jsonAddGroovyScriptedCertificateMapperRequest, _ := json.Marshal(dst.AddGroovyScriptedCertificateMapperRequest)
		if string(jsonAddGroovyScriptedCertificateMapperRequest) == "{}" { // empty struct
			dst.AddGroovyScriptedCertificateMapperRequest = nil
		} else {
			match++
		}
	} else {
		dst.AddGroovyScriptedCertificateMapperRequest = nil
	}

	// try to unmarshal data into AddSubjectAttributeToUserAttributeCertificateMapperRequest
	err = newStrictDecoder(data).Decode(&dst.AddSubjectAttributeToUserAttributeCertificateMapperRequest)
	if err == nil {
		jsonAddSubjectAttributeToUserAttributeCertificateMapperRequest, _ := json.Marshal(dst.AddSubjectAttributeToUserAttributeCertificateMapperRequest)
		if string(jsonAddSubjectAttributeToUserAttributeCertificateMapperRequest) == "{}" { // empty struct
			dst.AddSubjectAttributeToUserAttributeCertificateMapperRequest = nil
		} else {
			match++
		}
	} else {
		dst.AddSubjectAttributeToUserAttributeCertificateMapperRequest = nil
	}

	// try to unmarshal data into AddSubjectDnToUserAttributeCertificateMapperRequest
	err = newStrictDecoder(data).Decode(&dst.AddSubjectDnToUserAttributeCertificateMapperRequest)
	if err == nil {
		jsonAddSubjectDnToUserAttributeCertificateMapperRequest, _ := json.Marshal(dst.AddSubjectDnToUserAttributeCertificateMapperRequest)
		if string(jsonAddSubjectDnToUserAttributeCertificateMapperRequest) == "{}" { // empty struct
			dst.AddSubjectDnToUserAttributeCertificateMapperRequest = nil
		} else {
			match++
		}
	} else {
		dst.AddSubjectDnToUserAttributeCertificateMapperRequest = nil
	}

	// try to unmarshal data into AddSubjectEqualsDnCertificateMapperRequest
	err = newStrictDecoder(data).Decode(&dst.AddSubjectEqualsDnCertificateMapperRequest)
	if err == nil {
		jsonAddSubjectEqualsDnCertificateMapperRequest, _ := json.Marshal(dst.AddSubjectEqualsDnCertificateMapperRequest)
		if string(jsonAddSubjectEqualsDnCertificateMapperRequest) == "{}" { // empty struct
			dst.AddSubjectEqualsDnCertificateMapperRequest = nil
		} else {
			match++
		}
	} else {
		dst.AddSubjectEqualsDnCertificateMapperRequest = nil
	}

	// try to unmarshal data into AddThirdPartyCertificateMapperRequest
	err = newStrictDecoder(data).Decode(&dst.AddThirdPartyCertificateMapperRequest)
	if err == nil {
		jsonAddThirdPartyCertificateMapperRequest, _ := json.Marshal(dst.AddThirdPartyCertificateMapperRequest)
		if string(jsonAddThirdPartyCertificateMapperRequest) == "{}" { // empty struct
			dst.AddThirdPartyCertificateMapperRequest = nil
		} else {
			match++
		}
	} else {
		dst.AddThirdPartyCertificateMapperRequest = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.AddFingerprintCertificateMapperRequest = nil
		dst.AddGroovyScriptedCertificateMapperRequest = nil
		dst.AddSubjectAttributeToUserAttributeCertificateMapperRequest = nil
		dst.AddSubjectDnToUserAttributeCertificateMapperRequest = nil
		dst.AddSubjectEqualsDnCertificateMapperRequest = nil
		dst.AddThirdPartyCertificateMapperRequest = nil

		return fmt.Errorf("data matches more than one schema in oneOf(AddCertificateMapperRequest)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(AddCertificateMapperRequest)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src AddCertificateMapperRequest) MarshalJSON() ([]byte, error) {
	if src.AddFingerprintCertificateMapperRequest != nil {
		return json.Marshal(&src.AddFingerprintCertificateMapperRequest)
	}

	if src.AddGroovyScriptedCertificateMapperRequest != nil {
		return json.Marshal(&src.AddGroovyScriptedCertificateMapperRequest)
	}

	if src.AddSubjectAttributeToUserAttributeCertificateMapperRequest != nil {
		return json.Marshal(&src.AddSubjectAttributeToUserAttributeCertificateMapperRequest)
	}

	if src.AddSubjectDnToUserAttributeCertificateMapperRequest != nil {
		return json.Marshal(&src.AddSubjectDnToUserAttributeCertificateMapperRequest)
	}

	if src.AddSubjectEqualsDnCertificateMapperRequest != nil {
		return json.Marshal(&src.AddSubjectEqualsDnCertificateMapperRequest)
	}

	if src.AddThirdPartyCertificateMapperRequest != nil {
		return json.Marshal(&src.AddThirdPartyCertificateMapperRequest)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *AddCertificateMapperRequest) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.AddFingerprintCertificateMapperRequest != nil {
		return obj.AddFingerprintCertificateMapperRequest
	}

	if obj.AddGroovyScriptedCertificateMapperRequest != nil {
		return obj.AddGroovyScriptedCertificateMapperRequest
	}

	if obj.AddSubjectAttributeToUserAttributeCertificateMapperRequest != nil {
		return obj.AddSubjectAttributeToUserAttributeCertificateMapperRequest
	}

	if obj.AddSubjectDnToUserAttributeCertificateMapperRequest != nil {
		return obj.AddSubjectDnToUserAttributeCertificateMapperRequest
	}

	if obj.AddSubjectEqualsDnCertificateMapperRequest != nil {
		return obj.AddSubjectEqualsDnCertificateMapperRequest
	}

	if obj.AddThirdPartyCertificateMapperRequest != nil {
		return obj.AddThirdPartyCertificateMapperRequest
	}

	// all schemas are nil
	return nil
}

type NullableAddCertificateMapperRequest struct {
	value *AddCertificateMapperRequest
	isSet bool
}

func (v NullableAddCertificateMapperRequest) Get() *AddCertificateMapperRequest {
	return v.value
}

func (v *NullableAddCertificateMapperRequest) Set(val *AddCertificateMapperRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddCertificateMapperRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddCertificateMapperRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddCertificateMapperRequest(val *AddCertificateMapperRequest) *NullableAddCertificateMapperRequest {
	return &NullableAddCertificateMapperRequest{value: val, isSet: true}
}

func (v NullableAddCertificateMapperRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddCertificateMapperRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

