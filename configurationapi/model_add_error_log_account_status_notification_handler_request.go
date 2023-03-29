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

// checks if the AddErrorLogAccountStatusNotificationHandlerRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddErrorLogAccountStatusNotificationHandlerRequest{}

// AddErrorLogAccountStatusNotificationHandlerRequest struct for AddErrorLogAccountStatusNotificationHandlerRequest
type AddErrorLogAccountStatusNotificationHandlerRequest struct {
	// Name of the new Account Status Notification Handler
	HandlerName string                                                  `json:"handlerName"`
	Schemas     []EnumerrorLogAccountStatusNotificationHandlerSchemaUrn `json:"schemas"`
	// Indicates which types of event can trigger an account status notification.
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

// NewAddErrorLogAccountStatusNotificationHandlerRequest instantiates a new AddErrorLogAccountStatusNotificationHandlerRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddErrorLogAccountStatusNotificationHandlerRequest(handlerName string, schemas []EnumerrorLogAccountStatusNotificationHandlerSchemaUrn, accountStatusNotificationType []EnumaccountStatusNotificationHandlerAccountStatusNotificationTypeProp, enabled bool) *AddErrorLogAccountStatusNotificationHandlerRequest {
	this := AddErrorLogAccountStatusNotificationHandlerRequest{}
	this.HandlerName = handlerName
	this.Schemas = schemas
	this.AccountStatusNotificationType = accountStatusNotificationType
	this.Enabled = enabled
	return &this
}

// NewAddErrorLogAccountStatusNotificationHandlerRequestWithDefaults instantiates a new AddErrorLogAccountStatusNotificationHandlerRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddErrorLogAccountStatusNotificationHandlerRequestWithDefaults() *AddErrorLogAccountStatusNotificationHandlerRequest {
	this := AddErrorLogAccountStatusNotificationHandlerRequest{}
	return &this
}

// GetHandlerName returns the HandlerName field value
func (o *AddErrorLogAccountStatusNotificationHandlerRequest) GetHandlerName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HandlerName
}

// GetHandlerNameOk returns a tuple with the HandlerName field value
// and a boolean to check if the value has been set.
func (o *AddErrorLogAccountStatusNotificationHandlerRequest) GetHandlerNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HandlerName, true
}

// SetHandlerName sets field value
func (o *AddErrorLogAccountStatusNotificationHandlerRequest) SetHandlerName(v string) {
	o.HandlerName = v
}

// GetSchemas returns the Schemas field value
func (o *AddErrorLogAccountStatusNotificationHandlerRequest) GetSchemas() []EnumerrorLogAccountStatusNotificationHandlerSchemaUrn {
	if o == nil {
		var ret []EnumerrorLogAccountStatusNotificationHandlerSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddErrorLogAccountStatusNotificationHandlerRequest) GetSchemasOk() ([]EnumerrorLogAccountStatusNotificationHandlerSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddErrorLogAccountStatusNotificationHandlerRequest) SetSchemas(v []EnumerrorLogAccountStatusNotificationHandlerSchemaUrn) {
	o.Schemas = v
}

// GetAccountStatusNotificationType returns the AccountStatusNotificationType field value
func (o *AddErrorLogAccountStatusNotificationHandlerRequest) GetAccountStatusNotificationType() []EnumaccountStatusNotificationHandlerAccountStatusNotificationTypeProp {
	if o == nil {
		var ret []EnumaccountStatusNotificationHandlerAccountStatusNotificationTypeProp
		return ret
	}

	return o.AccountStatusNotificationType
}

// GetAccountStatusNotificationTypeOk returns a tuple with the AccountStatusNotificationType field value
// and a boolean to check if the value has been set.
func (o *AddErrorLogAccountStatusNotificationHandlerRequest) GetAccountStatusNotificationTypeOk() ([]EnumaccountStatusNotificationHandlerAccountStatusNotificationTypeProp, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccountStatusNotificationType, true
}

// SetAccountStatusNotificationType sets field value
func (o *AddErrorLogAccountStatusNotificationHandlerRequest) SetAccountStatusNotificationType(v []EnumaccountStatusNotificationHandlerAccountStatusNotificationTypeProp) {
	o.AccountStatusNotificationType = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddErrorLogAccountStatusNotificationHandlerRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddErrorLogAccountStatusNotificationHandlerRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddErrorLogAccountStatusNotificationHandlerRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddErrorLogAccountStatusNotificationHandlerRequest) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *AddErrorLogAccountStatusNotificationHandlerRequest) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AddErrorLogAccountStatusNotificationHandlerRequest) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AddErrorLogAccountStatusNotificationHandlerRequest) SetEnabled(v bool) {
	o.Enabled = v
}

// GetAsynchronous returns the Asynchronous field value if set, zero value otherwise.
func (o *AddErrorLogAccountStatusNotificationHandlerRequest) GetAsynchronous() bool {
	if o == nil || IsNil(o.Asynchronous) {
		var ret bool
		return ret
	}
	return *o.Asynchronous
}

// GetAsynchronousOk returns a tuple with the Asynchronous field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddErrorLogAccountStatusNotificationHandlerRequest) GetAsynchronousOk() (*bool, bool) {
	if o == nil || IsNil(o.Asynchronous) {
		return nil, false
	}
	return o.Asynchronous, true
}

// HasAsynchronous returns a boolean if a field has been set.
func (o *AddErrorLogAccountStatusNotificationHandlerRequest) HasAsynchronous() bool {
	if o != nil && !IsNil(o.Asynchronous) {
		return true
	}

	return false
}

// SetAsynchronous gets a reference to the given bool and assigns it to the Asynchronous field.
func (o *AddErrorLogAccountStatusNotificationHandlerRequest) SetAsynchronous(v bool) {
	o.Asynchronous = &v
}

// GetAccountCreationNotificationRequestCriteria returns the AccountCreationNotificationRequestCriteria field value if set, zero value otherwise.
func (o *AddErrorLogAccountStatusNotificationHandlerRequest) GetAccountCreationNotificationRequestCriteria() string {
	if o == nil || IsNil(o.AccountCreationNotificationRequestCriteria) {
		var ret string
		return ret
	}
	return *o.AccountCreationNotificationRequestCriteria
}

// GetAccountCreationNotificationRequestCriteriaOk returns a tuple with the AccountCreationNotificationRequestCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddErrorLogAccountStatusNotificationHandlerRequest) GetAccountCreationNotificationRequestCriteriaOk() (*string, bool) {
	if o == nil || IsNil(o.AccountCreationNotificationRequestCriteria) {
		return nil, false
	}
	return o.AccountCreationNotificationRequestCriteria, true
}

// HasAccountCreationNotificationRequestCriteria returns a boolean if a field has been set.
func (o *AddErrorLogAccountStatusNotificationHandlerRequest) HasAccountCreationNotificationRequestCriteria() bool {
	if o != nil && !IsNil(o.AccountCreationNotificationRequestCriteria) {
		return true
	}

	return false
}

// SetAccountCreationNotificationRequestCriteria gets a reference to the given string and assigns it to the AccountCreationNotificationRequestCriteria field.
func (o *AddErrorLogAccountStatusNotificationHandlerRequest) SetAccountCreationNotificationRequestCriteria(v string) {
	o.AccountCreationNotificationRequestCriteria = &v
}

// GetAccountUpdateNotificationRequestCriteria returns the AccountUpdateNotificationRequestCriteria field value if set, zero value otherwise.
func (o *AddErrorLogAccountStatusNotificationHandlerRequest) GetAccountUpdateNotificationRequestCriteria() string {
	if o == nil || IsNil(o.AccountUpdateNotificationRequestCriteria) {
		var ret string
		return ret
	}
	return *o.AccountUpdateNotificationRequestCriteria
}

// GetAccountUpdateNotificationRequestCriteriaOk returns a tuple with the AccountUpdateNotificationRequestCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddErrorLogAccountStatusNotificationHandlerRequest) GetAccountUpdateNotificationRequestCriteriaOk() (*string, bool) {
	if o == nil || IsNil(o.AccountUpdateNotificationRequestCriteria) {
		return nil, false
	}
	return o.AccountUpdateNotificationRequestCriteria, true
}

// HasAccountUpdateNotificationRequestCriteria returns a boolean if a field has been set.
func (o *AddErrorLogAccountStatusNotificationHandlerRequest) HasAccountUpdateNotificationRequestCriteria() bool {
	if o != nil && !IsNil(o.AccountUpdateNotificationRequestCriteria) {
		return true
	}

	return false
}

// SetAccountUpdateNotificationRequestCriteria gets a reference to the given string and assigns it to the AccountUpdateNotificationRequestCriteria field.
func (o *AddErrorLogAccountStatusNotificationHandlerRequest) SetAccountUpdateNotificationRequestCriteria(v string) {
	o.AccountUpdateNotificationRequestCriteria = &v
}

func (o AddErrorLogAccountStatusNotificationHandlerRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddErrorLogAccountStatusNotificationHandlerRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["handlerName"] = o.HandlerName
	toSerialize["schemas"] = o.Schemas
	toSerialize["accountStatusNotificationType"] = o.AccountStatusNotificationType
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["enabled"] = o.Enabled
	if !IsNil(o.Asynchronous) {
		toSerialize["asynchronous"] = o.Asynchronous
	}
	if !IsNil(o.AccountCreationNotificationRequestCriteria) {
		toSerialize["accountCreationNotificationRequestCriteria"] = o.AccountCreationNotificationRequestCriteria
	}
	if !IsNil(o.AccountUpdateNotificationRequestCriteria) {
		toSerialize["accountUpdateNotificationRequestCriteria"] = o.AccountUpdateNotificationRequestCriteria
	}
	return toSerialize, nil
}

type NullableAddErrorLogAccountStatusNotificationHandlerRequest struct {
	value *AddErrorLogAccountStatusNotificationHandlerRequest
	isSet bool
}

func (v NullableAddErrorLogAccountStatusNotificationHandlerRequest) Get() *AddErrorLogAccountStatusNotificationHandlerRequest {
	return v.value
}

func (v *NullableAddErrorLogAccountStatusNotificationHandlerRequest) Set(val *AddErrorLogAccountStatusNotificationHandlerRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddErrorLogAccountStatusNotificationHandlerRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddErrorLogAccountStatusNotificationHandlerRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddErrorLogAccountStatusNotificationHandlerRequest(val *AddErrorLogAccountStatusNotificationHandlerRequest) *NullableAddErrorLogAccountStatusNotificationHandlerRequest {
	return &NullableAddErrorLogAccountStatusNotificationHandlerRequest{value: val, isSet: true}
}

func (v NullableAddErrorLogAccountStatusNotificationHandlerRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddErrorLogAccountStatusNotificationHandlerRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
