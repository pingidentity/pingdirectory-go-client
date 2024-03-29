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

// checks if the AddEncryptionSettingsDatabaseAccessibilityMonitorProviderRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddEncryptionSettingsDatabaseAccessibilityMonitorProviderRequest{}

// AddEncryptionSettingsDatabaseAccessibilityMonitorProviderRequest struct for AddEncryptionSettingsDatabaseAccessibilityMonitorProviderRequest
type AddEncryptionSettingsDatabaseAccessibilityMonitorProviderRequest struct {
	Schemas []EnumencryptionSettingsDatabaseAccessibilityMonitorProviderSchemaUrn `json:"schemas"`
	// The frequency with which this monitor provider should confirm the ability to access the server's encryption settings database.
	CheckFrequency *string `json:"checkFrequency,omitempty"`
	// The minimum length of time that an outage should persist before it is considered a prolonged outage. If an outage lasts at least as long as this duration, then the server will take the action indicated by the prolonged-outage-behavior property.
	ProlongedOutageDuration *string                                         `json:"prolongedOutageDuration,omitempty"`
	ProlongedOutageBehavior *EnummonitorProviderProlongedOutageBehaviorProp `json:"prolongedOutageBehavior,omitempty"`
	// A description for this Monitor Provider
	Description *string `json:"description,omitempty"`
	// Indicates whether the Monitor Provider is enabled for use.
	Enabled bool `json:"enabled"`
	// Name of the new Monitor Provider
	ProviderName string `json:"providerName"`
}

// NewAddEncryptionSettingsDatabaseAccessibilityMonitorProviderRequest instantiates a new AddEncryptionSettingsDatabaseAccessibilityMonitorProviderRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddEncryptionSettingsDatabaseAccessibilityMonitorProviderRequest(schemas []EnumencryptionSettingsDatabaseAccessibilityMonitorProviderSchemaUrn, enabled bool, providerName string) *AddEncryptionSettingsDatabaseAccessibilityMonitorProviderRequest {
	this := AddEncryptionSettingsDatabaseAccessibilityMonitorProviderRequest{}
	this.Schemas = schemas
	this.Enabled = enabled
	this.ProviderName = providerName
	return &this
}

// NewAddEncryptionSettingsDatabaseAccessibilityMonitorProviderRequestWithDefaults instantiates a new AddEncryptionSettingsDatabaseAccessibilityMonitorProviderRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddEncryptionSettingsDatabaseAccessibilityMonitorProviderRequestWithDefaults() *AddEncryptionSettingsDatabaseAccessibilityMonitorProviderRequest {
	this := AddEncryptionSettingsDatabaseAccessibilityMonitorProviderRequest{}
	return &this
}

// GetSchemas returns the Schemas field value
func (o *AddEncryptionSettingsDatabaseAccessibilityMonitorProviderRequest) GetSchemas() []EnumencryptionSettingsDatabaseAccessibilityMonitorProviderSchemaUrn {
	if o == nil {
		var ret []EnumencryptionSettingsDatabaseAccessibilityMonitorProviderSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddEncryptionSettingsDatabaseAccessibilityMonitorProviderRequest) GetSchemasOk() ([]EnumencryptionSettingsDatabaseAccessibilityMonitorProviderSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddEncryptionSettingsDatabaseAccessibilityMonitorProviderRequest) SetSchemas(v []EnumencryptionSettingsDatabaseAccessibilityMonitorProviderSchemaUrn) {
	o.Schemas = v
}

// GetCheckFrequency returns the CheckFrequency field value if set, zero value otherwise.
func (o *AddEncryptionSettingsDatabaseAccessibilityMonitorProviderRequest) GetCheckFrequency() string {
	if o == nil || IsNil(o.CheckFrequency) {
		var ret string
		return ret
	}
	return *o.CheckFrequency
}

// GetCheckFrequencyOk returns a tuple with the CheckFrequency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddEncryptionSettingsDatabaseAccessibilityMonitorProviderRequest) GetCheckFrequencyOk() (*string, bool) {
	if o == nil || IsNil(o.CheckFrequency) {
		return nil, false
	}
	return o.CheckFrequency, true
}

// HasCheckFrequency returns a boolean if a field has been set.
func (o *AddEncryptionSettingsDatabaseAccessibilityMonitorProviderRequest) HasCheckFrequency() bool {
	if o != nil && !IsNil(o.CheckFrequency) {
		return true
	}

	return false
}

// SetCheckFrequency gets a reference to the given string and assigns it to the CheckFrequency field.
func (o *AddEncryptionSettingsDatabaseAccessibilityMonitorProviderRequest) SetCheckFrequency(v string) {
	o.CheckFrequency = &v
}

// GetProlongedOutageDuration returns the ProlongedOutageDuration field value if set, zero value otherwise.
func (o *AddEncryptionSettingsDatabaseAccessibilityMonitorProviderRequest) GetProlongedOutageDuration() string {
	if o == nil || IsNil(o.ProlongedOutageDuration) {
		var ret string
		return ret
	}
	return *o.ProlongedOutageDuration
}

// GetProlongedOutageDurationOk returns a tuple with the ProlongedOutageDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddEncryptionSettingsDatabaseAccessibilityMonitorProviderRequest) GetProlongedOutageDurationOk() (*string, bool) {
	if o == nil || IsNil(o.ProlongedOutageDuration) {
		return nil, false
	}
	return o.ProlongedOutageDuration, true
}

// HasProlongedOutageDuration returns a boolean if a field has been set.
func (o *AddEncryptionSettingsDatabaseAccessibilityMonitorProviderRequest) HasProlongedOutageDuration() bool {
	if o != nil && !IsNil(o.ProlongedOutageDuration) {
		return true
	}

	return false
}

// SetProlongedOutageDuration gets a reference to the given string and assigns it to the ProlongedOutageDuration field.
func (o *AddEncryptionSettingsDatabaseAccessibilityMonitorProviderRequest) SetProlongedOutageDuration(v string) {
	o.ProlongedOutageDuration = &v
}

// GetProlongedOutageBehavior returns the ProlongedOutageBehavior field value if set, zero value otherwise.
func (o *AddEncryptionSettingsDatabaseAccessibilityMonitorProviderRequest) GetProlongedOutageBehavior() EnummonitorProviderProlongedOutageBehaviorProp {
	if o == nil || IsNil(o.ProlongedOutageBehavior) {
		var ret EnummonitorProviderProlongedOutageBehaviorProp
		return ret
	}
	return *o.ProlongedOutageBehavior
}

// GetProlongedOutageBehaviorOk returns a tuple with the ProlongedOutageBehavior field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddEncryptionSettingsDatabaseAccessibilityMonitorProviderRequest) GetProlongedOutageBehaviorOk() (*EnummonitorProviderProlongedOutageBehaviorProp, bool) {
	if o == nil || IsNil(o.ProlongedOutageBehavior) {
		return nil, false
	}
	return o.ProlongedOutageBehavior, true
}

// HasProlongedOutageBehavior returns a boolean if a field has been set.
func (o *AddEncryptionSettingsDatabaseAccessibilityMonitorProviderRequest) HasProlongedOutageBehavior() bool {
	if o != nil && !IsNil(o.ProlongedOutageBehavior) {
		return true
	}

	return false
}

// SetProlongedOutageBehavior gets a reference to the given EnummonitorProviderProlongedOutageBehaviorProp and assigns it to the ProlongedOutageBehavior field.
func (o *AddEncryptionSettingsDatabaseAccessibilityMonitorProviderRequest) SetProlongedOutageBehavior(v EnummonitorProviderProlongedOutageBehaviorProp) {
	o.ProlongedOutageBehavior = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddEncryptionSettingsDatabaseAccessibilityMonitorProviderRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddEncryptionSettingsDatabaseAccessibilityMonitorProviderRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddEncryptionSettingsDatabaseAccessibilityMonitorProviderRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddEncryptionSettingsDatabaseAccessibilityMonitorProviderRequest) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *AddEncryptionSettingsDatabaseAccessibilityMonitorProviderRequest) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AddEncryptionSettingsDatabaseAccessibilityMonitorProviderRequest) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AddEncryptionSettingsDatabaseAccessibilityMonitorProviderRequest) SetEnabled(v bool) {
	o.Enabled = v
}

// GetProviderName returns the ProviderName field value
func (o *AddEncryptionSettingsDatabaseAccessibilityMonitorProviderRequest) GetProviderName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProviderName
}

// GetProviderNameOk returns a tuple with the ProviderName field value
// and a boolean to check if the value has been set.
func (o *AddEncryptionSettingsDatabaseAccessibilityMonitorProviderRequest) GetProviderNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProviderName, true
}

// SetProviderName sets field value
func (o *AddEncryptionSettingsDatabaseAccessibilityMonitorProviderRequest) SetProviderName(v string) {
	o.ProviderName = v
}

func (o AddEncryptionSettingsDatabaseAccessibilityMonitorProviderRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddEncryptionSettingsDatabaseAccessibilityMonitorProviderRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["schemas"] = o.Schemas
	if !IsNil(o.CheckFrequency) {
		toSerialize["checkFrequency"] = o.CheckFrequency
	}
	if !IsNil(o.ProlongedOutageDuration) {
		toSerialize["prolongedOutageDuration"] = o.ProlongedOutageDuration
	}
	if !IsNil(o.ProlongedOutageBehavior) {
		toSerialize["prolongedOutageBehavior"] = o.ProlongedOutageBehavior
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["enabled"] = o.Enabled
	toSerialize["providerName"] = o.ProviderName
	return toSerialize, nil
}

type NullableAddEncryptionSettingsDatabaseAccessibilityMonitorProviderRequest struct {
	value *AddEncryptionSettingsDatabaseAccessibilityMonitorProviderRequest
	isSet bool
}

func (v NullableAddEncryptionSettingsDatabaseAccessibilityMonitorProviderRequest) Get() *AddEncryptionSettingsDatabaseAccessibilityMonitorProviderRequest {
	return v.value
}

func (v *NullableAddEncryptionSettingsDatabaseAccessibilityMonitorProviderRequest) Set(val *AddEncryptionSettingsDatabaseAccessibilityMonitorProviderRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddEncryptionSettingsDatabaseAccessibilityMonitorProviderRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddEncryptionSettingsDatabaseAccessibilityMonitorProviderRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddEncryptionSettingsDatabaseAccessibilityMonitorProviderRequest(val *AddEncryptionSettingsDatabaseAccessibilityMonitorProviderRequest) *NullableAddEncryptionSettingsDatabaseAccessibilityMonitorProviderRequest {
	return &NullableAddEncryptionSettingsDatabaseAccessibilityMonitorProviderRequest{value: val, isSet: true}
}

func (v NullableAddEncryptionSettingsDatabaseAccessibilityMonitorProviderRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddEncryptionSettingsDatabaseAccessibilityMonitorProviderRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
