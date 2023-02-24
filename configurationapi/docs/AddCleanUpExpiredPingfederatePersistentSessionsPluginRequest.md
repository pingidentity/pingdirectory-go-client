# AddCleanUpExpiredPingfederatePersistentSessionsPluginRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PluginName** | **string** | Name of the new Plugin | 
**Schemas** | [**[]EnumcleanUpExpiredPingfederatePersistentSessionsPluginSchemaUrn**](EnumcleanUpExpiredPingfederatePersistentSessionsPluginSchemaUrn.md) |  | 
**PollingInterval** | Pointer to **string** | This specifies how often the plugin should check for expired data. It also controls the offset of peer servers (see the peer-server-priority-index for more information). | [optional] 
**PeerServerPriorityIndex** | Pointer to **int32** | In a replicated environment, this determines the order in which peer servers should attempt to purge data. | [optional] 
**BaseDN** | Pointer to **string** | Only entries located within the subtree specified by this base DN are eligible for purging. | [optional] 
**MaxUpdatesPerSecond** | Pointer to **int32** | This setting smooths out the performance impact on the server by throttling the purging to the specified maximum number of updates per second. To avoid a large backlog, this value should be set comfortably above the average rate that expired data is generated. When purge-behavior is set to subtree-delete-entries, then deletion of the entire subtree is considered a single update for the purposes of throttling. | [optional] 
**NumDeleteThreads** | Pointer to **int32** | The number of threads used to delete expired entries. | [optional] 
**Enabled** | **bool** | Indicates whether the plug-in is enabled for use. | 

## Methods

### NewAddCleanUpExpiredPingfederatePersistentSessionsPluginRequest

`func NewAddCleanUpExpiredPingfederatePersistentSessionsPluginRequest(pluginName string, schemas []EnumcleanUpExpiredPingfederatePersistentSessionsPluginSchemaUrn, enabled bool, ) *AddCleanUpExpiredPingfederatePersistentSessionsPluginRequest`

NewAddCleanUpExpiredPingfederatePersistentSessionsPluginRequest instantiates a new AddCleanUpExpiredPingfederatePersistentSessionsPluginRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddCleanUpExpiredPingfederatePersistentSessionsPluginRequestWithDefaults

`func NewAddCleanUpExpiredPingfederatePersistentSessionsPluginRequestWithDefaults() *AddCleanUpExpiredPingfederatePersistentSessionsPluginRequest`

NewAddCleanUpExpiredPingfederatePersistentSessionsPluginRequestWithDefaults instantiates a new AddCleanUpExpiredPingfederatePersistentSessionsPluginRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPluginName

`func (o *AddCleanUpExpiredPingfederatePersistentSessionsPluginRequest) GetPluginName() string`

GetPluginName returns the PluginName field if non-nil, zero value otherwise.

### GetPluginNameOk

`func (o *AddCleanUpExpiredPingfederatePersistentSessionsPluginRequest) GetPluginNameOk() (*string, bool)`

GetPluginNameOk returns a tuple with the PluginName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginName

`func (o *AddCleanUpExpiredPingfederatePersistentSessionsPluginRequest) SetPluginName(v string)`

SetPluginName sets PluginName field to given value.


### GetSchemas

`func (o *AddCleanUpExpiredPingfederatePersistentSessionsPluginRequest) GetSchemas() []EnumcleanUpExpiredPingfederatePersistentSessionsPluginSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddCleanUpExpiredPingfederatePersistentSessionsPluginRequest) GetSchemasOk() (*[]EnumcleanUpExpiredPingfederatePersistentSessionsPluginSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddCleanUpExpiredPingfederatePersistentSessionsPluginRequest) SetSchemas(v []EnumcleanUpExpiredPingfederatePersistentSessionsPluginSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetPollingInterval

`func (o *AddCleanUpExpiredPingfederatePersistentSessionsPluginRequest) GetPollingInterval() string`

GetPollingInterval returns the PollingInterval field if non-nil, zero value otherwise.

### GetPollingIntervalOk

`func (o *AddCleanUpExpiredPingfederatePersistentSessionsPluginRequest) GetPollingIntervalOk() (*string, bool)`

GetPollingIntervalOk returns a tuple with the PollingInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPollingInterval

`func (o *AddCleanUpExpiredPingfederatePersistentSessionsPluginRequest) SetPollingInterval(v string)`

SetPollingInterval sets PollingInterval field to given value.

### HasPollingInterval

`func (o *AddCleanUpExpiredPingfederatePersistentSessionsPluginRequest) HasPollingInterval() bool`

HasPollingInterval returns a boolean if a field has been set.

### GetPeerServerPriorityIndex

`func (o *AddCleanUpExpiredPingfederatePersistentSessionsPluginRequest) GetPeerServerPriorityIndex() int32`

GetPeerServerPriorityIndex returns the PeerServerPriorityIndex field if non-nil, zero value otherwise.

### GetPeerServerPriorityIndexOk

`func (o *AddCleanUpExpiredPingfederatePersistentSessionsPluginRequest) GetPeerServerPriorityIndexOk() (*int32, bool)`

GetPeerServerPriorityIndexOk returns a tuple with the PeerServerPriorityIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerServerPriorityIndex

`func (o *AddCleanUpExpiredPingfederatePersistentSessionsPluginRequest) SetPeerServerPriorityIndex(v int32)`

SetPeerServerPriorityIndex sets PeerServerPriorityIndex field to given value.

### HasPeerServerPriorityIndex

`func (o *AddCleanUpExpiredPingfederatePersistentSessionsPluginRequest) HasPeerServerPriorityIndex() bool`

HasPeerServerPriorityIndex returns a boolean if a field has been set.

### GetBaseDN

`func (o *AddCleanUpExpiredPingfederatePersistentSessionsPluginRequest) GetBaseDN() string`

GetBaseDN returns the BaseDN field if non-nil, zero value otherwise.

### GetBaseDNOk

`func (o *AddCleanUpExpiredPingfederatePersistentSessionsPluginRequest) GetBaseDNOk() (*string, bool)`

GetBaseDNOk returns a tuple with the BaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseDN

`func (o *AddCleanUpExpiredPingfederatePersistentSessionsPluginRequest) SetBaseDN(v string)`

SetBaseDN sets BaseDN field to given value.

### HasBaseDN

`func (o *AddCleanUpExpiredPingfederatePersistentSessionsPluginRequest) HasBaseDN() bool`

HasBaseDN returns a boolean if a field has been set.

### GetMaxUpdatesPerSecond

`func (o *AddCleanUpExpiredPingfederatePersistentSessionsPluginRequest) GetMaxUpdatesPerSecond() int32`

GetMaxUpdatesPerSecond returns the MaxUpdatesPerSecond field if non-nil, zero value otherwise.

### GetMaxUpdatesPerSecondOk

`func (o *AddCleanUpExpiredPingfederatePersistentSessionsPluginRequest) GetMaxUpdatesPerSecondOk() (*int32, bool)`

GetMaxUpdatesPerSecondOk returns a tuple with the MaxUpdatesPerSecond field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxUpdatesPerSecond

`func (o *AddCleanUpExpiredPingfederatePersistentSessionsPluginRequest) SetMaxUpdatesPerSecond(v int32)`

SetMaxUpdatesPerSecond sets MaxUpdatesPerSecond field to given value.

### HasMaxUpdatesPerSecond

`func (o *AddCleanUpExpiredPingfederatePersistentSessionsPluginRequest) HasMaxUpdatesPerSecond() bool`

HasMaxUpdatesPerSecond returns a boolean if a field has been set.

### GetNumDeleteThreads

`func (o *AddCleanUpExpiredPingfederatePersistentSessionsPluginRequest) GetNumDeleteThreads() int32`

GetNumDeleteThreads returns the NumDeleteThreads field if non-nil, zero value otherwise.

### GetNumDeleteThreadsOk

`func (o *AddCleanUpExpiredPingfederatePersistentSessionsPluginRequest) GetNumDeleteThreadsOk() (*int32, bool)`

GetNumDeleteThreadsOk returns a tuple with the NumDeleteThreads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumDeleteThreads

`func (o *AddCleanUpExpiredPingfederatePersistentSessionsPluginRequest) SetNumDeleteThreads(v int32)`

SetNumDeleteThreads sets NumDeleteThreads field to given value.

### HasNumDeleteThreads

`func (o *AddCleanUpExpiredPingfederatePersistentSessionsPluginRequest) HasNumDeleteThreads() bool`

HasNumDeleteThreads returns a boolean if a field has been set.

### GetEnabled

`func (o *AddCleanUpExpiredPingfederatePersistentSessionsPluginRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddCleanUpExpiredPingfederatePersistentSessionsPluginRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddCleanUpExpiredPingfederatePersistentSessionsPluginRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


