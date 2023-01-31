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

// AddValidateTotpPasswordExtendedOperationHandlerRequest struct for AddValidateTotpPasswordExtendedOperationHandlerRequest
type AddValidateTotpPasswordExtendedOperationHandlerRequest struct {
	// Name of the new Extended Operation Handler
	HandlerName string                                                      `json:"handlerName"`
	Schemas     []EnumvalidateTotpPasswordExtendedOperationHandlerSchemaUrn `json:"schemas"`
	// The name or OID of the attribute that will be used to hold the shared secret key used during TOTP processing.
	SharedSecretAttributeType *string `json:"sharedSecretAttributeType,omitempty"`
	// The duration of the time interval used for TOTP processing.
	TimeIntervalDuration *string `json:"timeIntervalDuration,omitempty"`
	// The number of adjacent time intervals (both before and after the current time) that should be checked when performing authentication.
	AdjacentIntervalsToCheck *int32 `json:"adjacentIntervalsToCheck,omitempty"`
	// Indicates whether to prevent clients from re-using TOTP passwords.
	PreventTOTPReuse *bool `json:"preventTOTPReuse,omitempty"`
	// A description for this Extended Operation Handler
	Description *string `json:"description,omitempty"`
	// Indicates whether the Extended Operation Handler is enabled (that is, whether the types of extended operations are allowed in the server).
	Enabled bool `json:"enabled"`
}

// NewAddValidateTotpPasswordExtendedOperationHandlerRequest instantiates a new AddValidateTotpPasswordExtendedOperationHandlerRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddValidateTotpPasswordExtendedOperationHandlerRequest(handlerName string, schemas []EnumvalidateTotpPasswordExtendedOperationHandlerSchemaUrn, enabled bool) *AddValidateTotpPasswordExtendedOperationHandlerRequest {
	this := AddValidateTotpPasswordExtendedOperationHandlerRequest{}
	this.HandlerName = handlerName
	this.Schemas = schemas
	this.Enabled = enabled
	return &this
}

// NewAddValidateTotpPasswordExtendedOperationHandlerRequestWithDefaults instantiates a new AddValidateTotpPasswordExtendedOperationHandlerRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddValidateTotpPasswordExtendedOperationHandlerRequestWithDefaults() *AddValidateTotpPasswordExtendedOperationHandlerRequest {
	this := AddValidateTotpPasswordExtendedOperationHandlerRequest{}
	return &this
}

// GetHandlerName returns the HandlerName field value
func (o *AddValidateTotpPasswordExtendedOperationHandlerRequest) GetHandlerName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HandlerName
}

// GetHandlerNameOk returns a tuple with the HandlerName field value
// and a boolean to check if the value has been set.
func (o *AddValidateTotpPasswordExtendedOperationHandlerRequest) GetHandlerNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HandlerName, true
}

// SetHandlerName sets field value
func (o *AddValidateTotpPasswordExtendedOperationHandlerRequest) SetHandlerName(v string) {
	o.HandlerName = v
}

// GetSchemas returns the Schemas field value
func (o *AddValidateTotpPasswordExtendedOperationHandlerRequest) GetSchemas() []EnumvalidateTotpPasswordExtendedOperationHandlerSchemaUrn {
	if o == nil {
		var ret []EnumvalidateTotpPasswordExtendedOperationHandlerSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddValidateTotpPasswordExtendedOperationHandlerRequest) GetSchemasOk() ([]EnumvalidateTotpPasswordExtendedOperationHandlerSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddValidateTotpPasswordExtendedOperationHandlerRequest) SetSchemas(v []EnumvalidateTotpPasswordExtendedOperationHandlerSchemaUrn) {
	o.Schemas = v
}

// GetSharedSecretAttributeType returns the SharedSecretAttributeType field value if set, zero value otherwise.
func (o *AddValidateTotpPasswordExtendedOperationHandlerRequest) GetSharedSecretAttributeType() string {
	if o == nil || isNil(o.SharedSecretAttributeType) {
		var ret string
		return ret
	}
	return *o.SharedSecretAttributeType
}

// GetSharedSecretAttributeTypeOk returns a tuple with the SharedSecretAttributeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddValidateTotpPasswordExtendedOperationHandlerRequest) GetSharedSecretAttributeTypeOk() (*string, bool) {
	if o == nil || isNil(o.SharedSecretAttributeType) {
		return nil, false
	}
	return o.SharedSecretAttributeType, true
}

// HasSharedSecretAttributeType returns a boolean if a field has been set.
func (o *AddValidateTotpPasswordExtendedOperationHandlerRequest) HasSharedSecretAttributeType() bool {
	if o != nil && !isNil(o.SharedSecretAttributeType) {
		return true
	}

	return false
}

// SetSharedSecretAttributeType gets a reference to the given string and assigns it to the SharedSecretAttributeType field.
func (o *AddValidateTotpPasswordExtendedOperationHandlerRequest) SetSharedSecretAttributeType(v string) {
	o.SharedSecretAttributeType = &v
}

// GetTimeIntervalDuration returns the TimeIntervalDuration field value if set, zero value otherwise.
func (o *AddValidateTotpPasswordExtendedOperationHandlerRequest) GetTimeIntervalDuration() string {
	if o == nil || isNil(o.TimeIntervalDuration) {
		var ret string
		return ret
	}
	return *o.TimeIntervalDuration
}

// GetTimeIntervalDurationOk returns a tuple with the TimeIntervalDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddValidateTotpPasswordExtendedOperationHandlerRequest) GetTimeIntervalDurationOk() (*string, bool) {
	if o == nil || isNil(o.TimeIntervalDuration) {
		return nil, false
	}
	return o.TimeIntervalDuration, true
}

// HasTimeIntervalDuration returns a boolean if a field has been set.
func (o *AddValidateTotpPasswordExtendedOperationHandlerRequest) HasTimeIntervalDuration() bool {
	if o != nil && !isNil(o.TimeIntervalDuration) {
		return true
	}

	return false
}

// SetTimeIntervalDuration gets a reference to the given string and assigns it to the TimeIntervalDuration field.
func (o *AddValidateTotpPasswordExtendedOperationHandlerRequest) SetTimeIntervalDuration(v string) {
	o.TimeIntervalDuration = &v
}

// GetAdjacentIntervalsToCheck returns the AdjacentIntervalsToCheck field value if set, zero value otherwise.
func (o *AddValidateTotpPasswordExtendedOperationHandlerRequest) GetAdjacentIntervalsToCheck() int32 {
	if o == nil || isNil(o.AdjacentIntervalsToCheck) {
		var ret int32
		return ret
	}
	return *o.AdjacentIntervalsToCheck
}

// GetAdjacentIntervalsToCheckOk returns a tuple with the AdjacentIntervalsToCheck field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddValidateTotpPasswordExtendedOperationHandlerRequest) GetAdjacentIntervalsToCheckOk() (*int32, bool) {
	if o == nil || isNil(o.AdjacentIntervalsToCheck) {
		return nil, false
	}
	return o.AdjacentIntervalsToCheck, true
}

// HasAdjacentIntervalsToCheck returns a boolean if a field has been set.
func (o *AddValidateTotpPasswordExtendedOperationHandlerRequest) HasAdjacentIntervalsToCheck() bool {
	if o != nil && !isNil(o.AdjacentIntervalsToCheck) {
		return true
	}

	return false
}

// SetAdjacentIntervalsToCheck gets a reference to the given int32 and assigns it to the AdjacentIntervalsToCheck field.
func (o *AddValidateTotpPasswordExtendedOperationHandlerRequest) SetAdjacentIntervalsToCheck(v int32) {
	o.AdjacentIntervalsToCheck = &v
}

// GetPreventTOTPReuse returns the PreventTOTPReuse field value if set, zero value otherwise.
func (o *AddValidateTotpPasswordExtendedOperationHandlerRequest) GetPreventTOTPReuse() bool {
	if o == nil || isNil(o.PreventTOTPReuse) {
		var ret bool
		return ret
	}
	return *o.PreventTOTPReuse
}

// GetPreventTOTPReuseOk returns a tuple with the PreventTOTPReuse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddValidateTotpPasswordExtendedOperationHandlerRequest) GetPreventTOTPReuseOk() (*bool, bool) {
	if o == nil || isNil(o.PreventTOTPReuse) {
		return nil, false
	}
	return o.PreventTOTPReuse, true
}

// HasPreventTOTPReuse returns a boolean if a field has been set.
func (o *AddValidateTotpPasswordExtendedOperationHandlerRequest) HasPreventTOTPReuse() bool {
	if o != nil && !isNil(o.PreventTOTPReuse) {
		return true
	}

	return false
}

// SetPreventTOTPReuse gets a reference to the given bool and assigns it to the PreventTOTPReuse field.
func (o *AddValidateTotpPasswordExtendedOperationHandlerRequest) SetPreventTOTPReuse(v bool) {
	o.PreventTOTPReuse = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddValidateTotpPasswordExtendedOperationHandlerRequest) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddValidateTotpPasswordExtendedOperationHandlerRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddValidateTotpPasswordExtendedOperationHandlerRequest) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddValidateTotpPasswordExtendedOperationHandlerRequest) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *AddValidateTotpPasswordExtendedOperationHandlerRequest) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AddValidateTotpPasswordExtendedOperationHandlerRequest) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AddValidateTotpPasswordExtendedOperationHandlerRequest) SetEnabled(v bool) {
	o.Enabled = v
}

func (o AddValidateTotpPasswordExtendedOperationHandlerRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["handlerName"] = o.HandlerName
	}
	if true {
		toSerialize["schemas"] = o.Schemas
	}
	if !isNil(o.SharedSecretAttributeType) {
		toSerialize["sharedSecretAttributeType"] = o.SharedSecretAttributeType
	}
	if !isNil(o.TimeIntervalDuration) {
		toSerialize["timeIntervalDuration"] = o.TimeIntervalDuration
	}
	if !isNil(o.AdjacentIntervalsToCheck) {
		toSerialize["adjacentIntervalsToCheck"] = o.AdjacentIntervalsToCheck
	}
	if !isNil(o.PreventTOTPReuse) {
		toSerialize["preventTOTPReuse"] = o.PreventTOTPReuse
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableAddValidateTotpPasswordExtendedOperationHandlerRequest struct {
	value *AddValidateTotpPasswordExtendedOperationHandlerRequest
	isSet bool
}

func (v NullableAddValidateTotpPasswordExtendedOperationHandlerRequest) Get() *AddValidateTotpPasswordExtendedOperationHandlerRequest {
	return v.value
}

func (v *NullableAddValidateTotpPasswordExtendedOperationHandlerRequest) Set(val *AddValidateTotpPasswordExtendedOperationHandlerRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddValidateTotpPasswordExtendedOperationHandlerRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddValidateTotpPasswordExtendedOperationHandlerRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddValidateTotpPasswordExtendedOperationHandlerRequest(val *AddValidateTotpPasswordExtendedOperationHandlerRequest) *NullableAddValidateTotpPasswordExtendedOperationHandlerRequest {
	return &NullableAddValidateTotpPasswordExtendedOperationHandlerRequest{value: val, isSet: true}
}

func (v NullableAddValidateTotpPasswordExtendedOperationHandlerRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddValidateTotpPasswordExtendedOperationHandlerRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
