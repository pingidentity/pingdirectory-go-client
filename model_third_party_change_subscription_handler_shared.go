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

// ThirdPartyChangeSubscriptionHandlerShared struct for ThirdPartyChangeSubscriptionHandlerShared
type ThirdPartyChangeSubscriptionHandlerShared struct {
	Schemas []EnumthirdPartyChangeSubscriptionHandlerSchemaUrn `json:"schemas"`
	// The fully-qualified name of the Java class providing the logic for the Third Party Change Subscription Handler.
	ExtensionClass string `json:"extensionClass"`
	ExtensionArgument []string `json:"extensionArgument,omitempty"`
	// A description for this Change Subscription Handler
	Description *string `json:"description,omitempty"`
	// Indicates whether this change subscription handler is enabled within the server.
	Enabled bool `json:"enabled"`
	ChangeSubscription []string `json:"changeSubscription,omitempty"`
}

// NewThirdPartyChangeSubscriptionHandlerShared instantiates a new ThirdPartyChangeSubscriptionHandlerShared object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThirdPartyChangeSubscriptionHandlerShared(schemas []EnumthirdPartyChangeSubscriptionHandlerSchemaUrn, extensionClass string, enabled bool) *ThirdPartyChangeSubscriptionHandlerShared {
	this := ThirdPartyChangeSubscriptionHandlerShared{}
	this.Schemas = schemas
	this.ExtensionClass = extensionClass
	this.Enabled = enabled
	return &this
}

// NewThirdPartyChangeSubscriptionHandlerSharedWithDefaults instantiates a new ThirdPartyChangeSubscriptionHandlerShared object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThirdPartyChangeSubscriptionHandlerSharedWithDefaults() *ThirdPartyChangeSubscriptionHandlerShared {
	this := ThirdPartyChangeSubscriptionHandlerShared{}
	return &this
}

// GetSchemas returns the Schemas field value
func (o *ThirdPartyChangeSubscriptionHandlerShared) GetSchemas() []EnumthirdPartyChangeSubscriptionHandlerSchemaUrn {
	if o == nil {
		var ret []EnumthirdPartyChangeSubscriptionHandlerSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *ThirdPartyChangeSubscriptionHandlerShared) GetSchemasOk() ([]EnumthirdPartyChangeSubscriptionHandlerSchemaUrn, bool) {
	if o == nil {
    return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *ThirdPartyChangeSubscriptionHandlerShared) SetSchemas(v []EnumthirdPartyChangeSubscriptionHandlerSchemaUrn) {
	o.Schemas = v
}

// GetExtensionClass returns the ExtensionClass field value
func (o *ThirdPartyChangeSubscriptionHandlerShared) GetExtensionClass() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExtensionClass
}

// GetExtensionClassOk returns a tuple with the ExtensionClass field value
// and a boolean to check if the value has been set.
func (o *ThirdPartyChangeSubscriptionHandlerShared) GetExtensionClassOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ExtensionClass, true
}

// SetExtensionClass sets field value
func (o *ThirdPartyChangeSubscriptionHandlerShared) SetExtensionClass(v string) {
	o.ExtensionClass = v
}

// GetExtensionArgument returns the ExtensionArgument field value if set, zero value otherwise.
func (o *ThirdPartyChangeSubscriptionHandlerShared) GetExtensionArgument() []string {
	if o == nil || isNil(o.ExtensionArgument) {
		var ret []string
		return ret
	}
	return o.ExtensionArgument
}

// GetExtensionArgumentOk returns a tuple with the ExtensionArgument field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyChangeSubscriptionHandlerShared) GetExtensionArgumentOk() ([]string, bool) {
	if o == nil || isNil(o.ExtensionArgument) {
    return nil, false
	}
	return o.ExtensionArgument, true
}

// HasExtensionArgument returns a boolean if a field has been set.
func (o *ThirdPartyChangeSubscriptionHandlerShared) HasExtensionArgument() bool {
	if o != nil && !isNil(o.ExtensionArgument) {
		return true
	}

	return false
}

// SetExtensionArgument gets a reference to the given []string and assigns it to the ExtensionArgument field.
func (o *ThirdPartyChangeSubscriptionHandlerShared) SetExtensionArgument(v []string) {
	o.ExtensionArgument = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ThirdPartyChangeSubscriptionHandlerShared) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyChangeSubscriptionHandlerShared) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ThirdPartyChangeSubscriptionHandlerShared) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ThirdPartyChangeSubscriptionHandlerShared) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *ThirdPartyChangeSubscriptionHandlerShared) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *ThirdPartyChangeSubscriptionHandlerShared) GetEnabledOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *ThirdPartyChangeSubscriptionHandlerShared) SetEnabled(v bool) {
	o.Enabled = v
}

// GetChangeSubscription returns the ChangeSubscription field value if set, zero value otherwise.
func (o *ThirdPartyChangeSubscriptionHandlerShared) GetChangeSubscription() []string {
	if o == nil || isNil(o.ChangeSubscription) {
		var ret []string
		return ret
	}
	return o.ChangeSubscription
}

// GetChangeSubscriptionOk returns a tuple with the ChangeSubscription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyChangeSubscriptionHandlerShared) GetChangeSubscriptionOk() ([]string, bool) {
	if o == nil || isNil(o.ChangeSubscription) {
    return nil, false
	}
	return o.ChangeSubscription, true
}

// HasChangeSubscription returns a boolean if a field has been set.
func (o *ThirdPartyChangeSubscriptionHandlerShared) HasChangeSubscription() bool {
	if o != nil && !isNil(o.ChangeSubscription) {
		return true
	}

	return false
}

// SetChangeSubscription gets a reference to the given []string and assigns it to the ChangeSubscription field.
func (o *ThirdPartyChangeSubscriptionHandlerShared) SetChangeSubscription(v []string) {
	o.ChangeSubscription = v
}

func (o ThirdPartyChangeSubscriptionHandlerShared) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
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
	if !isNil(o.ChangeSubscription) {
		toSerialize["changeSubscription"] = o.ChangeSubscription
	}
	return json.Marshal(toSerialize)
}

type NullableThirdPartyChangeSubscriptionHandlerShared struct {
	value *ThirdPartyChangeSubscriptionHandlerShared
	isSet bool
}

func (v NullableThirdPartyChangeSubscriptionHandlerShared) Get() *ThirdPartyChangeSubscriptionHandlerShared {
	return v.value
}

func (v *NullableThirdPartyChangeSubscriptionHandlerShared) Set(val *ThirdPartyChangeSubscriptionHandlerShared) {
	v.value = val
	v.isSet = true
}

func (v NullableThirdPartyChangeSubscriptionHandlerShared) IsSet() bool {
	return v.isSet
}

func (v *NullableThirdPartyChangeSubscriptionHandlerShared) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThirdPartyChangeSubscriptionHandlerShared(val *ThirdPartyChangeSubscriptionHandlerShared) *NullableThirdPartyChangeSubscriptionHandlerShared {
	return &NullableThirdPartyChangeSubscriptionHandlerShared{value: val, isSet: true}
}

func (v NullableThirdPartyChangeSubscriptionHandlerShared) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThirdPartyChangeSubscriptionHandlerShared) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


