/*
PingData Config - OpenAPI 3.0

This is the PingData Configuration API

API version: 0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// AddDelayBindResponseFailureLockoutActionRequest struct for AddDelayBindResponseFailureLockoutActionRequest
type AddDelayBindResponseFailureLockoutActionRequest struct {
	// Name of the new Failure Lockout Action
	ActionName string `json:"actionName"`
	Schemas []EnumdelayBindResponseFailureLockoutActionSchemaUrn `json:"schemas"`
	// The length of time to delay the bind response for accounts with too many failed authentication attempts.
	Delay string `json:"delay"`
	// Indicates whether to delay the response for authentication attempts even if that delay may block the thread being used to process the attempt.
	AllowBlockingDelay *bool `json:"allowBlockingDelay,omitempty"`
	// Indicates whether to generate an account status notification for cases in which a bind response is delayed because of failure lockout.
	GenerateAccountStatusNotification *bool `json:"generateAccountStatusNotification,omitempty"`
	// A description for this Failure Lockout Action
	Description *string `json:"description,omitempty"`
}

// NewAddDelayBindResponseFailureLockoutActionRequest instantiates a new AddDelayBindResponseFailureLockoutActionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddDelayBindResponseFailureLockoutActionRequest(actionName string, schemas []EnumdelayBindResponseFailureLockoutActionSchemaUrn, delay string) *AddDelayBindResponseFailureLockoutActionRequest {
	this := AddDelayBindResponseFailureLockoutActionRequest{}
	this.ActionName = actionName
	this.Schemas = schemas
	this.Delay = delay
	return &this
}

// NewAddDelayBindResponseFailureLockoutActionRequestWithDefaults instantiates a new AddDelayBindResponseFailureLockoutActionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddDelayBindResponseFailureLockoutActionRequestWithDefaults() *AddDelayBindResponseFailureLockoutActionRequest {
	this := AddDelayBindResponseFailureLockoutActionRequest{}
	return &this
}

// GetActionName returns the ActionName field value
func (o *AddDelayBindResponseFailureLockoutActionRequest) GetActionName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ActionName
}

// GetActionNameOk returns a tuple with the ActionName field value
// and a boolean to check if the value has been set.
func (o *AddDelayBindResponseFailureLockoutActionRequest) GetActionNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ActionName, true
}

// SetActionName sets field value
func (o *AddDelayBindResponseFailureLockoutActionRequest) SetActionName(v string) {
	o.ActionName = v
}

// GetSchemas returns the Schemas field value
func (o *AddDelayBindResponseFailureLockoutActionRequest) GetSchemas() []EnumdelayBindResponseFailureLockoutActionSchemaUrn {
	if o == nil {
		var ret []EnumdelayBindResponseFailureLockoutActionSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddDelayBindResponseFailureLockoutActionRequest) GetSchemasOk() ([]EnumdelayBindResponseFailureLockoutActionSchemaUrn, bool) {
	if o == nil {
    return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddDelayBindResponseFailureLockoutActionRequest) SetSchemas(v []EnumdelayBindResponseFailureLockoutActionSchemaUrn) {
	o.Schemas = v
}

// GetDelay returns the Delay field value
func (o *AddDelayBindResponseFailureLockoutActionRequest) GetDelay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Delay
}

// GetDelayOk returns a tuple with the Delay field value
// and a boolean to check if the value has been set.
func (o *AddDelayBindResponseFailureLockoutActionRequest) GetDelayOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Delay, true
}

// SetDelay sets field value
func (o *AddDelayBindResponseFailureLockoutActionRequest) SetDelay(v string) {
	o.Delay = v
}

// GetAllowBlockingDelay returns the AllowBlockingDelay field value if set, zero value otherwise.
func (o *AddDelayBindResponseFailureLockoutActionRequest) GetAllowBlockingDelay() bool {
	if o == nil || isNil(o.AllowBlockingDelay) {
		var ret bool
		return ret
	}
	return *o.AllowBlockingDelay
}

// GetAllowBlockingDelayOk returns a tuple with the AllowBlockingDelay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddDelayBindResponseFailureLockoutActionRequest) GetAllowBlockingDelayOk() (*bool, bool) {
	if o == nil || isNil(o.AllowBlockingDelay) {
    return nil, false
	}
	return o.AllowBlockingDelay, true
}

// HasAllowBlockingDelay returns a boolean if a field has been set.
func (o *AddDelayBindResponseFailureLockoutActionRequest) HasAllowBlockingDelay() bool {
	if o != nil && !isNil(o.AllowBlockingDelay) {
		return true
	}

	return false
}

// SetAllowBlockingDelay gets a reference to the given bool and assigns it to the AllowBlockingDelay field.
func (o *AddDelayBindResponseFailureLockoutActionRequest) SetAllowBlockingDelay(v bool) {
	o.AllowBlockingDelay = &v
}

// GetGenerateAccountStatusNotification returns the GenerateAccountStatusNotification field value if set, zero value otherwise.
func (o *AddDelayBindResponseFailureLockoutActionRequest) GetGenerateAccountStatusNotification() bool {
	if o == nil || isNil(o.GenerateAccountStatusNotification) {
		var ret bool
		return ret
	}
	return *o.GenerateAccountStatusNotification
}

// GetGenerateAccountStatusNotificationOk returns a tuple with the GenerateAccountStatusNotification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddDelayBindResponseFailureLockoutActionRequest) GetGenerateAccountStatusNotificationOk() (*bool, bool) {
	if o == nil || isNil(o.GenerateAccountStatusNotification) {
    return nil, false
	}
	return o.GenerateAccountStatusNotification, true
}

// HasGenerateAccountStatusNotification returns a boolean if a field has been set.
func (o *AddDelayBindResponseFailureLockoutActionRequest) HasGenerateAccountStatusNotification() bool {
	if o != nil && !isNil(o.GenerateAccountStatusNotification) {
		return true
	}

	return false
}

// SetGenerateAccountStatusNotification gets a reference to the given bool and assigns it to the GenerateAccountStatusNotification field.
func (o *AddDelayBindResponseFailureLockoutActionRequest) SetGenerateAccountStatusNotification(v bool) {
	o.GenerateAccountStatusNotification = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddDelayBindResponseFailureLockoutActionRequest) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddDelayBindResponseFailureLockoutActionRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddDelayBindResponseFailureLockoutActionRequest) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddDelayBindResponseFailureLockoutActionRequest) SetDescription(v string) {
	o.Description = &v
}

func (o AddDelayBindResponseFailureLockoutActionRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["actionName"] = o.ActionName
	}
	if true {
		toSerialize["schemas"] = o.Schemas
	}
	if true {
		toSerialize["delay"] = o.Delay
	}
	if !isNil(o.AllowBlockingDelay) {
		toSerialize["allowBlockingDelay"] = o.AllowBlockingDelay
	}
	if !isNil(o.GenerateAccountStatusNotification) {
		toSerialize["generateAccountStatusNotification"] = o.GenerateAccountStatusNotification
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullableAddDelayBindResponseFailureLockoutActionRequest struct {
	value *AddDelayBindResponseFailureLockoutActionRequest
	isSet bool
}

func (v NullableAddDelayBindResponseFailureLockoutActionRequest) Get() *AddDelayBindResponseFailureLockoutActionRequest {
	return v.value
}

func (v *NullableAddDelayBindResponseFailureLockoutActionRequest) Set(val *AddDelayBindResponseFailureLockoutActionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddDelayBindResponseFailureLockoutActionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddDelayBindResponseFailureLockoutActionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddDelayBindResponseFailureLockoutActionRequest(val *AddDelayBindResponseFailureLockoutActionRequest) *NullableAddDelayBindResponseFailureLockoutActionRequest {
	return &NullableAddDelayBindResponseFailureLockoutActionRequest{value: val, isSet: true}
}

func (v NullableAddDelayBindResponseFailureLockoutActionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddDelayBindResponseFailureLockoutActionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


