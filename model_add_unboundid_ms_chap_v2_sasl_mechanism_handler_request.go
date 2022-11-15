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

// AddUnboundidMsChapV2SaslMechanismHandlerRequest struct for AddUnboundidMsChapV2SaslMechanismHandlerRequest
type AddUnboundidMsChapV2SaslMechanismHandlerRequest struct {
	// Name of the new SASL Mechanism Handler
	HandlerName string `json:"handlerName"`
	Schemas []EnumunboundidMsChapV2SaslMechanismHandlerSchemaUrn `json:"schemas"`
	// The identity mapper that should be used to identify the entry associated with the username provided in the bind request.
	IdentityMapper string `json:"identityMapper"`
	// A description for this SASL Mechanism Handler
	Description *string `json:"description,omitempty"`
	// Indicates whether the SASL mechanism handler is enabled for use.
	Enabled bool `json:"enabled"`
}

// NewAddUnboundidMsChapV2SaslMechanismHandlerRequest instantiates a new AddUnboundidMsChapV2SaslMechanismHandlerRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddUnboundidMsChapV2SaslMechanismHandlerRequest(handlerName string, schemas []EnumunboundidMsChapV2SaslMechanismHandlerSchemaUrn, identityMapper string, enabled bool) *AddUnboundidMsChapV2SaslMechanismHandlerRequest {
	this := AddUnboundidMsChapV2SaslMechanismHandlerRequest{}
	this.HandlerName = handlerName
	this.Schemas = schemas
	this.IdentityMapper = identityMapper
	this.Enabled = enabled
	return &this
}

// NewAddUnboundidMsChapV2SaslMechanismHandlerRequestWithDefaults instantiates a new AddUnboundidMsChapV2SaslMechanismHandlerRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddUnboundidMsChapV2SaslMechanismHandlerRequestWithDefaults() *AddUnboundidMsChapV2SaslMechanismHandlerRequest {
	this := AddUnboundidMsChapV2SaslMechanismHandlerRequest{}
	return &this
}

// GetHandlerName returns the HandlerName field value
func (o *AddUnboundidMsChapV2SaslMechanismHandlerRequest) GetHandlerName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HandlerName
}

// GetHandlerNameOk returns a tuple with the HandlerName field value
// and a boolean to check if the value has been set.
func (o *AddUnboundidMsChapV2SaslMechanismHandlerRequest) GetHandlerNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.HandlerName, true
}

// SetHandlerName sets field value
func (o *AddUnboundidMsChapV2SaslMechanismHandlerRequest) SetHandlerName(v string) {
	o.HandlerName = v
}

// GetSchemas returns the Schemas field value
func (o *AddUnboundidMsChapV2SaslMechanismHandlerRequest) GetSchemas() []EnumunboundidMsChapV2SaslMechanismHandlerSchemaUrn {
	if o == nil {
		var ret []EnumunboundidMsChapV2SaslMechanismHandlerSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddUnboundidMsChapV2SaslMechanismHandlerRequest) GetSchemasOk() ([]EnumunboundidMsChapV2SaslMechanismHandlerSchemaUrn, bool) {
	if o == nil {
    return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddUnboundidMsChapV2SaslMechanismHandlerRequest) SetSchemas(v []EnumunboundidMsChapV2SaslMechanismHandlerSchemaUrn) {
	o.Schemas = v
}

// GetIdentityMapper returns the IdentityMapper field value
func (o *AddUnboundidMsChapV2SaslMechanismHandlerRequest) GetIdentityMapper() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IdentityMapper
}

// GetIdentityMapperOk returns a tuple with the IdentityMapper field value
// and a boolean to check if the value has been set.
func (o *AddUnboundidMsChapV2SaslMechanismHandlerRequest) GetIdentityMapperOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.IdentityMapper, true
}

// SetIdentityMapper sets field value
func (o *AddUnboundidMsChapV2SaslMechanismHandlerRequest) SetIdentityMapper(v string) {
	o.IdentityMapper = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddUnboundidMsChapV2SaslMechanismHandlerRequest) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddUnboundidMsChapV2SaslMechanismHandlerRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddUnboundidMsChapV2SaslMechanismHandlerRequest) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddUnboundidMsChapV2SaslMechanismHandlerRequest) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *AddUnboundidMsChapV2SaslMechanismHandlerRequest) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AddUnboundidMsChapV2SaslMechanismHandlerRequest) GetEnabledOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AddUnboundidMsChapV2SaslMechanismHandlerRequest) SetEnabled(v bool) {
	o.Enabled = v
}

func (o AddUnboundidMsChapV2SaslMechanismHandlerRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["handlerName"] = o.HandlerName
	}
	if true {
		toSerialize["schemas"] = o.Schemas
	}
	if true {
		toSerialize["identityMapper"] = o.IdentityMapper
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableAddUnboundidMsChapV2SaslMechanismHandlerRequest struct {
	value *AddUnboundidMsChapV2SaslMechanismHandlerRequest
	isSet bool
}

func (v NullableAddUnboundidMsChapV2SaslMechanismHandlerRequest) Get() *AddUnboundidMsChapV2SaslMechanismHandlerRequest {
	return v.value
}

func (v *NullableAddUnboundidMsChapV2SaslMechanismHandlerRequest) Set(val *AddUnboundidMsChapV2SaslMechanismHandlerRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddUnboundidMsChapV2SaslMechanismHandlerRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddUnboundidMsChapV2SaslMechanismHandlerRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddUnboundidMsChapV2SaslMechanismHandlerRequest(val *AddUnboundidMsChapV2SaslMechanismHandlerRequest) *NullableAddUnboundidMsChapV2SaslMechanismHandlerRequest {
	return &NullableAddUnboundidMsChapV2SaslMechanismHandlerRequest{value: val, isSet: true}
}

func (v NullableAddUnboundidMsChapV2SaslMechanismHandlerRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddUnboundidMsChapV2SaslMechanismHandlerRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


