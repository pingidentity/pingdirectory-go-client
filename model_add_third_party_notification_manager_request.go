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

// AddThirdPartyNotificationManagerRequest struct for AddThirdPartyNotificationManagerRequest
type AddThirdPartyNotificationManagerRequest struct {
	// Name of the new Notification Manager
	ManagerName string `json:"managerName"`
	Schemas []EnumthirdPartyNotificationManagerSchemaUrn `json:"schemas,omitempty"`
	// The fully-qualified name of the Java class providing the logic for the Third Party Notification Manager.
	ExtensionClass string `json:"extensionClass"`
	ExtensionArgument []string `json:"extensionArgument,omitempty"`
	// A description for this Notification Manager
	Description *string `json:"description,omitempty"`
	// Indicates whether this Notification Manager is enabled within the server.
	Enabled bool `json:"enabled"`
	// Specifies the DN of the entry below which subscription data is stored for this Notification Manager. This needs to be in the backend that has the data to be notified on, and must not be the same entry as the backend base DN. The subscription base DN entry does not need to exist as it will be created by the server.
	SubscriptionBaseDN string `json:"subscriptionBaseDN"`
	TransactionNotification EnumnotificationManagerTransactionNotificationProp `json:"transactionNotification"`
	// Enables monitor entries for this Notification Manager.
	MonitorEntriesEnabled *bool `json:"monitorEntriesEnabled,omitempty"`
}

// NewAddThirdPartyNotificationManagerRequest instantiates a new AddThirdPartyNotificationManagerRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddThirdPartyNotificationManagerRequest(managerName string, extensionClass string, enabled bool, subscriptionBaseDN string, transactionNotification EnumnotificationManagerTransactionNotificationProp) *AddThirdPartyNotificationManagerRequest {
	this := AddThirdPartyNotificationManagerRequest{}
	this.ManagerName = managerName
	this.ExtensionClass = extensionClass
	this.Enabled = enabled
	this.SubscriptionBaseDN = subscriptionBaseDN
	this.TransactionNotification = transactionNotification
	return &this
}

// NewAddThirdPartyNotificationManagerRequestWithDefaults instantiates a new AddThirdPartyNotificationManagerRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddThirdPartyNotificationManagerRequestWithDefaults() *AddThirdPartyNotificationManagerRequest {
	this := AddThirdPartyNotificationManagerRequest{}
	return &this
}

// GetManagerName returns the ManagerName field value
func (o *AddThirdPartyNotificationManagerRequest) GetManagerName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ManagerName
}

// GetManagerNameOk returns a tuple with the ManagerName field value
// and a boolean to check if the value has been set.
func (o *AddThirdPartyNotificationManagerRequest) GetManagerNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ManagerName, true
}

// SetManagerName sets field value
func (o *AddThirdPartyNotificationManagerRequest) SetManagerName(v string) {
	o.ManagerName = v
}

// GetSchemas returns the Schemas field value if set, zero value otherwise.
func (o *AddThirdPartyNotificationManagerRequest) GetSchemas() []EnumthirdPartyNotificationManagerSchemaUrn {
	if o == nil || isNil(o.Schemas) {
		var ret []EnumthirdPartyNotificationManagerSchemaUrn
		return ret
	}
	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddThirdPartyNotificationManagerRequest) GetSchemasOk() ([]EnumthirdPartyNotificationManagerSchemaUrn, bool) {
	if o == nil || isNil(o.Schemas) {
    return nil, false
	}
	return o.Schemas, true
}

// HasSchemas returns a boolean if a field has been set.
func (o *AddThirdPartyNotificationManagerRequest) HasSchemas() bool {
	if o != nil && !isNil(o.Schemas) {
		return true
	}

	return false
}

// SetSchemas gets a reference to the given []EnumthirdPartyNotificationManagerSchemaUrn and assigns it to the Schemas field.
func (o *AddThirdPartyNotificationManagerRequest) SetSchemas(v []EnumthirdPartyNotificationManagerSchemaUrn) {
	o.Schemas = v
}

// GetExtensionClass returns the ExtensionClass field value
func (o *AddThirdPartyNotificationManagerRequest) GetExtensionClass() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExtensionClass
}

// GetExtensionClassOk returns a tuple with the ExtensionClass field value
// and a boolean to check if the value has been set.
func (o *AddThirdPartyNotificationManagerRequest) GetExtensionClassOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ExtensionClass, true
}

// SetExtensionClass sets field value
func (o *AddThirdPartyNotificationManagerRequest) SetExtensionClass(v string) {
	o.ExtensionClass = v
}

// GetExtensionArgument returns the ExtensionArgument field value if set, zero value otherwise.
func (o *AddThirdPartyNotificationManagerRequest) GetExtensionArgument() []string {
	if o == nil || isNil(o.ExtensionArgument) {
		var ret []string
		return ret
	}
	return o.ExtensionArgument
}

// GetExtensionArgumentOk returns a tuple with the ExtensionArgument field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddThirdPartyNotificationManagerRequest) GetExtensionArgumentOk() ([]string, bool) {
	if o == nil || isNil(o.ExtensionArgument) {
    return nil, false
	}
	return o.ExtensionArgument, true
}

// HasExtensionArgument returns a boolean if a field has been set.
func (o *AddThirdPartyNotificationManagerRequest) HasExtensionArgument() bool {
	if o != nil && !isNil(o.ExtensionArgument) {
		return true
	}

	return false
}

// SetExtensionArgument gets a reference to the given []string and assigns it to the ExtensionArgument field.
func (o *AddThirdPartyNotificationManagerRequest) SetExtensionArgument(v []string) {
	o.ExtensionArgument = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddThirdPartyNotificationManagerRequest) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddThirdPartyNotificationManagerRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddThirdPartyNotificationManagerRequest) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddThirdPartyNotificationManagerRequest) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *AddThirdPartyNotificationManagerRequest) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AddThirdPartyNotificationManagerRequest) GetEnabledOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AddThirdPartyNotificationManagerRequest) SetEnabled(v bool) {
	o.Enabled = v
}

// GetSubscriptionBaseDN returns the SubscriptionBaseDN field value
func (o *AddThirdPartyNotificationManagerRequest) GetSubscriptionBaseDN() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubscriptionBaseDN
}

// GetSubscriptionBaseDNOk returns a tuple with the SubscriptionBaseDN field value
// and a boolean to check if the value has been set.
func (o *AddThirdPartyNotificationManagerRequest) GetSubscriptionBaseDNOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.SubscriptionBaseDN, true
}

// SetSubscriptionBaseDN sets field value
func (o *AddThirdPartyNotificationManagerRequest) SetSubscriptionBaseDN(v string) {
	o.SubscriptionBaseDN = v
}

// GetTransactionNotification returns the TransactionNotification field value
func (o *AddThirdPartyNotificationManagerRequest) GetTransactionNotification() EnumnotificationManagerTransactionNotificationProp {
	if o == nil {
		var ret EnumnotificationManagerTransactionNotificationProp
		return ret
	}

	return o.TransactionNotification
}

// GetTransactionNotificationOk returns a tuple with the TransactionNotification field value
// and a boolean to check if the value has been set.
func (o *AddThirdPartyNotificationManagerRequest) GetTransactionNotificationOk() (*EnumnotificationManagerTransactionNotificationProp, bool) {
	if o == nil {
    return nil, false
	}
	return &o.TransactionNotification, true
}

// SetTransactionNotification sets field value
func (o *AddThirdPartyNotificationManagerRequest) SetTransactionNotification(v EnumnotificationManagerTransactionNotificationProp) {
	o.TransactionNotification = v
}

// GetMonitorEntriesEnabled returns the MonitorEntriesEnabled field value if set, zero value otherwise.
func (o *AddThirdPartyNotificationManagerRequest) GetMonitorEntriesEnabled() bool {
	if o == nil || isNil(o.MonitorEntriesEnabled) {
		var ret bool
		return ret
	}
	return *o.MonitorEntriesEnabled
}

// GetMonitorEntriesEnabledOk returns a tuple with the MonitorEntriesEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddThirdPartyNotificationManagerRequest) GetMonitorEntriesEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.MonitorEntriesEnabled) {
    return nil, false
	}
	return o.MonitorEntriesEnabled, true
}

// HasMonitorEntriesEnabled returns a boolean if a field has been set.
func (o *AddThirdPartyNotificationManagerRequest) HasMonitorEntriesEnabled() bool {
	if o != nil && !isNil(o.MonitorEntriesEnabled) {
		return true
	}

	return false
}

// SetMonitorEntriesEnabled gets a reference to the given bool and assigns it to the MonitorEntriesEnabled field.
func (o *AddThirdPartyNotificationManagerRequest) SetMonitorEntriesEnabled(v bool) {
	o.MonitorEntriesEnabled = &v
}

func (o AddThirdPartyNotificationManagerRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["managerName"] = o.ManagerName
	}
	if !isNil(o.Schemas) {
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
	if true {
		toSerialize["subscriptionBaseDN"] = o.SubscriptionBaseDN
	}
	if true {
		toSerialize["transactionNotification"] = o.TransactionNotification
	}
	if !isNil(o.MonitorEntriesEnabled) {
		toSerialize["monitorEntriesEnabled"] = o.MonitorEntriesEnabled
	}
	return json.Marshal(toSerialize)
}

type NullableAddThirdPartyNotificationManagerRequest struct {
	value *AddThirdPartyNotificationManagerRequest
	isSet bool
}

func (v NullableAddThirdPartyNotificationManagerRequest) Get() *AddThirdPartyNotificationManagerRequest {
	return v.value
}

func (v *NullableAddThirdPartyNotificationManagerRequest) Set(val *AddThirdPartyNotificationManagerRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddThirdPartyNotificationManagerRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddThirdPartyNotificationManagerRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddThirdPartyNotificationManagerRequest(val *AddThirdPartyNotificationManagerRequest) *NullableAddThirdPartyNotificationManagerRequest {
	return &NullableAddThirdPartyNotificationManagerRequest{value: val, isSet: true}
}

func (v NullableAddThirdPartyNotificationManagerRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddThirdPartyNotificationManagerRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


