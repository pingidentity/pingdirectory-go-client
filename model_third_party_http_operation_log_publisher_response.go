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

// ThirdPartyHttpOperationLogPublisherResponse struct for ThirdPartyHttpOperationLogPublisherResponse
type ThirdPartyHttpOperationLogPublisherResponse struct {
	// Name of the Log Publisher
	Id string `json:"id"`
	Schemas []EnumthirdPartyHttpOperationLogPublisherSchemaUrn `json:"schemas"`
	// The fully-qualified name of the Java class providing the logic for the Third Party HTTP Operation Log Publisher.
	ExtensionClass string `json:"extensionClass"`
	// The set of arguments used to customize the behavior for the Third Party HTTP Operation Log Publisher. Each configuration property should be given in the form 'name=value'.
	ExtensionArgument []string `json:"extensionArgument,omitempty"`
	// A description for this Log Publisher
	Description *string `json:"description,omitempty"`
	// Indicates whether the Log Publisher is enabled for use.
	Enabled bool `json:"enabled"`
	LoggingErrorBehavior *EnumlogPublisherLoggingErrorBehaviorProp `json:"loggingErrorBehavior,omitempty"`
	Meta *MetaMeta `json:"meta,omitempty"`
}

// NewThirdPartyHttpOperationLogPublisherResponse instantiates a new ThirdPartyHttpOperationLogPublisherResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThirdPartyHttpOperationLogPublisherResponse(id string, schemas []EnumthirdPartyHttpOperationLogPublisherSchemaUrn, extensionClass string, enabled bool) *ThirdPartyHttpOperationLogPublisherResponse {
	this := ThirdPartyHttpOperationLogPublisherResponse{}
	this.Id = id
	this.Schemas = schemas
	this.ExtensionClass = extensionClass
	this.Enabled = enabled
	return &this
}

// NewThirdPartyHttpOperationLogPublisherResponseWithDefaults instantiates a new ThirdPartyHttpOperationLogPublisherResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThirdPartyHttpOperationLogPublisherResponseWithDefaults() *ThirdPartyHttpOperationLogPublisherResponse {
	this := ThirdPartyHttpOperationLogPublisherResponse{}
	return &this
}

// GetId returns the Id field value
func (o *ThirdPartyHttpOperationLogPublisherResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ThirdPartyHttpOperationLogPublisherResponse) GetIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ThirdPartyHttpOperationLogPublisherResponse) SetId(v string) {
	o.Id = v
}

// GetSchemas returns the Schemas field value
func (o *ThirdPartyHttpOperationLogPublisherResponse) GetSchemas() []EnumthirdPartyHttpOperationLogPublisherSchemaUrn {
	if o == nil {
		var ret []EnumthirdPartyHttpOperationLogPublisherSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *ThirdPartyHttpOperationLogPublisherResponse) GetSchemasOk() ([]EnumthirdPartyHttpOperationLogPublisherSchemaUrn, bool) {
	if o == nil {
    return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *ThirdPartyHttpOperationLogPublisherResponse) SetSchemas(v []EnumthirdPartyHttpOperationLogPublisherSchemaUrn) {
	o.Schemas = v
}

// GetExtensionClass returns the ExtensionClass field value
func (o *ThirdPartyHttpOperationLogPublisherResponse) GetExtensionClass() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExtensionClass
}

// GetExtensionClassOk returns a tuple with the ExtensionClass field value
// and a boolean to check if the value has been set.
func (o *ThirdPartyHttpOperationLogPublisherResponse) GetExtensionClassOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ExtensionClass, true
}

// SetExtensionClass sets field value
func (o *ThirdPartyHttpOperationLogPublisherResponse) SetExtensionClass(v string) {
	o.ExtensionClass = v
}

// GetExtensionArgument returns the ExtensionArgument field value if set, zero value otherwise.
func (o *ThirdPartyHttpOperationLogPublisherResponse) GetExtensionArgument() []string {
	if o == nil || isNil(o.ExtensionArgument) {
		var ret []string
		return ret
	}
	return o.ExtensionArgument
}

// GetExtensionArgumentOk returns a tuple with the ExtensionArgument field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyHttpOperationLogPublisherResponse) GetExtensionArgumentOk() ([]string, bool) {
	if o == nil || isNil(o.ExtensionArgument) {
    return nil, false
	}
	return o.ExtensionArgument, true
}

// HasExtensionArgument returns a boolean if a field has been set.
func (o *ThirdPartyHttpOperationLogPublisherResponse) HasExtensionArgument() bool {
	if o != nil && !isNil(o.ExtensionArgument) {
		return true
	}

	return false
}

// SetExtensionArgument gets a reference to the given []string and assigns it to the ExtensionArgument field.
func (o *ThirdPartyHttpOperationLogPublisherResponse) SetExtensionArgument(v []string) {
	o.ExtensionArgument = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ThirdPartyHttpOperationLogPublisherResponse) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyHttpOperationLogPublisherResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ThirdPartyHttpOperationLogPublisherResponse) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ThirdPartyHttpOperationLogPublisherResponse) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *ThirdPartyHttpOperationLogPublisherResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *ThirdPartyHttpOperationLogPublisherResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *ThirdPartyHttpOperationLogPublisherResponse) SetEnabled(v bool) {
	o.Enabled = v
}

// GetLoggingErrorBehavior returns the LoggingErrorBehavior field value if set, zero value otherwise.
func (o *ThirdPartyHttpOperationLogPublisherResponse) GetLoggingErrorBehavior() EnumlogPublisherLoggingErrorBehaviorProp {
	if o == nil || isNil(o.LoggingErrorBehavior) {
		var ret EnumlogPublisherLoggingErrorBehaviorProp
		return ret
	}
	return *o.LoggingErrorBehavior
}

// GetLoggingErrorBehaviorOk returns a tuple with the LoggingErrorBehavior field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyHttpOperationLogPublisherResponse) GetLoggingErrorBehaviorOk() (*EnumlogPublisherLoggingErrorBehaviorProp, bool) {
	if o == nil || isNil(o.LoggingErrorBehavior) {
    return nil, false
	}
	return o.LoggingErrorBehavior, true
}

// HasLoggingErrorBehavior returns a boolean if a field has been set.
func (o *ThirdPartyHttpOperationLogPublisherResponse) HasLoggingErrorBehavior() bool {
	if o != nil && !isNil(o.LoggingErrorBehavior) {
		return true
	}

	return false
}

// SetLoggingErrorBehavior gets a reference to the given EnumlogPublisherLoggingErrorBehaviorProp and assigns it to the LoggingErrorBehavior field.
func (o *ThirdPartyHttpOperationLogPublisherResponse) SetLoggingErrorBehavior(v EnumlogPublisherLoggingErrorBehaviorProp) {
	o.LoggingErrorBehavior = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ThirdPartyHttpOperationLogPublisherResponse) GetMeta() MetaMeta {
	if o == nil || isNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyHttpOperationLogPublisherResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || isNil(o.Meta) {
    return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ThirdPartyHttpOperationLogPublisherResponse) HasMeta() bool {
	if o != nil && !isNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *ThirdPartyHttpOperationLogPublisherResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

func (o ThirdPartyHttpOperationLogPublisherResponse) MarshalJSON() ([]byte, error) {
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
	if !isNil(o.LoggingErrorBehavior) {
		toSerialize["loggingErrorBehavior"] = o.LoggingErrorBehavior
	}
	if !isNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return json.Marshal(toSerialize)
}

type NullableThirdPartyHttpOperationLogPublisherResponse struct {
	value *ThirdPartyHttpOperationLogPublisherResponse
	isSet bool
}

func (v NullableThirdPartyHttpOperationLogPublisherResponse) Get() *ThirdPartyHttpOperationLogPublisherResponse {
	return v.value
}

func (v *NullableThirdPartyHttpOperationLogPublisherResponse) Set(val *ThirdPartyHttpOperationLogPublisherResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableThirdPartyHttpOperationLogPublisherResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableThirdPartyHttpOperationLogPublisherResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThirdPartyHttpOperationLogPublisherResponse(val *ThirdPartyHttpOperationLogPublisherResponse) *NullableThirdPartyHttpOperationLogPublisherResponse {
	return &NullableThirdPartyHttpOperationLogPublisherResponse{value: val, isSet: true}
}

func (v NullableThirdPartyHttpOperationLogPublisherResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThirdPartyHttpOperationLogPublisherResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


