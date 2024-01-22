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

// EnumpluginPostCommitValidationProp Specifies the degree of validation that should be performed after a change has been successfully processed, in order to detect any conflicts that may have resulted from other changes processed simultaneously through this or another Directory Server instance. If a conflict is detected during post-commit validation, an administrative alert will be generated to notify administrators of the problem.
type EnumpluginPostCommitValidationProp string

// List of Enumplugin-postCommitValidationProp
const (
	ENUMPLUGINPOSTCOMMITVALIDATIONPROP_NONE                          EnumpluginPostCommitValidationProp = "none"
	ENUMPLUGINPOSTCOMMITVALIDATIONPROP_ALL_SUBTREE_VIEWS             EnumpluginPostCommitValidationProp = "all-subtree-views"
	ENUMPLUGINPOSTCOMMITVALIDATIONPROP_ALL_BACKEND_SETS              EnumpluginPostCommitValidationProp = "all-backend-sets"
	ENUMPLUGINPOSTCOMMITVALIDATIONPROP_ALL_AVAILABLE_BACKEND_SERVERS EnumpluginPostCommitValidationProp = "all-available-backend-servers"
)

// All allowed values of EnumpluginPostCommitValidationProp enum
var AllowedEnumpluginPostCommitValidationPropEnumValues = []EnumpluginPostCommitValidationProp{
	"none",
	"all-subtree-views",
	"all-backend-sets",
	"all-available-backend-servers",
}

func (v *EnumpluginPostCommitValidationProp) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumpluginPostCommitValidationProp(value)
	for _, existing := range AllowedEnumpluginPostCommitValidationPropEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumpluginPostCommitValidationProp", value)
}

// NewEnumpluginPostCommitValidationPropFromValue returns a pointer to a valid EnumpluginPostCommitValidationProp
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumpluginPostCommitValidationPropFromValue(v string) (*EnumpluginPostCommitValidationProp, error) {
	ev := EnumpluginPostCommitValidationProp(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumpluginPostCommitValidationProp: valid values are %v", v, AllowedEnumpluginPostCommitValidationPropEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumpluginPostCommitValidationProp) IsValid() bool {
	for _, existing := range AllowedEnumpluginPostCommitValidationPropEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumplugin-postCommitValidationProp value
func (v EnumpluginPostCommitValidationProp) Ptr() *EnumpluginPostCommitValidationProp {
	return &v
}

type NullableEnumpluginPostCommitValidationProp struct {
	value *EnumpluginPostCommitValidationProp
	isSet bool
}

func (v NullableEnumpluginPostCommitValidationProp) Get() *EnumpluginPostCommitValidationProp {
	return v.value
}

func (v *NullableEnumpluginPostCommitValidationProp) Set(val *EnumpluginPostCommitValidationProp) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumpluginPostCommitValidationProp) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumpluginPostCommitValidationProp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumpluginPostCommitValidationProp(val *EnumpluginPostCommitValidationProp) *NullableEnumpluginPostCommitValidationProp {
	return &NullableEnumpluginPostCommitValidationProp{value: val, isSet: true}
}

func (v NullableEnumpluginPostCommitValidationProp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumpluginPostCommitValidationProp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}