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

// GetVirtualAttribute200Response - struct for GetVirtualAttribute200Response
type GetVirtualAttribute200Response struct {
	ConstructedVirtualAttributeResponse             *ConstructedVirtualAttributeResponse
	CurrentTimeVirtualAttributeResponse             *CurrentTimeVirtualAttributeResponse
	DnJoinVirtualAttributeResponse                  *DnJoinVirtualAttributeResponse
	EntryChecksumVirtualAttributeResponse           *EntryChecksumVirtualAttributeResponse
	EntryDnVirtualAttributeResponse                 *EntryDnVirtualAttributeResponse
	EqualityJoinVirtualAttributeResponse            *EqualityJoinVirtualAttributeResponse
	GroovyScriptedVirtualAttributeResponse          *GroovyScriptedVirtualAttributeResponse
	HasSubordinatesVirtualAttributeResponse         *HasSubordinatesVirtualAttributeResponse
	IdentifyReferencesVirtualAttributeResponse      *IdentifyReferencesVirtualAttributeResponse
	InstanceNameVirtualAttributeResponse            *InstanceNameVirtualAttributeResponse
	IsMemberOfVirtualAttributeResponse              *IsMemberOfVirtualAttributeResponse
	MemberOfServerGroupVirtualAttributeResponse     *MemberOfServerGroupVirtualAttributeResponse
	MemberVirtualAttributeResponse                  *MemberVirtualAttributeResponse
	MirrorVirtualAttributeResponse                  *MirrorVirtualAttributeResponse
	NumSubordinatesVirtualAttributeResponse         *NumSubordinatesVirtualAttributeResponse
	PasswordPolicyStateJsonVirtualAttributeResponse *PasswordPolicyStateJsonVirtualAttributeResponse
	ReplicationStateDetailVirtualAttributeResponse  *ReplicationStateDetailVirtualAttributeResponse
	ReverseDnJoinVirtualAttributeResponse           *ReverseDnJoinVirtualAttributeResponse
	ShortUniqueIdVirtualAttributeResponse           *ShortUniqueIdVirtualAttributeResponse
	SubschemaSubentryVirtualAttributeResponse       *SubschemaSubentryVirtualAttributeResponse
	ThirdPartyVirtualAttributeResponse              *ThirdPartyVirtualAttributeResponse
	UserDefinedVirtualAttributeResponse             *UserDefinedVirtualAttributeResponse
}

// ConstructedVirtualAttributeResponseAsGetVirtualAttribute200Response is a convenience function that returns ConstructedVirtualAttributeResponse wrapped in GetVirtualAttribute200Response
func ConstructedVirtualAttributeResponseAsGetVirtualAttribute200Response(v *ConstructedVirtualAttributeResponse) GetVirtualAttribute200Response {
	return GetVirtualAttribute200Response{
		ConstructedVirtualAttributeResponse: v,
	}
}

// CurrentTimeVirtualAttributeResponseAsGetVirtualAttribute200Response is a convenience function that returns CurrentTimeVirtualAttributeResponse wrapped in GetVirtualAttribute200Response
func CurrentTimeVirtualAttributeResponseAsGetVirtualAttribute200Response(v *CurrentTimeVirtualAttributeResponse) GetVirtualAttribute200Response {
	return GetVirtualAttribute200Response{
		CurrentTimeVirtualAttributeResponse: v,
	}
}

// DnJoinVirtualAttributeResponseAsGetVirtualAttribute200Response is a convenience function that returns DnJoinVirtualAttributeResponse wrapped in GetVirtualAttribute200Response
func DnJoinVirtualAttributeResponseAsGetVirtualAttribute200Response(v *DnJoinVirtualAttributeResponse) GetVirtualAttribute200Response {
	return GetVirtualAttribute200Response{
		DnJoinVirtualAttributeResponse: v,
	}
}

// EntryChecksumVirtualAttributeResponseAsGetVirtualAttribute200Response is a convenience function that returns EntryChecksumVirtualAttributeResponse wrapped in GetVirtualAttribute200Response
func EntryChecksumVirtualAttributeResponseAsGetVirtualAttribute200Response(v *EntryChecksumVirtualAttributeResponse) GetVirtualAttribute200Response {
	return GetVirtualAttribute200Response{
		EntryChecksumVirtualAttributeResponse: v,
	}
}

// EntryDnVirtualAttributeResponseAsGetVirtualAttribute200Response is a convenience function that returns EntryDnVirtualAttributeResponse wrapped in GetVirtualAttribute200Response
func EntryDnVirtualAttributeResponseAsGetVirtualAttribute200Response(v *EntryDnVirtualAttributeResponse) GetVirtualAttribute200Response {
	return GetVirtualAttribute200Response{
		EntryDnVirtualAttributeResponse: v,
	}
}

// EqualityJoinVirtualAttributeResponseAsGetVirtualAttribute200Response is a convenience function that returns EqualityJoinVirtualAttributeResponse wrapped in GetVirtualAttribute200Response
func EqualityJoinVirtualAttributeResponseAsGetVirtualAttribute200Response(v *EqualityJoinVirtualAttributeResponse) GetVirtualAttribute200Response {
	return GetVirtualAttribute200Response{
		EqualityJoinVirtualAttributeResponse: v,
	}
}

// GroovyScriptedVirtualAttributeResponseAsGetVirtualAttribute200Response is a convenience function that returns GroovyScriptedVirtualAttributeResponse wrapped in GetVirtualAttribute200Response
func GroovyScriptedVirtualAttributeResponseAsGetVirtualAttribute200Response(v *GroovyScriptedVirtualAttributeResponse) GetVirtualAttribute200Response {
	return GetVirtualAttribute200Response{
		GroovyScriptedVirtualAttributeResponse: v,
	}
}

// HasSubordinatesVirtualAttributeResponseAsGetVirtualAttribute200Response is a convenience function that returns HasSubordinatesVirtualAttributeResponse wrapped in GetVirtualAttribute200Response
func HasSubordinatesVirtualAttributeResponseAsGetVirtualAttribute200Response(v *HasSubordinatesVirtualAttributeResponse) GetVirtualAttribute200Response {
	return GetVirtualAttribute200Response{
		HasSubordinatesVirtualAttributeResponse: v,
	}
}

// IdentifyReferencesVirtualAttributeResponseAsGetVirtualAttribute200Response is a convenience function that returns IdentifyReferencesVirtualAttributeResponse wrapped in GetVirtualAttribute200Response
func IdentifyReferencesVirtualAttributeResponseAsGetVirtualAttribute200Response(v *IdentifyReferencesVirtualAttributeResponse) GetVirtualAttribute200Response {
	return GetVirtualAttribute200Response{
		IdentifyReferencesVirtualAttributeResponse: v,
	}
}

// InstanceNameVirtualAttributeResponseAsGetVirtualAttribute200Response is a convenience function that returns InstanceNameVirtualAttributeResponse wrapped in GetVirtualAttribute200Response
func InstanceNameVirtualAttributeResponseAsGetVirtualAttribute200Response(v *InstanceNameVirtualAttributeResponse) GetVirtualAttribute200Response {
	return GetVirtualAttribute200Response{
		InstanceNameVirtualAttributeResponse: v,
	}
}

// IsMemberOfVirtualAttributeResponseAsGetVirtualAttribute200Response is a convenience function that returns IsMemberOfVirtualAttributeResponse wrapped in GetVirtualAttribute200Response
func IsMemberOfVirtualAttributeResponseAsGetVirtualAttribute200Response(v *IsMemberOfVirtualAttributeResponse) GetVirtualAttribute200Response {
	return GetVirtualAttribute200Response{
		IsMemberOfVirtualAttributeResponse: v,
	}
}

// MemberOfServerGroupVirtualAttributeResponseAsGetVirtualAttribute200Response is a convenience function that returns MemberOfServerGroupVirtualAttributeResponse wrapped in GetVirtualAttribute200Response
func MemberOfServerGroupVirtualAttributeResponseAsGetVirtualAttribute200Response(v *MemberOfServerGroupVirtualAttributeResponse) GetVirtualAttribute200Response {
	return GetVirtualAttribute200Response{
		MemberOfServerGroupVirtualAttributeResponse: v,
	}
}

// MemberVirtualAttributeResponseAsGetVirtualAttribute200Response is a convenience function that returns MemberVirtualAttributeResponse wrapped in GetVirtualAttribute200Response
func MemberVirtualAttributeResponseAsGetVirtualAttribute200Response(v *MemberVirtualAttributeResponse) GetVirtualAttribute200Response {
	return GetVirtualAttribute200Response{
		MemberVirtualAttributeResponse: v,
	}
}

// MirrorVirtualAttributeResponseAsGetVirtualAttribute200Response is a convenience function that returns MirrorVirtualAttributeResponse wrapped in GetVirtualAttribute200Response
func MirrorVirtualAttributeResponseAsGetVirtualAttribute200Response(v *MirrorVirtualAttributeResponse) GetVirtualAttribute200Response {
	return GetVirtualAttribute200Response{
		MirrorVirtualAttributeResponse: v,
	}
}

// NumSubordinatesVirtualAttributeResponseAsGetVirtualAttribute200Response is a convenience function that returns NumSubordinatesVirtualAttributeResponse wrapped in GetVirtualAttribute200Response
func NumSubordinatesVirtualAttributeResponseAsGetVirtualAttribute200Response(v *NumSubordinatesVirtualAttributeResponse) GetVirtualAttribute200Response {
	return GetVirtualAttribute200Response{
		NumSubordinatesVirtualAttributeResponse: v,
	}
}

// PasswordPolicyStateJsonVirtualAttributeResponseAsGetVirtualAttribute200Response is a convenience function that returns PasswordPolicyStateJsonVirtualAttributeResponse wrapped in GetVirtualAttribute200Response
func PasswordPolicyStateJsonVirtualAttributeResponseAsGetVirtualAttribute200Response(v *PasswordPolicyStateJsonVirtualAttributeResponse) GetVirtualAttribute200Response {
	return GetVirtualAttribute200Response{
		PasswordPolicyStateJsonVirtualAttributeResponse: v,
	}
}

// ReplicationStateDetailVirtualAttributeResponseAsGetVirtualAttribute200Response is a convenience function that returns ReplicationStateDetailVirtualAttributeResponse wrapped in GetVirtualAttribute200Response
func ReplicationStateDetailVirtualAttributeResponseAsGetVirtualAttribute200Response(v *ReplicationStateDetailVirtualAttributeResponse) GetVirtualAttribute200Response {
	return GetVirtualAttribute200Response{
		ReplicationStateDetailVirtualAttributeResponse: v,
	}
}

// ReverseDnJoinVirtualAttributeResponseAsGetVirtualAttribute200Response is a convenience function that returns ReverseDnJoinVirtualAttributeResponse wrapped in GetVirtualAttribute200Response
func ReverseDnJoinVirtualAttributeResponseAsGetVirtualAttribute200Response(v *ReverseDnJoinVirtualAttributeResponse) GetVirtualAttribute200Response {
	return GetVirtualAttribute200Response{
		ReverseDnJoinVirtualAttributeResponse: v,
	}
}

// ShortUniqueIdVirtualAttributeResponseAsGetVirtualAttribute200Response is a convenience function that returns ShortUniqueIdVirtualAttributeResponse wrapped in GetVirtualAttribute200Response
func ShortUniqueIdVirtualAttributeResponseAsGetVirtualAttribute200Response(v *ShortUniqueIdVirtualAttributeResponse) GetVirtualAttribute200Response {
	return GetVirtualAttribute200Response{
		ShortUniqueIdVirtualAttributeResponse: v,
	}
}

// SubschemaSubentryVirtualAttributeResponseAsGetVirtualAttribute200Response is a convenience function that returns SubschemaSubentryVirtualAttributeResponse wrapped in GetVirtualAttribute200Response
func SubschemaSubentryVirtualAttributeResponseAsGetVirtualAttribute200Response(v *SubschemaSubentryVirtualAttributeResponse) GetVirtualAttribute200Response {
	return GetVirtualAttribute200Response{
		SubschemaSubentryVirtualAttributeResponse: v,
	}
}

// ThirdPartyVirtualAttributeResponseAsGetVirtualAttribute200Response is a convenience function that returns ThirdPartyVirtualAttributeResponse wrapped in GetVirtualAttribute200Response
func ThirdPartyVirtualAttributeResponseAsGetVirtualAttribute200Response(v *ThirdPartyVirtualAttributeResponse) GetVirtualAttribute200Response {
	return GetVirtualAttribute200Response{
		ThirdPartyVirtualAttributeResponse: v,
	}
}

// UserDefinedVirtualAttributeResponseAsGetVirtualAttribute200Response is a convenience function that returns UserDefinedVirtualAttributeResponse wrapped in GetVirtualAttribute200Response
func UserDefinedVirtualAttributeResponseAsGetVirtualAttribute200Response(v *UserDefinedVirtualAttributeResponse) GetVirtualAttribute200Response {
	return GetVirtualAttribute200Response{
		UserDefinedVirtualAttributeResponse: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *GetVirtualAttribute200Response) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ConstructedVirtualAttributeResponse
	err = newStrictDecoder(data).Decode(&dst.ConstructedVirtualAttributeResponse)
	if err == nil {
		jsonConstructedVirtualAttributeResponse, _ := json.Marshal(dst.ConstructedVirtualAttributeResponse)
		if string(jsonConstructedVirtualAttributeResponse) == "{}" { // empty struct
			dst.ConstructedVirtualAttributeResponse = nil
		} else {
			match++
		}
	} else {
		dst.ConstructedVirtualAttributeResponse = nil
	}

	// try to unmarshal data into CurrentTimeVirtualAttributeResponse
	err = newStrictDecoder(data).Decode(&dst.CurrentTimeVirtualAttributeResponse)
	if err == nil {
		jsonCurrentTimeVirtualAttributeResponse, _ := json.Marshal(dst.CurrentTimeVirtualAttributeResponse)
		if string(jsonCurrentTimeVirtualAttributeResponse) == "{}" { // empty struct
			dst.CurrentTimeVirtualAttributeResponse = nil
		} else {
			match++
		}
	} else {
		dst.CurrentTimeVirtualAttributeResponse = nil
	}

	// try to unmarshal data into DnJoinVirtualAttributeResponse
	err = newStrictDecoder(data).Decode(&dst.DnJoinVirtualAttributeResponse)
	if err == nil {
		jsonDnJoinVirtualAttributeResponse, _ := json.Marshal(dst.DnJoinVirtualAttributeResponse)
		if string(jsonDnJoinVirtualAttributeResponse) == "{}" { // empty struct
			dst.DnJoinVirtualAttributeResponse = nil
		} else {
			match++
		}
	} else {
		dst.DnJoinVirtualAttributeResponse = nil
	}

	// try to unmarshal data into EntryChecksumVirtualAttributeResponse
	err = newStrictDecoder(data).Decode(&dst.EntryChecksumVirtualAttributeResponse)
	if err == nil {
		jsonEntryChecksumVirtualAttributeResponse, _ := json.Marshal(dst.EntryChecksumVirtualAttributeResponse)
		if string(jsonEntryChecksumVirtualAttributeResponse) == "{}" { // empty struct
			dst.EntryChecksumVirtualAttributeResponse = nil
		} else {
			match++
		}
	} else {
		dst.EntryChecksumVirtualAttributeResponse = nil
	}

	// try to unmarshal data into EntryDnVirtualAttributeResponse
	err = newStrictDecoder(data).Decode(&dst.EntryDnVirtualAttributeResponse)
	if err == nil {
		jsonEntryDnVirtualAttributeResponse, _ := json.Marshal(dst.EntryDnVirtualAttributeResponse)
		if string(jsonEntryDnVirtualAttributeResponse) == "{}" { // empty struct
			dst.EntryDnVirtualAttributeResponse = nil
		} else {
			match++
		}
	} else {
		dst.EntryDnVirtualAttributeResponse = nil
	}

	// try to unmarshal data into EqualityJoinVirtualAttributeResponse
	err = newStrictDecoder(data).Decode(&dst.EqualityJoinVirtualAttributeResponse)
	if err == nil {
		jsonEqualityJoinVirtualAttributeResponse, _ := json.Marshal(dst.EqualityJoinVirtualAttributeResponse)
		if string(jsonEqualityJoinVirtualAttributeResponse) == "{}" { // empty struct
			dst.EqualityJoinVirtualAttributeResponse = nil
		} else {
			match++
		}
	} else {
		dst.EqualityJoinVirtualAttributeResponse = nil
	}

	// try to unmarshal data into GroovyScriptedVirtualAttributeResponse
	err = newStrictDecoder(data).Decode(&dst.GroovyScriptedVirtualAttributeResponse)
	if err == nil {
		jsonGroovyScriptedVirtualAttributeResponse, _ := json.Marshal(dst.GroovyScriptedVirtualAttributeResponse)
		if string(jsonGroovyScriptedVirtualAttributeResponse) == "{}" { // empty struct
			dst.GroovyScriptedVirtualAttributeResponse = nil
		} else {
			match++
		}
	} else {
		dst.GroovyScriptedVirtualAttributeResponse = nil
	}

	// try to unmarshal data into HasSubordinatesVirtualAttributeResponse
	err = newStrictDecoder(data).Decode(&dst.HasSubordinatesVirtualAttributeResponse)
	if err == nil {
		jsonHasSubordinatesVirtualAttributeResponse, _ := json.Marshal(dst.HasSubordinatesVirtualAttributeResponse)
		if string(jsonHasSubordinatesVirtualAttributeResponse) == "{}" { // empty struct
			dst.HasSubordinatesVirtualAttributeResponse = nil
		} else {
			match++
		}
	} else {
		dst.HasSubordinatesVirtualAttributeResponse = nil
	}

	// try to unmarshal data into IdentifyReferencesVirtualAttributeResponse
	err = newStrictDecoder(data).Decode(&dst.IdentifyReferencesVirtualAttributeResponse)
	if err == nil {
		jsonIdentifyReferencesVirtualAttributeResponse, _ := json.Marshal(dst.IdentifyReferencesVirtualAttributeResponse)
		if string(jsonIdentifyReferencesVirtualAttributeResponse) == "{}" { // empty struct
			dst.IdentifyReferencesVirtualAttributeResponse = nil
		} else {
			match++
		}
	} else {
		dst.IdentifyReferencesVirtualAttributeResponse = nil
	}

	// try to unmarshal data into InstanceNameVirtualAttributeResponse
	err = newStrictDecoder(data).Decode(&dst.InstanceNameVirtualAttributeResponse)
	if err == nil {
		jsonInstanceNameVirtualAttributeResponse, _ := json.Marshal(dst.InstanceNameVirtualAttributeResponse)
		if string(jsonInstanceNameVirtualAttributeResponse) == "{}" { // empty struct
			dst.InstanceNameVirtualAttributeResponse = nil
		} else {
			match++
		}
	} else {
		dst.InstanceNameVirtualAttributeResponse = nil
	}

	// try to unmarshal data into IsMemberOfVirtualAttributeResponse
	err = newStrictDecoder(data).Decode(&dst.IsMemberOfVirtualAttributeResponse)
	if err == nil {
		jsonIsMemberOfVirtualAttributeResponse, _ := json.Marshal(dst.IsMemberOfVirtualAttributeResponse)
		if string(jsonIsMemberOfVirtualAttributeResponse) == "{}" { // empty struct
			dst.IsMemberOfVirtualAttributeResponse = nil
		} else {
			match++
		}
	} else {
		dst.IsMemberOfVirtualAttributeResponse = nil
	}

	// try to unmarshal data into MemberOfServerGroupVirtualAttributeResponse
	err = newStrictDecoder(data).Decode(&dst.MemberOfServerGroupVirtualAttributeResponse)
	if err == nil {
		jsonMemberOfServerGroupVirtualAttributeResponse, _ := json.Marshal(dst.MemberOfServerGroupVirtualAttributeResponse)
		if string(jsonMemberOfServerGroupVirtualAttributeResponse) == "{}" { // empty struct
			dst.MemberOfServerGroupVirtualAttributeResponse = nil
		} else {
			match++
		}
	} else {
		dst.MemberOfServerGroupVirtualAttributeResponse = nil
	}

	// try to unmarshal data into MemberVirtualAttributeResponse
	err = newStrictDecoder(data).Decode(&dst.MemberVirtualAttributeResponse)
	if err == nil {
		jsonMemberVirtualAttributeResponse, _ := json.Marshal(dst.MemberVirtualAttributeResponse)
		if string(jsonMemberVirtualAttributeResponse) == "{}" { // empty struct
			dst.MemberVirtualAttributeResponse = nil
		} else {
			match++
		}
	} else {
		dst.MemberVirtualAttributeResponse = nil
	}

	// try to unmarshal data into MirrorVirtualAttributeResponse
	err = newStrictDecoder(data).Decode(&dst.MirrorVirtualAttributeResponse)
	if err == nil {
		jsonMirrorVirtualAttributeResponse, _ := json.Marshal(dst.MirrorVirtualAttributeResponse)
		if string(jsonMirrorVirtualAttributeResponse) == "{}" { // empty struct
			dst.MirrorVirtualAttributeResponse = nil
		} else {
			match++
		}
	} else {
		dst.MirrorVirtualAttributeResponse = nil
	}

	// try to unmarshal data into NumSubordinatesVirtualAttributeResponse
	err = newStrictDecoder(data).Decode(&dst.NumSubordinatesVirtualAttributeResponse)
	if err == nil {
		jsonNumSubordinatesVirtualAttributeResponse, _ := json.Marshal(dst.NumSubordinatesVirtualAttributeResponse)
		if string(jsonNumSubordinatesVirtualAttributeResponse) == "{}" { // empty struct
			dst.NumSubordinatesVirtualAttributeResponse = nil
		} else {
			match++
		}
	} else {
		dst.NumSubordinatesVirtualAttributeResponse = nil
	}

	// try to unmarshal data into PasswordPolicyStateJsonVirtualAttributeResponse
	err = newStrictDecoder(data).Decode(&dst.PasswordPolicyStateJsonVirtualAttributeResponse)
	if err == nil {
		jsonPasswordPolicyStateJsonVirtualAttributeResponse, _ := json.Marshal(dst.PasswordPolicyStateJsonVirtualAttributeResponse)
		if string(jsonPasswordPolicyStateJsonVirtualAttributeResponse) == "{}" { // empty struct
			dst.PasswordPolicyStateJsonVirtualAttributeResponse = nil
		} else {
			match++
		}
	} else {
		dst.PasswordPolicyStateJsonVirtualAttributeResponse = nil
	}

	// try to unmarshal data into ReplicationStateDetailVirtualAttributeResponse
	err = newStrictDecoder(data).Decode(&dst.ReplicationStateDetailVirtualAttributeResponse)
	if err == nil {
		jsonReplicationStateDetailVirtualAttributeResponse, _ := json.Marshal(dst.ReplicationStateDetailVirtualAttributeResponse)
		if string(jsonReplicationStateDetailVirtualAttributeResponse) == "{}" { // empty struct
			dst.ReplicationStateDetailVirtualAttributeResponse = nil
		} else {
			match++
		}
	} else {
		dst.ReplicationStateDetailVirtualAttributeResponse = nil
	}

	// try to unmarshal data into ReverseDnJoinVirtualAttributeResponse
	err = newStrictDecoder(data).Decode(&dst.ReverseDnJoinVirtualAttributeResponse)
	if err == nil {
		jsonReverseDnJoinVirtualAttributeResponse, _ := json.Marshal(dst.ReverseDnJoinVirtualAttributeResponse)
		if string(jsonReverseDnJoinVirtualAttributeResponse) == "{}" { // empty struct
			dst.ReverseDnJoinVirtualAttributeResponse = nil
		} else {
			match++
		}
	} else {
		dst.ReverseDnJoinVirtualAttributeResponse = nil
	}

	// try to unmarshal data into ShortUniqueIdVirtualAttributeResponse
	err = newStrictDecoder(data).Decode(&dst.ShortUniqueIdVirtualAttributeResponse)
	if err == nil {
		jsonShortUniqueIdVirtualAttributeResponse, _ := json.Marshal(dst.ShortUniqueIdVirtualAttributeResponse)
		if string(jsonShortUniqueIdVirtualAttributeResponse) == "{}" { // empty struct
			dst.ShortUniqueIdVirtualAttributeResponse = nil
		} else {
			match++
		}
	} else {
		dst.ShortUniqueIdVirtualAttributeResponse = nil
	}

	// try to unmarshal data into SubschemaSubentryVirtualAttributeResponse
	err = newStrictDecoder(data).Decode(&dst.SubschemaSubentryVirtualAttributeResponse)
	if err == nil {
		jsonSubschemaSubentryVirtualAttributeResponse, _ := json.Marshal(dst.SubschemaSubentryVirtualAttributeResponse)
		if string(jsonSubschemaSubentryVirtualAttributeResponse) == "{}" { // empty struct
			dst.SubschemaSubentryVirtualAttributeResponse = nil
		} else {
			match++
		}
	} else {
		dst.SubschemaSubentryVirtualAttributeResponse = nil
	}

	// try to unmarshal data into ThirdPartyVirtualAttributeResponse
	err = newStrictDecoder(data).Decode(&dst.ThirdPartyVirtualAttributeResponse)
	if err == nil {
		jsonThirdPartyVirtualAttributeResponse, _ := json.Marshal(dst.ThirdPartyVirtualAttributeResponse)
		if string(jsonThirdPartyVirtualAttributeResponse) == "{}" { // empty struct
			dst.ThirdPartyVirtualAttributeResponse = nil
		} else {
			match++
		}
	} else {
		dst.ThirdPartyVirtualAttributeResponse = nil
	}

	// try to unmarshal data into UserDefinedVirtualAttributeResponse
	err = newStrictDecoder(data).Decode(&dst.UserDefinedVirtualAttributeResponse)
	if err == nil {
		jsonUserDefinedVirtualAttributeResponse, _ := json.Marshal(dst.UserDefinedVirtualAttributeResponse)
		if string(jsonUserDefinedVirtualAttributeResponse) == "{}" { // empty struct
			dst.UserDefinedVirtualAttributeResponse = nil
		} else {
			match++
		}
	} else {
		dst.UserDefinedVirtualAttributeResponse = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ConstructedVirtualAttributeResponse = nil
		dst.CurrentTimeVirtualAttributeResponse = nil
		dst.DnJoinVirtualAttributeResponse = nil
		dst.EntryChecksumVirtualAttributeResponse = nil
		dst.EntryDnVirtualAttributeResponse = nil
		dst.EqualityJoinVirtualAttributeResponse = nil
		dst.GroovyScriptedVirtualAttributeResponse = nil
		dst.HasSubordinatesVirtualAttributeResponse = nil
		dst.IdentifyReferencesVirtualAttributeResponse = nil
		dst.InstanceNameVirtualAttributeResponse = nil
		dst.IsMemberOfVirtualAttributeResponse = nil
		dst.MemberOfServerGroupVirtualAttributeResponse = nil
		dst.MemberVirtualAttributeResponse = nil
		dst.MirrorVirtualAttributeResponse = nil
		dst.NumSubordinatesVirtualAttributeResponse = nil
		dst.PasswordPolicyStateJsonVirtualAttributeResponse = nil
		dst.ReplicationStateDetailVirtualAttributeResponse = nil
		dst.ReverseDnJoinVirtualAttributeResponse = nil
		dst.ShortUniqueIdVirtualAttributeResponse = nil
		dst.SubschemaSubentryVirtualAttributeResponse = nil
		dst.ThirdPartyVirtualAttributeResponse = nil
		dst.UserDefinedVirtualAttributeResponse = nil

		return fmt.Errorf("data matches more than one schema in oneOf(GetVirtualAttribute200Response)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(GetVirtualAttribute200Response)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GetVirtualAttribute200Response) MarshalJSON() ([]byte, error) {
	if src.ConstructedVirtualAttributeResponse != nil {
		return json.Marshal(&src.ConstructedVirtualAttributeResponse)
	}

	if src.CurrentTimeVirtualAttributeResponse != nil {
		return json.Marshal(&src.CurrentTimeVirtualAttributeResponse)
	}

	if src.DnJoinVirtualAttributeResponse != nil {
		return json.Marshal(&src.DnJoinVirtualAttributeResponse)
	}

	if src.EntryChecksumVirtualAttributeResponse != nil {
		return json.Marshal(&src.EntryChecksumVirtualAttributeResponse)
	}

	if src.EntryDnVirtualAttributeResponse != nil {
		return json.Marshal(&src.EntryDnVirtualAttributeResponse)
	}

	if src.EqualityJoinVirtualAttributeResponse != nil {
		return json.Marshal(&src.EqualityJoinVirtualAttributeResponse)
	}

	if src.GroovyScriptedVirtualAttributeResponse != nil {
		return json.Marshal(&src.GroovyScriptedVirtualAttributeResponse)
	}

	if src.HasSubordinatesVirtualAttributeResponse != nil {
		return json.Marshal(&src.HasSubordinatesVirtualAttributeResponse)
	}

	if src.IdentifyReferencesVirtualAttributeResponse != nil {
		return json.Marshal(&src.IdentifyReferencesVirtualAttributeResponse)
	}

	if src.InstanceNameVirtualAttributeResponse != nil {
		return json.Marshal(&src.InstanceNameVirtualAttributeResponse)
	}

	if src.IsMemberOfVirtualAttributeResponse != nil {
		return json.Marshal(&src.IsMemberOfVirtualAttributeResponse)
	}

	if src.MemberOfServerGroupVirtualAttributeResponse != nil {
		return json.Marshal(&src.MemberOfServerGroupVirtualAttributeResponse)
	}

	if src.MemberVirtualAttributeResponse != nil {
		return json.Marshal(&src.MemberVirtualAttributeResponse)
	}

	if src.MirrorVirtualAttributeResponse != nil {
		return json.Marshal(&src.MirrorVirtualAttributeResponse)
	}

	if src.NumSubordinatesVirtualAttributeResponse != nil {
		return json.Marshal(&src.NumSubordinatesVirtualAttributeResponse)
	}

	if src.PasswordPolicyStateJsonVirtualAttributeResponse != nil {
		return json.Marshal(&src.PasswordPolicyStateJsonVirtualAttributeResponse)
	}

	if src.ReplicationStateDetailVirtualAttributeResponse != nil {
		return json.Marshal(&src.ReplicationStateDetailVirtualAttributeResponse)
	}

	if src.ReverseDnJoinVirtualAttributeResponse != nil {
		return json.Marshal(&src.ReverseDnJoinVirtualAttributeResponse)
	}

	if src.ShortUniqueIdVirtualAttributeResponse != nil {
		return json.Marshal(&src.ShortUniqueIdVirtualAttributeResponse)
	}

	if src.SubschemaSubentryVirtualAttributeResponse != nil {
		return json.Marshal(&src.SubschemaSubentryVirtualAttributeResponse)
	}

	if src.ThirdPartyVirtualAttributeResponse != nil {
		return json.Marshal(&src.ThirdPartyVirtualAttributeResponse)
	}

	if src.UserDefinedVirtualAttributeResponse != nil {
		return json.Marshal(&src.UserDefinedVirtualAttributeResponse)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GetVirtualAttribute200Response) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.ConstructedVirtualAttributeResponse != nil {
		return obj.ConstructedVirtualAttributeResponse
	}

	if obj.CurrentTimeVirtualAttributeResponse != nil {
		return obj.CurrentTimeVirtualAttributeResponse
	}

	if obj.DnJoinVirtualAttributeResponse != nil {
		return obj.DnJoinVirtualAttributeResponse
	}

	if obj.EntryChecksumVirtualAttributeResponse != nil {
		return obj.EntryChecksumVirtualAttributeResponse
	}

	if obj.EntryDnVirtualAttributeResponse != nil {
		return obj.EntryDnVirtualAttributeResponse
	}

	if obj.EqualityJoinVirtualAttributeResponse != nil {
		return obj.EqualityJoinVirtualAttributeResponse
	}

	if obj.GroovyScriptedVirtualAttributeResponse != nil {
		return obj.GroovyScriptedVirtualAttributeResponse
	}

	if obj.HasSubordinatesVirtualAttributeResponse != nil {
		return obj.HasSubordinatesVirtualAttributeResponse
	}

	if obj.IdentifyReferencesVirtualAttributeResponse != nil {
		return obj.IdentifyReferencesVirtualAttributeResponse
	}

	if obj.InstanceNameVirtualAttributeResponse != nil {
		return obj.InstanceNameVirtualAttributeResponse
	}

	if obj.IsMemberOfVirtualAttributeResponse != nil {
		return obj.IsMemberOfVirtualAttributeResponse
	}

	if obj.MemberOfServerGroupVirtualAttributeResponse != nil {
		return obj.MemberOfServerGroupVirtualAttributeResponse
	}

	if obj.MemberVirtualAttributeResponse != nil {
		return obj.MemberVirtualAttributeResponse
	}

	if obj.MirrorVirtualAttributeResponse != nil {
		return obj.MirrorVirtualAttributeResponse
	}

	if obj.NumSubordinatesVirtualAttributeResponse != nil {
		return obj.NumSubordinatesVirtualAttributeResponse
	}

	if obj.PasswordPolicyStateJsonVirtualAttributeResponse != nil {
		return obj.PasswordPolicyStateJsonVirtualAttributeResponse
	}

	if obj.ReplicationStateDetailVirtualAttributeResponse != nil {
		return obj.ReplicationStateDetailVirtualAttributeResponse
	}

	if obj.ReverseDnJoinVirtualAttributeResponse != nil {
		return obj.ReverseDnJoinVirtualAttributeResponse
	}

	if obj.ShortUniqueIdVirtualAttributeResponse != nil {
		return obj.ShortUniqueIdVirtualAttributeResponse
	}

	if obj.SubschemaSubentryVirtualAttributeResponse != nil {
		return obj.SubschemaSubentryVirtualAttributeResponse
	}

	if obj.ThirdPartyVirtualAttributeResponse != nil {
		return obj.ThirdPartyVirtualAttributeResponse
	}

	if obj.UserDefinedVirtualAttributeResponse != nil {
		return obj.UserDefinedVirtualAttributeResponse
	}

	// all schemas are nil
	return nil
}

type NullableGetVirtualAttribute200Response struct {
	value *GetVirtualAttribute200Response
	isSet bool
}

func (v NullableGetVirtualAttribute200Response) Get() *GetVirtualAttribute200Response {
	return v.value
}

func (v *NullableGetVirtualAttribute200Response) Set(val *GetVirtualAttribute200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetVirtualAttribute200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetVirtualAttribute200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetVirtualAttribute200Response(val *GetVirtualAttribute200Response) *NullableGetVirtualAttribute200Response {
	return &NullableGetVirtualAttribute200Response{value: val, isSet: true}
}

func (v NullableGetVirtualAttribute200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetVirtualAttribute200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
