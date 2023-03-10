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

// AddVirtualAttributeRequest - struct for AddVirtualAttributeRequest
type AddVirtualAttributeRequest struct {
	AddConstructedVirtualAttributeRequest             *AddConstructedVirtualAttributeRequest
	AddDnJoinVirtualAttributeRequest                  *AddDnJoinVirtualAttributeRequest
	AddEntryDnVirtualAttributeRequest                 *AddEntryDnVirtualAttributeRequest
	AddEqualityJoinVirtualAttributeRequest            *AddEqualityJoinVirtualAttributeRequest
	AddGroovyScriptedVirtualAttributeRequest          *AddGroovyScriptedVirtualAttributeRequest
	AddIdentifyReferencesVirtualAttributeRequest      *AddIdentifyReferencesVirtualAttributeRequest
	AddIsMemberOfVirtualAttributeRequest              *AddIsMemberOfVirtualAttributeRequest
	AddMemberVirtualAttributeRequest                  *AddMemberVirtualAttributeRequest
	AddMirrorVirtualAttributeRequest                  *AddMirrorVirtualAttributeRequest
	AddPasswordPolicyStateJsonVirtualAttributeRequest *AddPasswordPolicyStateJsonVirtualAttributeRequest
	AddReverseDnJoinVirtualAttributeRequest           *AddReverseDnJoinVirtualAttributeRequest
	AddThirdPartyVirtualAttributeRequest              *AddThirdPartyVirtualAttributeRequest
	AddUserDefinedVirtualAttributeRequest             *AddUserDefinedVirtualAttributeRequest
}

// AddConstructedVirtualAttributeRequestAsAddVirtualAttributeRequest is a convenience function that returns AddConstructedVirtualAttributeRequest wrapped in AddVirtualAttributeRequest
func AddConstructedVirtualAttributeRequestAsAddVirtualAttributeRequest(v *AddConstructedVirtualAttributeRequest) AddVirtualAttributeRequest {
	return AddVirtualAttributeRequest{
		AddConstructedVirtualAttributeRequest: v,
	}
}

// AddDnJoinVirtualAttributeRequestAsAddVirtualAttributeRequest is a convenience function that returns AddDnJoinVirtualAttributeRequest wrapped in AddVirtualAttributeRequest
func AddDnJoinVirtualAttributeRequestAsAddVirtualAttributeRequest(v *AddDnJoinVirtualAttributeRequest) AddVirtualAttributeRequest {
	return AddVirtualAttributeRequest{
		AddDnJoinVirtualAttributeRequest: v,
	}
}

// AddEntryDnVirtualAttributeRequestAsAddVirtualAttributeRequest is a convenience function that returns AddEntryDnVirtualAttributeRequest wrapped in AddVirtualAttributeRequest
func AddEntryDnVirtualAttributeRequestAsAddVirtualAttributeRequest(v *AddEntryDnVirtualAttributeRequest) AddVirtualAttributeRequest {
	return AddVirtualAttributeRequest{
		AddEntryDnVirtualAttributeRequest: v,
	}
}

// AddEqualityJoinVirtualAttributeRequestAsAddVirtualAttributeRequest is a convenience function that returns AddEqualityJoinVirtualAttributeRequest wrapped in AddVirtualAttributeRequest
func AddEqualityJoinVirtualAttributeRequestAsAddVirtualAttributeRequest(v *AddEqualityJoinVirtualAttributeRequest) AddVirtualAttributeRequest {
	return AddVirtualAttributeRequest{
		AddEqualityJoinVirtualAttributeRequest: v,
	}
}

// AddGroovyScriptedVirtualAttributeRequestAsAddVirtualAttributeRequest is a convenience function that returns AddGroovyScriptedVirtualAttributeRequest wrapped in AddVirtualAttributeRequest
func AddGroovyScriptedVirtualAttributeRequestAsAddVirtualAttributeRequest(v *AddGroovyScriptedVirtualAttributeRequest) AddVirtualAttributeRequest {
	return AddVirtualAttributeRequest{
		AddGroovyScriptedVirtualAttributeRequest: v,
	}
}

// AddIdentifyReferencesVirtualAttributeRequestAsAddVirtualAttributeRequest is a convenience function that returns AddIdentifyReferencesVirtualAttributeRequest wrapped in AddVirtualAttributeRequest
func AddIdentifyReferencesVirtualAttributeRequestAsAddVirtualAttributeRequest(v *AddIdentifyReferencesVirtualAttributeRequest) AddVirtualAttributeRequest {
	return AddVirtualAttributeRequest{
		AddIdentifyReferencesVirtualAttributeRequest: v,
	}
}

// AddIsMemberOfVirtualAttributeRequestAsAddVirtualAttributeRequest is a convenience function that returns AddIsMemberOfVirtualAttributeRequest wrapped in AddVirtualAttributeRequest
func AddIsMemberOfVirtualAttributeRequestAsAddVirtualAttributeRequest(v *AddIsMemberOfVirtualAttributeRequest) AddVirtualAttributeRequest {
	return AddVirtualAttributeRequest{
		AddIsMemberOfVirtualAttributeRequest: v,
	}
}

// AddMemberVirtualAttributeRequestAsAddVirtualAttributeRequest is a convenience function that returns AddMemberVirtualAttributeRequest wrapped in AddVirtualAttributeRequest
func AddMemberVirtualAttributeRequestAsAddVirtualAttributeRequest(v *AddMemberVirtualAttributeRequest) AddVirtualAttributeRequest {
	return AddVirtualAttributeRequest{
		AddMemberVirtualAttributeRequest: v,
	}
}

// AddMirrorVirtualAttributeRequestAsAddVirtualAttributeRequest is a convenience function that returns AddMirrorVirtualAttributeRequest wrapped in AddVirtualAttributeRequest
func AddMirrorVirtualAttributeRequestAsAddVirtualAttributeRequest(v *AddMirrorVirtualAttributeRequest) AddVirtualAttributeRequest {
	return AddVirtualAttributeRequest{
		AddMirrorVirtualAttributeRequest: v,
	}
}

// AddPasswordPolicyStateJsonVirtualAttributeRequestAsAddVirtualAttributeRequest is a convenience function that returns AddPasswordPolicyStateJsonVirtualAttributeRequest wrapped in AddVirtualAttributeRequest
func AddPasswordPolicyStateJsonVirtualAttributeRequestAsAddVirtualAttributeRequest(v *AddPasswordPolicyStateJsonVirtualAttributeRequest) AddVirtualAttributeRequest {
	return AddVirtualAttributeRequest{
		AddPasswordPolicyStateJsonVirtualAttributeRequest: v,
	}
}

// AddReverseDnJoinVirtualAttributeRequestAsAddVirtualAttributeRequest is a convenience function that returns AddReverseDnJoinVirtualAttributeRequest wrapped in AddVirtualAttributeRequest
func AddReverseDnJoinVirtualAttributeRequestAsAddVirtualAttributeRequest(v *AddReverseDnJoinVirtualAttributeRequest) AddVirtualAttributeRequest {
	return AddVirtualAttributeRequest{
		AddReverseDnJoinVirtualAttributeRequest: v,
	}
}

// AddThirdPartyVirtualAttributeRequestAsAddVirtualAttributeRequest is a convenience function that returns AddThirdPartyVirtualAttributeRequest wrapped in AddVirtualAttributeRequest
func AddThirdPartyVirtualAttributeRequestAsAddVirtualAttributeRequest(v *AddThirdPartyVirtualAttributeRequest) AddVirtualAttributeRequest {
	return AddVirtualAttributeRequest{
		AddThirdPartyVirtualAttributeRequest: v,
	}
}

// AddUserDefinedVirtualAttributeRequestAsAddVirtualAttributeRequest is a convenience function that returns AddUserDefinedVirtualAttributeRequest wrapped in AddVirtualAttributeRequest
func AddUserDefinedVirtualAttributeRequestAsAddVirtualAttributeRequest(v *AddUserDefinedVirtualAttributeRequest) AddVirtualAttributeRequest {
	return AddVirtualAttributeRequest{
		AddUserDefinedVirtualAttributeRequest: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *AddVirtualAttributeRequest) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AddConstructedVirtualAttributeRequest
	err = newStrictDecoder(data).Decode(&dst.AddConstructedVirtualAttributeRequest)
	if err == nil {
		jsonAddConstructedVirtualAttributeRequest, _ := json.Marshal(dst.AddConstructedVirtualAttributeRequest)
		if string(jsonAddConstructedVirtualAttributeRequest) == "{}" { // empty struct
			dst.AddConstructedVirtualAttributeRequest = nil
		} else {
			match++
		}
	} else {
		dst.AddConstructedVirtualAttributeRequest = nil
	}

	// try to unmarshal data into AddDnJoinVirtualAttributeRequest
	err = newStrictDecoder(data).Decode(&dst.AddDnJoinVirtualAttributeRequest)
	if err == nil {
		jsonAddDnJoinVirtualAttributeRequest, _ := json.Marshal(dst.AddDnJoinVirtualAttributeRequest)
		if string(jsonAddDnJoinVirtualAttributeRequest) == "{}" { // empty struct
			dst.AddDnJoinVirtualAttributeRequest = nil
		} else {
			match++
		}
	} else {
		dst.AddDnJoinVirtualAttributeRequest = nil
	}

	// try to unmarshal data into AddEntryDnVirtualAttributeRequest
	err = newStrictDecoder(data).Decode(&dst.AddEntryDnVirtualAttributeRequest)
	if err == nil {
		jsonAddEntryDnVirtualAttributeRequest, _ := json.Marshal(dst.AddEntryDnVirtualAttributeRequest)
		if string(jsonAddEntryDnVirtualAttributeRequest) == "{}" { // empty struct
			dst.AddEntryDnVirtualAttributeRequest = nil
		} else {
			match++
		}
	} else {
		dst.AddEntryDnVirtualAttributeRequest = nil
	}

	// try to unmarshal data into AddEqualityJoinVirtualAttributeRequest
	err = newStrictDecoder(data).Decode(&dst.AddEqualityJoinVirtualAttributeRequest)
	if err == nil {
		jsonAddEqualityJoinVirtualAttributeRequest, _ := json.Marshal(dst.AddEqualityJoinVirtualAttributeRequest)
		if string(jsonAddEqualityJoinVirtualAttributeRequest) == "{}" { // empty struct
			dst.AddEqualityJoinVirtualAttributeRequest = nil
		} else {
			match++
		}
	} else {
		dst.AddEqualityJoinVirtualAttributeRequest = nil
	}

	// try to unmarshal data into AddGroovyScriptedVirtualAttributeRequest
	err = newStrictDecoder(data).Decode(&dst.AddGroovyScriptedVirtualAttributeRequest)
	if err == nil {
		jsonAddGroovyScriptedVirtualAttributeRequest, _ := json.Marshal(dst.AddGroovyScriptedVirtualAttributeRequest)
		if string(jsonAddGroovyScriptedVirtualAttributeRequest) == "{}" { // empty struct
			dst.AddGroovyScriptedVirtualAttributeRequest = nil
		} else {
			match++
		}
	} else {
		dst.AddGroovyScriptedVirtualAttributeRequest = nil
	}

	// try to unmarshal data into AddIdentifyReferencesVirtualAttributeRequest
	err = newStrictDecoder(data).Decode(&dst.AddIdentifyReferencesVirtualAttributeRequest)
	if err == nil {
		jsonAddIdentifyReferencesVirtualAttributeRequest, _ := json.Marshal(dst.AddIdentifyReferencesVirtualAttributeRequest)
		if string(jsonAddIdentifyReferencesVirtualAttributeRequest) == "{}" { // empty struct
			dst.AddIdentifyReferencesVirtualAttributeRequest = nil
		} else {
			match++
		}
	} else {
		dst.AddIdentifyReferencesVirtualAttributeRequest = nil
	}

	// try to unmarshal data into AddIsMemberOfVirtualAttributeRequest
	err = newStrictDecoder(data).Decode(&dst.AddIsMemberOfVirtualAttributeRequest)
	if err == nil {
		jsonAddIsMemberOfVirtualAttributeRequest, _ := json.Marshal(dst.AddIsMemberOfVirtualAttributeRequest)
		if string(jsonAddIsMemberOfVirtualAttributeRequest) == "{}" { // empty struct
			dst.AddIsMemberOfVirtualAttributeRequest = nil
		} else {
			match++
		}
	} else {
		dst.AddIsMemberOfVirtualAttributeRequest = nil
	}

	// try to unmarshal data into AddMemberVirtualAttributeRequest
	err = newStrictDecoder(data).Decode(&dst.AddMemberVirtualAttributeRequest)
	if err == nil {
		jsonAddMemberVirtualAttributeRequest, _ := json.Marshal(dst.AddMemberVirtualAttributeRequest)
		if string(jsonAddMemberVirtualAttributeRequest) == "{}" { // empty struct
			dst.AddMemberVirtualAttributeRequest = nil
		} else {
			match++
		}
	} else {
		dst.AddMemberVirtualAttributeRequest = nil
	}

	// try to unmarshal data into AddMirrorVirtualAttributeRequest
	err = newStrictDecoder(data).Decode(&dst.AddMirrorVirtualAttributeRequest)
	if err == nil {
		jsonAddMirrorVirtualAttributeRequest, _ := json.Marshal(dst.AddMirrorVirtualAttributeRequest)
		if string(jsonAddMirrorVirtualAttributeRequest) == "{}" { // empty struct
			dst.AddMirrorVirtualAttributeRequest = nil
		} else {
			match++
		}
	} else {
		dst.AddMirrorVirtualAttributeRequest = nil
	}

	// try to unmarshal data into AddPasswordPolicyStateJsonVirtualAttributeRequest
	err = newStrictDecoder(data).Decode(&dst.AddPasswordPolicyStateJsonVirtualAttributeRequest)
	if err == nil {
		jsonAddPasswordPolicyStateJsonVirtualAttributeRequest, _ := json.Marshal(dst.AddPasswordPolicyStateJsonVirtualAttributeRequest)
		if string(jsonAddPasswordPolicyStateJsonVirtualAttributeRequest) == "{}" { // empty struct
			dst.AddPasswordPolicyStateJsonVirtualAttributeRequest = nil
		} else {
			match++
		}
	} else {
		dst.AddPasswordPolicyStateJsonVirtualAttributeRequest = nil
	}

	// try to unmarshal data into AddReverseDnJoinVirtualAttributeRequest
	err = newStrictDecoder(data).Decode(&dst.AddReverseDnJoinVirtualAttributeRequest)
	if err == nil {
		jsonAddReverseDnJoinVirtualAttributeRequest, _ := json.Marshal(dst.AddReverseDnJoinVirtualAttributeRequest)
		if string(jsonAddReverseDnJoinVirtualAttributeRequest) == "{}" { // empty struct
			dst.AddReverseDnJoinVirtualAttributeRequest = nil
		} else {
			match++
		}
	} else {
		dst.AddReverseDnJoinVirtualAttributeRequest = nil
	}

	// try to unmarshal data into AddThirdPartyVirtualAttributeRequest
	err = newStrictDecoder(data).Decode(&dst.AddThirdPartyVirtualAttributeRequest)
	if err == nil {
		jsonAddThirdPartyVirtualAttributeRequest, _ := json.Marshal(dst.AddThirdPartyVirtualAttributeRequest)
		if string(jsonAddThirdPartyVirtualAttributeRequest) == "{}" { // empty struct
			dst.AddThirdPartyVirtualAttributeRequest = nil
		} else {
			match++
		}
	} else {
		dst.AddThirdPartyVirtualAttributeRequest = nil
	}

	// try to unmarshal data into AddUserDefinedVirtualAttributeRequest
	err = newStrictDecoder(data).Decode(&dst.AddUserDefinedVirtualAttributeRequest)
	if err == nil {
		jsonAddUserDefinedVirtualAttributeRequest, _ := json.Marshal(dst.AddUserDefinedVirtualAttributeRequest)
		if string(jsonAddUserDefinedVirtualAttributeRequest) == "{}" { // empty struct
			dst.AddUserDefinedVirtualAttributeRequest = nil
		} else {
			match++
		}
	} else {
		dst.AddUserDefinedVirtualAttributeRequest = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.AddConstructedVirtualAttributeRequest = nil
		dst.AddDnJoinVirtualAttributeRequest = nil
		dst.AddEntryDnVirtualAttributeRequest = nil
		dst.AddEqualityJoinVirtualAttributeRequest = nil
		dst.AddGroovyScriptedVirtualAttributeRequest = nil
		dst.AddIdentifyReferencesVirtualAttributeRequest = nil
		dst.AddIsMemberOfVirtualAttributeRequest = nil
		dst.AddMemberVirtualAttributeRequest = nil
		dst.AddMirrorVirtualAttributeRequest = nil
		dst.AddPasswordPolicyStateJsonVirtualAttributeRequest = nil
		dst.AddReverseDnJoinVirtualAttributeRequest = nil
		dst.AddThirdPartyVirtualAttributeRequest = nil
		dst.AddUserDefinedVirtualAttributeRequest = nil

		return fmt.Errorf("data matches more than one schema in oneOf(AddVirtualAttributeRequest)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(AddVirtualAttributeRequest)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src AddVirtualAttributeRequest) MarshalJSON() ([]byte, error) {
	if src.AddConstructedVirtualAttributeRequest != nil {
		return json.Marshal(&src.AddConstructedVirtualAttributeRequest)
	}

	if src.AddDnJoinVirtualAttributeRequest != nil {
		return json.Marshal(&src.AddDnJoinVirtualAttributeRequest)
	}

	if src.AddEntryDnVirtualAttributeRequest != nil {
		return json.Marshal(&src.AddEntryDnVirtualAttributeRequest)
	}

	if src.AddEqualityJoinVirtualAttributeRequest != nil {
		return json.Marshal(&src.AddEqualityJoinVirtualAttributeRequest)
	}

	if src.AddGroovyScriptedVirtualAttributeRequest != nil {
		return json.Marshal(&src.AddGroovyScriptedVirtualAttributeRequest)
	}

	if src.AddIdentifyReferencesVirtualAttributeRequest != nil {
		return json.Marshal(&src.AddIdentifyReferencesVirtualAttributeRequest)
	}

	if src.AddIsMemberOfVirtualAttributeRequest != nil {
		return json.Marshal(&src.AddIsMemberOfVirtualAttributeRequest)
	}

	if src.AddMemberVirtualAttributeRequest != nil {
		return json.Marshal(&src.AddMemberVirtualAttributeRequest)
	}

	if src.AddMirrorVirtualAttributeRequest != nil {
		return json.Marshal(&src.AddMirrorVirtualAttributeRequest)
	}

	if src.AddPasswordPolicyStateJsonVirtualAttributeRequest != nil {
		return json.Marshal(&src.AddPasswordPolicyStateJsonVirtualAttributeRequest)
	}

	if src.AddReverseDnJoinVirtualAttributeRequest != nil {
		return json.Marshal(&src.AddReverseDnJoinVirtualAttributeRequest)
	}

	if src.AddThirdPartyVirtualAttributeRequest != nil {
		return json.Marshal(&src.AddThirdPartyVirtualAttributeRequest)
	}

	if src.AddUserDefinedVirtualAttributeRequest != nil {
		return json.Marshal(&src.AddUserDefinedVirtualAttributeRequest)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *AddVirtualAttributeRequest) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.AddConstructedVirtualAttributeRequest != nil {
		return obj.AddConstructedVirtualAttributeRequest
	}

	if obj.AddDnJoinVirtualAttributeRequest != nil {
		return obj.AddDnJoinVirtualAttributeRequest
	}

	if obj.AddEntryDnVirtualAttributeRequest != nil {
		return obj.AddEntryDnVirtualAttributeRequest
	}

	if obj.AddEqualityJoinVirtualAttributeRequest != nil {
		return obj.AddEqualityJoinVirtualAttributeRequest
	}

	if obj.AddGroovyScriptedVirtualAttributeRequest != nil {
		return obj.AddGroovyScriptedVirtualAttributeRequest
	}

	if obj.AddIdentifyReferencesVirtualAttributeRequest != nil {
		return obj.AddIdentifyReferencesVirtualAttributeRequest
	}

	if obj.AddIsMemberOfVirtualAttributeRequest != nil {
		return obj.AddIsMemberOfVirtualAttributeRequest
	}

	if obj.AddMemberVirtualAttributeRequest != nil {
		return obj.AddMemberVirtualAttributeRequest
	}

	if obj.AddMirrorVirtualAttributeRequest != nil {
		return obj.AddMirrorVirtualAttributeRequest
	}

	if obj.AddPasswordPolicyStateJsonVirtualAttributeRequest != nil {
		return obj.AddPasswordPolicyStateJsonVirtualAttributeRequest
	}

	if obj.AddReverseDnJoinVirtualAttributeRequest != nil {
		return obj.AddReverseDnJoinVirtualAttributeRequest
	}

	if obj.AddThirdPartyVirtualAttributeRequest != nil {
		return obj.AddThirdPartyVirtualAttributeRequest
	}

	if obj.AddUserDefinedVirtualAttributeRequest != nil {
		return obj.AddUserDefinedVirtualAttributeRequest
	}

	// all schemas are nil
	return nil
}

type NullableAddVirtualAttributeRequest struct {
	value *AddVirtualAttributeRequest
	isSet bool
}

func (v NullableAddVirtualAttributeRequest) Get() *AddVirtualAttributeRequest {
	return v.value
}

func (v *NullableAddVirtualAttributeRequest) Set(val *AddVirtualAttributeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddVirtualAttributeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddVirtualAttributeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddVirtualAttributeRequest(val *AddVirtualAttributeRequest) *NullableAddVirtualAttributeRequest {
	return &NullableAddVirtualAttributeRequest{value: val, isSet: true}
}

func (v NullableAddVirtualAttributeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddVirtualAttributeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
