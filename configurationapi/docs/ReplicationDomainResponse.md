# ReplicationDomainResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | Pointer to [**[]EnumreplicationDomainSchemaUrn**](EnumreplicationDomainSchemaUrn.md) |  | [optional] 
**ServerID** | **int32** | Specifies a unique identifier for the Directory Server within the Replication Domain. | 
**BaseDN** | **string** | Specifies the base DN of the replicated data. | 
**WindowSize** | Pointer to **int32** | Specifies the maximum number of replication updates the Directory Server can have outstanding from the Replication Server. | [optional] 
**HeartbeatInterval** | Pointer to **string** | Specifies the heartbeat interval that the Directory Server will use when communicating with Replication Servers. | [optional] 
**SyncHistPurgeDelay** | Pointer to **string** | The time in seconds after which historical information used in replication conflict resolution is purged. The information is removed from entries when they are modified after the purge delay has elapsed. | [optional] 
**Restricted** | Pointer to **bool** | When set to true, changes are only replicated with server instances that belong to the same replication set. | [optional] 
**OnReplayFailureWaitForDependentOpsTimeout** | Pointer to **string** | Defines the maximum time to retry a failed operation. An operation will be retried only if it appears that the failure might be dependent on an earlier operation from a different server that hasn&#39;t replicated yet. The frequency of the retry is determined by the dependent-ops-replay-failure-wait-time property. | [optional] 
**DependentOpsReplayFailureWaitTime** | Pointer to **string** | Defines how long to wait before retrying certain operations, specifically operations that might have failed because they depend on an operation from a different server that has not yet replicated to this instance. | [optional] 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 

## Methods

### NewReplicationDomainResponse

`func NewReplicationDomainResponse(serverID int32, baseDN string, ) *ReplicationDomainResponse`

NewReplicationDomainResponse instantiates a new ReplicationDomainResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReplicationDomainResponseWithDefaults

`func NewReplicationDomainResponseWithDefaults() *ReplicationDomainResponse`

NewReplicationDomainResponseWithDefaults instantiates a new ReplicationDomainResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *ReplicationDomainResponse) GetSchemas() []EnumreplicationDomainSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *ReplicationDomainResponse) GetSchemasOk() (*[]EnumreplicationDomainSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *ReplicationDomainResponse) SetSchemas(v []EnumreplicationDomainSchemaUrn)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *ReplicationDomainResponse) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetServerID

`func (o *ReplicationDomainResponse) GetServerID() int32`

GetServerID returns the ServerID field if non-nil, zero value otherwise.

### GetServerIDOk

`func (o *ReplicationDomainResponse) GetServerIDOk() (*int32, bool)`

GetServerIDOk returns a tuple with the ServerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerID

`func (o *ReplicationDomainResponse) SetServerID(v int32)`

SetServerID sets ServerID field to given value.


### GetBaseDN

`func (o *ReplicationDomainResponse) GetBaseDN() string`

GetBaseDN returns the BaseDN field if non-nil, zero value otherwise.

### GetBaseDNOk

`func (o *ReplicationDomainResponse) GetBaseDNOk() (*string, bool)`

GetBaseDNOk returns a tuple with the BaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseDN

`func (o *ReplicationDomainResponse) SetBaseDN(v string)`

SetBaseDN sets BaseDN field to given value.


### GetWindowSize

`func (o *ReplicationDomainResponse) GetWindowSize() int32`

GetWindowSize returns the WindowSize field if non-nil, zero value otherwise.

### GetWindowSizeOk

`func (o *ReplicationDomainResponse) GetWindowSizeOk() (*int32, bool)`

GetWindowSizeOk returns a tuple with the WindowSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowSize

`func (o *ReplicationDomainResponse) SetWindowSize(v int32)`

SetWindowSize sets WindowSize field to given value.

### HasWindowSize

`func (o *ReplicationDomainResponse) HasWindowSize() bool`

HasWindowSize returns a boolean if a field has been set.

### GetHeartbeatInterval

`func (o *ReplicationDomainResponse) GetHeartbeatInterval() string`

GetHeartbeatInterval returns the HeartbeatInterval field if non-nil, zero value otherwise.

### GetHeartbeatIntervalOk

`func (o *ReplicationDomainResponse) GetHeartbeatIntervalOk() (*string, bool)`

GetHeartbeatIntervalOk returns a tuple with the HeartbeatInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeartbeatInterval

`func (o *ReplicationDomainResponse) SetHeartbeatInterval(v string)`

SetHeartbeatInterval sets HeartbeatInterval field to given value.

### HasHeartbeatInterval

`func (o *ReplicationDomainResponse) HasHeartbeatInterval() bool`

HasHeartbeatInterval returns a boolean if a field has been set.

### GetSyncHistPurgeDelay

`func (o *ReplicationDomainResponse) GetSyncHistPurgeDelay() string`

GetSyncHistPurgeDelay returns the SyncHistPurgeDelay field if non-nil, zero value otherwise.

### GetSyncHistPurgeDelayOk

`func (o *ReplicationDomainResponse) GetSyncHistPurgeDelayOk() (*string, bool)`

GetSyncHistPurgeDelayOk returns a tuple with the SyncHistPurgeDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncHistPurgeDelay

`func (o *ReplicationDomainResponse) SetSyncHistPurgeDelay(v string)`

SetSyncHistPurgeDelay sets SyncHistPurgeDelay field to given value.

### HasSyncHistPurgeDelay

`func (o *ReplicationDomainResponse) HasSyncHistPurgeDelay() bool`

HasSyncHistPurgeDelay returns a boolean if a field has been set.

### GetRestricted

`func (o *ReplicationDomainResponse) GetRestricted() bool`

GetRestricted returns the Restricted field if non-nil, zero value otherwise.

### GetRestrictedOk

`func (o *ReplicationDomainResponse) GetRestrictedOk() (*bool, bool)`

GetRestrictedOk returns a tuple with the Restricted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestricted

`func (o *ReplicationDomainResponse) SetRestricted(v bool)`

SetRestricted sets Restricted field to given value.

### HasRestricted

`func (o *ReplicationDomainResponse) HasRestricted() bool`

HasRestricted returns a boolean if a field has been set.

### GetOnReplayFailureWaitForDependentOpsTimeout

`func (o *ReplicationDomainResponse) GetOnReplayFailureWaitForDependentOpsTimeout() string`

GetOnReplayFailureWaitForDependentOpsTimeout returns the OnReplayFailureWaitForDependentOpsTimeout field if non-nil, zero value otherwise.

### GetOnReplayFailureWaitForDependentOpsTimeoutOk

`func (o *ReplicationDomainResponse) GetOnReplayFailureWaitForDependentOpsTimeoutOk() (*string, bool)`

GetOnReplayFailureWaitForDependentOpsTimeoutOk returns a tuple with the OnReplayFailureWaitForDependentOpsTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnReplayFailureWaitForDependentOpsTimeout

`func (o *ReplicationDomainResponse) SetOnReplayFailureWaitForDependentOpsTimeout(v string)`

SetOnReplayFailureWaitForDependentOpsTimeout sets OnReplayFailureWaitForDependentOpsTimeout field to given value.

### HasOnReplayFailureWaitForDependentOpsTimeout

`func (o *ReplicationDomainResponse) HasOnReplayFailureWaitForDependentOpsTimeout() bool`

HasOnReplayFailureWaitForDependentOpsTimeout returns a boolean if a field has been set.

### GetDependentOpsReplayFailureWaitTime

`func (o *ReplicationDomainResponse) GetDependentOpsReplayFailureWaitTime() string`

GetDependentOpsReplayFailureWaitTime returns the DependentOpsReplayFailureWaitTime field if non-nil, zero value otherwise.

### GetDependentOpsReplayFailureWaitTimeOk

`func (o *ReplicationDomainResponse) GetDependentOpsReplayFailureWaitTimeOk() (*string, bool)`

GetDependentOpsReplayFailureWaitTimeOk returns a tuple with the DependentOpsReplayFailureWaitTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependentOpsReplayFailureWaitTime

`func (o *ReplicationDomainResponse) SetDependentOpsReplayFailureWaitTime(v string)`

SetDependentOpsReplayFailureWaitTime sets DependentOpsReplayFailureWaitTime field to given value.

### HasDependentOpsReplayFailureWaitTime

`func (o *ReplicationDomainResponse) HasDependentOpsReplayFailureWaitTime() bool`

HasDependentOpsReplayFailureWaitTime returns a boolean if a field has been set.

### GetMeta

`func (o *ReplicationDomainResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ReplicationDomainResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ReplicationDomainResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ReplicationDomainResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *ReplicationDomainResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *ReplicationDomainResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *ReplicationDomainResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *ReplicationDomainResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


