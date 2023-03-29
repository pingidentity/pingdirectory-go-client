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

// checks if the ReplicationSynchronizationProviderResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReplicationSynchronizationProviderResponse{}

// ReplicationSynchronizationProviderResponse struct for ReplicationSynchronizationProviderResponse
type ReplicationSynchronizationProviderResponse struct {
	Schemas []EnumreplicationSynchronizationProviderSchemaUrn `json:"schemas"`
	// Name of the Synchronization Provider
	Id string `json:"id"`
	// Specifies the number of update replay threads.
	NumUpdateReplayThreads *int32 `json:"numUpdateReplayThreads,omitempty"`
	// A description for this Synchronization Provider
	Description *string `json:"description,omitempty"`
	// Indicates whether the Synchronization Provider is enabled for use.
	Enabled                                       bool                                               `json:"enabled"`
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
}

// NewReplicationSynchronizationProviderResponse instantiates a new ReplicationSynchronizationProviderResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReplicationSynchronizationProviderResponse(schemas []EnumreplicationSynchronizationProviderSchemaUrn, id string, enabled bool) *ReplicationSynchronizationProviderResponse {
	this := ReplicationSynchronizationProviderResponse{}
	this.Schemas = schemas
	this.Id = id
	this.Enabled = enabled
	return &this
}

// NewReplicationSynchronizationProviderResponseWithDefaults instantiates a new ReplicationSynchronizationProviderResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReplicationSynchronizationProviderResponseWithDefaults() *ReplicationSynchronizationProviderResponse {
	this := ReplicationSynchronizationProviderResponse{}
	return &this
}

// GetSchemas returns the Schemas field value
func (o *ReplicationSynchronizationProviderResponse) GetSchemas() []EnumreplicationSynchronizationProviderSchemaUrn {
	if o == nil {
		var ret []EnumreplicationSynchronizationProviderSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *ReplicationSynchronizationProviderResponse) GetSchemasOk() ([]EnumreplicationSynchronizationProviderSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *ReplicationSynchronizationProviderResponse) SetSchemas(v []EnumreplicationSynchronizationProviderSchemaUrn) {
	o.Schemas = v
}

// GetId returns the Id field value
func (o *ReplicationSynchronizationProviderResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ReplicationSynchronizationProviderResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ReplicationSynchronizationProviderResponse) SetId(v string) {
	o.Id = v
}

// GetNumUpdateReplayThreads returns the NumUpdateReplayThreads field value if set, zero value otherwise.
func (o *ReplicationSynchronizationProviderResponse) GetNumUpdateReplayThreads() int32 {
	if o == nil || IsNil(o.NumUpdateReplayThreads) {
		var ret int32
		return ret
	}
	return *o.NumUpdateReplayThreads
}

// GetNumUpdateReplayThreadsOk returns a tuple with the NumUpdateReplayThreads field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplicationSynchronizationProviderResponse) GetNumUpdateReplayThreadsOk() (*int32, bool) {
	if o == nil || IsNil(o.NumUpdateReplayThreads) {
		return nil, false
	}
	return o.NumUpdateReplayThreads, true
}

// HasNumUpdateReplayThreads returns a boolean if a field has been set.
func (o *ReplicationSynchronizationProviderResponse) HasNumUpdateReplayThreads() bool {
	if o != nil && !IsNil(o.NumUpdateReplayThreads) {
		return true
	}

	return false
}

// SetNumUpdateReplayThreads gets a reference to the given int32 and assigns it to the NumUpdateReplayThreads field.
func (o *ReplicationSynchronizationProviderResponse) SetNumUpdateReplayThreads(v int32) {
	o.NumUpdateReplayThreads = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ReplicationSynchronizationProviderResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplicationSynchronizationProviderResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ReplicationSynchronizationProviderResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ReplicationSynchronizationProviderResponse) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *ReplicationSynchronizationProviderResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *ReplicationSynchronizationProviderResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *ReplicationSynchronizationProviderResponse) SetEnabled(v bool) {
	o.Enabled = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ReplicationSynchronizationProviderResponse) GetMeta() MetaMeta {
	if o == nil || IsNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplicationSynchronizationProviderResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ReplicationSynchronizationProviderResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *ReplicationSynchronizationProviderResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *ReplicationSynchronizationProviderResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplicationSynchronizationProviderResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *ReplicationSynchronizationProviderResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *ReplicationSynchronizationProviderResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

func (o ReplicationSynchronizationProviderResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReplicationSynchronizationProviderResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["schemas"] = o.Schemas
	toSerialize["id"] = o.Id
	if !IsNil(o.NumUpdateReplayThreads) {
		toSerialize["numUpdateReplayThreads"] = o.NumUpdateReplayThreads
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["enabled"] = o.Enabled
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		toSerialize["urn:pingidentity:schemas:configuration:messages:2.0"] = o.Urnpingidentityschemasconfigurationmessages20
	}
	return toSerialize, nil
}

type NullableReplicationSynchronizationProviderResponse struct {
	value *ReplicationSynchronizationProviderResponse
	isSet bool
}

func (v NullableReplicationSynchronizationProviderResponse) Get() *ReplicationSynchronizationProviderResponse {
	return v.value
}

func (v *NullableReplicationSynchronizationProviderResponse) Set(val *ReplicationSynchronizationProviderResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableReplicationSynchronizationProviderResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableReplicationSynchronizationProviderResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReplicationSynchronizationProviderResponse(val *ReplicationSynchronizationProviderResponse) *NullableReplicationSynchronizationProviderResponse {
	return &NullableReplicationSynchronizationProviderResponse{value: val, isSet: true}
}

func (v NullableReplicationSynchronizationProviderResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReplicationSynchronizationProviderResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
