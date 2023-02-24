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

// AddCleanUpInactivePingfederatePersistentSessionsPluginRequest struct for AddCleanUpInactivePingfederatePersistentSessionsPluginRequest
type AddCleanUpInactivePingfederatePersistentSessionsPluginRequest struct {
	// Name of the new Plugin
	PluginName string                                                             `json:"pluginName"`
	Schemas    []EnumcleanUpInactivePingfederatePersistentSessionsPluginSchemaUrn `json:"schemas"`
	// Sessions whose last activity timestamp is older than this offset will be removed.
	ExpirationOffset string `json:"expirationOffset"`
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

// NewAddCleanUpInactivePingfederatePersistentSessionsPluginRequest instantiates a new AddCleanUpInactivePingfederatePersistentSessionsPluginRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddCleanUpInactivePingfederatePersistentSessionsPluginRequest(pluginName string, schemas []EnumcleanUpInactivePingfederatePersistentSessionsPluginSchemaUrn, expirationOffset string, enabled bool) *AddCleanUpInactivePingfederatePersistentSessionsPluginRequest {
	this := AddCleanUpInactivePingfederatePersistentSessionsPluginRequest{}
	this.PluginName = pluginName
	this.Schemas = schemas
	this.ExpirationOffset = expirationOffset
	this.Enabled = enabled
	return &this
}

// NewAddCleanUpInactivePingfederatePersistentSessionsPluginRequestWithDefaults instantiates a new AddCleanUpInactivePingfederatePersistentSessionsPluginRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddCleanUpInactivePingfederatePersistentSessionsPluginRequestWithDefaults() *AddCleanUpInactivePingfederatePersistentSessionsPluginRequest {
	this := AddCleanUpInactivePingfederatePersistentSessionsPluginRequest{}
	return &this
}

// GetPluginName returns the PluginName field value
func (o *AddCleanUpInactivePingfederatePersistentSessionsPluginRequest) GetPluginName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PluginName
}

// GetPluginNameOk returns a tuple with the PluginName field value
// and a boolean to check if the value has been set.
func (o *AddCleanUpInactivePingfederatePersistentSessionsPluginRequest) GetPluginNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PluginName, true
}

// SetPluginName sets field value
func (o *AddCleanUpInactivePingfederatePersistentSessionsPluginRequest) SetPluginName(v string) {
	o.PluginName = v
}

// GetSchemas returns the Schemas field value
func (o *AddCleanUpInactivePingfederatePersistentSessionsPluginRequest) GetSchemas() []EnumcleanUpInactivePingfederatePersistentSessionsPluginSchemaUrn {
	if o == nil {
		var ret []EnumcleanUpInactivePingfederatePersistentSessionsPluginSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddCleanUpInactivePingfederatePersistentSessionsPluginRequest) GetSchemasOk() ([]EnumcleanUpInactivePingfederatePersistentSessionsPluginSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddCleanUpInactivePingfederatePersistentSessionsPluginRequest) SetSchemas(v []EnumcleanUpInactivePingfederatePersistentSessionsPluginSchemaUrn) {
	o.Schemas = v
}

// GetExpirationOffset returns the ExpirationOffset field value
func (o *AddCleanUpInactivePingfederatePersistentSessionsPluginRequest) GetExpirationOffset() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExpirationOffset
}

// GetExpirationOffsetOk returns a tuple with the ExpirationOffset field value
// and a boolean to check if the value has been set.
func (o *AddCleanUpInactivePingfederatePersistentSessionsPluginRequest) GetExpirationOffsetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpirationOffset, true
}

// SetExpirationOffset sets field value
func (o *AddCleanUpInactivePingfederatePersistentSessionsPluginRequest) SetExpirationOffset(v string) {
	o.ExpirationOffset = v
}

// GetPollingInterval returns the PollingInterval field value if set, zero value otherwise.
func (o *AddCleanUpInactivePingfederatePersistentSessionsPluginRequest) GetPollingInterval() string {
	if o == nil || isNil(o.PollingInterval) {
		var ret string
		return ret
	}
	return *o.PollingInterval
}

// GetPollingIntervalOk returns a tuple with the PollingInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddCleanUpInactivePingfederatePersistentSessionsPluginRequest) GetPollingIntervalOk() (*string, bool) {
	if o == nil || isNil(o.PollingInterval) {
		return nil, false
	}
	return o.PollingInterval, true
}

// HasPollingInterval returns a boolean if a field has been set.
func (o *AddCleanUpInactivePingfederatePersistentSessionsPluginRequest) HasPollingInterval() bool {
	if o != nil && !isNil(o.PollingInterval) {
		return true
	}

	return false
}

// SetPollingInterval gets a reference to the given string and assigns it to the PollingInterval field.
func (o *AddCleanUpInactivePingfederatePersistentSessionsPluginRequest) SetPollingInterval(v string) {
	o.PollingInterval = &v
}

// GetPeerServerPriorityIndex returns the PeerServerPriorityIndex field value if set, zero value otherwise.
func (o *AddCleanUpInactivePingfederatePersistentSessionsPluginRequest) GetPeerServerPriorityIndex() int32 {
	if o == nil || isNil(o.PeerServerPriorityIndex) {
		var ret int32
		return ret
	}
	return *o.PeerServerPriorityIndex
}

// GetPeerServerPriorityIndexOk returns a tuple with the PeerServerPriorityIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddCleanUpInactivePingfederatePersistentSessionsPluginRequest) GetPeerServerPriorityIndexOk() (*int32, bool) {
	if o == nil || isNil(o.PeerServerPriorityIndex) {
		return nil, false
	}
	return o.PeerServerPriorityIndex, true
}

// HasPeerServerPriorityIndex returns a boolean if a field has been set.
func (o *AddCleanUpInactivePingfederatePersistentSessionsPluginRequest) HasPeerServerPriorityIndex() bool {
	if o != nil && !isNil(o.PeerServerPriorityIndex) {
		return true
	}

	return false
}

// SetPeerServerPriorityIndex gets a reference to the given int32 and assigns it to the PeerServerPriorityIndex field.
func (o *AddCleanUpInactivePingfederatePersistentSessionsPluginRequest) SetPeerServerPriorityIndex(v int32) {
	o.PeerServerPriorityIndex = &v
}

// GetBaseDN returns the BaseDN field value if set, zero value otherwise.
func (o *AddCleanUpInactivePingfederatePersistentSessionsPluginRequest) GetBaseDN() string {
	if o == nil || isNil(o.BaseDN) {
		var ret string
		return ret
	}
	return *o.BaseDN
}

// GetBaseDNOk returns a tuple with the BaseDN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddCleanUpInactivePingfederatePersistentSessionsPluginRequest) GetBaseDNOk() (*string, bool) {
	if o == nil || isNil(o.BaseDN) {
		return nil, false
	}
	return o.BaseDN, true
}

// HasBaseDN returns a boolean if a field has been set.
func (o *AddCleanUpInactivePingfederatePersistentSessionsPluginRequest) HasBaseDN() bool {
	if o != nil && !isNil(o.BaseDN) {
		return true
	}

	return false
}

// SetBaseDN gets a reference to the given string and assigns it to the BaseDN field.
func (o *AddCleanUpInactivePingfederatePersistentSessionsPluginRequest) SetBaseDN(v string) {
	o.BaseDN = &v
}

// GetMaxUpdatesPerSecond returns the MaxUpdatesPerSecond field value if set, zero value otherwise.
func (o *AddCleanUpInactivePingfederatePersistentSessionsPluginRequest) GetMaxUpdatesPerSecond() int32 {
	if o == nil || isNil(o.MaxUpdatesPerSecond) {
		var ret int32
		return ret
	}
	return *o.MaxUpdatesPerSecond
}

// GetMaxUpdatesPerSecondOk returns a tuple with the MaxUpdatesPerSecond field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddCleanUpInactivePingfederatePersistentSessionsPluginRequest) GetMaxUpdatesPerSecondOk() (*int32, bool) {
	if o == nil || isNil(o.MaxUpdatesPerSecond) {
		return nil, false
	}
	return o.MaxUpdatesPerSecond, true
}

// HasMaxUpdatesPerSecond returns a boolean if a field has been set.
func (o *AddCleanUpInactivePingfederatePersistentSessionsPluginRequest) HasMaxUpdatesPerSecond() bool {
	if o != nil && !isNil(o.MaxUpdatesPerSecond) {
		return true
	}

	return false
}

// SetMaxUpdatesPerSecond gets a reference to the given int32 and assigns it to the MaxUpdatesPerSecond field.
func (o *AddCleanUpInactivePingfederatePersistentSessionsPluginRequest) SetMaxUpdatesPerSecond(v int32) {
	o.MaxUpdatesPerSecond = &v
}

// GetNumDeleteThreads returns the NumDeleteThreads field value if set, zero value otherwise.
func (o *AddCleanUpInactivePingfederatePersistentSessionsPluginRequest) GetNumDeleteThreads() int32 {
	if o == nil || isNil(o.NumDeleteThreads) {
		var ret int32
		return ret
	}
	return *o.NumDeleteThreads
}

// GetNumDeleteThreadsOk returns a tuple with the NumDeleteThreads field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddCleanUpInactivePingfederatePersistentSessionsPluginRequest) GetNumDeleteThreadsOk() (*int32, bool) {
	if o == nil || isNil(o.NumDeleteThreads) {
		return nil, false
	}
	return o.NumDeleteThreads, true
}

// HasNumDeleteThreads returns a boolean if a field has been set.
func (o *AddCleanUpInactivePingfederatePersistentSessionsPluginRequest) HasNumDeleteThreads() bool {
	if o != nil && !isNil(o.NumDeleteThreads) {
		return true
	}

	return false
}

// SetNumDeleteThreads gets a reference to the given int32 and assigns it to the NumDeleteThreads field.
func (o *AddCleanUpInactivePingfederatePersistentSessionsPluginRequest) SetNumDeleteThreads(v int32) {
	o.NumDeleteThreads = &v
}

// GetEnabled returns the Enabled field value
func (o *AddCleanUpInactivePingfederatePersistentSessionsPluginRequest) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AddCleanUpInactivePingfederatePersistentSessionsPluginRequest) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AddCleanUpInactivePingfederatePersistentSessionsPluginRequest) SetEnabled(v bool) {
	o.Enabled = v
}

func (o AddCleanUpInactivePingfederatePersistentSessionsPluginRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pluginName"] = o.PluginName
	}
	if true {
		toSerialize["schemas"] = o.Schemas
	}
	if true {
		toSerialize["expirationOffset"] = o.ExpirationOffset
	}
	if !isNil(o.PollingInterval) {
		toSerialize["pollingInterval"] = o.PollingInterval
	}
	if !isNil(o.PeerServerPriorityIndex) {
		toSerialize["peerServerPriorityIndex"] = o.PeerServerPriorityIndex
	}
	if !isNil(o.BaseDN) {
		toSerialize["baseDN"] = o.BaseDN
	}
	if !isNil(o.MaxUpdatesPerSecond) {
		toSerialize["maxUpdatesPerSecond"] = o.MaxUpdatesPerSecond
	}
	if !isNil(o.NumDeleteThreads) {
		toSerialize["numDeleteThreads"] = o.NumDeleteThreads
	}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableAddCleanUpInactivePingfederatePersistentSessionsPluginRequest struct {
	value *AddCleanUpInactivePingfederatePersistentSessionsPluginRequest
	isSet bool
}

func (v NullableAddCleanUpInactivePingfederatePersistentSessionsPluginRequest) Get() *AddCleanUpInactivePingfederatePersistentSessionsPluginRequest {
	return v.value
}

func (v *NullableAddCleanUpInactivePingfederatePersistentSessionsPluginRequest) Set(val *AddCleanUpInactivePingfederatePersistentSessionsPluginRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddCleanUpInactivePingfederatePersistentSessionsPluginRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddCleanUpInactivePingfederatePersistentSessionsPluginRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddCleanUpInactivePingfederatePersistentSessionsPluginRequest(val *AddCleanUpInactivePingfederatePersistentSessionsPluginRequest) *NullableAddCleanUpInactivePingfederatePersistentSessionsPluginRequest {
	return &NullableAddCleanUpInactivePingfederatePersistentSessionsPluginRequest{value: val, isSet: true}
}

func (v NullableAddCleanUpInactivePingfederatePersistentSessionsPluginRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddCleanUpInactivePingfederatePersistentSessionsPluginRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
