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

// checks if the AddStringArrayTokenClaimValidationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddStringArrayTokenClaimValidationRequest{}

// AddStringArrayTokenClaimValidationRequest struct for AddStringArrayTokenClaimValidationRequest
type AddStringArrayTokenClaimValidationRequest struct {
	// Name of the new Token Claim Validation
	ValidationName string                                         `json:"validationName"`
	Schemas        []EnumstringArrayTokenClaimValidationSchemaUrn `json:"schemas"`
	// The set of all values that the claim must have to be considered valid.
	AllRequiredValue []string `json:"allRequiredValue,omitempty"`
	// The set of values that the claim may have to be considered valid.
	AnyRequiredValue []string `json:"anyRequiredValue,omitempty"`
	// A description for this Token Claim Validation
	Description *string `json:"description,omitempty"`
	// The name of the claim to be validated.
	ClaimName string `json:"claimName"`
}

// NewAddStringArrayTokenClaimValidationRequest instantiates a new AddStringArrayTokenClaimValidationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddStringArrayTokenClaimValidationRequest(validationName string, schemas []EnumstringArrayTokenClaimValidationSchemaUrn, claimName string) *AddStringArrayTokenClaimValidationRequest {
	this := AddStringArrayTokenClaimValidationRequest{}
	this.ValidationName = validationName
	this.Schemas = schemas
	this.ClaimName = claimName
	return &this
}

// NewAddStringArrayTokenClaimValidationRequestWithDefaults instantiates a new AddStringArrayTokenClaimValidationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddStringArrayTokenClaimValidationRequestWithDefaults() *AddStringArrayTokenClaimValidationRequest {
	this := AddStringArrayTokenClaimValidationRequest{}
	return &this
}

// GetValidationName returns the ValidationName field value
func (o *AddStringArrayTokenClaimValidationRequest) GetValidationName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ValidationName
}

// GetValidationNameOk returns a tuple with the ValidationName field value
// and a boolean to check if the value has been set.
func (o *AddStringArrayTokenClaimValidationRequest) GetValidationNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ValidationName, true
}

// SetValidationName sets field value
func (o *AddStringArrayTokenClaimValidationRequest) SetValidationName(v string) {
	o.ValidationName = v
}

// GetSchemas returns the Schemas field value
func (o *AddStringArrayTokenClaimValidationRequest) GetSchemas() []EnumstringArrayTokenClaimValidationSchemaUrn {
	if o == nil {
		var ret []EnumstringArrayTokenClaimValidationSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddStringArrayTokenClaimValidationRequest) GetSchemasOk() ([]EnumstringArrayTokenClaimValidationSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddStringArrayTokenClaimValidationRequest) SetSchemas(v []EnumstringArrayTokenClaimValidationSchemaUrn) {
	o.Schemas = v
}

// GetAllRequiredValue returns the AllRequiredValue field value if set, zero value otherwise.
func (o *AddStringArrayTokenClaimValidationRequest) GetAllRequiredValue() []string {
	if o == nil || IsNil(o.AllRequiredValue) {
		var ret []string
		return ret
	}
	return o.AllRequiredValue
}

// GetAllRequiredValueOk returns a tuple with the AllRequiredValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddStringArrayTokenClaimValidationRequest) GetAllRequiredValueOk() ([]string, bool) {
	if o == nil || IsNil(o.AllRequiredValue) {
		return nil, false
	}
	return o.AllRequiredValue, true
}

// HasAllRequiredValue returns a boolean if a field has been set.
func (o *AddStringArrayTokenClaimValidationRequest) HasAllRequiredValue() bool {
	if o != nil && !IsNil(o.AllRequiredValue) {
		return true
	}

	return false
}

// SetAllRequiredValue gets a reference to the given []string and assigns it to the AllRequiredValue field.
func (o *AddStringArrayTokenClaimValidationRequest) SetAllRequiredValue(v []string) {
	o.AllRequiredValue = v
}

// GetAnyRequiredValue returns the AnyRequiredValue field value if set, zero value otherwise.
func (o *AddStringArrayTokenClaimValidationRequest) GetAnyRequiredValue() []string {
	if o == nil || IsNil(o.AnyRequiredValue) {
		var ret []string
		return ret
	}
	return o.AnyRequiredValue
}

// GetAnyRequiredValueOk returns a tuple with the AnyRequiredValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddStringArrayTokenClaimValidationRequest) GetAnyRequiredValueOk() ([]string, bool) {
	if o == nil || IsNil(o.AnyRequiredValue) {
		return nil, false
	}
	return o.AnyRequiredValue, true
}

// HasAnyRequiredValue returns a boolean if a field has been set.
func (o *AddStringArrayTokenClaimValidationRequest) HasAnyRequiredValue() bool {
	if o != nil && !IsNil(o.AnyRequiredValue) {
		return true
	}

	return false
}

// SetAnyRequiredValue gets a reference to the given []string and assigns it to the AnyRequiredValue field.
func (o *AddStringArrayTokenClaimValidationRequest) SetAnyRequiredValue(v []string) {
	o.AnyRequiredValue = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddStringArrayTokenClaimValidationRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddStringArrayTokenClaimValidationRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddStringArrayTokenClaimValidationRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddStringArrayTokenClaimValidationRequest) SetDescription(v string) {
	o.Description = &v
}

// GetClaimName returns the ClaimName field value
func (o *AddStringArrayTokenClaimValidationRequest) GetClaimName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClaimName
}

// GetClaimNameOk returns a tuple with the ClaimName field value
// and a boolean to check if the value has been set.
func (o *AddStringArrayTokenClaimValidationRequest) GetClaimNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClaimName, true
}

// SetClaimName sets field value
func (o *AddStringArrayTokenClaimValidationRequest) SetClaimName(v string) {
	o.ClaimName = v
}

func (o AddStringArrayTokenClaimValidationRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddStringArrayTokenClaimValidationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["validationName"] = o.ValidationName
	toSerialize["schemas"] = o.Schemas
	if !IsNil(o.AllRequiredValue) {
		toSerialize["allRequiredValue"] = o.AllRequiredValue
	}
	if !IsNil(o.AnyRequiredValue) {
		toSerialize["anyRequiredValue"] = o.AnyRequiredValue
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["claimName"] = o.ClaimName
	return toSerialize, nil
}

type NullableAddStringArrayTokenClaimValidationRequest struct {
	value *AddStringArrayTokenClaimValidationRequest
	isSet bool
}

func (v NullableAddStringArrayTokenClaimValidationRequest) Get() *AddStringArrayTokenClaimValidationRequest {
	return v.value
}

func (v *NullableAddStringArrayTokenClaimValidationRequest) Set(val *AddStringArrayTokenClaimValidationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddStringArrayTokenClaimValidationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddStringArrayTokenClaimValidationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddStringArrayTokenClaimValidationRequest(val *AddStringArrayTokenClaimValidationRequest) *NullableAddStringArrayTokenClaimValidationRequest {
	return &NullableAddStringArrayTokenClaimValidationRequest{value: val, isSet: true}
}

func (v NullableAddStringArrayTokenClaimValidationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddStringArrayTokenClaimValidationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
