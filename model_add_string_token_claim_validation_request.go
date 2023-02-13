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

// AddStringTokenClaimValidationRequest struct for AddStringTokenClaimValidationRequest
type AddStringTokenClaimValidationRequest struct {
	// Name of the new Token Claim Validation
	ValidationName string                                    `json:"validationName"`
	Schemas        []EnumstringTokenClaimValidationSchemaUrn `json:"schemas"`
	// The set of values that the claim may have to be considered valid.
	AnyRequiredValue []string `json:"anyRequiredValue"`
	// A description for this Token Claim Validation
	Description *string `json:"description,omitempty"`
	// The name of the claim to be validated.
	ClaimName string `json:"claimName"`
}

// NewAddStringTokenClaimValidationRequest instantiates a new AddStringTokenClaimValidationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddStringTokenClaimValidationRequest(validationName string, schemas []EnumstringTokenClaimValidationSchemaUrn, anyRequiredValue []string, claimName string) *AddStringTokenClaimValidationRequest {
	this := AddStringTokenClaimValidationRequest{}
	this.ValidationName = validationName
	this.Schemas = schemas
	this.AnyRequiredValue = anyRequiredValue
	this.ClaimName = claimName
	return &this
}

// NewAddStringTokenClaimValidationRequestWithDefaults instantiates a new AddStringTokenClaimValidationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddStringTokenClaimValidationRequestWithDefaults() *AddStringTokenClaimValidationRequest {
	this := AddStringTokenClaimValidationRequest{}
	return &this
}

// GetValidationName returns the ValidationName field value
func (o *AddStringTokenClaimValidationRequest) GetValidationName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ValidationName
}

// GetValidationNameOk returns a tuple with the ValidationName field value
// and a boolean to check if the value has been set.
func (o *AddStringTokenClaimValidationRequest) GetValidationNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ValidationName, true
}

// SetValidationName sets field value
func (o *AddStringTokenClaimValidationRequest) SetValidationName(v string) {
	o.ValidationName = v
}

// GetSchemas returns the Schemas field value
func (o *AddStringTokenClaimValidationRequest) GetSchemas() []EnumstringTokenClaimValidationSchemaUrn {
	if o == nil {
		var ret []EnumstringTokenClaimValidationSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddStringTokenClaimValidationRequest) GetSchemasOk() ([]EnumstringTokenClaimValidationSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddStringTokenClaimValidationRequest) SetSchemas(v []EnumstringTokenClaimValidationSchemaUrn) {
	o.Schemas = v
}

// GetAnyRequiredValue returns the AnyRequiredValue field value
func (o *AddStringTokenClaimValidationRequest) GetAnyRequiredValue() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.AnyRequiredValue
}

// GetAnyRequiredValueOk returns a tuple with the AnyRequiredValue field value
// and a boolean to check if the value has been set.
func (o *AddStringTokenClaimValidationRequest) GetAnyRequiredValueOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AnyRequiredValue, true
}

// SetAnyRequiredValue sets field value
func (o *AddStringTokenClaimValidationRequest) SetAnyRequiredValue(v []string) {
	o.AnyRequiredValue = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddStringTokenClaimValidationRequest) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddStringTokenClaimValidationRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddStringTokenClaimValidationRequest) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddStringTokenClaimValidationRequest) SetDescription(v string) {
	o.Description = &v
}

// GetClaimName returns the ClaimName field value
func (o *AddStringTokenClaimValidationRequest) GetClaimName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClaimName
}

// GetClaimNameOk returns a tuple with the ClaimName field value
// and a boolean to check if the value has been set.
func (o *AddStringTokenClaimValidationRequest) GetClaimNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClaimName, true
}

// SetClaimName sets field value
func (o *AddStringTokenClaimValidationRequest) SetClaimName(v string) {
	o.ClaimName = v
}

func (o AddStringTokenClaimValidationRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["validationName"] = o.ValidationName
	}
	if true {
		toSerialize["schemas"] = o.Schemas
	}
	if true {
		toSerialize["anyRequiredValue"] = o.AnyRequiredValue
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["claimName"] = o.ClaimName
	}
	return json.Marshal(toSerialize)
}

type NullableAddStringTokenClaimValidationRequest struct {
	value *AddStringTokenClaimValidationRequest
	isSet bool
}

func (v NullableAddStringTokenClaimValidationRequest) Get() *AddStringTokenClaimValidationRequest {
	return v.value
}

func (v *NullableAddStringTokenClaimValidationRequest) Set(val *AddStringTokenClaimValidationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddStringTokenClaimValidationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddStringTokenClaimValidationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddStringTokenClaimValidationRequest(val *AddStringTokenClaimValidationRequest) *NullableAddStringTokenClaimValidationRequest {
	return &NullableAddStringTokenClaimValidationRequest{value: val, isSet: true}
}

func (v NullableAddStringTokenClaimValidationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddStringTokenClaimValidationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
