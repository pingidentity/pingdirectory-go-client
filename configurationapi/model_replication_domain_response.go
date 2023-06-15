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

// checks if the ReplicationDomainResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReplicationDomainResponse{}

// ReplicationDomainResponse struct for ReplicationDomainResponse
type ReplicationDomainResponse struct {
	Schemas []EnumreplicationDomainSchemaUrn `json:"schemas,omitempty"`
	// Name of the Replication Domain
	Id string `json:"id"`
	// Specifies a unique identifier for the Directory Server within the Replication Domain.
	ServerID int64 `json:"serverID"`
	// Specifies the base DN of the replicated data.
	BaseDN string `json:"baseDN"`
	// Specifies the maximum number of replication updates the Directory Server can have outstanding from the Replication Server.
	WindowSize *int64 `json:"windowSize,omitempty"`
	// Specifies the heartbeat interval that the Directory Server will use when communicating with Replication Servers.
	HeartbeatInterval *string `json:"heartbeatInterval,omitempty"`
	// The time in seconds after which historical information used in replication conflict resolution is purged. The information is removed from entries when they are modified after the purge delay has elapsed.
	SyncHistPurgeDelay *string `json:"syncHistPurgeDelay,omitempty"`
	// When set to true, changes are only replicated with server instances that belong to the same replication set.
	Restricted *bool `json:"restricted,omitempty"`
	// Defines the maximum time to retry a failed operation. An operation will be retried only if it appears that the failure might be dependent on an earlier operation from a different server that hasn't replicated yet. The frequency of the retry is determined by the dependent-ops-replay-failure-wait-time property.
	OnReplayFailureWaitForDependentOpsTimeout *string `json:"onReplayFailureWaitForDependentOpsTimeout,omitempty"`
	// Defines how long to wait before retrying certain operations, specifically operations that might have failed because they depend on an operation from a different server that has not yet replicated to this instance.
	DependentOpsReplayFailureWaitTime             *string                                            `json:"dependentOpsReplayFailureWaitTime,omitempty"`
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
}

// NewReplicationDomainResponse instantiates a new ReplicationDomainResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReplicationDomainResponse(id string, serverID int64, baseDN string) *ReplicationDomainResponse {
	this := ReplicationDomainResponse{}
	this.Id = id
	this.ServerID = serverID
	this.BaseDN = baseDN
	return &this
}

// NewReplicationDomainResponseWithDefaults instantiates a new ReplicationDomainResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReplicationDomainResponseWithDefaults() *ReplicationDomainResponse {
	this := ReplicationDomainResponse{}
	return &this
}

// GetSchemas returns the Schemas field value if set, zero value otherwise.
func (o *ReplicationDomainResponse) GetSchemas() []EnumreplicationDomainSchemaUrn {
	if o == nil || IsNil(o.Schemas) {
		var ret []EnumreplicationDomainSchemaUrn
		return ret
	}
	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplicationDomainResponse) GetSchemasOk() ([]EnumreplicationDomainSchemaUrn, bool) {
	if o == nil || IsNil(o.Schemas) {
		return nil, false
	}
	return o.Schemas, true
}

// HasSchemas returns a boolean if a field has been set.
func (o *ReplicationDomainResponse) HasSchemas() bool {
	if o != nil && !IsNil(o.Schemas) {
		return true
	}

	return false
}

// SetSchemas gets a reference to the given []EnumreplicationDomainSchemaUrn and assigns it to the Schemas field.
func (o *ReplicationDomainResponse) SetSchemas(v []EnumreplicationDomainSchemaUrn) {
	o.Schemas = v
}

// GetId returns the Id field value
func (o *ReplicationDomainResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ReplicationDomainResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ReplicationDomainResponse) SetId(v string) {
	o.Id = v
}

// GetServerID returns the ServerID field value
func (o *ReplicationDomainResponse) GetServerID() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ServerID
}

// GetServerIDOk returns a tuple with the ServerID field value
// and a boolean to check if the value has been set.
func (o *ReplicationDomainResponse) GetServerIDOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServerID, true
}

// SetServerID sets field value
func (o *ReplicationDomainResponse) SetServerID(v int64) {
	o.ServerID = v
}

// GetBaseDN returns the BaseDN field value
func (o *ReplicationDomainResponse) GetBaseDN() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BaseDN
}

// GetBaseDNOk returns a tuple with the BaseDN field value
// and a boolean to check if the value has been set.
func (o *ReplicationDomainResponse) GetBaseDNOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BaseDN, true
}

// SetBaseDN sets field value
func (o *ReplicationDomainResponse) SetBaseDN(v string) {
	o.BaseDN = v
}

// GetWindowSize returns the WindowSize field value if set, zero value otherwise.
func (o *ReplicationDomainResponse) GetWindowSize() int64 {
	if o == nil || IsNil(o.WindowSize) {
		var ret int64
		return ret
	}
	return *o.WindowSize
}

// GetWindowSizeOk returns a tuple with the WindowSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplicationDomainResponse) GetWindowSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.WindowSize) {
		return nil, false
	}
	return o.WindowSize, true
}

// HasWindowSize returns a boolean if a field has been set.
func (o *ReplicationDomainResponse) HasWindowSize() bool {
	if o != nil && !IsNil(o.WindowSize) {
		return true
	}

	return false
}

// SetWindowSize gets a reference to the given int64 and assigns it to the WindowSize field.
func (o *ReplicationDomainResponse) SetWindowSize(v int64) {
	o.WindowSize = &v
}

// GetHeartbeatInterval returns the HeartbeatInterval field value if set, zero value otherwise.
func (o *ReplicationDomainResponse) GetHeartbeatInterval() string {
	if o == nil || IsNil(o.HeartbeatInterval) {
		var ret string
		return ret
	}
	return *o.HeartbeatInterval
}

// GetHeartbeatIntervalOk returns a tuple with the HeartbeatInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplicationDomainResponse) GetHeartbeatIntervalOk() (*string, bool) {
	if o == nil || IsNil(o.HeartbeatInterval) {
		return nil, false
	}
	return o.HeartbeatInterval, true
}

// HasHeartbeatInterval returns a boolean if a field has been set.
func (o *ReplicationDomainResponse) HasHeartbeatInterval() bool {
	if o != nil && !IsNil(o.HeartbeatInterval) {
		return true
	}

	return false
}

// SetHeartbeatInterval gets a reference to the given string and assigns it to the HeartbeatInterval field.
func (o *ReplicationDomainResponse) SetHeartbeatInterval(v string) {
	o.HeartbeatInterval = &v
}

// GetSyncHistPurgeDelay returns the SyncHistPurgeDelay field value if set, zero value otherwise.
func (o *ReplicationDomainResponse) GetSyncHistPurgeDelay() string {
	if o == nil || IsNil(o.SyncHistPurgeDelay) {
		var ret string
		return ret
	}
	return *o.SyncHistPurgeDelay
}

// GetSyncHistPurgeDelayOk returns a tuple with the SyncHistPurgeDelay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplicationDomainResponse) GetSyncHistPurgeDelayOk() (*string, bool) {
	if o == nil || IsNil(o.SyncHistPurgeDelay) {
		return nil, false
	}
	return o.SyncHistPurgeDelay, true
}

// HasSyncHistPurgeDelay returns a boolean if a field has been set.
func (o *ReplicationDomainResponse) HasSyncHistPurgeDelay() bool {
	if o != nil && !IsNil(o.SyncHistPurgeDelay) {
		return true
	}

	return false
}

// SetSyncHistPurgeDelay gets a reference to the given string and assigns it to the SyncHistPurgeDelay field.
func (o *ReplicationDomainResponse) SetSyncHistPurgeDelay(v string) {
	o.SyncHistPurgeDelay = &v
}

// GetRestricted returns the Restricted field value if set, zero value otherwise.
func (o *ReplicationDomainResponse) GetRestricted() bool {
	if o == nil || IsNil(o.Restricted) {
		var ret bool
		return ret
	}
	return *o.Restricted
}

// GetRestrictedOk returns a tuple with the Restricted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplicationDomainResponse) GetRestrictedOk() (*bool, bool) {
	if o == nil || IsNil(o.Restricted) {
		return nil, false
	}
	return o.Restricted, true
}

// HasRestricted returns a boolean if a field has been set.
func (o *ReplicationDomainResponse) HasRestricted() bool {
	if o != nil && !IsNil(o.Restricted) {
		return true
	}

	return false
}

// SetRestricted gets a reference to the given bool and assigns it to the Restricted field.
func (o *ReplicationDomainResponse) SetRestricted(v bool) {
	o.Restricted = &v
}

// GetOnReplayFailureWaitForDependentOpsTimeout returns the OnReplayFailureWaitForDependentOpsTimeout field value if set, zero value otherwise.
func (o *ReplicationDomainResponse) GetOnReplayFailureWaitForDependentOpsTimeout() string {
	if o == nil || IsNil(o.OnReplayFailureWaitForDependentOpsTimeout) {
		var ret string
		return ret
	}
	return *o.OnReplayFailureWaitForDependentOpsTimeout
}

// GetOnReplayFailureWaitForDependentOpsTimeoutOk returns a tuple with the OnReplayFailureWaitForDependentOpsTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplicationDomainResponse) GetOnReplayFailureWaitForDependentOpsTimeoutOk() (*string, bool) {
	if o == nil || IsNil(o.OnReplayFailureWaitForDependentOpsTimeout) {
		return nil, false
	}
	return o.OnReplayFailureWaitForDependentOpsTimeout, true
}

// HasOnReplayFailureWaitForDependentOpsTimeout returns a boolean if a field has been set.
func (o *ReplicationDomainResponse) HasOnReplayFailureWaitForDependentOpsTimeout() bool {
	if o != nil && !IsNil(o.OnReplayFailureWaitForDependentOpsTimeout) {
		return true
	}

	return false
}

// SetOnReplayFailureWaitForDependentOpsTimeout gets a reference to the given string and assigns it to the OnReplayFailureWaitForDependentOpsTimeout field.
func (o *ReplicationDomainResponse) SetOnReplayFailureWaitForDependentOpsTimeout(v string) {
	o.OnReplayFailureWaitForDependentOpsTimeout = &v
}

// GetDependentOpsReplayFailureWaitTime returns the DependentOpsReplayFailureWaitTime field value if set, zero value otherwise.
func (o *ReplicationDomainResponse) GetDependentOpsReplayFailureWaitTime() string {
	if o == nil || IsNil(o.DependentOpsReplayFailureWaitTime) {
		var ret string
		return ret
	}
	return *o.DependentOpsReplayFailureWaitTime
}

// GetDependentOpsReplayFailureWaitTimeOk returns a tuple with the DependentOpsReplayFailureWaitTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplicationDomainResponse) GetDependentOpsReplayFailureWaitTimeOk() (*string, bool) {
	if o == nil || IsNil(o.DependentOpsReplayFailureWaitTime) {
		return nil, false
	}
	return o.DependentOpsReplayFailureWaitTime, true
}

// HasDependentOpsReplayFailureWaitTime returns a boolean if a field has been set.
func (o *ReplicationDomainResponse) HasDependentOpsReplayFailureWaitTime() bool {
	if o != nil && !IsNil(o.DependentOpsReplayFailureWaitTime) {
		return true
	}

	return false
}

// SetDependentOpsReplayFailureWaitTime gets a reference to the given string and assigns it to the DependentOpsReplayFailureWaitTime field.
func (o *ReplicationDomainResponse) SetDependentOpsReplayFailureWaitTime(v string) {
	o.DependentOpsReplayFailureWaitTime = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ReplicationDomainResponse) GetMeta() MetaMeta {
	if o == nil || IsNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplicationDomainResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ReplicationDomainResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *ReplicationDomainResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *ReplicationDomainResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplicationDomainResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *ReplicationDomainResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *ReplicationDomainResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

func (o ReplicationDomainResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReplicationDomainResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Schemas) {
		toSerialize["schemas"] = o.Schemas
	}
	toSerialize["id"] = o.Id
	toSerialize["serverID"] = o.ServerID
	toSerialize["baseDN"] = o.BaseDN
	if !IsNil(o.WindowSize) {
		toSerialize["windowSize"] = o.WindowSize
	}
	if !IsNil(o.HeartbeatInterval) {
		toSerialize["heartbeatInterval"] = o.HeartbeatInterval
	}
	if !IsNil(o.SyncHistPurgeDelay) {
		toSerialize["syncHistPurgeDelay"] = o.SyncHistPurgeDelay
	}
	if !IsNil(o.Restricted) {
		toSerialize["restricted"] = o.Restricted
	}
	if !IsNil(o.OnReplayFailureWaitForDependentOpsTimeout) {
		toSerialize["onReplayFailureWaitForDependentOpsTimeout"] = o.OnReplayFailureWaitForDependentOpsTimeout
	}
	if !IsNil(o.DependentOpsReplayFailureWaitTime) {
		toSerialize["dependentOpsReplayFailureWaitTime"] = o.DependentOpsReplayFailureWaitTime
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		toSerialize["urn:pingidentity:schemas:configuration:messages:2.0"] = o.Urnpingidentityschemasconfigurationmessages20
	}
	return toSerialize, nil
}

type NullableReplicationDomainResponse struct {
	value *ReplicationDomainResponse
	isSet bool
}

func (v NullableReplicationDomainResponse) Get() *ReplicationDomainResponse {
	return v.value
}

func (v *NullableReplicationDomainResponse) Set(val *ReplicationDomainResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableReplicationDomainResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableReplicationDomainResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReplicationDomainResponse(val *ReplicationDomainResponse) *NullableReplicationDomainResponse {
	return &NullableReplicationDomainResponse{value: val, isSet: true}
}

func (v NullableReplicationDomainResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReplicationDomainResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
