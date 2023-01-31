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

// EnumaccessControlHandlerAllowedBindControlProp Specifies a set of controls that clients should be allowed to include in bind requests. As bind requests are evaluated as the unauthenticated user, any controls included in this set will be permitted for any bind attempt. If you wish to grant permission for any bind controls not listed here, then the allowed-bind-control-oid property may be used to accomplish that.
type EnumaccessControlHandlerAllowedBindControlProp string

// List of Enumaccess-control-handler-allowedBindControlProp
const (
	ENUMACCESSCONTROLHANDLERALLOWEDBINDCONTROLPROP_ADMINISTRATIVE_OPERATION              EnumaccessControlHandlerAllowedBindControlProp = "administrative-operation"
	ENUMACCESSCONTROLHANDLERALLOWEDBINDCONTROLPROP_AUTHORIZATION_IDENTITY                EnumaccessControlHandlerAllowedBindControlProp = "authorization-identity"
	ENUMACCESSCONTROLHANDLERALLOWEDBINDCONTROLPROP_GET_AUTHORIZATION_ENTRY               EnumaccessControlHandlerAllowedBindControlProp = "get-authorization-entry"
	ENUMACCESSCONTROLHANDLERALLOWEDBINDCONTROLPROP_GET_BACKEND_SET_ID                    EnumaccessControlHandlerAllowedBindControlProp = "get-backend-set-id"
	ENUMACCESSCONTROLHANDLERALLOWEDBINDCONTROLPROP_GET_PASSWORD_POLICY_STATE_ISSUES      EnumaccessControlHandlerAllowedBindControlProp = "get-password-policy-state-issues"
	ENUMACCESSCONTROLHANDLERALLOWEDBINDCONTROLPROP_GET_RECENT_LOGIN_HISTORY              EnumaccessControlHandlerAllowedBindControlProp = "get-recent-login-history"
	ENUMACCESSCONTROLHANDLERALLOWEDBINDCONTROLPROP_GET_SERVER_ID                         EnumaccessControlHandlerAllowedBindControlProp = "get-server-id"
	ENUMACCESSCONTROLHANDLERALLOWEDBINDCONTROLPROP_GET_USER_RESOURCE_LIMITS              EnumaccessControlHandlerAllowedBindControlProp = "get-user-resource-limits"
	ENUMACCESSCONTROLHANDLERALLOWEDBINDCONTROLPROP_INTERMEDIATE_CLIENT                   EnumaccessControlHandlerAllowedBindControlProp = "intermediate-client"
	ENUMACCESSCONTROLHANDLERALLOWEDBINDCONTROLPROP_OPERATION_PURPOSE                     EnumaccessControlHandlerAllowedBindControlProp = "operation-purpose"
	ENUMACCESSCONTROLHANDLERALLOWEDBINDCONTROLPROP_PASSWORD_POLICY                       EnumaccessControlHandlerAllowedBindControlProp = "password-policy"
	ENUMACCESSCONTROLHANDLERALLOWEDBINDCONTROLPROP_RETAIN_IDENTITY                       EnumaccessControlHandlerAllowedBindControlProp = "retain-identity"
	ENUMACCESSCONTROLHANDLERALLOWEDBINDCONTROLPROP_ROUTE_TO_BACKEND_SET                  EnumaccessControlHandlerAllowedBindControlProp = "route-to-backend-set"
	ENUMACCESSCONTROLHANDLERALLOWEDBINDCONTROLPROP_ROUTE_TO_SERVER                       EnumaccessControlHandlerAllowedBindControlProp = "route-to-server"
	ENUMACCESSCONTROLHANDLERALLOWEDBINDCONTROLPROP_SUPPRESS_OPERATIONAL_ATTRIBUTE_UPDATE EnumaccessControlHandlerAllowedBindControlProp = "suppress-operational-attribute-update"
)

// All allowed values of EnumaccessControlHandlerAllowedBindControlProp enum
var AllowedEnumaccessControlHandlerAllowedBindControlPropEnumValues = []EnumaccessControlHandlerAllowedBindControlProp{
	"administrative-operation",
	"authorization-identity",
	"get-authorization-entry",
	"get-backend-set-id",
	"get-password-policy-state-issues",
	"get-recent-login-history",
	"get-server-id",
	"get-user-resource-limits",
	"intermediate-client",
	"operation-purpose",
	"password-policy",
	"retain-identity",
	"route-to-backend-set",
	"route-to-server",
	"suppress-operational-attribute-update",
}

func (v *EnumaccessControlHandlerAllowedBindControlProp) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumaccessControlHandlerAllowedBindControlProp(value)
	for _, existing := range AllowedEnumaccessControlHandlerAllowedBindControlPropEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumaccessControlHandlerAllowedBindControlProp", value)
}

// NewEnumaccessControlHandlerAllowedBindControlPropFromValue returns a pointer to a valid EnumaccessControlHandlerAllowedBindControlProp
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumaccessControlHandlerAllowedBindControlPropFromValue(v string) (*EnumaccessControlHandlerAllowedBindControlProp, error) {
	ev := EnumaccessControlHandlerAllowedBindControlProp(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumaccessControlHandlerAllowedBindControlProp: valid values are %v", v, AllowedEnumaccessControlHandlerAllowedBindControlPropEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumaccessControlHandlerAllowedBindControlProp) IsValid() bool {
	for _, existing := range AllowedEnumaccessControlHandlerAllowedBindControlPropEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumaccess-control-handler-allowedBindControlProp value
func (v EnumaccessControlHandlerAllowedBindControlProp) Ptr() *EnumaccessControlHandlerAllowedBindControlProp {
	return &v
}

type NullableEnumaccessControlHandlerAllowedBindControlProp struct {
	value *EnumaccessControlHandlerAllowedBindControlProp
	isSet bool
}

func (v NullableEnumaccessControlHandlerAllowedBindControlProp) Get() *EnumaccessControlHandlerAllowedBindControlProp {
	return v.value
}

func (v *NullableEnumaccessControlHandlerAllowedBindControlProp) Set(val *EnumaccessControlHandlerAllowedBindControlProp) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumaccessControlHandlerAllowedBindControlProp) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumaccessControlHandlerAllowedBindControlProp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumaccessControlHandlerAllowedBindControlProp(val *EnumaccessControlHandlerAllowedBindControlProp) *NullableEnumaccessControlHandlerAllowedBindControlProp {
	return &NullableEnumaccessControlHandlerAllowedBindControlProp{value: val, isSet: true}
}

func (v NullableEnumaccessControlHandlerAllowedBindControlProp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumaccessControlHandlerAllowedBindControlProp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
