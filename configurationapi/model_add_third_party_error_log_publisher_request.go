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

// checks if the AddThirdPartyErrorLogPublisherRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddThirdPartyErrorLogPublisherRequest{}

// AddThirdPartyErrorLogPublisherRequest struct for AddThirdPartyErrorLogPublisherRequest
type AddThirdPartyErrorLogPublisherRequest struct {
	Schemas []EnumthirdPartyErrorLogPublisherSchemaUrn `json:"schemas"`
	// The fully-qualified name of the Java class providing the logic for the Third Party Error Log Publisher.
	ExtensionClass string `json:"extensionClass"`
	// The set of arguments used to customize the behavior for the Third Party Error Log Publisher. Each configuration property should be given in the form 'name=value'.
	ExtensionArgument []string                              `json:"extensionArgument,omitempty"`
	DefaultSeverity   []EnumlogPublisherDefaultSeverityProp `json:"defaultSeverity,omitempty"`
	// Specifies the override severity levels for the logger based on the category of the messages.
	OverrideSeverity []string `json:"overrideSeverity,omitempty"`
	// A description for this Log Publisher
	Description *string `json:"description,omitempty"`
	// Indicates whether the Log Publisher is enabled for use.
	Enabled              bool                                      `json:"enabled"`
	LoggingErrorBehavior *EnumlogPublisherLoggingErrorBehaviorProp `json:"loggingErrorBehavior,omitempty"`
	// Name of the new Log Publisher
	PublisherName string `json:"publisherName"`
}

// NewAddThirdPartyErrorLogPublisherRequest instantiates a new AddThirdPartyErrorLogPublisherRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddThirdPartyErrorLogPublisherRequest(schemas []EnumthirdPartyErrorLogPublisherSchemaUrn, extensionClass string, enabled bool, publisherName string) *AddThirdPartyErrorLogPublisherRequest {
	this := AddThirdPartyErrorLogPublisherRequest{}
	this.Schemas = schemas
	this.ExtensionClass = extensionClass
	this.Enabled = enabled
	this.PublisherName = publisherName
	return &this
}

// NewAddThirdPartyErrorLogPublisherRequestWithDefaults instantiates a new AddThirdPartyErrorLogPublisherRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddThirdPartyErrorLogPublisherRequestWithDefaults() *AddThirdPartyErrorLogPublisherRequest {
	this := AddThirdPartyErrorLogPublisherRequest{}
	return &this
}

// GetSchemas returns the Schemas field value
func (o *AddThirdPartyErrorLogPublisherRequest) GetSchemas() []EnumthirdPartyErrorLogPublisherSchemaUrn {
	if o == nil {
		var ret []EnumthirdPartyErrorLogPublisherSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddThirdPartyErrorLogPublisherRequest) GetSchemasOk() ([]EnumthirdPartyErrorLogPublisherSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddThirdPartyErrorLogPublisherRequest) SetSchemas(v []EnumthirdPartyErrorLogPublisherSchemaUrn) {
	o.Schemas = v
}

// GetExtensionClass returns the ExtensionClass field value
func (o *AddThirdPartyErrorLogPublisherRequest) GetExtensionClass() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExtensionClass
}

// GetExtensionClassOk returns a tuple with the ExtensionClass field value
// and a boolean to check if the value has been set.
func (o *AddThirdPartyErrorLogPublisherRequest) GetExtensionClassOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExtensionClass, true
}

// SetExtensionClass sets field value
func (o *AddThirdPartyErrorLogPublisherRequest) SetExtensionClass(v string) {
	o.ExtensionClass = v
}

// GetExtensionArgument returns the ExtensionArgument field value if set, zero value otherwise.
func (o *AddThirdPartyErrorLogPublisherRequest) GetExtensionArgument() []string {
	if o == nil || IsNil(o.ExtensionArgument) {
		var ret []string
		return ret
	}
	return o.ExtensionArgument
}

// GetExtensionArgumentOk returns a tuple with the ExtensionArgument field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddThirdPartyErrorLogPublisherRequest) GetExtensionArgumentOk() ([]string, bool) {
	if o == nil || IsNil(o.ExtensionArgument) {
		return nil, false
	}
	return o.ExtensionArgument, true
}

// HasExtensionArgument returns a boolean if a field has been set.
func (o *AddThirdPartyErrorLogPublisherRequest) HasExtensionArgument() bool {
	if o != nil && !IsNil(o.ExtensionArgument) {
		return true
	}

	return false
}

// SetExtensionArgument gets a reference to the given []string and assigns it to the ExtensionArgument field.
func (o *AddThirdPartyErrorLogPublisherRequest) SetExtensionArgument(v []string) {
	o.ExtensionArgument = v
}

// GetDefaultSeverity returns the DefaultSeverity field value if set, zero value otherwise.
func (o *AddThirdPartyErrorLogPublisherRequest) GetDefaultSeverity() []EnumlogPublisherDefaultSeverityProp {
	if o == nil || IsNil(o.DefaultSeverity) {
		var ret []EnumlogPublisherDefaultSeverityProp
		return ret
	}
	return o.DefaultSeverity
}

// GetDefaultSeverityOk returns a tuple with the DefaultSeverity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddThirdPartyErrorLogPublisherRequest) GetDefaultSeverityOk() ([]EnumlogPublisherDefaultSeverityProp, bool) {
	if o == nil || IsNil(o.DefaultSeverity) {
		return nil, false
	}
	return o.DefaultSeverity, true
}

// HasDefaultSeverity returns a boolean if a field has been set.
func (o *AddThirdPartyErrorLogPublisherRequest) HasDefaultSeverity() bool {
	if o != nil && !IsNil(o.DefaultSeverity) {
		return true
	}

	return false
}

// SetDefaultSeverity gets a reference to the given []EnumlogPublisherDefaultSeverityProp and assigns it to the DefaultSeverity field.
func (o *AddThirdPartyErrorLogPublisherRequest) SetDefaultSeverity(v []EnumlogPublisherDefaultSeverityProp) {
	o.DefaultSeverity = v
}

// GetOverrideSeverity returns the OverrideSeverity field value if set, zero value otherwise.
func (o *AddThirdPartyErrorLogPublisherRequest) GetOverrideSeverity() []string {
	if o == nil || IsNil(o.OverrideSeverity) {
		var ret []string
		return ret
	}
	return o.OverrideSeverity
}

// GetOverrideSeverityOk returns a tuple with the OverrideSeverity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddThirdPartyErrorLogPublisherRequest) GetOverrideSeverityOk() ([]string, bool) {
	if o == nil || IsNil(o.OverrideSeverity) {
		return nil, false
	}
	return o.OverrideSeverity, true
}

// HasOverrideSeverity returns a boolean if a field has been set.
func (o *AddThirdPartyErrorLogPublisherRequest) HasOverrideSeverity() bool {
	if o != nil && !IsNil(o.OverrideSeverity) {
		return true
	}

	return false
}

// SetOverrideSeverity gets a reference to the given []string and assigns it to the OverrideSeverity field.
func (o *AddThirdPartyErrorLogPublisherRequest) SetOverrideSeverity(v []string) {
	o.OverrideSeverity = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddThirdPartyErrorLogPublisherRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddThirdPartyErrorLogPublisherRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddThirdPartyErrorLogPublisherRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddThirdPartyErrorLogPublisherRequest) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *AddThirdPartyErrorLogPublisherRequest) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AddThirdPartyErrorLogPublisherRequest) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AddThirdPartyErrorLogPublisherRequest) SetEnabled(v bool) {
	o.Enabled = v
}

// GetLoggingErrorBehavior returns the LoggingErrorBehavior field value if set, zero value otherwise.
func (o *AddThirdPartyErrorLogPublisherRequest) GetLoggingErrorBehavior() EnumlogPublisherLoggingErrorBehaviorProp {
	if o == nil || IsNil(o.LoggingErrorBehavior) {
		var ret EnumlogPublisherLoggingErrorBehaviorProp
		return ret
	}
	return *o.LoggingErrorBehavior
}

// GetLoggingErrorBehaviorOk returns a tuple with the LoggingErrorBehavior field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddThirdPartyErrorLogPublisherRequest) GetLoggingErrorBehaviorOk() (*EnumlogPublisherLoggingErrorBehaviorProp, bool) {
	if o == nil || IsNil(o.LoggingErrorBehavior) {
		return nil, false
	}
	return o.LoggingErrorBehavior, true
}

// HasLoggingErrorBehavior returns a boolean if a field has been set.
func (o *AddThirdPartyErrorLogPublisherRequest) HasLoggingErrorBehavior() bool {
	if o != nil && !IsNil(o.LoggingErrorBehavior) {
		return true
	}

	return false
}

// SetLoggingErrorBehavior gets a reference to the given EnumlogPublisherLoggingErrorBehaviorProp and assigns it to the LoggingErrorBehavior field.
func (o *AddThirdPartyErrorLogPublisherRequest) SetLoggingErrorBehavior(v EnumlogPublisherLoggingErrorBehaviorProp) {
	o.LoggingErrorBehavior = &v
}

// GetPublisherName returns the PublisherName field value
func (o *AddThirdPartyErrorLogPublisherRequest) GetPublisherName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PublisherName
}

// GetPublisherNameOk returns a tuple with the PublisherName field value
// and a boolean to check if the value has been set.
func (o *AddThirdPartyErrorLogPublisherRequest) GetPublisherNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PublisherName, true
}

// SetPublisherName sets field value
func (o *AddThirdPartyErrorLogPublisherRequest) SetPublisherName(v string) {
	o.PublisherName = v
}

func (o AddThirdPartyErrorLogPublisherRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddThirdPartyErrorLogPublisherRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["schemas"] = o.Schemas
	toSerialize["extensionClass"] = o.ExtensionClass
	if !IsNil(o.ExtensionArgument) {
		toSerialize["extensionArgument"] = o.ExtensionArgument
	}
	if !IsNil(o.DefaultSeverity) {
		toSerialize["defaultSeverity"] = o.DefaultSeverity
	}
	if !IsNil(o.OverrideSeverity) {
		toSerialize["overrideSeverity"] = o.OverrideSeverity
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["enabled"] = o.Enabled
	if !IsNil(o.LoggingErrorBehavior) {
		toSerialize["loggingErrorBehavior"] = o.LoggingErrorBehavior
	}
	toSerialize["publisherName"] = o.PublisherName
	return toSerialize, nil
}

type NullableAddThirdPartyErrorLogPublisherRequest struct {
	value *AddThirdPartyErrorLogPublisherRequest
	isSet bool
}

func (v NullableAddThirdPartyErrorLogPublisherRequest) Get() *AddThirdPartyErrorLogPublisherRequest {
	return v.value
}

func (v *NullableAddThirdPartyErrorLogPublisherRequest) Set(val *AddThirdPartyErrorLogPublisherRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddThirdPartyErrorLogPublisherRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddThirdPartyErrorLogPublisherRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddThirdPartyErrorLogPublisherRequest(val *AddThirdPartyErrorLogPublisherRequest) *NullableAddThirdPartyErrorLogPublisherRequest {
	return &NullableAddThirdPartyErrorLogPublisherRequest{value: val, isSet: true}
}

func (v NullableAddThirdPartyErrorLogPublisherRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddThirdPartyErrorLogPublisherRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
