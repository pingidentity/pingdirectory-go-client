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

// AddThirdPartyUncachedAttributeCriteriaRequest struct for AddThirdPartyUncachedAttributeCriteriaRequest
type AddThirdPartyUncachedAttributeCriteriaRequest struct {
	// Name of the new Uncached Attribute Criteria
	CriteriaName string `json:"criteriaName"`
	Schemas []EnumthirdPartyUncachedAttributeCriteriaSchemaUrn `json:"schemas"`
	// The fully-qualified name of the Java class providing the logic for the Third Party Uncached Attribute Criteria.
	ExtensionClass string `json:"extensionClass"`
	ExtensionArgument []string `json:"extensionArgument,omitempty"`
	// A description for this Uncached Attribute Criteria
	Description *string `json:"description,omitempty"`
	// Indicates whether this Uncached Attribute Criteria is enabled for use in the server.
	Enabled bool `json:"enabled"`
}

// NewAddThirdPartyUncachedAttributeCriteriaRequest instantiates a new AddThirdPartyUncachedAttributeCriteriaRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddThirdPartyUncachedAttributeCriteriaRequest(criteriaName string, schemas []EnumthirdPartyUncachedAttributeCriteriaSchemaUrn, extensionClass string, enabled bool) *AddThirdPartyUncachedAttributeCriteriaRequest {
	this := AddThirdPartyUncachedAttributeCriteriaRequest{}
	this.CriteriaName = criteriaName
	this.Schemas = schemas
	this.ExtensionClass = extensionClass
	this.Enabled = enabled
	return &this
}

// NewAddThirdPartyUncachedAttributeCriteriaRequestWithDefaults instantiates a new AddThirdPartyUncachedAttributeCriteriaRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddThirdPartyUncachedAttributeCriteriaRequestWithDefaults() *AddThirdPartyUncachedAttributeCriteriaRequest {
	this := AddThirdPartyUncachedAttributeCriteriaRequest{}
	return &this
}

// GetCriteriaName returns the CriteriaName field value
func (o *AddThirdPartyUncachedAttributeCriteriaRequest) GetCriteriaName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CriteriaName
}

// GetCriteriaNameOk returns a tuple with the CriteriaName field value
// and a boolean to check if the value has been set.
func (o *AddThirdPartyUncachedAttributeCriteriaRequest) GetCriteriaNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.CriteriaName, true
}

// SetCriteriaName sets field value
func (o *AddThirdPartyUncachedAttributeCriteriaRequest) SetCriteriaName(v string) {
	o.CriteriaName = v
}

// GetSchemas returns the Schemas field value
func (o *AddThirdPartyUncachedAttributeCriteriaRequest) GetSchemas() []EnumthirdPartyUncachedAttributeCriteriaSchemaUrn {
	if o == nil {
		var ret []EnumthirdPartyUncachedAttributeCriteriaSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddThirdPartyUncachedAttributeCriteriaRequest) GetSchemasOk() ([]EnumthirdPartyUncachedAttributeCriteriaSchemaUrn, bool) {
	if o == nil {
    return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddThirdPartyUncachedAttributeCriteriaRequest) SetSchemas(v []EnumthirdPartyUncachedAttributeCriteriaSchemaUrn) {
	o.Schemas = v
}

// GetExtensionClass returns the ExtensionClass field value
func (o *AddThirdPartyUncachedAttributeCriteriaRequest) GetExtensionClass() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExtensionClass
}

// GetExtensionClassOk returns a tuple with the ExtensionClass field value
// and a boolean to check if the value has been set.
func (o *AddThirdPartyUncachedAttributeCriteriaRequest) GetExtensionClassOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ExtensionClass, true
}

// SetExtensionClass sets field value
func (o *AddThirdPartyUncachedAttributeCriteriaRequest) SetExtensionClass(v string) {
	o.ExtensionClass = v
}

// GetExtensionArgument returns the ExtensionArgument field value if set, zero value otherwise.
func (o *AddThirdPartyUncachedAttributeCriteriaRequest) GetExtensionArgument() []string {
	if o == nil || isNil(o.ExtensionArgument) {
		var ret []string
		return ret
	}
	return o.ExtensionArgument
}

// GetExtensionArgumentOk returns a tuple with the ExtensionArgument field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddThirdPartyUncachedAttributeCriteriaRequest) GetExtensionArgumentOk() ([]string, bool) {
	if o == nil || isNil(o.ExtensionArgument) {
    return nil, false
	}
	return o.ExtensionArgument, true
}

// HasExtensionArgument returns a boolean if a field has been set.
func (o *AddThirdPartyUncachedAttributeCriteriaRequest) HasExtensionArgument() bool {
	if o != nil && !isNil(o.ExtensionArgument) {
		return true
	}

	return false
}

// SetExtensionArgument gets a reference to the given []string and assigns it to the ExtensionArgument field.
func (o *AddThirdPartyUncachedAttributeCriteriaRequest) SetExtensionArgument(v []string) {
	o.ExtensionArgument = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddThirdPartyUncachedAttributeCriteriaRequest) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddThirdPartyUncachedAttributeCriteriaRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddThirdPartyUncachedAttributeCriteriaRequest) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddThirdPartyUncachedAttributeCriteriaRequest) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *AddThirdPartyUncachedAttributeCriteriaRequest) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AddThirdPartyUncachedAttributeCriteriaRequest) GetEnabledOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AddThirdPartyUncachedAttributeCriteriaRequest) SetEnabled(v bool) {
	o.Enabled = v
}

func (o AddThirdPartyUncachedAttributeCriteriaRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["criteriaName"] = o.CriteriaName
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

type NullableAddThirdPartyUncachedAttributeCriteriaRequest struct {
	value *AddThirdPartyUncachedAttributeCriteriaRequest
	isSet bool
}

func (v NullableAddThirdPartyUncachedAttributeCriteriaRequest) Get() *AddThirdPartyUncachedAttributeCriteriaRequest {
	return v.value
}

func (v *NullableAddThirdPartyUncachedAttributeCriteriaRequest) Set(val *AddThirdPartyUncachedAttributeCriteriaRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddThirdPartyUncachedAttributeCriteriaRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddThirdPartyUncachedAttributeCriteriaRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddThirdPartyUncachedAttributeCriteriaRequest(val *AddThirdPartyUncachedAttributeCriteriaRequest) *NullableAddThirdPartyUncachedAttributeCriteriaRequest {
	return &NullableAddThirdPartyUncachedAttributeCriteriaRequest{value: val, isSet: true}
}

func (v NullableAddThirdPartyUncachedAttributeCriteriaRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddThirdPartyUncachedAttributeCriteriaRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


