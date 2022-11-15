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

// SyslogJsonAuditLogPublisherResponseAllOf struct for SyslogJsonAuditLogPublisherResponseAllOf
type SyslogJsonAuditLogPublisherResponseAllOf struct {
	// Name of the Log Publisher
	Id string `json:"id"`
}

// NewSyslogJsonAuditLogPublisherResponseAllOf instantiates a new SyslogJsonAuditLogPublisherResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSyslogJsonAuditLogPublisherResponseAllOf(id string) *SyslogJsonAuditLogPublisherResponseAllOf {
	this := SyslogJsonAuditLogPublisherResponseAllOf{}
	this.Id = id
	return &this
}

// NewSyslogJsonAuditLogPublisherResponseAllOfWithDefaults instantiates a new SyslogJsonAuditLogPublisherResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSyslogJsonAuditLogPublisherResponseAllOfWithDefaults() *SyslogJsonAuditLogPublisherResponseAllOf {
	this := SyslogJsonAuditLogPublisherResponseAllOf{}
	return &this
}

// GetId returns the Id field value
func (o *SyslogJsonAuditLogPublisherResponseAllOf) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SyslogJsonAuditLogPublisherResponseAllOf) GetIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SyslogJsonAuditLogPublisherResponseAllOf) SetId(v string) {
	o.Id = v
}

func (o SyslogJsonAuditLogPublisherResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableSyslogJsonAuditLogPublisherResponseAllOf struct {
	value *SyslogJsonAuditLogPublisherResponseAllOf
	isSet bool
}

func (v NullableSyslogJsonAuditLogPublisherResponseAllOf) Get() *SyslogJsonAuditLogPublisherResponseAllOf {
	return v.value
}

func (v *NullableSyslogJsonAuditLogPublisherResponseAllOf) Set(val *SyslogJsonAuditLogPublisherResponseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSyslogJsonAuditLogPublisherResponseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSyslogJsonAuditLogPublisherResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyslogJsonAuditLogPublisherResponseAllOf(val *SyslogJsonAuditLogPublisherResponseAllOf) *NullableSyslogJsonAuditLogPublisherResponseAllOf {
	return &NullableSyslogJsonAuditLogPublisherResponseAllOf{value: val, isSet: true}
}

func (v NullableSyslogJsonAuditLogPublisherResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyslogJsonAuditLogPublisherResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


