# CleanUpExpiredPingfederatePersistentAccessGrantsPluginResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Id** | **string** | Name of the Plugin | 
**Schemas** | [**[]EnumcleanUpExpiredPingfederatePersistentAccessGrantsPluginSchemaUrn**](EnumcleanUpExpiredPingfederatePersistentAccessGrantsPluginSchemaUrn.md) |  | 
**PollingInterval** | **string** | This specifies how often the plugin should check for expired data. It also controls the offset of peer servers (see the peer-server-priority-index for more information). | 
**PeerServerPriorityIndex** | Pointer to **int32** | In a replicated environment, this determines the order in which peer servers should attempt to purge data. | [optional] 
**BaseDN** | Pointer to **string** | Only entries located within the subtree specified by this base DN are eligible for purging. | [optional] 
**MaxUpdatesPerSecond** | **int32** | This setting smooths out the performance impact on the server by throttling the purging to the specified maximum number of updates per second. To avoid a large backlog, this value should be set comfortably above the average rate that expired data is generated. When purge-behavior is set to subtree-delete-entries, then deletion of the entire subtree is considered a single update for the purposes of throttling. | 
**NumDeleteThreads** | **int32** | The number of threads used to delete expired entries. | 
**Enabled** | **bool** | Indicates whether the plug-in is enabled for use. | 

## Methods

### NewCleanUpExpiredPingfederatePersistentAccessGrantsPluginResponse

`func NewCleanUpExpiredPingfederatePersistentAccessGrantsPluginResponse(id string, schemas []EnumcleanUpExpiredPingfederatePersistentAccessGrantsPluginSchemaUrn, pollingInterval string, maxUpdatesPerSecond int32, numDeleteThreads int32, enabled bool, ) *CleanUpExpiredPingfederatePersistentAccessGrantsPluginResponse`

NewCleanUpExpiredPingfederatePersistentAccessGrantsPluginResponse instantiates a new CleanUpExpiredPingfederatePersistentAccessGrantsPluginResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCleanUpExpiredPingfederatePersistentAccessGrantsPluginResponseWithDefaults

`func NewCleanUpExpiredPingfederatePersistentAccessGrantsPluginResponseWithDefaults() *CleanUpExpiredPingfederatePersistentAccessGrantsPluginResponse`

NewCleanUpExpiredPingfederatePersistentAccessGrantsPluginResponseWithDefaults instantiates a new CleanUpExpiredPingfederatePersistentAccessGrantsPluginResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *CleanUpExpiredPingfederatePersistentAccessGrantsPluginResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *CleanUpExpiredPingfederatePersistentAccessGrantsPluginResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *CleanUpExpiredPingfederatePersistentAccessGrantsPluginResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *CleanUpExpiredPingfederatePersistentAccessGrantsPluginResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *CleanUpExpiredPingfederatePersistentAccessGrantsPluginResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *CleanUpExpiredPingfederatePersistentAccessGrantsPluginResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *CleanUpExpiredPingfederatePersistentAccessGrantsPluginResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *CleanUpExpiredPingfederatePersistentAccessGrantsPluginResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetId

`func (o *CleanUpExpiredPingfederatePersistentAccessGrantsPluginResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CleanUpExpiredPingfederatePersistentAccessGrantsPluginResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CleanUpExpiredPingfederatePersistentAccessGrantsPluginResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *CleanUpExpiredPingfederatePersistentAccessGrantsPluginResponse) GetSchemas() []EnumcleanUpExpiredPingfederatePersistentAccessGrantsPluginSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *CleanUpExpiredPingfederatePersistentAccessGrantsPluginResponse) GetSchemasOk() (*[]EnumcleanUpExpiredPingfederatePersistentAccessGrantsPluginSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *CleanUpExpiredPingfederatePersistentAccessGrantsPluginResponse) SetSchemas(v []EnumcleanUpExpiredPingfederatePersistentAccessGrantsPluginSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetPollingInterval

`func (o *CleanUpExpiredPingfederatePersistentAccessGrantsPluginResponse) GetPollingInterval() string`

GetPollingInterval returns the PollingInterval field if non-nil, zero value otherwise.

### GetPollingIntervalOk

`func (o *CleanUpExpiredPingfederatePersistentAccessGrantsPluginResponse) GetPollingIntervalOk() (*string, bool)`

GetPollingIntervalOk returns a tuple with the PollingInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPollingInterval

`func (o *CleanUpExpiredPingfederatePersistentAccessGrantsPluginResponse) SetPollingInterval(v string)`

SetPollingInterval sets PollingInterval field to given value.


### GetPeerServerPriorityIndex

`func (o *CleanUpExpiredPingfederatePersistentAccessGrantsPluginResponse) GetPeerServerPriorityIndex() int32`

GetPeerServerPriorityIndex returns the PeerServerPriorityIndex field if non-nil, zero value otherwise.

### GetPeerServerPriorityIndexOk

`func (o *CleanUpExpiredPingfederatePersistentAccessGrantsPluginResponse) GetPeerServerPriorityIndexOk() (*int32, bool)`

GetPeerServerPriorityIndexOk returns a tuple with the PeerServerPriorityIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerServerPriorityIndex

`func (o *CleanUpExpiredPingfederatePersistentAccessGrantsPluginResponse) SetPeerServerPriorityIndex(v int32)`

SetPeerServerPriorityIndex sets PeerServerPriorityIndex field to given value.

### HasPeerServerPriorityIndex

`func (o *CleanUpExpiredPingfederatePersistentAccessGrantsPluginResponse) HasPeerServerPriorityIndex() bool`

HasPeerServerPriorityIndex returns a boolean if a field has been set.

### GetBaseDN

`func (o *CleanUpExpiredPingfederatePersistentAccessGrantsPluginResponse) GetBaseDN() string`

GetBaseDN returns the BaseDN field if non-nil, zero value otherwise.

### GetBaseDNOk

`func (o *CleanUpExpiredPingfederatePersistentAccessGrantsPluginResponse) GetBaseDNOk() (*string, bool)`

GetBaseDNOk returns a tuple with the BaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseDN

`func (o *CleanUpExpiredPingfederatePersistentAccessGrantsPluginResponse) SetBaseDN(v string)`

SetBaseDN sets BaseDN field to given value.

### HasBaseDN

`func (o *CleanUpExpiredPingfederatePersistentAccessGrantsPluginResponse) HasBaseDN() bool`

HasBaseDN returns a boolean if a field has been set.

### GetMaxUpdatesPerSecond

`func (o *CleanUpExpiredPingfederatePersistentAccessGrantsPluginResponse) GetMaxUpdatesPerSecond() int32`

GetMaxUpdatesPerSecond returns the MaxUpdatesPerSecond field if non-nil, zero value otherwise.

### GetMaxUpdatesPerSecondOk

`func (o *CleanUpExpiredPingfederatePersistentAccessGrantsPluginResponse) GetMaxUpdatesPerSecondOk() (*int32, bool)`

GetMaxUpdatesPerSecondOk returns a tuple with the MaxUpdatesPerSecond field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxUpdatesPerSecond

`func (o *CleanUpExpiredPingfederatePersistentAccessGrantsPluginResponse) SetMaxUpdatesPerSecond(v int32)`

SetMaxUpdatesPerSecond sets MaxUpdatesPerSecond field to given value.


### GetNumDeleteThreads

`func (o *CleanUpExpiredPingfederatePersistentAccessGrantsPluginResponse) GetNumDeleteThreads() int32`

GetNumDeleteThreads returns the NumDeleteThreads field if non-nil, zero value otherwise.

### GetNumDeleteThreadsOk

`func (o *CleanUpExpiredPingfederatePersistentAccessGrantsPluginResponse) GetNumDeleteThreadsOk() (*int32, bool)`

GetNumDeleteThreadsOk returns a tuple with the NumDeleteThreads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumDeleteThreads

`func (o *CleanUpExpiredPingfederatePersistentAccessGrantsPluginResponse) SetNumDeleteThreads(v int32)`

SetNumDeleteThreads sets NumDeleteThreads field to given value.


### GetEnabled

`func (o *CleanUpExpiredPingfederatePersistentAccessGrantsPluginResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CleanUpExpiredPingfederatePersistentAccessGrantsPluginResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CleanUpExpiredPingfederatePersistentAccessGrantsPluginResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


