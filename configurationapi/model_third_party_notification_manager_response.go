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

// checks if the ThirdPartyNotificationManagerResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ThirdPartyNotificationManagerResponse{}

// ThirdPartyNotificationManagerResponse struct for ThirdPartyNotificationManagerResponse
type ThirdPartyNotificationManagerResponse struct {
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
	// Name of the Notification Manager
	Id      string                                       `json:"id"`
	Schemas []EnumthirdPartyNotificationManagerSchemaUrn `json:"schemas,omitempty"`
	// The fully-qualified name of the Java class providing the logic for the Third Party Notification Manager.
	ExtensionClass string `json:"extensionClass"`
	// The set of arguments used to customize the behavior for the Third Party Notification Manager. Each configuration property should be given in the form 'name=value'.
	ExtensionArgument []string `json:"extensionArgument,omitempty"`
	// A description for this Notification Manager
	Description *string `json:"description,omitempty"`
	// Indicates whether this Notification Manager is enabled within the server.
	Enabled bool `json:"enabled"`
	// Specifies the DN of the entry below which subscription data is stored for this Notification Manager. This needs to be in the backend that has the data to be notified on, and must not be the same entry as the backend base DN. The subscription base DN entry does not need to exist as it will be created by the server.
	SubscriptionBaseDN      string                                             `json:"subscriptionBaseDN"`
	TransactionNotification EnumnotificationManagerTransactionNotificationProp `json:"transactionNotification"`
	// Enables monitor entries for this Notification Manager.
	MonitorEntriesEnabled *bool `json:"monitorEntriesEnabled,omitempty"`
}

// NewThirdPartyNotificationManagerResponse instantiates a new ThirdPartyNotificationManagerResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThirdPartyNotificationManagerResponse(id string, extensionClass string, enabled bool, subscriptionBaseDN string, transactionNotification EnumnotificationManagerTransactionNotificationProp) *ThirdPartyNotificationManagerResponse {
	this := ThirdPartyNotificationManagerResponse{}
	this.Id = id
	this.ExtensionClass = extensionClass
	this.Enabled = enabled
	this.SubscriptionBaseDN = subscriptionBaseDN
	this.TransactionNotification = transactionNotification
	return &this
}

// NewThirdPartyNotificationManagerResponseWithDefaults instantiates a new ThirdPartyNotificationManagerResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThirdPartyNotificationManagerResponseWithDefaults() *ThirdPartyNotificationManagerResponse {
	this := ThirdPartyNotificationManagerResponse{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ThirdPartyNotificationManagerResponse) GetMeta() MetaMeta {
	if o == nil || IsNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyNotificationManagerResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ThirdPartyNotificationManagerResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *ThirdPartyNotificationManagerResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *ThirdPartyNotificationManagerResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyNotificationManagerResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *ThirdPartyNotificationManagerResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *ThirdPartyNotificationManagerResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetId returns the Id field value
func (o *ThirdPartyNotificationManagerResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ThirdPartyNotificationManagerResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ThirdPartyNotificationManagerResponse) SetId(v string) {
	o.Id = v
}

// GetSchemas returns the Schemas field value if set, zero value otherwise.
func (o *ThirdPartyNotificationManagerResponse) GetSchemas() []EnumthirdPartyNotificationManagerSchemaUrn {
	if o == nil || IsNil(o.Schemas) {
		var ret []EnumthirdPartyNotificationManagerSchemaUrn
		return ret
	}
	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyNotificationManagerResponse) GetSchemasOk() ([]EnumthirdPartyNotificationManagerSchemaUrn, bool) {
	if o == nil || IsNil(o.Schemas) {
		return nil, false
	}
	return o.Schemas, true
}

// HasSchemas returns a boolean if a field has been set.
func (o *ThirdPartyNotificationManagerResponse) HasSchemas() bool {
	if o != nil && !IsNil(o.Schemas) {
		return true
	}

	return false
}

// SetSchemas gets a reference to the given []EnumthirdPartyNotificationManagerSchemaUrn and assigns it to the Schemas field.
func (o *ThirdPartyNotificationManagerResponse) SetSchemas(v []EnumthirdPartyNotificationManagerSchemaUrn) {
	o.Schemas = v
}

// GetExtensionClass returns the ExtensionClass field value
func (o *ThirdPartyNotificationManagerResponse) GetExtensionClass() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExtensionClass
}

// GetExtensionClassOk returns a tuple with the ExtensionClass field value
// and a boolean to check if the value has been set.
func (o *ThirdPartyNotificationManagerResponse) GetExtensionClassOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExtensionClass, true
}

// SetExtensionClass sets field value
func (o *ThirdPartyNotificationManagerResponse) SetExtensionClass(v string) {
	o.ExtensionClass = v
}

// GetExtensionArgument returns the ExtensionArgument field value if set, zero value otherwise.
func (o *ThirdPartyNotificationManagerResponse) GetExtensionArgument() []string {
	if o == nil || IsNil(o.ExtensionArgument) {
		var ret []string
		return ret
	}
	return o.ExtensionArgument
}

// GetExtensionArgumentOk returns a tuple with the ExtensionArgument field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyNotificationManagerResponse) GetExtensionArgumentOk() ([]string, bool) {
	if o == nil || IsNil(o.ExtensionArgument) {
		return nil, false
	}
	return o.ExtensionArgument, true
}

// HasExtensionArgument returns a boolean if a field has been set.
func (o *ThirdPartyNotificationManagerResponse) HasExtensionArgument() bool {
	if o != nil && !IsNil(o.ExtensionArgument) {
		return true
	}

	return false
}

// SetExtensionArgument gets a reference to the given []string and assigns it to the ExtensionArgument field.
func (o *ThirdPartyNotificationManagerResponse) SetExtensionArgument(v []string) {
	o.ExtensionArgument = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ThirdPartyNotificationManagerResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyNotificationManagerResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ThirdPartyNotificationManagerResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ThirdPartyNotificationManagerResponse) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *ThirdPartyNotificationManagerResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *ThirdPartyNotificationManagerResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *ThirdPartyNotificationManagerResponse) SetEnabled(v bool) {
	o.Enabled = v
}

// GetSubscriptionBaseDN returns the SubscriptionBaseDN field value
func (o *ThirdPartyNotificationManagerResponse) GetSubscriptionBaseDN() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubscriptionBaseDN
}

// GetSubscriptionBaseDNOk returns a tuple with the SubscriptionBaseDN field value
// and a boolean to check if the value has been set.
func (o *ThirdPartyNotificationManagerResponse) GetSubscriptionBaseDNOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubscriptionBaseDN, true
}

// SetSubscriptionBaseDN sets field value
func (o *ThirdPartyNotificationManagerResponse) SetSubscriptionBaseDN(v string) {
	o.SubscriptionBaseDN = v
}

// GetTransactionNotification returns the TransactionNotification field value
func (o *ThirdPartyNotificationManagerResponse) GetTransactionNotification() EnumnotificationManagerTransactionNotificationProp {
	if o == nil {
		var ret EnumnotificationManagerTransactionNotificationProp
		return ret
	}

	return o.TransactionNotification
}

// GetTransactionNotificationOk returns a tuple with the TransactionNotification field value
// and a boolean to check if the value has been set.
func (o *ThirdPartyNotificationManagerResponse) GetTransactionNotificationOk() (*EnumnotificationManagerTransactionNotificationProp, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransactionNotification, true
}

// SetTransactionNotification sets field value
func (o *ThirdPartyNotificationManagerResponse) SetTransactionNotification(v EnumnotificationManagerTransactionNotificationProp) {
	o.TransactionNotification = v
}

// GetMonitorEntriesEnabled returns the MonitorEntriesEnabled field value if set, zero value otherwise.
func (o *ThirdPartyNotificationManagerResponse) GetMonitorEntriesEnabled() bool {
	if o == nil || IsNil(o.MonitorEntriesEnabled) {
		var ret bool
		return ret
	}
	return *o.MonitorEntriesEnabled
}

// GetMonitorEntriesEnabledOk returns a tuple with the MonitorEntriesEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyNotificationManagerResponse) GetMonitorEntriesEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.MonitorEntriesEnabled) {
		return nil, false
	}
	return o.MonitorEntriesEnabled, true
}

// HasMonitorEntriesEnabled returns a boolean if a field has been set.
func (o *ThirdPartyNotificationManagerResponse) HasMonitorEntriesEnabled() bool {
	if o != nil && !IsNil(o.MonitorEntriesEnabled) {
		return true
	}

	return false
}

// SetMonitorEntriesEnabled gets a reference to the given bool and assigns it to the MonitorEntriesEnabled field.
func (o *ThirdPartyNotificationManagerResponse) SetMonitorEntriesEnabled(v bool) {
	o.MonitorEntriesEnabled = &v
}

func (o ThirdPartyNotificationManagerResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ThirdPartyNotificationManagerResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		toSerialize["urn:pingidentity:schemas:configuration:messages:2.0"] = o.Urnpingidentityschemasconfigurationmessages20
	}
	toSerialize["id"] = o.Id
	if !IsNil(o.Schemas) {
		toSerialize["schemas"] = o.Schemas
	}
	toSerialize["extensionClass"] = o.ExtensionClass
	if !IsNil(o.ExtensionArgument) {
		toSerialize["extensionArgument"] = o.ExtensionArgument
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["enabled"] = o.Enabled
	toSerialize["subscriptionBaseDN"] = o.SubscriptionBaseDN
	toSerialize["transactionNotification"] = o.TransactionNotification
	if !IsNil(o.MonitorEntriesEnabled) {
		toSerialize["monitorEntriesEnabled"] = o.MonitorEntriesEnabled
	}
	return toSerialize, nil
}

type NullableThirdPartyNotificationManagerResponse struct {
	value *ThirdPartyNotificationManagerResponse
	isSet bool
}

func (v NullableThirdPartyNotificationManagerResponse) Get() *ThirdPartyNotificationManagerResponse {
	return v.value
}

func (v *NullableThirdPartyNotificationManagerResponse) Set(val *ThirdPartyNotificationManagerResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableThirdPartyNotificationManagerResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableThirdPartyNotificationManagerResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThirdPartyNotificationManagerResponse(val *ThirdPartyNotificationManagerResponse) *NullableThirdPartyNotificationManagerResponse {
	return &NullableThirdPartyNotificationManagerResponse{value: val, isSet: true}
}

func (v NullableThirdPartyNotificationManagerResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThirdPartyNotificationManagerResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
