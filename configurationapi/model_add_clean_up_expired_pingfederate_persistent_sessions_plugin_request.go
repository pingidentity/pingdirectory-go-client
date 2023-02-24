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

// checks if the AddCleanUpExpiredPingfederatePersistentSessionsPluginRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddCleanUpExpiredPingfederatePersistentSessionsPluginRequest{}

// AddCleanUpExpiredPingfederatePersistentSessionsPluginRequest struct for AddCleanUpExpiredPingfederatePersistentSessionsPluginRequest
type AddCleanUpExpiredPingfederatePersistentSessionsPluginRequest struct {
	// Name of the new Plugin
	PluginName string                                                            `json:"pluginName"`
	Schemas    []EnumcleanUpExpiredPingfederatePersistentSessionsPluginSchemaUrn `json:"schemas"`
	// This specifies how often the plugin should check for expired data. It also controls the offset of peer servers (see the peer-server-priority-index for more information).
	PollingInterval *string `json:"pollingInterval,omitempty"`
	// In a replicated environment, this determines the order in which peer servers should attempt to purge data.
	PeerServerPriorityIndex *int32 `json:"peerServerPriorityIndex,omitempty"`
	// Only entries located within the subtree specified by this base DN are eligible for purging.
	BaseDN *string `json:"baseDN,omitempty"`
	// This setting smooths out the performance impact on the server by throttling the purging to the specified maximum number of updates per second. To avoid a large backlog, this value should be set comfortably above the average rate that expired data is generated. When purge-behavior is set to subtree-delete-entries, then deletion of the entire subtree is considered a single update for the purposes of throttling.
	MaxUpdatesPerSecond *int32 `json:"maxUpdatesPerSecond,omitempty"`
	// The number of threads used to delete expired entries.
	NumDeleteThreads *int32 `json:"numDeleteThreads,omitempty"`
	// Indicates whether the plug-in is enabled for use.
	Enabled bool `json:"enabled"`
}

// NewAddCleanUpExpiredPingfederatePersistentSessionsPluginRequest instantiates a new AddCleanUpExpiredPingfederatePersistentSessionsPluginRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddCleanUpExpiredPingfederatePersistentSessionsPluginRequest(pluginName string, schemas []EnumcleanUpExpiredPingfederatePersistentSessionsPluginSchemaUrn, enabled bool) *AddCleanUpExpiredPingfederatePersistentSessionsPluginRequest {
	this := AddCleanUpExpiredPingfederatePersistentSessionsPluginRequest{}
	this.PluginName = pluginName
	this.Schemas = schemas
	this.Enabled = enabled
	return &this
}

// NewAddCleanUpExpiredPingfederatePersistentSessionsPluginRequestWithDefaults instantiates a new AddCleanUpExpiredPingfederatePersistentSessionsPluginRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddCleanUpExpiredPingfederatePersistentSessionsPluginRequestWithDefaults() *AddCleanUpExpiredPingfederatePersistentSessionsPluginRequest {
	this := AddCleanUpExpiredPingfederatePersistentSessionsPluginRequest{}
	return &this
}

// GetPluginName returns the PluginName field value
func (o *AddCleanUpExpiredPingfederatePersistentSessionsPluginRequest) GetPluginName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PluginName
}

// GetPluginNameOk returns a tuple with the PluginName field value
// and a boolean to check if the value has been set.
func (o *AddCleanUpExpiredPingfederatePersistentSessionsPluginRequest) GetPluginNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PluginName, true
}

// SetPluginName sets field value
func (o *AddCleanUpExpiredPingfederatePersistentSessionsPluginRequest) SetPluginName(v string) {
	o.PluginName = v
}

// GetSchemas returns the Schemas field value
func (o *AddCleanUpExpiredPingfederatePersistentSessionsPluginRequest) GetSchemas() []EnumcleanUpExpiredPingfederatePersistentSessionsPluginSchemaUrn {
	if o == nil {
		var ret []EnumcleanUpExpiredPingfederatePersistentSessionsPluginSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddCleanUpExpiredPingfederatePersistentSessionsPluginRequest) GetSchemasOk() ([]EnumcleanUpExpiredPingfederatePersistentSessionsPluginSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddCleanUpExpiredPingfederatePersistentSessionsPluginRequest) SetSchemas(v []EnumcleanUpExpiredPingfederatePersistentSessionsPluginSchemaUrn) {
	o.Schemas = v
}

// GetPollingInterval returns the PollingInterval field value if set, zero value otherwise.
func (o *AddCleanUpExpiredPingfederatePersistentSessionsPluginRequest) GetPollingInterval() string {
	if o == nil || IsNil(o.PollingInterval) {
		var ret string
		return ret
	}
	return *o.PollingInterval
}

// GetPollingIntervalOk returns a tuple with the PollingInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddCleanUpExpiredPingfederatePersistentSessionsPluginRequest) GetPollingIntervalOk() (*string, bool) {
	if o == nil || IsNil(o.PollingInterval) {
		return nil, false
	}
	return o.PollingInterval, true
}

// HasPollingInterval returns a boolean if a field has been set.
func (o *AddCleanUpExpiredPingfederatePersistentSessionsPluginRequest) HasPollingInterval() bool {
	if o != nil && !IsNil(o.PollingInterval) {
		return true
	}

	return false
}

// SetPollingInterval gets a reference to the given string and assigns it to the PollingInterval field.
func (o *AddCleanUpExpiredPingfederatePersistentSessionsPluginRequest) SetPollingInterval(v string) {
	o.PollingInterval = &v
}

// GetPeerServerPriorityIndex returns the PeerServerPriorityIndex field value if set, zero value otherwise.
func (o *AddCleanUpExpiredPingfederatePersistentSessionsPluginRequest) GetPeerServerPriorityIndex() int32 {
	if o == nil || IsNil(o.PeerServerPriorityIndex) {
		var ret int32
		return ret
	}
	return *o.PeerServerPriorityIndex
}

// GetPeerServerPriorityIndexOk returns a tuple with the PeerServerPriorityIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddCleanUpExpiredPingfederatePersistentSessionsPluginRequest) GetPeerServerPriorityIndexOk() (*int32, bool) {
	if o == nil || IsNil(o.PeerServerPriorityIndex) {
		return nil, false
	}
	return o.PeerServerPriorityIndex, true
}

// HasPeerServerPriorityIndex returns a boolean if a field has been set.
func (o *AddCleanUpExpiredPingfederatePersistentSessionsPluginRequest) HasPeerServerPriorityIndex() bool {
	if o != nil && !IsNil(o.PeerServerPriorityIndex) {
		return true
	}

	return false
}

// SetPeerServerPriorityIndex gets a reference to the given int32 and assigns it to the PeerServerPriorityIndex field.
func (o *AddCleanUpExpiredPingfederatePersistentSessionsPluginRequest) SetPeerServerPriorityIndex(v int32) {
	o.PeerServerPriorityIndex = &v
}

// GetBaseDN returns the BaseDN field value if set, zero value otherwise.
func (o *AddCleanUpExpiredPingfederatePersistentSessionsPluginRequest) GetBaseDN() string {
	if o == nil || IsNil(o.BaseDN) {
		var ret string
		return ret
	}
	return *o.BaseDN
}

// GetBaseDNOk returns a tuple with the BaseDN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddCleanUpExpiredPingfederatePersistentSessionsPluginRequest) GetBaseDNOk() (*string, bool) {
	if o == nil || IsNil(o.BaseDN) {
		return nil, false
	}
	return o.BaseDN, true
}

// HasBaseDN returns a boolean if a field has been set.
func (o *AddCleanUpExpiredPingfederatePersistentSessionsPluginRequest) HasBaseDN() bool {
	if o != nil && !IsNil(o.BaseDN) {
		return true
	}

	return false
}

// SetBaseDN gets a reference to the given string and assigns it to the BaseDN field.
func (o *AddCleanUpExpiredPingfederatePersistentSessionsPluginRequest) SetBaseDN(v string) {
	o.BaseDN = &v
}

// GetMaxUpdatesPerSecond returns the MaxUpdatesPerSecond field value if set, zero value otherwise.
func (o *AddCleanUpExpiredPingfederatePersistentSessionsPluginRequest) GetMaxUpdatesPerSecond() int32 {
	if o == nil || IsNil(o.MaxUpdatesPerSecond) {
		var ret int32
		return ret
	}
	return *o.MaxUpdatesPerSecond
}

// GetMaxUpdatesPerSecondOk returns a tuple with the MaxUpdatesPerSecond field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddCleanUpExpiredPingfederatePersistentSessionsPluginRequest) GetMaxUpdatesPerSecondOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxUpdatesPerSecond) {
		return nil, false
	}
	return o.MaxUpdatesPerSecond, true
}

// HasMaxUpdatesPerSecond returns a boolean if a field has been set.
func (o *AddCleanUpExpiredPingfederatePersistentSessionsPluginRequest) HasMaxUpdatesPerSecond() bool {
	if o != nil && !IsNil(o.MaxUpdatesPerSecond) {
		return true
	}

	return false
}

// SetMaxUpdatesPerSecond gets a reference to the given int32 and assigns it to the MaxUpdatesPerSecond field.
func (o *AddCleanUpExpiredPingfederatePersistentSessionsPluginRequest) SetMaxUpdatesPerSecond(v int32) {
	o.MaxUpdatesPerSecond = &v
}

// GetNumDeleteThreads returns the NumDeleteThreads field value if set, zero value otherwise.
func (o *AddCleanUpExpiredPingfederatePersistentSessionsPluginRequest) GetNumDeleteThreads() int32 {
	if o == nil || IsNil(o.NumDeleteThreads) {
		var ret int32
		return ret
	}
	return *o.NumDeleteThreads
}

// GetNumDeleteThreadsOk returns a tuple with the NumDeleteThreads field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddCleanUpExpiredPingfederatePersistentSessionsPluginRequest) GetNumDeleteThreadsOk() (*int32, bool) {
	if o == nil || IsNil(o.NumDeleteThreads) {
		return nil, false
	}
	return o.NumDeleteThreads, true
}

// HasNumDeleteThreads returns a boolean if a field has been set.
func (o *AddCleanUpExpiredPingfederatePersistentSessionsPluginRequest) HasNumDeleteThreads() bool {
	if o != nil && !IsNil(o.NumDeleteThreads) {
		return true
	}

	return false
}

// SetNumDeleteThreads gets a reference to the given int32 and assigns it to the NumDeleteThreads field.
func (o *AddCleanUpExpiredPingfederatePersistentSessionsPluginRequest) SetNumDeleteThreads(v int32) {
	o.NumDeleteThreads = &v
}

// GetEnabled returns the Enabled field value
func (o *AddCleanUpExpiredPingfederatePersistentSessionsPluginRequest) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AddCleanUpExpiredPingfederatePersistentSessionsPluginRequest) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AddCleanUpExpiredPingfederatePersistentSessionsPluginRequest) SetEnabled(v bool) {
	o.Enabled = v
}

func (o AddCleanUpExpiredPingfederatePersistentSessionsPluginRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddCleanUpExpiredPingfederatePersistentSessionsPluginRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pluginName"] = o.PluginName
	toSerialize["schemas"] = o.Schemas
	if !IsNil(o.PollingInterval) {
		toSerialize["pollingInterval"] = o.PollingInterval
	}
	if !IsNil(o.PeerServerPriorityIndex) {
		toSerialize["peerServerPriorityIndex"] = o.PeerServerPriorityIndex
	}
	if !IsNil(o.BaseDN) {
		toSerialize["baseDN"] = o.BaseDN
	}
	if !IsNil(o.MaxUpdatesPerSecond) {
		toSerialize["maxUpdatesPerSecond"] = o.MaxUpdatesPerSecond
	}
	if !IsNil(o.NumDeleteThreads) {
		toSerialize["numDeleteThreads"] = o.NumDeleteThreads
	}
	toSerialize["enabled"] = o.Enabled
	return toSerialize, nil
}

type NullableAddCleanUpExpiredPingfederatePersistentSessionsPluginRequest struct {
	value *AddCleanUpExpiredPingfederatePersistentSessionsPluginRequest
	isSet bool
}

func (v NullableAddCleanUpExpiredPingfederatePersistentSessionsPluginRequest) Get() *AddCleanUpExpiredPingfederatePersistentSessionsPluginRequest {
	return v.value
}

func (v *NullableAddCleanUpExpiredPingfederatePersistentSessionsPluginRequest) Set(val *AddCleanUpExpiredPingfederatePersistentSessionsPluginRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddCleanUpExpiredPingfederatePersistentSessionsPluginRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddCleanUpExpiredPingfederatePersistentSessionsPluginRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddCleanUpExpiredPingfederatePersistentSessionsPluginRequest(val *AddCleanUpExpiredPingfederatePersistentSessionsPluginRequest) *NullableAddCleanUpExpiredPingfederatePersistentSessionsPluginRequest {
	return &NullableAddCleanUpExpiredPingfederatePersistentSessionsPluginRequest{value: val, isSet: true}
}

func (v NullableAddCleanUpExpiredPingfederatePersistentSessionsPluginRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddCleanUpExpiredPingfederatePersistentSessionsPluginRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
