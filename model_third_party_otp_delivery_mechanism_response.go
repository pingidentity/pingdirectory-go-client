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

// ThirdPartyOtpDeliveryMechanismResponse struct for ThirdPartyOtpDeliveryMechanismResponse
type ThirdPartyOtpDeliveryMechanismResponse struct {
	// Name of the OTP Delivery Mechanism
	Id string `json:"id"`
	Schemas []EnumthirdPartyOtpDeliveryMechanismSchemaUrn `json:"schemas"`
	// The fully-qualified name of the Java class providing the logic for the Third Party OTP Delivery Mechanism.
	ExtensionClass string `json:"extensionClass"`
	ExtensionArgument []string `json:"extensionArgument,omitempty"`
	// A description for this OTP Delivery Mechanism
	Description *string `json:"description,omitempty"`
	// Indicates whether this OTP Delivery Mechanism is enabled for use in the server.
	Enabled bool `json:"enabled"`
}

// NewThirdPartyOtpDeliveryMechanismResponse instantiates a new ThirdPartyOtpDeliveryMechanismResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThirdPartyOtpDeliveryMechanismResponse(id string, schemas []EnumthirdPartyOtpDeliveryMechanismSchemaUrn, extensionClass string, enabled bool) *ThirdPartyOtpDeliveryMechanismResponse {
	this := ThirdPartyOtpDeliveryMechanismResponse{}
	this.Id = id
	this.Schemas = schemas
	this.ExtensionClass = extensionClass
	this.Enabled = enabled
	return &this
}

// NewThirdPartyOtpDeliveryMechanismResponseWithDefaults instantiates a new ThirdPartyOtpDeliveryMechanismResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThirdPartyOtpDeliveryMechanismResponseWithDefaults() *ThirdPartyOtpDeliveryMechanismResponse {
	this := ThirdPartyOtpDeliveryMechanismResponse{}
	return &this
}

// GetId returns the Id field value
func (o *ThirdPartyOtpDeliveryMechanismResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ThirdPartyOtpDeliveryMechanismResponse) GetIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ThirdPartyOtpDeliveryMechanismResponse) SetId(v string) {
	o.Id = v
}

// GetSchemas returns the Schemas field value
func (o *ThirdPartyOtpDeliveryMechanismResponse) GetSchemas() []EnumthirdPartyOtpDeliveryMechanismSchemaUrn {
	if o == nil {
		var ret []EnumthirdPartyOtpDeliveryMechanismSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *ThirdPartyOtpDeliveryMechanismResponse) GetSchemasOk() ([]EnumthirdPartyOtpDeliveryMechanismSchemaUrn, bool) {
	if o == nil {
    return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *ThirdPartyOtpDeliveryMechanismResponse) SetSchemas(v []EnumthirdPartyOtpDeliveryMechanismSchemaUrn) {
	o.Schemas = v
}

// GetExtensionClass returns the ExtensionClass field value
func (o *ThirdPartyOtpDeliveryMechanismResponse) GetExtensionClass() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExtensionClass
}

// GetExtensionClassOk returns a tuple with the ExtensionClass field value
// and a boolean to check if the value has been set.
func (o *ThirdPartyOtpDeliveryMechanismResponse) GetExtensionClassOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ExtensionClass, true
}

// SetExtensionClass sets field value
func (o *ThirdPartyOtpDeliveryMechanismResponse) SetExtensionClass(v string) {
	o.ExtensionClass = v
}

// GetExtensionArgument returns the ExtensionArgument field value if set, zero value otherwise.
func (o *ThirdPartyOtpDeliveryMechanismResponse) GetExtensionArgument() []string {
	if o == nil || isNil(o.ExtensionArgument) {
		var ret []string
		return ret
	}
	return o.ExtensionArgument
}

// GetExtensionArgumentOk returns a tuple with the ExtensionArgument field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyOtpDeliveryMechanismResponse) GetExtensionArgumentOk() ([]string, bool) {
	if o == nil || isNil(o.ExtensionArgument) {
    return nil, false
	}
	return o.ExtensionArgument, true
}

// HasExtensionArgument returns a boolean if a field has been set.
func (o *ThirdPartyOtpDeliveryMechanismResponse) HasExtensionArgument() bool {
	if o != nil && !isNil(o.ExtensionArgument) {
		return true
	}

	return false
}

// SetExtensionArgument gets a reference to the given []string and assigns it to the ExtensionArgument field.
func (o *ThirdPartyOtpDeliveryMechanismResponse) SetExtensionArgument(v []string) {
	o.ExtensionArgument = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ThirdPartyOtpDeliveryMechanismResponse) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyOtpDeliveryMechanismResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ThirdPartyOtpDeliveryMechanismResponse) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ThirdPartyOtpDeliveryMechanismResponse) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *ThirdPartyOtpDeliveryMechanismResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *ThirdPartyOtpDeliveryMechanismResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *ThirdPartyOtpDeliveryMechanismResponse) SetEnabled(v bool) {
	o.Enabled = v
}

func (o ThirdPartyOtpDeliveryMechanismResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
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

type NullableThirdPartyOtpDeliveryMechanismResponse struct {
	value *ThirdPartyOtpDeliveryMechanismResponse
	isSet bool
}

func (v NullableThirdPartyOtpDeliveryMechanismResponse) Get() *ThirdPartyOtpDeliveryMechanismResponse {
	return v.value
}

func (v *NullableThirdPartyOtpDeliveryMechanismResponse) Set(val *ThirdPartyOtpDeliveryMechanismResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableThirdPartyOtpDeliveryMechanismResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableThirdPartyOtpDeliveryMechanismResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThirdPartyOtpDeliveryMechanismResponse(val *ThirdPartyOtpDeliveryMechanismResponse) *NullableThirdPartyOtpDeliveryMechanismResponse {
	return &NullableThirdPartyOtpDeliveryMechanismResponse{value: val, isSet: true}
}

func (v NullableThirdPartyOtpDeliveryMechanismResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThirdPartyOtpDeliveryMechanismResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


