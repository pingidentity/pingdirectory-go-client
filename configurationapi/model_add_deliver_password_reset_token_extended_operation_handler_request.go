/*
PingData Config - OpenAPI 3.0

This is the PingData Configuration API

API version: 0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configurationapi

import (
	"encoding/json"
)

// checks if the AddDeliverPasswordResetTokenExtendedOperationHandlerRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddDeliverPasswordResetTokenExtendedOperationHandlerRequest{}

// AddDeliverPasswordResetTokenExtendedOperationHandlerRequest struct for AddDeliverPasswordResetTokenExtendedOperationHandlerRequest
type AddDeliverPasswordResetTokenExtendedOperationHandlerRequest struct {
	Schemas []EnumdeliverPasswordResetTokenExtendedOperationHandlerSchemaUrn `json:"schemas"`
	// The password generator that will be used to create the password reset token values to be delivered to the end user.
	PasswordGenerator string `json:"passwordGenerator"`
	// The set of delivery mechanisms that may be used to deliver password reset tokens to users for requests that do not specify one or more preferred delivery mechanisms.
	DefaultTokenDeliveryMechanism []string `json:"defaultTokenDeliveryMechanism"`
	// The maximum length of time that a password reset token should be considered valid.
	PasswordResetTokenValidityDuration *string `json:"passwordResetTokenValidityDuration,omitempty"`
	// A description for this Extended Operation Handler
	Description *string `json:"description,omitempty"`
	// Indicates whether the Extended Operation Handler is enabled (that is, whether the types of extended operations are allowed in the server).
	Enabled bool `json:"enabled"`
	// Name of the new Extended Operation Handler
	HandlerName string `json:"handlerName"`
}

// NewAddDeliverPasswordResetTokenExtendedOperationHandlerRequest instantiates a new AddDeliverPasswordResetTokenExtendedOperationHandlerRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddDeliverPasswordResetTokenExtendedOperationHandlerRequest(schemas []EnumdeliverPasswordResetTokenExtendedOperationHandlerSchemaUrn, passwordGenerator string, defaultTokenDeliveryMechanism []string, enabled bool, handlerName string) *AddDeliverPasswordResetTokenExtendedOperationHandlerRequest {
	this := AddDeliverPasswordResetTokenExtendedOperationHandlerRequest{}
	this.Schemas = schemas
	this.PasswordGenerator = passwordGenerator
	this.DefaultTokenDeliveryMechanism = defaultTokenDeliveryMechanism
	this.Enabled = enabled
	this.HandlerName = handlerName
	return &this
}

// NewAddDeliverPasswordResetTokenExtendedOperationHandlerRequestWithDefaults instantiates a new AddDeliverPasswordResetTokenExtendedOperationHandlerRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddDeliverPasswordResetTokenExtendedOperationHandlerRequestWithDefaults() *AddDeliverPasswordResetTokenExtendedOperationHandlerRequest {
	this := AddDeliverPasswordResetTokenExtendedOperationHandlerRequest{}
	return &this
}

// GetSchemas returns the Schemas field value
func (o *AddDeliverPasswordResetTokenExtendedOperationHandlerRequest) GetSchemas() []EnumdeliverPasswordResetTokenExtendedOperationHandlerSchemaUrn {
	if o == nil {
		var ret []EnumdeliverPasswordResetTokenExtendedOperationHandlerSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddDeliverPasswordResetTokenExtendedOperationHandlerRequest) GetSchemasOk() ([]EnumdeliverPasswordResetTokenExtendedOperationHandlerSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddDeliverPasswordResetTokenExtendedOperationHandlerRequest) SetSchemas(v []EnumdeliverPasswordResetTokenExtendedOperationHandlerSchemaUrn) {
	o.Schemas = v
}

// GetPasswordGenerator returns the PasswordGenerator field value
func (o *AddDeliverPasswordResetTokenExtendedOperationHandlerRequest) GetPasswordGenerator() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PasswordGenerator
}

// GetPasswordGeneratorOk returns a tuple with the PasswordGenerator field value
// and a boolean to check if the value has been set.
func (o *AddDeliverPasswordResetTokenExtendedOperationHandlerRequest) GetPasswordGeneratorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PasswordGenerator, true
}

// SetPasswordGenerator sets field value
func (o *AddDeliverPasswordResetTokenExtendedOperationHandlerRequest) SetPasswordGenerator(v string) {
	o.PasswordGenerator = v
}

// GetDefaultTokenDeliveryMechanism returns the DefaultTokenDeliveryMechanism field value
func (o *AddDeliverPasswordResetTokenExtendedOperationHandlerRequest) GetDefaultTokenDeliveryMechanism() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.DefaultTokenDeliveryMechanism
}

// GetDefaultTokenDeliveryMechanismOk returns a tuple with the DefaultTokenDeliveryMechanism field value
// and a boolean to check if the value has been set.
func (o *AddDeliverPasswordResetTokenExtendedOperationHandlerRequest) GetDefaultTokenDeliveryMechanismOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DefaultTokenDeliveryMechanism, true
}

// SetDefaultTokenDeliveryMechanism sets field value
func (o *AddDeliverPasswordResetTokenExtendedOperationHandlerRequest) SetDefaultTokenDeliveryMechanism(v []string) {
	o.DefaultTokenDeliveryMechanism = v
}

// GetPasswordResetTokenValidityDuration returns the PasswordResetTokenValidityDuration field value if set, zero value otherwise.
func (o *AddDeliverPasswordResetTokenExtendedOperationHandlerRequest) GetPasswordResetTokenValidityDuration() string {
	if o == nil || IsNil(o.PasswordResetTokenValidityDuration) {
		var ret string
		return ret
	}
	return *o.PasswordResetTokenValidityDuration
}

// GetPasswordResetTokenValidityDurationOk returns a tuple with the PasswordResetTokenValidityDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddDeliverPasswordResetTokenExtendedOperationHandlerRequest) GetPasswordResetTokenValidityDurationOk() (*string, bool) {
	if o == nil || IsNil(o.PasswordResetTokenValidityDuration) {
		return nil, false
	}
	return o.PasswordResetTokenValidityDuration, true
}

// HasPasswordResetTokenValidityDuration returns a boolean if a field has been set.
func (o *AddDeliverPasswordResetTokenExtendedOperationHandlerRequest) HasPasswordResetTokenValidityDuration() bool {
	if o != nil && !IsNil(o.PasswordResetTokenValidityDuration) {
		return true
	}

	return false
}

// SetPasswordResetTokenValidityDuration gets a reference to the given string and assigns it to the PasswordResetTokenValidityDuration field.
func (o *AddDeliverPasswordResetTokenExtendedOperationHandlerRequest) SetPasswordResetTokenValidityDuration(v string) {
	o.PasswordResetTokenValidityDuration = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddDeliverPasswordResetTokenExtendedOperationHandlerRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddDeliverPasswordResetTokenExtendedOperationHandlerRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddDeliverPasswordResetTokenExtendedOperationHandlerRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddDeliverPasswordResetTokenExtendedOperationHandlerRequest) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *AddDeliverPasswordResetTokenExtendedOperationHandlerRequest) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AddDeliverPasswordResetTokenExtendedOperationHandlerRequest) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AddDeliverPasswordResetTokenExtendedOperationHandlerRequest) SetEnabled(v bool) {
	o.Enabled = v
}

// GetHandlerName returns the HandlerName field value
func (o *AddDeliverPasswordResetTokenExtendedOperationHandlerRequest) GetHandlerName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HandlerName
}

// GetHandlerNameOk returns a tuple with the HandlerName field value
// and a boolean to check if the value has been set.
func (o *AddDeliverPasswordResetTokenExtendedOperationHandlerRequest) GetHandlerNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HandlerName, true
}

// SetHandlerName sets field value
func (o *AddDeliverPasswordResetTokenExtendedOperationHandlerRequest) SetHandlerName(v string) {
	o.HandlerName = v
}

func (o AddDeliverPasswordResetTokenExtendedOperationHandlerRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddDeliverPasswordResetTokenExtendedOperationHandlerRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["schemas"] = o.Schemas
	toSerialize["passwordGenerator"] = o.PasswordGenerator
	toSerialize["defaultTokenDeliveryMechanism"] = o.DefaultTokenDeliveryMechanism
	if !IsNil(o.PasswordResetTokenValidityDuration) {
		toSerialize["passwordResetTokenValidityDuration"] = o.PasswordResetTokenValidityDuration
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["enabled"] = o.Enabled
	toSerialize["handlerName"] = o.HandlerName
	return toSerialize, nil
}

type NullableAddDeliverPasswordResetTokenExtendedOperationHandlerRequest struct {
	value *AddDeliverPasswordResetTokenExtendedOperationHandlerRequest
	isSet bool
}

func (v NullableAddDeliverPasswordResetTokenExtendedOperationHandlerRequest) Get() *AddDeliverPasswordResetTokenExtendedOperationHandlerRequest {
	return v.value
}

func (v *NullableAddDeliverPasswordResetTokenExtendedOperationHandlerRequest) Set(val *AddDeliverPasswordResetTokenExtendedOperationHandlerRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddDeliverPasswordResetTokenExtendedOperationHandlerRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddDeliverPasswordResetTokenExtendedOperationHandlerRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddDeliverPasswordResetTokenExtendedOperationHandlerRequest(val *AddDeliverPasswordResetTokenExtendedOperationHandlerRequest) *NullableAddDeliverPasswordResetTokenExtendedOperationHandlerRequest {
	return &NullableAddDeliverPasswordResetTokenExtendedOperationHandlerRequest{value: val, isSet: true}
}

func (v NullableAddDeliverPasswordResetTokenExtendedOperationHandlerRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddDeliverPasswordResetTokenExtendedOperationHandlerRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
