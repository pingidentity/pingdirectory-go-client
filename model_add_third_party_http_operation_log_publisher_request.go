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

// AddThirdPartyHttpOperationLogPublisherRequest struct for AddThirdPartyHttpOperationLogPublisherRequest
type AddThirdPartyHttpOperationLogPublisherRequest struct {
	// Name of the new Log Publisher
	PublisherName string                                             `json:"publisherName"`
	Schemas       []EnumthirdPartyHttpOperationLogPublisherSchemaUrn `json:"schemas"`
	// The fully-qualified name of the Java class providing the logic for the Third Party HTTP Operation Log Publisher.
	ExtensionClass string `json:"extensionClass"`
	// The set of arguments used to customize the behavior for the Third Party HTTP Operation Log Publisher. Each configuration property should be given in the form 'name=value'.
	ExtensionArgument []string `json:"extensionArgument,omitempty"`
	// A description for this Log Publisher
	Description *string `json:"description,omitempty"`
	// Indicates whether the Log Publisher is enabled for use.
	Enabled              bool                                      `json:"enabled"`
	LoggingErrorBehavior *EnumlogPublisherLoggingErrorBehaviorProp `json:"loggingErrorBehavior,omitempty"`
}

// NewAddThirdPartyHttpOperationLogPublisherRequest instantiates a new AddThirdPartyHttpOperationLogPublisherRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddThirdPartyHttpOperationLogPublisherRequest(publisherName string, schemas []EnumthirdPartyHttpOperationLogPublisherSchemaUrn, extensionClass string, enabled bool) *AddThirdPartyHttpOperationLogPublisherRequest {
	this := AddThirdPartyHttpOperationLogPublisherRequest{}
	this.PublisherName = publisherName
	this.Schemas = schemas
	this.ExtensionClass = extensionClass
	this.Enabled = enabled
	return &this
}

// NewAddThirdPartyHttpOperationLogPublisherRequestWithDefaults instantiates a new AddThirdPartyHttpOperationLogPublisherRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddThirdPartyHttpOperationLogPublisherRequestWithDefaults() *AddThirdPartyHttpOperationLogPublisherRequest {
	this := AddThirdPartyHttpOperationLogPublisherRequest{}
	return &this
}

// GetPublisherName returns the PublisherName field value
func (o *AddThirdPartyHttpOperationLogPublisherRequest) GetPublisherName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PublisherName
}

// GetPublisherNameOk returns a tuple with the PublisherName field value
// and a boolean to check if the value has been set.
func (o *AddThirdPartyHttpOperationLogPublisherRequest) GetPublisherNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PublisherName, true
}

// SetPublisherName sets field value
func (o *AddThirdPartyHttpOperationLogPublisherRequest) SetPublisherName(v string) {
	o.PublisherName = v
}

// GetSchemas returns the Schemas field value
func (o *AddThirdPartyHttpOperationLogPublisherRequest) GetSchemas() []EnumthirdPartyHttpOperationLogPublisherSchemaUrn {
	if o == nil {
		var ret []EnumthirdPartyHttpOperationLogPublisherSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddThirdPartyHttpOperationLogPublisherRequest) GetSchemasOk() ([]EnumthirdPartyHttpOperationLogPublisherSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddThirdPartyHttpOperationLogPublisherRequest) SetSchemas(v []EnumthirdPartyHttpOperationLogPublisherSchemaUrn) {
	o.Schemas = v
}

// GetExtensionClass returns the ExtensionClass field value
func (o *AddThirdPartyHttpOperationLogPublisherRequest) GetExtensionClass() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExtensionClass
}

// GetExtensionClassOk returns a tuple with the ExtensionClass field value
// and a boolean to check if the value has been set.
func (o *AddThirdPartyHttpOperationLogPublisherRequest) GetExtensionClassOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExtensionClass, true
}

// SetExtensionClass sets field value
func (o *AddThirdPartyHttpOperationLogPublisherRequest) SetExtensionClass(v string) {
	o.ExtensionClass = v
}

// GetExtensionArgument returns the ExtensionArgument field value if set, zero value otherwise.
func (o *AddThirdPartyHttpOperationLogPublisherRequest) GetExtensionArgument() []string {
	if o == nil || isNil(o.ExtensionArgument) {
		var ret []string
		return ret
	}
	return o.ExtensionArgument
}

// GetExtensionArgumentOk returns a tuple with the ExtensionArgument field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddThirdPartyHttpOperationLogPublisherRequest) GetExtensionArgumentOk() ([]string, bool) {
	if o == nil || isNil(o.ExtensionArgument) {
		return nil, false
	}
	return o.ExtensionArgument, true
}

// HasExtensionArgument returns a boolean if a field has been set.
func (o *AddThirdPartyHttpOperationLogPublisherRequest) HasExtensionArgument() bool {
	if o != nil && !isNil(o.ExtensionArgument) {
		return true
	}

	return false
}

// SetExtensionArgument gets a reference to the given []string and assigns it to the ExtensionArgument field.
func (o *AddThirdPartyHttpOperationLogPublisherRequest) SetExtensionArgument(v []string) {
	o.ExtensionArgument = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddThirdPartyHttpOperationLogPublisherRequest) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddThirdPartyHttpOperationLogPublisherRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddThirdPartyHttpOperationLogPublisherRequest) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddThirdPartyHttpOperationLogPublisherRequest) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *AddThirdPartyHttpOperationLogPublisherRequest) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AddThirdPartyHttpOperationLogPublisherRequest) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AddThirdPartyHttpOperationLogPublisherRequest) SetEnabled(v bool) {
	o.Enabled = v
}

// GetLoggingErrorBehavior returns the LoggingErrorBehavior field value if set, zero value otherwise.
func (o *AddThirdPartyHttpOperationLogPublisherRequest) GetLoggingErrorBehavior() EnumlogPublisherLoggingErrorBehaviorProp {
	if o == nil || isNil(o.LoggingErrorBehavior) {
		var ret EnumlogPublisherLoggingErrorBehaviorProp
		return ret
	}
	return *o.LoggingErrorBehavior
}

// GetLoggingErrorBehaviorOk returns a tuple with the LoggingErrorBehavior field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddThirdPartyHttpOperationLogPublisherRequest) GetLoggingErrorBehaviorOk() (*EnumlogPublisherLoggingErrorBehaviorProp, bool) {
	if o == nil || isNil(o.LoggingErrorBehavior) {
		return nil, false
	}
	return o.LoggingErrorBehavior, true
}

// HasLoggingErrorBehavior returns a boolean if a field has been set.
func (o *AddThirdPartyHttpOperationLogPublisherRequest) HasLoggingErrorBehavior() bool {
	if o != nil && !isNil(o.LoggingErrorBehavior) {
		return true
	}

	return false
}

// SetLoggingErrorBehavior gets a reference to the given EnumlogPublisherLoggingErrorBehaviorProp and assigns it to the LoggingErrorBehavior field.
func (o *AddThirdPartyHttpOperationLogPublisherRequest) SetLoggingErrorBehavior(v EnumlogPublisherLoggingErrorBehaviorProp) {
	o.LoggingErrorBehavior = &v
}

func (o AddThirdPartyHttpOperationLogPublisherRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["publisherName"] = o.PublisherName
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
	if !isNil(o.LoggingErrorBehavior) {
		toSerialize["loggingErrorBehavior"] = o.LoggingErrorBehavior
	}
	return json.Marshal(toSerialize)
}

type NullableAddThirdPartyHttpOperationLogPublisherRequest struct {
	value *AddThirdPartyHttpOperationLogPublisherRequest
	isSet bool
}

func (v NullableAddThirdPartyHttpOperationLogPublisherRequest) Get() *AddThirdPartyHttpOperationLogPublisherRequest {
	return v.value
}

func (v *NullableAddThirdPartyHttpOperationLogPublisherRequest) Set(val *AddThirdPartyHttpOperationLogPublisherRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddThirdPartyHttpOperationLogPublisherRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddThirdPartyHttpOperationLogPublisherRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddThirdPartyHttpOperationLogPublisherRequest(val *AddThirdPartyHttpOperationLogPublisherRequest) *NullableAddThirdPartyHttpOperationLogPublisherRequest {
	return &NullableAddThirdPartyHttpOperationLogPublisherRequest{value: val, isSet: true}
}

func (v NullableAddThirdPartyHttpOperationLogPublisherRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddThirdPartyHttpOperationLogPublisherRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
