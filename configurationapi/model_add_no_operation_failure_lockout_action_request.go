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

// checks if the AddNoOperationFailureLockoutActionRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddNoOperationFailureLockoutActionRequest{}

// AddNoOperationFailureLockoutActionRequest struct for AddNoOperationFailureLockoutActionRequest
type AddNoOperationFailureLockoutActionRequest struct {
	Schemas []EnumnoOperationFailureLockoutActionSchemaUrn `json:"schemas"`
	// Indicates whether to generate an account status notification for cases in which this failure lockout action is invoked for a bind attempt with too many outstanding authentication failures.
	GenerateAccountStatusNotification *bool `json:"generateAccountStatusNotification,omitempty"`
	// A description for this Failure Lockout Action
	Description *string `json:"description,omitempty"`
	// Name of the new Failure Lockout Action
	ActionName string `json:"actionName"`
}

// NewAddNoOperationFailureLockoutActionRequest instantiates a new AddNoOperationFailureLockoutActionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddNoOperationFailureLockoutActionRequest(schemas []EnumnoOperationFailureLockoutActionSchemaUrn, actionName string) *AddNoOperationFailureLockoutActionRequest {
	this := AddNoOperationFailureLockoutActionRequest{}
	this.Schemas = schemas
	this.ActionName = actionName
	return &this
}

// NewAddNoOperationFailureLockoutActionRequestWithDefaults instantiates a new AddNoOperationFailureLockoutActionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddNoOperationFailureLockoutActionRequestWithDefaults() *AddNoOperationFailureLockoutActionRequest {
	this := AddNoOperationFailureLockoutActionRequest{}
	return &this
}

// GetSchemas returns the Schemas field value
func (o *AddNoOperationFailureLockoutActionRequest) GetSchemas() []EnumnoOperationFailureLockoutActionSchemaUrn {
	if o == nil {
		var ret []EnumnoOperationFailureLockoutActionSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddNoOperationFailureLockoutActionRequest) GetSchemasOk() ([]EnumnoOperationFailureLockoutActionSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddNoOperationFailureLockoutActionRequest) SetSchemas(v []EnumnoOperationFailureLockoutActionSchemaUrn) {
	o.Schemas = v
}

// GetGenerateAccountStatusNotification returns the GenerateAccountStatusNotification field value if set, zero value otherwise.
func (o *AddNoOperationFailureLockoutActionRequest) GetGenerateAccountStatusNotification() bool {
	if o == nil || IsNil(o.GenerateAccountStatusNotification) {
		var ret bool
		return ret
	}
	return *o.GenerateAccountStatusNotification
}

// GetGenerateAccountStatusNotificationOk returns a tuple with the GenerateAccountStatusNotification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddNoOperationFailureLockoutActionRequest) GetGenerateAccountStatusNotificationOk() (*bool, bool) {
	if o == nil || IsNil(o.GenerateAccountStatusNotification) {
		return nil, false
	}
	return o.GenerateAccountStatusNotification, true
}

// HasGenerateAccountStatusNotification returns a boolean if a field has been set.
func (o *AddNoOperationFailureLockoutActionRequest) HasGenerateAccountStatusNotification() bool {
	if o != nil && !IsNil(o.GenerateAccountStatusNotification) {
		return true
	}

	return false
}

// SetGenerateAccountStatusNotification gets a reference to the given bool and assigns it to the GenerateAccountStatusNotification field.
func (o *AddNoOperationFailureLockoutActionRequest) SetGenerateAccountStatusNotification(v bool) {
	o.GenerateAccountStatusNotification = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddNoOperationFailureLockoutActionRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddNoOperationFailureLockoutActionRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddNoOperationFailureLockoutActionRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddNoOperationFailureLockoutActionRequest) SetDescription(v string) {
	o.Description = &v
}

// GetActionName returns the ActionName field value
func (o *AddNoOperationFailureLockoutActionRequest) GetActionName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ActionName
}

// GetActionNameOk returns a tuple with the ActionName field value
// and a boolean to check if the value has been set.
func (o *AddNoOperationFailureLockoutActionRequest) GetActionNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActionName, true
}

// SetActionName sets field value
func (o *AddNoOperationFailureLockoutActionRequest) SetActionName(v string) {
	o.ActionName = v
}

func (o AddNoOperationFailureLockoutActionRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddNoOperationFailureLockoutActionRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["schemas"] = o.Schemas
	if !IsNil(o.GenerateAccountStatusNotification) {
		toSerialize["generateAccountStatusNotification"] = o.GenerateAccountStatusNotification
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["actionName"] = o.ActionName
	return toSerialize, nil
}

type NullableAddNoOperationFailureLockoutActionRequest struct {
	value *AddNoOperationFailureLockoutActionRequest
	isSet bool
}

func (v NullableAddNoOperationFailureLockoutActionRequest) Get() *AddNoOperationFailureLockoutActionRequest {
	return v.value
}

func (v *NullableAddNoOperationFailureLockoutActionRequest) Set(val *AddNoOperationFailureLockoutActionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddNoOperationFailureLockoutActionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddNoOperationFailureLockoutActionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddNoOperationFailureLockoutActionRequest(val *AddNoOperationFailureLockoutActionRequest) *NullableAddNoOperationFailureLockoutActionRequest {
	return &NullableAddNoOperationFailureLockoutActionRequest{value: val, isSet: true}
}

func (v NullableAddNoOperationFailureLockoutActionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddNoOperationFailureLockoutActionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
