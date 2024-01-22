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

// AddPostLdifExportTaskProcessor200Response - struct for AddPostLdifExportTaskProcessor200Response
type AddPostLdifExportTaskProcessor200Response struct {
	ThirdPartyPostLdifExportTaskProcessorResponse *ThirdPartyPostLdifExportTaskProcessorResponse
	UploadToS3PostLdifExportTaskProcessorResponse *UploadToS3PostLdifExportTaskProcessorResponse
}

// ThirdPartyPostLdifExportTaskProcessorResponseAsAddPostLdifExportTaskProcessor200Response is a convenience function that returns ThirdPartyPostLdifExportTaskProcessorResponse wrapped in AddPostLdifExportTaskProcessor200Response
func ThirdPartyPostLdifExportTaskProcessorResponseAsAddPostLdifExportTaskProcessor200Response(v *ThirdPartyPostLdifExportTaskProcessorResponse) AddPostLdifExportTaskProcessor200Response {
	return AddPostLdifExportTaskProcessor200Response{
		ThirdPartyPostLdifExportTaskProcessorResponse: v,
	}
}

// UploadToS3PostLdifExportTaskProcessorResponseAsAddPostLdifExportTaskProcessor200Response is a convenience function that returns UploadToS3PostLdifExportTaskProcessorResponse wrapped in AddPostLdifExportTaskProcessor200Response
func UploadToS3PostLdifExportTaskProcessorResponseAsAddPostLdifExportTaskProcessor200Response(v *UploadToS3PostLdifExportTaskProcessorResponse) AddPostLdifExportTaskProcessor200Response {
	return AddPostLdifExportTaskProcessor200Response{
		UploadToS3PostLdifExportTaskProcessorResponse: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *AddPostLdifExportTaskProcessor200Response) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ThirdPartyPostLdifExportTaskProcessorResponse
	err = newStrictDecoder(data).Decode(&dst.ThirdPartyPostLdifExportTaskProcessorResponse)
	if err == nil {
		jsonThirdPartyPostLdifExportTaskProcessorResponse, _ := json.Marshal(dst.ThirdPartyPostLdifExportTaskProcessorResponse)
		if string(jsonThirdPartyPostLdifExportTaskProcessorResponse) == "{}" { // empty struct
			dst.ThirdPartyPostLdifExportTaskProcessorResponse = nil
		} else {
			match++
		}
	} else {
		dst.ThirdPartyPostLdifExportTaskProcessorResponse = nil
	}

	// try to unmarshal data into UploadToS3PostLdifExportTaskProcessorResponse
	err = newStrictDecoder(data).Decode(&dst.UploadToS3PostLdifExportTaskProcessorResponse)
	if err == nil {
		jsonUploadToS3PostLdifExportTaskProcessorResponse, _ := json.Marshal(dst.UploadToS3PostLdifExportTaskProcessorResponse)
		if string(jsonUploadToS3PostLdifExportTaskProcessorResponse) == "{}" { // empty struct
			dst.UploadToS3PostLdifExportTaskProcessorResponse = nil
		} else {
			match++
		}
	} else {
		dst.UploadToS3PostLdifExportTaskProcessorResponse = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ThirdPartyPostLdifExportTaskProcessorResponse = nil
		dst.UploadToS3PostLdifExportTaskProcessorResponse = nil

		return fmt.Errorf("data matches more than one schema in oneOf(AddPostLdifExportTaskProcessor200Response)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(AddPostLdifExportTaskProcessor200Response)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src AddPostLdifExportTaskProcessor200Response) MarshalJSON() ([]byte, error) {
	if src.ThirdPartyPostLdifExportTaskProcessorResponse != nil {
		return json.Marshal(&src.ThirdPartyPostLdifExportTaskProcessorResponse)
	}

	if src.UploadToS3PostLdifExportTaskProcessorResponse != nil {
		return json.Marshal(&src.UploadToS3PostLdifExportTaskProcessorResponse)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *AddPostLdifExportTaskProcessor200Response) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.ThirdPartyPostLdifExportTaskProcessorResponse != nil {
		return obj.ThirdPartyPostLdifExportTaskProcessorResponse
	}

	if obj.UploadToS3PostLdifExportTaskProcessorResponse != nil {
		return obj.UploadToS3PostLdifExportTaskProcessorResponse
	}

	// all schemas are nil
	return nil
}

type NullableAddPostLdifExportTaskProcessor200Response struct {
	value *AddPostLdifExportTaskProcessor200Response
	isSet bool
}

func (v NullableAddPostLdifExportTaskProcessor200Response) Get() *AddPostLdifExportTaskProcessor200Response {
	return v.value
}

func (v *NullableAddPostLdifExportTaskProcessor200Response) Set(val *AddPostLdifExportTaskProcessor200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableAddPostLdifExportTaskProcessor200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableAddPostLdifExportTaskProcessor200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddPostLdifExportTaskProcessor200Response(val *AddPostLdifExportTaskProcessor200Response) *NullableAddPostLdifExportTaskProcessor200Response {
	return &NullableAddPostLdifExportTaskProcessor200Response{value: val, isSet: true}
}

func (v NullableAddPostLdifExportTaskProcessor200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddPostLdifExportTaskProcessor200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}