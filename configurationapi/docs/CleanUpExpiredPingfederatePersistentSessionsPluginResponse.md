# CleanUpExpiredPingfederatePersistentSessionsPluginResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Id** | **string** | Name of the Plugin | 
**Schemas** | [**[]EnumcleanUpExpiredPingfederatePersistentSessionsPluginSchemaUrn**](EnumcleanUpExpiredPingfederatePersistentSessionsPluginSchemaUrn.md) |  | 
**PollingInterval** | **string** | This specifies how often the plugin should check for expired data. It also controls the offset of peer servers (see the peer-server-priority-index for more information). | 
**PeerServerPriorityIndex** | Pointer to **int32** | In a replicated environment, this determines the order in which peer servers should attempt to purge data. | [optional] 
**BaseDN** | Pointer to **string** | Only entries located within the subtree specified by this base DN are eligible for purging. | [optional] 
**MaxUpdatesPerSecond** | **int32** | This setting smooths out the performance impact on the server by throttling the purging to the specified maximum number of updates per second. To avoid a large backlog, this value should be set comfortably above the average rate that expired data is generated. When purge-behavior is set to subtree-delete-entries, then deletion of the entire subtree is considered a single update for the purposes of throttling. | 
**NumDeleteThreads** | **int32** | The number of threads used to delete expired entries. | 
**Enabled** | **bool** | Indicates whether the plug-in is enabled for use. | 

## Methods

### NewCleanUpExpiredPingfederatePersistentSessionsPluginResponse

`func NewCleanUpExpiredPingfederatePersistentSessionsPluginResponse(id string, schemas []EnumcleanUpExpiredPingfederatePersistentSessionsPluginSchemaUrn, pollingInterval string, maxUpdatesPerSecond int32, numDeleteThreads int32, enabled bool, ) *CleanUpExpiredPingfederatePersistentSessionsPluginResponse`

NewCleanUpExpiredPingfederatePersistentSessionsPluginResponse instantiates a new CleanUpExpiredPingfederatePersistentSessionsPluginResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCleanUpExpiredPingfederatePersistentSessionsPluginResponseWithDefaults

`func NewCleanUpExpiredPingfederatePersistentSessionsPluginResponseWithDefaults() *CleanUpExpiredPingfederatePersistentSessionsPluginResponse`

NewCleanUpExpiredPingfederatePersistentSessionsPluginResponseWithDefaults instantiates a new CleanUpExpiredPingfederatePersistentSessionsPluginResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *CleanUpExpiredPingfederatePersistentSessionsPluginResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *CleanUpExpiredPingfederatePersistentSessionsPluginResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *CleanUpExpiredPingfederatePersistentSessionsPluginResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *CleanUpExpiredPingfederatePersistentSessionsPluginResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *CleanUpExpiredPingfederatePersistentSessionsPluginResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *CleanUpExpiredPingfederatePersistentSessionsPluginResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *CleanUpExpiredPingfederatePersistentSessionsPluginResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *CleanUpExpiredPingfederatePersistentSessionsPluginResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetId

`func (o *CleanUpExpiredPingfederatePersistentSessionsPluginResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CleanUpExpiredPingfederatePersistentSessionsPluginResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CleanUpExpiredPingfederatePersistentSessionsPluginResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *CleanUpExpiredPingfederatePersistentSessionsPluginResponse) GetSchemas() []EnumcleanUpExpiredPingfederatePersistentSessionsPluginSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *CleanUpExpiredPingfederatePersistentSessionsPluginResponse) GetSchemasOk() (*[]EnumcleanUpExpiredPingfederatePersistentSessionsPluginSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *CleanUpExpiredPingfederatePersistentSessionsPluginResponse) SetSchemas(v []EnumcleanUpExpiredPingfederatePersistentSessionsPluginSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetPollingInterval

`func (o *CleanUpExpiredPingfederatePersistentSessionsPluginResponse) GetPollingInterval() string`

GetPollingInterval returns the PollingInterval field if non-nil, zero value otherwise.

### GetPollingIntervalOk

`func (o *CleanUpExpiredPingfederatePersistentSessionsPluginResponse) GetPollingIntervalOk() (*string, bool)`

GetPollingIntervalOk returns a tuple with the PollingInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPollingInterval

`func (o *CleanUpExpiredPingfederatePersistentSessionsPluginResponse) SetPollingInterval(v string)`

SetPollingInterval sets PollingInterval field to given value.


### GetPeerServerPriorityIndex

`func (o *CleanUpExpiredPingfederatePersistentSessionsPluginResponse) GetPeerServerPriorityIndex() int32`

GetPeerServerPriorityIndex returns the PeerServerPriorityIndex field if non-nil, zero value otherwise.

### GetPeerServerPriorityIndexOk

`func (o *CleanUpExpiredPingfederatePersistentSessionsPluginResponse) GetPeerServerPriorityIndexOk() (*int32, bool)`

GetPeerServerPriorityIndexOk returns a tuple with the PeerServerPriorityIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerServerPriorityIndex

`func (o *CleanUpExpiredPingfederatePersistentSessionsPluginResponse) SetPeerServerPriorityIndex(v int32)`

SetPeerServerPriorityIndex sets PeerServerPriorityIndex field to given value.

### HasPeerServerPriorityIndex

`func (o *CleanUpExpiredPingfederatePersistentSessionsPluginResponse) HasPeerServerPriorityIndex() bool`

HasPeerServerPriorityIndex returns a boolean if a field has been set.

### GetBaseDN

`func (o *CleanUpExpiredPingfederatePersistentSessionsPluginResponse) GetBaseDN() string`

GetBaseDN returns the BaseDN field if non-nil, zero value otherwise.

### GetBaseDNOk

`func (o *CleanUpExpiredPingfederatePersistentSessionsPluginResponse) GetBaseDNOk() (*string, bool)`

GetBaseDNOk returns a tuple with the BaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseDN

`func (o *CleanUpExpiredPingfederatePersistentSessionsPluginResponse) SetBaseDN(v string)`

SetBaseDN sets BaseDN field to given value.

### HasBaseDN

`func (o *CleanUpExpiredPingfederatePersistentSessionsPluginResponse) HasBaseDN() bool`

HasBaseDN returns a boolean if a field has been set.

### GetMaxUpdatesPerSecond

`func (o *CleanUpExpiredPingfederatePersistentSessionsPluginResponse) GetMaxUpdatesPerSecond() int32`

GetMaxUpdatesPerSecond returns the MaxUpdatesPerSecond field if non-nil, zero value otherwise.

### GetMaxUpdatesPerSecondOk

`func (o *CleanUpExpiredPingfederatePersistentSessionsPluginResponse) GetMaxUpdatesPerSecondOk() (*int32, bool)`

GetMaxUpdatesPerSecondOk returns a tuple with the MaxUpdatesPerSecond field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxUpdatesPerSecond

`func (o *CleanUpExpiredPingfederatePersistentSessionsPluginResponse) SetMaxUpdatesPerSecond(v int32)`

SetMaxUpdatesPerSecond sets MaxUpdatesPerSecond field to given value.


### GetNumDeleteThreads

`func (o *CleanUpExpiredPingfederatePersistentSessionsPluginResponse) GetNumDeleteThreads() int32`

GetNumDeleteThreads returns the NumDeleteThreads field if non-nil, zero value otherwise.

### GetNumDeleteThreadsOk

`func (o *CleanUpExpiredPingfederatePersistentSessionsPluginResponse) GetNumDeleteThreadsOk() (*int32, bool)`

GetNumDeleteThreadsOk returns a tuple with the NumDeleteThreads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumDeleteThreads

`func (o *CleanUpExpiredPingfederatePersistentSessionsPluginResponse) SetNumDeleteThreads(v int32)`

SetNumDeleteThreads sets NumDeleteThreads field to given value.


### GetEnabled

`func (o *CleanUpExpiredPingfederatePersistentSessionsPluginResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CleanUpExpiredPingfederatePersistentSessionsPluginResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CleanUpExpiredPingfederatePersistentSessionsPluginResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


