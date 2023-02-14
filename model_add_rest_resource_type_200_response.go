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

// AddRestResourceType200Response - struct for AddRestResourceType200Response
type AddRestResourceType200Response struct {
	GenericRestResourceTypeResponse *GenericRestResourceTypeResponse
	GroupRestResourceTypeResponse   *GroupRestResourceTypeResponse
	UserRestResourceTypeResponse    *UserRestResourceTypeResponse
}

// GenericRestResourceTypeResponseAsAddRestResourceType200Response is a convenience function that returns GenericRestResourceTypeResponse wrapped in AddRestResourceType200Response
func GenericRestResourceTypeResponseAsAddRestResourceType200Response(v *GenericRestResourceTypeResponse) AddRestResourceType200Response {
	return AddRestResourceType200Response{
		GenericRestResourceTypeResponse: v,
	}
}

// GroupRestResourceTypeResponseAsAddRestResourceType200Response is a convenience function that returns GroupRestResourceTypeResponse wrapped in AddRestResourceType200Response
func GroupRestResourceTypeResponseAsAddRestResourceType200Response(v *GroupRestResourceTypeResponse) AddRestResourceType200Response {
	return AddRestResourceType200Response{
		GroupRestResourceTypeResponse: v,
	}
}

// UserRestResourceTypeResponseAsAddRestResourceType200Response is a convenience function that returns UserRestResourceTypeResponse wrapped in AddRestResourceType200Response
func UserRestResourceTypeResponseAsAddRestResourceType200Response(v *UserRestResourceTypeResponse) AddRestResourceType200Response {
	return AddRestResourceType200Response{
		UserRestResourceTypeResponse: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *AddRestResourceType200Response) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into GenericRestResourceTypeResponse
	err = newStrictDecoder(data).Decode(&dst.GenericRestResourceTypeResponse)
	if err == nil {
		jsonGenericRestResourceTypeResponse, _ := json.Marshal(dst.GenericRestResourceTypeResponse)
		if string(jsonGenericRestResourceTypeResponse) == "{}" { // empty struct
			dst.GenericRestResourceTypeResponse = nil
		} else {
			match++
		}
	} else {
		dst.GenericRestResourceTypeResponse = nil
	}

	// try to unmarshal data into GroupRestResourceTypeResponse
	err = newStrictDecoder(data).Decode(&dst.GroupRestResourceTypeResponse)
	if err == nil {
		jsonGroupRestResourceTypeResponse, _ := json.Marshal(dst.GroupRestResourceTypeResponse)
		if string(jsonGroupRestResourceTypeResponse) == "{}" { // empty struct
			dst.GroupRestResourceTypeResponse = nil
		} else {
			match++
		}
	} else {
		dst.GroupRestResourceTypeResponse = nil
	}

	// try to unmarshal data into UserRestResourceTypeResponse
	err = newStrictDecoder(data).Decode(&dst.UserRestResourceTypeResponse)
	if err == nil {
		jsonUserRestResourceTypeResponse, _ := json.Marshal(dst.UserRestResourceTypeResponse)
		if string(jsonUserRestResourceTypeResponse) == "{}" { // empty struct
			dst.UserRestResourceTypeResponse = nil
		} else {
			match++
		}
	} else {
		dst.UserRestResourceTypeResponse = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.GenericRestResourceTypeResponse = nil
		dst.GroupRestResourceTypeResponse = nil
		dst.UserRestResourceTypeResponse = nil

		return fmt.Errorf("data matches more than one schema in oneOf(AddRestResourceType200Response)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(AddRestResourceType200Response)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src AddRestResourceType200Response) MarshalJSON() ([]byte, error) {
	if src.GenericRestResourceTypeResponse != nil {
		return json.Marshal(&src.GenericRestResourceTypeResponse)
	}

	if src.GroupRestResourceTypeResponse != nil {
		return json.Marshal(&src.GroupRestResourceTypeResponse)
	}

	if src.UserRestResourceTypeResponse != nil {
		return json.Marshal(&src.UserRestResourceTypeResponse)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *AddRestResourceType200Response) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.GenericRestResourceTypeResponse != nil {
		return obj.GenericRestResourceTypeResponse
	}

	if obj.GroupRestResourceTypeResponse != nil {
		return obj.GroupRestResourceTypeResponse
	}

	if obj.UserRestResourceTypeResponse != nil {
		return obj.UserRestResourceTypeResponse
	}

	// all schemas are nil
	return nil
}

type NullableAddRestResourceType200Response struct {
	value *AddRestResourceType200Response
	isSet bool
}

func (v NullableAddRestResourceType200Response) Get() *AddRestResourceType200Response {
	return v.value
}

func (v *NullableAddRestResourceType200Response) Set(val *AddRestResourceType200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableAddRestResourceType200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableAddRestResourceType200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddRestResourceType200Response(val *AddRestResourceType200Response) *NullableAddRestResourceType200Response {
	return &NullableAddRestResourceType200Response{value: val, isSet: true}
}

func (v NullableAddRestResourceType200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddRestResourceType200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
