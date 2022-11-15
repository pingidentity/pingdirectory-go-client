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

// AddThirdPartyAccountStatusNotificationHandlerRequest struct for AddThirdPartyAccountStatusNotificationHandlerRequest
type AddThirdPartyAccountStatusNotificationHandlerRequest struct {
	// Name of the new Account Status Notification Handler
	HandlerName string `json:"handlerName"`
	Schemas []EnumthirdPartyAccountStatusNotificationHandlerSchemaUrn `json:"schemas"`
	// The fully-qualified name of the Java class providing the logic for the Third Party Account Status Notification Handler.
	ExtensionClass string `json:"extensionClass"`
	ExtensionArgument []string `json:"extensionArgument,omitempty"`
	// A description for this Account Status Notification Handler
	Description *string `json:"description,omitempty"`
	// Indicates whether the Account Status Notification Handler is enabled. Only enabled handlers are invoked whenever a related event occurs in the server.
	Enabled bool `json:"enabled"`
	// Indicates whether the server should attempt to invoke this Account Status Notification Handler in a background thread so that any potentially-expensive processing (e.g., performing network communication to deliver a message) will not delay processing for the operation that triggered the notification.
	Asynchronous *bool `json:"asynchronous,omitempty"`
	// A request criteria object that identifies which add requests should result in account creation notifications for this handler.
	AccountCreationNotificationRequestCriteria *string `json:"accountCreationNotificationRequestCriteria,omitempty"`
	// A request criteria object that identifies which modify and modify DN requests should result in account update notifications for this handler.
	AccountUpdateNotificationRequestCriteria *string `json:"accountUpdateNotificationRequestCriteria,omitempty"`
}

// NewAddThirdPartyAccountStatusNotificationHandlerRequest instantiates a new AddThirdPartyAccountStatusNotificationHandlerRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddThirdPartyAccountStatusNotificationHandlerRequest(handlerName string, schemas []EnumthirdPartyAccountStatusNotificationHandlerSchemaUrn, extensionClass string, enabled bool) *AddThirdPartyAccountStatusNotificationHandlerRequest {
	this := AddThirdPartyAccountStatusNotificationHandlerRequest{}
	this.HandlerName = handlerName
	this.Schemas = schemas
	this.ExtensionClass = extensionClass
	this.Enabled = enabled
	return &this
}

// NewAddThirdPartyAccountStatusNotificationHandlerRequestWithDefaults instantiates a new AddThirdPartyAccountStatusNotificationHandlerRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddThirdPartyAccountStatusNotificationHandlerRequestWithDefaults() *AddThirdPartyAccountStatusNotificationHandlerRequest {
	this := AddThirdPartyAccountStatusNotificationHandlerRequest{}
	return &this
}

// GetHandlerName returns the HandlerName field value
func (o *AddThirdPartyAccountStatusNotificationHandlerRequest) GetHandlerName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HandlerName
}

// GetHandlerNameOk returns a tuple with the HandlerName field value
// and a boolean to check if the value has been set.
func (o *AddThirdPartyAccountStatusNotificationHandlerRequest) GetHandlerNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.HandlerName, true
}

// SetHandlerName sets field value
func (o *AddThirdPartyAccountStatusNotificationHandlerRequest) SetHandlerName(v string) {
	o.HandlerName = v
}

// GetSchemas returns the Schemas field value
func (o *AddThirdPartyAccountStatusNotificationHandlerRequest) GetSchemas() []EnumthirdPartyAccountStatusNotificationHandlerSchemaUrn {
	if o == nil {
		var ret []EnumthirdPartyAccountStatusNotificationHandlerSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddThirdPartyAccountStatusNotificationHandlerRequest) GetSchemasOk() ([]EnumthirdPartyAccountStatusNotificationHandlerSchemaUrn, bool) {
	if o == nil {
    return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddThirdPartyAccountStatusNotificationHandlerRequest) SetSchemas(v []EnumthirdPartyAccountStatusNotificationHandlerSchemaUrn) {
	o.Schemas = v
}

// GetExtensionClass returns the ExtensionClass field value
func (o *AddThirdPartyAccountStatusNotificationHandlerRequest) GetExtensionClass() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExtensionClass
}

// GetExtensionClassOk returns a tuple with the ExtensionClass field value
// and a boolean to check if the value has been set.
func (o *AddThirdPartyAccountStatusNotificationHandlerRequest) GetExtensionClassOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ExtensionClass, true
}

// SetExtensionClass sets field value
func (o *AddThirdPartyAccountStatusNotificationHandlerRequest) SetExtensionClass(v string) {
	o.ExtensionClass = v
}

// GetExtensionArgument returns the ExtensionArgument field value if set, zero value otherwise.
func (o *AddThirdPartyAccountStatusNotificationHandlerRequest) GetExtensionArgument() []string {
	if o == nil || isNil(o.ExtensionArgument) {
		var ret []string
		return ret
	}
	return o.ExtensionArgument
}

// GetExtensionArgumentOk returns a tuple with the ExtensionArgument field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddThirdPartyAccountStatusNotificationHandlerRequest) GetExtensionArgumentOk() ([]string, bool) {
	if o == nil || isNil(o.ExtensionArgument) {
    return nil, false
	}
	return o.ExtensionArgument, true
}

// HasExtensionArgument returns a boolean if a field has been set.
func (o *AddThirdPartyAccountStatusNotificationHandlerRequest) HasExtensionArgument() bool {
	if o != nil && !isNil(o.ExtensionArgument) {
		return true
	}

	return false
}

// SetExtensionArgument gets a reference to the given []string and assigns it to the ExtensionArgument field.
func (o *AddThirdPartyAccountStatusNotificationHandlerRequest) SetExtensionArgument(v []string) {
	o.ExtensionArgument = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddThirdPartyAccountStatusNotificationHandlerRequest) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddThirdPartyAccountStatusNotificationHandlerRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddThirdPartyAccountStatusNotificationHandlerRequest) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddThirdPartyAccountStatusNotificationHandlerRequest) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *AddThirdPartyAccountStatusNotificationHandlerRequest) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AddThirdPartyAccountStatusNotificationHandlerRequest) GetEnabledOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AddThirdPartyAccountStatusNotificationHandlerRequest) SetEnabled(v bool) {
	o.Enabled = v
}

// GetAsynchronous returns the Asynchronous field value if set, zero value otherwise.
func (o *AddThirdPartyAccountStatusNotificationHandlerRequest) GetAsynchronous() bool {
	if o == nil || isNil(o.Asynchronous) {
		var ret bool
		return ret
	}
	return *o.Asynchronous
}

// GetAsynchronousOk returns a tuple with the Asynchronous field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddThirdPartyAccountStatusNotificationHandlerRequest) GetAsynchronousOk() (*bool, bool) {
	if o == nil || isNil(o.Asynchronous) {
    return nil, false
	}
	return o.Asynchronous, true
}

// HasAsynchronous returns a boolean if a field has been set.
func (o *AddThirdPartyAccountStatusNotificationHandlerRequest) HasAsynchronous() bool {
	if o != nil && !isNil(o.Asynchronous) {
		return true
	}

	return false
}

// SetAsynchronous gets a reference to the given bool and assigns it to the Asynchronous field.
func (o *AddThirdPartyAccountStatusNotificationHandlerRequest) SetAsynchronous(v bool) {
	o.Asynchronous = &v
}

// GetAccountCreationNotificationRequestCriteria returns the AccountCreationNotificationRequestCriteria field value if set, zero value otherwise.
func (o *AddThirdPartyAccountStatusNotificationHandlerRequest) GetAccountCreationNotificationRequestCriteria() string {
	if o == nil || isNil(o.AccountCreationNotificationRequestCriteria) {
		var ret string
		return ret
	}
	return *o.AccountCreationNotificationRequestCriteria
}

// GetAccountCreationNotificationRequestCriteriaOk returns a tuple with the AccountCreationNotificationRequestCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddThirdPartyAccountStatusNotificationHandlerRequest) GetAccountCreationNotificationRequestCriteriaOk() (*string, bool) {
	if o == nil || isNil(o.AccountCreationNotificationRequestCriteria) {
    return nil, false
	}
	return o.AccountCreationNotificationRequestCriteria, true
}

// HasAccountCreationNotificationRequestCriteria returns a boolean if a field has been set.
func (o *AddThirdPartyAccountStatusNotificationHandlerRequest) HasAccountCreationNotificationRequestCriteria() bool {
	if o != nil && !isNil(o.AccountCreationNotificationRequestCriteria) {
		return true
	}

	return false
}

// SetAccountCreationNotificationRequestCriteria gets a reference to the given string and assigns it to the AccountCreationNotificationRequestCriteria field.
func (o *AddThirdPartyAccountStatusNotificationHandlerRequest) SetAccountCreationNotificationRequestCriteria(v string) {
	o.AccountCreationNotificationRequestCriteria = &v
}

// GetAccountUpdateNotificationRequestCriteria returns the AccountUpdateNotificationRequestCriteria field value if set, zero value otherwise.
func (o *AddThirdPartyAccountStatusNotificationHandlerRequest) GetAccountUpdateNotificationRequestCriteria() string {
	if o == nil || isNil(o.AccountUpdateNotificationRequestCriteria) {
		var ret string
		return ret
	}
	return *o.AccountUpdateNotificationRequestCriteria
}

// GetAccountUpdateNotificationRequestCriteriaOk returns a tuple with the AccountUpdateNotificationRequestCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddThirdPartyAccountStatusNotificationHandlerRequest) GetAccountUpdateNotificationRequestCriteriaOk() (*string, bool) {
	if o == nil || isNil(o.AccountUpdateNotificationRequestCriteria) {
    return nil, false
	}
	return o.AccountUpdateNotificationRequestCriteria, true
}

// HasAccountUpdateNotificationRequestCriteria returns a boolean if a field has been set.
func (o *AddThirdPartyAccountStatusNotificationHandlerRequest) HasAccountUpdateNotificationRequestCriteria() bool {
	if o != nil && !isNil(o.AccountUpdateNotificationRequestCriteria) {
		return true
	}

	return false
}

// SetAccountUpdateNotificationRequestCriteria gets a reference to the given string and assigns it to the AccountUpdateNotificationRequestCriteria field.
func (o *AddThirdPartyAccountStatusNotificationHandlerRequest) SetAccountUpdateNotificationRequestCriteria(v string) {
	o.AccountUpdateNotificationRequestCriteria = &v
}

func (o AddThirdPartyAccountStatusNotificationHandlerRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["handlerName"] = o.HandlerName
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
	if !isNil(o.Asynchronous) {
		toSerialize["asynchronous"] = o.Asynchronous
	}
	if !isNil(o.AccountCreationNotificationRequestCriteria) {
		toSerialize["accountCreationNotificationRequestCriteria"] = o.AccountCreationNotificationRequestCriteria
	}
	if !isNil(o.AccountUpdateNotificationRequestCriteria) {
		toSerialize["accountUpdateNotificationRequestCriteria"] = o.AccountUpdateNotificationRequestCriteria
	}
	return json.Marshal(toSerialize)
}

type NullableAddThirdPartyAccountStatusNotificationHandlerRequest struct {
	value *AddThirdPartyAccountStatusNotificationHandlerRequest
	isSet bool
}

func (v NullableAddThirdPartyAccountStatusNotificationHandlerRequest) Get() *AddThirdPartyAccountStatusNotificationHandlerRequest {
	return v.value
}

func (v *NullableAddThirdPartyAccountStatusNotificationHandlerRequest) Set(val *AddThirdPartyAccountStatusNotificationHandlerRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddThirdPartyAccountStatusNotificationHandlerRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddThirdPartyAccountStatusNotificationHandlerRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddThirdPartyAccountStatusNotificationHandlerRequest(val *AddThirdPartyAccountStatusNotificationHandlerRequest) *NullableAddThirdPartyAccountStatusNotificationHandlerRequest {
	return &NullableAddThirdPartyAccountStatusNotificationHandlerRequest{value: val, isSet: true}
}

func (v NullableAddThirdPartyAccountStatusNotificationHandlerRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddThirdPartyAccountStatusNotificationHandlerRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


