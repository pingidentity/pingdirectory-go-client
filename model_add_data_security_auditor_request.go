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

// AddDataSecurityAuditorRequest - struct for AddDataSecurityAuditorRequest
type AddDataSecurityAuditorRequest struct {
	AddAccessControlDataSecurityAuditorRequest         *AddAccessControlDataSecurityAuditorRequest
	AddDisabledAccountDataSecurityAuditorRequest       *AddDisabledAccountDataSecurityAuditorRequest
	AddExpiredPasswordDataSecurityAuditorRequest       *AddExpiredPasswordDataSecurityAuditorRequest
	AddLockedAccountDataSecurityAuditorRequest         *AddLockedAccountDataSecurityAuditorRequest
	AddMultiplePasswordDataSecurityAuditorRequest      *AddMultiplePasswordDataSecurityAuditorRequest
	AddPrivilegeDataSecurityAuditorRequest             *AddPrivilegeDataSecurityAuditorRequest
	AddWeaklyEncodedPasswordDataSecurityAuditorRequest *AddWeaklyEncodedPasswordDataSecurityAuditorRequest
}

// AddAccessControlDataSecurityAuditorRequestAsAddDataSecurityAuditorRequest is a convenience function that returns AddAccessControlDataSecurityAuditorRequest wrapped in AddDataSecurityAuditorRequest
func AddAccessControlDataSecurityAuditorRequestAsAddDataSecurityAuditorRequest(v *AddAccessControlDataSecurityAuditorRequest) AddDataSecurityAuditorRequest {
	return AddDataSecurityAuditorRequest{
		AddAccessControlDataSecurityAuditorRequest: v,
	}
}

// AddDisabledAccountDataSecurityAuditorRequestAsAddDataSecurityAuditorRequest is a convenience function that returns AddDisabledAccountDataSecurityAuditorRequest wrapped in AddDataSecurityAuditorRequest
func AddDisabledAccountDataSecurityAuditorRequestAsAddDataSecurityAuditorRequest(v *AddDisabledAccountDataSecurityAuditorRequest) AddDataSecurityAuditorRequest {
	return AddDataSecurityAuditorRequest{
		AddDisabledAccountDataSecurityAuditorRequest: v,
	}
}

// AddExpiredPasswordDataSecurityAuditorRequestAsAddDataSecurityAuditorRequest is a convenience function that returns AddExpiredPasswordDataSecurityAuditorRequest wrapped in AddDataSecurityAuditorRequest
func AddExpiredPasswordDataSecurityAuditorRequestAsAddDataSecurityAuditorRequest(v *AddExpiredPasswordDataSecurityAuditorRequest) AddDataSecurityAuditorRequest {
	return AddDataSecurityAuditorRequest{
		AddExpiredPasswordDataSecurityAuditorRequest: v,
	}
}

// AddLockedAccountDataSecurityAuditorRequestAsAddDataSecurityAuditorRequest is a convenience function that returns AddLockedAccountDataSecurityAuditorRequest wrapped in AddDataSecurityAuditorRequest
func AddLockedAccountDataSecurityAuditorRequestAsAddDataSecurityAuditorRequest(v *AddLockedAccountDataSecurityAuditorRequest) AddDataSecurityAuditorRequest {
	return AddDataSecurityAuditorRequest{
		AddLockedAccountDataSecurityAuditorRequest: v,
	}
}

// AddMultiplePasswordDataSecurityAuditorRequestAsAddDataSecurityAuditorRequest is a convenience function that returns AddMultiplePasswordDataSecurityAuditorRequest wrapped in AddDataSecurityAuditorRequest
func AddMultiplePasswordDataSecurityAuditorRequestAsAddDataSecurityAuditorRequest(v *AddMultiplePasswordDataSecurityAuditorRequest) AddDataSecurityAuditorRequest {
	return AddDataSecurityAuditorRequest{
		AddMultiplePasswordDataSecurityAuditorRequest: v,
	}
}

// AddPrivilegeDataSecurityAuditorRequestAsAddDataSecurityAuditorRequest is a convenience function that returns AddPrivilegeDataSecurityAuditorRequest wrapped in AddDataSecurityAuditorRequest
func AddPrivilegeDataSecurityAuditorRequestAsAddDataSecurityAuditorRequest(v *AddPrivilegeDataSecurityAuditorRequest) AddDataSecurityAuditorRequest {
	return AddDataSecurityAuditorRequest{
		AddPrivilegeDataSecurityAuditorRequest: v,
	}
}

// AddWeaklyEncodedPasswordDataSecurityAuditorRequestAsAddDataSecurityAuditorRequest is a convenience function that returns AddWeaklyEncodedPasswordDataSecurityAuditorRequest wrapped in AddDataSecurityAuditorRequest
func AddWeaklyEncodedPasswordDataSecurityAuditorRequestAsAddDataSecurityAuditorRequest(v *AddWeaklyEncodedPasswordDataSecurityAuditorRequest) AddDataSecurityAuditorRequest {
	return AddDataSecurityAuditorRequest{
		AddWeaklyEncodedPasswordDataSecurityAuditorRequest: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *AddDataSecurityAuditorRequest) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AddAccessControlDataSecurityAuditorRequest
	err = newStrictDecoder(data).Decode(&dst.AddAccessControlDataSecurityAuditorRequest)
	if err == nil {
		jsonAddAccessControlDataSecurityAuditorRequest, _ := json.Marshal(dst.AddAccessControlDataSecurityAuditorRequest)
		if string(jsonAddAccessControlDataSecurityAuditorRequest) == "{}" { // empty struct
			dst.AddAccessControlDataSecurityAuditorRequest = nil
		} else {
			match++
		}
	} else {
		dst.AddAccessControlDataSecurityAuditorRequest = nil
	}

	// try to unmarshal data into AddDisabledAccountDataSecurityAuditorRequest
	err = newStrictDecoder(data).Decode(&dst.AddDisabledAccountDataSecurityAuditorRequest)
	if err == nil {
		jsonAddDisabledAccountDataSecurityAuditorRequest, _ := json.Marshal(dst.AddDisabledAccountDataSecurityAuditorRequest)
		if string(jsonAddDisabledAccountDataSecurityAuditorRequest) == "{}" { // empty struct
			dst.AddDisabledAccountDataSecurityAuditorRequest = nil
		} else {
			match++
		}
	} else {
		dst.AddDisabledAccountDataSecurityAuditorRequest = nil
	}

	// try to unmarshal data into AddExpiredPasswordDataSecurityAuditorRequest
	err = newStrictDecoder(data).Decode(&dst.AddExpiredPasswordDataSecurityAuditorRequest)
	if err == nil {
		jsonAddExpiredPasswordDataSecurityAuditorRequest, _ := json.Marshal(dst.AddExpiredPasswordDataSecurityAuditorRequest)
		if string(jsonAddExpiredPasswordDataSecurityAuditorRequest) == "{}" { // empty struct
			dst.AddExpiredPasswordDataSecurityAuditorRequest = nil
		} else {
			match++
		}
	} else {
		dst.AddExpiredPasswordDataSecurityAuditorRequest = nil
	}

	// try to unmarshal data into AddLockedAccountDataSecurityAuditorRequest
	err = newStrictDecoder(data).Decode(&dst.AddLockedAccountDataSecurityAuditorRequest)
	if err == nil {
		jsonAddLockedAccountDataSecurityAuditorRequest, _ := json.Marshal(dst.AddLockedAccountDataSecurityAuditorRequest)
		if string(jsonAddLockedAccountDataSecurityAuditorRequest) == "{}" { // empty struct
			dst.AddLockedAccountDataSecurityAuditorRequest = nil
		} else {
			match++
		}
	} else {
		dst.AddLockedAccountDataSecurityAuditorRequest = nil
	}

	// try to unmarshal data into AddMultiplePasswordDataSecurityAuditorRequest
	err = newStrictDecoder(data).Decode(&dst.AddMultiplePasswordDataSecurityAuditorRequest)
	if err == nil {
		jsonAddMultiplePasswordDataSecurityAuditorRequest, _ := json.Marshal(dst.AddMultiplePasswordDataSecurityAuditorRequest)
		if string(jsonAddMultiplePasswordDataSecurityAuditorRequest) == "{}" { // empty struct
			dst.AddMultiplePasswordDataSecurityAuditorRequest = nil
		} else {
			match++
		}
	} else {
		dst.AddMultiplePasswordDataSecurityAuditorRequest = nil
	}

	// try to unmarshal data into AddPrivilegeDataSecurityAuditorRequest
	err = newStrictDecoder(data).Decode(&dst.AddPrivilegeDataSecurityAuditorRequest)
	if err == nil {
		jsonAddPrivilegeDataSecurityAuditorRequest, _ := json.Marshal(dst.AddPrivilegeDataSecurityAuditorRequest)
		if string(jsonAddPrivilegeDataSecurityAuditorRequest) == "{}" { // empty struct
			dst.AddPrivilegeDataSecurityAuditorRequest = nil
		} else {
			match++
		}
	} else {
		dst.AddPrivilegeDataSecurityAuditorRequest = nil
	}

	// try to unmarshal data into AddWeaklyEncodedPasswordDataSecurityAuditorRequest
	err = newStrictDecoder(data).Decode(&dst.AddWeaklyEncodedPasswordDataSecurityAuditorRequest)
	if err == nil {
		jsonAddWeaklyEncodedPasswordDataSecurityAuditorRequest, _ := json.Marshal(dst.AddWeaklyEncodedPasswordDataSecurityAuditorRequest)
		if string(jsonAddWeaklyEncodedPasswordDataSecurityAuditorRequest) == "{}" { // empty struct
			dst.AddWeaklyEncodedPasswordDataSecurityAuditorRequest = nil
		} else {
			match++
		}
	} else {
		dst.AddWeaklyEncodedPasswordDataSecurityAuditorRequest = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.AddAccessControlDataSecurityAuditorRequest = nil
		dst.AddDisabledAccountDataSecurityAuditorRequest = nil
		dst.AddExpiredPasswordDataSecurityAuditorRequest = nil
		dst.AddLockedAccountDataSecurityAuditorRequest = nil
		dst.AddMultiplePasswordDataSecurityAuditorRequest = nil
		dst.AddPrivilegeDataSecurityAuditorRequest = nil
		dst.AddWeaklyEncodedPasswordDataSecurityAuditorRequest = nil

		return fmt.Errorf("data matches more than one schema in oneOf(AddDataSecurityAuditorRequest)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(AddDataSecurityAuditorRequest)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src AddDataSecurityAuditorRequest) MarshalJSON() ([]byte, error) {
	if src.AddAccessControlDataSecurityAuditorRequest != nil {
		return json.Marshal(&src.AddAccessControlDataSecurityAuditorRequest)
	}

	if src.AddDisabledAccountDataSecurityAuditorRequest != nil {
		return json.Marshal(&src.AddDisabledAccountDataSecurityAuditorRequest)
	}

	if src.AddExpiredPasswordDataSecurityAuditorRequest != nil {
		return json.Marshal(&src.AddExpiredPasswordDataSecurityAuditorRequest)
	}

	if src.AddLockedAccountDataSecurityAuditorRequest != nil {
		return json.Marshal(&src.AddLockedAccountDataSecurityAuditorRequest)
	}

	if src.AddMultiplePasswordDataSecurityAuditorRequest != nil {
		return json.Marshal(&src.AddMultiplePasswordDataSecurityAuditorRequest)
	}

	if src.AddPrivilegeDataSecurityAuditorRequest != nil {
		return json.Marshal(&src.AddPrivilegeDataSecurityAuditorRequest)
	}

	if src.AddWeaklyEncodedPasswordDataSecurityAuditorRequest != nil {
		return json.Marshal(&src.AddWeaklyEncodedPasswordDataSecurityAuditorRequest)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *AddDataSecurityAuditorRequest) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.AddAccessControlDataSecurityAuditorRequest != nil {
		return obj.AddAccessControlDataSecurityAuditorRequest
	}

	if obj.AddDisabledAccountDataSecurityAuditorRequest != nil {
		return obj.AddDisabledAccountDataSecurityAuditorRequest
	}

	if obj.AddExpiredPasswordDataSecurityAuditorRequest != nil {
		return obj.AddExpiredPasswordDataSecurityAuditorRequest
	}

	if obj.AddLockedAccountDataSecurityAuditorRequest != nil {
		return obj.AddLockedAccountDataSecurityAuditorRequest
	}

	if obj.AddMultiplePasswordDataSecurityAuditorRequest != nil {
		return obj.AddMultiplePasswordDataSecurityAuditorRequest
	}

	if obj.AddPrivilegeDataSecurityAuditorRequest != nil {
		return obj.AddPrivilegeDataSecurityAuditorRequest
	}

	if obj.AddWeaklyEncodedPasswordDataSecurityAuditorRequest != nil {
		return obj.AddWeaklyEncodedPasswordDataSecurityAuditorRequest
	}

	// all schemas are nil
	return nil
}

type NullableAddDataSecurityAuditorRequest struct {
	value *AddDataSecurityAuditorRequest
	isSet bool
}

func (v NullableAddDataSecurityAuditorRequest) Get() *AddDataSecurityAuditorRequest {
	return v.value
}

func (v *NullableAddDataSecurityAuditorRequest) Set(val *AddDataSecurityAuditorRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddDataSecurityAuditorRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddDataSecurityAuditorRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddDataSecurityAuditorRequest(val *AddDataSecurityAuditorRequest) *NullableAddDataSecurityAuditorRequest {
	return &NullableAddDataSecurityAuditorRequest{value: val, isSet: true}
}

func (v NullableAddDataSecurityAuditorRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddDataSecurityAuditorRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
