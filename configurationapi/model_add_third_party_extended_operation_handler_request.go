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

// checks if the AddThirdPartyExtendedOperationHandlerRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddThirdPartyExtendedOperationHandlerRequest{}

// AddThirdPartyExtendedOperationHandlerRequest struct for AddThirdPartyExtendedOperationHandlerRequest
type AddThirdPartyExtendedOperationHandlerRequest struct {
	// Name of the new Extended Operation Handler
	HandlerName string                                            `json:"handlerName"`
	Schemas     []EnumthirdPartyExtendedOperationHandlerSchemaUrn `json:"schemas"`
	// The fully-qualified name of the Java class providing the logic for the Third Party Extended Operation Handler.
	ExtensionClass string `json:"extensionClass"`
	// The set of arguments used to customize the behavior for the Third Party Extended Operation Handler. Each configuration property should be given in the form 'name=value'.
	ExtensionArgument []string `json:"extensionArgument,omitempty"`
	// A description for this Extended Operation Handler
	Description *string `json:"description,omitempty"`
	// Indicates whether the Extended Operation Handler is enabled (that is, whether the types of extended operations are allowed in the server).
	Enabled bool `json:"enabled"`
}

// NewAddThirdPartyExtendedOperationHandlerRequest instantiates a new AddThirdPartyExtendedOperationHandlerRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddThirdPartyExtendedOperationHandlerRequest(handlerName string, schemas []EnumthirdPartyExtendedOperationHandlerSchemaUrn, extensionClass string, enabled bool) *AddThirdPartyExtendedOperationHandlerRequest {
	this := AddThirdPartyExtendedOperationHandlerRequest{}
	this.HandlerName = handlerName
	this.Schemas = schemas
	this.ExtensionClass = extensionClass
	this.Enabled = enabled
	return &this
}

// NewAddThirdPartyExtendedOperationHandlerRequestWithDefaults instantiates a new AddThirdPartyExtendedOperationHandlerRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddThirdPartyExtendedOperationHandlerRequestWithDefaults() *AddThirdPartyExtendedOperationHandlerRequest {
	this := AddThirdPartyExtendedOperationHandlerRequest{}
	return &this
}

// GetHandlerName returns the HandlerName field value
func (o *AddThirdPartyExtendedOperationHandlerRequest) GetHandlerName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HandlerName
}

// GetHandlerNameOk returns a tuple with the HandlerName field value
// and a boolean to check if the value has been set.
func (o *AddThirdPartyExtendedOperationHandlerRequest) GetHandlerNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HandlerName, true
}

// SetHandlerName sets field value
func (o *AddThirdPartyExtendedOperationHandlerRequest) SetHandlerName(v string) {
	o.HandlerName = v
}

// GetSchemas returns the Schemas field value
func (o *AddThirdPartyExtendedOperationHandlerRequest) GetSchemas() []EnumthirdPartyExtendedOperationHandlerSchemaUrn {
	if o == nil {
		var ret []EnumthirdPartyExtendedOperationHandlerSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddThirdPartyExtendedOperationHandlerRequest) GetSchemasOk() ([]EnumthirdPartyExtendedOperationHandlerSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddThirdPartyExtendedOperationHandlerRequest) SetSchemas(v []EnumthirdPartyExtendedOperationHandlerSchemaUrn) {
	o.Schemas = v
}

// GetExtensionClass returns the ExtensionClass field value
func (o *AddThirdPartyExtendedOperationHandlerRequest) GetExtensionClass() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExtensionClass
}

// GetExtensionClassOk returns a tuple with the ExtensionClass field value
// and a boolean to check if the value has been set.
func (o *AddThirdPartyExtendedOperationHandlerRequest) GetExtensionClassOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExtensionClass, true
}

// SetExtensionClass sets field value
func (o *AddThirdPartyExtendedOperationHandlerRequest) SetExtensionClass(v string) {
	o.ExtensionClass = v
}

// GetExtensionArgument returns the ExtensionArgument field value if set, zero value otherwise.
func (o *AddThirdPartyExtendedOperationHandlerRequest) GetExtensionArgument() []string {
	if o == nil || IsNil(o.ExtensionArgument) {
		var ret []string
		return ret
	}
	return o.ExtensionArgument
}

// GetExtensionArgumentOk returns a tuple with the ExtensionArgument field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddThirdPartyExtendedOperationHandlerRequest) GetExtensionArgumentOk() ([]string, bool) {
	if o == nil || IsNil(o.ExtensionArgument) {
		return nil, false
	}
	return o.ExtensionArgument, true
}

// HasExtensionArgument returns a boolean if a field has been set.
func (o *AddThirdPartyExtendedOperationHandlerRequest) HasExtensionArgument() bool {
	if o != nil && !IsNil(o.ExtensionArgument) {
		return true
	}

	return false
}

// SetExtensionArgument gets a reference to the given []string and assigns it to the ExtensionArgument field.
func (o *AddThirdPartyExtendedOperationHandlerRequest) SetExtensionArgument(v []string) {
	o.ExtensionArgument = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddThirdPartyExtendedOperationHandlerRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddThirdPartyExtendedOperationHandlerRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddThirdPartyExtendedOperationHandlerRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddThirdPartyExtendedOperationHandlerRequest) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *AddThirdPartyExtendedOperationHandlerRequest) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AddThirdPartyExtendedOperationHandlerRequest) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AddThirdPartyExtendedOperationHandlerRequest) SetEnabled(v bool) {
	o.Enabled = v
}

func (o AddThirdPartyExtendedOperationHandlerRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddThirdPartyExtendedOperationHandlerRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["handlerName"] = o.HandlerName
	toSerialize["schemas"] = o.Schemas
	toSerialize["extensionClass"] = o.ExtensionClass
	if !IsNil(o.ExtensionArgument) {
		toSerialize["extensionArgument"] = o.ExtensionArgument
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["enabled"] = o.Enabled
	return toSerialize, nil
}

type NullableAddThirdPartyExtendedOperationHandlerRequest struct {
	value *AddThirdPartyExtendedOperationHandlerRequest
	isSet bool
}

func (v NullableAddThirdPartyExtendedOperationHandlerRequest) Get() *AddThirdPartyExtendedOperationHandlerRequest {
	return v.value
}

func (v *NullableAddThirdPartyExtendedOperationHandlerRequest) Set(val *AddThirdPartyExtendedOperationHandlerRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddThirdPartyExtendedOperationHandlerRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddThirdPartyExtendedOperationHandlerRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddThirdPartyExtendedOperationHandlerRequest(val *AddThirdPartyExtendedOperationHandlerRequest) *NullableAddThirdPartyExtendedOperationHandlerRequest {
	return &NullableAddThirdPartyExtendedOperationHandlerRequest{value: val, isSet: true}
}

func (v NullableAddThirdPartyExtendedOperationHandlerRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddThirdPartyExtendedOperationHandlerRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
