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

// checks if the AddBooleanTokenClaimValidationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddBooleanTokenClaimValidationRequest{}

// AddBooleanTokenClaimValidationRequest struct for AddBooleanTokenClaimValidationRequest
type AddBooleanTokenClaimValidationRequest struct {
	Schemas       []EnumbooleanTokenClaimValidationSchemaUrn `json:"schemas"`
	RequiredValue EnumtokenClaimValidationRequiredValueProp  `json:"requiredValue"`
	// A description for this Token Claim Validation
	Description *string `json:"description,omitempty"`
	// The name of the claim to be validated.
	ClaimName string `json:"claimName"`
	// Name of the new Token Claim Validation
	ValidationName string `json:"validationName"`
}

// NewAddBooleanTokenClaimValidationRequest instantiates a new AddBooleanTokenClaimValidationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddBooleanTokenClaimValidationRequest(schemas []EnumbooleanTokenClaimValidationSchemaUrn, requiredValue EnumtokenClaimValidationRequiredValueProp, claimName string, validationName string) *AddBooleanTokenClaimValidationRequest {
	this := AddBooleanTokenClaimValidationRequest{}
	this.Schemas = schemas
	this.RequiredValue = requiredValue
	this.ClaimName = claimName
	this.ValidationName = validationName
	return &this
}

// NewAddBooleanTokenClaimValidationRequestWithDefaults instantiates a new AddBooleanTokenClaimValidationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddBooleanTokenClaimValidationRequestWithDefaults() *AddBooleanTokenClaimValidationRequest {
	this := AddBooleanTokenClaimValidationRequest{}
	return &this
}

// GetSchemas returns the Schemas field value
func (o *AddBooleanTokenClaimValidationRequest) GetSchemas() []EnumbooleanTokenClaimValidationSchemaUrn {
	if o == nil {
		var ret []EnumbooleanTokenClaimValidationSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddBooleanTokenClaimValidationRequest) GetSchemasOk() ([]EnumbooleanTokenClaimValidationSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddBooleanTokenClaimValidationRequest) SetSchemas(v []EnumbooleanTokenClaimValidationSchemaUrn) {
	o.Schemas = v
}

// GetRequiredValue returns the RequiredValue field value
func (o *AddBooleanTokenClaimValidationRequest) GetRequiredValue() EnumtokenClaimValidationRequiredValueProp {
	if o == nil {
		var ret EnumtokenClaimValidationRequiredValueProp
		return ret
	}

	return o.RequiredValue
}

// GetRequiredValueOk returns a tuple with the RequiredValue field value
// and a boolean to check if the value has been set.
func (o *AddBooleanTokenClaimValidationRequest) GetRequiredValueOk() (*EnumtokenClaimValidationRequiredValueProp, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequiredValue, true
}

// SetRequiredValue sets field value
func (o *AddBooleanTokenClaimValidationRequest) SetRequiredValue(v EnumtokenClaimValidationRequiredValueProp) {
	o.RequiredValue = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddBooleanTokenClaimValidationRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddBooleanTokenClaimValidationRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddBooleanTokenClaimValidationRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddBooleanTokenClaimValidationRequest) SetDescription(v string) {
	o.Description = &v
}

// GetClaimName returns the ClaimName field value
func (o *AddBooleanTokenClaimValidationRequest) GetClaimName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClaimName
}

// GetClaimNameOk returns a tuple with the ClaimName field value
// and a boolean to check if the value has been set.
func (o *AddBooleanTokenClaimValidationRequest) GetClaimNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClaimName, true
}

// SetClaimName sets field value
func (o *AddBooleanTokenClaimValidationRequest) SetClaimName(v string) {
	o.ClaimName = v
}

// GetValidationName returns the ValidationName field value
func (o *AddBooleanTokenClaimValidationRequest) GetValidationName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ValidationName
}

// GetValidationNameOk returns a tuple with the ValidationName field value
// and a boolean to check if the value has been set.
func (o *AddBooleanTokenClaimValidationRequest) GetValidationNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ValidationName, true
}

// SetValidationName sets field value
func (o *AddBooleanTokenClaimValidationRequest) SetValidationName(v string) {
	o.ValidationName = v
}

func (o AddBooleanTokenClaimValidationRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddBooleanTokenClaimValidationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["schemas"] = o.Schemas
	toSerialize["requiredValue"] = o.RequiredValue
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["claimName"] = o.ClaimName
	toSerialize["validationName"] = o.ValidationName
	return toSerialize, nil
}

type NullableAddBooleanTokenClaimValidationRequest struct {
	value *AddBooleanTokenClaimValidationRequest
	isSet bool
}

func (v NullableAddBooleanTokenClaimValidationRequest) Get() *AddBooleanTokenClaimValidationRequest {
	return v.value
}

func (v *NullableAddBooleanTokenClaimValidationRequest) Set(val *AddBooleanTokenClaimValidationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddBooleanTokenClaimValidationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddBooleanTokenClaimValidationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddBooleanTokenClaimValidationRequest(val *AddBooleanTokenClaimValidationRequest) *NullableAddBooleanTokenClaimValidationRequest {
	return &NullableAddBooleanTokenClaimValidationRequest{value: val, isSet: true}
}

func (v NullableAddBooleanTokenClaimValidationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddBooleanTokenClaimValidationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
