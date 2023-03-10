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

// checks if the MetaUrnPingidentitySchemasConfigurationMessages20 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MetaUrnPingidentitySchemasConfigurationMessages20{}

// MetaUrnPingidentitySchemasConfigurationMessages20 struct for MetaUrnPingidentitySchemasConfigurationMessages20
type MetaUrnPingidentitySchemasConfigurationMessages20 struct {
	Notifications   []string                                                                `json:"notifications,omitempty"`
	RequiredActions []MetaUrnPingidentitySchemasConfigurationMessages20RequiredActionsInner `json:"requiredActions,omitempty"`
}

// NewMetaUrnPingidentitySchemasConfigurationMessages20 instantiates a new MetaUrnPingidentitySchemasConfigurationMessages20 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetaUrnPingidentitySchemasConfigurationMessages20() *MetaUrnPingidentitySchemasConfigurationMessages20 {
	this := MetaUrnPingidentitySchemasConfigurationMessages20{}
	return &this
}

// NewMetaUrnPingidentitySchemasConfigurationMessages20WithDefaults instantiates a new MetaUrnPingidentitySchemasConfigurationMessages20 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetaUrnPingidentitySchemasConfigurationMessages20WithDefaults() *MetaUrnPingidentitySchemasConfigurationMessages20 {
	this := MetaUrnPingidentitySchemasConfigurationMessages20{}
	return &this
}

// GetNotifications returns the Notifications field value if set, zero value otherwise.
func (o *MetaUrnPingidentitySchemasConfigurationMessages20) GetNotifications() []string {
	if o == nil || IsNil(o.Notifications) {
		var ret []string
		return ret
	}
	return o.Notifications
}

// GetNotificationsOk returns a tuple with the Notifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaUrnPingidentitySchemasConfigurationMessages20) GetNotificationsOk() ([]string, bool) {
	if o == nil || IsNil(o.Notifications) {
		return nil, false
	}
	return o.Notifications, true
}

// HasNotifications returns a boolean if a field has been set.
func (o *MetaUrnPingidentitySchemasConfigurationMessages20) HasNotifications() bool {
	if o != nil && !IsNil(o.Notifications) {
		return true
	}

	return false
}

// SetNotifications gets a reference to the given []string and assigns it to the Notifications field.
func (o *MetaUrnPingidentitySchemasConfigurationMessages20) SetNotifications(v []string) {
	o.Notifications = v
}

// GetRequiredActions returns the RequiredActions field value if set, zero value otherwise.
func (o *MetaUrnPingidentitySchemasConfigurationMessages20) GetRequiredActions() []MetaUrnPingidentitySchemasConfigurationMessages20RequiredActionsInner {
	if o == nil || IsNil(o.RequiredActions) {
		var ret []MetaUrnPingidentitySchemasConfigurationMessages20RequiredActionsInner
		return ret
	}
	return o.RequiredActions
}

// GetRequiredActionsOk returns a tuple with the RequiredActions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaUrnPingidentitySchemasConfigurationMessages20) GetRequiredActionsOk() ([]MetaUrnPingidentitySchemasConfigurationMessages20RequiredActionsInner, bool) {
	if o == nil || IsNil(o.RequiredActions) {
		return nil, false
	}
	return o.RequiredActions, true
}

// HasRequiredActions returns a boolean if a field has been set.
func (o *MetaUrnPingidentitySchemasConfigurationMessages20) HasRequiredActions() bool {
	if o != nil && !IsNil(o.RequiredActions) {
		return true
	}

	return false
}

// SetRequiredActions gets a reference to the given []MetaUrnPingidentitySchemasConfigurationMessages20RequiredActionsInner and assigns it to the RequiredActions field.
func (o *MetaUrnPingidentitySchemasConfigurationMessages20) SetRequiredActions(v []MetaUrnPingidentitySchemasConfigurationMessages20RequiredActionsInner) {
	o.RequiredActions = v
}

func (o MetaUrnPingidentitySchemasConfigurationMessages20) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MetaUrnPingidentitySchemasConfigurationMessages20) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Notifications) {
		toSerialize["notifications"] = o.Notifications
	}
	if !IsNil(o.RequiredActions) {
		toSerialize["requiredActions"] = o.RequiredActions
	}
	return toSerialize, nil
}

type NullableMetaUrnPingidentitySchemasConfigurationMessages20 struct {
	value *MetaUrnPingidentitySchemasConfigurationMessages20
	isSet bool
}

func (v NullableMetaUrnPingidentitySchemasConfigurationMessages20) Get() *MetaUrnPingidentitySchemasConfigurationMessages20 {
	return v.value
}

func (v *NullableMetaUrnPingidentitySchemasConfigurationMessages20) Set(val *MetaUrnPingidentitySchemasConfigurationMessages20) {
	v.value = val
	v.isSet = true
}

func (v NullableMetaUrnPingidentitySchemasConfigurationMessages20) IsSet() bool {
	return v.isSet
}

func (v *NullableMetaUrnPingidentitySchemasConfigurationMessages20) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetaUrnPingidentitySchemasConfigurationMessages20(val *MetaUrnPingidentitySchemasConfigurationMessages20) *NullableMetaUrnPingidentitySchemasConfigurationMessages20 {
	return &NullableMetaUrnPingidentitySchemasConfigurationMessages20{value: val, isSet: true}
}

func (v NullableMetaUrnPingidentitySchemasConfigurationMessages20) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetaUrnPingidentitySchemasConfigurationMessages20) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
