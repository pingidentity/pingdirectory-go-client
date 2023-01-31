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

// AddDataSecurityAuditor200Response - struct for AddDataSecurityAuditor200Response
type AddDataSecurityAuditor200Response struct {
	AccessControlDataSecurityAuditorResponse         *AccessControlDataSecurityAuditorResponse
	DisabledAccountDataSecurityAuditorResponse       *DisabledAccountDataSecurityAuditorResponse
	ExpiredPasswordDataSecurityAuditorResponse       *ExpiredPasswordDataSecurityAuditorResponse
	LockedAccountDataSecurityAuditorResponse         *LockedAccountDataSecurityAuditorResponse
	MultiplePasswordDataSecurityAuditorResponse      *MultiplePasswordDataSecurityAuditorResponse
	PrivilegeDataSecurityAuditorResponse             *PrivilegeDataSecurityAuditorResponse
	WeaklyEncodedPasswordDataSecurityAuditorResponse *WeaklyEncodedPasswordDataSecurityAuditorResponse
}

// AccessControlDataSecurityAuditorResponseAsAddDataSecurityAuditor200Response is a convenience function that returns AccessControlDataSecurityAuditorResponse wrapped in AddDataSecurityAuditor200Response
func AccessControlDataSecurityAuditorResponseAsAddDataSecurityAuditor200Response(v *AccessControlDataSecurityAuditorResponse) AddDataSecurityAuditor200Response {
	return AddDataSecurityAuditor200Response{
		AccessControlDataSecurityAuditorResponse: v,
	}
}

// DisabledAccountDataSecurityAuditorResponseAsAddDataSecurityAuditor200Response is a convenience function that returns DisabledAccountDataSecurityAuditorResponse wrapped in AddDataSecurityAuditor200Response
func DisabledAccountDataSecurityAuditorResponseAsAddDataSecurityAuditor200Response(v *DisabledAccountDataSecurityAuditorResponse) AddDataSecurityAuditor200Response {
	return AddDataSecurityAuditor200Response{
		DisabledAccountDataSecurityAuditorResponse: v,
	}
}

// ExpiredPasswordDataSecurityAuditorResponseAsAddDataSecurityAuditor200Response is a convenience function that returns ExpiredPasswordDataSecurityAuditorResponse wrapped in AddDataSecurityAuditor200Response
func ExpiredPasswordDataSecurityAuditorResponseAsAddDataSecurityAuditor200Response(v *ExpiredPasswordDataSecurityAuditorResponse) AddDataSecurityAuditor200Response {
	return AddDataSecurityAuditor200Response{
		ExpiredPasswordDataSecurityAuditorResponse: v,
	}
}

// LockedAccountDataSecurityAuditorResponseAsAddDataSecurityAuditor200Response is a convenience function that returns LockedAccountDataSecurityAuditorResponse wrapped in AddDataSecurityAuditor200Response
func LockedAccountDataSecurityAuditorResponseAsAddDataSecurityAuditor200Response(v *LockedAccountDataSecurityAuditorResponse) AddDataSecurityAuditor200Response {
	return AddDataSecurityAuditor200Response{
		LockedAccountDataSecurityAuditorResponse: v,
	}
}

// MultiplePasswordDataSecurityAuditorResponseAsAddDataSecurityAuditor200Response is a convenience function that returns MultiplePasswordDataSecurityAuditorResponse wrapped in AddDataSecurityAuditor200Response
func MultiplePasswordDataSecurityAuditorResponseAsAddDataSecurityAuditor200Response(v *MultiplePasswordDataSecurityAuditorResponse) AddDataSecurityAuditor200Response {
	return AddDataSecurityAuditor200Response{
		MultiplePasswordDataSecurityAuditorResponse: v,
	}
}

// PrivilegeDataSecurityAuditorResponseAsAddDataSecurityAuditor200Response is a convenience function that returns PrivilegeDataSecurityAuditorResponse wrapped in AddDataSecurityAuditor200Response
func PrivilegeDataSecurityAuditorResponseAsAddDataSecurityAuditor200Response(v *PrivilegeDataSecurityAuditorResponse) AddDataSecurityAuditor200Response {
	return AddDataSecurityAuditor200Response{
		PrivilegeDataSecurityAuditorResponse: v,
	}
}

// WeaklyEncodedPasswordDataSecurityAuditorResponseAsAddDataSecurityAuditor200Response is a convenience function that returns WeaklyEncodedPasswordDataSecurityAuditorResponse wrapped in AddDataSecurityAuditor200Response
func WeaklyEncodedPasswordDataSecurityAuditorResponseAsAddDataSecurityAuditor200Response(v *WeaklyEncodedPasswordDataSecurityAuditorResponse) AddDataSecurityAuditor200Response {
	return AddDataSecurityAuditor200Response{
		WeaklyEncodedPasswordDataSecurityAuditorResponse: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *AddDataSecurityAuditor200Response) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AccessControlDataSecurityAuditorResponse
	err = newStrictDecoder(data).Decode(&dst.AccessControlDataSecurityAuditorResponse)
	if err == nil {
		jsonAccessControlDataSecurityAuditorResponse, _ := json.Marshal(dst.AccessControlDataSecurityAuditorResponse)
		if string(jsonAccessControlDataSecurityAuditorResponse) == "{}" { // empty struct
			dst.AccessControlDataSecurityAuditorResponse = nil
		} else {
			match++
		}
	} else {
		dst.AccessControlDataSecurityAuditorResponse = nil
	}

	// try to unmarshal data into DisabledAccountDataSecurityAuditorResponse
	err = newStrictDecoder(data).Decode(&dst.DisabledAccountDataSecurityAuditorResponse)
	if err == nil {
		jsonDisabledAccountDataSecurityAuditorResponse, _ := json.Marshal(dst.DisabledAccountDataSecurityAuditorResponse)
		if string(jsonDisabledAccountDataSecurityAuditorResponse) == "{}" { // empty struct
			dst.DisabledAccountDataSecurityAuditorResponse = nil
		} else {
			match++
		}
	} else {
		dst.DisabledAccountDataSecurityAuditorResponse = nil
	}

	// try to unmarshal data into ExpiredPasswordDataSecurityAuditorResponse
	err = newStrictDecoder(data).Decode(&dst.ExpiredPasswordDataSecurityAuditorResponse)
	if err == nil {
		jsonExpiredPasswordDataSecurityAuditorResponse, _ := json.Marshal(dst.ExpiredPasswordDataSecurityAuditorResponse)
		if string(jsonExpiredPasswordDataSecurityAuditorResponse) == "{}" { // empty struct
			dst.ExpiredPasswordDataSecurityAuditorResponse = nil
		} else {
			match++
		}
	} else {
		dst.ExpiredPasswordDataSecurityAuditorResponse = nil
	}

	// try to unmarshal data into LockedAccountDataSecurityAuditorResponse
	err = newStrictDecoder(data).Decode(&dst.LockedAccountDataSecurityAuditorResponse)
	if err == nil {
		jsonLockedAccountDataSecurityAuditorResponse, _ := json.Marshal(dst.LockedAccountDataSecurityAuditorResponse)
		if string(jsonLockedAccountDataSecurityAuditorResponse) == "{}" { // empty struct
			dst.LockedAccountDataSecurityAuditorResponse = nil
		} else {
			match++
		}
	} else {
		dst.LockedAccountDataSecurityAuditorResponse = nil
	}

	// try to unmarshal data into MultiplePasswordDataSecurityAuditorResponse
	err = newStrictDecoder(data).Decode(&dst.MultiplePasswordDataSecurityAuditorResponse)
	if err == nil {
		jsonMultiplePasswordDataSecurityAuditorResponse, _ := json.Marshal(dst.MultiplePasswordDataSecurityAuditorResponse)
		if string(jsonMultiplePasswordDataSecurityAuditorResponse) == "{}" { // empty struct
			dst.MultiplePasswordDataSecurityAuditorResponse = nil
		} else {
			match++
		}
	} else {
		dst.MultiplePasswordDataSecurityAuditorResponse = nil
	}

	// try to unmarshal data into PrivilegeDataSecurityAuditorResponse
	err = newStrictDecoder(data).Decode(&dst.PrivilegeDataSecurityAuditorResponse)
	if err == nil {
		jsonPrivilegeDataSecurityAuditorResponse, _ := json.Marshal(dst.PrivilegeDataSecurityAuditorResponse)
		if string(jsonPrivilegeDataSecurityAuditorResponse) == "{}" { // empty struct
			dst.PrivilegeDataSecurityAuditorResponse = nil
		} else {
			match++
		}
	} else {
		dst.PrivilegeDataSecurityAuditorResponse = nil
	}

	// try to unmarshal data into WeaklyEncodedPasswordDataSecurityAuditorResponse
	err = newStrictDecoder(data).Decode(&dst.WeaklyEncodedPasswordDataSecurityAuditorResponse)
	if err == nil {
		jsonWeaklyEncodedPasswordDataSecurityAuditorResponse, _ := json.Marshal(dst.WeaklyEncodedPasswordDataSecurityAuditorResponse)
		if string(jsonWeaklyEncodedPasswordDataSecurityAuditorResponse) == "{}" { // empty struct
			dst.WeaklyEncodedPasswordDataSecurityAuditorResponse = nil
		} else {
			match++
		}
	} else {
		dst.WeaklyEncodedPasswordDataSecurityAuditorResponse = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.AccessControlDataSecurityAuditorResponse = nil
		dst.DisabledAccountDataSecurityAuditorResponse = nil
		dst.ExpiredPasswordDataSecurityAuditorResponse = nil
		dst.LockedAccountDataSecurityAuditorResponse = nil
		dst.MultiplePasswordDataSecurityAuditorResponse = nil
		dst.PrivilegeDataSecurityAuditorResponse = nil
		dst.WeaklyEncodedPasswordDataSecurityAuditorResponse = nil

		return fmt.Errorf("data matches more than one schema in oneOf(AddDataSecurityAuditor200Response)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(AddDataSecurityAuditor200Response)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src AddDataSecurityAuditor200Response) MarshalJSON() ([]byte, error) {
	if src.AccessControlDataSecurityAuditorResponse != nil {
		return json.Marshal(&src.AccessControlDataSecurityAuditorResponse)
	}

	if src.DisabledAccountDataSecurityAuditorResponse != nil {
		return json.Marshal(&src.DisabledAccountDataSecurityAuditorResponse)
	}

	if src.ExpiredPasswordDataSecurityAuditorResponse != nil {
		return json.Marshal(&src.ExpiredPasswordDataSecurityAuditorResponse)
	}

	if src.LockedAccountDataSecurityAuditorResponse != nil {
		return json.Marshal(&src.LockedAccountDataSecurityAuditorResponse)
	}

	if src.MultiplePasswordDataSecurityAuditorResponse != nil {
		return json.Marshal(&src.MultiplePasswordDataSecurityAuditorResponse)
	}

	if src.PrivilegeDataSecurityAuditorResponse != nil {
		return json.Marshal(&src.PrivilegeDataSecurityAuditorResponse)
	}

	if src.WeaklyEncodedPasswordDataSecurityAuditorResponse != nil {
		return json.Marshal(&src.WeaklyEncodedPasswordDataSecurityAuditorResponse)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *AddDataSecurityAuditor200Response) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.AccessControlDataSecurityAuditorResponse != nil {
		return obj.AccessControlDataSecurityAuditorResponse
	}

	if obj.DisabledAccountDataSecurityAuditorResponse != nil {
		return obj.DisabledAccountDataSecurityAuditorResponse
	}

	if obj.ExpiredPasswordDataSecurityAuditorResponse != nil {
		return obj.ExpiredPasswordDataSecurityAuditorResponse
	}

	if obj.LockedAccountDataSecurityAuditorResponse != nil {
		return obj.LockedAccountDataSecurityAuditorResponse
	}

	if obj.MultiplePasswordDataSecurityAuditorResponse != nil {
		return obj.MultiplePasswordDataSecurityAuditorResponse
	}

	if obj.PrivilegeDataSecurityAuditorResponse != nil {
		return obj.PrivilegeDataSecurityAuditorResponse
	}

	if obj.WeaklyEncodedPasswordDataSecurityAuditorResponse != nil {
		return obj.WeaklyEncodedPasswordDataSecurityAuditorResponse
	}

	// all schemas are nil
	return nil
}

type NullableAddDataSecurityAuditor200Response struct {
	value *AddDataSecurityAuditor200Response
	isSet bool
}

func (v NullableAddDataSecurityAuditor200Response) Get() *AddDataSecurityAuditor200Response {
	return v.value
}

func (v *NullableAddDataSecurityAuditor200Response) Set(val *AddDataSecurityAuditor200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableAddDataSecurityAuditor200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableAddDataSecurityAuditor200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddDataSecurityAuditor200Response(val *AddDataSecurityAuditor200Response) *NullableAddDataSecurityAuditor200Response {
	return &NullableAddDataSecurityAuditor200Response{value: val, isSet: true}
}

func (v NullableAddDataSecurityAuditor200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddDataSecurityAuditor200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
