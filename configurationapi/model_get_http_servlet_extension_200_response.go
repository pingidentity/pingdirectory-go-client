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

// GetHttpServletExtension200Response - struct for GetHttpServletExtension200Response
type GetHttpServletExtension200Response struct {
	AvailabilityStateHttpServletExtensionResponse *AvailabilityStateHttpServletExtensionResponse
	ConfigHttpServletExtensionResponse            *ConfigHttpServletExtensionResponse
	ConsentHttpServletExtensionResponse           *ConsentHttpServletExtensionResponse
	DelegatedAdminHttpServletExtensionResponse    *DelegatedAdminHttpServletExtensionResponse
	DirectoryRestApiHttpServletExtensionResponse  *DirectoryRestApiHttpServletExtensionResponse
	FileServerHttpServletExtensionResponse        *FileServerHttpServletExtensionResponse
	GroovyScriptedHttpServletExtensionResponse    *GroovyScriptedHttpServletExtensionResponse
	LdapMappedScimHttpServletExtensionResponse    *LdapMappedScimHttpServletExtensionResponse
	QuickstartHttpServletExtensionResponse        *QuickstartHttpServletExtensionResponse
	Scim2HttpServletExtensionResponse             *Scim2HttpServletExtensionResponse
	ThirdPartyHttpServletExtensionResponse        *ThirdPartyHttpServletExtensionResponse
	VelocityHttpServletExtensionResponse          *VelocityHttpServletExtensionResponse
}

// AvailabilityStateHttpServletExtensionResponseAsGetHttpServletExtension200Response is a convenience function that returns AvailabilityStateHttpServletExtensionResponse wrapped in GetHttpServletExtension200Response
func AvailabilityStateHttpServletExtensionResponseAsGetHttpServletExtension200Response(v *AvailabilityStateHttpServletExtensionResponse) GetHttpServletExtension200Response {
	return GetHttpServletExtension200Response{
		AvailabilityStateHttpServletExtensionResponse: v,
	}
}

// ConfigHttpServletExtensionResponseAsGetHttpServletExtension200Response is a convenience function that returns ConfigHttpServletExtensionResponse wrapped in GetHttpServletExtension200Response
func ConfigHttpServletExtensionResponseAsGetHttpServletExtension200Response(v *ConfigHttpServletExtensionResponse) GetHttpServletExtension200Response {
	return GetHttpServletExtension200Response{
		ConfigHttpServletExtensionResponse: v,
	}
}

// ConsentHttpServletExtensionResponseAsGetHttpServletExtension200Response is a convenience function that returns ConsentHttpServletExtensionResponse wrapped in GetHttpServletExtension200Response
func ConsentHttpServletExtensionResponseAsGetHttpServletExtension200Response(v *ConsentHttpServletExtensionResponse) GetHttpServletExtension200Response {
	return GetHttpServletExtension200Response{
		ConsentHttpServletExtensionResponse: v,
	}
}

// DelegatedAdminHttpServletExtensionResponseAsGetHttpServletExtension200Response is a convenience function that returns DelegatedAdminHttpServletExtensionResponse wrapped in GetHttpServletExtension200Response
func DelegatedAdminHttpServletExtensionResponseAsGetHttpServletExtension200Response(v *DelegatedAdminHttpServletExtensionResponse) GetHttpServletExtension200Response {
	return GetHttpServletExtension200Response{
		DelegatedAdminHttpServletExtensionResponse: v,
	}
}

// DirectoryRestApiHttpServletExtensionResponseAsGetHttpServletExtension200Response is a convenience function that returns DirectoryRestApiHttpServletExtensionResponse wrapped in GetHttpServletExtension200Response
func DirectoryRestApiHttpServletExtensionResponseAsGetHttpServletExtension200Response(v *DirectoryRestApiHttpServletExtensionResponse) GetHttpServletExtension200Response {
	return GetHttpServletExtension200Response{
		DirectoryRestApiHttpServletExtensionResponse: v,
	}
}

// FileServerHttpServletExtensionResponseAsGetHttpServletExtension200Response is a convenience function that returns FileServerHttpServletExtensionResponse wrapped in GetHttpServletExtension200Response
func FileServerHttpServletExtensionResponseAsGetHttpServletExtension200Response(v *FileServerHttpServletExtensionResponse) GetHttpServletExtension200Response {
	return GetHttpServletExtension200Response{
		FileServerHttpServletExtensionResponse: v,
	}
}

// GroovyScriptedHttpServletExtensionResponseAsGetHttpServletExtension200Response is a convenience function that returns GroovyScriptedHttpServletExtensionResponse wrapped in GetHttpServletExtension200Response
func GroovyScriptedHttpServletExtensionResponseAsGetHttpServletExtension200Response(v *GroovyScriptedHttpServletExtensionResponse) GetHttpServletExtension200Response {
	return GetHttpServletExtension200Response{
		GroovyScriptedHttpServletExtensionResponse: v,
	}
}

// LdapMappedScimHttpServletExtensionResponseAsGetHttpServletExtension200Response is a convenience function that returns LdapMappedScimHttpServletExtensionResponse wrapped in GetHttpServletExtension200Response
func LdapMappedScimHttpServletExtensionResponseAsGetHttpServletExtension200Response(v *LdapMappedScimHttpServletExtensionResponse) GetHttpServletExtension200Response {
	return GetHttpServletExtension200Response{
		LdapMappedScimHttpServletExtensionResponse: v,
	}
}

// QuickstartHttpServletExtensionResponseAsGetHttpServletExtension200Response is a convenience function that returns QuickstartHttpServletExtensionResponse wrapped in GetHttpServletExtension200Response
func QuickstartHttpServletExtensionResponseAsGetHttpServletExtension200Response(v *QuickstartHttpServletExtensionResponse) GetHttpServletExtension200Response {
	return GetHttpServletExtension200Response{
		QuickstartHttpServletExtensionResponse: v,
	}
}

// Scim2HttpServletExtensionResponseAsGetHttpServletExtension200Response is a convenience function that returns Scim2HttpServletExtensionResponse wrapped in GetHttpServletExtension200Response
func Scim2HttpServletExtensionResponseAsGetHttpServletExtension200Response(v *Scim2HttpServletExtensionResponse) GetHttpServletExtension200Response {
	return GetHttpServletExtension200Response{
		Scim2HttpServletExtensionResponse: v,
	}
}

// ThirdPartyHttpServletExtensionResponseAsGetHttpServletExtension200Response is a convenience function that returns ThirdPartyHttpServletExtensionResponse wrapped in GetHttpServletExtension200Response
func ThirdPartyHttpServletExtensionResponseAsGetHttpServletExtension200Response(v *ThirdPartyHttpServletExtensionResponse) GetHttpServletExtension200Response {
	return GetHttpServletExtension200Response{
		ThirdPartyHttpServletExtensionResponse: v,
	}
}

// VelocityHttpServletExtensionResponseAsGetHttpServletExtension200Response is a convenience function that returns VelocityHttpServletExtensionResponse wrapped in GetHttpServletExtension200Response
func VelocityHttpServletExtensionResponseAsGetHttpServletExtension200Response(v *VelocityHttpServletExtensionResponse) GetHttpServletExtension200Response {
	return GetHttpServletExtension200Response{
		VelocityHttpServletExtensionResponse: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *GetHttpServletExtension200Response) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AvailabilityStateHttpServletExtensionResponse
	err = newStrictDecoder(data).Decode(&dst.AvailabilityStateHttpServletExtensionResponse)
	if err == nil {
		jsonAvailabilityStateHttpServletExtensionResponse, _ := json.Marshal(dst.AvailabilityStateHttpServletExtensionResponse)
		if string(jsonAvailabilityStateHttpServletExtensionResponse) == "{}" { // empty struct
			dst.AvailabilityStateHttpServletExtensionResponse = nil
		} else {
			match++
		}
	} else {
		dst.AvailabilityStateHttpServletExtensionResponse = nil
	}

	// try to unmarshal data into ConfigHttpServletExtensionResponse
	err = newStrictDecoder(data).Decode(&dst.ConfigHttpServletExtensionResponse)
	if err == nil {
		jsonConfigHttpServletExtensionResponse, _ := json.Marshal(dst.ConfigHttpServletExtensionResponse)
		if string(jsonConfigHttpServletExtensionResponse) == "{}" { // empty struct
			dst.ConfigHttpServletExtensionResponse = nil
		} else {
			match++
		}
	} else {
		dst.ConfigHttpServletExtensionResponse = nil
	}

	// try to unmarshal data into ConsentHttpServletExtensionResponse
	err = newStrictDecoder(data).Decode(&dst.ConsentHttpServletExtensionResponse)
	if err == nil {
		jsonConsentHttpServletExtensionResponse, _ := json.Marshal(dst.ConsentHttpServletExtensionResponse)
		if string(jsonConsentHttpServletExtensionResponse) == "{}" { // empty struct
			dst.ConsentHttpServletExtensionResponse = nil
		} else {
			match++
		}
	} else {
		dst.ConsentHttpServletExtensionResponse = nil
	}

	// try to unmarshal data into DelegatedAdminHttpServletExtensionResponse
	err = newStrictDecoder(data).Decode(&dst.DelegatedAdminHttpServletExtensionResponse)
	if err == nil {
		jsonDelegatedAdminHttpServletExtensionResponse, _ := json.Marshal(dst.DelegatedAdminHttpServletExtensionResponse)
		if string(jsonDelegatedAdminHttpServletExtensionResponse) == "{}" { // empty struct
			dst.DelegatedAdminHttpServletExtensionResponse = nil
		} else {
			match++
		}
	} else {
		dst.DelegatedAdminHttpServletExtensionResponse = nil
	}

	// try to unmarshal data into DirectoryRestApiHttpServletExtensionResponse
	err = newStrictDecoder(data).Decode(&dst.DirectoryRestApiHttpServletExtensionResponse)
	if err == nil {
		jsonDirectoryRestApiHttpServletExtensionResponse, _ := json.Marshal(dst.DirectoryRestApiHttpServletExtensionResponse)
		if string(jsonDirectoryRestApiHttpServletExtensionResponse) == "{}" { // empty struct
			dst.DirectoryRestApiHttpServletExtensionResponse = nil
		} else {
			match++
		}
	} else {
		dst.DirectoryRestApiHttpServletExtensionResponse = nil
	}

	// try to unmarshal data into FileServerHttpServletExtensionResponse
	err = newStrictDecoder(data).Decode(&dst.FileServerHttpServletExtensionResponse)
	if err == nil {
		jsonFileServerHttpServletExtensionResponse, _ := json.Marshal(dst.FileServerHttpServletExtensionResponse)
		if string(jsonFileServerHttpServletExtensionResponse) == "{}" { // empty struct
			dst.FileServerHttpServletExtensionResponse = nil
		} else {
			match++
		}
	} else {
		dst.FileServerHttpServletExtensionResponse = nil
	}

	// try to unmarshal data into GroovyScriptedHttpServletExtensionResponse
	err = newStrictDecoder(data).Decode(&dst.GroovyScriptedHttpServletExtensionResponse)
	if err == nil {
		jsonGroovyScriptedHttpServletExtensionResponse, _ := json.Marshal(dst.GroovyScriptedHttpServletExtensionResponse)
		if string(jsonGroovyScriptedHttpServletExtensionResponse) == "{}" { // empty struct
			dst.GroovyScriptedHttpServletExtensionResponse = nil
		} else {
			match++
		}
	} else {
		dst.GroovyScriptedHttpServletExtensionResponse = nil
	}

	// try to unmarshal data into LdapMappedScimHttpServletExtensionResponse
	err = newStrictDecoder(data).Decode(&dst.LdapMappedScimHttpServletExtensionResponse)
	if err == nil {
		jsonLdapMappedScimHttpServletExtensionResponse, _ := json.Marshal(dst.LdapMappedScimHttpServletExtensionResponse)
		if string(jsonLdapMappedScimHttpServletExtensionResponse) == "{}" { // empty struct
			dst.LdapMappedScimHttpServletExtensionResponse = nil
		} else {
			match++
		}
	} else {
		dst.LdapMappedScimHttpServletExtensionResponse = nil
	}

	// try to unmarshal data into QuickstartHttpServletExtensionResponse
	err = newStrictDecoder(data).Decode(&dst.QuickstartHttpServletExtensionResponse)
	if err == nil {
		jsonQuickstartHttpServletExtensionResponse, _ := json.Marshal(dst.QuickstartHttpServletExtensionResponse)
		if string(jsonQuickstartHttpServletExtensionResponse) == "{}" { // empty struct
			dst.QuickstartHttpServletExtensionResponse = nil
		} else {
			match++
		}
	} else {
		dst.QuickstartHttpServletExtensionResponse = nil
	}

	// try to unmarshal data into Scim2HttpServletExtensionResponse
	err = newStrictDecoder(data).Decode(&dst.Scim2HttpServletExtensionResponse)
	if err == nil {
		jsonScim2HttpServletExtensionResponse, _ := json.Marshal(dst.Scim2HttpServletExtensionResponse)
		if string(jsonScim2HttpServletExtensionResponse) == "{}" { // empty struct
			dst.Scim2HttpServletExtensionResponse = nil
		} else {
			match++
		}
	} else {
		dst.Scim2HttpServletExtensionResponse = nil
	}

	// try to unmarshal data into ThirdPartyHttpServletExtensionResponse
	err = newStrictDecoder(data).Decode(&dst.ThirdPartyHttpServletExtensionResponse)
	if err == nil {
		jsonThirdPartyHttpServletExtensionResponse, _ := json.Marshal(dst.ThirdPartyHttpServletExtensionResponse)
		if string(jsonThirdPartyHttpServletExtensionResponse) == "{}" { // empty struct
			dst.ThirdPartyHttpServletExtensionResponse = nil
		} else {
			match++
		}
	} else {
		dst.ThirdPartyHttpServletExtensionResponse = nil
	}

	// try to unmarshal data into VelocityHttpServletExtensionResponse
	err = newStrictDecoder(data).Decode(&dst.VelocityHttpServletExtensionResponse)
	if err == nil {
		jsonVelocityHttpServletExtensionResponse, _ := json.Marshal(dst.VelocityHttpServletExtensionResponse)
		if string(jsonVelocityHttpServletExtensionResponse) == "{}" { // empty struct
			dst.VelocityHttpServletExtensionResponse = nil
		} else {
			match++
		}
	} else {
		dst.VelocityHttpServletExtensionResponse = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.AvailabilityStateHttpServletExtensionResponse = nil
		dst.ConfigHttpServletExtensionResponse = nil
		dst.ConsentHttpServletExtensionResponse = nil
		dst.DelegatedAdminHttpServletExtensionResponse = nil
		dst.DirectoryRestApiHttpServletExtensionResponse = nil
		dst.FileServerHttpServletExtensionResponse = nil
		dst.GroovyScriptedHttpServletExtensionResponse = nil
		dst.LdapMappedScimHttpServletExtensionResponse = nil
		dst.QuickstartHttpServletExtensionResponse = nil
		dst.Scim2HttpServletExtensionResponse = nil
		dst.ThirdPartyHttpServletExtensionResponse = nil
		dst.VelocityHttpServletExtensionResponse = nil

		return fmt.Errorf("data matches more than one schema in oneOf(GetHttpServletExtension200Response)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(GetHttpServletExtension200Response)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GetHttpServletExtension200Response) MarshalJSON() ([]byte, error) {
	if src.AvailabilityStateHttpServletExtensionResponse != nil {
		return json.Marshal(&src.AvailabilityStateHttpServletExtensionResponse)
	}

	if src.ConfigHttpServletExtensionResponse != nil {
		return json.Marshal(&src.ConfigHttpServletExtensionResponse)
	}

	if src.ConsentHttpServletExtensionResponse != nil {
		return json.Marshal(&src.ConsentHttpServletExtensionResponse)
	}

	if src.DelegatedAdminHttpServletExtensionResponse != nil {
		return json.Marshal(&src.DelegatedAdminHttpServletExtensionResponse)
	}

	if src.DirectoryRestApiHttpServletExtensionResponse != nil {
		return json.Marshal(&src.DirectoryRestApiHttpServletExtensionResponse)
	}

	if src.FileServerHttpServletExtensionResponse != nil {
		return json.Marshal(&src.FileServerHttpServletExtensionResponse)
	}

	if src.GroovyScriptedHttpServletExtensionResponse != nil {
		return json.Marshal(&src.GroovyScriptedHttpServletExtensionResponse)
	}

	if src.LdapMappedScimHttpServletExtensionResponse != nil {
		return json.Marshal(&src.LdapMappedScimHttpServletExtensionResponse)
	}

	if src.QuickstartHttpServletExtensionResponse != nil {
		return json.Marshal(&src.QuickstartHttpServletExtensionResponse)
	}

	if src.Scim2HttpServletExtensionResponse != nil {
		return json.Marshal(&src.Scim2HttpServletExtensionResponse)
	}

	if src.ThirdPartyHttpServletExtensionResponse != nil {
		return json.Marshal(&src.ThirdPartyHttpServletExtensionResponse)
	}

	if src.VelocityHttpServletExtensionResponse != nil {
		return json.Marshal(&src.VelocityHttpServletExtensionResponse)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GetHttpServletExtension200Response) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.AvailabilityStateHttpServletExtensionResponse != nil {
		return obj.AvailabilityStateHttpServletExtensionResponse
	}

	if obj.ConfigHttpServletExtensionResponse != nil {
		return obj.ConfigHttpServletExtensionResponse
	}

	if obj.ConsentHttpServletExtensionResponse != nil {
		return obj.ConsentHttpServletExtensionResponse
	}

	if obj.DelegatedAdminHttpServletExtensionResponse != nil {
		return obj.DelegatedAdminHttpServletExtensionResponse
	}

	if obj.DirectoryRestApiHttpServletExtensionResponse != nil {
		return obj.DirectoryRestApiHttpServletExtensionResponse
	}

	if obj.FileServerHttpServletExtensionResponse != nil {
		return obj.FileServerHttpServletExtensionResponse
	}

	if obj.GroovyScriptedHttpServletExtensionResponse != nil {
		return obj.GroovyScriptedHttpServletExtensionResponse
	}

	if obj.LdapMappedScimHttpServletExtensionResponse != nil {
		return obj.LdapMappedScimHttpServletExtensionResponse
	}

	if obj.QuickstartHttpServletExtensionResponse != nil {
		return obj.QuickstartHttpServletExtensionResponse
	}

	if obj.Scim2HttpServletExtensionResponse != nil {
		return obj.Scim2HttpServletExtensionResponse
	}

	if obj.ThirdPartyHttpServletExtensionResponse != nil {
		return obj.ThirdPartyHttpServletExtensionResponse
	}

	if obj.VelocityHttpServletExtensionResponse != nil {
		return obj.VelocityHttpServletExtensionResponse
	}

	// all schemas are nil
	return nil
}

type NullableGetHttpServletExtension200Response struct {
	value *GetHttpServletExtension200Response
	isSet bool
}

func (v NullableGetHttpServletExtension200Response) Get() *GetHttpServletExtension200Response {
	return v.value
}

func (v *NullableGetHttpServletExtension200Response) Set(val *GetHttpServletExtension200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetHttpServletExtension200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetHttpServletExtension200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetHttpServletExtension200Response(val *GetHttpServletExtension200Response) *NullableGetHttpServletExtension200Response {
	return &NullableGetHttpServletExtension200Response{value: val, isSet: true}
}

func (v NullableGetHttpServletExtension200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetHttpServletExtension200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}