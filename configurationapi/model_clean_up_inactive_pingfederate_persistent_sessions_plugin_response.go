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

// checks if the CleanUpInactivePingfederatePersistentSessionsPluginResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CleanUpInactivePingfederatePersistentSessionsPluginResponse{}

// CleanUpInactivePingfederatePersistentSessionsPluginResponse struct for CleanUpInactivePingfederatePersistentSessionsPluginResponse
type CleanUpInactivePingfederatePersistentSessionsPluginResponse struct {
	Schemas []EnumcleanUpInactivePingfederatePersistentSessionsPluginSchemaUrn `json:"schemas"`
	// Sessions whose last activity timestamp is older than this offset will be removed.
	ExpirationOffset string `json:"expirationOffset"`
	// This specifies how often the plugin should check for expired data. It also controls the offset of peer servers (see the peer-server-priority-index for more information).
	PollingInterval string `json:"pollingInterval"`
	// In a replicated environment, this determines the order in which peer servers should attempt to purge data.
	PeerServerPriorityIndex *int64 `json:"peerServerPriorityIndex,omitempty"`
	// Only entries located within the subtree specified by this base DN are eligible for purging.
	BaseDN *string `json:"baseDN,omitempty"`
	// This setting smooths out the performance impact on the server by throttling the purging to the specified maximum number of updates per second. To avoid a large backlog, this value should be set comfortably above the average rate that expired data is generated. When purge-behavior is set to subtree-delete-entries, then deletion of the entire subtree is considered a single update for the purposes of throttling.
	MaxUpdatesPerSecond int64 `json:"maxUpdatesPerSecond"`
	// The number of threads used to delete expired entries.
	NumDeleteThreads int64 `json:"numDeleteThreads"`
	// Indicates whether the plug-in is enabled for use.
	Enabled                                       bool                                               `json:"enabled"`
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
	// Name of the Plugin
	Id string `json:"id"`
}

// NewCleanUpInactivePingfederatePersistentSessionsPluginResponse instantiates a new CleanUpInactivePingfederatePersistentSessionsPluginResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCleanUpInactivePingfederatePersistentSessionsPluginResponse(schemas []EnumcleanUpInactivePingfederatePersistentSessionsPluginSchemaUrn, expirationOffset string, pollingInterval string, maxUpdatesPerSecond int64, numDeleteThreads int64, enabled bool, id string) *CleanUpInactivePingfederatePersistentSessionsPluginResponse {
	this := CleanUpInactivePingfederatePersistentSessionsPluginResponse{}
	this.Schemas = schemas
	this.ExpirationOffset = expirationOffset
	this.PollingInterval = pollingInterval
	this.MaxUpdatesPerSecond = maxUpdatesPerSecond
	this.NumDeleteThreads = numDeleteThreads
	this.Enabled = enabled
	this.Id = id
	return &this
}

// NewCleanUpInactivePingfederatePersistentSessionsPluginResponseWithDefaults instantiates a new CleanUpInactivePingfederatePersistentSessionsPluginResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCleanUpInactivePingfederatePersistentSessionsPluginResponseWithDefaults() *CleanUpInactivePingfederatePersistentSessionsPluginResponse {
	this := CleanUpInactivePingfederatePersistentSessionsPluginResponse{}
	return &this
}

// GetSchemas returns the Schemas field value
func (o *CleanUpInactivePingfederatePersistentSessionsPluginResponse) GetSchemas() []EnumcleanUpInactivePingfederatePersistentSessionsPluginSchemaUrn {
	if o == nil {
		var ret []EnumcleanUpInactivePingfederatePersistentSessionsPluginSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *CleanUpInactivePingfederatePersistentSessionsPluginResponse) GetSchemasOk() ([]EnumcleanUpInactivePingfederatePersistentSessionsPluginSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *CleanUpInactivePingfederatePersistentSessionsPluginResponse) SetSchemas(v []EnumcleanUpInactivePingfederatePersistentSessionsPluginSchemaUrn) {
	o.Schemas = v
}

// GetExpirationOffset returns the ExpirationOffset field value
func (o *CleanUpInactivePingfederatePersistentSessionsPluginResponse) GetExpirationOffset() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExpirationOffset
}

// GetExpirationOffsetOk returns a tuple with the ExpirationOffset field value
// and a boolean to check if the value has been set.
func (o *CleanUpInactivePingfederatePersistentSessionsPluginResponse) GetExpirationOffsetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpirationOffset, true
}

// SetExpirationOffset sets field value
func (o *CleanUpInactivePingfederatePersistentSessionsPluginResponse) SetExpirationOffset(v string) {
	o.ExpirationOffset = v
}

// GetPollingInterval returns the PollingInterval field value
func (o *CleanUpInactivePingfederatePersistentSessionsPluginResponse) GetPollingInterval() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PollingInterval
}

// GetPollingIntervalOk returns a tuple with the PollingInterval field value
// and a boolean to check if the value has been set.
func (o *CleanUpInactivePingfederatePersistentSessionsPluginResponse) GetPollingIntervalOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PollingInterval, true
}

// SetPollingInterval sets field value
func (o *CleanUpInactivePingfederatePersistentSessionsPluginResponse) SetPollingInterval(v string) {
	o.PollingInterval = v
}

// GetPeerServerPriorityIndex returns the PeerServerPriorityIndex field value if set, zero value otherwise.
func (o *CleanUpInactivePingfederatePersistentSessionsPluginResponse) GetPeerServerPriorityIndex() int64 {
	if o == nil || IsNil(o.PeerServerPriorityIndex) {
		var ret int64
		return ret
	}
	return *o.PeerServerPriorityIndex
}

// GetPeerServerPriorityIndexOk returns a tuple with the PeerServerPriorityIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CleanUpInactivePingfederatePersistentSessionsPluginResponse) GetPeerServerPriorityIndexOk() (*int64, bool) {
	if o == nil || IsNil(o.PeerServerPriorityIndex) {
		return nil, false
	}
	return o.PeerServerPriorityIndex, true
}

// HasPeerServerPriorityIndex returns a boolean if a field has been set.
func (o *CleanUpInactivePingfederatePersistentSessionsPluginResponse) HasPeerServerPriorityIndex() bool {
	if o != nil && !IsNil(o.PeerServerPriorityIndex) {
		return true
	}

	return false
}

// SetPeerServerPriorityIndex gets a reference to the given int64 and assigns it to the PeerServerPriorityIndex field.
func (o *CleanUpInactivePingfederatePersistentSessionsPluginResponse) SetPeerServerPriorityIndex(v int64) {
	o.PeerServerPriorityIndex = &v
}

// GetBaseDN returns the BaseDN field value if set, zero value otherwise.
func (o *CleanUpInactivePingfederatePersistentSessionsPluginResponse) GetBaseDN() string {
	if o == nil || IsNil(o.BaseDN) {
		var ret string
		return ret
	}
	return *o.BaseDN
}

// GetBaseDNOk returns a tuple with the BaseDN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CleanUpInactivePingfederatePersistentSessionsPluginResponse) GetBaseDNOk() (*string, bool) {
	if o == nil || IsNil(o.BaseDN) {
		return nil, false
	}
	return o.BaseDN, true
}

// HasBaseDN returns a boolean if a field has been set.
func (o *CleanUpInactivePingfederatePersistentSessionsPluginResponse) HasBaseDN() bool {
	if o != nil && !IsNil(o.BaseDN) {
		return true
	}

	return false
}

// SetBaseDN gets a reference to the given string and assigns it to the BaseDN field.
func (o *CleanUpInactivePingfederatePersistentSessionsPluginResponse) SetBaseDN(v string) {
	o.BaseDN = &v
}

// GetMaxUpdatesPerSecond returns the MaxUpdatesPerSecond field value
func (o *CleanUpInactivePingfederatePersistentSessionsPluginResponse) GetMaxUpdatesPerSecond() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.MaxUpdatesPerSecond
}

// GetMaxUpdatesPerSecondOk returns a tuple with the MaxUpdatesPerSecond field value
// and a boolean to check if the value has been set.
func (o *CleanUpInactivePingfederatePersistentSessionsPluginResponse) GetMaxUpdatesPerSecondOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxUpdatesPerSecond, true
}

// SetMaxUpdatesPerSecond sets field value
func (o *CleanUpInactivePingfederatePersistentSessionsPluginResponse) SetMaxUpdatesPerSecond(v int64) {
	o.MaxUpdatesPerSecond = v
}

// GetNumDeleteThreads returns the NumDeleteThreads field value
func (o *CleanUpInactivePingfederatePersistentSessionsPluginResponse) GetNumDeleteThreads() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.NumDeleteThreads
}

// GetNumDeleteThreadsOk returns a tuple with the NumDeleteThreads field value
// and a boolean to check if the value has been set.
func (o *CleanUpInactivePingfederatePersistentSessionsPluginResponse) GetNumDeleteThreadsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NumDeleteThreads, true
}

// SetNumDeleteThreads sets field value
func (o *CleanUpInactivePingfederatePersistentSessionsPluginResponse) SetNumDeleteThreads(v int64) {
	o.NumDeleteThreads = v
}

// GetEnabled returns the Enabled field value
func (o *CleanUpInactivePingfederatePersistentSessionsPluginResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *CleanUpInactivePingfederatePersistentSessionsPluginResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *CleanUpInactivePingfederatePersistentSessionsPluginResponse) SetEnabled(v bool) {
	o.Enabled = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *CleanUpInactivePingfederatePersistentSessionsPluginResponse) GetMeta() MetaMeta {
	if o == nil || IsNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CleanUpInactivePingfederatePersistentSessionsPluginResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *CleanUpInactivePingfederatePersistentSessionsPluginResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *CleanUpInactivePingfederatePersistentSessionsPluginResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *CleanUpInactivePingfederatePersistentSessionsPluginResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CleanUpInactivePingfederatePersistentSessionsPluginResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *CleanUpInactivePingfederatePersistentSessionsPluginResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *CleanUpInactivePingfederatePersistentSessionsPluginResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetId returns the Id field value
func (o *CleanUpInactivePingfederatePersistentSessionsPluginResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CleanUpInactivePingfederatePersistentSessionsPluginResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CleanUpInactivePingfederatePersistentSessionsPluginResponse) SetId(v string) {
	o.Id = v
}

func (o CleanUpInactivePingfederatePersistentSessionsPluginResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CleanUpInactivePingfederatePersistentSessionsPluginResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["schemas"] = o.Schemas
	toSerialize["expirationOffset"] = o.ExpirationOffset
	toSerialize["pollingInterval"] = o.PollingInterval
	if !IsNil(o.PeerServerPriorityIndex) {
		toSerialize["peerServerPriorityIndex"] = o.PeerServerPriorityIndex
	}
	if !IsNil(o.BaseDN) {
		toSerialize["baseDN"] = o.BaseDN
	}
	toSerialize["maxUpdatesPerSecond"] = o.MaxUpdatesPerSecond
	toSerialize["numDeleteThreads"] = o.NumDeleteThreads
	toSerialize["enabled"] = o.Enabled
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		toSerialize["urn:pingidentity:schemas:configuration:messages:2.0"] = o.Urnpingidentityschemasconfigurationmessages20
	}
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

type NullableCleanUpInactivePingfederatePersistentSessionsPluginResponse struct {
	value *CleanUpInactivePingfederatePersistentSessionsPluginResponse
	isSet bool
}

func (v NullableCleanUpInactivePingfederatePersistentSessionsPluginResponse) Get() *CleanUpInactivePingfederatePersistentSessionsPluginResponse {
	return v.value
}

func (v *NullableCleanUpInactivePingfederatePersistentSessionsPluginResponse) Set(val *CleanUpInactivePingfederatePersistentSessionsPluginResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCleanUpInactivePingfederatePersistentSessionsPluginResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCleanUpInactivePingfederatePersistentSessionsPluginResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCleanUpInactivePingfederatePersistentSessionsPluginResponse(val *CleanUpInactivePingfederatePersistentSessionsPluginResponse) *NullableCleanUpInactivePingfederatePersistentSessionsPluginResponse {
	return &NullableCleanUpInactivePingfederatePersistentSessionsPluginResponse{value: val, isSet: true}
}

func (v NullableCleanUpInactivePingfederatePersistentSessionsPluginResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCleanUpInactivePingfederatePersistentSessionsPluginResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
