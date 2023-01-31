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

// LastAccessTimeUncachedEntryCriteriaResponse struct for LastAccessTimeUncachedEntryCriteriaResponse
type LastAccessTimeUncachedEntryCriteriaResponse struct {
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
	// Name of the Uncached Entry Criteria
	Id      string                                             `json:"id"`
	Schemas []EnumlastAccessTimeUncachedEntryCriteriaSchemaUrn `json:"schemas"`
	// Specifies the maximum length of time that has passed since an entry was last accessed that it should still be included in the id2entry database. Entries that have not been accessed in more than this length of time may be written into the uncached-id2entry database.
	AccessTimeThreshold string `json:"accessTimeThreshold"`
	// A description for this Uncached Entry Criteria
	Description *string `json:"description,omitempty"`
	// Indicates whether this Uncached Entry Criteria is enabled for use in the server.
	Enabled bool `json:"enabled"`
}

// NewLastAccessTimeUncachedEntryCriteriaResponse instantiates a new LastAccessTimeUncachedEntryCriteriaResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLastAccessTimeUncachedEntryCriteriaResponse(id string, schemas []EnumlastAccessTimeUncachedEntryCriteriaSchemaUrn, accessTimeThreshold string, enabled bool) *LastAccessTimeUncachedEntryCriteriaResponse {
	this := LastAccessTimeUncachedEntryCriteriaResponse{}
	this.Id = id
	this.Schemas = schemas
	this.AccessTimeThreshold = accessTimeThreshold
	this.Enabled = enabled
	return &this
}

// NewLastAccessTimeUncachedEntryCriteriaResponseWithDefaults instantiates a new LastAccessTimeUncachedEntryCriteriaResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLastAccessTimeUncachedEntryCriteriaResponseWithDefaults() *LastAccessTimeUncachedEntryCriteriaResponse {
	this := LastAccessTimeUncachedEntryCriteriaResponse{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *LastAccessTimeUncachedEntryCriteriaResponse) GetMeta() MetaMeta {
	if o == nil || isNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LastAccessTimeUncachedEntryCriteriaResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || isNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *LastAccessTimeUncachedEntryCriteriaResponse) HasMeta() bool {
	if o != nil && !isNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *LastAccessTimeUncachedEntryCriteriaResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *LastAccessTimeUncachedEntryCriteriaResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LastAccessTimeUncachedEntryCriteriaResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *LastAccessTimeUncachedEntryCriteriaResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *LastAccessTimeUncachedEntryCriteriaResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetId returns the Id field value
func (o *LastAccessTimeUncachedEntryCriteriaResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *LastAccessTimeUncachedEntryCriteriaResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *LastAccessTimeUncachedEntryCriteriaResponse) SetId(v string) {
	o.Id = v
}

// GetSchemas returns the Schemas field value
func (o *LastAccessTimeUncachedEntryCriteriaResponse) GetSchemas() []EnumlastAccessTimeUncachedEntryCriteriaSchemaUrn {
	if o == nil {
		var ret []EnumlastAccessTimeUncachedEntryCriteriaSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *LastAccessTimeUncachedEntryCriteriaResponse) GetSchemasOk() ([]EnumlastAccessTimeUncachedEntryCriteriaSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *LastAccessTimeUncachedEntryCriteriaResponse) SetSchemas(v []EnumlastAccessTimeUncachedEntryCriteriaSchemaUrn) {
	o.Schemas = v
}

// GetAccessTimeThreshold returns the AccessTimeThreshold field value
func (o *LastAccessTimeUncachedEntryCriteriaResponse) GetAccessTimeThreshold() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccessTimeThreshold
}

// GetAccessTimeThresholdOk returns a tuple with the AccessTimeThreshold field value
// and a boolean to check if the value has been set.
func (o *LastAccessTimeUncachedEntryCriteriaResponse) GetAccessTimeThresholdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessTimeThreshold, true
}

// SetAccessTimeThreshold sets field value
func (o *LastAccessTimeUncachedEntryCriteriaResponse) SetAccessTimeThreshold(v string) {
	o.AccessTimeThreshold = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *LastAccessTimeUncachedEntryCriteriaResponse) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LastAccessTimeUncachedEntryCriteriaResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *LastAccessTimeUncachedEntryCriteriaResponse) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *LastAccessTimeUncachedEntryCriteriaResponse) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *LastAccessTimeUncachedEntryCriteriaResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *LastAccessTimeUncachedEntryCriteriaResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *LastAccessTimeUncachedEntryCriteriaResponse) SetEnabled(v bool) {
	o.Enabled = v
}

func (o LastAccessTimeUncachedEntryCriteriaResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		toSerialize["urn:pingidentity:schemas:configuration:messages:2.0"] = o.Urnpingidentityschemasconfigurationmessages20
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["schemas"] = o.Schemas
	}
	if true {
		toSerialize["accessTimeThreshold"] = o.AccessTimeThreshold
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableLastAccessTimeUncachedEntryCriteriaResponse struct {
	value *LastAccessTimeUncachedEntryCriteriaResponse
	isSet bool
}

func (v NullableLastAccessTimeUncachedEntryCriteriaResponse) Get() *LastAccessTimeUncachedEntryCriteriaResponse {
	return v.value
}

func (v *NullableLastAccessTimeUncachedEntryCriteriaResponse) Set(val *LastAccessTimeUncachedEntryCriteriaResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableLastAccessTimeUncachedEntryCriteriaResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableLastAccessTimeUncachedEntryCriteriaResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLastAccessTimeUncachedEntryCriteriaResponse(val *LastAccessTimeUncachedEntryCriteriaResponse) *NullableLastAccessTimeUncachedEntryCriteriaResponse {
	return &NullableLastAccessTimeUncachedEntryCriteriaResponse{value: val, isSet: true}
}

func (v NullableLastAccessTimeUncachedEntryCriteriaResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLastAccessTimeUncachedEntryCriteriaResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
