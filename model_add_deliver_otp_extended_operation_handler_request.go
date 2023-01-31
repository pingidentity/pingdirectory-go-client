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

// AddDeliverOtpExtendedOperationHandlerRequest struct for AddDeliverOtpExtendedOperationHandlerRequest
type AddDeliverOtpExtendedOperationHandlerRequest struct {
	// Name of the new Extended Operation Handler
	HandlerName string                                            `json:"handlerName"`
	Schemas     []EnumdeliverOtpExtendedOperationHandlerSchemaUrn `json:"schemas"`
	// The identity mapper that should be used to identify the user(s) targeted by the authentication identity contained in the extended request. This will only be used for \"u:\"-style authentication identities.
	IdentityMapper string `json:"identityMapper"`
	// The password generator that will be used to create the one-time password values to be delivered to the end user.
	PasswordGenerator string `json:"passwordGenerator"`
	// The set of delivery mechanisms that may be used to deliver one-time passwords to users in requests that do not specify one or more preferred delivery mechanisms.
	DefaultOTPDeliveryMechanism []string `json:"defaultOTPDeliveryMechanism"`
	// A description for this Extended Operation Handler
	Description *string `json:"description,omitempty"`
	// Indicates whether the Extended Operation Handler is enabled (that is, whether the types of extended operations are allowed in the server).
	Enabled bool `json:"enabled"`
}

// NewAddDeliverOtpExtendedOperationHandlerRequest instantiates a new AddDeliverOtpExtendedOperationHandlerRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddDeliverOtpExtendedOperationHandlerRequest(handlerName string, schemas []EnumdeliverOtpExtendedOperationHandlerSchemaUrn, identityMapper string, passwordGenerator string, defaultOTPDeliveryMechanism []string, enabled bool) *AddDeliverOtpExtendedOperationHandlerRequest {
	this := AddDeliverOtpExtendedOperationHandlerRequest{}
	this.HandlerName = handlerName
	this.Schemas = schemas
	this.IdentityMapper = identityMapper
	this.PasswordGenerator = passwordGenerator
	this.DefaultOTPDeliveryMechanism = defaultOTPDeliveryMechanism
	this.Enabled = enabled
	return &this
}

// NewAddDeliverOtpExtendedOperationHandlerRequestWithDefaults instantiates a new AddDeliverOtpExtendedOperationHandlerRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddDeliverOtpExtendedOperationHandlerRequestWithDefaults() *AddDeliverOtpExtendedOperationHandlerRequest {
	this := AddDeliverOtpExtendedOperationHandlerRequest{}
	return &this
}

// GetHandlerName returns the HandlerName field value
func (o *AddDeliverOtpExtendedOperationHandlerRequest) GetHandlerName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HandlerName
}

// GetHandlerNameOk returns a tuple with the HandlerName field value
// and a boolean to check if the value has been set.
func (o *AddDeliverOtpExtendedOperationHandlerRequest) GetHandlerNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HandlerName, true
}

// SetHandlerName sets field value
func (o *AddDeliverOtpExtendedOperationHandlerRequest) SetHandlerName(v string) {
	o.HandlerName = v
}

// GetSchemas returns the Schemas field value
func (o *AddDeliverOtpExtendedOperationHandlerRequest) GetSchemas() []EnumdeliverOtpExtendedOperationHandlerSchemaUrn {
	if o == nil {
		var ret []EnumdeliverOtpExtendedOperationHandlerSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddDeliverOtpExtendedOperationHandlerRequest) GetSchemasOk() ([]EnumdeliverOtpExtendedOperationHandlerSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddDeliverOtpExtendedOperationHandlerRequest) SetSchemas(v []EnumdeliverOtpExtendedOperationHandlerSchemaUrn) {
	o.Schemas = v
}

// GetIdentityMapper returns the IdentityMapper field value
func (o *AddDeliverOtpExtendedOperationHandlerRequest) GetIdentityMapper() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IdentityMapper
}

// GetIdentityMapperOk returns a tuple with the IdentityMapper field value
// and a boolean to check if the value has been set.
func (o *AddDeliverOtpExtendedOperationHandlerRequest) GetIdentityMapperOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IdentityMapper, true
}

// SetIdentityMapper sets field value
func (o *AddDeliverOtpExtendedOperationHandlerRequest) SetIdentityMapper(v string) {
	o.IdentityMapper = v
}

// GetPasswordGenerator returns the PasswordGenerator field value
func (o *AddDeliverOtpExtendedOperationHandlerRequest) GetPasswordGenerator() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PasswordGenerator
}

// GetPasswordGeneratorOk returns a tuple with the PasswordGenerator field value
// and a boolean to check if the value has been set.
func (o *AddDeliverOtpExtendedOperationHandlerRequest) GetPasswordGeneratorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PasswordGenerator, true
}

// SetPasswordGenerator sets field value
func (o *AddDeliverOtpExtendedOperationHandlerRequest) SetPasswordGenerator(v string) {
	o.PasswordGenerator = v
}

// GetDefaultOTPDeliveryMechanism returns the DefaultOTPDeliveryMechanism field value
func (o *AddDeliverOtpExtendedOperationHandlerRequest) GetDefaultOTPDeliveryMechanism() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.DefaultOTPDeliveryMechanism
}

// GetDefaultOTPDeliveryMechanismOk returns a tuple with the DefaultOTPDeliveryMechanism field value
// and a boolean to check if the value has been set.
func (o *AddDeliverOtpExtendedOperationHandlerRequest) GetDefaultOTPDeliveryMechanismOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DefaultOTPDeliveryMechanism, true
}

// SetDefaultOTPDeliveryMechanism sets field value
func (o *AddDeliverOtpExtendedOperationHandlerRequest) SetDefaultOTPDeliveryMechanism(v []string) {
	o.DefaultOTPDeliveryMechanism = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddDeliverOtpExtendedOperationHandlerRequest) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddDeliverOtpExtendedOperationHandlerRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddDeliverOtpExtendedOperationHandlerRequest) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddDeliverOtpExtendedOperationHandlerRequest) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *AddDeliverOtpExtendedOperationHandlerRequest) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AddDeliverOtpExtendedOperationHandlerRequest) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AddDeliverOtpExtendedOperationHandlerRequest) SetEnabled(v bool) {
	o.Enabled = v
}

func (o AddDeliverOtpExtendedOperationHandlerRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["handlerName"] = o.HandlerName
	}
	if true {
		toSerialize["schemas"] = o.Schemas
	}
	if true {
		toSerialize["identityMapper"] = o.IdentityMapper
	}
	if true {
		toSerialize["passwordGenerator"] = o.PasswordGenerator
	}
	if true {
		toSerialize["defaultOTPDeliveryMechanism"] = o.DefaultOTPDeliveryMechanism
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableAddDeliverOtpExtendedOperationHandlerRequest struct {
	value *AddDeliverOtpExtendedOperationHandlerRequest
	isSet bool
}

func (v NullableAddDeliverOtpExtendedOperationHandlerRequest) Get() *AddDeliverOtpExtendedOperationHandlerRequest {
	return v.value
}

func (v *NullableAddDeliverOtpExtendedOperationHandlerRequest) Set(val *AddDeliverOtpExtendedOperationHandlerRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddDeliverOtpExtendedOperationHandlerRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddDeliverOtpExtendedOperationHandlerRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddDeliverOtpExtendedOperationHandlerRequest(val *AddDeliverOtpExtendedOperationHandlerRequest) *NullableAddDeliverOtpExtendedOperationHandlerRequest {
	return &NullableAddDeliverOtpExtendedOperationHandlerRequest{value: val, isSet: true}
}

func (v NullableAddDeliverOtpExtendedOperationHandlerRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddDeliverOtpExtendedOperationHandlerRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
