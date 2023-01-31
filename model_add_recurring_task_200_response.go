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

// AddRecurringTask200Response - struct for AddRecurringTask200Response
type AddRecurringTask200Response struct {
	BackupRecurringTaskResponse                *BackupRecurringTaskResponse
	CollectSupportDataRecurringTaskResponse    *CollectSupportDataRecurringTaskResponse
	DelayRecurringTaskResponse                 *DelayRecurringTaskResponse
	EnterLockdownModeRecurringTaskResponse     *EnterLockdownModeRecurringTaskResponse
	ExecRecurringTaskResponse                  *ExecRecurringTaskResponse
	FileRetentionRecurringTaskResponse         *FileRetentionRecurringTaskResponse
	GenerateServerProfileRecurringTaskResponse *GenerateServerProfileRecurringTaskResponse
	LdifExportRecurringTaskResponse            *LdifExportRecurringTaskResponse
	LeaveLockdownModeRecurringTaskResponse     *LeaveLockdownModeRecurringTaskResponse
	StaticallyDefinedRecurringTaskResponse     *StaticallyDefinedRecurringTaskResponse
	ThirdPartyRecurringTaskResponse            *ThirdPartyRecurringTaskResponse
}

// BackupRecurringTaskResponseAsAddRecurringTask200Response is a convenience function that returns BackupRecurringTaskResponse wrapped in AddRecurringTask200Response
func BackupRecurringTaskResponseAsAddRecurringTask200Response(v *BackupRecurringTaskResponse) AddRecurringTask200Response {
	return AddRecurringTask200Response{
		BackupRecurringTaskResponse: v,
	}
}

// CollectSupportDataRecurringTaskResponseAsAddRecurringTask200Response is a convenience function that returns CollectSupportDataRecurringTaskResponse wrapped in AddRecurringTask200Response
func CollectSupportDataRecurringTaskResponseAsAddRecurringTask200Response(v *CollectSupportDataRecurringTaskResponse) AddRecurringTask200Response {
	return AddRecurringTask200Response{
		CollectSupportDataRecurringTaskResponse: v,
	}
}

// DelayRecurringTaskResponseAsAddRecurringTask200Response is a convenience function that returns DelayRecurringTaskResponse wrapped in AddRecurringTask200Response
func DelayRecurringTaskResponseAsAddRecurringTask200Response(v *DelayRecurringTaskResponse) AddRecurringTask200Response {
	return AddRecurringTask200Response{
		DelayRecurringTaskResponse: v,
	}
}

// EnterLockdownModeRecurringTaskResponseAsAddRecurringTask200Response is a convenience function that returns EnterLockdownModeRecurringTaskResponse wrapped in AddRecurringTask200Response
func EnterLockdownModeRecurringTaskResponseAsAddRecurringTask200Response(v *EnterLockdownModeRecurringTaskResponse) AddRecurringTask200Response {
	return AddRecurringTask200Response{
		EnterLockdownModeRecurringTaskResponse: v,
	}
}

// ExecRecurringTaskResponseAsAddRecurringTask200Response is a convenience function that returns ExecRecurringTaskResponse wrapped in AddRecurringTask200Response
func ExecRecurringTaskResponseAsAddRecurringTask200Response(v *ExecRecurringTaskResponse) AddRecurringTask200Response {
	return AddRecurringTask200Response{
		ExecRecurringTaskResponse: v,
	}
}

// FileRetentionRecurringTaskResponseAsAddRecurringTask200Response is a convenience function that returns FileRetentionRecurringTaskResponse wrapped in AddRecurringTask200Response
func FileRetentionRecurringTaskResponseAsAddRecurringTask200Response(v *FileRetentionRecurringTaskResponse) AddRecurringTask200Response {
	return AddRecurringTask200Response{
		FileRetentionRecurringTaskResponse: v,
	}
}

// GenerateServerProfileRecurringTaskResponseAsAddRecurringTask200Response is a convenience function that returns GenerateServerProfileRecurringTaskResponse wrapped in AddRecurringTask200Response
func GenerateServerProfileRecurringTaskResponseAsAddRecurringTask200Response(v *GenerateServerProfileRecurringTaskResponse) AddRecurringTask200Response {
	return AddRecurringTask200Response{
		GenerateServerProfileRecurringTaskResponse: v,
	}
}

// LdifExportRecurringTaskResponseAsAddRecurringTask200Response is a convenience function that returns LdifExportRecurringTaskResponse wrapped in AddRecurringTask200Response
func LdifExportRecurringTaskResponseAsAddRecurringTask200Response(v *LdifExportRecurringTaskResponse) AddRecurringTask200Response {
	return AddRecurringTask200Response{
		LdifExportRecurringTaskResponse: v,
	}
}

// LeaveLockdownModeRecurringTaskResponseAsAddRecurringTask200Response is a convenience function that returns LeaveLockdownModeRecurringTaskResponse wrapped in AddRecurringTask200Response
func LeaveLockdownModeRecurringTaskResponseAsAddRecurringTask200Response(v *LeaveLockdownModeRecurringTaskResponse) AddRecurringTask200Response {
	return AddRecurringTask200Response{
		LeaveLockdownModeRecurringTaskResponse: v,
	}
}

// StaticallyDefinedRecurringTaskResponseAsAddRecurringTask200Response is a convenience function that returns StaticallyDefinedRecurringTaskResponse wrapped in AddRecurringTask200Response
func StaticallyDefinedRecurringTaskResponseAsAddRecurringTask200Response(v *StaticallyDefinedRecurringTaskResponse) AddRecurringTask200Response {
	return AddRecurringTask200Response{
		StaticallyDefinedRecurringTaskResponse: v,
	}
}

// ThirdPartyRecurringTaskResponseAsAddRecurringTask200Response is a convenience function that returns ThirdPartyRecurringTaskResponse wrapped in AddRecurringTask200Response
func ThirdPartyRecurringTaskResponseAsAddRecurringTask200Response(v *ThirdPartyRecurringTaskResponse) AddRecurringTask200Response {
	return AddRecurringTask200Response{
		ThirdPartyRecurringTaskResponse: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *AddRecurringTask200Response) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into BackupRecurringTaskResponse
	err = newStrictDecoder(data).Decode(&dst.BackupRecurringTaskResponse)
	if err == nil {
		jsonBackupRecurringTaskResponse, _ := json.Marshal(dst.BackupRecurringTaskResponse)
		if string(jsonBackupRecurringTaskResponse) == "{}" { // empty struct
			dst.BackupRecurringTaskResponse = nil
		} else {
			match++
		}
	} else {
		dst.BackupRecurringTaskResponse = nil
	}

	// try to unmarshal data into CollectSupportDataRecurringTaskResponse
	err = newStrictDecoder(data).Decode(&dst.CollectSupportDataRecurringTaskResponse)
	if err == nil {
		jsonCollectSupportDataRecurringTaskResponse, _ := json.Marshal(dst.CollectSupportDataRecurringTaskResponse)
		if string(jsonCollectSupportDataRecurringTaskResponse) == "{}" { // empty struct
			dst.CollectSupportDataRecurringTaskResponse = nil
		} else {
			match++
		}
	} else {
		dst.CollectSupportDataRecurringTaskResponse = nil
	}

	// try to unmarshal data into DelayRecurringTaskResponse
	err = newStrictDecoder(data).Decode(&dst.DelayRecurringTaskResponse)
	if err == nil {
		jsonDelayRecurringTaskResponse, _ := json.Marshal(dst.DelayRecurringTaskResponse)
		if string(jsonDelayRecurringTaskResponse) == "{}" { // empty struct
			dst.DelayRecurringTaskResponse = nil
		} else {
			match++
		}
	} else {
		dst.DelayRecurringTaskResponse = nil
	}

	// try to unmarshal data into EnterLockdownModeRecurringTaskResponse
	err = newStrictDecoder(data).Decode(&dst.EnterLockdownModeRecurringTaskResponse)
	if err == nil {
		jsonEnterLockdownModeRecurringTaskResponse, _ := json.Marshal(dst.EnterLockdownModeRecurringTaskResponse)
		if string(jsonEnterLockdownModeRecurringTaskResponse) == "{}" { // empty struct
			dst.EnterLockdownModeRecurringTaskResponse = nil
		} else {
			match++
		}
	} else {
		dst.EnterLockdownModeRecurringTaskResponse = nil
	}

	// try to unmarshal data into ExecRecurringTaskResponse
	err = newStrictDecoder(data).Decode(&dst.ExecRecurringTaskResponse)
	if err == nil {
		jsonExecRecurringTaskResponse, _ := json.Marshal(dst.ExecRecurringTaskResponse)
		if string(jsonExecRecurringTaskResponse) == "{}" { // empty struct
			dst.ExecRecurringTaskResponse = nil
		} else {
			match++
		}
	} else {
		dst.ExecRecurringTaskResponse = nil
	}

	// try to unmarshal data into FileRetentionRecurringTaskResponse
	err = newStrictDecoder(data).Decode(&dst.FileRetentionRecurringTaskResponse)
	if err == nil {
		jsonFileRetentionRecurringTaskResponse, _ := json.Marshal(dst.FileRetentionRecurringTaskResponse)
		if string(jsonFileRetentionRecurringTaskResponse) == "{}" { // empty struct
			dst.FileRetentionRecurringTaskResponse = nil
		} else {
			match++
		}
	} else {
		dst.FileRetentionRecurringTaskResponse = nil
	}

	// try to unmarshal data into GenerateServerProfileRecurringTaskResponse
	err = newStrictDecoder(data).Decode(&dst.GenerateServerProfileRecurringTaskResponse)
	if err == nil {
		jsonGenerateServerProfileRecurringTaskResponse, _ := json.Marshal(dst.GenerateServerProfileRecurringTaskResponse)
		if string(jsonGenerateServerProfileRecurringTaskResponse) == "{}" { // empty struct
			dst.GenerateServerProfileRecurringTaskResponse = nil
		} else {
			match++
		}
	} else {
		dst.GenerateServerProfileRecurringTaskResponse = nil
	}

	// try to unmarshal data into LdifExportRecurringTaskResponse
	err = newStrictDecoder(data).Decode(&dst.LdifExportRecurringTaskResponse)
	if err == nil {
		jsonLdifExportRecurringTaskResponse, _ := json.Marshal(dst.LdifExportRecurringTaskResponse)
		if string(jsonLdifExportRecurringTaskResponse) == "{}" { // empty struct
			dst.LdifExportRecurringTaskResponse = nil
		} else {
			match++
		}
	} else {
		dst.LdifExportRecurringTaskResponse = nil
	}

	// try to unmarshal data into LeaveLockdownModeRecurringTaskResponse
	err = newStrictDecoder(data).Decode(&dst.LeaveLockdownModeRecurringTaskResponse)
	if err == nil {
		jsonLeaveLockdownModeRecurringTaskResponse, _ := json.Marshal(dst.LeaveLockdownModeRecurringTaskResponse)
		if string(jsonLeaveLockdownModeRecurringTaskResponse) == "{}" { // empty struct
			dst.LeaveLockdownModeRecurringTaskResponse = nil
		} else {
			match++
		}
	} else {
		dst.LeaveLockdownModeRecurringTaskResponse = nil
	}

	// try to unmarshal data into StaticallyDefinedRecurringTaskResponse
	err = newStrictDecoder(data).Decode(&dst.StaticallyDefinedRecurringTaskResponse)
	if err == nil {
		jsonStaticallyDefinedRecurringTaskResponse, _ := json.Marshal(dst.StaticallyDefinedRecurringTaskResponse)
		if string(jsonStaticallyDefinedRecurringTaskResponse) == "{}" { // empty struct
			dst.StaticallyDefinedRecurringTaskResponse = nil
		} else {
			match++
		}
	} else {
		dst.StaticallyDefinedRecurringTaskResponse = nil
	}

	// try to unmarshal data into ThirdPartyRecurringTaskResponse
	err = newStrictDecoder(data).Decode(&dst.ThirdPartyRecurringTaskResponse)
	if err == nil {
		jsonThirdPartyRecurringTaskResponse, _ := json.Marshal(dst.ThirdPartyRecurringTaskResponse)
		if string(jsonThirdPartyRecurringTaskResponse) == "{}" { // empty struct
			dst.ThirdPartyRecurringTaskResponse = nil
		} else {
			match++
		}
	} else {
		dst.ThirdPartyRecurringTaskResponse = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.BackupRecurringTaskResponse = nil
		dst.CollectSupportDataRecurringTaskResponse = nil
		dst.DelayRecurringTaskResponse = nil
		dst.EnterLockdownModeRecurringTaskResponse = nil
		dst.ExecRecurringTaskResponse = nil
		dst.FileRetentionRecurringTaskResponse = nil
		dst.GenerateServerProfileRecurringTaskResponse = nil
		dst.LdifExportRecurringTaskResponse = nil
		dst.LeaveLockdownModeRecurringTaskResponse = nil
		dst.StaticallyDefinedRecurringTaskResponse = nil
		dst.ThirdPartyRecurringTaskResponse = nil

		return fmt.Errorf("data matches more than one schema in oneOf(AddRecurringTask200Response)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(AddRecurringTask200Response)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src AddRecurringTask200Response) MarshalJSON() ([]byte, error) {
	if src.BackupRecurringTaskResponse != nil {
		return json.Marshal(&src.BackupRecurringTaskResponse)
	}

	if src.CollectSupportDataRecurringTaskResponse != nil {
		return json.Marshal(&src.CollectSupportDataRecurringTaskResponse)
	}

	if src.DelayRecurringTaskResponse != nil {
		return json.Marshal(&src.DelayRecurringTaskResponse)
	}

	if src.EnterLockdownModeRecurringTaskResponse != nil {
		return json.Marshal(&src.EnterLockdownModeRecurringTaskResponse)
	}

	if src.ExecRecurringTaskResponse != nil {
		return json.Marshal(&src.ExecRecurringTaskResponse)
	}

	if src.FileRetentionRecurringTaskResponse != nil {
		return json.Marshal(&src.FileRetentionRecurringTaskResponse)
	}

	if src.GenerateServerProfileRecurringTaskResponse != nil {
		return json.Marshal(&src.GenerateServerProfileRecurringTaskResponse)
	}

	if src.LdifExportRecurringTaskResponse != nil {
		return json.Marshal(&src.LdifExportRecurringTaskResponse)
	}

	if src.LeaveLockdownModeRecurringTaskResponse != nil {
		return json.Marshal(&src.LeaveLockdownModeRecurringTaskResponse)
	}

	if src.StaticallyDefinedRecurringTaskResponse != nil {
		return json.Marshal(&src.StaticallyDefinedRecurringTaskResponse)
	}

	if src.ThirdPartyRecurringTaskResponse != nil {
		return json.Marshal(&src.ThirdPartyRecurringTaskResponse)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *AddRecurringTask200Response) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.BackupRecurringTaskResponse != nil {
		return obj.BackupRecurringTaskResponse
	}

	if obj.CollectSupportDataRecurringTaskResponse != nil {
		return obj.CollectSupportDataRecurringTaskResponse
	}

	if obj.DelayRecurringTaskResponse != nil {
		return obj.DelayRecurringTaskResponse
	}

	if obj.EnterLockdownModeRecurringTaskResponse != nil {
		return obj.EnterLockdownModeRecurringTaskResponse
	}

	if obj.ExecRecurringTaskResponse != nil {
		return obj.ExecRecurringTaskResponse
	}

	if obj.FileRetentionRecurringTaskResponse != nil {
		return obj.FileRetentionRecurringTaskResponse
	}

	if obj.GenerateServerProfileRecurringTaskResponse != nil {
		return obj.GenerateServerProfileRecurringTaskResponse
	}

	if obj.LdifExportRecurringTaskResponse != nil {
		return obj.LdifExportRecurringTaskResponse
	}

	if obj.LeaveLockdownModeRecurringTaskResponse != nil {
		return obj.LeaveLockdownModeRecurringTaskResponse
	}

	if obj.StaticallyDefinedRecurringTaskResponse != nil {
		return obj.StaticallyDefinedRecurringTaskResponse
	}

	if obj.ThirdPartyRecurringTaskResponse != nil {
		return obj.ThirdPartyRecurringTaskResponse
	}

	// all schemas are nil
	return nil
}

type NullableAddRecurringTask200Response struct {
	value *AddRecurringTask200Response
	isSet bool
}

func (v NullableAddRecurringTask200Response) Get() *AddRecurringTask200Response {
	return v.value
}

func (v *NullableAddRecurringTask200Response) Set(val *AddRecurringTask200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableAddRecurringTask200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableAddRecurringTask200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddRecurringTask200Response(val *AddRecurringTask200Response) *NullableAddRecurringTask200Response {
	return &NullableAddRecurringTask200Response{value: val, isSet: true}
}

func (v NullableAddRecurringTask200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddRecurringTask200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
