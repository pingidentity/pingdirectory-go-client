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

// AddThirdPartyOtpDeliveryMechanismRequest struct for AddThirdPartyOtpDeliveryMechanismRequest
type AddThirdPartyOtpDeliveryMechanismRequest struct {
	// Name of the new OTP Delivery Mechanism
	MechanismName string                                        `json:"mechanismName"`
	Schemas       []EnumthirdPartyOtpDeliveryMechanismSchemaUrn `json:"schemas"`
	// The fully-qualified name of the Java class providing the logic for the Third Party OTP Delivery Mechanism.
	ExtensionClass string `json:"extensionClass"`
	// The set of arguments used to customize the behavior for the Third Party OTP Delivery Mechanism. Each configuration property should be given in the form 'name=value'.
	ExtensionArgument []string `json:"extensionArgument,omitempty"`
	// A description for this OTP Delivery Mechanism
	Description *string `json:"description,omitempty"`
	// Indicates whether this OTP Delivery Mechanism is enabled for use in the server.
	Enabled bool `json:"enabled"`
}

// NewAddThirdPartyOtpDeliveryMechanismRequest instantiates a new AddThirdPartyOtpDeliveryMechanismRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddThirdPartyOtpDeliveryMechanismRequest(mechanismName string, schemas []EnumthirdPartyOtpDeliveryMechanismSchemaUrn, extensionClass string, enabled bool) *AddThirdPartyOtpDeliveryMechanismRequest {
	this := AddThirdPartyOtpDeliveryMechanismRequest{}
	this.MechanismName = mechanismName
	this.Schemas = schemas
	this.ExtensionClass = extensionClass
	this.Enabled = enabled
	return &this
}

// NewAddThirdPartyOtpDeliveryMechanismRequestWithDefaults instantiates a new AddThirdPartyOtpDeliveryMechanismRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddThirdPartyOtpDeliveryMechanismRequestWithDefaults() *AddThirdPartyOtpDeliveryMechanismRequest {
	this := AddThirdPartyOtpDeliveryMechanismRequest{}
	return &this
}

// GetMechanismName returns the MechanismName field value
func (o *AddThirdPartyOtpDeliveryMechanismRequest) GetMechanismName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MechanismName
}

// GetMechanismNameOk returns a tuple with the MechanismName field value
// and a boolean to check if the value has been set.
func (o *AddThirdPartyOtpDeliveryMechanismRequest) GetMechanismNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MechanismName, true
}

// SetMechanismName sets field value
func (o *AddThirdPartyOtpDeliveryMechanismRequest) SetMechanismName(v string) {
	o.MechanismName = v
}

// GetSchemas returns the Schemas field value
func (o *AddThirdPartyOtpDeliveryMechanismRequest) GetSchemas() []EnumthirdPartyOtpDeliveryMechanismSchemaUrn {
	if o == nil {
		var ret []EnumthirdPartyOtpDeliveryMechanismSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddThirdPartyOtpDeliveryMechanismRequest) GetSchemasOk() ([]EnumthirdPartyOtpDeliveryMechanismSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddThirdPartyOtpDeliveryMechanismRequest) SetSchemas(v []EnumthirdPartyOtpDeliveryMechanismSchemaUrn) {
	o.Schemas = v
}

// GetExtensionClass returns the ExtensionClass field value
func (o *AddThirdPartyOtpDeliveryMechanismRequest) GetExtensionClass() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExtensionClass
}

// GetExtensionClassOk returns a tuple with the ExtensionClass field value
// and a boolean to check if the value has been set.
func (o *AddThirdPartyOtpDeliveryMechanismRequest) GetExtensionClassOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExtensionClass, true
}

// SetExtensionClass sets field value
func (o *AddThirdPartyOtpDeliveryMechanismRequest) SetExtensionClass(v string) {
	o.ExtensionClass = v
}

// GetExtensionArgument returns the ExtensionArgument field value if set, zero value otherwise.
func (o *AddThirdPartyOtpDeliveryMechanismRequest) GetExtensionArgument() []string {
	if o == nil || isNil(o.ExtensionArgument) {
		var ret []string
		return ret
	}
	return o.ExtensionArgument
}

// GetExtensionArgumentOk returns a tuple with the ExtensionArgument field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddThirdPartyOtpDeliveryMechanismRequest) GetExtensionArgumentOk() ([]string, bool) {
	if o == nil || isNil(o.ExtensionArgument) {
		return nil, false
	}
	return o.ExtensionArgument, true
}

// HasExtensionArgument returns a boolean if a field has been set.
func (o *AddThirdPartyOtpDeliveryMechanismRequest) HasExtensionArgument() bool {
	if o != nil && !isNil(o.ExtensionArgument) {
		return true
	}

	return false
}

// SetExtensionArgument gets a reference to the given []string and assigns it to the ExtensionArgument field.
func (o *AddThirdPartyOtpDeliveryMechanismRequest) SetExtensionArgument(v []string) {
	o.ExtensionArgument = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddThirdPartyOtpDeliveryMechanismRequest) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddThirdPartyOtpDeliveryMechanismRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddThirdPartyOtpDeliveryMechanismRequest) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddThirdPartyOtpDeliveryMechanismRequest) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *AddThirdPartyOtpDeliveryMechanismRequest) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AddThirdPartyOtpDeliveryMechanismRequest) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AddThirdPartyOtpDeliveryMechanismRequest) SetEnabled(v bool) {
	o.Enabled = v
}

func (o AddThirdPartyOtpDeliveryMechanismRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["mechanismName"] = o.MechanismName
	}
	if true {
		toSerialize["schemas"] = o.Schemas
	}
	if true {
		toSerialize["extensionClass"] = o.ExtensionClass
	}
	if !isNil(o.ExtensionArgument) {
		toSerialize["extensionArgument"] = o.ExtensionArgument
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableAddThirdPartyOtpDeliveryMechanismRequest struct {
	value *AddThirdPartyOtpDeliveryMechanismRequest
	isSet bool
}

func (v NullableAddThirdPartyOtpDeliveryMechanismRequest) Get() *AddThirdPartyOtpDeliveryMechanismRequest {
	return v.value
}

func (v *NullableAddThirdPartyOtpDeliveryMechanismRequest) Set(val *AddThirdPartyOtpDeliveryMechanismRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddThirdPartyOtpDeliveryMechanismRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddThirdPartyOtpDeliveryMechanismRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddThirdPartyOtpDeliveryMechanismRequest(val *AddThirdPartyOtpDeliveryMechanismRequest) *NullableAddThirdPartyOtpDeliveryMechanismRequest {
	return &NullableAddThirdPartyOtpDeliveryMechanismRequest{value: val, isSet: true}
}

func (v NullableAddThirdPartyOtpDeliveryMechanismRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddThirdPartyOtpDeliveryMechanismRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
