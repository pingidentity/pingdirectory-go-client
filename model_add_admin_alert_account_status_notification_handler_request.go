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

// AddAdminAlertAccountStatusNotificationHandlerRequest struct for AddAdminAlertAccountStatusNotificationHandlerRequest
type AddAdminAlertAccountStatusNotificationHandlerRequest struct {
	// Name of the new Account Status Notification Handler
	HandlerName string `json:"handlerName"`
	Schemas []EnumadminAlertAccountStatusNotificationHandlerSchemaUrn `json:"schemas"`
	AccountStatusNotificationType []EnumaccountStatusNotificationHandlerAccountStatusNotificationTypeProp `json:"accountStatusNotificationType"`
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

// NewAddAdminAlertAccountStatusNotificationHandlerRequest instantiates a new AddAdminAlertAccountStatusNotificationHandlerRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddAdminAlertAccountStatusNotificationHandlerRequest(handlerName string, schemas []EnumadminAlertAccountStatusNotificationHandlerSchemaUrn, accountStatusNotificationType []EnumaccountStatusNotificationHandlerAccountStatusNotificationTypeProp, enabled bool) *AddAdminAlertAccountStatusNotificationHandlerRequest {
	this := AddAdminAlertAccountStatusNotificationHandlerRequest{}
	this.HandlerName = handlerName
	this.Schemas = schemas
	this.AccountStatusNotificationType = accountStatusNotificationType
	this.Enabled = enabled
	return &this
}

// NewAddAdminAlertAccountStatusNotificationHandlerRequestWithDefaults instantiates a new AddAdminAlertAccountStatusNotificationHandlerRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddAdminAlertAccountStatusNotificationHandlerRequestWithDefaults() *AddAdminAlertAccountStatusNotificationHandlerRequest {
	this := AddAdminAlertAccountStatusNotificationHandlerRequest{}
	return &this
}

// GetHandlerName returns the HandlerName field value
func (o *AddAdminAlertAccountStatusNotificationHandlerRequest) GetHandlerName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HandlerName
}

// GetHandlerNameOk returns a tuple with the HandlerName field value
// and a boolean to check if the value has been set.
func (o *AddAdminAlertAccountStatusNotificationHandlerRequest) GetHandlerNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.HandlerName, true
}

// SetHandlerName sets field value
func (o *AddAdminAlertAccountStatusNotificationHandlerRequest) SetHandlerName(v string) {
	o.HandlerName = v
}

// GetSchemas returns the Schemas field value
func (o *AddAdminAlertAccountStatusNotificationHandlerRequest) GetSchemas() []EnumadminAlertAccountStatusNotificationHandlerSchemaUrn {
	if o == nil {
		var ret []EnumadminAlertAccountStatusNotificationHandlerSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddAdminAlertAccountStatusNotificationHandlerRequest) GetSchemasOk() ([]EnumadminAlertAccountStatusNotificationHandlerSchemaUrn, bool) {
	if o == nil {
    return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddAdminAlertAccountStatusNotificationHandlerRequest) SetSchemas(v []EnumadminAlertAccountStatusNotificationHandlerSchemaUrn) {
	o.Schemas = v
}

// GetAccountStatusNotificationType returns the AccountStatusNotificationType field value
func (o *AddAdminAlertAccountStatusNotificationHandlerRequest) GetAccountStatusNotificationType() []EnumaccountStatusNotificationHandlerAccountStatusNotificationTypeProp {
	if o == nil {
		var ret []EnumaccountStatusNotificationHandlerAccountStatusNotificationTypeProp
		return ret
	}

	return o.AccountStatusNotificationType
}

// GetAccountStatusNotificationTypeOk returns a tuple with the AccountStatusNotificationType field value
// and a boolean to check if the value has been set.
func (o *AddAdminAlertAccountStatusNotificationHandlerRequest) GetAccountStatusNotificationTypeOk() ([]EnumaccountStatusNotificationHandlerAccountStatusNotificationTypeProp, bool) {
	if o == nil {
    return nil, false
	}
	return o.AccountStatusNotificationType, true
}

// SetAccountStatusNotificationType sets field value
func (o *AddAdminAlertAccountStatusNotificationHandlerRequest) SetAccountStatusNotificationType(v []EnumaccountStatusNotificationHandlerAccountStatusNotificationTypeProp) {
	o.AccountStatusNotificationType = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddAdminAlertAccountStatusNotificationHandlerRequest) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddAdminAlertAccountStatusNotificationHandlerRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddAdminAlertAccountStatusNotificationHandlerRequest) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddAdminAlertAccountStatusNotificationHandlerRequest) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *AddAdminAlertAccountStatusNotificationHandlerRequest) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AddAdminAlertAccountStatusNotificationHandlerRequest) GetEnabledOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AddAdminAlertAccountStatusNotificationHandlerRequest) SetEnabled(v bool) {
	o.Enabled = v
}

// GetAsynchronous returns the Asynchronous field value if set, zero value otherwise.
func (o *AddAdminAlertAccountStatusNotificationHandlerRequest) GetAsynchronous() bool {
	if o == nil || isNil(o.Asynchronous) {
		var ret bool
		return ret
	}
	return *o.Asynchronous
}

// GetAsynchronousOk returns a tuple with the Asynchronous field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddAdminAlertAccountStatusNotificationHandlerRequest) GetAsynchronousOk() (*bool, bool) {
	if o == nil || isNil(o.Asynchronous) {
    return nil, false
	}
	return o.Asynchronous, true
}

// HasAsynchronous returns a boolean if a field has been set.
func (o *AddAdminAlertAccountStatusNotificationHandlerRequest) HasAsynchronous() bool {
	if o != nil && !isNil(o.Asynchronous) {
		return true
	}

	return false
}

// SetAsynchronous gets a reference to the given bool and assigns it to the Asynchronous field.
func (o *AddAdminAlertAccountStatusNotificationHandlerRequest) SetAsynchronous(v bool) {
	o.Asynchronous = &v
}

// GetAccountCreationNotificationRequestCriteria returns the AccountCreationNotificationRequestCriteria field value if set, zero value otherwise.
func (o *AddAdminAlertAccountStatusNotificationHandlerRequest) GetAccountCreationNotificationRequestCriteria() string {
	if o == nil || isNil(o.AccountCreationNotificationRequestCriteria) {
		var ret string
		return ret
	}
	return *o.AccountCreationNotificationRequestCriteria
}

// GetAccountCreationNotificationRequestCriteriaOk returns a tuple with the AccountCreationNotificationRequestCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddAdminAlertAccountStatusNotificationHandlerRequest) GetAccountCreationNotificationRequestCriteriaOk() (*string, bool) {
	if o == nil || isNil(o.AccountCreationNotificationRequestCriteria) {
    return nil, false
	}
	return o.AccountCreationNotificationRequestCriteria, true
}

// HasAccountCreationNotificationRequestCriteria returns a boolean if a field has been set.
func (o *AddAdminAlertAccountStatusNotificationHandlerRequest) HasAccountCreationNotificationRequestCriteria() bool {
	if o != nil && !isNil(o.AccountCreationNotificationRequestCriteria) {
		return true
	}

	return false
}

// SetAccountCreationNotificationRequestCriteria gets a reference to the given string and assigns it to the AccountCreationNotificationRequestCriteria field.
func (o *AddAdminAlertAccountStatusNotificationHandlerRequest) SetAccountCreationNotificationRequestCriteria(v string) {
	o.AccountCreationNotificationRequestCriteria = &v
}

// GetAccountUpdateNotificationRequestCriteria returns the AccountUpdateNotificationRequestCriteria field value if set, zero value otherwise.
func (o *AddAdminAlertAccountStatusNotificationHandlerRequest) GetAccountUpdateNotificationRequestCriteria() string {
	if o == nil || isNil(o.AccountUpdateNotificationRequestCriteria) {
		var ret string
		return ret
	}
	return *o.AccountUpdateNotificationRequestCriteria
}

// GetAccountUpdateNotificationRequestCriteriaOk returns a tuple with the AccountUpdateNotificationRequestCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddAdminAlertAccountStatusNotificationHandlerRequest) GetAccountUpdateNotificationRequestCriteriaOk() (*string, bool) {
	if o == nil || isNil(o.AccountUpdateNotificationRequestCriteria) {
    return nil, false
	}
	return o.AccountUpdateNotificationRequestCriteria, true
}

// HasAccountUpdateNotificationRequestCriteria returns a boolean if a field has been set.
func (o *AddAdminAlertAccountStatusNotificationHandlerRequest) HasAccountUpdateNotificationRequestCriteria() bool {
	if o != nil && !isNil(o.AccountUpdateNotificationRequestCriteria) {
		return true
	}

	return false
}

// SetAccountUpdateNotificationRequestCriteria gets a reference to the given string and assigns it to the AccountUpdateNotificationRequestCriteria field.
func (o *AddAdminAlertAccountStatusNotificationHandlerRequest) SetAccountUpdateNotificationRequestCriteria(v string) {
	o.AccountUpdateNotificationRequestCriteria = &v
}

func (o AddAdminAlertAccountStatusNotificationHandlerRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["handlerName"] = o.HandlerName
	}
	if true {
		toSerialize["schemas"] = o.Schemas
	}
	if true {
		toSerialize["accountStatusNotificationType"] = o.AccountStatusNotificationType
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

type NullableAddAdminAlertAccountStatusNotificationHandlerRequest struct {
	value *AddAdminAlertAccountStatusNotificationHandlerRequest
	isSet bool
}

func (v NullableAddAdminAlertAccountStatusNotificationHandlerRequest) Get() *AddAdminAlertAccountStatusNotificationHandlerRequest {
	return v.value
}

func (v *NullableAddAdminAlertAccountStatusNotificationHandlerRequest) Set(val *AddAdminAlertAccountStatusNotificationHandlerRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddAdminAlertAccountStatusNotificationHandlerRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddAdminAlertAccountStatusNotificationHandlerRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddAdminAlertAccountStatusNotificationHandlerRequest(val *AddAdminAlertAccountStatusNotificationHandlerRequest) *NullableAddAdminAlertAccountStatusNotificationHandlerRequest {
	return &NullableAddAdminAlertAccountStatusNotificationHandlerRequest{value: val, isSet: true}
}

func (v NullableAddAdminAlertAccountStatusNotificationHandlerRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddAdminAlertAccountStatusNotificationHandlerRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


